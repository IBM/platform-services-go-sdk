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
 * IBM OpenAPI SDK Code Generator Version: 3.53.0-9710cac3-20220713-193508
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

// GetProjectFile : Get file from the project git repo
// Get file from the project git repo as base64 content.
func (projects *ProjectsV1) GetProjectFile(getProjectFileOptions *GetProjectFileOptions) (result *string, response *core.DetailedResponse, err error) {
	return projects.GetProjectFileWithContext(context.Background(), getProjectFileOptions)
}

// GetProjectFileWithContext is an alternate form of the GetProjectFile method which supports a Context parameter
func (projects *ProjectsV1) GetProjectFileWithContext(ctx context.Context, getProjectFileOptions *GetProjectFileOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectFileOptions, "getProjectFileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectFileOptions, "getProjectFileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectFileOptions.ID,
		"file_path": *getProjectFileOptions.FilePath,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/files/{file_path}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectFileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProjectFile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain")

	if getProjectFileOptions.Branch != nil {
		builder.AddQuery("branch", fmt.Sprint(*getProjectFileOptions.Branch))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, &result)

	return
}

// MergeProject : Merge a Project Definition
// Update project status when a PR merge happens. Install or update existing configs.
func (projects *ProjectsV1) MergeProject(mergeProjectOptions *MergeProjectOptions) (result *ProjectResponse, response *core.DetailedResponse, err error) {
	return projects.MergeProjectWithContext(context.Background(), mergeProjectOptions)
}

