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

// Package iamaccessgroupsv2 : Operations and models for the IamAccessGroupsV2 service
package iamaccessgroupsv2

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// IamAccessGroupsV2 : The IAM Access Groups API allows for the management of Access Groups (Create, Read, Update,
// Delete) as well as the management of memberships and rules within the group container.
//
// Version: 1.0
type IamAccessGroupsV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.cloud.ibm.com/v2"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "iam_access_groups"

// IamAccessGroupsV2Options : Service options
type IamAccessGroupsV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIamAccessGroupsV2UsingExternalConfig : constructs an instance of IamAccessGroupsV2 with passed in options and external configuration.
func NewIamAccessGroupsV2UsingExternalConfig(options *IamAccessGroupsV2Options) (iamAccessGroups *IamAccessGroupsV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	iamAccessGroups, err = NewIamAccessGroupsV2(options)
	if err != nil {
		return
	}

	err = iamAccessGroups.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = iamAccessGroups.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIamAccessGroupsV2 : constructs an instance of IamAccessGroupsV2 with passed in options.
func NewIamAccessGroupsV2(options *IamAccessGroupsV2Options) (service *IamAccessGroupsV2, err error) {
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

	service = &IamAccessGroupsV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (iamAccessGroups *IamAccessGroupsV2) SetServiceURL(url string) error {
	return iamAccessGroups.Service.SetServiceURL(url)
}

// CreateAccessGroup : Create an Access Group
// Create a new Access Group to assign multiple users and service ids to multiple policies. The group will be created in
// the account specified by the `account_id` parameter. The group name is a required field, but a description is
// optional. Because the group's name does not have to be unique, it is possible to create multiple groups with the same
// name.
func (iamAccessGroups *IamAccessGroupsV2) CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAccessGroupOptions, "createAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAccessGroupOptions, "createAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "CreateAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createAccessGroupOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*createAccessGroupOptions.AccountID))

	body := make(map[string]interface{})
	if createAccessGroupOptions.Name != nil {
		body["name"] = createAccessGroupOptions.Name
	}
	if createAccessGroupOptions.Description != nil {
		body["description"] = createAccessGroupOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalGroup(m)
		response.Result = result
	}

	return
}

// ListAccessGroups : List Access Groups
// This API lists Access Groups within an account. Parameters for pagination and sorting can be used to filter the
// results. The `account_id` query parameter determines which account to retrieve groups from. Only the groups you have
// access to are returned (either because of a policy on a specific group or account level access (admin, editor, or
// viewer)). There may be more groups in the account that aren't shown if you lack the aforementioned permissions.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) (result *GroupsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupsOptions, "listAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupsOptions, "listAccessGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAccessGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListAccessGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAccessGroupsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAccessGroupsOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listAccessGroupsOptions.AccountID))
	if listAccessGroupsOptions.IamID != nil {
		builder.AddQuery("iam_id", fmt.Sprint(*listAccessGroupsOptions.IamID))
	}
	if listAccessGroupsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAccessGroupsOptions.Limit))
	}
	if listAccessGroupsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listAccessGroupsOptions.Offset))
	}
	if listAccessGroupsOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listAccessGroupsOptions.Sort))
	}
	if listAccessGroupsOptions.ShowFederated != nil {
		builder.AddQuery("show_federated", fmt.Sprint(*listAccessGroupsOptions.ShowFederated))
	}
	if listAccessGroupsOptions.HidePublicAccess != nil {
		builder.AddQuery("hide_public_access", fmt.Sprint(*listAccessGroupsOptions.HidePublicAccess))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalGroupsList(m)
		response.Result = result
	}

	return
}

// GetAccessGroup : Get an Access Group
// Retrieve an Access Group by its `access_group_id`. Only the groups data is returned (group name, description,
// account_id, ...), not membership or rule information. A revision number is returned in the `Etag` header, which is
// needed when updating the Access Group.
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccessGroupOptions, "getAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccessGroupOptions, "getAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups"}
	pathParameters := []string{*getAccessGroupOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAccessGroupOptions.TransactionID))
	}

	if getAccessGroupOptions.ShowFederated != nil {
		builder.AddQuery("show_federated", fmt.Sprint(*getAccessGroupOptions.ShowFederated))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalGroup(m)
		response.Result = result
	}

	return
}

// UpdateAccessGroup : Update an Access Group
// Update the group name or description of an existing Access Group using this API. An `If-Match` header must be
// populated with the group's most recent revision number (which can be acquired in the `Get an Access Group` API).
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccessGroupOptions, "updateAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccessGroupOptions, "updateAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups"}
	pathParameters := []string{*updateAccessGroupOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "UpdateAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateAccessGroupOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateAccessGroupOptions.IfMatch))
	}
	if updateAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateAccessGroupOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateAccessGroupOptions.Name != nil {
		body["name"] = updateAccessGroupOptions.Name
	}
	if updateAccessGroupOptions.Description != nil {
		body["description"] = updateAccessGroupOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalGroup(m)
		response.Result = result
	}

	return
}

// DeleteAccessGroup : Delete an Access Group
// This API is used for deleting an Access Group. If the Access Group has no members or rules associated with it, the
// group and its policies will be deleted. However, if rules or members do exist, set the `force` parameter to true to
// delete the group as well as its associated members, rules, and policies.
func (iamAccessGroups *IamAccessGroupsV2) DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccessGroupOptions, "deleteAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccessGroupOptions, "deleteAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups"}
	pathParameters := []string{*deleteAccessGroupOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "DeleteAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteAccessGroupOptions.TransactionID))
	}

	if deleteAccessGroupOptions.Force != nil {
		builder.AddQuery("force", fmt.Sprint(*deleteAccessGroupOptions.Force))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// GetAccountSettings : Get Account Settings
// Retrieve the Access Groups settings for a specific account.
func (iamAccessGroups *IamAccessGroupsV2) GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSettingsOptions, "getAccountSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSettingsOptions, "getAccountSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups/settings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetAccountSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAccountSettingsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAccountSettingsOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*getAccountSettingsOptions.AccountID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAccountSettings(m)
		response.Result = result
	}

	return
}

// UpdateAccountSettings : Update Account Settings
// Update the Access Groups settings for a specific account. Note: When the `public_access_enabled` setting is set to
// false, all policies within the account attached to the Public Access group will be deleted. Only set
// `public_access_enabled` to false if you are sure that you want those policies to be removed.
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountSettingsOptions, "updateAccountSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountSettingsOptions, "updateAccountSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups/settings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "UpdateAccountSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateAccountSettingsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateAccountSettingsOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*updateAccountSettingsOptions.AccountID))

	body := make(map[string]interface{})
	if updateAccountSettingsOptions.PublicAccessEnabled != nil {
		body["public_access_enabled"] = updateAccountSettingsOptions.PublicAccessEnabled
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAccountSettings(m)
		response.Result = result
	}

	return
}

// IsMemberOfAccessGroup : Check membership in an Access Group
// This HEAD operation determines if a given `iam_id` is present in a group. No response body is returned with this
// request. If the membership exists, a `204 - No Content` status code is returned. If the membership or the group does
// not exist, a `404 - Not Found` status code is returned.
func (iamAccessGroups *IamAccessGroupsV2) IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(isMemberOfAccessGroupOptions, "isMemberOfAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(isMemberOfAccessGroupOptions, "isMemberOfAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "members"}
	pathParameters := []string{*isMemberOfAccessGroupOptions.AccessGroupID, *isMemberOfAccessGroupOptions.IamID}

	builder := core.NewRequestBuilder(core.HEAD)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range isMemberOfAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "IsMemberOfAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if isMemberOfAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*isMemberOfAccessGroupOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// AddMembersToAccessGroup : Add members to an Access Group
// Use this API to add users (`IBMid-...`) or service IDs (`iam-ServiceId-...`) to an Access Group. Any member added
// gains access to resources defined in the group's policies. To revoke a given user's access, simply remove them from
// the group. There is no limit to the number of members one group can have, but each `iam_id` can only be added to 50
// groups. Additionally, this API request payload can add up to 50 members per call.
func (iamAccessGroups *IamAccessGroupsV2) AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions) (result *AddGroupMembersResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addMembersToAccessGroupOptions, "addMembersToAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addMembersToAccessGroupOptions, "addMembersToAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "members"}
	pathParameters := []string{*addMembersToAccessGroupOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addMembersToAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "AddMembersToAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addMembersToAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addMembersToAccessGroupOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if addMembersToAccessGroupOptions.Members != nil {
		body["members"] = addMembersToAccessGroupOptions.Members
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAddGroupMembersResponse(m)
		response.Result = result
	}

	return
}

// ListAccessGroupMembers : List Access Group members
// List all members of a given group using this API. Parameters for pagination and sorting can be used to filter the
// results. The most useful query parameter may be the `verbose` flag. If `verbose=true`, user and service ID names will
// be retrieved for each `iam_id`. If performance is a concern, leave the `verbose` parameter off so that name
// information does not get retrieved.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions) (result *GroupMembersList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupMembersOptions, "listAccessGroupMembersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupMembersOptions, "listAccessGroupMembersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "members"}
	pathParameters := []string{*listAccessGroupMembersOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAccessGroupMembersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListAccessGroupMembers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAccessGroupMembersOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAccessGroupMembersOptions.TransactionID))
	}

	if listAccessGroupMembersOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAccessGroupMembersOptions.Limit))
	}
	if listAccessGroupMembersOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listAccessGroupMembersOptions.Offset))
	}
	if listAccessGroupMembersOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listAccessGroupMembersOptions.Type))
	}
	if listAccessGroupMembersOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*listAccessGroupMembersOptions.Verbose))
	}
	if listAccessGroupMembersOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listAccessGroupMembersOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalGroupMembersList(m)
		response.Result = result
	}

	return
}

// RemoveMemberFromAccessGroup : Delete member from an Access Group
// Remove one member from a group using this API. If the operation is successful, only a `204 - No Content` response
// with no body is returned. However, if any error occurs, the standard error format will be returned.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMemberFromAccessGroupOptions, "removeMemberFromAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMemberFromAccessGroupOptions, "removeMemberFromAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "members"}
	pathParameters := []string{*removeMemberFromAccessGroupOptions.AccessGroupID, *removeMemberFromAccessGroupOptions.IamID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeMemberFromAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "RemoveMemberFromAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if removeMemberFromAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*removeMemberFromAccessGroupOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// RemoveMembersFromAccessGroup : Delete members from an Access Group
// Remove multiple members from a group using this API. On a successful call, this API will always return 207. It is the
// caller's responsibility to iterate across the body to determine successful deletion of each member. This API request
// payload can delete up to 50 members per call.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions) (result *DeleteGroupBulkMembersResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMembersFromAccessGroupOptions, "removeMembersFromAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMembersFromAccessGroupOptions, "removeMembersFromAccessGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "members/delete"}
	pathParameters := []string{*removeMembersFromAccessGroupOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeMembersFromAccessGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "RemoveMembersFromAccessGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if removeMembersFromAccessGroupOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*removeMembersFromAccessGroupOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if removeMembersFromAccessGroupOptions.Members != nil {
		body["members"] = removeMembersFromAccessGroupOptions.Members
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalDeleteGroupBulkMembersResponse(m)
		response.Result = result
	}

	return
}

