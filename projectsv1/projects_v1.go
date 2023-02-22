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

// Package projectsv1 : Operations and models for the ProjectsV1 service
package projectsv1

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

// ProjectsV1 : This document is the **REST API specification** for the Projects Service. The Projects service provides
// the capability to manage infrastructure as code in IBM Cloud.
//
// API Version: 1.0.0
type ProjectsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://projects-api.projects-stage-us-south-c-324a0a89696c3783407d3a435ca143c0-0001.us-south.containers.appdomain.cloud"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "projects"

// ProjectsV1Options : Service options
type ProjectsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewProjectsV1UsingExternalConfig : constructs an instance of ProjectsV1 with passed in options and external configuration.
func NewProjectsV1UsingExternalConfig(options *ProjectsV1Options) (projects *ProjectsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	projects, err = NewProjectsV1(options)
	if err != nil {
		return
	}

	err = projects.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = projects.Service.SetServiceURL(options.URL)
	}
	return
}

// NewProjectsV1 : constructs an instance of ProjectsV1 with passed in options.
func NewProjectsV1(options *ProjectsV1Options) (service *ProjectsV1, err error) {
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

	service = &ProjectsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "projects" suitable for processing requests.
func (projects *ProjectsV1) Clone() *ProjectsV1 {
	if core.IsNil(projects) {
		return nil
	}
	clone := *projects
	clone.Service = projects.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (projects *ProjectsV1) SetServiceURL(url string) error {
	return projects.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (projects *ProjectsV1) GetServiceURL() string {
	return projects.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (projects *ProjectsV1) SetDefaultHeaders(headers http.Header) {
	projects.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (projects *ProjectsV1) SetEnableGzipCompression(enableGzip bool) {
	projects.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (projects *ProjectsV1) GetEnableGzipCompression() bool {
	return projects.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (projects *ProjectsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	projects.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (projects *ProjectsV1) DisableRetries() {
	projects.Service.DisableRetries()
}

// CreateProject : Create a Project
// Create a new Project, which asynchronously setup the tools to manage it and creates the initial Pull Request on the
// Project git repo. After approving the PR, the user can deploy the resources that the Project configures.
func (projects *ProjectsV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	return projects.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (projects *ProjectsV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "CreateProject")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// ListProjects : List Projects
// List existing Projects. Projects are sorted by id.
func (projects *ProjectsV1) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ProjectListResponseSchema, response *core.DetailedResponse, err error) {
	return projects.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (projects *ProjectsV1) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ProjectListResponseSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listProjectsOptions, "listProjectsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ListProjects")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetProject : Get project by id
// Get a project definition document by id.
func (projects *ProjectsV1) GetProject(getProjectOptions *GetProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
	return projects.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (projects *ProjectsV1) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *GetProjectResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProject")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// UpdateProject : Update a project by id
// Update a project.
func (projects *ProjectsV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *ProjectUpdate, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *ProjectUpdate, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectOptions.Name != nil {
		body["name"] = updateProjectOptions.Name
	}
	if updateProjectOptions.Description != nil {
		body["description"] = updateProjectOptions.Description
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// DeleteProject : Delete a project by id
// Delete a project document. A project can only be deleted after deleting all its artifacts.
func (projects *ProjectsV1) DeleteProject(deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
	return projects.DeleteProjectWithContext(context.Background(), deleteProjectOptions)
}

// DeleteProjectWithContext is an alternate form of the DeleteProject method which supports a Context parameter
func (projects *ProjectsV1) DeleteProjectWithContext(ctx context.Context, deleteProjectOptions *DeleteProjectOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeleteProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// CreateConfig : Add a new config to a project
// Add a new config to a project.
func (projects *ProjectsV1) CreateConfig(createConfigOptions *CreateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return projects.CreateConfigWithContext(context.Background(), createConfigOptions)
}

// CreateConfigWithContext is an alternate form of the CreateConfig method which supports a Context parameter
func (projects *ProjectsV1) CreateConfigWithContext(ctx context.Context, createConfigOptions *CreateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "CreateConfig")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// ListConfigs : List all project configs
// Returns all project configs for a given project.
func (projects *ProjectsV1) ListConfigs(listConfigsOptions *ListConfigsOptions) (result *ProjectConfigList, response *core.DetailedResponse, err error) {
	return projects.ListConfigsWithContext(context.Background(), listConfigsOptions)
}

// ListConfigsWithContext is an alternate form of the ListConfigs method which supports a Context parameter
func (projects *ProjectsV1) ListConfigsWithContext(ctx context.Context, listConfigsOptions *ListConfigsOptions) (result *ProjectConfigList, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ListConfigs")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetConfig : Get a project config
// Returns the specified project config in a given project.
func (projects *ProjectsV1) GetConfig(getConfigOptions *GetConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return projects.GetConfigWithContext(context.Background(), getConfigOptions)
}

// GetConfigWithContext is an alternate form of the GetConfig method which supports a Context parameter
func (projects *ProjectsV1) GetConfigWithContext(ctx context.Context, getConfigOptions *GetConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetConfig")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// UpdateConfig : Update a config in a project by id
// Update a config in a project.
func (projects *ProjectsV1) UpdateConfig(updateConfigOptions *UpdateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return projects.UpdateConfigWithContext(context.Background(), updateConfigOptions)
}

// UpdateConfigWithContext is an alternate form of the UpdateConfig method which supports a Context parameter
func (projects *ProjectsV1) UpdateConfigWithContext(ctx context.Context, updateConfigOptions *UpdateConfigOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateConfigOptions.ProjectConfig)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = projects.Service.Request(request, &rawResponse)
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

// DeleteConfig : Delete a config in a project by id
// Delete a config in a project. Deleting the config will also destroy all the resources deployed by the config.
func (projects *ProjectsV1) DeleteConfig(deleteConfigOptions *DeleteConfigOptions) (result *DeleteProjectConfigResponse, response *core.DetailedResponse, err error) {
	return projects.DeleteConfigWithContext(context.Background(), deleteConfigOptions)
}

// DeleteConfigWithContext is an alternate form of the DeleteConfig method which supports a Context parameter
func (projects *ProjectsV1) DeleteConfigWithContext(ctx context.Context, deleteConfigOptions *DeleteConfigOptions) (result *DeleteProjectConfigResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeleteConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetConfigDiff : Get a diff summary of a project config
// Returns a diff summary of the specified project config between its current draft and active version of a given
// project.
func (projects *ProjectsV1) GetConfigDiff(getConfigDiffOptions *GetConfigDiffOptions) (result *ProjectConfigDiff, response *core.DetailedResponse, err error) {
	return projects.GetConfigDiffWithContext(context.Background(), getConfigDiffOptions)
}

// GetConfigDiffWithContext is an alternate form of the GetConfigDiff method which supports a Context parameter
func (projects *ProjectsV1) GetConfigDiffWithContext(ctx context.Context, getConfigDiffOptions *GetConfigDiffOptions) (result *ProjectConfigDiff, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/diff`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConfigDiffOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetConfigDiff")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = projects.Service.Request(request, &rawResponse)
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

// CreateDraftAction : Merge or discard a project config draft
// If a merge action is requested, the changes from the current active draft will be merged to the active config. If a
// discard action is requested, the current draft will be set to discarded state.
func (projects *ProjectsV1) CreateDraftAction(createDraftActionOptions *CreateDraftActionOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
	return projects.CreateDraftActionWithContext(context.Background(), createDraftActionOptions)
}

// CreateDraftActionWithContext is an alternate form of the CreateDraftAction method which supports a Context parameter
func (projects *ProjectsV1) CreateDraftActionWithContext(ctx context.Context, createDraftActionOptions *CreateDraftActionOptions) (result *ProjectConfig, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/draft/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDraftActionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "CreateDraftAction")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

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
	response, err = projects.Service.Request(request, &rawResponse)
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

// CheckConfig : Run a validation check to a given configuration in project
// Run a validation check to a given configuration in project. The check includes creating or updating the associated
// schematics workspace with a plan job, running the CRA scans and cost estimate.
func (projects *ProjectsV1) CheckConfig(checkConfigOptions *CheckConfigOptions) (response *core.DetailedResponse, err error) {
	return projects.CheckConfigWithContext(context.Background(), checkConfigOptions)
}

// CheckConfigWithContext is an alternate form of the CheckConfig method which supports a Context parameter
func (projects *ProjectsV1) CheckConfigWithContext(ctx context.Context, checkConfigOptions *CheckConfigOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/check`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "CheckConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if checkConfigOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*checkConfigOptions.XAuthRefreshToken))
	}

	if checkConfigOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*checkConfigOptions.Version))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// InstallConfig : Install a Config
// Install a project's configuration. It is an asynchronous operation that can be tracked using the project status api.
func (projects *ProjectsV1) InstallConfig(installConfigOptions *InstallConfigOptions) (response *core.DetailedResponse, err error) {
	return projects.InstallConfigWithContext(context.Background(), installConfigOptions)
}

// InstallConfigWithContext is an alternate form of the InstallConfig method which supports a Context parameter
func (projects *ProjectsV1) InstallConfigWithContext(ctx context.Context, installConfigOptions *InstallConfigOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/install`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range installConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "InstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// UninstallConfig : Uninstall a Config
// Uninstall a project's configuration. The operation uninstall all the resources deployed with the given configuration.
// You can track it by using the project status api.
func (projects *ProjectsV1) UninstallConfig(uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
	return projects.UninstallConfigWithContext(context.Background(), uninstallConfigOptions)
}

// UninstallConfigWithContext is an alternate form of the UninstallConfig method which supports a Context parameter
func (projects *ProjectsV1) UninstallConfigWithContext(ctx context.Context, uninstallConfigOptions *UninstallConfigOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/uninstall`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uninstallConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UninstallConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// GetSchematicsJob : Fetch and find the latest schematics job corresponds to a given config action
// Fetch and find the latest schematics job corresponds to a plan, install or uninstall action.
func (projects *ProjectsV1) GetSchematicsJob(getSchematicsJobOptions *GetSchematicsJobOptions) (result *GetActionJobResponse, response *core.DetailedResponse, err error) {
	return projects.GetSchematicsJobWithContext(context.Background(), getSchematicsJobOptions)
}

// GetSchematicsJobWithContext is an alternate form of the GetSchematicsJob method which supports a Context parameter
func (projects *ProjectsV1) GetSchematicsJobWithContext(ctx context.Context, getSchematicsJobOptions *GetSchematicsJobOptions) (result *GetActionJobResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/job/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSchematicsJobOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetSchematicsJob")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetCostEstimate : Fetch the cost estimate for a given configuraton
// Fetch the cost estimate for a given configuraton.
func (projects *ProjectsV1) GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions) (result *GetCostEstimateResponse, response *core.DetailedResponse, err error) {
	return projects.GetCostEstimateWithContext(context.Background(), getCostEstimateOptions)
}

// GetCostEstimateWithContext is an alternate form of the GetCostEstimate method which supports a Context parameter
func (projects *ProjectsV1) GetCostEstimateWithContext(ctx context.Context, getCostEstimateOptions *GetCostEstimateOptions) (result *GetCostEstimateResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/cost_estimate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCostEstimateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetCostEstimate")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// PostNotification : Add some notifications
// Creates a notification event to be stored on the project definition.
func (projects *ProjectsV1) PostNotification(postNotificationOptions *PostNotificationOptions) (result *PostNotificationsResponse, response *core.DetailedResponse, err error) {
	return projects.PostNotificationWithContext(context.Background(), postNotificationOptions)
}

// PostNotificationWithContext is an alternate form of the PostNotification method which supports a Context parameter
func (projects *ProjectsV1) PostNotificationWithContext(ctx context.Context, postNotificationOptions *PostNotificationOptions) (result *PostNotificationsResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "PostNotification")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetNotifications : Get events by project id
// Get all the notification events from a specific project id.
func (projects *ProjectsV1) GetNotifications(getNotificationsOptions *GetNotificationsOptions) (result *GetNotificationsResponse, response *core.DetailedResponse, err error) {
	return projects.GetNotificationsWithContext(context.Background(), getNotificationsOptions)
}

// GetNotificationsWithContext is an alternate form of the GetNotifications method which supports a Context parameter
func (projects *ProjectsV1) GetNotificationsWithContext(ctx context.Context, getNotificationsOptions *GetNotificationsOptions) (result *GetNotificationsResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = projects.Service.Request(request, &rawResponse)
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

// DeleteNotification : Delete a notification from a project
// Delete a project notification.
// - in: query
//   name: notification_id
//   description: The id of the project, which uniquely identifies it.
//   required: true
//   schema:
//     $ref: "#/components/schemas/Identifier".
func (projects *ProjectsV1) DeleteNotification(deleteNotificationOptions *DeleteNotificationOptions) (response *core.DetailedResponse, err error) {
	return projects.DeleteNotificationWithContext(context.Background(), deleteNotificationOptions)
}

// DeleteNotificationWithContext is an alternate form of the DeleteNotification method which supports a Context parameter
func (projects *ProjectsV1) DeleteNotificationWithContext(ctx context.Context, deleteNotificationOptions *DeleteNotificationOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/event`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeleteNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// ReceivePulsarCatalogEvents : Webhook for catalog events
// This is a webhook for pulsar catalog events.
func (projects *ProjectsV1) ReceivePulsarCatalogEvents(receivePulsarCatalogEventsOptions *ReceivePulsarCatalogEventsOptions) (response *core.DetailedResponse, err error) {
	return projects.ReceivePulsarCatalogEventsWithContext(context.Background(), receivePulsarCatalogEventsOptions)
}

// ReceivePulsarCatalogEventsWithContext is an alternate form of the ReceivePulsarCatalogEvents method which supports a Context parameter
func (projects *ProjectsV1) ReceivePulsarCatalogEventsWithContext(ctx context.Context, receivePulsarCatalogEventsOptions *ReceivePulsarCatalogEventsOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/pulsar/catalog_events`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range receivePulsarCatalogEventsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ReceivePulsarCatalogEvents")
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

	response, err = projects.Service.Request(request, nil)

	return
}

// GetHealth : Get service health information
// Get service health information.
func (projects *ProjectsV1) GetHealth(getHealthOptions *GetHealthOptions) (result *Health, response *core.DetailedResponse, err error) {
	return projects.GetHealthWithContext(context.Background(), getHealthOptions)
}

// GetHealthWithContext is an alternate form of the GetHealth method which supports a Context parameter
func (projects *ProjectsV1) GetHealthWithContext(ctx context.Context, getHealthOptions *GetHealthOptions) (result *Health, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getHealthOptions, "getHealthOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/health`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getHealthOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetHealth")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// ReplaceServiceInstance : create a new service instance
// Provision a new service instance Create a service instance. When the service broker receives a provision request from
// the IBM Cloud platform, it MUST take whatever action is necessary to create a new resource. When a user creates a
// service instance from the IBM Cloud console or the IBM Cloud CLI, the IBM Cloud platform validates that the user has
// permission to create the service instance using IBM Cloud IAM. After this validation occurs, your service broker's
// provision endpoint (PUT /v2/resource_instances/:instance_id) will be invoked. When provisioning occurs, the IBM Cloud
// platform provides the following values:
// - The IBM Cloud context is included in the context variable - The X-Broker-API-Originating-Identity will have the IBM
// IAM ID of the user that initiated the request - The parameters section will include the requested location (and
// additional parameters required by your service).
func (projects *ProjectsV1) ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *CreateResult, response *core.DetailedResponse, err error) {
	return projects.ReplaceServiceInstanceWithContext(context.Background(), replaceServiceInstanceOptions)
}

// ReplaceServiceInstanceWithContext is an alternate form of the ReplaceServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) ReplaceServiceInstanceWithContext(ctx context.Context, replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *CreateResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ReplaceServiceInstance")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// DeleteServiceInstance : delete a project service instance
// Delete (deprovision) a project service instance by GUID. When a service broker receives a deprovision request from
// the IBM Cloud platform, it MUST delete any resources it created during the provision. Usually this means that all
// resources are immediately reclaimed for future provisions.
func (projects *ProjectsV1) DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *DeleteResult, response *core.DetailedResponse, err error) {
	return projects.DeleteServiceInstanceWithContext(context.Background(), deleteServiceInstanceOptions)
}

// DeleteServiceInstanceWithContext is an alternate form of the DeleteServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) DeleteServiceInstanceWithContext(ctx context.Context, deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *DeleteResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeleteServiceInstance")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// UpdateServiceInstance : allow to change plans and service parameters in a provisioned service instance
// Update plans and service parameters in a provisioned service instance.
func (projects *ProjectsV1) UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *UpdateResult, response *core.DetailedResponse, err error) {
	return projects.UpdateServiceInstanceWithContext(context.Background(), updateServiceInstanceOptions)
}

// UpdateServiceInstanceWithContext is an alternate form of the UpdateServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) UpdateServiceInstanceWithContext(ctx context.Context, updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *UpdateResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v2/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateServiceInstanceOptions.XBrokerApiVersion != nil {
		builder.AddHeader("X-Broker-Api-Version", fmt.Sprint(*updateServiceInstanceOptions.XBrokerApiVersion))
	}
	if updateServiceInstanceOptions.XBrokerApiOriginatingIdentity != nil {
		builder.AddHeader("X-Broker-Api-Originating-Identity", fmt.Sprint(*updateServiceInstanceOptions.XBrokerApiOriginatingIdentity))
	}

	if updateServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*updateServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if updateServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = updateServiceInstanceOptions.ServiceID
	}
	if updateServiceInstanceOptions.Context != nil {
		body["context"] = updateServiceInstanceOptions.Context
	}
	if updateServiceInstanceOptions.Parameters != nil {
		body["parameters"] = updateServiceInstanceOptions.Parameters
	}
	if updateServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = updateServiceInstanceOptions.PlanID
	}
	if updateServiceInstanceOptions.PreviousValues != nil {
		body["previous_values"] = updateServiceInstanceOptions.PreviousValues
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetLastOperation : Get last_operation for instance by GUID (for asynchronous provision calls)
// Retrieve previous operation for service instance by GUID.
func (projects *ProjectsV1) GetLastOperation(getLastOperationOptions *GetLastOperationOptions) (result *GetLastOperationResult, response *core.DetailedResponse, err error) {
	return projects.GetLastOperationWithContext(context.Background(), getLastOperationOptions)
}

// GetLastOperationWithContext is an alternate form of the GetLastOperation method which supports a Context parameter
func (projects *ProjectsV1) GetLastOperationWithContext(ctx context.Context, getLastOperationOptions *GetLastOperationOptions) (result *GetLastOperationResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v2/service_instances/{instance_id}/last_operation`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLastOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetLastOperation")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// ReplaceServiceInstanceState : Update (disable or enable) the state of a provisioned service instance
// Update the state of a provisioned service instance.
func (projects *ProjectsV1) ReplaceServiceInstanceState(replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	return projects.ReplaceServiceInstanceStateWithContext(context.Background(), replaceServiceInstanceStateOptions)
}

// ReplaceServiceInstanceStateWithContext is an alternate form of the ReplaceServiceInstanceState method which supports a Context parameter
func (projects *ProjectsV1) ReplaceServiceInstanceStateWithContext(ctx context.Context, replaceServiceInstanceStateOptions *ReplaceServiceInstanceStateOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ReplaceServiceInstanceState")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetServiceInstance : Get the current state information associated with the service instance
// Retrieve current state for the specified the service instance.
func (projects *ProjectsV1) GetServiceInstance(getServiceInstanceOptions *GetServiceInstanceOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
	return projects.GetServiceInstanceWithContext(context.Background(), getServiceInstanceOptions)
}

// GetServiceInstanceWithContext is an alternate form of the GetServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) GetServiceInstanceWithContext(ctx context.Context, getServiceInstanceOptions *GetServiceInstanceOptions) (result *BrokerResult, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/bluemix_v1/service_instances/{instance_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetServiceInstance")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetCatalog : Get the catalog metadata stored within the broker
// Fetch catalog metadata.
func (projects *ProjectsV1) GetCatalog(getCatalogOptions *GetCatalogOptions) (result *CatalogResponse, response *core.DetailedResponse, err error) {
	return projects.GetCatalogWithContext(context.Background(), getCatalogOptions)
}

// GetCatalogWithContext is an alternate form of the GetCatalog method which supports a Context parameter
func (projects *ProjectsV1) GetCatalogWithContext(ctx context.Context, getCatalogOptions *GetCatalogOptions) (result *CatalogResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCatalogOptions, "getCatalogOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v2/catalog`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetCatalog")
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// PostEventNotificationsIntegration : connect to a event notifications instance
// connects a project instance to an event notifications instance.
func (projects *ProjectsV1) PostEventNotificationsIntegration(postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions) (result *PostEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	return projects.PostEventNotificationsIntegrationWithContext(context.Background(), postEventNotificationsIntegrationOptions)
}

// PostEventNotificationsIntegrationWithContext is an alternate form of the PostEventNotificationsIntegration method which supports a Context parameter
func (projects *ProjectsV1) PostEventNotificationsIntegrationWithContext(ctx context.Context, postEventNotificationsIntegrationOptions *PostEventNotificationsIntegrationOptions) (result *PostEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "PostEventNotificationsIntegration")
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
	if postEventNotificationsIntegrationOptions.Name != nil {
		body["name"] = postEventNotificationsIntegrationOptions.Name
	}
	if postEventNotificationsIntegrationOptions.Enabled != nil {
		body["enabled"] = postEventNotificationsIntegrationOptions.Enabled
	}
	if postEventNotificationsIntegrationOptions.Source != nil {
		body["source"] = postEventNotificationsIntegrationOptions.Source
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
	response, err = projects.Service.Request(request, &rawResponse)
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

// GetEventNotificationsIntegration : Get event notification source details by project id
// gets the source details of the project from the connect event notifications instance.
func (projects *ProjectsV1) GetEventNotificationsIntegration(getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions) (result *GetEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
	return projects.GetEventNotificationsIntegrationWithContext(context.Background(), getEventNotificationsIntegrationOptions)
}

// GetEventNotificationsIntegrationWithContext is an alternate form of the GetEventNotificationsIntegration method which supports a Context parameter
func (projects *ProjectsV1) GetEventNotificationsIntegrationWithContext(ctx context.Context, getEventNotificationsIntegrationOptions *GetEventNotificationsIntegrationOptions) (result *GetEventNotificationsIntegrationResponse, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetEventNotificationsIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = projects.Service.Request(request, &rawResponse)
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

// DeleteEventNotificationsIntegration : Delete a event notifications connection to this project
// Removes the event notifications integration if the project was onboarded to one.
func (projects *ProjectsV1) DeleteEventNotificationsIntegration(deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions) (response *core.DetailedResponse, err error) {
	return projects.DeleteEventNotificationsIntegrationWithContext(context.Background(), deleteEventNotificationsIntegrationOptions)
}

// DeleteEventNotificationsIntegrationWithContext is an alternate form of the DeleteEventNotificationsIntegration method which supports a Context parameter
func (projects *ProjectsV1) DeleteEventNotificationsIntegrationWithContext(ctx context.Context, deleteEventNotificationsIntegrationOptions *DeleteEventNotificationsIntegrationOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteEventNotificationsIntegrationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeleteEventNotificationsIntegration")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// PostEventNotification : send notification to event notifications instance
// sends notification to event notifications instance.
func (projects *ProjectsV1) PostEventNotification(postEventNotificationOptions *PostEventNotificationOptions) (result *PostEventNotificationRequest, response *core.DetailedResponse, err error) {
	return projects.PostEventNotificationWithContext(context.Background(), postEventNotificationOptions)
}

// PostEventNotificationWithContext is an alternate form of the PostEventNotification method which supports a Context parameter
func (projects *ProjectsV1) PostEventNotificationWithContext(ctx context.Context, postEventNotificationOptions *PostEventNotificationOptions) (result *PostEventNotificationRequest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postEventNotificationOptions, "postEventNotificationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postEventNotificationOptions, "postEventNotificationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *postEventNotificationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/integrations/event_notifications/notifications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postEventNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "PostEventNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postEventNotificationOptions.NewID != nil {
		body["id"] = postEventNotificationOptions.NewID
	}
	if postEventNotificationOptions.NewSource != nil {
		body["source"] = postEventNotificationOptions.NewSource
	}
	if postEventNotificationOptions.NewDatacontenttype != nil {
		body["datacontenttype"] = postEventNotificationOptions.NewDatacontenttype
	}
	if postEventNotificationOptions.NewIbmendefaultlong != nil {
		body["ibmendefaultlong"] = postEventNotificationOptions.NewIbmendefaultlong
	}
	if postEventNotificationOptions.NewIbmendefaultshort != nil {
		body["ibmendefaultshort"] = postEventNotificationOptions.NewIbmendefaultshort
	}
	if postEventNotificationOptions.NewIbmensourceid != nil {
		body["ibmensourceid"] = postEventNotificationOptions.NewIbmensourceid
	}
	if postEventNotificationOptions.NewSpecversion != nil {
		body["specversion"] = postEventNotificationOptions.NewSpecversion
	}
	if postEventNotificationOptions.NewType != nil {
		body["type"] = postEventNotificationOptions.NewType
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
	response, err = projects.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPostEventNotificationRequest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BrokerResult : Result of Get instance status call.
type BrokerResult struct {
	// Indicates whether the service instance is active and is meaningful if enabled is true. The default value is true if
	// not specified.
	Active *string `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *string `json:"enabled,omitempty"`

	// Indicates when the service instance was last accessed/modified/etc., and is meaningful if enabled is true AND active
	// is false. Represented as milliseconds since the epoch, but does not need to be accurate to the second/hour.
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

// CatalogResponse : CatalogResponse struct
type CatalogResponse struct {
	Services []CatalogResponseServicesItem `json:"services,omitempty"`
}

// UnmarshalCatalogResponse unmarshals an instance of CatalogResponse from the specified map of raw messages.
func UnmarshalCatalogResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponse)
	err = core.UnmarshalModel(m, "services", &obj.Services, UnmarshalCatalogResponseServicesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServicesItem : CatalogResponseServicesItem struct
type CatalogResponseServicesItem struct {
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

	Metadata *CatalogResponseServicesItemMetadata `json:"metadata,omitempty"`

	// The service name is not your display name. Your service name must follow the follow these rules: It must be all
	// lowercase. It can't include spaces but may include hyphens (-). It must be less than 32 characters. Your service
	// name should include your company name. If your company has more then one offering your service name should include
	// both company and offering as part of the name. For example, the Compose company has offerings for Redis and
	// Elasticsearch. Sample service names on IBM Cloud for these offerings would be compose-redis and
	// compose-elasticsearch. Each of these service names have associated display names that are shown in the IBM Cloud
	// catalog: Compose Redis and Compose Elasticsearch. Another company (e.g. FastJetMail) may only have the single
	// JetMail offering, in which case the service name should be fastjetmail. Recommended: If you define your service in
	// RMC, you can export a catalog.json that will include the service name you defined within the RMC.
	Name *string `json:"name,omitempty"`

	// The Default is false. This specifices whether or not you support plan changes for provisioned instances. If your
	// offering supports multiple plans, and you want users to be able to change plans for a provisioned instance, you will
	// need to enable the ability for users to update their service instance by using /v2/service_instances/{instance_id}
	// PATCH.
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`

	Tags []string `json:"tags,omitempty"`

	// A list of plans for this service that must contain at least one plan.
	Plans []CatalogResponseServicesItemPlansItem `json:"plans,omitempty"`
}

// UnmarshalCatalogResponseServicesItem unmarshals an instance of CatalogResponseServicesItem from the specified map of raw messages.
func UnmarshalCatalogResponseServicesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesItem)
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
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalCatalogResponseServicesItemMetadata)
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
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalCatalogResponseServicesItemPlansItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogResponseServicesItemMetadata : CatalogResponseServicesItemMetadata struct
type CatalogResponseServicesItemMetadata struct {
	DisplayName *string `json:"display_name,omitempty"`

	DocumentationURL *string `json:"documentation_url,omitempty"`

	ImageURL *string `json:"image_url,omitempty"`

	InstructionsURL *string `json:"instructions_url,omitempty"`

	LongDescription *string `json:"long_description,omitempty"`

	ProviderDisplayName *string `json:"provider_display_name,omitempty"`

	SupportURL *string `json:"support_url,omitempty"`

	TermsURL *string `json:"terms_url,omitempty"`
}

// UnmarshalCatalogResponseServicesItemMetadata unmarshals an instance of CatalogResponseServicesItemMetadata from the specified map of raw messages.
func UnmarshalCatalogResponseServicesItemMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesItemMetadata)
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

// CatalogResponseServicesItemPlansItem : CatalogResponseServicesItemPlansItem struct
type CatalogResponseServicesItemPlansItem struct {
	Description *string `json:"description,omitempty"`

	Free *bool `json:"free,omitempty"`

	ID *string `json:"id,omitempty"`

	Metadata *CatalogResponseServicesItemPlansItemMetadata `json:"metadata,omitempty"`

	Name *string `json:"name,omitempty"`
}

// UnmarshalCatalogResponseServicesItemPlansItem unmarshals an instance of CatalogResponseServicesItemPlansItem from the specified map of raw messages.
func UnmarshalCatalogResponseServicesItemPlansItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesItemPlansItem)
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
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalCatalogResponseServicesItemPlansItemMetadata)
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

// CatalogResponseServicesItemPlansItemMetadata : CatalogResponseServicesItemPlansItemMetadata struct
type CatalogResponseServicesItemPlansItemMetadata struct {
	Bullets []string `json:"bullets,omitempty"`

	DisplayName *string `json:"display_name,omitempty"`
}

// UnmarshalCatalogResponseServicesItemPlansItemMetadata unmarshals an instance of CatalogResponseServicesItemPlansItemMetadata from the specified map of raw messages.
func UnmarshalCatalogResponseServicesItemPlansItemMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogResponseServicesItemPlansItemMetadata)
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config to trigger a validation check.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The IAM refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token,omitempty"`

	// The version of the config that the validation check should trigger against.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckConfigOptions : Instantiate CheckConfigOptions
func (*ProjectsV1) NewCheckConfigOptions(id string, configID string) *CheckConfigOptions {
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

// SetHeaders : Allow user to set Headers
func (options *CheckConfigOptions) SetHeaders(param map[string]string) *CheckConfigOptions {
	options.Headers = param
	return options
}

// ConfigSettingItems : ConfigSettingItems struct
type ConfigSettingItems struct {
	// The name of the config setting.
	Name *string `json:"name" validate:"required"`

	// The value of a the config setting.
	Value *string `json:"value" validate:"required"`
}

// NewConfigSettingItems : Instantiate ConfigSettingItems (Generic Model Constructor)
func (*ProjectsV1) NewConfigSettingItems(name string, value string) (_model *ConfigSettingItems, err error) {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"-" validate:"required,ne="`

	// The config name.
	NewName *string `json:"name" validate:"required"`

	// The location id of a Project Config Manual Property.
	NewLocatorID *string `json:"locator_id" validate:"required"`

	// The unique id of a project.
	NewID *string `json:"id,omitempty"`

	// collection of config labels.
	NewLabels []string `json:"labels,omitempty"`

	// A project config description.
	NewDescription *string `json:"description,omitempty"`

	// The inputs of a Schematics Template Property.
	NewInput []InputVariableInput `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	NewSetting []ConfigSettingItems `json:"setting,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigOptions : Instantiate CreateConfigOptions
func (*ProjectsV1) NewCreateConfigOptions(id string, newName string, newLocatorID string) *CreateConfigOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The action to the draft.
	Action *string `json:"action" validate:"required,ne="`

	// Notes on the project draft action.
	Comment *string `json:"comment,omitempty"`

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
func (*ProjectsV1) NewCreateDraftActionOptions(id string, configID string, action string) *CreateDraftActionOptions {
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

	Configs []ProjectConfigInput `json:"configs,omitempty"`

	// Group name of the customized collection of resources.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// Data center locations for resource deployment.
	Location *string `json:"location,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*ProjectsV1) NewCreateProjectOptions(name string) *CreateProjectOptions {
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

// CreateResult : Result of Provision call.
type CreateResult struct {
	// The URL of a web-based management user interface for the service instance. The URL MUST contain enough information
	// for the dashboard to identify the resource being accessed.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// For asynchronous responses, service brokers MAY return an identifier representing the operation. The value of this
	// field MUST be provided by the platform with requests to the last_operation endpoint in a URL encoded query
	// parameter. If present, MUST be a non-empty string.
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

// CumulativeNeedsAttentionView : CumulativeNeedsAttentionView struct
type CumulativeNeedsAttentionView struct {
	// The event name.
	Event *string `json:"event,omitempty"`

	// The unique id of a project.
	EventID *string `json:"event_id,omitempty"`

	// The unique id of a project.
	ConfigID *string `json:"config_id,omitempty"`

	// The version number of the config.
	ConfigVersion *int64 `json:"config_version,omitempty"`
}

// UnmarshalCumulativeNeedsAttentionView unmarshals an instance of CumulativeNeedsAttentionView from the specified map of raw messages.
func UnmarshalCumulativeNeedsAttentionView(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CumulativeNeedsAttentionView)
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteConfigOptions : Instantiate DeleteConfigOptions
func (*ProjectsV1) NewDeleteConfigOptions(id string, configID string) *DeleteConfigOptions {
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

// SetHeaders : Allow user to set Headers
func (options *DeleteConfigOptions) SetHeaders(param map[string]string) *DeleteConfigOptions {
	options.Headers = param
	return options
}

// DeleteEventNotificationsIntegrationOptions : The DeleteEventNotificationsIntegration options.
type DeleteEventNotificationsIntegrationOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteEventNotificationsIntegrationOptions : Instantiate DeleteEventNotificationsIntegrationOptions
func (*ProjectsV1) NewDeleteEventNotificationsIntegrationOptions(id string) *DeleteEventNotificationsIntegrationOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteNotificationOptions : Instantiate DeleteNotificationOptions
func (*ProjectsV1) NewDeleteNotificationOptions(id string) *DeleteNotificationOptions {
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

// DeleteProjectConfigResponse : Delete config response.
type DeleteProjectConfigResponse struct {
	// The unique id of a project.
	ID *string `json:"id,omitempty"`

	// The name of the config being deleted.
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteProjectOptions : Instantiate DeleteProjectOptions
func (*ProjectsV1) NewDeleteProjectOptions(id string) *DeleteProjectOptions {
	return &DeleteProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteProjectOptions) SetID(id string) *DeleteProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteProjectOptions) SetHeaders(param map[string]string) *DeleteProjectOptions {
	options.Headers = param
	return options
}

// DeleteResult : Result of deprovisioning service instance.
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
	// broker. This value should be a GUID. MUST be a non-empty string.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. MUST be a non-empty
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
func (*ProjectsV1) NewDeleteServiceInstanceOptions(instanceID string, planID string, serviceID string) *DeleteServiceInstanceOptions {
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

// GetActionJobResponse : The response of a fetching an Action Job.
type GetActionJobResponse struct {
	// The unique id of a project.
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
	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogOptions : Instantiate GetCatalogOptions
func (*ProjectsV1) NewGetCatalogOptions() *GetCatalogOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigDiffOptions : Instantiate GetConfigDiffOptions
func (*ProjectsV1) NewGetConfigDiffOptions(id string, configID string) *GetConfigDiffOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The version of the config to return.
	Version *string `json:"version,omitempty"`

	// The flag to tell if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetConfigOptions : Instantiate GetConfigOptions
func (*ProjectsV1) NewGetConfigOptions(id string, configID string) *GetConfigOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config of the cost estimate to fetch.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The version of the config that the cost estimate will be fetched.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCostEstimateOptions : Instantiate GetCostEstimateOptions
func (*ProjectsV1) NewGetCostEstimateOptions(id string, configID string) *GetCostEstimateOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEventNotificationsIntegrationOptions : Instantiate GetEventNotificationsIntegrationOptions
func (*ProjectsV1) NewGetEventNotificationsIntegrationOptions(id string) *GetEventNotificationsIntegrationOptions {
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

// GetEventNotificationsIntegrationResponse : the resulting response of getting the source details of the event notifications integration.
type GetEventNotificationsIntegrationResponse struct {
	// description of the instance of event.
	Description *string `json:"description,omitempty"`

	// name of the instance of event.
	Name *string `json:"name,omitempty"`

	// status of instance of event.
	Enabled *bool `json:"enabled,omitempty"`

	ID *string `json:"id,omitempty"`

	// type of the instance of event.
	Type *string `json:"type,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	// topic count of the instance of event.
	TopicCount *int64 `json:"topic_count,omitempty"`

	// topic names of the instance of event.
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
func (*ProjectsV1) NewGetHealthOptions() *GetHealthOptions {
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
func (*ProjectsV1) NewGetLastOperationOptions(instanceID string) *GetLastOperationOptions {
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

// GetLastOperationResult : Result of get_last_operation call.
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNotificationsOptions : Instantiate GetNotificationsOptions
func (*ProjectsV1) NewGetNotificationsOptions(id string) *GetNotificationsOptions {
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

// GetNotificationsResponse : GetNotificationsResponse struct
type GetNotificationsResponse struct {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Only return with the active configs, no drafts.
	ExcludeConfigs *bool `json:"exclude_configs,omitempty"`

	// The flag to tell if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectOptions : Instantiate GetProjectOptions
func (*ProjectsV1) NewGetProjectOptions(id string) *GetProjectOptions {
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

// GetProjectResponse : The project returned in response body.
type GetProjectResponse struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`

	// The unique id of a project.
	ID *string `json:"id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	Crn *string `json:"crn,omitempty"`

	Configs []ProjectConfig `json:"configs,omitempty"`

	Metadata *GetProjectResponseMetadata `json:"metadata,omitempty"`
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
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalGetProjectResponseMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetProjectResponseMetadata : GetProjectResponseMetadata struct
type GetProjectResponseMetadata struct {
	// An IBM Cloud Resource Name, which uniquely identify a resource.
	Crn *string `json:"crn,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	CumulativeNeedsAttentionView *CumulativeNeedsAttentionView `json:"cumulative_needs_attention_view,omitempty"`

	// True to indicate the fetch of needs attention items failed.
	CumulativeNeedsAttentionViewErr *string `json:"cumulative_needs_attention_view_err,omitempty"`

	// The location of where the project created.
	Location *string `json:"location,omitempty"`

	// The resource group of where the project created.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// The project status value.
	State *string `json:"state,omitempty"`

	// The crn of the event notifications instance if one is connected to this project.
	EventNotificationsCrn *string `json:"event_notifications_crn,omitempty"`
}

// UnmarshalGetProjectResponseMetadata unmarshals an instance of GetProjectResponseMetadata from the specified map of raw messages.
func UnmarshalGetProjectResponseMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetProjectResponseMetadata)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cumulative_needs_attention_view", &obj.CumulativeNeedsAttentionView, UnmarshalCumulativeNeedsAttentionView)
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

// GetSchematicsJobOptions : The GetSchematicsJob options.
type GetSchematicsJobOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config that triggered the action.
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
func (*ProjectsV1) NewGetSchematicsJobOptions(id string, configID string, action string) *GetSchematicsJobOptions {
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

	// Broker Api Version.
	XBrokerApiVersion *string `json:"X-Broker-Api-Version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceInstanceOptions : Instantiate GetServiceInstanceOptions
func (*ProjectsV1) NewGetServiceInstanceOptions(instanceID string) *GetServiceInstanceOptions {
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
func (*ProjectsV1) NewInputVariableInput(name string) (_model *InputVariableInput, err error) {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config to install.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallConfigOptions : Instantiate InstallConfigOptions
func (*ProjectsV1) NewInstallConfigOptions(id string, configID string) *InstallConfigOptions {
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

// SetHeaders : Allow user to set Headers
func (options *InstallConfigOptions) SetHeaders(param map[string]string) *InstallConfigOptions {
	options.Headers = param
	return options
}

// ListConfigsOptions : The ListConfigs options.
type ListConfigsOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The version of configs to return.
	Version *string `json:"version,omitempty"`

	// The flag to tell if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListConfigsOptions.Version property.
// The version of configs to return.
const (
	ListConfigsOptions_Version_Active = "active"
	ListConfigsOptions_Version_Draft = "draft"
	ListConfigsOptions_Version_Mixed = "mixed"
)

// NewListConfigsOptions : Instantiate ListConfigsOptions
func (*ProjectsV1) NewListConfigsOptions(id string) *ListConfigsOptions {
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

	// The flag to tell if full metadata should be returned.
	Complete *bool `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*ProjectsV1) NewListProjectsOptions() *ListProjectsOptions {
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
	// Type of event.
	Event *string `json:"event" validate:"required"`

	// The unique id of a project.
	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewNotificationEvent : Instantiate NotificationEvent (Generic Model Constructor)
func (*ProjectsV1) NewNotificationEvent(event string, target string) (_model *NotificationEvent, err error) {
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
	// Type of event.
	Event *string `json:"event" validate:"required"`

	// The unique id of a project.
	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`

	// The unique id of a project.
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
	// Type of event.
	Event *string `json:"event" validate:"required"`

	// The unique id of a project.
	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data map[string]interface{} `json:"data,omitempty"`

	// The unique id of a project.
	ID *string `json:"id" validate:"required"`

	// whether or not the event successfully posted.
	Status *string `json:"status,omitempty"`

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

	// A descriptive of the output value.
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

// PaginationLink : PaginationLink struct
type PaginationLink struct {
	// The url of the PR, which uniquely identifies it.
	Href *string `json:"href" validate:"required"`

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

// PostEventNotificationOptions : The PostEventNotification options.
type PostEventNotificationOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"-" validate:"required,ne="`

	NewID *string `json:"id" validate:"required"`

	// source of instance of event.
	NewSource *string `json:"source" validate:"required"`

	// data content type of the instance of event.
	NewDatacontenttype *string `json:"datacontenttype,omitempty"`

	// ibm default long message of the instance of event.
	NewIbmendefaultlong *string `json:"ibmendefaultlong,omitempty"`

	// ibm default short message of the instance of event.
	NewIbmendefaultshort *string `json:"ibmendefaultshort,omitempty"`

	// ibm source id of the instance of event.
	NewIbmensourceid *string `json:"ibmensourceid,omitempty"`

	// spec version of instance of event.
	NewSpecversion *string `json:"specversion,omitempty"`

	// type of instance of event.
	NewType *string `json:"type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostEventNotificationOptions : Instantiate PostEventNotificationOptions
func (*ProjectsV1) NewPostEventNotificationOptions(id string, newID string, newSource string) *PostEventNotificationOptions {
	return &PostEventNotificationOptions{
		ID: core.StringPtr(id),
		NewID: core.StringPtr(newID),
		NewSource: core.StringPtr(newSource),
	}
}

// SetID : Allow user to set ID
func (_options *PostEventNotificationOptions) SetID(id string) *PostEventNotificationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetNewID : Allow user to set NewID
func (_options *PostEventNotificationOptions) SetNewID(newID string) *PostEventNotificationOptions {
	_options.NewID = core.StringPtr(newID)
	return _options
}

// SetNewSource : Allow user to set NewSource
func (_options *PostEventNotificationOptions) SetNewSource(newSource string) *PostEventNotificationOptions {
	_options.NewSource = core.StringPtr(newSource)
	return _options
}

// SetNewDatacontenttype : Allow user to set NewDatacontenttype
func (_options *PostEventNotificationOptions) SetNewDatacontenttype(newDatacontenttype string) *PostEventNotificationOptions {
	_options.NewDatacontenttype = core.StringPtr(newDatacontenttype)
	return _options
}

// SetNewIbmendefaultlong : Allow user to set NewIbmendefaultlong
func (_options *PostEventNotificationOptions) SetNewIbmendefaultlong(newIbmendefaultlong string) *PostEventNotificationOptions {
	_options.NewIbmendefaultlong = core.StringPtr(newIbmendefaultlong)
	return _options
}

// SetNewIbmendefaultshort : Allow user to set NewIbmendefaultshort
func (_options *PostEventNotificationOptions) SetNewIbmendefaultshort(newIbmendefaultshort string) *PostEventNotificationOptions {
	_options.NewIbmendefaultshort = core.StringPtr(newIbmendefaultshort)
	return _options
}

// SetNewIbmensourceid : Allow user to set NewIbmensourceid
func (_options *PostEventNotificationOptions) SetNewIbmensourceid(newIbmensourceid string) *PostEventNotificationOptions {
	_options.NewIbmensourceid = core.StringPtr(newIbmensourceid)
	return _options
}

// SetNewSpecversion : Allow user to set NewSpecversion
func (_options *PostEventNotificationOptions) SetNewSpecversion(newSpecversion string) *PostEventNotificationOptions {
	_options.NewSpecversion = core.StringPtr(newSpecversion)
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *PostEventNotificationOptions) SetNewType(newType string) *PostEventNotificationOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostEventNotificationOptions) SetHeaders(param map[string]string) *PostEventNotificationOptions {
	options.Headers = param
	return options
}

// PostEventNotificationRequest : the request to post an event notification to the event notifications API.
type PostEventNotificationRequest struct {
	// data content type of the instance of event.
	Datacontenttype *string `json:"datacontenttype,omitempty"`

	// ibm default long message of the instance of event.
	Ibmendefaultlong *string `json:"ibmendefaultlong,omitempty"`

	// ibm default short message of the instance of event.
	Ibmendefaultshort *string `json:"ibmendefaultshort,omitempty"`

	// ibm source id of the instance of event.
	Ibmensourceid *string `json:"ibmensourceid,omitempty"`

	ID *string `json:"id" validate:"required"`

	// source of instance of event.
	Source *string `json:"source" validate:"required"`

	// spec version of instance of event.
	Specversion *string `json:"specversion,omitempty"`

	// type of instance of event.
	Type *string `json:"type,omitempty"`
}

// NewPostEventNotificationRequest : Instantiate PostEventNotificationRequest (Generic Model Constructor)
func (*ProjectsV1) NewPostEventNotificationRequest(id string, source string) (_model *PostEventNotificationRequest, err error) {
	_model = &PostEventNotificationRequest{
		ID: core.StringPtr(id),
		Source: core.StringPtr(source),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPostEventNotificationRequest unmarshals an instance of PostEventNotificationRequest from the specified map of raw messages.
func UnmarshalPostEventNotificationRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PostEventNotificationRequest)
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

// PostEventNotificationsIntegrationOptions : The PostEventNotificationsIntegration options.
type PostEventNotificationsIntegrationOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	InstanceCrn *string `json:"instance_crn" validate:"required"`

	// description of the instance of event.
	Description *string `json:"description,omitempty"`

	// name of the instance of event.
	Name *string `json:"name,omitempty"`

	// status of instance of event.
	Enabled *bool `json:"enabled,omitempty"`

	// source of instance of event.
	Source *string `json:"source,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostEventNotificationsIntegrationOptions : Instantiate PostEventNotificationsIntegrationOptions
func (*ProjectsV1) NewPostEventNotificationsIntegrationOptions(id string, instanceCrn string) *PostEventNotificationsIntegrationOptions {
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

// SetName : Allow user to set Name
func (_options *PostEventNotificationsIntegrationOptions) SetName(name string) *PostEventNotificationsIntegrationOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *PostEventNotificationsIntegrationOptions) SetEnabled(enabled bool) *PostEventNotificationsIntegrationOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetSource : Allow user to set Source
func (_options *PostEventNotificationsIntegrationOptions) SetSource(source string) *PostEventNotificationsIntegrationOptions {
	_options.Source = core.StringPtr(source)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PostEventNotificationsIntegrationOptions) SetHeaders(param map[string]string) *PostEventNotificationsIntegrationOptions {
	options.Headers = param
	return options
}

// PostEventNotificationsIntegrationResponse : the resulting response of connecting a project to a event notifications instance.
type PostEventNotificationsIntegrationResponse struct {
	// description of the instance of event.
	Description *string `json:"description,omitempty"`

	// name of the instance of event.
	Name *string `json:"name,omitempty"`

	// status of instance of event.
	Enabled *bool `json:"enabled,omitempty"`

	ID *string `json:"id,omitempty"`

	// type of instance of event.
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	Notifications []NotificationEvent `json:"notifications,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostNotificationOptions : Instantiate PostNotificationOptions
func (*ProjectsV1) NewPostNotificationOptions(id string) *PostNotificationOptions {
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

// ProjectConfig : The Project Config.
type ProjectConfig struct {
	// The unique id of a project.
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name" validate:"required"`

	// collection of config labels.
	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	// The location id of a Project Config Manual Property.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The type of a Project Config Manual Property.
	Type *string `json:"type" validate:"required"`

	// The outputs of a Schematics Template Property.
	Input []InputVariable `json:"input,omitempty"`

	// The outputs of a Schematics Template Property.
	Output []OutputValue `json:"output,omitempty"`

	// Optional setting object we can pass to the cart api.
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

// ProjectConfigDiff : The Project Config diff summary.
type ProjectConfigDiff struct {
	// The additions to configs in the diff summary.
	Added *ProjectConfigDiffAdded `json:"added,omitempty"`

	// The changes to configs in the diff summary.
	Changed *ProjectConfigDiffChanged `json:"changed,omitempty"`

	// The deletions to configs in the diff summary.
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

// ProjectConfigDiffAdded : The additions to configs in the diff summary.
type ProjectConfigDiffAdded struct {
	// collection of additions to configs in the diff summary.
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

// ProjectConfigDiffChanged : The changes to configs in the diff summary.
type ProjectConfigDiffChanged struct {
	// collection of changess to configs in the diff summary.
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

// ProjectConfigDiffRemoved : The deletions to configs in the diff summary.
type ProjectConfigDiffRemoved struct {
	// collection of deletions to configs in the diff summary.
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

// ProjectConfigInput : The input of a Project Config.
type ProjectConfigInput struct {
	// The unique id of a project.
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name" validate:"required"`

	// collection of config labels.
	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	// The location id of a Project Config Manual Property.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The inputs of a Schematics Template Property.
	Input []InputVariableInput `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItems `json:"setting,omitempty"`
}

// NewProjectConfigInput : Instantiate ProjectConfigInput (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigInput(name string, locatorID string) (_model *ProjectConfigInput, err error) {
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

// ProjectConfigList : The Project Config List.
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
	// The unique id of a project.
	ID *string `json:"id,omitempty"`

	// The project name.
	Name *string `json:"name,omitempty"`

	// The project description.
	Description *string `json:"description,omitempty"`

	Metadata *ProjectListItemMetadata `json:"metadata,omitempty"`
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
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalProjectListItemMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListItemMetadata : ProjectListItemMetadata struct
type ProjectListItemMetadata struct {
	// An IBM Cloud Resource Name, which uniquely identify a resource.
	Crn *string `json:"crn,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	CumulativeNeedsAttentionView *CumulativeNeedsAttentionView `json:"cumulative_needs_attention_view,omitempty"`

	// True to indicate the fetch of needs attention items failed.
	CumulativeNeedsAttentionViewErr *string `json:"cumulative_needs_attention_view_err,omitempty"`

	// The location of where the project created.
	Location *string `json:"location,omitempty"`

	// The resource group of where the project created.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// The project status value.
	State *string `json:"state,omitempty"`
}

// UnmarshalProjectListItemMetadata unmarshals an instance of ProjectListItemMetadata from the specified map of raw messages.
func UnmarshalProjectListItemMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListItemMetadata)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cumulative_needs_attention_view", &obj.CumulativeNeedsAttentionView, UnmarshalCumulativeNeedsAttentionView)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListResponseSchema : Projects list.
type ProjectListResponseSchema struct {
	Limit *int64 `json:"limit" validate:"required"`

	// Get the occurrencies of the total Projects.
	TotalCount *int64 `json:"total_count" validate:"required"`

	First *PaginationLink `json:"first" validate:"required"`

	Last *PaginationLink `json:"last,omitempty"`

	Previous *PaginationLink `json:"previous,omitempty"`

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

// ProjectUpdate : ProjectUpdate struct
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

// PulsarEventItem : PulsarEventItem struct
type PulsarEventItem struct {
	// The type of the event that is published and written in dot notation.
	EventType *string `json:"event_type" validate:"required"`

	// The time at which the event occurred written as a date-time string.
	Timestamp *string `json:"timestamp" validate:"required"`

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

// NewPulsarEventItem : Instantiate PulsarEventItem (Generic Model Constructor)
func (*ProjectsV1) NewPulsarEventItem(eventType string, timestamp string, publisher string, accountID string, version string) (_model *PulsarEventItem, err error) {
	_model = &PulsarEventItem{
		EventType: core.StringPtr(eventType),
		Timestamp: core.StringPtr(timestamp),
		Publisher: core.StringPtr(publisher),
		AccountID: core.StringPtr(accountID),
		Version: core.StringPtr(version),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of PulsarEventItem
func (o *PulsarEventItem) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of PulsarEventItem
func (o *PulsarEventItem) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of PulsarEventItem
func (o *PulsarEventItem) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of PulsarEventItem
func (o *PulsarEventItem) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of PulsarEventItem
func (o *PulsarEventItem) MarshalJSON() (buffer []byte, err error) {
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

// UnmarshalPulsarEventItem unmarshals an instance of PulsarEventItem from the specified map of raw messages.
func UnmarshalPulsarEventItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PulsarEventItem)
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
	PulsarCatalogEvents []PulsarEventItem `json:"pulsar_catalog_events" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReceivePulsarCatalogEventsOptions : Instantiate ReceivePulsarCatalogEventsOptions
func (*ProjectsV1) NewReceivePulsarCatalogEventsOptions(pulsarCatalogEvents []PulsarEventItem) *ReceivePulsarCatalogEventsOptions {
	return &ReceivePulsarCatalogEventsOptions{
		PulsarCatalogEvents: pulsarCatalogEvents,
	}
}

// SetPulsarCatalogEvents : Allow user to set PulsarCatalogEvents
func (_options *ReceivePulsarCatalogEventsOptions) SetPulsarCatalogEvents(pulsarCatalogEvents []PulsarEventItem) *ReceivePulsarCatalogEventsOptions {
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

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID and it MUST be a
	// non-empty string.
	ServiceID *string `json:"service_id" validate:"required"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
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
func (*ProjectsV1) NewReplaceServiceInstanceOptions(instanceID string, serviceID string, planID string) *ReplaceServiceInstanceOptions {
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

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. It MUST be a non-empty
	// string.
	Enabled *bool `json:"enabled" validate:"required"`

	// Optional string that shows the user ID that is initiating the call.
	InitiatorID *string `json:"initiator_id,omitempty"`

	// Optional string that states the reason code for the service instance state change. Valid values are
	// IBMCLOUD_ACCT_ACTIVATE, IBMCLOUD_RECLAMATION_RESTORE, or IBMCLOUD_SERVICE_INSTANCE_BELOW_CAP for enable calls;
	// IBMCLOUD_ACCT_SUSPEND, IBMCLOUD_RECLAMATION_SCHEDULE, or IBMCLOUD_SERVICE_INSTANCE_ABOVE_CAP for disable calls; and
	// IBMCLOUD_ADMIN_REQUEST for enable and disable calls.
	ReasonCode map[string]interface{} `json:"reason_code,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
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
func (*ProjectsV1) NewReplaceServiceInstanceStateOptions(instanceID string, enabled bool) *ReplaceServiceInstanceStateOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config to uninstall.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUninstallConfigOptions : Instantiate UninstallConfigOptions
func (*ProjectsV1) NewUninstallConfigOptions(id string, configID string) *UninstallConfigOptions {
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config, which uniquely identifies it.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The change delta of the project config to update.
	ProjectConfig UpdateProjectConfigInputIntf `json:"project_config" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigOptions : Instantiate UpdateConfigOptions
func (*ProjectsV1) NewUpdateConfigOptions(id string, configID string, projectConfig UpdateProjectConfigInputIntf) *UpdateConfigOptions {
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
func (_options *UpdateConfigOptions) SetProjectConfig(projectConfig UpdateProjectConfigInputIntf) *UpdateConfigOptions {
	_options.ProjectConfig = projectConfig
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigOptions) SetHeaders(param map[string]string) *UpdateConfigOptions {
	options.Headers = param
	return options
}

// UpdateProjectConfigInput : The Project Config input.
// Models which "extend" this model:
// - UpdateProjectConfigInputProjectConfigManualProperty
// - UpdateProjectConfigInputSchematicsTemplate
type UpdateProjectConfigInput struct {
	// The config name.
	Name *string `json:"name,omitempty"`

	// The config labels.
	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	// The type of a Project Config Manual Property.
	Type *string `json:"type,omitempty"`

	// The external resource account id in project config.
	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	// The location id of a Project Config Manual Property.
	LocatorID *string `json:"locator_id,omitempty"`

	// The inputs of a Schematics Template Property.
	Input []InputVariableInput `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItems `json:"setting,omitempty"`
}

// Constants associated with the UpdateProjectConfigInput.Type property.
// The type of a Project Config Manual Property.
const (
	UpdateProjectConfigInput_Type_Manual = "manual"
)
func (*UpdateProjectConfigInput) isaUpdateProjectConfigInput() bool {
	return true
}

type UpdateProjectConfigInputIntf interface {
	isaUpdateProjectConfigInput() bool
}

// UnmarshalUpdateProjectConfigInput unmarshals an instance of UpdateProjectConfigInput from the specified map of raw messages.
func UnmarshalUpdateProjectConfigInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateProjectConfigInput)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "external_resources_account", &obj.ExternalResourcesAccount)
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

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The project name.
	Name *string `json:"name,omitempty"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectsV1) NewUpdateProjectOptions(id string) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectOptions) SetID(id string) *UpdateProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateProjectOptions) SetName(name string) *UpdateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateProjectOptions) SetDescription(description string) *UpdateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectOptions) SetHeaders(param map[string]string) *UpdateProjectOptions {
	options.Headers = param
	return options
}

// UpdateResult : Result of deprovisioning service instance.
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

// UpdateServiceInstanceOptions : The UpdateServiceInstance options.
type UpdateServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. It MUST be a non-empty
	// string.
	ServiceID []string `json:"service_id" validate:"required"`

	// Platform specific contextual information under which the service instance is to be provisioned.
	Context []string `json:"context,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker.
	PlanID *string `json:"plan_id,omitempty"`

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

// NewUpdateServiceInstanceOptions : Instantiate UpdateServiceInstanceOptions
func (*ProjectsV1) NewUpdateServiceInstanceOptions(instanceID string, serviceID []string) *UpdateServiceInstanceOptions {
	return &UpdateServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
		ServiceID: serviceID,
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateServiceInstanceOptions) SetInstanceID(instanceID string) *UpdateServiceInstanceOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetServiceID : Allow user to set ServiceID
func (_options *UpdateServiceInstanceOptions) SetServiceID(serviceID []string) *UpdateServiceInstanceOptions {
	_options.ServiceID = serviceID
	return _options
}

// SetContext : Allow user to set Context
func (_options *UpdateServiceInstanceOptions) SetContext(context []string) *UpdateServiceInstanceOptions {
	_options.Context = context
	return _options
}

// SetParameters : Allow user to set Parameters
func (_options *UpdateServiceInstanceOptions) SetParameters(parameters map[string]interface{}) *UpdateServiceInstanceOptions {
	_options.Parameters = parameters
	return _options
}

// SetPlanID : Allow user to set PlanID
func (_options *UpdateServiceInstanceOptions) SetPlanID(planID string) *UpdateServiceInstanceOptions {
	_options.PlanID = core.StringPtr(planID)
	return _options
}

// SetPreviousValues : Allow user to set PreviousValues
func (_options *UpdateServiceInstanceOptions) SetPreviousValues(previousValues []string) *UpdateServiceInstanceOptions {
	_options.PreviousValues = previousValues
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

// UpdateProjectConfigInputSchematicsTemplate : The Schematics Template Property.
// This model "extends" UpdateProjectConfigInput
type UpdateProjectConfigInputSchematicsTemplate struct {
	// The config name.
	Name *string `json:"name,omitempty"`

	// The config labels.
	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	// The location id of a Project Config Manual Property.
	LocatorID *string `json:"locator_id,omitempty"`

	// The inputs of a Schematics Template Property.
	Input []InputVariableInput `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItems `json:"setting,omitempty"`
}

func (*UpdateProjectConfigInputSchematicsTemplate) isaUpdateProjectConfigInput() bool {
	return true
}

// UnmarshalUpdateProjectConfigInputSchematicsTemplate unmarshals an instance of UpdateProjectConfigInputSchematicsTemplate from the specified map of raw messages.
func UnmarshalUpdateProjectConfigInputSchematicsTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateProjectConfigInputSchematicsTemplate)
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

// UpdateProjectConfigInputProjectConfigManualProperty : The Project Config Manual Type.
// This model "extends" UpdateProjectConfigInput
type UpdateProjectConfigInputProjectConfigManualProperty struct {
	// The config name.
	Name *string `json:"name,omitempty"`

	// The config labels.
	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	// The type of a Project Config Manual Property.
	Type *string `json:"type" validate:"required"`

	// The external resource account id in project config.
	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`
}

// Constants associated with the UpdateProjectConfigInputProjectConfigManualProperty.Type property.
// The type of a Project Config Manual Property.
const (
	UpdateProjectConfigInputProjectConfigManualProperty_Type_Manual = "manual"
)

// NewUpdateProjectConfigInputProjectConfigManualProperty : Instantiate UpdateProjectConfigInputProjectConfigManualProperty (Generic Model Constructor)
func (*ProjectsV1) NewUpdateProjectConfigInputProjectConfigManualProperty(typeVar string) (_model *UpdateProjectConfigInputProjectConfigManualProperty, err error) {
	_model = &UpdateProjectConfigInputProjectConfigManualProperty{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*UpdateProjectConfigInputProjectConfigManualProperty) isaUpdateProjectConfigInput() bool {
	return true
}

// UnmarshalUpdateProjectConfigInputProjectConfigManualProperty unmarshals an instance of UpdateProjectConfigInputProjectConfigManualProperty from the specified map of raw messages.
func UnmarshalUpdateProjectConfigInputProjectConfigManualProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateProjectConfigInputProjectConfigManualProperty)
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "external_resources_account", &obj.ExternalResourcesAccount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// ProjectsPager can be used to simplify the use of the "ListProjects" method.
//
type ProjectsPager struct {
	hasNext bool
	options *ListProjectsOptions
	client  *ProjectsV1
	pageContext struct {
		next *string
	}
}

// NewProjectsPager returns a new ProjectsPager instance.
func (projects *ProjectsV1) NewProjectsPager(options *ListProjectsOptions) (pager *ProjectsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = fmt.Errorf("the 'options.Start' field should not be set")
		return
	}

	var optionsCopy ListProjectsOptions = *options
	pager = &ProjectsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  projects,
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
