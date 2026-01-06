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

// Package distributionlistapiv1 : Operations and models for the DistributionListApiV1 service
package distributionlistapiv1

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

// DistributionListApiV1 : API for managing distribution lists for IBM Cloud accounts.
//
// API Version: 1.0.0
type DistributionListApiV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://cloud.ibm.com/notification-api"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "distribution_list_api"

// DistributionListApiV1Options : Service options
type DistributionListApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDistributionListApiV1UsingExternalConfig : constructs an instance of DistributionListApiV1 with passed in options and external configuration.
func NewDistributionListApiV1UsingExternalConfig(options *DistributionListApiV1Options) (distributionListApi *DistributionListApiV1, err error) {
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

	distributionListApi, err = NewDistributionListApiV1(options)
	err = core.RepurposeSDKProblem(err, "new-client-error")
	if err != nil {
		return
	}

	err = distributionListApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		err = core.SDKErrorf(err, "", "client-config-error", common.GetComponentInfo())
		return
	}

	if options.URL != "" {
		err = distributionListApi.Service.SetServiceURL(options.URL)
		err = core.RepurposeSDKProblem(err, "url-set-error")
	}
	return
}

// NewDistributionListApiV1 : constructs an instance of DistributionListApiV1 with passed in options.
func NewDistributionListApiV1(options *DistributionListApiV1Options) (service *DistributionListApiV1, err error) {
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

	service = &DistributionListApiV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", core.SDKErrorf(nil, "service does not support regional URLs", "no-regional-support", common.GetComponentInfo())
}

// Clone makes a copy of "distributionListApi" suitable for processing requests.
func (distributionListApi *DistributionListApiV1) Clone() *DistributionListApiV1 {
	if core.IsNil(distributionListApi) {
		return nil
	}
	clone := *distributionListApi
	clone.Service = distributionListApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (distributionListApi *DistributionListApiV1) SetServiceURL(url string) error {
	err := distributionListApi.Service.SetServiceURL(url)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-set-error", common.GetComponentInfo())
	}
	return err
}