// RemoveMemberFromAllAccessGroups : Delete member from all Access Groups
// This API removes a given member from every group they are a member of within the specified account. By using one
// operation, you can revoke one member's access to all Access Groups in the account. If a partial failure occurs on
// deletion, the response will be shown in the body.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) (result *DeleteFromAllGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMemberFromAllAccessGroupsOptions, "removeMemberFromAllAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMemberFromAllAccessGroupsOptions, "removeMemberFromAllAccessGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups/_allgroups/members"}
	pathParameters := []string{*removeMemberFromAllAccessGroupsOptions.IamID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeMemberFromAllAccessGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "RemoveMemberFromAllAccessGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if removeMemberFromAllAccessGroupsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*removeMemberFromAllAccessGroupsOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*removeMemberFromAllAccessGroupsOptions.AccountID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalDeleteFromAllGroupsResponse(m)
		response.Result = result
	}

	return
}

// AddMemberToMultipleAccessGroups : Add member to multiple Access Groups
// This API will add a member to multiple Access Groups in an account. The limit of how many groups that can be in the
// request is 50. The response is a list of results that show if adding the member to each group was successful or not.
func (iamAccessGroups *IamAccessGroupsV2) AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) (result *AddMembershipMultipleGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addMemberToMultipleAccessGroupsOptions, "addMemberToMultipleAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addMemberToMultipleAccessGroupsOptions, "addMemberToMultipleAccessGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups/_allgroups/members"}
	pathParameters := []string{*addMemberToMultipleAccessGroupsOptions.IamID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addMemberToMultipleAccessGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "AddMemberToMultipleAccessGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addMemberToMultipleAccessGroupsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addMemberToMultipleAccessGroupsOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*addMemberToMultipleAccessGroupsOptions.AccountID))

	body := make(map[string]interface{})
	if addMemberToMultipleAccessGroupsOptions.Type != nil {
		body["type"] = addMemberToMultipleAccessGroupsOptions.Type
	}
	if addMemberToMultipleAccessGroupsOptions.Groups != nil {
		body["groups"] = addMemberToMultipleAccessGroupsOptions.Groups
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAddMembershipMultipleGroupsResponse(m)
		response.Result = result
	}

	return
}

// AddAccessGroupRule : Create rule for an Access Group
// Rules can be used to dynamically add users to an Access Group. If a user's SAML assertions match the rule's
// conditions during login, the user will be dynamically added to the group. The duration of the user's access to the
// group is determined by the `expiration` field. After access expires, the user will need to log in again to regain
// access. Note that the condition's value field must be a stringified JSON value. [Consult this documentation for
// further explanation of dynamic rules.](/docs/iam/accessgroup_rules.html#rules).
func (iamAccessGroups *IamAccessGroupsV2) AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addAccessGroupRuleOptions, "addAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addAccessGroupRuleOptions, "addAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "rules"}
	pathParameters := []string{*addAccessGroupRuleOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range addAccessGroupRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "AddAccessGroupRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addAccessGroupRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*addAccessGroupRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if addAccessGroupRuleOptions.Expiration != nil {
		body["expiration"] = addAccessGroupRuleOptions.Expiration
	}
	if addAccessGroupRuleOptions.RealmName != nil {
		body["realm_name"] = addAccessGroupRuleOptions.RealmName
	}
	if addAccessGroupRuleOptions.Conditions != nil {
		body["conditions"] = addAccessGroupRuleOptions.Conditions
	}
	if addAccessGroupRuleOptions.Name != nil {
		body["name"] = addAccessGroupRuleOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalRule(m)
		response.Result = result
	}

	return
}

// ListAccessGroupRules : List Access Group rules
// This API lists all rules in a given Access Group. Because only a few rules are created on each group, there is no
// pagination or sorting support on this API.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions) (result *RulesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupRulesOptions, "listAccessGroupRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupRulesOptions, "listAccessGroupRulesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "rules"}
	pathParameters := []string{*listAccessGroupRulesOptions.AccessGroupID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAccessGroupRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListAccessGroupRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAccessGroupRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAccessGroupRulesOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalRulesList(m)
		response.Result = result
	}

	return
}

// GetAccessGroupRule : Get an Access Group rule
// Retrieve a rule from an Access Group. A revision number is returned in the `Etag` header, which is needed when
// updating the rule.
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccessGroupRuleOptions, "getAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccessGroupRuleOptions, "getAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "rules"}
	pathParameters := []string{*getAccessGroupRuleOptions.AccessGroupID, *getAccessGroupRuleOptions.RuleID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccessGroupRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetAccessGroupRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAccessGroupRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAccessGroupRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalRule(m)
		response.Result = result
	}

	return
}

// ReplaceAccessGroupRule : Replace an Access Group rule
// Update the body of an existing rule using this API. An `If-Match` header must be populated with the rule's most
// recent revision number (which can be acquired in the `Get an Access Group rule` API).
func (iamAccessGroups *IamAccessGroupsV2) ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceAccessGroupRuleOptions, "replaceAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceAccessGroupRuleOptions, "replaceAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "rules"}
	pathParameters := []string{*replaceAccessGroupRuleOptions.AccessGroupID, *replaceAccessGroupRuleOptions.RuleID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceAccessGroupRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ReplaceAccessGroupRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceAccessGroupRuleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*replaceAccessGroupRuleOptions.IfMatch))
	}
	if replaceAccessGroupRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceAccessGroupRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceAccessGroupRuleOptions.Expiration != nil {
		body["expiration"] = replaceAccessGroupRuleOptions.Expiration
	}
	if replaceAccessGroupRuleOptions.RealmName != nil {
		body["realm_name"] = replaceAccessGroupRuleOptions.RealmName
	}
	if replaceAccessGroupRuleOptions.Conditions != nil {
		body["conditions"] = replaceAccessGroupRuleOptions.Conditions
	}
	if replaceAccessGroupRuleOptions.Name != nil {
		body["name"] = replaceAccessGroupRuleOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalRule(m)
		response.Result = result
	}

	return
}

// RemoveAccessGroupRule : Delete an Access Group rule
// Remove one rule from a group using this API. If the operation is successful, only a `204 - No Content` response with
// no body is returned. However, if any error occurs, the standard error format will be returned.
func (iamAccessGroups *IamAccessGroupsV2) RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeAccessGroupRuleOptions, "removeAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeAccessGroupRuleOptions, "removeAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"groups", "rules"}
	pathParameters := []string{*removeAccessGroupRuleOptions.AccessGroupID, *removeAccessGroupRuleOptions.RuleID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(iamAccessGroups.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeAccessGroupRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "RemoveAccessGroupRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if removeAccessGroupRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*removeAccessGroupRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// AccountSettings : The Access Groups settings for a specific account.
type AccountSettings struct {
	// The account id of the settings being shown.
	AccountID *string `json:"account_id,omitempty"`

	// The timestamp the settings were last edited at.
	LastModifiedAt *string `json:"last_modified_at,omitempty"`

	// The `iam_id` of the entity that last modified the settings.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// This flag controls the public access feature within the account. It is set to true by default. Note: When this flag
	// is set to false, all policies within the account attached to the Public Access group will be deleted.
	PublicAccessEnabled *bool `json:"public_access_enabled,omitempty"`
}


// UnmarshalAccountSettings constructs an instance of AccountSettings from the specified map.
func UnmarshalAccountSettings(m map[string]interface{}) (result *AccountSettings, err error) {
	obj := new(AccountSettings)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.LastModifiedAt, err = core.UnmarshalString(m, "last_modified_at")
	if err != nil {
		return
	}
	obj.LastModifiedByID, err = core.UnmarshalString(m, "last_modified_by_id")
	if err != nil {
		return
	}
	obj.PublicAccessEnabled, err = core.UnmarshalBool(m, "public_access_enabled")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAccountSettingsSlice unmarshals a slice of AccountSettings instances from the specified list of maps.
func UnmarshalAccountSettingsSlice(s []interface{}) (slice []AccountSettings, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AccountSettings'")
			return
		}
		obj, e := UnmarshalAccountSettings(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAccountSettingsAsProperty unmarshals an instance of AccountSettings that is stored as a property
// within the specified map.
func UnmarshalAccountSettingsAsProperty(m map[string]interface{}, propertyName string) (result *AccountSettings, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AccountSettings'", propertyName)
			return
		}
		result, err = UnmarshalAccountSettings(objMap)
	}
	return
}

// UnmarshalAccountSettingsSliceAsProperty unmarshals a slice of AccountSettings instances that are stored as a property
// within the specified map.
func UnmarshalAccountSettingsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AccountSettings, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AccountSettings'", propertyName)
			return
		}
		slice, err = UnmarshalAccountSettingsSlice(vSlice)
	}
	return
}

