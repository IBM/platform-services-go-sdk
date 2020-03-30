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

// Package iamidentityservicesv1 : Operations and models for the IamIdentityServicesV1 service
package iamidentityservicesv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// IamIdentityServicesV1 : The IAM Identity Service API allows for the management of Identities (Service IDs, ApiKeys).
//
// Version: 1.0
type IamIdentityServicesV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.test.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "iam_identity_services"

// IamIdentityServicesV1Options : Service options
type IamIdentityServicesV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIamIdentityServicesV1UsingExternalConfig : constructs an instance of IamIdentityServicesV1 with passed in options and external configuration.
func NewIamIdentityServicesV1UsingExternalConfig(options *IamIdentityServicesV1Options) (iamIdentityServices *IamIdentityServicesV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	iamIdentityServices, err = NewIamIdentityServicesV1(options)
	if err != nil {
		return
	}

	err = iamIdentityServices.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = iamIdentityServices.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIamIdentityServicesV1 : constructs an instance of IamIdentityServicesV1 with passed in options.
func NewIamIdentityServicesV1(options *IamIdentityServicesV1Options) (service *IamIdentityServicesV1, err error) {
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

	service = &IamIdentityServicesV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (iamIdentityServices *IamIdentityServicesV1) SetServiceURL(url string) error {
	return iamIdentityServices.Service.SetServiceURL(url)
}

// ListApiKeys : Get API keys for a given a service or user IAM ID and account ID
// Returns the list of API key details for a given service or user IAM ID and account ID. Users can manage user API keys
// for themself, or service ID API keys for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) ListApiKeys(listApiKeysOptions *ListApiKeysOptions) (result *ListApiKeysResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listApiKeysOptions, "listApiKeysOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listApiKeysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "ListApiKeys")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listApiKeysOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listApiKeysOptions.AccountID))
	}
	if listApiKeysOptions.IamID != nil {
		builder.AddQuery("iam_id", fmt.Sprint(*listApiKeysOptions.IamID))
	}
	if listApiKeysOptions.Pagesize != nil {
		builder.AddQuery("pagesize", fmt.Sprint(*listApiKeysOptions.Pagesize))
	}
	if listApiKeysOptions.Pagetoken != nil {
		builder.AddQuery("pagetoken", fmt.Sprint(*listApiKeysOptions.Pagetoken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalListApiKeysResponse(m)
		response.Result = result
	}

	return
}

// CreateApiKey : Create an ApiKey
// Creates an API key for a UserID or service ID. Users can manage user API keys for themself, or service ID API keys
// for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) CreateApiKey(createApiKeyOptions *CreateApiKeyOptions) (result *ApiKeyDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createApiKeyOptions, "createApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createApiKeyOptions, "createApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "CreateApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createApiKeyOptions.EntityLock != nil {
		builder.AddHeader("Entity-Lock", fmt.Sprint(*createApiKeyOptions.EntityLock))
	}

	body := make(map[string]interface{})
	if createApiKeyOptions.Name != nil {
		body["name"] = createApiKeyOptions.Name
	}
	if createApiKeyOptions.IamID != nil {
		body["iam_id"] = createApiKeyOptions.IamID
	}
	if createApiKeyOptions.Description != nil {
		body["description"] = createApiKeyOptions.Description
	}
	if createApiKeyOptions.AccountID != nil {
		body["account_id"] = createApiKeyOptions.AccountID
	}
	if createApiKeyOptions.Apikey != nil {
		body["apikey"] = createApiKeyOptions.Apikey
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalApiKeyDetails(m)
		response.Result = result
	}

	return
}

// GetApiKeyDetails : Get details of an API key by its value
// Returns the details of an API key by its value. Users can manage user API keys for themself, or service ID API keys
// for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) GetApiKeyDetails(getApiKeyDetailsOptions *GetApiKeyDetailsOptions) (result *ApiKeyDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getApiKeyDetailsOptions, "getApiKeyDetailsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys/details"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApiKeyDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "GetApiKeyDetails")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getApiKeyDetailsOptions.IAMApiKey != nil {
		builder.AddHeader("IAM-ApiKey", fmt.Sprint(*getApiKeyDetailsOptions.IAMApiKey))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalApiKeyDetails(m)
		response.Result = result
	}

	return
}

// GetApiKey : Get details of an API key
// Returns the details of an API key. Users can manage user API keys for themself, or service ID API keys for service
// IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) GetApiKey(getApiKeyOptions *GetApiKeyOptions) (result *ApiKeyDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApiKeyOptions, "getApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApiKeyOptions, "getApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys"}
	pathParameters := []string{*getApiKeyOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "GetApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalApiKeyDetails(m)
		response.Result = result
	}

	return
}

// UpdateApiKey : Updates an ApiKey
// Updates properties of an API key. This does NOT affect existing access tokens. Their token content will stay
// unchanged until the access token is refreshed. To update an API key, pass the property to be modified. To delete one
// property's value, pass the property with an empty value "".Users can manage user API keys for themself, or service ID
// API keys for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions) (result *ApiKeyDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateApiKeyOptions, "updateApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateApiKeyOptions, "updateApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys"}
	pathParameters := []string{*updateApiKeyOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "UpdateApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateApiKeyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateApiKeyOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateApiKeyOptions.Name != nil {
		body["name"] = updateApiKeyOptions.Name
	}
	if updateApiKeyOptions.Description != nil {
		body["description"] = updateApiKeyOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalApiKeyDetails(m)
		response.Result = result
	}

	return
}

// DeleteApiKey : Deletes an ApiKey
// Deletes an API key. Existing tokens will remain valid until expired. Refresh tokens will not work any more for this
// API key. Users can manage user API keys for themself, or service ID API keys for service IDs that are bound to an
// entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteApiKeyOptions, "deleteApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteApiKeyOptions, "deleteApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys"}
	pathParameters := []string{*deleteApiKeyOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "DeleteApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, nil)

	return
}

// LockApiKey : Lock the API key
// Locks an API key by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that are
// bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) LockApiKey(lockApiKeyOptions *LockApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(lockApiKeyOptions, "lockApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(lockApiKeyOptions, "lockApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys", "lock"}
	pathParameters := []string{*lockApiKeyOptions.ID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range lockApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "LockApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, nil)

	return
}

