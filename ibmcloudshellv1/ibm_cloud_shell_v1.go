/**
 * (C) Copyright IBM Corp. 2021.
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
 * IBM OpenAPI SDK Code Generator Version: 3.33.0-caf29bd0-20210603-225214
 */

// Package ibmcloudshellv1 : Operations and models for the IbmCloudShellV1 service
package ibmcloudshellv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
)

// IbmCloudShellV1 : API docs for IBM Cloud Shell repository
//
// Version: 1.0
type IbmCloudShellV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.shell.test.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "ibm_cloud_shell"

// IbmCloudShellV1Options : Service options
type IbmCloudShellV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewIbmCloudShellV1UsingExternalConfig : constructs an instance of IbmCloudShellV1 with passed in options and external configuration.
func NewIbmCloudShellV1UsingExternalConfig(options *IbmCloudShellV1Options) (ibmCloudShell *IbmCloudShellV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	ibmCloudShell, err = NewIbmCloudShellV1(options)
	if err != nil {
		return
	}

	err = ibmCloudShell.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = ibmCloudShell.Service.SetServiceURL(options.URL)
	}
	return
}

// NewIbmCloudShellV1 : constructs an instance of IbmCloudShellV1 with passed in options.
func NewIbmCloudShellV1(options *IbmCloudShellV1Options) (service *IbmCloudShellV1, err error) {
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

	service = &IbmCloudShellV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "ibmCloudShell" suitable for processing requests.
func (ibmCloudShell *IbmCloudShellV1) Clone() *IbmCloudShellV1 {
	if core.IsNil(ibmCloudShell) {
		return nil
	}
	clone := *ibmCloudShell
	clone.Service = ibmCloudShell.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (ibmCloudShell *IbmCloudShellV1) SetServiceURL(url string) error {
	return ibmCloudShell.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (ibmCloudShell *IbmCloudShellV1) GetServiceURL() string {
	return ibmCloudShell.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (ibmCloudShell *IbmCloudShellV1) SetDefaultHeaders(headers http.Header) {
	ibmCloudShell.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (ibmCloudShell *IbmCloudShellV1) SetEnableGzipCompression(enableGzip bool) {
	ibmCloudShell.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (ibmCloudShell *IbmCloudShellV1) GetEnableGzipCompression() bool {
	return ibmCloudShell.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (ibmCloudShell *IbmCloudShellV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	ibmCloudShell.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (ibmCloudShell *IbmCloudShellV1) DisableRetries() {
	ibmCloudShell.Service.DisableRetries()
}

// GetAccountSettingsByID : Get account settings
// Retrieve account settings for the given account ID. Call this method to get details about a particular account
// setting, whether Cloud Shell is enabled, the list of enabled regions and the list of enabled features. Users need to
// be an account owner or users need to be assigned an IAM policy with the Administrator role for the Cloud Shell
// account management service.
func (ibmCloudShell *IbmCloudShellV1) GetAccountSettingsByID(getAccountSettingsByIdOptions *GetAccountSettingsByIdOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	return ibmCloudShell.GetAccountSettingsByIDWithContext(context.Background(), getAccountSettingsByIdOptions)
}

// GetAccountSettingsByIDWithContext is an alternate form of the GetAccountSettingsByID method which supports a Context parameter
func (ibmCloudShell *IbmCloudShellV1) GetAccountSettingsByIDWithContext(ctx context.Context, getAccountSettingsByIdOptions *GetAccountSettingsByIdOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSettingsByIdOptions, "getAccountSettingsByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSettingsByIdOptions, "getAccountSettingsByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getAccountSettingsByIdOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudShell.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudShell.Service.Options.URL, `/api/v1/user/accounts/{account_id}/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountSettingsByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_shell", "V1", "GetAccountSettingsByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = ibmCloudShell.Service.Request(request, &rawResponse)
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

// UpdateAccountSettingsByID : Update account settings
// Update account settings for the given account ID. Call this method to update account settings configuration, you can
// enable or disable Cloud Shell, enable or disable available regions and enable and disable features. To update account
// settings, users need to be an account owner or users need to be assigned an IAM policy with the Administrator role
// for the Cloud Shell account management service.
func (ibmCloudShell *IbmCloudShellV1) UpdateAccountSettingsByID(updateAccountSettingsByIdOptions *UpdateAccountSettingsByIdOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	return ibmCloudShell.UpdateAccountSettingsByIDWithContext(context.Background(), updateAccountSettingsByIdOptions)
}

// UpdateAccountSettingsByIDWithContext is an alternate form of the UpdateAccountSettingsByID method which supports a Context parameter
func (ibmCloudShell *IbmCloudShellV1) UpdateAccountSettingsByIDWithContext(ctx context.Context, updateAccountSettingsByIdOptions *UpdateAccountSettingsByIdOptions) (result *AccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAccountSettingsByIdOptions, "updateAccountSettingsByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAccountSettingsByIdOptions, "updateAccountSettingsByIdOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *updateAccountSettingsByIdOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = ibmCloudShell.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(ibmCloudShell.Service.Options.URL, `/api/v1/user/accounts/{account_id}/settings`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAccountSettingsByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("ibm_cloud_shell", "V1", "UpdateAccountSettingsByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateAccountSettingsByIdOptions.NewID != nil {
		body["_id"] = updateAccountSettingsByIdOptions.NewID
	}
	if updateAccountSettingsByIdOptions.NewRev != nil {
		body["_rev"] = updateAccountSettingsByIdOptions.NewRev
	}
	if updateAccountSettingsByIdOptions.NewAccountID != nil {
		body["account_id"] = updateAccountSettingsByIdOptions.NewAccountID
	}
	if updateAccountSettingsByIdOptions.NewCreatedAt != nil {
		body["created_at"] = updateAccountSettingsByIdOptions.NewCreatedAt
	}
	if updateAccountSettingsByIdOptions.NewCreatedBy != nil {
		body["created_by"] = updateAccountSettingsByIdOptions.NewCreatedBy
	}
	if updateAccountSettingsByIdOptions.NewDefaultEnableNewFeatures != nil {
		body["default_enable_new_features"] = updateAccountSettingsByIdOptions.NewDefaultEnableNewFeatures
	}
	if updateAccountSettingsByIdOptions.NewDefaultEnableNewRegions != nil {
		body["default_enable_new_regions"] = updateAccountSettingsByIdOptions.NewDefaultEnableNewRegions
	}
	if updateAccountSettingsByIdOptions.NewEnabled != nil {
		body["enabled"] = updateAccountSettingsByIdOptions.NewEnabled
	}
	if updateAccountSettingsByIdOptions.NewFeatures != nil {
		body["features"] = updateAccountSettingsByIdOptions.NewFeatures
	}
	if updateAccountSettingsByIdOptions.NewRegions != nil {
		body["regions"] = updateAccountSettingsByIdOptions.NewRegions
	}
	if updateAccountSettingsByIdOptions.NewType != nil {
		body["type"] = updateAccountSettingsByIdOptions.NewType
	}
	if updateAccountSettingsByIdOptions.NewUpdatedAt != nil {
		body["updated_at"] = updateAccountSettingsByIdOptions.NewUpdatedAt
	}
	if updateAccountSettingsByIdOptions.NewUpdatedBy != nil {
		body["updated_by"] = updateAccountSettingsByIdOptions.NewUpdatedBy
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
	response, err = ibmCloudShell.Service.Request(request, &rawResponse)
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

// AccountSettings : Definition of Cloud Shell account settings.
type AccountSettings struct {
	// Unique id of the settings object.
	ID *string `json:"_id,omitempty"`

	// Unique revision number for the settings object.
	Rev *string `json:"_rev,omitempty"`

	// The id of the account the settings belong to.
	AccountID *string `json:"account_id,omitempty"`

	// Creation timestamp.
	CreatedAt *int64 `json:"created_at,omitempty"`

	// IAM ID of creator.
	CreatedBy *string `json:"created_by,omitempty"`

	// You can choose which Cloud Shell features are available in the account and whether any new features are enabled as
	// they become available. The feature settings apply only to the enabled Cloud Shell locations.
	DefaultEnableNewFeatures *bool `json:"default_enable_new_features,omitempty"`

	// Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and
	// session data are stored. By default, users are routed to the nearest available location.
	DefaultEnableNewRegions *bool `json:"default_enable_new_regions,omitempty"`

	// When enabled, Cloud Shell is available to all users in the account.
	Enabled *bool `json:"enabled,omitempty"`

	// List of Cloud Shell features.
	Features []Feature `json:"features,omitempty"`

	// List of Cloud Shell region settings.
	Regions []RegionSetting `json:"regions,omitempty"`

	// Type of api response object.
	Type *string `json:"type,omitempty"`

	// Timestamp of last update.
	UpdatedAt *int64 `json:"updated_at,omitempty"`

	// IAM ID of last updater.
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// UnmarshalAccountSettings unmarshals an instance of AccountSettings from the specified map of raw messages.
func UnmarshalAccountSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountSettings)
	err = core.UnmarshalPrimitive(m, "_id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "_rev", &obj.Rev)
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
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_enable_new_features", &obj.DefaultEnableNewFeatures)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_enable_new_regions", &obj.DefaultEnableNewRegions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "regions", &obj.Regions, UnmarshalRegionSetting)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// Feature : Describes a Cloud Shell feature.
type Feature struct {
	// State of the feature.
	Enabled *bool `json:"enabled,omitempty"`

	// Name of the feature.
	Key *string `json:"key,omitempty"`
}

// UnmarshalFeature unmarshals an instance of Feature from the specified map of raw messages.
func UnmarshalFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feature)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAccountSettingsByIdOptions : The GetAccountSettingsByID options.
type GetAccountSettingsByIdOptions struct {
	// The account ID in which the account settings belong to.
	AccountID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountSettingsByIdOptions : Instantiate GetAccountSettingsByIdOptions
func (*IbmCloudShellV1) NewGetAccountSettingsByIdOptions(accountID string) *GetAccountSettingsByIdOptions {
	return &GetAccountSettingsByIdOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetAccountSettingsByIdOptions) SetAccountID(accountID string) *GetAccountSettingsByIdOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountSettingsByIdOptions) SetHeaders(param map[string]string) *GetAccountSettingsByIdOptions {
	options.Headers = param
	return options
}

// RegionSetting : Describes a Cloud Shell region setting.
type RegionSetting struct {
	// State of the region.
	Enabled *bool `json:"enabled,omitempty"`

	// Name of the region.
	Key *string `json:"key,omitempty"`
}

// UnmarshalRegionSetting unmarshals an instance of RegionSetting from the specified map of raw messages.
func UnmarshalRegionSetting(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegionSetting)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAccountSettingsByIdOptions : The UpdateAccountSettingsByID options.
type UpdateAccountSettingsByIdOptions struct {
	// The account ID in which the account settings belong to.
	AccountID *string `validate:"required,ne="`

	// Unique id of the settings object.
	NewID *string

	// Unique revision number for the settings object.
	NewRev *string

	// The id of the account the settings belong to.
	NewAccountID *string

	// Creation timestamp.
	NewCreatedAt *int64

	// IAM ID of creator.
	NewCreatedBy *string

	// You can choose which Cloud Shell features are available in the account and whether any new features are enabled as
	// they become available. The feature settings apply only to the enabled Cloud Shell locations.
	NewDefaultEnableNewFeatures *bool

	// Set whether Cloud Shell is enabled in a specific location for the account. The location determines where user and
	// session data are stored. By default, users are routed to the nearest available location.
	NewDefaultEnableNewRegions *bool

	// When enabled, Cloud Shell is available to all users in the account.
	NewEnabled *bool

	// List of Cloud Shell features.
	NewFeatures []Feature

	// List of Cloud Shell region settings.
	NewRegions []RegionSetting

	// Type of api response object.
	NewType *string

	// Timestamp of last update.
	NewUpdatedAt *int64

	// IAM ID of last updater.
	NewUpdatedBy *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAccountSettingsByIdOptions : Instantiate UpdateAccountSettingsByIdOptions
func (*IbmCloudShellV1) NewUpdateAccountSettingsByIdOptions(accountID string) *UpdateAccountSettingsByIdOptions {
	return &UpdateAccountSettingsByIdOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *UpdateAccountSettingsByIdOptions) SetAccountID(accountID string) *UpdateAccountSettingsByIdOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetNewID : Allow user to set NewID
func (_options *UpdateAccountSettingsByIdOptions) SetNewID(newID string) *UpdateAccountSettingsByIdOptions {
	_options.NewID = core.StringPtr(newID)
	return _options
}

// SetNewRev : Allow user to set NewRev
func (_options *UpdateAccountSettingsByIdOptions) SetNewRev(newRev string) *UpdateAccountSettingsByIdOptions {
	_options.NewRev = core.StringPtr(newRev)
	return _options
}

// SetNewAccountID : Allow user to set NewAccountID
func (_options *UpdateAccountSettingsByIdOptions) SetNewAccountID(newAccountID string) *UpdateAccountSettingsByIdOptions {
	_options.NewAccountID = core.StringPtr(newAccountID)
	return _options
}

// SetNewCreatedAt : Allow user to set NewCreatedAt
func (_options *UpdateAccountSettingsByIdOptions) SetNewCreatedAt(newCreatedAt int64) *UpdateAccountSettingsByIdOptions {
	_options.NewCreatedAt = core.Int64Ptr(newCreatedAt)
	return _options
}

// SetNewCreatedBy : Allow user to set NewCreatedBy
func (_options *UpdateAccountSettingsByIdOptions) SetNewCreatedBy(newCreatedBy string) *UpdateAccountSettingsByIdOptions {
	_options.NewCreatedBy = core.StringPtr(newCreatedBy)
	return _options
}

// SetNewDefaultEnableNewFeatures : Allow user to set NewDefaultEnableNewFeatures
func (_options *UpdateAccountSettingsByIdOptions) SetNewDefaultEnableNewFeatures(newDefaultEnableNewFeatures bool) *UpdateAccountSettingsByIdOptions {
	_options.NewDefaultEnableNewFeatures = core.BoolPtr(newDefaultEnableNewFeatures)
	return _options
}

// SetNewDefaultEnableNewRegions : Allow user to set NewDefaultEnableNewRegions
func (_options *UpdateAccountSettingsByIdOptions) SetNewDefaultEnableNewRegions(newDefaultEnableNewRegions bool) *UpdateAccountSettingsByIdOptions {
	_options.NewDefaultEnableNewRegions = core.BoolPtr(newDefaultEnableNewRegions)
	return _options
}

// SetNewEnabled : Allow user to set NewEnabled
func (_options *UpdateAccountSettingsByIdOptions) SetNewEnabled(newEnabled bool) *UpdateAccountSettingsByIdOptions {
	_options.NewEnabled = core.BoolPtr(newEnabled)
	return _options
}

// SetNewFeatures : Allow user to set NewFeatures
func (_options *UpdateAccountSettingsByIdOptions) SetNewFeatures(newFeatures []Feature) *UpdateAccountSettingsByIdOptions {
	_options.NewFeatures = newFeatures
	return _options
}

// SetNewRegions : Allow user to set NewRegions
func (_options *UpdateAccountSettingsByIdOptions) SetNewRegions(newRegions []RegionSetting) *UpdateAccountSettingsByIdOptions {
	_options.NewRegions = newRegions
	return _options
}

// SetNewType : Allow user to set NewType
func (_options *UpdateAccountSettingsByIdOptions) SetNewType(newType string) *UpdateAccountSettingsByIdOptions {
	_options.NewType = core.StringPtr(newType)
	return _options
}

// SetNewUpdatedAt : Allow user to set NewUpdatedAt
func (_options *UpdateAccountSettingsByIdOptions) SetNewUpdatedAt(newUpdatedAt int64) *UpdateAccountSettingsByIdOptions {
	_options.NewUpdatedAt = core.Int64Ptr(newUpdatedAt)
	return _options
}

// SetNewUpdatedBy : Allow user to set NewUpdatedBy
func (_options *UpdateAccountSettingsByIdOptions) SetNewUpdatedBy(newUpdatedBy string) *UpdateAccountSettingsByIdOptions {
	_options.NewUpdatedBy = core.StringPtr(newUpdatedBy)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAccountSettingsByIdOptions) SetHeaders(param map[string]string) *UpdateAccountSettingsByIdOptions {
	options.Headers = param
	return options
}
