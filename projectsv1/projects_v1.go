/**
 * (C) Copyright IBM Corp. 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.54.0-af6d2126-20220803-151219
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
func (projects *ProjectsV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *ProjectResponse, response *core.DetailedResponse, err error) {
	return projects.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (projects *ProjectsV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *ProjectResponse, response *core.DetailedResponse, err error) {
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
	if createProjectOptions.XIamApi != nil {
		builder.AddHeader("X-Iam-Api", fmt.Sprint(*createProjectOptions.XIamApi))
	}
	if createProjectOptions.Prefer != nil {
		builder.AddHeader("prefer", fmt.Sprint(*createProjectOptions.Prefer))
	}

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
	if createProjectOptions.Dashboard != nil {
		body["dashboard"] = createProjectOptions.Dashboard
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

// ListProjects : List Projects
// List existing Projects. Projects are sorted by id.
func (projects *ProjectsV1) ListProjects(listProjectsOptions *ListProjectsOptions) (result *ListProjectsResponse, response *core.DetailedResponse, err error) {
	return projects.ListProjectsWithContext(context.Background(), listProjectsOptions)
}

// ListProjectsWithContext is an alternate form of the ListProjects method which supports a Context parameter
func (projects *ProjectsV1) ListProjectsWithContext(ctx context.Context, listProjectsOptions *ListProjectsOptions) (result *ListProjectsResponse, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListProjectsResponse)
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

	if getProjectOptions.Branch != nil {
		builder.AddQuery("branch", fmt.Sprint(*getProjectOptions.Branch))
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
func (projects *ProjectsV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *UpdateResponsePullRequest, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *UpdateResponsePullRequest, response *core.DetailedResponse, err error) {
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

	builder := core.NewRequestBuilder(core.PUT)
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

	if updateProjectOptions.Branch != nil {
		builder.AddQuery("branch", fmt.Sprint(*updateProjectOptions.Branch))
	}

	body := make(map[string]interface{})
	if updateProjectOptions.Name != nil {
		body["name"] = updateProjectOptions.Name
	}
	if updateProjectOptions.Description != nil {
		body["description"] = updateProjectOptions.Description
	}
	if updateProjectOptions.Configs != nil {
		body["configs"] = updateProjectOptions.Configs
	}
	if updateProjectOptions.Dashboard != nil {
		body["dashboard"] = updateProjectOptions.Dashboard
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateResponsePullRequest)
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

// CheckProject : Check Project
// Check one or more project's configurations. The operation verifies that all the config resources are installed, and
// compliant to the configured security postures. You can track the operation and view the results using the project
// status api.
func (projects *ProjectsV1) CheckProject(checkProjectOptions *CheckProjectOptions) (response *core.DetailedResponse, err error) {
	return projects.CheckProjectWithContext(context.Background(), checkProjectOptions)
}

// CheckProjectWithContext is an alternate form of the CheckProject method which supports a Context parameter
func (projects *ProjectsV1) CheckProjectWithContext(ctx context.Context, checkProjectOptions *CheckProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(checkProjectOptions, "checkProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(checkProjectOptions, "checkProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *checkProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/check`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range checkProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "CheckProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if checkProjectOptions.ConfigNames != nil {
		body["config_names"] = checkProjectOptions.ConfigNames
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
	if mergeProjectOptions.Dashboard != nil {
		body["dashboard"] = mergeProjectOptions.Dashboard
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
	if validateProjectOptions.Dashboard != nil {
		body["dashboard"] = validateProjectOptions.Dashboard
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

// InitProjectConfig : Initialize the configs from a catalog solution
// Initialize the configs from a catalog solution.
func (projects *ProjectsV1) InitProjectConfig(initProjectConfigOptions *InitProjectConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
	return projects.InitProjectConfigWithContext(context.Background(), initProjectConfigOptions)
}

// InitProjectConfigWithContext is an alternate form of the InitProjectConfig method which supports a Context parameter
func (projects *ProjectsV1) InitProjectConfigWithContext(ctx context.Context, initProjectConfigOptions *InitProjectConfigOptions) (result ProjectConfigIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(initProjectConfigOptions, "initProjectConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(initProjectConfigOptions, "initProjectConfigOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/init_config`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range initProjectConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "InitProjectConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if initProjectConfigOptions.LocatorID != nil {
		body["locator_id"] = initProjectConfigOptions.LocatorID
	}
	if initProjectConfigOptions.ConfigName != nil {
		body["config_name"] = initProjectConfigOptions.ConfigName
	}
	if initProjectConfigOptions.Description != nil {
		body["description"] = initProjectConfigOptions.Description
	}
	if initProjectConfigOptions.Labels != nil {
		body["labels"] = initProjectConfigOptions.Labels
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

// GetProjectStatus : Get project status information
// Returns the detailed project status, including all the config statuses and all the computed statuses sent by external
// tools.
func (projects *ProjectsV1) GetProjectStatus(getProjectStatusOptions *GetProjectStatusOptions) (result *ProjectStatus, response *core.DetailedResponse, err error) {
	return projects.GetProjectStatusWithContext(context.Background(), getProjectStatusOptions)
}

// GetProjectStatusWithContext is an alternate form of the GetProjectStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectStatusWithContext(ctx context.Context, getProjectStatusOptions *GetProjectStatusOptions) (result *ProjectStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectStatusOptions, "getProjectStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectStatusOptions, "getProjectStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectStatusOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/status`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProjectStatus")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectStatus : Update project status information
// Update the project status.
func (projects *ProjectsV1) UpdateProjectStatus(updateProjectStatusOptions *UpdateProjectStatusOptions) (result *ProjectStatus, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectStatusWithContext(context.Background(), updateProjectStatusOptions)
}

// UpdateProjectStatusWithContext is an alternate form of the UpdateProjectStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectStatusWithContext(ctx context.Context, updateProjectStatusOptions *UpdateProjectStatusOptions) (result *ProjectStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectStatusOptions, "updateProjectStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectStatusOptions, "updateProjectStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectStatusOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/status`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateProjectStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectStatusOptions.State != nil {
		body["state"] = updateProjectStatusOptions.State
	}
	if updateProjectStatusOptions.History != nil {
		body["history"] = updateProjectStatusOptions.History
	}
	if updateProjectStatusOptions.GitRepo != nil {
		body["git_repo"] = updateProjectStatusOptions.GitRepo
	}
	if updateProjectStatusOptions.Toolchain != nil {
		body["toolchain"] = updateProjectStatusOptions.Toolchain
	}
	if updateProjectStatusOptions.Schematics != nil {
		body["schematics"] = updateProjectStatusOptions.Schematics
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProjectComputedStatus : Get computed status information
// Returns a project computed status.
func (projects *ProjectsV1) GetProjectComputedStatus(getProjectComputedStatusOptions *GetProjectComputedStatusOptions) (result *ProjectComputedStatus, response *core.DetailedResponse, err error) {
	return projects.GetProjectComputedStatusWithContext(context.Background(), getProjectComputedStatusOptions)
}

// GetProjectComputedStatusWithContext is an alternate form of the GetProjectComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectComputedStatusWithContext(ctx context.Context, getProjectComputedStatusOptions *GetProjectComputedStatusOptions) (result *ProjectComputedStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectComputedStatusOptions, "getProjectComputedStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectComputedStatusOptions, "getProjectComputedStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectComputedStatusOptions.ID,
		"computed_status": *getProjectComputedStatusOptions.ComputedStatus,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/status/{computed_status}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectComputedStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProjectComputedStatus")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectComputedStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectComputedStatus : Update computed status information
// Update a project computed status.
func (projects *ProjectsV1) UpdateProjectComputedStatus(updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions) (result *ProjectComputedStatus, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectComputedStatusWithContext(context.Background(), updateProjectComputedStatusOptions)
}

// UpdateProjectComputedStatusWithContext is an alternate form of the UpdateProjectComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectComputedStatusWithContext(ctx context.Context, updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions) (result *ProjectComputedStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectComputedStatusOptions, "updateProjectComputedStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectComputedStatusOptions, "updateProjectComputedStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectComputedStatusOptions.ID,
		"computed_status": *updateProjectComputedStatusOptions.ComputedStatus,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/status/{computed_status}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectComputedStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateProjectComputedStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectComputedStatusOptions.ComputedStatuses != nil {
		body["computed_statuses"] = updateProjectComputedStatusOptions.ComputedStatuses
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectComputedStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProjectConfigStatuses : List project configs status information
// Returns the status of all project configs.
func (projects *ProjectsV1) ListProjectConfigStatuses(listProjectConfigStatusesOptions *ListProjectConfigStatusesOptions) (result *ProjectConfigStatuses, response *core.DetailedResponse, err error) {
	return projects.ListProjectConfigStatusesWithContext(context.Background(), listProjectConfigStatusesOptions)
}

// ListProjectConfigStatusesWithContext is an alternate form of the ListProjectConfigStatuses method which supports a Context parameter
func (projects *ProjectsV1) ListProjectConfigStatusesWithContext(ctx context.Context, listProjectConfigStatusesOptions *ListProjectConfigStatusesOptions) (result *ProjectConfigStatuses, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listProjectConfigStatusesOptions, "listProjectConfigStatusesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listProjectConfigStatusesOptions, "listProjectConfigStatusesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *listProjectConfigStatusesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_statuses`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProjectConfigStatusesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "ListProjectConfigStatuses")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigStatuses)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProjectConfigStatus : Get project config status information
// Update a config status and eventually the output values.
func (projects *ProjectsV1) GetProjectConfigStatus(getProjectConfigStatusOptions *GetProjectConfigStatusOptions) (result *ProjectConfigStatus, response *core.DetailedResponse, err error) {
	return projects.GetProjectConfigStatusWithContext(context.Background(), getProjectConfigStatusOptions)
}

// GetProjectConfigStatusWithContext is an alternate form of the GetProjectConfigStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectConfigStatusWithContext(ctx context.Context, getProjectConfigStatusOptions *GetProjectConfigStatusOptions) (result *ProjectConfigStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectConfigStatusOptions, "getProjectConfigStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectConfigStatusOptions, "getProjectConfigStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectConfigStatusOptions.ID,
		"config_name": *getProjectConfigStatusOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_statuses/{config_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectConfigStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProjectConfigStatus")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectConfigStatus : Update project config status information
// Update a config status and eventually the output values.
func (projects *ProjectsV1) UpdateProjectConfigStatus(updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions) (result *ProjectConfigStatus, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectConfigStatusWithContext(context.Background(), updateProjectConfigStatusOptions)
}

// UpdateProjectConfigStatusWithContext is an alternate form of the UpdateProjectConfigStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectConfigStatusWithContext(ctx context.Context, updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions) (result *ProjectConfigStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectConfigStatusOptions, "updateProjectConfigStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectConfigStatusOptions, "updateProjectConfigStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectConfigStatusOptions.ID,
		"config_name": *updateProjectConfigStatusOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_statuses/{config_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectConfigStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateProjectConfigStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectConfigStatusOptions.Status != nil {
		body["status"] = updateProjectConfigStatusOptions.Status
	}
	if updateProjectConfigStatusOptions.Messages != nil {
		body["messages"] = updateProjectConfigStatusOptions.Messages
	}
	if updateProjectConfigStatusOptions.PipelineRun != nil {
		body["pipeline_run"] = updateProjectConfigStatusOptions.PipelineRun
	}
	if updateProjectConfigStatusOptions.SchematicsResourceID != nil {
		body["schematics_resource_id"] = updateProjectConfigStatusOptions.SchematicsResourceID
	}
	if updateProjectConfigStatusOptions.Output != nil {
		body["output"] = updateProjectConfigStatusOptions.Output
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetProjectConfigComputedStatus : Get config computed status information
// Returns the given computed status for the config.
func (projects *ProjectsV1) GetProjectConfigComputedStatus(getProjectConfigComputedStatusOptions *GetProjectConfigComputedStatusOptions) (result *ProjectConfigComputedStatus, response *core.DetailedResponse, err error) {
	return projects.GetProjectConfigComputedStatusWithContext(context.Background(), getProjectConfigComputedStatusOptions)
}

// GetProjectConfigComputedStatusWithContext is an alternate form of the GetProjectConfigComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectConfigComputedStatusWithContext(ctx context.Context, getProjectConfigComputedStatusOptions *GetProjectConfigComputedStatusOptions) (result *ProjectConfigComputedStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getProjectConfigComputedStatusOptions, "getProjectConfigComputedStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getProjectConfigComputedStatusOptions, "getProjectConfigComputedStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getProjectConfigComputedStatusOptions.ID,
		"config_name": *getProjectConfigComputedStatusOptions.ConfigName,
		"computed_status": *getProjectConfigComputedStatusOptions.ComputedStatus,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_statuses/{config_name}/{computed_status}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getProjectConfigComputedStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "GetProjectConfigComputedStatus")
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigComputedStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectConfigComputedStatus : Update a project computed status information
// Update a computed status with content to which the projects service is agnostic.
func (projects *ProjectsV1) UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions) (result *ProjectConfigComputedStatus, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectConfigComputedStatusWithContext(context.Background(), updateProjectConfigComputedStatusOptions)
}

// UpdateProjectConfigComputedStatusWithContext is an alternate form of the UpdateProjectConfigComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectConfigComputedStatusWithContext(ctx context.Context, updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions) (result *ProjectConfigComputedStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateProjectConfigComputedStatusOptions, "updateProjectConfigComputedStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateProjectConfigComputedStatusOptions, "updateProjectConfigComputedStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateProjectConfigComputedStatusOptions.ID,
		"config_name": *updateProjectConfigComputedStatusOptions.ConfigName,
		"computed_status": *updateProjectConfigComputedStatusOptions.ComputedStatus,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/config_statuses/{config_name}/{computed_status}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateProjectConfigComputedStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdateProjectConfigComputedStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateProjectConfigComputedStatusOptions.RequestBody)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectConfigComputedStatus)
		if err != nil {
			return
		}
		response.Result = result
	}

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

// Notify : Notify a project change
// Notify the Projects Service that something has changed in a Project status.
func (projects *ProjectsV1) Notify(notifyOptions *NotifyOptions) (response *core.DetailedResponse, err error) {
	return projects.NotifyWithContext(context.Background(), notifyOptions)
}

// NotifyWithContext is an alternate form of the Notify method which supports a Context parameter
func (projects *ProjectsV1) NotifyWithContext(ctx context.Context, notifyOptions *NotifyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(notifyOptions, "notifyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(notifyOptions, "notifyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/notify`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range notifyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "Notify")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if notifyOptions.ID != nil {
		body["id"] = notifyOptions.ID
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

// RegisterPullRequest : Register a pull request
// Registers a pull request on the associated git repo of the project.
func (projects *ProjectsV1) RegisterPullRequest(registerPullRequestOptions *RegisterPullRequestOptions) (response *core.DetailedResponse, err error) {
	return projects.RegisterPullRequestWithContext(context.Background(), registerPullRequestOptions)
}

// RegisterPullRequestWithContext is an alternate form of the RegisterPullRequest method which supports a Context parameter
func (projects *ProjectsV1) RegisterPullRequestWithContext(ctx context.Context, registerPullRequestOptions *RegisterPullRequestOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(registerPullRequestOptions, "registerPullRequestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(registerPullRequestOptions, "registerPullRequestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *registerPullRequestOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/pullrequest`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range registerPullRequestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "RegisterPullRequest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if registerPullRequestOptions.Branch != nil {
		body["branch"] = registerPullRequestOptions.Branch
	}
	if registerPullRequestOptions.URL != nil {
		body["url"] = registerPullRequestOptions.URL
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

// DeregisterPullRequest : Deregister a pull request
// Deregisters a pull request on the associated git repo of the project.
func (projects *ProjectsV1) DeregisterPullRequest(deregisterPullRequestOptions *DeregisterPullRequestOptions) (response *core.DetailedResponse, err error) {
	return projects.DeregisterPullRequestWithContext(context.Background(), deregisterPullRequestOptions)
}

// DeregisterPullRequestWithContext is an alternate form of the DeregisterPullRequest method which supports a Context parameter
func (projects *ProjectsV1) DeregisterPullRequestWithContext(ctx context.Context, deregisterPullRequestOptions *DeregisterPullRequestOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deregisterPullRequestOptions, "deregisterPullRequestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deregisterPullRequestOptions, "deregisterPullRequestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deregisterPullRequestOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/pullrequest`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deregisterPullRequestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "DeregisterPullRequest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("url", fmt.Sprint(*deregisterPullRequestOptions.URL))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = projects.Service.Request(request, nil)

	return
}

// UpdatePullRequestConfigs : Update configs associated to the given pull request
// Updates configs referenced by a pull request. Dereference the pull request from the configs if they are newly added.
// Otherwise update or remove the configs from the project status document wrt the pull request changes.
func (projects *ProjectsV1) UpdatePullRequestConfigs(updatePullRequestConfigsOptions *UpdatePullRequestConfigsOptions) (response *core.DetailedResponse, err error) {
	return projects.UpdatePullRequestConfigsWithContext(context.Background(), updatePullRequestConfigsOptions)
}

// UpdatePullRequestConfigsWithContext is an alternate form of the UpdatePullRequestConfigs method which supports a Context parameter
func (projects *ProjectsV1) UpdatePullRequestConfigsWithContext(ctx context.Context, updatePullRequestConfigsOptions *UpdatePullRequestConfigsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePullRequestConfigsOptions, "updatePullRequestConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePullRequestConfigsOptions, "updatePullRequestConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updatePullRequestConfigsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/pullrequest/configs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePullRequestConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UpdatePullRequestConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	if updatePullRequestConfigsOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*updatePullRequestConfigsOptions.State))
	}

	body := make(map[string]interface{})
	if updatePullRequestConfigsOptions.PullRequest != nil {
		body["pull_request"] = updatePullRequestConfigsOptions.PullRequest
	}
	if updatePullRequestConfigsOptions.ProjectDefinition != nil {
		body["project_definition"] = updatePullRequestConfigsOptions.ProjectDefinition
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

// PlanPullRequestConfigs : Trigger a plan job for the added or updated configs in a given pull request
// Trigger a plan job for all newly added or updated configs found in a pull request. If a config is new, the config
// will be added to the status document with a reference to the pull request.
func (projects *ProjectsV1) PlanPullRequestConfigs(planPullRequestConfigsOptions *PlanPullRequestConfigsOptions) (result *PlanPullRequestConfigsResponse, response *core.DetailedResponse, err error) {
	return projects.PlanPullRequestConfigsWithContext(context.Background(), planPullRequestConfigsOptions)
}

// PlanPullRequestConfigsWithContext is an alternate form of the PlanPullRequestConfigs method which supports a Context parameter
func (projects *ProjectsV1) PlanPullRequestConfigsWithContext(ctx context.Context, planPullRequestConfigsOptions *PlanPullRequestConfigsOptions) (result *PlanPullRequestConfigsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(planPullRequestConfigsOptions, "planPullRequestConfigsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(planPullRequestConfigsOptions, "planPullRequestConfigsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *planPullRequestConfigsOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/pullrequest/configs/plan`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range planPullRequestConfigsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "PlanPullRequestConfigs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if planPullRequestConfigsOptions.PullRequest != nil {
		body["pull_request"] = planPullRequestConfigsOptions.PullRequest
	}
	if planPullRequestConfigsOptions.ProjectDefinition != nil {
		body["project_definition"] = planPullRequestConfigsOptions.ProjectDefinition
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPlanPullRequestConfigsResponse)
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
		"config_name": *planConfigOptions.ConfigName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{config_name}/plan`, pathParamsMap)
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

// ActivePR : Info about an active pull request (source branch and url).
type ActivePR struct {
	Branch *string `json:"branch,omitempty"`

	URL *string `json:"url,omitempty"`
}

// UnmarshalActivePR unmarshals an instance of ActivePR from the specified map of raw messages.
func UnmarshalActivePR(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActivePR)
	err = core.UnmarshalPrimitive(m, "branch", &obj.Branch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
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

// CheckProjectOptions : The CheckProject options.
type CheckProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The configs to check. Leave the array empty to check all the configs.
	ConfigNames []string `json:"config_names,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCheckProjectOptions : Instantiate CheckProjectOptions
func (*ProjectsV1) NewCheckProjectOptions(id string) *CheckProjectOptions {
	return &CheckProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *CheckProjectOptions) SetID(id string) *CheckProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigNames : Allow user to set ConfigNames
func (_options *CheckProjectOptions) SetConfigNames(configNames []string) *CheckProjectOptions {
	_options.ConfigNames = configNames
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CheckProjectOptions) SetHeaders(param map[string]string) *CheckProjectOptions {
	options.Headers = param
	return options
}

// ConfigChangesOptions : The ConfigChanges options.
type ConfigChangesOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	Source *ProjectPrototype `json:"source" validate:"required"`

	PullRequest *string `json:"pull_request,omitempty"`

	Target *ProjectPrototype `json:"target,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewConfigChangesOptions : Instantiate ConfigChangesOptions
func (*ProjectsV1) NewConfigChangesOptions(id string, source *ProjectPrototype) *ConfigChangesOptions {
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
func (_options *ConfigChangesOptions) SetSource(source *ProjectPrototype) *ConfigChangesOptions {
	_options.Source = source
	return _options
}

// SetPullRequest : Allow user to set PullRequest
func (_options *ConfigChangesOptions) SetPullRequest(pullRequest string) *ConfigChangesOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *ConfigChangesOptions) SetTarget(target *ProjectPrototype) *ConfigChangesOptions {
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

	// The error returned by fetching catalog.
	CatalogError *string `json:"catalog_error,omitempty"`

	// The error status code returned by fetching catalog.
	CatalogStatusCode *int64 `json:"catalog_status_code,omitempty"`

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
	err = core.UnmarshalPrimitive(m, "catalog_error", &obj.CatalogError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_status_code", &obj.CatalogStatusCode)
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

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`

	XIamApi *string `json:"X-Iam-Api,omitempty"`

	// Specifying the optional "return=representation" includes the original input payload in the successful response.
	Prefer *string `json:"prefer,omitempty"`

	ResourceGroup *string `json:"resource_group,omitempty"`

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
func (_options *CreateProjectOptions) SetConfigs(configs []ProjectConfigIntf) *CreateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetDashboard : Allow user to set Dashboard
func (_options *CreateProjectOptions) SetDashboard(dashboard *ProjectPrototypeDashboard) *CreateProjectOptions {
	_options.Dashboard = dashboard
	return _options
}

// SetXIamApi : Allow user to set XIamApi
func (_options *CreateProjectOptions) SetXIamApi(xIamApi string) *CreateProjectOptions {
	_options.XIamApi = core.StringPtr(xIamApi)
	return _options
}

// SetPrefer : Allow user to set Prefer
func (_options *CreateProjectOptions) SetPrefer(prefer string) *CreateProjectOptions {
	_options.Prefer = core.StringPtr(prefer)
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

// DeregisterPullRequestOptions : The DeregisterPullRequest options.
type DeregisterPullRequestOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The url of the PR, which uniquely identifies it.
	URL *string `json:"url" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeregisterPullRequestOptions : Instantiate DeregisterPullRequestOptions
func (*ProjectsV1) NewDeregisterPullRequestOptions(id string, url string) *DeregisterPullRequestOptions {
	return &DeregisterPullRequestOptions{
		ID: core.StringPtr(id),
		URL: core.StringPtr(url),
	}
}

// SetID : Allow user to set ID
func (_options *DeregisterPullRequestOptions) SetID(id string) *DeregisterPullRequestOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetURL : Allow user to set URL
func (_options *DeregisterPullRequestOptions) SetURL(url string) *DeregisterPullRequestOptions {
	_options.URL = core.StringPtr(url)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeregisterPullRequestOptions) SetHeaders(param map[string]string) *DeregisterPullRequestOptions {
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

// GetProjectComputedStatusOptions : The GetProjectComputedStatus options.
type GetProjectComputedStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the computed status.
	ComputedStatus *string `json:"computed_status" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectComputedStatusOptions : Instantiate GetProjectComputedStatusOptions
func (*ProjectsV1) NewGetProjectComputedStatusOptions(id string, computedStatus string) *GetProjectComputedStatusOptions {
	return &GetProjectComputedStatusOptions{
		ID: core.StringPtr(id),
		ComputedStatus: core.StringPtr(computedStatus),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectComputedStatusOptions) SetID(id string) *GetProjectComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetComputedStatus : Allow user to set ComputedStatus
func (_options *GetProjectComputedStatusOptions) SetComputedStatus(computedStatus string) *GetProjectComputedStatusOptions {
	_options.ComputedStatus = core.StringPtr(computedStatus)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectComputedStatusOptions) SetHeaders(param map[string]string) *GetProjectComputedStatusOptions {
	options.Headers = param
	return options
}

// GetProjectConfigComputedStatusOptions : The GetProjectConfigComputedStatus options.
type GetProjectConfigComputedStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config, which must be unique within the project.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// The name of the computed status, which must be unique within the config.
	ComputedStatus *string `json:"computed_status" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectConfigComputedStatusOptions : Instantiate GetProjectConfigComputedStatusOptions
func (*ProjectsV1) NewGetProjectConfigComputedStatusOptions(id string, configName string, computedStatus string) *GetProjectConfigComputedStatusOptions {
	return &GetProjectConfigComputedStatusOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
		ComputedStatus: core.StringPtr(computedStatus),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectConfigComputedStatusOptions) SetID(id string) *GetProjectConfigComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *GetProjectConfigComputedStatusOptions) SetConfigName(configName string) *GetProjectConfigComputedStatusOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetComputedStatus : Allow user to set ComputedStatus
func (_options *GetProjectConfigComputedStatusOptions) SetComputedStatus(computedStatus string) *GetProjectConfigComputedStatusOptions {
	_options.ComputedStatus = core.StringPtr(computedStatus)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectConfigComputedStatusOptions) SetHeaders(param map[string]string) *GetProjectConfigComputedStatusOptions {
	options.Headers = param
	return options
}

// GetProjectConfigStatusOptions : The GetProjectConfigStatus options.
type GetProjectConfigStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config, which must be unique within the project.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectConfigStatusOptions : Instantiate GetProjectConfigStatusOptions
func (*ProjectsV1) NewGetProjectConfigStatusOptions(id string, configName string) *GetProjectConfigStatusOptions {
	return &GetProjectConfigStatusOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectConfigStatusOptions) SetID(id string) *GetProjectConfigStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *GetProjectConfigStatusOptions) SetConfigName(configName string) *GetProjectConfigStatusOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectConfigStatusOptions) SetHeaders(param map[string]string) *GetProjectConfigStatusOptions {
	options.Headers = param
	return options
}

// GetProjectOptions : The GetProject options.
type GetProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Set this parameter if you want to get the project definition from a specific branch.
	Branch *string `json:"branch,omitempty"`

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

// SetBranch : Allow user to set Branch
func (_options *GetProjectOptions) SetBranch(branch string) *GetProjectOptions {
	_options.Branch = core.StringPtr(branch)
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

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *GetProjectResponseDashboard `json:"dashboard,omitempty"`
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
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dashboard", &obj.Dashboard, UnmarshalGetProjectResponseDashboard)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetProjectResponseDashboard : GetProjectResponseDashboard struct
type GetProjectResponseDashboard struct {
	Widgets []Widget `json:"widgets" validate:"required"`
}

// UnmarshalGetProjectResponseDashboard unmarshals an instance of GetProjectResponseDashboard from the specified map of raw messages.
func UnmarshalGetProjectResponseDashboard(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetProjectResponseDashboard)
	err = core.UnmarshalModel(m, "widgets", &obj.Widgets, UnmarshalWidget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetProjectStatusOptions : The GetProjectStatus options.
type GetProjectStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectStatusOptions : Instantiate GetProjectStatusOptions
func (*ProjectsV1) NewGetProjectStatusOptions(id string) *GetProjectStatusOptions {
	return &GetProjectStatusOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectStatusOptions) SetID(id string) *GetProjectStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetProjectStatusOptions) SetHeaders(param map[string]string) *GetProjectStatusOptions {
	options.Headers = param
	return options
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

// History : History struct
type History struct {
	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	Timestamp *strfmt.DateTime `json:"timestamp,omitempty"`

	Code *string `json:"code,omitempty"`

	Type *string `json:"type,omitempty"`
}

// Constants associated with the History.Type property.
const (
	History_Type_GitRepo = "git_repo"
	History_Type_Project = "project"
	History_Type_Schematics = "schematics"
	History_Type_Toolchain = "toolchain"
)

// UnmarshalHistory unmarshals an instance of History from the specified map of raw messages.
func UnmarshalHistory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(History)
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
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

// InitProjectConfigOptions : The InitProjectConfig options.
type InitProjectConfigOptions struct {
	// The version_locator of the offer. It is a dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id" validate:"required"`

	// The config name.
	ConfigName *string `json:"config_name" validate:"required"`

	// The config description.
	Description *string `json:"description,omitempty"`

	Labels []string `json:"labels,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInitProjectConfigOptions : Instantiate InitProjectConfigOptions
func (*ProjectsV1) NewInitProjectConfigOptions(locatorID string, configName string) *InitProjectConfigOptions {
	return &InitProjectConfigOptions{
		LocatorID: core.StringPtr(locatorID),
		ConfigName: core.StringPtr(configName),
	}
}

// SetLocatorID : Allow user to set LocatorID
func (_options *InitProjectConfigOptions) SetLocatorID(locatorID string) *InitProjectConfigOptions {
	_options.LocatorID = core.StringPtr(locatorID)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *InitProjectConfigOptions) SetConfigName(configName string) *InitProjectConfigOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *InitProjectConfigOptions) SetDescription(description string) *InitProjectConfigOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLabels : Allow user to set Labels
func (_options *InitProjectConfigOptions) SetLabels(labels []string) *InitProjectConfigOptions {
	_options.Labels = labels
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InitProjectConfigOptions) SetHeaders(param map[string]string) *InitProjectConfigOptions {
	options.Headers = param
	return options
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

// ListProjectConfigStatusesOptions : The ListProjectConfigStatuses options.
type ListProjectConfigStatusesOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectConfigStatusesOptions : Instantiate ListProjectConfigStatusesOptions
func (*ProjectsV1) NewListProjectConfigStatusesOptions(id string) *ListProjectConfigStatusesOptions {
	return &ListProjectConfigStatusesOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ListProjectConfigStatusesOptions) SetID(id string) *ListProjectConfigStatusesOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListProjectConfigStatusesOptions) SetHeaders(param map[string]string) *ListProjectConfigStatusesOptions {
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

// SetHeaders : Allow user to set Headers
func (options *ListProjectsOptions) SetHeaders(param map[string]string) *ListProjectsOptions {
	options.Headers = param
	return options
}

// ListProjectsResponse : Projects list.
type ListProjectsResponse struct {
	Limit *int64 `json:"limit" validate:"required"`

	// Get the occurrencies of the total Projects.
	TotalCount *int64 `json:"total_count" validate:"required"`

	First *PaginationLink `json:"first" validate:"required"`

	Last *PaginationLink `json:"last,omitempty"`

	Previous *PaginationLink `json:"previous,omitempty"`

	Next *PaginationLink `json:"next,omitempty"`

	// An array of projects.
	Projects []ProjectListResponse `json:"projects,omitempty"`
}

// UnmarshalListProjectsResponse unmarshals an instance of ListProjectsResponse from the specified map of raw messages.
func UnmarshalListProjectsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListProjectsResponse)
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
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProjectListResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListProjectsResponse) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// MergeProjectOptions : The MergeProject options.
type MergeProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`

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
func (_options *MergeProjectOptions) SetConfigs(configs []ProjectConfigIntf) *MergeProjectOptions {
	_options.Configs = configs
	return _options
}

// SetDashboard : Allow user to set Dashboard
func (_options *MergeProjectOptions) SetDashboard(dashboard *ProjectPrototypeDashboard) *MergeProjectOptions {
	_options.Dashboard = dashboard
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *MergeProjectOptions) SetHeaders(param map[string]string) *MergeProjectOptions {
	options.Headers = param
	return options
}

// NotifyOptions : The Notify options.
type NotifyOptions struct {
	// The project id.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewNotifyOptions : Instantiate NotifyOptions
func (*ProjectsV1) NewNotifyOptions(id string) *NotifyOptions {
	return &NotifyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *NotifyOptions) SetID(id string) *NotifyOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *NotifyOptions) SetHeaders(param map[string]string) *NotifyOptions {
	options.Headers = param
	return options
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
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config to trigger a plan.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPlanConfigOptions : Instantiate PlanConfigOptions
func (*ProjectsV1) NewPlanConfigOptions(id string, configName string) *PlanConfigOptions {
	return &PlanConfigOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
	}
}

// SetID : Allow user to set ID
func (_options *PlanConfigOptions) SetID(id string) *PlanConfigOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *PlanConfigOptions) SetConfigName(configName string) *PlanConfigOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PlanConfigOptions) SetHeaders(param map[string]string) *PlanConfigOptions {
	options.Headers = param
	return options
}

// PlanPullRequestConfigsOptions : The PlanPullRequestConfigs options.
type PlanPullRequestConfigsOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	PullRequest *string `json:"pull_request" validate:"required"`

	ProjectDefinition *ProjectPrototype `json:"project_definition" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPlanPullRequestConfigsOptions : Instantiate PlanPullRequestConfigsOptions
func (*ProjectsV1) NewPlanPullRequestConfigsOptions(id string, pullRequest string, projectDefinition *ProjectPrototype) *PlanPullRequestConfigsOptions {
	return &PlanPullRequestConfigsOptions{
		ID: core.StringPtr(id),
		PullRequest: core.StringPtr(pullRequest),
		ProjectDefinition: projectDefinition,
	}
}

// SetID : Allow user to set ID
func (_options *PlanPullRequestConfigsOptions) SetID(id string) *PlanPullRequestConfigsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetPullRequest : Allow user to set PullRequest
func (_options *PlanPullRequestConfigsOptions) SetPullRequest(pullRequest string) *PlanPullRequestConfigsOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
	return _options
}

// SetProjectDefinition : Allow user to set ProjectDefinition
func (_options *PlanPullRequestConfigsOptions) SetProjectDefinition(projectDefinition *ProjectPrototype) *PlanPullRequestConfigsOptions {
	_options.ProjectDefinition = projectDefinition
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *PlanPullRequestConfigsOptions) SetHeaders(param map[string]string) *PlanPullRequestConfigsOptions {
	options.Headers = param
	return options
}

// PlanPullRequestConfigsResponse : PlanPullRequestConfigsResponse struct
type PlanPullRequestConfigsResponse struct {
	Configs []ConfigJobResponse `json:"configs,omitempty"`
}

// UnmarshalPlanPullRequestConfigsResponse unmarshals an instance of PlanPullRequestConfigsResponse from the specified map of raw messages.
func UnmarshalPlanPullRequestConfigsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PlanPullRequestConfigsResponse)
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalConfigJobResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectComputedStatus : The project computed statuses.
type ProjectComputedStatus struct {
	ProjectID *string `json:"project_id" validate:"required"`

	Href *string `json:"href" validate:"required"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses" validate:"required"`
}

// UnmarshalProjectComputedStatus unmarshals an instance of ProjectComputedStatus from the specified map of raw messages.
func UnmarshalProjectComputedStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectComputedStatus)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computed_statuses", &obj.ComputedStatuses)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfig : ProjectConfig struct
// Models which "extend" this model:
// - ProjectConfigManualProperty
// - ProjectConfigTerraformTemplateProperty
// - ProjectConfigSchematicsBlueprintProperty
type ProjectConfig struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type,omitempty"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	Input []InputVariable `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`

	// A Terraform blueprint to use for provisioning a set of project resources.
	Template *TerraformTemplate `json:"template,omitempty"`

	// A Schematics blueprint to use for provisioning a set of project resources.
	Blueprint *SchematicsBlueprint `json:"blueprint,omitempty"`
}

// Constants associated with the ProjectConfig.Type property.
const (
	ProjectConfig_Type_Manual = "manual"
)
func (*ProjectConfig) isaProjectConfig() bool {
	return true
}

type ProjectConfigIntf interface {
	isaProjectConfig() bool
}

// UnmarshalProjectConfig unmarshals an instance of ProjectConfig from the specified map of raw messages.
func UnmarshalProjectConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfig)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "template", &obj.Template, UnmarshalTerraformTemplate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "blueprint", &obj.Blueprint, UnmarshalSchematicsBlueprint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigChangesResponse : ProjectConfigChangesResponse struct
type ProjectConfigChangesResponse struct {
	Added []string `json:"added" validate:"required"`

	Deleted []string `json:"deleted" validate:"required"`

	Changed []string `json:"changed" validate:"required"`
}

// UnmarshalProjectConfigChangesResponse unmarshals an instance of ProjectConfigChangesResponse from the specified map of raw messages.
func UnmarshalProjectConfigChangesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigChangesResponse)
	err = core.UnmarshalPrimitive(m, "added", &obj.Added)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deleted", &obj.Deleted)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "changed", &obj.Changed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigComputedStatus : ProjectConfigComputedStatus struct
type ProjectConfigComputedStatus struct {
	Href *string `json:"href" validate:"required"`

	Name *string `json:"name" validate:"required"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`
}

// UnmarshalProjectConfigComputedStatus unmarshals an instance of ProjectConfigComputedStatus from the specified map of raw messages.
func UnmarshalProjectConfigComputedStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigComputedStatus)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computed_statuses", &obj.ComputedStatuses)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigStatus : ProjectConfigStatus struct
type ProjectConfigStatus struct {
	Href *string `json:"href" validate:"required"`

	Name *string `json:"name" validate:"required"`

	Status *string `json:"status" validate:"required"`

	Messages []string `json:"messages" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	SchematicsResourceID *string `json:"schematics_resource_id,omitempty"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	Output []OutputValue `json:"output,omitempty"`
}

// Constants associated with the ProjectConfigStatus.Status property.
const (
	ProjectConfigStatus_Status_Installing = "INSTALLING"
	ProjectConfigStatus_Status_InstallingFailed = "INSTALLING_FAILED"
	ProjectConfigStatus_Status_Ready = "READY"
	ProjectConfigStatus_Status_Updating = "UPDATING"
	ProjectConfigStatus_Status_UpdatingFailed = "UPDATING_FAILED"
)

// UnmarshalProjectConfigStatus unmarshals an instance of ProjectConfigStatus from the specified map of raw messages.
func UnmarshalProjectConfigStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigStatus)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pipeline_run", &obj.PipelineRun)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schematics_resource_id", &obj.SchematicsResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computed_statuses", &obj.ComputedStatuses)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigStatuses : The status of all project configs.
type ProjectConfigStatuses struct {
	ProjectID *string `json:"project_id" validate:"required"`

	Href *string `json:"href" validate:"required"`

	ConfigStatuses []ProjectConfigStatus `json:"config_statuses" validate:"required"`
}

// UnmarshalProjectConfigStatuses unmarshals an instance of ProjectConfigStatuses from the specified map of raw messages.
func UnmarshalProjectConfigStatuses(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigStatuses)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config_statuses", &obj.ConfigStatuses, UnmarshalProjectConfigStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectListResponse : ProjectListResponse struct
type ProjectListResponse struct {
	// The project name.
	Name *string `json:"name,omitempty"`

	ID *string `json:"id,omitempty"`

	// The branch name of the PR that is created.
	PrBranch *string `json:"pr_branch,omitempty"`

	Definition *GetProjectResponse `json:"definition,omitempty"`

	State *string `json:"state,omitempty"`
}

// Constants associated with the ProjectListResponse.State property.
const (
	ProjectListResponse_State_Creating = "CREATING"
	ProjectListResponse_State_CreatingFailed = "CREATING_FAILED"
	ProjectListResponse_State_Ready = "READY"
	ProjectListResponse_State_Updating = "UPDATING"
	ProjectListResponse_State_UpdatingFailed = "UPDATING_FAILED"
)

// UnmarshalProjectListResponse unmarshals an instance of ProjectListResponse from the specified map of raw messages.
func UnmarshalProjectListResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectListResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pr_branch", &obj.PrBranch)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "definition", &obj.Definition, UnmarshalGetProjectResponse)
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

// ProjectPrototype : ProjectPrototype struct
type ProjectPrototype struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`
}

// NewProjectPrototype : Instantiate ProjectPrototype (Generic Model Constructor)
func (*ProjectsV1) NewProjectPrototype(name string) (_model *ProjectPrototype, err error) {
	_model = &ProjectPrototype{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectPrototype unmarshals an instance of ProjectPrototype from the specified map of raw messages.
func UnmarshalProjectPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectPrototype)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dashboard", &obj.Dashboard, UnmarshalProjectPrototypeDashboard)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectPrototypeDashboard : ProjectPrototypeDashboard struct
type ProjectPrototypeDashboard struct {
	Widgets []Widget `json:"widgets" validate:"required"`
}

// NewProjectPrototypeDashboard : Instantiate ProjectPrototypeDashboard (Generic Model Constructor)
func (*ProjectsV1) NewProjectPrototypeDashboard(widgets []Widget) (_model *ProjectPrototypeDashboard, err error) {
	_model = &ProjectPrototypeDashboard{
		Widgets: widgets,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectPrototypeDashboard unmarshals an instance of ProjectPrototypeDashboard from the specified map of raw messages.
func UnmarshalProjectPrototypeDashboard(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectPrototypeDashboard)
	err = core.UnmarshalModel(m, "widgets", &obj.Widgets, UnmarshalWidget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectResponse : ProjectResponse struct
type ProjectResponse struct {
	// The project name.
	Name *string `json:"name,omitempty"`

	ID *string `json:"id,omitempty"`

	// The branch name of the PR that is created.
	PrBranch *string `json:"pr_branch,omitempty"`

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
	err = core.UnmarshalPrimitive(m, "pr_branch", &obj.PrBranch)
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

// ProjectStatus : The project status, including plumbing services and computed statuses information.
type ProjectStatus struct {
	ProjectID *string `json:"project_id" validate:"required"`

	Href *string `json:"href,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	ProjectCrn *string `json:"project_crn" validate:"required"`

	// The project name.
	ProjectName *string `json:"project_name" validate:"required"`

	Location *string `json:"location" validate:"required"`

	// The resource group id (or Default for the default resource group).
	ResourceGroup *string `json:"resource_group" validate:"required"`

	State *string `json:"state" validate:"required"`

	// Project plumbing service info.
	GitRepo *ServiceInfoGit `json:"git_repo,omitempty"`

	// Project plumbing service info.
	Toolchain *ServiceInfoToolchain `json:"toolchain,omitempty"`

	// Project plumbing service info.
	Schematics *ServiceInfoSchematics `json:"schematics,omitempty"`

	Credentials *ProjectStatusCredentials `json:"credentials,omitempty"`

	Configs []StatusConfigs `json:"configs,omitempty"`

	Dashboard *ProjectStatusDashboard `json:"dashboard,omitempty"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	ActivePrs []ActivePR `json:"active_prs,omitempty"`

	History []History `json:"history,omitempty"`
}

// Constants associated with the ProjectStatus.Location property.
const (
	ProjectStatus_Location_AuSyd = "au-syd"
	ProjectStatus_Location_EuDe = "eu-de"
	ProjectStatus_Location_EuGb = "eu-gb"
	ProjectStatus_Location_UsEast = "us-east"
	ProjectStatus_Location_UsSouth = "us-south"
)

// Constants associated with the ProjectStatus.State property.
const (
	ProjectStatus_State_Creating = "CREATING"
	ProjectStatus_State_CreatingFailed = "CREATING_FAILED"
	ProjectStatus_State_Ready = "READY"
	ProjectStatus_State_Updating = "UPDATING"
	ProjectStatus_State_UpdatingFailed = "UPDATING_FAILED"
)

// UnmarshalProjectStatus unmarshals an instance of ProjectStatus from the specified map of raw messages.
func UnmarshalProjectStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectStatus)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_crn", &obj.ProjectCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_name", &obj.ProjectName)
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
	err = core.UnmarshalModel(m, "git_repo", &obj.GitRepo, UnmarshalServiceInfoGit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "toolchain", &obj.Toolchain, UnmarshalServiceInfoToolchain)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schematics", &obj.Schematics, UnmarshalServiceInfoSchematics)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "credentials", &obj.Credentials, UnmarshalProjectStatusCredentials)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalStatusConfigs)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dashboard", &obj.Dashboard, UnmarshalProjectStatusDashboard)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computed_statuses", &obj.ComputedStatuses)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "active_prs", &obj.ActivePrs, UnmarshalActivePR)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "history", &obj.History, UnmarshalHistory)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectStatusCredentials : ProjectStatusCredentials struct
type ProjectStatusCredentials struct {
	// A valid IAM api key with the permissions to manage the plumbing services (Schematics and Toolchain).
	ApiKeyRef *string `json:"api_key_ref,omitempty"`
}

// UnmarshalProjectStatusCredentials unmarshals an instance of ProjectStatusCredentials from the specified map of raw messages.
func UnmarshalProjectStatusCredentials(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectStatusCredentials)
	err = core.UnmarshalPrimitive(m, "api_key_ref", &obj.ApiKeyRef)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectStatusDashboard : ProjectStatusDashboard struct
type ProjectStatusDashboard struct {
	Widgets []Widget `json:"widgets" validate:"required"`
}

// UnmarshalProjectStatusDashboard unmarshals an instance of ProjectStatusDashboard from the specified map of raw messages.
func UnmarshalProjectStatusDashboard(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectStatusDashboard)
	err = core.UnmarshalModel(m, "widgets", &obj.Widgets, UnmarshalWidget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RegisterPullRequestOptions : The RegisterPullRequest options.
type RegisterPullRequestOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the branch.
	Branch *string `json:"branch,omitempty"`

	URL *string `json:"url,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRegisterPullRequestOptions : Instantiate RegisterPullRequestOptions
func (*ProjectsV1) NewRegisterPullRequestOptions(id string) *RegisterPullRequestOptions {
	return &RegisterPullRequestOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *RegisterPullRequestOptions) SetID(id string) *RegisterPullRequestOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetBranch : Allow user to set Branch
func (_options *RegisterPullRequestOptions) SetBranch(branch string) *RegisterPullRequestOptions {
	_options.Branch = core.StringPtr(branch)
	return _options
}

// SetURL : Allow user to set URL
func (_options *RegisterPullRequestOptions) SetURL(url string) *RegisterPullRequestOptions {
	_options.URL = core.StringPtr(url)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterPullRequestOptions) SetHeaders(param map[string]string) *RegisterPullRequestOptions {
	options.Headers = param
	return options
}

// SchematicsBlueprint : A Schematics blueprint to use for provisioning a set of project resources.
type SchematicsBlueprint struct {
	RepoURL *string `json:"repo_url" validate:"required"`

	CatalogRef *string `json:"catalog_ref,omitempty"`

	DefinitionFile *string `json:"definition_file" validate:"required"`

	PersonalAccessToken *string `json:"personal_access_token,omitempty"`
}

// NewSchematicsBlueprint : Instantiate SchematicsBlueprint (Generic Model Constructor)
func (*ProjectsV1) NewSchematicsBlueprint(repoURL string, definitionFile string) (_model *SchematicsBlueprint, err error) {
	_model = &SchematicsBlueprint{
		RepoURL: core.StringPtr(repoURL),
		DefinitionFile: core.StringPtr(definitionFile),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSchematicsBlueprint unmarshals an instance of SchematicsBlueprint from the specified map of raw messages.
func UnmarshalSchematicsBlueprint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsBlueprint)
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_ref", &obj.CatalogRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "definition_file", &obj.DefinitionFile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "personal_access_token", &obj.PersonalAccessToken)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceInfoGit : Project plumbing service info.
type ServiceInfoGit struct {
	URL *string `json:"url,omitempty"`

	Branch *string `json:"branch,omitempty"`

	ProjectID *string `json:"project_id,omitempty"`
}

// UnmarshalServiceInfoGit unmarshals an instance of ServiceInfoGit from the specified map of raw messages.
func UnmarshalServiceInfoGit(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceInfoGit)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "branch", &obj.Branch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceInfoSchematics : Project plumbing service info.
type ServiceInfoSchematics struct {
	CartOrderID *string `json:"cart_order_id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	CartItemName *string `json:"cart_item_name,omitempty"`
}

// UnmarshalServiceInfoSchematics unmarshals an instance of ServiceInfoSchematics from the specified map of raw messages.
func UnmarshalServiceInfoSchematics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceInfoSchematics)
	err = core.UnmarshalPrimitive(m, "cart_order_id", &obj.CartOrderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cart_item_name", &obj.CartItemName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceInfoToolchain : Project plumbing service info.
type ServiceInfoToolchain struct {
	ID *string `json:"id,omitempty"`
}

// UnmarshalServiceInfoToolchain unmarshals an instance of ServiceInfoToolchain from the specified map of raw messages.
func UnmarshalServiceInfoToolchain(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceInfoToolchain)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// StatusConfigs : StatusConfigs struct
type StatusConfigs struct {
	Name *string `json:"name" validate:"required"`

	PullRequest *string `json:"pull_request,omitempty"`

	CatalogID *string `json:"catalog_id,omitempty"`

	OfferingID *string `json:"offering_id,omitempty"`

	OfferingKindID *string `json:"offering_kind_id,omitempty"`

	VersionID *string `json:"version_id,omitempty"`

	// Version of the catalog offer.
	OfferingVersion *string `json:"offering_version,omitempty"`

	OfferingFulfilmentKind *string `json:"offering_fulfilment_kind,omitempty"`

	CartOrderID *string `json:"cart_order_id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	CartItemName *string `json:"cart_item_name,omitempty"`
}

// Constants associated with the StatusConfigs.OfferingFulfilmentKind property.
const (
	StatusConfigs_OfferingFulfilmentKind_Terraform = "terraform"
)

// UnmarshalStatusConfigs unmarshals an instance of StatusConfigs from the specified map of raw messages.
func UnmarshalStatusConfigs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(StatusConfigs)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pull_request", &obj.PullRequest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_id", &obj.OfferingID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_kind_id", &obj.OfferingKindID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_id", &obj.VersionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_version", &obj.OfferingVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_fulfilment_kind", &obj.OfferingFulfilmentKind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cart_order_id", &obj.CartOrderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cart_item_name", &obj.CartItemName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TerraformTemplate : A Terraform blueprint to use for provisioning a set of project resources.
type TerraformTemplate struct {
	// The version_locator of the offer. It is a dotted value of catalogID.versionID.
	LocatorID *string `json:"locator_id" validate:"required"`
}

// NewTerraformTemplate : Instantiate TerraformTemplate (Generic Model Constructor)
func (*ProjectsV1) NewTerraformTemplate(locatorID string) (_model *TerraformTemplate, err error) {
	_model = &TerraformTemplate{
		LocatorID: core.StringPtr(locatorID),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTerraformTemplate unmarshals an instance of TerraformTemplate from the specified map of raw messages.
func UnmarshalTerraformTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TerraformTemplate)
	err = core.UnmarshalPrimitive(m, "locator_id", &obj.LocatorID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
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

// UpdateProjectComputedStatusOptions : The UpdateProjectComputedStatus options.
type UpdateProjectComputedStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the computed status.
	ComputedStatus *string `json:"computed_status" validate:"required,ne="`

	ComputedStatuses map[string]interface{} `json:"computed_statuses" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectComputedStatusOptions : Instantiate UpdateProjectComputedStatusOptions
func (*ProjectsV1) NewUpdateProjectComputedStatusOptions(id string, computedStatus string, computedStatuses map[string]interface{}) *UpdateProjectComputedStatusOptions {
	return &UpdateProjectComputedStatusOptions{
		ID: core.StringPtr(id),
		ComputedStatus: core.StringPtr(computedStatus),
		ComputedStatuses: computedStatuses,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectComputedStatusOptions) SetID(id string) *UpdateProjectComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetComputedStatus : Allow user to set ComputedStatus
func (_options *UpdateProjectComputedStatusOptions) SetComputedStatus(computedStatus string) *UpdateProjectComputedStatusOptions {
	_options.ComputedStatus = core.StringPtr(computedStatus)
	return _options
}

// SetComputedStatuses : Allow user to set ComputedStatuses
func (_options *UpdateProjectComputedStatusOptions) SetComputedStatuses(computedStatuses map[string]interface{}) *UpdateProjectComputedStatusOptions {
	_options.ComputedStatuses = computedStatuses
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectComputedStatusOptions) SetHeaders(param map[string]string) *UpdateProjectComputedStatusOptions {
	options.Headers = param
	return options
}

// UpdateProjectConfigComputedStatusOptions : The UpdateProjectConfigComputedStatus options.
type UpdateProjectConfigComputedStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config, which must be unique in the project.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	// The name of the computed status, which must be unique within the config.
	ComputedStatus *string `json:"computed_status" validate:"required,ne="`

	// The computed status to set.
	RequestBody map[string]interface{} `json:"request_body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectConfigComputedStatusOptions : Instantiate UpdateProjectConfigComputedStatusOptions
func (*ProjectsV1) NewUpdateProjectConfigComputedStatusOptions(id string, configName string, computedStatus string, requestBody map[string]interface{}) *UpdateProjectConfigComputedStatusOptions {
	return &UpdateProjectConfigComputedStatusOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
		ComputedStatus: core.StringPtr(computedStatus),
		RequestBody: requestBody,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectConfigComputedStatusOptions) SetID(id string) *UpdateProjectConfigComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *UpdateProjectConfigComputedStatusOptions) SetConfigName(configName string) *UpdateProjectConfigComputedStatusOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetComputedStatus : Allow user to set ComputedStatus
func (_options *UpdateProjectConfigComputedStatusOptions) SetComputedStatus(computedStatus string) *UpdateProjectConfigComputedStatusOptions {
	_options.ComputedStatus = core.StringPtr(computedStatus)
	return _options
}

// SetRequestBody : Allow user to set RequestBody
func (_options *UpdateProjectConfigComputedStatusOptions) SetRequestBody(requestBody map[string]interface{}) *UpdateProjectConfigComputedStatusOptions {
	_options.RequestBody = requestBody
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectConfigComputedStatusOptions) SetHeaders(param map[string]string) *UpdateProjectConfigComputedStatusOptions {
	options.Headers = param
	return options
}

// UpdateProjectConfigStatusOptions : The UpdateProjectConfigStatus options.
type UpdateProjectConfigStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config, which must be unique within the project.
	ConfigName *string `json:"config_name" validate:"required,ne="`

	Status *string `json:"status" validate:"required"`

	Messages []string `json:"messages" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	SchematicsResourceID *string `json:"schematics_resource_id,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateProjectConfigStatusOptions.Status property.
const (
	UpdateProjectConfigStatusOptions_Status_Installing = "INSTALLING"
	UpdateProjectConfigStatusOptions_Status_InstallingFailed = "INSTALLING_FAILED"
	UpdateProjectConfigStatusOptions_Status_Ready = "READY"
	UpdateProjectConfigStatusOptions_Status_Updating = "UPDATING"
	UpdateProjectConfigStatusOptions_Status_UpdatingFailed = "UPDATING_FAILED"
)

// NewUpdateProjectConfigStatusOptions : Instantiate UpdateProjectConfigStatusOptions
func (*ProjectsV1) NewUpdateProjectConfigStatusOptions(id string, configName string, status string, messages []string) *UpdateProjectConfigStatusOptions {
	return &UpdateProjectConfigStatusOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
		Status: core.StringPtr(status),
		Messages: messages,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectConfigStatusOptions) SetID(id string) *UpdateProjectConfigStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigName : Allow user to set ConfigName
func (_options *UpdateProjectConfigStatusOptions) SetConfigName(configName string) *UpdateProjectConfigStatusOptions {
	_options.ConfigName = core.StringPtr(configName)
	return _options
}

// SetStatus : Allow user to set Status
func (_options *UpdateProjectConfigStatusOptions) SetStatus(status string) *UpdateProjectConfigStatusOptions {
	_options.Status = core.StringPtr(status)
	return _options
}

// SetMessages : Allow user to set Messages
func (_options *UpdateProjectConfigStatusOptions) SetMessages(messages []string) *UpdateProjectConfigStatusOptions {
	_options.Messages = messages
	return _options
}

// SetPipelineRun : Allow user to set PipelineRun
func (_options *UpdateProjectConfigStatusOptions) SetPipelineRun(pipelineRun string) *UpdateProjectConfigStatusOptions {
	_options.PipelineRun = core.StringPtr(pipelineRun)
	return _options
}

// SetSchematicsResourceID : Allow user to set SchematicsResourceID
func (_options *UpdateProjectConfigStatusOptions) SetSchematicsResourceID(schematicsResourceID string) *UpdateProjectConfigStatusOptions {
	_options.SchematicsResourceID = core.StringPtr(schematicsResourceID)
	return _options
}

// SetOutput : Allow user to set Output
func (_options *UpdateProjectConfigStatusOptions) SetOutput(output []OutputValue) *UpdateProjectConfigStatusOptions {
	_options.Output = output
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectConfigStatusOptions) SetHeaders(param map[string]string) *UpdateProjectConfigStatusOptions {
	options.Headers = param
	return options
}

// UpdateProjectOptions : The UpdateProject options.
type UpdateProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`

	// Set this parameter if you want to update the project definition of a specific branch.
	Branch *string `json:"branch,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectsV1) NewUpdateProjectOptions(id string, name string) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
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

// SetConfigs : Allow user to set Configs
func (_options *UpdateProjectOptions) SetConfigs(configs []ProjectConfigIntf) *UpdateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetDashboard : Allow user to set Dashboard
func (_options *UpdateProjectOptions) SetDashboard(dashboard *ProjectPrototypeDashboard) *UpdateProjectOptions {
	_options.Dashboard = dashboard
	return _options
}

// SetBranch : Allow user to set Branch
func (_options *UpdateProjectOptions) SetBranch(branch string) *UpdateProjectOptions {
	_options.Branch = core.StringPtr(branch)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectOptions) SetHeaders(param map[string]string) *UpdateProjectOptions {
	options.Headers = param
	return options
}

// UpdateProjectStatusOptions : The UpdateProjectStatus options.
type UpdateProjectStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	State *string `json:"state" validate:"required"`

	History []History `json:"history" validate:"required"`

	// Project plumbing service info.
	GitRepo *ServiceInfoGit `json:"git_repo,omitempty"`

	// Project plumbing service info.
	Toolchain *ServiceInfoToolchain `json:"toolchain,omitempty"`

	// Project plumbing service info.
	Schematics *ServiceInfoSchematics `json:"schematics,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateProjectStatusOptions.State property.
const (
	UpdateProjectStatusOptions_State_Creating = "CREATING"
	UpdateProjectStatusOptions_State_CreatingFailed = "CREATING_FAILED"
	UpdateProjectStatusOptions_State_Ready = "READY"
	UpdateProjectStatusOptions_State_Updating = "UPDATING"
	UpdateProjectStatusOptions_State_UpdatingFailed = "UPDATING_FAILED"
)

// NewUpdateProjectStatusOptions : Instantiate UpdateProjectStatusOptions
func (*ProjectsV1) NewUpdateProjectStatusOptions(id string, state string, history []History) *UpdateProjectStatusOptions {
	return &UpdateProjectStatusOptions{
		ID: core.StringPtr(id),
		State: core.StringPtr(state),
		History: history,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectStatusOptions) SetID(id string) *UpdateProjectStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetState : Allow user to set State
func (_options *UpdateProjectStatusOptions) SetState(state string) *UpdateProjectStatusOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHistory : Allow user to set History
func (_options *UpdateProjectStatusOptions) SetHistory(history []History) *UpdateProjectStatusOptions {
	_options.History = history
	return _options
}

// SetGitRepo : Allow user to set GitRepo
func (_options *UpdateProjectStatusOptions) SetGitRepo(gitRepo *ServiceInfoGit) *UpdateProjectStatusOptions {
	_options.GitRepo = gitRepo
	return _options
}

// SetToolchain : Allow user to set Toolchain
func (_options *UpdateProjectStatusOptions) SetToolchain(toolchain *ServiceInfoToolchain) *UpdateProjectStatusOptions {
	_options.Toolchain = toolchain
	return _options
}

// SetSchematics : Allow user to set Schematics
func (_options *UpdateProjectStatusOptions) SetSchematics(schematics *ServiceInfoSchematics) *UpdateProjectStatusOptions {
	_options.Schematics = schematics
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectStatusOptions) SetHeaders(param map[string]string) *UpdateProjectStatusOptions {
	options.Headers = param
	return options
}

// UpdatePullRequestConfigsOptions : The UpdatePullRequestConfigs options.
type UpdatePullRequestConfigsOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	PullRequest *string `json:"pull_request" validate:"required"`

	ProjectDefinition *ProjectPrototype `json:"project_definition,omitempty"`

	// The state of the PR.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdatePullRequestConfigsOptions.State property.
// The state of the PR.
const (
	UpdatePullRequestConfigsOptions_State_Close = "close"
	UpdatePullRequestConfigsOptions_State_Merge = "merge"
)

// NewUpdatePullRequestConfigsOptions : Instantiate UpdatePullRequestConfigsOptions
func (*ProjectsV1) NewUpdatePullRequestConfigsOptions(id string, pullRequest string) *UpdatePullRequestConfigsOptions {
	return &UpdatePullRequestConfigsOptions{
		ID: core.StringPtr(id),
		PullRequest: core.StringPtr(pullRequest),
	}
}

// SetID : Allow user to set ID
func (_options *UpdatePullRequestConfigsOptions) SetID(id string) *UpdatePullRequestConfigsOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetPullRequest : Allow user to set PullRequest
func (_options *UpdatePullRequestConfigsOptions) SetPullRequest(pullRequest string) *UpdatePullRequestConfigsOptions {
	_options.PullRequest = core.StringPtr(pullRequest)
	return _options
}

// SetProjectDefinition : Allow user to set ProjectDefinition
func (_options *UpdatePullRequestConfigsOptions) SetProjectDefinition(projectDefinition *ProjectPrototype) *UpdatePullRequestConfigsOptions {
	_options.ProjectDefinition = projectDefinition
	return _options
}

// SetState : Allow user to set State
func (_options *UpdatePullRequestConfigsOptions) SetState(state string) *UpdatePullRequestConfigsOptions {
	_options.State = core.StringPtr(state)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePullRequestConfigsOptions) SetHeaders(param map[string]string) *UpdatePullRequestConfigsOptions {
	options.Headers = param
	return options
}

// UpdateResponsePullRequest : UpdateResponsePullRequest struct
type UpdateResponsePullRequest struct {
	// The name of the branch.
	Branch *string `json:"branch,omitempty"`

	PrURL *string `json:"pr_url,omitempty"`
}

// UnmarshalUpdateResponsePullRequest unmarshals an instance of UpdateResponsePullRequest from the specified map of raw messages.
func UnmarshalUpdateResponsePullRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateResponsePullRequest)
	err = core.UnmarshalPrimitive(m, "branch", &obj.Branch)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pr_url", &obj.PrURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ValidateProjectOptions : The ValidateProject options.
type ValidateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`

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
func (_options *ValidateProjectOptions) SetConfigs(configs []ProjectConfigIntf) *ValidateProjectOptions {
	_options.Configs = configs
	return _options
}

// SetDashboard : Allow user to set Dashboard
func (_options *ValidateProjectOptions) SetDashboard(dashboard *ProjectPrototypeDashboard) *ValidateProjectOptions {
	_options.Dashboard = dashboard
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ValidateProjectOptions) SetHeaders(param map[string]string) *ValidateProjectOptions {
	options.Headers = param
	return options
}

// Widget : Widget struct
type Widget struct {
	// The name of the widget, which must be unique within the project.
	Name *string `json:"name" validate:"required"`
}

// NewWidget : Instantiate Widget (Generic Model Constructor)
func (*ProjectsV1) NewWidget(name string) (_model *Widget, err error) {
	_model = &Widget{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalWidget unmarshals an instance of Widget from the specified map of raw messages.
func UnmarshalWidget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Widget)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigManualProperty : ProjectConfigManualProperty struct
// This model "extends" ProjectConfig
type ProjectConfigManualProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type" validate:"required"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`
}

// Constants associated with the ProjectConfigManualProperty.Type property.
const (
	ProjectConfigManualProperty_Type_Manual = "manual"
)

// NewProjectConfigManualProperty : Instantiate ProjectConfigManualProperty (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigManualProperty(name string, typeVar string) (_model *ProjectConfigManualProperty, err error) {
	_model = &ProjectConfigManualProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ProjectConfigManualProperty) isaProjectConfig() bool {
	return true
}

// UnmarshalProjectConfigManualProperty unmarshals an instance of ProjectConfigManualProperty from the specified map of raw messages.
func UnmarshalProjectConfigManualProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigManualProperty)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
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

// ProjectConfigSchematicsBlueprintProperty : ProjectConfigSchematicsBlueprintProperty struct
// This model "extends" ProjectConfig
type ProjectConfigSchematicsBlueprintProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input,omitempty"`

	// A Schematics blueprint to use for provisioning a set of project resources.
	Blueprint *SchematicsBlueprint `json:"blueprint" validate:"required"`
}

// Constants associated with the ProjectConfigSchematicsBlueprintProperty.Type property.
const (
	ProjectConfigSchematicsBlueprintProperty_Type_SchematicsBlueprint = "schematics_blueprint"
)

// NewProjectConfigSchematicsBlueprintProperty : Instantiate ProjectConfigSchematicsBlueprintProperty (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigSchematicsBlueprintProperty(name string, typeVar string, blueprint *SchematicsBlueprint) (_model *ProjectConfigSchematicsBlueprintProperty, err error) {
	_model = &ProjectConfigSchematicsBlueprintProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Blueprint: blueprint,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ProjectConfigSchematicsBlueprintProperty) isaProjectConfig() bool {
	return true
}

// UnmarshalProjectConfigSchematicsBlueprintProperty unmarshals an instance of ProjectConfigSchematicsBlueprintProperty from the specified map of raw messages.
func UnmarshalProjectConfigSchematicsBlueprintProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigSchematicsBlueprintProperty)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "blueprint", &obj.Blueprint, UnmarshalSchematicsBlueprint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectConfigTerraformTemplateProperty : ProjectConfigTerraformTemplateProperty struct
// This model "extends" ProjectConfig
type ProjectConfigTerraformTemplateProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// A project config description.
	Description *string `json:"description,omitempty"`

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input,omitempty"`

	// Optional setting object we can pass to the cart api.
	Setting []ConfigSettingItem `json:"setting,omitempty"`

	// A Terraform blueprint to use for provisioning a set of project resources.
	Template *TerraformTemplate `json:"template" validate:"required"`
}

// Constants associated with the ProjectConfigTerraformTemplateProperty.Type property.
const (
	ProjectConfigTerraformTemplateProperty_Type_TerraformTemplate = "terraform_template"
)

// NewProjectConfigTerraformTemplateProperty : Instantiate ProjectConfigTerraformTemplateProperty (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigTerraformTemplateProperty(name string, typeVar string, template *TerraformTemplate) (_model *ProjectConfigTerraformTemplateProperty, err error) {
	_model = &ProjectConfigTerraformTemplateProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Template: template,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ProjectConfigTerraformTemplateProperty) isaProjectConfig() bool {
	return true
}

// UnmarshalProjectConfigTerraformTemplateProperty unmarshals an instance of ProjectConfigTerraformTemplateProperty from the specified map of raw messages.
func UnmarshalProjectConfigTerraformTemplateProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectConfigTerraformTemplateProperty)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "output", &obj.Output, UnmarshalOutputValue)
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
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "setting", &obj.Setting, UnmarshalConfigSettingItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "template", &obj.Template, UnmarshalTerraformTemplate)
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
func (pager *ProjectsPager) GetNextWithContext(ctx context.Context) (page []ProjectListResponse, err error) {
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
func (pager *ProjectsPager) GetAllWithContext(ctx context.Context) (allItems []ProjectListResponse, err error) {
	for pager.HasNext() {
		var nextPage []ProjectListResponse
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetNext() (page []ProjectListResponse, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *ProjectsPager) GetAll() (allItems []ProjectListResponse, err error) {
	return pager.GetAllWithContext(context.Background())
}
