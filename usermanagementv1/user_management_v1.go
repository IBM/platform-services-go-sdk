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

// Package usermanagementv1 : Operations and models for the UserManagementV1 service
package usermanagementv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// UserManagementV1 : Manage lifecycle of your cloud users using User Management APIs.
//
// Version: 1.0
type UserManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://user-management.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "user_management"

// UserManagementV1Options : Service options
type UserManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUserManagementV1UsingExternalConfig : constructs an instance of UserManagementV1 with passed in options and external configuration.
func NewUserManagementV1UsingExternalConfig(options *UserManagementV1Options) (userManagement *UserManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	userManagement, err = NewUserManagementV1(options)
	if err != nil {
		return
	}

	err = userManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = userManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUserManagementV1 : constructs an instance of UserManagementV1 with passed in options.
func NewUserManagementV1(options *UserManagementV1Options) (service *UserManagementV1, err error) {
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

	service = &UserManagementV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (userManagement *UserManagementV1) SetServiceURL(url string) error {
	return userManagement.Service.SetServiceURL(url)
}

// GetUserLinkages : Get user linkages
// Retrieve a user's linkages by user's iam id.
func (userManagement *UserManagementV1) GetUserLinkages(getUserLinkagesOptions *GetUserLinkagesOptions) (result *UserLinkages, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserLinkagesOptions, "getUserLinkagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserLinkagesOptions, "getUserLinkagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users", "linkages"}
	pathParameters := []string{*getUserLinkagesOptions.AccountID, *getUserLinkagesOptions.IamID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserLinkagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetUserLinkages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserLinkages(m)
		response.Result = result
	}

	return
}

// CreateUserLinkages : create user linkages
// create a linakge for user by user's iam id.It's a system operation, only with System role/policy can invoke this api.
func (userManagement *UserManagementV1) CreateUserLinkages(createUserLinkagesOptions *CreateUserLinkagesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createUserLinkagesOptions, "createUserLinkagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createUserLinkagesOptions, "createUserLinkagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users", "linkages", ""}
	pathParameters := []string{*createUserLinkagesOptions.AccountID, *createUserLinkagesOptions.IamID, *createUserLinkagesOptions.Origin, *createUserLinkagesOptions.IdFromOrigin}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createUserLinkagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "CreateUserLinkages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// RemoveUserLinkages : remove a user linkages
// remove a user's linkage by user's iam id.It's a system operation, only with System role/policy can invoke this api.
func (userManagement *UserManagementV1) RemoveUserLinkages(removeUserLinkagesOptions *RemoveUserLinkagesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeUserLinkagesOptions, "removeUserLinkagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeUserLinkagesOptions, "removeUserLinkagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users", "linkages", ""}
	pathParameters := []string{*removeUserLinkagesOptions.AccountID, *removeUserLinkagesOptions.IamID, *removeUserLinkagesOptions.Origin, *removeUserLinkagesOptions.IdFromOrigin}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeUserLinkagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "RemoveUserLinkages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// GetUserProfile : Get user profile
// Retrieve a user profile by user's iam id or cloudant guid.
func (userManagement *UserManagementV1) GetUserProfile(getUserProfileOptions *GetUserProfileOptions) (result *UserProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserProfileOptions, "getUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserProfileOptions, "getUserProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*getUserProfileOptions.AccountID, *getUserProfileOptions.IamID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getUserProfileOptions.IncludeLinkages != nil {
		builder.AddQuery("include_linkages", fmt.Sprint(*getUserProfileOptions.IncludeLinkages))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserProfile(m)
		response.Result = result
	}

	return
}

// CreateUserProfile : create or replace user profile
// Create a new user or replace user if user already exist by user's iam id.We enforce schema validation, some fields
// are required.Only allow System to call to create user and sync user as a whole object. User update need to use
// Partial update user profile.
func (userManagement *UserManagementV1) CreateUserProfile(createUserProfileOptions *CreateUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createUserProfileOptions, "createUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createUserProfileOptions, "createUserProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*createUserProfileOptions.AccountID, *createUserProfileOptions.IamID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "CreateUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createUserProfileOptions.Realm != nil {
		body["realm"] = createUserProfileOptions.Realm
	}
	if createUserProfileOptions.UserID != nil {
		body["user_id"] = createUserProfileOptions.UserID
	}
	if createUserProfileOptions.Firstname != nil {
		body["firstname"] = createUserProfileOptions.Firstname
	}
	if createUserProfileOptions.Lastname != nil {
		body["lastname"] = createUserProfileOptions.Lastname
	}
	if createUserProfileOptions.State != nil {
		body["state"] = createUserProfileOptions.State
	}
	if createUserProfileOptions.Email != nil {
		body["email"] = createUserProfileOptions.Email
	}
	if createUserProfileOptions.Phonenumber != nil {
		body["phonenumber"] = createUserProfileOptions.Phonenumber
	}
	if createUserProfileOptions.Altphonenumber != nil {
		body["altphonenumber"] = createUserProfileOptions.Altphonenumber
	}
	if createUserProfileOptions.Photo != nil {
		body["photo"] = createUserProfileOptions.Photo
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// UpdateUserProfile : partial update user profile
// Partial update a user's profile by user's iam id.We enforce schema validations.User can disable/activate another
// user, as long as the user has user-management access, but user can not change state to "PROCESSING" or "PENDING",
// which are system states.
func (userManagement *UserManagementV1) UpdateUserProfile(updateUserProfileOptions *UpdateUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserProfileOptions, "updateUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserProfileOptions, "updateUserProfileOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*updateUserProfileOptions.AccountID, *updateUserProfileOptions.IamID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "UpdateUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateUserProfileOptions.UserID != nil {
		body["user_id"] = updateUserProfileOptions.UserID
	}
	if updateUserProfileOptions.Firstname != nil {
		body["firstname"] = updateUserProfileOptions.Firstname
	}
	if updateUserProfileOptions.Lastname != nil {
		body["lastname"] = updateUserProfileOptions.Lastname
	}
	if updateUserProfileOptions.State != nil {
		body["state"] = updateUserProfileOptions.State
	}
	if updateUserProfileOptions.Email != nil {
		body["email"] = updateUserProfileOptions.Email
	}
	if updateUserProfileOptions.Phonenumber != nil {
		body["phonenumber"] = updateUserProfileOptions.Phonenumber
	}
	if updateUserProfileOptions.Altphonenumber != nil {
		body["altphonenumber"] = updateUserProfileOptions.Altphonenumber
	}
	if updateUserProfileOptions.Photo != nil {
		body["photo"] = updateUserProfileOptions.Photo
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// GetUserSettings : Get user settings
// Retrieve a user's settings by user's iam id.
func (userManagement *UserManagementV1) GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions) (result *UserSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserSettingsOptions, "getUserSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserSettingsOptions, "getUserSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users", "settings"}
	pathParameters := []string{*getUserSettingsOptions.AccountID, *getUserSettingsOptions.IamID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetUserSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserSettings(m)
		response.Result = result
	}

	return
}

// UpdateUserSettings : Partial update user settings
// Update a user's settings by user's iam id.User can update "language", "notification_language" and can update
// "allowed_ip_addresses" if "self_manage" is true, but user can not update "allowed_ip_addresses" if "self_manage" is
// false.And Update "self_manage" requires user-management policy.
func (userManagement *UserManagementV1) UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserSettingsOptions, "updateUserSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserSettingsOptions, "updateUserSettingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users", "settings"}
	pathParameters := []string{*updateUserSettingsOptions.AccountID, *updateUserSettingsOptions.IamID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "UpdateUserSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateUserSettingsOptions.Language != nil {
		body["language"] = updateUserSettingsOptions.Language
	}
	if updateUserSettingsOptions.NotificationLanguage != nil {
		body["notification_language"] = updateUserSettingsOptions.NotificationLanguage
	}
	if updateUserSettingsOptions.AllowedIpAddresses != nil {
		body["allowed_ip_addresses"] = updateUserSettingsOptions.AllowedIpAddresses
	}
	if updateUserSettingsOptions.SelfManage != nil {
		body["self_manage"] = updateUserSettingsOptions.SelfManage
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// RemoveUserFromAccount : remove user from account
// IAM user management policy is required to perform this action.If the caller does not have proper IAM user management
// policy, then if the user is a decendent of the caller in IMS heirarchy, then allow as well.Do not support self
// delete.
func (userManagement *UserManagementV1) RemoveUserFromAccount(removeUserFromAccountOptions *RemoveUserFromAccountOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(removeUserFromAccountOptions, "removeUserFromAccountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(removeUserFromAccountOptions, "removeUserFromAccountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*removeUserFromAccountOptions.AccountID, *removeUserFromAccountOptions.IamID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range removeUserFromAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "RemoveUserFromAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// ListUsers : Get users
// Retrieve users in the account.If team directory enabled, return all users in the account.If team directory disbaled,
// and user has IAM viewer role on user-management service, then return all users in the account.If team directory
// disabled, and user does not have IAM viewer role on user-management service, then return only current user.
func (userManagement *UserManagementV1) ListUsers(listUsersOptions *ListUsersOptions) (result *UserList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listUsersOptions, "listUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listUsersOptions, "listUsersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*listUsersOptions.AccountID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "ListUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listUsersOptions.IAMid != nil {
		builder.AddQuery("IAMid", fmt.Sprint(*listUsersOptions.IAMid))
	}
	if listUsersOptions.Firstname != nil {
		builder.AddQuery("firstname", fmt.Sprint(*listUsersOptions.Firstname))
	}
	if listUsersOptions.Lastname != nil {
		builder.AddQuery("lastname", fmt.Sprint(*listUsersOptions.Lastname))
	}
	if listUsersOptions.Email != nil {
		builder.AddQuery("email", fmt.Sprint(*listUsersOptions.Email))
	}
	if listUsersOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listUsersOptions.State))
	}
	if listUsersOptions.Realm != nil {
		builder.AddQuery("realm", fmt.Sprint(*listUsersOptions.Realm))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserList(m)
		response.Result = result
	}

	return
}

// InviteUsers : Invite users
// Invite users to the account.
func (userManagement *UserManagementV1) InviteUsers(inviteUsersOptions *InviteUsersOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(inviteUsersOptions, "inviteUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(inviteUsersOptions, "inviteUsersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "users"}
	pathParameters := []string{*inviteUsersOptions.AccountID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range inviteUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "InviteUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if inviteUsersOptions.Users != nil {
		body["users"] = inviteUsersOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, nil)

	return
}

// GetImsUsers : Get users in account and filtering results based on IMS user heirarchy
// Retrieve a user profile by user's iam id or cloudant guid.
func (userManagement *UserManagementV1) GetImsUsers(getImsUsersOptions *GetImsUsersOptions) (result *UserList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getImsUsersOptions, "getImsUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getImsUsersOptions, "getImsUsersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "ims/users"}
	pathParameters := []string{*getImsUsersOptions.AccountID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getImsUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetImsUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getImsUsersOptions.IAMid != nil {
		builder.AddQuery("IAMid", fmt.Sprint(*getImsUsersOptions.IAMid))
	}
	if getImsUsersOptions.Firstname != nil {
		builder.AddQuery("firstname", fmt.Sprint(*getImsUsersOptions.Firstname))
	}
	if getImsUsersOptions.Lastname != nil {
		builder.AddQuery("lastname", fmt.Sprint(*getImsUsersOptions.Lastname))
	}
	if getImsUsersOptions.Email != nil {
		builder.AddQuery("email", fmt.Sprint(*getImsUsersOptions.Email))
	}
	if getImsUsersOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*getImsUsersOptions.State))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserList(m)
		response.Result = result
	}

	return
}

// GetCfUsers : Get users
// Get CF organizations Users in account organization.
func (userManagement *UserManagementV1) GetCfUsers(getCfUsersOptions *GetCfUsersOptions) (result *UserList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCfUsersOptions, "getCfUsersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCfUsersOptions, "getCfUsersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/accounts", "organizations", "users"}
	pathParameters := []string{*getCfUsersOptions.AccountID, *getCfUsersOptions.OrganizationGuid}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(userManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCfUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("user_management", "V1", "GetCfUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = userManagement.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalUserList(m)
		response.Result = result
	}

	return
}

// CreateUserLinkagesOptions : The CreateUserLinkages options.
type CreateUserLinkagesOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// origin is "IMS" OR "UAA".
	Origin *string `json:"origin" validate:"required"`

	// An alpha-numeric value identifying the origin.
	IdFromOrigin *string `json:"id_from_origin" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateUserLinkagesOptions : Instantiate CreateUserLinkagesOptions
func (*UserManagementV1) NewCreateUserLinkagesOptions(accountID string, iamID string, origin string, idFromOrigin string) *CreateUserLinkagesOptions {
	return &CreateUserLinkagesOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
		Origin: core.StringPtr(origin),
		IdFromOrigin: core.StringPtr(idFromOrigin),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateUserLinkagesOptions) SetAccountID(accountID string) *CreateUserLinkagesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *CreateUserLinkagesOptions) SetIamID(iamID string) *CreateUserLinkagesOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetOrigin : Allow user to set Origin
func (options *CreateUserLinkagesOptions) SetOrigin(origin string) *CreateUserLinkagesOptions {
	options.Origin = core.StringPtr(origin)
	return options
}

// SetIdFromOrigin : Allow user to set IdFromOrigin
func (options *CreateUserLinkagesOptions) SetIdFromOrigin(idFromOrigin string) *CreateUserLinkagesOptions {
	options.IdFromOrigin = core.StringPtr(idFromOrigin)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateUserLinkagesOptions) SetHeaders(param map[string]string) *CreateUserLinkagesOptions {
	options.Headers = param
	return options
}

// CreateUserProfileOptions : The CreateUserProfile options.
type CreateUserProfileOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// The real of the user, only for new user, this field can not be updated.
	Realm *string `json:"realm,omitempty"`

	// The user id of the user.
	UserID *string `json:"user_id,omitempty"`

	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`

	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`

	// The state of the user,Possible values "PROCESSING" | "PENDING" | "ACTIVE" | "DISABLED" | "VPN_ONLY".
	State *string `json:"state,omitempty"`

	// The email of the user.
	Email *string `json:"email,omitempty"`

	// The phone number of the user.
	Phonenumber *string `json:"phonenumber,omitempty"`

	// The altphonenumber of the user, new field to add (optional).
	Altphonenumber *string `json:"altphonenumber,omitempty"`

	// The phone link of the user.
	Photo *string `json:"photo,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateUserProfileOptions : Instantiate CreateUserProfileOptions
func (*UserManagementV1) NewCreateUserProfileOptions(accountID string, iamID string) *CreateUserProfileOptions {
	return &CreateUserProfileOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateUserProfileOptions) SetAccountID(accountID string) *CreateUserProfileOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *CreateUserProfileOptions) SetIamID(iamID string) *CreateUserProfileOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetRealm : Allow user to set Realm
func (options *CreateUserProfileOptions) SetRealm(realm string) *CreateUserProfileOptions {
	options.Realm = core.StringPtr(realm)
	return options
}

// SetUserID : Allow user to set UserID
func (options *CreateUserProfileOptions) SetUserID(userID string) *CreateUserProfileOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetFirstname : Allow user to set Firstname
func (options *CreateUserProfileOptions) SetFirstname(firstname string) *CreateUserProfileOptions {
	options.Firstname = core.StringPtr(firstname)
	return options
}

// SetLastname : Allow user to set Lastname
func (options *CreateUserProfileOptions) SetLastname(lastname string) *CreateUserProfileOptions {
	options.Lastname = core.StringPtr(lastname)
	return options
}

// SetState : Allow user to set State
func (options *CreateUserProfileOptions) SetState(state string) *CreateUserProfileOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetEmail : Allow user to set Email
func (options *CreateUserProfileOptions) SetEmail(email string) *CreateUserProfileOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetPhonenumber : Allow user to set Phonenumber
func (options *CreateUserProfileOptions) SetPhonenumber(phonenumber string) *CreateUserProfileOptions {
	options.Phonenumber = core.StringPtr(phonenumber)
	return options
}

// SetAltphonenumber : Allow user to set Altphonenumber
func (options *CreateUserProfileOptions) SetAltphonenumber(altphonenumber string) *CreateUserProfileOptions {
	options.Altphonenumber = core.StringPtr(altphonenumber)
	return options
}

// SetPhoto : Allow user to set Photo
func (options *CreateUserProfileOptions) SetPhoto(photo string) *CreateUserProfileOptions {
	options.Photo = core.StringPtr(photo)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateUserProfileOptions) SetHeaders(param map[string]string) *CreateUserProfileOptions {
	options.Headers = param
	return options
}

// GetCfUsersOptions : The GetCfUsers options.
type GetCfUsersOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The organization id.
	OrganizationGuid *string `json:"organization_guid" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCfUsersOptions : Instantiate GetCfUsersOptions
func (*UserManagementV1) NewGetCfUsersOptions(accountID string, organizationGuid string) *GetCfUsersOptions {
	return &GetCfUsersOptions{
		AccountID: core.StringPtr(accountID),
		OrganizationGuid: core.StringPtr(organizationGuid),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetCfUsersOptions) SetAccountID(accountID string) *GetCfUsersOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetOrganizationGuid : Allow user to set OrganizationGuid
func (options *GetCfUsersOptions) SetOrganizationGuid(organizationGuid string) *GetCfUsersOptions {
	options.OrganizationGuid = core.StringPtr(organizationGuid)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCfUsersOptions) SetHeaders(param map[string]string) *GetCfUsersOptions {
	options.Headers = param
	return options
}

// GetImsUsersOptions : The GetImsUsers options.
type GetImsUsersOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The realm of the user.
	IAMid *string `json:"IAMid,omitempty"`

	// The firstname of user.
	Firstname *string `json:"firstname,omitempty"`

	// The lastname of user.
	Lastname *string `json:"lastname,omitempty"`

	// The email of user.
	Email *string `json:"email,omitempty"`

	// The state.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetImsUsersOptions : Instantiate GetImsUsersOptions
func (*UserManagementV1) NewGetImsUsersOptions(accountID string) *GetImsUsersOptions {
	return &GetImsUsersOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetImsUsersOptions) SetAccountID(accountID string) *GetImsUsersOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIAMid : Allow user to set IAMid
func (options *GetImsUsersOptions) SetIAMid(iAMid string) *GetImsUsersOptions {
	options.IAMid = core.StringPtr(iAMid)
	return options
}

// SetFirstname : Allow user to set Firstname
func (options *GetImsUsersOptions) SetFirstname(firstname string) *GetImsUsersOptions {
	options.Firstname = core.StringPtr(firstname)
	return options
}

// SetLastname : Allow user to set Lastname
func (options *GetImsUsersOptions) SetLastname(lastname string) *GetImsUsersOptions {
	options.Lastname = core.StringPtr(lastname)
	return options
}

// SetEmail : Allow user to set Email
func (options *GetImsUsersOptions) SetEmail(email string) *GetImsUsersOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetState : Allow user to set State
func (options *GetImsUsersOptions) SetState(state string) *GetImsUsersOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetImsUsersOptions) SetHeaders(param map[string]string) *GetImsUsersOptions {
	options.Headers = param
	return options
}

// GetUserLinkagesOptions : The GetUserLinkages options.
type GetUserLinkagesOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserLinkagesOptions : Instantiate GetUserLinkagesOptions
func (*UserManagementV1) NewGetUserLinkagesOptions(accountID string, iamID string) *GetUserLinkagesOptions {
	return &GetUserLinkagesOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetUserLinkagesOptions) SetAccountID(accountID string) *GetUserLinkagesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *GetUserLinkagesOptions) SetIamID(iamID string) *GetUserLinkagesOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserLinkagesOptions) SetHeaders(param map[string]string) *GetUserLinkagesOptions {
	options.Headers = param
	return options
}

// GetUserProfileOptions : The GetUserProfile options.
type GetUserProfileOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// Indicate include linkages.
	IncludeLinkages *bool `json:"include_linkages,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserProfileOptions : Instantiate GetUserProfileOptions
func (*UserManagementV1) NewGetUserProfileOptions(accountID string, iamID string) *GetUserProfileOptions {
	return &GetUserProfileOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetUserProfileOptions) SetAccountID(accountID string) *GetUserProfileOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *GetUserProfileOptions) SetIamID(iamID string) *GetUserProfileOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetIncludeLinkages : Allow user to set IncludeLinkages
