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
 * IBM OpenAPI SDK Code Generator Version: 3.38.1-1037b405-20210908-184149
 */

// Package contextbasedrestrictionsv1 : Operations and models for the ContextBasedRestrictionsV1 service
package contextbasedrestrictionsv1

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

// ContextBasedRestrictionsV1 : With the Context Based Restrictions Admin API, you can:
// * Create, list, get, update, and delete zones
// * Create, list, get, update, and delete rules
// * Get account settings
//
// API Version: 1.0.0
type ContextBasedRestrictionsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://cbr.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "context_based_restrictions"

// Options : Service options
type Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewContextBasedRestrictionsV1UsingExternalConfig : constructs an instance of ContextBasedRestrictionsV1 with passed in options and external configuration.
func NewContextBasedRestrictionsV1UsingExternalConfig(options *Options) (contextBasedRestrictions *ContextBasedRestrictionsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	contextBasedRestrictions, err = NewContextBasedRestrictionsV1(options)
	if err != nil {
		return
	}

	err = contextBasedRestrictions.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = contextBasedRestrictions.Service.SetServiceURL(options.URL)
	}
	return
}

// NewContextBasedRestrictionsV1 : constructs an instance of ContextBasedRestrictionsV1 with passed in options.
func NewContextBasedRestrictionsV1(options *Options) (service *ContextBasedRestrictionsV1, err error) {
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

	service = &ContextBasedRestrictionsV1{
		Service: baseService,
	}

	return
}

// Clone makes a copy of "contextBasedRestrictions" suitable for processing requests.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) Clone() *ContextBasedRestrictionsV1 {
	if core.IsNil(contextBasedRestrictions) {
		return nil
	}
	clone := *contextBasedRestrictions
	clone.Service = contextBasedRestrictions.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (contextBasedRestrictions *ContextBasedRestrictionsV1) SetServiceURL(url string) error {
	return contextBasedRestrictions.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetServiceURL() string {
	return contextBasedRestrictions.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (contextBasedRestrictions *ContextBasedRestrictionsV1) SetDefaultHeaders(headers http.Header) {
	contextBasedRestrictions.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (contextBasedRestrictions *ContextBasedRestrictionsV1) SetEnableGzipCompression(enableGzip bool) {
	contextBasedRestrictions.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetEnableGzipCompression() bool {
	return contextBasedRestrictions.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	contextBasedRestrictions.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) DisableRetries() {
	contextBasedRestrictions.Service.DisableRetries()
}

// CreateZone : Create a zone
// Creates a zone for the specified account.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) CreateZone(createZoneOptions *CreateZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.CreateZoneWithContext(context.Background(), createZoneOptions)
}

// CreateZoneWithContext is an alternate form of the CreateZone method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) CreateZoneWithContext(ctx context.Context, createZoneOptions *CreateZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createZoneOptions, "createZoneOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "CreateZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createZoneOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createZoneOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createZoneOptions.Name != nil {
		body["name"] = createZoneOptions.Name
	}
	if createZoneOptions.AccountID != nil {
		body["account_id"] = createZoneOptions.AccountID
	}
	if createZoneOptions.Description != nil {
		body["description"] = createZoneOptions.Description
	}
	if createZoneOptions.Addresses != nil {
		body["addresses"] = createZoneOptions.Addresses
	}
	if createZoneOptions.Excluded != nil {
		body["excluded"] = createZoneOptions.Excluded
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
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutZone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListZones : List zones
// Lists zones for the specified account.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListZones(listZonesOptions *ListZonesOptions) (result *OutZonePage, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.ListZonesWithContext(context.Background(), listZonesOptions)
}

// ListZonesWithContext is an alternate form of the ListZones method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListZonesWithContext(ctx context.Context, listZonesOptions *ListZonesOptions) (result *OutZonePage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listZonesOptions, "listZonesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listZonesOptions, "listZonesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listZonesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "ListZones")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listZonesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listZonesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listZonesOptions.AccountID))
	if listZonesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listZonesOptions.Name))
	}
	if listZonesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listZonesOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutZonePage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetZone : Get the specified zone
// Gets the zone for the specified ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetZone(getZoneOptions *GetZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.GetZoneWithContext(context.Background(), getZoneOptions)
}

// GetZoneWithContext is an alternate form of the GetZone method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetZoneWithContext(ctx context.Context, getZoneOptions *GetZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneOptions, "getZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneOptions, "getZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"zone_id": *getZoneOptions.ZoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones/{zone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "GetZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getZoneOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getZoneOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutZone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceZone : Update the specified zone
// Updates the zone for the specified ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.ReplaceZoneWithContext(context.Background(), replaceZoneOptions)
}

// ReplaceZoneWithContext is an alternate form of the ReplaceZone method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ReplaceZoneWithContext(ctx context.Context, replaceZoneOptions *ReplaceZoneOptions) (result *OutZone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceZoneOptions, "replaceZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceZoneOptions, "replaceZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"zone_id": *replaceZoneOptions.ZoneID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones/{zone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "ReplaceZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceZoneOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*replaceZoneOptions.IfMatch))
	}
	if replaceZoneOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceZoneOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceZoneOptions.Name != nil {
		body["name"] = replaceZoneOptions.Name
	}
	if replaceZoneOptions.AccountID != nil {
		body["account_id"] = replaceZoneOptions.AccountID
	}
	if replaceZoneOptions.Description != nil {
		body["description"] = replaceZoneOptions.Description
	}
	if replaceZoneOptions.Addresses != nil {
		body["addresses"] = replaceZoneOptions.Addresses
	}
	if replaceZoneOptions.Excluded != nil {
		body["excluded"] = replaceZoneOptions.Excluded
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
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutZone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteZone : Delete the specified zone
// Deletes the zone for the specified home ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) DeleteZone(deleteZoneOptions *DeleteZoneOptions) (response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.DeleteZoneWithContext(context.Background(), deleteZoneOptions)
}

