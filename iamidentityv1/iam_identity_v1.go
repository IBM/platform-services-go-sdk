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

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-c934890e-20200918-141729
 */
 

// Package iamidentityv1 : Operations and models for the IamIdentityV1 service
package iamidentityv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"reflect"
)

// IamIdentityV1 : The IAM Identity Service API allows for the management of Identities (Service IDs, ApiKeys).
//
// Version: 1.0.0
type IamIdentityV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.test.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "iam_identity"

// IamIdentityV1Options : Service options
type IamIdentityV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIamIdentityV1UsingExternalConfig : constructs an instance of IamIdentityV1 with passed in options and external configuration.
func NewIamIdentityV1UsingExternalConfig(options *IamIdentityV1Options) (iamIdentity *IamIdentityV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	iamIdentity, err = NewIamIdentityV1(options)
	if err != nil {
		return
	}

	err = iamIdentity.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = iamIdentity.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIamIdentityV1 : constructs an instance of IamIdentityV1 with passed in options.
func NewIamIdentityV1(options *IamIdentityV1Options) (service *IamIdentityV1, err error) {
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

	service = &IamIdentityV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (iamIdentity *IamIdentityV1) SetServiceURL(url string) error {
	return iamIdentity.Service.SetServiceURL(url)
}

// ListApiKeys : Get API keys for a given service or user IAM ID and account ID
// Returns the list of API key details for a given service or user IAM ID and account ID. Users can manage user API keys
// for themself, or service ID API keys for  service IDs that are bound to an entity they have access to. In case of
// service IDs and their API keys, a user must be either an account owner,  a IBM Cloud org manager or IBM Cloud space
// developer in order to manage  service IDs of the entity.
func (iamIdentity *IamIdentityV1) ListApiKeys(listApiKeysOptions *ListApiKeysOptions) (result *ApiKeyList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listApiKeysOptions, "listApiKeysOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listApiKeysOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "ListApiKeys")
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
	if listApiKeysOptions.Scope != nil {
		builder.AddQuery("scope", fmt.Sprint(*listApiKeysOptions.Scope))
	}
	if listApiKeysOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listApiKeysOptions.Type))
	}
	if listApiKeysOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listApiKeysOptions.Sort))
	}
	if listApiKeysOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listApiKeysOptions.Order))
	}
	if listApiKeysOptions.IncludeHistory != nil {
		builder.AddQuery("include_history", fmt.Sprint(*listApiKeysOptions.IncludeHistory))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiKeyList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateApiKey : Create an API key
// Creates an API key for a UserID or service ID. Users can manage user API keys for themself, or service ID API keys
// for  service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) CreateApiKey(createApiKeyOptions *CreateApiKeyOptions) (result *ApiKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createApiKeyOptions, "createApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createApiKeyOptions, "createApiKeyOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "CreateApiKey")
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
	if createApiKeyOptions.StoreValue != nil {
		body["store_value"] = createApiKeyOptions.StoreValue
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
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiKey)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetApiKeysDetails : Get details of an API key by its value
// Returns the details of an API key by its value. Users can manage user API keys for themself, or service ID API keys
// for service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) GetApiKeysDetails(getApiKeysDetailsOptions *GetApiKeysDetailsOptions) (result *ApiKey, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getApiKeysDetailsOptions, "getApiKeysDetailsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/details`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApiKeysDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "GetApiKeysDetails")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getApiKeysDetailsOptions.IAMApiKey != nil {
		builder.AddHeader("IAM-ApiKey", fmt.Sprint(*getApiKeysDetailsOptions.IAMApiKey))
	}

	if getApiKeysDetailsOptions.IncludeHistory != nil {
		builder.AddQuery("include_history", fmt.Sprint(*getApiKeysDetailsOptions.IncludeHistory))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiKey)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetApiKey : Get details of an API key
