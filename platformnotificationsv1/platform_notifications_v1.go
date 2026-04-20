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
 * IBM OpenAPI SDK Code Generator Version: 3.108.0-56772134-20251111-102802
 */

// Package platformnotificationsv1 : Operations and models for the PlatformNotificationsV1 service
package platformnotificationsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// PlatformNotificationsV1 : **This API is currently in beta and subject to change.**
//
// API for managing notification distribution lists for IBM Cloud accounts.
//
// API Version: 1.0.0
type PlatformNotificationsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://notifications.cloud.ibm.com/api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "platform_notifications"

// PlatformNotificationsV1Options : Service options
type PlatformNotificationsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPlatformNotificationsV1UsingExternalConfig : constructs an instance of PlatformNotificationsV1 with passed in options and external configuration.
func NewPlatformNotificationsV1UsingExternalConfig(options *PlatformNotificationsV1Options) (platformNotifications *PlatformNotificationsV1, err error) {
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

	platformNotifications, err = NewPlatformNotificationsV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = platformNotifications.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = platformNotifications.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewPlatformNotificationsV1 : constructs an instance of PlatformNotificationsV1 with passed in options.
func NewPlatformNotificationsV1(options *PlatformNotificationsV1Options) (service *PlatformNotificationsV1, err error) {
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

	service = &PlatformNotificationsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "platformNotifications" suitable for processing requests.
func (platformNotifications *PlatformNotificationsV1) Clone() *PlatformNotificationsV1 {
	if core.IsNil(platformNotifications) {
		return nil
	}
	clone := *platformNotifications
	clone.Service = platformNotifications.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (platformNotifications *PlatformNotificationsV1) SetServiceURL(url string) error {
	err := platformNotifications.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (platformNotifications *PlatformNotificationsV1) GetServiceURL() string {
	return platformNotifications.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (platformNotifications *PlatformNotificationsV1) SetDefaultHeaders(headers http.Header) {
	platformNotifications.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (platformNotifications *PlatformNotificationsV1) SetEnableGzipCompression(enableGzip bool) {
	platformNotifications.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (platformNotifications *PlatformNotificationsV1) GetEnableGzipCompression() bool {
	return platformNotifications.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (platformNotifications *PlatformNotificationsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	platformNotifications.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (platformNotifications *PlatformNotificationsV1) DisableRetries() {
	platformNotifications.Service.DisableRetries()
}

// ListDistributionListDestinations : Get all destination entries
// Retrieve all destinations in the distribution list for the specified account.
func (platformNotifications *PlatformNotificationsV1) ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions) (result *AddDestinationCollection, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.ListDistributionListDestinationsWithContext(context.Background(), listDistributionListDestinationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDistributionListDestinationsWithContext is an alternate form of the ListDistributionListDestinations method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) ListDistributionListDestinationsWithContext(ctx context.Context, listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions) (result *AddDestinationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDistributionListDestinationsOptions, "listDistributionListDestinationsOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(listDistributionListDestinationsOptions, "listDistributionListDestinationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *listDistributionListDestinationsOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "ListDistributionListDestinations")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listDistributionListDestinationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_distribution_list_destinations", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddDestinationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreateDistributionListDestination : Add a destination entry
// Add a destination entry to the distribution list. A maximum of 10 destination entries per destination type. In terms
// of enterprise accounts, you can provide an Event Notifications destination that is from a different account than the
// distribution list account, provided these two accounts are from the same enterprise and the user has permission to
// manage the Event Notifications destinations on both accounts.
func (platformNotifications *PlatformNotificationsV1) CreateDistributionListDestination(createDistributionListDestinationOptions *CreateDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.CreateDistributionListDestinationWithContext(context.Background(), createDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDistributionListDestinationWithContext is an alternate form of the CreateDistributionListDestination method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) CreateDistributionListDestinationWithContext(ctx context.Context, createDistributionListDestinationOptions *CreateDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDistributionListDestinationOptions, "createDistributionListDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createDistributionListDestinationOptions, "createDistributionListDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *createDistributionListDestinationOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "CreateDistributionListDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createDistributionListDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(createDistributionListDestinationOptions.AddDestinationPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_distribution_list_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetDistributionListDestination : Get a destination entry
// Retrieve a specific destination from the distribution list of the account.
func (platformNotifications *PlatformNotificationsV1) GetDistributionListDestination(getDistributionListDestinationOptions *GetDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.GetDistributionListDestinationWithContext(context.Background(), getDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDistributionListDestinationWithContext is an alternate form of the GetDistributionListDestination method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) GetDistributionListDestinationWithContext(ctx context.Context, getDistributionListDestinationOptions *GetDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDistributionListDestinationOptions, "getDistributionListDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getDistributionListDestinationOptions, "getDistributionListDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getDistributionListDestinationOptions.AccountID,
		"destination_id": *getDistributionListDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "GetDistributionListDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getDistributionListDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_distribution_list_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAddDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteDistributionListDestination : Delete destination entry
// Remove a destination entry.
func (platformNotifications *PlatformNotificationsV1) DeleteDistributionListDestination(deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = platformNotifications.DeleteDistributionListDestinationWithContext(context.Background(), deleteDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDistributionListDestinationWithContext is an alternate form of the DeleteDistributionListDestination method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) DeleteDistributionListDestinationWithContext(ctx context.Context, deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDistributionListDestinationOptions, "deleteDistributionListDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteDistributionListDestinationOptions, "deleteDistributionListDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *deleteDistributionListDestinationOptions.AccountID,
		"destination_id": *deleteDistributionListDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{destination_id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "DeleteDistributionListDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteDistributionListDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = platformNotifications.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_distribution_list_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// TestDistributionListDestination : Test destination entry
// Send a test notification to a destination in the distribution list. This allows you to verify that the destination is
// properly configured and can receive notifications.
func (platformNotifications *PlatformNotificationsV1) TestDistributionListDestination(testDistributionListDestinationOptions *TestDistributionListDestinationOptions) (result *TestDestinationResponseBody, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.TestDistributionListDestinationWithContext(context.Background(), testDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// TestDistributionListDestinationWithContext is an alternate form of the TestDistributionListDestination method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) TestDistributionListDestinationWithContext(ctx context.Context, testDistributionListDestinationOptions *TestDistributionListDestinationOptions) (result *TestDestinationResponseBody, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(testDistributionListDestinationOptions, "testDistributionListDestinationOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(testDistributionListDestinationOptions, "testDistributionListDestinationOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *testDistributionListDestinationOptions.AccountID,
		"destination_id": *testDistributionListDestinationOptions.DestinationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{destination_id}/test`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "TestDistributionListDestination")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range testDistributionListDestinationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(testDistributionListDestinationOptions.TestDestinationRequestBodyPrototype)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "test_distribution_list_destination", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTestDestinationResponseBody)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// CreatePreferences : Create communication preferences
// Create communication preferences for the specified account.
func (platformNotifications *PlatformNotificationsV1) CreatePreferences(createPreferencesOptions *CreatePreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.CreatePreferencesWithContext(context.Background(), createPreferencesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreatePreferencesWithContext is an alternate form of the CreatePreferences method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) CreatePreferencesWithContext(ctx context.Context, createPreferencesOptions *CreatePreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPreferencesOptions, "createPreferencesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(createPreferencesOptions, "createPreferencesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *createPreferencesOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/{iam_id}/preferences`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "CreatePreferences")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range createPreferencesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createPreferencesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*createPreferencesOptions.AccountID))
	}

	body := make(map[string]interface{})
	if createPreferencesOptions.IncidentSeverity1 != nil {
		body["incident_severity1"] = createPreferencesOptions.IncidentSeverity1
	}
	if createPreferencesOptions.IncidentSeverity2 != nil {
		body["incident_severity2"] = createPreferencesOptions.IncidentSeverity2
	}
	if createPreferencesOptions.IncidentSeverity3 != nil {
		body["incident_severity3"] = createPreferencesOptions.IncidentSeverity3
	}
	if createPreferencesOptions.IncidentSeverity4 != nil {
		body["incident_severity4"] = createPreferencesOptions.IncidentSeverity4
	}
	if createPreferencesOptions.MaintenanceHigh != nil {
		body["maintenance_high"] = createPreferencesOptions.MaintenanceHigh
	}
	if createPreferencesOptions.MaintenanceMedium != nil {
		body["maintenance_medium"] = createPreferencesOptions.MaintenanceMedium
	}
	if createPreferencesOptions.MaintenanceLow != nil {
		body["maintenance_low"] = createPreferencesOptions.MaintenanceLow
	}
	if createPreferencesOptions.AnnouncementsMajor != nil {
		body["announcements_major"] = createPreferencesOptions.AnnouncementsMajor
	}
	if createPreferencesOptions.AnnouncementsMinor != nil {
		body["announcements_minor"] = createPreferencesOptions.AnnouncementsMinor
	}
	if createPreferencesOptions.SecurityNormal != nil {
		body["security_normal"] = createPreferencesOptions.SecurityNormal
	}
	if createPreferencesOptions.AccountNormal != nil {
		body["account_normal"] = createPreferencesOptions.AccountNormal
	}
	if createPreferencesOptions.BillingAndUsageOrder != nil {
		body["billing_and_usage_order"] = createPreferencesOptions.BillingAndUsageOrder
	}
	if createPreferencesOptions.BillingAndUsageInvoices != nil {
		body["billing_and_usage_invoices"] = createPreferencesOptions.BillingAndUsageInvoices
	}
	if createPreferencesOptions.BillingAndUsagePayments != nil {
		body["billing_and_usage_payments"] = createPreferencesOptions.BillingAndUsagePayments
	}
	if createPreferencesOptions.BillingAndUsageSubscriptionsAndPromoCodes != nil {
		body["billing_and_usage_subscriptions_and_promo_codes"] = createPreferencesOptions.BillingAndUsageSubscriptionsAndPromoCodes
	}
	if createPreferencesOptions.BillingAndUsageSpendingAlerts != nil {
		body["billing_and_usage_spending_alerts"] = createPreferencesOptions.BillingAndUsageSpendingAlerts
	}
	if createPreferencesOptions.ResourceactivityNormal != nil {
		body["resourceactivity_normal"] = createPreferencesOptions.ResourceactivityNormal
	}
	if createPreferencesOptions.OrderingReview != nil {
		body["ordering_review"] = createPreferencesOptions.OrderingReview
	}
	if createPreferencesOptions.OrderingApproved != nil {
		body["ordering_approved"] = createPreferencesOptions.OrderingApproved
	}
	if createPreferencesOptions.OrderingApprovedVsi != nil {
		body["ordering_approved_vsi"] = createPreferencesOptions.OrderingApprovedVsi
	}
	if createPreferencesOptions.OrderingApprovedServer != nil {
		body["ordering_approved_server"] = createPreferencesOptions.OrderingApprovedServer
	}
	if createPreferencesOptions.ProvisioningReloadComplete != nil {
		body["provisioning_reload_complete"] = createPreferencesOptions.ProvisioningReloadComplete
	}
	if createPreferencesOptions.ProvisioningCompleteVsi != nil {
		body["provisioning_complete_vsi"] = createPreferencesOptions.ProvisioningCompleteVsi
	}
	if createPreferencesOptions.ProvisioningCompleteServer != nil {
		body["provisioning_complete_server"] = createPreferencesOptions.ProvisioningCompleteServer
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "create_preferences", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPreferencesObject)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetPreferences : Get all communication preferences for a user in an account
// Retrieve all communication preferences of a user in an account.
func (platformNotifications *PlatformNotificationsV1) GetPreferences(getPreferencesOptions *GetPreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.GetPreferencesWithContext(context.Background(), getPreferencesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetPreferencesWithContext is an alternate form of the GetPreferences method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) GetPreferencesWithContext(ctx context.Context, getPreferencesOptions *GetPreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPreferencesOptions, "getPreferencesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(getPreferencesOptions, "getPreferencesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *getPreferencesOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/{iam_id}/preferences`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "GetPreferences")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getPreferencesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getPreferencesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getPreferencesOptions.AccountID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_preferences", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPreferencesObject)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceNotificationPreferences : Update communication preferences
// Update communication preferences for the specified account.
func (platformNotifications *PlatformNotificationsV1) ReplaceNotificationPreferences(replaceNotificationPreferencesOptions *ReplaceNotificationPreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.ReplaceNotificationPreferencesWithContext(context.Background(), replaceNotificationPreferencesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceNotificationPreferencesWithContext is an alternate form of the ReplaceNotificationPreferences method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) ReplaceNotificationPreferencesWithContext(ctx context.Context, replaceNotificationPreferencesOptions *ReplaceNotificationPreferencesOptions) (result *PreferencesObject, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceNotificationPreferencesOptions, "replaceNotificationPreferencesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceNotificationPreferencesOptions, "replaceNotificationPreferencesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *replaceNotificationPreferencesOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/{iam_id}/preferences`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "ReplaceNotificationPreferences")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceNotificationPreferencesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if replaceNotificationPreferencesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*replaceNotificationPreferencesOptions.AccountID))
	}

	body := make(map[string]interface{})
	if replaceNotificationPreferencesOptions.IncidentSeverity1 != nil {
		body["incident_severity1"] = replaceNotificationPreferencesOptions.IncidentSeverity1
	}
	if replaceNotificationPreferencesOptions.IncidentSeverity2 != nil {
		body["incident_severity2"] = replaceNotificationPreferencesOptions.IncidentSeverity2
	}
	if replaceNotificationPreferencesOptions.IncidentSeverity3 != nil {
		body["incident_severity3"] = replaceNotificationPreferencesOptions.IncidentSeverity3
	}
	if replaceNotificationPreferencesOptions.IncidentSeverity4 != nil {
		body["incident_severity4"] = replaceNotificationPreferencesOptions.IncidentSeverity4
	}
	if replaceNotificationPreferencesOptions.MaintenanceHigh != nil {
		body["maintenance_high"] = replaceNotificationPreferencesOptions.MaintenanceHigh
	}
	if replaceNotificationPreferencesOptions.MaintenanceMedium != nil {
		body["maintenance_medium"] = replaceNotificationPreferencesOptions.MaintenanceMedium
	}
	if replaceNotificationPreferencesOptions.MaintenanceLow != nil {
		body["maintenance_low"] = replaceNotificationPreferencesOptions.MaintenanceLow
	}
	if replaceNotificationPreferencesOptions.AnnouncementsMajor != nil {
		body["announcements_major"] = replaceNotificationPreferencesOptions.AnnouncementsMajor
	}
	if replaceNotificationPreferencesOptions.AnnouncementsMinor != nil {
		body["announcements_minor"] = replaceNotificationPreferencesOptions.AnnouncementsMinor
	}
	if replaceNotificationPreferencesOptions.SecurityNormal != nil {
		body["security_normal"] = replaceNotificationPreferencesOptions.SecurityNormal
	}
	if replaceNotificationPreferencesOptions.AccountNormal != nil {
		body["account_normal"] = replaceNotificationPreferencesOptions.AccountNormal
	}
	if replaceNotificationPreferencesOptions.BillingAndUsageOrder != nil {
		body["billing_and_usage_order"] = replaceNotificationPreferencesOptions.BillingAndUsageOrder
	}
	if replaceNotificationPreferencesOptions.BillingAndUsageInvoices != nil {
		body["billing_and_usage_invoices"] = replaceNotificationPreferencesOptions.BillingAndUsageInvoices
	}
	if replaceNotificationPreferencesOptions.BillingAndUsagePayments != nil {
		body["billing_and_usage_payments"] = replaceNotificationPreferencesOptions.BillingAndUsagePayments
	}
	if replaceNotificationPreferencesOptions.BillingAndUsageSubscriptionsAndPromoCodes != nil {
		body["billing_and_usage_subscriptions_and_promo_codes"] = replaceNotificationPreferencesOptions.BillingAndUsageSubscriptionsAndPromoCodes
	}
	if replaceNotificationPreferencesOptions.BillingAndUsageSpendingAlerts != nil {
		body["billing_and_usage_spending_alerts"] = replaceNotificationPreferencesOptions.BillingAndUsageSpendingAlerts
	}
	if replaceNotificationPreferencesOptions.ResourceactivityNormal != nil {
		body["resourceactivity_normal"] = replaceNotificationPreferencesOptions.ResourceactivityNormal
	}
	if replaceNotificationPreferencesOptions.OrderingReview != nil {
		body["ordering_review"] = replaceNotificationPreferencesOptions.OrderingReview
	}
	if replaceNotificationPreferencesOptions.OrderingApproved != nil {
		body["ordering_approved"] = replaceNotificationPreferencesOptions.OrderingApproved
	}
	if replaceNotificationPreferencesOptions.OrderingApprovedVsi != nil {
		body["ordering_approved_vsi"] = replaceNotificationPreferencesOptions.OrderingApprovedVsi
	}
	if replaceNotificationPreferencesOptions.OrderingApprovedServer != nil {
		body["ordering_approved_server"] = replaceNotificationPreferencesOptions.OrderingApprovedServer
	}
	if replaceNotificationPreferencesOptions.ProvisioningReloadComplete != nil {
		body["provisioning_reload_complete"] = replaceNotificationPreferencesOptions.ProvisioningReloadComplete
	}
	if replaceNotificationPreferencesOptions.ProvisioningCompleteVsi != nil {
		body["provisioning_complete_vsi"] = replaceNotificationPreferencesOptions.ProvisioningCompleteVsi
	}
	if replaceNotificationPreferencesOptions.ProvisioningCompleteServer != nil {
		body["provisioning_complete_server"] = replaceNotificationPreferencesOptions.ProvisioningCompleteServer
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "replace_notification_preferences", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPreferencesObject)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// DeleteNotificationPreferences : Resets all preferences to their default values
// Delete all communication preferences for the specified account, and resets all preferences to their default values.
func (platformNotifications *PlatformNotificationsV1) DeleteNotificationPreferences(deleteNotificationPreferencesOptions *DeleteNotificationPreferencesOptions) (response *core.DetailedResponse, err error) {
	response, err = platformNotifications.DeleteNotificationPreferencesWithContext(context.Background(), deleteNotificationPreferencesOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteNotificationPreferencesWithContext is an alternate form of the DeleteNotificationPreferences method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) DeleteNotificationPreferencesWithContext(ctx context.Context, deleteNotificationPreferencesOptions *DeleteNotificationPreferencesOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteNotificationPreferencesOptions, "deleteNotificationPreferencesOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(deleteNotificationPreferencesOptions, "deleteNotificationPreferencesOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	pathParamsMap := map[string]string{
		"iam_id": *deleteNotificationPreferencesOptions.IamID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/{iam_id}/preferences`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "DeleteNotificationPreferences")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range deleteNotificationPreferencesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteNotificationPreferencesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*deleteNotificationPreferencesOptions.AccountID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	response, err = platformNotifications.Service.Request(request, nil)
	if err != nil {
		core.EnrichHTTPProblem(err, "delete_notification_preferences", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}

	return
}

// ListNotifications : Get user notifications
// Retrieve all notifications for the requested user.
func (platformNotifications *PlatformNotificationsV1) ListNotifications(listNotificationsOptions *ListNotificationsOptions) (result *NotificationCollection, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.ListNotificationsWithContext(context.Background(), listNotificationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListNotificationsWithContext is an alternate form of the ListNotifications method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) ListNotificationsWithContext(ctx context.Context, listNotificationsOptions *ListNotificationsOptions) (result *NotificationCollection, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listNotificationsOptions, "listNotificationsOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "ListNotifications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range listNotificationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listNotificationsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listNotificationsOptions.AccountID))
	}
	if listNotificationsOptions.Start != nil {
		builder.AddQuery("start", fmt.Sprint(*listNotificationsOptions.Start))
	}
	if listNotificationsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listNotificationsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "list_notifications", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNotificationCollection)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// GetAcknowledgment : Get user's last acknowledged notification Id
// Retrieve the ID of the last notification acknowledged by the user for a specific account.
func (platformNotifications *PlatformNotificationsV1) GetAcknowledgment(getAcknowledgmentOptions *GetAcknowledgmentOptions) (result *Acknowledgment, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.GetAcknowledgmentWithContext(context.Background(), getAcknowledgmentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAcknowledgmentWithContext is an alternate form of the GetAcknowledgment method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) GetAcknowledgmentWithContext(ctx context.Context, getAcknowledgmentOptions *GetAcknowledgmentOptions) (result *Acknowledgment, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAcknowledgmentOptions, "getAcknowledgmentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/acknowledgment`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "GetAcknowledgment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range getAcknowledgmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getAcknowledgmentOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getAcknowledgmentOptions.AccountID))
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "get_acknowledgment", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcknowledgment)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}

// ReplaceNotificationAcknowledgment : Update user's last acknowledged notification
// Update the ID of the last notification acknowledged by the user for a specific account.
func (platformNotifications *PlatformNotificationsV1) ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions *ReplaceNotificationAcknowledgmentOptions) (result *Acknowledgment, response *core.DetailedResponse, err error) {
	result, response, err = platformNotifications.ReplaceNotificationAcknowledgmentWithContext(context.Background(), replaceNotificationAcknowledgmentOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ReplaceNotificationAcknowledgmentWithContext is an alternate form of the ReplaceNotificationAcknowledgment method which supports a Context parameter
func (platformNotifications *PlatformNotificationsV1) ReplaceNotificationAcknowledgmentWithContext(ctx context.Context, replaceNotificationAcknowledgmentOptions *ReplaceNotificationAcknowledgmentOptions) (result *Acknowledgment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceNotificationAcknowledgmentOptions, "replaceNotificationAcknowledgmentOptions cannot be nil")
	if err != nil {
		err = core.SDKErrorf(err, "", "unexpected-nil-param", common.GetComponentInfo())
		return
	}
	err = core.ValidateStruct(replaceNotificationAcknowledgmentOptions, "replaceNotificationAcknowledgmentOptions")
	if err != nil {
		err = core.SDKErrorf(err, "", "struct-validation-error", common.GetComponentInfo())
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = platformNotifications.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(platformNotifications.Service.Options.URL, `/v1/notifications/acknowledgment`, nil)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("platform_notifications", "V1", "ReplaceNotificationAcknowledgment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	for headerName, headerValue := range replaceNotificationAcknowledgmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if replaceNotificationAcknowledgmentOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*replaceNotificationAcknowledgmentOptions.AccountID))
	}

	body := make(map[string]interface{})
	if replaceNotificationAcknowledgmentOptions.LastAcknowledgedID != nil {
		body["last_acknowledged_id"] = replaceNotificationAcknowledgmentOptions.LastAcknowledgedID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		err = core.SDKErrorf(err, "", "set-json-body-error", common.GetComponentInfo())
		return
	}

	request, err := builder.Build()
	if err != nil {
		err = core.SDKErrorf(err, "", "build-error", common.GetComponentInfo())
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = platformNotifications.Service.Request(request, &rawResponse)
	if err != nil {
		core.EnrichHTTPProblem(err, "replace_notification_acknowledgment", getServiceComponentInfo())
		err = core.SDKErrorf(err, "", "http-request-err", common.GetComponentInfo())
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAcknowledgment)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-resp-error", common.GetComponentInfo())
			return
		}
		response.Result = result
	}

	return
}
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// Acknowledgment : Status indicating whether the user has unread notifications.
type Acknowledgment struct {
	// Indicates whether the user has unread notifications.
	HasUnread *bool `json:"has_unread" validate:"required"`

	// The ID of the most recent notification available to the user.
	LatestNotificationID *string `json:"latest_notification_id" validate:"required"`

	// The ID of the last notification acknowledged by the user.
	LastAcknowledgedID *string `json:"last_acknowledged_id" validate:"required"`
}

// UnmarshalAcknowledgment unmarshals an instance of Acknowledgment from the specified map of raw messages.
func UnmarshalAcknowledgment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Acknowledgment)
	err = core.UnmarshalPrimitive(m, "has_unread", &obj.HasUnread)
	if err != nil {
		err = core.SDKErrorf(err, "", "has_unread-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "latest_notification_id", &obj.LatestNotificationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "latest_notification_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "last_acknowledged_id", &obj.LastAcknowledgedID)
	if err != nil {
		err = core.SDKErrorf(err, "", "last_acknowledged_id-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddDestination : AddDestination struct
// Models which "extend" this model:
// - AddDestinationEventNotificationDestination
type AddDestination struct {
	// The GUID of the Event Notifications instance.
	DestinationID *strfmt.UUID `json:"destination_id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
}

// Constants associated with the AddDestination.DestinationType property.
// The type of the destination.
const (
	AddDestination_DestinationType_EventNotifications = "event_notifications"
)
func (*AddDestination) isaAddDestination() bool {
	return true
}

type AddDestinationIntf interface {
	isaAddDestination() bool
}

// UnmarshalAddDestination unmarshals an instance of AddDestination from the specified map of raw messages.
func UnmarshalAddDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", common.GetComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", common.GetComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddDestinationEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-AddDestinationEventNotificationDestination-error", common.GetComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// AddDestinationCollection : List of destinations in the distribution list.
type AddDestinationCollection struct {
	// Array of destination entries.
	Destinations []AddDestinationIntf `json:"destinations" validate:"required"`
}

// UnmarshalAddDestinationCollection unmarshals an instance of AddDestinationCollection from the specified map of raw messages.
func UnmarshalAddDestinationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationCollection)
	err = core.UnmarshalModel(m, "destinations", &obj.Destinations, UnmarshalAddDestination)
	if err != nil {
		err = core.SDKErrorf(err, "", "destinations-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddDestinationPrototype : AddDestinationPrototype struct
// Models which "extend" this model:
// - AddDestinationPrototypeEventNotificationDestinationPrototype
type AddDestinationPrototype struct {
	// The GUID of the Event Notifications instance.
	DestinationID *strfmt.UUID `json:"destination_id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationPrototype.DestinationType property.
// The type of the destination.
const (
	AddDestinationPrototype_DestinationType_EventNotifications = "event_notifications"
)
func (*AddDestinationPrototype) isaAddDestinationPrototype() bool {
	return true
}

type AddDestinationPrototypeIntf interface {
	isaAddDestinationPrototype() bool
}

// UnmarshalAddDestinationPrototype unmarshals an instance of AddDestinationPrototype from the specified map of raw messages.
func UnmarshalAddDestinationPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", common.GetComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", common.GetComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddDestinationPrototypeEventNotificationDestinationPrototype)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-AddDestinationPrototypeEventNotificationDestinationPrototype-error", common.GetComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// CreateDistributionListDestinationOptions : The CreateDistributionListDestination options.
type CreateDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	AddDestinationPrototype AddDestinationPrototypeIntf `json:"AddDestinationPrototype" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreateDistributionListDestinationOptions : Instantiate CreateDistributionListDestinationOptions
func (*PlatformNotificationsV1) NewCreateDistributionListDestinationOptions(accountID string, addDestinationPrototype AddDestinationPrototypeIntf) *CreateDistributionListDestinationOptions {
	return &CreateDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		AddDestinationPrototype: addDestinationPrototype,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *CreateDistributionListDestinationOptions) SetAccountID(accountID string) *CreateDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetAddDestinationPrototype : Allow user to set AddDestinationPrototype
func (_options *CreateDistributionListDestinationOptions) SetAddDestinationPrototype(addDestinationPrototype AddDestinationPrototypeIntf) *CreateDistributionListDestinationOptions {
	_options.AddDestinationPrototype = addDestinationPrototype
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDistributionListDestinationOptions) SetHeaders(param map[string]string) *CreateDistributionListDestinationOptions {
	options.Headers = param
	return options
}

// CreatePreferencesOptions : The CreatePreferences options.
type CreatePreferencesOptions struct {
	// The IAM ID of the user. Must match the IAM ID in the bearer token.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Preference settings for notification types that support updates.
	IncidentSeverity1 *PreferenceValueWithUpdates `json:"incident_severity1,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity2 *PreferenceValueWithUpdates `json:"incident_severity2,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity3 *PreferenceValueWithUpdates `json:"incident_severity3,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity4 *PreferenceValueWithUpdates `json:"incident_severity4,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceHigh *PreferenceValueWithUpdates `json:"maintenance_high,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceMedium *PreferenceValueWithUpdates `json:"maintenance_medium,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceLow *PreferenceValueWithUpdates `json:"maintenance_low,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMajor *PreferenceValueWithoutUpdates `json:"announcements_major,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMinor *PreferenceValueWithoutUpdates `json:"announcements_minor,omitempty"`

	// Preference settings for notification types that do not support updates.
	SecurityNormal *PreferenceValueWithoutUpdates `json:"security_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	AccountNormal *PreferenceValueWithoutUpdates `json:"account_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageOrder *PreferenceValueWithoutUpdates `json:"billing_and_usage_order,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageInvoices *PreferenceValueWithoutUpdates `json:"billing_and_usage_invoices,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsagePayments *PreferenceValueWithoutUpdates `json:"billing_and_usage_payments,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSubscriptionsAndPromoCodes *PreferenceValueWithoutUpdates `json:"billing_and_usage_subscriptions_and_promo_codes,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSpendingAlerts *PreferenceValueWithoutUpdates `json:"billing_and_usage_spending_alerts,omitempty"`

	// Preference settings for notification types that do not support updates.
	ResourceactivityNormal *PreferenceValueWithoutUpdates `json:"resourceactivity_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingReview *PreferenceValueWithoutUpdates `json:"ordering_review,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApproved *PreferenceValueWithoutUpdates `json:"ordering_approved,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedVsi *PreferenceValueWithoutUpdates `json:"ordering_approved_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedServer *PreferenceValueWithoutUpdates `json:"ordering_approved_server,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningReloadComplete *PreferenceValueWithoutUpdates `json:"provisioning_reload_complete,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteVsi *PreferenceValueWithoutUpdates `json:"provisioning_complete_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteServer *PreferenceValueWithoutUpdates `json:"provisioning_complete_server,omitempty"`

	// The IBM Cloud account ID. If not provided, the account ID from the bearer token will be used.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewCreatePreferencesOptions : Instantiate CreatePreferencesOptions
func (*PlatformNotificationsV1) NewCreatePreferencesOptions(iamID string) *CreatePreferencesOptions {
	return &CreatePreferencesOptions{
		IamID: core.StringPtr(iamID),
	}
}

// SetIamID : Allow user to set IamID
func (_options *CreatePreferencesOptions) SetIamID(iamID string) *CreatePreferencesOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetIncidentSeverity1 : Allow user to set IncidentSeverity1
func (_options *CreatePreferencesOptions) SetIncidentSeverity1(incidentSeverity1 *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.IncidentSeverity1 = incidentSeverity1
	return _options
}

// SetIncidentSeverity2 : Allow user to set IncidentSeverity2
func (_options *CreatePreferencesOptions) SetIncidentSeverity2(incidentSeverity2 *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.IncidentSeverity2 = incidentSeverity2
	return _options
}

// SetIncidentSeverity3 : Allow user to set IncidentSeverity3
func (_options *CreatePreferencesOptions) SetIncidentSeverity3(incidentSeverity3 *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.IncidentSeverity3 = incidentSeverity3
	return _options
}

// SetIncidentSeverity4 : Allow user to set IncidentSeverity4
func (_options *CreatePreferencesOptions) SetIncidentSeverity4(incidentSeverity4 *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.IncidentSeverity4 = incidentSeverity4
	return _options
}

// SetMaintenanceHigh : Allow user to set MaintenanceHigh
func (_options *CreatePreferencesOptions) SetMaintenanceHigh(maintenanceHigh *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.MaintenanceHigh = maintenanceHigh
	return _options
}

// SetMaintenanceMedium : Allow user to set MaintenanceMedium
func (_options *CreatePreferencesOptions) SetMaintenanceMedium(maintenanceMedium *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.MaintenanceMedium = maintenanceMedium
	return _options
}

// SetMaintenanceLow : Allow user to set MaintenanceLow
func (_options *CreatePreferencesOptions) SetMaintenanceLow(maintenanceLow *PreferenceValueWithUpdates) *CreatePreferencesOptions {
	_options.MaintenanceLow = maintenanceLow
	return _options
}

// SetAnnouncementsMajor : Allow user to set AnnouncementsMajor
func (_options *CreatePreferencesOptions) SetAnnouncementsMajor(announcementsMajor *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.AnnouncementsMajor = announcementsMajor
	return _options
}

// SetAnnouncementsMinor : Allow user to set AnnouncementsMinor
func (_options *CreatePreferencesOptions) SetAnnouncementsMinor(announcementsMinor *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.AnnouncementsMinor = announcementsMinor
	return _options
}

// SetSecurityNormal : Allow user to set SecurityNormal
func (_options *CreatePreferencesOptions) SetSecurityNormal(securityNormal *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.SecurityNormal = securityNormal
	return _options
}

// SetAccountNormal : Allow user to set AccountNormal
func (_options *CreatePreferencesOptions) SetAccountNormal(accountNormal *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.AccountNormal = accountNormal
	return _options
}

// SetBillingAndUsageOrder : Allow user to set BillingAndUsageOrder
func (_options *CreatePreferencesOptions) SetBillingAndUsageOrder(billingAndUsageOrder *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.BillingAndUsageOrder = billingAndUsageOrder
	return _options
}

// SetBillingAndUsageInvoices : Allow user to set BillingAndUsageInvoices
func (_options *CreatePreferencesOptions) SetBillingAndUsageInvoices(billingAndUsageInvoices *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.BillingAndUsageInvoices = billingAndUsageInvoices
	return _options
}

// SetBillingAndUsagePayments : Allow user to set BillingAndUsagePayments
func (_options *CreatePreferencesOptions) SetBillingAndUsagePayments(billingAndUsagePayments *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.BillingAndUsagePayments = billingAndUsagePayments
	return _options
}

// SetBillingAndUsageSubscriptionsAndPromoCodes : Allow user to set BillingAndUsageSubscriptionsAndPromoCodes
func (_options *CreatePreferencesOptions) SetBillingAndUsageSubscriptionsAndPromoCodes(billingAndUsageSubscriptionsAndPromoCodes *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.BillingAndUsageSubscriptionsAndPromoCodes = billingAndUsageSubscriptionsAndPromoCodes
	return _options
}

// SetBillingAndUsageSpendingAlerts : Allow user to set BillingAndUsageSpendingAlerts
func (_options *CreatePreferencesOptions) SetBillingAndUsageSpendingAlerts(billingAndUsageSpendingAlerts *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.BillingAndUsageSpendingAlerts = billingAndUsageSpendingAlerts
	return _options
}

// SetResourceactivityNormal : Allow user to set ResourceactivityNormal
func (_options *CreatePreferencesOptions) SetResourceactivityNormal(resourceactivityNormal *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.ResourceactivityNormal = resourceactivityNormal
	return _options
}

// SetOrderingReview : Allow user to set OrderingReview
func (_options *CreatePreferencesOptions) SetOrderingReview(orderingReview *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.OrderingReview = orderingReview
	return _options
}

// SetOrderingApproved : Allow user to set OrderingApproved
func (_options *CreatePreferencesOptions) SetOrderingApproved(orderingApproved *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.OrderingApproved = orderingApproved
	return _options
}

// SetOrderingApprovedVsi : Allow user to set OrderingApprovedVsi
func (_options *CreatePreferencesOptions) SetOrderingApprovedVsi(orderingApprovedVsi *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.OrderingApprovedVsi = orderingApprovedVsi
	return _options
}

// SetOrderingApprovedServer : Allow user to set OrderingApprovedServer
func (_options *CreatePreferencesOptions) SetOrderingApprovedServer(orderingApprovedServer *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.OrderingApprovedServer = orderingApprovedServer
	return _options
}

// SetProvisioningReloadComplete : Allow user to set ProvisioningReloadComplete
func (_options *CreatePreferencesOptions) SetProvisioningReloadComplete(provisioningReloadComplete *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.ProvisioningReloadComplete = provisioningReloadComplete
	return _options
}

// SetProvisioningCompleteVsi : Allow user to set ProvisioningCompleteVsi
func (_options *CreatePreferencesOptions) SetProvisioningCompleteVsi(provisioningCompleteVsi *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.ProvisioningCompleteVsi = provisioningCompleteVsi
	return _options
}

// SetProvisioningCompleteServer : Allow user to set ProvisioningCompleteServer
func (_options *CreatePreferencesOptions) SetProvisioningCompleteServer(provisioningCompleteServer *PreferenceValueWithoutUpdates) *CreatePreferencesOptions {
	_options.ProvisioningCompleteServer = provisioningCompleteServer
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *CreatePreferencesOptions) SetAccountID(accountID string) *CreatePreferencesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePreferencesOptions) SetHeaders(param map[string]string) *CreatePreferencesOptions {
	options.Headers = param
	return options
}

// DeleteDistributionListDestinationOptions : The DeleteDistributionListDestination options.
type DeleteDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The ID of the destination.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDistributionListDestinationOptions : Instantiate DeleteDistributionListDestinationOptions
func (*PlatformNotificationsV1) NewDeleteDistributionListDestinationOptions(accountID string, destinationID string) *DeleteDistributionListDestinationOptions {
	return &DeleteDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *DeleteDistributionListDestinationOptions) SetAccountID(accountID string) *DeleteDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *DeleteDistributionListDestinationOptions) SetDestinationID(destinationID string) *DeleteDistributionListDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDistributionListDestinationOptions) SetHeaders(param map[string]string) *DeleteDistributionListDestinationOptions {
	options.Headers = param
	return options
}

// DeleteNotificationPreferencesOptions : The DeleteNotificationPreferences options.
type DeleteNotificationPreferencesOptions struct {
	// The IAM ID of the user. Must match the IAM ID in the bearer token.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// The IBM Cloud account ID. If not provided, the account ID from the bearer token will be used.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteNotificationPreferencesOptions : Instantiate DeleteNotificationPreferencesOptions
func (*PlatformNotificationsV1) NewDeleteNotificationPreferencesOptions(iamID string) *DeleteNotificationPreferencesOptions {
	return &DeleteNotificationPreferencesOptions{
		IamID: core.StringPtr(iamID),
	}
}

// SetIamID : Allow user to set IamID
func (_options *DeleteNotificationPreferencesOptions) SetIamID(iamID string) *DeleteNotificationPreferencesOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *DeleteNotificationPreferencesOptions) SetAccountID(accountID string) *DeleteNotificationPreferencesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteNotificationPreferencesOptions) SetHeaders(param map[string]string) *DeleteNotificationPreferencesOptions {
	options.Headers = param
	return options
}

// GetAcknowledgmentOptions : The GetAcknowledgment options.
type GetAcknowledgmentOptions struct {
	// The account ID to retrieve acknowledgment for.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetAcknowledgmentOptions : Instantiate GetAcknowledgmentOptions
func (*PlatformNotificationsV1) NewGetAcknowledgmentOptions() *GetAcknowledgmentOptions {
	return &GetAcknowledgmentOptions{}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetAcknowledgmentOptions) SetAccountID(accountID string) *GetAcknowledgmentOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetAcknowledgmentOptions) SetHeaders(param map[string]string) *GetAcknowledgmentOptions {
	options.Headers = param
	return options
}

// GetDistributionListDestinationOptions : The GetDistributionListDestination options.
type GetDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The ID of the destination.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDistributionListDestinationOptions : Instantiate GetDistributionListDestinationOptions
func (*PlatformNotificationsV1) NewGetDistributionListDestinationOptions(accountID string, destinationID string) *GetDistributionListDestinationOptions {
	return &GetDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		DestinationID: core.StringPtr(destinationID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetDistributionListDestinationOptions) SetAccountID(accountID string) *GetDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *GetDistributionListDestinationOptions) SetDestinationID(destinationID string) *GetDistributionListDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDistributionListDestinationOptions) SetHeaders(param map[string]string) *GetDistributionListDestinationOptions {
	options.Headers = param
	return options
}

// GetPreferencesOptions : The GetPreferences options.
type GetPreferencesOptions struct {
	// The IAM ID of the user. Must match the IAM ID in the bearer token.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// The IBM Cloud account ID. If not provided, the account ID from the bearer token will be used.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetPreferencesOptions : Instantiate GetPreferencesOptions
func (*PlatformNotificationsV1) NewGetPreferencesOptions(iamID string) *GetPreferencesOptions {
	return &GetPreferencesOptions{
		IamID: core.StringPtr(iamID),
	}
}

// SetIamID : Allow user to set IamID
func (_options *GetPreferencesOptions) SetIamID(iamID string) *GetPreferencesOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *GetPreferencesOptions) SetAccountID(accountID string) *GetPreferencesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPreferencesOptions) SetHeaders(param map[string]string) *GetPreferencesOptions {
	options.Headers = param
	return options
}

// ListDistributionListDestinationsOptions : The ListDistributionListDestinations options.
type ListDistributionListDestinationsOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListDistributionListDestinationsOptions : Instantiate ListDistributionListDestinationsOptions
func (*PlatformNotificationsV1) NewListDistributionListDestinationsOptions(accountID string) *ListDistributionListDestinationsOptions {
	return &ListDistributionListDestinationsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListDistributionListDestinationsOptions) SetAccountID(accountID string) *ListDistributionListDestinationsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDistributionListDestinationsOptions) SetHeaders(param map[string]string) *ListDistributionListDestinationsOptions {
	options.Headers = param
	return options
}

// ListNotificationsOptions : The ListNotifications options.
type ListNotificationsOptions struct {
	// The IBM Cloud account ID. If not provided, the account ID from the bearer token will be used.
	AccountID *string `json:"account_id,omitempty"`

	// An opaque page token that specifies the resource to start the page on or after. If unspecified, the first page of
	// results is returned.
	Start *string `json:"start,omitempty"`

	// The maximum number of items to return per page. If unspecified, a default limit of 50 is used.
	Limit *int64 `json:"limit,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewListNotificationsOptions : Instantiate ListNotificationsOptions
func (*PlatformNotificationsV1) NewListNotificationsOptions() *ListNotificationsOptions {
	return &ListNotificationsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (_options *ListNotificationsOptions) SetAccountID(accountID string) *ListNotificationsOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetStart : Allow user to set Start
func (_options *ListNotificationsOptions) SetStart(start string) *ListNotificationsOptions {
	_options.Start = core.StringPtr(start)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListNotificationsOptions) SetLimit(limit int64) *ListNotificationsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListNotificationsOptions) SetHeaders(param map[string]string) *ListNotificationsOptions {
	options.Headers = param
	return options
}

// Notification : A notification entry.
type Notification struct {
	// The title of the notification.
	Title *string `json:"title" validate:"required"`

	// The body content of the notification.
	Body *string `json:"body" validate:"required"`

	// The unique identifier for the notification.
	ID *string `json:"id" validate:"required"`

	// The category of the notification.
	Category *string `json:"category" validate:"required"`

	// Array of component/service names affected by this notification.
	ComponentNames []string `json:"component_names" validate:"required"`

	// The start time of the notification in Unix timestamp (milliseconds).
	StartTime *int64 `json:"start_time" validate:"required"`

	// Indicates if the notification is global.
	IsGlobal *bool `json:"is_global" validate:"required"`

	// The current state of the notification.
	State *string `json:"state" validate:"required"`

	// Array of region identifiers affected by this notification.
	Regions []string `json:"regions" validate:"required"`

	// Array of CRN masks that define the scope of affected resources.
	CrnMasks []string `json:"crn_masks" validate:"required"`

	// The record identifier for tracking purposes.
	RecordID *string `json:"record_id,omitempty"`

	// The source identifier of the notification.
	SourceID *string `json:"source_id,omitempty"`

	// The completion code of the notification.
	CompletionCode *string `json:"completion_code" validate:"required"`

	// The end time of the notification in Unix timestamp (milliseconds).
	EndTime *int64 `json:"end_time,omitempty"`

	// The last update time of the notification in Unix timestamp (milliseconds).
	UpdateTime *int64 `json:"update_time" validate:"required"`

	// The severity level of the notification (0-3). The display value depends on the notification type:
	//
	// **Incidents:**
	// - 1 = Severity 1
	// - 2 = Severity 2
	// - 3 = Severity 3
	// - 0 = Severity 4
	//
	// **Maintenance:**
	// - 1 = High
	// - 2 = Medium
	// - 3 = Low
	//
	// **Announcements:**
	// - 1 = Major
	// - 0 = Minor.
	Severity *int64 `json:"severity" validate:"required"`

	// Lucene query string for filtering affected resources. Only present when instance targets are specified and
	// resource_link is not available. Mutually exclusive with resource_link.
	LuceneQuery *string `json:"lucene_query,omitempty"`

	// Link to additional resource information or documentation. Takes precedence over lucene_query when both are
	// available. Mutually exclusive with lucene_query.
	ResourceLink *string `json:"resource_link,omitempty"`
}

// Constants associated with the Notification.Category property.
// The category of the notification.
const (
	Notification_Category_Account = "account"
	Notification_Category_Announcements = "announcements"
	Notification_Category_BillingAndUsage = "billing_and_usage"
	Notification_Category_Incident = "incident"
	Notification_Category_Maintenance = "maintenance"
	Notification_Category_Ordering = "ordering"
	Notification_Category_Provisioning = "provisioning"
	Notification_Category_Resource = "resource"
	Notification_Category_Security = "security"
	Notification_Category_SecurityBulletins = "security_bulletins"
)

// Constants associated with the Notification.State property.
// The current state of the notification.
const (
	Notification_State_Complete = "complete"
	Notification_State_InProgress = "in-progress"
	Notification_State_New = "new"
	Notification_State_Resolved = "resolved"
)

// Constants associated with the Notification.CompletionCode property.
// The completion code of the notification.
const (
	Notification_CompletionCode_Cancelled = "cancelled"
	Notification_CompletionCode_Failed = "failed"
	Notification_CompletionCode_Successful = "successful"
)

// UnmarshalNotification unmarshals an instance of Notification from the specified map of raw messages.
func UnmarshalNotification(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Notification)
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		err = core.SDKErrorf(err, "", "title-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "body", &obj.Body)
	if err != nil {
		err = core.SDKErrorf(err, "", "body-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		err = core.SDKErrorf(err, "", "category-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "component_names", &obj.ComponentNames)
	if err != nil {
		err = core.SDKErrorf(err, "", "component_names-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start_time", &obj.StartTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "start_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "is_global", &obj.IsGlobal)
	if err != nil {
		err = core.SDKErrorf(err, "", "is_global-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		err = core.SDKErrorf(err, "", "state-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "regions", &obj.Regions)
	if err != nil {
		err = core.SDKErrorf(err, "", "regions-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "crn_masks", &obj.CrnMasks)
	if err != nil {
		err = core.SDKErrorf(err, "", "crn_masks-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "record_id", &obj.RecordID)
	if err != nil {
		err = core.SDKErrorf(err, "", "record_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "source_id", &obj.SourceID)
	if err != nil {
		err = core.SDKErrorf(err, "", "source_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "completion_code", &obj.CompletionCode)
	if err != nil {
		err = core.SDKErrorf(err, "", "completion_code-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "end_time", &obj.EndTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "end_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "update_time", &obj.UpdateTime)
	if err != nil {
		err = core.SDKErrorf(err, "", "update_time-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "severity", &obj.Severity)
	if err != nil {
		err = core.SDKErrorf(err, "", "severity-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "lucene_query", &obj.LuceneQuery)
	if err != nil {
		err = core.SDKErrorf(err, "", "lucene_query-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_link", &obj.ResourceLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "resource_link-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NotificationCollection : Collection of user notifications with token-based pagination metadata.
type NotificationCollection struct {
	// The maximum number of items returned in this response.
	Limit *int64 `json:"limit" validate:"required"`

	// The total number of notifications in the collection.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// A pagination link object containing the URL to a page.
	First *PaginationLink `json:"first" validate:"required"`

	// A pagination link object with a page token. Used for next, previous, and last page links.
	Previous *PaginationLinkWithToken `json:"previous,omitempty"`

	// A pagination link object with a page token. Used for next, previous, and last page links.
	Next *PaginationLinkWithToken `json:"next,omitempty"`

	// A pagination link object with a page token. Used for next, previous, and last page links.
	Last *PaginationLinkWithToken `json:"last,omitempty"`

	// Array of notification entries.
	Notifications []Notification `json:"notifications" validate:"required"`
}

// UnmarshalNotificationCollection unmarshals an instance of NotificationCollection from the specified map of raw messages.
func UnmarshalNotificationCollection(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NotificationCollection)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		err = core.SDKErrorf(err, "", "limit-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		err = core.SDKErrorf(err, "", "total_count-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPaginationLink)
	if err != nil {
		err = core.SDKErrorf(err, "", "first-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "previous", &obj.Previous, UnmarshalPaginationLinkWithToken)
	if err != nil {
		err = core.SDKErrorf(err, "", "previous-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPaginationLinkWithToken)
	if err != nil {
		err = core.SDKErrorf(err, "", "next-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalPaginationLinkWithToken)
	if err != nil {
		err = core.SDKErrorf(err, "", "last-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "notifications", &obj.Notifications, UnmarshalNotification)
	if err != nil {
		err = core.SDKErrorf(err, "", "notifications-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *NotificationCollection) GetNextStart() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Start, nil
}

// PaginationLink : A pagination link object containing the URL to a page.
type PaginationLink struct {
	// Complete URL to the page.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalPaginationLink unmarshals an instance of PaginationLink from the specified map of raw messages.
func UnmarshalPaginationLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLink)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PaginationLinkWithToken : A pagination link object with a page token. Used for next, previous, and last page links.
type PaginationLinkWithToken struct {
	// Complete URL to the page.
	Href *string `json:"href" validate:"required"`

	// Opaque page token that can be used to retrieve the page.
	Start *string `json:"start" validate:"required"`
}

// UnmarshalPaginationLinkWithToken unmarshals an instance of PaginationLinkWithToken from the specified map of raw messages.
func UnmarshalPaginationLinkWithToken(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PaginationLinkWithToken)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		err = core.SDKErrorf(err, "", "href-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		err = core.SDKErrorf(err, "", "start-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PreferenceValueWithUpdates : Preference settings for notification types that support updates.
type PreferenceValueWithUpdates struct {
	// Array of communication channels for this preference.
	Channels []string `json:"channels" validate:"required"`

	// Whether to receive updates for this preference. Optional, defaults to false if not provided.
	Updates *bool `json:"updates,omitempty"`
}

// Constants associated with the PreferenceValueWithUpdates.Channels property.
const (
	PreferenceValueWithUpdates_Channels_Email = "email"
)

// NewPreferenceValueWithUpdates : Instantiate PreferenceValueWithUpdates (Generic Model Constructor)
func (*PlatformNotificationsV1) NewPreferenceValueWithUpdates(channels []string) (_model *PreferenceValueWithUpdates, err error) {
	_model = &PreferenceValueWithUpdates{
		Channels: channels,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalPreferenceValueWithUpdates unmarshals an instance of PreferenceValueWithUpdates from the specified map of raw messages.
func UnmarshalPreferenceValueWithUpdates(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PreferenceValueWithUpdates)
	err = core.UnmarshalPrimitive(m, "channels", &obj.Channels)
	if err != nil {
		err = core.SDKErrorf(err, "", "channels-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "updates", &obj.Updates)
	if err != nil {
		err = core.SDKErrorf(err, "", "updates-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PreferenceValueWithoutUpdates : Preference settings for notification types that do not support updates.
type PreferenceValueWithoutUpdates struct {
	// Array of communication channels for this preference.
	Channels []string `json:"channels" validate:"required"`
}

// Constants associated with the PreferenceValueWithoutUpdates.Channels property.
const (
	PreferenceValueWithoutUpdates_Channels_Email = "email"
)

// NewPreferenceValueWithoutUpdates : Instantiate PreferenceValueWithoutUpdates (Generic Model Constructor)
func (*PlatformNotificationsV1) NewPreferenceValueWithoutUpdates(channels []string) (_model *PreferenceValueWithoutUpdates, err error) {
	_model = &PreferenceValueWithoutUpdates{
		Channels: channels,
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

// UnmarshalPreferenceValueWithoutUpdates unmarshals an instance of PreferenceValueWithoutUpdates from the specified map of raw messages.
func UnmarshalPreferenceValueWithoutUpdates(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PreferenceValueWithoutUpdates)
	err = core.UnmarshalPrimitive(m, "channels", &obj.Channels)
	if err != nil {
		err = core.SDKErrorf(err, "", "channels-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PreferencesObject : User communication preferences object. Only include preferences where communication is desired; absence of a key
// means no communication for that preference type.
type PreferencesObject struct {
	// Preference settings for notification types that support updates.
	IncidentSeverity1 *PreferenceValueWithUpdates `json:"incident_severity1,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity2 *PreferenceValueWithUpdates `json:"incident_severity2,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity3 *PreferenceValueWithUpdates `json:"incident_severity3,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity4 *PreferenceValueWithUpdates `json:"incident_severity4,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceHigh *PreferenceValueWithUpdates `json:"maintenance_high,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceMedium *PreferenceValueWithUpdates `json:"maintenance_medium,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceLow *PreferenceValueWithUpdates `json:"maintenance_low,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMajor *PreferenceValueWithoutUpdates `json:"announcements_major,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMinor *PreferenceValueWithoutUpdates `json:"announcements_minor,omitempty"`

	// Preference settings for notification types that do not support updates.
	SecurityNormal *PreferenceValueWithoutUpdates `json:"security_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	AccountNormal *PreferenceValueWithoutUpdates `json:"account_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageOrder *PreferenceValueWithoutUpdates `json:"billing_and_usage_order,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageInvoices *PreferenceValueWithoutUpdates `json:"billing_and_usage_invoices,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsagePayments *PreferenceValueWithoutUpdates `json:"billing_and_usage_payments,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSubscriptionsAndPromoCodes *PreferenceValueWithoutUpdates `json:"billing_and_usage_subscriptions_and_promo_codes,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSpendingAlerts *PreferenceValueWithoutUpdates `json:"billing_and_usage_spending_alerts,omitempty"`

	// Preference settings for notification types that do not support updates.
	ResourceactivityNormal *PreferenceValueWithoutUpdates `json:"resourceactivity_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingReview *PreferenceValueWithoutUpdates `json:"ordering_review,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApproved *PreferenceValueWithoutUpdates `json:"ordering_approved,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedVsi *PreferenceValueWithoutUpdates `json:"ordering_approved_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedServer *PreferenceValueWithoutUpdates `json:"ordering_approved_server,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningReloadComplete *PreferenceValueWithoutUpdates `json:"provisioning_reload_complete,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteVsi *PreferenceValueWithoutUpdates `json:"provisioning_complete_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteServer *PreferenceValueWithoutUpdates `json:"provisioning_complete_server,omitempty"`
}

// UnmarshalPreferencesObject unmarshals an instance of PreferencesObject from the specified map of raw messages.
func UnmarshalPreferencesObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PreferencesObject)
	err = core.UnmarshalModel(m, "incident_severity1", &obj.IncidentSeverity1, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "incident_severity1-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "incident_severity2", &obj.IncidentSeverity2, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "incident_severity2-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "incident_severity3", &obj.IncidentSeverity3, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "incident_severity3-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "incident_severity4", &obj.IncidentSeverity4, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "incident_severity4-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "maintenance_high", &obj.MaintenanceHigh, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "maintenance_high-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "maintenance_medium", &obj.MaintenanceMedium, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "maintenance_medium-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "maintenance_low", &obj.MaintenanceLow, UnmarshalPreferenceValueWithUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "maintenance_low-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "announcements_major", &obj.AnnouncementsMajor, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "announcements_major-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "announcements_minor", &obj.AnnouncementsMinor, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "announcements_minor-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "security_normal", &obj.SecurityNormal, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "security_normal-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "account_normal", &obj.AccountNormal, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "account_normal-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "billing_and_usage_order", &obj.BillingAndUsageOrder, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_and_usage_order-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "billing_and_usage_invoices", &obj.BillingAndUsageInvoices, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_and_usage_invoices-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "billing_and_usage_payments", &obj.BillingAndUsagePayments, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_and_usage_payments-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "billing_and_usage_subscriptions_and_promo_codes", &obj.BillingAndUsageSubscriptionsAndPromoCodes, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_and_usage_subscriptions_and_promo_codes-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "billing_and_usage_spending_alerts", &obj.BillingAndUsageSpendingAlerts, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "billing_and_usage_spending_alerts-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "resourceactivity_normal", &obj.ResourceactivityNormal, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "resourceactivity_normal-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ordering_review", &obj.OrderingReview, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "ordering_review-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ordering_approved", &obj.OrderingApproved, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "ordering_approved-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ordering_approved_vsi", &obj.OrderingApprovedVsi, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "ordering_approved_vsi-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "ordering_approved_server", &obj.OrderingApprovedServer, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "ordering_approved_server-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "provisioning_reload_complete", &obj.ProvisioningReloadComplete, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_reload_complete-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "provisioning_complete_vsi", &obj.ProvisioningCompleteVsi, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_complete_vsi-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalModel(m, "provisioning_complete_server", &obj.ProvisioningCompleteServer, UnmarshalPreferenceValueWithoutUpdates)
	if err != nil {
		err = core.SDKErrorf(err, "", "provisioning_complete_server-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceNotificationAcknowledgmentOptions : The ReplaceNotificationAcknowledgment options.
type ReplaceNotificationAcknowledgmentOptions struct {
	// The ID of a notification.
	LastAcknowledgedID *string `json:"last_acknowledged_id" validate:"required"`

	// The account ID to update acknowledgment for.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceNotificationAcknowledgmentOptions : Instantiate ReplaceNotificationAcknowledgmentOptions
func (*PlatformNotificationsV1) NewReplaceNotificationAcknowledgmentOptions(lastAcknowledgedID string) *ReplaceNotificationAcknowledgmentOptions {
	return &ReplaceNotificationAcknowledgmentOptions{
		LastAcknowledgedID: core.StringPtr(lastAcknowledgedID),
	}
}

// SetLastAcknowledgedID : Allow user to set LastAcknowledgedID
func (_options *ReplaceNotificationAcknowledgmentOptions) SetLastAcknowledgedID(lastAcknowledgedID string) *ReplaceNotificationAcknowledgmentOptions {
	_options.LastAcknowledgedID = core.StringPtr(lastAcknowledgedID)
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceNotificationAcknowledgmentOptions) SetAccountID(accountID string) *ReplaceNotificationAcknowledgmentOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceNotificationAcknowledgmentOptions) SetHeaders(param map[string]string) *ReplaceNotificationAcknowledgmentOptions {
	options.Headers = param
	return options
}

// ReplaceNotificationPreferencesOptions : The ReplaceNotificationPreferences options.
type ReplaceNotificationPreferencesOptions struct {
	// The IAM ID of the user. Must match the IAM ID in the bearer token.
	IamID *string `json:"iam_id" validate:"required,ne="`

	// Preference settings for notification types that support updates.
	IncidentSeverity1 *PreferenceValueWithUpdates `json:"incident_severity1,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity2 *PreferenceValueWithUpdates `json:"incident_severity2,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity3 *PreferenceValueWithUpdates `json:"incident_severity3,omitempty"`

	// Preference settings for notification types that support updates.
	IncidentSeverity4 *PreferenceValueWithUpdates `json:"incident_severity4,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceHigh *PreferenceValueWithUpdates `json:"maintenance_high,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceMedium *PreferenceValueWithUpdates `json:"maintenance_medium,omitempty"`

	// Preference settings for notification types that support updates.
	MaintenanceLow *PreferenceValueWithUpdates `json:"maintenance_low,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMajor *PreferenceValueWithoutUpdates `json:"announcements_major,omitempty"`

	// Preference settings for notification types that do not support updates.
	AnnouncementsMinor *PreferenceValueWithoutUpdates `json:"announcements_minor,omitempty"`

	// Preference settings for notification types that do not support updates.
	SecurityNormal *PreferenceValueWithoutUpdates `json:"security_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	AccountNormal *PreferenceValueWithoutUpdates `json:"account_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageOrder *PreferenceValueWithoutUpdates `json:"billing_and_usage_order,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageInvoices *PreferenceValueWithoutUpdates `json:"billing_and_usage_invoices,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsagePayments *PreferenceValueWithoutUpdates `json:"billing_and_usage_payments,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSubscriptionsAndPromoCodes *PreferenceValueWithoutUpdates `json:"billing_and_usage_subscriptions_and_promo_codes,omitempty"`

	// Preference settings for notification types that do not support updates.
	BillingAndUsageSpendingAlerts *PreferenceValueWithoutUpdates `json:"billing_and_usage_spending_alerts,omitempty"`

	// Preference settings for notification types that do not support updates.
	ResourceactivityNormal *PreferenceValueWithoutUpdates `json:"resourceactivity_normal,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingReview *PreferenceValueWithoutUpdates `json:"ordering_review,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApproved *PreferenceValueWithoutUpdates `json:"ordering_approved,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedVsi *PreferenceValueWithoutUpdates `json:"ordering_approved_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	OrderingApprovedServer *PreferenceValueWithoutUpdates `json:"ordering_approved_server,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningReloadComplete *PreferenceValueWithoutUpdates `json:"provisioning_reload_complete,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteVsi *PreferenceValueWithoutUpdates `json:"provisioning_complete_vsi,omitempty"`

	// Preference settings for notification types that do not support updates.
	ProvisioningCompleteServer *PreferenceValueWithoutUpdates `json:"provisioning_complete_server,omitempty"`

	// The IBM Cloud account ID. If not provided, the account ID from the bearer token will be used.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewReplaceNotificationPreferencesOptions : Instantiate ReplaceNotificationPreferencesOptions
func (*PlatformNotificationsV1) NewReplaceNotificationPreferencesOptions(iamID string) *ReplaceNotificationPreferencesOptions {
	return &ReplaceNotificationPreferencesOptions{
		IamID: core.StringPtr(iamID),
	}
}

// SetIamID : Allow user to set IamID
func (_options *ReplaceNotificationPreferencesOptions) SetIamID(iamID string) *ReplaceNotificationPreferencesOptions {
	_options.IamID = core.StringPtr(iamID)
	return _options
}

// SetIncidentSeverity1 : Allow user to set IncidentSeverity1
func (_options *ReplaceNotificationPreferencesOptions) SetIncidentSeverity1(incidentSeverity1 *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.IncidentSeverity1 = incidentSeverity1
	return _options
}

// SetIncidentSeverity2 : Allow user to set IncidentSeverity2
func (_options *ReplaceNotificationPreferencesOptions) SetIncidentSeverity2(incidentSeverity2 *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.IncidentSeverity2 = incidentSeverity2
	return _options
}

// SetIncidentSeverity3 : Allow user to set IncidentSeverity3
func (_options *ReplaceNotificationPreferencesOptions) SetIncidentSeverity3(incidentSeverity3 *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.IncidentSeverity3 = incidentSeverity3
	return _options
}

// SetIncidentSeverity4 : Allow user to set IncidentSeverity4
func (_options *ReplaceNotificationPreferencesOptions) SetIncidentSeverity4(incidentSeverity4 *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.IncidentSeverity4 = incidentSeverity4
	return _options
}

// SetMaintenanceHigh : Allow user to set MaintenanceHigh
func (_options *ReplaceNotificationPreferencesOptions) SetMaintenanceHigh(maintenanceHigh *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.MaintenanceHigh = maintenanceHigh
	return _options
}

// SetMaintenanceMedium : Allow user to set MaintenanceMedium
func (_options *ReplaceNotificationPreferencesOptions) SetMaintenanceMedium(maintenanceMedium *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.MaintenanceMedium = maintenanceMedium
	return _options
}

// SetMaintenanceLow : Allow user to set MaintenanceLow
func (_options *ReplaceNotificationPreferencesOptions) SetMaintenanceLow(maintenanceLow *PreferenceValueWithUpdates) *ReplaceNotificationPreferencesOptions {
	_options.MaintenanceLow = maintenanceLow
	return _options
}

// SetAnnouncementsMajor : Allow user to set AnnouncementsMajor
func (_options *ReplaceNotificationPreferencesOptions) SetAnnouncementsMajor(announcementsMajor *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.AnnouncementsMajor = announcementsMajor
	return _options
}

// SetAnnouncementsMinor : Allow user to set AnnouncementsMinor
func (_options *ReplaceNotificationPreferencesOptions) SetAnnouncementsMinor(announcementsMinor *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.AnnouncementsMinor = announcementsMinor
	return _options
}

// SetSecurityNormal : Allow user to set SecurityNormal
func (_options *ReplaceNotificationPreferencesOptions) SetSecurityNormal(securityNormal *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.SecurityNormal = securityNormal
	return _options
}

// SetAccountNormal : Allow user to set AccountNormal
func (_options *ReplaceNotificationPreferencesOptions) SetAccountNormal(accountNormal *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.AccountNormal = accountNormal
	return _options
}

// SetBillingAndUsageOrder : Allow user to set BillingAndUsageOrder
func (_options *ReplaceNotificationPreferencesOptions) SetBillingAndUsageOrder(billingAndUsageOrder *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.BillingAndUsageOrder = billingAndUsageOrder
	return _options
}

// SetBillingAndUsageInvoices : Allow user to set BillingAndUsageInvoices
func (_options *ReplaceNotificationPreferencesOptions) SetBillingAndUsageInvoices(billingAndUsageInvoices *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.BillingAndUsageInvoices = billingAndUsageInvoices
	return _options
}

// SetBillingAndUsagePayments : Allow user to set BillingAndUsagePayments
func (_options *ReplaceNotificationPreferencesOptions) SetBillingAndUsagePayments(billingAndUsagePayments *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.BillingAndUsagePayments = billingAndUsagePayments
	return _options
}

// SetBillingAndUsageSubscriptionsAndPromoCodes : Allow user to set BillingAndUsageSubscriptionsAndPromoCodes
func (_options *ReplaceNotificationPreferencesOptions) SetBillingAndUsageSubscriptionsAndPromoCodes(billingAndUsageSubscriptionsAndPromoCodes *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.BillingAndUsageSubscriptionsAndPromoCodes = billingAndUsageSubscriptionsAndPromoCodes
	return _options
}

// SetBillingAndUsageSpendingAlerts : Allow user to set BillingAndUsageSpendingAlerts
func (_options *ReplaceNotificationPreferencesOptions) SetBillingAndUsageSpendingAlerts(billingAndUsageSpendingAlerts *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.BillingAndUsageSpendingAlerts = billingAndUsageSpendingAlerts
	return _options
}

// SetResourceactivityNormal : Allow user to set ResourceactivityNormal
func (_options *ReplaceNotificationPreferencesOptions) SetResourceactivityNormal(resourceactivityNormal *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.ResourceactivityNormal = resourceactivityNormal
	return _options
}

// SetOrderingReview : Allow user to set OrderingReview
func (_options *ReplaceNotificationPreferencesOptions) SetOrderingReview(orderingReview *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.OrderingReview = orderingReview
	return _options
}

// SetOrderingApproved : Allow user to set OrderingApproved
func (_options *ReplaceNotificationPreferencesOptions) SetOrderingApproved(orderingApproved *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.OrderingApproved = orderingApproved
	return _options
}

// SetOrderingApprovedVsi : Allow user to set OrderingApprovedVsi
func (_options *ReplaceNotificationPreferencesOptions) SetOrderingApprovedVsi(orderingApprovedVsi *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.OrderingApprovedVsi = orderingApprovedVsi
	return _options
}

// SetOrderingApprovedServer : Allow user to set OrderingApprovedServer
func (_options *ReplaceNotificationPreferencesOptions) SetOrderingApprovedServer(orderingApprovedServer *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.OrderingApprovedServer = orderingApprovedServer
	return _options
}

// SetProvisioningReloadComplete : Allow user to set ProvisioningReloadComplete
func (_options *ReplaceNotificationPreferencesOptions) SetProvisioningReloadComplete(provisioningReloadComplete *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.ProvisioningReloadComplete = provisioningReloadComplete
	return _options
}

// SetProvisioningCompleteVsi : Allow user to set ProvisioningCompleteVsi
func (_options *ReplaceNotificationPreferencesOptions) SetProvisioningCompleteVsi(provisioningCompleteVsi *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.ProvisioningCompleteVsi = provisioningCompleteVsi
	return _options
}

// SetProvisioningCompleteServer : Allow user to set ProvisioningCompleteServer
func (_options *ReplaceNotificationPreferencesOptions) SetProvisioningCompleteServer(provisioningCompleteServer *PreferenceValueWithoutUpdates) *ReplaceNotificationPreferencesOptions {
	_options.ProvisioningCompleteServer = provisioningCompleteServer
	return _options
}

// SetAccountID : Allow user to set AccountID
func (_options *ReplaceNotificationPreferencesOptions) SetAccountID(accountID string) *ReplaceNotificationPreferencesOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceNotificationPreferencesOptions) SetHeaders(param map[string]string) *ReplaceNotificationPreferencesOptions {
	options.Headers = param
	return options
}

// TestDestinationRequestBodyPrototype : TestDestinationRequestBodyPrototype struct
// Models which "extend" this model:
// - TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype
type TestDestinationRequestBodyPrototype struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`

	// The type of the notification to test.
	NotificationType *string `json:"notification_type,omitempty"`
}

// Constants associated with the TestDestinationRequestBodyPrototype.DestinationType property.
// The type of the destination.
const (
	TestDestinationRequestBodyPrototype_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationRequestBodyPrototype.NotificationType property.
// The type of the notification to test.
const (
	TestDestinationRequestBodyPrototype_NotificationType_Announcements = "announcements"
	TestDestinationRequestBodyPrototype_NotificationType_BillingAndUsage = "billing_and_usage"
	TestDestinationRequestBodyPrototype_NotificationType_Incident = "incident"
	TestDestinationRequestBodyPrototype_NotificationType_Maintenance = "maintenance"
	TestDestinationRequestBodyPrototype_NotificationType_Resource = "resource"
	TestDestinationRequestBodyPrototype_NotificationType_SecurityBulletins = "security_bulletins"
)
func (*TestDestinationRequestBodyPrototype) isaTestDestinationRequestBodyPrototype() bool {
	return true
}

type TestDestinationRequestBodyPrototypeIntf interface {
	isaTestDestinationRequestBodyPrototype() bool
}

// UnmarshalTestDestinationRequestBodyPrototype unmarshals an instance of TestDestinationRequestBodyPrototype from the specified map of raw messages.
func UnmarshalTestDestinationRequestBodyPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "destination_type", &discValue)
	if err != nil {
		errMsg := fmt.Sprintf("error unmarshalling discriminator property 'destination_type': %s", err.Error())
		err = core.SDKErrorf(err, errMsg, "discriminator-unmarshal-error", common.GetComponentInfo())
		return
	}
	if discValue == "" {
		err = core.SDKErrorf(err, "required discriminator property 'destination_type' not found in JSON object", "missing-discriminator", common.GetComponentInfo())
		return
	}
	if discValue == "event_notifications" {
		err = core.UnmarshalModel(m, "", result, UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype-error", common.GetComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// TestDestinationResponseBody : Response from the test notification endpoint.
type TestDestinationResponseBody struct {
	// The status message that indicates the test result.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalTestDestinationResponseBody unmarshals an instance of TestDestinationResponseBody from the specified map of raw messages.
func UnmarshalTestDestinationResponseBody(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDestinationResponseBody)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		err = core.SDKErrorf(err, "", "message-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestDistributionListDestinationOptions : The TestDistributionListDestination options.
type TestDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The ID of the destination.
	DestinationID *string `json:"destination_id" validate:"required,ne="`

	TestDestinationRequestBodyPrototype TestDestinationRequestBodyPrototypeIntf `json:"TestDestinationRequestBodyPrototype" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewTestDistributionListDestinationOptions : Instantiate TestDistributionListDestinationOptions
func (*PlatformNotificationsV1) NewTestDistributionListDestinationOptions(accountID string, destinationID string, testDestinationRequestBodyPrototype TestDestinationRequestBodyPrototypeIntf) *TestDistributionListDestinationOptions {
	return &TestDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		DestinationID: core.StringPtr(destinationID),
		TestDestinationRequestBodyPrototype: testDestinationRequestBodyPrototype,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *TestDistributionListDestinationOptions) SetAccountID(accountID string) *TestDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetDestinationID : Allow user to set DestinationID
func (_options *TestDistributionListDestinationOptions) SetDestinationID(destinationID string) *TestDistributionListDestinationOptions {
	_options.DestinationID = core.StringPtr(destinationID)
	return _options
}

// SetTestDestinationRequestBodyPrototype : Allow user to set TestDestinationRequestBodyPrototype
func (_options *TestDistributionListDestinationOptions) SetTestDestinationRequestBodyPrototype(testDestinationRequestBodyPrototype TestDestinationRequestBodyPrototypeIntf) *TestDistributionListDestinationOptions {
	_options.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototype
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *TestDistributionListDestinationOptions) SetHeaders(param map[string]string) *TestDistributionListDestinationOptions {
	options.Headers = param
	return options
}

// AddDestinationPrototypeEventNotificationDestinationPrototype : Prototype for creating an Event Notifications destination entry.
// This model "extends" AddDestinationPrototype
type AddDestinationPrototypeEventNotificationDestinationPrototype struct {
	// The GUID of the Event Notifications instance.
	DestinationID *strfmt.UUID `json:"destination_id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationPrototypeEventNotificationDestinationPrototype.DestinationType property.
// The type of the destination.
const (
	AddDestinationPrototypeEventNotificationDestinationPrototype_DestinationType_EventNotifications = "event_notifications"
)

// NewAddDestinationPrototypeEventNotificationDestinationPrototype : Instantiate AddDestinationPrototypeEventNotificationDestinationPrototype (Generic Model Constructor)
func (*PlatformNotificationsV1) NewAddDestinationPrototypeEventNotificationDestinationPrototype(destinationID *strfmt.UUID, destinationType string) (_model *AddDestinationPrototypeEventNotificationDestinationPrototype, err error) {
	_model = &AddDestinationPrototypeEventNotificationDestinationPrototype{
		DestinationID: destinationID,
		DestinationType: core.StringPtr(destinationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*AddDestinationPrototypeEventNotificationDestinationPrototype) isaAddDestinationPrototype() bool {
	return true
}

// UnmarshalAddDestinationPrototypeEventNotificationDestinationPrototype unmarshals an instance of AddDestinationPrototypeEventNotificationDestinationPrototype from the specified map of raw messages.
func UnmarshalAddDestinationPrototypeEventNotificationDestinationPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationPrototypeEventNotificationDestinationPrototype)
	err = core.UnmarshalPrimitive(m, "destination_id", &obj.DestinationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddDestinationEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" AddDestination
type AddDestinationEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	DestinationID *strfmt.UUID `json:"destination_id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	AddDestinationEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

func (*AddDestinationEventNotificationDestination) isaAddDestination() bool {
	return true
}

// UnmarshalAddDestinationEventNotificationDestination unmarshals an instance of AddDestinationEventNotificationDestination from the specified map of raw messages.
func UnmarshalAddDestinationEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "destination_id", &obj.DestinationID)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_id-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype : Request body for testing an Event Notifications destination.
// This model "extends" TestDestinationRequestBodyPrototype
type TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`

	// The type of the notification to test.
	NotificationType *string `json:"notification_type" validate:"required"`
}

// Constants associated with the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype.DestinationType property.
// The type of the destination.
const (
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype.NotificationType property.
// The type of the notification to test.
const (
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_Announcements = "announcements"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_BillingAndUsage = "billing_and_usage"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_Incident = "incident"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_Maintenance = "maintenance"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_Resource = "resource"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype_NotificationType_SecurityBulletins = "security_bulletins"
)

// NewTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype : Instantiate TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype (Generic Model Constructor)
func (*PlatformNotificationsV1) NewTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype(destinationType string, notificationType string) (_model *TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype, err error) {
	_model = &TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype{
		DestinationType: core.StringPtr(destinationType),
		NotificationType: core.StringPtr(notificationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype) isaTestDestinationRequestBodyPrototype() bool {
	return true
}

// UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype unmarshals an instance of TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype from the specified map of raw messages.
func UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
	err = core.UnmarshalPrimitive(m, "destination_type", &obj.DestinationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "destination_type-error", common.GetComponentInfo())
		return
	}
	err = core.UnmarshalPrimitive(m, "notification_type", &obj.NotificationType)
	if err != nil {
		err = core.SDKErrorf(err, "", "notification_type-error", common.GetComponentInfo())
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

//
// NotificationsPager can be used to simplify the use of the "ListNotifications" method.
//
type NotificationsPager struct {
	hasNext bool
	options *ListNotificationsOptions
	client  *PlatformNotificationsV1
	pageContext struct {
		next *string
	}
}

// NewNotificationsPager returns a new NotificationsPager instance.
func (platformNotifications *PlatformNotificationsV1) NewNotificationsPager(options *ListNotificationsOptions) (pager *NotificationsPager, err error) {
	if options.Start != nil && *options.Start != "" {
		err = core.SDKErrorf(nil, "the 'options.Start' field should not be set", "no-query-setting", common.GetComponentInfo())
		return
	}

	var optionsCopy ListNotificationsOptions = *options
	pager = &NotificationsPager{
		hasNext: true,
		options: &optionsCopy,
		client:  platformNotifications,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *NotificationsPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *NotificationsPager) GetNextWithContext(ctx context.Context) (page []Notification, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Start = pager.pageContext.next

	result, _, err := pager.client.ListNotificationsWithContext(ctx, pager.options)
	if err != nil {
		err = core.RepurposeSDKProblem(err, "error-getting-next-page")
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Start
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Notifications

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *NotificationsPager) GetAllWithContext(ctx context.Context) (allItems []Notification, err error) {
	for pager.HasNext() {
		var nextPage []Notification
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			err = core.RepurposeSDKProblem(err, "error-getting-next-page")
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *NotificationsPager) GetNext() (page []Notification, err error) {
	page, err = pager.GetNextWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *NotificationsPager) GetAll() (allItems []Notification, err error) {
	allItems, err = pager.GetAllWithContext(context.Background())
	err = core.RepurposeSDKProblem(err, "")
	return
}
