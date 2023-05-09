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
 * IBM OpenAPI SDK Code Generator Version: 3.70.0-7df966bf-20230419-195904
 */

// Package iamaccessgroupsv2 : Operations and models for the IamAccessGroupsV2 service
package iamaccessgroupsv2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// IamAccessGroupsV2 : The IAM Access Groups API allows for the management of access groups (Create, Read, Update,
// Delete) as well as the management of memberships and rules within the group container.
//
// API Version: 2.0
type IamAccessGroupsV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://iam.cloud.ibm.com"

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

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "iamAccessGroups" suitable for processing requests.
func (iamAccessGroups *IamAccessGroupsV2) Clone() *IamAccessGroupsV2 {
	if core.IsNil(iamAccessGroups) {
		return nil
	}
	clone := *iamAccessGroups
	clone.Service = iamAccessGroups.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (iamAccessGroups *IamAccessGroupsV2) SetServiceURL(url string) error {
	return iamAccessGroups.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (iamAccessGroups *IamAccessGroupsV2) GetServiceURL() string {
	return iamAccessGroups.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (iamAccessGroups *IamAccessGroupsV2) SetDefaultHeaders(headers http.Header) {
	iamAccessGroups.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (iamAccessGroups *IamAccessGroupsV2) SetEnableGzipCompression(enableGzip bool) {
	iamAccessGroups.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (iamAccessGroups *IamAccessGroupsV2) GetEnableGzipCompression() bool {
	return iamAccessGroups.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (iamAccessGroups *IamAccessGroupsV2) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	iamAccessGroups.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (iamAccessGroups *IamAccessGroupsV2) DisableRetries() {
	iamAccessGroups.Service.DisableRetries()
}

// CreateAccessGroup : Create an access group
// Create a new access group to assign multiple users and service ids to multiple policies. The group will be created in
// the account specified by the `account_id` parameter. The group name is a required field, but a description is
// optional. Because the group's name does not have to be unique, it is possible to create multiple groups with the same
// name.
func (iamAccessGroups *IamAccessGroupsV2) CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	return iamAccessGroups.CreateAccessGroupWithContext(context.Background(), createAccessGroupOptions)
}

// CreateAccessGroupWithContext is an alternate form of the CreateAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) CreateAccessGroupWithContext(ctx context.Context, createAccessGroupOptions *CreateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAccessGroupOptions, "createAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAccessGroupOptions, "createAccessGroupOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups`, nil)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroup)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAccessGroups : List access groups
// This API lists access groups within an account. Parameters for pagination and sorting can be used to filter the
// results. The `account_id` query parameter determines which account to retrieve groups from. Only the groups you have
// access to are returned (either because of a policy on a specific group or account level access (admin, editor, or
// viewer)). There may be more groups in the account that aren't shown if you lack the aforementioned permissions.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) (result *GroupsList, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListAccessGroupsWithContext(context.Background(), listAccessGroupsOptions)
}

// ListAccessGroupsWithContext is an alternate form of the ListAccessGroups method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupsWithContext(ctx context.Context, listAccessGroupsOptions *ListAccessGroupsOptions) (result *GroupsList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupsOptions, "listAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupsOptions, "listAccessGroupsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups`, nil)
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
	if listAccessGroupsOptions.Search != nil {
		builder.AddQuery("search", fmt.Sprint(*listAccessGroupsOptions.Search))
	}
	if listAccessGroupsOptions.MembershipType != nil {
		builder.AddQuery("membership_type", fmt.Sprint(*listAccessGroupsOptions.MembershipType))
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroupsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccessGroup : Get an access group
// Retrieve an access group by its `access_group_id`. Only the groups data is returned (group name, description,
// account_id, ...), not membership or rule information. A revision number is returned in the `ETag` header, which is
// needed when updating the access group.
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetAccessGroupWithContext(context.Background(), getAccessGroupOptions)
}

// GetAccessGroupWithContext is an alternate form of the GetAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroupWithContext(ctx context.Context, getAccessGroupOptions *GetAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccessGroupOptions, "getAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccessGroupOptions, "getAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *getAccessGroupOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroup)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateAccessGroup : Update an access group
// Update the group name or description of an existing access group using this API. An `If-Match` header must be
// populated with the group's most recent revision number (which can be acquired in the `Get an access group` API).
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	return iamAccessGroups.UpdateAccessGroupWithContext(context.Background(), updateAccessGroupOptions)
}

// UpdateAccessGroupWithContext is an alternate form of the UpdateAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccessGroupWithContext(ctx context.Context, updateAccessGroupOptions *UpdateAccessGroupOptions) (result *Group, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccessGroupOptions, "updateAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccessGroupOptions, "updateAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *updateAccessGroupOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroup)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAccessGroup : Delete an access group
// This API is used for deleting an access group. If the access group has no members or rules associated with it, the
// group and its policies will be deleted. However, if rules or members do exist, set the `force` parameter to true to
// delete the group as well as its associated members, rules, and policies.
func (iamAccessGroups *IamAccessGroupsV2) DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.DeleteAccessGroupWithContext(context.Background(), deleteAccessGroupOptions)
}

// DeleteAccessGroupWithContext is an alternate form of the DeleteAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) DeleteAccessGroupWithContext(ctx context.Context, deleteAccessGroupOptions *DeleteAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAccessGroupOptions, "deleteAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAccessGroupOptions, "deleteAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *deleteAccessGroupOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}`, pathParamsMap)
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

// IsMemberOfAccessGroup : Check membership in an access group
// This HEAD operation determines if a given `iam_id` is present in a group either explicitly or via dynamic rules. No
// response body is returned with this request. If the membership exists, a `204 - No Content` status code is returned.
// If the membership or the group does not exist, a `404 - Not Found` status code is returned.
func (iamAccessGroups *IamAccessGroupsV2) IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.IsMemberOfAccessGroupWithContext(context.Background(), isMemberOfAccessGroupOptions)
}

// IsMemberOfAccessGroupWithContext is an alternate form of the IsMemberOfAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) IsMemberOfAccessGroupWithContext(ctx context.Context, isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(isMemberOfAccessGroupOptions, "isMemberOfAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(isMemberOfAccessGroupOptions, "isMemberOfAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *isMemberOfAccessGroupOptions.AccessGroupID,
		"iam_id": *isMemberOfAccessGroupOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.HEAD)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/members/{iam_id}`, pathParamsMap)
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

// AddMembersToAccessGroup : Add members to an access group
// Use this API to add users (`IBMid-...`), service IDs (`iam-ServiceId-...`) or trusted profiles (`iam-Profile-...`) to
// an access group. Any member added gains access to resources defined in the group's policies. To revoke a given
// members's access, simply remove them from the group. There is no limit to the number of members one group can have,
// but each `iam_id` can only be added to 50 groups. Additionally, this API request payload can add up to 50 members per
// call.
func (iamAccessGroups *IamAccessGroupsV2) AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions) (result *AddGroupMembersResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.AddMembersToAccessGroupWithContext(context.Background(), addMembersToAccessGroupOptions)
}

// AddMembersToAccessGroupWithContext is an alternate form of the AddMembersToAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) AddMembersToAccessGroupWithContext(ctx context.Context, addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions) (result *AddGroupMembersResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addMembersToAccessGroupOptions, "addMembersToAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addMembersToAccessGroupOptions, "addMembersToAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *addMembersToAccessGroupOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/members`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddGroupMembersResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAccessGroupMembers : List access group members
// List all members of a given group using this API. Parameters for pagination and sorting can be used to filter the
// results. The most useful query parameter may be the `verbose` flag. If `verbose=true`, user, service ID and trusted
// profile names will be retrieved for each `iam_id`. If performance is a concern, leave the `verbose` parameter off so
// that name information does not get retrieved.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions) (result *GroupMembersList, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListAccessGroupMembersWithContext(context.Background(), listAccessGroupMembersOptions)
}

// ListAccessGroupMembersWithContext is an alternate form of the ListAccessGroupMembers method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupMembersWithContext(ctx context.Context, listAccessGroupMembersOptions *ListAccessGroupMembersOptions) (result *GroupMembersList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupMembersOptions, "listAccessGroupMembersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupMembersOptions, "listAccessGroupMembersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *listAccessGroupMembersOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/members`, pathParamsMap)
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

	if listAccessGroupMembersOptions.MembershipType != nil {
		builder.AddQuery("membership_type", fmt.Sprint(*listAccessGroupMembersOptions.MembershipType))
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGroupMembersList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoveMemberFromAccessGroup : Delete member from an access group
// Remove one member from a group using this API. If the operation is successful, only a `204 - No Content` response
// with no body is returned. However, if any error occurs, the standard error format will be returned. Dynamic member
// cannot be deleted using this API. Dynamic rules needs to be adjusted to delete dynamic members.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.RemoveMemberFromAccessGroupWithContext(context.Background(), removeMemberFromAccessGroupOptions)
}

// RemoveMemberFromAccessGroupWithContext is an alternate form of the RemoveMemberFromAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAccessGroupWithContext(ctx context.Context, removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMemberFromAccessGroupOptions, "removeMemberFromAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMemberFromAccessGroupOptions, "removeMemberFromAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *removeMemberFromAccessGroupOptions.AccessGroupID,
		"iam_id": *removeMemberFromAccessGroupOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/members/{iam_id}`, pathParamsMap)
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

// RemoveMembersFromAccessGroup : Delete members from an access group
// Remove multiple members from a group using this API. On a successful call, this API will always return 207. It is the
// caller's responsibility to iterate across the body to determine successful deletion of each member. This API request
// payload can delete up to 50 members per call. This API doesnt delete dynamic members accessing the access group via
// dynamic rules.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions) (result *DeleteGroupBulkMembersResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.RemoveMembersFromAccessGroupWithContext(context.Background(), removeMembersFromAccessGroupOptions)
}

// RemoveMembersFromAccessGroupWithContext is an alternate form of the RemoveMembersFromAccessGroup method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) RemoveMembersFromAccessGroupWithContext(ctx context.Context, removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions) (result *DeleteGroupBulkMembersResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMembersFromAccessGroupOptions, "removeMembersFromAccessGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMembersFromAccessGroupOptions, "removeMembersFromAccessGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *removeMembersFromAccessGroupOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/members/delete`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteGroupBulkMembersResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoveMemberFromAllAccessGroups : Delete member from all access groups
// This API removes a given member from every group they are a member of within the specified account. By using one
// operation, you can revoke one member's access to all access groups in the account. If a partial failure occurs on
// deletion, the response will be shown in the body.
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) (result *DeleteFromAllGroupsResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.RemoveMemberFromAllAccessGroupsWithContext(context.Background(), removeMemberFromAllAccessGroupsOptions)
}

// RemoveMemberFromAllAccessGroupsWithContext is an alternate form of the RemoveMemberFromAllAccessGroups method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) RemoveMemberFromAllAccessGroupsWithContext(ctx context.Context, removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) (result *DeleteFromAllGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeMemberFromAllAccessGroupsOptions, "removeMemberFromAllAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeMemberFromAllAccessGroupsOptions, "removeMemberFromAllAccessGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *removeMemberFromAllAccessGroupsOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/_allgroups/members/{iam_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteFromAllGroupsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddMemberToMultipleAccessGroups : Add member to multiple access groups
// This API will add a member to multiple access groups in an account. The limit of how many groups that can be in the
// request is 50. The response is a list of results that show if adding the member to each group was successful or not.
func (iamAccessGroups *IamAccessGroupsV2) AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) (result *AddMembershipMultipleGroupsResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.AddMemberToMultipleAccessGroupsWithContext(context.Background(), addMemberToMultipleAccessGroupsOptions)
}

// AddMemberToMultipleAccessGroupsWithContext is an alternate form of the AddMemberToMultipleAccessGroups method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) AddMemberToMultipleAccessGroupsWithContext(ctx context.Context, addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) (result *AddMembershipMultipleGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addMemberToMultipleAccessGroupsOptions, "addMemberToMultipleAccessGroupsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addMemberToMultipleAccessGroupsOptions, "addMemberToMultipleAccessGroupsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *addMemberToMultipleAccessGroupsOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/_allgroups/members/{iam_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddMembershipMultipleGroupsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddAccessGroupRule : Create rule for an access group
// Rules can be used to dynamically add users to an access group. If a user's SAML assertions match the rule's
// conditions during login, the user will be dynamically added to the group. The duration of the user's access to the
// group is determined by the `expiration` field. After access expires, the user will need to log in again to regain
// access. Note that the condition's value field must be a stringified JSON value. [Consult this documentation for
// further explanation of dynamic rules.](/docs/account?topic=account-rules).
func (iamAccessGroups *IamAccessGroupsV2) AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return iamAccessGroups.AddAccessGroupRuleWithContext(context.Background(), addAccessGroupRuleOptions)
}

// AddAccessGroupRuleWithContext is an alternate form of the AddAccessGroupRule method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) AddAccessGroupRuleWithContext(ctx context.Context, addAccessGroupRuleOptions *AddAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addAccessGroupRuleOptions, "addAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addAccessGroupRuleOptions, "addAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *addAccessGroupRuleOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/rules`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAccessGroupRules : List access group rules
// This API lists all rules in a given access group. Because only a few rules are created on each group, there is no
// pagination or sorting support on this API.
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions) (result *RulesList, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListAccessGroupRulesWithContext(context.Background(), listAccessGroupRulesOptions)
}

// ListAccessGroupRulesWithContext is an alternate form of the ListAccessGroupRules method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListAccessGroupRulesWithContext(ctx context.Context, listAccessGroupRulesOptions *ListAccessGroupRulesOptions) (result *RulesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAccessGroupRulesOptions, "listAccessGroupRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAccessGroupRulesOptions, "listAccessGroupRulesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *listAccessGroupRulesOptions.AccessGroupID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/rules`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRulesList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAccessGroupRule : Get an access group rule
// Retrieve a rule from an access group. A revision number is returned in the `ETag` header, which is needed when
// updating the rule.
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetAccessGroupRuleWithContext(context.Background(), getAccessGroupRuleOptions)
}