// UnlockApiKey : Unlock the API key
// Unlocks an API key by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unlockApiKeyOptions, "unlockApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unlockApiKeyOptions, "unlockApiKeyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/apikeys", "lock"}
	pathParameters := []string{*unlockApiKeyOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range unlockApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "UnlockApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, nil)

	return
}

// ListServiceIds : List service IDs
// Returns a list of service IDs. Users can manage user API keys for themself, or service ID API keys for service IDs
// that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions) (result *ServiceIdsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listServiceIdsOptions, "listServiceIdsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listServiceIdsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "ListServiceIds")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listServiceIdsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listServiceIdsOptions.AccountID))
	}
	if listServiceIdsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listServiceIdsOptions.Name))
	}
	if listServiceIdsOptions.Pagesize != nil {
		builder.AddQuery("pagesize", fmt.Sprint(*listServiceIdsOptions.Pagesize))
	}
	if listServiceIdsOptions.Pagetoken != nil {
		builder.AddQuery("pagetoken", fmt.Sprint(*listServiceIdsOptions.Pagetoken))
	}
	if listServiceIdsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listServiceIdsOptions.Sort))
	}
	if listServiceIdsOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listServiceIdsOptions.Order))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdsList(m)
		response.Result = result
	}

	return
}

// CreateServiceID : Create a service ID
// Creates a service ID for an IBM Cloud account. Users can manage user API keys for themself, or service ID API keys
// for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) CreateServiceID(createServiceIdOptions *CreateServiceIdOptions) (result *ServiceIdDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createServiceIdOptions, "createServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createServiceIdOptions, "createServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "CreateServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createServiceIdOptions.EntityLock != nil {
		builder.AddHeader("Entity-Lock", fmt.Sprint(*createServiceIdOptions.EntityLock))
	}

	body := make(map[string]interface{})
	if createServiceIdOptions.AccountID != nil {
		body["account_id"] = createServiceIdOptions.AccountID
	}
	if createServiceIdOptions.Name != nil {
		body["name"] = createServiceIdOptions.Name
	}
	if createServiceIdOptions.Description != nil {
		body["description"] = createServiceIdOptions.Description
	}
	if createServiceIdOptions.UniqueInstanceCrns != nil {
		body["unique_instance_crns"] = createServiceIdOptions.UniqueInstanceCrns
	}
	if createServiceIdOptions.Apikey != nil {
		body["apikey"] = createServiceIdOptions.Apikey
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdDetails(m)
		response.Result = result
	}

	return
}

// GetServiceID : Get details of a service ID
// Returns the details of a service ID. Users can manage user API keys for themself, or service ID API keys for service
// IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) GetServiceID(getServiceIdOptions *GetServiceIdOptions) (result *ServiceIdDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getServiceIdOptions, "getServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getServiceIdOptions, "getServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids"}
	pathParameters := []string{*getServiceIdOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "GetServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdDetails(m)
		response.Result = result
	}

	return
}

// UpdateServiceID : Update service ID
// Updates properties of a service ID. This does NOT affect existing access tokens. Their token content will stay
// unchanged until the access token is refreshed. To update a service ID, pass the property to be modified. To delete
// one property's value, pass the property with an empty value "".Users can manage user API keys for themself, or
// service ID API keys for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions) (result *ServiceIdDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateServiceIdOptions, "updateServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateServiceIdOptions, "updateServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids"}
	pathParameters := []string{*updateServiceIdOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "UpdateServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateServiceIdOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateServiceIdOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateServiceIdOptions.Name != nil {
		body["name"] = updateServiceIdOptions.Name
	}
	if updateServiceIdOptions.Description != nil {
		body["description"] = updateServiceIdOptions.Description
	}
	if updateServiceIdOptions.UniqueInstanceCrns != nil {
		body["unique_instance_crns"] = updateServiceIdOptions.UniqueInstanceCrns
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdDetails(m)
		response.Result = result
	}

	return
}

// DeleteServiceID : Deletes a service ID and associated API keys
// Deletes a service ID and all API keys associated to it. Before deleting the service ID, all associated API keys are
// deleted. In case a Delete Conflict (status code 409) a retry of the request may help as the service ID is only
// deleted if the associated API keys were successfully deleted before. Users can manage user API keys for themself, or
// service ID API keys for service IDs that are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceIdOptions, "deleteServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceIdOptions, "deleteServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids"}
	pathParameters := []string{*deleteServiceIdOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "DeleteServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, nil)

	return
}

// LockServiceID : Lock the service ID
// Locks a service ID by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) LockServiceID(lockServiceIdOptions *LockServiceIdOptions) (result *ServiceIdDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(lockServiceIdOptions, "lockServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(lockServiceIdOptions, "lockServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids", "lock"}
	pathParameters := []string{*lockServiceIdOptions.ID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range lockServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "LockServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdDetails(m)
		response.Result = result
	}

	return
}

// UnlockServiceID : Unlock the service ID
// Unlocks a service ID by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to.
func (iamIdentityServices *IamIdentityServicesV1) UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions) (result *ServiceIdDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unlockServiceIdOptions, "unlockServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unlockServiceIdOptions, "unlockServiceIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/serviceids", "lock"}
	pathParameters := []string{*unlockServiceIdOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamIdentityServices.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range unlockServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity_services", "V1", "UnlockServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentityServices.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalServiceIdDetails(m)
		response.Result = result
	}

	return
}

// ApiKeyDetails : Response body format for API key V1 REST requests.
type ApiKeyDetails struct {
	// Context with key properties for problem determination.
	Context *ResponseContext `json:"context,omitempty"`

	// Unique identifier of this API Key.
	ID *string `json:"id" validate:"required"`

	// Version of the API Key details object. You need to specify this value when updating the API key to avoid stale
	// updates.
	EntityTag *string `json:"entity_tag,omitempty"`

	// Cloud Resource Name of the item. Example Cloud Resource Name:
	// 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::apikey:1234-9012-5678'.
	Crn *string `json:"crn" validate:"required"`

	// The API key cannot be changed if set to true.
	Locked *bool `json:"locked" validate:"required"`

	// If set contains a date time string of the creation date in ISO format.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// If set contains a date time string of the last modification date in ISO format.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist.
	// Access is done via the UUID of the API key.
	Name *string `json:"name" validate:"required"`

	// The optional description of the API key. The 'description' property is only available if a description was provided
	// during a create of an API key.
	Description *string `json:"description,omitempty"`

	// The iam_id that this API key authenticates.
	IamID *string `json:"iam_id" validate:"required"`

	// ID of the account that this API key authenticates for.
	AccountID *string `json:"account_id" validate:"required"`

	// The API key value. For API keys representing a user, this is only returned in the response to the create API key
	// request.
	Apikey *string `json:"apikey" validate:"required"`
}