func (options *GetUserProfileOptions) SetIncludeLinkages(includeLinkages bool) *GetUserProfileOptions {
	options.IncludeLinkages = core.BoolPtr(includeLinkages)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserProfileOptions) SetHeaders(param map[string]string) *GetUserProfileOptions {
	options.Headers = param
	return options
}

// GetUserSettingsOptions : The GetUserSettings options.
type GetUserSettingsOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserSettingsOptions : Instantiate GetUserSettingsOptions
func (*UserManagementV1) NewGetUserSettingsOptions(accountID string, iamID string) *GetUserSettingsOptions {
	return &GetUserSettingsOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetUserSettingsOptions) SetAccountID(accountID string) *GetUserSettingsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *GetUserSettingsOptions) SetIamID(iamID string) *GetUserSettingsOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserSettingsOptions) SetHeaders(param map[string]string) *GetUserSettingsOptions {
	options.Headers = param
	return options
}

// InviteUsersOptions : The InviteUsers options.
type InviteUsersOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// list of users to be invited.
	Users []InviteUser `json:"users,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInviteUsersOptions : Instantiate InviteUsersOptions
func (*UserManagementV1) NewInviteUsersOptions(accountID string) *InviteUsersOptions {
	return &InviteUsersOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *InviteUsersOptions) SetAccountID(accountID string) *InviteUsersOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetUsers : Allow user to set Users
func (options *InviteUsersOptions) SetUsers(users []InviteUser) *InviteUsersOptions {
	options.Users = users
	return options
}

// SetHeaders : Allow user to set Headers
func (options *InviteUsersOptions) SetHeaders(param map[string]string) *InviteUsersOptions {
	options.Headers = param
	return options
}

// ListUsersOptions : The ListUsers options.
type ListUsersOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The realm of the user.
	IAMid *string `json:"IAMid,omitempty"`

	// The firstname of user.
	Firstname *string `json:"firstname,omitempty"`

	// The lastname of user.
	Lastname *string `json:"lastname,omitempty"`

	// The email of user.
	Email *string `json:"email,omitempty"`

	// The state.
	State *string `json:"state,omitempty"`

	// The realm of the user.
	Realm *string `json:"realm,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListUsersOptions : Instantiate ListUsersOptions