// AddAccessGroupRuleOptions : The AddAccessGroupRule options.
type AddAccessGroupRuleOptions struct {
	// The id of the group that the rule will be added to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The number of hours that the rule lives for (Must be between 1 and 24).
	Expiration *int64 `json:"expiration" validate:"required"`

	// The url of the identity provider.
	RealmName *string `json:"realm_name" validate:"required"`

	// A list of conditions the rule must satisfy.
	Conditions []RuleConditions `json:"conditions" validate:"required"`

	// The name of the rule.
	Name *string `json:"name,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddAccessGroupRuleOptions : Instantiate AddAccessGroupRuleOptions
func (*IamAccessGroupsV2) NewAddAccessGroupRuleOptions(accessGroupID string, expiration int64, realmName string, conditions []RuleConditions) *AddAccessGroupRuleOptions {
	return &AddAccessGroupRuleOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		Expiration: core.Int64Ptr(expiration),
		RealmName: core.StringPtr(realmName),
		Conditions: conditions,
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *AddAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *AddAccessGroupRuleOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetExpiration : Allow user to set Expiration
func (options *AddAccessGroupRuleOptions) SetExpiration(expiration int64) *AddAccessGroupRuleOptions {
	options.Expiration = core.Int64Ptr(expiration)
	return options
}

// SetRealmName : Allow user to set RealmName
func (options *AddAccessGroupRuleOptions) SetRealmName(realmName string) *AddAccessGroupRuleOptions {
	options.RealmName = core.StringPtr(realmName)
	return options
}

// SetConditions : Allow user to set Conditions
func (options *AddAccessGroupRuleOptions) SetConditions(conditions []RuleConditions) *AddAccessGroupRuleOptions {
	options.Conditions = conditions
	return options
}

// SetName : Allow user to set Name
func (options *AddAccessGroupRuleOptions) SetName(name string) *AddAccessGroupRuleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *AddAccessGroupRuleOptions) SetTransactionID(transactionID string) *AddAccessGroupRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddAccessGroupRuleOptions) SetHeaders(param map[string]string) *AddAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// AddGroupMembersRequestMembersItem : AddGroupMembersRequestMembersItem struct
type AddGroupMembersRequestMembersItem struct {
	// The IBMid or Service Id of the member.
	IamID *string `json:"iam_id" validate:"required"`

	// The type of the member, must be either "user" or "service".
	Type *string `json:"type" validate:"required"`
}


// NewAddGroupMembersRequestMembersItem : Instantiate AddGroupMembersRequestMembersItem (Generic Model Constructor)
func (*IamAccessGroupsV2) NewAddGroupMembersRequestMembersItem(iamID string, typeVar string) (model *AddGroupMembersRequestMembersItem, err error) {
	model = &AddGroupMembersRequestMembersItem{
		IamID: core.StringPtr(iamID),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAddGroupMembersRequestMembersItem constructs an instance of AddGroupMembersRequestMembersItem from the specified map.
func UnmarshalAddGroupMembersRequestMembersItem(m map[string]interface{}) (result *AddGroupMembersRequestMembersItem, err error) {
	obj := new(AddGroupMembersRequestMembersItem)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAddGroupMembersRequestMembersItemSlice unmarshals a slice of AddGroupMembersRequestMembersItem instances from the specified list of maps.
func UnmarshalAddGroupMembersRequestMembersItemSlice(s []interface{}) (slice []AddGroupMembersRequestMembersItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AddGroupMembersRequestMembersItem'")
			return
		}
		obj, e := UnmarshalAddGroupMembersRequestMembersItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAddGroupMembersRequestMembersItemAsProperty unmarshals an instance of AddGroupMembersRequestMembersItem that is stored as a property
// within the specified map.
func UnmarshalAddGroupMembersRequestMembersItemAsProperty(m map[string]interface{}, propertyName string) (result *AddGroupMembersRequestMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AddGroupMembersRequestMembersItem'", propertyName)
			return
		}
		result, err = UnmarshalAddGroupMembersRequestMembersItem(objMap)
	}
	return
}

// UnmarshalAddGroupMembersRequestMembersItemSliceAsProperty unmarshals a slice of AddGroupMembersRequestMembersItem instances that are stored as a property
// within the specified map.
func UnmarshalAddGroupMembersRequestMembersItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AddGroupMembersRequestMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AddGroupMembersRequestMembersItem'", propertyName)
			return
		}
		slice, err = UnmarshalAddGroupMembersRequestMembersItemSlice(vSlice)
	}
	return
}

// AddGroupMembersResponse : The members added to an access group.
type AddGroupMembersResponse struct {
	// The members added to an access group.
	Members []AddGroupMembersResponseMembersItem `json:"members,omitempty"`
}


// UnmarshalAddGroupMembersResponse constructs an instance of AddGroupMembersResponse from the specified map.
func UnmarshalAddGroupMembersResponse(m map[string]interface{}) (result *AddGroupMembersResponse, err error) {
	obj := new(AddGroupMembersResponse)
	obj.Members, err = UnmarshalAddGroupMembersResponseMembersItemSliceAsProperty(m, "members")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAddGroupMembersResponseSlice unmarshals a slice of AddGroupMembersResponse instances from the specified list of maps.
func UnmarshalAddGroupMembersResponseSlice(s []interface{}) (slice []AddGroupMembersResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AddGroupMembersResponse'")
			return
		}
		obj, e := UnmarshalAddGroupMembersResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAddGroupMembersResponseAsProperty unmarshals an instance of AddGroupMembersResponse that is stored as a property
// within the specified map.
func UnmarshalAddGroupMembersResponseAsProperty(m map[string]interface{}, propertyName string) (result *AddGroupMembersResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AddGroupMembersResponse'", propertyName)
			return
		}
		result, err = UnmarshalAddGroupMembersResponse(objMap)
	}
	return
}

// UnmarshalAddGroupMembersResponseSliceAsProperty unmarshals a slice of AddGroupMembersResponse instances that are stored as a property
// within the specified map.
func UnmarshalAddGroupMembersResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AddGroupMembersResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AddGroupMembersResponse'", propertyName)
			return
		}
		slice, err = UnmarshalAddGroupMembersResponseSlice(vSlice)
	}
	return
}

// AddGroupMembersResponseMembersItem : AddGroupMembersResponseMembersItem struct
type AddGroupMembersResponseMembersItem struct {
	// The IBMid or Service Id of the member.
	IamID *string `json:"iam_id,omitempty"`

	// The member type - either `user` or `service`.
	Type *string `json:"type,omitempty"`

	// The timestamp the membership was created at.
	CreatedAt *string `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the membership.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The outcome of the operation on this `iam_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// A list of errors that occurred when trying to add members to a group.
	Errors []Error `json:"errors,omitempty"`
}


// UnmarshalAddGroupMembersResponseMembersItem constructs an instance of AddGroupMembersResponseMembersItem from the specified map.
func UnmarshalAddGroupMembersResponseMembersItem(m map[string]interface{}) (result *AddGroupMembersResponseMembersItem, err error) {
	obj := new(AddGroupMembersResponseMembersItem)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.CreatedByID, err = core.UnmarshalString(m, "created_by_id")
	if err != nil {
		return
	}
	obj.StatusCode, err = core.UnmarshalInt64(m, "status_code")
	if err != nil {
		return
	}
	obj.Trace, err = core.UnmarshalString(m, "trace")
	if err != nil {
		return
	}
	obj.Errors, err = UnmarshalErrorSliceAsProperty(m, "errors")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAddGroupMembersResponseMembersItemSlice unmarshals a slice of AddGroupMembersResponseMembersItem instances from the specified list of maps.
func UnmarshalAddGroupMembersResponseMembersItemSlice(s []interface{}) (slice []AddGroupMembersResponseMembersItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AddGroupMembersResponseMembersItem'")
			return
		}
		obj, e := UnmarshalAddGroupMembersResponseMembersItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAddGroupMembersResponseMembersItemAsProperty unmarshals an instance of AddGroupMembersResponseMembersItem that is stored as a property
// within the specified map.
func UnmarshalAddGroupMembersResponseMembersItemAsProperty(m map[string]interface{}, propertyName string) (result *AddGroupMembersResponseMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AddGroupMembersResponseMembersItem'", propertyName)
			return
		}
		result, err = UnmarshalAddGroupMembersResponseMembersItem(objMap)
	}
	return
}

// UnmarshalAddGroupMembersResponseMembersItemSliceAsProperty unmarshals a slice of AddGroupMembersResponseMembersItem instances that are stored as a property
// within the specified map.
func UnmarshalAddGroupMembersResponseMembersItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AddGroupMembersResponseMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AddGroupMembersResponseMembersItem'", propertyName)
			return
		}
		slice, err = UnmarshalAddGroupMembersResponseMembersItemSlice(vSlice)
	}
	return
}

// AddMemberToMultipleAccessGroupsOptions : The AddMemberToMultipleAccessGroups options.
type AddMemberToMultipleAccessGroupsOptions struct {
	// IBM Cloud account id of the groups that the member will be added to.
	AccountID *string `json:"account_id" validate:"required"`

	// The iam_id to be added to the groups.
	IamID *string `json:"iam_id" validate:"required"`

	// The type of the member, must be either "user" or "service".
	Type *string `json:"type,omitempty"`

	// The ids of the access groups a given member is to be added to.
	Groups []string `json:"groups,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddMemberToMultipleAccessGroupsOptions : Instantiate AddMemberToMultipleAccessGroupsOptions
func (*IamAccessGroupsV2) NewAddMemberToMultipleAccessGroupsOptions(accountID string, iamID string) *AddMemberToMultipleAccessGroupsOptions {
	return &AddMemberToMultipleAccessGroupsOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *AddMemberToMultipleAccessGroupsOptions) SetAccountID(accountID string) *AddMemberToMultipleAccessGroupsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *AddMemberToMultipleAccessGroupsOptions) SetIamID(iamID string) *AddMemberToMultipleAccessGroupsOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetType : Allow user to set Type
func (options *AddMemberToMultipleAccessGroupsOptions) SetType(typeVar string) *AddMemberToMultipleAccessGroupsOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetGroups : Allow user to set Groups
func (options *AddMemberToMultipleAccessGroupsOptions) SetGroups(groups []string) *AddMemberToMultipleAccessGroupsOptions {
	options.Groups = groups
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *AddMemberToMultipleAccessGroupsOptions) SetTransactionID(transactionID string) *AddMemberToMultipleAccessGroupsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddMemberToMultipleAccessGroupsOptions) SetHeaders(param map[string]string) *AddMemberToMultipleAccessGroupsOptions {
	options.Headers = param
	return options
}

