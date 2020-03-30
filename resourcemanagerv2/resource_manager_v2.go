/**
 * (C) Copyright IBM Corp. 2020.
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

// Package resourcemanagerv2 : Operations and models for the ResourceManagerV2 service
package resourcemanagerv2

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// ResourceManagerV2 : Manage lifecycle of your Cloud resource groups using Resource Manager APIs.
//
// Version: 2.0
type ResourceManagerV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://resource-controller.test.cloud.ibm.com/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "resource_manager"

// ResourceManagerV2Options : Service options
type ResourceManagerV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewResourceManagerV2UsingExternalConfig : constructs an instance of ResourceManagerV2 with passed in options and external configuration.
func NewResourceManagerV2UsingExternalConfig(options *ResourceManagerV2Options) (resourceManager *ResourceManagerV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	resourceManager, err = NewResourceManagerV2(options)
	if err != nil {
		return
	}

	err = resourceManager.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = resourceManager.Service.SetServiceURL(options.URL)
	}
	return
}

// NewResourceManagerV2 : constructs an instance of ResourceManagerV2 with passed in options.
func NewResourceManagerV2(options *ResourceManagerV2Options) (service *ResourceManagerV2, err error) {
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

	service = &ResourceManagerV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (resourceManager *ResourceManagerV2) SetServiceURL(url string) error {
	return resourceManager.Service.SetServiceURL(url)
}

// GetAccountQuotaList : Retrieve a list of quotas in an account
// Retrieve a list of all "merged" quota definition for the specified account.You can use the IAM service token or a
// user token for authorization. To use this method, the requesting user or service ID must have at least the viewer,
// editor, operator, or administrator role on the Resource Controller service. The merge operation takes the default
// quota definition for each resource type, and if there is a delta override quota definition defined for the specified
// account, the override values will replace any corresponding values in the default quota specification.
func (resourceManager *ResourceManagerV2) GetAccountQuotaList(getAccountQuotaListOptions *GetAccountQuotaListOptions) (result *QuotaDefinitions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountQuotaListOptions, "getAccountQuotaListOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountQuotaListOptions, "getAccountQuotaListOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/accounts"}
	pathParameters := []string{*getAccountQuotaListOptions.AccountID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountQuotaListOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "GetAccountQuotaList")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalQuotaDefinitions(m)
		response.Result = result
	}

	return
}

// GetResourceQuota : Retrieve a quota of a resource in an account
// Retrieve a "merged" quota definition for the specified account and type of resource. You can use the IAM service
// token or a user token for authorization. To use this method, the requesting user or service ID must have at least the
// viewer, editor, operator, or administrator role on the Resource Controller service. The merge operation takes the
// default quota definition for the specified resource type, and if there is a delta override quota definition defined
// for the specified account, the override values will replace any corresponding values in the default quota
// specification.
func (resourceManager *ResourceManagerV2) GetResourceQuota(getResourceQuotaOptions *GetResourceQuotaOptions) (result *ResourceQuota, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceQuotaOptions, "getResourceQuotaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceQuotaOptions, "getResourceQuotaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/accounts", "resource_types"}
	pathParameters := []string{*getResourceQuotaOptions.AccountID, *getResourceQuotaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "GetResourceQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceQuota(m)
		response.Result = result
	}

	return
}

// UpdateResourceQuota : Update a quota for a resource in an account
// Override a quota definition for a specified account and type of resource, replacing any previous quota definition
// defined for the account and resource type. You can use the IAM service token or a user token for authorization. To
// use this method, the requesting user or service ID must have at least editor or administrator role on the Resource
// Controller service.
func (resourceManager *ResourceManagerV2) UpdateResourceQuota(updateResourceQuotaOptions *UpdateResourceQuotaOptions) (result *ErrorCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateResourceQuotaOptions, "updateResourceQuotaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateResourceQuotaOptions, "updateResourceQuotaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/accounts", "resource_types"}
	pathParameters := []string{*updateResourceQuotaOptions.AccountID, *updateResourceQuotaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateResourceQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "UpdateResourceQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalErrorCollection(m)
		response.Result = result
	}

	return
}

// DeleteResourceQuota : Delete a quota for a resource in an account
// Deletes a quota definition if one has been defined for a specified account and type of resource, the default quota
// values for the resource type will still be effective. You can use the IAM service token or a user token for
// authorization. To use this method, the requesting user or service ID must have administrator role on the Resource
// Controller service.
func (resourceManager *ResourceManagerV2) DeleteResourceQuota(deleteResourceQuotaOptions *DeleteResourceQuotaOptions) (result *ErrorCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteResourceQuotaOptions, "deleteResourceQuotaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteResourceQuotaOptions, "deleteResourceQuotaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/accounts", "resource_types"}
	pathParameters := []string{*deleteResourceQuotaOptions.AccountID, *deleteResourceQuotaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteResourceQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "DeleteResourceQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalErrorCollection(m)
		response.Result = result
	}

	return
}

// CreateDefaultResourceQuota : Create a default resource quota for a resource type
// Create a default resource quota for a resource type. You can use the IAM service token or a user token for
// authorization. To use this method, the requesting user or service ID must have at least editor or administrator role
// on the Resource Controller service.
func (resourceManager *ResourceManagerV2) CreateDefaultResourceQuota(createDefaultResourceQuotaOptions *CreateDefaultResourceQuotaOptions) (result *ErrorCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDefaultResourceQuotaOptions, "createDefaultResourceQuotaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDefaultResourceQuotaOptions, "createDefaultResourceQuotaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/resource_types"}
	pathParameters := []string{*createDefaultResourceQuotaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDefaultResourceQuotaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "CreateDefaultResourceQuota")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalErrorCollection(m)
		response.Result = result
	}

	return
}

// CreateSchema : Create a schema for a resource type
// Create a schema for a resource type. You can use the IAM service token or a user token for authorization. To use this
// method, the requesting user or service ID must have at least editor or administrator role on the Resource Controller
// service.
func (resourceManager *ResourceManagerV2) CreateSchema(createSchemaOptions *CreateSchemaOptions) (result *ErrorCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSchemaOptions, "createSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSchemaOptions, "createSchemaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/resource_types", "schemas"}
	pathParameters := []string{*createSchemaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "CreateSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalErrorCollection(m)
		response.Result = result
	}

	return
}

// GetSchema : Retrieve a schema for a resource type
// Retrieve a schema for a resource type. You can use the IAM service token or a user token for authorization. To use
// this method, the requesting user or service ID must have at least editor or administrator role on the Resource
// Controller service.
func (resourceManager *ResourceManagerV2) GetSchema(getSchemaOptions *GetSchemaOptions) (result *ResourceQuota, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSchemaOptions, "getSchemaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSchemaOptions, "getSchemaOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions/resource_types", "schemas"}
	pathParameters := []string{*getSchemaOptions.ResourceType}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSchemaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "GetSchema")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceQuota(m)
		response.Result = result
	}

	return
}

// ListQuotaDefinitions : Get a list of all quota definitions
// Get a list of all quota definitions.
func (resourceManager *ResourceManagerV2) ListQuotaDefinitions(listQuotaDefinitionsOptions *ListQuotaDefinitionsOptions) (result *QuotaDefinitionsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listQuotaDefinitionsOptions, "listQuotaDefinitionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listQuotaDefinitionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "ListQuotaDefinitions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalQuotaDefinitionsList(m)
		response.Result = result
	}

	return
}

// GetQuotaDefinition : Get a quota definition
// Get a a quota definition.
func (resourceManager *ResourceManagerV2) GetQuotaDefinition(getQuotaDefinitionOptions *GetQuotaDefinitionOptions) (result *QuotaDefinitions, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getQuotaDefinitionOptions, "getQuotaDefinitionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getQuotaDefinitionOptions, "getQuotaDefinitionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"quota_definitions"}
	pathParameters := []string{*getQuotaDefinitionOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getQuotaDefinitionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "GetQuotaDefinition")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalQuotaDefinitions(m)
		response.Result = result
	}

	return
}

// ListResourceGroups : Get a list of all resource groups
// Get a list of all resource groups in an account.
func (resourceManager *ResourceManagerV2) ListResourceGroups(listResourceGroupsOptions *ListResourceGroupsOptions) (result *ResourceGroupsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listResourceGroupsOptions, "listResourceGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourceGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "ListResourceGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listResourceGroupsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listResourceGroupsOptions.AccountID))
	}
	if listResourceGroupsOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*listResourceGroupsOptions.Date))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceGroupsList(m)
		response.Result = result
	}

	return
}

// CreateResourceGroup : Create a new resource group
// Create a new resource group in an account.
func (resourceManager *ResourceManagerV2) CreateResourceGroup(createResourceGroupOptions *CreateResourceGroupOptions) (result *ResCreateResourceGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createResourceGroupOptions, "createResourceGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "CreateResourceGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createResourceGroupOptions.Name != nil {
		body["name"] = createResourceGroupOptions.Name
	}
	if createResourceGroupOptions.AccountID != nil {
		body["account_id"] = createResourceGroupOptions.AccountID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResCreateResourceGroup(m)
		response.Result = result
	}

	return
}

// GetResourceGroup : Get a resource group
// Retrieve a resource group by ID.
func (resourceManager *ResourceManagerV2) GetResourceGroup(getResourceGroupOptions *GetResourceGroupOptions) (result *ResourceGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceGroupOptions, "getResourceGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceGroupOptions, "getResourceGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_groups"}
	pathParameters := []string{*getResourceGroupOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "GetResourceGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceGroup(m)
		response.Result = result
	}

	return
}

// UpdateResourceGroup : Update a resource group
// Update a resource group by ID.
func (resourceManager *ResourceManagerV2) UpdateResourceGroup(updateResourceGroupOptions *UpdateResourceGroupOptions) (result *ResourceGroup, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateResourceGroupOptions, "updateResourceGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateResourceGroupOptions, "updateResourceGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_groups"}
	pathParameters := []string{*updateResourceGroupOptions.ID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateResourceGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "UpdateResourceGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateResourceGroupOptions.Name != nil {
		body["name"] = updateResourceGroupOptions.Name
	}
	if updateResourceGroupOptions.State != nil {
		body["state"] = updateResourceGroupOptions.State
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceGroup(m)
		response.Result = result
	}

	return
}

// DeleteResourceGroup : Delete a resource group
// Delete a resource group by ID.
func (resourceManager *ResourceManagerV2) DeleteResourceGroup(deleteResourceGroupOptions *DeleteResourceGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteResourceGroupOptions, "deleteResourceGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteResourceGroupOptions, "deleteResourceGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"resource_groups"}
	pathParameters := []string{*deleteResourceGroupOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(resourceManager.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteResourceGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("resource_manager", "V2", "DeleteResourceGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = resourceManager.Service.Request(request, nil)

	return
}

// CreateDefaultResourceQuotaOptions : The CreateDefaultResourceQuota options.
type CreateDefaultResourceQuotaOptions struct {
	// The resource type.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDefaultResourceQuotaOptions : Instantiate CreateDefaultResourceQuotaOptions
func (*ResourceManagerV2) NewCreateDefaultResourceQuotaOptions(resourceType string) *CreateDefaultResourceQuotaOptions {
	return &CreateDefaultResourceQuotaOptions{
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetResourceType : Allow user to set ResourceType
func (options *CreateDefaultResourceQuotaOptions) SetResourceType(resourceType string) *CreateDefaultResourceQuotaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDefaultResourceQuotaOptions) SetHeaders(param map[string]string) *CreateDefaultResourceQuotaOptions {
	options.Headers = param
	return options
}

// CreateResourceGroupOptions : The CreateResourceGroup options.
type CreateResourceGroupOptions struct {
	// The new name of the resource group.
	Name *string `json:"name,omitempty"`

	// The account id of the resource group.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateResourceGroupOptions : Instantiate CreateResourceGroupOptions
func (*ResourceManagerV2) NewCreateResourceGroupOptions() *CreateResourceGroupOptions {
	return &CreateResourceGroupOptions{}
}

// SetName : Allow user to set Name
func (options *CreateResourceGroupOptions) SetName(name string) *CreateResourceGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *CreateResourceGroupOptions) SetAccountID(accountID string) *CreateResourceGroupOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceGroupOptions) SetHeaders(param map[string]string) *CreateResourceGroupOptions {
	options.Headers = param
	return options
}

// CreateSchemaOptions : The CreateSchema options.
type CreateSchemaOptions struct {
	// The resource type.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSchemaOptions : Instantiate CreateSchemaOptions
func (*ResourceManagerV2) NewCreateSchemaOptions(resourceType string) *CreateSchemaOptions {
	return &CreateSchemaOptions{
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetResourceType : Allow user to set ResourceType
func (options *CreateSchemaOptions) SetResourceType(resourceType string) *CreateSchemaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSchemaOptions) SetHeaders(param map[string]string) *CreateSchemaOptions {
	options.Headers = param
	return options
}

// DeleteResourceGroupOptions : The DeleteResourceGroup options.
type DeleteResourceGroupOptions struct {
	// The short or long ID of the alias.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteResourceGroupOptions : Instantiate DeleteResourceGroupOptions
func (*ResourceManagerV2) NewDeleteResourceGroupOptions(id string) *DeleteResourceGroupOptions {
	return &DeleteResourceGroupOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteResourceGroupOptions) SetID(id string) *DeleteResourceGroupOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteResourceGroupOptions) SetHeaders(param map[string]string) *DeleteResourceGroupOptions {
	options.Headers = param
	return options
}

// DeleteResourceQuotaOptions : The DeleteResourceQuota options.
type DeleteResourceQuotaOptions struct {
	// The account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The type of resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteResourceQuotaOptions : Instantiate DeleteResourceQuotaOptions
func (*ResourceManagerV2) NewDeleteResourceQuotaOptions(accountID string, resourceType string) *DeleteResourceQuotaOptions {
	return &DeleteResourceQuotaOptions{
		AccountID: core.StringPtr(accountID),
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *DeleteResourceQuotaOptions) SetAccountID(accountID string) *DeleteResourceQuotaOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceType : Allow user to set ResourceType
func (options *DeleteResourceQuotaOptions) SetResourceType(resourceType string) *DeleteResourceQuotaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteResourceQuotaOptions) SetHeaders(param map[string]string) *DeleteResourceQuotaOptions {
	options.Headers = param
	return options
}

// ErrorCollection : Collection of errors.
type ErrorCollection struct {
	// The error code encountered.
	ErrorCode *string `json:"error_code,omitempty"`

	// The error message.
	Message *string `json:"message,omitempty"`

	// The status code.
	StatusCode *string `json:"status_code,omitempty"`

	// The transaction-id of the request.
	TransactionID *string `json:"transaction_id,omitempty"`
}


// UnmarshalErrorCollection constructs an instance of ErrorCollection from the specified map.
func UnmarshalErrorCollection(m map[string]interface{}) (result *ErrorCollection, err error) {
	obj := new(ErrorCollection)
	obj.ErrorCode, err = core.UnmarshalString(m, "error_code")
	if err != nil {
		return
	}
	obj.Message, err = core.UnmarshalString(m, "message")
	if err != nil {
		return
	}
	obj.StatusCode, err = core.UnmarshalString(m, "status_code")
	if err != nil {
		return
	}
	obj.TransactionID, err = core.UnmarshalString(m, "transaction_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalErrorCollectionSlice unmarshals a slice of ErrorCollection instances from the specified list of maps.
func UnmarshalErrorCollectionSlice(s []interface{}) (slice []ErrorCollection, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ErrorCollection'")
			return
		}
		obj, e := UnmarshalErrorCollection(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalErrorCollectionAsProperty unmarshals an instance of ErrorCollection that is stored as a property
// within the specified map.
func UnmarshalErrorCollectionAsProperty(m map[string]interface{}, propertyName string) (result *ErrorCollection, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ErrorCollection'", propertyName)
			return
		}
		result, err = UnmarshalErrorCollection(objMap)
	}
	return
}

// UnmarshalErrorCollectionSliceAsProperty unmarshals a slice of ErrorCollection instances that are stored as a property
// within the specified map.
func UnmarshalErrorCollectionSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ErrorCollection, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ErrorCollection'", propertyName)
			return
		}
		slice, err = UnmarshalErrorCollectionSlice(vSlice)
	}
	return
}

// GetAccountQuotaListOptions : The GetAccountQuotaList options.
type GetAccountQuotaListOptions struct {
	// The account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountQuotaListOptions : Instantiate GetAccountQuotaListOptions
func (*ResourceManagerV2) NewGetAccountQuotaListOptions(accountID string) *GetAccountQuotaListOptions {
	return &GetAccountQuotaListOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountQuotaListOptions) SetAccountID(accountID string) *GetAccountQuotaListOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountQuotaListOptions) SetHeaders(param map[string]string) *GetAccountQuotaListOptions {
	options.Headers = param
	return options
}

// GetQuotaDefinitionOptions : The GetQuotaDefinition options.
type GetQuotaDefinitionOptions struct {
	// The id of the quota.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetQuotaDefinitionOptions : Instantiate GetQuotaDefinitionOptions
func (*ResourceManagerV2) NewGetQuotaDefinitionOptions(id string) *GetQuotaDefinitionOptions {
	return &GetQuotaDefinitionOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetQuotaDefinitionOptions) SetID(id string) *GetQuotaDefinitionOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetQuotaDefinitionOptions) SetHeaders(param map[string]string) *GetQuotaDefinitionOptions {
	options.Headers = param
	return options
}

// GetResourceGroupOptions : The GetResourceGroup options.
type GetResourceGroupOptions struct {
	// The short or long ID of the alias.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceGroupOptions : Instantiate GetResourceGroupOptions
func (*ResourceManagerV2) NewGetResourceGroupOptions(id string) *GetResourceGroupOptions {
	return &GetResourceGroupOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetResourceGroupOptions) SetID(id string) *GetResourceGroupOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceGroupOptions) SetHeaders(param map[string]string) *GetResourceGroupOptions {
	options.Headers = param
	return options
}

// GetResourceQuotaOptions : The GetResourceQuota options.
type GetResourceQuotaOptions struct {
	// The account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The type of resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceQuotaOptions : Instantiate GetResourceQuotaOptions
func (*ResourceManagerV2) NewGetResourceQuotaOptions(accountID string, resourceType string) *GetResourceQuotaOptions {
	return &GetResourceQuotaOptions{
		AccountID: core.StringPtr(accountID),
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceQuotaOptions) SetAccountID(accountID string) *GetResourceQuotaOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceType : Allow user to set ResourceType
func (options *GetResourceQuotaOptions) SetResourceType(resourceType string) *GetResourceQuotaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceQuotaOptions) SetHeaders(param map[string]string) *GetResourceQuotaOptions {
	options.Headers = param
	return options
}

// GetSchemaOptions : The GetSchema options.
type GetSchemaOptions struct {
	// The resource type.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSchemaOptions : Instantiate GetSchemaOptions
func (*ResourceManagerV2) NewGetSchemaOptions(resourceType string) *GetSchemaOptions {
	return &GetSchemaOptions{
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetResourceType : Allow user to set ResourceType
func (options *GetSchemaOptions) SetResourceType(resourceType string) *GetSchemaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSchemaOptions) SetHeaders(param map[string]string) *GetSchemaOptions {
	options.Headers = param
	return options
}

// ListQuotaDefinitionsOptions : The ListQuotaDefinitions options.
type ListQuotaDefinitionsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListQuotaDefinitionsOptions : Instantiate ListQuotaDefinitionsOptions
func (*ResourceManagerV2) NewListQuotaDefinitionsOptions() *ListQuotaDefinitionsOptions {
	return &ListQuotaDefinitionsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListQuotaDefinitionsOptions) SetHeaders(param map[string]string) *ListQuotaDefinitionsOptions {
	options.Headers = param
	return options
}

// ListResourceGroupsOptions : The ListResourceGroups options.
type ListResourceGroupsOptions struct {
	// The ID of the account that contains the resource groups that you want to get.
	AccountID *string `json:"account_id,omitempty"`

	// The date would be in a format of YYYY-MM which returns resource groups exclude the deleted ones before this month.
	Date *string `json:"date,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourceGroupsOptions : Instantiate ListResourceGroupsOptions
func (*ResourceManagerV2) NewListResourceGroupsOptions() *ListResourceGroupsOptions {
	return &ListResourceGroupsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (options *ListResourceGroupsOptions) SetAccountID(accountID string) *ListResourceGroupsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetDate : Allow user to set Date
func (options *ListResourceGroupsOptions) SetDate(date string) *ListResourceGroupsOptions {
	options.Date = core.StringPtr(date)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourceGroupsOptions) SetHeaders(param map[string]string) *ListResourceGroupsOptions {
	options.Headers = param
	return options
}

// QuotaDefinitions : Returned quota definitions.
type QuotaDefinitions struct {
	// An alpha-numeric value identifying the quota.
	ID *string `json:"id,omitempty"`

	// The human-readable name of the quota.
	Name *string `json:"name,omitempty"`

	// The type of the quota.
	Type *string `json:"type,omitempty"`

	// The total app limit.
	NumberOfApps *float64 `json:"number_of_apps,omitempty"`

	// The total service instances limit per app.
	NumberOfServiceInstances *float64 `json:"number_of_service_instances,omitempty"`

	// Default number of instances per lite plan.
	DefaultNumberOfInstancesPerLitePlan *float64 `json:"default_number_of_instances_per_lite_plan,omitempty"`

	// The total instances limit per app.
	InstancesPerApp *float64 `json:"instances_per_app,omitempty"`

	// The total memory of app instance.
	InstanceMemory *string `json:"instance_memory,omitempty"`

	// The total app memory capacity.
	TotalAppMemory *string `json:"total_app_memory,omitempty"`

	// The VSI limit.
	VsiLimit *float64 `json:"vsi_limit,omitempty"`

	// Resource quota.
	ResourceQuotas *ResourceQuota `json:"resource_quotas,omitempty"`

	// The date when the quota was initially created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The date when the quota was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}


// UnmarshalQuotaDefinitions constructs an instance of QuotaDefinitions from the specified map.
func UnmarshalQuotaDefinitions(m map[string]interface{}) (result *QuotaDefinitions, err error) {
	obj := new(QuotaDefinitions)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.NumberOfApps, err = core.UnmarshalFloat64(m, "number_of_apps")
	if err != nil {
		return
	}
	obj.NumberOfServiceInstances, err = core.UnmarshalFloat64(m, "number_of_service_instances")
	if err != nil {
		return
	}
	obj.DefaultNumberOfInstancesPerLitePlan, err = core.UnmarshalFloat64(m, "default_number_of_instances_per_lite_plan")
	if err != nil {
		return
	}
	obj.InstancesPerApp, err = core.UnmarshalFloat64(m, "instances_per_app")
	if err != nil {
		return
	}
	obj.InstanceMemory, err = core.UnmarshalString(m, "instance_memory")
	if err != nil {
		return
	}
	obj.TotalAppMemory, err = core.UnmarshalString(m, "total_app_memory")
	if err != nil {
		return
	}
	obj.VsiLimit, err = core.UnmarshalFloat64(m, "vsi_limit")
	if err != nil {
		return
	}
	obj.ResourceQuotas, err = UnmarshalResourceQuotaAsProperty(m, "resource_quotas")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.UpdatedAt, err = core.UnmarshalString(m, "updated_at")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalQuotaDefinitionsSlice unmarshals a slice of QuotaDefinitions instances from the specified list of maps.
func UnmarshalQuotaDefinitionsSlice(s []interface{}) (slice []QuotaDefinitions, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'QuotaDefinitions'")
			return
		}
		obj, e := UnmarshalQuotaDefinitions(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalQuotaDefinitionsAsProperty unmarshals an instance of QuotaDefinitions that is stored as a property
// within the specified map.
func UnmarshalQuotaDefinitionsAsProperty(m map[string]interface{}, propertyName string) (result *QuotaDefinitions, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'QuotaDefinitions'", propertyName)
			return
		}
		result, err = UnmarshalQuotaDefinitions(objMap)
	}
	return
}

// UnmarshalQuotaDefinitionsSliceAsProperty unmarshals a slice of QuotaDefinitions instances that are stored as a property
// within the specified map.
func UnmarshalQuotaDefinitionsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []QuotaDefinitions, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'QuotaDefinitions'", propertyName)
			return
		}
		slice, err = UnmarshalQuotaDefinitionsSlice(vSlice)
	}
	return
}

// QuotaDefinitionsList : A list of quota definitions.
type QuotaDefinitionsList struct {
	// List of resources.
	Resources []QuotaDefinitions `json:"resources" validate:"required"`
}


// UnmarshalQuotaDefinitionsList constructs an instance of QuotaDefinitionsList from the specified map.
func UnmarshalQuotaDefinitionsList(m map[string]interface{}) (result *QuotaDefinitionsList, err error) {
	obj := new(QuotaDefinitionsList)
	obj.Resources, err = UnmarshalQuotaDefinitionsSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalQuotaDefinitionsListSlice unmarshals a slice of QuotaDefinitionsList instances from the specified list of maps.
func UnmarshalQuotaDefinitionsListSlice(s []interface{}) (slice []QuotaDefinitionsList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'QuotaDefinitionsList'")
			return
		}
		obj, e := UnmarshalQuotaDefinitionsList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalQuotaDefinitionsListAsProperty unmarshals an instance of QuotaDefinitionsList that is stored as a property
// within the specified map.
func UnmarshalQuotaDefinitionsListAsProperty(m map[string]interface{}, propertyName string) (result *QuotaDefinitionsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'QuotaDefinitionsList'", propertyName)
			return
		}
		result, err = UnmarshalQuotaDefinitionsList(objMap)
	}
	return
}

// UnmarshalQuotaDefinitionsListSliceAsProperty unmarshals a slice of QuotaDefinitionsList instances that are stored as a property
// within the specified map.
func UnmarshalQuotaDefinitionsListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []QuotaDefinitionsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'QuotaDefinitionsList'", propertyName)
			return
		}
		slice, err = UnmarshalQuotaDefinitionsListSlice(vSlice)
	}
	return
}

// ResCreateResourceGroup : Create a resource group.
type ResCreateResourceGroup struct {
	// An alpha-numeric value identifying the resource group.
	ID *string `json:"id,omitempty"`

	// The full CRN (cloud resource name) associated with the resource group. For more on this format, see [Cloud Resource
	// Names](https://cloud.ibm.com/docs/resources?topic=resources-crn).
	Crn *string `json:"crn,omitempty"`
}


// UnmarshalResCreateResourceGroup constructs an instance of ResCreateResourceGroup from the specified map.
func UnmarshalResCreateResourceGroup(m map[string]interface{}) (result *ResCreateResourceGroup, err error) {
	obj := new(ResCreateResourceGroup)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResCreateResourceGroupSlice unmarshals a slice of ResCreateResourceGroup instances from the specified list of maps.
func UnmarshalResCreateResourceGroupSlice(s []interface{}) (slice []ResCreateResourceGroup, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResCreateResourceGroup'")
			return
		}
		obj, e := UnmarshalResCreateResourceGroup(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResCreateResourceGroupAsProperty unmarshals an instance of ResCreateResourceGroup that is stored as a property
// within the specified map.
func UnmarshalResCreateResourceGroupAsProperty(m map[string]interface{}, propertyName string) (result *ResCreateResourceGroup, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResCreateResourceGroup'", propertyName)
			return
		}
		result, err = UnmarshalResCreateResourceGroup(objMap)
	}
	return
}

// UnmarshalResCreateResourceGroupSliceAsProperty unmarshals a slice of ResCreateResourceGroup instances that are stored as a property
// within the specified map.
func UnmarshalResCreateResourceGroupSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResCreateResourceGroup, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResCreateResourceGroup'", propertyName)
			return
		}
		slice, err = UnmarshalResCreateResourceGroupSlice(vSlice)
	}
	return
}

// ResourceGroup : A resource group.
type ResourceGroup struct {
	// An alpha-numeric value identifying the resource group.
	ID *string `json:"id,omitempty"`

	// The full CRN (cloud resource name) associated with the resource group. For more on this format, see [Cloud Resource
	// Names](https://cloud.ibm.com/docs/resources?topic=resources-crn).
	Crn *string `json:"crn,omitempty"`

	// An alpha-numeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The human-readable name of the resource group.
	Name *string `json:"name,omitempty"`

	// The state of the resource group.
	State *string `json:"state,omitempty"`

	// Identify if this resource group is default of the account or not.
	Default *bool `json:"default,omitempty"`

	// An alpha-numeric value identifying the quota ID associated with the resource group.
	QuotaID *string `json:"quota_id,omitempty"`

	// The URL to access the quota details that associated with the resource group.
	QuotaURL *string `json:"quota_url,omitempty"`

	// The URL to access the payment methods details that associated with the resource group.
	PaymentMethodsURL *string `json:"payment_methods_url,omitempty"`

	// An array of the resources that linked to the resource group.
	ResourceLinkages []interface{} `json:"resource_linkages,omitempty"`

	// The URL to access the team details that associated with the resource group.
	TeamsURL *string `json:"teams_url,omitempty"`

	// The date when the resource group was initially created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The date when the resource group was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}


// UnmarshalResourceGroup constructs an instance of ResourceGroup from the specified map.
func UnmarshalResourceGroup(m map[string]interface{}) (result *ResourceGroup, err error) {
	obj := new(ResourceGroup)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.Default, err = core.UnmarshalBool(m, "default")
	if err != nil {
		return
	}
	obj.QuotaID, err = core.UnmarshalString(m, "quota_id")
	if err != nil {
		return
	}
	obj.QuotaURL, err = core.UnmarshalString(m, "quota_url")
	if err != nil {
		return
	}
	obj.PaymentMethodsURL, err = core.UnmarshalString(m, "payment_methods_url")
	if err != nil {
		return
	}
	obj.ResourceLinkages, err = core.UnmarshalAnySlice(m, "resource_linkages")
	if err != nil {
		return
	}
	obj.TeamsURL, err = core.UnmarshalString(m, "teams_url")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.UpdatedAt, err = core.UnmarshalString(m, "updated_at")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceGroupSlice unmarshals a slice of ResourceGroup instances from the specified list of maps.
func UnmarshalResourceGroupSlice(s []interface{}) (slice []ResourceGroup, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceGroup'")
			return
		}
		obj, e := UnmarshalResourceGroup(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceGroupAsProperty unmarshals an instance of ResourceGroup that is stored as a property
// within the specified map.
func UnmarshalResourceGroupAsProperty(m map[string]interface{}, propertyName string) (result *ResourceGroup, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceGroup'", propertyName)
			return
		}
		result, err = UnmarshalResourceGroup(objMap)
	}
	return
}

// UnmarshalResourceGroupSliceAsProperty unmarshals a slice of ResourceGroup instances that are stored as a property
// within the specified map.
func UnmarshalResourceGroupSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceGroup, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceGroup'", propertyName)
			return
		}
		slice, err = UnmarshalResourceGroupSlice(vSlice)
	}
	return
}

// ResourceGroupsList : A list of resource groups.
type ResourceGroupsList struct {
	// List of resources.
	Resources []ResourceGroup `json:"resources" validate:"required"`
}


// UnmarshalResourceGroupsList constructs an instance of ResourceGroupsList from the specified map.
func UnmarshalResourceGroupsList(m map[string]interface{}) (result *ResourceGroupsList, err error) {
	obj := new(ResourceGroupsList)
	obj.Resources, err = UnmarshalResourceGroupSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceGroupsListSlice unmarshals a slice of ResourceGroupsList instances from the specified list of maps.
func UnmarshalResourceGroupsListSlice(s []interface{}) (slice []ResourceGroupsList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceGroupsList'")
			return
		}
		obj, e := UnmarshalResourceGroupsList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceGroupsListAsProperty unmarshals an instance of ResourceGroupsList that is stored as a property
// within the specified map.
func UnmarshalResourceGroupsListAsProperty(m map[string]interface{}, propertyName string) (result *ResourceGroupsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceGroupsList'", propertyName)
			return
		}
		result, err = UnmarshalResourceGroupsList(objMap)
	}
	return
}

// UnmarshalResourceGroupsListSliceAsProperty unmarshals a slice of ResourceGroupsList instances that are stored as a property
// within the specified map.
func UnmarshalResourceGroupsListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceGroupsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceGroupsList'", propertyName)
			return
		}
		slice, err = UnmarshalResourceGroupsListSlice(vSlice)
	}
	return
}

// ResourceQuota : Resource quota.
type ResourceQuota struct {
	// An alpha-numeric value identifying the quota.
	ID *string `json:"_id,omitempty"`

	// The human-readable name of the quota.
	ResourceID *string `json:"resource_id,omitempty"`

	// The full CRN (cloud resource name) associated with the quota. For more on this format, see
	// https://console.bluemix.net/docs/overview/crn.html#crn.
	Crn *string `json:"crn,omitempty"`

	// The limit number of this resource.
	Limit *float64 `json:"limit,omitempty"`
}


// UnmarshalResourceQuota constructs an instance of ResourceQuota from the specified map.
func UnmarshalResourceQuota(m map[string]interface{}) (result *ResourceQuota, err error) {
	obj := new(ResourceQuota)
	obj.ID, err = core.UnmarshalString(m, "_id")
	if err != nil {
		return
	}
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalFloat64(m, "limit")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceQuotaSlice unmarshals a slice of ResourceQuota instances from the specified list of maps.
func UnmarshalResourceQuotaSlice(s []interface{}) (slice []ResourceQuota, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceQuota'")
			return
		}
		obj, e := UnmarshalResourceQuota(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceQuotaAsProperty unmarshals an instance of ResourceQuota that is stored as a property
// within the specified map.
func UnmarshalResourceQuotaAsProperty(m map[string]interface{}, propertyName string) (result *ResourceQuota, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceQuota'", propertyName)
			return
		}
		result, err = UnmarshalResourceQuota(objMap)
	}
	return
}

// UnmarshalResourceQuotaSliceAsProperty unmarshals a slice of ResourceQuota instances that are stored as a property
// within the specified map.
func UnmarshalResourceQuotaSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceQuota, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceQuota'", propertyName)
			return
		}
		slice, err = UnmarshalResourceQuotaSlice(vSlice)
	}
	return
}

// UpdateResourceGroupOptions : The UpdateResourceGroup options.
type UpdateResourceGroupOptions struct {
	// The short or long ID of the alias.
	ID *string `json:"id" validate:"required"`

	// The new name of the resource group.
	Name *string `json:"name,omitempty"`

	// The state of the resource group.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateResourceGroupOptions : Instantiate UpdateResourceGroupOptions
func (*ResourceManagerV2) NewUpdateResourceGroupOptions(id string) *UpdateResourceGroupOptions {
	return &UpdateResourceGroupOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateResourceGroupOptions) SetID(id string) *UpdateResourceGroupOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateResourceGroupOptions) SetName(name string) *UpdateResourceGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetState : Allow user to set State
func (options *UpdateResourceGroupOptions) SetState(state string) *UpdateResourceGroupOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResourceGroupOptions) SetHeaders(param map[string]string) *UpdateResourceGroupOptions {
	options.Headers = param
	return options
}

// UpdateResourceQuotaOptions : The UpdateResourceQuota options.
type UpdateResourceQuotaOptions struct {
	// The account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The type of resource.
	ResourceType *string `json:"resource_type" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateResourceQuotaOptions : Instantiate UpdateResourceQuotaOptions
func (*ResourceManagerV2) NewUpdateResourceQuotaOptions(accountID string, resourceType string) *UpdateResourceQuotaOptions {
	return &UpdateResourceQuotaOptions{
		AccountID: core.StringPtr(accountID),
		ResourceType: core.StringPtr(resourceType),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateResourceQuotaOptions) SetAccountID(accountID string) *UpdateResourceQuotaOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceType : Allow user to set ResourceType
func (options *UpdateResourceQuotaOptions) SetResourceType(resourceType string) *UpdateResourceQuotaOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResourceQuotaOptions) SetHeaders(param map[string]string) *UpdateResourceQuotaOptions {
	options.Headers = param
	return options
}
