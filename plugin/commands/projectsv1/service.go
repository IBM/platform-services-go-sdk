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

package projectsv1

import (
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/bluemix/terminal"
	"github.com/IBM-Cloud/ibm-cloud-cli-sdk/plugin"
	"github.com/IBM/go-sdk-core/v5/core"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"github.com/IBM/platform-services-go-sdk/projectsv1"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
)

type Utilities interface {
	HandleError(error, string)
	ConfirmRunningCommand()
	GetServiceURL(func(string) (string, error)) string
	GetJsonStringAsBytes(string) []byte
	ProcessResponse(*core.DetailedResponse, error)
	ProcessEmptyResponse(*core.DetailedResponse, error)
	ProcessBinaryResponse(*core.DetailedResponse, error, string)
	ExposeOutputFormatVar() *string
	ExposeJMESQueryVar() *string
	SetJMESQuery(string)
	GetJMESQuery() string
	SetTableHeaderOrder([]string)
	CheckResponseForError(*core.DetailedResponse, error) (bool, interface{})
	NonZeroExit()
	Say(string)
	Ok()
	Prompt(string, *terminal.PromptOptions) *terminal.Prompt
	ConfirmDelete(bool) bool
	WriteFile(interface{}, string) error
	PrintOutput(interface{}, io.Writer)
	OutputIsNotMachineReadable() bool
	GetAuthenticator(string) (core.Authenticator, error)
	GetRegionFromContext() string
	PostProcessServiceConfiguration(*core.BaseService, string) error
	InitializeLogger(bool)
	ValidateRequiredFlags([]string, *pflag.FlagSet, string) error
	CreateErrorWithMessage(error, string) error
	SetServiceErrorMessages(map[string]string)
	GetPluginConfig() plugin.PluginConfig
}

var ServiceInstance *projectsv1.ProjectsV1

type ProjectsV1CommandHelper struct {
	RequiredFlags []string
	utils Utilities
}

type ServiceCommandHelper interface {
	InitializeServiceInstance(*pflag.FlagSet)
}

var Service ServiceCommandHelper

var serviceErrors = map[string]string{
	"badURL": translation.T("project-bad-url-error-message"),
}

// add a function to return the super-command
func GetProjectsV1Command(utils Utilities) *cobra.Command {
	InitializeService(utils)

	serviceCommands := []*cobra.Command{
		GetListProjectsCommand(NewListProjectsCommandRunner(utils, ListProjectsRequestSender{}, ListProjectsAllPagesRequestSender{})),
		GetGetProjectCommand(NewGetProjectCommandRunner(utils, GetProjectRequestSender{})),
		GetDeleteProjectCommand(NewDeleteProjectCommandRunner(utils, DeleteProjectRequestSender{})),
		GetListConfigsCommand(NewListConfigsCommandRunner(utils, ListConfigsRequestSender{})),
		GetGetConfigCommand(NewGetConfigCommandRunner(utils, GetConfigRequestSender{})),
		GetDeleteConfigCommand(NewDeleteConfigCommandRunner(utils, DeleteConfigRequestSender{})),
		GetGetConfigDiffCommand(NewGetConfigDiffCommandRunner(utils, GetConfigDiffRequestSender{})),
		GetInstallConfigCommand(NewInstallConfigCommandRunner(utils, InstallConfigRequestSender{})),
		GetUninstallConfigCommand(NewUninstallConfigCommandRunner(utils, UninstallConfigRequestSender{})),
		GetGetSchematicsJobCommand(NewGetSchematicsJobCommandRunner(utils, GetSchematicsJobRequestSender{})),
		GetGetCostEstimateCommand(NewGetCostEstimateCommandRunner(utils, GetCostEstimateRequestSender{})),
		GetPostNotificationCommand(NewPostNotificationCommandRunner(utils, PostNotificationRequestSender{})),
		GetGetHealthCommand(NewGetHealthCommandRunner(utils, GetHealthRequestSender{})),
		GetPostEventNotificationsIntegrationCommand(NewPostEventNotificationsIntegrationCommandRunner(utils, PostEventNotificationsIntegrationRequestSender{})),
		GetPostEventNotificationCommand(NewPostEventNotificationCommandRunner(utils, PostEventNotificationRequestSender{})),
	}

	projectsCommand := &cobra.Command{
		Use: "project [command] [options]",
		Short: translation.T("project-short-description"),
		Long: translation.T("project-long-description"),
		DisableFlagsInUseLine: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// ignore the error passed here - it just checks for a faulty implementation of the quiet flag
			quiet, _ := cmd.Flags().GetBool("quiet")
			utils.InitializeLogger(quiet)

			// these must only be set once the service command is actually executed
			utils.SetServiceErrorMessages(serviceErrors)
		},
	}

	// these flags pertain to all commands
	projectsCommand.PersistentFlags().StringVarP(utils.ExposeOutputFormatVar(), "output", "", "table", translation.T("output-global-flag-description"))
	projectsCommand.PersistentFlags().StringVarP(utils.ExposeJMESQueryVar(), "jmes-query", "j", "", translation.T("jmes-query-global-flag-description"))
	projectsCommand.PersistentFlags().BoolP("quiet", "q", false, translation.T("quiet-global-flag-description"))

	projectsCommand.AddCommand(serviceCommands...)

	return projectsCommand
}

func InitializeService(utils Utilities) {
	Service = &ProjectsV1CommandHelper{utils: utils}
}