// Returns the details of an API key. Users can manage user API keys for themself, or service ID API keys for  service
// IDs that are bound to an entity they have access to. In case of  service IDs and their API keys, a user must be
// either an account owner,  a IBM Cloud org manager or IBM Cloud space developer in order to manage  service IDs of the
// entity.
func (iamIdentity *IamIdentityV1) GetApiKey(getApiKeyOptions *GetApiKeyOptions) (result *ApiKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApiKeyOptions, "getApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApiKeyOptions, "getApiKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getApiKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "GetApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getApiKeyOptions.IncludeHistory != nil {
		builder.AddQuery("include_history", fmt.Sprint(*getApiKeyOptions.IncludeHistory))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiKey)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateApiKey : Updates an API key
// Updates properties of an API key. This does NOT affect existing access tokens. Their token content will stay
// unchanged until the access token is refreshed. To update an API key, pass the property to be modified. To delete one
// property's value, pass the property with an empty value "".Users can manage user API keys for themself, or service ID
// API keys for service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) UpdateApiKey(updateApiKeyOptions *UpdateApiKeyOptions) (result *ApiKey, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateApiKeyOptions, "updateApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateApiKeyOptions, "updateApiKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateApiKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "UpdateApiKey")
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

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApiKey)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteApiKey : Deletes an API key
// Deletes an API key. Existing tokens will remain valid until expired. Refresh tokens  will not work any more for this
// API key. Users can manage user API keys for themself, or service ID API  keys for service IDs that are bound to an
// entity they have access  to.
func (iamIdentity *IamIdentityV1) DeleteApiKey(deleteApiKeyOptions *DeleteApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteApiKeyOptions, "deleteApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteApiKeyOptions, "deleteApiKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteApiKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "DeleteApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentity.Service.Request(request, nil)

	return
}

// LockApiKey : Lock the API key
// Locks an API key by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that are
// bound to an entity they have access to. In case of service IDs and their API keys, a user must be either an account
// owner, a IBM Cloud org manager or IBM Cloud space developer in order to manage service IDs of the entity.
func (iamIdentity *IamIdentityV1) LockApiKey(lockApiKeyOptions *LockApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(lockApiKeyOptions, "lockApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(lockApiKeyOptions, "lockApiKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *lockApiKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/{id}/lock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range lockApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "LockApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentity.Service.Request(request, nil)

	return
}

// UnlockApiKey : Unlock the API key
// Unlocks an API key by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to. In case of service IDs and their API keys, a user must be either an
// account owner, a IBM Cloud org manager or IBM Cloud space developer in order to manage service IDs of the entity.
func (iamIdentity *IamIdentityV1) UnlockApiKey(unlockApiKeyOptions *UnlockApiKeyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unlockApiKeyOptions, "unlockApiKeyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unlockApiKeyOptions, "unlockApiKeyOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *unlockApiKeyOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/apikeys/{id}/lock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range unlockApiKeyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "UnlockApiKey")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentity.Service.Request(request, nil)

	return
}

// ListServiceIds : List service IDs
// Returns a list of service IDs. Users can manage user API keys for themself, or service ID API keys for service IDs
// that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) ListServiceIds(listServiceIdsOptions *ListServiceIdsOptions) (result *ServiceIdList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listServiceIdsOptions, "listServiceIdsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listServiceIdsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "ListServiceIds")
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
	if listServiceIdsOptions.IncludeHistory != nil {
		builder.AddQuery("include_history", fmt.Sprint(*listServiceIdsOptions.IncludeHistory))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceIdList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateServiceID : Create a service ID
// Creates a service ID for an IBM Cloud account. Users can manage user API keys for themself, or service ID API keys
// for service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) CreateServiceID(createServiceIdOptions *CreateServiceIdOptions) (result *ServiceID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createServiceIdOptions, "createServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createServiceIdOptions, "createServiceIdOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "CreateServiceID")
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

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceID)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetServiceID : Get details of a service ID
// Returns the details of a service ID. Users can manage user API keys for themself, or service ID API keys for service
// IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) GetServiceID(getServiceIdOptions *GetServiceIdOptions) (result *ServiceID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getServiceIdOptions, "getServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getServiceIdOptions, "getServiceIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getServiceIdOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "GetServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getServiceIdOptions.IncludeHistory != nil {
		builder.AddQuery("include_history", fmt.Sprint(*getServiceIdOptions.IncludeHistory))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceID)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateServiceID : Update service ID
