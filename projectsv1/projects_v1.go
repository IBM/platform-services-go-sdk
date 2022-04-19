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
 * IBM OpenAPI SDK Code Generator Version: 3.46.1-a5569134-20220316-164819
 */

// Package projectsv1 : Operations and models for the ProjectsV1 service
package projectsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
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
func (projects *ProjectsV1) CreateProject(createProjectOptions *CreateProjectOptions) (result *CreateProjectResponse, response *core.DetailedResponse, err error) {
	return projects.CreateProjectWithContext(context.Background(), createProjectOptions)
}

// CreateProjectWithContext is an alternate form of the CreateProject method which supports a Context parameter
func (projects *ProjectsV1) CreateProjectWithContext(ctx context.Context, createProjectOptions *CreateProjectOptions) (result *CreateProjectResponse, response *core.DetailedResponse, err error) {
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
	if createProjectOptions.Metadata != nil {
		body["metadata"] = createProjectOptions.Metadata
	}
	if createProjectOptions.Spec != nil {
		body["spec"] = createProjectOptions.Spec
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateProjectResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListProjects : List installed Projects
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

	if listProjectsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listProjectsOptions.Offset))
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
func (projects *ProjectsV1) GetProject(getProjectOptions *GetProjectOptions) (result *ProjectDefinitionResult, response *core.DetailedResponse, err error) {
	return projects.GetProjectWithContext(context.Background(), getProjectOptions)
}

// GetProjectWithContext is an alternate form of the GetProject method which supports a Context parameter
func (projects *ProjectsV1) GetProjectWithContext(ctx context.Context, getProjectOptions *GetProjectOptions) (result *ProjectDefinitionResult, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectDefinitionResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProject : Update a project by id
// Update a project.
func (projects *ProjectsV1) UpdateProject(updateProjectOptions *UpdateProjectOptions) (result *ProjectDefinitionResult, response *core.DetailedResponse, err error) {
	return projects.UpdateProjectWithContext(context.Background(), updateProjectOptions)
}

// UpdateProjectWithContext is an alternate form of the UpdateProject method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectWithContext(ctx context.Context, updateProjectOptions *UpdateProjectOptions) (result *ProjectDefinitionResult, response *core.DetailedResponse, err error) {
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
	if updateProjectOptions.Metadata != nil {
		body["metadata"] = updateProjectOptions.Metadata
	}
	if updateProjectOptions.Spec != nil {
		body["spec"] = updateProjectOptions.Spec
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectDefinitionResult)
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
func (projects *ProjectsV1) GetProjectStatus(getProjectStatusOptions *GetProjectStatusOptions) (result *ConfigStatusResult, response *core.DetailedResponse, err error) {
	return projects.GetProjectStatusWithContext(context.Background(), getProjectStatusOptions)
}

// GetProjectStatusWithContext is an alternate form of the GetProjectStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectStatusWithContext(ctx context.Context, getProjectStatusOptions *GetProjectStatusOptions) (result *ConfigStatusResult, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfigStatusResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectComputedStatus : Update a project computed status
// Update a computed status with content to which the projects service is agnostic.
func (projects *ProjectsV1) UpdateProjectComputedStatus(updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions) (response *core.DetailedResponse, err error) {
	return projects.UpdateProjectComputedStatusWithContext(context.Background(), updateProjectComputedStatusOptions)
}

// UpdateProjectComputedStatusWithContext is an alternate form of the UpdateProjectComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectComputedStatusWithContext(ctx context.Context, updateProjectComputedStatusOptions *UpdateProjectComputedStatusOptions) (response *core.DetailedResponse, err error) {
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
		"status_name": *updateProjectComputedStatusOptions.StatusName,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/computed_status/{status_name}`, pathParamsMap)
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
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateProjectComputedStatusOptions.RequestBody)
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

// GetProjectConfigStatus : Get project status information
// Returns the detailed project status, including all the config statuses and all the computed statuses sent by external
// tools.
func (projects *ProjectsV1) GetProjectConfigStatus(getProjectConfigStatusOptions *GetProjectConfigStatusOptions) (result *ProjectStatusResult, response *core.DetailedResponse, err error) {
	return projects.GetProjectConfigStatusWithContext(context.Background(), getProjectConfigStatusOptions)
}

// GetProjectConfigStatusWithContext is an alternate form of the GetProjectConfigStatus method which supports a Context parameter
func (projects *ProjectsV1) GetProjectConfigStatusWithContext(ctx context.Context, getProjectConfigStatusOptions *GetProjectConfigStatusOptions) (result *ProjectStatusResult, response *core.DetailedResponse, err error) {
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
		"name": *getProjectConfigStatusOptions.Name,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{name}/status`, pathParamsMap)
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProjectStatusResult)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateProjectConfigStatus : Update a project config status
// Update a config status and eventually the output values.
func (projects *ProjectsV1) UpdateProjectConfigStatus(updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions) (response *core.DetailedResponse, err error) {
	return projects.UpdateProjectConfigStatusWithContext(context.Background(), updateProjectConfigStatusOptions)
}