// AddMembersToAccessGroupOptions : The AddMembersToAccessGroup options.
type AddMembersToAccessGroupOptions struct {
	// The Access Group to add the members to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// An array of member objects to add to an access group.
	Members []AddGroupMembersRequestMembersItem `json:"members,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddMembersToAccessGroupOptions : Instantiate AddMembersToAccessGroupOptions
func (*IamAccessGroupsV2) NewAddMembersToAccessGroupOptions(accessGroupID string) *AddMembersToAccessGroupOptions {
	return &AddMembersToAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *AddMembersToAccessGroupOptions) SetAccessGroupID(accessGroupID string) *AddMembersToAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetMembers : Allow user to set Members
func (options *AddMembersToAccessGroupOptions) SetMembers(members []AddGroupMembersRequestMembersItem) *AddMembersToAccessGroupOptions {
	options.Members = members
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *AddMembersToAccessGroupOptions) SetTransactionID(transactionID string) *AddMembersToAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AddMembersToAccessGroupOptions) SetHeaders(param map[string]string) *AddMembersToAccessGroupOptions {
	options.Headers = param
	return options
}

// AddMembershipMultipleGroupsResponse : The response from the add member to multiple access groups request.
type AddMembershipMultipleGroupsResponse struct {
	// The iam_id of a member.
	IamID *string `json:"iam_id,omitempty"`

	// The list of access groups a member was added to.
	Groups []AddMembershipMultipleGroupsResponseGroupsItem `json:"groups,omitempty"`
}


// UnmarshalAddMembershipMultipleGroupsResponse constructs an instance of AddMembershipMultipleGroupsResponse from the specified map.
func UnmarshalAddMembershipMultipleGroupsResponse(m map[string]interface{}) (result *AddMembershipMultipleGroupsResponse, err error) {
	obj := new(AddMembershipMultipleGroupsResponse)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Groups, err = UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSliceAsProperty(m, "groups")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseSlice unmarshals a slice of AddMembershipMultipleGroupsResponse instances from the specified list of maps.
func UnmarshalAddMembershipMultipleGroupsResponseSlice(s []interface{}) (slice []AddMembershipMultipleGroupsResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AddMembershipMultipleGroupsResponse'")
			return
		}
		obj, e := UnmarshalAddMembershipMultipleGroupsResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseAsProperty unmarshals an instance of AddMembershipMultipleGroupsResponse that is stored as a property
// within the specified map.
func UnmarshalAddMembershipMultipleGroupsResponseAsProperty(m map[string]interface{}, propertyName string) (result *AddMembershipMultipleGroupsResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AddMembershipMultipleGroupsResponse'", propertyName)
			return
		}
		result, err = UnmarshalAddMembershipMultipleGroupsResponse(objMap)
	}
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseSliceAsProperty unmarshals a slice of AddMembershipMultipleGroupsResponse instances that are stored as a property
// within the specified map.
func UnmarshalAddMembershipMultipleGroupsResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AddMembershipMultipleGroupsResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AddMembershipMultipleGroupsResponse'", propertyName)
			return
		}
		slice, err = UnmarshalAddMembershipMultipleGroupsResponseSlice(vSlice)
	}
	return
}

// AddMembershipMultipleGroupsResponseGroupsItem : AddMembershipMultipleGroupsResponseGroupsItem struct
type AddMembershipMultipleGroupsResponseGroupsItem struct {
	// The Access Group that the member is to be added to.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The outcome of the add membership operation on this `access_group_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// List of errors encountered when adding member to access group.
	Errors []Error `json:"errors,omitempty"`
}


// UnmarshalAddMembershipMultipleGroupsResponseGroupsItem constructs an instance of AddMembershipMultipleGroupsResponseGroupsItem from the specified map.
func UnmarshalAddMembershipMultipleGroupsResponseGroupsItem(m map[string]interface{}) (result *AddMembershipMultipleGroupsResponseGroupsItem, err error) {
	obj := new(AddMembershipMultipleGroupsResponseGroupsItem)
	obj.AccessGroupID, err = core.UnmarshalString(m, "access_group_id")
	if err != nil {
		return
	}
	obj.StatusCode, err = core.UnmarshalInt64(m, "status_code")
	if err != nil {
		return
	}
	obj.Trace, err = core.UnmarshalString(m, "trace")
	if err != nil {
		return
	}
	obj.Errors, err = UnmarshalErrorSliceAsProperty(m, "errors")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSlice unmarshals a slice of AddMembershipMultipleGroupsResponseGroupsItem instances from the specified list of maps.
func UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSlice(s []interface{}) (slice []AddMembershipMultipleGroupsResponseGroupsItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AddMembershipMultipleGroupsResponseGroupsItem'")
			return
		}
		obj, e := UnmarshalAddMembershipMultipleGroupsResponseGroupsItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseGroupsItemAsProperty unmarshals an instance of AddMembershipMultipleGroupsResponseGroupsItem that is stored as a property
// within the specified map.
func UnmarshalAddMembershipMultipleGroupsResponseGroupsItemAsProperty(m map[string]interface{}, propertyName string) (result *AddMembershipMultipleGroupsResponseGroupsItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AddMembershipMultipleGroupsResponseGroupsItem'", propertyName)
			return
		}
		result, err = UnmarshalAddMembershipMultipleGroupsResponseGroupsItem(objMap)
	}
	return
}

// UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSliceAsProperty unmarshals a slice of AddMembershipMultipleGroupsResponseGroupsItem instances that are stored as a property
// within the specified map.
func UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AddMembershipMultipleGroupsResponseGroupsItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AddMembershipMultipleGroupsResponseGroupsItem'", propertyName)
			return
		}
		slice, err = UnmarshalAddMembershipMultipleGroupsResponseGroupsItemSlice(vSlice)
	}
	return
}

// CreateAccessGroupOptions : The CreateAccessGroup options.
type CreateAccessGroupOptions struct {
	// IBM Cloud account id under which the group is created.
	AccountID *string `json:"account_id" validate:"required"`

	// Assign the specified name to the Access Group. This field has a limit of 100 characters.
	Name *string `json:"name" validate:"required"`

	// Assign a description for the Access Group. This field has a limit of 250 characters.
	Description *string `json:"description,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAccessGroupOptions : Instantiate CreateAccessGroupOptions
func (*IamAccessGroupsV2) NewCreateAccessGroupOptions(accountID string, name string) *CreateAccessGroupOptions {
	return &CreateAccessGroupOptions{
		AccountID: core.StringPtr(accountID),
		Name: core.StringPtr(name),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateAccessGroupOptions) SetAccountID(accountID string) *CreateAccessGroupOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateAccessGroupOptions) SetName(name string) *CreateAccessGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateAccessGroupOptions) SetDescription(description string) *CreateAccessGroupOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateAccessGroupOptions) SetTransactionID(transactionID string) *CreateAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccessGroupOptions) SetHeaders(param map[string]string) *CreateAccessGroupOptions {
	options.Headers = param
	return options
}

// DeleteAccessGroupOptions : The DeleteAccessGroup options.
type DeleteAccessGroupOptions struct {
	// The Access group to delete.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// If force is true, delete the group as well as its associated members and rules.
	Force *bool `json:"force,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAccessGroupOptions : Instantiate DeleteAccessGroupOptions
func (*IamAccessGroupsV2) NewDeleteAccessGroupOptions(accessGroupID string) *DeleteAccessGroupOptions {
	return &DeleteAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *DeleteAccessGroupOptions) SetAccessGroupID(accessGroupID string) *DeleteAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteAccessGroupOptions) SetTransactionID(transactionID string) *DeleteAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetForce : Allow user to set Force
func (options *DeleteAccessGroupOptions) SetForce(force bool) *DeleteAccessGroupOptions {
	options.Force = core.BoolPtr(force)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccessGroupOptions) SetHeaders(param map[string]string) *DeleteAccessGroupOptions {
	options.Headers = param
	return options
}

// DeleteFromAllGroupsResponse : The response from the delete member from access groups request.
type DeleteFromAllGroupsResponse struct {
	// The `iam_id` of the member to removed from groups.
	IamID *string `json:"iam_id,omitempty"`

	// The groups the member was removed from.
	Groups []DeleteFromAllGroupsResponseGroupsItem `json:"groups,omitempty"`
}


// UnmarshalDeleteFromAllGroupsResponse constructs an instance of DeleteFromAllGroupsResponse from the specified map.
func UnmarshalDeleteFromAllGroupsResponse(m map[string]interface{}) (result *DeleteFromAllGroupsResponse, err error) {
	obj := new(DeleteFromAllGroupsResponse)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Groups, err = UnmarshalDeleteFromAllGroupsResponseGroupsItemSliceAsProperty(m, "groups")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeleteFromAllGroupsResponseSlice unmarshals a slice of DeleteFromAllGroupsResponse instances from the specified list of maps.
func UnmarshalDeleteFromAllGroupsResponseSlice(s []interface{}) (slice []DeleteFromAllGroupsResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeleteFromAllGroupsResponse'")
			return
		}
		obj, e := UnmarshalDeleteFromAllGroupsResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeleteFromAllGroupsResponseAsProperty unmarshals an instance of DeleteFromAllGroupsResponse that is stored as a property
// within the specified map.
func UnmarshalDeleteFromAllGroupsResponseAsProperty(m map[string]interface{}, propertyName string) (result *DeleteFromAllGroupsResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeleteFromAllGroupsResponse'", propertyName)
			return
		}
		result, err = UnmarshalDeleteFromAllGroupsResponse(objMap)
	}
	return
}

// UnmarshalDeleteFromAllGroupsResponseSliceAsProperty unmarshals a slice of DeleteFromAllGroupsResponse instances that are stored as a property
// within the specified map.
func UnmarshalDeleteFromAllGroupsResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeleteFromAllGroupsResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeleteFromAllGroupsResponse'", propertyName)
			return
		}
		slice, err = UnmarshalDeleteFromAllGroupsResponseSlice(vSlice)
	}
	return
}

// DeleteFromAllGroupsResponseGroupsItem : DeleteFromAllGroupsResponseGroupsItem struct
type DeleteFromAllGroupsResponseGroupsItem struct {
	// The Access Group that the member is to be deleted from.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The outcome of the delete operation on this `access_group_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// A list of errors that occurred when trying to remove a member from groups.
	Errors []Error `json:"errors,omitempty"`
}


// UnmarshalDeleteFromAllGroupsResponseGroupsItem constructs an instance of DeleteFromAllGroupsResponseGroupsItem from the specified map.
func UnmarshalDeleteFromAllGroupsResponseGroupsItem(m map[string]interface{}) (result *DeleteFromAllGroupsResponseGroupsItem, err error) {
	obj := new(DeleteFromAllGroupsResponseGroupsItem)
	obj.AccessGroupID, err = core.UnmarshalString(m, "access_group_id")
	if err != nil {
		return
	}
	obj.StatusCode, err = core.UnmarshalInt64(m, "status_code")
	if err != nil {
		return
	}
	obj.Trace, err = core.UnmarshalString(m, "trace")
	if err != nil {
		return
	}
	obj.Errors, err = UnmarshalErrorSliceAsProperty(m, "errors")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeleteFromAllGroupsResponseGroupsItemSlice unmarshals a slice of DeleteFromAllGroupsResponseGroupsItem instances from the specified list of maps.
func UnmarshalDeleteFromAllGroupsResponseGroupsItemSlice(s []interface{}) (slice []DeleteFromAllGroupsResponseGroupsItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeleteFromAllGroupsResponseGroupsItem'")
			return
		}
		obj, e := UnmarshalDeleteFromAllGroupsResponseGroupsItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeleteFromAllGroupsResponseGroupsItemAsProperty unmarshals an instance of DeleteFromAllGroupsResponseGroupsItem that is stored as a property
// within the specified map.
func UnmarshalDeleteFromAllGroupsResponseGroupsItemAsProperty(m map[string]interface{}, propertyName string) (result *DeleteFromAllGroupsResponseGroupsItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeleteFromAllGroupsResponseGroupsItem'", propertyName)
			return
		}
		result, err = UnmarshalDeleteFromAllGroupsResponseGroupsItem(objMap)
	}
	return
}