// DeleteZoneWithContext is an alternate form of the DeleteZone method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) DeleteZoneWithContext(ctx context.Context, deleteZoneOptions *DeleteZoneOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneOptions, "deleteZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteZoneOptions, "deleteZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"zone_id": *deleteZoneOptions.ZoneID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones/{zone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "DeleteZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteZoneOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteZoneOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = contextBasedRestrictions.Service.Request(request, nil)

	return
}

// ListAvailableServiceRefTargets : List available service reference targets
// Lists all available service reference targets.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions) (result *ServiceRefTargetPage, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.ListAvailableServiceRefTargetsWithContext(context.Background(), listAvailableServiceRefTargetsOptions)
}

// ListAvailableServiceRefTargetsWithContext is an alternate form of the ListAvailableServiceRefTargets method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListAvailableServiceRefTargetsWithContext(ctx context.Context, listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions) (result *ServiceRefTargetPage, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAvailableServiceRefTargetsOptions, "listAvailableServiceRefTargetsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/zones/serviceref_targets`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAvailableServiceRefTargetsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "ListAvailableServiceRefTargets")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAvailableServiceRefTargetsOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listAvailableServiceRefTargetsOptions.Type))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalServiceRefTargetPage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateRule : Create a rule
// Creates a rule for the specified account.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) CreateRule(createRuleOptions *CreateRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.CreateRuleWithContext(context.Background(), createRuleOptions)
}

// CreateRuleWithContext is an alternate form of the CreateRule method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) CreateRuleWithContext(ctx context.Context, createRuleOptions *CreateRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createRuleOptions, "createRuleOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "CreateRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createRuleOptions.Description != nil {
		body["description"] = createRuleOptions.Description
	}
	if createRuleOptions.Contexts != nil {
		body["contexts"] = createRuleOptions.Contexts
	}
	if createRuleOptions.Resources != nil {
		body["resources"] = createRuleOptions.Resources
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
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRules : List rules
// Lists rules for the specified account.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListRules(listRulesOptions *ListRulesOptions) (result *OutRulePage, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.ListRulesWithContext(context.Background(), listRulesOptions)
}

// ListRulesWithContext is an alternate form of the ListRules method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ListRulesWithContext(ctx context.Context, listRulesOptions *ListRulesOptions) (result *OutRulePage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listRulesOptions, "listRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listRulesOptions, "listRulesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "ListRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listRulesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listRulesOptions.AccountID))
	if listRulesOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*listRulesOptions.Region))
	}
	if listRulesOptions.Resource != nil {
		builder.AddQuery("resource", fmt.Sprint(*listRulesOptions.Resource))
	}
	if listRulesOptions.ResourceType != nil {
		builder.AddQuery("resource_type", fmt.Sprint(*listRulesOptions.ResourceType))
	}
	if listRulesOptions.ServiceInstance != nil {
		builder.AddQuery("service_instance", fmt.Sprint(*listRulesOptions.ServiceInstance))
	}
	if listRulesOptions.ServiceName != nil {
		builder.AddQuery("service_name", fmt.Sprint(*listRulesOptions.ServiceName))
	}
	if listRulesOptions.ServiceType != nil {
		builder.AddQuery("service_type", fmt.Sprint(*listRulesOptions.ServiceType))
	}
	if listRulesOptions.ZoneID != nil {
		builder.AddQuery("zone_id", fmt.Sprint(*listRulesOptions.ZoneID))
	}
	if listRulesOptions.Sort != nil {
		builder.AddQuery("sort", fmt.Sprint(*listRulesOptions.Sort))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutRulePage)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRule : Get the specified rule
// Gets the rule for the specified ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetRule(getRuleOptions *GetRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.GetRuleWithContext(context.Background(), getRuleOptions)
}

// GetRuleWithContext is an alternate form of the GetRule method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetRuleWithContext(ctx context.Context, getRuleOptions *GetRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRuleOptions, "getRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRuleOptions, "getRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *getRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "GetRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceRule : Update the specified rule
// Updates the rule for the specified ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.ReplaceRuleWithContext(context.Background(), replaceRuleOptions)
}

// ReplaceRuleWithContext is an alternate form of the ReplaceRule method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) ReplaceRuleWithContext(ctx context.Context, replaceRuleOptions *ReplaceRuleOptions) (result *OutRule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceRuleOptions, "replaceRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceRuleOptions, "replaceRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *replaceRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "ReplaceRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceRuleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*replaceRuleOptions.IfMatch))
	}
	if replaceRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*replaceRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if replaceRuleOptions.Description != nil {
		body["description"] = replaceRuleOptions.Description
	}
	if replaceRuleOptions.Contexts != nil {
		body["contexts"] = replaceRuleOptions.Contexts
	}
	if replaceRuleOptions.Resources != nil {
		body["resources"] = replaceRuleOptions.Resources
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
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRule : Delete the specified rule
// Deletes the rule for the specified home ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) DeleteRule(deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.DeleteRuleWithContext(context.Background(), deleteRuleOptions)
}

// DeleteRuleWithContext is an alternate form of the DeleteRule method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) DeleteRuleWithContext(ctx context.Context, deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRuleOptions, "deleteRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRuleOptions, "deleteRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *deleteRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "DeleteRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteRuleOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = contextBasedRestrictions.Service.Request(request, nil)

	return
}