// UnmarshalApiKeyDetails constructs an instance of ApiKeyDetails from the specified map.
func UnmarshalApiKeyDetails(m map[string]interface{}) (result *ApiKeyDetails, err error) {
	obj := new(ApiKeyDetails)
	obj.Context, err = UnmarshalResponseContextAsProperty(m, "context")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.EntityTag, err = core.UnmarshalString(m, "entity_tag")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.Locked, err = core.UnmarshalBool(m, "locked")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalDateTime(m, "created_at")
	if err != nil {
		return
	}
	obj.ModifiedAt, err = core.UnmarshalDateTime(m, "modified_at")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Apikey, err = core.UnmarshalString(m, "apikey")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalApiKeyDetailsSlice unmarshals a slice of ApiKeyDetails instances from the specified list of maps.
func UnmarshalApiKeyDetailsSlice(s []interface{}) (slice []ApiKeyDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ApiKeyDetails'")
			return
		}
		obj, e := UnmarshalApiKeyDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalApiKeyDetailsAsProperty unmarshals an instance of ApiKeyDetails that is stored as a property
// within the specified map.
func UnmarshalApiKeyDetailsAsProperty(m map[string]interface{}, propertyName string) (result *ApiKeyDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ApiKeyDetails'", propertyName)
			return
		}
		result, err = UnmarshalApiKeyDetails(objMap)
	}
	return
}

// UnmarshalApiKeyDetailsSliceAsProperty unmarshals a slice of ApiKeyDetails instances that are stored as a property
// within the specified map.
func UnmarshalApiKeyDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ApiKeyDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ApiKeyDetails'", propertyName)
			return
		}
		slice, err = UnmarshalApiKeyDetailsSlice(vSlice)
	}
	return
}

// CreateApiKeyOptions : The CreateApiKey options.
type CreateApiKeyOptions struct {
	// Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist.
	// Access is done via the UUID of the API key.
	Name *string `json:"name" validate:"required"`

	// The iam_id that this API key authenticates.
	IamID *string `json:"iam_id" validate:"required"`

	// The optional description of the API key. The 'description' property is only available if a description was provided
	// during a create of an API key.
	Description *string `json:"description,omitempty"`

	// The account ID of the API key.
	AccountID *string `json:"account_id,omitempty"`

	// You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is
	// done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key
	// value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this
	// value.
	Apikey *string `json:"apikey,omitempty"`

	// Indicates if the API key is locked for further write operations. False by default.
	EntityLock *string `json:"Entity-Lock,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateApiKeyOptions : Instantiate CreateApiKeyOptions
func (*IamIdentityServicesV1) NewCreateApiKeyOptions(name string, iamID string) *CreateApiKeyOptions {
	return &CreateApiKeyOptions{
		Name: core.StringPtr(name),
		IamID: core.StringPtr(iamID),
	}
}

// SetName : Allow user to set Name
func (options *CreateApiKeyOptions) SetName(name string) *CreateApiKeyOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetIamID : Allow user to set IamID
func (options *CreateApiKeyOptions) SetIamID(iamID string) *CreateApiKeyOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateApiKeyOptions) SetDescription(description string) *CreateApiKeyOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *CreateApiKeyOptions) SetAccountID(accountID string) *CreateApiKeyOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetApikey : Allow user to set Apikey
func (options *CreateApiKeyOptions) SetApikey(apikey string) *CreateApiKeyOptions {
	options.Apikey = core.StringPtr(apikey)
	return options
}

// SetEntityLock : Allow user to set EntityLock
func (options *CreateApiKeyOptions) SetEntityLock(entityLock string) *CreateApiKeyOptions {
	options.EntityLock = core.StringPtr(entityLock)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateApiKeyOptions) SetHeaders(param map[string]string) *CreateApiKeyOptions {
	options.Headers = param
	return options
}

// CreateApiKeyRequest : Input body parameters for the Create API key V1 REST request.
type CreateApiKeyRequest struct {
	// Name of the API key. The name is not checked for uniqueness. Therefore multiple names with the same value can exist.
	// Access is done via the UUID of the API key.
	Name *string `json:"name" validate:"required"`

	// The optional description of the API key. The 'description' property is only available if a description was provided
	// during a create of an API key.
	Description *string `json:"description,omitempty"`

	// The iam_id that this API key authenticates.
	IamID *string `json:"iam_id" validate:"required"`

	// The account ID of the API key.
	AccountID *string `json:"account_id,omitempty"`

	// You can optionally passthrough the API key value for this API key. If passed, NO validation of that apiKey value is
	// done, i.e. the value can be non-URL safe. If omitted, the API key management will create an URL safe opaque API key
	// value. The value of the API key is checked for uniqueness. Please ensure enough variations when passing in this
	// value.
	Apikey *string `json:"apikey,omitempty"`
}


// NewCreateApiKeyRequest : Instantiate CreateApiKeyRequest (Generic Model Constructor)
func (*IamIdentityServicesV1) NewCreateApiKeyRequest(name string, iamID string) (model *CreateApiKeyRequest, err error) {
	model = &CreateApiKeyRequest{
		Name: core.StringPtr(name),
		IamID: core.StringPtr(iamID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateApiKeyRequest constructs an instance of CreateApiKeyRequest from the specified map.
func UnmarshalCreateApiKeyRequest(m map[string]interface{}) (result *CreateApiKeyRequest, err error) {
	obj := new(CreateApiKeyRequest)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Apikey, err = core.UnmarshalString(m, "apikey")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCreateApiKeyRequestSlice unmarshals a slice of CreateApiKeyRequest instances from the specified list of maps.
func UnmarshalCreateApiKeyRequestSlice(s []interface{}) (slice []CreateApiKeyRequest, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'CreateApiKeyRequest'")
			return
		}
		obj, e := UnmarshalCreateApiKeyRequest(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCreateApiKeyRequestAsProperty unmarshals an instance of CreateApiKeyRequest that is stored as a property
// within the specified map.
func UnmarshalCreateApiKeyRequestAsProperty(m map[string]interface{}, propertyName string) (result *CreateApiKeyRequest, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'CreateApiKeyRequest'", propertyName)
			return
		}
		result, err = UnmarshalCreateApiKeyRequest(objMap)
	}
	return
}

// UnmarshalCreateApiKeyRequestSliceAsProperty unmarshals a slice of CreateApiKeyRequest instances that are stored as a property
// within the specified map.
func UnmarshalCreateApiKeyRequestSliceAsProperty(m map[string]interface{}, propertyName string) (slice []CreateApiKeyRequest, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'CreateApiKeyRequest'", propertyName)
			return
		}
		slice, err = UnmarshalCreateApiKeyRequestSlice(vSlice)
	}
	return
}

// CreateServiceIdOptions : The CreateServiceID options.
type CreateServiceIdOptions struct {
	// ID of the account the service ID belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// Name of the service ID. The name is not checked for uniqueness. Therefore multiple names with the same value can
	// exist. Access is done via the UUID of the Service Id.
	Name *string `json:"name" validate:"required"`

	// The optional description of the Service Id. The 'description' property is only available if a description was
	// provided during a create of a ServiceId.
	Description *string `json:"description,omitempty"`

	// Optional list of CRNs (string array) which point to the services connected to the service ID.
	UniqueInstanceCrns []string `json:"unique_instance_crns,omitempty"`

	// Input body parameters for the Create API key V1 REST request.
	Apikey *CreateApiKeyRequest `json:"apikey,omitempty"`

	// Indicates if the service ID is locked for further write operations. False by default.
	EntityLock *string `json:"Entity-Lock,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateServiceIdOptions : Instantiate CreateServiceIdOptions
func (*IamIdentityServicesV1) NewCreateServiceIdOptions(accountID string, name string) *CreateServiceIdOptions {
	return &CreateServiceIdOptions{
		AccountID: core.StringPtr(accountID),
		Name: core.StringPtr(name),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateServiceIdOptions) SetAccountID(accountID string) *CreateServiceIdOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateServiceIdOptions) SetName(name string) *CreateServiceIdOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateServiceIdOptions) SetDescription(description string) *CreateServiceIdOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetUniqueInstanceCrns : Allow user to set UniqueInstanceCrns