// UpdateProjectConfigStatusWithContext is an alternate form of the UpdateProjectConfigStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectConfigStatusWithContext(ctx context.Context, updateProjectConfigStatusOptions *UpdateProjectConfigStatusOptions) (response *core.DetailedResponse, err error) {
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
		"name": *updateProjectConfigStatusOptions.Name,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{name}/status`, pathParamsMap)
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
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateProjectConfigStatusOptions.NewName != nil {
		body["name"] = updateProjectConfigStatusOptions.NewName
	}
	if updateProjectConfigStatusOptions.NewStatus != nil {
		body["status"] = updateProjectConfigStatusOptions.NewStatus
	}
	if updateProjectConfigStatusOptions.NewMessage != nil {
		body["message"] = updateProjectConfigStatusOptions.NewMessage
	}
	if updateProjectConfigStatusOptions.NewPipelineRun != nil {
		body["pipeline_run"] = updateProjectConfigStatusOptions.NewPipelineRun
	}
	if updateProjectConfigStatusOptions.NewSchematics != nil {
		body["schematics"] = updateProjectConfigStatusOptions.NewSchematics
	}
	if updateProjectConfigStatusOptions.NewComputedStatuses != nil {
		body["computed_statuses"] = updateProjectConfigStatusOptions.NewComputedStatuses
	}
	if updateProjectConfigStatusOptions.NewOutput != nil {
		body["output"] = updateProjectConfigStatusOptions.NewOutput
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

// UpdateProjectConfigComputedStatus : Update a project computed config status
// Update a computed config status with a content to which the projects service is agnostic.
func (projects *ProjectsV1) UpdateProjectConfigComputedStatus(updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions) (response *core.DetailedResponse, err error) {
	return projects.UpdateProjectConfigComputedStatusWithContext(context.Background(), updateProjectConfigComputedStatusOptions)
}

// UpdateProjectConfigComputedStatusWithContext is an alternate form of the UpdateProjectConfigComputedStatus method which supports a Context parameter
func (projects *ProjectsV1) UpdateProjectConfigComputedStatusWithContext(ctx context.Context, updateProjectConfigComputedStatusOptions *UpdateProjectConfigComputedStatusOptions) (response *core.DetailedResponse, err error) {
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
		"name": *updateProjectConfigComputedStatusOptions.Name,
		"status_name": *updateProjectConfigComputedStatusOptions.StatusName,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = projects.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(projects.Service.Options.URL, `/v1/projects/{id}/configs/{name}/computed_status/{status_name}`, pathParamsMap)
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
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateProjectConfigComputedStatusOptions.RequestBody)
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
func (projects *ProjectsV1) GetHealth(getHealthOptions *GetHealthOptions) (result *HealthResponse, response *core.DetailedResponse, err error) {
	return projects.GetHealthWithContext(context.Background(), getHealthOptions)
}

// GetHealthWithContext is an alternate form of the GetHealth method which supports a Context parameter
func (projects *ProjectsV1) GetHealthWithContext(ctx context.Context, getHealthOptions *GetHealthOptions) (result *HealthResponse, response *core.DetailedResponse, err error) {
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
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHealthResponse)
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
	// Metadata attributes about the project.
	Metadata *Metadata `json:"metadata" validate:"required"`

	// The spec section includes project definition information, which must be given as input to create the project.
	Spec *Spec `json:"spec" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateProjectOptions : Instantiate CreateProjectOptions