// GetAccountSettings : Get the specified account settings
// Gets the settings for the specified account ID.
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) (result *OutAccountSettings, response *core.DetailedResponse, err error) {
	return contextBasedRestrictions.GetAccountSettingsWithContext(context.Background(), getAccountSettingsOptions)
}

// GetAccountSettingsWithContext is an alternate form of the GetAccountSettings method which supports a Context parameter
func (contextBasedRestrictions *ContextBasedRestrictionsV1) GetAccountSettingsWithContext(ctx context.Context, getAccountSettingsOptions *GetAccountSettingsOptions) (result *OutAccountSettings, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSettingsOptions, "getAccountSettingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSettingsOptions, "getAccountSettingsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id": *getAccountSettingsOptions.AccountID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = contextBasedRestrictions.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(contextBasedRestrictions.Service.Options.URL, `/v1/account_settings/{account_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountSettingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("context_based_restrictions", "V1", "GetAccountSettings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAccountSettingsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAccountSettingsOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = contextBasedRestrictions.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOutAccountSettings)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Address : A zone address.
// Models which "extend" this model:
// - AddressIPAddress
// - AddressIPAddressRange
// - AddressSubnet
// - AddressVPC
// - AddressServiceRef
type Address struct {
	// The type of address.
	Type *string `json:"type,omitempty"`

	// The IP address.
	Value *string `json:"value,omitempty"`

	// A service reference value.
	Ref *ServiceRefValue `json:"ref,omitempty"`
}

// Constants associated with the Address.Type property.
// The type of address.
const (
	AddressTypeIPAddress  = "ipAddress"
	AddressTypeIPRange    = "ipRange"
	AddressTypeServiceRef = "serviceRef"
	AddressTypeSubnet     = "subnet"
	AddressTypeVPC        = "vpc"
)

func (*Address) isaAddress() bool {
	return true
}

type AddressIntf interface {
	isaAddress() bool
}

// UnmarshalAddress unmarshals an instance of Address from the specified map of raw messages.
func UnmarshalAddress(m map[string]json.RawMessage, result interface{}) (err error) {
	// Retrieve discriminator value to determine correct "subclass".
	var discValue string
	err = core.UnmarshalPrimitive(m, "type", &discValue)
	if err != nil {
		err = fmt.Errorf("error unmarshalling discriminator property 'type': %s", err.Error())
		return
	}
	if discValue == "" {
		err = fmt.Errorf("required discriminator property 'type' not found in JSON object")
		return
	}
	if discValue == "ipAddress" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddressIPAddress)
	} else if discValue == "ipRange" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddressIPAddressRange)
	} else if discValue == "subnet" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddressSubnet)
	} else if discValue == "vpc" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddressVPC)
	} else if discValue == "serviceRef" {
		err = core.UnmarshalModel(m, "", result, UnmarshalAddressServiceRef)
	} else {
		err = fmt.Errorf("unrecognized value for discriminator property 'type': %s", discValue)
	}
	return
}

// CreateRuleOptions : The CreateRule options.
type CreateRuleOptions struct {
	// The description of the rule.
	Description *string `json:"description,omitempty"`

	// The contexts this rule applies to.
	Contexts []RuleContext `json:"contexts,omitempty"`

	// The resources this rule apply to.
	Resources []Resource `json:"resources,omitempty"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRuleOptions : Instantiate CreateRuleOptions
func (*ContextBasedRestrictionsV1) NewCreateRuleOptions() *CreateRuleOptions {
	return &CreateRuleOptions{}
}

// SetDescription : Allow user to set Description
func (options *CreateRuleOptions) SetDescription(description string) *CreateRuleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetContexts : Allow user to set Contexts
func (options *CreateRuleOptions) SetContexts(contexts []RuleContext) *CreateRuleOptions {
	options.Contexts = contexts
	return options
}

// SetResources : Allow user to set Resources
func (options *CreateRuleOptions) SetResources(resources []Resource) *CreateRuleOptions {
	options.Resources = resources
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateRuleOptions) SetTransactionID(transactionID string) *CreateRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRuleOptions) SetHeaders(param map[string]string) *CreateRuleOptions {
	options.Headers = param
	return options
}

// CreateZoneOptions : The CreateZone options.
type CreateZoneOptions struct {
	// The name of the zone.
	Name *string `json:"name,omitempty"`

	// The id of the account owning this zone.
	AccountID *string `json:"account_id,omitempty"`

	// The description of the zone.
	Description *string `json:"description,omitempty"`

	// The list of addresses in the zone.
	Addresses []AddressIntf `json:"addresses,omitempty"`

	// The list of excluded addresses in the zone.
	Excluded []AddressIntf `json:"excluded,omitempty"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateZoneOptions : Instantiate CreateZoneOptions
func (*ContextBasedRestrictionsV1) NewCreateZoneOptions() *CreateZoneOptions {
	return &CreateZoneOptions{}
}

// SetName : Allow user to set Name
func (options *CreateZoneOptions) SetName(name string) *CreateZoneOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *CreateZoneOptions) SetAccountID(accountID string) *CreateZoneOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateZoneOptions) SetDescription(description string) *CreateZoneOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetAddresses : Allow user to set Addresses
func (options *CreateZoneOptions) SetAddresses(addresses []AddressIntf) *CreateZoneOptions {
	options.Addresses = addresses
	return options
}

// SetExcluded : Allow user to set Excluded
func (options *CreateZoneOptions) SetExcluded(excluded []AddressIntf) *CreateZoneOptions {
	options.Excluded = excluded
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateZoneOptions) SetTransactionID(transactionID string) *CreateZoneOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateZoneOptions) SetHeaders(param map[string]string) *CreateZoneOptions {
	options.Headers = param
	return options
}