// UnmarshalDeleteFromAllGroupsResponseGroupsItemSliceAsProperty unmarshals a slice of DeleteFromAllGroupsResponseGroupsItem instances that are stored as a property
// within the specified map.
func UnmarshalDeleteFromAllGroupsResponseGroupsItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeleteFromAllGroupsResponseGroupsItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeleteFromAllGroupsResponseGroupsItem'", propertyName)
			return
		}
		slice, err = UnmarshalDeleteFromAllGroupsResponseGroupsItemSlice(vSlice)
	}
	return
}

// DeleteGroupBulkMembersResponse : The access group id and the members removed from it.
type DeleteGroupBulkMembersResponse struct {
	// The access group id.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The `iam_id`s removed from the access group.
	Members []DeleteGroupBulkMembersResponseMembersItem `json:"members,omitempty"`
}


// UnmarshalDeleteGroupBulkMembersResponse constructs an instance of DeleteGroupBulkMembersResponse from the specified map.
func UnmarshalDeleteGroupBulkMembersResponse(m map[string]interface{}) (result *DeleteGroupBulkMembersResponse, err error) {
	obj := new(DeleteGroupBulkMembersResponse)
	obj.AccessGroupID, err = core.UnmarshalString(m, "access_group_id")
	if err != nil {
		return
	}
	obj.Members, err = UnmarshalDeleteGroupBulkMembersResponseMembersItemSliceAsProperty(m, "members")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeleteGroupBulkMembersResponseSlice unmarshals a slice of DeleteGroupBulkMembersResponse instances from the specified list of maps.
func UnmarshalDeleteGroupBulkMembersResponseSlice(s []interface{}) (slice []DeleteGroupBulkMembersResponse, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeleteGroupBulkMembersResponse'")
			return
		}
		obj, e := UnmarshalDeleteGroupBulkMembersResponse(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeleteGroupBulkMembersResponseAsProperty unmarshals an instance of DeleteGroupBulkMembersResponse that is stored as a property
// within the specified map.
func UnmarshalDeleteGroupBulkMembersResponseAsProperty(m map[string]interface{}, propertyName string) (result *DeleteGroupBulkMembersResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeleteGroupBulkMembersResponse'", propertyName)
			return
		}
		result, err = UnmarshalDeleteGroupBulkMembersResponse(objMap)
	}
	return
}

// UnmarshalDeleteGroupBulkMembersResponseSliceAsProperty unmarshals a slice of DeleteGroupBulkMembersResponse instances that are stored as a property
// within the specified map.
func UnmarshalDeleteGroupBulkMembersResponseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeleteGroupBulkMembersResponse, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeleteGroupBulkMembersResponse'", propertyName)
			return
		}
		slice, err = UnmarshalDeleteGroupBulkMembersResponseSlice(vSlice)
	}
	return
}

// DeleteGroupBulkMembersResponseMembersItem : DeleteGroupBulkMembersResponseMembersItem struct
type DeleteGroupBulkMembersResponseMembersItem struct {
	// The `iam_id` to be deleted.
	IamID *string `json:"iam_id,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// The outcome of the delete membership operation on this `access_group_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A list of errors that occurred when trying to remove a member from groups.
	Errors []Error `json:"errors,omitempty"`
}


// UnmarshalDeleteGroupBulkMembersResponseMembersItem constructs an instance of DeleteGroupBulkMembersResponseMembersItem from the specified map.
func UnmarshalDeleteGroupBulkMembersResponseMembersItem(m map[string]interface{}) (result *DeleteGroupBulkMembersResponseMembersItem, err error) {
	obj := new(DeleteGroupBulkMembersResponseMembersItem)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Trace, err = core.UnmarshalString(m, "trace")
	if err != nil {
		return
	}
	obj.StatusCode, err = core.UnmarshalInt64(m, "status_code")
	if err != nil {
		return
	}
	obj.Errors, err = UnmarshalErrorSliceAsProperty(m, "errors")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeleteGroupBulkMembersResponseMembersItemSlice unmarshals a slice of DeleteGroupBulkMembersResponseMembersItem instances from the specified list of maps.
func UnmarshalDeleteGroupBulkMembersResponseMembersItemSlice(s []interface{}) (slice []DeleteGroupBulkMembersResponseMembersItem, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeleteGroupBulkMembersResponseMembersItem'")
			return
		}
		obj, e := UnmarshalDeleteGroupBulkMembersResponseMembersItem(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeleteGroupBulkMembersResponseMembersItemAsProperty unmarshals an instance of DeleteGroupBulkMembersResponseMembersItem that is stored as a property
// within the specified map.
func UnmarshalDeleteGroupBulkMembersResponseMembersItemAsProperty(m map[string]interface{}, propertyName string) (result *DeleteGroupBulkMembersResponseMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeleteGroupBulkMembersResponseMembersItem'", propertyName)
			return
		}
		result, err = UnmarshalDeleteGroupBulkMembersResponseMembersItem(objMap)
	}
	return
}

// UnmarshalDeleteGroupBulkMembersResponseMembersItemSliceAsProperty unmarshals a slice of DeleteGroupBulkMembersResponseMembersItem instances that are stored as a property
// within the specified map.
func UnmarshalDeleteGroupBulkMembersResponseMembersItemSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeleteGroupBulkMembersResponseMembersItem, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeleteGroupBulkMembersResponseMembersItem'", propertyName)
			return
		}
		slice, err = UnmarshalDeleteGroupBulkMembersResponseMembersItemSlice(vSlice)
	}
	return
}

// Error : Error contains the code and message for an error returned to the user code is a string identifying the problem,
// examples "missing_field", "reserved_value" message is a string explaining the solution to the problem that was
// encountered.
type Error struct {
	// A human-readable error code represented by a snake case string.
	Code *string `json:"code,omitempty"`

	// A specific error message that details the issue or an action to take.
	Message *string `json:"message,omitempty"`
}


// UnmarshalError constructs an instance of Error from the specified map.
func UnmarshalError(m map[string]interface{}) (result *Error, err error) {
	obj := new(Error)
	obj.Code, err = core.UnmarshalString(m, "code")
	if err != nil {
		return
	}
	obj.Message, err = core.UnmarshalString(m, "message")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalErrorSlice unmarshals a slice of Error instances from the specified list of maps.
func UnmarshalErrorSlice(s []interface{}) (slice []Error, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Error'")
			return
		}
		obj, e := UnmarshalError(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalErrorAsProperty unmarshals an instance of Error that is stored as a property
// within the specified map.
func UnmarshalErrorAsProperty(m map[string]interface{}, propertyName string) (result *Error, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Error'", propertyName)
			return
		}
		result, err = UnmarshalError(objMap)
	}
	return
}

// UnmarshalErrorSliceAsProperty unmarshals a slice of Error instances that are stored as a property
// within the specified map.
func UnmarshalErrorSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Error, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Error'", propertyName)
			return
		}
		slice, err = UnmarshalErrorSlice(vSlice)
	}
	return
}

// GetAccessGroupOptions : The GetAccessGroup options.
type GetAccessGroupOptions struct {
	// The Access Group to get.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// If show_federated is true, the group will return an is_federated value that is set to true if rules exist for the
	// group.
	ShowFederated *bool `json:"show_federated,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccessGroupOptions : Instantiate GetAccessGroupOptions
func (*IamAccessGroupsV2) NewGetAccessGroupOptions(accessGroupID string) *GetAccessGroupOptions {
	return &GetAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *GetAccessGroupOptions) SetAccessGroupID(accessGroupID string) *GetAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetAccessGroupOptions) SetTransactionID(transactionID string) *GetAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetShowFederated : Allow user to set ShowFederated
func (options *GetAccessGroupOptions) SetShowFederated(showFederated bool) *GetAccessGroupOptions {
	options.ShowFederated = core.BoolPtr(showFederated)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccessGroupOptions) SetHeaders(param map[string]string) *GetAccessGroupOptions {
	options.Headers = param
	return options
}

// GetAccessGroupRuleOptions : The GetAccessGroupRule options.
type GetAccessGroupRuleOptions struct {
	// The group id that the rule is bound to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The rule to get.
	RuleID *string `json:"rule_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccessGroupRuleOptions : Instantiate GetAccessGroupRuleOptions
func (*IamAccessGroupsV2) NewGetAccessGroupRuleOptions(accessGroupID string, ruleID string) *GetAccessGroupRuleOptions {
	return &GetAccessGroupRuleOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		RuleID: core.StringPtr(ruleID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *GetAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *GetAccessGroupRuleOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetRuleID : Allow user to set RuleID
func (options *GetAccessGroupRuleOptions) SetRuleID(ruleID string) *GetAccessGroupRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetAccessGroupRuleOptions) SetTransactionID(transactionID string) *GetAccessGroupRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccessGroupRuleOptions) SetHeaders(param map[string]string) *GetAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// GetAccountSettingsOptions : The GetAccountSettings options.
type GetAccountSettingsOptions struct {
	// The account id of the settings being retrieved.
	AccountID *string `json:"account_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountSettingsOptions : Instantiate GetAccountSettingsOptions
func (*IamAccessGroupsV2) NewGetAccountSettingsOptions(accountID string) *GetAccountSettingsOptions {
	return &GetAccountSettingsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountSettingsOptions) SetAccountID(accountID string) *GetAccountSettingsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetAccountSettingsOptions) SetTransactionID(transactionID string) *GetAccountSettingsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountSettingsOptions) SetHeaders(param map[string]string) *GetAccountSettingsOptions {
	options.Headers = param
	return options
}

// Group : An IAM access group.
type Group struct {
	// The group's Access Group ID.
	ID *string `json:"id,omitempty"`

	// The group's name.
	Name *string `json:"name,omitempty"`

	// The group's description - if defined.
	Description *string `json:"description,omitempty"`

	// The account id where the group was created.
	AccountID *string `json:"account_id,omitempty"`

	// The timestamp the group was created at.
	CreatedAt *string `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the group.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The timestamp the group was last edited at.
	LastModifiedAt *string `json:"last_modified_at,omitempty"`

	// The `iam_id` of the entity that last modified the group name or description.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// A url to the given group resource.
	Href *string `json:"href,omitempty"`

	// This is set to true if rules exist for the group.
	IsFederated *bool `json:"is_federated,omitempty"`
}


// UnmarshalGroup constructs an instance of Group from the specified map.
func UnmarshalGroup(m map[string]interface{}) (result *Group, err error) {
	obj := new(Group)
	obj.ID, err = core.UnmarshalString(m, "id")
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
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.CreatedByID, err = core.UnmarshalString(m, "created_by_id")
	if err != nil {
		return
	}
	obj.LastModifiedAt, err = core.UnmarshalString(m, "last_modified_at")
	if err != nil {
		return
	}
	obj.LastModifiedByID, err = core.UnmarshalString(m, "last_modified_by_id")
	if err != nil {
		return
	}
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	obj.IsFederated, err = core.UnmarshalBool(m, "is_federated")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalGroupSlice unmarshals a slice of Group instances from the specified list of maps.
func UnmarshalGroupSlice(s []interface{}) (slice []Group, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Group'")
			return
		}
		obj, e := UnmarshalGroup(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalGroupAsProperty unmarshals an instance of Group that is stored as a property
// within the specified map.
func UnmarshalGroupAsProperty(m map[string]interface{}, propertyName string) (result *Group, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Group'", propertyName)
			return
		}
		result, err = UnmarshalGroup(objMap)
	}
	return
}

// UnmarshalGroupSliceAsProperty unmarshals a slice of Group instances that are stored as a property
// within the specified map.
func UnmarshalGroupSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Group, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Group'", propertyName)
			return
		}
		slice, err = UnmarshalGroupSlice(vSlice)
	}
	return
}