func (*ProjectsV1) NewCreateProjectOptions(metadata *Metadata, spec *Spec) *CreateProjectOptions {
	return &CreateProjectOptions{
		Metadata: metadata,
		Spec: spec,
	}
}

// SetMetadata : Allow user to set Metadata
func (_options *CreateProjectOptions) SetMetadata(metadata *Metadata) *CreateProjectOptions {
	_options.Metadata = metadata
	return _options
}

// SetSpec : Allow user to set Spec
func (_options *CreateProjectOptions) SetSpec(spec *Spec) *CreateProjectOptions {
	_options.Spec = spec
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateProjectOptions) SetHeaders(param map[string]string) *CreateProjectOptions {
	options.Headers = param
	return options
}

// CreateProjectResponse : The project id.
type CreateProjectResponse struct {
	ProjectID *string `json:"project_id" validate:"required"`
}

// UnmarshalCreateProjectResponse unmarshals an instance of CreateProjectResponse from the specified map of raw messages.
func UnmarshalCreateProjectResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateProjectResponse)
	err = core.UnmarshalPrimitive(m, "project_id", &obj.ProjectID)
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

// GetProjectConfigStatusOptions : The GetProjectConfigStatus options.
type GetProjectConfigStatusOptions struct {
	// The id of the project, which uniquely identifies it.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the config, which must be unique within the project.
	Name *string `json:"name" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetProjectConfigStatusOptions : Instantiate GetProjectConfigStatusOptions
func (*ProjectsV1) NewGetProjectConfigStatusOptions(id string, name string) *GetProjectConfigStatusOptions {
	return &GetProjectConfigStatusOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
	}
}

// SetID : Allow user to set ID
func (_options *GetProjectConfigStatusOptions) SetID(id string) *GetProjectConfigStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *GetProjectConfigStatusOptions) SetName(name string) *GetProjectConfigStatusOptions {
	_options.Name = core.StringPtr(name)
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

// ListProjectsOptions : The ListProjects options.
type ListProjectsOptions struct {
	// Used to determine how many projects to skip over, given the order of the collection. If offset is unspecified, it
	// defaults to 0.
	Offset *int64 `json:"offset,omitempty"`

	// Determine how many resources are returned, unless the offset and limit are such that the response is the last page
	// of resources.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProjectsOptions : Instantiate ListProjectsOptions
func (*ProjectsV1) NewListProjectsOptions() *ListProjectsOptions {
	return &ListProjectsOptions{}
}

// SetOffset : Allow user to set Offset
func (_options *ListProjectsOptions) SetOffset(offset int64) *ListProjectsOptions {
	_options.Offset = core.Int64Ptr(offset)
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
	// Replies the input offset parameter.
	Offset *int64 `json:"offset" validate:"required"`

	// Replies the input limit parameter.
	Limit *int64 `json:"limit" validate:"required"`

	// Get the occurrencies of the total Projects.
	TotalCount *int64 `json:"total_count" validate:"required"`

	First *Href `json:"first,omitempty"`

	Last *Href `json:"last,omitempty"`

	Previous *Href `json:"previous,omitempty"`

	Next *Href `json:"next,omitempty"`

	// An array of projects.
	Projects []ProjectMetadataResult `json:"projects,omitempty"`
}

// UnmarshalListProjectsResponse unmarshals an instance of ListProjectsResponse from the specified map of raw messages.
func UnmarshalListProjectsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListProjectsResponse)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "projects", &obj.Projects, UnmarshalProjectMetadataResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListProjectsResponse) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// SpecDashboard : SpecDashboard struct
type SpecDashboard struct {
	Widgets []Widget `json:"widgets" validate:"required"`
}

// NewSpecDashboard : Instantiate SpecDashboard (Generic Model Constructor)
func (*ProjectsV1) NewSpecDashboard(widgets []Widget) (_model *SpecDashboard, err error) {
	_model = &SpecDashboard{
		Widgets: widgets,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSpecDashboard unmarshals an instance of SpecDashboard from the specified map of raw messages.
func UnmarshalSpecDashboard(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpecDashboard)
	err = core.UnmarshalModel(m, "widgets", &obj.Widgets, UnmarshalWidget)
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

	// The name of the computed status, which must be unique within the config.
	StatusName *string `json:"status_name" validate:"required,ne="`

	// The computed status to set.
	RequestBody map[string]interface{} `json:"request_body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectComputedStatusOptions : Instantiate UpdateProjectComputedStatusOptions
func (*ProjectsV1) NewUpdateProjectComputedStatusOptions(id string, statusName string, requestBody map[string]interface{}) *UpdateProjectComputedStatusOptions {
	return &UpdateProjectComputedStatusOptions{
		ID: core.StringPtr(id),
		StatusName: core.StringPtr(statusName),
		RequestBody: requestBody,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectComputedStatusOptions) SetID(id string) *UpdateProjectComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetStatusName : Allow user to set StatusName
func (_options *UpdateProjectComputedStatusOptions) SetStatusName(statusName string) *UpdateProjectComputedStatusOptions {
	_options.StatusName = core.StringPtr(statusName)
	return _options
}

// SetRequestBody : Allow user to set RequestBody
func (_options *UpdateProjectComputedStatusOptions) SetRequestBody(requestBody map[string]interface{}) *UpdateProjectComputedStatusOptions {
	_options.RequestBody = requestBody
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

	// The name of the config, which must be unique within the project.
	Name *string `json:"name" validate:"required,ne="`

	// The name of the computed status, which must be unique within the config.
	StatusName *string `json:"status_name" validate:"required,ne="`

	// The config computed status to set.
	RequestBody map[string]interface{} `json:"request_body" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateProjectConfigComputedStatusOptions : Instantiate UpdateProjectConfigComputedStatusOptions
func (*ProjectsV1) NewUpdateProjectConfigComputedStatusOptions(id string, name string, statusName string, requestBody map[string]interface{}) *UpdateProjectConfigComputedStatusOptions {
	return &UpdateProjectConfigComputedStatusOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
		StatusName: core.StringPtr(statusName),
		RequestBody: requestBody,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectConfigComputedStatusOptions) SetID(id string) *UpdateProjectConfigComputedStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateProjectConfigComputedStatusOptions) SetName(name string) *UpdateProjectConfigComputedStatusOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetStatusName : Allow user to set StatusName
func (_options *UpdateProjectConfigComputedStatusOptions) SetStatusName(statusName string) *UpdateProjectConfigComputedStatusOptions {
	_options.StatusName = core.StringPtr(statusName)
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
	Name *string `json:"-" validate:"required,ne="`

	NewName *string `json:"name" validate:"required"`

	NewStatus *string `json:"status" validate:"required"`

	// A detailed status message when applicable.
	NewMessage *string `json:"message" validate:"required"`

	NewPipelineRun *string `json:"pipeline_run,omitempty"`

	NewSchematics *string `json:"schematics,omitempty"`

	NewComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	NewOutput []OutputValue `json:"output,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateProjectConfigStatusOptions.NewStatus property.
const (
	UpdateProjectConfigStatusOptions_NewStatus_CheckInProgress = "check_in_progress"
	UpdateProjectConfigStatusOptions_NewStatus_InError = "in_error"
	UpdateProjectConfigStatusOptions_NewStatus_InstallInProgress = "install_in_progress"
	UpdateProjectConfigStatusOptions_NewStatus_Installed = "installed"
	UpdateProjectConfigStatusOptions_NewStatus_NotInstalled = "not_installed"
	UpdateProjectConfigStatusOptions_NewStatus_UninstallInProgress = "uninstall_in_progress"
)

// NewUpdateProjectConfigStatusOptions : Instantiate UpdateProjectConfigStatusOptions
func (*ProjectsV1) NewUpdateProjectConfigStatusOptions(id string, name string, newName string, newStatus string, newMessage string) *UpdateProjectConfigStatusOptions {
	return &UpdateProjectConfigStatusOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
		NewName: core.StringPtr(newName),
		NewStatus: core.StringPtr(newStatus),
		NewMessage: core.StringPtr(newMessage),
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectConfigStatusOptions) SetID(id string) *UpdateProjectConfigStatusOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateProjectConfigStatusOptions) SetName(name string) *UpdateProjectConfigStatusOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetNewName : Allow user to set NewName
func (_options *UpdateProjectConfigStatusOptions) SetNewName(newName string) *UpdateProjectConfigStatusOptions {
	_options.NewName = core.StringPtr(newName)
	return _options
}

// SetNewStatus : Allow user to set NewStatus
func (_options *UpdateProjectConfigStatusOptions) SetNewStatus(newStatus string) *UpdateProjectConfigStatusOptions {
	_options.NewStatus = core.StringPtr(newStatus)
	return _options
}

// SetNewMessage : Allow user to set NewMessage
func (_options *UpdateProjectConfigStatusOptions) SetNewMessage(newMessage string) *UpdateProjectConfigStatusOptions {
	_options.NewMessage = core.StringPtr(newMessage)
	return _options
}

// SetNewPipelineRun : Allow user to set NewPipelineRun
func (_options *UpdateProjectConfigStatusOptions) SetNewPipelineRun(newPipelineRun string) *UpdateProjectConfigStatusOptions {
	_options.NewPipelineRun = core.StringPtr(newPipelineRun)
	return _options
}

// SetNewSchematics : Allow user to set NewSchematics
func (_options *UpdateProjectConfigStatusOptions) SetNewSchematics(newSchematics string) *UpdateProjectConfigStatusOptions {
	_options.NewSchematics = core.StringPtr(newSchematics)
	return _options
}

// SetNewComputedStatuses : Allow user to set NewComputedStatuses
func (_options *UpdateProjectConfigStatusOptions) SetNewComputedStatuses(newComputedStatuses map[string]interface{}) *UpdateProjectConfigStatusOptions {
	_options.NewComputedStatuses = newComputedStatuses
	return _options
}

// SetNewOutput : Allow user to set NewOutput
func (_options *UpdateProjectConfigStatusOptions) SetNewOutput(newOutput []OutputValue) *UpdateProjectConfigStatusOptions {
	_options.NewOutput = newOutput
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

	// Metadata attributes about the project.
	Metadata *Metadata `json:"metadata" validate:"required"`

	// The spec section includes project definition information, which must be given as input to create the project.
	Spec *Spec `json:"spec" validate:"required"`

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
func (*ProjectsV1) NewUpdateProjectOptions(id string, metadata *Metadata, spec *Spec) *UpdateProjectOptions {
	return &UpdateProjectOptions{
		ID: core.StringPtr(id),
		Metadata: metadata,
		Spec: spec,
	}
}

// SetID : Allow user to set ID
func (_options *UpdateProjectOptions) SetID(id string) *UpdateProjectOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *UpdateProjectOptions) SetMetadata(metadata *Metadata) *UpdateProjectOptions {
	_options.Metadata = metadata
	return _options
}

// SetSpec : Allow user to set Spec
func (_options *UpdateProjectOptions) SetSpec(spec *Spec) *UpdateProjectOptions {
	_options.Spec = spec
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

// Config : Config struct
// Models which "extend" this model:
// - ConfigManualProperty
// - ConfigTerraformTemplateProperty
// - ConfigSchematicsBlueprintProperty
type Config struct {
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

// Constants associated with the Config.Type property.
const (
	Config_Type_Manual = "manual"
)
func (*Config) isaConfig() bool {
	return true
}

type ConfigIntf interface {
	isaConfig() bool
}

// UnmarshalConfig unmarshals an instance of Config from the specified map of raw messages.
func UnmarshalConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Config)
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

// ConfigStatus : ConfigStatus struct
type ConfigStatus struct {
	Name *string `json:"name" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	Schematics *string `json:"schematics,omitempty"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	Output []OutputValue `json:"output,omitempty"`
}

// Constants associated with the ConfigStatus.Status property.
const (
	ConfigStatus_Status_CheckInProgress = "check_in_progress"
	ConfigStatus_Status_InError = "in_error"
	ConfigStatus_Status_InstallInProgress = "install_in_progress"
	ConfigStatus_Status_Installed = "installed"
	ConfigStatus_Status_NotInstalled = "not_installed"
	ConfigStatus_Status_UninstallInProgress = "uninstall_in_progress"
)

// NewConfigStatus : Instantiate ConfigStatus (Generic Model Constructor)
func (*ProjectsV1) NewConfigStatus(name string, status string, message string) (_model *ConfigStatus, err error) {
	_model = &ConfigStatus{
		Name: core.StringPtr(name),
		Status: core.StringPtr(status),
		Message: core.StringPtr(message),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalConfigStatus unmarshals an instance of ConfigStatus from the specified map of raw messages.
func UnmarshalConfigStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigStatus)
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
	err = core.UnmarshalPrimitive(m, "schematics", &obj.Schematics)
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

// ConfigStatusResult : Project config status result.
type ConfigStatusResult struct {
	Href *string `json:"href,omitempty"`

	Name *string `json:"name" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	PipelineRun *string `json:"pipeline_run,omitempty"`

	Schematics *string `json:"schematics,omitempty"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	Output []OutputValue `json:"output,omitempty"`
}

// Constants associated with the ConfigStatusResult.Status property.
const (
	ConfigStatusResult_Status_CheckInProgress = "check_in_progress"
	ConfigStatusResult_Status_InError = "in_error"
	ConfigStatusResult_Status_InstallInProgress = "install_in_progress"
	ConfigStatusResult_Status_Installed = "installed"
	ConfigStatusResult_Status_NotInstalled = "not_installed"
	ConfigStatusResult_Status_UninstallInProgress = "uninstall_in_progress"
)

// UnmarshalConfigStatusResult unmarshals an instance of ConfigStatusResult from the specified map of raw messages.
func UnmarshalConfigStatusResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigStatusResult)
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
	err = core.UnmarshalPrimitive(m, "schematics", &obj.Schematics)
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

// HealthResponse : HealthResponse struct
type HealthResponse struct {
	// The name of the service.
	Name *string `json:"name,omitempty"`

	// The running version of the service.
	Version *string `json:"version,omitempty"`

	// The status of service dependencies.
	Dependencies map[string]interface{} `json:"dependencies,omitempty"`
}

// UnmarshalHealthResponse unmarshals an instance of HealthResponse from the specified map of raw messages.
func UnmarshalHealthResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthResponse)
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

// Href : Href struct
type Href struct {
	Href *string `json:"href,omitempty"`
}

// UnmarshalHref unmarshals an instance of Href from the specified map of raw messages.
func UnmarshalHref(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Href)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
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

// Metadata : Metadata attributes about the project.
type Metadata struct {
	// The project name.
	Name *string `json:"name" validate:"required"`

	// A project's descriptive text.
	Description *string `json:"description,omitempty"`

	ID *string `json:"id,omitempty"`

	CreatedBy *string `json:"created_by,omitempty"`

	// UTC format YYYY-MM-DDTHH:mm:ss.sssZ.
	CreatedAt *string `json:"created_at,omitempty"`

	// UTC format YYYY-MM-DDTHH:mm:ss.sssZ.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// User tags that the System automatically attach to the project.
	Tags []string `json:"tags,omitempty"`

	// The project git repo url.
	RepoURL *string `json:"repo_url,omitempty"`
}

// NewMetadata : Instantiate Metadata (Generic Model Constructor)
func (*ProjectsV1) NewMetadata(name string) (_model *Metadata, err error) {
	_model = &Metadata{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalMetadata unmarshals an instance of Metadata from the specified map of raw messages.
func UnmarshalMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metadata)
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
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
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

// ProjectDefinitionResult : ProjectDefinitionResult struct
type ProjectDefinitionResult struct {
	Href *string `json:"href,omitempty"`

	// Metadata attributes about the project.
	Metadata *Metadata `json:"metadata" validate:"required"`

	// The spec section includes project definition information, which must be given as input to create the project.
	Spec *Spec `json:"spec" validate:"required"`
}

// UnmarshalProjectDefinitionResult unmarshals an instance of ProjectDefinitionResult from the specified map of raw messages.
func UnmarshalProjectDefinitionResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectDefinitionResult)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "spec", &obj.Spec, UnmarshalSpec)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectMetadataResult : ProjectMetadataResult struct
type ProjectMetadataResult struct {
	Href *string `json:"href,omitempty"`

	// Metadata attributes about the project.
	Metadata *Metadata `json:"metadata" validate:"required"`
}

// UnmarshalProjectMetadataResult unmarshals an instance of ProjectMetadataResult from the specified map of raw messages.
func UnmarshalProjectMetadataResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectMetadataResult)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProjectStatusResult : Project status result.
type ProjectStatusResult struct {
	Href *string `json:"href,omitempty"`

	// Metadata attributes about the project.
	Metadata *Metadata `json:"metadata" validate:"required"`

	// The status section includes the deployment status of the project. It is returned by the status API.
	Status *Status `json:"status" validate:"required"`
}

