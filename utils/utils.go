/**
 * (C) Copyright IBM Corp. 2023.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"encoding/json"
	"errors"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/configuration/core_config"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM/go-sdk-core/v5/core"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"github.com/ghodss/yaml"
	JmesPath "github.com/jmespath/go-jmespath"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"strings"
)

const EmptyString = ""

type CustomPrinterFunc func(interface{}) bool

type Utils struct {
	ui                       terminal.UI
	OutputFormat             string
	JMESQuery                string
	serviceSpecificErrorMsgs map[string]string
	// provide plugin points for users to customize output printing
	customOutputPrinter        CustomPrinterFunc
	customJsonPrinter          CustomPrinterFunc
	customErrorResponseHandler func(interface{})
}

var context plugin.PluginContext
var currentCommand string

var TableHeaderOrder []string

func NewUtils(ui terminal.UI) *Utils {
	return &Utils{ui: ui}
}

func (u *Utils) SetTableHeaderOrder(order []string) {
	TableHeaderOrder = order
}

func (u *Utils) GetServiceURL(GetServiceURLForRegion func(string) (string, error)) string {
	// if GetServiceURLForRegion returns an error, it doesn't matter. just use the empty string
	// it returns and continue - the SDK will use the default URL
	region := u.GetRegionFromContext()
	serviceUrl, _ := GetServiceURLForRegion(region)
	if serviceUrl == "" {
		// check for private endpoint
		serviceUrl, _ = GetServiceURLForRegion("private." + region)
	}

	return serviceUrl
}

func checkForBadURL(err error, additionalMessage string) error {
	message := err.Error()
	isBadURLError := strings.Contains(message, "no such host")
	if isBadURLError {
		err = errors.New(message + "\n\n" + additionalMessage)
	}

	return err
}

func (u *Utils) HandleErrorResponse(errorBody interface{}) {
	if u.customErrorResponseHandler != nil {
		u.customErrorResponseHandler(errorBody)
		return
	}

	u.HandleErrorResponseImpl(errorBody)
}

func (u *Utils) HandleErrorResponseImpl(errorBody interface{}) {
	// error messages should not be returned in machine readable format
	u.SetOutputFormat("table")
	u.SetJMESQuery(EmptyString)
	u.PrintOutput(errorBody, terminal.ErrOutput)
}

func (u *Utils) CheckResponseForError(response *core.DetailedResponse, err error) (bool, interface{}) {
	// get json body of the response to print - may be success or error body
	successBody, errorBody, errorGettingResult := u.GetResult(response, err)

	if errorGettingResult != nil {
		errorGettingResult = checkForBadURL(errorGettingResult, u.serviceSpecificErrorMsgs["badURL"])
	}

	u.HandleError(errorGettingResult, translation.T("response-processing-error"))

	if errorBody != nil {
		u.HandleErrorResponse(errorBody)
		return false, nil
	}

	return true, successBody
}

func (u *Utils) ProcessResponse(response *core.DetailedResponse, err error) {
	if ok, result := u.CheckResponseForError(response, err); ok {
		u.PrintOutput(result, terminal.Output)
	} else {
		u.NonZeroExit()
	}
}

func (u *Utils) ProcessBinaryResponse(response *core.DetailedResponse, err error, outputFilename string) {
	if ok, result := u.CheckResponseForError(response, err); ok {
		// write the binary data to the file
		err := u.WriteFile(result, outputFilename)
		u.HandleError(err, translation.T("file-response-error", map[string]interface{}{
			"FILENAME": outputFilename,
		}))

		u.Ok()
		if u.OutputIsNotMachineReadable() {
			// this is silenced in quiet mode
			u.ui.Verbose(translation.T("output-file-confirmation", map[string]interface{}{
				"FILENAME": outputFilename,
			}))
		} else {
			u.PrintOutput(EmptyString, terminal.Output)
		}
	} else {
		u.NonZeroExit()
	}
}

func (u *Utils) ProcessEmptyResponse(response *core.DetailedResponse, err error) {
	if ok, _ := u.CheckResponseForError(response, err); ok {
		u.Ok()
		if !u.OutputIsNotMachineReadable() {
			u.PrintOutput(EmptyString, terminal.Output)
		}
	} else {
		u.NonZeroExit()
	}
}

func (u *Utils) PrintOutput(result interface{}, tableWriter io.Writer) {
	if u.customOutputPrinter != nil {
		if ok := u.customOutputPrinter(result); ok {
			return
		}
	}

	u.PrintOutputImpl(result, tableWriter)
}

func (u *Utils) PrintOutputImpl(result interface{}, tableWriter io.Writer) {
	// this eliminates any knowledge of structs, leaving the result to match the json
	// structure and key names the user expects
	result = u.MakeResultGeneric(result)

	// the jmes query applies to everything, so do it first
	if u.JMESQuery != EmptyString {
		jmes, jErr := JmesPath.Compile(u.JMESQuery)
		u.HandleError(jErr, translation.T("jmespath-compile-error"))

		newJson, searchErr := jmes.Search(result)
		u.HandleError(searchErr, translation.T("jmespath-application-error"))

		result = newJson
	}

	// print something based on the output format
	switch strings.ToLower(u.OutputFormat) {
	case "yaml":
		yamlified, yErr := yaml.Marshal(result)
		u.HandleError(yErr, translation.T("yaml-conversion-error"))

		u.ui.Print(string(yamlified))

	case "json":
		u.PrintJSON(result)

	default:
		// default to "table" - this will dynamically generate a table
		tableData := FormatTableData(result, u.JMESQuery)
		table := CreateTable(tableData, tableWriter)
		if table == nil {
			u.ui.Print(translation.T("no-data-for-table"))
		} else {
			table.Print()
		}
	}
}

func (u *Utils) PrintJSON(result interface{}) {
	if u.customJsonPrinter != nil {
		if ok := u.customJsonPrinter(result); ok {
			return
		}
	}

	u.PrintJSONImpl(result)
}

func (u *Utils) PrintJSONImpl(result interface{}) {
	// this will print raw json
	b, _ := json.MarshalIndent(result, EmptyString, "  ")
	u.ui.Print(string(b))
}

func (u *Utils) HandleError(err error, message string) {
	if err != nil {
		// if we get here, the output will not be machine readble
		// so its okay to print "Failed"
		u.ui.Failed(message + ":\n" + err.Error())
		u.NonZeroExit()
	}
}

func (u *Utils) ConfirmRunningCommand() {
	if u.OutputIsNotMachineReadable() {
		// this will be silenced in "quiet" mode
		// otherwise, it prints independently of log level
		u.ui.Verbose("...")
	}
}

func (u *Utils) ReadAsFile(userInput string) bool {
	return strings.HasPrefix(userInput, "@")
}

func (u *Utils) Ok() {
	if u.OutputIsNotMachineReadable() {
		u.ui.Ok()
	}
}

func (u *Utils) Say(message string) {
	// the "Say" method in the ui package is deprecated
	// the behavior is replaced with "Verbose"
	u.ui.Verbose(message)
}

func (u *Utils) Prompt(message string, options *terminal.PromptOptions) *terminal.Prompt {
	return u.ui.Prompt(message, options)
}

func (u *Utils) WriteFile(fileInterface interface{}, filename string) error {
	// open a new file
	outFile, outFileErr := os.Create(filename)
	if outFileErr != nil {
		return outFileErr
	}
	defer outFile.Close()

	file, ok := fileInterface.(io.ReadCloser)
	if !ok {
		return errors.New(translation.T("file-conversion-error"))
	}

	_, FileWriteErr := io.Copy(outFile, file)
	return FileWriteErr
}

func (u *Utils) GetJsonStringAsBytes(json string) (stringAsBytes []byte) {
	if u.ReadAsFile(json) {
		// read the json from a file
		// the [1:] removes the @ symbol from the string, used to designate a file
		fileContents, fileErr := os.ReadFile(json[1:])
		u.HandleError(fileErr, translation.T("file-reading-error", map[string]interface{}{
			"FILENAME": json[1:],
		}))
		stringAsBytes = fileContents
	} else {
		stringAsBytes = []byte(json)
	}

	return
}

func (u *Utils) MakeResultGeneric(result interface{}) interface{} {
	bytes, err := json.Marshal(result)
	u.HandleError(err, translation.T("json-conversion-error"))

	var data interface{}
	err = json.Unmarshal(bytes, &data)
	u.HandleError(err, translation.T("json-conversion-error"))

	return data
}

func (u *Utils) GetResult(response *core.DetailedResponse, err error) (interface{}, interface{}, error) {
	// based on the current code in the go sdk core, this situation is impossible.
	// adding this check as a failsafe should external behavior ever change in the future
	if response == nil && err == nil {
		return nil, nil, errors.New(translation.T("no-service-response-error"))
	}

	// this means there was an error response
	if err != nil {
		if response != nil {
			if response.GetResult() != nil {
				// we want to let the user know that the request failed,
				// even though we are going to print the body
				u.ui.Failed(EmptyString)
				return nil, response.GetResult(), nil
			} else if response.GetRawResult() != nil {
				// convert the raw result to a string and return as an error
				return nil, nil, errors.New(string(response.GetRawResult()))
			}
		}

		// if the response is nil, or if both the result and raw result are nil,
		// return the error
		return nil, nil, err
	}

	// this means there was a success response
	return response.GetResult(), nil, nil
}

// Exit the program with a non-zero exit code.
// This should only be called in an error situation.
func (u *Utils) NonZeroExit() {
	os.Exit(1)
}

func (u *Utils) OutputIsNotMachineReadable() bool {
	return strings.ToLower(u.OutputFormat) != "json" && strings.ToLower(u.OutputFormat) != "yaml"
}

// The following methods allow access to the properties from
// instances of the Utilities *interface* that this
// struct satisfies

func (u *Utils) ExposeOutputFormatVar() *string {
	return &u.OutputFormat
}

func (u *Utils) ExposeJMESQueryVar() *string {
	return &u.JMESQuery
}

func (u *Utils) SetOutputFormat(value string) {
	u.OutputFormat = value
}

func (u *Utils) SetJMESQuery(value string) {
	u.JMESQuery = value
}

func (u *Utils) GetOutputFormat() string {
	return u.OutputFormat
}

func (u *Utils) GetJMESQuery() string {
	return u.JMESQuery
}

func (u *Utils) SetServiceErrorMessages(msgs map[string]string) {
	u.serviceSpecificErrorMsgs = msgs
}

// store the name of the currently executed command
// to enable context-aware logic in the utilities
func (u *Utils) SetCommandName(args []string) {
	// the first arg should always be the command name
	// but make sure the array isnt empty
	if len(args) > 0 {
		currentCommand = args[0]
	}
}

func (u *Utils) GetCommandName() string {
	return currentCommand
}

func (u *Utils) SetContext(c plugin.PluginContext) {
	context = c
}

func (u *Utils) GetPluginConfig() plugin.PluginConfig {
    return context.PluginConfig()
}

func (u *Utils) GetAuthenticator(serviceName string) (core.Authenticator, error) {
	authenticator, err := core.GetAuthenticatorFromEnvironment(serviceName)
	if authenticator != nil && err == nil {
		return authenticator, err
	}

	if token := getActiveIAMToken(); token != "" {
		authenticator, err = core.NewBearerTokenAuthenticator(token)
		return authenticator, err
	}

	if err == nil {
		// if there are no credentials in the environment at all,
		// there wont be an error, just a nil authenticator,
		// but we dont want the code to progress - we want to trigger
		// the error message that addresses credentials
		err = errors.New(translation.T("no-credentials"))
	}

	return nil, err
}

func (u *Utils) GetRegionFromContext() string {
	return context.CurrentRegion().Name // e.g. us-south
}

func (u *Utils) PostProcessServiceConfiguration(service *core.BaseService, serviceName string) error {
	externalVars, err := core.GetServiceProperties(serviceName)
	if err != nil {
		return err
	}

	// if the url is set in the environment, it would have been overwritten by the
	// programatically-set URL from the plugin context. however, the external
	// variables should take priority so we need to check for it again here
	if url, ok := externalVars[core.PROPNAME_SVC_URL]; ok && url != "" {
		err = service.SetServiceURL(url)
		if err != nil {
			return err
		}
	}

	// if disableSSL is set in the environment, everything will be how it should be already
	// if it is NOT set in the environment, we should check to see if SSL verification is
	// disabled in the plugin context and handle that information as needed
	if _, set := externalVars[core.PROPNAME_SVC_DISABLE_SSL]; !set {
		// check disable ssl verification - this is currently the only service level parameter
		// supported in the plugin context
		if context.IsSSLDisabled() {
			service.DisableSSLVerification()
			// propagate this to request-based authenticators
			authType := service.Options.Authenticator.AuthenticationType()
			if authType == core.AUTHTYPE_IAM {
				authenticator := service.Options.Authenticator.(*core.IamAuthenticator)
				authenticator.DisableSSLVerification = true
			} else if authType == core.AUTHTYPE_CP4D {
				authenticator := service.Options.Authenticator.(*core.CloudPakForDataAuthenticator)
				authenticator.DisableSSLVerification = true
			}
		}
	}

	return nil
}

func getActiveIAMToken() string {
	// no point in looking for a token if the user isn't logged in
	if !context.IsLoggedIn() {
		return ""
	}

	// read current token
	token := sanitizeToken(context.IAMToken())
	// the token should never be empty while logged in,
	// but check for that just in case
	if token == "" {
		return token
	}

	// check if token is still active
	tokenInfo := core_config.NewIAMTokenInfo(token)
	expireTime := tokenInfo.Expiry.Unix()
	thirtySeconds := int64(30)
	// if token is nearing expiration, refresh it
	// allow a 30 second buffer to ensure the token does
	// not expire while the rest of the code is executing
	if core.GetCurrentTime() > (expireTime - thirtySeconds) {
		newToken, err := context.RefreshIAMToken()
		if err != nil {
			return ""
		}

		token = sanitizeToken(newToken)
	}

	return token
}

func sanitizeToken(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func (u *Utils) InitializeLogger(quiet bool) {
	if quiet {
		u.ui.SetQuiet(true)
		// the CLI SDK still logs errors in quiet mode, so match that behavior for now
		core.SetLoggingLevel(core.LevelError)
	}

	// in the ibm-cloud-cli-sdk package, the quiet flag has no bearing
	// on trace logs. this is handled the same way here, for consistency
	trace := context.Trace()
	var logDestination io.Writer

	switch strings.ToLower(trace) {
	case "", "false":
		// do nothing
	case "true":
		logDestination = terminal.ErrOutput
	default:
		// assume it's a file and try to use it
		file, err := os.OpenFile(trace, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
		if err != nil {
			u.ui.Warn(translation.T("log-file-creation-error", map[string]interface{}{
				"PATH":  trace,
				"ERROR": err.Error(),
			}))

			// if the file cannot be opened, still log the trace output to stderr.
			// this matches the behavior in the ibm-cloud-cli-sdk package
			logDestination = terminal.ErrOutput
		} else {
			logDestination = file
		}
	}

	// the trace logger exposed by the ibm-cloud-cli-sdk package creates a problem -
	// it requires that the "Client" field in the base service be overridden. this would
	// prevent any retries or disabling of ssl verification. to maintain those features,
	// the logger in the go core will be used instead of the trace logger. it outputs nearly
	// identical information
	if logDestination != nil {
		goLogger := log.New(logDestination, "", log.LstdFlags)
		core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
	}
}

func (u *Utils) ConfirmDelete(force bool) bool {
	if force {
		return true
	}

	var confirmed bool

	// require a value from the user
	options := &terminal.PromptOptions{
		Required: true,
	}

	err := u.ui.Prompt(translation.T("confirm-delete"), options).Resolve(&confirmed)
	u.HandleError(err, translation.T("confirmation-error"))

	return confirmed
}

func (u *Utils) ValidateRequiredFlags(required []string, flags *pflag.FlagSet, serviceName string) error {
	config := u.GetPluginConfig()
	missingFlags := make([]string, 0)
	for _, flagName := range required {
		if !flags.Changed(flagName) && !config.Exists(serviceName + "-" + flagName) {
			missingFlags = append(missingFlags, `"` + flagName + `"`)
		}
	}

	if len(missingFlags) > 0 {
		return errors.New(translation.T("missing-required-flags-error", map[string]interface{}{
			"FLAGS": strings.Join(missingFlags, ", "),
		}))
	}

	return nil
}

// Creates a new error with a descriptive message placed before the original
// error message. This is only used for interactive mode.
func (u *Utils) CreateErrorWithMessage(err error, msg string) error {
	if err != nil {
		err = errors.New(msg + ":\n" + err.Error())
	}

	return err
}