// DeleteRuleOptions : The DeleteRule options.
type DeleteRuleOptions struct {
	// The ID of a rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleOptions : Instantiate DeleteRuleOptions
func (*ContextBasedRestrictionsV1) NewDeleteRuleOptions(ruleID string) *DeleteRuleOptions {
	return &DeleteRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *DeleteRuleOptions) SetRuleID(ruleID string) *DeleteRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteRuleOptions) SetTransactionID(transactionID string) *DeleteRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRuleOptions) SetHeaders(param map[string]string) *DeleteRuleOptions {
	options.Headers = param
	return options
}

// DeleteZoneOptions : The DeleteZone options.
type DeleteZoneOptions struct {
	// The ID of a zone.
	ZoneID *string `json:"-" validate:"required,ne="`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteZoneOptions : Instantiate DeleteZoneOptions
func (*ContextBasedRestrictionsV1) NewDeleteZoneOptions(zoneID string) *DeleteZoneOptions {
	return &DeleteZoneOptions{
		ZoneID: core.StringPtr(zoneID),
	}
}

// SetZoneID : Allow user to set ZoneID
func (options *DeleteZoneOptions) SetZoneID(zoneID string) *DeleteZoneOptions {
	options.ZoneID = core.StringPtr(zoneID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteZoneOptions) SetTransactionID(transactionID string) *DeleteZoneOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneOptions) SetHeaders(param map[string]string) *DeleteZoneOptions {
	options.Headers = param
	return options
}

// GetAccountSettingsOptions : The GetAccountSettings options.
type GetAccountSettingsOptions struct {
	// The ID of the account the settings are for.
	AccountID *string `json:"-" validate:"required,ne="`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountSettingsOptions : Instantiate GetAccountSettingsOptions
func (*ContextBasedRestrictionsV1) NewGetAccountSettingsOptions(accountID string) *GetAccountSettingsOptions {
	return &GetAccountSettingsOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountSettingsOptions) SetAccountID(accountID string) *GetAccountSettingsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetAccountSettingsOptions) SetTransactionID(transactionID string) *GetAccountSettingsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountSettingsOptions) SetHeaders(param map[string]string) *GetAccountSettingsOptions {
	options.Headers = param
	return options
}

// GetRuleOptions : The GetRule options.
type GetRuleOptions struct {
	// The ID of a rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleOptions : Instantiate GetRuleOptions
func (*ContextBasedRestrictionsV1) NewGetRuleOptions(ruleID string) *GetRuleOptions {
	return &GetRuleOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *GetRuleOptions) SetRuleID(ruleID string) *GetRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetRuleOptions) SetTransactionID(transactionID string) *GetRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRuleOptions) SetHeaders(param map[string]string) *GetRuleOptions {
	options.Headers = param
	return options
}