// Updates properties of a service ID. This does NOT affect existing access tokens. Their token content will stay
// unchanged until the access token is refreshed. To update a service ID, pass the property to be modified. To delete
// one property's value, pass the property with an empty value "".Users can manage user API keys for themself, or
// service ID API keys for service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) UpdateServiceID(updateServiceIdOptions *UpdateServiceIdOptions) (result *ServiceID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateServiceIdOptions, "updateServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateServiceIdOptions, "updateServiceIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *updateServiceIdOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "UpdateServiceID")
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

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceID)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteServiceID : Deletes a service ID and associated API keys
// Deletes a service ID and all API keys associated to it. Before deleting the service ID, all associated API keys are
// deleted. In case a Delete Conflict (status code 409) a retry of the request may help as the service ID is only
// deleted if the associated API keys were successfully deleted before. Users can manage user API keys for themself, or
// service ID API keys for service IDs that are bound to an entity they have access to.
func (iamIdentity *IamIdentityV1) DeleteServiceID(deleteServiceIdOptions *DeleteServiceIdOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceIdOptions, "deleteServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceIdOptions, "deleteServiceIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteServiceIdOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "DeleteServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamIdentity.Service.Request(request, nil)

	return
}

// LockServiceID : Lock the service ID
// Locks a service ID by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to. In case of service IDs and their API keys, a user must be either an
// account owner, a IBM Cloud org manager or IBM Cloud space developer in order to manage service IDs of the entity.
func (iamIdentity *IamIdentityV1) LockServiceID(lockServiceIdOptions *LockServiceIdOptions) (result *ServiceID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(lockServiceIdOptions, "lockServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(lockServiceIdOptions, "lockServiceIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *lockServiceIdOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/{id}/lock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range lockServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "LockServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceID)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UnlockServiceID : Unlock the service ID
// Unlocks a service ID by ID. Users can manage user API keys for themself, or service ID API keys for service IDs that
// are bound to an entity they have access to. In case of service IDs and their API keys, a user must be either an
// account owner, a IBM Cloud org manager or IBM Cloud space developer in order to manage service IDs of the entity.
func (iamIdentity *IamIdentityV1) UnlockServiceID(unlockServiceIdOptions *UnlockServiceIdOptions) (result *ServiceID, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(unlockServiceIdOptions, "unlockServiceIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(unlockServiceIdOptions, "unlockServiceIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *unlockServiceIdOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(iamIdentity.Service.Options.URL, `/v1/serviceids/{id}/lock`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range unlockServiceIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_identity", "V1", "UnlockServiceID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamIdentity.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceID)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ApiKey : Response body format for API key V1 REST requests.
type ApiKey struct {
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
	CreatedAt *string `json:"created_at,omitempty"`

	// IAM ID of the user or service which created the API key.
	CreatedBy *string `json:"created_by" validate:"required"`

	// If set contains a date time string of the last modification date in ISO format.
	ModifiedAt *string `json:"modified_at,omitempty"`

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

	// The API key value. This property only contains the API key value for the following cases: create an API key, update
	// a service ID API key that stores the API key value as retrievable, or get a service ID API key that stores the API
	// key value as retrievable. All other operations don't return the API key value, for example all user API key related
	// operations, except for create, don't contain the API key value.
	Apikey *string `json:"apikey" validate:"required"`

	// History of the API key.
	History []EnityHistoryRecord `json:"history,omitempty"`
}


// UnmarshalApiKey unmarshals an instance of ApiKey from the specified map of raw messages.
func UnmarshalApiKey(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiKey)
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalResponseContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_tag", &obj.EntityTag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
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
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "apikey", &obj.Apikey)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "history", &obj.History, UnmarshalEnityHistoryRecord)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApiKeyList : Response body format for the List API keys V1 REST request.
type ApiKeyList struct {
	// Context with key properties for problem determination.
	Context *ResponseContext `json:"context,omitempty"`

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
	Apikeys []ApiKey `json:"apikeys" validate:"required"`
}


// UnmarshalApiKeyList unmarshals an instance of ApiKeyList from the specified map of raw messages.
func UnmarshalApiKeyList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApiKeyList)
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalResponseContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "previous", &obj.Previous)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "apikeys", &obj.Apikeys, UnmarshalApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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

	// Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API
	// key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing
	// of API keys for users.
	StoreValue *bool `json:"store_value,omitempty"`

	// Indicates if the API key is locked for further write operations. False by default.
	EntityLock *string `json:"Entity-Lock,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateApiKeyOptions : Instantiate CreateApiKeyOptions
func (*IamIdentityV1) NewCreateApiKeyOptions(name string, iamID string) *CreateApiKeyOptions {
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

// SetStoreValue : Allow user to set StoreValue
func (options *CreateApiKeyOptions) SetStoreValue(storeValue bool) *CreateApiKeyOptions {
	options.StoreValue = core.BoolPtr(storeValue)
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

	// Send true or false to set whether the API key value is retrievable in the future by using the Get details of an API
	// key request. If you create an API key for a user, you must specify `false` or omit the value. We don't allow storing
	// of API keys for users.
	StoreValue *bool `json:"store_value,omitempty"`
}


// NewCreateApiKeyRequest : Instantiate CreateApiKeyRequest (Generic Model Constructor)
func (*IamIdentityV1) NewCreateApiKeyRequest(name string, iamID string) (model *CreateApiKeyRequest, err error) {
	model = &CreateApiKeyRequest{
		Name: core.StringPtr(name),
		IamID: core.StringPtr(iamID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateApiKeyRequest unmarshals an instance of CreateApiKeyRequest from the specified map of raw messages.
func UnmarshalCreateApiKeyRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateApiKeyRequest)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "apikey", &obj.Apikey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "store_value", &obj.StoreValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateServiceIdOptions : The CreateServiceID options.
type CreateServiceIdOptions struct {
	// ID of the account the service ID belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// Name of the Service Id. The name is not checked for uniqueness. Therefore multiple names with the same value can
	// exist. Access is done via the UUID of the Service Id.
	Name *string `json:"name" validate:"required"`

	// The optional description of the Service Id. The 'description' property is only available if a description was
	// provided during a create of a Service Id.
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
func (*IamIdentityV1) NewCreateServiceIdOptions(accountID string, name string) *CreateServiceIdOptions {
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
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteApiKeyOptions : Instantiate DeleteApiKeyOptions
func (*IamIdentityV1) NewDeleteApiKeyOptions(id string) *DeleteApiKeyOptions {
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
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceIdOptions : Instantiate DeleteServiceIdOptions
func (*IamIdentityV1) NewDeleteServiceIdOptions(id string) *DeleteServiceIdOptions {
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

// EnityHistoryRecord : Response body format for an entity history record.
type EnityHistoryRecord struct {
	// Timestamp when the action was triggered.
	Timestamp *string `json:"timestamp" validate:"required"`

	// IAM ID of the identity which triggered the action.
	IamID *string `json:"iam_id" validate:"required"`

	// Account of the identity which triggered the action.
	IamIdAccount *string `json:"iam_id_account" validate:"required"`

	// Action of the history entry.
	Action *string `json:"action" validate:"required"`

	// Params of the history entry.
	Params []string `json:"params" validate:"required"`

	// Message which summarizes the executed action.
	Message *string `json:"message" validate:"required"`
}


// UnmarshalEnityHistoryRecord unmarshals an instance of EnityHistoryRecord from the specified map of raw messages.
func UnmarshalEnityHistoryRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnityHistoryRecord)
	err = core.UnmarshalPrimitive(m, "timestamp", &obj.Timestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_id_account", &obj.IamIdAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "params", &obj.Params)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetApiKeyOptions : The GetApiKey options.
type GetApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required,ne="`

	// Defines if the entity history is included in the response.
	IncludeHistory *bool `json:"include_history,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApiKeyOptions : Instantiate GetApiKeyOptions
func (*IamIdentityV1) NewGetApiKeyOptions(id string) *GetApiKeyOptions {
	return &GetApiKeyOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetApiKeyOptions) SetID(id string) *GetApiKeyOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIncludeHistory : Allow user to set IncludeHistory
func (options *GetApiKeyOptions) SetIncludeHistory(includeHistory bool) *GetApiKeyOptions {
	options.IncludeHistory = core.BoolPtr(includeHistory)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApiKeyOptions) SetHeaders(param map[string]string) *GetApiKeyOptions {
	options.Headers = param
	return options
}

// GetApiKeysDetailsOptions : The GetApiKeysDetails options.
type GetApiKeysDetailsOptions struct {
	// API key value.
	IAMApiKey *string `json:"IAM-ApiKey,omitempty"`

	// Defines if the entity history is included in the response.
	IncludeHistory *bool `json:"include_history,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApiKeysDetailsOptions : Instantiate GetApiKeysDetailsOptions
