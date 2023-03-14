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
 * IBM OpenAPI SDK Code Generator Version: 3.66.0-d6c2d7e0-20230215-221247
 */

// Package projectv1 : Operations and models for the ProjectV1 service
package projectv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// ProjectV1 : This document is the **REST API specification** for the Projects Service. The Projects service provides
// the capability to manage infrastructure as code in IBM Cloud.
//
// API Version: 1.0.0
type ProjectV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://projects.api.test.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "project"

// ProjectV1Options : Service options
type ProjectV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewProjectV1UsingExternalConfig : constructs an instance of ProjectV1 with passed in options and external configuration.
func NewProjectV1UsingExternalConfig(options *ProjectV1Options) (project *ProjectV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	project, err = NewProjectV1(options)
	if err != nil {
		return
	}

	err = project.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = project.Service.SetServiceURL(options.URL)
	}
	return
}

// NewProjectV1 : constructs an instance of ProjectV1 with passed in options.
func NewProjectV1(options *ProjectV1Options) (service *ProjectV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &ProjectV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "project" suitable for processing requests.
func (project *ProjectV1) Clone() *ProjectV1 {
	if core.IsNil(project) {
		return nil
	}
	clone := *project
	clone.Service = project.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (project *ProjectV1) SetServiceURL(url string) error {
	return project.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (project *ProjectV1) GetServiceURL() string {
	return project.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (project *ProjectV1) SetDefaultHeaders(headers http.Header) {
	project.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (project *ProjectV1) SetEnableGzipCompression(enableGzip bool) {
	project.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (project *ProjectV1) GetEnableGzipCompression() bool {
	return project.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (project *ProjectV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	project.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (project *ProjectV1) DisableRetries() {
	project.Service.DisableRetries()
}

// CreateProject : Create a project
// Create a new project and asynchronously setup the tools to manage it. An initial pull request is created on the
// project Git repo. After approving the pull request, the user can deploy the resources that the project configures.
func (project *ProjectV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	return project.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (project *ProjectV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createProjectOptions, "createProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createProjectOptions, "createProjectOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CreateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createProjectOptions.ResourceGroup != nil {
		builder.AddQuery("resource_group", fmt.Sprint(*createProjectOptions.ResourceGroup))
	}
	if createProjectOptions.Location != nil {
		builder.AddQuery("location", fmt.Sprint(*createProjectOptions.Location))
	}

	body := make(map[string]interface{})
	if createProjectOptions.Name != nil {
		body["name"] = createProjectOptions.Name
	}
	if createProjectOptions.Description != nil {
		body["description"] = createProjectOptions.Description
	}
	if createProjectOptions.Configs != nil {
		body["configs"] = createProjectOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetProjectResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProjects : List projects
// List existing projects. Projects are sorted by ID.
func (project *ProjectV1) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ProjectListResponseSchema, response *core.DetailedResponse, err error) {
	return project.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (project *ProjectV1) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ProjectListResponseSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProjectsOptions, "listProjectsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListProjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listProjectsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listProjectsOptions.Start))
	}
	if listProjectsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listProjectsOptions.Limit))
	}
	if listProjectsOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*listProjectsOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectListResponseSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProject : Get project by ID
// Get a project definition document by the ID.
func (project *ProjectV1) GetProject(getProjectOptions *GetProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	return project.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (project *ProjectV1) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectOptions, "getProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectOptions, "getProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getProjectOptions.ExcludeConfigs != nil {
		builder.AddQuery("exclude_configs", fmt.Sprint(*getProjectOptions.ExcludeConfigs))
	}
	if getProjectOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*getProjectOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetProjectResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProject : Update a project
// Update a project by the ID.
func (project *ProjectV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *ProjectUpdate, response *core.DetailedResponse, err error) {
	return project.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (project *ProjectV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *ProjectUpdate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectOptions, "updateProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectOptions, "updateProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UpdateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	_, err = builder.SetBodyContentJSON(updateProjectOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectUpdate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteProject : Delete a project
// Delete a project document by the ID. A project can only be deleted after deleting all of its artifacts.
func (project *ProjectV1) DeleteProject(deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	return project.DeleteProjectWithContext(context.Background(), deleteProjectOptions)
}

// DeleteProjectWithContext is an alternate form of the DeleteProject method which supports a Context parameter
func (project *ProjectV1) DeleteProjectWithContext(ctx context.Context, deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteProjectOptions, "deleteProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteProjectOptions, "deleteProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteProjectOptions.Destroy != nil {
		builder.AddQuery("destroy", fmt.Sprint(*deleteProjectOptions.Destroy))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// CreateConfig : Add a new configuration
// Add a new configuration to a project.
func (project *ProjectV1) CreateConfig(createConfigOptions *CreateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.CreateConfigWithContext(context.Background(), createConfigOptions)
}

// CreateConfigWithContext is an alternate form of the CreateConfig method which supports a Context parameter
func (project *ProjectV1) CreateConfigWithContext(ctx context.Context, createConfigOptions *CreateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createConfigOptions, "createConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createConfigOptions, "createConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *createConfigOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CreateConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createConfigOptions.NewName != nil {
		body["name"] = createConfigOptions.NewName
	}
	if createConfigOptions.NewLocatorID != nil {
		body["locator_id"] = createConfigOptions.NewLocatorID
	}
	if createConfigOptions.NewID != nil {
		body["id"] = createConfigOptions.NewID
	}
	if createConfigOptions.NewLabels != nil {
		body["labels"] = createConfigOptions.NewLabels
	}
	if createConfigOptions.NewDescription != nil {
		body["description"] = createConfigOptions.NewDescription
	}
	if createConfigOptions.NewInput != nil {
		body["input"] = createConfigOptions.NewInput
	}
	if createConfigOptions.NewSetting != nil {
		body["setting"] = createConfigOptions.NewSetting
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListConfigs : List all project configuration
// Lists all of the project configurations for a specific project.
func (project *ProjectV1) ListConfigs(listConfigsOptions *ListConfigsOptions) (result *ProjectConfigList, response *core.DetailedResponse, err error) {
	return project.ListConfigsWithContext(context.Background(), listConfigsOptions)
}

// ListConfigsWithContext is an alternate form of the ListConfigs method which supports a Context parameter
func (project *ProjectV1) ListConfigsWithContext(ctx context.Context, listConfigsOptions *ListConfigsOptions) (result *ProjectConfigList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listConfigsOptions, "listConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listConfigsOptions, "listConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listConfigsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ListConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listConfigsOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*listConfigsOptions.Version))
	}
	if listConfigsOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*listConfigsOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfig : Get a project configuration
// Returns the specified project configuration in a specific project.
func (project *ProjectV1) GetConfig(getConfigOptions *GetConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.GetConfigWithContext(context.Background(), getConfigOptions)
}

// GetConfigWithContext is an alternate form of the GetConfig method which supports a Context parameter
func (project *ProjectV1) GetConfigWithContext(ctx context.Context, getConfigOptions *GetConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigOptions, "getConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigOptions, "getConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getConfigOptions.ID,
		"config_id": *getConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getConfigOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*getConfigOptions.Version))
	}
	if getConfigOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*getConfigOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateConfig : Update a configuration
// Update a configuration in a project by the ID.
func (project *ProjectV1) UpdateConfig(updateConfigOptions *UpdateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.UpdateConfigWithContext(context.Background(), updateConfigOptions)
}

// UpdateConfigWithContext is an alternate form of the UpdateConfig method which supports a Context parameter
func (project *ProjectV1) UpdateConfigWithContext(ctx context.Context, updateConfigOptions *UpdateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateConfigOptions, "updateConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateConfigOptions, "updateConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateConfigOptions.ID,
		"config_id": *updateConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UpdateConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")

	if updateConfigOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*updateConfigOptions.Complete))
	}

	_, err = builder.SetBodyContentJSON(updateConfigOptions.ProjectConfig)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteConfig : Delete a configuration in a project by ID
// Delete a configuration in a project. Deleting the configuration will also destroy all the resources deployed by the
// configuration.
func (project *ProjectV1) DeleteConfig(deleteConfigOptions *DeleteConfigOptions) (result *DeleteProjectConfigResponse, response *core.DetailedResponse, err error) {
	return project.DeleteConfigWithContext(context.Background(), deleteConfigOptions)
}

// DeleteConfigWithContext is an alternate form of the DeleteConfig method which supports a Context parameter
func (project *ProjectV1) DeleteConfigWithContext(ctx context.Context, deleteConfigOptions *DeleteConfigOptions) (result *DeleteProjectConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteConfigOptions, "deleteConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteConfigOptions, "deleteConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteConfigOptions.ID,
		"config_id": *deleteConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if deleteConfigOptions.DraftOnly != nil {
		builder.AddQuery("draft_only", fmt.Sprint(*deleteConfigOptions.DraftOnly))
	}
	if deleteConfigOptions.Destroy != nil {
		builder.AddQuery("destroy", fmt.Sprint(*deleteConfigOptions.Destroy))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteProjectConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetConfigDiff : Get a diff summary of a project configuration
// Returns a diff summary of the specified project configuration between its current draft and active version of a
// specific project.
func (project *ProjectV1) GetConfigDiff(getConfigDiffOptions *GetConfigDiffOptions) (result *ProjectConfigDiff, response *core.DetailedResponse, err error) {
	return project.GetConfigDiffWithContext(context.Background(), getConfigDiffOptions)
}

// GetConfigDiffWithContext is an alternate form of the GetConfigDiff method which supports a Context parameter
func (project *ProjectV1) GetConfigDiffWithContext(ctx context.Context, getConfigDiffOptions *GetConfigDiffOptions) (result *ProjectConfigDiff, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getConfigDiffOptions, "getConfigDiffOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getConfigDiffOptions, "getConfigDiffOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getConfigDiffOptions.ID,
		"config_id": *getConfigDiffOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/diff`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigDiffOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetConfigDiff")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigDiff)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ForceMerge : Force merge a project configuration draft
// Force the merge of the changes from the current active draft to the active configuration with an approving comment.
func (project *ProjectV1) ForceMerge(forceMergeOptions *ForceMergeOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.ForceMergeWithContext(context.Background(), forceMergeOptions)
}

// ForceMergeWithContext is an alternate form of the ForceMerge method which supports a Context parameter
func (project *ProjectV1) ForceMergeWithContext(ctx context.Context, forceMergeOptions *ForceMergeOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(forceMergeOptions, "forceMergeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(forceMergeOptions, "forceMergeOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *forceMergeOptions.ID,
		"config_id": *forceMergeOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/draft/force_merge`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range forceMergeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ForceMerge")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if forceMergeOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*forceMergeOptions.Complete))
	}

	body := make(map[string]interface{})
	if forceMergeOptions.Comment != nil {
		body["comment"] = forceMergeOptions.Comment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDraftAction : Merge or discard a configuration draft
// If a merge action is requested, the changes from the current active draft are merged to the active configuration. If
// a discard action is requested, the current draft is set to the discarded state.
func (project *ProjectV1) CreateDraftAction(createDraftActionOptions *CreateDraftActionOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.CreateDraftActionWithContext(context.Background(), createDraftActionOptions)
}

// CreateDraftActionWithContext is an alternate form of the CreateDraftAction method which supports a Context parameter
func (project *ProjectV1) CreateDraftActionWithContext(ctx context.Context, createDraftActionOptions *CreateDraftActionOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDraftActionOptions, "createDraftActionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDraftActionOptions, "createDraftActionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *createDraftActionOptions.ID,
		"config_id": *createDraftActionOptions.ConfigID,
		"action": *createDraftActionOptions.Action,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/draft/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDraftActionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CreateDraftAction")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createDraftActionOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*createDraftActionOptions.Complete))
	}

	body := make(map[string]interface{})
	if createDraftActionOptions.Comment != nil {
		body["comment"] = createDraftActionOptions.Comment
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CheckConfig : Run a validation check
// Run a validation check on a given configuration in project. The check includes creating or updating the associated
// schematics workspace with a plan job, running the CRA scans, and cost estimatation.
func (project *ProjectV1) CheckConfig(checkConfigOptions *CheckConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.CheckConfigWithContext(context.Background(), checkConfigOptions)
}

// CheckConfigWithContext is an alternate form of the CheckConfig method which supports a Context parameter
func (project *ProjectV1) CheckConfigWithContext(ctx context.Context, checkConfigOptions *CheckConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(checkConfigOptions, "checkConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(checkConfigOptions, "checkConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *checkConfigOptions.ID,
		"config_id": *checkConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/check`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "CheckConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if checkConfigOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*checkConfigOptions.XAuthRefreshToken))
	}

	if checkConfigOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*checkConfigOptions.Version))
	}
	if checkConfigOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*checkConfigOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// InstallConfig : Deploy a configuration
// Deploy a project's configuration. It is an asynchronous operation that can be tracked using the project status API.
func (project *ProjectV1) InstallConfig(installConfigOptions *InstallConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return project.InstallConfigWithContext(context.Background(), installConfigOptions)
}

// InstallConfigWithContext is an alternate form of the InstallConfig method which supports a Context parameter
func (project *ProjectV1) InstallConfigWithContext(ctx context.Context, installConfigOptions *InstallConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(installConfigOptions, "installConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(installConfigOptions, "installConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *installConfigOptions.ID,
		"config_id": *installConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/install`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range installConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "InstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if installConfigOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*installConfigOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UninstallConfig : Destroy configuration resources
// Destroy a project's configuration resources. The operation destroys all the resources that are deployed with the
// specific configuration. You can track it by using the project status API.
func (project *ProjectV1) UninstallConfig(uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
	return project.UninstallConfigWithContext(context.Background(), uninstallConfigOptions)
}

// UninstallConfigWithContext is an alternate form of the UninstallConfig method which supports a Context parameter
func (project *ProjectV1) UninstallConfigWithContext(ctx context.Context, uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uninstallConfigOptions, "uninstallConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uninstallConfigOptions, "uninstallConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *uninstallConfigOptions.ID,
		"config_id": *uninstallConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/uninstall`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uninstallConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UninstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// GetSchematicsJob : View the latest schematics job
// Fetch and find the latest schematics job that corresponds to a plan, deploy, or destroy configuration resource
// action.
func (project *ProjectV1) GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions) (result *GetActionJobResponse, response *core.DetailedResponse, err error) {
	return project.GetSchematicsJobWithContext(context.Background(), getSchematicsJobOptions)
}

// GetSchematicsJobWithContext is an alternate form of the GetSchematicsJob method which supports a Context parameter
func (project *ProjectV1) GetSchematicsJobWithContext(ctx context.Context, getSchematicsJobOptions *GetSchematicsJobOptions) (result *GetActionJobResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSchematicsJobOptions, "getSchematicsJobOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSchematicsJobOptions, "getSchematicsJobOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getSchematicsJobOptions.ID,
		"config_id": *getSchematicsJobOptions.ConfigID,
		"action": *getSchematicsJobOptions.Action,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/job/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSchematicsJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetSchematicsJob")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getSchematicsJobOptions.Since != nil {
		builder.AddQuery("since", fmt.Sprint(*getSchematicsJobOptions.Since))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetActionJobResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCostEstimate : Get the cost estimate
// Retrieve the cost estimate for a configuraton.
func (project *ProjectV1) GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions) (result *GetCostEstimateResponse, response *core.DetailedResponse, err error) {
	return project.GetCostEstimateWithContext(context.Background(), getCostEstimateOptions)
}

// GetCostEstimateWithContext is an alternate form of the GetCostEstimate method which supports a Context parameter
func (project *ProjectV1) GetCostEstimateWithContext(ctx context.Context, getCostEstimateOptions *GetCostEstimateOptions) (result *GetCostEstimateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCostEstimateOptions, "getCostEstimateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCostEstimateOptions, "getCostEstimateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getCostEstimateOptions.ID,
		"config_id": *getCostEstimateOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/cost_estimate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCostEstimateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetCostEstimate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCostEstimateOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*getCostEstimateOptions.Version))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetCostEstimateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostNotification : Add notifications
// Creates a notification event to be stored on the project definition.
func (project *ProjectV1) PostNotification(postNotificationOptions *PostNotificationOptions) (result *PostNotificationsResponse, response *core.DetailedResponse, err error) {
	return project.PostNotificationWithContext(context.Background(), postNotificationOptions)
}

// PostNotificationWithContext is an alternate form of the PostNotification method which supports a Context parameter
func (project *ProjectV1) PostNotificationWithContext(ctx context.Context, postNotificationOptions *PostNotificationOptions) (result *PostNotificationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postNotificationOptions, "postNotificationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postNotificationOptions, "postNotificationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *postNotificationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "PostNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postNotificationOptions.Notifications != nil {
		body["notifications"] = postNotificationOptions.Notifications
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostNotificationsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetNotifications : Get events by project ID
// Get all the notification events from a specific project ID.
func (project *ProjectV1) GetNotifications(getNotificationsOptions *GetNotificationsOptions) (result *GetNotificationsResponse, response *core.DetailedResponse, err error) {
	return project.GetNotificationsWithContext(context.Background(), getNotificationsOptions)
}

// GetNotificationsWithContext is an alternate form of the GetNotifications method which supports a Context parameter
func (project *ProjectV1) GetNotificationsWithContext(ctx context.Context, getNotificationsOptions *GetNotificationsOptions) (result *GetNotificationsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getNotificationsOptions, "getNotificationsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getNotificationsOptions, "getNotificationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getNotificationsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetNotificationsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteNotification : Delete a notification
// Delete a notification from a project.
// - in: query
//   name: notification_id
//   description: The ID of the project, which uniquely identifies it.
//   required: true
//   schema:
//     $ref: "#/components/schemas/Identifier".
func (project *ProjectV1) DeleteNotification(deleteNotificationOptions *DeleteNotificationOptions) (response *core.DetailedResponse, err error) {
	return project.DeleteNotificationWithContext(context.Background(), deleteNotificationOptions)
}

// DeleteNotificationWithContext is an alternate form of the DeleteNotification method which supports a Context parameter
func (project *ProjectV1) DeleteNotificationWithContext(ctx context.Context, deleteNotificationOptions *DeleteNotificationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNotificationOptions, "deleteNotificationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteNotificationOptions, "deleteNotificationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteNotificationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// ReceivePulsarCatalogEvents : Webhook for catalog events
// This is a webhook for pulsar catalog events.
func (project *ProjectV1) ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptions *ReceivePulsarCatalogEventsOptions) (response *core.DetailedResponse, err error) {
	return project.ReceivePulsarCatalogEventsWithContext(context.Background(), receivePulsarCatalogEventsOptions)
}

// ReceivePulsarCatalogEventsWithContext is an alternate form of the ReceivePulsarCatalogEvents method which supports a Context parameter
func (project *ProjectV1) ReceivePulsarCatalogEventsWithContext(ctx context.Context, receivePulsarCatalogEventsOptions *ReceivePulsarCatalogEventsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(receivePulsarCatalogEventsOptions, "receivePulsarCatalogEventsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(receivePulsarCatalogEventsOptions, "receivePulsarCatalogEventsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/pulsar/catalog_events`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range receivePulsarCatalogEventsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ReceivePulsarCatalogEvents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(receivePulsarCatalogEventsOptions.PulsarCatalogEvents)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// GetHealth : Get service health information
// Get service health information.
func (project *ProjectV1) GetHealth(getHealthOptions *GetHealthOptions) (result *Health, response *core.DetailedResponse, err error) {
	return project.GetHealthWithContext(context.Background(), getHealthOptions)
}

// GetHealthWithContext is an alternate form of the GetHealth method which supports a Context parameter
func (project *ProjectV1) GetHealthWithContext(ctx context.Context, getHealthOptions *GetHealthOptions) (result *Health, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHealthOptions, "getHealthOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/health`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHealthOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetHealth")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getHealthOptions.Info != nil {
		builder.AddQuery("info", fmt.Sprint(*getHealthOptions.Info))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHealth)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceServiceInstance : Create a new service instance
// Create a new service instance Create a service instance. When the service broker receives a provision request from
// the IBM Cloud platform, it MUST take whatever action is necessary to create a new resource. When a user creates a
// service instance from the IBM Cloud console or the IBM Cloud CLI, the IBM Cloud platform validates that the user has
// permission to create the service instance by using IBM Cloud IAM. After this validation occurs, your service broker's
// provision endpoint (PUT /v2/resource_instances/:instance_id) will be invoked. When provisioning occurs, the IBM Cloud
// platform provides the following values:
// - The IBM Cloud context is included in the context variable - The X-Broker-API-Originating-Identity will have the IBM
// IAM ID of the user that initiated the request - The parameters section will include the requested location (and
// additional parameters required by your service).
func (project *ProjectV1) ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *CreateResult, response *core.DetailedResponse, err error) {
	return project.ReplaceServiceInstanceWithContext(context.Background(), replaceServiceInstanceOptions)
}

// ReplaceServiceInstanceWithContext is an alternate form of the ReplaceServiceInstance method which supports a Context parameter
func (project *ProjectV1) ReplaceServiceInstanceWithContext(ctx context.Context, replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *CreateResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceInstanceOptions, "replaceServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceInstanceOptions, "replaceServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ReplaceServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceServiceInstanceOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*replaceServiceInstanceOptions.XBrokerApiVersion))
	}
	if replaceServiceInstanceOptions.XBrokerApiOriginatingIdentity != nil {
		builder.AddHeader("X-Broker-Api-Originating-Identity", fmt.Sprint(*replaceServiceInstanceOptions.XBrokerApiOriginatingIdentity))
	}

	if replaceServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*replaceServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if replaceServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = replaceServiceInstanceOptions.ServiceID
	}
	if replaceServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = replaceServiceInstanceOptions.PlanID
	}
	if replaceServiceInstanceOptions.Context != nil {
		body["context"] = replaceServiceInstanceOptions.Context
	}
	if replaceServiceInstanceOptions.Parameters != nil {
		body["parameters"] = replaceServiceInstanceOptions.Parameters
	}
	if replaceServiceInstanceOptions.PreviousValues != nil {
		body["previous_values"] = replaceServiceInstanceOptions.PreviousValues
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteServiceInstance : Delete a project service instance
// Delete (deprovision) a project service instance by GUID. When a service broker receives a delete request from the IBM
// Cloud platform, it MUST delete any resources it created during the provision. Usually this means that all resources
// are immediately reclaimed for future provisions.
func (project *ProjectV1) DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *DeleteResult, response *core.DetailedResponse, err error) {
	return project.DeleteServiceInstanceWithContext(context.Background(), deleteServiceInstanceOptions)
}

// DeleteServiceInstanceWithContext is an alternate form of the DeleteServiceInstance method which supports a Context parameter
func (project *ProjectV1) DeleteServiceInstanceWithContext(ctx context.Context, deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *DeleteResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceInstanceOptions, "deleteServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceInstanceOptions, "deleteServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteServiceInstanceOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*deleteServiceInstanceOptions.XBrokerApiVersion))
	}
	if deleteServiceInstanceOptions.XBrokerApiOriginatingIdentity != nil {
		builder.AddHeader("X-Broker-Api-Originating-Identity", fmt.Sprint(*deleteServiceInstanceOptions.XBrokerApiOriginatingIdentity))
	}

	builder.AddQuery("plan_id", fmt.Sprint(*deleteServiceInstanceOptions.PlanID))
	builder.AddQuery("service_id", fmt.Sprint(*deleteServiceInstanceOptions.ServiceID))
	if deleteServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*deleteServiceInstanceOptions.AcceptsIncomplete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateServiceInstance : Change of plans and service parameters in a provisioned service instance
// Allows an update to the plans and service parameters in a provisioned service instance.
func (project *ProjectV1) UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *UpdateResult, response *core.DetailedResponse, err error) {
	return project.UpdateServiceInstanceWithContext(context.Background(), updateServiceInstanceOptions)
}

// UpdateServiceInstanceWithContext is an alternate form of the UpdateServiceInstance method which supports a Context parameter
func (project *ProjectV1) UpdateServiceInstanceWithContext(ctx context.Context, updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *UpdateResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateServiceInstanceOptions, "updateServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateServiceInstanceOptions, "updateServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "UpdateServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateServiceInstanceOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*updateServiceInstanceOptions.XBrokerApiVersion))
	}
	if updateServiceInstanceOptions.XBrokerApiOriginatingIdentity != nil {
		builder.AddHeader("X-Broker-Api-Originating-Identity", fmt.Sprint(*updateServiceInstanceOptions.XBrokerApiOriginatingIdentity))
	}

	if updateServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*updateServiceInstanceOptions.AcceptsIncomplete))
	}

	_, err = builder.SetBodyContentJSON(updateServiceInstanceOptions.JSONPatchOperation)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLastOperation : Get last_operation for instance by GUID
// Retrieve previous operation for service instance by GUID (for asynchronous provision calls).
func (project *ProjectV1) GetLastOperation(getLastOperationOptions *GetLastOperationOptions) (result *GetLastOperationResult, response *core.DetailedResponse, err error) {
	return project.GetLastOperationWithContext(context.Background(), getLastOperationOptions)
}

// GetLastOperationWithContext is an alternate form of the GetLastOperation method which supports a Context parameter
func (project *ProjectV1) GetLastOperationWithContext(ctx context.Context, getLastOperationOptions *GetLastOperationOptions) (result *GetLastOperationResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLastOperationOptions, "getLastOperationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLastOperationOptions, "getLastOperationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getLastOperationOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v2/service_instances/{instance_id}/last_operation`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLastOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetLastOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLastOperationOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*getLastOperationOptions.XBrokerApiVersion))
	}

	if getLastOperationOptions.Operation != nil {
		builder.AddQuery("operation", fmt.Sprint(*getLastOperationOptions.Operation))
	}
	if getLastOperationOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getLastOperationOptions.PlanID))
	}
	if getLastOperationOptions.ServiceID != nil {
		builder.AddQuery("service_id", fmt.Sprint(*getLastOperationOptions.ServiceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetLastOperationResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceServiceInstanceState : Update the state of a provisioned service instance
// Update (disable or enable) the state of a provisioned service instance.
func (project *ProjectV1) ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	return project.ReplaceServiceInstanceStateWithContext(context.Background(), replaceServiceInstanceStateOptions)
}

// ReplaceServiceInstanceStateWithContext is an alternate form of the ReplaceServiceInstanceState method which supports a Context parameter
func (project *ProjectV1) ReplaceServiceInstanceStateWithContext(ctx context.Context, replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceInstanceStateOptions, "replaceServiceInstanceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceInstanceStateOptions, "replaceServiceInstanceStateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *replaceServiceInstanceStateOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "ReplaceServiceInstanceState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceServiceInstanceStateOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*replaceServiceInstanceStateOptions.XBrokerApiVersion))
	}

	body := make(map[string]interface{})
	if replaceServiceInstanceStateOptions.Enabled != nil {
		body["enabled"] = replaceServiceInstanceStateOptions.Enabled
	}
	if replaceServiceInstanceStateOptions.InitiatorID != nil {
		body["initiator_id"] = replaceServiceInstanceStateOptions.InitiatorID
	}
	if replaceServiceInstanceStateOptions.ReasonCode != nil {
		body["reason_code"] = replaceServiceInstanceStateOptions.ReasonCode
	}
	if replaceServiceInstanceStateOptions.PlanID != nil {
		body["plan_id"] = replaceServiceInstanceStateOptions.PlanID
	}
	if replaceServiceInstanceStateOptions.PreviousValues != nil {
		body["previous_values"] = replaceServiceInstanceStateOptions.PreviousValues
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrokerResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetServiceInstance : Get the current state information
// Retrieve the current state for the specified service instance.
func (project *ProjectV1) GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	return project.GetServiceInstanceWithContext(context.Background(), getServiceInstanceOptions)
}

// GetServiceInstanceWithContext is an alternate form of the GetServiceInstance method which supports a Context parameter
func (project *ProjectV1) GetServiceInstanceWithContext(ctx context.Context, getServiceInstanceOptions *GetServiceInstanceOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getServiceInstanceOptions, "getServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getServiceInstanceOptions, "getServiceInstanceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getServiceInstanceOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getServiceInstanceOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*getServiceInstanceOptions.XBrokerApiVersion))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBrokerResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCatalog : Get the catalog metadata
// Fetch the catalog metadata that's stored within the broker.
func (project *ProjectV1) GetCatalog(getCatalogOptions *GetCatalogOptions) (result *CatalogResponse, response *core.DetailedResponse, err error) {
	return project.GetCatalogWithContext(context.Background(), getCatalogOptions)
}

// GetCatalogWithContext is an alternate form of the GetCatalog method which supports a Context parameter
func (project *ProjectV1) GetCatalogWithContext(ctx context.Context, getCatalogOptions *GetCatalogOptions) (result *CatalogResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCatalogOptions, "getCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v2/catalog`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getCatalogOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*getCatalogOptions.XBrokerApiVersion))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostEventNotificationsIntegration : Connect to a event notifications instance
// Connects a project instance to an event notifications instance.
func (project *ProjectV1) PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions) (result *PostEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	return project.PostEventNotificationsIntegrationWithContext(context.Background(), postEventNotificationsIntegrationOptions)
}

// PostEventNotificationsIntegrationWithContext is an alternate form of the PostEventNotificationsIntegration method which supports a Context parameter
func (project *ProjectV1) PostEventNotificationsIntegrationWithContext(ctx context.Context, postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions) (result *PostEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postEventNotificationsIntegrationOptions, "postEventNotificationsIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postEventNotificationsIntegrationOptions, "postEventNotificationsIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *postEventNotificationsIntegrationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "PostEventNotificationsIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postEventNotificationsIntegrationOptions.InstanceCrn != nil {
		body["instance_crn"] = postEventNotificationsIntegrationOptions.InstanceCrn
	}
	if postEventNotificationsIntegrationOptions.Description != nil {
		body["description"] = postEventNotificationsIntegrationOptions.Description
	}
	if postEventNotificationsIntegrationOptions.EventNotificationsSourceName != nil {
		body["event_notifications_source_name"] = postEventNotificationsIntegrationOptions.EventNotificationsSourceName
	}
	if postEventNotificationsIntegrationOptions.Enabled != nil {
		body["enabled"] = postEventNotificationsIntegrationOptions.Enabled
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostEventNotificationsIntegrationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetEventNotificationsIntegration : Get event notification source details by project ID
// Gets the source details of the project from the connect event notifications instance.
func (project *ProjectV1) GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions) (result *GetEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	return project.GetEventNotificationsIntegrationWithContext(context.Background(), getEventNotificationsIntegrationOptions)
}

// GetEventNotificationsIntegrationWithContext is an alternate form of the GetEventNotificationsIntegration method which supports a Context parameter
func (project *ProjectV1) GetEventNotificationsIntegrationWithContext(ctx context.Context, getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions) (result *GetEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEventNotificationsIntegrationOptions, "getEventNotificationsIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEventNotificationsIntegrationOptions, "getEventNotificationsIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getEventNotificationsIntegrationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "GetEventNotificationsIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetEventNotificationsIntegrationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteEventNotificationsIntegration : Delete a event notifications connection
// Deletes the event notifications integration if that is where the project was onboarded to.
func (project *ProjectV1) DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions) (response *core.DetailedResponse, err error) {
	return project.DeleteEventNotificationsIntegrationWithContext(context.Background(), deleteEventNotificationsIntegrationOptions)
}

// DeleteEventNotificationsIntegrationWithContext is an alternate form of the DeleteEventNotificationsIntegration method which supports a Context parameter
func (project *ProjectV1) DeleteEventNotificationsIntegrationWithContext(ctx context.Context, deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteEventNotificationsIntegrationOptions, "deleteEventNotificationsIntegrationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteEventNotificationsIntegrationOptions, "deleteEventNotificationsIntegrationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteEventNotificationsIntegrationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "DeleteEventNotificationsIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = project.Service.Request(request, nil)

	return
}

// PostTestEventNotification : Send notification to event notifications instance
// Sends a notification to the event notifications instance.
func (project *ProjectV1) PostTestEventNotification(postTestEventNotificationOptions *PostTestEventNotificationOptions) (result *PostTestEventNotificationResponse, response *core.DetailedResponse, err error) {
	return project.PostTestEventNotificationWithContext(context.Background(), postTestEventNotificationOptions)
}

// PostTestEventNotificationWithContext is an alternate form of the PostTestEventNotification method which supports a Context parameter
func (project *ProjectV1) PostTestEventNotificationWithContext(ctx context.Context, postTestEventNotificationOptions *PostTestEventNotificationOptions) (result *PostTestEventNotificationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postTestEventNotificationOptions, "postTestEventNotificationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postTestEventNotificationOptions, "postTestEventNotificationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *postTestEventNotificationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = project.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(project.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postTestEventNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("project", "V1", "PostTestEventNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postTestEventNotificationOptions.Ibmendefaultlong != nil {
		body["ibmendefaultlong"] = postTestEventNotificationOptions.Ibmendefaultlong
	}
	if postTestEventNotificationOptions.Ibmendefaultshort != nil {
		body["ibmendefaultshort"] = postTestEventNotificationOptions.Ibmendefaultshort
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = project.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostTestEventNotificationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BrokerResult : The result of Get instance status call.
type BrokerResult struct {
	// Indicates whether the service instance is active and is meaningful if enabled is true. The default value is true if
	// not specified.
	Active *string `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *string `json:"enabled,omitempty"`

	// Indicates when the service instance was last accessed/modified/etc., and it is meaningful if enabled is true and
	// active is false. Represented as milliseconds since the epoch, but does not need to be accurate to the second/hour.
	LastActive *string `json:"last_active,omitempty"`
}

// UnmarshalBrokerResult unmarshals an instance of BrokerResult from the specified map of raw messages.
func UnmarshalBrokerResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BrokerResult)
	err = core.UnmarshalPrimitive(m, "active", &obj.Active)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_active", &obj.LastActive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponse : Response from fetching the catalog metadata stored within the broker.
type CatalogResponse struct {
	// collection of catalog services.
	Services []CatalogResponseServices `json:"services,omitempty"`
}

// UnmarshalCatalogResponse unmarshals an instance of CatalogResponse from the specified map of raw messages.
func UnmarshalCatalogResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponse)
	err = core.UnmarshalModel(m, "services", &obj.Services, UnmarshalCatalogResponseServices)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServices : Catalog service structure.
type CatalogResponseServices struct {
	// Specifies whether or not your service can be bound to applications in IBM Cloud. If bindable, it must be able to
	// return API endpoints and credentials to your service consumers.
	Bindable *bool `json:"bindable,omitempty"`

	// A short description of the service. It MUST be a non-empty string. Note that this description is not displayed by
	// the the IBM Cloud console or IBM Cloud CLI.
	Description *string `json:"description,omitempty"`

	// An identifier used to correlate this service in future requests to the broker. This MUST be globally unique within
	// the IBM Cloud platform. It MUST be a non-empty string, and using a GUID is recommended. Recommended: If you define
	// your service in the RMC, the RMC will generate a globally unique GUID service ID that you can use in your service
	// broker.
	ID *string `json:"id,omitempty"`

	// catalog service metadata.
	Metadata *CatalogResponseServicesMetadata `json:"metadata,omitempty"`

	// The service name is not your display name. Your service name must follow the follow these rules: It must be all
	// lowercase. It can't include spaces but may include hyphens (-). It must be less than 32 characters. Your service
	// name should include your company name. If your company has more then one offering your service name should include
	// both company and offering as part of the name. For example, the Compose company has offerings for Redis and
	// Elasticsearch. Sample service names on IBM Cloud for these offerings would be compose-redis and
	// compose-elasticsearch. Each of these service names have associated display names that are shown in the IBM Cloud
	// catalog: Compose Redis and Compose Elasticsearch. Another company (e.g. FastJetMail) may only have the single
	// JetMail offering, in which case the service name should be fastjetmail. Recommended: If you define your service in
	// RMC, you can export a catalog.j-son that will include the service name you defined within the RMC.
	Name *string `json:"name,omitempty"`

	// The Default is false. This specifices whether or not you support plan changes for provisioned instances. If your
	// offering supports multiple plans, and you want users to be able to change plans for a provisioned instance, you will
	// need to enable the ability for users to update their service instance by using /v2/service_instances/{instance_id}
	// PATCH.
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`

	// collection of catalog service tags.
	Tags []string `json:"tags,omitempty"`

	// A list of plans for this service that must contain at least one plan.
	Plans []CatalogResponseServicesPlans `json:"plans,omitempty"`
}

// UnmarshalCatalogResponseServices unmarshals an instance of CatalogResponseServices from the specified map of raw messages.
func UnmarshalCatalogResponseServices(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServices)
	err = core.UnmarshalPrimitive(m, "bindable", &obj.Bindable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalCatalogResponseServicesMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_updateable", &obj.PlanUpdateable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalCatalogResponseServicesPlans)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServicesMetadata : catalog service metadata.
type CatalogResponseServicesMetadata struct {
	// catalog service name.
	DisplayName *string `json:"display_name,omitempty"`

	// catalog service documentation url.
	DocumentationURL *string `json:"documentation_url,omitempty"`

	// catalog service image url.
	ImageURL *string `json:"image_url,omitempty"`

	// catalog service instructions url.
	InstructionsURL *string `json:"instructions_url,omitempty"`

	// catalog service long description.
	LongDescription *string `json:"long_description,omitempty"`

	// catalog service provider display name.
	ProviderDisplayName *string `json:"provider_display_name,omitempty"`

	// catalog service support url.
	SupportURL *string `json:"support_url,omitempty"`

	// catalog service terms url.
	TermsURL *string `json:"terms_url,omitempty"`
}

// UnmarshalCatalogResponseServicesMetadata unmarshals an instance of CatalogResponseServicesMetadata from the specified map of raw messages.
func UnmarshalCatalogResponseServicesMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesMetadata)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "documentation_url", &obj.DocumentationURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_url", &obj.ImageURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instructions_url", &obj.InstructionsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "long_description", &obj.LongDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider_display_name", &obj.ProviderDisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "support_url", &obj.SupportURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "terms_url", &obj.TermsURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServicesPlans : catalog service plan.
type CatalogResponseServicesPlans struct {
	// catalog service plan description.
	Description *string `json:"description,omitempty"`

	// catalog service plan subscription level.
	Free *bool `json:"free,omitempty"`

	// catalog service plan subscription id.
	ID *string `json:"id,omitempty"`

	// catalog service plan metadata.
	Metadata *CatalogResponseServicesPlansMetadata `json:"metadata,omitempty"`

	// catalog service plan name.
	Name *string `json:"name,omitempty"`
}

// UnmarshalCatalogResponseServicesPlans unmarshals an instance of CatalogResponseServicesPlans from the specified map of raw messages.
func UnmarshalCatalogResponseServicesPlans(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesPlans)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "free", &obj.Free)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalCatalogResponseServicesPlansMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServicesPlansMetadata : catalog service plan metadata.
type CatalogResponseServicesPlansMetadata struct {
	// catalog service plan metadata bullets.
	Bullets []string `json:"bullets,omitempty"`

	// catalog service plan metadata name.
	DisplayName *string `json:"display_name,omitempty"`
}

// UnmarshalCatalogResponseServicesPlansMetadata unmarshals an instance of CatalogResponseServicesPlansMetadata from the specified map of raw messages.
func UnmarshalCatalogResponseServicesPlansMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesPlansMetadata)
	err = core.UnmarshalPrimitive(m, "bullets", &obj.Bullets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CheckConfigOptions : The CheckConfig options.
type CheckConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration to trigger a validation check.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The IAM refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token,omitempty"`

	// The version of the configuration that the validation check should trigger against.
	Version *string `json:"version,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckConfigOptions : Instantiate CheckConfigOptions
func (*ProjectV1) NewCheckConfigOptions(id string, configID string) *CheckConfigOptions {
	return &CheckConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *CheckConfigOptions) SetID(id string) *CheckConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *CheckConfigOptions) SetConfigID(configID string) *CheckConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (_options *CheckConfigOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *CheckConfigOptions {
	_options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *CheckConfigOptions) SetVersion(version string) *CheckConfigOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *CheckConfigOptions) SetComplete(complete bool) *CheckConfigOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CheckConfigOptions) SetHeaders(param map[string]string) *CheckConfigOptions {
	options.Headers = param
	return options
}

// ConfigSettingItems : ConfigSettingItems struct
type ConfigSettingItems struct {
	// The name of the configuration setting.
	Name *string `json:"name" validate:"required"`

	// The value of a the configuration setting.
	Value *string `json:"value" validate:"required"`
}

// NewConfigSettingItems : Instantiate ConfigSettingItems (Generic Model Constructor)
func (*ProjectV1) NewConfigSettingItems(name string, value string) (_model *ConfigSettingItems, err error) {
	_model = &ConfigSettingItems{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalConfigSettingItems unmarshals an instance of ConfigSettingItems from the specified map of raw messages.
func UnmarshalConfigSettingItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigSettingItems)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateConfigOptions : The CreateConfig options.
type CreateConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"-" validate:"required,ne="`

	// The configuration name.
	NewName *string `json:"name" validate:"required"`

	// The location ID of a project configuration manual property.
	NewLocatorID *string `json:"locator_id" validate:"required"`

	// The unique ID of a project.
	NewID *string `json:"id,omitempty"`

	// A collection of configuration labels.
	NewLabels []string `json:"labels,omitempty"`

	// A project configuration description.
	NewDescription *string `json:"description,omitempty"`

	// The inputs of a Schematics template property.
	NewInput []InputVariableInput `json:"input,omitempty"`

	// An optional setting object That is passed to the cart API.
	NewSetting []ConfigSettingItems `json:"setting,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigOptions : Instantiate CreateConfigOptions
func (*ProjectV1) NewCreateConfigOptions(id string, newName string, newLocatorID string) *CreateConfigOptions {
	return &CreateConfigOptions{
		ID: core.StringPtr(id),
		NewName: core.StringPtr(newName),
		NewLocatorID: core.StringPtr(newLocatorID),
	}
}

// SetID : Allow user to set ID
func (_options *CreateConfigOptions) SetID(id string) *CreateConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetNewName : Allow user to set NewName
func (_options *CreateConfigOptions) SetNewName(newName string) *CreateConfigOptions {
	_options.NewName = core.StringPtr(newName)
	return _options
}

// SetNewLocatorID : Allow user to set NewLocatorID
func (_options *CreateConfigOptions) SetNewLocatorID(newLocatorID string) *CreateConfigOptions {
	_options.NewLocatorID = core.StringPtr(newLocatorID)
	return _options
}

// SetNewID : Allow user to set NewID
func (_options *CreateConfigOptions) SetNewID(newID string) *CreateConfigOptions {
	_options.NewID = core.StringPtr(newID)
	return _options
}

// SetNewLabels : Allow user to set NewLabels
func (_options *CreateConfigOptions) SetNewLabels(newLabels []string) *CreateConfigOptions {
	_options.NewLabels = newLabels
	return _options
}

// SetNewDescription : Allow user to set NewDescription
func (_options *CreateConfigOptions) SetNewDescription(newDescription string) *CreateConfigOptions {
	_options.NewDescription = core.StringPtr(newDescription)
	return _options
}

// SetNewInput : Allow user to set NewInput
func (_options *CreateConfigOptions) SetNewInput(newInput []InputVariableInput) *CreateConfigOptions {
	_options.NewInput = newInput
	return _options
}

// SetNewSetting : Allow user to set NewSetting
func (_options *CreateConfigOptions) SetNewSetting(newSetting []ConfigSettingItems) *CreateConfigOptions {
	_options.NewSetting = newSetting
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigOptions) SetHeaders(param map[string]string) *CreateConfigOptions {
	options.Headers = param
	return options
}

// CreateDraftActionOptions : The CreateDraftAction options.
type CreateDraftActionOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The action to the draft.
	Action *string `json:"action" validate:"required,ne="`

	// Notes on the project draft action.
	Comment *string `json:"comment,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDraftActionOptions.Action property.
// The action to the draft.
const (
	CreateDraftActionOptions_Action_Discard = "discard"
	CreateDraftActionOptions_Action_Merge = "merge"
)

// NewCreateDraftActionOptions : Instantiate CreateDraftActionOptions
func (*ProjectV1) NewCreateDraftActionOptions(id string, configID string, action string) *CreateDraftActionOptions {
	return &CreateDraftActionOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
		Action: core.StringPtr(action),
	}
}

// SetID : Allow user to set ID
func (_options *CreateDraftActionOptions) SetID(id string) *CreateDraftActionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *CreateDraftActionOptions) SetConfigID(configID string) *CreateDraftActionOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *CreateDraftActionOptions) SetAction(action string) *CreateDraftActionOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetComment : Allow user to set Comment
func (_options *CreateDraftActionOptions) SetComment(comment string) *CreateDraftActionOptions {
	_options.Comment = core.StringPtr(comment)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *CreateDraftActionOptions) SetComplete(complete bool) *CreateDraftActionOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDraftActionOptions) SetHeaders(param map[string]string) *CreateDraftActionOptions {
	options.Headers = param
	return options
}

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	// The project configurations.
	Configs []ProjectConfigInput `json:"configs,omitempty"`

	// Group name of the customized collection of resources.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// Data center locations for resource deployment.
	Location *string `json:"location,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*ProjectV1) NewCreateProjectOptions(name string) *CreateProjectOptions {
	return &CreateProjectOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *CreateProjectOptions) SetName(name string) *CreateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateProjectOptions) SetDescription(description string) *CreateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *CreateProjectOptions) SetConfigs(configs []ProjectConfigInput) *CreateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateProjectOptions) SetResourceGroup(resourceGroup string) *CreateProjectOptions {
	_options.ResourceGroup = core.StringPtr(resourceGroup)
	return _options
}

// SetLocation : Allow user to set Location
func (_options *CreateProjectOptions) SetLocation(location string) *CreateProjectOptions {
	_options.Location = core.StringPtr(location)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectOptions) SetHeaders(param map[string]string) *CreateProjectOptions {
	options.Headers = param
	return options
}

// CreateResult : Result of the provision call.
type CreateResult struct {
	// The URL of a web-based management user interface for the service instance. The URL MUST contain enough information
	// for the dashboard to identify the resource being accessed.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// For asynchronous responses, service brokers can return an identifier representing the operation. The value of this
	// field MUST be provided by the platform with requests to the last_operation endpoint in a URL encoded query
	// parameter. If present, it MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`
}

