/**
 * (C) Copyright IBM Corp. 2022.
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
 * IBM OpenAPI SDK Code Generator Version: 3.61.0-1667892a-20221109-194550
 */

// Package metricsrouterv3 : Operations and models for the MetricsRouterV3 service
package metricsrouterv3

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

// MetricsRouterV3 : IBM Cloud Metrics Router allows you to configure how to route platform metrics in your account.
//
// API Version: 3.0.0.preview
type MetricsRouterV3 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://au-syd.metrics-router.dev.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "metrics_router"

// MetricsRouterV3Options : Service options
type MetricsRouterV3Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewMetricsRouterV3UsingExternalConfig : constructs an instance of MetricsRouterV3 with passed in options and external configuration.
func NewMetricsRouterV3UsingExternalConfig(options *MetricsRouterV3Options) (metricsRouter *MetricsRouterV3, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	metricsRouter, err = NewMetricsRouterV3(options)
	if err != nil {
		return
	}

	err = metricsRouter.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = metricsRouter.Service.SetServiceURL(options.URL)
	}
	return
}

// NewMetricsRouterV3 : constructs an instance of MetricsRouterV3 with passed in options.
func NewMetricsRouterV3(options *MetricsRouterV3Options) (service *MetricsRouterV3, err error) {
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

	service = &MetricsRouterV3{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	var endpoints = map[string]string{
		"au-syd.dev":       "https://au-syd.metrics-router.dev.cloud.ibm.com",       // The server for IBM Cloud Metrics Router Service in the au-syd region dev environment.
		"us-south":         "https://us-south.metrics-router.cloud.ibm.com",         // The server for IBM Cloud Metrics Router Service in the us-south region.
		"private.us-south": "https://private.us-south.metrics-router.cloud.ibm.com", // The server for IBM Cloud Metrics Router Service in the us-south region.
		"us-east":          "https://us-east.metrics-router.cloud.ibm.com",          // The server for IBM Cloud Metrics Router Service in the us-east region.
		"private.us-east":  "https://private.us-east.metrics-router.cloud.ibm.com",  // The server for IBM Cloud Metrics Router Service in the us-east region.
		"eu-de":            "https://eu-de.metrics-router.cloud.ibm.com",            // The server for IBM Cloud Metrics Router Service in the eu-de region.
		"private.eu-de":    "https://private.eu-de.metrics-router.cloud.ibm.com",    // The server for IBM Cloud Metrics Router Service in the eu-de region.
		"eu-gb":            "https://eu-gb.metrics-router.cloud.ibm.com",            // The server for IBM Cloud Metrics Router Service in the eu-gb region.
		"private.eu-gb":    "https://private.eu-gb.metrics-router.cloud.ibm.com",    // The server for IBM Cloud Metrics Router Service in the eu-gb region.
		"au-syd":           "https://au-syd.metrics-router.cloud.ibm.com",           // The server for IBM Cloud Metrics Router Service in the au-syd region.
		"private.au-syd":   "https://private.au-syd.metrics-router.cloud.ibm.com",   // The server for IBM Cloud Metrics Router Service in the au-syd region.
	}

	if url, ok := endpoints[region]; ok {
		return url, nil
	}
	return "", fmt.Errorf("service URL for region '%s' not found", region)
}

// Clone makes a copy of "metricsRouter" suitable for processing requests.
func (metricsRouter *MetricsRouterV3) Clone() *MetricsRouterV3 {
	if core.IsNil(metricsRouter) {
		return nil
	}
	clone := *metricsRouter
	clone.Service = metricsRouter.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (metricsRouter *MetricsRouterV3) SetServiceURL(url string) error {
	return metricsRouter.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (metricsRouter *MetricsRouterV3) GetServiceURL() string {
	return metricsRouter.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (metricsRouter *MetricsRouterV3) SetDefaultHeaders(headers http.Header) {
	metricsRouter.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (metricsRouter *MetricsRouterV3) SetEnableGzipCompression(enableGzip bool) {
	metricsRouter.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (metricsRouter *MetricsRouterV3) GetEnableGzipCompression() bool {
	return metricsRouter.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (metricsRouter *MetricsRouterV3) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	metricsRouter.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (metricsRouter *MetricsRouterV3) DisableRetries() {
	metricsRouter.Service.DisableRetries()
}

// CreateTarget : Create a target
// Creates a target that includes information about the destination required to write platform metrics to that target.
// You can send your platform metrics from all regions to a single target, different targets or multiple targets. One
// target per region is not required. You can define up to 16 targets per account.
func (metricsRouter *MetricsRouterV3) CreateTarget(createTargetOptions *CreateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return metricsRouter.CreateTargetWithContext(context.Background(), createTargetOptions)
}

// CreateTargetWithContext is an alternate form of the CreateTarget method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) CreateTargetWithContext(ctx context.Context, createTargetOptions *CreateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "CreateTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createTargetOptions.Name != nil {
		body["name"] = createTargetOptions.Name
	}
	if createTargetOptions.DestinationCRN != nil {
		body["destination_crn"] = createTargetOptions.DestinationCRN
	}
	if createTargetOptions.Region != nil {
		body["region"] = createTargetOptions.Region
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
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListTargets : List targets
// List all targets that are defined for your account.
func (metricsRouter *MetricsRouterV3) ListTargets(listTargetsOptions *ListTargetsOptions) (result *TargetList, response *core.DetailedResponse, err error) {
	return metricsRouter.ListTargetsWithContext(context.Background(), listTargetsOptions)
}

// ListTargetsWithContext is an alternate form of the ListTargets method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ListTargetsWithContext(ctx context.Context, listTargetsOptions *ListTargetsOptions) (result *TargetList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTargetsOptions, "listTargetsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTargetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ListTargets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTargetList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetTarget : Get details of a target
// Retrieve the configuration details of a target.
func (metricsRouter *MetricsRouterV3) GetTarget(getTargetOptions *GetTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return metricsRouter.GetTargetWithContext(context.Background(), getTargetOptions)
}

// GetTargetWithContext is an alternate form of the GetTarget method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) GetTargetWithContext(ctx context.Context, getTargetOptions *GetTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "GetTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceTarget : Update a target
// Update the configuration details of a target.
func (metricsRouter *MetricsRouterV3) ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return metricsRouter.ReplaceTargetWithContext(context.Background(), replaceTargetOptions)
}

// ReplaceTargetWithContext is an alternate form of the ReplaceTarget method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ReplaceTargetWithContext(ctx context.Context, replaceTargetOptions *ReplaceTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ReplaceTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceTargetOptions.Name != nil {
		body["name"] = replaceTargetOptions.Name
	}
	if replaceTargetOptions.DestinationCRN != nil {
		body["destination_crn"] = replaceTargetOptions.DestinationCRN
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
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTarget : Delete a target
// Delete a target.
func (metricsRouter *MetricsRouterV3) DeleteTarget(deleteTargetOptions *DeleteTargetOptions) (result *WarningReport, response *core.DetailedResponse, err error) {
	return metricsRouter.DeleteTargetWithContext(context.Background(), deleteTargetOptions)
}

// DeleteTargetWithContext is an alternate form of the DeleteTarget method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) DeleteTargetWithContext(ctx context.Context, deleteTargetOptions *DeleteTargetOptions) (result *WarningReport, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "DeleteTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWarningReport)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ValidateTarget : Validate a target
// Validate a target by checking the credential to write to the target. The result is included as additional data of the
// target in the section "write_status".
func (metricsRouter *MetricsRouterV3) ValidateTarget(validateTargetOptions *ValidateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	return metricsRouter.ValidateTargetWithContext(context.Background(), validateTargetOptions)
}

// ValidateTargetWithContext is an alternate form of the ValidateTarget method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ValidateTargetWithContext(ctx context.Context, validateTargetOptions *ValidateTargetOptions) (result *Target, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(validateTargetOptions, "validateTargetOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(validateTargetOptions, "validateTargetOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"id": *validateTargetOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/targets/{id}/validate`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range validateTargetOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ValidateTarget")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTarget)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateRoute : Create a route
// Create a route with rules that specify how to manage platform metrics routing.
func (metricsRouter *MetricsRouterV3) CreateRoute(createRouteOptions *CreateRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return metricsRouter.CreateRouteWithContext(context.Background(), createRouteOptions)
}

// CreateRouteWithContext is an alternate form of the CreateRoute method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) CreateRouteWithContext(ctx context.Context, createRouteOptions *CreateRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/routes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "CreateRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createRouteOptions.Name != nil {
		body["name"] = createRouteOptions.Name
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
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRoutes : List routes
// List the route that is configured for an account.
func (metricsRouter *MetricsRouterV3) ListRoutes(listRoutesOptions *ListRoutesOptions) (result *RouteList, response *core.DetailedResponse, err error) {
	return metricsRouter.ListRoutesWithContext(context.Background(), listRoutesOptions)
}

// ListRoutesWithContext is an alternate form of the ListRoutes method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ListRoutesWithContext(ctx context.Context, listRoutesOptions *ListRoutesOptions) (result *RouteList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRoutesOptions, "listRoutesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/routes`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRoutesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ListRoutes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRouteList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRoute : Get details of a route
// Get the configuration details of a route.
func (metricsRouter *MetricsRouterV3) GetRoute(getRouteOptions *GetRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return metricsRouter.GetRouteWithContext(context.Background(), getRouteOptions)
}

// GetRouteWithContext is an alternate form of the GetRoute method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) GetRouteWithContext(ctx context.Context, getRouteOptions *GetRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "GetRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceRoute : Update a route
// Update the configuration details of a route.
func (metricsRouter *MetricsRouterV3) ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
	return metricsRouter.ReplaceRouteWithContext(context.Background(), replaceRouteOptions)
}

// ReplaceRouteWithContext is an alternate form of the ReplaceRoute method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ReplaceRouteWithContext(ctx context.Context, replaceRouteOptions *ReplaceRouteOptions) (result *Route, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ReplaceRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceRouteOptions.Name != nil {
		body["name"] = replaceRouteOptions.Name
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
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRoute)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRoute : Delete a route
// Deletes a route.
func (metricsRouter *MetricsRouterV3) DeleteRoute(deleteRouteOptions *DeleteRouteOptions) (response *core.DetailedResponse, err error) {
	return metricsRouter.DeleteRouteWithContext(context.Background(), deleteRouteOptions)
}

// DeleteRouteWithContext is an alternate form of the DeleteRoute method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) DeleteRouteWithContext(ctx context.Context, deleteRouteOptions *DeleteRouteOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/routes/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRouteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "DeleteRoute")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = metricsRouter.Service.Request(request, nil)

	return
}

// GetSettings : Get settings
// Get information about the current account level settings for Metrics Router service.
func (metricsRouter *MetricsRouterV3) GetSettings(getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return metricsRouter.GetSettingsWithContext(context.Background(), getSettingsOptions)
}

// GetSettingsWithContext is an alternate form of the GetSettings method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) GetSettingsWithContext(ctx context.Context, getSettingsOptions *GetSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSettingsOptions, "getSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "GetSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceSettings : Modify settings
// Modify the current account level settings such as default targets, permitted target regions, metadata region primary
// and secondary.
func (metricsRouter *MetricsRouterV3) ReplaceSettings(replaceSettingsOptions *ReplaceSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	return metricsRouter.ReplaceSettingsWithContext(context.Background(), replaceSettingsOptions)
}

// ReplaceSettingsWithContext is an alternate form of the ReplaceSettings method which supports a Context parameter
func (metricsRouter *MetricsRouterV3) ReplaceSettingsWithContext(ctx context.Context, replaceSettingsOptions *ReplaceSettingsOptions) (result *Settings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceSettingsOptions, "replaceSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceSettingsOptions, "replaceSettingsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = metricsRouter.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(metricsRouter.Service.Options.URL, `/api/v3/settings`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("metrics_router", "V3", "ReplaceSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceSettingsOptions.MetadataRegionPrimary != nil {
		body["metadata_region_primary"] = replaceSettingsOptions.MetadataRegionPrimary
	}
	if replaceSettingsOptions.PrivateAPIEndpointOnly != nil {
		body["private_api_endpoint_only"] = replaceSettingsOptions.PrivateAPIEndpointOnly
	}
	if replaceSettingsOptions.DefaultTargets != nil {
		body["default_targets"] = replaceSettingsOptions.DefaultTargets
	}
	if replaceSettingsOptions.PermittedTargetRegions != nil {
		body["permitted_target_regions"] = replaceSettingsOptions.PermittedTargetRegions
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
	response, err = metricsRouter.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateRouteOptions : The CreateRoute options.
type CreateRouteOptions struct {
	// The name of the route. The name must be 1000 characters or less and cannot include any special characters other than
	// `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.
	Name *string `json:"name" validate:"required"`

	// Routing rules that will be evaluated in their order of the array.
	Rules []RulePrototype `json:"rules" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRouteOptions : Instantiate CreateRouteOptions
func (*MetricsRouterV3) NewCreateRouteOptions(name string, rules []RulePrototype) *CreateRouteOptions {
	return &CreateRouteOptions{
		Name:  core.StringPtr(name),
		Rules: rules,
	}
}

// SetName : Allow user to set Name
func (_options *CreateRouteOptions) SetName(name string) *CreateRouteOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *CreateRouteOptions) SetRules(rules []RulePrototype) *CreateRouteOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRouteOptions) SetHeaders(param map[string]string) *CreateRouteOptions {
	options.Headers = param
	return options
}

// CreateTargetOptions : The CreateTarget options.
type CreateTargetOptions struct {
	// The name of the target. The name must be 1000 characters or less, and cannot include any special characters other
	// than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.
	Name *string `json:"name" validate:"required"`

	// The CRN of a destination service instance or resource.
	DestinationCRN *string `json:"destination_crn" validate:"required"`

	// Include this optional field if you want to create a target in a different region other than the one you are
	// connected.
	Region *string `json:"region,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateTargetOptions : Instantiate CreateTargetOptions
func (*MetricsRouterV3) NewCreateTargetOptions(name string, destinationCRN string) *CreateTargetOptions {
	return &CreateTargetOptions{
		Name:           core.StringPtr(name),
		DestinationCRN: core.StringPtr(destinationCRN),
	}
}

// SetName : Allow user to set Name
func (_options *CreateTargetOptions) SetName(name string) *CreateTargetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDestinationCRN : Allow user to set DestinationCRN
func (_options *CreateTargetOptions) SetDestinationCRN(destinationCRN string) *CreateTargetOptions {
	_options.DestinationCRN = core.StringPtr(destinationCRN)
	return _options
}

// SetRegion : Allow user to set Region
func (_options *CreateTargetOptions) SetRegion(region string) *CreateTargetOptions {
	_options.Region = core.StringPtr(region)
	return _options
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
func (*MetricsRouterV3) NewDeleteRouteOptions(id string) *DeleteRouteOptions {
	return &DeleteRouteOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteRouteOptions) SetID(id string) *DeleteRouteOptions {
	_options.ID = core.StringPtr(id)
	return _options
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
func (*MetricsRouterV3) NewDeleteTargetOptions(id string) *DeleteTargetOptions {
	return &DeleteTargetOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *DeleteTargetOptions) SetID(id string) *DeleteTargetOptions {
	_options.ID = core.StringPtr(id)
	return _options
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
func (*MetricsRouterV3) NewGetRouteOptions(id string) *GetRouteOptions {
	return &GetRouteOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetRouteOptions) SetID(id string) *GetRouteOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetRouteOptions) SetHeaders(param map[string]string) *GetRouteOptions {
	options.Headers = param
	return options
}

// GetSettingsOptions : The GetSettings options.
type GetSettingsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSettingsOptions : Instantiate GetSettingsOptions
func (*MetricsRouterV3) NewGetSettingsOptions() *GetSettingsOptions {
	return &GetSettingsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSettingsOptions) SetHeaders(param map[string]string) *GetSettingsOptions {
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
func (*MetricsRouterV3) NewGetTargetOptions(id string) *GetTargetOptions {
	return &GetTargetOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *GetTargetOptions) SetID(id string) *GetTargetOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetTargetOptions) SetHeaders(param map[string]string) *GetTargetOptions {
	options.Headers = param
	return options
}

// InclusionFilter : A list of conditions to be satisfied for routing events to pre-defined target.
type InclusionFilter struct {
	// Part of CRN that can be compared with values.
	Operand *string `json:"operand" validate:"required"`

	// The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can
	// support upto 20 values in the array.
	Operator *string `json:"operator" validate:"required"`

	// The provided values of the in operand to be compared with. This can be string or array of string.
	Value interface{} `json:"value" validate:"required"`
}

// Constants associated with the InclusionFilter.Operand property.
// Part of CRN that can be compared with values.
const (
	InclusionFilterOperandLocationConst        = "location"
	InclusionFilterOperandResourceConst        = "resource"
	InclusionFilterOperandResourceTypeConst    = "resource_type"
	InclusionFilterOperandServiceInstanceConst = "service_instance"
	InclusionFilterOperandServiceNameConst     = "service_name"
)

// Constants associated with the InclusionFilter.Operator property.
// The operation to be performed between operand and the provided values. 'is' to be used with one value and 'in' can
// support upto 20 values in the array.
const (
	InclusionFilterOperatorInConst = "in"
	InclusionFilterOperatorIsConst = "is"
)

// NewInclusionFilter : Instantiate InclusionFilter (Generic Model Constructor)
func (*MetricsRouterV3) NewInclusionFilter(operand string, operator string, value []string) (_model *InclusionFilter, err error) {
	_model = &InclusionFilter{
		Operand:  core.StringPtr(operand),
		Operator: core.StringPtr(operator),
		Value:    value,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalInclusionFilter unmarshals an instance of InclusionFilter from the specified map of raw messages.
func UnmarshalInclusionFilter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InclusionFilter)
	err = core.UnmarshalPrimitive(m, "operand", &obj.Operand)
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

// ListRoutesOptions : The ListRoutes options.
type ListRoutesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRoutesOptions : Instantiate ListRoutesOptions
func (*MetricsRouterV3) NewListRoutesOptions() *ListRoutesOptions {
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
func (*MetricsRouterV3) NewListTargetsOptions() *ListTargetsOptions {
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

	// The name of the route. The name must be 1000 characters or less and cannot include any special characters other than
	// `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.
	Name *string `json:"name" validate:"required"`

	// Routing rules that will be evaluated in their order of the array.
	Rules []RulePrototype `json:"rules" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceRouteOptions : Instantiate ReplaceRouteOptions
func (*MetricsRouterV3) NewReplaceRouteOptions(id string, name string, rules []RulePrototype) *ReplaceRouteOptions {
	return &ReplaceRouteOptions{
		ID:    core.StringPtr(id),
		Name:  core.StringPtr(name),
		Rules: rules,
	}
}

// SetID : Allow user to set ID
func (_options *ReplaceRouteOptions) SetID(id string) *ReplaceRouteOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *ReplaceRouteOptions) SetName(name string) *ReplaceRouteOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRules : Allow user to set Rules
func (_options *ReplaceRouteOptions) SetRules(rules []RulePrototype) *ReplaceRouteOptions {
	_options.Rules = rules
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceRouteOptions) SetHeaders(param map[string]string) *ReplaceRouteOptions {
	options.Headers = param
	return options
}

// ReplaceSettingsOptions : The ReplaceSettings options.
type ReplaceSettingsOptions struct {
	// To store all your meta data in a single region.
	MetadataRegionPrimary *string `json:"metadata_region_primary" validate:"required"`

	// If you set this true then you cannot access api through public network.
	PrivateAPIEndpointOnly *bool `json:"private_api_endpoint_only" validate:"required"`

	// The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will
	// receive the event.
	DefaultTargets []string `json:"default_targets,omitempty"`

	// If present then only these regions may be used to define a target.
	PermittedTargetRegions []string `json:"permitted_target_regions,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceSettingsOptions : Instantiate ReplaceSettingsOptions
func (*MetricsRouterV3) NewReplaceSettingsOptions(metadataRegionPrimary string, privateAPIEndpointOnly bool) *ReplaceSettingsOptions {
	return &ReplaceSettingsOptions{
		MetadataRegionPrimary:  core.StringPtr(metadataRegionPrimary),
		PrivateAPIEndpointOnly: core.BoolPtr(privateAPIEndpointOnly),
	}
}

// SetMetadataRegionPrimary : Allow user to set MetadataRegionPrimary
func (_options *ReplaceSettingsOptions) SetMetadataRegionPrimary(metadataRegionPrimary string) *ReplaceSettingsOptions {
	_options.MetadataRegionPrimary = core.StringPtr(metadataRegionPrimary)
	return _options
}

// SetPrivateAPIEndpointOnly : Allow user to set PrivateAPIEndpointOnly
func (_options *ReplaceSettingsOptions) SetPrivateAPIEndpointOnly(privateAPIEndpointOnly bool) *ReplaceSettingsOptions {
	_options.PrivateAPIEndpointOnly = core.BoolPtr(privateAPIEndpointOnly)
	return _options
}

// SetDefaultTargets : Allow user to set DefaultTargets
func (_options *ReplaceSettingsOptions) SetDefaultTargets(defaultTargets []string) *ReplaceSettingsOptions {
	_options.DefaultTargets = defaultTargets
	return _options
}

// SetPermittedTargetRegions : Allow user to set PermittedTargetRegions
func (_options *ReplaceSettingsOptions) SetPermittedTargetRegions(permittedTargetRegions []string) *ReplaceSettingsOptions {
	_options.PermittedTargetRegions = permittedTargetRegions
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceSettingsOptions) SetHeaders(param map[string]string) *ReplaceSettingsOptions {
	options.Headers = param
	return options
}

// ReplaceTargetOptions : The ReplaceTarget options.
type ReplaceTargetOptions struct {
	// The v4 UUID that uniquely identifies the target.
	ID *string `json:"id" validate:"required,ne="`

	// The name of the target. The name must be 1000 characters or less, and cannot include any special characters other
	// than `(space) - . _ :`. Do not include any personal identifying information (PII) in any resource names.
	Name *string `json:"name" validate:"required"`

	// The CRN of the destination service instance or resource.
	DestinationCRN *string `json:"destination_crn" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceTargetOptions : Instantiate ReplaceTargetOptions
func (*MetricsRouterV3) NewReplaceTargetOptions(id string, name string, destinationCRN string) *ReplaceTargetOptions {
	return &ReplaceTargetOptions{
		ID:             core.StringPtr(id),
		Name:           core.StringPtr(name),
		DestinationCRN: core.StringPtr(destinationCRN),
	}
}

// SetID : Allow user to set ID
func (_options *ReplaceTargetOptions) SetID(id string) *ReplaceTargetOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetName : Allow user to set Name
func (_options *ReplaceTargetOptions) SetName(name string) *ReplaceTargetOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDestinationCRN : Allow user to set DestinationCRN
func (_options *ReplaceTargetOptions) SetDestinationCRN(destinationCRN string) *ReplaceTargetOptions {
	_options.DestinationCRN = core.StringPtr(destinationCRN)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceTargetOptions) SetHeaders(param map[string]string) *ReplaceTargetOptions {
	options.Headers = param
	return options
}

// Route : The route resource. The scope of the route is account wide. That means all the routes are evaluated in all regions,
// except the ones limited by region.
type Route struct {
	// The UUID of the route resource.
	ID *string `json:"id" validate:"required"`

	// The name of the route.
	Name *string `json:"name" validate:"required"`

	// The crn of the route resource.
	CRN *string `json:"crn" validate:"required"`

	// The version of the route.
	Version *int64 `json:"version,omitempty"`

	// The routing rules that will be evaluated in their order of the array. Once a rule is matched, the remaining rules in
	// the route definition will be skipped.
	Rules []Rule `json:"rules" validate:"required"`

	// The timestamp of the route creation time.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The timestamp of the route last updated time.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The API version of the route.
	APIVersion *int64 `json:"api_version" validate:"required"`

	// An optional message containing information about the route.
	Message *string `json:"message,omitempty"`
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
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_version", &obj.APIVersion)
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

// Rule : A configuration to route events to pre-defined target.
type Rule struct {
	// The target ID List. All the events will be sent to all targets listed in the rule. You can include targets from
	// other regions.
	TargetIds []string `json:"target_ids" validate:"required"`

	// A list of conditions to be satisfied for routing events to pre-defined target.
	InclusionFilters []InclusionFilter `json:"inclusion_filters" validate:"required"`
}

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "target_ids", &obj.TargetIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "inclusion_filters", &obj.InclusionFilters, UnmarshalInclusionFilter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RulePrototype : A configuration to route events to pre-defined target.
type RulePrototype struct {
	// The target ID List. All the events will be sent to all targets listed in the rule. You can include targets from
	// other regions.
	TargetIds []string `json:"target_ids" validate:"required"`

	// A list of conditions to be satisfied for routing events to pre-defined target.
	InclusionFilters []InclusionFilter `json:"inclusion_filters" validate:"required"`
}

// NewRulePrototype : Instantiate RulePrototype (Generic Model Constructor)
func (*MetricsRouterV3) NewRulePrototype(targetIds []string, inclusionFilters []InclusionFilter) (_model *RulePrototype, err error) {
	_model = &RulePrototype{
		TargetIds:        targetIds,
		InclusionFilters: inclusionFilters,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRulePrototype unmarshals an instance of RulePrototype from the specified map of raw messages.
func UnmarshalRulePrototype(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RulePrototype)
	err = core.UnmarshalPrimitive(m, "target_ids", &obj.TargetIds)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "inclusion_filters", &obj.InclusionFilters, UnmarshalInclusionFilter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Settings : Metrics Router settings response.
type Settings struct {
	// The target ID List. In the event that no routing rule causes the event to be sent to a target, these targets will
	// receive the event.
	DefaultTargets []string `json:"default_targets" validate:"required"`

	// If present then only these regions may be used to define a target.
	PermittedTargetRegions []string `json:"permitted_target_regions" validate:"required"`

	// To store all your meta data in a single region.
	MetadataRegionPrimary *string `json:"metadata_region_primary" validate:"required"`

	// If you set this true then you cannot access api through public network.
	PrivateAPIEndpointOnly *bool `json:"private_api_endpoint_only" validate:"required"`

	// The lowest API version of targets or routes that customer might have under his or her account.
	APIVersion *int64 `json:"api_version" validate:"required"`

	// An optional message containing information about the audit log locations.
	Message *string `json:"message,omitempty"`
}

// UnmarshalSettings unmarshals an instance of Settings from the specified map of raw messages.
func UnmarshalSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Settings)
	err = core.UnmarshalPrimitive(m, "default_targets", &obj.DefaultTargets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permitted_target_regions", &obj.PermittedTargetRegions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata_region_primary", &obj.MetadataRegionPrimary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "private_api_endpoint_only", &obj.PrivateAPIEndpointOnly)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_version", &obj.APIVersion)
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

// Target : Property values for a target in responses.
type Target struct {
	// The UUID of the target resource.
	ID *string `json:"id" validate:"required"`

	// The name of the target resource.
	Name *string `json:"name" validate:"required"`

	// The crn of the target resource.
	CRN *string `json:"crn" validate:"required"`

	// The CRN of the destination service instance or resource.
	DestinationCRN *string `json:"destination_crn" validate:"required"`

	// The type of the target.
	TargetType *string `json:"target_type" validate:"required"`

	// Included this optional field if you used it to create a target in a different region other than the one you are
	// connected.
	Region *string `json:"region,omitempty"`

	// The status of the write attempt to the target with the provided destination info.
	WriteStatus *WriteStatus `json:"write_status" validate:"required"`

	// The timestamp of the target creation time.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// The timestamp of the target last updated time.
	UpdatedAt *strfmt.DateTime `json:"updated_at" validate:"required"`

	// The API version of the target.
	APIVersion *int64 `json:"api_version" validate:"required"`
}

// Constants associated with the Target.TargetType property.
// The type of the target.
const (
	TargetTargetTypeSysdigMonitorConst = "sysdig_monitor"
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
	err = core.UnmarshalPrimitive(m, "crn", &obj.CRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destination_crn", &obj.DestinationCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_type", &obj.TargetType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "write_status", &obj.WriteStatus, UnmarshalWriteStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "api_version", &obj.APIVersion)
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

// ValidateTargetOptions : The ValidateTarget options.
type ValidateTargetOptions struct {
	// The v4 UUID that uniquely identifies the target.
	ID *string `json:"id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewValidateTargetOptions : Instantiate ValidateTargetOptions
func (*MetricsRouterV3) NewValidateTargetOptions(id string) *ValidateTargetOptions {
	return &ValidateTargetOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (_options *ValidateTargetOptions) SetID(id string) *ValidateTargetOptions {
	_options.ID = core.StringPtr(id)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ValidateTargetOptions) SetHeaders(param map[string]string) *ValidateTargetOptions {
	options.Headers = param
	return options
}

// Warning : The warning object.
type Warning struct {
	// The warning code.
	Code *string `json:"code" validate:"required"`

	// The warning message.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalWarning unmarshals an instance of Warning from the specified map of raw messages.
func UnmarshalWarning(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Warning)
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

// WarningReport : Description of a warning that occurred in a service request.
type WarningReport struct {
	// The status code.
	StatusCode *int64 `json:"status_code,omitempty"`

	// The transaction-id of the API request.
	Trace *string `json:"trace,omitempty"`

	// The warning array triggered by the API request.
	Warnings []Warning `json:"warnings,omitempty"`
}

// UnmarshalWarningReport unmarshals an instance of WarningReport from the specified map of raw messages.
func UnmarshalWarningReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WarningReport)
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "warnings", &obj.Warnings, UnmarshalWarning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// WriteStatus : The status of the write attempt to the target with the provided destination info.
type WriteStatus struct {
	// The status such as failed or success.
	Status *string `json:"status" validate:"required"`

	// The timestamp of the failure.
	LastFailure *strfmt.DateTime `json:"last_failure,omitempty"`

	// Detailed description of the cause of the failure.
	ReasonForLastFailure *string `json:"reason_for_last_failure,omitempty"`
}

// UnmarshalWriteStatus unmarshals an instance of WriteStatus from the specified map of raw messages.
func UnmarshalWriteStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(WriteStatus)
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_failure", &obj.LastFailure)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reason_for_last_failure", &obj.ReasonForLastFailure)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