// UnmarshalProjectStatusResult unmarshals an instance of ProjectStatusResult from the specified map of raw messages.
func UnmarshalProjectStatusResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProjectStatusResult)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "status", &obj.Status, UnmarshalStatus)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsBlueprint : A Schematics blueprint to use for provisioning a set of project resources.
type SchematicsBlueprint struct {
	RepoURL *string `json:"repo_url" validate:"required"`

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
	// An IBM Cloud CRN.
	ID *string `json:"id" validate:"required"`

	Status *string `json:"status" validate:"required"`

	// A detailed status message when applicable.
	Message *string `json:"message" validate:"required"`

	Schematics *string `json:"schematics,omitempty"`
}

// Constants associated with the ServiceStatus.Status property.
const (
	ServiceStatus_Status_CheckInProgress = "check_in_progress"
	ServiceStatus_Status_InError = "in_error"
	ServiceStatus_Status_InstallInProgress = "install_in_progress"
	ServiceStatus_Status_Installed = "installed"
	ServiceStatus_Status_NotInstalled = "not_installed"
	ServiceStatus_Status_UninstallInProgress = "uninstall_in_progress"
)

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
	err = core.UnmarshalPrimitive(m, "schematics", &obj.Schematics)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Services : Project plumbing services.
type Services struct {
	// The deployment status of a plumbing service.
	Toolchain *ServiceStatus `json:"toolchain" validate:"required"`

	// The deployment status of a plumbing service.
	Schematics *ServiceStatus `json:"schematics" validate:"required"`

	// The deployment status of a plumbing service.
	GitRepo *ServiceStatus `json:"git_repo,omitempty"`
}