// UnmarshalCreateResult unmarshals an instance of CreateResult from the specified map of raw messages.
func UnmarshalCreateResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateResult)
	err = core.UnmarshalPrimitive(m, "dashboard_url", &obj.DashboardURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CumulativeNeedsAttention : CumulativeNeedsAttention struct
type CumulativeNeedsAttention struct {
	// The event name.
	Event *string `json:"event,omitempty"`

	// The unique ID of a project.
	EventID *string `json:"event_id,omitempty"`

	// The unique ID of a project.
	ConfigID *string `json:"config_id,omitempty"`

	// The version number of the configuration.
	ConfigVersion *int64 `json:"config_version,omitempty"`
}

// UnmarshalCumulativeNeedsAttention unmarshals an instance of CumulativeNeedsAttention from the specified map of raw messages.
func UnmarshalCumulativeNeedsAttention(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CumulativeNeedsAttention)
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_id", &obj.EventID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_id", &obj.ConfigID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config_version", &obj.ConfigVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteConfigOptions : The DeleteConfig options.
type DeleteConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The flag to determine if only the draft version should be deleted.
	DraftOnly *bool `json:"draft_only,omitempty"`

	// The flag that indicates if the resources deployed by schematics should be destroyed.
	Destroy *bool `json:"destroy,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigOptions : Instantiate DeleteConfigOptions
func (*ProjectV1) NewDeleteConfigOptions(id string, configID string) *DeleteConfigOptions {
	return &DeleteConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteConfigOptions) SetID(id string) *DeleteConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *DeleteConfigOptions) SetConfigID(configID string) *DeleteConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetDraftOnly : Allow user to set DraftOnly
func (_options *DeleteConfigOptions) SetDraftOnly(draftOnly bool) *DeleteConfigOptions {
	_options.DraftOnly = core.BoolPtr(draftOnly)
	return _options
}

// SetDestroy : Allow user to set Destroy
func (_options *DeleteConfigOptions) SetDestroy(destroy bool) *DeleteConfigOptions {
	_options.Destroy = core.BoolPtr(destroy)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigOptions) SetHeaders(param map[string]string) *DeleteConfigOptions {
	options.Headers = param
	return options
}

// DeleteEventNotificationsIntegrationOptions : The DeleteEventNotificationsIntegration options.
type DeleteEventNotificationsIntegrationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEventNotificationsIntegrationOptions : Instantiate DeleteEventNotificationsIntegrationOptions
func (*ProjectV1) NewDeleteEventNotificationsIntegrationOptions(id string) *DeleteEventNotificationsIntegrationOptions {
	return &DeleteEventNotificationsIntegrationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteEventNotificationsIntegrationOptions) SetID(id string) *DeleteEventNotificationsIntegrationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteEventNotificationsIntegrationOptions) SetHeaders(param map[string]string) *DeleteEventNotificationsIntegrationOptions {
	options.Headers = param
	return options
}

// DeleteNotificationOptions : The DeleteNotification options.
type DeleteNotificationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNotificationOptions : Instantiate DeleteNotificationOptions
func (*ProjectV1) NewDeleteNotificationOptions(id string) *DeleteNotificationOptions {
	return &DeleteNotificationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteNotificationOptions) SetID(id string) *DeleteNotificationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNotificationOptions) SetHeaders(param map[string]string) *DeleteNotificationOptions {
	options.Headers = param
	return options
}

// DeleteProjectConfigResponse : Delete configuration response.
type DeleteProjectConfigResponse struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The name of the configuration being deleted.
	Name *string `json:"name,omitempty"`
}

// UnmarshalDeleteProjectConfigResponse unmarshals an instance of DeleteProjectConfigResponse from the specified map of raw messages.
func UnmarshalDeleteProjectConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteProjectConfigResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteProjectOptions : The DeleteProject options.
type DeleteProjectOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The flag that indicates if the resources deployed by schematics should be destroyed.
	Destroy *bool `json:"destroy,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectOptions : Instantiate DeleteProjectOptions
func (*ProjectV1) NewDeleteProjectOptions(id string) *DeleteProjectOptions {
	return &DeleteProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteProjectOptions) SetID(id string) *DeleteProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetDestroy : Allow user to set Destroy
func (_options *DeleteProjectOptions) SetDestroy(destroy bool) *DeleteProjectOptions {
	_options.Destroy = core.BoolPtr(destroy)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectOptions) SetHeaders(param map[string]string) *DeleteProjectOptions {
	options.Headers = param
	return options
}

// DeleteResult : The result of deprovisioning a service instance.
type DeleteResult struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of DeleteResult
func (o *DeleteResult) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of DeleteResult
func (o *DeleteResult) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of DeleteResult
func (o *DeleteResult) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of DeleteResult
func (o *DeleteResult) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of DeleteResult
func (o *DeleteResult) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalDeleteResult unmarshals an instance of DeleteResult from the specified map of raw messages.
func UnmarshalDeleteResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteResult)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteServiceInstanceOptions : The DeleteServiceInstance options.
type DeleteServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID. It MUST be a non-empty string.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. It MUST be a non-empty
	// string.
	ServiceID *string `json:"service_id" validate:"required"`

	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Broker Api Originating Identity.
	XBrokerApiOriginatingIdentity *string `json:"X-Broker-Api-Originating-Identity,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a 422 Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceInstanceOptions : Instantiate DeleteServiceInstanceOptions
func (*ProjectV1) NewDeleteServiceInstanceOptions(instanceID string, planID string, serviceID string) *DeleteServiceInstanceOptions {
	return &DeleteServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
		PlanID: core.StringPtr(planID),
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteServiceInstanceOptions) SetInstanceID(instanceID string) *DeleteServiceInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetPlanID : Allow user to set PlanID
func (_options *DeleteServiceInstanceOptions) SetPlanID(planID string) *DeleteServiceInstanceOptions {
	_options.PlanID = core.StringPtr(planID)
	return _options
}

// SetServiceID : Allow user to set ServiceID
func (_options *DeleteServiceInstanceOptions) SetServiceID(serviceID string) *DeleteServiceInstanceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *DeleteServiceInstanceOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *DeleteServiceInstanceOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetXBrokerApiOriginatingIdentity : Allow user to set XBrokerApiOriginatingIdentity
func (_options *DeleteServiceInstanceOptions) SetXBrokerApiOriginatingIdentity(xBrokerApiOriginatingIdentity string) *DeleteServiceInstanceOptions {
	_options.XBrokerApiOriginatingIdentity = core.StringPtr(xBrokerApiOriginatingIdentity)
	return _options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (_options *DeleteServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *DeleteServiceInstanceOptions {
	_options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceInstanceOptions) SetHeaders(param map[string]string) *DeleteServiceInstanceOptions {
	options.Headers = param
	return options
}

// ForceMergeOptions : The ForceMerge options.
type ForceMergeOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Notes on the project draft action.
	Comment *string `json:"comment,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewForceMergeOptions : Instantiate ForceMergeOptions