// GetZoneOptions : The GetZone options.
type GetZoneOptions struct {
	// The ID of a zone.
	ZoneID *string `json:"-" validate:"required,ne="`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneOptions : Instantiate GetZoneOptions
func (*ContextBasedRestrictionsV1) NewGetZoneOptions(zoneID string) *GetZoneOptions {
	return &GetZoneOptions{
		ZoneID: core.StringPtr(zoneID),
	}
}

// SetZoneID : Allow user to set ZoneID
func (options *GetZoneOptions) SetZoneID(zoneID string) *GetZoneOptions {
	options.ZoneID = core.StringPtr(zoneID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetZoneOptions) SetTransactionID(transactionID string) *GetZoneOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneOptions) SetHeaders(param map[string]string) *GetZoneOptions {
	options.Headers = param
	return options
}

// ListAvailableServiceRefTargetsOptions : The ListAvailableServiceRefTargets options.
type ListAvailableServiceRefTargetsOptions struct {
	// Specifies the types of services to retrieve.
	Type *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListAvailableServiceRefTargetsOptions.Type property.
// Specifies the types of services to retrieve.
const (
	ListAvailableServiceRefTargetsOptionsTypeAll             = "all"
	ListAvailableServiceRefTargetsOptionsTypePlatformService = "platform_service"
)

// NewListAvailableServiceRefTargetsOptions : Instantiate ListAvailableServiceRefTargetsOptions
func (*ContextBasedRestrictionsV1) NewListAvailableServiceRefTargetsOptions() *ListAvailableServiceRefTargetsOptions {
	return &ListAvailableServiceRefTargetsOptions{}
}

// SetType : Allow user to set Type
func (options *ListAvailableServiceRefTargetsOptions) SetType(typeVar string) *ListAvailableServiceRefTargetsOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAvailableServiceRefTargetsOptions) SetHeaders(param map[string]string) *ListAvailableServiceRefTargetsOptions {
	options.Headers = param
	return options
}

// ListRulesOptions : The ListRules options.
type ListRulesOptions struct {
	// The ID of the managing account.
	AccountID *string `json:"-" validate:"required"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// The `region` resource attribute.
	Region *string `json:"-"`

	// The `resource` resource attribute.
	Resource *string `json:"-"`

	// The `resourceType` resource attribute.
	ResourceType *string `json:"-"`

	// The `serviceInstance` resource attribute.
	ServiceInstance *string `json:"-"`

	// The `serviceName` resource attribute.
	ServiceName *string `json:"-"`

	// The rule's `serviceType` resource attribute.
	ServiceType *string `json:"-"`

	// The globally unique ID of the zone.
	ZoneID *string `json:"-"`

	// Sorts results by using a valid sort field. To learn more, see
	// [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
	Sort *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRulesOptions : Instantiate ListRulesOptions
func (*ContextBasedRestrictionsV1) NewListRulesOptions(accountID string) *ListRulesOptions {
	return &ListRulesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListRulesOptions) SetAccountID(accountID string) *ListRulesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListRulesOptions) SetTransactionID(transactionID string) *ListRulesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetRegion : Allow user to set Region
func (options *ListRulesOptions) SetRegion(region string) *ListRulesOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetResource : Allow user to set Resource
func (options *ListRulesOptions) SetResource(resource string) *ListRulesOptions {
	options.Resource = core.StringPtr(resource)
	return options
}

// SetResourceType : Allow user to set ResourceType
func (options *ListRulesOptions) SetResourceType(resourceType string) *ListRulesOptions {
	options.ResourceType = core.StringPtr(resourceType)
	return options
}

// SetServiceInstance : Allow user to set ServiceInstance
func (options *ListRulesOptions) SetServiceInstance(serviceInstance string) *ListRulesOptions {
	options.ServiceInstance = core.StringPtr(serviceInstance)
	return options
}

// SetServiceName : Allow user to set ServiceName
func (options *ListRulesOptions) SetServiceName(serviceName string) *ListRulesOptions {
	options.ServiceName = core.StringPtr(serviceName)
	return options
}

// SetServiceType : Allow user to set ServiceType
func (options *ListRulesOptions) SetServiceType(serviceType string) *ListRulesOptions {
	options.ServiceType = core.StringPtr(serviceType)
	return options
}

// SetZoneID : Allow user to set ZoneID
func (options *ListRulesOptions) SetZoneID(zoneID string) *ListRulesOptions {
	options.ZoneID = core.StringPtr(zoneID)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListRulesOptions) SetSort(sort string) *ListRulesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListRulesOptions) SetHeaders(param map[string]string) *ListRulesOptions {
	options.Headers = param
	return options
}

// ListZonesOptions : The ListZones options.
type ListZonesOptions struct {
	// The ID of the managing account.
	AccountID *string `json:"-" validate:"required"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// The name of the zone.
	Name *string `json:"-"`

	// Sorts results by using a valid sort field. To learn more, see
	// [Sorting](https://cloud.ibm.com/docs/api-handbook?topic=api-handbook-sorting).
	Sort *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListZonesOptions : Instantiate ListZonesOptions
func (*ContextBasedRestrictionsV1) NewListZonesOptions(accountID string) *ListZonesOptions {
	return &ListZonesOptions{
		AccountID: core.StringPtr(accountID),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *ListZonesOptions) SetAccountID(accountID string) *ListZonesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListZonesOptions) SetTransactionID(transactionID string) *ListZonesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetName : Allow user to set Name
func (options *ListZonesOptions) SetName(name string) *ListZonesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetSort : Allow user to set Sort
func (options *ListZonesOptions) SetSort(sort string) *ListZonesOptions {
	options.Sort = core.StringPtr(sort)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListZonesOptions) SetHeaders(param map[string]string) *ListZonesOptions {
	options.Headers = param
	return options
}

// OutAccountSettings : An output account settings.
type OutAccountSettings struct {
	// The globally unique ID of the account settings.
	ID *string `json:"id" validate:"required"`

	// The account settings CRN.
	Crn *string `json:"crn" validate:"required"`

	// the max number of rules allowed for the account.
	RuleCountLimit *int64 `json:"rule_count_limit" validate:"required"`

	// the max number of zones allowed for the account.
	ZoneCountLimit *int64 `json:"zone_count_limit" validate:"required"`

	// the current number of rules used by the account.
	CurrentRuleCount *int64 `json:"current_rule_count" validate:"required"`

	// the current number of zones used by the account.
	CurrentZoneCount *int64 `json:"current_zone_count" validate:"required"`

	// The href link to the resource.
	Href *string `json:"href" validate:"required"`

	// The time the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// IAM ID of the user or service which created the resource.
	CreatedByID *string `json:"created_by_id" validate:"required"`

	// The last time the resource was modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at" validate:"required"`

	// IAM ID of the user or service which modified the resource.
	LastModifiedByID *string `json:"last_modified_by_id" validate:"required"`
}

// UnmarshalOutAccountSettings unmarshals an instance of OutAccountSettings from the specified map of raw messages.
func UnmarshalOutAccountSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutAccountSettings)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_count_limit", &obj.RuleCountLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone_count_limit", &obj.ZoneCountLimit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current_rule_count", &obj.CurrentRuleCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current_zone_count", &obj.CurrentZoneCount)
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

// OutRule : An output rule.
type OutRule struct {
	// The globally unique ID of the rule.
	ID *string `json:"id" validate:"required"`

	// The rule CRN.
	Crn *string `json:"crn" validate:"required"`

	// The description of the rule.
	Description *string `json:"description,omitempty"`

	// The contexts this rule applies to.
	Contexts []RuleContext `json:"contexts" validate:"required"`

	// The resources this rule apply to.
	Resources []Resource `json:"resources" validate:"required"`

	// The href link to the resource.
	Href *string `json:"href" validate:"required"`

	// The time the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// IAM ID of the user or service which created the resource.
	CreatedByID *string `json:"created_by_id" validate:"required"`

	// The last time the resource was modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at" validate:"required"`

	// IAM ID of the user or service which modified the resource.
	LastModifiedByID *string `json:"last_modified_by_id" validate:"required"`
}

// UnmarshalOutRule unmarshals an instance of OutRule from the specified map of raw messages.
func UnmarshalOutRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutRule)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "contexts", &obj.Contexts, UnmarshalRuleContext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
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

// OutRulePage : The response object of the ListRules operation.
type OutRulePage struct {
	// The number of returned results.
	Count *int64 `json:"count" validate:"required"`

	// The returned rules.
	Rules []OutRule `json:"rules" validate:"required"`
}