func (*UserManagementV1) NewListUsersOptions(accountID string) *ListUsersOptions {
	return &ListUsersOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListUsersOptions) SetAccountID(accountID string) *ListUsersOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIAMid : Allow user to set IAMid
func (options *ListUsersOptions) SetIAMid(iAMid string) *ListUsersOptions {
	options.IAMid = core.StringPtr(iAMid)
	return options
}

// SetFirstname : Allow user to set Firstname
func (options *ListUsersOptions) SetFirstname(firstname string) *ListUsersOptions {
	options.Firstname = core.StringPtr(firstname)
	return options
}

// SetLastname : Allow user to set Lastname
func (options *ListUsersOptions) SetLastname(lastname string) *ListUsersOptions {
	options.Lastname = core.StringPtr(lastname)
	return options
}

// SetEmail : Allow user to set Email
func (options *ListUsersOptions) SetEmail(email string) *ListUsersOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetState : Allow user to set State
func (options *ListUsersOptions) SetState(state string) *ListUsersOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetRealm : Allow user to set Realm
func (options *ListUsersOptions) SetRealm(realm string) *ListUsersOptions {
	options.Realm = core.StringPtr(realm)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListUsersOptions) SetHeaders(param map[string]string) *ListUsersOptions {
	options.Headers = param
	return options
}