// GetAccessGroupRuleWithContext is an alternate form of the GetAccessGroupRule method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetAccessGroupRuleWithContext(ctx context.Context, getAccessGroupRuleOptions *GetAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccessGroupRuleOptions, "getAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccessGroupRuleOptions, "getAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *getAccessGroupRuleOptions.AccessGroupID,
		"rule_id": *getAccessGroupRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/rules/{rule_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceAccessGroupRule : Replace an access group rule
// Update the body of an existing rule using this API. An `If-Match` header must be populated with the rule's most
// recent revision number (which can be acquired in the `Get an access group rule` API).
func (iamAccessGroups *IamAccessGroupsV2) ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ReplaceAccessGroupRuleWithContext(context.Background(), replaceAccessGroupRuleOptions)
}

// ReplaceAccessGroupRuleWithContext is an alternate form of the ReplaceAccessGroupRule method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ReplaceAccessGroupRuleWithContext(ctx context.Context, replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceAccessGroupRuleOptions, "replaceAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceAccessGroupRuleOptions, "replaceAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *replaceAccessGroupRuleOptions.AccessGroupID,
		"rule_id": *replaceAccessGroupRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/rules/{rule_id}`, pathParamsMap)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RemoveAccessGroupRule : Delete an access group rule
// Remove one rule from a group using this API. If the operation is successful, only a `204 - No Content` response with
// no body is returned. However, if any error occurs, the standard error format will be returned.
func (iamAccessGroups *IamAccessGroupsV2) RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.RemoveAccessGroupRuleWithContext(context.Background(), removeAccessGroupRuleOptions)
}

// RemoveAccessGroupRuleWithContext is an alternate form of the RemoveAccessGroupRule method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) RemoveAccessGroupRuleWithContext(ctx context.Context, removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeAccessGroupRuleOptions, "removeAccessGroupRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeAccessGroupRuleOptions, "removeAccessGroupRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"access_group_id": *removeAccessGroupRuleOptions.AccessGroupID,
		"rule_id": *removeAccessGroupRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/{access_group_id}/rules/{rule_id}`, pathParamsMap)
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

// GetAccountSettings : Get account settings
// Retrieve the access groups settings for a specific account.
func (iamAccessGroups *IamAccessGroupsV2) GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetAccountSettingsWithContext(context.Background(), getAccountSettingsOptions)
}

// GetAccountSettingsWithContext is an alternate form of the GetAccountSettings method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetAccountSettingsWithContext(ctx context.Context, getAccountSettingsOptions *GetAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSettingsOptions, "getAccountSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSettingsOptions, "getAccountSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/settings`, nil)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateAccountSettings : Update account settings
// Update the access groups settings for a specific account. Note: When the `public_access_enabled` setting is set to
// false, all policies within the account attached to the Public Access group will be deleted. Only set
// `public_access_enabled` to false if you are sure that you want those policies to be removed.
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	return iamAccessGroups.UpdateAccountSettingsWithContext(context.Background(), updateAccountSettingsOptions)
}

// UpdateAccountSettingsWithContext is an alternate form of the UpdateAccountSettings method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) UpdateAccountSettingsWithContext(ctx context.Context, updateAccountSettingsOptions *UpdateAccountSettingsOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountSettingsOptions, "updateAccountSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountSettingsOptions, "updateAccountSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v2/groups/settings`, nil)
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

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTemplate : Create Template
// Endpoint to create an access group template.
func (iamAccessGroups *IamAccessGroupsV2) CreateTemplate(createTemplateOptions *CreateTemplateOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.CreateTemplateWithContext(context.Background(), createTemplateOptions)
}

// CreateTemplateWithContext is an alternate form of the CreateTemplate method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) CreateTemplateWithContext(ctx context.Context, createTemplateOptions *CreateTemplateOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTemplateOptions, "createTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTemplateOptions, "createTemplateOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "CreateTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createTemplateOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createTemplateOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createTemplateOptions.Name != nil {
		body["name"] = createTemplateOptions.Name
	}
	if createTemplateOptions.Description != nil {
		body["description"] = createTemplateOptions.Description
	}
	if createTemplateOptions.AccountID != nil {
		body["account_id"] = createTemplateOptions.AccountID
	}
	if createTemplateOptions.AccessGroup != nil {
		body["access_group"] = createTemplateOptions.AccessGroup
	}
	if createTemplateOptions.PolicyTemplates != nil {
		body["policy_templates"] = createTemplateOptions.PolicyTemplates
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
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTemplates : List Templates
// Endpoint to list access group templates in a given account.
func (iamAccessGroups *IamAccessGroupsV2) ListTemplates(listTemplatesOptions *ListTemplatesOptions) (result *ListTemplatesResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListTemplatesWithContext(context.Background(), listTemplatesOptions)
}

// ListTemplatesWithContext is an alternate form of the ListTemplates method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListTemplatesWithContext(ctx context.Context, listTemplatesOptions *ListTemplatesOptions) (result *ListTemplatesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTemplatesOptions, "listTemplatesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTemplatesOptions, "listTemplatesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTemplatesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListTemplates")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listTemplatesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listTemplatesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listTemplatesOptions.AccountID))
	if listTemplatesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listTemplatesOptions.Limit))
	}
	if listTemplatesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listTemplatesOptions.Offset))
	}
	if listTemplatesOptions.Verbose != nil {
		builder.AddQuery("verbose", fmt.Sprint(*listTemplatesOptions.Verbose))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListTemplatesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateTemplateVersion : Create template version
// Endpoint to create a new template version.
func (iamAccessGroups *IamAccessGroupsV2) CreateTemplateVersion(createTemplateVersionOptions *CreateTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.CreateTemplateVersionWithContext(context.Background(), createTemplateVersionOptions)
}

// CreateTemplateVersionWithContext is an alternate form of the CreateTemplateVersion method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) CreateTemplateVersionWithContext(ctx context.Context, createTemplateVersionOptions *CreateTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTemplateVersionOptions, "createTemplateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTemplateVersionOptions, "createTemplateVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *createTemplateVersionOptions.TemplateID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTemplateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "CreateTemplateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createTemplateVersionOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createTemplateVersionOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createTemplateVersionOptions.Name != nil {
		body["name"] = createTemplateVersionOptions.Name
	}
	if createTemplateVersionOptions.Description != nil {
		body["description"] = createTemplateVersionOptions.Description
	}
	if createTemplateVersionOptions.AccessGroup != nil {
		body["access_group"] = createTemplateVersionOptions.AccessGroup
	}
	if createTemplateVersionOptions.PolicyTemplates != nil {
		body["policy_templates"] = createTemplateVersionOptions.PolicyTemplates
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
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTemplateVersions : List template versions
// Endpoint to list all the versions of a template.
func (iamAccessGroups *IamAccessGroupsV2) ListTemplateVersions(listTemplateVersionsOptions *ListTemplateVersionsOptions) (result *ListTemplateVersionsResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListTemplateVersionsWithContext(context.Background(), listTemplateVersionsOptions)
}

// ListTemplateVersionsWithContext is an alternate form of the ListTemplateVersions method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListTemplateVersionsWithContext(ctx context.Context, listTemplateVersionsOptions *ListTemplateVersionsOptions) (result *ListTemplateVersionsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listTemplateVersionsOptions, "listTemplateVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listTemplateVersionsOptions, "listTemplateVersionsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *listTemplateVersionsOptions.TemplateID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTemplateVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListTemplateVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listTemplateVersionsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listTemplateVersionsOptions.Limit))
	}
	if listTemplateVersionsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listTemplateVersionsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListTemplateVersionsResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTemplateVersion : Get template version
// Endpoint to get a specific template version.
func (iamAccessGroups *IamAccessGroupsV2) GetTemplateVersion(getTemplateVersionOptions *GetTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetTemplateVersionWithContext(context.Background(), getTemplateVersionOptions)
}

// GetTemplateVersionWithContext is an alternate form of the GetTemplateVersion method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetTemplateVersionWithContext(ctx context.Context, getTemplateVersionOptions *GetTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTemplateVersionOptions, "getTemplateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTemplateVersionOptions, "getTemplateVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *getTemplateVersionOptions.TemplateID,
		"version_num": *getTemplateVersionOptions.VersionNum,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions/{version_num}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTemplateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetTemplateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getTemplateVersionOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getTemplateVersionOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateTemplateVersion : Update template version
// Endpoint to update a template version.
func (iamAccessGroups *IamAccessGroupsV2) UpdateTemplateVersion(updateTemplateVersionOptions *UpdateTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.UpdateTemplateVersionWithContext(context.Background(), updateTemplateVersionOptions)
}

// UpdateTemplateVersionWithContext is an alternate form of the UpdateTemplateVersion method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) UpdateTemplateVersionWithContext(ctx context.Context, updateTemplateVersionOptions *UpdateTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTemplateVersionOptions, "updateTemplateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTemplateVersionOptions, "updateTemplateVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *updateTemplateVersionOptions.TemplateID,
		"version_num": *updateTemplateVersionOptions.VersionNum,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions/{version_num}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTemplateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "UpdateTemplateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateTemplateVersionOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateTemplateVersionOptions.IfMatch))
	}
	if updateTemplateVersionOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateTemplateVersionOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateTemplateVersionOptions.ID != nil {
		body["id"] = updateTemplateVersionOptions.ID
	}
	if updateTemplateVersionOptions.Name != nil {
		body["name"] = updateTemplateVersionOptions.Name
	}
	if updateTemplateVersionOptions.Description != nil {
		body["description"] = updateTemplateVersionOptions.Description
	}
	if updateTemplateVersionOptions.AccountID != nil {
		body["account_id"] = updateTemplateVersionOptions.AccountID
	}
	if updateTemplateVersionOptions.Version != nil {
		body["version"] = updateTemplateVersionOptions.Version
	}
	if updateTemplateVersionOptions.Committed != nil {
		body["committed"] = updateTemplateVersionOptions.Committed
	}
	if updateTemplateVersionOptions.AccessGroup != nil {
		body["access_group"] = updateTemplateVersionOptions.AccessGroup
	}
	if updateTemplateVersionOptions.PolicyTemplates != nil {
		body["policy_templates"] = updateTemplateVersionOptions.PolicyTemplates
	}
	if updateTemplateVersionOptions.Href != nil {
		body["href"] = updateTemplateVersionOptions.Href
	}
	if updateTemplateVersionOptions.CreatedAt != nil {
		body["created_at"] = updateTemplateVersionOptions.CreatedAt
	}
	if updateTemplateVersionOptions.CreatedByID != nil {
		body["created_by_id"] = updateTemplateVersionOptions.CreatedByID
	}
	if updateTemplateVersionOptions.LastModifiedAt != nil {
		body["last_modified_at"] = updateTemplateVersionOptions.LastModifiedAt
	}
	if updateTemplateVersionOptions.LastModifiedByID != nil {
		body["last_modified_by_id"] = updateTemplateVersionOptions.LastModifiedByID
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
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTemplateVersion : Delete template version
// Endpoint to delete a template version.
func (iamAccessGroups *IamAccessGroupsV2) DeleteTemplateVersion(deleteTemplateVersionOptions *DeleteTemplateVersionOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.DeleteTemplateVersionWithContext(context.Background(), deleteTemplateVersionOptions)
}

// DeleteTemplateVersionWithContext is an alternate form of the DeleteTemplateVersion method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) DeleteTemplateVersionWithContext(ctx context.Context, deleteTemplateVersionOptions *DeleteTemplateVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTemplateVersionOptions, "deleteTemplateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTemplateVersionOptions, "deleteTemplateVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *deleteTemplateVersionOptions.TemplateID,
		"version_num": *deleteTemplateVersionOptions.VersionNum,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions/{version_num}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTemplateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "DeleteTemplateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTemplateVersionOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteTemplateVersionOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// CommitTemplate : Commit a template
// Endpoint to commit a template. After you commit the template, the version is immutable.
func (iamAccessGroups *IamAccessGroupsV2) CommitTemplate(commitTemplateOptions *CommitTemplateOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.CommitTemplateWithContext(context.Background(), commitTemplateOptions)
}

// CommitTemplateWithContext is an alternate form of the CommitTemplate method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) CommitTemplateWithContext(ctx context.Context, commitTemplateOptions *CommitTemplateOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(commitTemplateOptions, "commitTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(commitTemplateOptions, "commitTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *commitTemplateOptions.TemplateID,
		"version_num": *commitTemplateOptions.VersionNum,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}/versions/{version_num}/commit`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range commitTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "CommitTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if commitTemplateOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*commitTemplateOptions.IfMatch))
	}
	if commitTemplateOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*commitTemplateOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetLatestTemplateVersion : Get latest template version
// Endpoint to get the latest version of a template.
func (iamAccessGroups *IamAccessGroupsV2) GetLatestTemplateVersion(getLatestTemplateVersionOptions *GetLatestTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetLatestTemplateVersionWithContext(context.Background(), getLatestTemplateVersionOptions)
}

// GetLatestTemplateVersionWithContext is an alternate form of the GetLatestTemplateVersion method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetLatestTemplateVersionWithContext(ctx context.Context, getLatestTemplateVersionOptions *GetLatestTemplateVersionOptions) (result *CreateTemplateResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLatestTemplateVersionOptions, "getLatestTemplateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLatestTemplateVersionOptions, "getLatestTemplateVersionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *getLatestTemplateVersionOptions.TemplateID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLatestTemplateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetLatestTemplateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLatestTemplateVersionOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getLatestTemplateVersionOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateTemplateResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTemplate : Delete template
// Endpoint to delete a template. All access assigned by that template is deleted from all of the accounts where the
// template was assigned.
func (iamAccessGroups *IamAccessGroupsV2) DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.DeleteTemplateWithContext(context.Background(), deleteTemplateOptions)
}

// DeleteTemplateWithContext is an alternate form of the DeleteTemplate method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) DeleteTemplateWithContext(ctx context.Context, deleteTemplateOptions *DeleteTemplateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTemplateOptions, "deleteTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTemplateOptions, "deleteTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"template_id": *deleteTemplateOptions.TemplateID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_templates/{template_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "DeleteTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteTemplateOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteTemplateOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// CreateAssignment : Assign template to account
// Endpoint to assign a template to an account/account group.
func (iamAccessGroups *IamAccessGroupsV2) CreateAssignment(createAssignmentOptions *CreateAssignmentOptions) (result *TemplateCreateAssignmentResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.CreateAssignmentWithContext(context.Background(), createAssignmentOptions)
}

// CreateAssignmentWithContext is an alternate form of the CreateAssignment method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) CreateAssignmentWithContext(ctx context.Context, createAssignmentOptions *CreateAssignmentOptions) (result *TemplateCreateAssignmentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAssignmentOptions, "createAssignmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAssignmentOptions, "createAssignmentOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_assignments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAssignmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "CreateAssignment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createAssignmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createAssignmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createAssignmentOptions.TemplateID != nil {
		body["template_id"] = createAssignmentOptions.TemplateID
	}
	if createAssignmentOptions.TemplateVersion != nil {
		body["template_version"] = createAssignmentOptions.TemplateVersion
	}
	if createAssignmentOptions.TargetType != nil {
		body["target_type"] = createAssignmentOptions.TargetType
	}
	if createAssignmentOptions.Target != nil {
		body["target"] = createAssignmentOptions.Target
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
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplateCreateAssignmentResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListAssignments : List Assignment
// Endpoint to list template assignments.
func (iamAccessGroups *IamAccessGroupsV2) ListAssignments(listAssignmentsOptions *ListAssignmentsOptions) (result *TemplatesListAssignmentResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.ListAssignmentsWithContext(context.Background(), listAssignmentsOptions)
}

// ListAssignmentsWithContext is an alternate form of the ListAssignments method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) ListAssignmentsWithContext(ctx context.Context, listAssignmentsOptions *ListAssignmentsOptions) (result *TemplatesListAssignmentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAssignmentsOptions, "listAssignmentsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_assignments`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAssignmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "ListAssignments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAssignmentsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listAssignmentsOptions.AccountID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTemplatesListAssignmentResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAssignment : Get assignment
// Get a specific template assignment.
func (iamAccessGroups *IamAccessGroupsV2) GetAssignment(getAssignmentOptions *GetAssignmentOptions) (result *GetTemplateAssignmentResponse, response *core.DetailedResponse, err error) {
	return iamAccessGroups.GetAssignmentWithContext(context.Background(), getAssignmentOptions)
}