// UnmarshalOutRulePage unmarshals an instance of OutRulePage from the specified map of raw messages.
func UnmarshalOutRulePage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutRulePage)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalOutRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OutZone : An output zone.
type OutZone struct {
	// The globally unique ID of the zone.
	ID *string `json:"id" validate:"required"`

	// The zone CRN.
	Crn *string `json:"crn" validate:"required"`

	// The name of the zone.
	Name *string `json:"name" validate:"required"`

	// The id of the account owning this zone.
	AccountID *string `json:"account_id" validate:"required"`

	// The description of the zone.
	Description *string `json:"description,omitempty"`

	// The list of addresses in the zone.
	Addresses []AddressIntf `json:"addresses" validate:"required"`

	// The list of excluded addresses in the zone.
	Excluded []AddressIntf `json:"excluded,omitempty"`

	// The href link to the resource.
	Href *string `json:"href" validate:"required"`

	// The time the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// IAM ID of the user or service which created the resource.
	CreatedByID *string `json:"created_by_id" validate:"required"`

	// The last time the resource was modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at" validate:"required"`

	// IAM ID of the user or service which modified the resource.
	LastModifiedByID *string `json:"last_modified_by_id" validate:"required"`
}

// UnmarshalOutZone unmarshals an instance of OutZone from the specified map of raw messages.
func UnmarshalOutZone(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutZone)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "addresses", &obj.Addresses, UnmarshalAddress)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded", &obj.Excluded, UnmarshalAddress)
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

// OutZonePage : The response object of the ListZones operation.
type OutZonePage struct {
	// The number of returned results.
	Count *int64 `json:"count" validate:"required"`

	// The returned zones.
	Zones []OutZoneSummary `json:"zones" validate:"required"`
}

// UnmarshalOutZonePage unmarshals an instance of OutZonePage from the specified map of raw messages.
func UnmarshalOutZonePage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutZonePage)
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "zones", &obj.Zones, UnmarshalOutZoneSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OutZoneSummary : An output zone summary.
type OutZoneSummary struct {
	// The globally unique ID of the zone.
	ID *string `json:"id" validate:"required"`

	// The zone CRN.
	Crn *string `json:"crn" validate:"required"`

	// The name of the zone.
	Name *string `json:"name" validate:"required"`

	// The description of the zone.
	Description *string `json:"description,omitempty"`

	// A preview of addresses in the zone (3 addresses maximum).
	AddressesPreview []AddressIntf `json:"addresses_preview" validate:"required"`

	// The number of addresses in the zone.
	AddressCount *int64 `json:"address_count" validate:"required"`

	// The number of excluded addresses in the zone.
	ExcludedCount *int64 `json:"excluded_count" validate:"required"`

	// The href link to the resource.
	Href *string `json:"href" validate:"required"`

	// The time the resource was created.
	CreatedAt *strfmt.DateTime `json:"created_at" validate:"required"`

	// IAM ID of the user or service which created the resource.
	CreatedByID *string `json:"created_by_id" validate:"required"`

	// The last time the resource was modified.
	LastModifiedAt *strfmt.DateTime `json:"last_modified_at" validate:"required"`

	// IAM ID of the user or service which modified the resource.
	LastModifiedByID *string `json:"last_modified_by_id" validate:"required"`
}

// UnmarshalOutZoneSummary unmarshals an instance of OutZoneSummary from the specified map of raw messages.
func UnmarshalOutZoneSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OutZoneSummary)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
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
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "addresses_preview", &obj.AddressesPreview, UnmarshalAddress)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address_count", &obj.AddressCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "excluded_count", &obj.ExcludedCount)
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