// RemoveUserFromAccountOptions : The RemoveUserFromAccount options.
type RemoveUserFromAccountOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveUserFromAccountOptions : Instantiate RemoveUserFromAccountOptions
func (*UserManagementV1) NewRemoveUserFromAccountOptions(accountID string, iamID string) *RemoveUserFromAccountOptions {
	return &RemoveUserFromAccountOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *RemoveUserFromAccountOptions) SetAccountID(accountID string) *RemoveUserFromAccountOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *RemoveUserFromAccountOptions) SetIamID(iamID string) *RemoveUserFromAccountOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveUserFromAccountOptions) SetHeaders(param map[string]string) *RemoveUserFromAccountOptions {
	options.Headers = param
	return options
}

// RemoveUserLinkagesOptions : The RemoveUserLinkages options.
type RemoveUserLinkagesOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// origin is "IMS" OR "UAA".
	Origin *string `json:"origin" validate:"required"`

	// An alpha-numeric value identifying the origin.
	IdFromOrigin *string `json:"id_from_origin" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRemoveUserLinkagesOptions : Instantiate RemoveUserLinkagesOptions
func (*UserManagementV1) NewRemoveUserLinkagesOptions(accountID string, iamID string, origin string, idFromOrigin string) *RemoveUserLinkagesOptions {
	return &RemoveUserLinkagesOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
		Origin: core.StringPtr(origin),
		IdFromOrigin: core.StringPtr(idFromOrigin),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *RemoveUserLinkagesOptions) SetAccountID(accountID string) *RemoveUserLinkagesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *RemoveUserLinkagesOptions) SetIamID(iamID string) *RemoveUserLinkagesOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetOrigin : Allow user to set Origin
func (options *RemoveUserLinkagesOptions) SetOrigin(origin string) *RemoveUserLinkagesOptions {
	options.Origin = core.StringPtr(origin)
	return options
}

// SetIdFromOrigin : Allow user to set IdFromOrigin
func (options *RemoveUserLinkagesOptions) SetIdFromOrigin(idFromOrigin string) *RemoveUserLinkagesOptions {
	options.IdFromOrigin = core.StringPtr(idFromOrigin)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RemoveUserLinkagesOptions) SetHeaders(param map[string]string) *RemoveUserLinkagesOptions {
	options.Headers = param
	return options
}

// UpdateUserProfileOptions : The UpdateUserProfile options.
type UpdateUserProfileOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// The user id of the user.
	UserID *string `json:"user_id,omitempty"`

	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`

	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`

	// The state of the user,Possible values "PROCESSING" | "PENDING" | "ACTIVE" | "DISABLED" | "VPN_ONLY".
	State *string `json:"state,omitempty"`

	// The email of the user.
	Email *string `json:"email,omitempty"`

	// The phone number of the user.
	Phonenumber *string `json:"phonenumber,omitempty"`

	// The altphonenumber of the user, new field to add (optional).
	Altphonenumber *string `json:"altphonenumber,omitempty"`

	// The phone link of the user.
	Photo *string `json:"photo,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserProfileOptions : Instantiate UpdateUserProfileOptions