// MergeProjectWithContext is an alternate form of the MergeProject method which supports a Context parameter
func (projects *ProjectsV1) MergeProjectWithContext(ctx context.Context, mergeProjectOptions *MergeProjectOptions) (result *ProjectResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(mergeProjectOptions, "mergeProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(mergeProjectOptions, "mergeProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *mergeProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/merge`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range mergeProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "MergeProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if mergeProjectOptions.Name != nil {
		body["name"] = mergeProjectOptions.Name
	}
	if mergeProjectOptions.Description != nil {
		body["description"] = mergeProjectOptions.Description
	}
	if mergeProjectOptions.Configs != nil {
		body["configs"] = mergeProjectOptions.Configs
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ValidateProject : Validate the project for correctness
// Validates the project for correctness.
func (projects *ProjectsV1) ValidateProject(validateProjectOptions *ValidateProjectOptions) (response *core.DetailedResponse, err error) {
	return projects.ValidateProjectWithContext(context.Background(), validateProjectOptions)
}

// ValidateProjectWithContext is an alternate form of the ValidateProject method which supports a Context parameter
func (projects *ProjectsV1) ValidateProjectWithContext(ctx context.Context, validateProjectOptions *ValidateProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(validateProjectOptions, "validateProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(validateProjectOptions, "validateProjectOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/validate`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range validateProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ValidateProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if validateProjectOptions.Name != nil {
		body["name"] = validateProjectOptions.Name
	}
	if validateProjectOptions.Description != nil {
		body["description"] = validateProjectOptions.Description
	}
	if validateProjectOptions.Configs != nil {
		body["configs"] = validateProjectOptions.Configs
	}
	_, err = builder.SetBodyContentJSON(body)
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

// ConfigChanges : Detect changes in project configs
// Detect changes in project configs.
func (projects *ProjectsV1) ConfigChanges(configChangesOptions *ConfigChangesOptions) (result *ProjectConfigChangesResponse, response *core.DetailedResponse, err error) {
	return projects.ConfigChangesWithContext(context.Background(), configChangesOptions)
}

// ConfigChangesWithContext is an alternate form of the ConfigChanges method which supports a Context parameter
func (projects *ProjectsV1) ConfigChangesWithContext(ctx context.Context, configChangesOptions *ConfigChangesOptions) (result *ProjectConfigChangesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(configChangesOptions, "configChangesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(configChangesOptions, "configChangesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *configChangesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_changes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range configChangesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ConfigChanges")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if configChangesOptions.Source != nil {
		body["source"] = configChangesOptions.Source
	}
	if configChangesOptions.PullRequest != nil {
		body["pull_request"] = configChangesOptions.PullRequest
	}
	if configChangesOptions.Target != nil {
		body["target"] = configChangesOptions.Target
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigChangesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateConfig : Add a new config to a project
// Add a new config to a project.
func (projects *ProjectsV1) CreateConfig(createConfigOptions *CreateConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
	return projects.CreateConfigWithContext(context.Background(), createConfigOptions)
}

// CreateConfigWithContext is an alternate form of the CreateConfig method which supports a Context parameter
func (projects *ProjectsV1) CreateConfigWithContext(ctx context.Context, createConfigOptions *CreateConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
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

	_, err = builder.SetBodyContentJSON(createConfigOptions.ProjectConfig)
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
func (projects *ProjectsV1) ListConfigs(listConfigsOptions *ListConfigsOptions) (result []ProjectConfigIntf, response *core.DetailedResponse, err error) {
	return projects.ListConfigsWithContext(context.Background(), listConfigsOptions)
}

// ListConfigsWithContext is an alternate form of the ListConfigs method which supports a Context parameter
func (projects *ProjectsV1) ListConfigsWithContext(ctx context.Context, listConfigsOptions *ListConfigsOptions) (result []ProjectConfigIntf, response *core.DetailedResponse, err error) {
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

	var rawResponse []json.RawMessage
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

// GetConfig : Get a project config
// Returns the specified project config in a given project.
func (projects *ProjectsV1) GetConfig(getConfigOptions *GetConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
	return projects.GetConfigWithContext(context.Background(), getConfigOptions)
}

// GetConfigWithContext is an alternate form of the GetConfig method which supports a Context parameter
func (projects *ProjectsV1) GetConfigWithContext(ctx context.Context, getConfigOptions *GetConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
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
func (projects *ProjectsV1) UpdateConfig(updateConfigOptions *UpdateConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
	return projects.UpdateConfigWithContext(context.Background(), updateConfigOptions)
}

// UpdateConfigWithContext is an alternate form of the UpdateConfig method which supports a Context parameter
func (projects *ProjectsV1) UpdateConfigWithContext(ctx context.Context, updateConfigOptions *UpdateConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
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
	builder.AddHeader("Content-Type", "application/json-patch+json")

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

// PlanConfig : Run a plan job on a given configuration in project
// Run a plan job on a given configuration in project.
func (projects *ProjectsV1) PlanConfig(planConfigOptions *PlanConfigOptions) (result *ConfigJobResponse, response *core.DetailedResponse, err error) {
	return projects.PlanConfigWithContext(context.Background(), planConfigOptions)
}

// PlanConfigWithContext is an alternate form of the PlanConfig method which supports a Context parameter
func (projects *ProjectsV1) PlanConfigWithContext(ctx context.Context, planConfigOptions *PlanConfigOptions) (result *ConfigJobResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(planConfigOptions, "planConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(planConfigOptions, "planConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *planConfigOptions.ID,
		"config_id": *planConfigOptions.ConfigID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_id}/plan`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range planConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "PlanConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if planConfigOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*planConfigOptions.XAuthRefreshToken))
	}

	if planConfigOptions.Version != nil {
		builder.AddQuery("version", fmt.Sprint(*planConfigOptions.Version))
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigJobResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

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
		"config_name": *installConfigOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_name}/install`, pathParamsMap)
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
		"config_name": *uninstallConfigOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_name}/uninstall`, pathParamsMap)
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
		"config_name": *getSchematicsJobOptions.ConfigName,
		"action": *getSchematicsJobOptions.Action,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_name}/{action}/job`, pathParamsMap)
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
	if getSchematicsJobOptions.PullRequest != nil {
		builder.AddQuery("pull_request", fmt.Sprint(*getSchematicsJobOptions.PullRequest))
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
func (projects *ProjectsV1) GetCostEstimate(getCostEstimateOptions *GetCostEstimateOptions) (response *core.DetailedResponse, err error) {
	return projects.GetCostEstimateWithContext(context.Background(), getCostEstimateOptions)
}

// GetCostEstimateWithContext is an alternate form of the GetCostEstimate method which supports a Context parameter
func (projects *ProjectsV1) GetCostEstimateWithContext(ctx context.Context, getCostEstimateOptions *GetCostEstimateOptions) (response *core.DetailedResponse, err error) {
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
		"config_name": *getCostEstimateOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_name}/cost_estimate`, pathParamsMap)
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

	if getCostEstimateOptions.PullRequest != nil {
		builder.AddQuery("pull_request", fmt.Sprint(*getCostEstimateOptions.PullRequest))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

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

// ReceiveGitlabEvents : Webhook for gitlab webhook events
// This is a webhook for gitlab webhook merge request and push events.
func (projects *ProjectsV1) ReceiveGitlabEvents(receiveGitlabEventsOptions *ReceiveGitlabEventsOptions) (response *core.DetailedResponse, err error) {
	return projects.ReceiveGitlabEventsWithContext(context.Background(), receiveGitlabEventsOptions)
}

// ReceiveGitlabEventsWithContext is an alternate form of the ReceiveGitlabEvents method which supports a Context parameter
func (projects *ProjectsV1) ReceiveGitlabEventsWithContext(ctx context.Context, receiveGitlabEventsOptions *ReceiveGitlabEventsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(receiveGitlabEventsOptions, "receiveGitlabEventsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(receiveGitlabEventsOptions, "receiveGitlabEventsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *receiveGitlabEventsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/gitlab_webhook/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range receiveGitlabEventsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ReceiveGitlabEvents")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(receiveGitlabEventsOptions.GitLabEvent)
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
func (projects *ProjectsV1) DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (response *core.DetailedResponse, err error) {
	return projects.DeleteServiceInstanceWithContext(context.Background(), deleteServiceInstanceOptions)
}

// DeleteServiceInstanceWithContext is an alternate form of the DeleteServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) DeleteServiceInstanceWithContext(ctx context.Context, deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (response *core.DetailedResponse, err error) {
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

	response, err = projects.Service.Request(request, nil)

	return
}

// UpdateServiceInstance : allow to change plans and service parameters in a provisioned service instance
// Update plans and service parameters in a provisioned service instance.
func (projects *ProjectsV1) UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) (response *core.DetailedResponse, err error) {
	return projects.UpdateServiceInstanceWithContext(context.Background(), updateServiceInstanceOptions)
}

// UpdateServiceInstanceWithContext is an alternate form of the UpdateServiceInstance method which supports a Context parameter
func (projects *ProjectsV1) UpdateServiceInstanceWithContext(ctx context.Context, updateServiceInstanceOptions *UpdateServiceInstanceOptions) (response *core.DetailedResponse, err error) {
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

	response, err = projects.Service.Request(request, nil)

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

// ConfigChangesOptions : The ConfigChanges options.
type ConfigChangesOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	Source *ProjectInput `json:"source" validate:"required"`

	PullRequest *string `json:"pull_request,omitempty"`

	Target *ProjectInput `json:"target,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewConfigChangesOptions : Instantiate ConfigChangesOptions
func (*ProjectsV1) NewConfigChangesOptions(id string, source *ProjectInput) *ConfigChangesOptions {
	return &ConfigChangesOptions{
		ID: core.StringPtr(id),
		Source: source,
	}
}

// SetID : Allow user to set ID
func (_options *ConfigChangesOptions) SetID(id string) *ConfigChangesOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetSource : Allow user to set Source
func (_options *ConfigChangesOptions) SetSource(source *ProjectInput) *ConfigChangesOptions {
	_options.Source = source
	return _options
}

// SetPullRequest : Allow user to set PullRequest
func (_options *ConfigChangesOptions) SetPullRequest(pullRequest string) *ConfigChangesOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *ConfigChangesOptions) SetTarget(target *ProjectInput) *ConfigChangesOptions {
	_options.Target = target
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ConfigChangesOptions) SetHeaders(param map[string]string) *ConfigChangesOptions {
	options.Headers = param
	return options
}

// ConfigJobResponse : ConfigJobResponse struct
type ConfigJobResponse struct {
	Name *string `json:"name,omitempty"`

	Job *string `json:"job,omitempty"`

	Workspace *string `json:"workspace,omitempty"`

	CartOrder *string `json:"cart_order,omitempty"`

	// The error returned by schematics.
	SchematicsError *string `json:"schematics_error,omitempty"`

	// The error status code returned by schematics.
	SchematicsStatusCode *int64 `json:"schematics_status_code,omitempty"`

	// The timestamp of when the plan job was submitted.
	SchematicsSubmittedAt *int64 `json:"schematics_submitted_at,omitempty"`
}

// UnmarshalConfigJobResponse unmarshals an instance of ConfigJobResponse from the specified map of raw messages.
func UnmarshalConfigJobResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigJobResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "job", &obj.Job)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace", &obj.Workspace)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cart_order", &obj.CartOrder)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schematics_error", &obj.SchematicsError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schematics_status_code", &obj.SchematicsStatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schematics_submitted_at", &obj.SchematicsSubmittedAt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfigSettingItem : ConfigSettingItem struct
type ConfigSettingItem struct {
	Name *string `json:"name" validate:"required"`

	Value *string `json:"value" validate:"required"`
}

// NewConfigSettingItem : Instantiate ConfigSettingItem (Generic Model Constructor)
func (*ProjectsV1) NewConfigSettingItem(name string, value string) (_model *ConfigSettingItem, err error) {
	_model = &ConfigSettingItem{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalConfigSettingItem unmarshals an instance of ConfigSettingItem from the specified map of raw messages.
func UnmarshalConfigSettingItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigSettingItem)
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
	ID *string `json:"id" validate:"required,ne="`

	// The new project definition document.
	ProjectConfig ProjectConfigInputIntf `json:"project_config" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateConfigOptions : Instantiate CreateConfigOptions
func (*ProjectsV1) NewCreateConfigOptions(id string, projectConfig ProjectConfigInputIntf) *CreateConfigOptions {
	return &CreateConfigOptions{
		ID: core.StringPtr(id),
		ProjectConfig: projectConfig,
	}
}

// SetID : Allow user to set ID
func (_options *CreateConfigOptions) SetID(id string) *CreateConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetProjectConfig : Allow user to set ProjectConfig
func (_options *CreateConfigOptions) SetProjectConfig(projectConfig ProjectConfigInputIntf) *CreateConfigOptions {
	_options.ProjectConfig = projectConfig
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateConfigOptions) SetHeaders(param map[string]string) *CreateConfigOptions {
	options.Headers = param
	return options
}

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigInputIntf `json:"configs,omitempty"`

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
func (_options *CreateProjectOptions) SetConfigs(configs []ProjectConfigInputIntf) *CreateProjectOptions {
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

	EventID *string `json:"event_id,omitempty"`

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

// GetActionJobResponse : GetActionJobResponse struct
type GetActionJobResponse struct {
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

	// The name of the config that triggered the action.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// The pull request url associated to where the action was triggered.
	PullRequest *string `json:"pull_request,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCostEstimateOptions : Instantiate GetCostEstimateOptions
func (*ProjectsV1) NewGetCostEstimateOptions(id string, configName string) *GetCostEstimateOptions {
	return &GetCostEstimateOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
	}
}

// SetID : Allow user to set ID
func (_options *GetCostEstimateOptions) SetID(id string) *GetCostEstimateOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *GetCostEstimateOptions) SetConfigName(configName string) *GetCostEstimateOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetPullRequest : Allow user to set PullRequest
func (_options *GetCostEstimateOptions) SetPullRequest(pullRequest string) *GetCostEstimateOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCostEstimateOptions) SetHeaders(param map[string]string) *GetCostEstimateOptions {
	options.Headers = param
	return options
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

// GetProjectFileOptions : The GetProjectFile options.
type GetProjectFileOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// URL-encoded full path to new file, such as lib%2Fclass%2Erb.
	FilePath *string `json:"file_path" validate:"required,ne="`

	// Set this parameter if you want to get the file from a specific branch.
	Branch *string `json:"branch,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectFileOptions : Instantiate GetProjectFileOptions
func (*ProjectsV1) NewGetProjectFileOptions(id string, filePath string) *GetProjectFileOptions {
	return &GetProjectFileOptions{
		ID: core.StringPtr(id),
		FilePath: core.StringPtr(filePath),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectFileOptions) SetID(id string) *GetProjectFileOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetFilePath : Allow user to set FilePath
func (_options *GetProjectFileOptions) SetFilePath(filePath string) *GetProjectFileOptions {
	_options.FilePath = core.StringPtr(filePath)
	return _options
}

// SetBranch : Allow user to set Branch
func (_options *GetProjectFileOptions) SetBranch(branch string) *GetProjectFileOptions {
	_options.Branch = core.StringPtr(branch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectFileOptions) SetHeaders(param map[string]string) *GetProjectFileOptions {
	options.Headers = param
	return options
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

// GetProjectResponse : GetProjectResponse struct
type GetProjectResponse struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	Crn *string `json:"crn,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

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

	State *string `json:"state,omitempty"`
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSchematicsJobOptions : The GetSchematicsJob options.
type GetSchematicsJobOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config that triggered the action.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// The triggered action.
	Action *string `json:"action" validate:"required,ne="`

	// The timestamp of when the action was triggered.
	Since *int64 `json:"since,omitempty"`

	// The pull request url associated to where the action was triggered.
	PullRequest *string `json:"pull_request,omitempty"`

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
func (*ProjectsV1) NewGetSchematicsJobOptions(id string, configName string, action string) *GetSchematicsJobOptions {
	return &GetSchematicsJobOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
		Action: core.StringPtr(action),
	}
}

// SetID : Allow user to set ID
func (_options *GetSchematicsJobOptions) SetID(id string) *GetSchematicsJobOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *GetSchematicsJobOptions) SetConfigName(configName string) *GetSchematicsJobOptions {
	_options.ConfigName = core.StringPtr(configName)
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

// SetPullRequest : Allow user to set PullRequest
func (_options *GetSchematicsJobOptions) SetPullRequest(pullRequest string) *GetSchematicsJobOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
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

// GitLabEvent : A GitLab event payload.
type GitLabEvent struct {
	// The event kind.
	ObjectKind *string `json:"object_kind" validate:"required"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewGitLabEvent : Instantiate GitLabEvent (Generic Model Constructor)
func (*ProjectsV1) NewGitLabEvent(objectKind string) (_model *GitLabEvent, err error) {
	_model = &GitLabEvent{
		ObjectKind: core.StringPtr(objectKind),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of GitLabEvent
func (o *GitLabEvent) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// SetProperties allows the user to set a map of arbitrary properties on an instance of GitLabEvent
func (o *GitLabEvent) SetProperties(m map[string]interface{}) {
	o.additionalProperties = make(map[string]interface{})
	for k, v := range m {
		o.additionalProperties[k] = v
	}
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of GitLabEvent
func (o *GitLabEvent) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of GitLabEvent
func (o *GitLabEvent) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of GitLabEvent
func (o *GitLabEvent) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.ObjectKind != nil {
		m["object_kind"] = o.ObjectKind
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalGitLabEvent unmarshals an instance of GitLabEvent from the specified map of raw messages.
func UnmarshalGitLabEvent(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GitLabEvent)
	err = core.UnmarshalPrimitive(m, "object_kind", &obj.ObjectKind)
	if err != nil {
		return
	}
	delete(m, "object_kind")
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

// Health : Health struct
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

	// The variable value.
	Value interface{} `json:"value,omitempty"`

	// The variable default value.
	Default interface{} `json:"default,omitempty"`
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

// NewInputVariable : Instantiate InputVariable (Generic Model Constructor)
func (*ProjectsV1) NewInputVariable(name string, typeVar string) (_model *InputVariable, err error) {
	_model = &InputVariable{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

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
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
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

	// The name of the config to install.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallConfigOptions : Instantiate InstallConfigOptions
func (*ProjectsV1) NewInstallConfigOptions(id string, configName string) *InstallConfigOptions {
	return &InstallConfigOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
	}
}

// SetID : Allow user to set ID
func (_options *InstallConfigOptions) SetID(id string) *InstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *InstallConfigOptions) SetConfigName(configName string) *InstallConfigOptions {
	_options.ConfigName = core.StringPtr(configName)
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
func (*ProjectsV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
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

// MergeProjectOptions : The MergeProject options.
type MergeProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigInputIntf `json:"configs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewMergeProjectOptions : Instantiate MergeProjectOptions
func (*ProjectsV1) NewMergeProjectOptions(id string, name string) *MergeProjectOptions {
	return &MergeProjectOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
	}
}

// SetID : Allow user to set ID
func (_options *MergeProjectOptions) SetID(id string) *MergeProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *MergeProjectOptions) SetName(name string) *MergeProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *MergeProjectOptions) SetDescription(description string) *MergeProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *MergeProjectOptions) SetConfigs(configs []ProjectConfigInputIntf) *MergeProjectOptions {
	_options.Configs = configs
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MergeProjectOptions) SetHeaders(param map[string]string) *MergeProjectOptions {
	options.Headers = param
	return options
}