func (*ProjectV1) NewForceMergeOptions(id string, configID string) *ForceMergeOptions {
	return &ForceMergeOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *ForceMergeOptions) SetID(id string) *ForceMergeOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *ForceMergeOptions) SetConfigID(configID string) *ForceMergeOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetComment : Allow user to set Comment
func (_options *ForceMergeOptions) SetComment(comment string) *ForceMergeOptions {
	_options.Comment = core.StringPtr(comment)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *ForceMergeOptions) SetComplete(complete bool) *ForceMergeOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ForceMergeOptions) SetHeaders(param map[string]string) *ForceMergeOptions {
	options.Headers = param
	return options
}

// GetActionJobResponse : The response of a fetching an action job.
type GetActionJobResponse struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`
}

// UnmarshalGetActionJobResponse unmarshals an instance of GetActionJobResponse from the specified map of raw messages.
func UnmarshalGetActionJobResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetActionJobResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCatalogOptions : The GetCatalog options.
type GetCatalogOptions struct {
	// Broker API Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogOptions : Instantiate GetCatalogOptions
func (*ProjectV1) NewGetCatalogOptions() *GetCatalogOptions {
	return &GetCatalogOptions{}
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *GetCatalogOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *GetCatalogOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogOptions) SetHeaders(param map[string]string) *GetCatalogOptions {
	options.Headers = param
	return options
}

// GetConfigDiffOptions : The GetConfigDiff options.
type GetConfigDiffOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigDiffOptions : Instantiate GetConfigDiffOptions
func (*ProjectV1) NewGetConfigDiffOptions(id string, configID string) *GetConfigDiffOptions {
	return &GetConfigDiffOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *GetConfigDiffOptions) SetID(id string) *GetConfigDiffOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *GetConfigDiffOptions) SetConfigID(configID string) *GetConfigDiffOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigDiffOptions) SetHeaders(param map[string]string) *GetConfigDiffOptions {
	options.Headers = param
	return options
}

// GetConfigOptions : The GetConfig options.
type GetConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The version of the configuration to return.
	Version *string `json:"version,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigOptions : Instantiate GetConfigOptions
