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

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-818b4742-20201113-111146
 */
 

// Package atrackerv1 : Operations and models for the AtrackerV1 service
package atrackerv1

import (
	"context"
	"encoding/json"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"net/http"
	"reflect"
	"time"
)

// AtrackerV1 : IBM Cloud Activity Tracker Service (ATracker Service for short) is an activity tracker service for your
// application events as well as events from IBM services under your account. It is designed to enable you to route
// activity tracker events to your designated Cloud Object Storage location in different regions.
//
// Version: 1.0.0
type AtrackerV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://private.us-south.atracker.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "atracker"

// AtrackerV1Options : Service options
type AtrackerV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewAtrackerV1UsingExternalConfig : constructs an instance of AtrackerV1 with passed in options and external configuration.
func NewAtrackerV1UsingExternalConfig(options *AtrackerV1Options) (atracker *AtrackerV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	atracker, err = NewAtrackerV1(options)
	if err != nil {
		return
	}

	err = atracker.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = atracker.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAtrackerV1 : constructs an instance of AtrackerV1 with passed in options.
func NewAtrackerV1(options *AtrackerV1Options) (service *AtrackerV1, err error) {
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

	service = &AtrackerV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (atracker *AtrackerV1) SetServiceURL(url string) error {
	return atracker.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (atracker *AtrackerV1) GetServiceURL() string {
	return atracker.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (atracker *AtrackerV1) SetDefaultHeaders(headers http.Header) {
	atracker.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (atracker *AtrackerV1) SetEnableGzipCompression(enableGzip bool) {
	atracker.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (atracker *AtrackerV1) GetEnableGzipCompression() bool {
	return atracker.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (atracker *AtrackerV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	atracker.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (atracker *AtrackerV1) DisableRetries() {
	atracker.Service.DisableRetries()
}

// CreateTarget : Create a Cloud Object Storage target for a region
// Creates a new Cloud Object Storage (COS) target with specified COS endpoint information and credentials.  Commonly
// the COS endpoint should be on the same region as ATracker Services where this API is invoked. The  Target definition
// could only be referenced by the routing rules defined in the same region through the same  API endpoint. If a COS
// endpoint to be used across multiple regions, you must define a target for each region's API endpoint.
func (atracker *AtrackerV1) CreateTarget(createTargetOptions *CreateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return atracker.CreateTargetWithContext(context.Background(), createTargetOptions)
}

// CreateTargetWithContext is an alternate form of the CreateTarget method which supports a Context parameter
func (atracker *AtrackerV1) CreateTargetWithContext(ctx context.Context, createTargetOptions *CreateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createTargetOptions, "createTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createTargetOptions, "createTargetOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/targets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "CreateTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createTargetOptions.Name != nil {
		body["name"] = createTargetOptions.Name
	}
	if createTargetOptions.TargetType != nil {
		body["target_type"] = createTargetOptions.TargetType
	}
	if createTargetOptions.CosEndpoint != nil {
		body["cos_endpoint"] = createTargetOptions.CosEndpoint
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
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListTargets : List Cloud Object Storage targets for the region
// List all Cloud Object Storage (COS) targets defined under this region.
func (atracker *AtrackerV1) ListTargets(listTargetsOptions *ListTargetsOptions) (result *TargetList, response *core.DetailedResponse, err error) {
	return atracker.ListTargetsWithContext(context.Background(), listTargetsOptions)
}

// ListTargetsWithContext is an alternate form of the ListTargets method which supports a Context parameter
func (atracker *AtrackerV1) ListTargetsWithContext(ctx context.Context, listTargetsOptions *ListTargetsOptions) (result *TargetList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTargetsOptions, "listTargetsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/targets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTargetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "ListTargets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTargetList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetTarget : Retrieve a target
// Retrieves a target and its details by specifying the ID of the target.
func (atracker *AtrackerV1) GetTarget(getTargetOptions *GetTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return atracker.GetTargetWithContext(context.Background(), getTargetOptions)
}

// GetTargetWithContext is an alternate form of the GetTarget method which supports a Context parameter
func (atracker *AtrackerV1) GetTargetWithContext(ctx context.Context, getTargetOptions *GetTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTargetOptions, "getTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTargetOptions, "getTargetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getTargetOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "GetTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceTarget : Update a target
// Update a target details by specifying the ID of the target.
func (atracker *AtrackerV1) ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return atracker.ReplaceTargetWithContext(context.Background(), replaceTargetOptions)
}

// ReplaceTargetWithContext is an alternate form of the ReplaceTarget method which supports a Context parameter
func (atracker *AtrackerV1) ReplaceTargetWithContext(ctx context.Context, replaceTargetOptions *ReplaceTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceTargetOptions, "replaceTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceTargetOptions, "replaceTargetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *replaceTargetOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "ReplaceTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceTargetOptions.Name != nil {
		body["name"] = replaceTargetOptions.Name
	}
	if replaceTargetOptions.TargetType != nil {
		body["target_type"] = replaceTargetOptions.TargetType
	}
	if replaceTargetOptions.CosEndpoint != nil {
		body["cos_endpoint"] = replaceTargetOptions.CosEndpoint
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
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteTarget : Delete a target
// Deletes a target by specifying the ID of the target.
func (atracker *AtrackerV1) DeleteTarget(deleteTargetOptions *DeleteTargetOptions) (response *core.DetailedResponse, err error) {
	return atracker.DeleteTargetWithContext(context.Background(), deleteTargetOptions)
}

// DeleteTargetWithContext is an alternate form of the DeleteTarget method which supports a Context parameter
func (atracker *AtrackerV1) DeleteTargetWithContext(ctx context.Context, deleteTargetOptions *DeleteTargetOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTargetOptions, "deleteTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTargetOptions, "deleteTargetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteTargetOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "DeleteTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = atracker.Service.Request(request, nil)

	return
}

// CreateRoute : Create a Route for the region
// Creates a route with rules defined how to route AT events to targets for a region.  For each account and region, only
// one route could be defined. A route could contain multiple rules which enable atracker service to match incoming AT
// events based on the source crn and forward the events to customer configured targets.
func (atracker *AtrackerV1) CreateRoute(createRouteOptions *CreateRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return atracker.CreateRouteWithContext(context.Background(), createRouteOptions)
}

// CreateRouteWithContext is an alternate form of the CreateRoute method which supports a Context parameter
func (atracker *AtrackerV1) CreateRouteWithContext(ctx context.Context, createRouteOptions *CreateRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRouteOptions, "createRouteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRouteOptions, "createRouteOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/routes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "CreateRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createRouteOptions.Name != nil {
		body["name"] = createRouteOptions.Name
	}
	if createRouteOptions.ReceiveGlobalEvents != nil {
		body["receive_global_events"] = createRouteOptions.ReceiveGlobalEvents
	}
	if createRouteOptions.Rules != nil {
		body["rules"] = createRouteOptions.Rules
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
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListRoutes : List routes for the region
// List routes defined under this region.
func (atracker *AtrackerV1) ListRoutes(listRoutesOptions *ListRoutesOptions) (result *RouteList, response *core.DetailedResponse, err error) {
	return atracker.ListRoutesWithContext(context.Background(), listRoutesOptions)
}

// ListRoutesWithContext is an alternate form of the ListRoutes method which supports a Context parameter
func (atracker *AtrackerV1) ListRoutesWithContext(ctx context.Context, listRoutesOptions *ListRoutesOptions) (result *RouteList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRoutesOptions, "listRoutesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/routes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRoutesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "ListRoutes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRouteList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRoute : Retrieve a route
// Retrieves a route and its details by specifying the ID of the route.
func (atracker *AtrackerV1) GetRoute(getRouteOptions *GetRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return atracker.GetRouteWithContext(context.Background(), getRouteOptions)
}

// GetRouteWithContext is an alternate form of the GetRoute method which supports a Context parameter
func (atracker *AtrackerV1) GetRouteWithContext(ctx context.Context, getRouteOptions *GetRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRouteOptions, "getRouteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRouteOptions, "getRouteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *getRouteOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "GetRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceRoute : Replace a route
// Replace a route details by specifying the ID of the route.
func (atracker *AtrackerV1) ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return atracker.ReplaceRouteWithContext(context.Background(), replaceRouteOptions)
}

// ReplaceRouteWithContext is an alternate form of the ReplaceRoute method which supports a Context parameter
func (atracker *AtrackerV1) ReplaceRouteWithContext(ctx context.Context, replaceRouteOptions *ReplaceRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceRouteOptions, "replaceRouteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceRouteOptions, "replaceRouteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *replaceRouteOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "ReplaceRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceRouteOptions.Name != nil {
		body["name"] = replaceRouteOptions.Name
	}
	if replaceRouteOptions.ReceiveGlobalEvents != nil {
		body["receive_global_events"] = replaceRouteOptions.ReceiveGlobalEvents
	}
	if replaceRouteOptions.Rules != nil {
		body["rules"] = replaceRouteOptions.Rules
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
	response, err = atracker.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteRoute : Delete a route
// Deletes a route by specifying the ID of the route.
func (atracker *AtrackerV1) DeleteRoute(deleteRouteOptions *DeleteRouteOptions) (response *core.DetailedResponse, err error) {
	return atracker.DeleteRouteWithContext(context.Background(), deleteRouteOptions)
}

// DeleteRouteWithContext is an alternate form of the DeleteRoute method which supports a Context parameter
func (atracker *AtrackerV1) DeleteRouteWithContext(ctx context.Context, deleteRouteOptions *DeleteRouteOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRouteOptions, "deleteRouteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRouteOptions, "deleteRouteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *deleteRouteOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = atracker.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(atracker.Service.Options.URL, `/api/v1/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("atracker", "V1", "DeleteRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = atracker.Service.Request(request, nil)

	return
}

// CreateRouteOptions : The CreateRoute options.
type CreateRouteOptions struct {
	// The name of the route. Must be 180 characters or less and cannot include any special characters other than `(space)
	// - . _ :`.
	Name *string `json:"name" validate:"required"`

	// Whether or not all global events should be forwarded to this region.
	ReceiveGlobalEvents *bool `json:"receive_global_events" validate:"required"`

	// Routing rules that will be evaluated in their order of the array.
	Rules []Rule `json:"rules" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRouteOptions : Instantiate CreateRouteOptions
func (*AtrackerV1) NewCreateRouteOptions(name string, receiveGlobalEvents bool, rules []Rule) *CreateRouteOptions {
	return &CreateRouteOptions{
		Name: core.StringPtr(name),
		ReceiveGlobalEvents: core.BoolPtr(receiveGlobalEvents),
		Rules: rules,
	}
}

// SetName : Allow user to set Name
func (options *CreateRouteOptions) SetName(name string) *CreateRouteOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetReceiveGlobalEvents : Allow user to set ReceiveGlobalEvents
func (options *CreateRouteOptions) SetReceiveGlobalEvents(receiveGlobalEvents bool) *CreateRouteOptions {
	options.ReceiveGlobalEvents = core.BoolPtr(receiveGlobalEvents)
	return options
}

// SetRules : Allow user to set Rules
func (options *CreateRouteOptions) SetRules(rules []Rule) *CreateRouteOptions {
	options.Rules = rules
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRouteOptions) SetHeaders(param map[string]string) *CreateRouteOptions {
	options.Headers = param
	return options
}

// CreateTargetOptions : The CreateTarget options.
type CreateTargetOptions struct {
	// The name of the target. Must be 256 characters or less.
	Name *string `json:"name" validate:"required"`

	// The type of the target.
	TargetType *string `json:"target_type" validate:"required"`

	// Property values for a Cloud Object Storage Endpoint.
	CosEndpoint *CosEndpoint `json:"cos_endpoint" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateTargetOptions.TargetType property.
// The type of the target.
const (
	CreateTargetOptions_TargetType_Cos = "cos"
)

// NewCreateTargetOptions : Instantiate CreateTargetOptions
func (*AtrackerV1) NewCreateTargetOptions(name string, targetType string, cosEndpoint *CosEndpoint) *CreateTargetOptions {
	return &CreateTargetOptions{
		Name: core.StringPtr(name),
		TargetType: core.StringPtr(targetType),
		CosEndpoint: cosEndpoint,
	}
}

// SetName : Allow user to set Name
func (options *CreateTargetOptions) SetName(name string) *CreateTargetOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTargetType : Allow user to set TargetType
func (options *CreateTargetOptions) SetTargetType(targetType string) *CreateTargetOptions {
	options.TargetType = core.StringPtr(targetType)
	return options
}

// SetCosEndpoint : Allow user to set CosEndpoint
func (options *CreateTargetOptions) SetCosEndpoint(cosEndpoint *CosEndpoint) *CreateTargetOptions {
	options.CosEndpoint = cosEndpoint
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateTargetOptions) SetHeaders(param map[string]string) *CreateTargetOptions {
	options.Headers = param
	return options
}

// DeleteRouteOptions : The DeleteRoute options.
type DeleteRouteOptions struct {
	// The v4 UUID that uniquely identifies the route.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRouteOptions : Instantiate DeleteRouteOptions
func (*AtrackerV1) NewDeleteRouteOptions(id string) *DeleteRouteOptions {
	return &DeleteRouteOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteRouteOptions) SetID(id string) *DeleteRouteOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRouteOptions) SetHeaders(param map[string]string) *DeleteRouteOptions {
	options.Headers = param
	return options
}

// DeleteTargetOptions : The DeleteTarget options.
type DeleteTargetOptions struct {
	// The v4 UUID that uniquely identifies the target.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteTargetOptions : Instantiate DeleteTargetOptions
func (*AtrackerV1) NewDeleteTargetOptions(id string) *DeleteTargetOptions {
	return &DeleteTargetOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteTargetOptions) SetID(id string) *DeleteTargetOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTargetOptions) SetHeaders(param map[string]string) *DeleteTargetOptions {
	options.Headers = param
	return options
}

// GetRouteOptions : The GetRoute options.
type GetRouteOptions struct {
	// The v4 UUID that uniquely identifies the route.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRouteOptions : Instantiate GetRouteOptions
func (*AtrackerV1) NewGetRouteOptions(id string) *GetRouteOptions {
	return &GetRouteOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetRouteOptions) SetID(id string) *GetRouteOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRouteOptions) SetHeaders(param map[string]string) *GetRouteOptions {
	options.Headers = param
	return options
}

// GetTargetOptions : The GetTarget options.
type GetTargetOptions struct {
	// The v4 UUID that uniquely identifies the target.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTargetOptions : Instantiate GetTargetOptions
func (*AtrackerV1) NewGetTargetOptions(id string) *GetTargetOptions {
	return &GetTargetOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetTargetOptions) SetID(id string) *GetTargetOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTargetOptions) SetHeaders(param map[string]string) *GetTargetOptions {
	options.Headers = param
	return options
}

// ListRoutesOptions : The ListRoutes options.
type ListRoutesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRoutesOptions : Instantiate ListRoutesOptions
func (*AtrackerV1) NewListRoutesOptions() *ListRoutesOptions {
	return &ListRoutesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListRoutesOptions) SetHeaders(param map[string]string) *ListRoutesOptions {
	options.Headers = param
	return options
}

// ListTargetsOptions : The ListTargets options.
type ListTargetsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListTargetsOptions : Instantiate ListTargetsOptions
func (*AtrackerV1) NewListTargetsOptions() *ListTargetsOptions {
	return &ListTargetsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListTargetsOptions) SetHeaders(param map[string]string) *ListTargetsOptions {
	options.Headers = param
	return options
}

// ReplaceRouteOptions : The ReplaceRoute options.
type ReplaceRouteOptions struct {
	// The v4 UUID that uniquely identifies the route.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the route. Must be 180 characters or less and cannot include any special characters other than `(space)
	// - . _ :`.
	Name *string `json:"name" validate:"required"`

	// Whether or not all global events should be forwarded to this region.
	ReceiveGlobalEvents *bool `json:"receive_global_events" validate:"required"`

	// Routing rules that will be evaluated in their order of the array.
	Rules []Rule `json:"rules" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceRouteOptions : Instantiate ReplaceRouteOptions
func (*AtrackerV1) NewReplaceRouteOptions(id string, name string, receiveGlobalEvents bool, rules []Rule) *ReplaceRouteOptions {
	return &ReplaceRouteOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
		ReceiveGlobalEvents: core.BoolPtr(receiveGlobalEvents),
		Rules: rules,
	}
}

// SetID : Allow user to set ID
func (options *ReplaceRouteOptions) SetID(id string) *ReplaceRouteOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *ReplaceRouteOptions) SetName(name string) *ReplaceRouteOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetReceiveGlobalEvents : Allow user to set ReceiveGlobalEvents
func (options *ReplaceRouteOptions) SetReceiveGlobalEvents(receiveGlobalEvents bool) *ReplaceRouteOptions {
	options.ReceiveGlobalEvents = core.BoolPtr(receiveGlobalEvents)
	return options
}

// SetRules : Allow user to set Rules
func (options *ReplaceRouteOptions) SetRules(rules []Rule) *ReplaceRouteOptions {
	options.Rules = rules
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceRouteOptions) SetHeaders(param map[string]string) *ReplaceRouteOptions {
	options.Headers = param
	return options
}

// ReplaceTargetOptions : The ReplaceTarget options.
type ReplaceTargetOptions struct {
	// The v4 UUID that uniquely identifies the target.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the target. Must be 256 characters or less.
	Name *string `json:"name" validate:"required"`

	// The type of the target.
	TargetType *string `json:"target_type" validate:"required"`

	// Property values for a Cloud Object Storage Endpoint.
	CosEndpoint *CosEndpoint `json:"cos_endpoint" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ReplaceTargetOptions.TargetType property.
// The type of the target.
const (
	ReplaceTargetOptions_TargetType_Cos = "cos"
)

// NewReplaceTargetOptions : Instantiate ReplaceTargetOptions
func (*AtrackerV1) NewReplaceTargetOptions(id string, name string, targetType string, cosEndpoint *CosEndpoint) *ReplaceTargetOptions {
	return &ReplaceTargetOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
		TargetType: core.StringPtr(targetType),
		CosEndpoint: cosEndpoint,
	}
}

// SetID : Allow user to set ID
func (options *ReplaceTargetOptions) SetID(id string) *ReplaceTargetOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *ReplaceTargetOptions) SetName(name string) *ReplaceTargetOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetTargetType : Allow user to set TargetType
func (options *ReplaceTargetOptions) SetTargetType(targetType string) *ReplaceTargetOptions {
	options.TargetType = core.StringPtr(targetType)
	return options
}

// SetCosEndpoint : Allow user to set CosEndpoint
func (options *ReplaceTargetOptions) SetCosEndpoint(cosEndpoint *CosEndpoint) *ReplaceTargetOptions {
	options.CosEndpoint = cosEndpoint
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceTargetOptions) SetHeaders(param map[string]string) *ReplaceTargetOptions {
	options.Headers = param
	return options
}

// Route : The route resource.
type Route struct {
	// The uuid of this route resource.
	ID *string `json:"id,omitempty"`

	// The name of this route.
	Name *string `json:"name,omitempty"`

	// The uuid of ATracker services in this region.
	InstanceID *string `json:"instance_id,omitempty"`

	// The crn of this route type resource.
	Crn *string `json:"crn,omitempty"`

	// The version of this route.
	Version *int64 `json:"version,omitempty"`

	// Whether or not all global events should be forwarded to this region.
	ReceiveGlobalEvents *bool `json:"receive_global_events,omitempty"`

	// The routing rules that will be evaluated in their order of the array.
	Rules []Rule `json:"rules,omitempty"`
}


// UnmarshalRoute unmarshals an instance of Route from the specified map of raw messages.
func UnmarshalRoute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Route)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "receive_global_events", &obj.ReceiveGlobalEvents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RouteList : A list of route resources.
type RouteList struct {
	// A list of route resources.
	Routes []Route `json:"routes" validate:"required"`
}


// UnmarshalRouteList unmarshals an instance of RouteList from the specified map of raw messages.
func UnmarshalRouteList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RouteList)
	err = core.UnmarshalModel(m, "routes", &obj.Routes, UnmarshalRoute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Rule : The request payload to create a regional route.
type Rule struct {
	// The target ID List. Only one target id is supported. For regional route, the id must be V4 uuid of a target in the
	// same region. For global route, it will be region-code and target-id separated by colon.
	TargetIds []string `json:"target_ids" validate:"required"`
}


// NewRule : Instantiate Rule (Generic Model Constructor)
func (*AtrackerV1) NewRule(targetIds []string) (model *Rule, err error) {
	model = &Rule{
		TargetIds: targetIds,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "target_ids", &obj.TargetIds)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Target : Property values for a target in response. Credentials associated with the target are encrypted and masked as REDACTED
// in the response.
type Target struct {
	// The uuid of this target resource.
	ID *string `json:"id,omitempty"`

	// The name of this target resource.
	Name *string `json:"name,omitempty"`

	// The uuid of ATracker services in this region.
	InstanceID *string `json:"instance_id,omitempty"`

	// The crn of this target type resource.
	Crn *string `json:"crn,omitempty"`

	// The type of this target.
	TargetType *string `json:"target_type,omitempty"`

	// The encryption key used to encrypt events before ATracker services buffer them on storage. This credential will be
	// masked in the response.
	EncryptKey *string `json:"encrypt_key,omitempty"`

	// The COS endpoint information.
	CosEndpoint *TargetCosEndpoint `json:"cos_endpoint,omitempty"`
}

// Constants associated with the Target.TargetType property.
// The type of this target.
const (
	Target_TargetType_Cos = "cos"
)


// UnmarshalTarget unmarshals an instance of Target from the specified map of raw messages.
func UnmarshalTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Target)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "encrypt_key", &obj.EncryptKey)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cos_endpoint", &obj.CosEndpoint, UnmarshalTargetCosEndpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetCosEndpoint : The COS endpoint information.
type TargetCosEndpoint struct {
	// The host name of this COS endpoint.
	Endpoint *string `json:"endpoint,omitempty"`

	// The CRN of this COS instance.
	TargetCrn *string `json:"target_crn,omitempty"`

	// The bucket name under this COS instance.
	Bucket *string `json:"bucket,omitempty"`

	// The IAM Api key that have writer access to this cos instance. This credential will be masked in the response.
	ApiKey *string `json:"api_key,omitempty"`
}


// UnmarshalTargetCosEndpoint unmarshals an instance of TargetCosEndpoint from the specified map of raw messages.
func UnmarshalTargetCosEndpoint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetCosEndpoint)
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_crn", &obj.TargetCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TargetList : A list of target resources.
type TargetList struct {
	// A list of target resources.
	Targets []Target `json:"targets" validate:"required"`
}


// UnmarshalTargetList unmarshals an instance of TargetList from the specified map of raw messages.
func UnmarshalTargetList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetList)
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalTarget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CosEndpoint : Property values for a Cloud Object Storage Endpoint.
type CosEndpoint struct {
	// The host name of this COS endpoint.
	Endpoint *string `json:"endpoint" validate:"required"`

	// The CRN of this COS instance.
	TargetCrn *string `json:"target_crn" validate:"required"`

	// The bucket name under this COS instance.
	Bucket *string `json:"bucket" validate:"required"`

	// The IAM Api key that have writer access to this cos instance.
	ApiKey *string `json:"api_key" validate:"required"`
}


// NewCosEndpoint : Instantiate CosEndpoint (Generic Model Constructor)
func (*AtrackerV1) NewCosEndpoint(endpoint string, targetCrn string, bucket string, apiKey string) (model *CosEndpoint, err error) {
	model = &CosEndpoint{
		Endpoint: core.StringPtr(endpoint),
		TargetCrn: core.StringPtr(targetCrn),
		Bucket: core.StringPtr(bucket),
		ApiKey: core.StringPtr(apiKey),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCosEndpoint unmarshals an instance of CosEndpoint from the specified map of raw messages.
func UnmarshalCosEndpoint(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CosEndpoint)
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_crn", &obj.TargetCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "bucket", &obj.Bucket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_key", &obj.ApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
