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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.64.1-cee95189-20230124-211647
 */

package projectsv1

import (
	"github.com/IBM/go-sdk-core/v5/core"
	translation "github.com/IBM/platform-services-go-sdk/i18n"
	"github.com/IBM/platform-services-go-sdk/plugin/version"
	"github.com/IBM/platform-services-go-sdk/projectsv1"
	"github.com/IBM/platform-services-go-sdk/utils/deserialize"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"net/http"
)

var serviceName string = "projects"

func (r *ProjectsV1CommandHelper) GetAuthenticatorAndURL() (core.Authenticator, string) {
	authenticator, err := r.utils.GetAuthenticator(serviceName)
	r.utils.HandleError(err, translation.T("credentials-error"))

	serviceUrl := r.utils.GetServiceURL(projectsv1.GetServiceURLForRegion)

	return authenticator, serviceUrl
}

func (r *ProjectsV1CommandHelper) CreateServiceInstance(options projectsv1.ProjectsV1Options) {
	configurationErrorMessage := translation.T("config-error")

	projects, projectsErr := projectsv1.NewProjectsV1UsingExternalConfig(&options)
	r.utils.HandleError(projectsErr, configurationErrorMessage)

	// the cli differs from the sdk on configuration priority
	// ensure the correct priority is being used
	configErr := r.utils.PostProcessServiceConfiguration(projects.Service, serviceName)
	r.utils.HandleError(configErr, configurationErrorMessage)

	// set custom analytics header for the CLI
	customHeaders := http.Header{}
	customHeaders.Add("X-Original-User-Agent", "github.com/IBM/platform-services-go-sdk/" + version.GetPluginVersion().String())
	projects.SetDefaultHeaders(customHeaders)

	ServiceInstance = projects
}

func (r *ProjectsV1CommandHelper) InitializeServiceInstance(parentFlags *pflag.FlagSet) {
	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, parentFlags, serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	authenticator, serviceUrl := r.GetAuthenticatorAndURL()
	options := projectsv1.ProjectsV1Options{
		Authenticator: authenticator,
		// default to the contextual url, it may be overridden by an environment variable
		URL: serviceUrl,
	}

	r.CreateServiceInstance(options)
}

type RequestSender interface {
	Send(interface{}) (interface{}, *core.DetailedResponse, error)
}

// RequestSender for ListProjects command
type ListProjectsRequestSender struct {}

func (s ListProjectsRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.ListProjects(optionsModel.(*projectsv1.ListProjectsOptions))
}

type ListProjectsAllPagesRequestSender struct {}

func (s ListProjectsAllPagesRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	options := optionsModel.(*projectsv1.ListProjectsOptions)

	pager, err := ServiceInstance.NewProjectsPager(options)
	if err != nil {
		return nil, nil, err
	}

	allResults, err := pager.GetAll()
	if err != nil {
		return nil, nil, err
	}

	finalResponse := &core.DetailedResponse{}
	finalResponse.Result = map[string]interface{}{ // simulate structure of full response
		"projects": allResults,
	}

	return nil, finalResponse, nil
}

// Command Runner for ListProjects command
func NewListProjectsCommandRunner(utils Utilities, sender RequestSender, allPagesSender RequestSender) *ListProjectsCommandRunner {
	return &ListProjectsCommandRunner{utils: utils, sender: sender, allPagesSender: allPagesSender}
}

type ListProjectsCommandRunner struct {
	Start string
	Limit int64
	Complete bool
	RequiredFlags []string
	sender RequestSender
	utils Utilities
	GetAllPages bool
	allPagesSender RequestSender
}

// Command mapping: list, GetListProjectsCommand
func GetListProjectsCommand(r *ListProjectsCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list [--start START] [--limit LIMIT] [--complete COMPLETE]",
		Short: translation.T("project-list-command-short-description"),
		Long: translation.T("project-list-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "list",
		},
		Example: `  ibmcloud project list \
    --start=exampleString \
    --limit=10 \
    --complete=false`,
	}

	cmd.Flags().StringVarP(&r.Start, "start", "", "", translation.T("project-list-start-flag-description"))
	cmd.Flags().Int64VarP(&r.Limit, "limit", "", 0, translation.T("project-list-limit-flag-description"))
	cmd.Flags().BoolVarP(&r.Complete, "complete", "", false, translation.T("project-list-complete-flag-description"))
	cmd.Flags().Bool("all-pages", false, translation.T("project-list-all-pages-flag-description"))

	return cmd
}