func (*ProjectV1) NewGetConfigOptions(id string, configID string) *GetConfigOptions {
	return &GetConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *GetConfigOptions) SetID(id string) *GetConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *GetConfigOptions) SetConfigID(configID string) *GetConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *GetConfigOptions) SetVersion(version string) *GetConfigOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *GetConfigOptions) SetComplete(complete bool) *GetConfigOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetConfigOptions) SetHeaders(param map[string]string) *GetConfigOptions {
	options.Headers = param
	return options
}

// GetCostEstimateOptions : The GetCostEstimate options.
type GetCostEstimateOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration of the cost estimate to fetch.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The version of the configuration that the cost estimate will fetch.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCostEstimateOptions : Instantiate GetCostEstimateOptions
func (*ProjectV1) NewGetCostEstimateOptions(id string, configID string) *GetCostEstimateOptions {
	return &GetCostEstimateOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *GetCostEstimateOptions) SetID(id string) *GetCostEstimateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *GetCostEstimateOptions) SetConfigID(configID string) *GetCostEstimateOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *GetCostEstimateOptions) SetVersion(version string) *GetCostEstimateOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCostEstimateOptions) SetHeaders(param map[string]string) *GetCostEstimateOptions {
	options.Headers = param
	return options
}

// GetCostEstimateResponse : The cost estimate for the given configuration.
type GetCostEstimateResponse struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of GetCostEstimateResponse
func (o *GetCostEstimateResponse) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of GetCostEstimateResponse
func (o *GetCostEstimateResponse) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of GetCostEstimateResponse
func (o *GetCostEstimateResponse) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of GetCostEstimateResponse
func (o *GetCostEstimateResponse) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of GetCostEstimateResponse
func (o *GetCostEstimateResponse) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalGetCostEstimateResponse unmarshals an instance of GetCostEstimateResponse from the specified map of raw messages.
func UnmarshalGetCostEstimateResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetCostEstimateResponse)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetEventNotificationsIntegrationOptions : The GetEventNotificationsIntegration options.
type GetEventNotificationsIntegrationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEventNotificationsIntegrationOptions : Instantiate GetEventNotificationsIntegrationOptions
func (*ProjectV1) NewGetEventNotificationsIntegrationOptions(id string) *GetEventNotificationsIntegrationOptions {
	return &GetEventNotificationsIntegrationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetEventNotificationsIntegrationOptions) SetID(id string) *GetEventNotificationsIntegrationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetEventNotificationsIntegrationOptions) SetHeaders(param map[string]string) *GetEventNotificationsIntegrationOptions {
	options.Headers = param
	return options
}