// GroupMembersList : The members of a group.
type GroupMembersList struct {
	// Limit on how many items can be returned.
	Limit *int64 `json:"limit,omitempty"`

	// The number of items to skip over in the result set.
	Offset *int64 `json:"offset,omitempty"`

	// The total number of items that match the query.
	TotalCount *int64 `json:"total_count,omitempty"`

	// A link object.
	First *HrefStruct `json:"first,omitempty"`

	// A link object.
	Previous *HrefStruct `json:"previous,omitempty"`

	// A link object.
	Next *HrefStruct `json:"next,omitempty"`

	// A link object.
	Last *HrefStruct `json:"last,omitempty"`

	// The members of an access group.
	Members []ListGroupMembersResponseMember `json:"members,omitempty"`
}


// UnmarshalGroupMembersList constructs an instance of GroupMembersList from the specified map.
func UnmarshalGroupMembersList(m map[string]interface{}) (result *GroupMembersList, err error) {
	obj := new(GroupMembersList)
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalHrefStructAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Previous, err = UnmarshalHrefStructAsProperty(m, "previous")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalHrefStructAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Last, err = UnmarshalHrefStructAsProperty(m, "last")
	if err != nil {
		return
	}
	obj.Members, err = UnmarshalListGroupMembersResponseMemberSliceAsProperty(m, "members")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalGroupMembersListSlice unmarshals a slice of GroupMembersList instances from the specified list of maps.
func UnmarshalGroupMembersListSlice(s []interface{}) (slice []GroupMembersList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'GroupMembersList'")
			return
		}
		obj, e := UnmarshalGroupMembersList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalGroupMembersListAsProperty unmarshals an instance of GroupMembersList that is stored as a property
// within the specified map.
func UnmarshalGroupMembersListAsProperty(m map[string]interface{}, propertyName string) (result *GroupMembersList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'GroupMembersList'", propertyName)
			return
		}
		result, err = UnmarshalGroupMembersList(objMap)
	}
	return
}

// UnmarshalGroupMembersListSliceAsProperty unmarshals a slice of GroupMembersList instances that are stored as a property
// within the specified map.
func UnmarshalGroupMembersListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []GroupMembersList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'GroupMembersList'", propertyName)
			return
		}
		slice, err = UnmarshalGroupMembersListSlice(vSlice)
	}
	return
}

// GroupsList : The list of access groups returned as part of a response.
type GroupsList struct {
	// Limit on how many items can be returned.
	Limit *int64 `json:"limit,omitempty"`

	// The number of items to skip over in the result set.
	Offset *int64 `json:"offset,omitempty"`

	// The total number of items that match the query.
	TotalCount *int64 `json:"total_count,omitempty"`

	// A link object.
	First *HrefStruct `json:"first,omitempty"`

	// A link object.
	Previous *HrefStruct `json:"previous,omitempty"`

	// A link object.
	Next *HrefStruct `json:"next,omitempty"`

	// A link object.
	Last *HrefStruct `json:"last,omitempty"`

	// An array of access groups.
	Groups []Group `json:"groups,omitempty"`
}


// UnmarshalGroupsList constructs an instance of GroupsList from the specified map.
func UnmarshalGroupsList(m map[string]interface{}) (result *GroupsList, err error) {
	obj := new(GroupsList)
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalInt64(m, "offset")
	if err != nil {
		return
	}
	obj.TotalCount, err = core.UnmarshalInt64(m, "total_count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalHrefStructAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Previous, err = UnmarshalHrefStructAsProperty(m, "previous")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalHrefStructAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Last, err = UnmarshalHrefStructAsProperty(m, "last")
	if err != nil {
		return
	}
	obj.Groups, err = UnmarshalGroupSliceAsProperty(m, "groups")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalGroupsListSlice unmarshals a slice of GroupsList instances from the specified list of maps.
func UnmarshalGroupsListSlice(s []interface{}) (slice []GroupsList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'GroupsList'")
			return
		}
		obj, e := UnmarshalGroupsList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalGroupsListAsProperty unmarshals an instance of GroupsList that is stored as a property
// within the specified map.
func UnmarshalGroupsListAsProperty(m map[string]interface{}, propertyName string) (result *GroupsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'GroupsList'", propertyName)
			return
		}
		result, err = UnmarshalGroupsList(objMap)
	}
	return
}

// UnmarshalGroupsListSliceAsProperty unmarshals a slice of GroupsList instances that are stored as a property
// within the specified map.
func UnmarshalGroupsListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []GroupsList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'GroupsList'", propertyName)
			return
		}
		slice, err = UnmarshalGroupsListSlice(vSlice)
	}
	return
}

// HrefStruct : A link object.
type HrefStruct struct {
	// A string containing the link’s URL.
	Href *string `json:"href,omitempty"`
}


// UnmarshalHrefStruct constructs an instance of HrefStruct from the specified map.
func UnmarshalHrefStruct(m map[string]interface{}) (result *HrefStruct, err error) {
	obj := new(HrefStruct)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalHrefStructSlice unmarshals a slice of HrefStruct instances from the specified list of maps.
func UnmarshalHrefStructSlice(s []interface{}) (slice []HrefStruct, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'HrefStruct'")
			return
		}
		obj, e := UnmarshalHrefStruct(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalHrefStructAsProperty unmarshals an instance of HrefStruct that is stored as a property
// within the specified map.
func UnmarshalHrefStructAsProperty(m map[string]interface{}, propertyName string) (result *HrefStruct, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'HrefStruct'", propertyName)
			return
		}
		result, err = UnmarshalHrefStruct(objMap)
	}
	return
}

// UnmarshalHrefStructSliceAsProperty unmarshals a slice of HrefStruct instances that are stored as a property
// within the specified map.
func UnmarshalHrefStructSliceAsProperty(m map[string]interface{}, propertyName string) (slice []HrefStruct, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'HrefStruct'", propertyName)
			return
		}
		slice, err = UnmarshalHrefStructSlice(vSlice)
	}
	return
}

// IsMemberOfAccessGroupOptions : The IsMemberOfAccessGroup options.
type IsMemberOfAccessGroupOptions struct {
	// The access_group_id to check for membership in.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The iam_id to look for within the group.
	IamID *string `json:"iam_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIsMemberOfAccessGroupOptions : Instantiate IsMemberOfAccessGroupOptions
func (*IamAccessGroupsV2) NewIsMemberOfAccessGroupOptions(accessGroupID string, iamID string) *IsMemberOfAccessGroupOptions {
	return &IsMemberOfAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *IsMemberOfAccessGroupOptions) SetAccessGroupID(accessGroupID string) *IsMemberOfAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *IsMemberOfAccessGroupOptions) SetIamID(iamID string) *IsMemberOfAccessGroupOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *IsMemberOfAccessGroupOptions) SetTransactionID(transactionID string) *IsMemberOfAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *IsMemberOfAccessGroupOptions) SetHeaders(param map[string]string) *IsMemberOfAccessGroupOptions {
	options.Headers = param
	return options
}

// ListAccessGroupMembersOptions : The ListAccessGroupMembers options.
type ListAccessGroupMembersOptions struct {
	// The access_group_id to list members of.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Return up to this limit of results where limit is between 0 and 100.
	Limit *float64 `json:"limit,omitempty"`

	// Offset the results using this query parameter.
	Offset *float64 `json:"offset,omitempty"`

	// Filter the results by member type.
	Type *string `json:"type,omitempty"`

	// Return user's email and name for each user id or the name for each service id.
	Verbose *bool `json:"verbose,omitempty"`

	// If verbose is true, sort the results by id, name, or email.
	Sort *string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAccessGroupMembersOptions : Instantiate ListAccessGroupMembersOptions
func (*IamAccessGroupsV2) NewListAccessGroupMembersOptions(accessGroupID string) *ListAccessGroupMembersOptions {
	return &ListAccessGroupMembersOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *ListAccessGroupMembersOptions) SetAccessGroupID(accessGroupID string) *ListAccessGroupMembersOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListAccessGroupMembersOptions) SetTransactionID(transactionID string) *ListAccessGroupMembersOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAccessGroupMembersOptions) SetLimit(limit float64) *ListAccessGroupMembersOptions {
	options.Limit = core.Float64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListAccessGroupMembersOptions) SetOffset(offset float64) *ListAccessGroupMembersOptions {
	options.Offset = core.Float64Ptr(offset)
	return options
}

// SetType : Allow user to set Type
func (options *ListAccessGroupMembersOptions) SetType(typeVar string) *ListAccessGroupMembersOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetVerbose : Allow user to set Verbose
func (options *ListAccessGroupMembersOptions) SetVerbose(verbose bool) *ListAccessGroupMembersOptions {
	options.Verbose = core.BoolPtr(verbose)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListAccessGroupMembersOptions) SetSort(sort string) *ListAccessGroupMembersOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupMembersOptions) SetHeaders(param map[string]string) *ListAccessGroupMembersOptions {
	options.Headers = param
	return options
}

// ListAccessGroupRulesOptions : The ListAccessGroupRules options.
type ListAccessGroupRulesOptions struct {
	// The group id that the rules are bound to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAccessGroupRulesOptions : Instantiate ListAccessGroupRulesOptions
func (*IamAccessGroupsV2) NewListAccessGroupRulesOptions(accessGroupID string) *ListAccessGroupRulesOptions {
	return &ListAccessGroupRulesOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *ListAccessGroupRulesOptions) SetAccessGroupID(accessGroupID string) *ListAccessGroupRulesOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListAccessGroupRulesOptions) SetTransactionID(transactionID string) *ListAccessGroupRulesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupRulesOptions) SetHeaders(param map[string]string) *ListAccessGroupRulesOptions {
	options.Headers = param
	return options
}

// ListAccessGroupsOptions : The ListAccessGroups options.
type ListAccessGroupsOptions struct {
	// IBM Cloud account id under which the groups are listed.
	AccountID *string `json:"account_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Return groups for member id (IBMid or Service Id).
	IamID *string `json:"iam_id,omitempty"`

	// Return up to this limit of results where limit is between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// Offset the results using this query parameter.
	Offset *int64 `json:"offset,omitempty"`

	// Sort the results by id, name, description, or is_federated flag.
	Sort *string `json:"sort,omitempty"`

	// If show_federated is true, each group listed will return an is_federated value that is set to true if rules exist
	// for the group.
	ShowFederated *bool `json:"show_federated,omitempty"`

	// If hide_public_access is true, do not include the Public Access Group in the results.
	HidePublicAccess *bool `json:"hide_public_access,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAccessGroupsOptions : Instantiate ListAccessGroupsOptions
func (*IamAccessGroupsV2) NewListAccessGroupsOptions(accountID string) *ListAccessGroupsOptions {
	return &ListAccessGroupsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListAccessGroupsOptions) SetAccountID(accountID string) *ListAccessGroupsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListAccessGroupsOptions) SetTransactionID(transactionID string) *ListAccessGroupsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *ListAccessGroupsOptions) SetIamID(iamID string) *ListAccessGroupsOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAccessGroupsOptions) SetLimit(limit int64) *ListAccessGroupsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListAccessGroupsOptions) SetOffset(offset int64) *ListAccessGroupsOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListAccessGroupsOptions) SetSort(sort string) *ListAccessGroupsOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetShowFederated : Allow user to set ShowFederated
func (options *ListAccessGroupsOptions) SetShowFederated(showFederated bool) *ListAccessGroupsOptions {
	options.ShowFederated = core.BoolPtr(showFederated)
	return options
}

// SetHidePublicAccess : Allow user to set HidePublicAccess
func (options *ListAccessGroupsOptions) SetHidePublicAccess(hidePublicAccess bool) *ListAccessGroupsOptions {
	options.HidePublicAccess = core.BoolPtr(hidePublicAccess)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupsOptions) SetHeaders(param map[string]string) *ListAccessGroupsOptions {
	options.Headers = param
	return options
}

// ListGroupMembersResponseMember : A single member of an access group in a list.
type ListGroupMembersResponseMember struct {
	// The IBMid or Service Id of the member.
	IamID *string `json:"iam_id,omitempty"`

	// The member type - either `user` or `service`.
	Type *string `json:"type,omitempty"`

	// The user's or service id's name.
	Name *string `json:"name,omitempty"`

	// If the member type is user, this is the user's email.
	Email *string `json:"email,omitempty"`

	// If the member type is service, this is the service id's description.
	Description *string `json:"description,omitempty"`

	// A url to the given member resource.
	Href *string `json:"href,omitempty"`

	// The timestamp the membership was created at.
	CreatedAt *string `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the membership.
	CreatedByID *string `json:"created_by_id,omitempty"`
}