// ReplaceRuleOptions : The ReplaceRule options.
type ReplaceRuleOptions struct {
	// The ID of a rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// The current revision of the resource being updated. This can be found in the Create/Get/Update resource response
	// ETag header.
	IfMatch *string `json:"-" validate:"required"`

	// The description of the rule.
	Description *string `json:"description,omitempty"`

	// The contexts this rule applies to.
	Contexts []RuleContext `json:"contexts,omitempty"`

	// The resources this rule apply to.
	Resources []Resource `json:"resources,omitempty"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceRuleOptions : Instantiate ReplaceRuleOptions
func (*ContextBasedRestrictionsV1) NewReplaceRuleOptions(ruleID string, ifMatch string) *ReplaceRuleOptions {
	return &ReplaceRuleOptions{
		RuleID:  core.StringPtr(ruleID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *ReplaceRuleOptions) SetRuleID(ruleID string) *ReplaceRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *ReplaceRuleOptions) SetIfMatch(ifMatch string) *ReplaceRuleOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetDescription : Allow user to set Description
func (options *ReplaceRuleOptions) SetDescription(description string) *ReplaceRuleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetContexts : Allow user to set Contexts
func (options *ReplaceRuleOptions) SetContexts(contexts []RuleContext) *ReplaceRuleOptions {
	options.Contexts = contexts
	return options
}

// SetResources : Allow user to set Resources
func (options *ReplaceRuleOptions) SetResources(resources []Resource) *ReplaceRuleOptions {
	options.Resources = resources
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ReplaceRuleOptions) SetTransactionID(transactionID string) *ReplaceRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceRuleOptions) SetHeaders(param map[string]string) *ReplaceRuleOptions {
	options.Headers = param
	return options
}

// ReplaceZoneOptions : The ReplaceZone options.
type ReplaceZoneOptions struct {
	// The ID of a zone.
	ZoneID *string `json:"-" validate:"required,ne="`

	// The current revision of the resource being updated. This can be found in the Create/Get/Update resource response
	// ETag header.
	IfMatch *string `json:"-" validate:"required"`

	// The name of the zone.
	Name *string `json:"name,omitempty"`

	// The id of the account owning this zone.
	AccountID *string `json:"account_id,omitempty"`

	// The description of the zone.
	Description *string `json:"description,omitempty"`

	// The list of addresses in the zone.
	Addresses []AddressIntf `json:"addresses,omitempty"`

	// The list of excluded addresses in the zone.
	Excluded []AddressIntf `json:"excluded,omitempty"`

	// The UUID that is used to correlate and track transactions. If you omit this field, the service generates and sends a
	// transaction ID in the response.
	// **Note:** To help with debugging, we strongly recommend that you generate and supply a `Transaction-Id` with each
	// request.
	TransactionID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceZoneOptions : Instantiate ReplaceZoneOptions
func (*ContextBasedRestrictionsV1) NewReplaceZoneOptions(zoneID string, ifMatch string) *ReplaceZoneOptions {
	return &ReplaceZoneOptions{
		ZoneID:  core.StringPtr(zoneID),
		IfMatch: core.StringPtr(ifMatch),
	}
}

// SetZoneID : Allow user to set ZoneID
func (options *ReplaceZoneOptions) SetZoneID(zoneID string) *ReplaceZoneOptions {
	options.ZoneID = core.StringPtr(zoneID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *ReplaceZoneOptions) SetIfMatch(ifMatch string) *ReplaceZoneOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *ReplaceZoneOptions) SetName(name string) *ReplaceZoneOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *ReplaceZoneOptions) SetAccountID(accountID string) *ReplaceZoneOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetDescription : Allow user to set Description
func (options *ReplaceZoneOptions) SetDescription(description string) *ReplaceZoneOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetAddresses : Allow user to set Addresses
func (options *ReplaceZoneOptions) SetAddresses(addresses []AddressIntf) *ReplaceZoneOptions {
	options.Addresses = addresses
	return options
}

// SetExcluded : Allow user to set Excluded
func (options *ReplaceZoneOptions) SetExcluded(excluded []AddressIntf) *ReplaceZoneOptions {
	options.Excluded = excluded
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ReplaceZoneOptions) SetTransactionID(transactionID string) *ReplaceZoneOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceZoneOptions) SetHeaders(param map[string]string) *ReplaceZoneOptions {
	options.Headers = param
	return options
}

// Resource : An rule resource.
type Resource struct {
	// The resource attributes.
	Attributes []ResourceAttribute `json:"attributes" validate:"required"`

	// The optional resource tags.
	Tags []ResourceTagAttribute `json:"tags,omitempty"`
}

// NewResource : Instantiate Resource (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewResource(attributes []ResourceAttribute) (_model *Resource, err error) {
	_model = &Resource{
		Attributes: attributes,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalResourceAttribute)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "tags", &obj.Tags, UnmarshalResourceTagAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceAttribute : A rule resource attribute.
type ResourceAttribute struct {
	// The attribute name.
	Name *string `json:"name" validate:"required"`

	// The attribute value.
	Value *string `json:"value" validate:"required"`

	// The attribute operator.
	Operator *string `json:"operator,omitempty"`
}

// NewResourceAttribute : Instantiate ResourceAttribute (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewResourceAttribute(name string, value string) (_model *ResourceAttribute, err error) {
	_model = &ResourceAttribute{
		Name:  core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResourceAttribute unmarshals an instance of ResourceAttribute from the specified map of raw messages.
func UnmarshalResourceAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceTagAttribute : A rule resource tag attribute.
type ResourceTagAttribute struct {
	// The tag attribute name.
	Name *string `json:"name" validate:"required"`

	// The tag attribute value.
	Value *string `json:"value" validate:"required"`

	// The attribute operator.
	Operator *string `json:"operator,omitempty"`
}

// NewResourceTagAttribute : Instantiate ResourceTagAttribute (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewResourceTagAttribute(name string, value string) (_model *ResourceTagAttribute, err error) {
	_model = &ResourceTagAttribute{
		Name:  core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalResourceTagAttribute unmarshals an instance of ResourceTagAttribute from the specified map of raw messages.
func UnmarshalResourceTagAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceTagAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "operator", &obj.Operator)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleContext : A rule context.
type RuleContext struct {
	// The attributes.
	Attributes []RuleContextAttribute `json:"attributes" validate:"required"`
}

// NewRuleContext : Instantiate RuleContext (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewRuleContext(attributes []RuleContextAttribute) (_model *RuleContext, err error) {
	_model = &RuleContext{
		Attributes: attributes,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuleContext unmarshals an instance of RuleContext from the specified map of raw messages.
func UnmarshalRuleContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleContext)
	err = core.UnmarshalModel(m, "attributes", &obj.Attributes, UnmarshalRuleContextAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleContextAttribute : An rule context attribute.
type RuleContextAttribute struct {
	// The attribute name.
	Name *string `json:"name" validate:"required"`

	// The attribute value.
	Value *string `json:"value" validate:"required"`
}

// NewRuleContextAttribute : Instantiate RuleContextAttribute (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewRuleContextAttribute(name string, value string) (_model *RuleContextAttribute, err error) {
	_model = &RuleContextAttribute{
		Name:  core.StringPtr(name),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalRuleContextAttribute unmarshals an instance of RuleContextAttribute from the specified map of raw messages.
func UnmarshalRuleContextAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleContextAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// ServiceRefTarget : Summary information about a service reference target.
type ServiceRefTarget struct {
	// The name of the service.
	ServiceName *string `json:"service_name,omitempty"`

	// The type of the service.
	ServiceType *string `json:"service_type,omitempty"`
}

// UnmarshalServiceRefTarget unmarshals an instance of ServiceRefTarget from the specified map of raw messages.
func UnmarshalServiceRefTarget(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceRefTarget)
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_type", &obj.ServiceType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceRefTargetPage : A page of service reference targets.
type ServiceRefTargetPage struct {
	// The list of service reference targets.
	Targets []ServiceRefTarget `json:"targets,omitempty"`
}

// UnmarshalServiceRefTargetPage unmarshals an instance of ServiceRefTargetPage from the specified map of raw messages.
func UnmarshalServiceRefTargetPage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceRefTargetPage)
	err = core.UnmarshalModel(m, "targets", &obj.Targets, UnmarshalServiceRefTarget)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ServiceRefValue : A service reference value.
type ServiceRefValue struct {
	// The id of the account owning the service.
	AccountID *string `json:"account_id,omitempty"`

	// The service type.
	ServiceType *string `json:"service_type,omitempty"`

	// The service name.
	ServiceName *string `json:"service_name,omitempty"`

	// The service instance.
	ServiceInstance *string `json:"service_instance,omitempty"`
}

// UnmarshalServiceRefValue unmarshals an instance of ServiceRefValue from the specified map of raw messages.
func UnmarshalServiceRefValue(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ServiceRefValue)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_type", &obj.ServiceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service_instance", &obj.ServiceInstance)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddressIPAddress : A single IP address.
// This model "extends" Address
type AddressIPAddress struct {
	// The type of address.
	Type *string `json:"type" validate:"required"`

	// The IP address.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AddressIPAddress.Type property.
// The type of address.
const (
	AddressIPAddress_Type_Ipaddress = "ipAddress"
)

// NewAddressIPAddress : Instantiate AddressIPAddress (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewAddressIPAddress(typeVar string, value string) (_model *AddressIPAddress, err error) {
	_model = &AddressIPAddress{
		Type:  core.StringPtr(typeVar),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AddressIPAddress) isaAddress() bool {
	return true
}

// UnmarshalAddressIPAddress unmarshals an instance of AddressIPAddress from the specified map of raw messages.
func UnmarshalAddressIPAddress(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddressIPAddress)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// AddressIPAddressRange : An IP address range.
// This model "extends" Address
type AddressIPAddressRange struct {
	// The type of address.
	Type *string `json:"type" validate:"required"`

	// The ip range in <first-ip>-<last-ip> format.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AddressIPAddressRange.Type property.
// The type of address.
const (
	AddressIPAddressRange_Type_Iprange = "ipRange"
)

// NewAddressIPAddressRange : Instantiate AddressIPAddressRange (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewAddressIPAddressRange(typeVar string, value string) (_model *AddressIPAddressRange, err error) {
	_model = &AddressIPAddressRange{
		Type:  core.StringPtr(typeVar),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AddressIPAddressRange) isaAddress() bool {
	return true
}

// UnmarshalAddressIPAddressRange unmarshals an instance of AddressIPAddressRange from the specified map of raw messages.
func UnmarshalAddressIPAddressRange(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddressIPAddressRange)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// AddressServiceRef : A service reference.
// This model "extends" Address
type AddressServiceRef struct {
	// The type of address.
	Type *string `json:"type" validate:"required"`

	// A service reference value.
	Ref *ServiceRefValue `json:"ref" validate:"required"`
}

// Constants associated with the AddressServiceRef.Type property.
// The type of address.
const (
	AddressServiceRef_Type_ServiceRef = "serviceRef"
)

// NewAddressServiceRef : Instantiate AddressServiceRef (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewAddressServiceRef(typeVar string, ref *ServiceRefValue) (_model *AddressServiceRef, err error) {
	_model = &AddressServiceRef{
		Type: core.StringPtr(typeVar),
		Ref:  ref,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AddressServiceRef) isaAddress() bool {
	return true
}

// UnmarshalAddressServiceRef unmarshals an instance of AddressServiceRef from the specified map of raw messages.
func UnmarshalAddressServiceRef(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddressServiceRef)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "ref", &obj.Ref, UnmarshalServiceRefValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AddressSubnet : A subnet in CIDR format.
// This model "extends" Address
type AddressSubnet struct {
	// The type of address.
	Type *string `json:"type" validate:"required"`

	// The subnet in CIDR format.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AddressSubnet.Type property.
// The type of address.
const (
	AddressSubnet_Type_Subnet = "subnet"
)

// NewAddressSubnet : Instantiate AddressSubnet (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewAddressSubnet(typeVar string, value string) (_model *AddressSubnet, err error) {
	_model = &AddressSubnet{
		Type:  core.StringPtr(typeVar),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AddressSubnet) isaAddress() bool {
	return true
}

// UnmarshalAddressSubnet unmarshals an instance of AddressSubnet from the specified map of raw messages.
func UnmarshalAddressSubnet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddressSubnet)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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

// AddressVPC : A single VPC address.
// This model "extends" Address
type AddressVPC struct {
	// The type of address.
	Type *string `json:"type" validate:"required"`

	// The VPC CRN.
	Value *string `json:"value" validate:"required"`
}

// Constants associated with the AddressVPC.Type property.
// The type of address.
const (
	AddressVPC_Type_Vpc = "vpc"
)

// NewAddressVPC : Instantiate AddressVPC (Generic Model Constructor)
func (*ContextBasedRestrictionsV1) NewAddressVPC(typeVar string, value string) (_model *AddressVPC, err error) {
	_model = &AddressVPC{
		Type:  core.StringPtr(typeVar),
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*AddressVPC) isaAddress() bool {
	return true
}

// UnmarshalAddressVPC unmarshals an instance of AddressVPC from the specified map of raw messages.
func UnmarshalAddressVPC(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AddressVPC)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
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
