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

// Package enterprisemanagementv1 : Operations and models for the EnterpriseManagementV1 service
package enterprisemanagementv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"reflect"
)

// EnterpriseManagementV1 : The Enterprise Management API enables you to create and manage an enterprise, account
// groups, and accounts within the enterprise.
//
// Version: 1.0
type EnterpriseManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://enterprise.test.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "enterprise_management"

// EnterpriseManagementV1Options : Service options
type EnterpriseManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewEnterpriseManagementV1UsingExternalConfig : constructs an instance of EnterpriseManagementV1 with passed in options and external configuration.
func NewEnterpriseManagementV1UsingExternalConfig(options *EnterpriseManagementV1Options) (enterpriseManagement *EnterpriseManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	enterpriseManagement, err = NewEnterpriseManagementV1(options)
	if err != nil {
		return
	}

	err = enterpriseManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = enterpriseManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewEnterpriseManagementV1 : constructs an instance of EnterpriseManagementV1 with passed in options.
func NewEnterpriseManagementV1(options *EnterpriseManagementV1Options) (service *EnterpriseManagementV1, err error) {
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

	service = &EnterpriseManagementV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (enterpriseManagement *EnterpriseManagementV1) SetServiceURL(url string) error {
	return enterpriseManagement.Service.SetServiceURL(url)
}

// CreateAccountGroup : Create an account group
// Create a new account group, which can be used to group together multiple accounts. To create an account group, you
// must have an existing enterprise. The API creates an account group entity under the parent that is specified in the
// payload of the request. The request also takes in the name and the primary contact of this new account group.
func (enterpriseManagement *EnterpriseManagementV1) CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions) (result *CreateAccountGroupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAccountGroupOptions, "createAccountGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAccountGroupOptions, "createAccountGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"account-groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAccountGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "CreateAccountGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAccountGroupOptions.Parent != nil {
		body["parent"] = createAccountGroupOptions.Parent
	}
	if createAccountGroupOptions.Name != nil {
		body["name"] = createAccountGroupOptions.Name
	}
	if createAccountGroupOptions.PrimaryContactIamID != nil {
		body["primary_contact_iam_id"] = createAccountGroupOptions.PrimaryContactIamID
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
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateAccountGroupResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListAccountGroups : Get account groups by query parameter
// Retrieve all account groups based on the values that are passed in the query parameters. If no query parameter is
// passed, all of the account groups in the enterprise for which the calling identity has access are returned.
// <br/><br/>You can use pagination parameters to filter the results. The `limit` field can be used to limit the number
// of results that are displayed for this method.<br/><br/>This method ensures that only the account groups that the
// user has access to are returned. Access can be controlled either through a policy on a specific account group, or
// account-level platform services access roles, such as Administrator, Editor, Operator, or Viewer. When you call the
// method with the `enterprise_id`, `parent_account_group_id` or `parent` query parameter, all of the account groups
// that are immediate children of this entity are returned. Authentication is performed on all account groups before
// they are returned to the user to ensure that only those account groups are returned to which the calling identity has
// access.
func (enterpriseManagement *EnterpriseManagementV1) ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions) (result *ListAccountGroupsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAccountGroupsOptions, "listAccountGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"account-groups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAccountGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "ListAccountGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAccountGroupsOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*listAccountGroupsOptions.EnterpriseID))
	}
	if listAccountGroupsOptions.ParentAccountGroupID != nil {
		builder.AddQuery("parent_account_group_id", fmt.Sprint(*listAccountGroupsOptions.ParentAccountGroupID))
	}
	if listAccountGroupsOptions.Parent != nil {
		builder.AddQuery("parent", fmt.Sprint(*listAccountGroupsOptions.Parent))
	}
	if listAccountGroupsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAccountGroupsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListAccountGroupsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAccountGroupByID : Get account group by ID
// Retrieve an account by the `account_group_id` parameter. All data related to the account group is returned only if
// the caller has access to retrieve the account group.
func (enterpriseManagement *EnterpriseManagementV1) GetAccountGroupByID(getAccountGroupByIdOptions *GetAccountGroupByIdOptions) (result *AccountGroupResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountGroupByIdOptions, "getAccountGroupByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountGroupByIdOptions, "getAccountGroupByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"account-groups"}
	pathParameters := []string{*getAccountGroupByIdOptions.AccountGroupID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountGroupByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetAccountGroupByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountGroupResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateAccountGroup : Update an account group
// Update the name or IAM ID of the primary contact for an existing account group. The new primary contact must already
// be a user in the enterprise account.
func (enterpriseManagement *EnterpriseManagementV1) UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountGroupOptions, "updateAccountGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountGroupOptions, "updateAccountGroupOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"account-groups"}
	pathParameters := []string{*updateAccountGroupOptions.AccountGroupID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "UpdateAccountGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountGroupOptions.Name != nil {
		body["name"] = updateAccountGroupOptions.Name
	}
	if updateAccountGroupOptions.PrimaryContactIamID != nil {
		body["primary_contact_iam_id"] = updateAccountGroupOptions.PrimaryContactIamID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// GetAccountGroupPermissibleActions : Get permissible actions for an account group
// Return all the actions that are allowed on a particular account group. This method takes an array of IAM actions in
// the body of the request and returns those actions that can be performed by the caller. An authentication check is
// performed for each action that is passed in the payload.
func (enterpriseManagement *EnterpriseManagementV1) GetAccountGroupPermissibleActions(getAccountGroupPermissibleActionsOptions *GetAccountGroupPermissibleActionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountGroupPermissibleActionsOptions, "getAccountGroupPermissibleActionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountGroupPermissibleActionsOptions, "getAccountGroupPermissibleActionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"account-groups", "permissible-actions"}
	pathParameters := []string{*getAccountGroupPermissibleActionsOptions.AccountGroupID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountGroupPermissibleActionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetAccountGroupPermissibleActions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getAccountGroupPermissibleActionsOptions.Actions != nil {
		body["actions"] = getAccountGroupPermissibleActionsOptions.Actions
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// ImportAccountToEnterprise : Import an account into an enterprise
// Import an existing stand-alone account into an enterprise. The existing account can be any type: trial (`TRIAL`),
// Lite (`STANDARD`), Pay-As-You-Go (`PAYG`), or Subscription (`SUBSCRIPTION`). In the case of a `SUBSCRIPTION` account,
// the credits, promotional offers, and discounts are migrated to the billing unit of the enterprise. For a billable
// account (`PAYG` or `SUBSCRIPTION`), the country and currency code of the existing account and the billing unit of the
// enterprise must match. The API returns a `202` response and performs asynchronous operations to import the account
// into the enterprise. <br/></br>For more information about impacts to the account, see [Adding accounts to an
// enterprise](https://{DomainName}/docs/account?topic=account-enterprise-add).
func (enterpriseManagement *EnterpriseManagementV1) ImportAccountToEnterprise(importAccountToEnterpriseOptions *ImportAccountToEnterpriseOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importAccountToEnterpriseOptions, "importAccountToEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importAccountToEnterpriseOptions, "importAccountToEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises", "import/accounts"}
	pathParameters := []string{*importAccountToEnterpriseOptions.EnterpriseID, *importAccountToEnterpriseOptions.AccountID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importAccountToEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "ImportAccountToEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if importAccountToEnterpriseOptions.Parent != nil {
		body["parent"] = importAccountToEnterpriseOptions.Parent
	}
	if importAccountToEnterpriseOptions.BillingUnitID != nil {
		body["billing_unit_id"] = importAccountToEnterpriseOptions.BillingUnitID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// CreateAccount : Create a new account in an enterprise
// Create a new account as a part of an existing enterprise. The API creates an account entity under the parent that is
// specified in the payload of the request. The request also takes in the name and the owner of this new account. The
// owner must have a valid IBMid that's registered with {{site.data.keyword.cloud_notm}}, but they don't need to be a
// user in the enterprise account.
func (enterpriseManagement *EnterpriseManagementV1) CreateAccount(createAccountOptions *CreateAccountOptions) (result *CreateAccountResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAccountOptions, "createAccountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAccountOptions, "createAccountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"accounts"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "CreateAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createAccountOptions.Parent != nil {
		body["parent"] = createAccountOptions.Parent
	}
	if createAccountOptions.Name != nil {
		body["name"] = createAccountOptions.Name
	}
	if createAccountOptions.OwnerIamID != nil {
		body["owner_iam_id"] = createAccountOptions.OwnerIamID
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
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateAccountResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListAccounts : Get accounts by query parameter
// Retrieve all accounts based on the values that are passed in the query parameters. If no query parameter is passed,
// all of the accounts in the enterprise for which the calling identity has access are returned. <br/><br/>You can use
// pagination parameters to filter the results. The `limit` field can be used to limit the number of results that are
// displayed for this method.<br/><br/>This method ensures that only the accounts that the user has access to are
// returned. Access can be controlled either through a policy on a specific account, or account-level platform services
// access roles, such as Administrator, Editor, Operator, or Viewer. When you call the method with the `enterprise_id`,
// `account_group_id` or `parent` query parameter, all of the accounts that are immediate children of this entity are
// returned. Authentication is performed on all the accounts before they are returned to the user to ensure that only
// those accounts are returned to which the calling identity has access to.
func (enterpriseManagement *EnterpriseManagementV1) ListAccounts(listAccountsOptions *ListAccountsOptions) (result *ListAccountsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAccountsOptions, "listAccountsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"accounts"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAccountsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "ListAccounts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAccountsOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*listAccountsOptions.EnterpriseID))
	}
	if listAccountsOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*listAccountsOptions.AccountGroupID))
	}
	if listAccountsOptions.Parent != nil {
		builder.AddQuery("parent", fmt.Sprint(*listAccountsOptions.Parent))
	}
	if listAccountsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAccountsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListAccountsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAccountByID : Get account by ID
// Retrieve an account by the `account_id` parameter. All data related to the account is returned only if the caller has
// access to retrieve the account.
func (enterpriseManagement *EnterpriseManagementV1) GetAccountByID(getAccountByIdOptions *GetAccountByIdOptions) (result *AccountResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountByIdOptions, "getAccountByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountByIdOptions, "getAccountByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"accounts"}
	pathParameters := []string{*getAccountByIdOptions.AccountID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetAccountByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateAccount : Move an account with the enterprise
// Move an account to a different parent within the same enterprise.
func (enterpriseManagement *EnterpriseManagementV1) UpdateAccount(updateAccountOptions *UpdateAccountOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountOptions, "updateAccountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountOptions, "updateAccountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"accounts"}
	pathParameters := []string{*updateAccountOptions.AccountID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "UpdateAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountOptions.Parent != nil {
		body["parent"] = updateAccountOptions.Parent
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// GetAccountPermissibleActions : Get permissible actions for an account
// Return all the actions that are allowed on a particular account. This method takes an array of IAM actions in the
// body of the request and returns those actions which can be performed by the caller. An authentication check is
// performed for each action that is passed in the payload.
func (enterpriseManagement *EnterpriseManagementV1) GetAccountPermissibleActions(getAccountPermissibleActionsOptions *GetAccountPermissibleActionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountPermissibleActionsOptions, "getAccountPermissibleActionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountPermissibleActionsOptions, "getAccountPermissibleActionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"accounts", "permissible-actions"}
	pathParameters := []string{*getAccountPermissibleActionsOptions.AccountID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountPermissibleActionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetAccountPermissibleActions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getAccountPermissibleActionsOptions.Actions != nil {
		body["actions"] = getAccountPermissibleActionsOptions.Actions
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// CreateEnterprise : Create an enterprise
// Create a new enterprise, which you can use to centrally manage multiple accounts. To create an enterprise, you must
// have an active Subscription account. <br/><br/>The API creates an enterprise entity, which is the root of the
// enterprise hierarchy. It also creates a new enterprise account that is used to manage the enterprise. All
// subscriptions, support entitlements, credits, and discounts from the source subscription account are migrated to the
// enterprise account, and the source account becomes a child account in the hierarchy. The user that you assign as the
// enterprise primary contact is also assigned as the owner of the enterprise account.
func (enterpriseManagement *EnterpriseManagementV1) CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions) (result *CreateEnterpriseResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createEnterpriseOptions, "createEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createEnterpriseOptions, "createEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "CreateEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createEnterpriseOptions.SourceAccountID != nil {
		body["source_account_id"] = createEnterpriseOptions.SourceAccountID
	}
	if createEnterpriseOptions.Name != nil {
		body["name"] = createEnterpriseOptions.Name
	}
	if createEnterpriseOptions.PrimaryContactIamID != nil {
		body["primary_contact_iam_id"] = createEnterpriseOptions.PrimaryContactIamID
	}
	if createEnterpriseOptions.Domain != nil {
		body["domain"] = createEnterpriseOptions.Domain
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
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateEnterpriseResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListEnterprises : Get enterprise by query parameter
// Retrieve all enterprises for a given ID by passing the IDs on query parameters. If no ID is passed, the enterprises
// for which the calling identity is the primary contact are returned. You can use pagination parameters to filter the
// results. <br/><br/>This method ensures that only the enterprises that the user has access to are returned. Access can
// be controlled either through a policy on a specific enterprise, or account-level platform services access roles, such
// as Administrator, Editor, Operator, or Viewer. When you call the method with the `enterprise_account_id` or
// `account_id` query parameter, the account ID in the token is compared with that in the query parameter. If these
// account IDs match, authentication isn't performed and the enterprise information is returned. If the account IDs
// don't match, authentication is performed and only then is the enterprise information returned in the response.
func (enterpriseManagement *EnterpriseManagementV1) ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions) (result *ListEnterprisesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listEnterprisesOptions, "listEnterprisesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listEnterprisesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "ListEnterprises")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listEnterprisesOptions.EnterpriseAccountID != nil {
		builder.AddQuery("enterprise_account_id", fmt.Sprint(*listEnterprisesOptions.EnterpriseAccountID))
	}
	if listEnterprisesOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*listEnterprisesOptions.AccountGroupID))
	}
	if listEnterprisesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listEnterprisesOptions.AccountID))
	}
	if listEnterprisesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listEnterprisesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListEnterprisesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetEnterprise : Get enterprise by ID