// UnmarshalListGroupMembersResponseMember constructs an instance of ListGroupMembersResponseMember from the specified map.
func UnmarshalListGroupMembersResponseMember(m map[string]interface{}) (result *ListGroupMembersResponseMember, err error) {
	obj := new(ListGroupMembersResponseMember)
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Email, err = core.UnmarshalString(m, "email")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.CreatedByID, err = core.UnmarshalString(m, "created_by_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalListGroupMembersResponseMemberSlice unmarshals a slice of ListGroupMembersResponseMember instances from the specified list of maps.
func UnmarshalListGroupMembersResponseMemberSlice(s []interface{}) (slice []ListGroupMembersResponseMember, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ListGroupMembersResponseMember'")
			return
		}
		obj, e := UnmarshalListGroupMembersResponseMember(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalListGroupMembersResponseMemberAsProperty unmarshals an instance of ListGroupMembersResponseMember that is stored as a property
// within the specified map.
func UnmarshalListGroupMembersResponseMemberAsProperty(m map[string]interface{}, propertyName string) (result *ListGroupMembersResponseMember, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ListGroupMembersResponseMember'", propertyName)
			return
		}
		result, err = UnmarshalListGroupMembersResponseMember(objMap)
	}
	return
}

// UnmarshalListGroupMembersResponseMemberSliceAsProperty unmarshals a slice of ListGroupMembersResponseMember instances that are stored as a property
// within the specified map.
func UnmarshalListGroupMembersResponseMemberSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ListGroupMembersResponseMember, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ListGroupMembersResponseMember'", propertyName)
			return
		}
		slice, err = UnmarshalListGroupMembersResponseMemberSlice(vSlice)
	}
	return
}

// RemoveAccessGroupRuleOptions : The RemoveAccessGroupRule options.
type RemoveAccessGroupRuleOptions struct {
	// The group id that the rule is bound to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The rule to delete.
	RuleID *string `json:"rule_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveAccessGroupRuleOptions : Instantiate RemoveAccessGroupRuleOptions
func (*IamAccessGroupsV2) NewRemoveAccessGroupRuleOptions(accessGroupID string, ruleID string) *RemoveAccessGroupRuleOptions {
	return &RemoveAccessGroupRuleOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		RuleID: core.StringPtr(ruleID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *RemoveAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *RemoveAccessGroupRuleOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetRuleID : Allow user to set RuleID
func (options *RemoveAccessGroupRuleOptions) SetRuleID(ruleID string) *RemoveAccessGroupRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *RemoveAccessGroupRuleOptions) SetTransactionID(transactionID string) *RemoveAccessGroupRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveAccessGroupRuleOptions) SetHeaders(param map[string]string) *RemoveAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// RemoveMemberFromAccessGroupOptions : The RemoveMemberFromAccessGroup options.
type RemoveMemberFromAccessGroupOptions struct {
	// The access_group_id to find the membership in.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The iam_id to remove from the group.
	IamID *string `json:"iam_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveMemberFromAccessGroupOptions : Instantiate RemoveMemberFromAccessGroupOptions
func (*IamAccessGroupsV2) NewRemoveMemberFromAccessGroupOptions(accessGroupID string, iamID string) *RemoveMemberFromAccessGroupOptions {
	return &RemoveMemberFromAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *RemoveMemberFromAccessGroupOptions) SetAccessGroupID(accessGroupID string) *RemoveMemberFromAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *RemoveMemberFromAccessGroupOptions) SetIamID(iamID string) *RemoveMemberFromAccessGroupOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *RemoveMemberFromAccessGroupOptions) SetTransactionID(transactionID string) *RemoveMemberFromAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMemberFromAccessGroupOptions) SetHeaders(param map[string]string) *RemoveMemberFromAccessGroupOptions {
	options.Headers = param
	return options
}

// RemoveMemberFromAllAccessGroupsOptions : The RemoveMemberFromAllAccessGroups options.
type RemoveMemberFromAllAccessGroupsOptions struct {
	// IBM Cloud account id for the group membership deletion.
	AccountID *string `json:"account_id" validate:"required"`

	// The iam_id to remove from all groups.
	IamID *string `json:"iam_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveMemberFromAllAccessGroupsOptions : Instantiate RemoveMemberFromAllAccessGroupsOptions
func (*IamAccessGroupsV2) NewRemoveMemberFromAllAccessGroupsOptions(accountID string, iamID string) *RemoveMemberFromAllAccessGroupsOptions {
	return &RemoveMemberFromAllAccessGroupsOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *RemoveMemberFromAllAccessGroupsOptions) SetAccountID(accountID string) *RemoveMemberFromAllAccessGroupsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *RemoveMemberFromAllAccessGroupsOptions) SetIamID(iamID string) *RemoveMemberFromAllAccessGroupsOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *RemoveMemberFromAllAccessGroupsOptions) SetTransactionID(transactionID string) *RemoveMemberFromAllAccessGroupsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMemberFromAllAccessGroupsOptions) SetHeaders(param map[string]string) *RemoveMemberFromAllAccessGroupsOptions {
	options.Headers = param
	return options
}

// RemoveMembersFromAccessGroupOptions : The RemoveMembersFromAccessGroup options.
type RemoveMembersFromAccessGroupOptions struct {
	// The access_group_id to find the memberships in.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The `iam_id`s to remove from the access group. This field has a limit of 50 `iam_id`s.
	Members []string `json:"members,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveMembersFromAccessGroupOptions : Instantiate RemoveMembersFromAccessGroupOptions
func (*IamAccessGroupsV2) NewRemoveMembersFromAccessGroupOptions(accessGroupID string) *RemoveMembersFromAccessGroupOptions {
	return &RemoveMembersFromAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *RemoveMembersFromAccessGroupOptions) SetAccessGroupID(accessGroupID string) *RemoveMembersFromAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetMembers : Allow user to set Members
func (options *RemoveMembersFromAccessGroupOptions) SetMembers(members []string) *RemoveMembersFromAccessGroupOptions {
	options.Members = members
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *RemoveMembersFromAccessGroupOptions) SetTransactionID(transactionID string) *RemoveMembersFromAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMembersFromAccessGroupOptions) SetHeaders(param map[string]string) *RemoveMembersFromAccessGroupOptions {
	options.Headers = param
	return options
}

