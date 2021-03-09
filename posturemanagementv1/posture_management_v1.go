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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-e6cfc86e-20210308-121737
 */

// Package posturemanagementv1 : Operations and models for the PostureManagementV1 service
package posturemanagementv1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"net/http"
	"reflect"
	"time"
)

// PostureManagementV1 : With IBM Cloud® Security and Compliance Center, you can embed checks into your every day
// workflows to help manage your current security and compliance posture. By monitoring for risks, you can identify
// security vulnerabilities and quickly work to mitigate the impact.
//
// Version: 1.0.0
type PostureManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "posture_management"

// PostureManagementV1Options : Service options
type PostureManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPostureManagementV1UsingExternalConfig : constructs an instance of PostureManagementV1 with passed in options and external configuration.
func NewPostureManagementV1UsingExternalConfig(options *PostureManagementV1Options) (postureManagement *PostureManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	postureManagement, err = NewPostureManagementV1(options)
	if err != nil {
		return
	}

	err = postureManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = postureManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewPostureManagementV1 : constructs an instance of PostureManagementV1 with passed in options.
func NewPostureManagementV1(options *PostureManagementV1Options) (service *PostureManagementV1, err error) {
	serviceOptions := &core.ServiceOptions{
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

	service = &PostureManagementV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "postureManagement" suitable for processing requests.
func (postureManagement *PostureManagementV1) Clone() *PostureManagementV1 {
	if core.IsNil(postureManagement) {
		return nil
	}
	clone := *postureManagement
	clone.Service = postureManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (postureManagement *PostureManagementV1) SetServiceURL(url string) error {
	return postureManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (postureManagement *PostureManagementV1) GetServiceURL() string {
	return postureManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (postureManagement *PostureManagementV1) SetDefaultHeaders(headers http.Header) {
	postureManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (postureManagement *PostureManagementV1) SetEnableGzipCompression(enableGzip bool) {
	postureManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (postureManagement *PostureManagementV1) GetEnableGzipCompression() bool {
	return postureManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (postureManagement *PostureManagementV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	postureManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (postureManagement *PostureManagementV1) DisableRetries() {
	postureManagement.Service.DisableRetries()
}

// CreateValidationScan : Initiate a validation scan
// Validation scans determine a specified scope's adherence to regulatory controls by validating the configuration of
// the resources in your scope to the attached profile. To initiate a scan, you must have configured a collector,
// provided credentials, and completed both a fact collection and discovery scan. [Learn
// more](/docs/security-compliance?topic=security-compliance-schedule-scan).
func (postureManagement *PostureManagementV1) CreateValidationScan(createValidationScanOptions *CreateValidationScanOptions) (result *Result, response *core.DetailedResponse, err error) {
	return postureManagement.CreateValidationScanWithContext(context.Background(), createValidationScanOptions)
}

// CreateValidationScanWithContext is an alternate form of the CreateValidationScan method which supports a Context parameter
func (postureManagement *PostureManagementV1) CreateValidationScanWithContext(ctx context.Context, createValidationScanOptions *CreateValidationScanOptions) (result *Result, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createValidationScanOptions, "createValidationScanOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createValidationScanOptions, "createValidationScanOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scans/validation`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createValidationScanOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "CreateValidationScan")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("account_id", fmt.Sprint(*createValidationScanOptions.AccountID))

	body := make(map[string]interface{})
	if createValidationScanOptions.ScopeID != nil {
		body["scope_id"] = createValidationScanOptions.ScopeID
	}
	if createValidationScanOptions.ProfileID != nil {
		body["profile_id"] = createValidationScanOptions.ProfileID
	}
	if createValidationScanOptions.GroupProfileID != nil {
		body["group_profile_id"] = createValidationScanOptions.GroupProfileID
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
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListProfile : List profiles
// List all of the profiles that are available in your account. To view a specific profile, you can filter by name.
func (postureManagement *PostureManagementV1) ListProfile(listProfileOptions *ListProfileOptions) (result *ProfilesList, response *core.DetailedResponse, err error) {
	return postureManagement.ListProfileWithContext(context.Background(), listProfileOptions)
}

// ListProfileWithContext is an alternate form of the ListProfile method which supports a Context parameter
func (postureManagement *PostureManagementV1) ListProfileWithContext(ctx context.Context, listProfileOptions *ListProfileOptions) (result *ProfilesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listProfileOptions, "listProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listProfileOptions, "listProfileOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "ListProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("account_id", fmt.Sprint(*listProfileOptions.AccountID))
	if listProfileOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listProfileOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalProfilesList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListScopes : List scopes
// List all of the scopes that are available in your account. To view a specific scope, you can filter by name.
func (postureManagement *PostureManagementV1) ListScopes(listScopesOptions *ListScopesOptions) (result *ScopesList, response *core.DetailedResponse, err error) {
	return postureManagement.ListScopesWithContext(context.Background(), listScopesOptions)
}

// ListScopesWithContext is an alternate form of the ListScopes method which supports a Context parameter
func (postureManagement *PostureManagementV1) ListScopesWithContext(ctx context.Context, listScopesOptions *ListScopesOptions) (result *ScopesList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listScopesOptions, "listScopesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listScopesOptions, "listScopesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = postureManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(postureManagement.Service.Options.URL, `/posture/v1/scopes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listScopesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("posture_management", "V1", "ListScopes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("account_id", fmt.Sprint(*listScopesOptions.AccountID))
	if listScopesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listScopesOptions.Name))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = postureManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScopesList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ApplicabilityCriteria : The criteria that defines how a profile applies.
type ApplicabilityCriteria struct {
	// A list of environments that a profile can be applied to.
	Environment []string `json:"environment,omitempty"`

	// A list of resources that a profile can be used with.
	Resource []string `json:"resource,omitempty"`

	// The type of environment that a profile is able to be applied to.
	EnvironmentCategory []string `json:"environment_category,omitempty"`

	// The type of resource that a profile is able to be applied to.
	ResourceCategory []string `json:"resource_category,omitempty"`

	// The resource type that the profile applies to.
	ResourceType []string `json:"resource_type,omitempty"`

	// The software that the profile applies to.
	SoftwareDetails interface{} `json:"software_details,omitempty"`

	// The operatoring system that the profile applies to.
	OsDetails interface{} `json:"os_details,omitempty"`

	// Any additional details about the profile.
	AdditionalDetails interface{} `json:"additional_details,omitempty"`

	// The type of environment that your scope is targeted to.
	EnvironmentCategoryDescription map[string]string `json:"environment_category_description,omitempty"`

	// The environment that your scope is targeted to.
	EnvironmentDescription map[string]string `json:"environment_description,omitempty"`

	// The type of resource that your scope is targeted to.
	ResourceCategoryDescription map[string]string `json:"resource_category_description,omitempty"`

	// A further classification of the type of resource that your scope is targeted to.
	ResourceTypeDescription map[string]string `json:"resource_type_description,omitempty"`

	// The resource that is scanned as part of your scope.
	ResourceDescription map[string]string `json:"resource_description,omitempty"`
}

// UnmarshalApplicabilityCriteria unmarshals an instance of ApplicabilityCriteria from the specified map of raw messages.
func UnmarshalApplicabilityCriteria(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicabilityCriteria)
	err = core.UnmarshalPrimitive(m, "environment", &obj.Environment)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource", &obj.Resource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_category", &obj.EnvironmentCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_category", &obj.ResourceCategory)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "software_details", &obj.SoftwareDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "os_details", &obj.OsDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "additional_details", &obj.AdditionalDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_category_description", &obj.EnvironmentCategoryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_description", &obj.EnvironmentDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_category_description", &obj.ResourceCategoryDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type_description", &obj.ResourceTypeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_description", &obj.ResourceDescription)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateValidationScanOptions : The CreateValidationScan options.
type CreateValidationScanOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique ID of the scope.
	ScopeID *int64

	// The unique ID of the profile.
	ProfileID *int64

	// The ID of the profile group.
	GroupProfileID *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateValidationScanOptions : Instantiate CreateValidationScanOptions
func (*PostureManagementV1) NewCreateValidationScanOptions(accountID string) *CreateValidationScanOptions {
	return &CreateValidationScanOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *CreateValidationScanOptions) SetAccountID(accountID string) *CreateValidationScanOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetScopeID : Allow user to set ScopeID
func (options *CreateValidationScanOptions) SetScopeID(scopeID int64) *CreateValidationScanOptions {
	options.ScopeID = core.Int64Ptr(scopeID)
	return options
}

// SetProfileID : Allow user to set ProfileID
func (options *CreateValidationScanOptions) SetProfileID(profileID int64) *CreateValidationScanOptions {
	options.ProfileID = core.Int64Ptr(profileID)
	return options
}

// SetGroupProfileID : Allow user to set GroupProfileID
func (options *CreateValidationScanOptions) SetGroupProfileID(groupProfileID int64) *CreateValidationScanOptions {
	options.GroupProfileID = core.Int64Ptr(groupProfileID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateValidationScanOptions) SetHeaders(param map[string]string) *CreateValidationScanOptions {
	options.Headers = param
	return options
}

// ListProfileOptions : The ListProfile options.
type ListProfileOptions struct {
	// Your account ID.
	AccountID *string `validate:"required"`

	// The name of the profile.
	Name *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListProfileOptions : Instantiate ListProfileOptions
func (*PostureManagementV1) NewListProfileOptions(accountID string) *ListProfileOptions {
	return &ListProfileOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListProfileOptions) SetAccountID(accountID string) *ListProfileOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *ListProfileOptions) SetName(name string) *ListProfileOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListProfileOptions) SetHeaders(param map[string]string) *ListProfileOptions {
	options.Headers = param
	return options
}

// ListScopesOptions : The ListScopes options.
type ListScopesOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The name of the scope.
	Name *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListScopesOptions : Instantiate ListScopesOptions
func (*PostureManagementV1) NewListScopesOptions(accountID string) *ListScopesOptions {
	return &ListScopesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListScopesOptions) SetAccountID(accountID string) *ListScopesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *ListScopesOptions) SetName(name string) *ListScopesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListScopesOptions) SetHeaders(param map[string]string) *ListScopesOptions {
	options.Headers = param
	return options
}

// Profile : Profile.
type Profile struct {
	// The name of the profile.
	Name *string `json:"name,omitempty"`

	// The number of goals that are in the profile.
	NoOfGoals *int64 `json:"no_of_goals,omitempty"`

	// A description of the profile.
	Description *string `json:"description,omitempty"`

	// The version of the profile.
	Version *int64 `json:"version,omitempty"`

	// The user who created the profile.
	CreatedBy *string `json:"created_by,omitempty"`

	// The user who last modified the profile.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// A reason that you want to delete a profile.
	ReasonForDelete *string `json:"reason_for_delete,omitempty"`

	// The criteria that defines how a profile applies.
	ApplicabilityCriteria *ApplicabilityCriteria `json:"applicability_criteria,omitempty"`

	// An auto-generated unique identifying number of the profile.
	ProfileID *int64 `json:"profile_id,omitempty"`

	// The base profile that the controls are pulled from.
	BaseProfile *string `json:"base_profile,omitempty"`

	// The type of profile.
	ProfileType *string `json:"profile_type,omitempty"`

	// The time that the profile was created in UTC.
	CreatedTime *string `json:"created_time,omitempty"`

	// The time that the profile was most recently modified in UTC.
	ModifiedTime *string `json:"modified_time,omitempty"`

	// The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// Constants associated with the Profile.ProfileType property.
// The type of profile.
const (
	ProfileProfileTypeCustomConst = "custom"
	ProfileProfileTypeGroupConst = "group"
	ProfileProfileTypePredefinedConst = "predefined"
)

// UnmarshalProfile unmarshals an instance of Profile from the specified map of raw messages.
func UnmarshalProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Profile)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "no_of_goals", &obj.NoOfGoals)
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
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason_for_delete", &obj.ReasonForDelete)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "applicability_criteria", &obj.ApplicabilityCriteria, UnmarshalApplicabilityCriteria)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_id", &obj.ProfileID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base_profile", &obj.BaseProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "profile_type", &obj.ProfileType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_time", &obj.ModifiedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ProfilesList : A list of profiles.
type ProfilesList struct {
	// Profiles.
	Profiles []Profile `json:"profiles,omitempty"`
}

// UnmarshalProfilesList unmarshals an instance of ProfilesList from the specified map of raw messages.
func UnmarshalProfilesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ProfilesList)
	err = core.UnmarshalModel(m, "profiles", &obj.Profiles, UnmarshalProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Result : Result.
type Result struct {
	// Result.
	Result *bool `json:"result,omitempty"`

	// A message is returned.
	Message *string `json:"message,omitempty"`
}

// UnmarshalResult unmarshals an instance of Result from the specified map of raw messages.
func UnmarshalResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Result)
	err = core.UnmarshalPrimitive(m, "result", &obj.Result)
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

// Scan : Scan.
type Scan struct {
	// An auto-generated unique identifier for the scan.
	ScanID *int64 `json:"scan_id,omitempty"`

	// An auto-generated unique identifier for discovery.
	DiscoverID *int64 `json:"discover_id,omitempty"`

	// The status of the collector as it completes a scan.
	Status *string `json:"status,omitempty"`

	// The current status of the collector.
	StatusMessage *string `json:"status_message,omitempty"`
}

// Constants associated with the Scan.Status property.
// The status of the collector as it completes a scan.
const (
	ScanStatusAbortTaskRequestCompletedConst = "abort_task_request_completed"
	ScanStatusAbortTaskRequestFailedConst = "abort_task_request_failed"
	ScanStatusAbortTaskRequestReceivedConst = "abort_task_request_received"
	ScanStatusControllerAbortedConst = "controller_aborted"
	ScanStatusDiscoveryCompletedConst = "discovery_completed"
	ScanStatusDiscoveryInProgressConst = "discovery_in_progress"
	ScanStatusDiscoveryResultPostedNoErrorConst = "discovery_result_posted_no_error"
	ScanStatusDiscoveryResultPostedWithErrorConst = "discovery_result_posted_with_error"
	ScanStatusDiscoveryStartedConst = "discovery_started"
	ScanStatusErrorInAbortTaskRequestConst = "error_in_abort_task_request"
	ScanStatusErrorInDiscoveryConst = "error_in_discovery"
	ScanStatusErrorInFactCollectionConst = "error_in_fact_collection"
	ScanStatusErrorInFactValidationConst = "error_in_fact_validation"
	ScanStatusErrorInInventoryConst = "error_in_inventory"
	ScanStatusErrorInRemediationConst = "error_in_remediation"
	ScanStatusErrorInValidationConst = "error_in_validation"
	ScanStatusFactCollectionCompletedConst = "fact_collection_completed"
	ScanStatusFactCollectionInProgressConst = "fact_collection_in_progress"
	ScanStatusFactCollectionStartedConst = "fact_collection_started"
	ScanStatusFactValidationCompletedConst = "fact_validation_completed"
	ScanStatusFactValidationInProgressConst = "fact_validation_in_progress"
	ScanStatusFactValidationStartedConst = "fact_validation_started"
	ScanStatusGatewayAbortedConst = "gateway_aborted"
	ScanStatusInventoryCompletedConst = "inventory_completed"
	ScanStatusInventoryCompletedWithErrorConst = "inventory_completed_with_error"
	ScanStatusInventoryInProgressConst = "inventory_in_progress"
	ScanStatusInventoryStartedConst = "inventory_started"
	ScanStatusNotAcceptedConst = "not_accepted"
	ScanStatusPendingConst = "pending"
	ScanStatusRemediationCompletedConst = "remediation_completed"
	ScanStatusRemediationInProgressConst = "remediation_in_progress"
	ScanStatusRemediationStartedConst = "remediation_started"
	ScanStatusSentToCollectorConst = "sent_to_collector"
	ScanStatusUserAbortedConst = "user_aborted"
	ScanStatusValidationCompletedConst = "validation_completed"
	ScanStatusValidationInProgressConst = "validation_in_progress"
	ScanStatusValidationResultPostedNoErrorConst = "validation_result_posted_no_error"
	ScanStatusValidationResultPostedWithErrorConst = "validation_result_posted_with_error"
	ScanStatusValidationStartedConst = "validation_started"
	ScanStatusWaitingForRefineConst = "waiting_for_refine"
)

// UnmarshalScan unmarshals an instance of Scan from the specified map of raw messages.
func UnmarshalScan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scan)
	err = core.UnmarshalPrimitive(m, "scan_id", &obj.ScanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "discover_id", &obj.DiscoverID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_message", &obj.StatusMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Scope : Scope.
type Scope struct {
	// A detailed description of the scope.
	Description *string `json:"description,omitempty"`

	// The user who created the scope.
	CreatedBy *string `json:"created_by,omitempty"`

	// The user who most recently modified the scope.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// An auto-generated unique identifier for the scope.
	ScopeID *int64 `json:"scope_id,omitempty"`

	// A unique name for your scope.
	Name *string `json:"name,omitempty"`

	// Indicates whether scope is enabled/disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The environment that the scope is targeted to.
	EnvironmentType *string `json:"environment_type,omitempty"`

	// The time that the scope was created in UTC.
	CreatedTime *string `json:"created_time,omitempty"`

	// The time that the scope was last modified in UTC.
	ModifiedTime *string `json:"modified_time,omitempty"`

	// The last type of scan that was run on the scope.
	LastScanType *string `json:"last_scan_type,omitempty"`

	// A description of the last scan type.
	LastScanTypeDescription *string `json:"last_scan_type_description,omitempty"`

	// The last time that a scan status for a scope was updated in UTC.
	LastScanStatusUpdatedTime *string `json:"last_scan_status_updated_time,omitempty"`

	// The unique IDs of the collectors that are attached to the scope.
	CollectorsID []int64 `json:"collectors_id,omitempty"`

	// A list of the scans that have been run on the scope.
	Scans []Scan `json:"scans,omitempty"`
}

// Constants associated with the Scope.EnvironmentType property.
// The environment that the scope is targeted to.
const (
	ScopeEnvironmentTypeAwsConst = "aws"
	ScopeEnvironmentTypeAzureConst = "azure"
	ScopeEnvironmentTypeGcpConst = "gcp"
	ScopeEnvironmentTypeHostedConst = "hosted"
	ScopeEnvironmentTypeIBMConst = "ibm"
	ScopeEnvironmentTypeOnPremiseConst = "on_premise"
	ScopeEnvironmentTypeOpenstackConst = "openstack"
	ScopeEnvironmentTypeServicesConst = "services"
)

// Constants associated with the Scope.LastScanType property.
// The last type of scan that was run on the scope.
const (
	ScopeLastScanTypeAbortTasksConst = "abort_tasks"
	ScopeLastScanTypeDiscoveryConst = "discovery"
	ScopeLastScanTypeEvidenceConst = "evidence"
	ScopeLastScanTypeFactCollectionConst = "fact_collection"
	ScopeLastScanTypeFactValidationConst = "fact_validation"
	ScopeLastScanTypeInventoryConst = "inventory"
	ScopeLastScanTypeRemediationConst = "remediation"
	ScopeLastScanTypeScriptConst = "script"
	ScopeLastScanTypeValidationConst = "validation"
)

// UnmarshalScope unmarshals an instance of Scope from the specified map of raw messages.
func UnmarshalScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Scope)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "environment_type", &obj.EnvironmentType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_time", &obj.ModifiedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_type", &obj.LastScanType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_type_description", &obj.LastScanTypeDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_scan_status_updated_time", &obj.LastScanStatusUpdatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "collectors_id", &obj.CollectorsID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "scans", &obj.Scans, UnmarshalScan)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ScopesList : Scopes list.
type ScopesList struct {
	// Scopes.
	Scopes []Scope `json:"scopes,omitempty"`
}

// UnmarshalScopesList unmarshals an instance of ScopesList from the specified map of raw messages.
func UnmarshalScopesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScopesList)
	err = core.UnmarshalModel(m, "scopes", &obj.Scopes, UnmarshalScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