// GetEventNotificationsIntegrationResponse : The resulting response of getting the source details of the event notifications integration.
type GetEventNotificationsIntegrationResponse struct {
	// A description of the instance of the event.
	Description *string `json:"description,omitempty"`

	// The name of the instance of the event.
	Name *string `json:"name,omitempty"`

	// The status of instance of the event.
	Enabled *bool `json:"enabled,omitempty"`

	// A unique ID of the instance of the event.
	ID *string `json:"id,omitempty"`

	// The type of the instance of event.
	Type *string `json:"type,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// The topic count of the instance of the event.
	TopicCount *int64 `json:"topic_count,omitempty"`

	// The topic names of the instance of the event.
	TopicNames []string `json:"topic_names,omitempty"`
}

// UnmarshalGetEventNotificationsIntegrationResponse unmarshals an instance of GetEventNotificationsIntegrationResponse from the specified map of raw messages.
func UnmarshalGetEventNotificationsIntegrationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetEventNotificationsIntegrationResponse)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "topic_count", &obj.TopicCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "topic_names", &obj.TopicNames)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetHealthOptions : The GetHealth options.
type GetHealthOptions struct {
	// Set this parameter if you want to get the version information in the output response.
	Info *bool `json:"info,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetHealthOptions : Instantiate GetHealthOptions
func (*ProjectV1) NewGetHealthOptions() *GetHealthOptions {
	return &GetHealthOptions{}
}

// SetInfo : Allow user to set Info
func (_options *GetHealthOptions) SetInfo(info bool) *GetHealthOptions {
	_options.Info = core.BoolPtr(info)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetHealthOptions) SetHeaders(param map[string]string) *GetHealthOptions {
	options.Headers = param
	return options
}

// GetLastOperationOptions : The GetLastOperation options.
type GetLastOperationOptions struct {
	// The unique instance ID generated during provisioning by the IBM Cloud platform.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// A broker-provided identifier for the operation.
	Operation *string `json:"operation,omitempty"`

	// ID of the plan from the catalog.json in your broker.
	PlanID *string `json:"plan_id,omitempty"`

	// ID of the service from the catalog.json in your service broker.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLastOperationOptions : Instantiate GetLastOperationOptions
func (*ProjectV1) NewGetLastOperationOptions(instanceID string) *GetLastOperationOptions {
	return &GetLastOperationOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetLastOperationOptions) SetInstanceID(instanceID string) *GetLastOperationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *GetLastOperationOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *GetLastOperationOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetOperation : Allow user to set Operation
func (_options *GetLastOperationOptions) SetOperation(operation string) *GetLastOperationOptions {
	_options.Operation = core.StringPtr(operation)
	return _options
}

// SetPlanID : Allow user to set PlanID
func (_options *GetLastOperationOptions) SetPlanID(planID string) *GetLastOperationOptions {
	_options.PlanID = core.StringPtr(planID)
	return _options
}

// SetServiceID : Allow user to set ServiceID
func (_options *GetLastOperationOptions) SetServiceID(serviceID string) *GetLastOperationOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLastOperationOptions) SetHeaders(param map[string]string) *GetLastOperationOptions {
	options.Headers = param
	return options
}

// GetLastOperationResult : The result of get_last_operation call.
type GetLastOperationResult struct {
	// Valid values are in progress, succeeded, and failed.
	State *string `json:"state,omitempty"`

	// A user-facing message displayed to the platform API client. Can be used to tell the user details about the status of
	// the operation.
	Description *string `json:"description,omitempty"`
}

// UnmarshalGetLastOperationResult unmarshals an instance of GetLastOperationResult from the specified map of raw messages.
func UnmarshalGetLastOperationResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetLastOperationResult)
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetNotificationsOptions : The GetNotifications options.
type GetNotificationsOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNotificationsOptions : Instantiate GetNotificationsOptions
func (*ProjectV1) NewGetNotificationsOptions(id string) *GetNotificationsOptions {
	return &GetNotificationsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetNotificationsOptions) SetID(id string) *GetNotificationsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetNotificationsOptions) SetHeaders(param map[string]string) *GetNotificationsOptions {
	options.Headers = param
	return options
}

// GetNotificationsResponse : The response from fetching notifications.
type GetNotificationsResponse struct {
	// Collection of the notification events with an ID.
	Notifications []NotificationEventWithID `json:"notifications,omitempty"`
}

// UnmarshalGetNotificationsResponse unmarshals an instance of GetNotificationsResponse from the specified map of raw messages.
func UnmarshalGetNotificationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetNotificationsResponse)
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalNotificationEventWithID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetProjectOptions : The GetProject options.
type GetProjectOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Only return with the active configuration, no drafts.
	ExcludeConfigs *bool `json:"exclude_configs,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectOptions : Instantiate GetProjectOptions
func (*ProjectV1) NewGetProjectOptions(id string) *GetProjectOptions {
	return &GetProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectOptions) SetID(id string) *GetProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetExcludeConfigs : Allow user to set ExcludeConfigs
func (_options *GetProjectOptions) SetExcludeConfigs(excludeConfigs bool) *GetProjectOptions {
	_options.ExcludeConfigs = core.BoolPtr(excludeConfigs)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *GetProjectOptions) SetComplete(complete bool) *GetProjectOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectOptions) SetHeaders(param map[string]string) *GetProjectOptions {
	options.Headers = param
	return options
}

// GetProjectResponse : The project returned in the response body.
type GetProjectResponse struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`

	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// An IBM Cloud resource name, which uniquely identifies a resource.
	Crn *string `json:"crn,omitempty"`

	// The project configurations.
	Configs []ProjectConfig `json:"configs,omitempty"`

	// Metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
}

// UnmarshalGetProjectResponse unmarshals an instance of GetProjectResponse from the specified map of raw messages.
func UnmarshalGetProjectResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetProjectResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalProjectMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSchematicsJobOptions : The GetSchematicsJob options.
type GetSchematicsJobOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration that triggered the action.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The triggered action.
	Action *string `json:"action" validate:"required,ne="`

	// The timestamp of when the action was triggered.
	Since *int64 `json:"since,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetSchematicsJobOptions.Action property.
// The triggered action.
const (
	GetSchematicsJobOptions_Action_Install = "install"
	GetSchematicsJobOptions_Action_Plan = "plan"
	GetSchematicsJobOptions_Action_Uninstall = "uninstall"
)

// NewGetSchematicsJobOptions : Instantiate GetSchematicsJobOptions
func (*ProjectV1) NewGetSchematicsJobOptions(id string, configID string, action string) *GetSchematicsJobOptions {
	return &GetSchematicsJobOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
		Action: core.StringPtr(action),
	}
}

// SetID : Allow user to set ID
func (_options *GetSchematicsJobOptions) SetID(id string) *GetSchematicsJobOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *GetSchematicsJobOptions) SetConfigID(configID string) *GetSchematicsJobOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetAction : Allow user to set Action
func (_options *GetSchematicsJobOptions) SetAction(action string) *GetSchematicsJobOptions {
	_options.Action = core.StringPtr(action)
	return _options
}

// SetSince : Allow user to set Since
func (_options *GetSchematicsJobOptions) SetSince(since int64) *GetSchematicsJobOptions {
	_options.Since = core.Int64Ptr(since)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSchematicsJobOptions) SetHeaders(param map[string]string) *GetSchematicsJobOptions {
	options.Headers = param
	return options
}

// GetServiceInstanceOptions : The GetServiceInstance options.
type GetServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// Broker API Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceInstanceOptions : Instantiate GetServiceInstanceOptions
func (*ProjectV1) NewGetServiceInstanceOptions(instanceID string) *GetServiceInstanceOptions {
	return &GetServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetServiceInstanceOptions) SetInstanceID(instanceID string) *GetServiceInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *GetServiceInstanceOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *GetServiceInstanceOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceInstanceOptions) SetHeaders(param map[string]string) *GetServiceInstanceOptions {
	options.Headers = param
	return options
}

// Health : Response data from a health check request.
type Health struct {
	// The name of the service.
	Name *string `json:"name,omitempty"`

	// The running version of the service.
	Version *string `json:"version,omitempty"`

	// The status of service dependencies.
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
}

// UnmarshalHealth unmarshals an instance of Health from the specified map of raw messages.
func UnmarshalHealth(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Health)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dependencies", &obj.Dependencies)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InputVariable : InputVariable struct
type InputVariable struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// The variable type.
	Type *string `json:"type" validate:"required"`

	// Whether the variable is required or not.
	Required *bool `json:"required,omitempty"`
}

// Constants associated with the InputVariable.Type property.
// The variable type.
const (
	InputVariable_Type_Array = "array"
	InputVariable_Type_Boolean = "boolean"
	InputVariable_Type_Float = "float"
	InputVariable_Type_Int = "int"
	InputVariable_Type_Number = "number"
	InputVariable_Type_Object = "object"
	InputVariable_Type_Password = "password"
	InputVariable_Type_String = "string"
)

// UnmarshalInputVariable unmarshals an instance of InputVariable from the specified map of raw messages.
func UnmarshalInputVariable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InputVariable)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InputVariableInput : InputVariableInput struct
type InputVariableInput struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`
}