// ReplaceAccessGroupRuleOptions : The ReplaceAccessGroupRule options.
type ReplaceAccessGroupRuleOptions struct {
	// The group id that the rule is bound to.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The rule to update.
	RuleID *string `json:"rule_id" validate:"required"`

	// The current revision number of the rule being updated. This can be found in the Get Rule response Etag header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The number of hours that the rule lives for (Must be between 1 and 24).
	Expiration *int64 `json:"expiration" validate:"required"`

	// The url of the identity provider.
	RealmName *string `json:"realm_name" validate:"required"`

	// A list of conditions the rule must satisfy.
	Conditions []RuleConditions `json:"conditions" validate:"required"`

	// The name of the rule.
	Name *string `json:"name,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceAccessGroupRuleOptions : Instantiate ReplaceAccessGroupRuleOptions
func (*IamAccessGroupsV2) NewReplaceAccessGroupRuleOptions(accessGroupID string, ruleID string, ifMatch string, expiration int64, realmName string, conditions []RuleConditions) *ReplaceAccessGroupRuleOptions {
	return &ReplaceAccessGroupRuleOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		RuleID: core.StringPtr(ruleID),
		IfMatch: core.StringPtr(ifMatch),
		Expiration: core.Int64Ptr(expiration),
		RealmName: core.StringPtr(realmName),
		Conditions: conditions,
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *ReplaceAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *ReplaceAccessGroupRuleOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetRuleID : Allow user to set RuleID
func (options *ReplaceAccessGroupRuleOptions) SetRuleID(ruleID string) *ReplaceAccessGroupRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *ReplaceAccessGroupRuleOptions) SetIfMatch(ifMatch string) *ReplaceAccessGroupRuleOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetExpiration : Allow user to set Expiration
func (options *ReplaceAccessGroupRuleOptions) SetExpiration(expiration int64) *ReplaceAccessGroupRuleOptions {
	options.Expiration = core.Int64Ptr(expiration)
	return options
}

// SetRealmName : Allow user to set RealmName
func (options *ReplaceAccessGroupRuleOptions) SetRealmName(realmName string) *ReplaceAccessGroupRuleOptions {
	options.RealmName = core.StringPtr(realmName)
	return options
}

// SetConditions : Allow user to set Conditions
func (options *ReplaceAccessGroupRuleOptions) SetConditions(conditions []RuleConditions) *ReplaceAccessGroupRuleOptions {
	options.Conditions = conditions
	return options
}

// SetName : Allow user to set Name
func (options *ReplaceAccessGroupRuleOptions) SetName(name string) *ReplaceAccessGroupRuleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ReplaceAccessGroupRuleOptions) SetTransactionID(transactionID string) *ReplaceAccessGroupRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceAccessGroupRuleOptions) SetHeaders(param map[string]string) *ReplaceAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// Rule : A rule of an access group.
type Rule struct {
	// The rule id.
	ID *string `json:"id,omitempty"`

	// The name of the rule.
	Name *string `json:"name,omitempty"`

	// The number of hours that the rule lives for (Must be between 1 and 24).
	Expiration *int64 `json:"expiration,omitempty"`

	// The url of the identity provider.
	RealmName *string `json:"realm_name,omitempty"`

	// The group id that the rule is assigned to.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The account id that the group is in.
	AccountID *string `json:"account_id,omitempty"`

	// A list of conditions the rule must satisfy.
	Conditions []RuleConditions `json:"conditions,omitempty"`

	// The timestamp the rule was created at.
	CreatedAt *string `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the rule.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The timestamp the rule was last edited at.
	LastModifiedAt *string `json:"last_modified_at,omitempty"`

	// The IAM id that last modified the rule.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`
}


// UnmarshalRule constructs an instance of Rule from the specified map.
func UnmarshalRule(m map[string]interface{}) (result *Rule, err error) {
	obj := new(Rule)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Expiration, err = core.UnmarshalInt64(m, "expiration")
	if err != nil {
		return
	}
	obj.RealmName, err = core.UnmarshalString(m, "realm_name")
	if err != nil {
		return
	}
	obj.AccessGroupID, err = core.UnmarshalString(m, "access_group_id")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Conditions, err = UnmarshalRuleConditionsSliceAsProperty(m, "conditions")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	obj.CreatedByID, err = core.UnmarshalString(m, "created_by_id")
	if err != nil {
		return
	}
	obj.LastModifiedAt, err = core.UnmarshalString(m, "last_modified_at")
	if err != nil {
		return
	}
	obj.LastModifiedByID, err = core.UnmarshalString(m, "last_modified_by_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalRuleSlice unmarshals a slice of Rule instances from the specified list of maps.
func UnmarshalRuleSlice(s []interface{}) (slice []Rule, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Rule'")
			return
		}
		obj, e := UnmarshalRule(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalRuleAsProperty unmarshals an instance of Rule that is stored as a property
// within the specified map.
func UnmarshalRuleAsProperty(m map[string]interface{}, propertyName string) (result *Rule, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Rule'", propertyName)
			return
		}
		result, err = UnmarshalRule(objMap)
	}
	return
}

// UnmarshalRuleSliceAsProperty unmarshals a slice of Rule instances that are stored as a property
// within the specified map.
func UnmarshalRuleSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Rule, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Rule'", propertyName)
			return
		}
		slice, err = UnmarshalRuleSlice(vSlice)
	}
	return
}

// RuleConditions : The conditions of a rule.
type RuleConditions struct {
	// The claim to evaluate against. This will be found in the `ext` claims of a user's login request.
	Claim *string `json:"claim" validate:"required"`

	// The operation to perform on the claim. Valid operators are EQUALS, EQUALS_IGNORE_CASE, IN, NOT_EQUALS_IGNORE_CASE,
	// NOT_EQUALS, and CONTAINS.
	Operator *string `json:"operator" validate:"required"`

	// The stringified JSON value that the claim is compared to using the operator.
	Value *string `json:"value" validate:"required"`
}


// NewRuleConditions : Instantiate RuleConditions (Generic Model Constructor)
func (*IamAccessGroupsV2) NewRuleConditions(claim string, operator string, value string) (model *RuleConditions, err error) {
	model = &RuleConditions{
		Claim: core.StringPtr(claim),
		Operator: core.StringPtr(operator),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleConditions constructs an instance of RuleConditions from the specified map.
func UnmarshalRuleConditions(m map[string]interface{}) (result *RuleConditions, err error) {
	obj := new(RuleConditions)
	obj.Claim, err = core.UnmarshalString(m, "claim")
	if err != nil {
		return
	}
	obj.Operator, err = core.UnmarshalString(m, "operator")
	if err != nil {
		return
	}
	obj.Value, err = core.UnmarshalString(m, "value")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalRuleConditionsSlice unmarshals a slice of RuleConditions instances from the specified list of maps.
func UnmarshalRuleConditionsSlice(s []interface{}) (slice []RuleConditions, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'RuleConditions'")
			return
		}
		obj, e := UnmarshalRuleConditions(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalRuleConditionsAsProperty unmarshals an instance of RuleConditions that is stored as a property
// within the specified map.
func UnmarshalRuleConditionsAsProperty(m map[string]interface{}, propertyName string) (result *RuleConditions, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'RuleConditions'", propertyName)
			return
		}
		result, err = UnmarshalRuleConditions(objMap)
	}
	return
}

// UnmarshalRuleConditionsSliceAsProperty unmarshals a slice of RuleConditions instances that are stored as a property
// within the specified map.
func UnmarshalRuleConditionsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []RuleConditions, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'RuleConditions'", propertyName)
			return
		}
		slice, err = UnmarshalRuleConditionsSlice(vSlice)
	}
	return
}

// RulesList : A list of rules attached to the access group.
type RulesList struct {
	// A list of rules.
	Rules []Rule `json:"rules,omitempty"`
}


// UnmarshalRulesList constructs an instance of RulesList from the specified map.
func UnmarshalRulesList(m map[string]interface{}) (result *RulesList, err error) {
	obj := new(RulesList)
	obj.Rules, err = UnmarshalRuleSliceAsProperty(m, "rules")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalRulesListSlice unmarshals a slice of RulesList instances from the specified list of maps.
func UnmarshalRulesListSlice(s []interface{}) (slice []RulesList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'RulesList'")
			return
		}
		obj, e := UnmarshalRulesList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalRulesListAsProperty unmarshals an instance of RulesList that is stored as a property
// within the specified map.
func UnmarshalRulesListAsProperty(m map[string]interface{}, propertyName string) (result *RulesList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'RulesList'", propertyName)
			return
		}
		result, err = UnmarshalRulesList(objMap)
	}
	return
}

// UnmarshalRulesListSliceAsProperty unmarshals a slice of RulesList instances that are stored as a property
// within the specified map.
func UnmarshalRulesListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []RulesList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'RulesList'", propertyName)
			return
		}
		slice, err = UnmarshalRulesListSlice(vSlice)
	}
	return
}

// UpdateAccessGroupOptions : The UpdateAccessGroup options.
type UpdateAccessGroupOptions struct {
	// The Access group to update.
	AccessGroupID *string `json:"access_group_id" validate:"required"`

	// The current revision number of the group being updated. This can be found in the Create/Get Access Group response
	// Etag header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Assign the specified name to the Access Group. This field has a limit of 100 characters.
	Name *string `json:"name,omitempty"`

	// Assign a description for the Access Group. This field has a limit of 250 characters.
	Description *string `json:"description,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccessGroupOptions : Instantiate UpdateAccessGroupOptions
func (*IamAccessGroupsV2) NewUpdateAccessGroupOptions(accessGroupID string, ifMatch string) *UpdateAccessGroupOptions {
	return &UpdateAccessGroupOptions{
		AccessGroupID: core.StringPtr(accessGroupID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetAccessGroupID : Allow user to set AccessGroupID
func (options *UpdateAccessGroupOptions) SetAccessGroupID(accessGroupID string) *UpdateAccessGroupOptions {
	options.AccessGroupID = core.StringPtr(accessGroupID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateAccessGroupOptions) SetIfMatch(ifMatch string) *UpdateAccessGroupOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateAccessGroupOptions) SetName(name string) *UpdateAccessGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateAccessGroupOptions) SetDescription(description string) *UpdateAccessGroupOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateAccessGroupOptions) SetTransactionID(transactionID string) *UpdateAccessGroupOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccessGroupOptions) SetHeaders(param map[string]string) *UpdateAccessGroupOptions {
	options.Headers = param
	return options
}

// UpdateAccountSettingsOptions : The UpdateAccountSettings options.
type UpdateAccountSettingsOptions struct {
	// The account id of the settings being updated.
	AccountID *string `json:"account_id" validate:"required"`

	// This flag controls the public access feature within the account. It is set to true by default. Note: When this flag
	// is set to false, all policies within the account attached to the Public Access group will be deleted.
	PublicAccessEnabled *bool `json:"public_access_enabled,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccountSettingsOptions : Instantiate UpdateAccountSettingsOptions
func (*IamAccessGroupsV2) NewUpdateAccountSettingsOptions(accountID string) *UpdateAccountSettingsOptions {
	return &UpdateAccountSettingsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateAccountSettingsOptions) SetAccountID(accountID string) *UpdateAccountSettingsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetPublicAccessEnabled : Allow user to set PublicAccessEnabled
func (options *UpdateAccountSettingsOptions) SetPublicAccessEnabled(publicAccessEnabled bool) *UpdateAccountSettingsOptions {
	options.PublicAccessEnabled = core.BoolPtr(publicAccessEnabled)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateAccountSettingsOptions) SetTransactionID(transactionID string) *UpdateAccountSettingsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountSettingsOptions) SetHeaders(param map[string]string) *UpdateAccountSettingsOptions {
	options.Headers = param
	return options
}
