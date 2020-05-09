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

// Package iampolicymanagementv1 : Operations and models for the IamPolicyManagementV1 service
package iampolicymanagementv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
	"reflect"
)

// IamPolicyManagementV1 : IAM Policy Management API
//
// Version: 1.0.1
type IamPolicyManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.test.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "iam_policy_management"

// IamPolicyManagementV1Options : Service options
type IamPolicyManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIamPolicyManagementV1UsingExternalConfig : constructs an instance of IamPolicyManagementV1 with passed in options and external configuration.
func NewIamPolicyManagementV1UsingExternalConfig(options *IamPolicyManagementV1Options) (iamPolicyManagement *IamPolicyManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	iamPolicyManagement, err = NewIamPolicyManagementV1(options)
	if err != nil {
		return
	}

	err = iamPolicyManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = iamPolicyManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIamPolicyManagementV1 : constructs an instance of IamPolicyManagementV1 with passed in options.
func NewIamPolicyManagementV1(options *IamPolicyManagementV1Options) (service *IamPolicyManagementV1, err error) {
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

	service = &IamPolicyManagementV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (iamPolicyManagement *IamPolicyManagementV1) SetServiceURL(url string) error {
	return iamPolicyManagement.Service.SetServiceURL(url)
}

// ListPolicies : Get policies by attributes
// Get policies and filter by attributes. While managing policies, you may want to retrieve policies in the account and
// filter by attribute values. This can be done through query parameters. Currently, we only support the following
// attributes: account_id, iam_id, access_group_id, type, and service_type. account_id is a required query parameter.
// Only policies that have the specified attributes and that the caller has read access to are returned. If the caller
// does not have read access to any policies an empty array is returned.
func (iamPolicyManagement *IamPolicyManagementV1) ListPolicies(listPoliciesOptions *ListPoliciesOptions) (result *PolicyList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPoliciesOptions, "listPoliciesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPoliciesOptions, "listPoliciesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/policies"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPoliciesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "ListPolicies")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPoliciesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listPoliciesOptions.AcceptLanguage))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listPoliciesOptions.AccountID))
	if listPoliciesOptions.IamID != nil {
		builder.AddQuery("iam_id", fmt.Sprint(*listPoliciesOptions.IamID))
	}
	if listPoliciesOptions.AccessGroupID != nil {
		builder.AddQuery("access_group_id", fmt.Sprint(*listPoliciesOptions.AccessGroupID))
	}
	if listPoliciesOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listPoliciesOptions.Type))
	}
	if listPoliciesOptions.ServiceType != nil {
		builder.AddQuery("service_type", fmt.Sprint(*listPoliciesOptions.ServiceType))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicyList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreatePolicy : Create a policy
// Creates a policy to grant access between a subject and a resource. There are two types of policies: **access** and
// **authorization**. A policy administrator might want to create an access policy which grants access to a user,
// service-id, or an access group. They might also want to create an authorization policy and setup access between
// services.
// ### Access To create an access policy, use **`"type": "access"`** in the body. The possible subject attributes are
// **`iam_id`** and **`access_group_id`**. Use the **`iam_id`** subject attribute for assigning access for a user or
// service-id. Use the **`access_group_id`** subject attribute for assigning access for an access group. The roles must
// be a subset of a service's or the platform's supported roles. The resource attributes must be a subset of a service's
// or the platform's supported attributes. The policy resource must include either the **`serviceType`**,
// **`serviceName`**,  or **`resourceGroupId`** attribute and the **`accountId`** attribute.` If the subject is a locked
// service-id, the request will fail.
// ### Authorization Authorization policies are supported by services on a case by case basis. Refer to service
// documentation to verify their support of authorization policies. To create an authorization policy, use **`"type":
// "authorization"`** in the body. The subject attributes must match the supported authorization subjects of the
// resource. Multiple subject attributes might be provided. The following attributes are supported:
//   serviceName, serviceInstance, region, resourceType, resource, accountId The policy roles must be a subset of the
// supported authorization roles supported by the target service. The user must also have the same level of access or
// greater to the target resource in order to grant the role. The resource attributes must be a subset of a service's or
// the platform's supported attributes. Both the policy subject and the policy resource must include the
// **`serviceName`** and **`accountId`** attributes.
func (iamPolicyManagement *IamPolicyManagementV1) CreatePolicy(createPolicyOptions *CreatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPolicyOptions, "createPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPolicyOptions, "createPolicyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/policies"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "CreatePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPolicyOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createPolicyOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createPolicyOptions.Type != nil {
		body["type"] = createPolicyOptions.Type
	}
	if createPolicyOptions.Subjects != nil {
		body["subjects"] = createPolicyOptions.Subjects
	}
	if createPolicyOptions.Roles != nil {
		body["roles"] = createPolicyOptions.Roles
	}
	if createPolicyOptions.Resources != nil {
		body["resources"] = createPolicyOptions.Resources
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
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdatePolicy : Update a policy
// Update a policy to grant access between a subject and a resource. A policy administrator might want to update an
// existing policy. The policy type cannot be changed (You cannot change an access policy to an authorization policy).
// ### Access To update an access policy, use **`"type": "access"`** in the body. The possible subject attributes are
// **`iam_id`** and **`access_group_id`**. Use the **`iam_id`** subject attribute for assigning access for a user or
// service-id. Use the **`access_group_id`** subject attribute for assigning access for an access group. The roles must
// be a subset of a service's or the platform's supported roles. The resource attributes must be a subset of a service's
// or the platform's supported attributes. The policy resource must include either the **`serviceType`**,
// **`serviceName`**,  or **`resourceGroupId`** attribute and the **`accountId`** attribute.` If the subject is a locked
// service-id, the request will fail.
// ### Authorization To update an authorization policy, use **`"type": "authorization"`** in the body. The subject
// attributes must match the supported authorization subjects of the resource. Multiple subject attributes might be
// provided. The following attributes are supported:
//   serviceName, serviceInstance, region, resourceType, resource, accountId The policy roles must be a subset of the
// supported authorization roles supported by the target service. The user must also have the same level of access or
// greater to the target resource in order to grant the role. The resource attributes must be a subset of a service's or
// the platform's supported attributes. Both the policy subject and the policy resource must include the
// **`serviceName`** and **`accountId`** attributes.
func (iamPolicyManagement *IamPolicyManagementV1) UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePolicyOptions, "updatePolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePolicyOptions, "updatePolicyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/policies"}
	pathParameters := []string{*updatePolicyOptions.PolicyID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "UpdatePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updatePolicyOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updatePolicyOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updatePolicyOptions.Type != nil {
		body["type"] = updatePolicyOptions.Type
	}
	if updatePolicyOptions.Subjects != nil {
		body["subjects"] = updatePolicyOptions.Subjects
	}
	if updatePolicyOptions.Roles != nil {
		body["roles"] = updatePolicyOptions.Roles
	}
	if updatePolicyOptions.Resources != nil {
		body["resources"] = updatePolicyOptions.Resources
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
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetPolicy : Retrieve a policy by ID
// Retrieve a policy by providing a policy ID.
func (iamPolicyManagement *IamPolicyManagementV1) GetPolicy(getPolicyOptions *GetPolicyOptions) (result *Policy, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPolicyOptions, "getPolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPolicyOptions, "getPolicyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/policies"}
	pathParameters := []string{*getPolicyOptions.PolicyID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "GetPolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPolicy)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeletePolicy : Delete a policy by ID
// Delete a policy by providing a policy ID. A policy cannot be deleted if the subject ID contains a locked service ID.
// If the subject of the policy is a locked service-id, the request will fail.
func (iamPolicyManagement *IamPolicyManagementV1) DeletePolicy(deletePolicyOptions *DeletePolicyOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePolicyOptions, "deletePolicyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePolicyOptions, "deletePolicyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/policies"}
	pathParameters := []string{*deletePolicyOptions.PolicyID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePolicyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "DeletePolicy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamPolicyManagement.Service.Request(request, nil)

	return
}

// ListRoles : Get roles by filters
// Get roles based on the filters. While managing roles, you may want to retrieve roles and filter by usages. This can
// be done through query parameters. Currently, we only support the following attributes: account_id, and service_name.
// Only roles that match the filter and that the caller has read access to are returned. If the caller does not have
// read access to any roles an empty array is returned.
func (iamPolicyManagement *IamPolicyManagementV1) ListRoles(listRolesOptions *ListRolesOptions) (result *RoleList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRolesOptions, "listRolesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/roles"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "ListRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRolesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listRolesOptions.AcceptLanguage))
	}

	if listRolesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listRolesOptions.AccountID))
	}
	if listRolesOptions.ServiceName != nil {
		builder.AddQuery("service_name", fmt.Sprint(*listRolesOptions.ServiceName))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoleList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateRole : Create a role
// Creates a custom role for a specific service within the account. An account owner or a user assigned the
// Administrator role on the Role management service can create a custom role. Any number of actions for a single
// service can be mapped to the new role, but there must be at least one service-defined action to successfully create
// the new role.
func (iamPolicyManagement *IamPolicyManagementV1) CreateRole(createRoleOptions *CreateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRoleOptions, "createRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRoleOptions, "createRoleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/roles"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "CreateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRoleOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createRoleOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createRoleOptions.Name != nil {
		body["name"] = createRoleOptions.Name
	}
	if createRoleOptions.AccountID != nil {
		body["account_id"] = createRoleOptions.AccountID
	}
	if createRoleOptions.ServiceName != nil {
		body["service_name"] = createRoleOptions.ServiceName
	}
	if createRoleOptions.DisplayName != nil {
		body["display_name"] = createRoleOptions.DisplayName
	}
	if createRoleOptions.Actions != nil {
		body["actions"] = createRoleOptions.Actions
	}
	if createRoleOptions.Description != nil {
		body["description"] = createRoleOptions.Description
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
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateRole : Update a role
// Update a custom role. A role administrator might want to update an existing role by updating the display name,
// description, or the actions that are mapped to the role. The name, account_id, and service_name can't be changed.
func (iamPolicyManagement *IamPolicyManagementV1) UpdateRole(updateRoleOptions *UpdateRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRoleOptions, "updateRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRoleOptions, "updateRoleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/roles"}
	pathParameters := []string{*updateRoleOptions.RoleID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "UpdateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateRoleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateRoleOptions.IfMatch))
	}

	body := make(map[string]interface{})
	if updateRoleOptions.DisplayName != nil {
		body["display_name"] = updateRoleOptions.DisplayName
	}
	if updateRoleOptions.Description != nil {
		body["description"] = updateRoleOptions.Description
	}
	if updateRoleOptions.Actions != nil {
		body["actions"] = updateRoleOptions.Actions
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
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRole : Retrieve a role by ID
// Retrieve a role by providing a role ID.
func (iamPolicyManagement *IamPolicyManagementV1) GetRole(getRoleOptions *GetRoleOptions) (result *CustomRole, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRoleOptions, "getRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRoleOptions, "getRoleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/roles"}
	pathParameters := []string{*getRoleOptions.RoleID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "GetRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamPolicyManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomRole)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteRole : Delete a role by ID
// Delete a role by providing a role ID.
func (iamPolicyManagement *IamPolicyManagementV1) DeleteRole(deleteRoleOptions *DeleteRoleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRoleOptions, "deleteRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRoleOptions, "deleteRoleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/roles"}
	pathParameters := []string{*deleteRoleOptions.RoleID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamPolicyManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_policy_management", "V1", "DeleteRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamPolicyManagement.Service.Request(request, nil)

	return
}

// CreatePolicyOptions : The CreatePolicy options.
type CreatePolicyOptions struct {
	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type" validate:"required"`

	// The subject attribute values that must match in order for this policy to apply in a permission decision.
	Subjects []PolicyRequestSubjectsItem `json:"subjects" validate:"required"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRequestRolesItem `json:"roles" validate:"required"`

	// The attributes of the resource. Note that only one resource is allowed in a policy.
	Resources []PolicyRequestResourcesItem `json:"resources" validate:"required"`

	// Translation language code.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreatePolicyOptions : Instantiate CreatePolicyOptions
func (*IamPolicyManagementV1) NewCreatePolicyOptions(typeVar string, subjects []PolicyRequestSubjectsItem, roles []PolicyRequestRolesItem, resources []PolicyRequestResourcesItem) *CreatePolicyOptions {
	return &CreatePolicyOptions{
		Type: core.StringPtr(typeVar),
		Subjects: subjects,
		Roles: roles,
		Resources: resources,
	}
}

// SetType : Allow user to set Type
func (options *CreatePolicyOptions) SetType(typeVar string) *CreatePolicyOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetSubjects : Allow user to set Subjects
func (options *CreatePolicyOptions) SetSubjects(subjects []PolicyRequestSubjectsItem) *CreatePolicyOptions {
	options.Subjects = subjects
	return options
}

// SetRoles : Allow user to set Roles
func (options *CreatePolicyOptions) SetRoles(roles []PolicyRequestRolesItem) *CreatePolicyOptions {
	options.Roles = roles
	return options
}

// SetResources : Allow user to set Resources
func (options *CreatePolicyOptions) SetResources(resources []PolicyRequestResourcesItem) *CreatePolicyOptions {
	options.Resources = resources
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *CreatePolicyOptions) SetAcceptLanguage(acceptLanguage string) *CreatePolicyOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePolicyOptions) SetHeaders(param map[string]string) *CreatePolicyOptions {
	options.Headers = param
	return options
}

// CreateRoleOptions : The CreateRole options.
type CreateRoleOptions struct {
	// The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.
	Name *string `json:"name" validate:"required"`

	// The account GUID.
	AccountID *string `json:"account_id" validate:"required"`

	// The service name.
	ServiceName *string `json:"service_name" validate:"required"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name" validate:"required"`

	// The actions of the role.
	Actions []string `json:"actions" validate:"required"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// Translation language code.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRoleOptions : Instantiate CreateRoleOptions
func (*IamPolicyManagementV1) NewCreateRoleOptions(name string, accountID string, serviceName string, displayName string, actions []string) *CreateRoleOptions {
	return &CreateRoleOptions{
		Name: core.StringPtr(name),
		AccountID: core.StringPtr(accountID),
		ServiceName: core.StringPtr(serviceName),
		DisplayName: core.StringPtr(displayName),
		Actions: actions,
	}
}

// SetName : Allow user to set Name
func (options *CreateRoleOptions) SetName(name string) *CreateRoleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *CreateRoleOptions) SetAccountID(accountID string) *CreateRoleOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetServiceName : Allow user to set ServiceName
func (options *CreateRoleOptions) SetServiceName(serviceName string) *CreateRoleOptions {
	options.ServiceName = core.StringPtr(serviceName)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *CreateRoleOptions) SetDisplayName(displayName string) *CreateRoleOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetActions : Allow user to set Actions
func (options *CreateRoleOptions) SetActions(actions []string) *CreateRoleOptions {
	options.Actions = actions
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateRoleOptions) SetDescription(description string) *CreateRoleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *CreateRoleOptions) SetAcceptLanguage(acceptLanguage string) *CreateRoleOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRoleOptions) SetHeaders(param map[string]string) *CreateRoleOptions {
	options.Headers = param
	return options
}

// DeletePolicyOptions : The DeletePolicy options.
type DeletePolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePolicyOptions : Instantiate DeletePolicyOptions
func (*IamPolicyManagementV1) NewDeletePolicyOptions(policyID string) *DeletePolicyOptions {
	return &DeletePolicyOptions{
		PolicyID: core.StringPtr(policyID),
	}
}

// SetPolicyID : Allow user to set PolicyID
func (options *DeletePolicyOptions) SetPolicyID(policyID string) *DeletePolicyOptions {
	options.PolicyID = core.StringPtr(policyID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePolicyOptions) SetHeaders(param map[string]string) *DeletePolicyOptions {
	options.Headers = param
	return options
}

// DeleteRoleOptions : The DeleteRole options.
type DeleteRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRoleOptions : Instantiate DeleteRoleOptions
func (*IamPolicyManagementV1) NewDeleteRoleOptions(roleID string) *DeleteRoleOptions {
	return &DeleteRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (options *DeleteRoleOptions) SetRoleID(roleID string) *DeleteRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRoleOptions) SetHeaders(param map[string]string) *DeleteRoleOptions {
	options.Headers = param
	return options
}

// GetPolicyOptions : The GetPolicy options.
type GetPolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPolicyOptions : Instantiate GetPolicyOptions
func (*IamPolicyManagementV1) NewGetPolicyOptions(policyID string) *GetPolicyOptions {
	return &GetPolicyOptions{
		PolicyID: core.StringPtr(policyID),
	}
}

// SetPolicyID : Allow user to set PolicyID
func (options *GetPolicyOptions) SetPolicyID(policyID string) *GetPolicyOptions {
	options.PolicyID = core.StringPtr(policyID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPolicyOptions) SetHeaders(param map[string]string) *GetPolicyOptions {
	options.Headers = param
	return options
}

// GetRoleOptions : The GetRole options.
type GetRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRoleOptions : Instantiate GetRoleOptions
func (*IamPolicyManagementV1) NewGetRoleOptions(roleID string) *GetRoleOptions {
	return &GetRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (options *GetRoleOptions) SetRoleID(roleID string) *GetRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRoleOptions) SetHeaders(param map[string]string) *GetRoleOptions {
	options.Headers = param
	return options
}

// ListPoliciesOptions : The ListPolicies options.
type ListPoliciesOptions struct {
	// The account GUID in which the policies belong to.
	AccountID *string `json:"account_id" validate:"required"`

	// Translation language code.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// The IAM ID used to identify the subject.
	IamID *string `json:"iam_id,omitempty"`

	// The access group id.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The type of policy (access or authorization).
	Type *string `json:"type,omitempty"`

	// The type of service.
	ServiceType *string `json:"service_type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPoliciesOptions : Instantiate ListPoliciesOptions
func (*IamPolicyManagementV1) NewListPoliciesOptions(accountID string) *ListPoliciesOptions {
	return &ListPoliciesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListPoliciesOptions) SetAccountID(accountID string) *ListPoliciesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ListPoliciesOptions) SetAcceptLanguage(acceptLanguage string) *ListPoliciesOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetIamID : Allow user to set IamID
func (options *ListPoliciesOptions) SetIamID(iamID string) *ListPoliciesOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *ListPoliciesOptions) SetAccessGroupID(accessGroupID string) *ListPoliciesOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetType : Allow user to set Type
func (options *ListPoliciesOptions) SetType(typeVar string) *ListPoliciesOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetServiceType : Allow user to set ServiceType
func (options *ListPoliciesOptions) SetServiceType(serviceType string) *ListPoliciesOptions {
	options.ServiceType = core.StringPtr(serviceType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListPoliciesOptions) SetHeaders(param map[string]string) *ListPoliciesOptions {
	options.Headers = param
	return options
}

// ListRolesOptions : The ListRoles options.
type ListRolesOptions struct {
	// Translation language code.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// The account GUID in which the roles belong to.
	AccountID *string `json:"account_id,omitempty"`

	// The name of service.
	ServiceName *string `json:"service_name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRolesOptions : Instantiate ListRolesOptions
func (*IamPolicyManagementV1) NewListRolesOptions() *ListRolesOptions {
	return &ListRolesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *ListRolesOptions) SetAcceptLanguage(acceptLanguage string) *ListRolesOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *ListRolesOptions) SetAccountID(accountID string) *ListRolesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetServiceName : Allow user to set ServiceName
func (options *ListRolesOptions) SetServiceName(serviceName string) *ListRolesOptions {
	options.ServiceName = core.StringPtr(serviceName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListRolesOptions) SetHeaders(param map[string]string) *ListRolesOptions {
	options.Headers = param
	return options
}

// PolicyBaseResourcesItem : PolicyBaseResourcesItem struct
type PolicyBaseResourcesItem struct {
	// List of resource attributes.
	Attributes []PolicyBaseResourcesItemAttributesItem `json:"attributes,omitempty"`
}


// UnmarshalPolicyBaseResourcesItem unmarshals an instance of PolicyBaseResourcesItem from the specified map of raw messages.
func UnmarshalPolicyBaseResourcesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyBaseResourcesItem)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalPolicyBaseResourcesItemAttributesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyBaseResourcesItemAttributesItem : PolicyBaseResourcesItemAttributesItem struct
type PolicyBaseResourcesItemAttributesItem struct {
	// The name of an attribute.
	Name *string `json:"name,omitempty"`

	// The value of an attribute.
	Value *string `json:"value,omitempty"`

	// The operator of an attribute.
	Operator *string `json:"operator,omitempty"`
}


// UnmarshalPolicyBaseResourcesItemAttributesItem unmarshals an instance of PolicyBaseResourcesItemAttributesItem from the specified map of raw messages.
func UnmarshalPolicyBaseResourcesItemAttributesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyBaseResourcesItemAttributesItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyBaseSubjectsItem : PolicyBaseSubjectsItem struct
type PolicyBaseSubjectsItem struct {
	// List of subject attributes.
	Attributes []PolicyBaseSubjectsItemAttributesItem `json:"attributes,omitempty"`
}


// UnmarshalPolicyBaseSubjectsItem unmarshals an instance of PolicyBaseSubjectsItem from the specified map of raw messages.
func UnmarshalPolicyBaseSubjectsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyBaseSubjectsItem)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalPolicyBaseSubjectsItemAttributesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyBaseSubjectsItemAttributesItem : PolicyBaseSubjectsItemAttributesItem struct
type PolicyBaseSubjectsItemAttributesItem struct {
	// The name of an attribute.
	Name *string `json:"name,omitempty"`

	// The value of an attribute.
	Value *string `json:"value,omitempty"`
}


// UnmarshalPolicyBaseSubjectsItemAttributesItem unmarshals an instance of PolicyBaseSubjectsItemAttributesItem from the specified map of raw messages.
func UnmarshalPolicyBaseSubjectsItemAttributesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyBaseSubjectsItemAttributesItem)
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

// PolicyRequestResourcesItem : PolicyRequestResourcesItem struct
type PolicyRequestResourcesItem struct {
	// List of resource attributes.
	Attributes []PolicyRequestResourcesItemAttributesItem `json:"attributes" validate:"required"`
}


// NewPolicyRequestResourcesItem : Instantiate PolicyRequestResourcesItem (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRequestResourcesItem(attributes []PolicyRequestResourcesItemAttributesItem) (model *PolicyRequestResourcesItem, err error) {
	model = &PolicyRequestResourcesItem{
		Attributes: attributes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPolicyRequestResourcesItem unmarshals an instance of PolicyRequestResourcesItem from the specified map of raw messages.
func UnmarshalPolicyRequestResourcesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRequestResourcesItem)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalPolicyRequestResourcesItemAttributesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyRequestResourcesItemAttributesItem : PolicyRequestResourcesItemAttributesItem struct
type PolicyRequestResourcesItemAttributesItem struct {
	// The name of an attribute.
	Name *string `json:"name" validate:"required"`

	// The value of an attribute.
	Value *string `json:"value" validate:"required"`

	// The operator of an attribute.
	Operator *string `json:"operator,omitempty"`
}


// NewPolicyRequestResourcesItemAttributesItem : Instantiate PolicyRequestResourcesItemAttributesItem (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRequestResourcesItemAttributesItem(name string, value string) (model *PolicyRequestResourcesItemAttributesItem, err error) {
	model = &PolicyRequestResourcesItemAttributesItem{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPolicyRequestResourcesItemAttributesItem unmarshals an instance of PolicyRequestResourcesItemAttributesItem from the specified map of raw messages.
func UnmarshalPolicyRequestResourcesItemAttributesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRequestResourcesItemAttributesItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyRequestRolesItem : PolicyRequestRolesItem struct
type PolicyRequestRolesItem struct {
	// A role cloud resource name (CRN).
	RoleID *string `json:"role_id" validate:"required"`
}


// NewPolicyRequestRolesItem : Instantiate PolicyRequestRolesItem (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRequestRolesItem(roleID string) (model *PolicyRequestRolesItem, err error) {
	model = &PolicyRequestRolesItem{
		RoleID: core.StringPtr(roleID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPolicyRequestRolesItem unmarshals an instance of PolicyRequestRolesItem from the specified map of raw messages.
func UnmarshalPolicyRequestRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRequestRolesItem)
	err = core.UnmarshalPrimitive(m, "role_id", &obj.RoleID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyRequestSubjectsItem : PolicyRequestSubjectsItem struct
type PolicyRequestSubjectsItem struct {
	// List of subject attributes.
	Attributes []PolicyRequestSubjectsItemAttributesItem `json:"attributes" validate:"required"`
}


// NewPolicyRequestSubjectsItem : Instantiate PolicyRequestSubjectsItem (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRequestSubjectsItem(attributes []PolicyRequestSubjectsItemAttributesItem) (model *PolicyRequestSubjectsItem, err error) {
	model = &PolicyRequestSubjectsItem{
		Attributes: attributes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPolicyRequestSubjectsItem unmarshals an instance of PolicyRequestSubjectsItem from the specified map of raw messages.
func UnmarshalPolicyRequestSubjectsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRequestSubjectsItem)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalPolicyRequestSubjectsItemAttributesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyRequestSubjectsItemAttributesItem : PolicyRequestSubjectsItemAttributesItem struct
type PolicyRequestSubjectsItemAttributesItem struct {
	// The name of an attribute.
	Name *string `json:"name" validate:"required"`

	// The value of an attribute.
	Value *string `json:"value" validate:"required"`
}


// NewPolicyRequestSubjectsItemAttributesItem : Instantiate PolicyRequestSubjectsItemAttributesItem (Generic Model Constructor)
func (*IamPolicyManagementV1) NewPolicyRequestSubjectsItemAttributesItem(name string, value string) (model *PolicyRequestSubjectsItemAttributesItem, err error) {
	model = &PolicyRequestSubjectsItemAttributesItem{
		Name: core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalPolicyRequestSubjectsItemAttributesItem unmarshals an instance of PolicyRequestSubjectsItemAttributesItem from the specified map of raw messages.
func UnmarshalPolicyRequestSubjectsItemAttributesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRequestSubjectsItemAttributesItem)
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

// PolicyRolesItem : PolicyRolesItem struct
type PolicyRolesItem struct {
	// The role cloud resource name granted by the policy.
	RoleID *string `json:"role_id,omitempty"`

	// The display name of the role.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`
}


// UnmarshalPolicyRolesItem unmarshals an instance of PolicyRolesItem from the specified map of raw messages.
func UnmarshalPolicyRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyRolesItem)
	err = core.UnmarshalPrimitive(m, "role_id", &obj.RoleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
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

// RoleListCustomRolesItem : RoleListCustomRolesItem struct
type RoleListCustomRolesItem struct {
	// The role ID.
	ID *string `json:"id,omitempty"`

	// The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.
	Name *string `json:"name,omitempty"`

	// The account GUID.
	AccountID *string `json:"account_id,omitempty"`

	// The service name.
	ServiceName *string `json:"service_name,omitempty"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The role CRN.
	Crn *string `json:"crn,omitempty"`

	// The actions of the role.
	Actions []string `json:"actions,omitempty"`

	// The UTC timestamp when the role was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The iam ID of the entity that created the role.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The UTC timestamp when the role was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The iam ID of the entity that last modified the policy.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// The href link back to the role.
	Href *string `json:"href,omitempty"`
}


// UnmarshalRoleListCustomRolesItem unmarshals an instance of RoleListCustomRolesItem from the specified map of raw messages.
func UnmarshalRoleListCustomRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoleListCustomRolesItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_at", &obj.LastModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_by_id", &obj.LastModifiedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RoleListServiceRolesItem : RoleListServiceRolesItem struct
type RoleListServiceRolesItem struct {
	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The role CRN.
	Crn *string `json:"crn,omitempty"`

	// The actions of the role.
	Actions []string `json:"actions,omitempty"`
}


// UnmarshalRoleListServiceRolesItem unmarshals an instance of RoleListServiceRolesItem from the specified map of raw messages.
func UnmarshalRoleListServiceRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoleListServiceRolesItem)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RoleListSystemRolesItem : RoleListSystemRolesItem struct
type RoleListSystemRolesItem struct {
	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The role CRN.
	Crn *string `json:"crn,omitempty"`

	// The actions of the role.
	Actions []string `json:"actions,omitempty"`
}


// UnmarshalRoleListSystemRolesItem unmarshals an instance of RoleListSystemRolesItem from the specified map of raw messages.
func UnmarshalRoleListSystemRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoleListSystemRolesItem)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdatePolicyOptions : The UpdatePolicy options.
type UpdatePolicyOptions struct {
	// The policy ID.
	PolicyID *string `json:"policy_id" validate:"required"`

	// The revision number for updating a policy and must match the ETag value of the existing policy. The Etag can be
	// retrieved using the GET /v1/policies/{policy_id} API and looking at the ETag response header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type" validate:"required"`

	// The subject attribute values that must match in order for this policy to apply in a permission decision.
	Subjects []PolicyRequestSubjectsItem `json:"subjects" validate:"required"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRequestRolesItem `json:"roles" validate:"required"`

	// The attributes of the resource. Note that only one resource is allowed in a policy.
	Resources []PolicyRequestResourcesItem `json:"resources" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdatePolicyOptions : Instantiate UpdatePolicyOptions
func (*IamPolicyManagementV1) NewUpdatePolicyOptions(policyID string, ifMatch string, typeVar string, subjects []PolicyRequestSubjectsItem, roles []PolicyRequestRolesItem, resources []PolicyRequestResourcesItem) *UpdatePolicyOptions {
	return &UpdatePolicyOptions{
		PolicyID: core.StringPtr(policyID),
		IfMatch: core.StringPtr(ifMatch),
		Type: core.StringPtr(typeVar),
		Subjects: subjects,
		Roles: roles,
		Resources: resources,
	}
}

// SetPolicyID : Allow user to set PolicyID
func (options *UpdatePolicyOptions) SetPolicyID(policyID string) *UpdatePolicyOptions {
	options.PolicyID = core.StringPtr(policyID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdatePolicyOptions) SetIfMatch(ifMatch string) *UpdatePolicyOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetType : Allow user to set Type
func (options *UpdatePolicyOptions) SetType(typeVar string) *UpdatePolicyOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetSubjects : Allow user to set Subjects
func (options *UpdatePolicyOptions) SetSubjects(subjects []PolicyRequestSubjectsItem) *UpdatePolicyOptions {
	options.Subjects = subjects
	return options
}

// SetRoles : Allow user to set Roles
func (options *UpdatePolicyOptions) SetRoles(roles []PolicyRequestRolesItem) *UpdatePolicyOptions {
	options.Roles = roles
	return options
}

// SetResources : Allow user to set Resources
func (options *UpdatePolicyOptions) SetResources(resources []PolicyRequestResourcesItem) *UpdatePolicyOptions {
	options.Resources = resources
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePolicyOptions) SetHeaders(param map[string]string) *UpdatePolicyOptions {
	options.Headers = param
	return options
}

// UpdateRoleOptions : The UpdateRole options.
type UpdateRoleOptions struct {
	// The role ID.
	RoleID *string `json:"role_id" validate:"required"`

	// The revision number for updating a role and must match the ETag value of the existing role. The Etag can be
	// retrieved using the GET /v2/roles/{role_id} API and looking at the ETag response header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The actions of the role.
	Actions []string `json:"actions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRoleOptions : Instantiate UpdateRoleOptions
func (*IamPolicyManagementV1) NewUpdateRoleOptions(roleID string, ifMatch string) *UpdateRoleOptions {
	return &UpdateRoleOptions{
		RoleID: core.StringPtr(roleID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetRoleID : Allow user to set RoleID
func (options *UpdateRoleOptions) SetRoleID(roleID string) *UpdateRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateRoleOptions) SetIfMatch(ifMatch string) *UpdateRoleOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetDisplayName : Allow user to set DisplayName
func (options *UpdateRoleOptions) SetDisplayName(displayName string) *UpdateRoleOptions {
	options.DisplayName = core.StringPtr(displayName)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateRoleOptions) SetDescription(description string) *UpdateRoleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetActions : Allow user to set Actions
func (options *UpdateRoleOptions) SetActions(actions []string) *UpdateRoleOptions {
	options.Actions = actions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRoleOptions) SetHeaders(param map[string]string) *UpdateRoleOptions {
	options.Headers = param
	return options
}

// CustomRole : CustomRole struct
type CustomRole struct {
	// The role ID.
	ID *string `json:"id,omitempty"`

	// The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.
	Name *string `json:"name,omitempty"`

	// The account GUID.
	AccountID *string `json:"account_id,omitempty"`

	// The service name.
	ServiceName *string `json:"service_name,omitempty"`

	// The display name of the role that is shown in the console.
	DisplayName *string `json:"display_name,omitempty"`

	// The description of the role.
	Description *string `json:"description,omitempty"`

	// The role CRN.
	Crn *string `json:"crn,omitempty"`

	// The actions of the role.
	Actions []string `json:"actions,omitempty"`

	// The UTC timestamp when the role was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The iam ID of the entity that created the role.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The UTC timestamp when the role was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The iam ID of the entity that last modified the policy.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// The href link back to the role.
	Href *string `json:"href,omitempty"`
}


// UnmarshalCustomRole unmarshals an instance of CustomRole from the specified map of raw messages.
func UnmarshalCustomRole(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomRole)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "actions", &obj.Actions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_at", &obj.LastModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_by_id", &obj.LastModifiedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Policy : Policy struct
type Policy struct {
	// The policy ID.
	ID *string `json:"id,omitempty"`

	// The policy type; either 'access' or 'authorization'.
	Type *string `json:"type,omitempty"`

	// The subject attribute values that must match in order for this policy to apply in a permission decision.
	Subjects []PolicyBaseSubjectsItem `json:"subjects,omitempty"`

	// A set of role cloud resource names (CRNs) granted by the policy.
	Roles []PolicyRolesItem `json:"roles,omitempty"`

	// The attributes of the resource. Note that only one resource is allowed in a policy.
	Resources []PolicyBaseResourcesItem `json:"resources,omitempty"`

	// The href link back to the policy.
	Href *string `json:"href,omitempty"`

	// The UTC timestamp when the policy was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The iam ID of the entity that created the policy.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The UTC timestamp when the policy was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The iam ID of the entity that last modified the policy.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`
}


// UnmarshalPolicy unmarshals an instance of Policy from the specified map of raw messages.
func UnmarshalPolicy(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Policy)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subjects", &obj.Subjects, UnmarshalPolicyBaseSubjectsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalPolicyRolesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalPolicyBaseResourcesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_at", &obj.LastModifiedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified_by_id", &obj.LastModifiedByID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyList : PolicyList struct
type PolicyList struct {
	// List of policies.
	Policies []Policy `json:"policies,omitempty"`
}


// UnmarshalPolicyList unmarshals an instance of PolicyList from the specified map of raw messages.
func UnmarshalPolicyList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyList)
	err = core.UnmarshalModel(m, "policies", &obj.Policies, UnmarshalPolicy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RoleList : RoleList struct
type RoleList struct {
	// List of custom roles.
	CustomRoles []RoleListCustomRolesItem `json:"custom_roles,omitempty"`

	// List of service roles.
	ServiceRoles []RoleListServiceRolesItem `json:"service_roles,omitempty"`

	// List of system roles.
	SystemRoles []RoleListSystemRolesItem `json:"system_roles,omitempty"`
}


// UnmarshalRoleList unmarshals an instance of RoleList from the specified map of raw messages.
func UnmarshalRoleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RoleList)
	err = core.UnmarshalModel(m, "custom_roles", &obj.CustomRoles, UnmarshalRoleListCustomRolesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "service_roles", &obj.ServiceRoles, UnmarshalRoleListServiceRolesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "system_roles", &obj.SystemRoles, UnmarshalRoleListSystemRolesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