func (options *CreateServiceIdOptions) SetUniqueInstanceCrns(uniqueInstanceCrns []string) *CreateServiceIdOptions {
	options.UniqueInstanceCrns = uniqueInstanceCrns
	return options
}

// SetApikey : Allow user to set Apikey
func (options *CreateServiceIdOptions) SetApikey(apikey *CreateApiKeyRequest) *CreateServiceIdOptions {
	options.Apikey = apikey
	return options
}

// SetEntityLock : Allow user to set EntityLock
func (options *CreateServiceIdOptions) SetEntityLock(entityLock string) *CreateServiceIdOptions {
	options.EntityLock = core.StringPtr(entityLock)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateServiceIdOptions) SetHeaders(param map[string]string) *CreateServiceIdOptions {
	options.Headers = param
	return options
}

// DeleteApiKeyOptions : The DeleteApiKey options.
type DeleteApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteApiKeyOptions : Instantiate DeleteApiKeyOptions
func (*IamIdentityServicesV1) NewDeleteApiKeyOptions(id string) *DeleteApiKeyOptions {
	return &DeleteApiKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteApiKeyOptions) SetID(id string) *DeleteApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteApiKeyOptions) SetHeaders(param map[string]string) *DeleteApiKeyOptions {
	options.Headers = param
	return options
}

// DeleteServiceIdOptions : The DeleteServiceID options.
type DeleteServiceIdOptions struct {
	// Unique ID of the service ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceIdOptions : Instantiate DeleteServiceIdOptions
func (*IamIdentityServicesV1) NewDeleteServiceIdOptions(id string) *DeleteServiceIdOptions {
	return &DeleteServiceIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteServiceIdOptions) SetID(id string) *DeleteServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceIdOptions) SetHeaders(param map[string]string) *DeleteServiceIdOptions {
	options.Headers = param
	return options
}

// ExceptionResponseContext : Context fill with key properties for problem determination.
type ExceptionResponseContext struct {
	// The request ID of the inbound REST request.
	RequestID *string `json:"requestId,omitempty"`

	// The request type of the inbound REST request.
	RequestType *string `json:"requestType,omitempty"`

	// The user agent of the inbound REST request.
	UserAgent *string `json:"userAgent,omitempty"`

	// The client ip address of the inbound REST request.
	ClientIp *string `json:"clientIp,omitempty"`

	// The URL of that cluster.
	URL *string `json:"url,omitempty"`

	// The instance ID of the server instance processing the request.
	InstanceID *string `json:"instanceId,omitempty"`

	// The thread ID of the server instance processing the request.
	ThreadID *string `json:"threadId,omitempty"`

	// The host of the server instance processing the request.
	Host *string `json:"host,omitempty"`

	// The start time of the request.
	StartTime *string `json:"startTime,omitempty"`

	// The finish time of the request.
	EndTime *string `json:"endTime,omitempty"`

	// The elapsed time in msec.
	ElapsedTime *string `json:"elapsedTime,omitempty"`

	// The language used to present the error message.
	Locale *string `json:"locale,omitempty"`

	// The cluster name.
	ClusterName *string `json:"clusterName,omitempty"`
}


// UnmarshalExceptionResponseContext constructs an instance of ExceptionResponseContext from the specified map.
func UnmarshalExceptionResponseContext(m map[string]interface{}) (result *ExceptionResponseContext, err error) {
	obj := new(ExceptionResponseContext)
	obj.RequestID, err = core.UnmarshalString(m, "requestId")
	if err != nil {
		return
	}
	obj.RequestType, err = core.UnmarshalString(m, "requestType")
	if err != nil {
		return
	}
	obj.UserAgent, err = core.UnmarshalString(m, "userAgent")
	if err != nil {
		return
	}
	obj.ClientIp, err = core.UnmarshalString(m, "clientIp")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "url")
	if err != nil {
		return
	}
	obj.InstanceID, err = core.UnmarshalString(m, "instanceId")
	if err != nil {
		return
	}
	obj.ThreadID, err = core.UnmarshalString(m, "threadId")
	if err != nil {
		return
	}
	obj.Host, err = core.UnmarshalString(m, "host")
	if err != nil {
		return
	}
	obj.StartTime, err = core.UnmarshalString(m, "startTime")
	if err != nil {
		return
	}
	obj.EndTime, err = core.UnmarshalString(m, "endTime")
	if err != nil {
		return
	}
	obj.ElapsedTime, err = core.UnmarshalString(m, "elapsedTime")
	if err != nil {
		return
	}
	obj.Locale, err = core.UnmarshalString(m, "locale")
	if err != nil {
		return
	}
	obj.ClusterName, err = core.UnmarshalString(m, "clusterName")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalExceptionResponseContextSlice unmarshals a slice of ExceptionResponseContext instances from the specified list of maps.