func (*UserManagementV1) NewUpdateUserProfileOptions(accountID string, iamID string) *UpdateUserProfileOptions {
	return &UpdateUserProfileOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateUserProfileOptions) SetAccountID(accountID string) *UpdateUserProfileOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *UpdateUserProfileOptions) SetIamID(iamID string) *UpdateUserProfileOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetUserID : Allow user to set UserID
func (options *UpdateUserProfileOptions) SetUserID(userID string) *UpdateUserProfileOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetFirstname : Allow user to set Firstname
func (options *UpdateUserProfileOptions) SetFirstname(firstname string) *UpdateUserProfileOptions {
	options.Firstname = core.StringPtr(firstname)
	return options
}

// SetLastname : Allow user to set Lastname
func (options *UpdateUserProfileOptions) SetLastname(lastname string) *UpdateUserProfileOptions {
	options.Lastname = core.StringPtr(lastname)
	return options
}

// SetState : Allow user to set State
func (options *UpdateUserProfileOptions) SetState(state string) *UpdateUserProfileOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetEmail : Allow user to set Email
func (options *UpdateUserProfileOptions) SetEmail(email string) *UpdateUserProfileOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetPhonenumber : Allow user to set Phonenumber
func (options *UpdateUserProfileOptions) SetPhonenumber(phonenumber string) *UpdateUserProfileOptions {
	options.Phonenumber = core.StringPtr(phonenumber)
	return options
}