// GetAssignmentWithContext is an alternate form of the GetAssignment method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) GetAssignmentWithContext(ctx context.Context, getAssignmentOptions *GetAssignmentOptions) (result *GetTemplateAssignmentResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAssignmentOptions, "getAssignmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAssignmentOptions, "getAssignmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assignment_id": *getAssignmentOptions.AssignmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_assignments/{assignment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAssignmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "GetAssignment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAssignmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAssignmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = iamAccessGroups.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTemplateAssignmentResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteAssignment : Delete assignment
// Endpoint to delete template assignment.
func (iamAccessGroups *IamAccessGroupsV2) DeleteAssignment(deleteAssignmentOptions *DeleteAssignmentOptions) (response *core.DetailedResponse, err error) {
	return iamAccessGroups.DeleteAssignmentWithContext(context.Background(), deleteAssignmentOptions)
}

// DeleteAssignmentWithContext is an alternate form of the DeleteAssignment method which supports a Context parameter
func (iamAccessGroups *IamAccessGroupsV2) DeleteAssignmentWithContext(ctx context.Context, deleteAssignmentOptions *DeleteAssignmentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAssignmentOptions, "deleteAssignmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAssignmentOptions, "deleteAssignmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"assignment_id": *deleteAssignmentOptions.AssignmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = iamAccessGroups.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(iamAccessGroups.Service.Options.URL, `/v1/group_assignments/{assignment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAssignmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("iam_access_groups", "V2", "DeleteAssignment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteAssignmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteAssignmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = iamAccessGroups.Service.Request(request, nil)

	return
}

// AccessActionControls : Control whether or not child account administrators can add access policies to the enterprise-managed access group in
// their account.
type AccessActionControls struct {
	// Action control for adding access policies.
	Add *bool `json:"add,omitempty"`
}

// UnmarshalAccessActionControls unmarshals an instance of AccessActionControls from the specified map of raw messages.
func UnmarshalAccessActionControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccessActionControls)
	err = core.UnmarshalPrimitive(m, "add", &obj.Add)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccessGroupInput : Access Group Input Component.
type AccessGroupInput struct {
	// access group input name.
	Name *string `json:"name" validate:"required"`

	// access group input description.
	Description *string `json:"description,omitempty"`

	// Members Input component.
	Members *MembersInput `json:"members,omitempty"`

	// Assertions Input Component.
	Assertions *AssertionsInput `json:"assertions,omitempty"`

	// Access Group Action Controls component.
	ActionControls *GroupActionControls `json:"action_controls,omitempty"`
}

// NewAccessGroupInput : Instantiate AccessGroupInput (Generic Model Constructor)
func (*IamAccessGroupsV2) NewAccessGroupInput(name string) (_model *AccessGroupInput, err error) {
	_model = &AccessGroupInput{
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAccessGroupInput unmarshals an instance of AccessGroupInput from the specified map of raw messages.
func UnmarshalAccessGroupInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccessGroupInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalMembersInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "assertions", &obj.Assertions, UnmarshalAssertionsInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_controls", &obj.ActionControls, UnmarshalGroupActionControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountSettings : The access groups settings for a specific account.
type AccountSettings struct {
	// The account id of the settings being shown.
	AccountID *string `json:"account_id,omitempty"`

	// The timestamp the settings were last edited at.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The `iam_id` of the entity that last modified the settings.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// This flag controls the public access feature within the account. It is set to true by default. Note: When this flag
	// is set to false, all policies within the account attached to the Public Access group will be deleted.
	PublicAccessEnabled *bool `json:"public_access_enabled,omitempty"`
}

// UnmarshalAccountSettings unmarshals an instance of AccountSettings from the specified map of raw messages.
func UnmarshalAccountSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountSettings)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
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
	err = core.UnmarshalPrimitive(m, "public_access_enabled", &obj.PublicAccessEnabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddAccessGroupRuleOptions : The AddAccessGroupRule options.
type AddAccessGroupRuleOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// Session duration in hours. Access group membership is revoked after this time period expires. Users must log back in
	// to refresh their access group membership.
	Expiration *int64 `json:"expiration" validate:"required"`

	// The URL of the identity provider (IdP).
	RealmName *string `json:"realm_name" validate:"required"`

	// A list of conditions that identities must satisfy to gain access group membership.
	Conditions []RuleConditions `json:"conditions" validate:"required"`

	// The name of the dynaimic rule.
	Name *string `json:"name,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *AddAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *AddAccessGroupRuleOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetExpiration : Allow user to set Expiration
func (_options *AddAccessGroupRuleOptions) SetExpiration(expiration int64) *AddAccessGroupRuleOptions {
	_options.Expiration = core.Int64Ptr(expiration)
	return _options
}

// SetRealmName : Allow user to set RealmName
func (_options *AddAccessGroupRuleOptions) SetRealmName(realmName string) *AddAccessGroupRuleOptions {
	_options.RealmName = core.StringPtr(realmName)
	return _options
}

// SetConditions : Allow user to set Conditions
func (_options *AddAccessGroupRuleOptions) SetConditions(conditions []RuleConditions) *AddAccessGroupRuleOptions {
	_options.Conditions = conditions
	return _options
}

// SetName : Allow user to set Name
func (_options *AddAccessGroupRuleOptions) SetName(name string) *AddAccessGroupRuleOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddAccessGroupRuleOptions) SetTransactionID(transactionID string) *AddAccessGroupRuleOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddAccessGroupRuleOptions) SetHeaders(param map[string]string) *AddAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// AddGroupMembersRequestMembersItem : AddGroupMembersRequestMembersItem struct
type AddGroupMembersRequestMembersItem struct {
	// The IBMid, service ID or trusted profile ID of the member.
	IamID *string `json:"iam_id" validate:"required"`

	// The type of the member, must be either "user", "service" or "profile".
	Type *string `json:"type" validate:"required"`
}

// NewAddGroupMembersRequestMembersItem : Instantiate AddGroupMembersRequestMembersItem (Generic Model Constructor)
func (*IamAccessGroupsV2) NewAddGroupMembersRequestMembersItem(iamID string, typeVar string) (_model *AddGroupMembersRequestMembersItem, err error) {
	_model = &AddGroupMembersRequestMembersItem{
		IamID: core.StringPtr(iamID),
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalAddGroupMembersRequestMembersItem unmarshals an instance of AddGroupMembersRequestMembersItem from the specified map of raw messages.
func UnmarshalAddGroupMembersRequestMembersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddGroupMembersRequestMembersItem)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
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

// AddGroupMembersResponse : The members added to an access group.
type AddGroupMembersResponse struct {
	// The members added to an access group.
	Members []AddGroupMembersResponseMembersItem `json:"members,omitempty"`
}

// UnmarshalAddGroupMembersResponse unmarshals an instance of AddGroupMembersResponse from the specified map of raw messages.
func UnmarshalAddGroupMembersResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddGroupMembersResponse)
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalAddGroupMembersResponseMembersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddGroupMembersResponseMembersItem : AddGroupMembersResponseMembersItem struct
type AddGroupMembersResponseMembersItem struct {
	// The IBMid or Service Id of the member.
	IamID *string `json:"iam_id,omitempty"`

	// The member type - either `user`, `service` or `profile`.
	Type *string `json:"type,omitempty"`

	// The timestamp of when the membership was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the membership.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The outcome of the operation on this `iam_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// A list of errors that occurred when trying to add members to a group.
	Errors []Error `json:"errors,omitempty"`
}

// UnmarshalAddGroupMembersResponseMembersItem unmarshals an instance of AddGroupMembersResponseMembersItem from the specified map of raw messages.
func UnmarshalAddGroupMembersResponseMembersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddGroupMembersResponseMembersItem)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
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
	err = core.UnmarshalPrimitive(m, "created_by_id", &obj.CreatedByID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddMemberToMultipleAccessGroupsOptions : The AddMemberToMultipleAccessGroups options.
type AddMemberToMultipleAccessGroupsOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// The IAM identifier.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// The type of the member, must be either "user", "service" or "profile".
	Type *string `json:"type,omitempty"`

	// The ids of the access groups a given member is to be added to.
	Groups []string `json:"groups,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *AddMemberToMultipleAccessGroupsOptions) SetAccountID(accountID string) *AddMemberToMultipleAccessGroupsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *AddMemberToMultipleAccessGroupsOptions) SetIamID(iamID string) *AddMemberToMultipleAccessGroupsOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetType : Allow user to set Type
func (_options *AddMemberToMultipleAccessGroupsOptions) SetType(typeVar string) *AddMemberToMultipleAccessGroupsOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetGroups : Allow user to set Groups
func (_options *AddMemberToMultipleAccessGroupsOptions) SetGroups(groups []string) *AddMemberToMultipleAccessGroupsOptions {
	_options.Groups = groups
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddMemberToMultipleAccessGroupsOptions) SetTransactionID(transactionID string) *AddMemberToMultipleAccessGroupsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddMemberToMultipleAccessGroupsOptions) SetHeaders(param map[string]string) *AddMemberToMultipleAccessGroupsOptions {
	options.Headers = param
	return options
}

// AddMembersToAccessGroupOptions : The AddMembersToAccessGroup options.
type AddMembersToAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// An array of member objects to add to an access group.
	Members []AddGroupMembersRequestMembersItem `json:"members,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *AddMembersToAccessGroupOptions) SetAccessGroupID(accessGroupID string) *AddMembersToAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetMembers : Allow user to set Members
func (_options *AddMembersToAccessGroupOptions) SetMembers(members []AddGroupMembersRequestMembersItem) *AddMembersToAccessGroupOptions {
	_options.Members = members
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *AddMembersToAccessGroupOptions) SetTransactionID(transactionID string) *AddMembersToAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
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

// UnmarshalAddMembershipMultipleGroupsResponse unmarshals an instance of AddMembershipMultipleGroupsResponse from the specified map of raw messages.
func UnmarshalAddMembershipMultipleGroupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddMembershipMultipleGroupsResponse)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalAddMembershipMultipleGroupsResponseGroupsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddMembershipMultipleGroupsResponseGroupsItem : AddMembershipMultipleGroupsResponseGroupsItem struct
type AddMembershipMultipleGroupsResponseGroupsItem struct {
	// The access group that the member is to be added to.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The outcome of the add membership operation on this `access_group_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// List of errors encountered when adding member to access group.
	Errors []Error `json:"errors,omitempty"`
}