func UnmarshalExceptionResponseContextSlice(s []interface{}) (slice []ExceptionResponseContext, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ExceptionResponseContext'")
			return
		}
		obj, e := UnmarshalExceptionResponseContext(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalExceptionResponseContextAsProperty unmarshals an instance of ExceptionResponseContext that is stored as a property
// within the specified map.
func UnmarshalExceptionResponseContextAsProperty(m map[string]interface{}, propertyName string) (result *ExceptionResponseContext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ExceptionResponseContext'", propertyName)
			return
		}
		result, err = UnmarshalExceptionResponseContext(objMap)
	}
	return
}

// UnmarshalExceptionResponseContextSliceAsProperty unmarshals a slice of ExceptionResponseContext instances that are stored as a property
// within the specified map.
func UnmarshalExceptionResponseContextSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ExceptionResponseContext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ExceptionResponseContext'", propertyName)
			return
		}
		slice, err = UnmarshalExceptionResponseContextSlice(vSlice)
	}
	return
}

// GetApiKeyDetailsOptions : The GetApiKeyDetails options.
type GetApiKeyDetailsOptions struct {
	// API key value.
	IAMApiKey *string `json:"IAM-ApiKey,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApiKeyDetailsOptions : Instantiate GetApiKeyDetailsOptions
func (*IamIdentityServicesV1) NewGetApiKeyDetailsOptions() *GetApiKeyDetailsOptions {
	return &GetApiKeyDetailsOptions{}
}

// SetIAMApiKey : Allow user to set IAMApiKey
func (options *GetApiKeyDetailsOptions) SetIAMApiKey(iAMApiKey string) *GetApiKeyDetailsOptions {
	options.IAMApiKey = core.StringPtr(iAMApiKey)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApiKeyDetailsOptions) SetHeaders(param map[string]string) *GetApiKeyDetailsOptions {
	options.Headers = param
	return options
}

// GetApiKeyOptions : The GetApiKey options.
type GetApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApiKeyOptions : Instantiate GetApiKeyOptions
func (*IamIdentityServicesV1) NewGetApiKeyOptions(id string) *GetApiKeyOptions {
	return &GetApiKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetApiKeyOptions) SetID(id string) *GetApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApiKeyOptions) SetHeaders(param map[string]string) *GetApiKeyOptions {
	options.Headers = param
	return options
}

// GetServiceIdOptions : The GetServiceID options.
type GetServiceIdOptions struct {
	// Unique ID of the service ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceIdOptions : Instantiate GetServiceIdOptions
func (*IamIdentityServicesV1) NewGetServiceIdOptions(id string) *GetServiceIdOptions {
	return &GetServiceIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetServiceIdOptions) SetID(id string) *GetServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceIdOptions) SetHeaders(param map[string]string) *GetServiceIdOptions {
	options.Headers = param
	return options
}

// ListApiKeysOptions : The ListApiKeys options.
type ListApiKeysOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id,omitempty"`

	// IAM ID of the API key(s) to be queried. The IAM ID may be that of a user or a service. For a user IAM ID iam_id must
	// match the Authorization token.
	IamID *string `json:"iam_id,omitempty"`

	// Optional size of a single page. Default is 20 items per page. Valid range is 1 to 100.
	Pagesize *string `json:"pagesize,omitempty"`

	// Optional Prev or Next page token returned from a previous query execution. Default is start with first page.
	Pagetoken *string `json:"pagetoken,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListApiKeysOptions : Instantiate ListApiKeysOptions
func (*IamIdentityServicesV1) NewListApiKeysOptions() *ListApiKeysOptions {
	return &ListApiKeysOptions{}
}

// SetAccountID : Allow user to set AccountID
func (options *ListApiKeysOptions) SetAccountID(accountID string) *ListApiKeysOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *ListApiKeysOptions) SetIamID(iamID string) *ListApiKeysOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetPagesize : Allow user to set Pagesize
func (options *ListApiKeysOptions) SetPagesize(pagesize string) *ListApiKeysOptions {
	options.Pagesize = core.StringPtr(pagesize)
	return options
}

// SetPagetoken : Allow user to set Pagetoken
func (options *ListApiKeysOptions) SetPagetoken(pagetoken string) *ListApiKeysOptions {
	options.Pagetoken = core.StringPtr(pagetoken)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListApiKeysOptions) SetHeaders(param map[string]string) *ListApiKeysOptions {
	options.Headers = param
	return options
}

// ListApiKeysResponse : Response body format for the List API keys V1 REST request.
type ListApiKeysResponse struct {
	// Context fill with key properties for problem determination.
	Context *ExceptionResponseContext `json:"context,omitempty"`

	// The offset of the current page.
	Offset *int64 `json:"offset,omitempty"`

	// Optional size of a single page. Default is 20 items per page. Valid range is 1 to 100.
	Limit *int64 `json:"limit,omitempty"`

	// Link to the first page.
	First *string `json:"first,omitempty"`

	// Link to the previous available page. If 'previous' property is not part of the response no previous page is
	// available.
	Previous *string `json:"previous,omitempty"`

	// Link to the next available page. If 'next' property is not part of the response no next page is available.
	Next *string `json:"next,omitempty"`

	// List of API keys based on the query paramters and the page size. The apikeys array is always part of the response
	// but might be empty depending on the query parameters values provided.
	Apikeys []ApiKeyDetails `json:"apikeys" validate:"required"`
}


// UnmarshalListApiKeysResponse constructs an instance of ListApiKeysResponse from the specified map.
func UnmarshalListApiKeysResponse(m map[string]interface{}) (result *ListApiKeysResponse, err error) {
	obj := new(ListApiKeysResponse)
	obj.Context, err = UnmarshalExceptionResponseContextAsProperty(m, "context")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.First, err = core.UnmarshalString(m, "first")
	if err != nil {
		return
	}
	obj.Previous, err = core.UnmarshalString(m, "previous")
	if err != nil {
		return
	}
	obj.Next, err = core.UnmarshalString(m, "next")
	if err != nil {
		return
	}
	obj.Apikeys, err = UnmarshalApiKeyDetailsSliceAsProperty(m, "apikeys")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalListApiKeysResponseSlice unmarshals a slice of ListApiKeysResponse instances from the specified list of maps.
func UnmarshalListApiKeysResponseSlice(s []interface{}) (slice []ListApiKeysResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ListApiKeysResponse'")
			return
		}
		obj, e := UnmarshalListApiKeysResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalListApiKeysResponseAsProperty unmarshals an instance of ListApiKeysResponse that is stored as a property
// within the specified map.
func UnmarshalListApiKeysResponseAsProperty(m map[string]interface{}, propertyName string) (result *ListApiKeysResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ListApiKeysResponse'", propertyName)
			return
		}
		result, err = UnmarshalListApiKeysResponse(objMap)
	}
	return
}

// UnmarshalListApiKeysResponseSliceAsProperty unmarshals a slice of ListApiKeysResponse instances that are stored as a property
// within the specified map.
func UnmarshalListApiKeysResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ListApiKeysResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ListApiKeysResponse'", propertyName)
			return
		}
		slice, err = UnmarshalListApiKeysResponseSlice(vSlice)
	}
	return
}

// ListServiceIdsOptions : The ListServiceIds options.
type ListServiceIdsOptions struct {
	// Account ID of the service ID(s) to query. This parameter is required (unless using a pagetoken).
	AccountID *string `json:"account_id,omitempty"`

	// Name of the service ID(s) to query. Optional.20 items per page.
	Name *string `json:"name,omitempty"`

	// Optional size of a single page. Default is 20 items per page. Valid range is 1 to 100.
	Pagesize *string `json:"pagesize,omitempty"`

	// Optional Prev or Next page token returned from a previous query execution. Default is start with first page.
	Pagetoken *string `json:"pagetoken,omitempty"`

	// Optional sort property, valid values are name, description, createdAt and modifiedAt. If specified, the items are
	// sorted by the value of this property.
	Sort *string `json:"sort,omitempty"`

	// Optional sort order, valid values are asc and desc. Default: asc.
	Order *string `json:"order,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListServiceIdsOptions : Instantiate ListServiceIdsOptions
func (*IamIdentityServicesV1) NewListServiceIdsOptions() *ListServiceIdsOptions {
	return &ListServiceIdsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (options *ListServiceIdsOptions) SetAccountID(accountID string) *ListServiceIdsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *ListServiceIdsOptions) SetName(name string) *ListServiceIdsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetPagesize : Allow user to set Pagesize
func (options *ListServiceIdsOptions) SetPagesize(pagesize string) *ListServiceIdsOptions {
	options.Pagesize = core.StringPtr(pagesize)
	return options
}

// SetPagetoken : Allow user to set Pagetoken
func (options *ListServiceIdsOptions) SetPagetoken(pagetoken string) *ListServiceIdsOptions {
	options.Pagetoken = core.StringPtr(pagetoken)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListServiceIdsOptions) SetSort(sort string) *ListServiceIdsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetOrder : Allow user to set Order
func (options *ListServiceIdsOptions) SetOrder(order string) *ListServiceIdsOptions {
	options.Order = core.StringPtr(order)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListServiceIdsOptions) SetHeaders(param map[string]string) *ListServiceIdsOptions {
	options.Headers = param
	return options
}

// LockApiKeyOptions : The LockApiKey options.
type LockApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewLockApiKeyOptions : Instantiate LockApiKeyOptions
func (*IamIdentityServicesV1) NewLockApiKeyOptions(id string) *LockApiKeyOptions {
	return &LockApiKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *LockApiKeyOptions) SetID(id string) *LockApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *LockApiKeyOptions) SetHeaders(param map[string]string) *LockApiKeyOptions {
	options.Headers = param
	return options
}