// NotificationEvent : NotificationEvent struct
type NotificationEvent struct {
	// Type of event.
	Event *string `json:"event" validate:"required"`

	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data interface{} `json:"data,omitempty"`
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

	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data interface{} `json:"data,omitempty"`

	ID *string `json:"_id" validate:"required"`
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
	err = core.UnmarshalPrimitive(m, "_id", &obj.ID)
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

	Target *string `json:"target" validate:"required"`

	// The id of the event producer.
	Source *string `json:"source,omitempty"`

	// The URL you can go to as next steps.
	ActionURL *string `json:"action_url,omitempty"`

	// Any relevant metadata to be stored.
	Data interface{} `json:"data,omitempty"`

	ID *string `json:"_id" validate:"required"`

	// whether or not the event successfully posted.
	Status *string `json:"status,omitempty"`

	Reasons []interface{} `json:"reasons,omitempty"`
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
	err = core.UnmarshalPrimitive(m, "_id", &obj.ID)
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
	// The output value name.
	Name *string `json:"name" validate:"required"`

	// A descriptive of the output value.
	Description *string `json:"description,omitempty"`

	// The output value.
	Value []string `json:"value,omitempty"`
}

// NewOutputValue : Instantiate OutputValue (Generic Model Constructor)
func (*ProjectsV1) NewOutputValue(name string) (_model *OutputValue, err error) {
	_model = &OutputValue{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
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

// PlanConfigOptions : The PlanConfig options.
type PlanConfigOptions struct {
	// The IAM refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The id of the config to trigger a plan.
	ConfigID *string `json:"config_id" validate:"required,ne="`

	// The version of the config that the plan should trigger against.
	Version *string `json:"version,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPlanConfigOptions : Instantiate PlanConfigOptions
func (*ProjectsV1) NewPlanConfigOptions(xAuthRefreshToken string, id string, configID string) *PlanConfigOptions {
	return &PlanConfigOptions{
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
		ID: core.StringPtr(id),
		ConfigID: core.StringPtr(configID),
	}
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (_options *PlanConfigOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *PlanConfigOptions {
	_options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return _options
}

// SetID : Allow user to set ID
func (_options *PlanConfigOptions) SetID(id string) *PlanConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigID : Allow user to set ConfigID
func (_options *PlanConfigOptions) SetConfigID(configID string) *PlanConfigOptions {
	_options.ConfigID = core.StringPtr(configID)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *PlanConfigOptions) SetVersion(version string) *PlanConfigOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PlanConfigOptions) SetHeaders(param map[string]string) *PlanConfigOptions {
	options.Headers = param
	return options
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

// ProjectConfig : ProjectConfig struct
// Models which "extend" this model:
// - ProjectConfigProjectConfigCommon
// - ProjectConfigProp
type ProjectConfig struct {
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name,omitempty"`

	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	LocatorID *string `json:"locator_id,omitempty"`

	Input []InputVariable `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`
}
func (*ProjectConfig) isaProjectConfig() bool {
	return true
}

type ProjectConfigIntf interface {
	isaProjectConfig() bool
}

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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*ProjectsV1) NewProjectConfigPatch(projectConfig *ProjectConfig) (_patch []JSONPatchOperation) {
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
	if (projectConfig.Type != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/type"),
			Value: projectConfig.Type,
		})
	}
	if (projectConfig.ExternalResourcesAccount != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/external_resources_account"),
			Value: projectConfig.ExternalResourcesAccount,
		})
	}
	if (projectConfig.LocatorID != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/locator_id"),
			Value: projectConfig.LocatorID,
		})
	}
	if (projectConfig.Input != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/input"),
			Value: projectConfig.Input,
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

// ProjectConfigChangesResponse : ProjectConfigChangesResponse struct
type ProjectConfigChangesResponse struct {
	Added []ProjectConfigChangesResponseAddedItem `json:"added" validate:"required"`

	Deleted []ProjectConfigChangesResponseDeletedItem `json:"deleted" validate:"required"`

	Changed []ProjectConfigChangesResponseChangedItem `json:"changed" validate:"required"`
}

// UnmarshalProjectConfigChangesResponse unmarshals an instance of ProjectConfigChangesResponse from the specified map of raw messages.
func UnmarshalProjectConfigChangesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigChangesResponse)
	err = core.UnmarshalModel(m, "added", &obj.Added, UnmarshalProjectConfigChangesResponseAddedItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deleted", &obj.Deleted, UnmarshalProjectConfigChangesResponseDeletedItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "changed", &obj.Changed, UnmarshalProjectConfigChangesResponseChangedItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigChangesResponseAddedItem : ProjectConfigChangesResponseAddedItem struct
type ProjectConfigChangesResponseAddedItem struct {
	Name *string `json:"name,omitempty"`
}

// UnmarshalProjectConfigChangesResponseAddedItem unmarshals an instance of ProjectConfigChangesResponseAddedItem from the specified map of raw messages.
func UnmarshalProjectConfigChangesResponseAddedItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigChangesResponseAddedItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigChangesResponseChangedItem : ProjectConfigChangesResponseChangedItem struct
type ProjectConfigChangesResponseChangedItem struct {
	Name *string `json:"name,omitempty"`

	NewName *string `json:"new_name,omitempty"`
}

// UnmarshalProjectConfigChangesResponseChangedItem unmarshals an instance of ProjectConfigChangesResponseChangedItem from the specified map of raw messages.
func UnmarshalProjectConfigChangesResponseChangedItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigChangesResponseChangedItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "new_name", &obj.NewName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigChangesResponseDeletedItem : ProjectConfigChangesResponseDeletedItem struct
type ProjectConfigChangesResponseDeletedItem struct {
	Name *string `json:"name,omitempty"`
}

// UnmarshalProjectConfigChangesResponseDeletedItem unmarshals an instance of ProjectConfigChangesResponseDeletedItem from the specified map of raw messages.
func UnmarshalProjectConfigChangesResponseDeletedItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigChangesResponseDeletedItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigInput : ProjectConfigInput struct
// Models which "extend" this model:
// - ProjectConfigInputProjectConfigCommon
// - ProjectConfigInputProp
type ProjectConfigInput struct {
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name,omitempty"`

	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	LocatorID *string `json:"locator_id,omitempty"`

	Input []InputVariable `json:"input,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`
}
func (*ProjectConfigInput) isaProjectConfigInput() bool {
	return true
}

type ProjectConfigInputIntf interface {
	isaProjectConfigInput() bool
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectInput : ProjectInput struct
type ProjectInput struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigInputIntf `json:"configs,omitempty"`
}

// NewProjectInput : Instantiate ProjectInput (Generic Model Constructor)
func (*ProjectsV1) NewProjectInput(name string) (_model *ProjectInput, err error) {
	_model = &ProjectInput{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectInput unmarshals an instance of ProjectInput from the specified map of raw messages.
func UnmarshalProjectInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfigInput)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListItem : ProjectListItem struct
type ProjectListItem struct {
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

// ProjectResponse : ProjectResponse struct
type ProjectResponse struct {
	// The project name.
	Name *string `json:"name,omitempty"`

	ID *string `json:"id,omitempty"`

	Definition *GetProjectResponse `json:"definition,omitempty"`
}

// UnmarshalProjectResponse unmarshals an instance of ProjectResponse from the specified map of raw messages.
func UnmarshalProjectResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalGetProjectResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

func (*ProjectsV1) NewProjectUpdatePatch(projectUpdate *ProjectUpdate) (_patch []JSONPatchOperation) {
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
	EventProperties interface{} `json:"event_properties,omitempty"`

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

// ReceiveGitlabEventsOptions : The ReceiveGitlabEvents options.
type ReceiveGitlabEventsOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// A gitlab event.
	GitLabEvent *GitLabEvent `json:"GitLabEvent" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReceiveGitlabEventsOptions : Instantiate ReceiveGitlabEventsOptions
func (*ProjectsV1) NewReceiveGitlabEventsOptions(id string, gitLabEvent *GitLabEvent) *ReceiveGitlabEventsOptions {
	return &ReceiveGitlabEventsOptions{
		ID: core.StringPtr(id),
		GitLabEvent: gitLabEvent,
	}
}

// SetID : Allow user to set ID
func (_options *ReceiveGitlabEventsOptions) SetID(id string) *ReceiveGitlabEventsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetGitLabEvent : Allow user to set GitLabEvent
func (_options *ReceiveGitlabEventsOptions) SetGitLabEvent(gitLabEvent *GitLabEvent) *ReceiveGitlabEventsOptions {
	_options.GitLabEvent = gitLabEvent
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReceiveGitlabEventsOptions) SetHeaders(param map[string]string) *ReceiveGitlabEventsOptions {
	options.Headers = param
	return options
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
	Parameters interface{} `json:"parameters,omitempty"`

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
func (_options *ReplaceServiceInstanceOptions) SetParameters(parameters interface{}) *ReplaceServiceInstanceOptions {
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
	ReasonCode interface{} `json:"reason_code,omitempty"`

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
func (_options *ReplaceServiceInstanceStateOptions) SetReasonCode(reasonCode interface{}) *ReplaceServiceInstanceStateOptions {
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

	// The name of the config to uninstall.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUninstallConfigOptions : Instantiate UninstallConfigOptions
func (*ProjectsV1) NewUninstallConfigOptions(id string, configName string) *UninstallConfigOptions {
	return &UninstallConfigOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
	}
}

// SetID : Allow user to set ID
func (_options *UninstallConfigOptions) SetID(id string) *UninstallConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *UninstallConfigOptions) SetConfigName(configName string) *UninstallConfigOptions {
	_options.ConfigName = core.StringPtr(configName)
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
	ProjectConfig []JSONPatchOperation `json:"project_config" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateConfigOptions : Instantiate UpdateConfigOptions
func (*ProjectsV1) NewUpdateConfigOptions(id string, configID string, projectConfig []JSONPatchOperation) *UpdateConfigOptions {
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

// SetHeaders : Allow user to set Headers
func (options *UpdateConfigOptions) SetHeaders(param map[string]string) *UpdateConfigOptions {
	options.Headers = param
	return options
}

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The new project definition document.
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectsV1) NewUpdateProjectOptions(id string, jsonPatchOperation []JSONPatchOperation) *UpdateProjectOptions {
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

// UpdateResult : Result of deprovisioning service instance.
type UpdateResult struct {
}

// UnmarshalUpdateResult unmarshals an instance of UpdateResult from the specified map of raw messages.
func UnmarshalUpdateResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateResult)
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*ProjectsV1) NewUpdateResultPatch(updateResult *UpdateResult) (_patch []JSONPatchOperation) {
	return
}

// UpdateServiceInstanceOptions : The UpdateServiceInstance options.
type UpdateServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required,ne="`

	// It contains the query filters and the search token (initally set to null or undefined).
	JSONPatchOperation []JSONPatchOperation `json:"JsonPatchOperation" validate:"required"`

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
func (*ProjectsV1) NewUpdateServiceInstanceOptions(instanceID string, jsonPatchOperation []JSONPatchOperation) *UpdateServiceInstanceOptions {
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

// ValidateProjectOptions : The ValidateProject options.
type ValidateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigInputIntf `json:"configs,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewValidateProjectOptions : Instantiate ValidateProjectOptions
func (*ProjectsV1) NewValidateProjectOptions(name string) *ValidateProjectOptions {
	return &ValidateProjectOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *ValidateProjectOptions) SetName(name string) *ValidateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *ValidateProjectOptions) SetDescription(description string) *ValidateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetConfigs : Allow user to set Configs
func (_options *ValidateProjectOptions) SetConfigs(configs []ProjectConfigInputIntf) *ValidateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ValidateProjectOptions) SetHeaders(param map[string]string) *ValidateProjectOptions {
	options.Headers = param
	return options
}

// ProjectConfigInputProp : ProjectConfigInputProp struct
// This model "extends" ProjectConfigInput
type ProjectConfigInputProp struct {
	Type *string `json:"type,omitempty"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	LocatorID *string `json:"locator_id" validate:"required"`

	Input []InputVariable `json:"input,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`
}

// NewProjectConfigInputProp : Instantiate ProjectConfigInputProp (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigInputProp(locatorID string) (_model *ProjectConfigInputProp, err error) {
	_model = &ProjectConfigInputProp{
		LocatorID: core.StringPtr(locatorID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ProjectConfigInputProp) isaProjectConfigInput() bool {
	return true
}

// UnmarshalProjectConfigInputProp unmarshals an instance of ProjectConfigInputProp from the specified map of raw messages.
func UnmarshalProjectConfigInputProp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigInputProp)
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigInputProjectConfigCommon : ProjectConfigInputProjectConfigCommon struct
// This model "extends" ProjectConfigInput
type ProjectConfigInputProjectConfigCommon struct {
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`
}

// NewProjectConfigInputProjectConfigCommon : Instantiate ProjectConfigInputProjectConfigCommon (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigInputProjectConfigCommon(name string) (_model *ProjectConfigInputProjectConfigCommon, err error) {
	_model = &ProjectConfigInputProjectConfigCommon{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ProjectConfigInputProjectConfigCommon) isaProjectConfigInput() bool {
	return true
}

// UnmarshalProjectConfigInputProjectConfigCommon unmarshals an instance of ProjectConfigInputProjectConfigCommon from the specified map of raw messages.
func UnmarshalProjectConfigInputProjectConfigCommon(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigInputProjectConfigCommon)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigProp : ProjectConfigProp struct
// This model "extends" ProjectConfig
type ProjectConfigProp struct {
	Type *string `json:"type" validate:"required"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	LocatorID *string `json:"locator_id" validate:"required"`

	Input []InputVariable `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`
}

func (*ProjectConfigProp) isaProjectConfig() bool {
	return true
}

// UnmarshalProjectConfigProp unmarshals an instance of ProjectConfigProp from the specified map of raw messages.
func UnmarshalProjectConfigProp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigProp)
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigProjectConfigCommon : ProjectConfigProjectConfigCommon struct
// This model "extends" ProjectConfig
type ProjectConfigProjectConfigCommon struct {
	ID *string `json:"id,omitempty"`

	// The config name.
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`
}

func (*ProjectConfigProjectConfigCommon) isaProjectConfig() bool {
	return true
}

// UnmarshalProjectConfigProjectConfigCommon unmarshals an instance of ProjectConfigProjectConfigCommon from the specified map of raw messages.
func UnmarshalProjectConfigProjectConfigCommon(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigProjectConfigCommon)
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