// UnmarshalServices unmarshals an instance of Services from the specified map of raw messages.
func UnmarshalServices(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Services)
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

// Spec : The spec section includes project definition information, which must be given as input to create the project.
type Spec struct {
	// The list of configurations that make up the project. Each configuration can be managed on its own, augmenting the
	// visibility into the project and the flexibiliy of managing it. However, the user may still encode the entire logic
	// of a project in a big blueprint, in a single config.
	Configs []ConfigIntf `json:"configs" validate:"required"`

	Dashboard *SpecDashboard `json:"dashboard,omitempty"`
}

// NewSpec : Instantiate Spec (Generic Model Constructor)
func (*ProjectsV1) NewSpec(configs []ConfigIntf) (_model *Spec, err error) {
	_model = &Spec{
		Configs: configs,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalSpec unmarshals an instance of Spec from the specified map of raw messages.
func UnmarshalSpec(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Spec)
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "dashboard", &obj.Dashboard, UnmarshalSpecDashboard)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Status : The status section includes the deployment status of the project. It is returned by the status API.
type Status struct {
	Configs []ConfigStatus `json:"configs" validate:"required"`

	ComputedStatuses map[string]interface{} `json:"computed_statuses,omitempty"`

	// Project plumbing services.
	Services *Services `json:"services" validate:"required"`
}

// UnmarshalStatus unmarshals an instance of Status from the specified map of raw messages.
func UnmarshalStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Status)
	err = core.UnmarshalModel(m, "configs", &obj.Configs, UnmarshalConfigStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "computed_statuses", &obj.ComputedStatuses)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "services", &obj.Services, UnmarshalServices)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TerraformTemplate : A Terraform blueprint to use for provisioning a set of project resources.