// LockServiceIdOptions : The LockServiceID options.
type LockServiceIdOptions struct {
	// Unique ID of the service ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewLockServiceIdOptions : Instantiate LockServiceIdOptions
func (*IamIdentityServicesV1) NewLockServiceIdOptions(id string) *LockServiceIdOptions {
	return &LockServiceIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *LockServiceIdOptions) SetID(id string) *LockServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *LockServiceIdOptions) SetHeaders(param map[string]string) *LockServiceIdOptions {
	options.Headers = param
	return options
}

// ResponseContext : Context with key properties for problem determination.
type ResponseContext struct {
	// The transaction ID of the inbound REST request.
	TransactionID *string `json:"transaction_id,omitempty"`

	// The operation of the inbound REST request.
	Operation *string `json:"operation,omitempty"`

	// The user agent of the inbound REST request.
	UserAgent *string `json:"user_agent,omitempty"`

	// The client ip address of the inbound REST request.
	ClientIp *string `json:"client_ip,omitempty"`

	// The URL of that cluster.
	URL *string `json:"url,omitempty"`

	// The instance ID of the server instance processing the request.
	InstanceID *string `json:"instance_id,omitempty"`

	// The thread ID of the server instance processing the request.
	ThreadID *string `json:"thread_id,omitempty"`

	// The host of the server instance processing the request.
	Host *string `json:"host,omitempty"`

	// The start time of the request.
	StartTime *string `json:"start_time,omitempty"`

	// The finish time of the request.
	EndTime *string `json:"end_time,omitempty"`

	// The elapsed time in msec.
	ElapsedTime *string `json:"elapsed_time,omitempty"`

	// The cluster name.
	ClusterName *string `json:"cluster_name,omitempty"`
}