// UnmarshalAddMembershipMultipleGroupsResponseGroupsItem unmarshals an instance of AddMembershipMultipleGroupsResponseGroupsItem from the specified map of raw messages.
func UnmarshalAddMembershipMultipleGroupsResponseGroupsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddMembershipMultipleGroupsResponseGroupsItem)
	err = core.UnmarshalPrimitive(m, "access_group_id", &obj.AccessGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssertionsActionControls : Assertions Action Controls component.
type AssertionsActionControls struct {
	// assertions action controls add.
	Add *bool `json:"add,omitempty"`

	// assertions action controls remove.
	Remove *bool `json:"remove,omitempty"`

	// assertions action controls update.
	Update *bool `json:"update,omitempty"`
}

// UnmarshalAssertionsActionControls unmarshals an instance of AssertionsActionControls from the specified map of raw messages.
func UnmarshalAssertionsActionControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssertionsActionControls)
	err = core.UnmarshalPrimitive(m, "add", &obj.Add)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "remove", &obj.Remove)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update", &obj.Update)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssertionsInput : Assertions Input Component.
type AssertionsInput struct {
	// assertions input rules.
	Rules []RuleInput `json:"rules,omitempty"`

	// Assertions Action Controls component.
	ActionControls *AssertionsActionControls `json:"action_controls,omitempty"`
}

// UnmarshalAssertionsInput unmarshals an instance of AssertionsInput from the specified map of raw messages.
func UnmarshalAssertionsInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssertionsInput)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRuleInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_controls", &obj.ActionControls, UnmarshalAssertionsActionControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssignmentResourceAccessGroup : Assignment Resource Access Group.
type AssignmentResourceAccessGroup struct {
	// Assignment resource entry.
	Group *AssignmentResourceEntry `json:"group" validate:"required"`

	// List of member resources of the group.
	Members []AssignmentResourceEntry `json:"members" validate:"required"`

	// List of rules associated with the group.
	Rules []AssignmentResourceEntry `json:"rules" validate:"required"`
}

// UnmarshalAssignmentResourceAccessGroup unmarshals an instance of AssignmentResourceAccessGroup from the specified map of raw messages.
func UnmarshalAssignmentResourceAccessGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssignmentResourceAccessGroup)
	err = core.UnmarshalModel(m, "group", &obj.Group, UnmarshalAssignmentResourceEntry)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalAssignmentResourceEntry)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalAssignmentResourceEntry)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssignmentResourceEntry : Assignment resource entry.
type AssignmentResourceEntry struct {
	// Assignment Resource Entry Id.
	ID *string `json:"id" validate:"required"`

	// Optional name of the resource.
	Name *string `json:"name,omitempty"`

	// Optional version of the resource.
	Version *string `json:"version,omitempty"`

	// Resource in assignment resource entry.
	Resource *string `json:"resource" validate:"required"`

	// Error in assignment resource entry.
	Error *string `json:"error" validate:"required"`

	// Optional operation on the resource.
	Operation *string `json:"operation,omitempty"`

	// Status of assignment resource entry.
	Status *string `json:"status" validate:"required"`
}

// UnmarshalAssignmentResourceEntry unmarshals an instance of AssignmentResourceEntry from the specified map of raw messages.
func UnmarshalAssignmentResourceEntry(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssignmentResourceEntry)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource", &obj.Resource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommitTemplateOptions : The CommitTemplate options.
type CommitTemplateOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// version number in path.
	VersionNum *string `json:"version_num" validate:"required,ne="`

	// ETag value of the template version document.
	IfMatch *string `json:"If-Match" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCommitTemplateOptions : Instantiate CommitTemplateOptions
func (*IamAccessGroupsV2) NewCommitTemplateOptions(templateID string, versionNum string, ifMatch string) *CommitTemplateOptions {
	return &CommitTemplateOptions{
		TemplateID: core.StringPtr(templateID),
		VersionNum: core.StringPtr(versionNum),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *CommitTemplateOptions) SetTemplateID(templateID string) *CommitTemplateOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetVersionNum : Allow user to set VersionNum
func (_options *CommitTemplateOptions) SetVersionNum(versionNum string) *CommitTemplateOptions {
	_options.VersionNum = core.StringPtr(versionNum)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *CommitTemplateOptions) SetIfMatch(ifMatch string) *CommitTemplateOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CommitTemplateOptions) SetTransactionID(transactionID string) *CommitTemplateOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CommitTemplateOptions) SetHeaders(param map[string]string) *CommitTemplateOptions {
	options.Headers = param
	return options
}

// ConditionInput : Condition Input component.
type ConditionInput struct {
	// condition input claim.
	Claim *string `json:"claim,omitempty"`

	// condition input operator.
	Operator *string `json:"operator,omitempty"`

	// condition input value.
	Value *string `json:"value,omitempty"`
}

// UnmarshalConditionInput unmarshals an instance of ConditionInput from the specified map of raw messages.
func UnmarshalConditionInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConditionInput)
	err = core.UnmarshalPrimitive(m, "claim", &obj.Claim)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
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

// CreateAccessGroupOptions : The CreateAccessGroup options.
type CreateAccessGroupOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// Give the access group a unique name that doesn't conflict with an existing access group in the account. This field
	// is case-insensitive and has a limit of 100 characters.
	Name *string `json:"name" validate:"required"`

	// Assign an optional description for the access group. This field has a limit of 250 characters.
	Description *string `json:"description,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *CreateAccessGroupOptions) SetAccountID(accountID string) *CreateAccessGroupOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateAccessGroupOptions) SetName(name string) *CreateAccessGroupOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateAccessGroupOptions) SetDescription(description string) *CreateAccessGroupOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateAccessGroupOptions) SetTransactionID(transactionID string) *CreateAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccessGroupOptions) SetHeaders(param map[string]string) *CreateAccessGroupOptions {
	options.Headers = param
	return options
}

// CreateAssignmentOptions : The CreateAssignment options.
type CreateAssignmentOptions struct {
	// The unique identifier of the template to be assigned.
	TemplateID *string `json:"template_id" validate:"required"`

	// The version number of the template to be assigned.
	TemplateVersion *string `json:"template_version" validate:"required"`

	// The type of the entity to which the template should be assigned, e.g. 'account', 'accountGroup', etc.
	TargetType *string `json:"target_type" validate:"required"`

	// The unique identifier of the entity to which the template should be assigned.
	Target *string `json:"target" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAssignmentOptions : Instantiate CreateAssignmentOptions
func (*IamAccessGroupsV2) NewCreateAssignmentOptions(templateID string, templateVersion string, targetType string, target string) *CreateAssignmentOptions {
	return &CreateAssignmentOptions{
		TemplateID: core.StringPtr(templateID),
		TemplateVersion: core.StringPtr(templateVersion),
		TargetType: core.StringPtr(targetType),
		Target: core.StringPtr(target),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *CreateAssignmentOptions) SetTemplateID(templateID string) *CreateAssignmentOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetTemplateVersion : Allow user to set TemplateVersion
func (_options *CreateAssignmentOptions) SetTemplateVersion(templateVersion string) *CreateAssignmentOptions {
	_options.TemplateVersion = core.StringPtr(templateVersion)
	return _options
}

// SetTargetType : Allow user to set TargetType
func (_options *CreateAssignmentOptions) SetTargetType(targetType string) *CreateAssignmentOptions {
	_options.TargetType = core.StringPtr(targetType)
	return _options
}

// SetTarget : Allow user to set Target
func (_options *CreateAssignmentOptions) SetTarget(target string) *CreateAssignmentOptions {
	_options.Target = core.StringPtr(target)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateAssignmentOptions) SetTransactionID(transactionID string) *CreateAssignmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAssignmentOptions) SetHeaders(param map[string]string) *CreateAssignmentOptions {
	options.Headers = param
	return options
}

// CreateTemplateOptions : The CreateTemplate options.
type CreateTemplateOptions struct {
	// create template input name.
	Name *string `json:"name" validate:"required"`

	// create template input description.
	Description *string `json:"description" validate:"required"`

	// create template input account id.
	AccountID *string `json:"account_id" validate:"required"`

	// Access Group Input Component.
	AccessGroup *AccessGroupInput `json:"access_group,omitempty"`

	// policy template references.
	PolicyTemplates []PolicyTemplatesInput `json:"policy_templates,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTemplateOptions : Instantiate CreateTemplateOptions
func (*IamAccessGroupsV2) NewCreateTemplateOptions(name string, description string, accountID string) *CreateTemplateOptions {
	return &CreateTemplateOptions{
		Name: core.StringPtr(name),
		Description: core.StringPtr(description),
		AccountID: core.StringPtr(accountID),
	}
}

// SetName : Allow user to set Name
func (_options *CreateTemplateOptions) SetName(name string) *CreateTemplateOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateTemplateOptions) SetDescription(description string) *CreateTemplateOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *CreateTemplateOptions) SetAccountID(accountID string) *CreateTemplateOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetAccessGroup : Allow user to set AccessGroup
func (_options *CreateTemplateOptions) SetAccessGroup(accessGroup *AccessGroupInput) *CreateTemplateOptions {
	_options.AccessGroup = accessGroup
	return _options
}

// SetPolicyTemplates : Allow user to set PolicyTemplates
func (_options *CreateTemplateOptions) SetPolicyTemplates(policyTemplates []PolicyTemplatesInput) *CreateTemplateOptions {
	_options.PolicyTemplates = policyTemplates
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateTemplateOptions) SetTransactionID(transactionID string) *CreateTemplateOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTemplateOptions) SetHeaders(param map[string]string) *CreateTemplateOptions {
	options.Headers = param
	return options
}

// CreateTemplateResponse : Successful response output for create template.
type CreateTemplateResponse struct {
	// The ID of the access group template.
	ID *string `json:"id,omitempty"`

	// The name of the access group template.
	Name *string `json:"name,omitempty"`

	// The description of the access group template.
	Description *string `json:"description,omitempty"`

	// The ID of the account to which the access group template is assigned.
	AccountID *string `json:"account_id,omitempty"`

	// The version of the access group template.
	Version *string `json:"version,omitempty"`

	// A boolean indicating whether the access group template is committed.
	Committed *bool `json:"committed,omitempty"`

	// Access Group Input Component.
	AccessGroup *AccessGroupInput `json:"access_group,omitempty"`

	// References to policy templates assigned to the access group template.
	PolicyTemplates []PolicyTemplatesInput `json:"policy_templates,omitempty"`

	// The URL of the access group template resource.
	Href *string `json:"href,omitempty"`

	// The date and time when the access group template was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The ID of the user who created the access group template.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The date and time when the access group template was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The ID of the user who last modified the access group template.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`
}

// UnmarshalCreateTemplateResponse unmarshals an instance of CreateTemplateResponse from the specified map of raw messages.
func UnmarshalCreateTemplateResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateTemplateResponse)
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
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "committed", &obj.Committed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access_group", &obj.AccessGroup, UnmarshalAccessGroupInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "policy_templates", &obj.PolicyTemplates, UnmarshalPolicyTemplatesInput)
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

// CreateTemplateVersionOptions : The CreateTemplateVersion options.
type CreateTemplateVersionOptions struct {
	// parameter template id.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// The name of the template version.
	Name *string `json:"name,omitempty"`

	// The description of the template version.
	Description *string `json:"description,omitempty"`

	// Access Group Input Component.
	AccessGroup *AccessGroupInput `json:"access_group,omitempty"`

	// The policy templates associated with the template version.
	PolicyTemplates []PolicyTemplatesInput `json:"policy_templates,omitempty"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTemplateVersionOptions : Instantiate CreateTemplateVersionOptions
func (*IamAccessGroupsV2) NewCreateTemplateVersionOptions(templateID string) *CreateTemplateVersionOptions {
	return &CreateTemplateVersionOptions{
		TemplateID: core.StringPtr(templateID),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *CreateTemplateVersionOptions) SetTemplateID(templateID string) *CreateTemplateVersionOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateTemplateVersionOptions) SetName(name string) *CreateTemplateVersionOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateTemplateVersionOptions) SetDescription(description string) *CreateTemplateVersionOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAccessGroup : Allow user to set AccessGroup
func (_options *CreateTemplateVersionOptions) SetAccessGroup(accessGroup *AccessGroupInput) *CreateTemplateVersionOptions {
	_options.AccessGroup = accessGroup
	return _options
}

// SetPolicyTemplates : Allow user to set PolicyTemplates
func (_options *CreateTemplateVersionOptions) SetPolicyTemplates(policyTemplates []PolicyTemplatesInput) *CreateTemplateVersionOptions {
	_options.PolicyTemplates = policyTemplates
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *CreateTemplateVersionOptions) SetTransactionID(transactionID string) *CreateTemplateVersionOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTemplateVersionOptions) SetHeaders(param map[string]string) *CreateTemplateVersionOptions {
	options.Headers = param
	return options
}

// DeleteAccessGroupOptions : The DeleteAccessGroup options.
type DeleteAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *DeleteAccessGroupOptions) SetAccessGroupID(accessGroupID string) *DeleteAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteAccessGroupOptions) SetTransactionID(transactionID string) *DeleteAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetForce : Allow user to set Force
func (_options *DeleteAccessGroupOptions) SetForce(force bool) *DeleteAccessGroupOptions {
	_options.Force = core.BoolPtr(force)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAccessGroupOptions) SetHeaders(param map[string]string) *DeleteAccessGroupOptions {
	options.Headers = param
	return options
}

// DeleteAssignmentOptions : The DeleteAssignment options.
type DeleteAssignmentOptions struct {
	// assignment id path parameter.
	AssignmentID *string `json:"assignment_id" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAssignmentOptions : Instantiate DeleteAssignmentOptions
func (*IamAccessGroupsV2) NewDeleteAssignmentOptions(assignmentID string) *DeleteAssignmentOptions {
	return &DeleteAssignmentOptions{
		AssignmentID: core.StringPtr(assignmentID),
	}
}

// SetAssignmentID : Allow user to set AssignmentID
func (_options *DeleteAssignmentOptions) SetAssignmentID(assignmentID string) *DeleteAssignmentOptions {
	_options.AssignmentID = core.StringPtr(assignmentID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteAssignmentOptions) SetTransactionID(transactionID string) *DeleteAssignmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAssignmentOptions) SetHeaders(param map[string]string) *DeleteAssignmentOptions {
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

// UnmarshalDeleteFromAllGroupsResponse unmarshals an instance of DeleteFromAllGroupsResponse from the specified map of raw messages.
func UnmarshalDeleteFromAllGroupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteFromAllGroupsResponse)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalDeleteFromAllGroupsResponseGroupsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteFromAllGroupsResponseGroupsItem : DeleteFromAllGroupsResponseGroupsItem struct
type DeleteFromAllGroupsResponseGroupsItem struct {
	// The access group that the member is to be deleted from.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The outcome of the delete operation on this `access_group_id`.
	StatusCode *int64 `json:"status_code,omitempty"`

	// A transaction-id that can be used for debugging purposes.
	Trace *string `json:"trace,omitempty"`

	// A list of errors that occurred when trying to remove a member from groups.
	Errors []Error `json:"errors,omitempty"`
}

// UnmarshalDeleteFromAllGroupsResponseGroupsItem unmarshals an instance of DeleteFromAllGroupsResponseGroupsItem from the specified map of raw messages.
func UnmarshalDeleteFromAllGroupsResponseGroupsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteFromAllGroupsResponseGroupsItem)
	err = core.UnmarshalPrimitive(m, "access_group_id", &obj.AccessGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteGroupBulkMembersResponse : The access group id and the members removed from it.
type DeleteGroupBulkMembersResponse struct {
	// The access group id.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The `iam_id`s removed from the access group.
	Members []DeleteGroupBulkMembersResponseMembersItem `json:"members,omitempty"`
}

// UnmarshalDeleteGroupBulkMembersResponse unmarshals an instance of DeleteGroupBulkMembersResponse from the specified map of raw messages.
func UnmarshalDeleteGroupBulkMembersResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteGroupBulkMembersResponse)
	err = core.UnmarshalPrimitive(m, "access_group_id", &obj.AccessGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalDeleteGroupBulkMembersResponseMembersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
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

// UnmarshalDeleteGroupBulkMembersResponseMembersItem unmarshals an instance of DeleteGroupBulkMembersResponseMembersItem from the specified map of raw messages.
func UnmarshalDeleteGroupBulkMembersResponseMembersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteGroupBulkMembersResponseMembersItem)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteTemplateOptions : The DeleteTemplate options.
type DeleteTemplateOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTemplateOptions : Instantiate DeleteTemplateOptions
func (*IamAccessGroupsV2) NewDeleteTemplateOptions(templateID string) *DeleteTemplateOptions {
	return &DeleteTemplateOptions{
		TemplateID: core.StringPtr(templateID),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *DeleteTemplateOptions) SetTemplateID(templateID string) *DeleteTemplateOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteTemplateOptions) SetTransactionID(transactionID string) *DeleteTemplateOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTemplateOptions) SetHeaders(param map[string]string) *DeleteTemplateOptions {
	options.Headers = param
	return options
}