// SetAltphonenumber : Allow user to set Altphonenumber
func (options *UpdateUserProfileOptions) SetAltphonenumber(altphonenumber string) *UpdateUserProfileOptions {
	options.Altphonenumber = core.StringPtr(altphonenumber)
	return options
}

// SetPhoto : Allow user to set Photo
func (options *UpdateUserProfileOptions) SetPhoto(photo string) *UpdateUserProfileOptions {
	options.Photo = core.StringPtr(photo)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserProfileOptions) SetHeaders(param map[string]string) *UpdateUserProfileOptions {
	options.Headers = param
	return options
}

// UpdateUserSettingsOptions : The UpdateUserSettings options.
type UpdateUserSettingsOptions struct {
	// The account id.
	AccountID *string `json:"account_id" validate:"required"`

	// The user's iam id.
	IamID *string `json:"iam_id" validate:"required"`

	// UI language, default value empty.
	Language *string `json:"language,omitempty"`

	// For email, phone notification, default value empty.
	NotificationLanguage *string `json:"notification_language,omitempty"`

	// Ip address string use comma to separate string.
	AllowedIpAddresses *string `json:"allowed_ip_addresses,omitempty"`

	// a field set for user be able to self manage or not, default false.
	SelfManage *bool `json:"self_manage,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserSettingsOptions : Instantiate UpdateUserSettingsOptions
func (*UserManagementV1) NewUpdateUserSettingsOptions(accountID string, iamID string) *UpdateUserSettingsOptions {
	return &UpdateUserSettingsOptions{
		AccountID: core.StringPtr(accountID),
		IamID: core.StringPtr(iamID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateUserSettingsOptions) SetAccountID(accountID string) *UpdateUserSettingsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIamID : Allow user to set IamID
func (options *UpdateUserSettingsOptions) SetIamID(iamID string) *UpdateUserSettingsOptions {
	options.IamID = core.StringPtr(iamID)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateUserSettingsOptions) SetLanguage(language string) *UpdateUserSettingsOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetNotificationLanguage : Allow user to set NotificationLanguage
func (options *UpdateUserSettingsOptions) SetNotificationLanguage(notificationLanguage string) *UpdateUserSettingsOptions {
	options.NotificationLanguage = core.StringPtr(notificationLanguage)
	return options
}

// SetAllowedIpAddresses : Allow user to set AllowedIpAddresses
func (options *UpdateUserSettingsOptions) SetAllowedIpAddresses(allowedIpAddresses string) *UpdateUserSettingsOptions {
	options.AllowedIpAddresses = core.StringPtr(allowedIpAddresses)
	return options
}

// SetSelfManage : Allow user to set SelfManage
func (options *UpdateUserSettingsOptions) SetSelfManage(selfManage bool) *UpdateUserSettingsOptions {
	options.SelfManage = core.BoolPtr(selfManage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserSettingsOptions) SetHeaders(param map[string]string) *UpdateUserSettingsOptions {
	options.Headers = param
	return options
}

// UserLinkages : The returned list of user linkages.
type UserLinkages struct {
	// shows the origin of the user and id of that origin.
	Linkages []Linkage `json:"linkages,omitempty"`
}


// UnmarshalUserLinkages constructs an instance of UserLinkages from the specified map.
func UnmarshalUserLinkages(m map[string]interface{}) (result *UserLinkages, err error) {
	obj := new(UserLinkages)
	obj.Linkages, err = UnmarshalLinkageSliceAsProperty(m, "linkages")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUserLinkagesSlice unmarshals a slice of UserLinkages instances from the specified list of maps.
func UnmarshalUserLinkagesSlice(s []interface{}) (slice []UserLinkages, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UserLinkages'")
			return
		}
		obj, e := UnmarshalUserLinkages(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUserLinkagesAsProperty unmarshals an instance of UserLinkages that is stored as a property
// within the specified map.
func UnmarshalUserLinkagesAsProperty(m map[string]interface{}, propertyName string) (result *UserLinkages, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UserLinkages'", propertyName)
			return
		}
		result, err = UnmarshalUserLinkages(objMap)
	}
	return
}

// UnmarshalUserLinkagesSliceAsProperty unmarshals a slice of UserLinkages instances that are stored as a property
// within the specified map.
func UnmarshalUserLinkagesSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UserLinkages, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UserLinkages'", propertyName)
			return
		}
		slice, err = UnmarshalUserLinkagesSlice(vSlice)
	}
	return
}

// UserList : The returned list of users.
type UserList struct {
	// the number of users returned.
	TotalResults *float64 `json:"total_results,omitempty"`

	// limit of the users returned in a page.
	Limit *float64 `json:"limit,omitempty"`

	// the first url of the get users api.
	FirstURL *string `json:"first_url,omitempty"`

	// the next url of the get users api.
	NextURL *string `json:"next_url,omitempty"`

	// shows the users in the account.
	Resources []UserProfile `json:"resources,omitempty"`
}


// UnmarshalUserList constructs an instance of UserList from the specified map.
func UnmarshalUserList(m map[string]interface{}) (result *UserList, err error) {
	obj := new(UserList)
	obj.TotalResults, err = core.UnmarshalFloat64(m, "total_results")
	if err != nil {
		return
	}
	obj.Limit, err = core.UnmarshalFloat64(m, "limit")
	if err != nil {
		return
	}
	obj.FirstURL, err = core.UnmarshalString(m, "first_url")
	if err != nil {
		return
	}
	obj.NextURL, err = core.UnmarshalString(m, "next_url")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalUserProfileSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUserListSlice unmarshals a slice of UserList instances from the specified list of maps.
func UnmarshalUserListSlice(s []interface{}) (slice []UserList, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UserList'")
			return
		}
		obj, e := UnmarshalUserList(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUserListAsProperty unmarshals an instance of UserList that is stored as a property
// within the specified map.
func UnmarshalUserListAsProperty(m map[string]interface{}, propertyName string) (result *UserList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UserList'", propertyName)
			return
		}
		result, err = UnmarshalUserList(objMap)
	}
	return
}

// UnmarshalUserListSliceAsProperty unmarshals a slice of UserList instances that are stored as a property
// within the specified map.
func UnmarshalUserListSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UserList, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UserList'", propertyName)
			return
		}
		slice, err = UnmarshalUserListSlice(vSlice)
	}
	return
}

// UserProfile : The returned user profile.
type UserProfile struct {
	// An alpha-numeric value identifying the user profile.
	ID *string `json:"id,omitempty"`

	// An alpha-numeric value identifying the user's iam id.
	IamID *string `json:"iam_id,omitempty"`

	// The value would be IBMid or SL.
	Realm *string `json:"realm,omitempty"`

	// The user id used for login.
	UserID *string `json:"user_id,omitempty"`

	// The first name of the user.
	Firstname *string `json:"firstname,omitempty"`

	// The last name of the user.
	Lastname *string `json:"lastname,omitempty"`

	// The state of the user, Possible value:"PROCESSING" | "PENDING" | "ACTIVE" | "DISABLED" | "VPN_ONLY".
	State *string `json:"state,omitempty"`

	// The email of the user.
	Email *string `json:"email,omitempty"`

	// The phone for the user.
	Phonenumber *string `json:"phonenumber,omitempty"`

	// The altphonenumber of the user.
	Altphonenumber *string `json:"altphonenumber,omitempty"`

	// The link of the photo of user.
	Photo *string `json:"photo,omitempty"`

	// An alpha-numeric value identifying the account ID.
	AccountID *string `json:"account_id,omitempty"`

	// shows the origin of the user and id of that origin.
	Linkages []Linkage `json:"linkages,omitempty"`
}


// UnmarshalUserProfile constructs an instance of UserProfile from the specified map.
func UnmarshalUserProfile(m map[string]interface{}) (result *UserProfile, err error) {
	obj := new(UserProfile)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.IamID, err = core.UnmarshalString(m, "iam_id")
	if err != nil {
		return
	}
	obj.Realm, err = core.UnmarshalString(m, "realm")
	if err != nil {
		return
	}
	obj.UserID, err = core.UnmarshalString(m, "user_id")
	if err != nil {
		return
	}
	obj.Firstname, err = core.UnmarshalString(m, "firstname")
	if err != nil {
		return
	}
	obj.Lastname, err = core.UnmarshalString(m, "lastname")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.Email, err = core.UnmarshalString(m, "email")
	if err != nil {
		return
	}
	obj.Phonenumber, err = core.UnmarshalString(m, "phonenumber")
	if err != nil {
		return
	}
	obj.Altphonenumber, err = core.UnmarshalString(m, "altphonenumber")
	if err != nil {
		return
	}
	obj.Photo, err = core.UnmarshalString(m, "photo")
	if err != nil {
		return
	}
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Linkages, err = UnmarshalLinkageSliceAsProperty(m, "linkages")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUserProfileSlice unmarshals a slice of UserProfile instances from the specified list of maps.
func UnmarshalUserProfileSlice(s []interface{}) (slice []UserProfile, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UserProfile'")
			return
		}
		obj, e := UnmarshalUserProfile(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUserProfileAsProperty unmarshals an instance of UserProfile that is stored as a property
// within the specified map.
func UnmarshalUserProfileAsProperty(m map[string]interface{}, propertyName string) (result *UserProfile, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UserProfile'", propertyName)
			return
		}
		result, err = UnmarshalUserProfile(objMap)
	}
	return
}

// UnmarshalUserProfileSliceAsProperty unmarshals a slice of UserProfile instances that are stored as a property
// within the specified map.
func UnmarshalUserProfileSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UserProfile, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UserProfile'", propertyName)
			return
		}
		slice, err = UnmarshalUserProfileSlice(vSlice)
	}
	return
}

// UserSettings : The returned user settings.
type UserSettings struct {
	// UI language, default value empty.
	Language *string `json:"language,omitempty"`

	// For email, phone notification, default value empty.
	NotificationLanguage *string `json:"notification_language,omitempty"`

	// Ip address string use comma to separate string.
	AllowedIpAddresses *string `json:"allowed_ip_addresses,omitempty"`

	// a field set for user be able to self manage or not, default false.
	SelfManage *bool `json:"self_manage,omitempty"`
}


// UnmarshalUserSettings constructs an instance of UserSettings from the specified map.
func UnmarshalUserSettings(m map[string]interface{}) (result *UserSettings, err error) {
	obj := new(UserSettings)
	obj.Language, err = core.UnmarshalString(m, "language")
	if err != nil {
		return
	}
	obj.NotificationLanguage, err = core.UnmarshalString(m, "notification_language")
	if err != nil {
		return
	}
	obj.AllowedIpAddresses, err = core.UnmarshalString(m, "allowed_ip_addresses")
	if err != nil {
		return
	}
	obj.SelfManage, err = core.UnmarshalBool(m, "self_manage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUserSettingsSlice unmarshals a slice of UserSettings instances from the specified list of maps.
func UnmarshalUserSettingsSlice(s []interface{}) (slice []UserSettings, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UserSettings'")
			return
		}
		obj, e := UnmarshalUserSettings(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUserSettingsAsProperty unmarshals an instance of UserSettings that is stored as a property
// within the specified map.
func UnmarshalUserSettingsAsProperty(m map[string]interface{}, propertyName string) (result *UserSettings, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UserSettings'", propertyName)
			return
		}
		result, err = UnmarshalUserSettings(objMap)
	}
	return
}

// UnmarshalUserSettingsSliceAsProperty unmarshals a slice of UserSettings instances that are stored as a property
// within the specified map.
func UnmarshalUserSettingsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UserSettings, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UserSettings'", propertyName)
			return
		}
		slice, err = UnmarshalUserSettingsSlice(vSlice)
	}
	return
}

// InviteUser : Invite a user.
type InviteUser struct {
	// An email of the user to be invited.
	Email *string `json:"email,omitempty"`
}


// UnmarshalInviteUser constructs an instance of InviteUser from the specified map.
func UnmarshalInviteUser(m map[string]interface{}) (result *InviteUser, err error) {
	obj := new(InviteUser)
	obj.Email, err = core.UnmarshalString(m, "email")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalInviteUserSlice unmarshals a slice of InviteUser instances from the specified list of maps.
func UnmarshalInviteUserSlice(s []interface{}) (slice []InviteUser, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'InviteUser'")
			return
		}
		obj, e := UnmarshalInviteUser(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalInviteUserAsProperty unmarshals an instance of InviteUser that is stored as a property
// within the specified map.
func UnmarshalInviteUserAsProperty(m map[string]interface{}, propertyName string) (result *InviteUser, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'InviteUser'", propertyName)
			return
		}
		result, err = UnmarshalInviteUser(objMap)
	}
	return
}

// UnmarshalInviteUserSliceAsProperty unmarshals a slice of InviteUser instances that are stored as a property
// within the specified map.
func UnmarshalInviteUserSliceAsProperty(m map[string]interface{}, propertyName string) (slice []InviteUser, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'InviteUser'", propertyName)
			return
		}
		slice, err = UnmarshalInviteUserSlice(vSlice)
	}
	return
}

// Linkage : Origin of the user and its id.
type Linkage struct {
	// A string shows the name of the origin.
	Origin *string `json:"origin,omitempty"`

	// An alpha-numeric value identifying the origin.
	ID *string `json:"id,omitempty"`
}


// UnmarshalLinkage constructs an instance of Linkage from the specified map.
func UnmarshalLinkage(m map[string]interface{}) (result *Linkage, err error) {
	obj := new(Linkage)
	obj.Origin, err = core.UnmarshalString(m, "origin")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalLinkageSlice unmarshals a slice of Linkage instances from the specified list of maps.
func UnmarshalLinkageSlice(s []interface{}) (slice []Linkage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Linkage'")
			return
		}
		obj, e := UnmarshalLinkage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalLinkageAsProperty unmarshals an instance of Linkage that is stored as a property
// within the specified map.
func UnmarshalLinkageAsProperty(m map[string]interface{}, propertyName string) (result *Linkage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Linkage'", propertyName)
			return
		}
		result, err = UnmarshalLinkage(objMap)
	}
	return
}

// UnmarshalLinkageSliceAsProperty unmarshals a slice of Linkage instances that are stored as a property
// within the specified map.
func UnmarshalLinkageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Linkage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Linkage'", propertyName)
			return
		}
		slice, err = UnmarshalLinkageSlice(vSlice)
	}
	return
}