// Primary logic for running ListProjects
func (r *ListProjectsCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.ListProjectsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "start" {
			OptionsModel.SetStart(r.Start)
		}
		if flag.Name == "limit" {
			OptionsModel.SetLimit(r.Limit)
		}
		if flag.Name == "complete" {
			OptionsModel.SetComplete(r.Complete)
		}
		if flag.Name == "all-pages" {
			r.GetAllPages = true
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *ListProjectsCommandRunner) MakeRequest(OptionsModel projectsv1.ListProjectsOptions) {
	var sender RequestSender

	if r.GetAllPages {
		sender = r.allPagesSender
	} else {
		sender = r.sender
	}

	_, DetailedResponse, ResponseErr := sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"limit",
		"total_count",
		"first",
		"last",
		"previous",
		"next",
		"projects",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetProject command
type GetProjectRequestSender struct {}

func (s GetProjectRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetProject(optionsModel.(*projectsv1.GetProjectOptions))
}

// Command Runner for GetProject command
func NewGetProjectCommandRunner(utils Utilities, sender RequestSender) *GetProjectCommandRunner {
	return &GetProjectCommandRunner{utils: utils, sender: sender}
}

type GetProjectCommandRunner struct {
	ID string
	ExcludeConfigs bool
	Complete bool
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: get, GetGetProjectCommand
func GetGetProjectCommand(r *GetProjectCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get --id ID [--exclude-configs EXCLUDE-CONFIGS] [--complete COMPLETE]",
		Short: translation.T("project-get-command-short-description"),
		Long: translation.T("project-get-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "get",
		},
		Example: `  ibmcloud project get \
    --id=exampleString \
    --exclude-configs=false \
    --complete=false`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-get-id-flag-description"))
	cmd.Flags().BoolVarP(&r.ExcludeConfigs, "exclude-configs", "", false, translation.T("project-get-exclude-configs-flag-description"))
	cmd.Flags().BoolVarP(&r.Complete, "complete", "", false, translation.T("project-get-complete-flag-description"))
	r.RequiredFlags = []string{
		"id",
	}

	return cmd
}

// Primary logic for running GetProject
func (r *GetProjectCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetProjectOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "exclude-configs" {
			OptionsModel.SetExcludeConfigs(r.ExcludeConfigs)
		}
		if flag.Name == "complete" {
			OptionsModel.SetComplete(r.Complete)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetProjectCommandRunner) MakeRequest(OptionsModel projectsv1.GetProjectOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"name",
		"description",
		"id",
		"crn",
		"configs",
		"metadata",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for DeleteProject command
type DeleteProjectRequestSender struct {}

func (s DeleteProjectRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	res, err := ServiceInstance.DeleteProject(optionsModel.(*projectsv1.DeleteProjectOptions))
	// DeleteProject returns an empty response body
	return nil, res, err
}

// Command Runner for DeleteProject command
func NewDeleteProjectCommandRunner(utils Utilities, sender RequestSender) *DeleteProjectCommandRunner {
	return &DeleteProjectCommandRunner{utils: utils, sender: sender}
}

type DeleteProjectCommandRunner struct {
	ID string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: delete, GetDeleteProjectCommand
func GetDeleteProjectCommand(r *DeleteProjectCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete --id ID",
		Short: translation.T("project-delete-command-short-description"),
		Long: translation.T("project-delete-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "delete",
		},
		Example: `  ibmcloud project delete \
    --id=exampleString`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-delete-id-flag-description"))
	r.RequiredFlags = []string{
		"id",
	}

	return cmd
}

// Primary logic for running DeleteProject
func (r *DeleteProjectCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.DeleteProjectOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *DeleteProjectCommandRunner) MakeRequest(OptionsModel projectsv1.DeleteProjectOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)
	r.utils.ProcessEmptyResponse(DetailedResponse, ResponseErr)
}

// RequestSender for ListConfigs command
type ListConfigsRequestSender struct {}

func (s ListConfigsRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.ListConfigs(optionsModel.(*projectsv1.ListConfigsOptions))
}

// Command Runner for ListConfigs command
func NewListConfigsCommandRunner(utils Utilities, sender RequestSender) *ListConfigsCommandRunner {
	return &ListConfigsCommandRunner{utils: utils, sender: sender}
}

type ListConfigsCommandRunner struct {
	ID string
	Version string
	Complete bool
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: configs, GetListConfigsCommand
func GetListConfigsCommand(r *ListConfigsCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "configs --id ID [--version VERSION] [--complete COMPLETE]",
		Short: translation.T("project-configs-command-short-description"),
		Long: translation.T("project-configs-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Example: `  ibmcloud project configs \
    --id=exampleString \
    --version=active \
    --complete=false`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-configs-id-flag-description"))
	cmd.Flags().StringVarP(&r.Version, "version", "", "", translation.T("project-configs-version-flag-description"))
	cmd.Flags().BoolVarP(&r.Complete, "complete", "", false, translation.T("project-configs-complete-flag-description"))
	r.RequiredFlags = []string{
		"id",
	}

	return cmd
}

// Primary logic for running ListConfigs
func (r *ListConfigsCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.ListConfigsOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "version" {
			OptionsModel.SetVersion(r.Version)
		}
		if flag.Name == "complete" {
			OptionsModel.SetComplete(r.Complete)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *ListConfigsCommandRunner) MakeRequest(OptionsModel projectsv1.ListConfigsOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"configs",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetConfig command
type GetConfigRequestSender struct {}

func (s GetConfigRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetConfig(optionsModel.(*projectsv1.GetConfigOptions))
}

// Command Runner for GetConfig command
func NewGetConfigCommandRunner(utils Utilities, sender RequestSender) *GetConfigCommandRunner {
	return &GetConfigCommandRunner{utils: utils, sender: sender}
}

type GetConfigCommandRunner struct {
	ID string
	ConfigID string
	Version string
	Complete bool
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: config-operation, GetGetConfigCommand
func GetGetConfigCommand(r *GetConfigCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "config-operation --id ID --config-id CONFIG-ID [--version VERSION] [--complete COMPLETE]",
		Short: translation.T("project-config-operation-command-short-description"),
		Long: translation.T("project-config-operation-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Example: `  ibmcloud project config-operation \
    --id=exampleString \
    --config-id=exampleString \
    --version=active \
    --complete=false`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-config-operation-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-config-operation-config-id-flag-description"))
	cmd.Flags().StringVarP(&r.Version, "version", "", "", translation.T("project-config-operation-version-flag-description"))
	cmd.Flags().BoolVarP(&r.Complete, "complete", "", false, translation.T("project-config-operation-complete-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running GetConfig
func (r *GetConfigCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetConfigOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
		if flag.Name == "version" {
			OptionsModel.SetVersion(r.Version)
		}
		if flag.Name == "complete" {
			OptionsModel.SetComplete(r.Complete)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetConfigCommandRunner) MakeRequest(OptionsModel projectsv1.GetConfigOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"id",
		"name",
		"labels",
		"description",
		"locator_id",
		"type",
		"input",
		"output",
		"setting",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for DeleteConfig command
type DeleteConfigRequestSender struct {}

func (s DeleteConfigRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.DeleteConfig(optionsModel.(*projectsv1.DeleteConfigOptions))
}

// Command Runner for DeleteConfig command
func NewDeleteConfigCommandRunner(utils Utilities, sender RequestSender) *DeleteConfigCommandRunner {
	return &DeleteConfigCommandRunner{utils: utils, sender: sender}
}

type DeleteConfigCommandRunner struct {
	ID string
	ConfigID string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: delete_config, GetDeleteConfigCommand
func GetDeleteConfigCommand(r *DeleteConfigCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete_config --id ID --config-id CONFIG-ID",
		Short: translation.T("project-delete_config-command-short-description"),
		Long: translation.T("project-delete_config-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "delete_config",
		},
		Example: `  ibmcloud project delete_config \
    --id=exampleString \
    --config-id=exampleString`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-delete_config-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-delete_config-config-id-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running DeleteConfig
func (r *DeleteConfigCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.DeleteConfigOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *DeleteConfigCommandRunner) MakeRequest(OptionsModel projectsv1.DeleteConfigOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"id",
		"name",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetConfigDiff command
type GetConfigDiffRequestSender struct {}

func (s GetConfigDiffRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetConfigDiff(optionsModel.(*projectsv1.GetConfigDiffOptions))
}

// Command Runner for GetConfigDiff command
func NewGetConfigDiffCommandRunner(utils Utilities, sender RequestSender) *GetConfigDiffCommandRunner {
	return &GetConfigDiffCommandRunner{utils: utils, sender: sender}
}

type GetConfigDiffCommandRunner struct {
	ID string
	ConfigID string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: config-diff, GetGetConfigDiffCommand
func GetGetConfigDiffCommand(r *GetConfigDiffCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "config-diff --id ID --config-id CONFIG-ID",
		Short: translation.T("project-config-diff-command-short-description"),
		Long: translation.T("project-config-diff-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Example: `  ibmcloud project config-diff \
    --id=exampleString \
    --config-id=exampleString`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-config-diff-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-config-diff-config-id-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running GetConfigDiff
func (r *GetConfigDiffCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetConfigDiffOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetConfigDiffCommandRunner) MakeRequest(OptionsModel projectsv1.GetConfigDiffOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"added",
		"changed",
		"removed",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for InstallConfig command
type InstallConfigRequestSender struct {}

func (s InstallConfigRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	res, err := ServiceInstance.InstallConfig(optionsModel.(*projectsv1.InstallConfigOptions))
	// InstallConfig returns an empty response body
	return nil, res, err
}

// Command Runner for InstallConfig command
func NewInstallConfigCommandRunner(utils Utilities, sender RequestSender) *InstallConfigCommandRunner {
	return &InstallConfigCommandRunner{utils: utils, sender: sender}
}

type InstallConfigCommandRunner struct {
	ID string
	ConfigID string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: install, GetInstallConfigCommand
func GetInstallConfigCommand(r *InstallConfigCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "install --id ID --config-id CONFIG-ID",
		Short: translation.T("project-install-command-short-description"),
		Long: translation.T("project-install-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "install",
		},
		Example: `  ibmcloud project install \
    --id=exampleString \
    --config-id=exampleString`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-install-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-install-config-id-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running InstallConfig
func (r *InstallConfigCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.InstallConfigOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *InstallConfigCommandRunner) MakeRequest(OptionsModel projectsv1.InstallConfigOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)
	r.utils.ProcessEmptyResponse(DetailedResponse, ResponseErr)
}

// RequestSender for UninstallConfig command
type UninstallConfigRequestSender struct {}

func (s UninstallConfigRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	res, err := ServiceInstance.UninstallConfig(optionsModel.(*projectsv1.UninstallConfigOptions))
	// UninstallConfig returns an empty response body
	return nil, res, err
}

// Command Runner for UninstallConfig command
func NewUninstallConfigCommandRunner(utils Utilities, sender RequestSender) *UninstallConfigCommandRunner {
	return &UninstallConfigCommandRunner{utils: utils, sender: sender}
}

type UninstallConfigCommandRunner struct {
	ID string
	ConfigID string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: uninstall, GetUninstallConfigCommand
func GetUninstallConfigCommand(r *UninstallConfigCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "uninstall --id ID --config-id CONFIG-ID",
		Short: translation.T("project-uninstall-command-short-description"),
		Long: translation.T("project-uninstall-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "uninstall",
		},
		Example: `  ibmcloud project uninstall \
    --id=exampleString \
    --config-id=exampleString`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-uninstall-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-uninstall-config-id-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running UninstallConfig
func (r *UninstallConfigCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.UninstallConfigOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *UninstallConfigCommandRunner) MakeRequest(OptionsModel projectsv1.UninstallConfigOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)
	r.utils.ProcessEmptyResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetSchematicsJob command
type GetSchematicsJobRequestSender struct {}

func (s GetSchematicsJobRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetSchematicsJob(optionsModel.(*projectsv1.GetSchematicsJobOptions))
}

// Command Runner for GetSchematicsJob command
func NewGetSchematicsJobCommandRunner(utils Utilities, sender RequestSender) *GetSchematicsJobCommandRunner {
	return &GetSchematicsJobCommandRunner{utils: utils, sender: sender}
}

type GetSchematicsJobCommandRunner struct {
	ID string
	ConfigID string
	Action string
	Since int64
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: get_schematics_job, GetGetSchematicsJobCommand
func GetGetSchematicsJobCommand(r *GetSchematicsJobCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get_schematics_job --id ID --config-id CONFIG-ID --action ACTION [--since SINCE]",
		Short: translation.T("project-get_schematics_job-command-short-description"),
		Long: translation.T("project-get_schematics_job-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "get_schematics_job",
		},
		Example: `  ibmcloud project get_schematics_job \
    --id=exampleString \
    --config-id=exampleString \
    --action=plan \
    --since=38`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-get_schematics_job-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-get_schematics_job-config-id-flag-description"))
	cmd.Flags().StringVarP(&r.Action, "action", "", "", translation.T("project-get_schematics_job-action-flag-description"))
	cmd.Flags().Int64VarP(&r.Since, "since", "", 0, translation.T("project-get_schematics_job-since-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
		"action",
	}

	return cmd
}

// Primary logic for running GetSchematicsJob
func (r *GetSchematicsJobCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetSchematicsJobOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
		if flag.Name == "action" {
			OptionsModel.SetAction(r.Action)
		}
		if flag.Name == "since" {
			OptionsModel.SetSince(r.Since)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetSchematicsJobCommandRunner) MakeRequest(OptionsModel projectsv1.GetSchematicsJobOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"id",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetCostEstimate command
type GetCostEstimateRequestSender struct {}

func (s GetCostEstimateRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetCostEstimate(optionsModel.(*projectsv1.GetCostEstimateOptions))
}

// Command Runner for GetCostEstimate command
func NewGetCostEstimateCommandRunner(utils Utilities, sender RequestSender) *GetCostEstimateCommandRunner {
	return &GetCostEstimateCommandRunner{utils: utils, sender: sender}
}

type GetCostEstimateCommandRunner struct {
	ID string
	ConfigID string
	Version string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: get_cost_estimate, GetGetCostEstimateCommand
func GetGetCostEstimateCommand(r *GetCostEstimateCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "get_cost_estimate --id ID --config-id CONFIG-ID [--version VERSION]",
		Short: translation.T("project-get_cost_estimate-command-short-description"),
		Long: translation.T("project-get_cost_estimate-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "get_cost_estimate",
		},
		Example: `  ibmcloud project get_cost_estimate \
    --id=exampleString \
    --config-id=exampleString \
    --version=active`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-get_cost_estimate-id-flag-description"))
	cmd.Flags().StringVarP(&r.ConfigID, "config-id", "", "", translation.T("project-get_cost_estimate-config-id-flag-description"))
	cmd.Flags().StringVarP(&r.Version, "version", "", "", translation.T("project-get_cost_estimate-version-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"config-id",
	}

	return cmd
}

// Primary logic for running GetCostEstimate
func (r *GetCostEstimateCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetCostEstimateOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "config-id" {
			OptionsModel.SetConfigID(r.ConfigID)
		}
		if flag.Name == "version" {
			OptionsModel.SetVersion(r.Version)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetCostEstimateCommandRunner) MakeRequest(OptionsModel projectsv1.GetCostEstimateOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)
	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for PostNotification command
type PostNotificationRequestSender struct {}

func (s PostNotificationRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.PostNotification(optionsModel.(*projectsv1.PostNotificationOptions))
}

// Command Runner for PostNotification command
func NewPostNotificationCommandRunner(utils Utilities, sender RequestSender) *PostNotificationCommandRunner {
	return &PostNotificationCommandRunner{utils: utils, sender: sender}
}

type PostNotificationCommandRunner struct {
	ID string
	Notifications string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: post-notification, GetPostNotificationCommand
func GetPostNotificationCommand(r *PostNotificationCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "post-notification --id ID [--notifications NOTIFICATIONS]",
		Short: translation.T("project-post-notification-command-short-description"),
		Long: translation.T("project-post-notification-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-exclude": "false",
		},
		Example: `  ibmcloud project post-notification \
    --id=exampleString \
    --notifications='[{"event": "project.create.failed", "target": "234234324-3444-4556-224232432", "source": "id.of.project.service.instance", "action_url": "url.for.project.documentation", "data": {"anyKey": "anyValue"}}]'`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-post-notification-id-flag-description"))
	cmd.Flags().StringVarP(&r.Notifications, "notifications", "", "", translation.T("project-post-notification-notifications-flag-description"))
	r.RequiredFlags = []string{
		"id",
	}

	return cmd
}

// Primary logic for running PostNotification
func (r *PostNotificationCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.PostNotificationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "notifications" {
			var Notifications []projectsv1.NotificationEvent
			err, msg := deserialize.ModelSlice(
				r.Notifications,
				"notifications",
				"NotificationEvent",
				projectsv1.UnmarshalNotificationEvent,
				&Notifications,
			)
			r.utils.HandleError(err, msg)
			OptionsModel.SetNotifications(Notifications)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *PostNotificationCommandRunner) MakeRequest(OptionsModel projectsv1.PostNotificationOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"notifications",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for GetHealth command
type GetHealthRequestSender struct {}

func (s GetHealthRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.GetHealth(optionsModel.(*projectsv1.GetHealthOptions))
}

// Command Runner for GetHealth command
func NewGetHealthCommandRunner(utils Utilities, sender RequestSender) *GetHealthCommandRunner {
	return &GetHealthCommandRunner{utils: utils, sender: sender}
}

type GetHealthCommandRunner struct {
	Info bool
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: health, GetGetHealthCommand
func GetGetHealthCommand(r *GetHealthCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "health [--info INFO]",
		Short: translation.T("project-health-command-short-description"),
		Long: translation.T("project-health-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-command": "health",
		},
		Example: `  ibmcloud project health \
    --info=false`,
	}

	cmd.Flags().BoolVarP(&r.Info, "info", "", false, translation.T("project-health-info-flag-description"))

	return cmd
}

// Primary logic for running GetHealth
func (r *GetHealthCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.GetHealthOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "info" {
			OptionsModel.SetInfo(r.Info)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *GetHealthCommandRunner) MakeRequest(OptionsModel projectsv1.GetHealthOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"name",
		"version",
		"dependencies",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for PostEventNotificationsIntegration command
type PostEventNotificationsIntegrationRequestSender struct {}

func (s PostEventNotificationsIntegrationRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.PostEventNotificationsIntegration(optionsModel.(*projectsv1.PostEventNotificationsIntegrationOptions))
}

// Command Runner for PostEventNotificationsIntegration command
func NewPostEventNotificationsIntegrationCommandRunner(utils Utilities, sender RequestSender) *PostEventNotificationsIntegrationCommandRunner {
	return &PostEventNotificationsIntegrationCommandRunner{utils: utils, sender: sender}
}

type PostEventNotificationsIntegrationCommandRunner struct {
	ID string
	InstanceCrn string
	Description string
	Name string
	Enabled bool
	Source string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: post-event-notifications-integration, GetPostEventNotificationsIntegrationCommand
func GetPostEventNotificationsIntegrationCommand(r *PostEventNotificationsIntegrationCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "post-event-notifications-integration --id ID --instance-crn INSTANCE-CRN [--description DESCRIPTION] [--name NAME] [--enabled ENABLED] [--source SOURCE]",
		Short: translation.T("project-post-event-notifications-integration-command-short-description"),
		Long: translation.T("project-post-event-notifications-integration-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-exclude": "false",
		},
		Example: `  ibmcloud project post-event-notifications-integration \
    --id=exampleString \
    --instance-crn='crn of event notifications instance' \
    --description='A sample project source' \
    --name='Project name' \
    --enabled=true \
    --source='CRN of the project instance'`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-post-event-notifications-integration-id-flag-description"))
	cmd.Flags().StringVarP(&r.InstanceCrn, "instance-crn", "", "", translation.T("project-post-event-notifications-integration-instance-crn-flag-description"))
	cmd.Flags().StringVarP(&r.Description, "description", "", "", translation.T("project-post-event-notifications-integration-description-flag-description"))
	cmd.Flags().StringVarP(&r.Name, "name", "", "", translation.T("project-post-event-notifications-integration-name-flag-description"))
	cmd.Flags().BoolVarP(&r.Enabled, "enabled", "", false, translation.T("project-post-event-notifications-integration-enabled-flag-description"))
	cmd.Flags().StringVarP(&r.Source, "source", "", "", translation.T("project-post-event-notifications-integration-source-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"instance-crn",
	}

	return cmd
}

// Primary logic for running PostEventNotificationsIntegration
func (r *PostEventNotificationsIntegrationCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.PostEventNotificationsIntegrationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "instance-crn" {
			OptionsModel.SetInstanceCrn(r.InstanceCrn)
		}
		if flag.Name == "description" {
			OptionsModel.SetDescription(r.Description)
		}
		if flag.Name == "name" {
			OptionsModel.SetName(r.Name)
		}
		if flag.Name == "enabled" {
			OptionsModel.SetEnabled(r.Enabled)
		}
		if flag.Name == "source" {
			OptionsModel.SetSource(r.Source)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *PostEventNotificationsIntegrationCommandRunner) MakeRequest(OptionsModel projectsv1.PostEventNotificationsIntegrationOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"description",
		"name",
		"enabled",
		"id",
		"type",
		"created_at",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}

// RequestSender for PostEventNotification command
type PostEventNotificationRequestSender struct {}

func (s PostEventNotificationRequestSender) Send(optionsModel interface{}) (interface{}, *core.DetailedResponse, error) {
	return ServiceInstance.PostEventNotification(optionsModel.(*projectsv1.PostEventNotificationOptions))
}

// Command Runner for PostEventNotification command
func NewPostEventNotificationCommandRunner(utils Utilities, sender RequestSender) *PostEventNotificationCommandRunner {
	return &PostEventNotificationCommandRunner{utils: utils, sender: sender}
}

type PostEventNotificationCommandRunner struct {
	ID string
	NewID string
	NewSource string
	NewDatacontenttype string
	NewIbmendefaultlong string
	NewIbmendefaultshort string
	NewIbmensourceid string
	NewSpecversion string
	NewType string
	RequiredFlags []string
	sender RequestSender
	utils Utilities
}

// Command mapping: post-event-notification, GetPostEventNotificationCommand
func GetPostEventNotificationCommand(r *PostEventNotificationCommandRunner) *cobra.Command {
	cmd := &cobra.Command{
		Use: "post-event-notification --id ID --new-id NEW-ID --new-source NEW-SOURCE [--new-datacontenttype NEW-DATACONTENTTYPE] [--new-ibmendefaultlong NEW-IBMENDEFAULTLONG] [--new-ibmendefaultshort NEW-IBMENDEFAULTSHORT] [--new-ibmensourceid NEW-IBMENSOURCEID] [--new-specversion NEW-SPECVERSION] [--new-type NEW-TYPE]",
		Short: translation.T("project-post-event-notification-command-short-description"),
		Long: translation.T("project-post-event-notification-command-long-description"),
		Run: r.Run,
		DisableFlagsInUseLine: true,
		Annotations: map[string]string{
			"x-cli-exclude": "false",
		},
		Example: `  ibmcloud project post-event-notification \
    --id=exampleString \
    --new-id=5f208fef-6b64-413c-aa07-dfed0b46abc1236 \
    --new-source='crn of project' \
    --new-datacontenttype=application/json \
    --new-ibmendefaultlong='long test notification message' \
    --new-ibmendefaultshort='Test notification' \
    --new-ibmensourceid='crn of project' \
    --new-specversion=1.0 \
    --new-type=com.ibm.cloud.project.project.test_notification`,
	}

	cmd.Flags().StringVarP(&r.ID, "id", "", "", translation.T("project-post-event-notification-id-flag-description"))
	cmd.Flags().StringVarP(&r.NewID, "new-id", "", "", translation.T("project-post-event-notification-new-id-flag-description"))
	cmd.Flags().StringVarP(&r.NewSource, "new-source", "", "", translation.T("project-post-event-notification-new-source-flag-description"))
	cmd.Flags().StringVarP(&r.NewDatacontenttype, "new-datacontenttype", "", "", translation.T("project-post-event-notification-new-datacontenttype-flag-description"))
	cmd.Flags().StringVarP(&r.NewIbmendefaultlong, "new-ibmendefaultlong", "", "", translation.T("project-post-event-notification-new-ibmendefaultlong-flag-description"))
	cmd.Flags().StringVarP(&r.NewIbmendefaultshort, "new-ibmendefaultshort", "", "", translation.T("project-post-event-notification-new-ibmendefaultshort-flag-description"))
	cmd.Flags().StringVarP(&r.NewIbmensourceid, "new-ibmensourceid", "", "", translation.T("project-post-event-notification-new-ibmensourceid-flag-description"))
	cmd.Flags().StringVarP(&r.NewSpecversion, "new-specversion", "", "", translation.T("project-post-event-notification-new-specversion-flag-description"))
	cmd.Flags().StringVarP(&r.NewType, "new-type", "", "", translation.T("project-post-event-notification-new-type-flag-description"))
	r.RequiredFlags = []string{
		"id",
		"new-id",
		"new-source",
	}

	return cmd
}

// Primary logic for running PostEventNotification
func (r *PostEventNotificationCommandRunner) Run(cmd *cobra.Command, args []string) {
	Service.InitializeServiceInstance(cmd.Flags())

	err := r.utils.ValidateRequiredFlags(r.RequiredFlags, cmd.Flags(), serviceName)
	r.utils.HandleError(err, translation.T("root-command-error"))

	r.utils.ConfirmRunningCommand()
	OptionsModel := projectsv1.PostEventNotificationOptions{}

	// optional params should only be set when they are explicitly passed by the user
	// otherwise, the default type values will be sent to the service
	FlagSet := cmd.Flags()
	FlagSet.Visit(func(flag *pflag.Flag) {
		if flag.Name == "id" {
			OptionsModel.SetID(r.ID)
		}
		if flag.Name == "new-id" {
			OptionsModel.SetNewID(r.NewID)
		}
		if flag.Name == "new-source" {
			OptionsModel.SetNewSource(r.NewSource)
		}
		if flag.Name == "new-datacontenttype" {
			OptionsModel.SetNewDatacontenttype(r.NewDatacontenttype)
		}
		if flag.Name == "new-ibmendefaultlong" {
			OptionsModel.SetNewIbmendefaultlong(r.NewIbmendefaultlong)
		}
		if flag.Name == "new-ibmendefaultshort" {
			OptionsModel.SetNewIbmendefaultshort(r.NewIbmendefaultshort)
		}
		if flag.Name == "new-ibmensourceid" {
			OptionsModel.SetNewIbmensourceid(r.NewIbmensourceid)
		}
		if flag.Name == "new-specversion" {
			OptionsModel.SetNewSpecversion(r.NewSpecversion)
		}
		if flag.Name == "new-type" {
			OptionsModel.SetNewType(r.NewType)
		}
	})

	r.MakeRequest(OptionsModel)
}

func (r *PostEventNotificationCommandRunner) MakeRequest(OptionsModel projectsv1.PostEventNotificationOptions) {
	_, DetailedResponse, ResponseErr := r.sender.Send(&OptionsModel)

	r.utils.SetTableHeaderOrder([]string{
		"datacontenttype",
		"ibmendefaultlong",
		"ibmendefaultshort",
		"ibmensourceid",
		"id",
		"source",
		"specversion",
		"type",
	})

	r.utils.ProcessResponse(DetailedResponse, ResponseErr)
}