// NewInputVariableInput : Instantiate InputVariableInput (Generic Model Constructor)
func (*ProjectV1) NewInputVariableInput(name string) (_model *InputVariableInput, err error) {
	_model = &InputVariableInput{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalInputVariableInput unmarshals an instance of InputVariableInput from the specified map of raw messages.
func UnmarshalInputVariableInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InputVariableInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallConfigOptions : The InstallConfig options.
type InstallConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration to install.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallConfigOptions : Instantiate InstallConfigOptions
func (*ProjectV1) NewInstallConfigOptions(id string, configID string) *InstallConfigOptions {
	return &InstallConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *InstallConfigOptions) SetID(id string) *InstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *InstallConfigOptions) SetConfigID(configID string) *InstallConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *InstallConfigOptions) SetComplete(complete bool) *InstallConfigOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstallConfigOptions) SetHeaders(param map[string]string) *InstallConfigOptions {
	options.Headers = param
	return options
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add = "add"
	JSONPatchOperation_Op_Copy = "copy"
	JSONPatchOperation_Op_Move = "move"
	JSONPatchOperation_Op_Remove = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*ProjectV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListConfigsOptions : The ListConfigs options.
type ListConfigsOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The version of configuration to return.
	Version *string `json:"version,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListConfigsOptions.Version property.
// The version of configuration to return.
const (
	ListConfigsOptions_Version_Active = "active"
	ListConfigsOptions_Version_Draft = "draft"
	ListConfigsOptions_Version_Mixed = "mixed"
)

// NewListConfigsOptions : Instantiate ListConfigsOptions
func (*ProjectV1) NewListConfigsOptions(id string) *ListConfigsOptions {
	return &ListConfigsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListConfigsOptions) SetID(id string) *ListConfigsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *ListConfigsOptions) SetVersion(version string) *ListConfigsOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *ListConfigsOptions) SetComplete(complete bool) *ListConfigsOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListConfigsOptions) SetHeaders(param map[string]string) *ListConfigsOptions {
	options.Headers = param
	return options
}

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {
	// Page token query parameter that is used to determine what resource to start the page after. If not specified, the
	// logical first page is returned.
	Start *string `json:"start,omitempty"`

	// Determine the maximum number of resources to return. The number of resources returned is the same, with exception of
	// the last page.
	Limit *int64 `json:"limit,omitempty"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*ProjectV1) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetStart : Allow user to set Start
func (_options *ListProjectsOptions) SetStart(start string) *ListProjectsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListProjectsOptions) SetLimit(limit int64) *ListProjectsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *ListProjectsOptions) SetComplete(complete bool) *ListProjectsOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// NotificationEvent : NotificationEvent struct
type NotificationEvent struct {
	// The type of event.
	Event *string `json:"event" validate:"required"`

	// The target of the event.
	Target *string `json:"target" validate:"required"`

	// The source of the event.
	Source *string `json:"source,omitempty"`

	// Who triggered the flow that posted the event.
	TriggeredBy *string `json:"triggered_by,omitempty"`

	// Actionable URL that users can go to as a response to the event.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewNotificationEvent : Instantiate NotificationEvent (Generic Model Constructor)
func (*ProjectV1) NewNotificationEvent(event string, target string) (_model *NotificationEvent, err error) {
	_model = &NotificationEvent{
		Event: core.StringPtr(event),
		Target: core.StringPtr(target),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNotificationEvent unmarshals an instance of NotificationEvent from the specified map of raw messages.
func UnmarshalNotificationEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationEvent)
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "triggered_by", &obj.TriggeredBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action_url", &obj.ActionURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NotificationEventWithID : NotificationEventWithID struct
type NotificationEventWithID struct {
	// The type of event.
	Event *string `json:"event" validate:"required"`

	// The target of the event.
	Target *string `json:"target" validate:"required"`

	// The source of the event.
	Source *string `json:"source,omitempty"`

	// Who triggered the flow that posted the event.
	TriggeredBy *string `json:"triggered_by,omitempty"`

	// Actionable URL that users can go to as a response to the event.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`

	// The unique ID of a project.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalNotificationEventWithID unmarshals an instance of NotificationEventWithID from the specified map of raw messages.
func UnmarshalNotificationEventWithID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationEventWithID)
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "triggered_by", &obj.TriggeredBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action_url", &obj.ActionURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NotificationEventWithStatus : NotificationEventWithStatus struct
type NotificationEventWithStatus struct {
	// The type of event.
	Event *string `json:"event" validate:"required"`

	// The target of the event.
	Target *string `json:"target" validate:"required"`

	// The source of the event.
	Source *string `json:"source,omitempty"`

	// Who triggered the flow that posted the event.
	TriggeredBy *string `json:"triggered_by,omitempty"`

	// Actionable URL that users can go to as a response to the event.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`

	// The unique ID of a project.
	ID *string `json:"id" validate:"required"`

	// Whether or not the event successfully posted.
	Status *string `json:"status,omitempty"`

	// The reasons for the status of an event.
	Reasons []map[string]interface{} `json:"reasons,omitempty"`
}

// UnmarshalNotificationEventWithStatus unmarshals an instance of NotificationEventWithStatus from the specified map of raw messages.
func UnmarshalNotificationEventWithStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationEventWithStatus)
	err = core.UnmarshalPrimitive(m, "event", &obj.Event)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "triggered_by", &obj.TriggeredBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action_url", &obj.ActionURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reasons", &obj.Reasons)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OutputValue : OutputValue struct
type OutputValue struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// A short explanation of the output value.
	Description *string `json:"description,omitempty"`

	// The output value.
	Value []string `json:"value,omitempty"`
}

// UnmarshalOutputValue unmarshals an instance of OutputValue from the specified map of raw messages.
func UnmarshalOutputValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutputValue)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationLink : A pagination link.
type PaginationLink struct {
	// The url of the pull request, which uniquely identifies it.
	Href *string `json:"href" validate:"required"`

	// A pagination token.
	Start *string `json:"start,omitempty"`
}

// UnmarshalPaginationLink unmarshals an instance of PaginationLink from the specified map of raw messages.
func UnmarshalPaginationLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostEventNotificationsIntegrationOptions : The PostEventNotificationsIntegration options.
type PostEventNotificationsIntegrationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// A CRN of the instance of the event.
	InstanceCrn *string `json:"instance_crn" validate:"required"`

	// A description of the instance of the event.
	Description *string `json:"description,omitempty"`

	// The name of the source this project is on the event notifications instance.
	EventNotificationsSourceName *string `json:"event_notifications_source_name,omitempty"`

	// A status of the instance of the event.
	Enabled *bool `json:"enabled,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostEventNotificationsIntegrationOptions : Instantiate PostEventNotificationsIntegrationOptions
func (*ProjectV1) NewPostEventNotificationsIntegrationOptions(id string, instanceCrn string) *PostEventNotificationsIntegrationOptions {
	return &PostEventNotificationsIntegrationOptions{
		ID: core.StringPtr(id),
		InstanceCrn: core.StringPtr(instanceCrn),
	}
}

// SetID : Allow user to set ID
func (_options *PostEventNotificationsIntegrationOptions) SetID(id string) *PostEventNotificationsIntegrationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetInstanceCrn : Allow user to set InstanceCrn
func (_options *PostEventNotificationsIntegrationOptions) SetInstanceCrn(instanceCrn string) *PostEventNotificationsIntegrationOptions {
	_options.InstanceCrn = core.StringPtr(instanceCrn)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *PostEventNotificationsIntegrationOptions) SetDescription(description string) *PostEventNotificationsIntegrationOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEventNotificationsSourceName : Allow user to set EventNotificationsSourceName
func (_options *PostEventNotificationsIntegrationOptions) SetEventNotificationsSourceName(eventNotificationsSourceName string) *PostEventNotificationsIntegrationOptions {
	_options.EventNotificationsSourceName = core.StringPtr(eventNotificationsSourceName)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *PostEventNotificationsIntegrationOptions) SetEnabled(enabled bool) *PostEventNotificationsIntegrationOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostEventNotificationsIntegrationOptions) SetHeaders(param map[string]string) *PostEventNotificationsIntegrationOptions {
	options.Headers = param
	return options
}

// PostEventNotificationsIntegrationResponse : The resulting response of connecting a project to a event notifications instance.
type PostEventNotificationsIntegrationResponse struct {
	// A description of the instance of the event.
	Description *string `json:"description,omitempty"`

	// A name of the instance of the event.
	Name *string `json:"name,omitempty"`

	// A status of the instance of the event.
	Enabled *bool `json:"enabled,omitempty"`

	// A unique ID of the instance of the event.
	ID *string `json:"id,omitempty"`

	// The type of instance of the event.
	Type *string `json:"type,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`
}

// UnmarshalPostEventNotificationsIntegrationResponse unmarshals an instance of PostEventNotificationsIntegrationResponse from the specified map of raw messages.
func UnmarshalPostEventNotificationsIntegrationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostEventNotificationsIntegrationResponse)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostNotificationOptions : The PostNotification options.
type PostNotificationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Collection of the notification events to post.
	Notifications []NotificationEvent `json:"notifications,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostNotificationOptions : Instantiate PostNotificationOptions
func (*ProjectV1) NewPostNotificationOptions(id string) *PostNotificationOptions {
	return &PostNotificationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *PostNotificationOptions) SetID(id string) *PostNotificationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetNotifications : Allow user to set Notifications
func (_options *PostNotificationOptions) SetNotifications(notifications []NotificationEvent) *PostNotificationOptions {
	_options.Notifications = notifications
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostNotificationOptions) SetHeaders(param map[string]string) *PostNotificationOptions {
	options.Headers = param
	return options
}

// PostNotificationsResponse : The result of a notification post.
type PostNotificationsResponse struct {
	// The collection of the notification events with status.
	Notifications []NotificationEventWithStatus `json:"notifications,omitempty"`
}

// UnmarshalPostNotificationsResponse unmarshals an instance of PostNotificationsResponse from the specified map of raw messages.
func UnmarshalPostNotificationsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostNotificationsResponse)
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalNotificationEventWithStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostTestEventNotificationOptions : The PostTestEventNotification options.
type PostTestEventNotificationOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// IBM default long message of the instance of the event.
	Ibmendefaultlong *string `json:"ibmendefaultlong,omitempty"`

	// IBM default short message of the instance of the event.
	Ibmendefaultshort *string `json:"ibmendefaultshort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostTestEventNotificationOptions : Instantiate PostTestEventNotificationOptions
func (*ProjectV1) NewPostTestEventNotificationOptions(id string) *PostTestEventNotificationOptions {
	return &PostTestEventNotificationOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *PostTestEventNotificationOptions) SetID(id string) *PostTestEventNotificationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetIbmendefaultlong : Allow user to set Ibmendefaultlong
func (_options *PostTestEventNotificationOptions) SetIbmendefaultlong(ibmendefaultlong string) *PostTestEventNotificationOptions {
	_options.Ibmendefaultlong = core.StringPtr(ibmendefaultlong)
	return _options
}

// SetIbmendefaultshort : Allow user to set Ibmendefaultshort
func (_options *PostTestEventNotificationOptions) SetIbmendefaultshort(ibmendefaultshort string) *PostTestEventNotificationOptions {
	_options.Ibmendefaultshort = core.StringPtr(ibmendefaultshort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostTestEventNotificationOptions) SetHeaders(param map[string]string) *PostTestEventNotificationOptions {
	options.Headers = param
	return options
}

// PostTestEventNotificationResponse : The response for posting a test notification to the event notifications instance.
type PostTestEventNotificationResponse struct {
	// The data content type of the instance of the event.
	Datacontenttype *string `json:"datacontenttype,omitempty"`

	// IBM default long message of the instance of the event.
	Ibmendefaultlong *string `json:"ibmendefaultlong,omitempty"`

	// IBM default short message of the instance of the event.
	Ibmendefaultshort *string `json:"ibmendefaultshort,omitempty"`

	// IBM source ID of the instance of the event.
	Ibmensourceid *string `json:"ibmensourceid,omitempty"`

	// A unique ID of the instance of the event.
	ID *string `json:"id" validate:"required"`

	// The source of the instance of the event.
	Source *string `json:"source" validate:"required"`

	// The spec version of the instance of the event.
	Specversion *string `json:"specversion,omitempty"`

	// The type of instance of the event.
	Type *string `json:"type,omitempty"`
}

// UnmarshalPostTestEventNotificationResponse unmarshals an instance of PostTestEventNotificationResponse from the specified map of raw messages.
func UnmarshalPostTestEventNotificationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostTestEventNotificationResponse)
	err = core.UnmarshalPrimitive(m, "datacontenttype", &obj.Datacontenttype)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmendefaultlong", &obj.Ibmendefaultlong)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmendefaultshort", &obj.Ibmendefaultshort)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibmensourceid", &obj.Ibmensourceid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "specversion", &obj.Specversion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfig : The project configuration.