// DeleteTemplateVersionOptions : The DeleteTemplateVersion options.
type DeleteTemplateVersionOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// version number in path.
	VersionNum *string `json:"version_num" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTemplateVersionOptions : Instantiate DeleteTemplateVersionOptions
func (*IamAccessGroupsV2) NewDeleteTemplateVersionOptions(templateID string, versionNum string) *DeleteTemplateVersionOptions {
	return &DeleteTemplateVersionOptions{
		TemplateID: core.StringPtr(templateID),
		VersionNum: core.StringPtr(versionNum),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *DeleteTemplateVersionOptions) SetTemplateID(templateID string) *DeleteTemplateVersionOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetVersionNum : Allow user to set VersionNum
func (_options *DeleteTemplateVersionOptions) SetVersionNum(versionNum string) *DeleteTemplateVersionOptions {
	_options.VersionNum = core.StringPtr(versionNum)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *DeleteTemplateVersionOptions) SetTransactionID(transactionID string) *DeleteTemplateVersionOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTemplateVersionOptions) SetHeaders(param map[string]string) *DeleteTemplateVersionOptions {
	options.Headers = param
	return options
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

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
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

// GetAccessGroupOptions : The GetAccessGroup options.
type GetAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *GetAccessGroupOptions) SetAccessGroupID(accessGroupID string) *GetAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetAccessGroupOptions) SetTransactionID(transactionID string) *GetAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetShowFederated : Allow user to set ShowFederated
func (_options *GetAccessGroupOptions) SetShowFederated(showFederated bool) *GetAccessGroupOptions {
	_options.ShowFederated = core.BoolPtr(showFederated)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccessGroupOptions) SetHeaders(param map[string]string) *GetAccessGroupOptions {
	options.Headers = param
	return options
}

// GetAccessGroupRuleOptions : The GetAccessGroupRule options.
type GetAccessGroupRuleOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The rule to get.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *GetAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *GetAccessGroupRuleOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *GetAccessGroupRuleOptions) SetRuleID(ruleID string) *GetAccessGroupRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetAccessGroupRuleOptions) SetTransactionID(transactionID string) *GetAccessGroupRuleOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccessGroupRuleOptions) SetHeaders(param map[string]string) *GetAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// GetAccountSettingsOptions : The GetAccountSettings options.
type GetAccountSettingsOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *GetAccountSettingsOptions) SetAccountID(accountID string) *GetAccountSettingsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetAccountSettingsOptions) SetTransactionID(transactionID string) *GetAccountSettingsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountSettingsOptions) SetHeaders(param map[string]string) *GetAccountSettingsOptions {
	options.Headers = param
	return options
}

// GetAssignmentOptions : The GetAssignment options.
type GetAssignmentOptions struct {
	// assignment id parameter.
	AssignmentID *string `json:"assignment_id" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAssignmentOptions : Instantiate GetAssignmentOptions
func (*IamAccessGroupsV2) NewGetAssignmentOptions(assignmentID string) *GetAssignmentOptions {
	return &GetAssignmentOptions{
		AssignmentID: core.StringPtr(assignmentID),
	}
}

// SetAssignmentID : Allow user to set AssignmentID
func (_options *GetAssignmentOptions) SetAssignmentID(assignmentID string) *GetAssignmentOptions {
	_options.AssignmentID = core.StringPtr(assignmentID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetAssignmentOptions) SetTransactionID(transactionID string) *GetAssignmentOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAssignmentOptions) SetHeaders(param map[string]string) *GetAssignmentOptions {
	options.Headers = param
	return options
}

// GetLatestTemplateVersionOptions : The GetLatestTemplateVersion options.
type GetLatestTemplateVersionOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLatestTemplateVersionOptions : Instantiate GetLatestTemplateVersionOptions
func (*IamAccessGroupsV2) NewGetLatestTemplateVersionOptions(templateID string) *GetLatestTemplateVersionOptions {
	return &GetLatestTemplateVersionOptions{
		TemplateID: core.StringPtr(templateID),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *GetLatestTemplateVersionOptions) SetTemplateID(templateID string) *GetLatestTemplateVersionOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetLatestTemplateVersionOptions) SetTransactionID(transactionID string) *GetLatestTemplateVersionOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLatestTemplateVersionOptions) SetHeaders(param map[string]string) *GetLatestTemplateVersionOptions {
	options.Headers = param
	return options
}

// GetTemplateAssignmentResponse : Response object containing the details of a template assignment.
type GetTemplateAssignmentResponse struct {
	// The ID of the assignment.
	ID *string `json:"id" validate:"required"`

	// The ID of the account that the assignment belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the template that the assignment is based on.
	TemplateID *string `json:"template_id" validate:"required"`

	// The version of the template that the assignment is based on.
	TemplateVersion *string `json:"template_version" validate:"required"`

	// The type of the entity that the assignment applies to.
	TargetType *string `json:"target_type" validate:"required"`

	// The ID of the entity that the assignment applies to.
	Target *string `json:"target" validate:"required"`

	// The operation that the assignment applies to (e.g. 'create', 'update', 'delete').
	Operation *string `json:"operation" validate:"required"`

	// The status of the assignment (e.g. 'pending', 'success', 'failure').
	Status *string `json:"status" validate:"required"`

	// List of resources for the assignment.
	Resources []ResourceListWithTargetAccountID `json:"resources" validate:"required"`

	// The URL of the assignment resource.
	Href *string `json:"href" validate:"required"`

	// The date and time when the assignment was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The user or system that created the assignment.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The date and time when the assignment was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The user or system that last updated the assignment.
	UpdatedBy *string `json:"updated_by" validate:"required"`
}

// UnmarshalGetTemplateAssignmentResponse unmarshals an instance of GetTemplateAssignmentResponse from the specified map of raw messages.
func UnmarshalGetTemplateAssignmentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTemplateAssignmentResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_id", &obj.TemplateID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_version", &obj.TemplateVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceListWithTargetAccountID)
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
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTemplateVersionOptions : The GetTemplateVersion options.
type GetTemplateVersionOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// path parameter verison number.
	VersionNum *string `json:"version_num" validate:"required,ne="`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTemplateVersionOptions : Instantiate GetTemplateVersionOptions
func (*IamAccessGroupsV2) NewGetTemplateVersionOptions(templateID string, versionNum string) *GetTemplateVersionOptions {
	return &GetTemplateVersionOptions{
		TemplateID: core.StringPtr(templateID),
		VersionNum: core.StringPtr(versionNum),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *GetTemplateVersionOptions) SetTemplateID(templateID string) *GetTemplateVersionOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetVersionNum : Allow user to set VersionNum
func (_options *GetTemplateVersionOptions) SetVersionNum(versionNum string) *GetTemplateVersionOptions {
	_options.VersionNum = core.StringPtr(versionNum)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *GetTemplateVersionOptions) SetTransactionID(transactionID string) *GetTemplateVersionOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTemplateVersionOptions) SetHeaders(param map[string]string) *GetTemplateVersionOptions {
	options.Headers = param
	return options
}

// Group : An IAM access group.
type Group struct {
	// The group's access group ID.
	ID *string `json:"id,omitempty"`

	// The group's name.
	Name *string `json:"name,omitempty"`

	// The group's description - if defined.
	Description *string `json:"description,omitempty"`

	// The account id where the group was created.
	AccountID *string `json:"account_id,omitempty"`

	// The timestamp of when the group was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the group.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The timestamp of when the group was last edited.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The `iam_id` of the entity that last modified the group name or description.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// A url to the given group resource.
	Href *string `json:"href,omitempty"`

	// This is set to true if rules exist for the group.
	IsFederated *bool `json:"is_federated,omitempty"`
}

// UnmarshalGroup unmarshals an instance of Group from the specified map of raw messages.
func UnmarshalGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Group)
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
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
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
	err = core.UnmarshalPrimitive(m, "is_federated", &obj.IsFederated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupActionControls : Access Group Action Controls component.
type GroupActionControls struct {
	// Control whether or not child account administrators can add access policies to the enterprise-managed access group
	// in their account.
	Access *AccessActionControls `json:"access,omitempty"`
}

// UnmarshalGroupActionControls unmarshals an instance of GroupActionControls from the specified map of raw messages.
func UnmarshalGroupActionControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupActionControls)
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalAccessActionControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GroupMembersList : The members of a group.
type GroupMembersList struct {
	// Limit on how many items can be returned.
	Limit *int64 `json:"limit" validate:"required"`

	// The offset of the first item returned in the result set.
	Offset *int64 `json:"offset" validate:"required"`

	// The total number of items that match the query.
	TotalCount *int64 `json:"total_count" validate:"required"`

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

// UnmarshalGroupMembersList unmarshals an instance of GroupMembersList from the specified map of raw messages.
func UnmarshalGroupMembersList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupMembersList)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "members", &obj.Members, UnmarshalListGroupMembersResponseMember)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *GroupMembersList) GetNextOffset() (*int64, error) {
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

// GroupsList : The list of access groups returned as part of a response.
type GroupsList struct {
	// Limit on how many items can be returned.
	Limit *int64 `json:"limit" validate:"required"`

	// The offset of the first item returned in the result set.
	Offset *int64 `json:"offset" validate:"required"`

	// The total number of items that match the query.
	TotalCount *int64 `json:"total_count" validate:"required"`

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

// UnmarshalGroupsList unmarshals an instance of GroupsList from the specified map of raw messages.
func UnmarshalGroupsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GroupsList)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups", &obj.Groups, UnmarshalGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *GroupsList) GetNextOffset() (*int64, error) {
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

// HrefStruct : A link object.
type HrefStruct struct {
	// A string containing the link’s URL.
	Href *string `json:"href,omitempty"`
}

// UnmarshalHrefStruct unmarshals an instance of HrefStruct from the specified map of raw messages.
func UnmarshalHrefStruct(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HrefStruct)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IsMemberOfAccessGroupOptions : The IsMemberOfAccessGroup options.
type IsMemberOfAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The IAM identifier.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *IsMemberOfAccessGroupOptions) SetAccessGroupID(accessGroupID string) *IsMemberOfAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *IsMemberOfAccessGroupOptions) SetIamID(iamID string) *IsMemberOfAccessGroupOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *IsMemberOfAccessGroupOptions) SetTransactionID(transactionID string) *IsMemberOfAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *IsMemberOfAccessGroupOptions) SetHeaders(param map[string]string) *IsMemberOfAccessGroupOptions {
	options.Headers = param
	return options
}

