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

package testing_utilities

import (
	"errors"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	fakeTerm "github.com/IBM-Cloud/ibm-cloud-cli-sdk/testhelpers/terminal"
	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
	"io"
	"strings"
)

var ConfirmRunningCommandCalled bool

const (
	MockSuccessResponse = "SuccessfulResponse"
	MockErrorResponse = "ErrorResponse"
	MockErrorMessage = "ErrorMessage"
	MockFileContents = "This is a mock file"
)

// Returns a response with a mock string as the response body
func GetMockSuccessResponse() (interface{}, *core.DetailedResponse, error) {
	return GetMockSuccessResponseWithBody(MockSuccessResponse)
}

// Returns a response with a given success body
func GetMockSuccessResponseWithBody(body interface{}) (interface{}, *core.DetailedResponse, error) {
	res := new(core.DetailedResponse)
	res.Result = body
	return nil, res, nil
}

// Returns an error body with an attached error object to mimic the behavior
// of the Go SDK Core
func GetMockErrorResponse() (interface{}, *core.DetailedResponse, error) {
	res := new(core.DetailedResponse)
	res.Result = MockErrorResponse
	return nil, res, errors.New(MockErrorMessage)
}

// Used to test commands that expect a file response from the service
func GetMockFileResponse() (interface{}, *core.DetailedResponse, error) {
	res := new(core.DetailedResponse)
	res.Result = io.NopCloser(strings.NewReader(MockFileContents))
	return nil, res, nil
}

// Provide a stub for the InitializeServiceInstance function so that unit tests for
// individual commands do not need to provide mock values for all of the global
// parameters in addtion to their own parameters.
type TestServiceCommandHelper struct {}
func (t *TestServiceCommandHelper) InitializeServiceInstance(f *pflag.FlagSet) {}

// Positive Test Utilities
type PositiveTestUtilities struct {
	ExpectedJmespath string
	OutputFormat string
	JMESQuery string
	FakeUI *fakeTerm.FakeUI
	LastThingSaid string
}

func NewPositiveTestUtilities() *PositiveTestUtilities {
	return &PositiveTestUtilities{ExpectedJmespath: "", FakeUI: fakeTerm.NewFakeUI()}
}

func (u *PositiveTestUtilities) HandleError(err error, message string) {
	Expect(err).To(BeNil())
}

func (u *PositiveTestUtilities) ConfirmRunningCommand() {
	ConfirmRunningCommandCalled = true
}

func (u *PositiveTestUtilities) GetServiceURL(getURL func(string) (string, error)) (string) { return "" }

func (u *PositiveTestUtilities) GetJsonStringAsBytes(json string) []byte {
	return []byte(json)
}

func (u *PositiveTestUtilities) ProcessResponse(res *core.DetailedResponse, err error) {
	Expect(err).To(BeNil())
	result := res.GetResult()
	result, ok := result.(string)
	if ok {
		// if result is a string, it should be the mock success response
		Expect(result).To(Equal(MockSuccessResponse))
	} else {
		// if not, it is a scenario where the result needed to be a complex object
		Expect(result).NotTo(BeNil())
	}
}

func (u *PositiveTestUtilities) ProcessEmptyResponse(res *core.DetailedResponse, err error) {
	Expect(err).To(BeNil())
}

func (u *PositiveTestUtilities) ProcessBinaryResponse(res *core.DetailedResponse, err error, filename string) {
	Expect(err).To(BeNil())
	result := res.GetResult()
	_, ok := result.(io.ReadCloser)
	Expect(ok).To(Equal(true))
	Expect(filename).To(Equal("tempdir/test-output.txt"))
}

func (u *PositiveTestUtilities) ExposeOutputFormatVar() *string {
	return &u.OutputFormat
}

func (u *PositiveTestUtilities) ExposeJMESQueryVar() *string {
	return &u.JMESQuery
}

func (u *PositiveTestUtilities) SetJMESQuery(value string) {
	Expect(value).To(Equal(u.ExpectedJmespath))
	u.JMESQuery = value
}

func (u *PositiveTestUtilities) GetJMESQuery() string {
	return u.JMESQuery
}

func (u *PositiveTestUtilities) SetTableHeaderOrder([]string) {}

func (u *PositiveTestUtilities) CheckResponseForError(response *core.DetailedResponse, err error) (bool, interface{}) {
	return true, response.GetResult()
}

func (u *PositiveTestUtilities) NonZeroExit() {}

func (u *PositiveTestUtilities) Say(s string) {
	u.LastThingSaid = s
}

func (u *PositiveTestUtilities) Ok() {
	u.LastThingSaid = "OK"
}

func (u *PositiveTestUtilities) WriteFile(interface{}, string) error {
	return nil
}

func (u *PositiveTestUtilities) PrintOutput(interface{}, io.Writer) {}

func (u *PositiveTestUtilities) OutputIsNotMachineReadable() bool {
	return false
}

func (u *PositiveTestUtilities) Prompt(message string, opts *terminal.PromptOptions) *terminal.Prompt {
	return u.FakeUI.Prompt(message, opts)
}

func (u *PositiveTestUtilities) ConfirmDelete(force bool) bool {
	return true
}