// UnmarshalResponseContext constructs an instance of ResponseContext from the specified map.
func UnmarshalResponseContext(m map[string]interface{}) (result *ResponseContext, err error) {
	obj := new(ResponseContext)
	obj.TransactionID, err = core.UnmarshalString(m, "transaction_id")
	if err != nil {
		return
	}
	obj.Operation, err = core.UnmarshalString(m, "operation")
	if err != nil {
		return
	}
	obj.UserAgent, err = core.UnmarshalString(m, "user_agent")
	if err != nil {
		return
	}
	obj.ClientIp, err = core.UnmarshalString(m, "client_ip")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "url")
	if err != nil {
		return
	}
	obj.InstanceID, err = core.UnmarshalString(m, "instance_id")
	if err != nil {
		return
	}
	obj.ThreadID, err = core.UnmarshalString(m, "thread_id")
	if err != nil {
		return
	}
	obj.Host, err = core.UnmarshalString(m, "host")
	if err != nil {
		return
	}
	obj.StartTime, err = core.UnmarshalString(m, "start_time")
	if err != nil {
		return
	}
	obj.EndTime, err = core.UnmarshalString(m, "end_time")
	if err != nil {
		return
	}
	obj.ElapsedTime, err = core.UnmarshalString(m, "elapsed_time")
	if err != nil {
		return
	}
	obj.ClusterName, err = core.UnmarshalString(m, "cluster_name")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResponseContextSlice unmarshals a slice of ResponseContext instances from the specified list of maps.
func UnmarshalResponseContextSlice(s []interface{}) (slice []ResponseContext, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResponseContext'")
			return
		}
		obj, e := UnmarshalResponseContext(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResponseContextAsProperty unmarshals an instance of ResponseContext that is stored as a property
// within the specified map.
func UnmarshalResponseContextAsProperty(m map[string]interface{}, propertyName string) (result *ResponseContext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResponseContext'", propertyName)
			return
		}
		result, err = UnmarshalResponseContext(objMap)
	}
	return
}

// UnmarshalResponseContextSliceAsProperty unmarshals a slice of ResponseContext instances that are stored as a property
// within the specified map.
func UnmarshalResponseContextSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResponseContext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResponseContext'", propertyName)
			return
		}
		slice, err = UnmarshalResponseContextSlice(vSlice)
	}
	return
}

// ServiceIdDetails : Response body format for service ID V1 REST requests.
type ServiceIdDetails struct {
	// Context with key properties for problem determination.
	Context *ResponseContext `json:"context,omitempty"`

	// Unique identifier of this Service ID.
	ID *string `json:"id" validate:"required"`

	// Cloud wide identifier for identities of this service ID.
	IamID *string `json:"iam_id" validate:"required"`

	// Version of the service ID details object. You need to specify this value when updating the service ID to avoid stale
	// updates.
	EntityTag *string `json:"entity_tag,omitempty"`

	// Cloud Resource Name of the item. Example Cloud Resource Name:
	// 'crn:v1:bluemix:public:iam-identity:us-south:a/myaccount::serviceid:1234-5678-9012'.
	Crn *string `json:"crn" validate:"required"`

	// The service ID cannot be changed if set to true.
	Locked *bool `json:"locked" validate:"required"`

	// If set contains a date time string of the creation date in ISO format.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// If set contains a date time string of the last modification date in ISO format.
	ModifiedAt *strfmt.DateTime `json:"modified_at,omitempty"`

	// ID of the account the service ID belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// Name of the service ID. The name is not checked for uniqueness. Therefore multiple names with the same value can
	// exist. Access is done via the UUID of the service ID.
	Name *string `json:"name" validate:"required"`

	// The optional description of the service ID. The 'description' property is only available if a description was
	// provided during a create of a service ID.
	Description *string `json:"description,omitempty"`

	// Optional list of CRNs (string array) which point to the services connected to the service ID.
	UniqueInstanceCrns []string `json:"unique_instance_crns,omitempty"`

	// Response body format for API key V1 REST requests.
	Apikey *ApiKeyDetails `json:"apikey" validate:"required"`
}


// UnmarshalServiceIdDetails constructs an instance of ServiceIdDetails from the specified map.
func UnmarshalServiceIdDetails(m map[string]interface{}) (result *ServiceIdDetails, err error) {
	obj := new(ServiceIdDetails)
	obj.Context, err = UnmarshalResponseContextAsProperty(m, "context")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.EntityTag, err = core.UnmarshalString(m, "entity_tag")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.Locked, err = core.UnmarshalBool(m, "locked")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalDateTime(m, "created_at")
	if err != nil {
		return
	}
	obj.ModifiedAt, err = core.UnmarshalDateTime(m, "modified_at")
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
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.UniqueInstanceCrns, err = core.UnmarshalStringSlice(m, "unique_instance_crns")
	if err != nil {
		return
	}
	obj.Apikey, err = UnmarshalApiKeyDetailsAsProperty(m, "apikey")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalServiceIdDetailsSlice unmarshals a slice of ServiceIdDetails instances from the specified list of maps.
func UnmarshalServiceIdDetailsSlice(s []interface{}) (slice []ServiceIdDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ServiceIdDetails'")
			return
		}
		obj, e := UnmarshalServiceIdDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalServiceIdDetailsAsProperty unmarshals an instance of ServiceIdDetails that is stored as a property
// within the specified map.
func UnmarshalServiceIdDetailsAsProperty(m map[string]interface{}, propertyName string) (result *ServiceIdDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ServiceIdDetails'", propertyName)
			return
		}
		result, err = UnmarshalServiceIdDetails(objMap)
	}
	return
}

// UnmarshalServiceIdDetailsSliceAsProperty unmarshals a slice of ServiceIdDetails instances that are stored as a property
// within the specified map.
func UnmarshalServiceIdDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ServiceIdDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ServiceIdDetails'", propertyName)
			return
		}
		slice, err = UnmarshalServiceIdDetailsSlice(vSlice)
	}
	return
}

// ServiceIdsList : Response body format for the list service ID V1 REST request.
type ServiceIdsList struct {
	// Context fill with key properties for problem determination.
	Context *ExceptionResponseContext `json:"context,omitempty"`

	// The offset of the current page.
	Offset *int64 `json:"offset,omitempty"`

	// Optional size of a single page. Default is 20 items per page. Valid range is 1 to 100.
	Limit *int64 `json:"limit,omitempty"`

	// Link to the first page.
	First *string `json:"first,omitempty"`

	// Link to the previous available page. If 'previous' property is not part of the response no previous page is
	// available.
	Previous *string `json:"previous,omitempty"`

	// Link to the next available page. If 'next' property is not part of the response no next page is available.
	Next *string `json:"next,omitempty"`

	// List of service IDs based on the query paramters and the page size. The service IDs array is always part of the
	// response but might be empty depending on the query parameter values provided.
	Serviceids []ServiceIdDetails `json:"serviceids" validate:"required"`
}