func (*IamIdentityV1) NewGetApiKeysDetailsOptions() *GetApiKeysDetailsOptions {
	return &GetApiKeysDetailsOptions{}
}

// SetIAMApiKey : Allow user to set IAMApiKey
func (options *GetApiKeysDetailsOptions) SetIAMApiKey(iAMApiKey string) *GetApiKeysDetailsOptions {
	options.IAMApiKey = core.StringPtr(iAMApiKey)
	return options
}

// SetIncludeHistory : Allow user to set IncludeHistory
func (options *GetApiKeysDetailsOptions) SetIncludeHistory(includeHistory bool) *GetApiKeysDetailsOptions {
	options.IncludeHistory = core.BoolPtr(includeHistory)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApiKeysDetailsOptions) SetHeaders(param map[string]string) *GetApiKeysDetailsOptions {
	options.Headers = param
	return options
}

// GetServiceIdOptions : The GetServiceID options.
type GetServiceIdOptions struct {
	// Unique ID of the service ID.
	ID *string `json:"id" validate:"required,ne="`

	// Defines if the entity history is included in the response.
	IncludeHistory *bool `json:"include_history,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceIdOptions : Instantiate GetServiceIdOptions
func (*IamIdentityV1) NewGetServiceIdOptions(id string) *GetServiceIdOptions {
	return &GetServiceIdOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetServiceIdOptions) SetID(id string) *GetServiceIdOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetIncludeHistory : Allow user to set IncludeHistory
func (options *GetServiceIdOptions) SetIncludeHistory(includeHistory bool) *GetServiceIdOptions {
	options.IncludeHistory = core.BoolPtr(includeHistory)
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
	Pagesize *int64 `json:"pagesize,omitempty"`

	// Optional Prev or Next page token returned from a previous query execution. Default is start with first page.
	Pagetoken *string `json:"pagetoken,omitempty"`

	// Optional parameter to define the scope of the queried API Keys. Can be 'entity' (default) or 'account'.
	Scope *string `json:"scope,omitempty"`

	// Optional parameter to filter the type of the queried API Keys. Can be 'user' or 'serviceid'.
	Type *string `json:"type,omitempty"`

	// Optional sort property, valid values are name, description, created_at and created_by. If specified, the items are
	// sorted by the value of this property.
	Sort *string `json:"sort,omitempty"`

	// Optional sort order, valid values are asc and desc. Default: asc.
	Order *string `json:"order,omitempty"`

	// Defines if the entity history is included in the response.
	IncludeHistory *bool `json:"include_history,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListApiKeysOptions.Scope property.
// Optional parameter to define the scope of the queried API Keys. Can be 'entity' (default) or 'account'.
const (
	ListApiKeysOptions_Scope_Account = "account"
	ListApiKeysOptions_Scope_Entity = "entity"
)

// Constants associated with the ListApiKeysOptions.Type property.
// Optional parameter to filter the type of the queried API Keys. Can be 'user' or 'serviceid'.
const (
	ListApiKeysOptions_Type_Serviceid = "serviceid"
	ListApiKeysOptions_Type_User = "user"
)

// Constants associated with the ListApiKeysOptions.Order property.
// Optional sort order, valid values are asc and desc. Default: asc.
const (
	ListApiKeysOptions_Order_Asc = "asc"
	ListApiKeysOptions_Order_Desc = "desc"
)

// NewListApiKeysOptions : Instantiate ListApiKeysOptions
func (*IamIdentityV1) NewListApiKeysOptions() *ListApiKeysOptions {
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
func (options *ListApiKeysOptions) SetPagesize(pagesize int64) *ListApiKeysOptions {
	options.Pagesize = core.Int64Ptr(pagesize)
	return options
}

// SetPagetoken : Allow user to set Pagetoken
func (options *ListApiKeysOptions) SetPagetoken(pagetoken string) *ListApiKeysOptions {
	options.Pagetoken = core.StringPtr(pagetoken)
	return options
}

// SetScope : Allow user to set Scope
func (options *ListApiKeysOptions) SetScope(scope string) *ListApiKeysOptions {
	options.Scope = core.StringPtr(scope)
	return options
}

// SetType : Allow user to set Type
func (options *ListApiKeysOptions) SetType(typeVar string) *ListApiKeysOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListApiKeysOptions) SetSort(sort string) *ListApiKeysOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetOrder : Allow user to set Order
func (options *ListApiKeysOptions) SetOrder(order string) *ListApiKeysOptions {
	options.Order = core.StringPtr(order)
	return options
}

// SetIncludeHistory : Allow user to set IncludeHistory
func (options *ListApiKeysOptions) SetIncludeHistory(includeHistory bool) *ListApiKeysOptions {
	options.IncludeHistory = core.BoolPtr(includeHistory)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListApiKeysOptions) SetHeaders(param map[string]string) *ListApiKeysOptions {
	options.Headers = param
	return options
}

// ListServiceIdsOptions : The ListServiceIds options.
type ListServiceIdsOptions struct {
	// Account ID of the service ID(s) to query. This parameter is required (unless using a pagetoken).
	AccountID *string `json:"account_id,omitempty"`

	// Name of the service ID(s) to query. Optional.20 items per page. Valid range is 1 to 100.
	Name *string `json:"name,omitempty"`

	// Optional size of a single page. Default is 20 items per page. Valid range is 1 to 100.
	Pagesize *int64 `json:"pagesize,omitempty"`

	// Optional Prev or Next page token returned from a previous query execution. Default is start with first page.
	Pagetoken *string `json:"pagetoken,omitempty"`

	// Optional sort property, valid values are name, description, created_at and modified_at. If specified, the items are
	// sorted by the value of this property.
	Sort *string `json:"sort,omitempty"`

	// Optional sort order, valid values are asc and desc. Default: asc.
	Order *string `json:"order,omitempty"`

	// Defines if the entity history is included in the response.
	IncludeHistory *bool `json:"include_history,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListServiceIdsOptions.Order property.
// Optional sort order, valid values are asc and desc. Default: asc.
const (
	ListServiceIdsOptions_Order_Asc = "asc"
	ListServiceIdsOptions_Order_Desc = "desc"
)

// NewListServiceIdsOptions : Instantiate ListServiceIdsOptions
func (*IamIdentityV1) NewListServiceIdsOptions() *ListServiceIdsOptions {
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
func (options *ListServiceIdsOptions) SetPagesize(pagesize int64) *ListServiceIdsOptions {
	options.Pagesize = core.Int64Ptr(pagesize)
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

// SetIncludeHistory : Allow user to set IncludeHistory
func (options *ListServiceIdsOptions) SetIncludeHistory(includeHistory bool) *ListServiceIdsOptions {
	options.IncludeHistory = core.BoolPtr(includeHistory)
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
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewLockApiKeyOptions : Instantiate LockApiKeyOptions
func (*IamIdentityV1) NewLockApiKeyOptions(id string) *LockApiKeyOptions {
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
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewLockServiceIdOptions : Instantiate LockServiceIdOptions
func (*IamIdentityV1) NewLockServiceIdOptions(id string) *LockServiceIdOptions {
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


// UnmarshalResponseContext unmarshals an instance of ResponseContext from the specified map of raw messages.
func UnmarshalResponseContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResponseContext)
	err = core.UnmarshalPrimitive(m, "transaction_id", &obj.TransactionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user_agent", &obj.UserAgent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "thread_id", &obj.ThreadID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host", &obj.Host)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "elapsed_time", &obj.ElapsedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cluster_name", &obj.ClusterName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceID : Response body format for service ID V1 REST requests.
type ServiceID struct {
	// Context with key properties for problem determination.
	Context *ResponseContext `json:"context,omitempty"`

	// Unique identifier of this Service Id.
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
	CreatedAt *string `json:"created_at,omitempty"`

	// If set contains a date time string of the last modification date in ISO format.
	ModifiedAt *string `json:"modified_at,omitempty"`

	// ID of the account the service ID belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// Name of the Service Id. The name is not checked for uniqueness. Therefore multiple names with the same value can
	// exist. Access is done via the UUID of the Service Id.
	Name *string `json:"name" validate:"required"`

	// The optional description of the Service Id. The 'description' property is only available if a description was
	// provided during a create of a Service Id.
	Description *string `json:"description,omitempty"`

	// Optional list of CRNs (string array) which point to the services connected to the service ID.
	UniqueInstanceCrns []string `json:"unique_instance_crns,omitempty"`

	// History of the Service ID.
	History []EnityHistoryRecord `json:"history,omitempty"`

	// Response body format for API key V1 REST requests.
	Apikey *ApiKey `json:"apikey" validate:"required"`
}


// UnmarshalServiceID unmarshals an instance of ServiceID from the specified map of raw messages.
func UnmarshalServiceID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceID)
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalResponseContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_tag", &obj.EntityTag)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_at", &obj.ModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
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
	err = core.UnmarshalPrimitive(m, "unique_instance_crns", &obj.UniqueInstanceCrns)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "history", &obj.History, UnmarshalEnityHistoryRecord)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "apikey", &obj.Apikey, UnmarshalApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceIdList : Response body format for the list service ID V1 REST request.
type ServiceIdList struct {
	// Context with key properties for problem determination.
	Context *ResponseContext `json:"context,omitempty"`

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
	Serviceids []ServiceID `json:"serviceids" validate:"required"`
}


// UnmarshalServiceIdList unmarshals an instance of ServiceIdList from the specified map of raw messages.
func UnmarshalServiceIdList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceIdList)
	err = core.UnmarshalModel(m, "context", &obj.Context, UnmarshalResponseContext)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "previous", &obj.Previous)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "serviceids", &obj.Serviceids, UnmarshalServiceID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UnlockApiKeyOptions : The UnlockApiKey options.
type UnlockApiKeyOptions struct {
	// Unique ID of the API key.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnlockApiKeyOptions : Instantiate UnlockApiKeyOptions
func (*IamIdentityV1) NewUnlockApiKeyOptions(id string) *UnlockApiKeyOptions {
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
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUnlockServiceIdOptions : Instantiate UnlockServiceIdOptions
func (*IamIdentityV1) NewUnlockServiceIdOptions(id string) *UnlockServiceIdOptions {
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
	ID *string `json:"id" validate:"required,ne="`

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
func (*IamIdentityV1) NewUpdateApiKeyOptions(id string, ifMatch string) *UpdateApiKeyOptions {
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
	ID *string `json:"id" validate:"required,ne="`

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
func (*IamIdentityV1) NewUpdateServiceIdOptions(id string, ifMatch string) *UpdateServiceIdOptions {
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