type TerraformTemplate struct {
	RepoURL *string `json:"repo_url" validate:"required"`

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

// ConfigManualProperty : ConfigManualProperty struct
// This model "extends" Config
type ConfigManualProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	Type *string `json:"type" validate:"required"`

	ExternalResourcesAccount *string `json:"external_resources_account,omitempty"`
}

// Constants associated with the ConfigManualProperty.Type property.
const (
	ConfigManualProperty_Type_Manual = "manual"
)

// NewConfigManualProperty : Instantiate ConfigManualProperty (Generic Model Constructor)
func (*ProjectsV1) NewConfigManualProperty(name string, typeVar string) (_model *ConfigManualProperty, err error) {
	_model = &ConfigManualProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ConfigManualProperty) isaConfig() bool {
	return true
}

// UnmarshalConfigManualProperty unmarshals an instance of ConfigManualProperty from the specified map of raw messages.
func UnmarshalConfigManualProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigManualProperty)
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

// ConfigSchematicsBlueprintProperty : ConfigSchematicsBlueprintProperty struct
// This model "extends" Config
type ConfigSchematicsBlueprintProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input" validate:"required"`

	// A Schematics blueprint to use for provisioning a set of project resources.
	SchematicsBlueprint *SchematicsBlueprint `json:"schematics_blueprint,omitempty"`
}