// GetServiceURL returns the service URL
func (distributionListApi *DistributionListApiV1) GetServiceURL() string {
	return distributionListApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (distributionListApi *DistributionListApiV1) SetDefaultHeaders(headers http.Header) {
	distributionListApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (distributionListApi *DistributionListApiV1) SetEnableGzipCompression(enableGzip bool) {
	distributionListApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (distributionListApi *DistributionListApiV1) GetEnableGzipCompression() bool {
	return distributionListApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (distributionListApi *DistributionListApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	distributionListApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (distributionListApi *DistributionListApiV1) DisableRetries() {
	distributionListApi.Service.DisableRetries()
}

// ListDistributionListDestinations : Get all destination entries
// Retrieve all destinations in the distribution list for the specified account.
func (distributionListApi *DistributionListApiV1) ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions) (result *AddDestinationCollection, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.ListDistributionListDestinationsWithContext(context.Background(), listDistributionListDestinationsOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// ListDistributionListDestinationsWithContext is an alternate form of the ListDistributionListDestinations method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) ListDistributionListDestinationsWithContext(ctx context.Context, listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions) (result *AddDestinationCollection, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("distribution_list_api", "V1", "ListDistributionListDestinations")
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
	response, err = distributionListApi.Service.Request(request, &rawResponse)
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
// Add a destination entry to the distribution list. Maximum of 10 destination entries per destination type. In terms of
// enterprise accounts, you can provide an Event Notifications destination that is from a different account than the
// distribution list account, provided these two accounts are from the same enterprise and the user has access rights to
// manage the Event Notification destinations on both accounts.
func (distributionListApi *DistributionListApiV1) CreateDistributionListDestination(createDistributionListDestinationOptions *CreateDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.CreateDistributionListDestinationWithContext(context.Background(), createDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// CreateDistributionListDestinationWithContext is an alternate form of the CreateDistributionListDestination method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) CreateDistributionListDestinationWithContext(ctx context.Context, createDistributionListDestinationOptions *CreateDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("distribution_list_api", "V1", "CreateDistributionListDestination")
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
	response, err = distributionListApi.Service.Request(request, &rawResponse)
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
// Retrieve a specific destination from the distribution list of the given account.
func (distributionListApi *DistributionListApiV1) GetDistributionListDestination(getDistributionListDestinationOptions *GetDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.GetDistributionListDestinationWithContext(context.Background(), getDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// GetDistributionListDestinationWithContext is an alternate form of the GetDistributionListDestination method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) GetDistributionListDestinationWithContext(ctx context.Context, getDistributionListDestinationOptions *GetDistributionListDestinationOptions) (result AddDestinationIntf, response *core.DetailedResponse, err error) {
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
		"id": *getDistributionListDestinationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("distribution_list_api", "V1", "GetDistributionListDestination")
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
	response, err = distributionListApi.Service.Request(request, &rawResponse)
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
func (distributionListApi *DistributionListApiV1) DeleteDistributionListDestination(deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions) (response *core.DetailedResponse, err error) {
	response, err = distributionListApi.DeleteDistributionListDestinationWithContext(context.Background(), deleteDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// DeleteDistributionListDestinationWithContext is an alternate form of the DeleteDistributionListDestination method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) DeleteDistributionListDestinationWithContext(ctx context.Context, deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions) (response *core.DetailedResponse, err error) {
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
		"id": *deleteDistributionListDestinationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{id}`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("distribution_list_api", "V1", "DeleteDistributionListDestination")
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

	response, err = distributionListApi.Service.Request(request, nil)
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
func (distributionListApi *DistributionListApiV1) TestDistributionListDestination(testDistributionListDestinationOptions *TestDistributionListDestinationOptions) (result *TestDestinationResponseBody, response *core.DetailedResponse, err error) {
	result, response, err = distributionListApi.TestDistributionListDestinationWithContext(context.Background(), testDistributionListDestinationOptions)
	err = core.RepurposeSDKProblem(err, "")
	return
}

// TestDistributionListDestinationWithContext is an alternate form of the TestDistributionListDestination method which supports a Context parameter
func (distributionListApi *DistributionListApiV1) TestDistributionListDestinationWithContext(ctx context.Context, testDistributionListDestinationOptions *TestDistributionListDestinationOptions) (result *TestDestinationResponseBody, response *core.DetailedResponse, err error) {
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
		"id": *testDistributionListDestinationOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = distributionListApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(distributionListApi.Service.Options.URL, `/v1/distribution_lists/{account_id}/destinations/{id}/test`, pathParamsMap)
	if err != nil {
		err = core.SDKErrorf(err, "", "url-resolve-error", common.GetComponentInfo())
		return
	}

	sdkHeaders := common.GetSdkHeaders("distribution_list_api", "V1", "TestDistributionListDestination")
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
	response, err = distributionListApi.Service.Request(request, &rawResponse)
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
func getServiceComponentInfo() *core.ProblemComponent {
	return core.NewProblemComponent(DefaultServiceName, "1.0.0")
}

// AddDestination : AddDestination struct
// Models which "extend" this model:
// - AddDestinationEventNotificationDestination
type AddDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

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
// - AddDestinationPrototypeEventNotificationDestination
type AddDestinationPrototype struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id,omitempty"`

	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`
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
		err = core.UnmarshalModel(m, "", result, UnmarshalAddDestinationPrototypeEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-AddDestinationPrototypeEventNotificationDestination-error", common.GetComponentInfo())
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
func (*DistributionListApiV1) NewCreateDistributionListDestinationOptions(accountID string, addDestinationPrototype AddDestinationPrototypeIntf) *CreateDistributionListDestinationOptions {
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

// DeleteDistributionListDestinationOptions : The DeleteDistributionListDestination options.
type DeleteDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The ID of the destination.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewDeleteDistributionListDestinationOptions : Instantiate DeleteDistributionListDestinationOptions
func (*DistributionListApiV1) NewDeleteDistributionListDestinationOptions(accountID string, id string) *DeleteDistributionListDestinationOptions {
	return &DeleteDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		ID: core.StringPtr(id),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *DeleteDistributionListDestinationOptions) SetAccountID(accountID string) *DeleteDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetID : Allow user to set ID
func (_options *DeleteDistributionListDestinationOptions) SetID(id string) *DeleteDistributionListDestinationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDistributionListDestinationOptions) SetHeaders(param map[string]string) *DeleteDistributionListDestinationOptions {
	options.Headers = param
	return options
}

// GetDistributionListDestinationOptions : The GetDistributionListDestination options.
type GetDistributionListDestinationOptions struct {
	// The IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The ID of the destination.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewGetDistributionListDestinationOptions : Instantiate GetDistributionListDestinationOptions
func (*DistributionListApiV1) NewGetDistributionListDestinationOptions(accountID string, id string) *GetDistributionListDestinationOptions {
	return &GetDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		ID: core.StringPtr(id),
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *GetDistributionListDestinationOptions) SetAccountID(accountID string) *GetDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetID : Allow user to set ID
func (_options *GetDistributionListDestinationOptions) SetID(id string) *GetDistributionListDestinationOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDistributionListDestinationOptions) SetHeaders(param map[string]string) *GetDistributionListDestinationOptions {
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
func (*DistributionListApiV1) NewListDistributionListDestinationsOptions(accountID string) *ListDistributionListDestinationsOptions {
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

// TestDestinationRequestBodyPrototype : TestDestinationRequestBodyPrototype struct
// Models which "extend" this model:
// - TestDestinationRequestBodyPrototypeTestEventNotificationDestination
type TestDestinationRequestBodyPrototype struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type,omitempty"`

	// Type of notification to test.
	NotificationType *string `json:"notification_type,omitempty"`
}

// Constants associated with the TestDestinationRequestBodyPrototype.DestinationType property.
// The type of the destination.
const (
	TestDestinationRequestBodyPrototype_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationRequestBodyPrototype.NotificationType property.
// Type of notification to test.
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
		err = core.UnmarshalModel(m, "", result, UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestination)
		if err != nil {
			err = core.SDKErrorf(err, "", "unmarshal-TestDestinationRequestBodyPrototypeTestEventNotificationDestination-error", common.GetComponentInfo())
		}
	} else {
		errMsg := fmt.Sprintf("unrecognized value for discriminator property 'destination_type': %s", discValue)
		err = core.SDKErrorf(err, errMsg, "invalid-discriminator", common.GetComponentInfo())
	}
	return
}

// TestDestinationResponseBody : Response from test notification endpoint.
type TestDestinationResponseBody struct {
	// Status message indicating test result.
	Message *string `json:"message,omitempty"`
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
	ID *string `json:"id" validate:"required,ne="`

	TestDestinationRequestBodyPrototype TestDestinationRequestBodyPrototypeIntf `json:"TestDestinationRequestBodyPrototype" validate:"required"`

	// Allows users to set headers on API requests.
	Headers map[string]string
}

// NewTestDistributionListDestinationOptions : Instantiate TestDistributionListDestinationOptions
func (*DistributionListApiV1) NewTestDistributionListDestinationOptions(accountID string, id string, testDestinationRequestBodyPrototype TestDestinationRequestBodyPrototypeIntf) *TestDistributionListDestinationOptions {
	return &TestDistributionListDestinationOptions{
		AccountID: core.StringPtr(accountID),
		ID: core.StringPtr(id),
		TestDestinationRequestBodyPrototype: testDestinationRequestBodyPrototype,
	}
}

// SetAccountID : Allow user to set AccountID
func (_options *TestDistributionListDestinationOptions) SetAccountID(accountID string) *TestDistributionListDestinationOptions {
	_options.AccountID = core.StringPtr(accountID)
	return _options
}

// SetID : Allow user to set ID
func (_options *TestDistributionListDestinationOptions) SetID(id string) *TestDistributionListDestinationOptions {
	_options.ID = core.StringPtr(id)
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

// AddDestinationPrototypeEventNotificationDestination : An Event Notifications destination entry in the distribution list.
// This model "extends" AddDestinationPrototype
type AddDestinationPrototypeEventNotificationDestination struct {
	// The GUID of the Event Notifications instance.
	ID *strfmt.UUID `json:"id" validate:"required"`

	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`
}

// Constants associated with the AddDestinationPrototypeEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	AddDestinationPrototypeEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

// NewAddDestinationPrototypeEventNotificationDestination : Instantiate AddDestinationPrototypeEventNotificationDestination (Generic Model Constructor)
func (*DistributionListApiV1) NewAddDestinationPrototypeEventNotificationDestination(id *strfmt.UUID, destinationType string) (_model *AddDestinationPrototypeEventNotificationDestination, err error) {
	_model = &AddDestinationPrototypeEventNotificationDestination{
		ID: id,
		DestinationType: core.StringPtr(destinationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*AddDestinationPrototypeEventNotificationDestination) isaAddDestinationPrototype() bool {
	return true
}

// UnmarshalAddDestinationPrototypeEventNotificationDestination unmarshals an instance of AddDestinationPrototypeEventNotificationDestination from the specified map of raw messages.
func UnmarshalAddDestinationPrototypeEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddDestinationPrototypeEventNotificationDestination)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
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
	ID *strfmt.UUID `json:"id" validate:"required"`

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
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		err = core.SDKErrorf(err, "", "id-error", common.GetComponentInfo())
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

// TestDestinationRequestBodyPrototypeTestEventNotificationDestination : Request body for testing an Event Notifications destination.
// This model "extends" TestDestinationRequestBodyPrototype
type TestDestinationRequestBodyPrototypeTestEventNotificationDestination struct {
	// The type of the destination.
	DestinationType *string `json:"destination_type" validate:"required"`

	// Type of notification to test.
	NotificationType *string `json:"notification_type" validate:"required"`
}

// Constants associated with the TestDestinationRequestBodyPrototypeTestEventNotificationDestination.DestinationType property.
// The type of the destination.
const (
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_DestinationType_EventNotifications = "event_notifications"
)

// Constants associated with the TestDestinationRequestBodyPrototypeTestEventNotificationDestination.NotificationType property.
// Type of notification to test.
const (
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_Announcements = "announcements"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_BillingAndUsage = "billing_and_usage"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_Incident = "incident"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_Maintenance = "maintenance"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_Resource = "resource"
	TestDestinationRequestBodyPrototypeTestEventNotificationDestination_NotificationType_SecurityBulletins = "security_bulletins"
)

// NewTestDestinationRequestBodyPrototypeTestEventNotificationDestination : Instantiate TestDestinationRequestBodyPrototypeTestEventNotificationDestination (Generic Model Constructor)
func (*DistributionListApiV1) NewTestDestinationRequestBodyPrototypeTestEventNotificationDestination(destinationType string, notificationType string) (_model *TestDestinationRequestBodyPrototypeTestEventNotificationDestination, err error) {
	_model = &TestDestinationRequestBodyPrototypeTestEventNotificationDestination{
		DestinationType: core.StringPtr(destinationType),
		NotificationType: core.StringPtr(notificationType),
	}
	err = core.ValidateStruct(_model, "required parameters")
	if err != nil {
		err = core.SDKErrorf(err, "", "model-missing-required", common.GetComponentInfo())
	}
	return
}

func (*TestDestinationRequestBodyPrototypeTestEventNotificationDestination) isaTestDestinationRequestBodyPrototype() bool {
	return true
}

// UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestination unmarshals an instance of TestDestinationRequestBodyPrototypeTestEventNotificationDestination from the specified map of raw messages.
func UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestination(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TestDestinationRequestBodyPrototypeTestEventNotificationDestination)
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