// ListAccessGroupMembersOptions : The ListAccessGroupMembers options.
type ListAccessGroupMembersOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Filters members by membership type. Filter by `static`, `dynamic` or `all`. `static` lists the members explicitly
	// added to the access group, and `dynamic` lists the members that are part of the access group at that time via
	// dynamic rules. `all` lists both static and dynamic members.
	MembershipType *string `json:"membership_type,omitempty"`

	// Return up to this limit of results where limit is between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// The offset of the first result item to be returned.
	Offset *int64 `json:"offset,omitempty"`

	// Filter the results by member type.
	Type *string `json:"type,omitempty"`

	// Return user's email and name for each user ID or the name for each service ID or trusted profile.
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
func (_options *ListAccessGroupMembersOptions) SetAccessGroupID(accessGroupID string) *ListAccessGroupMembersOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListAccessGroupMembersOptions) SetTransactionID(transactionID string) *ListAccessGroupMembersOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetMembershipType : Allow user to set MembershipType
func (_options *ListAccessGroupMembersOptions) SetMembershipType(membershipType string) *ListAccessGroupMembersOptions {
	_options.MembershipType = core.StringPtr(membershipType)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAccessGroupMembersOptions) SetLimit(limit int64) *ListAccessGroupMembersOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListAccessGroupMembersOptions) SetOffset(offset int64) *ListAccessGroupMembersOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetType : Allow user to set Type
func (_options *ListAccessGroupMembersOptions) SetType(typeVar string) *ListAccessGroupMembersOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *ListAccessGroupMembersOptions) SetVerbose(verbose bool) *ListAccessGroupMembersOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListAccessGroupMembersOptions) SetSort(sort string) *ListAccessGroupMembersOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupMembersOptions) SetHeaders(param map[string]string) *ListAccessGroupMembersOptions {
	options.Headers = param
	return options
}

// ListAccessGroupRulesOptions : The ListAccessGroupRules options.
type ListAccessGroupRulesOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *ListAccessGroupRulesOptions) SetAccessGroupID(accessGroupID string) *ListAccessGroupRulesOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListAccessGroupRulesOptions) SetTransactionID(transactionID string) *ListAccessGroupRulesOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupRulesOptions) SetHeaders(param map[string]string) *ListAccessGroupRulesOptions {
	options.Headers = param
	return options
}

// ListAccessGroupsOptions : The ListAccessGroups options.
type ListAccessGroupsOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Return groups for member ID (IBMid, service ID or trusted profile ID).
	IamID *string `json:"iam_id,omitempty"`

	// Use search to filter access groups list by id, name or description.
	// * `search=id:<ACCESS_GROUP_ID>` - To list access groups by id
	// * `search=name:<ACCESS_GROUP_NAME>` - To list access groups by name
	// * `search=description:<ACCESS_GROUP_DESC>` - To list access groups by description.
	Search *string `json:"search,omitempty"`

	// Membership type need to be specified along with iam_id and must be either `static`, `dynamic` or `all`. If
	// membership type is `static`, members explicitly added to the group will be shown. If membership type is `dynamic`,
	// members accessing the access group at the moment via dynamic rules will be shown. If membership type is `all`, both
	// static and dynamic members will be shown.
	MembershipType *string `json:"membership_type,omitempty"`

	// Return up to this limit of results where limit is between 0 and 100.
	Limit *int64 `json:"limit,omitempty"`

	// The offset of the first result item to be returned.
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
func (_options *ListAccessGroupsOptions) SetAccountID(accountID string) *ListAccessGroupsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListAccessGroupsOptions) SetTransactionID(transactionID string) *ListAccessGroupsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *ListAccessGroupsOptions) SetIamID(iamID string) *ListAccessGroupsOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetSearch : Allow user to set Search
func (_options *ListAccessGroupsOptions) SetSearch(search string) *ListAccessGroupsOptions {
	_options.Search = core.StringPtr(search)
	return _options
}

// SetMembershipType : Allow user to set MembershipType
func (_options *ListAccessGroupsOptions) SetMembershipType(membershipType string) *ListAccessGroupsOptions {
	_options.MembershipType = core.StringPtr(membershipType)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListAccessGroupsOptions) SetLimit(limit int64) *ListAccessGroupsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListAccessGroupsOptions) SetOffset(offset int64) *ListAccessGroupsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetSort : Allow user to set Sort
func (_options *ListAccessGroupsOptions) SetSort(sort string) *ListAccessGroupsOptions {
	_options.Sort = core.StringPtr(sort)
	return _options
}

// SetShowFederated : Allow user to set ShowFederated
func (_options *ListAccessGroupsOptions) SetShowFederated(showFederated bool) *ListAccessGroupsOptions {
	_options.ShowFederated = core.BoolPtr(showFederated)
	return _options
}

// SetHidePublicAccess : Allow user to set HidePublicAccess
func (_options *ListAccessGroupsOptions) SetHidePublicAccess(hidePublicAccess bool) *ListAccessGroupsOptions {
	_options.HidePublicAccess = core.BoolPtr(hidePublicAccess)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccessGroupsOptions) SetHeaders(param map[string]string) *ListAccessGroupsOptions {
	options.Headers = param
	return options
}

// ListAssignmentsOptions : The ListAssignments options.
type ListAssignmentsOptions struct {
	// query parameter account id.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAssignmentsOptions : Instantiate ListAssignmentsOptions
func (*IamAccessGroupsV2) NewListAssignmentsOptions() *ListAssignmentsOptions {
	return &ListAssignmentsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListAssignmentsOptions) SetAccountID(accountID string) *ListAssignmentsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListAssignmentsOptions) SetHeaders(param map[string]string) *ListAssignmentsOptions {
	options.Headers = param
	return options
}

// ListGroupMembersResponseMember : A single member of an access group in a list.
type ListGroupMembersResponseMember struct {
	// The IBMid or Service Id of the member.
	IamID *string `json:"iam_id,omitempty"`

	// The member type - either `user`, `service` or `profile`.
	Type *string `json:"type,omitempty"`

	// The membership type - either `static` or `dynamic`.
	MembershipType *string `json:"membership_type,omitempty"`

	// The user's or service id's name.
	Name *string `json:"name,omitempty"`

	// If the member type is user, this is the user's email.
	Email *string `json:"email,omitempty"`

	// If the member type is service, this is the service id's description.
	Description *string `json:"description,omitempty"`

	// A url to the given member resource.
	Href *string `json:"href,omitempty"`

	// The timestamp the membership was created at.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the membership.
	CreatedByID *string `json:"created_by_id,omitempty"`
}

// UnmarshalListGroupMembersResponseMember unmarshals an instance of ListGroupMembersResponseMember from the specified map of raw messages.
func UnmarshalListGroupMembersResponseMember(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListGroupMembersResponseMember)
	err = core.UnmarshalPrimitive(m, "iam_id", &obj.IamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "membership_type", &obj.MembershipType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListTemplateVersionsOptions : The ListTemplateVersions options.
type ListTemplateVersionsOptions struct {
	// template id parameter.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// limit parameter.
	Limit *int64 `json:"limit,omitempty"`

	// offset parameter.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTemplateVersionsOptions : Instantiate ListTemplateVersionsOptions
func (*IamAccessGroupsV2) NewListTemplateVersionsOptions(templateID string) *ListTemplateVersionsOptions {
	return &ListTemplateVersionsOptions{
		TemplateID: core.StringPtr(templateID),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *ListTemplateVersionsOptions) SetTemplateID(templateID string) *ListTemplateVersionsOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListTemplateVersionsOptions) SetLimit(limit int64) *ListTemplateVersionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListTemplateVersionsOptions) SetOffset(offset int64) *ListTemplateVersionsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTemplateVersionsOptions) SetHeaders(param map[string]string) *ListTemplateVersionsOptions {
	options.Headers = param
	return options
}

// ListTemplateVersionsResponse : Response object for listing template versions.
type ListTemplateVersionsResponse struct {
	// The maximum number of resources to return.
	Limit *int64 `json:"limit" validate:"required"`

	// The offset of the first resource in the list.
	Offset *int64 `json:"offset" validate:"required"`

	// The total number of resources in the list.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// A link object.
	First *HrefStruct `json:"first,omitempty"`

	// A link object.
	Previous *HrefStruct `json:"previous,omitempty"`

	// A link object.
	Next *HrefStruct `json:"next,omitempty"`

	// A link object.
	Last *HrefStruct `json:"last,omitempty"`

	// A list of access group template versions.
	Versions []ListTemplatesVersionsResponse `json:"versions,omitempty"`
}

// UnmarshalListTemplateVersionsResponse unmarshals an instance of ListTemplateVersionsResponse from the specified map of raw messages.
func UnmarshalListTemplateVersionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListTemplateVersionsResponse)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalListTemplatesVersionsResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListTemplateVersionsResponse) GetNextOffset() (*int64, error) {
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

// ListTemplatesOptions : The ListTemplates options.
type ListTemplatesOptions struct {
	// query parameter account id.
	AccountID *string `json:"account_id" validate:"required"`

	// An optional transaction id for the request.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// limit parameter.
	Limit *int64 `json:"limit,omitempty"`

	// offset parameter.
	Offset *int64 `json:"offset,omitempty"`

	// query parameter verbose.
	Verbose *bool `json:"verbose,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTemplatesOptions : Instantiate ListTemplatesOptions
func (*IamAccessGroupsV2) NewListTemplatesOptions(accountID string) *ListTemplatesOptions {
	return &ListTemplatesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListTemplatesOptions) SetAccountID(accountID string) *ListTemplatesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ListTemplatesOptions) SetTransactionID(transactionID string) *ListTemplatesOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListTemplatesOptions) SetLimit(limit int64) *ListTemplatesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListTemplatesOptions) SetOffset(offset int64) *ListTemplatesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetVerbose : Allow user to set Verbose
func (_options *ListTemplatesOptions) SetVerbose(verbose bool) *ListTemplatesOptions {
	_options.Verbose = core.BoolPtr(verbose)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListTemplatesOptions) SetHeaders(param map[string]string) *ListTemplatesOptions {
	options.Headers = param
	return options
}

// ListTemplatesResponse : Response object for listing templates.
type ListTemplatesResponse struct {
	// The maximum number of resources to return.
	Limit *int64 `json:"limit" validate:"required"`

	// The offset of the first resource in the list.
	Offset *int64 `json:"offset" validate:"required"`

	// The total number of resources in the list.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// A link object.
	First *HrefStruct `json:"first,omitempty"`

	// A link object.
	Previous *HrefStruct `json:"previous,omitempty"`

	// A link object.
	Next *HrefStruct `json:"next,omitempty"`

	// A link object.
	Last *HrefStruct `json:"last,omitempty"`

	// A list of templates.
	GroupsTemplates []TemplateItem `json:"groups_templates,omitempty"`
}

// UnmarshalListTemplatesResponse unmarshals an instance of ListTemplatesResponse from the specified map of raw messages.
func UnmarshalListTemplatesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListTemplatesResponse)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalHrefStruct)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups_templates", &obj.GroupsTemplates, UnmarshalTemplateItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListTemplatesResponse) GetNextOffset() (*int64, error) {
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

// ListTemplatesVersionsResponse : Response object for a single access group template version.
type ListTemplatesVersionsResponse struct {
	// The name of the template.
	Name *string `json:"name,omitempty"`

	// The description of the template.
	Description *string `json:"description,omitempty"`

	// The ID of the account associated with the template.
	AccountID *string `json:"account_id,omitempty"`

	// The version number of the template.
	Version *string `json:"version,omitempty"`

	// A boolean indicating whether the template is committed or not.
	Committed *bool `json:"committed,omitempty"`

	// Access Group Input Component.
	AccessGroup *AccessGroupInput `json:"access_group,omitempty"`

	// A list of policy templates associated with the template.
	PolicyTemplates []PolicyTemplatesInput `json:"policy_templates,omitempty"`

	// The URL to the template resource.
	Href *string `json:"href,omitempty"`

	// The date and time the template was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The ID of the user who created the template.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The date and time the template was last modified.
	LastModifiedAt *string `json:"last_modified_at,omitempty"`

	// The ID of the user who last modified the template.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`
}

// UnmarshalListTemplatesVersionsResponse unmarshals an instance of ListTemplatesVersionsResponse from the specified map of raw messages.
func UnmarshalListTemplatesVersionsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListTemplatesVersionsResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "committed", &obj.Committed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access_group", &obj.AccessGroup, UnmarshalAccessGroupInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "policy_templates", &obj.PolicyTemplates, UnmarshalPolicyTemplatesInput)
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

// MembersActionControls : Control whether or not child account administrators can add and remove members from the enterprise-managed access
// group in their account.
type MembersActionControls struct {
	// Action control for adding members.
	Add *bool `json:"add,omitempty"`

	// Action control for removing members.
	Remove *bool `json:"remove,omitempty"`
}

// UnmarshalMembersActionControls unmarshals an instance of MembersActionControls from the specified map of raw messages.
func UnmarshalMembersActionControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MembersActionControls)
	err = core.UnmarshalPrimitive(m, "add", &obj.Add)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "remove", &obj.Remove)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MembersInput : Members Input component.
type MembersInput struct {
	// Users array.
	Users []string `json:"users,omitempty"`

	// Service ids array.
	ServiceIds []string `json:"service_ids,omitempty"`

	// Control whether or not child account administrators can add and remove members from the enterprise-managed access
	// group in their account.
	ActionControls *MembersActionControls `json:"action_controls,omitempty"`
}

// UnmarshalMembersInput unmarshals an instance of MembersInput from the specified map of raw messages.
func UnmarshalMembersInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MembersInput)
	err = core.UnmarshalPrimitive(m, "users", &obj.Users)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_ids", &obj.ServiceIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_controls", &obj.ActionControls, UnmarshalMembersActionControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PolicyTemplatesInput : Policy Templates Input component.
type PolicyTemplatesInput struct {
	// policy template input id.
	ID *string `json:"id,omitempty"`

	// policy template input version.
	Version *string `json:"version,omitempty"`
}