// UnmarshalServiceIdsList constructs an instance of ServiceIdsList from the specified map.
func UnmarshalServiceIdsList(m map[string]interface{}) (result *ServiceIdsList, err error) {
	obj := new(ServiceIdsList)
	obj.Context, err = UnmarshalExceptionResponseContextAsProperty(m, "context")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.First, err = core.UnmarshalString(m, "first")
	if err != nil {
		return
	}
	obj.Previous, err = core.UnmarshalString(m, "previous")
	if err != nil {
		return
	}
	obj.Next, err = core.UnmarshalString(m, "next")
	if err != nil {
		return
	}
	obj.Serviceids, err = UnmarshalServiceIdDetailsSliceAsProperty(m, "serviceids")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalServiceIdsListSlice unmarshals a slice of ServiceIdsList instances from the specified list of maps.
func UnmarshalServiceIdsListSlice(s []interface{}) (slice []ServiceIdsList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ServiceIdsList'")
			return
		}
		obj, e := UnmarshalServiceIdsList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalServiceIdsListAsProperty unmarshals an instance of ServiceIdsList that is stored as a property
// within the specified map.
func UnmarshalServiceIdsListAsProperty(m map[string]interface{}, propertyName string) (result *ServiceIdsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ServiceIdsList'", propertyName)
			return
		}
		result, err = UnmarshalServiceIdsList(objMap)
	}
	return
}

// UnmarshalServiceIdsListSliceAsProperty unmarshals a slice of ServiceIdsList instances that are stored as a property
// within the specified map.
func UnmarshalServiceIdsListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ServiceIdsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ServiceIdsList'", propertyName)
			return
		}
		slice, err = UnmarshalServiceIdsListSlice(vSlice)
	}
	return
}

// UnlockApiKeyOptions : The UnlockApiKey options.
type UnlockApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnlockApiKeyOptions : Instantiate UnlockApiKeyOptions
func (*IamIdentityServicesV1) NewUnlockApiKeyOptions(id string) *UnlockApiKeyOptions {
	return &UnlockApiKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UnlockApiKeyOptions) SetID(id string) *UnlockApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UnlockApiKeyOptions) SetHeaders(param map[string]string) *UnlockApiKeyOptions {
	options.Headers = param
	return options
}

// UnlockServiceIdOptions : The UnlockServiceID options.
type UnlockServiceIdOptions struct {
	// Unique ID of the service ID.
	ID *string `json:"id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnlockServiceIdOptions : Instantiate UnlockServiceIdOptions
func (*IamIdentityServicesV1) NewUnlockServiceIdOptions(id string) *UnlockServiceIdOptions {
	return &UnlockServiceIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UnlockServiceIdOptions) SetID(id string) *UnlockServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UnlockServiceIdOptions) SetHeaders(param map[string]string) *UnlockServiceIdOptions {
	options.Headers = param
	return options
}

// UpdateApiKeyOptions : The UpdateApiKey options.
type UpdateApiKeyOptions struct {
	// Unique ID of the API key to be updated.
	ID *string `json:"id" validate:"required"`

	// Version of the API key to be updated. Specify the version that you retrieved when reading the API key. This value
	// helps identifying parallel usage of this API. Pass * to indicate to update any version available. This might result
	// in stale updates.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The name of the API key to update. If specified in the request the parameter must not be empty. The name is not
	// checked for uniqueness. Failure to this will result in an Error condition.
	Name *string `json:"name,omitempty"`

	// The description of the API key to update. If specified an empty description will clear the description of the API
	// key. If a non empty value is provided the API key will be updated.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateApiKeyOptions : Instantiate UpdateApiKeyOptions
func (*IamIdentityServicesV1) NewUpdateApiKeyOptions(id string, ifMatch string) *UpdateApiKeyOptions {
	return &UpdateApiKeyOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (options *UpdateApiKeyOptions) SetID(id string) *UpdateApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateApiKeyOptions) SetIfMatch(ifMatch string) *UpdateApiKeyOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateApiKeyOptions) SetName(name string) *UpdateApiKeyOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateApiKeyOptions) SetDescription(description string) *UpdateApiKeyOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateApiKeyOptions) SetHeaders(param map[string]string) *UpdateApiKeyOptions {
	options.Headers = param
	return options
}

// UpdateServiceIdOptions : The UpdateServiceID options.
type UpdateServiceIdOptions struct {
	// Unique ID of the service ID to be updated.
	ID *string `json:"id" validate:"required"`

	// Version of the service ID to be updated. Specify the version that you retrieved as entity_tag (ETag header) when
	// reading the service ID. This value helps identifying parallel usage of this API. Pass * to indicate to update any
	// version available. This might result in stale updates.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The name of the service ID to update. If specified in the request the parameter must not be empty. The name is not
	// checked for uniqueness. Failure to this will result in an Error condition.
	Name *string `json:"name,omitempty"`

	// The description of the service ID to update. If specified an empty description will clear the description of the
	// service ID. If an non empty value is provided the service ID will be updated.
	Description *string `json:"description,omitempty"`

	// List of CRNs which point to the services connected to this service ID. If specified an empty list will clear all
	// existing unique instance crns of the service ID.
	UniqueInstanceCrns []string `json:"unique_instance_crns,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateServiceIdOptions : Instantiate UpdateServiceIdOptions
func (*IamIdentityServicesV1) NewUpdateServiceIdOptions(id string, ifMatch string) *UpdateServiceIdOptions {
	return &UpdateServiceIdOptions{
		ID: core.StringPtr(id),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetID : Allow user to set ID
func (options *UpdateServiceIdOptions) SetID(id string) *UpdateServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateServiceIdOptions) SetIfMatch(ifMatch string) *UpdateServiceIdOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateServiceIdOptions) SetName(name string) *UpdateServiceIdOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateServiceIdOptions) SetDescription(description string) *UpdateServiceIdOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetUniqueInstanceCrns : Allow user to set UniqueInstanceCrns
func (options *UpdateServiceIdOptions) SetUniqueInstanceCrns(uniqueInstanceCrns []string) *UpdateServiceIdOptions {
	options.UniqueInstanceCrns = uniqueInstanceCrns
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateServiceIdOptions) SetHeaders(param map[string]string) *UpdateServiceIdOptions {
	options.Headers = param
	return options
}
