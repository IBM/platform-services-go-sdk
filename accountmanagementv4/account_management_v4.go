/**
 * (C) Copyright IBM Corp. 2026.
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
 * IBM OpenAPI SDK Code Generator Version: 3.111.0-1bfb72c2-20260206-185521
 */

// Package accountmanagementv4 : Operations and models for the AccountManagementV4 service
package accountmanagementv4

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
)

// AccountManagementV4 : The Account Management API allows for the management of Account
//
// API Version: 4.0.0
type AccountManagementV4 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://accounts.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "account_management"

// AccountManagementV4Options : Service options
type AccountManagementV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAccountManagementV4UsingExternalConfig : constructs an instance of AccountManagementV4 with passed in options and external configuration.
func NewAccountManagementV4UsingExternalConfig(options *AccountManagementV4Options) (accountManagement *AccountManagementV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			err = core.SDKErrorf(err, "", "env-auth-error", common.GetComponentInfo())
			return
		}
	}

	accountManagement, err = NewAccountManagementV4(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = accountManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = accountManagement.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewAccountManagementV4 : constructs an instance of AccountManagementV4 with passed in options.
func NewAccountManagementV4(options *AccountManagementV4Options) (service *AccountManagementV4, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		err = core.SDKErrorf(err, "", "new-base-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			err = core.SDKErrorf(err, "", "set-url-error", common.GetComponentInfo())
			return
		}
	}

	service = &AccountManagementV4{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "accountManagement" suitable for processing requests.
func (accountManagement *AccountManagementV4) Clone() *AccountManagementV4 {
	if core.IsNil(accountManagement) {
		return nil
	}
	clone := *accountManagement
	clone.Service = accountManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (accountManagement *AccountManagementV4) SetServiceURL(url string) error {
	err := accountManagement.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (accountManagement *AccountManagementV4) GetServiceURL() string {
	return accountManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (accountManagement *AccountManagementV4) SetDefaultHeaders(headers http.Header) {
	accountManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (accountManagement *AccountManagementV4) SetEnableGzipCompression(enableGzip bool) {
	accountManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (accountManagement *AccountManagementV4) GetEnableGzipCompression() bool {
	return accountManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (accountManagement *AccountManagementV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	accountManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (accountManagement *AccountManagementV4) DisableRetries() {
	accountManagement.Service.DisableRetries()
}

// GetAccount : Get Account by Account ID
// Returns the details of an account.
func (accountManagement *AccountManagementV4) GetAccount(getAccountOptions *GetAccountOptions) (result *AccountResponse, response *core.DetailedResponse, err error) {
	result, response, err = accountManagement.GetAccountWithContext(context.Background(), getAccountOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAccountWithContext is an alternate form of the GetAccount method which supports a Context parameter
func (accountManagement *AccountManagementV4) GetAccountWithContext(ctx context.Context, getAccountOptions *GetAccountOptions) (result *AccountResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountOptions, "getAccountOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getAccountOptions, "getAccountOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getAccountOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = accountManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(accountManagement.Service.Options.URL, `/v4/accounts/{account_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("account_management", "V4", "GetAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = accountManagement.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "getAccount", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountResponse)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "4.0.0")
}

// AccountResponseTraits : AccountResponseTraits struct
type AccountResponseTraits struct {
	EuSupported *bool `json:"eu_supported" validate:"required"`

	Poc *bool `json:"poc" validate:"required"`

	Hippa *bool `json:"hippa" validate:"required"`
}

// UnmarshalAccountResponseTraits unmarshals an instance of AccountResponseTraits from the specified map of raw messages.
func UnmarshalAccountResponseTraits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountResponseTraits)
	err = core.UnmarshalPrimitive(m, "eu_supported", &obj.EuSupported)
	if err != nil {
		err = core.SDKErrorf(err, "", "eu_supported-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "poc", &obj.Poc)
	if err != nil {
		err = core.SDKErrorf(err, "", "poc-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "hippa", &obj.Hippa)
	if err != nil {
		err = core.SDKErrorf(err, "", "hippa-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAccountOptions : The GetAccount options.
type GetAccountOptions struct {
	// The unique identifier of the account you want to retrieve.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAccountOptions : Instantiate GetAccountOptions
func (*AccountManagementV4) NewGetAccountOptions(accountID string) *GetAccountOptions {
	return &GetAccountOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetAccountOptions) SetAccountID(accountID string) *GetAccountOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountOptions) SetHeaders(param map[string]string) *GetAccountOptions {
	options.Headers = param
	return options
}

// AccountResponse : AccountResponse struct
type AccountResponse struct {
	Name *string `json:"name" validate:"required"`

	ID *string `json:"id" validate:"required"`

	Owner *string `json:"owner" validate:"required"`

	OwnerUserid *string `json:"owner_userid" validate:"required"`

	OwnerIamid *string `json:"owner_iamid" validate:"required"`

	Type *string `json:"type" validate:"required"`

	Status *string `json:"status" validate:"required"`

	LinkedSoftlayerAccount *string `json:"linked_softlayer_account" validate:"required"`

	TeamDirectoryEnabled bool `json:"team_directory_enabled" validate:"required"`

	Traits *AccountResponseTraits `json:"traits" validate:"required"`
}

// UnmarshalAccountResponse unmarshals an instance of AccountResponse from the specified map of raw messages.
func UnmarshalAccountResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		err = core.SDKErrorf(err, "", "name-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		err = core.SDKErrorf(err, "", "owner-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_userid", &obj.OwnerUserid)
	if err != nil {
		err = core.SDKErrorf(err, "", "owner_userid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_iamid", &obj.OwnerIamid)
	if err != nil {
		err = core.SDKErrorf(err, "", "owner_iamid-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		err = core.SDKErrorf(err, "", "type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		err = core.SDKErrorf(err, "", "status-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "linked_softlayer_account", &obj.LinkedSoftlayerAccount)
	if err != nil {
		err = core.SDKErrorf(err, "", "linked_softlayer_account-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "team_directory_enabled", &obj.TeamDirectoryEnabled)
	if err != nil {
		err = core.SDKErrorf(err, "", "team_directory_enabled-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "traits", &obj.Traits, UnmarshalAccountResponseTraits)
	if err != nil {
		err = core.SDKErrorf(err, "", "traits-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