// UnmarshalPolicyTemplatesInput unmarshals an instance of PolicyTemplatesInput from the specified map of raw messages.
func UnmarshalPolicyTemplatesInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PolicyTemplatesInput)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RemoveAccessGroupRuleOptions : The RemoveAccessGroupRule options.
type RemoveAccessGroupRuleOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The rule to get.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *RemoveAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *RemoveAccessGroupRuleOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *RemoveAccessGroupRuleOptions) SetRuleID(ruleID string) *RemoveAccessGroupRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *RemoveAccessGroupRuleOptions) SetTransactionID(transactionID string) *RemoveAccessGroupRuleOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveAccessGroupRuleOptions) SetHeaders(param map[string]string) *RemoveAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// RemoveMemberFromAccessGroupOptions : The RemoveMemberFromAccessGroup options.
type RemoveMemberFromAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The IAM identifier.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *RemoveMemberFromAccessGroupOptions) SetAccessGroupID(accessGroupID string) *RemoveMemberFromAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *RemoveMemberFromAccessGroupOptions) SetIamID(iamID string) *RemoveMemberFromAccessGroupOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *RemoveMemberFromAccessGroupOptions) SetTransactionID(transactionID string) *RemoveMemberFromAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMemberFromAccessGroupOptions) SetHeaders(param map[string]string) *RemoveMemberFromAccessGroupOptions {
	options.Headers = param
	return options
}

// RemoveMemberFromAllAccessGroupsOptions : The RemoveMemberFromAllAccessGroups options.
type RemoveMemberFromAllAccessGroupsOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// The IAM identifier.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *RemoveMemberFromAllAccessGroupsOptions) SetAccountID(accountID string) *RemoveMemberFromAllAccessGroupsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetIamID : Allow user to set IamID
func (_options *RemoveMemberFromAllAccessGroupsOptions) SetIamID(iamID string) *RemoveMemberFromAllAccessGroupsOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *RemoveMemberFromAllAccessGroupsOptions) SetTransactionID(transactionID string) *RemoveMemberFromAllAccessGroupsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMemberFromAllAccessGroupsOptions) SetHeaders(param map[string]string) *RemoveMemberFromAllAccessGroupsOptions {
	options.Headers = param
	return options
}

// RemoveMembersFromAccessGroupOptions : The RemoveMembersFromAccessGroup options.
type RemoveMembersFromAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The `iam_id`s to remove from the access group. This field has a limit of 50 `iam_id`s.
	Members []string `json:"members,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *RemoveMembersFromAccessGroupOptions) SetAccessGroupID(accessGroupID string) *RemoveMembersFromAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetMembers : Allow user to set Members
func (_options *RemoveMembersFromAccessGroupOptions) SetMembers(members []string) *RemoveMembersFromAccessGroupOptions {
	_options.Members = members
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *RemoveMembersFromAccessGroupOptions) SetTransactionID(transactionID string) *RemoveMembersFromAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveMembersFromAccessGroupOptions) SetHeaders(param map[string]string) *RemoveMembersFromAccessGroupOptions {
	options.Headers = param
	return options
}

// ReplaceAccessGroupRuleOptions : The ReplaceAccessGroupRule options.
type ReplaceAccessGroupRuleOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The rule to get.
	RuleID *string `json:"rule_id" validate:"required,ne="`

	// The current revision number of the rule being updated. This can be found in the Get Rule response ETag header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Session duration in hours. Access group membership is revoked after this time period expires. Users must log back in
	// to refresh their access group membership.
	Expiration *int64 `json:"expiration" validate:"required"`

	// The URL of the identity provider (IdP).
	RealmName *string `json:"realm_name" validate:"required"`

	// A list of conditions that identities must satisfy to gain access group membership.
	Conditions []RuleConditions `json:"conditions" validate:"required"`

	// The name of the dynaimic rule.
	Name *string `json:"name,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *ReplaceAccessGroupRuleOptions) SetAccessGroupID(accessGroupID string) *ReplaceAccessGroupRuleOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *ReplaceAccessGroupRuleOptions) SetRuleID(ruleID string) *ReplaceAccessGroupRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *ReplaceAccessGroupRuleOptions) SetIfMatch(ifMatch string) *ReplaceAccessGroupRuleOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetExpiration : Allow user to set Expiration
func (_options *ReplaceAccessGroupRuleOptions) SetExpiration(expiration int64) *ReplaceAccessGroupRuleOptions {
	_options.Expiration = core.Int64Ptr(expiration)
	return _options
}

// SetRealmName : Allow user to set RealmName
func (_options *ReplaceAccessGroupRuleOptions) SetRealmName(realmName string) *ReplaceAccessGroupRuleOptions {
	_options.RealmName = core.StringPtr(realmName)
	return _options
}

// SetConditions : Allow user to set Conditions
func (_options *ReplaceAccessGroupRuleOptions) SetConditions(conditions []RuleConditions) *ReplaceAccessGroupRuleOptions {
	_options.Conditions = conditions
	return _options
}

// SetName : Allow user to set Name
func (_options *ReplaceAccessGroupRuleOptions) SetName(name string) *ReplaceAccessGroupRuleOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *ReplaceAccessGroupRuleOptions) SetTransactionID(transactionID string) *ReplaceAccessGroupRuleOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceAccessGroupRuleOptions) SetHeaders(param map[string]string) *ReplaceAccessGroupRuleOptions {
	options.Headers = param
	return options
}

// ResourceListWithTargetAccountID : Object containing details of a resource list with target account ID.
type ResourceListWithTargetAccountID struct {
	// The ID of the entity that the resource list applies to.
	Target *string `json:"target,omitempty"`

	// Assignment Resource Access Group.
	AccessGroup *AssignmentResourceAccessGroup `json:"access_group,omitempty"`

	// List of policy template references for the resource list.
	PolicyTemplates []AssignmentResourceEntry `json:"policy_templates,omitempty"`
}

