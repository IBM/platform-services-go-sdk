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
 * IBM OpenAPI SDK Code Generator Version: 3.48.0-e80b60a1-20220414-145125
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
const DefaultServiceURL = "http://localhost:9989"

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
func (projects *ProjectsV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
	return projects.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (projects *ProjectsV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
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

	body := make(map[string]interface{})
	if createProjectOptions.Name != nil {
		body["name"] = createProjectOptions.Name
	}
	if createProjectOptions.ApiKey != nil {
		body["api_key"] = createProjectOptions.ApiKey
	}
	if createProjectOptions.Description != nil {
		body["description"] = createProjectOptions.Description
	}
	if createProjectOptions.RepoURL != nil {
		body["repo_url"] = createProjectOptions.RepoURL
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProject)
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
func (projects *ProjectsV1) GetProject(getProjectOptions *GetProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
	return projects.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (projects *ProjectsV1) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProject)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProject : Update a project by id
// Update a project.
func (projects *ProjectsV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *Project, response *core.DetailedResponse, err error) {
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
	if updateProjectOptions.Prefer != nil {
		builder.AddHeader("Prefer", fmt.Sprint(*updateProjectOptions.Prefer))
	}

	body := make(map[string]interface{})
	if updateProjectOptions.NewName != nil {
		body["name"] = updateProjectOptions.NewName
	}
	if updateProjectOptions.NewDescription != nil {
		body["description"] = updateProjectOptions.NewDescription
	}
	if updateProjectOptions.NewID != nil {
		body["id"] = updateProjectOptions.NewID
	}
	if updateProjectOptions.NewCrn != nil {
		body["crn"] = updateProjectOptions.NewCrn
	}
	if updateProjectOptions.NewCreatedBy != nil {
		body["created_by"] = updateProjectOptions.NewCreatedBy
	}
	if updateProjectOptions.NewCreatedAt != nil {
		body["created_at"] = updateProjectOptions.NewCreatedAt
	}
	if updateProjectOptions.NewUpdatedAt != nil {
		body["updated_at"] = updateProjectOptions.NewUpdatedAt
	}
	if updateProjectOptions.NewRepoURL != nil {
		body["repo_url"] = updateProjectOptions.NewRepoURL
	}
	if updateProjectOptions.NewHref != nil {
		body["href"] = updateProjectOptions.NewHref
	}
	if updateProjectOptions.NewConfigs != nil {
		body["configs"] = updateProjectOptions.NewConfigs
	}
	if updateProjectOptions.NewDashboard != nil {
		body["dashboard"] = updateProjectOptions.NewDashboard
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProject)
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

// InstallProject : Install Project
// Install one or more project's configurations. It is an asynchronous operation that can be tracked using the project
// status api.
func (projects *ProjectsV1) InstallProject(installProjectOptions *InstallProjectOptions) (response *core.DetailedResponse, err error) {
	return projects.InstallProjectWithContext(context.Background(), installProjectOptions)
}

// InstallProjectWithContext is an alternate form of the InstallProject method which supports a Context parameter
func (projects *ProjectsV1) InstallProjectWithContext(ctx context.Context, installProjectOptions *InstallProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(installProjectOptions, "installProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(installProjectOptions, "installProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *installProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/install`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range installProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "InstallProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if installProjectOptions.ConfigNames != nil {
		body["config_names"] = installProjectOptions.ConfigNames
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

// UninstallProject : Uninstall Project
// Uninstall one or more project's configurations. The operation uninstall all the resources deployed with the given
// configurations. You can track it by using the project status api.
func (projects *ProjectsV1) UninstallProject(uninstallProjectOptions *UninstallProjectOptions) (response *core.DetailedResponse, err error) {
	return projects.UninstallProjectWithContext(context.Background(), uninstallProjectOptions)
}

// UninstallProjectWithContext is an alternate form of the UninstallProject method which supports a Context parameter
func (projects *ProjectsV1) UninstallProjectWithContext(ctx context.Context, uninstallProjectOptions *UninstallProjectOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uninstallProjectOptions, "uninstallProjectOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uninstallProjectOptions, "uninstallProjectOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *uninstallProjectOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/uninstall`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range uninstallProjectOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("projects", "V1", "UninstallProject")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if uninstallProjectOptions.ConfigNames != nil {
		body["config_names"] = uninstallProjectOptions.ConfigNames
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
	if updateProjectStatusOptions.ServicesStatus != nil {
		body["services_status"] = updateProjectStatusOptions.ServicesStatus
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
	if updateProjectConfigStatusOptions.Message != nil {
		body["message"] = updateProjectConfigStatusOptions.Message
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

// CreateProjectOptions : The CreateProject options.
type CreateProjectOptions struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A valid IAM api key with the permissions to manage the plumbing services; Schematics and Toolchain.
	ApiKey *string `json:"api_key" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	RepoURL *string `json:"repo_url,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectPrototypeDashboard `json:"dashboard,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*ProjectsV1) NewCreateProjectOptions(name string, apiKey string) *CreateProjectOptions {
	return &CreateProjectOptions{
		Name: core.StringPtr(name),
		ApiKey: core.StringPtr(apiKey),
	}
}

// SetName : Allow user to set Name
func (_options *CreateProjectOptions) SetName(name string) *CreateProjectOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetApiKey : Allow user to set ApiKey
func (_options *CreateProjectOptions) SetApiKey(apiKey string) *CreateProjectOptions {
	_options.ApiKey = core.StringPtr(apiKey)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateProjectOptions) SetDescription(description string) *CreateProjectOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetRepoURL : Allow user to set RepoURL
func (_options *CreateProjectOptions) SetRepoURL(repoURL string) *CreateProjectOptions {
	_options.RepoURL = core.StringPtr(repoURL)
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

// SetHeaders : Allow user to set Headers
func (options *GetProjectOptions) SetHeaders(param map[string]string) *GetProjectOptions {
	options.Headers = param
	return options
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

	// A descriptive of the variable.
	Description *string `json:"description,omitempty"`

	// The variable type.
	Type *string `json:"type,omitempty"`

	// Whether the variable is a secret.
	Sensitive *bool `json:"sensitive,omitempty"`

	// The variable value.
	Value interface{} `json:"value,omitempty"`

	// The variable default value.
	Default interface{} `json:"default,omitempty"`
}

// Constants associated with the InputVariable.Type property.
// The variable type.
const (
	InputVariable_Type_Bool = "bool"
	InputVariable_Type_List = "list"
	InputVariable_Type_Map = "map"
	InputVariable_Type_Number = "number"
	InputVariable_Type_Object = "object"
	InputVariable_Type_Set = "set"
	InputVariable_Type_String = "string"
	InputVariable_Type_Tuple = "tuple"
)

// NewInputVariable : Instantiate InputVariable (Generic Model Constructor)
func (*ProjectsV1) NewInputVariable(name string) (_model *InputVariable, err error) {
	_model = &InputVariable{
		Name: core.StringPtr(name),
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
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sensitive", &obj.Sensitive)
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

// InstallProjectOptions : The InstallProject options.
type InstallProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The configs to install. Leave the array empty to install all the configs.
	ConfigNames []string `json:"config_names,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallProjectOptions : Instantiate InstallProjectOptions
func (*ProjectsV1) NewInstallProjectOptions(id string) *InstallProjectOptions {
	return &InstallProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *InstallProjectOptions) SetID(id string) *InstallProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigNames : Allow user to set ConfigNames
func (_options *InstallProjectOptions) SetConfigNames(configNames []string) *InstallProjectOptions {
	_options.ConfigNames = configNames
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *InstallProjectOptions) SetHeaders(param map[string]string) *InstallProjectOptions {
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
	Projects []Project `json:"projects,omitempty"`
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
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProject)
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

// OutputValue : OutputValue struct
type OutputValue struct {
	// The output value name.
	Name *string `json:"name" validate:"required"`

	// A descriptive of the output value.
	Description *string `json:"description,omitempty"`

	// The output value.
	Value interface{} `json:"value,omitempty"`
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

// Project : Project struct
type Project struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project descriptive text.
	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	Crn *string `json:"crn,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	RepoURL *string `json:"repo_url,omitempty"`

	Href *string `json:"href,omitempty"`

	Configs []ProjectConfigIntf `json:"configs,omitempty"`

	Dashboard *ProjectDashboard `json:"dashboard,omitempty"`
}

// NewProject : Instantiate Project (Generic Model Constructor)
func (*ProjectsV1) NewProject(name string) (_model *Project, err error) {
	_model = &Project{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProject unmarshals an instance of Project from the specified map of raw messages.
func UnmarshalProject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Project)
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
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalProjectConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dashboard", &obj.Dashboard, UnmarshalProjectDashboard)
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

	Type *string `json:"type,omitempty"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`

	Input []InputVariable `json:"input,omitempty"`

	// A Terraform blueprint to use for provisioning a set of project resources.
	TerraformTemplate *TerraformTemplate `json:"terraform_template,omitempty"`

	// A Schematics blueprint to use for provisioning a set of project resources.
	SchematicsBlueprint *SchematicsBlueprint `json:"schematics_blueprint,omitempty"`
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
	err = core.UnmarshalModel(m, "terraform_template", &obj.TerraformTemplate, UnmarshalTerraformTemplate)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schematics_blueprint", &obj.SchematicsBlueprint, UnmarshalSchematicsBlueprint)
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

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	SchematicsResourceID *string `json:"schematics_resource_id,omitempty"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	Output []OutputValue `json:"output,omitempty"`
}

// Constants associated with the ProjectConfigStatus.Status property.
const (
	ProjectConfigStatus_Status_CheckInProgress = "check_in_progress"
	ProjectConfigStatus_Status_CheckSubmitted = "check_submitted"
	ProjectConfigStatus_Status_InError = "in_error"
	ProjectConfigStatus_Status_InstallInProgress = "install_in_progress"
	ProjectConfigStatus_Status_InstallSubmitted = "install_submitted"
	ProjectConfigStatus_Status_Installed = "installed"
	ProjectConfigStatus_Status_NotInstalled = "not_installed"
	ProjectConfigStatus_Status_PendingCheck = "pending_check"
	ProjectConfigStatus_Status_PendingInstall = "pending_install"
	ProjectConfigStatus_Status_PendingUninstall = "pending_uninstall"
	ProjectConfigStatus_Status_ToReinstall = "to_reinstall"
	ProjectConfigStatus_Status_UninstallInProgress = "uninstall_in_progress"
	ProjectConfigStatus_Status_UninstallSubmitted = "uninstall_submitted"
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
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
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

// ProjectDashboard : ProjectDashboard struct
type ProjectDashboard struct {
	Widgets []Widget `json:"widgets" validate:"required"`
}

// NewProjectDashboard : Instantiate ProjectDashboard (Generic Model Constructor)
func (*ProjectsV1) NewProjectDashboard(widgets []Widget) (_model *ProjectDashboard, err error) {
	_model = &ProjectDashboard{
		Widgets: widgets,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalProjectDashboard unmarshals an instance of ProjectDashboard from the specified map of raw messages.
func UnmarshalProjectDashboard(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectDashboard)
	err = core.UnmarshalModel(m, "widgets", &obj.Widgets, UnmarshalWidget)
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

// ProjectStatus : The project status, including plumbing services and computed statuses information.
type ProjectStatus struct {
	ProjectID *string `json:"project_id" validate:"required"`

	Href *string `json:"href" validate:"required"`

	// Project plumbing services and their status.
	ServicesStatus *ServicesStatus `json:"services_status" validate:"required"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`
}

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
	err = core.UnmarshalModel(m, "services_status", &obj.ServicesStatus, UnmarshalServicesStatus)
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

// ServiceStatus : The deployment status of a plumbing service.
type ServiceStatus struct {
	// An IBM Cloud Resource Name, which uniquely identify a resource.
	ID *string `json:"id" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	SchematicsResourceID *string `json:"schematics_resource_id,omitempty"`
}

// Constants associated with the ServiceStatus.Status property.
const (
	ServiceStatus_Status_CheckInProgress = "check_in_progress"
	ServiceStatus_Status_CheckSubmitted = "check_submitted"
	ServiceStatus_Status_InError = "in_error"
	ServiceStatus_Status_InstallInProgress = "install_in_progress"
	ServiceStatus_Status_InstallSubmitted = "install_submitted"
	ServiceStatus_Status_Installed = "installed"
	ServiceStatus_Status_NotInstalled = "not_installed"
	ServiceStatus_Status_PendingCheck = "pending_check"
	ServiceStatus_Status_PendingInstall = "pending_install"
	ServiceStatus_Status_PendingUninstall = "pending_uninstall"
	ServiceStatus_Status_ToReinstall = "to_reinstall"
	ServiceStatus_Status_UninstallInProgress = "uninstall_in_progress"
	ServiceStatus_Status_UninstallSubmitted = "uninstall_submitted"
)

// NewServiceStatus : Instantiate ServiceStatus (Generic Model Constructor)
func (*ProjectsV1) NewServiceStatus(id string, status string, message string) (_model *ServiceStatus, err error) {
	_model = &ServiceStatus{
		ID: core.StringPtr(id),
		Status: core.StringPtr(status),
		Message: core.StringPtr(message),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalServiceStatus unmarshals an instance of ServiceStatus from the specified map of raw messages.
func UnmarshalServiceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceStatus)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "schematics_resource_id", &obj.SchematicsResourceID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServicesStatus : Project plumbing services and their status.
type ServicesStatus struct {
	// The deployment status of a plumbing service.
	Toolchain *ServiceStatus `json:"toolchain" validate:"required"`

	// The deployment status of a plumbing service.
	Schematics *ServiceStatus `json:"schematics" validate:"required"`

	// The deployment status of a plumbing service.
	GitRepo *ServiceStatus `json:"git_repo,omitempty"`
}

// NewServicesStatus : Instantiate ServicesStatus (Generic Model Constructor)
func (*ProjectsV1) NewServicesStatus(toolchain *ServiceStatus, schematics *ServiceStatus) (_model *ServicesStatus, err error) {
	_model = &ServicesStatus{
		Toolchain: toolchain,
		Schematics: schematics,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalServicesStatus unmarshals an instance of ServicesStatus from the specified map of raw messages.
func UnmarshalServicesStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServicesStatus)
	err = core.UnmarshalModel(m, "toolchain", &obj.Toolchain, UnmarshalServiceStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schematics", &obj.Schematics, UnmarshalServiceStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "git_repo", &obj.GitRepo, UnmarshalServiceStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TerraformTemplate : A Terraform blueprint to use for provisioning a set of project resources.
type TerraformTemplate struct {
	RepoURL *string `json:"repo_url" validate:"required"`

	CatalogRef *string `json:"catalog_ref,omitempty"`

	TerraformVersion *string `json:"terraform_version" validate:"required"`

	PersonalAccessToken *string `json:"personal_access_token,omitempty"`
}

// NewTerraformTemplate : Instantiate TerraformTemplate (Generic Model Constructor)
func (*ProjectsV1) NewTerraformTemplate(repoURL string, terraformVersion string) (_model *TerraformTemplate, err error) {
	_model = &TerraformTemplate{
		RepoURL: core.StringPtr(repoURL),
		TerraformVersion: core.StringPtr(terraformVersion),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTerraformTemplate unmarshals an instance of TerraformTemplate from the specified map of raw messages.
func UnmarshalTerraformTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TerraformTemplate)
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_ref", &obj.CatalogRef)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "terraform_version", &obj.TerraformVersion)
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

// UninstallProjectOptions : The UninstallProject options.
type UninstallProjectOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The configs to uninstall. Leave the array empty to install all the configs.
	ConfigNames []string `json:"config_names,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUninstallProjectOptions : Instantiate UninstallProjectOptions
func (*ProjectsV1) NewUninstallProjectOptions(id string) *UninstallProjectOptions {
	return &UninstallProjectOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *UninstallProjectOptions) SetID(id string) *UninstallProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetConfigNames : Allow user to set ConfigNames
func (_options *UninstallProjectOptions) SetConfigNames(configNames []string) *UninstallProjectOptions {
	_options.ConfigNames = configNames
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UninstallProjectOptions) SetHeaders(param map[string]string) *UninstallProjectOptions {
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

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	SchematicsResourceID *string `json:"schematics_resource_id,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateProjectConfigStatusOptions.Status property.
const (
	UpdateProjectConfigStatusOptions_Status_CheckInProgress = "check_in_progress"
	UpdateProjectConfigStatusOptions_Status_CheckSubmitted = "check_submitted"
	UpdateProjectConfigStatusOptions_Status_InError = "in_error"
	UpdateProjectConfigStatusOptions_Status_InstallInProgress = "install_in_progress"
	UpdateProjectConfigStatusOptions_Status_InstallSubmitted = "install_submitted"
	UpdateProjectConfigStatusOptions_Status_Installed = "installed"
	UpdateProjectConfigStatusOptions_Status_NotInstalled = "not_installed"
	UpdateProjectConfigStatusOptions_Status_PendingCheck = "pending_check"
	UpdateProjectConfigStatusOptions_Status_PendingInstall = "pending_install"
	UpdateProjectConfigStatusOptions_Status_PendingUninstall = "pending_uninstall"
	UpdateProjectConfigStatusOptions_Status_ToReinstall = "to_reinstall"
	UpdateProjectConfigStatusOptions_Status_UninstallInProgress = "uninstall_in_progress"
	UpdateProjectConfigStatusOptions_Status_UninstallSubmitted = "uninstall_submitted"
)

// NewUpdateProjectConfigStatusOptions : Instantiate UpdateProjectConfigStatusOptions
func (*ProjectsV1) NewUpdateProjectConfigStatusOptions(id string, configName string, status string, message string) *UpdateProjectConfigStatusOptions {
	return &UpdateProjectConfigStatusOptions{
		ID: core.StringPtr(id),
		ConfigName: core.StringPtr(configName),
		Status: core.StringPtr(status),
		Message: core.StringPtr(message),
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

// SetMessage : Allow user to set Message
func (_options *UpdateProjectConfigStatusOptions) SetMessage(message string) *UpdateProjectConfigStatusOptions {
	_options.Message = core.StringPtr(message)
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
	ID *string `json:"-" validate:"required,ne="`

	// The project name.
	NewName *string `json:"name" validate:"required"`

	// A project descriptive text.
	NewDescription *string `json:"description,omitempty"`

	NewID *string `json:"id,omitempty"`

	// An IBM Cloud Resource Name, which uniquely identify a resource.
	NewCrn *string `json:"crn,omitempty"`

	NewCreatedBy *string `json:"created_by,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	NewCreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as
	// specified by RFC 3339.
	NewUpdatedAt *strfmt.DateTime `json:"updated_at,omitempty"`

	NewRepoURL *string `json:"repo_url,omitempty"`

	NewHref *string `json:"href,omitempty"`

	NewConfigs []ProjectConfigIntf `json:"configs,omitempty"`

	NewDashboard *ProjectDashboard `json:"dashboard,omitempty"`

	// Set this header to control the return of the Project. If return=minimal is set, a successful response has a 201
	// Created or 204 No Content status and includes no body. If return=representation is set, a successful response
	// includes the updated Project. Default includes no body.
	Prefer *string `json:"Prefer,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateProjectOptions.Prefer property.
// Set this header to control the return of the Project. If return=minimal is set, a successful response has a 201
// Created or 204 No Content status and includes no body. If return=representation is set, a successful response
// includes the updated Project. Default includes no body.
const (
	UpdateProjectOptions_Prefer_ReturnMinimal = "return=minimal"
	UpdateProjectOptions_Prefer_ReturnRepresentation = "return=representation"
)

// NewUpdateProjectOptions : Instantiate UpdateProjectOptions
func (*ProjectsV1) NewUpdateProjectOptions(id string, newName string) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
		NewName: core.StringPtr(newName),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectOptions) SetID(id string) *UpdateProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetNewName : Allow user to set NewName
func (_options *UpdateProjectOptions) SetNewName(newName string) *UpdateProjectOptions {
	_options.NewName = core.StringPtr(newName)
	return _options
}

// SetNewDescription : Allow user to set NewDescription
func (_options *UpdateProjectOptions) SetNewDescription(newDescription string) *UpdateProjectOptions {
	_options.NewDescription = core.StringPtr(newDescription)
	return _options
}

// SetNewID : Allow user to set NewID
func (_options *UpdateProjectOptions) SetNewID(newID string) *UpdateProjectOptions {
	_options.NewID = core.StringPtr(newID)
	return _options
}

// SetNewCrn : Allow user to set NewCrn
func (_options *UpdateProjectOptions) SetNewCrn(newCrn string) *UpdateProjectOptions {
	_options.NewCrn = core.StringPtr(newCrn)
	return _options
}

// SetNewCreatedBy : Allow user to set NewCreatedBy
func (_options *UpdateProjectOptions) SetNewCreatedBy(newCreatedBy string) *UpdateProjectOptions {
	_options.NewCreatedBy = core.StringPtr(newCreatedBy)
	return _options
}

// SetNewCreatedAt : Allow user to set NewCreatedAt
func (_options *UpdateProjectOptions) SetNewCreatedAt(newCreatedAt *strfmt.DateTime) *UpdateProjectOptions {
	_options.NewCreatedAt = newCreatedAt
	return _options
}

// SetNewUpdatedAt : Allow user to set NewUpdatedAt
func (_options *UpdateProjectOptions) SetNewUpdatedAt(newUpdatedAt *strfmt.DateTime) *UpdateProjectOptions {
	_options.NewUpdatedAt = newUpdatedAt
	return _options
}

// SetNewRepoURL : Allow user to set NewRepoURL
func (_options *UpdateProjectOptions) SetNewRepoURL(newRepoURL string) *UpdateProjectOptions {
	_options.NewRepoURL = core.StringPtr(newRepoURL)
	return _options
}

// SetNewHref : Allow user to set NewHref
func (_options *UpdateProjectOptions) SetNewHref(newHref string) *UpdateProjectOptions {
	_options.NewHref = core.StringPtr(newHref)
	return _options
}

// SetNewConfigs : Allow user to set NewConfigs
func (_options *UpdateProjectOptions) SetNewConfigs(newConfigs []ProjectConfigIntf) *UpdateProjectOptions {
	_options.NewConfigs = newConfigs
	return _options
}

// SetNewDashboard : Allow user to set NewDashboard
func (_options *UpdateProjectOptions) SetNewDashboard(newDashboard *ProjectDashboard) *UpdateProjectOptions {
	_options.NewDashboard = newDashboard
	return _options
}

// SetPrefer : Allow user to set Prefer
func (_options *UpdateProjectOptions) SetPrefer(prefer string) *UpdateProjectOptions {
	_options.Prefer = core.StringPtr(prefer)
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

	// Project plumbing services and their status.
	ServicesStatus *ServicesStatus `json:"services_status" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectStatusOptions : Instantiate UpdateProjectStatusOptions
func (*ProjectsV1) NewUpdateProjectStatusOptions(id string, servicesStatus *ServicesStatus) *UpdateProjectStatusOptions {
	return &UpdateProjectStatusOptions{
		ID: core.StringPtr(id),
		ServicesStatus: servicesStatus,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectStatusOptions) SetID(id string) *UpdateProjectStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetServicesStatus : Allow user to set ServicesStatus
func (_options *UpdateProjectStatusOptions) SetServicesStatus(servicesStatus *ServicesStatus) *UpdateProjectStatusOptions {
	_options.ServicesStatus = servicesStatus
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateProjectStatusOptions) SetHeaders(param map[string]string) *UpdateProjectStatusOptions {
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

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input" validate:"required"`

	// A Schematics blueprint to use for provisioning a set of project resources.
	SchematicsBlueprint *SchematicsBlueprint `json:"schematics_blueprint,omitempty"`
}

// Constants associated with the ProjectConfigSchematicsBlueprintProperty.Type property.
const (
	ProjectConfigSchematicsBlueprintProperty_Type_SchematicsBlueprint = "schematics_blueprint"
)

// NewProjectConfigSchematicsBlueprintProperty : Instantiate ProjectConfigSchematicsBlueprintProperty (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigSchematicsBlueprintProperty(name string, typeVar string, input []InputVariable) (_model *ProjectConfigSchematicsBlueprintProperty, err error) {
	_model = &ProjectConfigSchematicsBlueprintProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Input: input,
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "schematics_blueprint", &obj.SchematicsBlueprint, UnmarshalSchematicsBlueprint)
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

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input" validate:"required"`

	// A Terraform blueprint to use for provisioning a set of project resources.
	TerraformTemplate *TerraformTemplate `json:"terraform_template,omitempty"`
}

// Constants associated with the ProjectConfigTerraformTemplateProperty.Type property.
const (
	ProjectConfigTerraformTemplateProperty_Type_TerraformTemplate = "terraform_template"
)

// NewProjectConfigTerraformTemplateProperty : Instantiate ProjectConfigTerraformTemplateProperty (Generic Model Constructor)
func (*ProjectsV1) NewProjectConfigTerraformTemplateProperty(name string, typeVar string, input []InputVariable) (_model *ProjectConfigTerraformTemplateProperty, err error) {
	_model = &ProjectConfigTerraformTemplateProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Input: input,
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
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "input", &obj.Input, UnmarshalInputVariable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "terraform_template", &obj.TerraformTemplate, UnmarshalTerraformTemplate)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