type ProjectConfig struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The configuration name.
	Name *string `json:"name" validate:"required"`

	// A collection of configuration labels.
	Labels []string `json:"labels,omitempty"`

	// A project configuration description.
	Description *string `json:"description,omitempty"`

	// The location ID of a Project configuration manual property.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The type of a Project Config Manual Property.
	Type *string `json:"type" validate:"required"`

	// The outputs of a Schematics template property.
	Input []InputVariable `json:"input,omitempty"`

	// The outputs of a Schematics template property.
	Output []OutputValue `json:"output,omitempty"`

	// An optional setting object That is passed to the cart API.
	Setting []ConfigSettingItems `json:"setting,omitempty"`
}

// Constants associated with the ProjectConfig.Type property.
// The type of a Project Config Manual Property.
const (
	ProjectConfig_Type_SchematicsBlueprint = "schematics_blueprint"
	ProjectConfig_Type_TerraformTemplate = "terraform_template"
)

// UnmarshalProjectConfig unmarshals an instance of ProjectConfig from the specified map of raw messages.
func UnmarshalProjectConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfig)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locator_id", &obj.LocatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItems)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*ProjectV1) NewProjectConfigPatch(projectConfig *ProjectConfig) (_patch []JSONPatchOperation) {
	if (projectConfig.ID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/id"),
			Value: projectConfig.ID,
		})
	}
	if (projectConfig.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/name"),
			Value: projectConfig.Name,
		})
	}
	if (projectConfig.Labels != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/labels"),
			Value: projectConfig.Labels,
		})
	}
	if (projectConfig.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: projectConfig.Description,
		})
	}
	if (projectConfig.LocatorID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/locator_id"),
			Value: projectConfig.LocatorID,
		})
	}
	if (projectConfig.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: projectConfig.Type,
		})
	}
	if (projectConfig.Input != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/input"),
			Value: projectConfig.Input,
		})
	}
	if (projectConfig.Output != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/output"),
			Value: projectConfig.Output,
		})
	}
	if (projectConfig.Setting != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/setting"),
			Value: projectConfig.Setting,
		})
	}
	return
}

// ProjectConfigDiff : The project configuration diff summary.
type ProjectConfigDiff struct {
	// The additions to configurations in the diff summary.
	Added *ProjectConfigDiffAdded `json:"added,omitempty"`

	// The changes to configurations in the diff summary.
	Changed *ProjectConfigDiffChanged `json:"changed,omitempty"`

	// The deletions to configurations in the diff summary.
	Removed *ProjectConfigDiffRemoved `json:"removed,omitempty"`
}

// UnmarshalProjectConfigDiff unmarshals an instance of ProjectConfigDiff from the specified map of raw messages.
func UnmarshalProjectConfigDiff(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDiff)
	err = core.UnmarshalModel(m, "added", &obj.Added, UnmarshalProjectConfigDiffAdded)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "changed", &obj.Changed, UnmarshalProjectConfigDiffChanged)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "removed", &obj.Removed, UnmarshalProjectConfigDiffRemoved)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDiffAdded : The additions to configurations in the diff summary.
type ProjectConfigDiffAdded struct {
	// Collection of additions to configurations in the diff summary.
	Input []ProjectConfigDiffInputVariable `json:"input,omitempty"`
}

// UnmarshalProjectConfigDiffAdded unmarshals an instance of ProjectConfigDiffAdded from the specified map of raw messages.
func UnmarshalProjectConfigDiffAdded(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDiffAdded)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalProjectConfigDiffInputVariable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDiffChanged : The changes to configurations in the diff summary.
type ProjectConfigDiffChanged struct {
	// Collection of changes to configurations in the diff summary.
	Input []ProjectConfigDiffInputVariable `json:"input,omitempty"`
}

// UnmarshalProjectConfigDiffChanged unmarshals an instance of ProjectConfigDiffChanged from the specified map of raw messages.
func UnmarshalProjectConfigDiffChanged(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDiffChanged)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalProjectConfigDiffInputVariable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDiffInputVariable : ProjectConfigDiffInputVariable struct
type ProjectConfigDiffInputVariable struct {
	// The variable name.
	Name *string `json:"name" validate:"required"`

	// The variable type.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the ProjectConfigDiffInputVariable.Type property.
// The variable type.
const (
	ProjectConfigDiffInputVariable_Type_Array = "array"
	ProjectConfigDiffInputVariable_Type_Boolean = "boolean"
	ProjectConfigDiffInputVariable_Type_Float = "float"
	ProjectConfigDiffInputVariable_Type_Int = "int"
	ProjectConfigDiffInputVariable_Type_Number = "number"
	ProjectConfigDiffInputVariable_Type_Object = "object"
	ProjectConfigDiffInputVariable_Type_Password = "password"
	ProjectConfigDiffInputVariable_Type_String = "string"
)

// UnmarshalProjectConfigDiffInputVariable unmarshals an instance of ProjectConfigDiffInputVariable from the specified map of raw messages.
func UnmarshalProjectConfigDiffInputVariable(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDiffInputVariable)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigDiffRemoved : The deletions to configurations in the diff summary.
type ProjectConfigDiffRemoved struct {
	// Collection of deletions to configurations in the diff summary.
	Input []ProjectConfigDiffInputVariable `json:"input,omitempty"`
}

// UnmarshalProjectConfigDiffRemoved unmarshals an instance of ProjectConfigDiffRemoved from the specified map of raw messages.
func UnmarshalProjectConfigDiffRemoved(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigDiffRemoved)
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalProjectConfigDiffInputVariable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigInput : The input of a project configuration.
type ProjectConfigInput struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The configuration name.
	Name *string `json:"name" validate:"required"`

	// A collection of configuration labels.
	Labels []string `json:"labels,omitempty"`

	// A project configuration description.
	Description *string `json:"description,omitempty"`

	// The location ID of a project configuration manual property.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The inputs of a Schematics template property.
	Input []InputVariableInput `json:"input,omitempty"`

	// An optional setting object That is passed to the cart API.
	Setting []ConfigSettingItems `json:"setting,omitempty"`
}