// Retrieve an enterprise by the `enterprise_id` parameter. All data related to the enterprise is returned only if the
// caller has access to retrieve the enterprise.
func (enterpriseManagement *EnterpriseManagementV1) GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions) (result *EnterpriseResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEnterpriseOptions, "getEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEnterpriseOptions, "getEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{*getEnterpriseOptions.EnterpriseID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = enterpriseManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnterpriseResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateEnterprise : Update an enterprise
// Update the name, domain, or IAM ID of the primary contact for an existing enterprise. The new primary contact must
// already be a user in the enterprise account.
func (enterpriseManagement *EnterpriseManagementV1) UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateEnterpriseOptions, "updateEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateEnterpriseOptions, "updateEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{*updateEnterpriseOptions.EnterpriseID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "UpdateEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateEnterpriseOptions.Name != nil {
		body["name"] = updateEnterpriseOptions.Name
	}
	if updateEnterpriseOptions.Domain != nil {
		body["domain"] = updateEnterpriseOptions.Domain
	}
	if updateEnterpriseOptions.PrimaryContactIamID != nil {
		body["primary_contact_iam_id"] = updateEnterpriseOptions.PrimaryContactIamID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// GetEnterprisePermissibleActions : Get permissible actions for an enterprise
// Return all the actions that are allowed on a particular enterprise. This method takes an array of IAM actions in the
// body of the request and returns those actions which can be performed by the caller. An authentication check is
// performed for each action that is passed in the payload.
func (enterpriseManagement *EnterpriseManagementV1) GetEnterprisePermissibleActions(getEnterprisePermissibleActionsOptions *GetEnterprisePermissibleActionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEnterprisePermissibleActionsOptions, "getEnterprisePermissibleActionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEnterprisePermissibleActionsOptions, "getEnterprisePermissibleActionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises", "permissible-actions"}
	pathParameters := []string{*getEnterprisePermissibleActionsOptions.EnterpriseID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(enterpriseManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnterprisePermissibleActionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_management", "V1", "GetEnterprisePermissibleActions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if getEnterprisePermissibleActionsOptions.Actions != nil {
		body["actions"] = getEnterprisePermissibleActionsOptions.Actions
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseManagement.Service.Request(request, nil)

	return
}

// AccountGroupResponse : An object that represents account group.
type AccountGroupResponse struct {
	// The URL of the account group.
	URL *string `json:"url,omitempty"`

	// The account group ID.
	ID *string `json:"id,omitempty"`

	// The CRN of the account group.
	Crn *string `json:"crn,omitempty"`

	// The CRN of the parent of the account group.
	Parent *string `json:"parent,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The enterprise ID that the account group is a part of.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The path from the enterprise to this particular account group.
	EnterprisePath *string `json:"enterprise_path,omitempty"`

	// The name of the account group.
	Name *string `json:"name,omitempty"`

	// The state of the account group.
	State *string `json:"state,omitempty"`

	// The IAM ID of the primary contact of the account group.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// The email address of the primary contact of the account group.
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`

	// The time stamp at which the account group was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the account group.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the account group was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the account group.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalAccountGroupResponse unmarshals an instance of AccountGroupResponse from the specified map of raw messages.
func UnmarshalAccountGroupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountGroupResponse)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
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
	err = core.UnmarshalPrimitive(m, "parent", &obj.Parent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_path", &obj.EnterprisePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_iam_id", &obj.PrimaryContactIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_email", &obj.PrimaryContactEmail)
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

// AccountResponse : The response from successfully calling get account.
type AccountResponse struct {
	// The URL of the account.
	URL *string `json:"url,omitempty"`

	// The account ID.
	ID *string `json:"id,omitempty"`

	// The Cloud Resource Name (CRN) of the account.
	Crn *string `json:"crn,omitempty"`

	// The CRN of the parent of the account.
	Parent *string `json:"parent,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The enterprise ID that the account is a part of.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The path from the enterprise to this particular account.
	EnterprisePath *string `json:"enterprise_path,omitempty"`

	// The name of the account.
	Name *string `json:"name,omitempty"`

	// The state of the account.
	State *string `json:"state,omitempty"`

	// The IAM ID of the owner of the account.
	OwnerIamID *string `json:"owner_iam_id,omitempty"`

	// The type of account - whether it is free or paid.
	Paid *bool `json:"paid,omitempty"`

	// The email address of the owner of the account.
	OwnerEmail *string `json:"owner_email,omitempty"`

	// The flag to indicate whether the account is an enterprise account or not.
	IsEnterpriseAccount *bool `json:"is_enterprise_account,omitempty"`

	// The time stamp at which the account was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the account.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the account was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the account.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalAccountResponse unmarshals an instance of AccountResponse from the specified map of raw messages.
func UnmarshalAccountResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountResponse)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
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
	err = core.UnmarshalPrimitive(m, "parent", &obj.Parent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_path", &obj.EnterprisePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_iam_id", &obj.OwnerIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "paid", &obj.Paid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_email", &obj.OwnerEmail)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_enterprise_account", &obj.IsEnterpriseAccount)
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

// CreateAccountGroupOptions : The CreateAccountGroup options.
type CreateAccountGroupOptions struct {
	// The CRN of the parent under which the account group will be created. The parent can be an existing account group or
	// the enterprise itself.
	Parent *string `json:"parent" validate:"required"`

	// The name of the account group. This field must have 3 - 60 characters.
	Name *string `json:"name" validate:"required"`

	// The IAM ID of the primary contact for this account group, such as `IBMid-0123ABC`. The IAM ID must already exist.
	PrimaryContactIamID *string `json:"primary_contact_iam_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAccountGroupOptions : Instantiate CreateAccountGroupOptions
func (*EnterpriseManagementV1) NewCreateAccountGroupOptions(parent string, name string, primaryContactIamID string) *CreateAccountGroupOptions {
	return &CreateAccountGroupOptions{
		Parent: core.StringPtr(parent),
		Name: core.StringPtr(name),
		PrimaryContactIamID: core.StringPtr(primaryContactIamID),
	}
}

// SetParent : Allow user to set Parent
func (options *CreateAccountGroupOptions) SetParent(parent string) *CreateAccountGroupOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetName : Allow user to set Name
func (options *CreateAccountGroupOptions) SetName(name string) *CreateAccountGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetPrimaryContactIamID : Allow user to set PrimaryContactIamID
func (options *CreateAccountGroupOptions) SetPrimaryContactIamID(primaryContactIamID string) *CreateAccountGroupOptions {
	options.PrimaryContactIamID = core.StringPtr(primaryContactIamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccountGroupOptions) SetHeaders(param map[string]string) *CreateAccountGroupOptions {
	options.Headers = param
	return options
}

// CreateAccountGroupResponse : Create account group request completed successfully.
type CreateAccountGroupResponse struct {
	// The ID of the account group entity that was created.
	AccountGroupID *string `json:"account_group_id,omitempty"`
}


// UnmarshalCreateAccountGroupResponse unmarshals an instance of CreateAccountGroupResponse from the specified map of raw messages.
func UnmarshalCreateAccountGroupResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateAccountGroupResponse)
	err = core.UnmarshalPrimitive(m, "account_group_id", &obj.AccountGroupID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateAccountOptions : The CreateAccount options.
type CreateAccountOptions struct {
	// The CRN of the parent under which the account will be created. The parent can be an existing account group or the
	// enterprise itself.
	Parent *string `json:"parent" validate:"required"`

	// The name of the account. This field must have 3 - 60 characters.
	Name *string `json:"name" validate:"required"`

	// The IAM ID of the account owner, such as `IBMid-0123ABC`. The IAM ID must already exist.
	OwnerIamID *string `json:"owner_iam_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAccountOptions : Instantiate CreateAccountOptions
func (*EnterpriseManagementV1) NewCreateAccountOptions(parent string, name string, ownerIamID string) *CreateAccountOptions {
	return &CreateAccountOptions{
		Parent: core.StringPtr(parent),
		Name: core.StringPtr(name),
		OwnerIamID: core.StringPtr(ownerIamID),
	}
}

// SetParent : Allow user to set Parent
func (options *CreateAccountOptions) SetParent(parent string) *CreateAccountOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetName : Allow user to set Name
func (options *CreateAccountOptions) SetName(name string) *CreateAccountOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetOwnerIamID : Allow user to set OwnerIamID
func (options *CreateAccountOptions) SetOwnerIamID(ownerIamID string) *CreateAccountOptions {
	options.OwnerIamID = core.StringPtr(ownerIamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAccountOptions) SetHeaders(param map[string]string) *CreateAccountOptions {
	options.Headers = param
	return options
}

// CreateAccountResponse : The create account request completed successfully.
type CreateAccountResponse struct {
	// The ID of the account group entity that was created.
	AccountGroupID *string `json:"account_group_id,omitempty"`
}


// UnmarshalCreateAccountResponse unmarshals an instance of CreateAccountResponse from the specified map of raw messages.
func UnmarshalCreateAccountResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateAccountResponse)
	err = core.UnmarshalPrimitive(m, "account_group_id", &obj.AccountGroupID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateEnterpriseOptions : The CreateEnterprise options.
type CreateEnterpriseOptions struct {
	// The ID of the account that is used to create the enterprise.
	SourceAccountID *string `json:"source_account_id" validate:"required"`

	// The name of the enterprise. This field must have 3 - 60 characters.
	Name *string `json:"name" validate:"required"`

	// The IAM ID of the enterprise primary contact, such as `IBMid-0123ABC`. The IAM ID must already exist.
	PrimaryContactIamID *string `json:"primary_contact_iam_id" validate:"required"`

	// A domain or subdomain for the enterprise, such as `example.com` or `my.example.com`.
	Domain *string `json:"domain,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateEnterpriseOptions : Instantiate CreateEnterpriseOptions
func (*EnterpriseManagementV1) NewCreateEnterpriseOptions(sourceAccountID string, name string, primaryContactIamID string) *CreateEnterpriseOptions {
	return &CreateEnterpriseOptions{
		SourceAccountID: core.StringPtr(sourceAccountID),
		Name: core.StringPtr(name),
		PrimaryContactIamID: core.StringPtr(primaryContactIamID),
	}
}

// SetSourceAccountID : Allow user to set SourceAccountID
func (options *CreateEnterpriseOptions) SetSourceAccountID(sourceAccountID string) *CreateEnterpriseOptions {
	options.SourceAccountID = core.StringPtr(sourceAccountID)
	return options
}

// SetName : Allow user to set Name
func (options *CreateEnterpriseOptions) SetName(name string) *CreateEnterpriseOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetPrimaryContactIamID : Allow user to set PrimaryContactIamID
func (options *CreateEnterpriseOptions) SetPrimaryContactIamID(primaryContactIamID string) *CreateEnterpriseOptions {
	options.PrimaryContactIamID = core.StringPtr(primaryContactIamID)
	return options
}

// SetDomain : Allow user to set Domain
func (options *CreateEnterpriseOptions) SetDomain(domain string) *CreateEnterpriseOptions {
	options.Domain = core.StringPtr(domain)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateEnterpriseOptions) SetHeaders(param map[string]string) *CreateEnterpriseOptions {
	options.Headers = param
	return options
}

// CreateEnterpriseResponse : The response from calling create enterprise.
type CreateEnterpriseResponse struct {
	// The ID of the enterprise entity that was created. This entity is the root of the hierarchy.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The ID of the enterprise account that was created. The enterprise account is used to manage billing and access to
	// the enterprise management.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`
}


// UnmarshalCreateEnterpriseResponse unmarshals an instance of CreateEnterpriseResponse from the specified map of raw messages.
func UnmarshalCreateEnterpriseResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateEnterpriseResponse)
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnterpriseResponse : The response from calling get enterprise.
type EnterpriseResponse struct {
	// The URL of the enterprise.
	URL *string `json:"url,omitempty"`

	// The enterprise ID.
	ID *string `json:"id,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The Cloud Resource Name (CRN) of the enterprise.
	Crn *string `json:"crn,omitempty"`

	// The name of the enterprise.
	Name *string `json:"name,omitempty"`

	// The domain of the enterprise.
	Domain *string `json:"domain,omitempty"`

	// The state of the enterprise.
	State *string `json:"state,omitempty"`

	// The IAM ID of the primary contact of the enterprise, such as `IBMid-0123ABC`.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// The email of the primary contact of the enterprise.
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`

	// The time stamp at which the enterprise was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the enterprise.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the enterprise was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the enterprise.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalEnterpriseResponse unmarshals an instance of EnterpriseResponse from the specified map of raw messages.
func UnmarshalEnterpriseResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnterpriseResponse)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_iam_id", &obj.PrimaryContactIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_email", &obj.PrimaryContactEmail)
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

// GetAccountByIdOptions : The GetAccountByID options.
type GetAccountByIdOptions struct {
	// The ID of the account to retrieve.
	AccountID *string `json:"account_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountByIdOptions : Instantiate GetAccountByIdOptions
func (*EnterpriseManagementV1) NewGetAccountByIdOptions(accountID string) *GetAccountByIdOptions {
	return &GetAccountByIdOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountByIdOptions) SetAccountID(accountID string) *GetAccountByIdOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountByIdOptions) SetHeaders(param map[string]string) *GetAccountByIdOptions {
	options.Headers = param
	return options
}

// GetAccountGroupByIdOptions : The GetAccountGroupByID options.
type GetAccountGroupByIdOptions struct {
	// The ID of the account group to retrieve.
	AccountGroupID *string `json:"account_group_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountGroupByIdOptions : Instantiate GetAccountGroupByIdOptions
func (*EnterpriseManagementV1) NewGetAccountGroupByIdOptions(accountGroupID string) *GetAccountGroupByIdOptions {
	return &GetAccountGroupByIdOptions{
		AccountGroupID: core.StringPtr(accountGroupID),
	}
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *GetAccountGroupByIdOptions) SetAccountGroupID(accountGroupID string) *GetAccountGroupByIdOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountGroupByIdOptions) SetHeaders(param map[string]string) *GetAccountGroupByIdOptions {
	options.Headers = param
	return options
}

// GetAccountGroupPermissibleActionsOptions : The GetAccountGroupPermissibleActions options.
type GetAccountGroupPermissibleActionsOptions struct {
	// The ID of the account group to check for permissible actions.
	AccountGroupID *string `json:"account_group_id" validate:"required"`

	// A list of names of permissible actions.
	Actions []string `json:"actions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountGroupPermissibleActionsOptions : Instantiate GetAccountGroupPermissibleActionsOptions
func (*EnterpriseManagementV1) NewGetAccountGroupPermissibleActionsOptions(accountGroupID string) *GetAccountGroupPermissibleActionsOptions {
	return &GetAccountGroupPermissibleActionsOptions{
		AccountGroupID: core.StringPtr(accountGroupID),
	}
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *GetAccountGroupPermissibleActionsOptions) SetAccountGroupID(accountGroupID string) *GetAccountGroupPermissibleActionsOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetActions : Allow user to set Actions
func (options *GetAccountGroupPermissibleActionsOptions) SetActions(actions []string) *GetAccountGroupPermissibleActionsOptions {
	options.Actions = actions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountGroupPermissibleActionsOptions) SetHeaders(param map[string]string) *GetAccountGroupPermissibleActionsOptions {
	options.Headers = param
	return options
}

// GetAccountPermissibleActionsOptions : The GetAccountPermissibleActions options.
type GetAccountPermissibleActionsOptions struct {
	// The ID of the account to check for permissible actions.
	AccountID *string `json:"account_id" validate:"required"`

	// A list of names of permissible actions.
	Actions []string `json:"actions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountPermissibleActionsOptions : Instantiate GetAccountPermissibleActionsOptions
func (*EnterpriseManagementV1) NewGetAccountPermissibleActionsOptions(accountID string) *GetAccountPermissibleActionsOptions {
	return &GetAccountPermissibleActionsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountPermissibleActionsOptions) SetAccountID(accountID string) *GetAccountPermissibleActionsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetActions : Allow user to set Actions
func (options *GetAccountPermissibleActionsOptions) SetActions(actions []string) *GetAccountPermissibleActionsOptions {
	options.Actions = actions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountPermissibleActionsOptions) SetHeaders(param map[string]string) *GetAccountPermissibleActionsOptions {
	options.Headers = param
	return options
}

// GetEnterpriseOptions : The GetEnterprise options.
type GetEnterpriseOptions struct {
	// The ID of the enterprise to retrieve.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnterpriseOptions : Instantiate GetEnterpriseOptions
func (*EnterpriseManagementV1) NewGetEnterpriseOptions(enterpriseID string) *GetEnterpriseOptions {
	return &GetEnterpriseOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *GetEnterpriseOptions) SetEnterpriseID(enterpriseID string) *GetEnterpriseOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnterpriseOptions) SetHeaders(param map[string]string) *GetEnterpriseOptions {
	options.Headers = param
	return options
}

// GetEnterprisePermissibleActionsOptions : The GetEnterprisePermissibleActions options.
type GetEnterprisePermissibleActionsOptions struct {
	// The ID of the enterprise to check for permissible actions.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// A list of names of permissible actions.
	Actions []string `json:"actions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnterprisePermissibleActionsOptions : Instantiate GetEnterprisePermissibleActionsOptions
func (*EnterpriseManagementV1) NewGetEnterprisePermissibleActionsOptions(enterpriseID string) *GetEnterprisePermissibleActionsOptions {
	return &GetEnterprisePermissibleActionsOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *GetEnterprisePermissibleActionsOptions) SetEnterpriseID(enterpriseID string) *GetEnterprisePermissibleActionsOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetActions : Allow user to set Actions
func (options *GetEnterprisePermissibleActionsOptions) SetActions(actions []string) *GetEnterprisePermissibleActionsOptions {
	options.Actions = actions
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnterprisePermissibleActionsOptions) SetHeaders(param map[string]string) *GetEnterprisePermissibleActionsOptions {
	options.Headers = param
	return options
}

// ImportAccountToEnterpriseOptions : The ImportAccountToEnterprise options.
type ImportAccountToEnterpriseOptions struct {
	// The ID of the enterprise to import the stand-alone account into.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// The ID of the existing stand-alone account to be imported.
	AccountID *string `json:"account_id" validate:"required"`

	// The CRN of the expected parent of the imported account. The parent is the enterprise or account group that the
	// account is added to.
	Parent *string `json:"parent,omitempty"`

	// The ID of the [billing unit](/apidocs/enterprise-apis/billing-unit) to use for billing this account in the
	// enterprise.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportAccountToEnterpriseOptions : Instantiate ImportAccountToEnterpriseOptions
func (*EnterpriseManagementV1) NewImportAccountToEnterpriseOptions(enterpriseID string, accountID string) *ImportAccountToEnterpriseOptions {
	return &ImportAccountToEnterpriseOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
		AccountID: core.StringPtr(accountID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *ImportAccountToEnterpriseOptions) SetEnterpriseID(enterpriseID string) *ImportAccountToEnterpriseOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *ImportAccountToEnterpriseOptions) SetAccountID(accountID string) *ImportAccountToEnterpriseOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetParent : Allow user to set Parent
func (options *ImportAccountToEnterpriseOptions) SetParent(parent string) *ImportAccountToEnterpriseOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (options *ImportAccountToEnterpriseOptions) SetBillingUnitID(billingUnitID string) *ImportAccountToEnterpriseOptions {
	options.BillingUnitID = core.StringPtr(billingUnitID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportAccountToEnterpriseOptions) SetHeaders(param map[string]string) *ImportAccountToEnterpriseOptions {
	options.Headers = param
	return options
}

// ListAccountGroupsOptions : The ListAccountGroups options.
type ListAccountGroupsOptions struct {
	// Get account groups that are either immediate children or are a part of the hierarchy for a given enterprise ID.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// Get account groups that are either immediate children or are a part of the hierarchy for a given account group ID.
	ParentAccountGroupID *string `json:"parent_account_group_id,omitempty"`

	// Get account groups that are either immediate children or are a part of the hierarchy for a given parent CRN.
	Parent *string `json:"parent,omitempty"`

	// Return results up to this limit. Valid values are between `0` and `100`.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAccountGroupsOptions : Instantiate ListAccountGroupsOptions
func (*EnterpriseManagementV1) NewListAccountGroupsOptions() *ListAccountGroupsOptions {
	return &ListAccountGroupsOptions{}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *ListAccountGroupsOptions) SetEnterpriseID(enterpriseID string) *ListAccountGroupsOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetParentAccountGroupID : Allow user to set ParentAccountGroupID
func (options *ListAccountGroupsOptions) SetParentAccountGroupID(parentAccountGroupID string) *ListAccountGroupsOptions {
	options.ParentAccountGroupID = core.StringPtr(parentAccountGroupID)
	return options
}

// SetParent : Allow user to set Parent
func (options *ListAccountGroupsOptions) SetParent(parent string) *ListAccountGroupsOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAccountGroupsOptions) SetLimit(limit int64) *ListAccountGroupsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccountGroupsOptions) SetHeaders(param map[string]string) *ListAccountGroupsOptions {
	options.Headers = param
	return options
}

// ListAccountGroupsResources : An object that represents account groups resource.
type ListAccountGroupsResources struct {
	// The URL of the account group.
	URL *string `json:"url,omitempty"`

	// The account group ID.
	ID *string `json:"id,omitempty"`

	// The Cloud Resource Name (CRN) of the account group.
	Crn *string `json:"crn,omitempty"`

	// The CRN of the parent of the account group.
	Parent *string `json:"parent,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The enterprise ID that the account group is a part of.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The path from the enterprise to this particular account group.
	EnterprisePath *string `json:"enterprise_path,omitempty"`

	// The name of the account group.
	Name *string `json:"name,omitempty"`

	// The state of the account group.
	State *string `json:"state,omitempty"`

	// The IAM ID of the primary contact of the account group.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// The email address of the primary contact of the account group.
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`

	// The time stamp at which the account group was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the account group.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the account group was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the account group.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalListAccountGroupsResources unmarshals an instance of ListAccountGroupsResources from the specified map of raw messages.
func UnmarshalListAccountGroupsResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountGroupsResources)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
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
	err = core.UnmarshalPrimitive(m, "parent", &obj.Parent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_path", &obj.EnterprisePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_iam_id", &obj.PrimaryContactIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_email", &obj.PrimaryContactEmail)
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

// ListAccountGroupsResponse : The response from successfully calling list account groups.
type ListAccountGroupsResponse struct {
	// The number of enterprises returned from calling list account groups.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// A string that represents the link to the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// A list of account groups.
	Resources []ListAccountGroupsResources `json:"resources,omitempty"`
}


// UnmarshalListAccountGroupsResponse unmarshals an instance of ListAccountGroupsResponse from the specified map of raw messages.
func UnmarshalListAccountGroupsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountGroupsResponse)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalListAccountGroupsResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAccountResources : An object that represents account resource.
type ListAccountResources struct {
	// The URL of the account.
	URL *string `json:"url,omitempty"`

	// The account ID.
	ID *string `json:"id,omitempty"`

	// The Cloud Resource Name (CRN) of the account.
	Crn *string `json:"crn,omitempty"`

	// The CRN of the parent of the account.
	Parent *string `json:"parent,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The enterprise ID that the account is a part of.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The path from the enterprise to this particular account.
	EnterprisePath *string `json:"enterprise_path,omitempty"`

	// The name of the account.
	Name *string `json:"name,omitempty"`

	// The state of the account.
	State *string `json:"state,omitempty"`

	// The IAM ID of the owner of the account.
	OwnerIamID *string `json:"owner_iam_id,omitempty"`

	// The type of account - whether it is free or paid.
	Paid *bool `json:"paid,omitempty"`

	// The email address of the owner of the account.
	OwnerEmail *string `json:"owner_email,omitempty"`

	// The flag to indicate whether the account is an enterprise account or not.
	IsEnterpriseAccount *bool `json:"is_enterprise_account,omitempty"`

	// The time stamp at which the account was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the account.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the account was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the account.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalListAccountResources unmarshals an instance of ListAccountResources from the specified map of raw messages.
func UnmarshalListAccountResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountResources)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
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
	err = core.UnmarshalPrimitive(m, "parent", &obj.Parent)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_id", &obj.EnterpriseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_path", &obj.EnterprisePath)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_iam_id", &obj.OwnerIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "paid", &obj.Paid)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_email", &obj.OwnerEmail)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_enterprise_account", &obj.IsEnterpriseAccount)
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

// ListAccountsOptions : The ListAccounts options.
type ListAccountsOptions struct {
	// Get accounts that are either immediate children or are a part of the hierarchy for a given enterprise ID.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// Get accounts that are either immediate children or are a part of the hierarchy for a given account group ID.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// Get accounts that are either immediate children or are a part of the hierarchy for a given parent CRN.
	Parent *string `json:"parent,omitempty"`

	// Return results up to this limit. Valid values are between `0` and `100`.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAccountsOptions : Instantiate ListAccountsOptions
func (*EnterpriseManagementV1) NewListAccountsOptions() *ListAccountsOptions {
	return &ListAccountsOptions{}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *ListAccountsOptions) SetEnterpriseID(enterpriseID string) *ListAccountsOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *ListAccountsOptions) SetAccountGroupID(accountGroupID string) *ListAccountsOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetParent : Allow user to set Parent
func (options *ListAccountsOptions) SetParent(parent string) *ListAccountsOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAccountsOptions) SetLimit(limit int64) *ListAccountsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAccountsOptions) SetHeaders(param map[string]string) *ListAccountsOptions {
	options.Headers = param
	return options
}

// ListAccountsResponse : The response from successfully calling list accounts.
type ListAccountsResponse struct {
	// The number of enterprises returned from calling list accounts.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// A string that represents the link to the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// A list of accounts.
	Resources []ListAccountResources `json:"resources,omitempty"`
}


// UnmarshalListAccountsResponse unmarshals an instance of ListAccountsResponse from the specified map of raw messages.
func UnmarshalListAccountsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListAccountsResponse)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalListAccountResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListEnterpriseResources : An object that represents an enterprise.
type ListEnterpriseResources struct {
	// The URL of the enterprise.
	URL *string `json:"url,omitempty"`

	// The enterprise ID.
	ID *string `json:"id,omitempty"`

	// The enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// The Cloud Resource Name (CRN) of the enterprise.
	Crn *string `json:"crn,omitempty"`

	// The name of the enterprise.
	Name *string `json:"name,omitempty"`

	// The domain of the enterprise.
	Domain *string `json:"domain,omitempty"`

	// The state of the enterprise.
	State *string `json:"state,omitempty"`

	// The IAM ID of the primary contact of the enterprise, such as `IBMid-0123ABC`.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// The email of the primary contact of the enterprise.
	PrimaryContactEmail *string `json:"primary_contact_email,omitempty"`

	// The time stamp at which the enterprise was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The IAM ID of the user or service that created the enterprise.
	CreatedBy *string `json:"created_by,omitempty"`

	// The time stamp at which the enterprise was last updated.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// The IAM ID of the user or service that updated the enterprise.
	UpdatedBy *string `json:"updated_by,omitempty"`
}


// UnmarshalListEnterpriseResources unmarshals an instance of ListEnterpriseResources from the specified map of raw messages.
func UnmarshalListEnterpriseResources(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEnterpriseResources)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enterprise_account_id", &obj.EnterpriseAccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "domain", &obj.Domain)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_iam_id", &obj.PrimaryContactIamID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary_contact_email", &obj.PrimaryContactEmail)
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

// ListEnterprisesOptions : The ListEnterprises options.
type ListEnterprisesOptions struct {
	// Get enterprises for a given enterprise account ID.
	EnterpriseAccountID *string `json:"enterprise_account_id,omitempty"`

	// Get enterprises for a given account group ID.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// Get enterprises for a given account ID.
	AccountID *string `json:"account_id,omitempty"`

	// Return results up to this limit. Valid values are between `0` and `100`.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListEnterprisesOptions : Instantiate ListEnterprisesOptions
func (*EnterpriseManagementV1) NewListEnterprisesOptions() *ListEnterprisesOptions {
	return &ListEnterprisesOptions{}
}

// SetEnterpriseAccountID : Allow user to set EnterpriseAccountID
func (options *ListEnterprisesOptions) SetEnterpriseAccountID(enterpriseAccountID string) *ListEnterprisesOptions {
	options.EnterpriseAccountID = core.StringPtr(enterpriseAccountID)
	return options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *ListEnterprisesOptions) SetAccountGroupID(accountGroupID string) *ListEnterprisesOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *ListEnterprisesOptions) SetAccountID(accountID string) *ListEnterprisesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListEnterprisesOptions) SetLimit(limit int64) *ListEnterprisesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListEnterprisesOptions) SetHeaders(param map[string]string) *ListEnterprisesOptions {
	options.Headers = param
	return options
}

// ListEnterprisesResponse : The response from calling list enterprises.
type ListEnterprisesResponse struct {
	// The number of enterprises returned from calling list enterprise.
	RowsCount *int64 `json:"rows_count,omitempty"`

	// A string that represents the link to the next page of results.
	NextURL *string `json:"next_url,omitempty"`

	// A list of enterprise objects.
	Resources []ListEnterpriseResources `json:"resources,omitempty"`
}


// UnmarshalListEnterprisesResponse unmarshals an instance of ListEnterprisesResponse from the specified map of raw messages.
func UnmarshalListEnterprisesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListEnterprisesResponse)
	err = core.UnmarshalPrimitive(m, "rows_count", &obj.RowsCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalListEnterpriseResources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAccountGroupOptions : The UpdateAccountGroup options.
type UpdateAccountGroupOptions struct {
	// The ID of the account group to retrieve.
	AccountGroupID *string `json:"account_group_id" validate:"required"`

	// The new name of the account group. This field must have 3 - 60 characters.
	Name *string `json:"name,omitempty"`

	// The IAM ID of the user to be the new primary contact for the account group.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccountGroupOptions : Instantiate UpdateAccountGroupOptions
func (*EnterpriseManagementV1) NewUpdateAccountGroupOptions(accountGroupID string) *UpdateAccountGroupOptions {
	return &UpdateAccountGroupOptions{
		AccountGroupID: core.StringPtr(accountGroupID),
	}
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *UpdateAccountGroupOptions) SetAccountGroupID(accountGroupID string) *UpdateAccountGroupOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateAccountGroupOptions) SetName(name string) *UpdateAccountGroupOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetPrimaryContactIamID : Allow user to set PrimaryContactIamID
func (options *UpdateAccountGroupOptions) SetPrimaryContactIamID(primaryContactIamID string) *UpdateAccountGroupOptions {
	options.PrimaryContactIamID = core.StringPtr(primaryContactIamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountGroupOptions) SetHeaders(param map[string]string) *UpdateAccountGroupOptions {
	options.Headers = param
	return options
}

// UpdateAccountOptions : The UpdateAccount options.
type UpdateAccountOptions struct {
	// The ID of the account to retrieve.
	AccountID *string `json:"account_id" validate:"required"`

	// The CRN of the new parent within the enterprise.
	Parent *string `json:"parent" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccountOptions : Instantiate UpdateAccountOptions
func (*EnterpriseManagementV1) NewUpdateAccountOptions(accountID string, parent string) *UpdateAccountOptions {
	return &UpdateAccountOptions{
		AccountID: core.StringPtr(accountID),
		Parent: core.StringPtr(parent),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateAccountOptions) SetAccountID(accountID string) *UpdateAccountOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetParent : Allow user to set Parent
func (options *UpdateAccountOptions) SetParent(parent string) *UpdateAccountOptions {
	options.Parent = core.StringPtr(parent)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountOptions) SetHeaders(param map[string]string) *UpdateAccountOptions {
	options.Headers = param
	return options
}

// UpdateEnterpriseOptions : The UpdateEnterprise options.
type UpdateEnterpriseOptions struct {
	// The ID of the enterprise to retrieve.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// The new name of the enterprise. This field must have 3 - 60 characters.
	Name *string `json:"name,omitempty"`

	// The new domain of the enterprise. This field has a limit of 60 characters.
	Domain *string `json:"domain,omitempty"`

	// The IAM ID of the user to be the new primary contact for the enterprise.
	PrimaryContactIamID *string `json:"primary_contact_iam_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateEnterpriseOptions : Instantiate UpdateEnterpriseOptions
func (*EnterpriseManagementV1) NewUpdateEnterpriseOptions(enterpriseID string) *UpdateEnterpriseOptions {
	return &UpdateEnterpriseOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *UpdateEnterpriseOptions) SetEnterpriseID(enterpriseID string) *UpdateEnterpriseOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateEnterpriseOptions) SetName(name string) *UpdateEnterpriseOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDomain : Allow user to set Domain
func (options *UpdateEnterpriseOptions) SetDomain(domain string) *UpdateEnterpriseOptions {
	options.Domain = core.StringPtr(domain)
	return options
}

// SetPrimaryContactIamID : Allow user to set PrimaryContactIamID
func (options *UpdateEnterpriseOptions) SetPrimaryContactIamID(primaryContactIamID string) *UpdateEnterpriseOptions {
	options.PrimaryContactIamID = core.StringPtr(primaryContactIamID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateEnterpriseOptions) SetHeaders(param map[string]string) *UpdateEnterpriseOptions {
	options.Headers = param
	return options
}