func (u *PositiveTestUtilities) GetAuthenticator(string) (core.Authenticator, error) {
	return &core.NoAuthAuthenticator{}, nil
}

func (u *PositiveTestUtilities) GetRegionFromContext() string {
	return "us-south"
}

func (u *PositiveTestUtilities) PostProcessServiceConfiguration(*core.BaseService, string) error {
	return nil
}

func (u *PositiveTestUtilities) InitializeLogger(bool) {}

func (u *PositiveTestUtilities) ValidateRequiredFlags(required []string, flags *pflag.FlagSet, name string) error {
	// required flags will always be provided in the tests
	return nil
}

func (u *PositiveTestUtilities) CreateErrorWithMessage(err error, msg string) error {
	Expect(err).To(BeNil())
	Expect(msg).To(Equal(""))
	return nil
}

func (u *PositiveTestUtilities) SetServiceErrorMessages(map[string]string) {}

func (u *PositiveTestUtilities) GetPluginConfig() plugin.PluginConfig {
	ctx := plugin.InitPluginContext("test")
	return ctx.PluginConfig()
}

// Negative Test Utilities
type NegativeTestUtilities struct {
	ExpectedJmespath string
	OutputFormat string
	JMESQuery string
}

func NewNegativeTestUtilities() *NegativeTestUtilities {
	return &NegativeTestUtilities{}
}

func (u *NegativeTestUtilities) HandleError(err error, message string) {
	if err != nil {
		// figure out a way to verify it was called with an error once
		Expect(err.Error()).To(Equal(MockErrorMessage))
	}
}

func (u *NegativeTestUtilities) ConfirmRunningCommand() {
	ConfirmRunningCommandCalled = true
}

func (u *NegativeTestUtilities) GetServiceURL(getURL func(string) (string, error)) (string) { return "" }

func (u *NegativeTestUtilities) GetJsonStringAsBytes(json string) []byte {
	return []byte(json)
}

func (u *NegativeTestUtilities) ProcessResponse(res *core.DetailedResponse, err error) {
	Expect(err.Error()).To(Equal(MockErrorMessage))
	Expect(res.GetResult()).To(Equal(MockErrorResponse))
}

func (u *NegativeTestUtilities) ProcessEmptyResponse(res *core.DetailedResponse, err error) {
	Expect(err.Error()).To(Equal(MockErrorMessage))
	Expect(res.GetResult()).To(Equal(MockErrorResponse))
}

func (u *NegativeTestUtilities) ProcessBinaryResponse(res *core.DetailedResponse, err error, filename string) {
	Expect(err.Error()).To(Equal(MockErrorMessage))
	Expect(res.GetResult()).To(Equal(MockErrorResponse))
}

func (u *NegativeTestUtilities) ExposeOutputFormatVar() *string {
	return &u.OutputFormat
}

func (u *NegativeTestUtilities) ExposeJMESQueryVar() *string {
	return &u.JMESQuery
}

func (u *NegativeTestUtilities) SetJMESQuery(value string) {
	Expect(value).To(Equal(u.ExpectedJmespath))
	u.JMESQuery = value
}

func (u *NegativeTestUtilities) GetJMESQuery() string {
	return u.JMESQuery
}

func (u *NegativeTestUtilities) SetTableHeaderOrder([]string) {}

func (u *NegativeTestUtilities) CheckResponseForError(response *core.DetailedResponse, err error) (bool, interface{}) {
	return false, response.GetResult()
}

func (u *NegativeTestUtilities) NonZeroExit() {}

func (u *NegativeTestUtilities) Say(string) {}

func (u *NegativeTestUtilities) Ok() {}

func (u *NegativeTestUtilities) WriteFile(interface{}, string) error {
	return nil
}

func (u *NegativeTestUtilities) PrintOutput(interface{}, io.Writer) {}

func (u *NegativeTestUtilities) OutputIsNotMachineReadable() bool {
	return false
}

func (u *NegativeTestUtilities) Prompt(message string, opts *terminal.PromptOptions) *terminal.Prompt {
	return fakeTerm.NewFakeUI().Prompt(message, opts)
}

func (u *NegativeTestUtilities) ConfirmDelete(force bool) bool {
	return true
}

func (u *NegativeTestUtilities) GetAuthenticator(string) (core.Authenticator, error) {
	return &core.NoAuthAuthenticator{}, nil
}

func (u *NegativeTestUtilities) GetRegionFromContext() string {
	return "us-south"
}

func (u *NegativeTestUtilities) PostProcessServiceConfiguration(*core.BaseService, string) error {
	return nil
}

func (u *NegativeTestUtilities) InitializeLogger(bool) {}

func (u *NegativeTestUtilities) ValidateRequiredFlags(required []string, flags *pflag.FlagSet, name string) error {
	// required flags will always be provided in the tests
	return nil
}

func (u *NegativeTestUtilities) CreateErrorWithMessage(err error, msg string) error {
	return nil
}

func (u *NegativeTestUtilities) SetServiceErrorMessages(map[string]string) {}

func (u *NegativeTestUtilities) GetPluginConfig() plugin.PluginConfig {
	ctx := plugin.InitPluginContext("test")
	return ctx.PluginConfig()
}