// NewProjectConfigInput : Instantiate ProjectConfigInput (Generic Model Constructor)
func (*ProjectV1) NewProjectConfigInput(name string, locatorID string) (_model *ProjectConfigInput, err error) {
	_model = &ProjectConfigInput{
		Name: core.StringPtr(name),
		LocatorID: core.StringPtr(locatorID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectConfigInput unmarshals an instance of ProjectConfigInput from the specified map of raw messages.
func UnmarshalProjectConfigInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigInput)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locator_id", &obj.LocatorID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariableInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItems)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigList : The project configuration list.
type ProjectConfigList struct {
	// Collection list operation response schema should define array property with name "configs".
	Configs []ProjectConfig `json:"configs,omitempty"`
}

// UnmarshalProjectConfigList unmarshals an instance of ProjectConfigList from the specified map of raw messages.
func UnmarshalProjectConfigList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigList)
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListItem : ProjectListItem struct
type ProjectListItem struct {
	// The unique ID of a project.
	ID *string `json:"id,omitempty"`

	// The project name.
	Name *string `json:"name,omitempty"`

	// The project description.
	Description *string `json:"description,omitempty"`

	// Metadata of the project.
	Metadata *ProjectMetadata `json:"metadata,omitempty"`
}

// UnmarshalProjectListItem unmarshals an instance of ProjectListItem from the specified map of raw messages.
func UnmarshalProjectListItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalProjectMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListResponseSchema : Projects list.
type ProjectListResponseSchema struct {
	// A pagination limit.
	Limit *int64 `json:"limit" validate:"required"`

	// Get the occurrencies of the total projects.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// A pagination link.
	First *PaginationLink `json:"first" validate:"required"`

	// A pagination link.
	Last *PaginationLink `json:"last,omitempty"`

	// A pagination link.
	Previous *PaginationLink `json:"previous,omitempty"`

	// A pagination link.
	Next *PaginationLink `json:"next,omitempty"`

	// An array of projects.
	Projects []ProjectListItem `json:"projects,omitempty"`
}

// UnmarshalProjectListResponseSchema unmarshals an instance of ProjectListResponseSchema from the specified map of raw messages.
func UnmarshalProjectListResponseSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListResponseSchema)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProjectListItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ProjectListResponseSchema) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// ProjectMetadata : Metadata of the project.
type ProjectMetadata struct {
	// An IBM Cloud resource name, which uniquely identifies a resource.
	Crn *string `json:"crn,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The cumulative list of needs attention items of a project.
	CumulativeNeedsAttentionView []CumulativeNeedsAttention `json:"cumulative_needs_attention_view,omitempty"`

	// True to indicate the fetch of needs attention items that failed.
	CumulativeNeedsAttentionViewErr *string `json:"cumulative_needs_attention_view_err,omitempty"`

	// The location of where the project was created.
	Location *string `json:"location,omitempty"`

	// The resource group of where the project was created.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// The project status value.
	State *string `json:"state,omitempty"`

	// The CRN of the event notifications instance if one is connected to this project.
	EventNotificationsCrn *string `json:"event_notifications_crn,omitempty"`
}

// UnmarshalProjectMetadata unmarshals an instance of ProjectMetadata from the specified map of raw messages.
func UnmarshalProjectMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectMetadata)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cumulative_needs_attention_view", &obj.CumulativeNeedsAttentionView, UnmarshalCumulativeNeedsAttention)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cumulative_needs_attention_view_err", &obj.CumulativeNeedsAttentionViewErr)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "location", &obj.Location)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group", &obj.ResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "event_notifications_crn", &obj.EventNotificationsCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectUpdate : The project update request.
type ProjectUpdate struct {
	// The project name.
	Name *string `json:"name,omitempty"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`
}

// UnmarshalProjectUpdate unmarshals an instance of ProjectUpdate from the specified map of raw messages.
func UnmarshalProjectUpdate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectUpdate)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*ProjectV1) NewProjectUpdatePatch(projectUpdate *ProjectUpdate) (_patch []JSONPatchOperation) {
	if (projectUpdate.Name != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/name"),
			Value: projectUpdate.Name,
		})
	}
	if (projectUpdate.Description != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/description"),
			Value: projectUpdate.Description,
		})
	}
	return
}

// PulsarEventItems : PulsarEventItems struct
type PulsarEventItems struct {
	// The type of the event that is published and written in dot notation.
	EventType *string `json:"event_type" validate:"required"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	Timestamp *strfmt.DateTime `json:"timestamp" validate:"required"`

	// The publisher of the events, preferably written as the service's CRN.
	Publisher *string `json:"publisher" validate:"required"`

	// The IBM Cloud ID that the event is scoped to.
	AccountID *string `json:"account_id" validate:"required"`

	// The version of the payload.
	Version *string `json:"version" validate:"required"`

	// Custom event properties for a specific event.
	EventProperties map[string]interface{} `json:"event_properties,omitempty"`

	// A unique identifier for that individual event.
	EventID *string `json:"event_id,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewPulsarEventItems : Instantiate PulsarEventItems (Generic Model Constructor)
func (*ProjectV1) NewPulsarEventItems(eventType string, timestamp *strfmt.DateTime, publisher string, accountID string, version string) (_model *PulsarEventItems, err error) {
	_model = &PulsarEventItems{
		EventType: core.StringPtr(eventType),
		Timestamp: timestamp,
		Publisher: core.StringPtr(publisher),
		AccountID: core.StringPtr(accountID),
		Version: core.StringPtr(version),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of PulsarEventItems
func (o *PulsarEventItems) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of PulsarEventItems
func (o *PulsarEventItems) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of PulsarEventItems
func (o *PulsarEventItems) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of PulsarEventItems
func (o *PulsarEventItems) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of PulsarEventItems
func (o *PulsarEventItems) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.EventType != nil {
		m["event_type"] = o.EventType
	}
	if o.Timestamp != nil {
		m["timestamp"] = o.Timestamp
	}
	if o.Publisher != nil {
		m["publisher"] = o.Publisher
	}
	if o.AccountID != nil {
		m["account_id"] = o.AccountID
	}
	if o.Version != nil {
		m["version"] = o.Version
	}
	if o.EventProperties != nil {
		m["event_properties"] = o.EventProperties
	}
	if o.EventID != nil {
		m["event_id"] = o.EventID
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalPulsarEventItems unmarshals an instance of PulsarEventItems from the specified map of raw messages.
func UnmarshalPulsarEventItems(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PulsarEventItems)
	err = core.UnmarshalPrimitive(m, "event_type", &obj.EventType)
	if err != nil {
		return
	}
	delete(m, "event_type")
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	delete(m, "timestamp")
	err = core.UnmarshalPrimitive(m, "publisher", &obj.Publisher)
	if err != nil {
		return
	}
	delete(m, "publisher")
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	delete(m, "account_id")
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	delete(m, "version")
	err = core.UnmarshalPrimitive(m, "event_properties", &obj.EventProperties)
	if err != nil {
		return
	}
	delete(m, "event_properties")
	err = core.UnmarshalPrimitive(m, "event_id", &obj.EventID)
	if err != nil {
		return
	}
	delete(m, "event_id")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReceivePulsarCatalogEventsOptions : The ReceivePulsarCatalogEvents options.
type ReceivePulsarCatalogEventsOptions struct {
	// A pulsar event.
	PulsarCatalogEvents []PulsarEventItems `json:"pulsar_catalog_events" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReceivePulsarCatalogEventsOptions : Instantiate ReceivePulsarCatalogEventsOptions
func (*ProjectV1) NewReceivePulsarCatalogEventsOptions(pulsarCatalogEvents []PulsarEventItems) *ReceivePulsarCatalogEventsOptions {
	return &ReceivePulsarCatalogEventsOptions{
		PulsarCatalogEvents: pulsarCatalogEvents,
	}
}

// SetPulsarCatalogEvents : Allow user to set PulsarCatalogEvents
func (_options *ReceivePulsarCatalogEventsOptions) SetPulsarCatalogEvents(pulsarCatalogEvents []PulsarEventItems) *ReceivePulsarCatalogEventsOptions {
	_options.PulsarCatalogEvents = pulsarCatalogEvents
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReceivePulsarCatalogEventsOptions) SetHeaders(param map[string]string) *ReceivePulsarCatalogEventsOptions {
	options.Headers = param
	return options
}

// ReplaceServiceInstanceOptions : The ReplaceServiceInstance options.
type ReplaceServiceInstanceOptions struct {
	// The instance_id of a service instance is provided by the IBM Cloud platform.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the service stored in the catalog.j-son of your broker. This value should be a GUID and it MUST be a
	// non-empty string.
	ServiceID *string `json:"service_id" validate:"required"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.j-son of your
	// broker.
	PlanID *string `json:"plan_id" validate:"required"`

	// Platform specific contextual information under which the service instance is to be provisioned.
	Context []string `json:"context,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// Information about the service instance prior to the update.
	PreviousValues []string `json:"previous_values,omitempty"`

	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Broker Api Originating Identity.
	XBrokerApiOriginatingIdentity *string `json:"X-Broker-Api-Originating-Identity,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a 422 Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceInstanceOptions : Instantiate ReplaceServiceInstanceOptions
func (*ProjectV1) NewReplaceServiceInstanceOptions(instanceID string, serviceID string, planID string) *ReplaceServiceInstanceOptions {
	return &ReplaceServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
		ServiceID: core.StringPtr(serviceID),
		PlanID: core.StringPtr(planID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceServiceInstanceOptions) SetInstanceID(instanceID string) *ReplaceServiceInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetServiceID : Allow user to set ServiceID
func (_options *ReplaceServiceInstanceOptions) SetServiceID(serviceID string) *ReplaceServiceInstanceOptions {
	_options.ServiceID = core.StringPtr(serviceID)
	return _options
}

// SetPlanID : Allow user to set PlanID
func (_options *ReplaceServiceInstanceOptions) SetPlanID(planID string) *ReplaceServiceInstanceOptions {
	_options.PlanID = core.StringPtr(planID)
	return _options
}

// SetContext : Allow user to set Context
func (_options *ReplaceServiceInstanceOptions) SetContext(context []string) *ReplaceServiceInstanceOptions {
	_options.Context = context
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *ReplaceServiceInstanceOptions) SetParameters(parameters map[string]interface{}) *ReplaceServiceInstanceOptions {
	_options.Parameters = parameters
	return _options
}

// SetPreviousValues : Allow user to set PreviousValues
func (_options *ReplaceServiceInstanceOptions) SetPreviousValues(previousValues []string) *ReplaceServiceInstanceOptions {
	_options.PreviousValues = previousValues
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *ReplaceServiceInstanceOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *ReplaceServiceInstanceOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetXBrokerApiOriginatingIdentity : Allow user to set XBrokerApiOriginatingIdentity
func (_options *ReplaceServiceInstanceOptions) SetXBrokerApiOriginatingIdentity(xBrokerApiOriginatingIdentity string) *ReplaceServiceInstanceOptions {
	_options.XBrokerApiOriginatingIdentity = core.StringPtr(xBrokerApiOriginatingIdentity)
	return _options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (_options *ReplaceServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *ReplaceServiceInstanceOptions {
	_options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceInstanceOptions) SetHeaders(param map[string]string) *ReplaceServiceInstanceOptions {
	options.Headers = param
	return options
}

// ReplaceServiceInstanceStateOptions : The ReplaceServiceInstanceState options.
type ReplaceServiceInstanceStateOptions struct {
	// The instance_id of a service instance is provided by the IBM Cloud platform.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the service stored in the catalog.j-son of your broker. This value should be a GUID. It MUST be a
	// non-empty string.
	Enabled *bool `json:"enabled" validate:"required"`

	// Optional string that shows the user ID that is initiating the call.
	InitiatorID *string `json:"initiator_id,omitempty"`

	// Optional string that states the reason code for the service instance state change. Valid values are
	// IBMCLOUD_ACCT_ACTIVATE, IBMCLOUD_RECLAMATION_RESTORE, or IBMCLOUD_SERVICE_INSTANCE_BELOW_CAP for enable calls;
	// IBMCLOUD_ACCT_SUSPEND, IBMCLOUD_RECLAMATION_SCHEDULE, or IBMCLOUD_SERVICE_INSTANCE_ABOVE_CAP for disable calls; and
	// IBMCLOUD_ADMIN_REQUEST for enable and disable calls.
	ReasonCode map[string]interface{} `json:"reason_code,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.j-son of your
	// broker.
	PlanID *string `json:"plan_id,omitempty"`

	// Information about the service instance prior to the update.
	PreviousValues []string `json:"previous_values,omitempty"`

	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceInstanceStateOptions : Instantiate ReplaceServiceInstanceStateOptions
func (*ProjectV1) NewReplaceServiceInstanceStateOptions(instanceID string, enabled bool) *ReplaceServiceInstanceStateOptions {
	return &ReplaceServiceInstanceStateOptions{
		InstanceID: core.StringPtr(instanceID),
		Enabled: core.BoolPtr(enabled),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ReplaceServiceInstanceStateOptions) SetInstanceID(instanceID string) *ReplaceServiceInstanceStateOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *ReplaceServiceInstanceStateOptions) SetEnabled(enabled bool) *ReplaceServiceInstanceStateOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetInitiatorID : Allow user to set InitiatorID
func (_options *ReplaceServiceInstanceStateOptions) SetInitiatorID(initiatorID string) *ReplaceServiceInstanceStateOptions {
	_options.InitiatorID = core.StringPtr(initiatorID)
	return _options
}

// SetReasonCode : Allow user to set ReasonCode
func (_options *ReplaceServiceInstanceStateOptions) SetReasonCode(reasonCode map[string]interface{}) *ReplaceServiceInstanceStateOptions {
	_options.ReasonCode = reasonCode
	return _options
}

// SetPlanID : Allow user to set PlanID
func (_options *ReplaceServiceInstanceStateOptions) SetPlanID(planID string) *ReplaceServiceInstanceStateOptions {
	_options.PlanID = core.StringPtr(planID)
	return _options
}

// SetPreviousValues : Allow user to set PreviousValues
func (_options *ReplaceServiceInstanceStateOptions) SetPreviousValues(previousValues []string) *ReplaceServiceInstanceStateOptions {
	_options.PreviousValues = previousValues
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *ReplaceServiceInstanceStateOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *ReplaceServiceInstanceStateOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceInstanceStateOptions) SetHeaders(param map[string]string) *ReplaceServiceInstanceStateOptions {
	options.Headers = param
	return options
}

// UninstallConfigOptions : The UninstallConfig options.
type UninstallConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration to destroy configuration resources.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUninstallConfigOptions : Instantiate UninstallConfigOptions
func (*ProjectV1) NewUninstallConfigOptions(id string, configID string) *UninstallConfigOptions {
	return &UninstallConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetID : Allow user to set ID
func (_options *UninstallConfigOptions) SetID(id string) *UninstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *UninstallConfigOptions) SetConfigID(configID string) *UninstallConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UninstallConfigOptions) SetHeaders(param map[string]string) *UninstallConfigOptions {
	options.Headers = param
	return options
}

// UpdateConfigOptions : The UpdateConfig options.
type UpdateConfigOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The ID of the configuration, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The change delta of the project configuration to update.
	ProjectConfig []JSONPatchOperation `json:"project_config" validate:"required"`

	// The flag to determine if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigOptions : Instantiate UpdateConfigOptions
func (*ProjectV1) NewUpdateConfigOptions(id string, configID string, projectConfig []JSONPatchOperation) *UpdateConfigOptions {
	return &UpdateConfigOptions{
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
		ProjectConfig: projectConfig,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateConfigOptions) SetID(id string) *UpdateConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *UpdateConfigOptions) SetConfigID(configID string) *UpdateConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetProjectConfig : Allow user to set ProjectConfig
func (_options *UpdateConfigOptions) SetProjectConfig(projectConfig []JSONPatchOperation) *UpdateConfigOptions {
	_options.ProjectConfig = projectConfig
	return _options
}

// SetComplete : Allow user to set Complete
func (_options *UpdateConfigOptions) SetComplete(complete bool) *UpdateConfigOptions {
	_options.Complete = core.BoolPtr(complete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigOptions) SetHeaders(param map[string]string) *UpdateConfigOptions {
	options.Headers = param
	return options
}

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The ID of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The new project definition document.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectV1) NewUpdateProjectOptions(id string, jsonPatchOperation []JSONPatchOperation) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectOptions) SetID(id string) *UpdateProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *UpdateProjectOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *UpdateProjectOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectOptions) SetHeaders(param map[string]string) *UpdateProjectOptions {
	options.Headers = param
	return options
}

// UpdateResult : The result of deprovisioning service instance.
type UpdateResult struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of UpdateResult
func (o *UpdateResult) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of UpdateResult
func (o *UpdateResult) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of UpdateResult
func (o *UpdateResult) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of UpdateResult
func (o *UpdateResult) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of UpdateResult
func (o *UpdateResult) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalUpdateResult unmarshals an instance of UpdateResult from the specified map of raw messages.
func UnmarshalUpdateResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateResult)
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*ProjectV1) NewUpdateResultPatch(updateResult *UpdateResult) (_patch []JSONPatchOperation) {
	for key, value := range updateResult.additionalProperties {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/" + key),
			Value: value,
		})
	}
	return
}

// UpdateServiceInstanceOptions : The UpdateServiceInstance options.
type UpdateServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// It contains the query filters and the search token that is initally set to null or undefined.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Broker API Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Broker Api Originating Identity.
	XBrokerApiOriginatingIdentity *string `json:"X-Broker-Api-Originating-Identity,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a 422 Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateServiceInstanceOptions : Instantiate UpdateServiceInstanceOptions
func (*ProjectV1) NewUpdateServiceInstanceOptions(instanceID string, jsonPatchOperation []JSONPatchOperation) *UpdateServiceInstanceOptions {
	return &UpdateServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
		JSONPatchOperation: jsonPatchOperation,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateServiceInstanceOptions) SetInstanceID(instanceID string) *UpdateServiceInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetJSONPatchOperation : Allow user to set JSONPatchOperation
func (_options *UpdateServiceInstanceOptions) SetJSONPatchOperation(jsonPatchOperation []JSONPatchOperation) *UpdateServiceInstanceOptions {
	_options.JSONPatchOperation = jsonPatchOperation
	return _options
}

// SetXBrokerApiVersion : Allow user to set XBrokerApiVersion
func (_options *UpdateServiceInstanceOptions) SetXBrokerApiVersion(xBrokerApiVersion string) *UpdateServiceInstanceOptions {
	_options.XBrokerApiVersion = core.StringPtr(xBrokerApiVersion)
	return _options
}

// SetXBrokerApiOriginatingIdentity : Allow user to set XBrokerApiOriginatingIdentity
func (_options *UpdateServiceInstanceOptions) SetXBrokerApiOriginatingIdentity(xBrokerApiOriginatingIdentity string) *UpdateServiceInstanceOptions {
	_options.XBrokerApiOriginatingIdentity = core.StringPtr(xBrokerApiOriginatingIdentity)
	return _options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (_options *UpdateServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *UpdateServiceInstanceOptions {
	_options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateServiceInstanceOptions) SetHeaders(param map[string]string) *UpdateServiceInstanceOptions {
	options.Headers = param
	return options
}

//
// ProjectsPager can be used to simplify the use of the "ListProjects" method.
//
type ProjectsPager struct {
	hasNext bool
	options *ListProjectsOptions
	client  *ProjectV1
	pageContext struct {
		next *string
	}
}

// NewProjectsPager returns a new ProjectsPager instance.
func (project *ProjectV1) NewProjectsPager(options *ListProjectsOptions) (pager *ProjectsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListProjectsOptions = *options
	pager = &ProjectsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  project,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *ProjectsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *ProjectsPager) GetNextWithContext(ctx context.Context) (page []ProjectListItem, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListProjectsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Projects

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *ProjectsPager) GetAllWithContext(ctx context.Context) (allItems []ProjectListItem, err error) {
	for pager.HasNext() {
		var nextPage []ProjectListItem
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetNext() (page []ProjectListItem, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetAll() (allItems []ProjectListItem, err error) {
	return pager.GetAllWithContext(context.Background())
}