// Constants associated with the ConfigSchematicsBlueprintProperty.Type property.
const (
	ConfigSchematicsBlueprintProperty_Type_SchematicsBlueprint = "schematics_blueprint"
)

// NewConfigSchematicsBlueprintProperty : Instantiate ConfigSchematicsBlueprintProperty (Generic Model Constructor)
func (*ProjectsV1) NewConfigSchematicsBlueprintProperty(name string, typeVar string, input []InputVariable) (_model *ConfigSchematicsBlueprintProperty, err error) {
	_model = &ConfigSchematicsBlueprintProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Input: input,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ConfigSchematicsBlueprintProperty) isaConfig() bool {
	return true
}

// UnmarshalConfigSchematicsBlueprintProperty unmarshals an instance of ConfigSchematicsBlueprintProperty from the specified map of raw messages.
func UnmarshalConfigSchematicsBlueprintProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigSchematicsBlueprintProperty)
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

// ConfigTerraformTemplateProperty : ConfigTerraformTemplateProperty struct
// This model "extends" Config
type ConfigTerraformTemplateProperty struct {
	Name *string `json:"name" validate:"required"`

	Labels []string `json:"labels,omitempty"`

	Output []OutputValue `json:"output,omitempty"`

	Type *string `json:"type" validate:"required"`

	Input []InputVariable `json:"input" validate:"required"`

	// A Terraform blueprint to use for provisioning a set of project resources.
	TerraformTemplate *TerraformTemplate `json:"terraform_template,omitempty"`
}

// Constants associated with the ConfigTerraformTemplateProperty.Type property.
const (
	ConfigTerraformTemplateProperty_Type_TerraformTemplate = "terraform_template"
)

// NewConfigTerraformTemplateProperty : Instantiate ConfigTerraformTemplateProperty (Generic Model Constructor)
func (*ProjectsV1) NewConfigTerraformTemplateProperty(name string, typeVar string, input []InputVariable) (_model *ConfigTerraformTemplateProperty, err error) {
	_model = &ConfigTerraformTemplateProperty{
		Name: core.StringPtr(name),
		Type: core.StringPtr(typeVar),
		Input: input,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ConfigTerraformTemplateProperty) isaConfig() bool {
	return true
}

// UnmarshalConfigTerraformTemplateProperty unmarshals an instance of ConfigTerraformTemplateProperty from the specified map of raw messages.
func UnmarshalConfigTerraformTemplateProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfigTerraformTemplateProperty)
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