// UnmarshalResourceListWithTargetAccountID unmarshals an instance of ResourceListWithTargetAccountID from the specified map of raw messages.
func UnmarshalResourceListWithTargetAccountID(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceListWithTargetAccountID)
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access_group", &obj.AccessGroup, UnmarshalAssignmentResourceAccessGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "policy_templates", &obj.PolicyTemplates, UnmarshalAssignmentResourceEntry)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : A dynamic rule of an access group.
type Rule struct {
	// The rule id.
	ID *string `json:"id,omitempty"`

	// The name of the rule.
	Name *string `json:"name,omitempty"`

	// Session duration in hours. Access group membership is revoked after this time period expires. Users must log back in
	// to refresh their access group membership. Must be between 1 and 24.
	Expiration *int64 `json:"expiration,omitempty"`

	// The URL of the identity provider.
	RealmName *string `json:"realm_name,omitempty"`

	// The group id that the dynamic rule is assigned to.
	AccessGroupID *string `json:"access_group_id,omitempty"`

	// The account id that the group is in.
	AccountID *string `json:"account_id,omitempty"`

	// A list of conditions that identities must satisfy to gain access group membership.
	Conditions []RuleConditions `json:"conditions,omitempty"`

	// The timestamp for when the rule was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The `iam_id` of the entity that created the dynamic rule.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The timestamp for when the dynamic rule was last edited.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The IAM id that last modified the rule.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration", &obj.Expiration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "realm_name", &obj.RealmName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "access_group_id", &obj.AccessGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conditions", &obj.Conditions, UnmarshalRuleConditions)
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

// RuleConditions : The conditions of a dynamic rule.
type RuleConditions struct {
	// The claim to evaluate against. This will be found in the `ext` claims of a user's login request.
	Claim *string `json:"claim" validate:"required"`

	// The operation to perform on the claim.
	Operator *string `json:"operator" validate:"required"`

	// The stringified JSON value that the claim is compared to using the operator.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the RuleConditions.Operator property.
// The operation to perform on the claim.
const (
	RuleConditionsOperatorContainsConst = "CONTAINS"
	RuleConditionsOperatorEqualsConst = "EQUALS"
	RuleConditionsOperatorEqualsIgnoreCaseConst = "EQUALS_IGNORE_CASE"
	RuleConditionsOperatorInConst = "IN"
	RuleConditionsOperatorNotEqualsConst = "NOT_EQUALS"
	RuleConditionsOperatorNotEqualsIgnoreCaseConst = "NOT_EQUALS_IGNORE_CASE"
)

// NewRuleConditions : Instantiate RuleConditions (Generic Model Constructor)
func (*IamAccessGroupsV2) NewRuleConditions(claim string, operator string, value string) (_model *RuleConditions, err error) {
	_model = &RuleConditions{
		Claim: core.StringPtr(claim),
		Operator: core.StringPtr(operator),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuleConditions unmarshals an instance of RuleConditions from the specified map of raw messages.
func UnmarshalRuleConditions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditions)
	err = core.UnmarshalPrimitive(m, "claim", &obj.Claim)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
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

// RuleInput : Rule Input component.
type RuleInput struct {
	// rule input name.
	Name *string `json:"name,omitempty"`

	// rule input expiration.
	Expiration *int64 `json:"expiration,omitempty"`

	// rule input realm name.
	RealmName *string `json:"realm_name,omitempty"`

	// rule input conditions.
	Conditions []ConditionInput `json:"conditions,omitempty"`

	// Control whether or not child account administrators can update and remove dynamic rules from the enterprise-managed
	// access group in their account.
	ActionControls *RulesActionControls `json:"action_controls,omitempty"`
}

// UnmarshalRuleInput unmarshals an instance of RuleInput from the specified map of raw messages.
func UnmarshalRuleInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration", &obj.Expiration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "realm_name", &obj.RealmName)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "conditions", &obj.Conditions, UnmarshalConditionInput)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "action_controls", &obj.ActionControls, UnmarshalRulesActionControls)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulesActionControls : Control whether or not child account administrators can update and remove dynamic rules from the enterprise-managed
// access group in their account.
type RulesActionControls struct {
	// Action control for removing dynamic rules.
	Remove *bool `json:"remove,omitempty"`

	// Action control for updating dynamic rules.
	Update *bool `json:"update,omitempty"`
}

// UnmarshalRulesActionControls unmarshals an instance of RulesActionControls from the specified map of raw messages.
func UnmarshalRulesActionControls(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulesActionControls)
	err = core.UnmarshalPrimitive(m, "remove", &obj.Remove)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update", &obj.Update)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulesList : A list of dynamic rules attached to the access group.
type RulesList struct {
	// A list of dynamic rules.
	Rules []Rule `json:"rules,omitempty"`
}

// UnmarshalRulesList unmarshals an instance of RulesList from the specified map of raw messages.
func UnmarshalRulesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulesList)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TemplateCreateAssignmentResponse : Response object containing the details of a template assignment.
type TemplateCreateAssignmentResponse struct {
	// The ID of the assignment.
	ID *string `json:"id" validate:"required"`

	// The ID of the account that the assignment belongs to.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the template that the assignment is based on.
	TemplateID *string `json:"template_id" validate:"required"`

	// The version of the template that the assignment is based on.
	TemplateVersion *string `json:"template_version" validate:"required"`

	// The type of the entity that the assignment applies to.
	TargetType *string `json:"target_type" validate:"required"`

	// The ID of the entity that the assignment applies to.
	Target *string `json:"target" validate:"required"`

	// The operation that the assignment applies to (e.g. 'create', 'update', 'delete').
	Operation *string `json:"operation" validate:"required"`

	// The status of the assignment (e.g. 'pending', 'success', 'failure').
	Status *string `json:"status" validate:"required"`

	// The URL of the assignment resource.
	Href *string `json:"href" validate:"required"`

	// The date and time when the assignment was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The user or system that created the assignment.
	CreatedBy *string `json:"created_by" validate:"required"`

	// The date and time when the assignment was last updated.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The user or system that last updated the assignment.
	UpdatedBy *string `json:"updated_by" validate:"required"`
}

// UnmarshalTemplateCreateAssignmentResponse unmarshals an instance of TemplateCreateAssignmentResponse from the specified map of raw messages.
func UnmarshalTemplateCreateAssignmentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateCreateAssignmentResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_id", &obj.TemplateID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_version", &obj.TemplateVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operation", &obj.Operation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
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
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_by", &obj.UpdatedBy)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TemplateItem : TemplateItem struct
type TemplateItem struct {
	// The ID of the template.
	ID *string `json:"id" validate:"required"`

	// The name of the template.
	Name *string `json:"name" validate:"required"`

	// The description of the template.
	Description *string `json:"description,omitempty"`

	// The version of the template.
	Version *string `json:"version" validate:"required"`

	// The timestamp when the template was created.
	CreatedAt *string `json:"created_at" validate:"required"`

	// The ID of the user who created the template.
	CreatedByID *string `json:"created_by_id" validate:"required"`

	// The timestamp when the template was last modified.
	LastModifiedAt *string `json:"last_modified_at" validate:"required"`

	// The ID of the user who last modified the template.
	LastModifiedByID *string `json:"last_modified_by_id" validate:"required"`

	// The URL to access the template resource.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalTemplateItem unmarshals an instance of TemplateItem from the specified map of raw messages.
func UnmarshalTemplateItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplateItem)
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
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
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

// TemplatesListAssignmentResponse : Response object containing a list of template assignments.
type TemplatesListAssignmentResponse struct {
	// Total number of items matching the query.
	Total *int64 `json:"total,omitempty"`

	// Maximum number of items returned in the response.
	Limit *int64 `json:"limit,omitempty"`

	// Index of the first item returned in the response.
	Offset *int64 `json:"offset,omitempty"`

	// List of template assignments.
	GroupsAssignment []TemplateCreateAssignmentResponse `json:"groups_assignment,omitempty"`
}

// UnmarshalTemplatesListAssignmentResponse unmarshals an instance of TemplatesListAssignmentResponse from the specified map of raw messages.
func UnmarshalTemplatesListAssignmentResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TemplatesListAssignmentResponse)
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "groups_assignment", &obj.GroupsAssignment, UnmarshalTemplateCreateAssignmentResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAccessGroupOptions : The UpdateAccessGroup options.
type UpdateAccessGroupOptions struct {
	// The access group identifier.
	AccessGroupID *string `json:"access_group_id" validate:"required,ne="`

	// The current revision number of the group being updated. This can be found in the Create/Get access group response
	// ETag header.
	IfMatch *string `json:"If-Match" validate:"required"`

	// Give the access group a unique name that doesn't conflict with an existing access group in the account. This field
	// is case-insensitive and has a limit of 100 characters.
	Name *string `json:"name,omitempty"`

	// Assign an optional description for the access group. This field has a limit of 250 characters.
	Description *string `json:"description,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *UpdateAccessGroupOptions) SetAccessGroupID(accessGroupID string) *UpdateAccessGroupOptions {
	_options.AccessGroupID = core.StringPtr(accessGroupID)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateAccessGroupOptions) SetIfMatch(ifMatch string) *UpdateAccessGroupOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateAccessGroupOptions) SetName(name string) *UpdateAccessGroupOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateAccessGroupOptions) SetDescription(description string) *UpdateAccessGroupOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *UpdateAccessGroupOptions) SetTransactionID(transactionID string) *UpdateAccessGroupOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccessGroupOptions) SetHeaders(param map[string]string) *UpdateAccessGroupOptions {
	options.Headers = param
	return options
}

// UpdateAccountSettingsOptions : The UpdateAccountSettings options.
type UpdateAccountSettingsOptions struct {
	// Account ID of the API keys(s) to query. If a service IAM ID is specified in iam_id then account_id must match the
	// account of the IAM ID. If a user IAM ID is specified in iam_id then then account_id must match the account of the
	// Authorization token.
	AccountID *string `json:"account_id" validate:"required"`

	// This flag controls the public access feature within the account. It is set to true by default. Note: When this flag
	// is set to false, all policies within the account attached to the Public Access group will be deleted.
	PublicAccessEnabled *bool `json:"public_access_enabled,omitempty"`

	// An optional transaction ID can be passed to your request, which can be useful for tracking calls through multiple
	// services by using one identifier. The header key must be set to Transaction-Id and the value is anything that you
	// choose. If no transaction ID is passed in, then a random ID is generated.
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
func (_options *UpdateAccountSettingsOptions) SetAccountID(accountID string) *UpdateAccountSettingsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetPublicAccessEnabled : Allow user to set PublicAccessEnabled
func (_options *UpdateAccountSettingsOptions) SetPublicAccessEnabled(publicAccessEnabled bool) *UpdateAccountSettingsOptions {
	_options.PublicAccessEnabled = core.BoolPtr(publicAccessEnabled)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *UpdateAccountSettingsOptions) SetTransactionID(transactionID string) *UpdateAccountSettingsOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountSettingsOptions) SetHeaders(param map[string]string) *UpdateAccountSettingsOptions {
	options.Headers = param
	return options
}

// UpdateTemplateVersionOptions : The UpdateTemplateVersion options.
type UpdateTemplateVersionOptions struct {
	// ID of the template.
	TemplateID *string `json:"template_id" validate:"required,ne="`

	// Version number of the template.
	VersionNum *string `json:"version_num" validate:"required,ne="`

	// ETag value of the template version document.
	IfMatch *string `json:"If-Match" validate:"required"`

	// The ID of the access group template.
	ID *string `json:"id,omitempty"`

	// The name of the access group template.
	Name *string `json:"name,omitempty"`

	// The description of the access group template.
	Description *string `json:"description,omitempty"`

	// The ID of the account to which the access group template is assigned.
	AccountID *string `json:"account_id,omitempty"`

	// The version of the access group template.
	Version *string `json:"version,omitempty"`

	// A boolean indicating whether the access group template is committed.
	Committed *bool `json:"committed,omitempty"`

	// Access Group Input Component.
	AccessGroup *AccessGroupInput `json:"access_group,omitempty"`

	// References to policy templates assigned to the access group template.
	PolicyTemplates []PolicyTemplatesInput `json:"policy_templates,omitempty"`

	// The URL of the access group template resource.
	Href *string `json:"href,omitempty"`

	// The date and time when the access group template was created.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// The ID of the user who created the access group template.
	CreatedByID *string `json:"created_by_id,omitempty"`

	// The date and time when the access group template was last modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at,omitempty"`

	// The ID of the user who last modified the access group template.
	LastModifiedByID *string `json:"last_modified_by_id,omitempty"`

	// transaction id in header.
	TransactionID *string `json:"Transaction-Id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateTemplateVersionOptions : Instantiate UpdateTemplateVersionOptions
func (*IamAccessGroupsV2) NewUpdateTemplateVersionOptions(templateID string, versionNum string, ifMatch string) *UpdateTemplateVersionOptions {
	return &UpdateTemplateVersionOptions{
		TemplateID: core.StringPtr(templateID),
		VersionNum: core.StringPtr(versionNum),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetTemplateID : Allow user to set TemplateID
func (_options *UpdateTemplateVersionOptions) SetTemplateID(templateID string) *UpdateTemplateVersionOptions {
	_options.TemplateID = core.StringPtr(templateID)
	return _options
}

// SetVersionNum : Allow user to set VersionNum
func (_options *UpdateTemplateVersionOptions) SetVersionNum(versionNum string) *UpdateTemplateVersionOptions {
	_options.VersionNum = core.StringPtr(versionNum)
	return _options
}

// SetIfMatch : Allow user to set IfMatch
func (_options *UpdateTemplateVersionOptions) SetIfMatch(ifMatch string) *UpdateTemplateVersionOptions {
	_options.IfMatch = core.StringPtr(ifMatch)
	return _options
}

// SetID : Allow user to set ID
func (_options *UpdateTemplateVersionOptions) SetID(id string) *UpdateTemplateVersionOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateTemplateVersionOptions) SetName(name string) *UpdateTemplateVersionOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateTemplateVersionOptions) SetDescription(description string) *UpdateTemplateVersionOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *UpdateTemplateVersionOptions) SetAccountID(accountID string) *UpdateTemplateVersionOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetVersion : Allow user to set Version
func (_options *UpdateTemplateVersionOptions) SetVersion(version string) *UpdateTemplateVersionOptions {
	_options.Version = core.StringPtr(version)
	return _options
}

// SetCommitted : Allow user to set Committed
func (_options *UpdateTemplateVersionOptions) SetCommitted(committed bool) *UpdateTemplateVersionOptions {
	_options.Committed = core.BoolPtr(committed)
	return _options
}

// SetAccessGroup : Allow user to set AccessGroup
func (_options *UpdateTemplateVersionOptions) SetAccessGroup(accessGroup *AccessGroupInput) *UpdateTemplateVersionOptions {
	_options.AccessGroup = accessGroup
	return _options
}

// SetPolicyTemplates : Allow user to set PolicyTemplates
func (_options *UpdateTemplateVersionOptions) SetPolicyTemplates(policyTemplates []PolicyTemplatesInput) *UpdateTemplateVersionOptions {
	_options.PolicyTemplates = policyTemplates
	return _options
}

// SetHref : Allow user to set Href
func (_options *UpdateTemplateVersionOptions) SetHref(href string) *UpdateTemplateVersionOptions {
	_options.Href = core.StringPtr(href)
	return _options
}

// SetCreatedAt : Allow user to set CreatedAt
func (_options *UpdateTemplateVersionOptions) SetCreatedAt(createdAt *strfmt.DateTime) *UpdateTemplateVersionOptions {
	_options.CreatedAt = createdAt
	return _options
}

// SetCreatedByID : Allow user to set CreatedByID
func (_options *UpdateTemplateVersionOptions) SetCreatedByID(createdByID string) *UpdateTemplateVersionOptions {
	_options.CreatedByID = core.StringPtr(createdByID)
	return _options
}

// SetLastModifiedAt : Allow user to set LastModifiedAt
func (_options *UpdateTemplateVersionOptions) SetLastModifiedAt(lastModifiedAt *strfmt.DateTime) *UpdateTemplateVersionOptions {
	_options.LastModifiedAt = lastModifiedAt
	return _options
}

// SetLastModifiedByID : Allow user to set LastModifiedByID
func (_options *UpdateTemplateVersionOptions) SetLastModifiedByID(lastModifiedByID string) *UpdateTemplateVersionOptions {
	_options.LastModifiedByID = core.StringPtr(lastModifiedByID)
	return _options
}

// SetTransactionID : Allow user to set TransactionID
func (_options *UpdateTemplateVersionOptions) SetTransactionID(transactionID string) *UpdateTemplateVersionOptions {
	_options.TransactionID = core.StringPtr(transactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTemplateVersionOptions) SetHeaders(param map[string]string) *UpdateTemplateVersionOptions {
	options.Headers = param
	return options
}

//
// AccessGroupsPager can be used to simplify the use of the "ListAccessGroups" method.
//
type AccessGroupsPager struct {
	hasNext bool
	options *ListAccessGroupsOptions
	client  *IamAccessGroupsV2
	pageContext struct {
		next *int64
	}
}

// NewAccessGroupsPager returns a new AccessGroupsPager instance.
func (iamAccessGroups *IamAccessGroupsV2) NewAccessGroupsPager(options *ListAccessGroupsOptions) (pager *AccessGroupsPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListAccessGroupsOptions = *options
	pager = &AccessGroupsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  iamAccessGroups,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AccessGroupsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AccessGroupsPager) GetNextWithContext(ctx context.Context) (page []Group, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListAccessGroupsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Groups

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AccessGroupsPager) GetAllWithContext(ctx context.Context) (allItems []Group, err error) {
	for pager.HasNext() {
		var nextPage []Group
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AccessGroupsPager) GetNext() (page []Group, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AccessGroupsPager) GetAll() (allItems []Group, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// AccessGroupMembersPager can be used to simplify the use of the "ListAccessGroupMembers" method.
//
type AccessGroupMembersPager struct {
	hasNext bool
	options *ListAccessGroupMembersOptions
	client  *IamAccessGroupsV2
	pageContext struct {
		next *int64
	}
}

// NewAccessGroupMembersPager returns a new AccessGroupMembersPager instance.
func (iamAccessGroups *IamAccessGroupsV2) NewAccessGroupMembersPager(options *ListAccessGroupMembersOptions) (pager *AccessGroupMembersPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListAccessGroupMembersOptions = *options
	pager = &AccessGroupMembersPager{
		hasNext: true,
		options: &optionsCopy,
		client:  iamAccessGroups,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *AccessGroupMembersPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *AccessGroupMembersPager) GetNextWithContext(ctx context.Context) (page []ListGroupMembersResponseMember, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListAccessGroupMembersWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Members

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *AccessGroupMembersPager) GetAllWithContext(ctx context.Context) (allItems []ListGroupMembersResponseMember, err error) {
	for pager.HasNext() {
		var nextPage []ListGroupMembersResponseMember
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *AccessGroupMembersPager) GetNext() (page []ListGroupMembersResponseMember, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *AccessGroupMembersPager) GetAll() (allItems []ListGroupMembersResponseMember, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// TemplatesPager can be used to simplify the use of the "ListTemplates" method.
//
type TemplatesPager struct {
	hasNext bool
	options *ListTemplatesOptions
	client  *IamAccessGroupsV2
	pageContext struct {
		next *int64
	}
}

// NewTemplatesPager returns a new TemplatesPager instance.
func (iamAccessGroups *IamAccessGroupsV2) NewTemplatesPager(options *ListTemplatesOptions) (pager *TemplatesPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListTemplatesOptions = *options
	pager = &TemplatesPager{
		hasNext: true,
		options: &optionsCopy,
		client:  iamAccessGroups,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *TemplatesPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *TemplatesPager) GetNextWithContext(ctx context.Context) (page []TemplateItem, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListTemplatesWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.GroupsTemplates

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *TemplatesPager) GetAllWithContext(ctx context.Context) (allItems []TemplateItem, err error) {
	for pager.HasNext() {
		var nextPage []TemplateItem
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *TemplatesPager) GetNext() (page []TemplateItem, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *TemplatesPager) GetAll() (allItems []TemplateItem, err error) {
	return pager.GetAllWithContext(context.Background())
}

//
// TemplateVersionsPager can be used to simplify the use of the "ListTemplateVersions" method.
//
type TemplateVersionsPager struct {
	hasNext bool
	options *ListTemplateVersionsOptions
	client  *IamAccessGroupsV2
	pageContext struct {
		next *int64
	}
}

// NewTemplateVersionsPager returns a new TemplateVersionsPager instance.
func (iamAccessGroups *IamAccessGroupsV2) NewTemplateVersionsPager(options *ListTemplateVersionsOptions) (pager *TemplateVersionsPager, err error) {
	if options.Offset != nil && *options.Offset != 0 {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy ListTemplateVersionsOptions = *options
	pager = &TemplateVersionsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  iamAccessGroups,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *TemplateVersionsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *TemplateVersionsPager) GetNextWithContext(ctx context.Context) (page []ListTemplatesVersionsResponse, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.ListTemplateVersionsWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *int64
	if result.Next != nil {
		var offset *int64
		offset, err = core.GetQueryParamAsInt(result.Next.Href, "offset")
		if err != nil {
			err = fmt.Errorf("error retrieving 'offset' query parameter from URL '%s': %s", *result.Next.Href, err.Error())
			return
		}
		next = offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Versions

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *TemplateVersionsPager) GetAllWithContext(ctx context.Context) (allItems []ListTemplatesVersionsResponse, err error) {
	for pager.HasNext() {
		var nextPage []ListTemplatesVersionsResponse
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *TemplateVersionsPager) GetNext() (page []ListTemplatesVersionsResponse, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *TemplateVersionsPager) GetAll() (allItems []ListTemplatesVersionsResponse, err error) {
	return pager.GetAllWithContext(context.Background())
}
