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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-c6db7f4a-20210114-141015
 */

// Package configurationgovernancev1 : Operations and models for the ConfigurationGovernanceV1 service
package configurationgovernancev1

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

// ConfigurationGovernanceV1 : API specification for the Configuration Governance service.
//
// Version: 1.0.0
type ConfigurationGovernanceV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://compliance.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "configuration_governance"

// ConfigurationGovernanceV1Options : Service options
type ConfigurationGovernanceV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewConfigurationGovernanceV1UsingExternalConfig : constructs an instance of ConfigurationGovernanceV1 with passed in options and external configuration.
func NewConfigurationGovernanceV1UsingExternalConfig(options *ConfigurationGovernanceV1Options) (configurationGovernance *ConfigurationGovernanceV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	configurationGovernance, err = NewConfigurationGovernanceV1(options)
	if err != nil {
		return
	}

	err = configurationGovernance.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = configurationGovernance.Service.SetServiceURL(options.URL)
	}
	return
}

// NewConfigurationGovernanceV1 : constructs an instance of ConfigurationGovernanceV1 with passed in options.
func NewConfigurationGovernanceV1(options *ConfigurationGovernanceV1Options) (service *ConfigurationGovernanceV1, err error) {
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

	service = &ConfigurationGovernanceV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "configurationGovernance" suitable for processing requests.
func (configurationGovernance *ConfigurationGovernanceV1) Clone() *ConfigurationGovernanceV1 {
	if core.IsNil(configurationGovernance) {
		return nil
	}
	clone := *configurationGovernance
	clone.Service = configurationGovernance.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (configurationGovernance *ConfigurationGovernanceV1) SetServiceURL(url string) error {
	return configurationGovernance.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (configurationGovernance *ConfigurationGovernanceV1) GetServiceURL() string {
	return configurationGovernance.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (configurationGovernance *ConfigurationGovernanceV1) SetDefaultHeaders(headers http.Header) {
	configurationGovernance.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (configurationGovernance *ConfigurationGovernanceV1) SetEnableGzipCompression(enableGzip bool) {
	configurationGovernance.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (configurationGovernance *ConfigurationGovernanceV1) GetEnableGzipCompression() bool {
	return configurationGovernance.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (configurationGovernance *ConfigurationGovernanceV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	configurationGovernance.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (configurationGovernance *ConfigurationGovernanceV1) DisableRetries() {
	configurationGovernance.Service.DisableRetries()
}

// CreateRules : Create rules
// Creates one or more rules that you can use to govern the way that IBM Cloud resources can be provisioned and
// configured.
//
// A successful `POST /config/rules` request defines a rule based on the target, conditions, and enforcement actions
// that you specify. The response returns the ID value for your rule, along with other metadata.
func (configurationGovernance *ConfigurationGovernanceV1) CreateRules(createRulesOptions *CreateRulesOptions) (result *CreateRulesResponse, response *core.DetailedResponse, err error) {
	return configurationGovernance.CreateRulesWithContext(context.Background(), createRulesOptions)
}

// CreateRulesWithContext is an alternate form of the CreateRules method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) CreateRulesWithContext(ctx context.Context, createRulesOptions *CreateRulesOptions) (result *CreateRulesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRulesOptions, "createRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRulesOptions, "createRulesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "CreateRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createRulesOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createRulesOptions.Rules != nil {
		body["rules"] = createRulesOptions.Rules
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
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateRulesResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListRules : List rules
// Retrieves a list of the rules that are available in your account.
func (configurationGovernance *ConfigurationGovernanceV1) ListRules(listRulesOptions *ListRulesOptions) (result *RuleList, response *core.DetailedResponse, err error) {
	return configurationGovernance.ListRulesWithContext(context.Background(), listRulesOptions)
}

// ListRulesWithContext is an alternate form of the ListRules method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) ListRulesWithContext(ctx context.Context, listRulesOptions *ListRulesOptions) (result *RuleList, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "ListRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listRulesOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listRulesOptions.TransactionID))
	}

	builder.AddQuery("account_id", fmt.Sprint(*listRulesOptions.AccountID))
	if listRulesOptions.Attached != nil {
		builder.AddQuery("attached", fmt.Sprint(*listRulesOptions.Attached))
	}
	if listRulesOptions.Labels != nil {
		builder.AddQuery("labels", fmt.Sprint(*listRulesOptions.Labels))
	}
	if listRulesOptions.Scopes != nil {
		builder.AddQuery("scopes", fmt.Sprint(*listRulesOptions.Scopes))
	}
	if listRulesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listRulesOptions.Limit))
	}
	if listRulesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listRulesOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRuleList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRule : Get a rule
// Retrieves an existing rule and its details.
func (configurationGovernance *ConfigurationGovernanceV1) GetRule(getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configurationGovernance.GetRuleWithContext(context.Background(), getRuleOptions)
}

// GetRuleWithContext is an alternate form of the GetRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) GetRuleWithContext(ctx context.Context, getRuleOptions *GetRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "GetRule")
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
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateRule : Update a rule
// Updates an existing rule based on the properties that you specify.
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRule(updateRuleOptions *UpdateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	return configurationGovernance.UpdateRuleWithContext(context.Background(), updateRuleOptions)
}

// UpdateRuleWithContext is an alternate form of the UpdateRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) UpdateRuleWithContext(ctx context.Context, updateRuleOptions *UpdateRuleOptions) (result *Rule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRuleOptions, "updateRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRuleOptions, "updateRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *updateRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "UpdateRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateRuleOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateRuleOptions.IfMatch))
	}
	if updateRuleOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateRuleOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateRuleOptions.Name != nil {
		body["name"] = updateRuleOptions.Name
	}
	if updateRuleOptions.Description != nil {
		body["description"] = updateRuleOptions.Description
	}
	if updateRuleOptions.Target != nil {
		body["target"] = updateRuleOptions.Target
	}
	if updateRuleOptions.RequiredConfig != nil {
		body["required_config"] = updateRuleOptions.RequiredConfig
	}
	if updateRuleOptions.EnforcementActions != nil {
		body["enforcement_actions"] = updateRuleOptions.EnforcementActions
	}
	if updateRuleOptions.AccountID != nil {
		body["account_id"] = updateRuleOptions.AccountID
	}
	if updateRuleOptions.RuleType != nil {
		body["rule_type"] = updateRuleOptions.RuleType
	}
	if updateRuleOptions.Labels != nil {
		body["labels"] = updateRuleOptions.Labels
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
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRule)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteRule : Delete a rule
// Deletes an existing rule.
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRule(deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
	return configurationGovernance.DeleteRuleWithContext(context.Background(), deleteRuleOptions)
}

// DeleteRuleWithContext is an alternate form of the DeleteRule method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) DeleteRuleWithContext(ctx context.Context, deleteRuleOptions *DeleteRuleOptions) (response *core.DetailedResponse, err error) {
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
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "DeleteRule")
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

	response, err = configurationGovernance.Service.Request(request, nil)

	return
}

// CreateAttachments : Create attachments
// Creates one or more scope attachments for an existing rule.
//
// You can attach an existing rule to a scope, such as a specific IBM Cloud account, to start evaluating the rule for
// compliance. A successful
// `POST /config/v1/rules/{rule_id}/attachments` returns the ID value for the attachment, along with other metadata.
func (configurationGovernance *ConfigurationGovernanceV1) CreateAttachments(createAttachmentsOptions *CreateAttachmentsOptions) (result *CreateAttachmentsResponse, response *core.DetailedResponse, err error) {
	return configurationGovernance.CreateAttachmentsWithContext(context.Background(), createAttachmentsOptions)
}

// CreateAttachmentsWithContext is an alternate form of the CreateAttachments method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) CreateAttachmentsWithContext(ctx context.Context, createAttachmentsOptions *CreateAttachmentsOptions) (result *CreateAttachmentsResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createAttachmentsOptions, "createAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createAttachmentsOptions, "createAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *createAttachmentsOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "CreateAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createAttachmentsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*createAttachmentsOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if createAttachmentsOptions.Attachments != nil {
		body["attachments"] = createAttachmentsOptions.Attachments
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
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateAttachmentsResponse)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListAttachments : List attachments
// Retrieves a list of scope attachments that are associated with the specified rule.
func (configurationGovernance *ConfigurationGovernanceV1) ListAttachments(listAttachmentsOptions *ListAttachmentsOptions) (result *AttachmentList, response *core.DetailedResponse, err error) {
	return configurationGovernance.ListAttachmentsWithContext(context.Background(), listAttachmentsOptions)
}

// ListAttachmentsWithContext is an alternate form of the ListAttachments method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) ListAttachmentsWithContext(ctx context.Context, listAttachmentsOptions *ListAttachmentsOptions) (result *AttachmentList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listAttachmentsOptions, "listAttachmentsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listAttachmentsOptions, "listAttachmentsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id": *listAttachmentsOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAttachmentsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "ListAttachments")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listAttachmentsOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*listAttachmentsOptions.TransactionID))
	}

	if listAttachmentsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listAttachmentsOptions.Limit))
	}
	if listAttachmentsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listAttachmentsOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachmentList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAttachment : Get an attachment
// Retrieves an existing scope attachment for a rule.
func (configurationGovernance *ConfigurationGovernanceV1) GetAttachment(getAttachmentOptions *GetAttachmentOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	return configurationGovernance.GetAttachmentWithContext(context.Background(), getAttachmentOptions)
}

// GetAttachmentWithContext is an alternate form of the GetAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) GetAttachmentWithContext(ctx context.Context, getAttachmentOptions *GetAttachmentOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAttachmentOptions, "getAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAttachmentOptions, "getAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *getAttachmentOptions.RuleID,
		"attachment_id": *getAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "GetAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*getAttachmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateAttachment : Update an attachment
// Updates an existing scope attachment based on the properties that you specify.
func (configurationGovernance *ConfigurationGovernanceV1) UpdateAttachment(updateAttachmentOptions *UpdateAttachmentOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	return configurationGovernance.UpdateAttachmentWithContext(context.Background(), updateAttachmentOptions)
}

// UpdateAttachmentWithContext is an alternate form of the UpdateAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) UpdateAttachmentWithContext(ctx context.Context, updateAttachmentOptions *UpdateAttachmentOptions) (result *Attachment, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateAttachmentOptions, "updateAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateAttachmentOptions, "updateAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *updateAttachmentOptions.RuleID,
		"attachment_id": *updateAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "UpdateAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateAttachmentOptions.IfMatch != nil {
		builder.AddHeader("If-Match", fmt.Sprint(*updateAttachmentOptions.IfMatch))
	}
	if updateAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*updateAttachmentOptions.TransactionID))
	}

	body := make(map[string]interface{})
	if updateAttachmentOptions.AccountID != nil {
		body["account_id"] = updateAttachmentOptions.AccountID
	}
	if updateAttachmentOptions.IncludedScope != nil {
		body["included_scope"] = updateAttachmentOptions.IncludedScope
	}
	if updateAttachmentOptions.ExcludedScopes != nil {
		body["excluded_scopes"] = updateAttachmentOptions.ExcludedScopes
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
	response, err = configurationGovernance.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAttachment)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteAttachment : Delete an attachment
// Deletes an existing scope attachment.
func (configurationGovernance *ConfigurationGovernanceV1) DeleteAttachment(deleteAttachmentOptions *DeleteAttachmentOptions) (response *core.DetailedResponse, err error) {
	return configurationGovernance.DeleteAttachmentWithContext(context.Background(), deleteAttachmentOptions)
}

// DeleteAttachmentWithContext is an alternate form of the DeleteAttachment method which supports a Context parameter
func (configurationGovernance *ConfigurationGovernanceV1) DeleteAttachmentWithContext(ctx context.Context, deleteAttachmentOptions *DeleteAttachmentOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteAttachmentOptions, "deleteAttachmentOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteAttachmentOptions, "deleteAttachmentOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"rule_id":       *deleteAttachmentOptions.RuleID,
		"attachment_id": *deleteAttachmentOptions.AttachmentID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = configurationGovernance.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(configurationGovernance.Service.Options.URL, `/config/v1/rules/{rule_id}/attachments/{attachment_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteAttachmentOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("configuration_governance", "V1", "DeleteAttachment")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteAttachmentOptions.TransactionID != nil {
		builder.AddHeader("Transaction-Id", fmt.Sprint(*deleteAttachmentOptions.TransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = configurationGovernance.Service.Request(request, nil)

	return
}

// Attachment : The scopes to attach to the rule.
type Attachment struct {
	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `json:"attachment_id" validate:"required"`

	// The UUID that uniquely identifies the rule.
	RuleID *string `json:"rule_id" validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `json:"included_scope" validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope `json:"excluded_scopes,omitempty"`
}

// UnmarshalAttachment unmarshals an instance of Attachment from the specified map of raw messages.
func UnmarshalAttachment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Attachment)
	err = core.UnmarshalPrimitive(m, "attachment_id", &obj.AttachmentID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalRuleScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded_scopes", &obj.ExcludedScopes, UnmarshalRuleScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentList : A list of attachments.
type AttachmentList struct {
	// The requested offset for the returned items.
	Offset *int64 `json:"offset" validate:"required"`

	// The requested limit for the returned items.
	Limit *int64 `json:"limit" validate:"required"`

	// The total number of available items.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The first page of available items.
	First *Link `json:"first" validate:"required"`

	// The last page of available items.
	Last *Link `json:"last" validate:"required"`

	Attachments []Attachment `json:"attachments" validate:"required"`
}

// UnmarshalAttachmentList unmarshals an instance of AttachmentList from the specified map of raw messages.
func UnmarshalAttachmentList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentList)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AttachmentRequest : The scopes to attach to the rule.
type AttachmentRequest struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id" validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `json:"included_scope" validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope `json:"excluded_scopes,omitempty"`
}

// NewAttachmentRequest : Instantiate AttachmentRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewAttachmentRequest(accountID string, includedScope *RuleScope) (model *AttachmentRequest, err error) {
	model = &AttachmentRequest{
		AccountID:     core.StringPtr(accountID),
		IncludedScope: includedScope,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalAttachmentRequest unmarshals an instance of AttachmentRequest from the specified map of raw messages.
func UnmarshalAttachmentRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AttachmentRequest)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "included_scope", &obj.IncludedScope, UnmarshalRuleScope)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "excluded_scopes", &obj.ExcludedScopes, UnmarshalRuleScope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateAttachmentsOptions : The CreateAttachments options.
type CreateAttachmentsOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	Attachments []AttachmentRequest `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateAttachmentsOptions : Instantiate CreateAttachmentsOptions
func (*ConfigurationGovernanceV1) NewCreateAttachmentsOptions(ruleID string, attachments []AttachmentRequest) *CreateAttachmentsOptions {
	return &CreateAttachmentsOptions{
		RuleID:      core.StringPtr(ruleID),
		Attachments: attachments,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *CreateAttachmentsOptions) SetRuleID(ruleID string) *CreateAttachmentsOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachments : Allow user to set Attachments
func (options *CreateAttachmentsOptions) SetAttachments(attachments []AttachmentRequest) *CreateAttachmentsOptions {
	options.Attachments = attachments
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateAttachmentsOptions) SetTransactionID(transactionID string) *CreateAttachmentsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateAttachmentsOptions) SetHeaders(param map[string]string) *CreateAttachmentsOptions {
	options.Headers = param
	return options
}

// CreateAttachmentsResponse : CreateAttachmentsResponse struct
type CreateAttachmentsResponse struct {
	Attachments []Attachment `json:"attachments" validate:"required"`
}

// UnmarshalCreateAttachmentsResponse unmarshals an instance of CreateAttachmentsResponse from the specified map of raw messages.
func UnmarshalCreateAttachmentsResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateAttachmentsResponse)
	err = core.UnmarshalModel(m, "attachments", &obj.Attachments, UnmarshalAttachment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleRequest : A rule to be created.
type CreateRuleRequest struct {
	// A field that you can use in bulk operations to store a custom identifier for an individual request. If you omit this
	// field, the service generates and sends a `request_id` string for each new rule. The generated string corresponds
	// with the numerical order of the rules request array. For example, `"request_id": "1"`, `"request_id": "2"`.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `request_id` with
	// each request.
	RequestID *string `json:"request_id,omitempty"`

	// User-settable properties associated with a rule to be created or updated.
	Rule *RuleRequest `json:"rule" validate:"required"`
}

// NewCreateRuleRequest : Instantiate CreateRuleRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewCreateRuleRequest(rule *RuleRequest) (model *CreateRuleRequest, err error) {
	model = &CreateRuleRequest{
		Rule: rule,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateRuleRequest unmarshals an instance of CreateRuleRequest from the specified map of raw messages.
func UnmarshalCreateRuleRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRuleRequest)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rule", &obj.Rule, UnmarshalRuleRequest)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRuleResponse : Response information for a rule request.
//
// If the 'status_code' property indicates success, the 'request_id' and 'rule' properties will be present.   If the
// 'status_code' property indicates an error, the 'request_id', 'errors', and 'trace' fields will be present.
type CreateRuleResponse struct {
	// The identifier that is used to correlate an individual request.
	//
	// To assist with debugging, you can use this ID to identify and inspect only one request that was made as part of a
	// bulk operation.
	RequestID *string `json:"request_id,omitempty"`

	// The HTTP response status code.
	StatusCode *int64 `json:"status_code,omitempty"`

	// Information about a newly-created rule.
	//
	// This field will be present for a successful request.
	Rule *Rule `json:"rule,omitempty"`

	// The error contents of the multi-status response.
	//
	// This field will be present for a failed rule request.
	Errors []RuleResponseError `json:"errors,omitempty"`

	// The UUID that uniquely identifies the request.
	//
	// This field will be present for a failed rule request.
	Trace *string `json:"trace,omitempty"`
}

// UnmarshalCreateRuleResponse unmarshals an instance of CreateRuleResponse from the specified map of raw messages.
func UnmarshalCreateRuleResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRuleResponse)
	err = core.UnmarshalPrimitive(m, "request_id", &obj.RequestID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status_code", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rule", &obj.Rule, UnmarshalRule)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalRuleResponseError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "trace", &obj.Trace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRulesOptions : The CreateRules options.
type CreateRulesOptions struct {
	// A list of rules to be created.
	Rules []CreateRuleRequest `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRulesOptions : Instantiate CreateRulesOptions
func (*ConfigurationGovernanceV1) NewCreateRulesOptions(rules []CreateRuleRequest) *CreateRulesOptions {
	return &CreateRulesOptions{
		Rules: rules,
	}
}

// SetRules : Allow user to set Rules
func (options *CreateRulesOptions) SetRules(rules []CreateRuleRequest) *CreateRulesOptions {
	options.Rules = rules
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *CreateRulesOptions) SetTransactionID(transactionID string) *CreateRulesOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRulesOptions) SetHeaders(param map[string]string) *CreateRulesOptions {
	options.Headers = param
	return options
}

// CreateRulesResponse : The response associated with a request to create one or more rules.
type CreateRulesResponse struct {
	// An array of rule responses.
	Rules []CreateRuleResponse `json:"rules" validate:"required"`
}

// UnmarshalCreateRulesResponse unmarshals an instance of CreateRulesResponse from the specified map of raw messages.
func UnmarshalCreateRulesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRulesResponse)
	err = core.UnmarshalModel(m, "rules", &obj.Rules, UnmarshalCreateRuleResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteAttachmentOptions : The DeleteAttachment options.
type DeleteAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteAttachmentOptions : Instantiate DeleteAttachmentOptions
func (*ConfigurationGovernanceV1) NewDeleteAttachmentOptions(ruleID string, attachmentID string) *DeleteAttachmentOptions {
	return &DeleteAttachmentOptions{
		RuleID:       core.StringPtr(ruleID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *DeleteAttachmentOptions) SetRuleID(ruleID string) *DeleteAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *DeleteAttachmentOptions) SetAttachmentID(attachmentID string) *DeleteAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *DeleteAttachmentOptions) SetTransactionID(transactionID string) *DeleteAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteAttachmentOptions) SetHeaders(param map[string]string) *DeleteAttachmentOptions {
	options.Headers = param
	return options
}

// DeleteRuleOptions : The DeleteRule options.
type DeleteRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRuleOptions : Instantiate DeleteRuleOptions
func (*ConfigurationGovernanceV1) NewDeleteRuleOptions(ruleID string) *DeleteRuleOptions {
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

// EnforcementAction : EnforcementAction struct
type EnforcementAction struct {
	// To block a request from completing, use `disallow`. To log the request to Activity Tracker with LogDNA, use
	// `audit_log`.
	Action *string `json:"action" validate:"required"`
}

// Constants associated with the EnforcementAction.Action property.
// To block a request from completing, use `disallow`. To log the request to Activity Tracker with LogDNA, use
// `audit_log`.
const (
	EnforcementActionActionAuditLogConst = "audit_log"
	EnforcementActionActionDisallowConst = "disallow"
)

// NewEnforcementAction : Instantiate EnforcementAction (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewEnforcementAction(action string) (model *EnforcementAction, err error) {
	model = &EnforcementAction{
		Action: core.StringPtr(action),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEnforcementAction unmarshals an instance of EnforcementAction from the specified map of raw messages.
func UnmarshalEnforcementAction(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnforcementAction)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAttachmentOptions : The GetAttachment options.
type GetAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAttachmentOptions : Instantiate GetAttachmentOptions
func (*ConfigurationGovernanceV1) NewGetAttachmentOptions(ruleID string, attachmentID string) *GetAttachmentOptions {
	return &GetAttachmentOptions{
		RuleID:       core.StringPtr(ruleID),
		AttachmentID: core.StringPtr(attachmentID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *GetAttachmentOptions) SetRuleID(ruleID string) *GetAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *GetAttachmentOptions) SetAttachmentID(attachmentID string) *GetAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *GetAttachmentOptions) SetTransactionID(transactionID string) *GetAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAttachmentOptions) SetHeaders(param map[string]string) *GetAttachmentOptions {
	options.Headers = param
	return options
}

// GetRuleOptions : The GetRule options.
type GetRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRuleOptions : Instantiate GetRuleOptions
func (*ConfigurationGovernanceV1) NewGetRuleOptions(ruleID string) *GetRuleOptions {
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

// Link : A link that is used to paginate through available resources.
type Link struct {
	// The URL for the first, previous, next, or last page of resources.
	Href *string `json:"href" validate:"required"`
}

// UnmarshalLink unmarshals an instance of Link from the specified map of raw messages.
func UnmarshalLink(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Link)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListAttachmentsOptions : The ListAttachments options.
type ListAttachmentsOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// The number of resources to retrieve. By default, list operations return the first 100 items. To retrieve a different
	// set of items, use `limit` with `offset` to page through your available resources.
	//
	// **Usage:** If you have 20 rules, and you want to retrieve only the first 5 rules, use
	// `../rules?account_id={account_id}&limit=5`.
	Limit *int64

	// The number of resources to skip. By specifying `offset`, you retrieve a subset of resources that starts with the
	// `offset` value. Use `offset` with `limit` to page through your available resources.
	//
	// **Usage:** If you have 100 rules, and you want to retrieve rules 26 through 50, use
	// `../rules?account_id={account_id}&offset=25&limit=5`.
	Offset *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListAttachmentsOptions : Instantiate ListAttachmentsOptions
func (*ConfigurationGovernanceV1) NewListAttachmentsOptions(ruleID string) *ListAttachmentsOptions {
	return &ListAttachmentsOptions{
		RuleID: core.StringPtr(ruleID),
	}
}

// SetRuleID : Allow user to set RuleID
func (options *ListAttachmentsOptions) SetRuleID(ruleID string) *ListAttachmentsOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *ListAttachmentsOptions) SetTransactionID(transactionID string) *ListAttachmentsOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListAttachmentsOptions) SetLimit(limit int64) *ListAttachmentsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListAttachmentsOptions) SetOffset(offset int64) *ListAttachmentsOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAttachmentsOptions) SetHeaders(param map[string]string) *ListAttachmentsOptions {
	options.Headers = param
	return options
}

// ListRulesOptions : The ListRules options.
type ListRulesOptions struct {
	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Retrieves a list of rules that have scope attachments.
	Attached *bool

	// Retrieves a list of rules that match the labels that you specify.
	Labels *string

	// Retrieves a list of rules that match the scope ID that you specify.
	Scopes *string

	// The number of resources to retrieve. By default, list operations return the first 100 items. To retrieve a different
	// set of items, use `limit` with `offset` to page through your available resources.
	//
	// **Usage:** If you have 20 rules, and you want to retrieve only the first 5 rules, use
	// `../rules?account_id={account_id}&limit=5`.
	Limit *int64

	// The number of resources to skip. By specifying `offset`, you retrieve a subset of resources that starts with the
	// `offset` value. Use `offset` with `limit` to page through your available resources.
	//
	// **Usage:** If you have 100 rules, and you want to retrieve rules 26 through 50, use
	// `../rules?account_id={account_id}&offset=25&limit=5`.
	Offset *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRulesOptions : Instantiate ListRulesOptions
func (*ConfigurationGovernanceV1) NewListRulesOptions(accountID string) *ListRulesOptions {
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

// SetAttached : Allow user to set Attached
func (options *ListRulesOptions) SetAttached(attached bool) *ListRulesOptions {
	options.Attached = core.BoolPtr(attached)
	return options
}

// SetLabels : Allow user to set Labels
func (options *ListRulesOptions) SetLabels(labels string) *ListRulesOptions {
	options.Labels = core.StringPtr(labels)
	return options
}

// SetScopes : Allow user to set Scopes
func (options *ListRulesOptions) SetScopes(scopes string) *ListRulesOptions {
	options.Scopes = core.StringPtr(scopes)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListRulesOptions) SetLimit(limit int64) *ListRulesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListRulesOptions) SetOffset(offset int64) *ListRulesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListRulesOptions) SetHeaders(param map[string]string) *ListRulesOptions {
	options.Headers = param
	return options
}

// Rule : Properties associated with a rule, including both user-settable and server-populated properties.
type Rule struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id,omitempty"`

	// A human-readable alias to assign to your rule.
	Name *string `json:"name" validate:"required"`

	// An extended description of your rule.
	Description *string `json:"description" validate:"required"`

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string `json:"rule_type,omitempty"`

	// The properties that describe the resource that you want to target
	// with the rule.
	Target *TargetResource `json:"target" validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `json:"required_config" validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `json:"enforcement_actions" validate:"required"`

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string `json:"labels,omitempty"`

	// The UUID that uniquely identifies the rule.
	RuleID *string `json:"rule_id,omitempty"`

	// The date the resource was created.
	CreationDate *strfmt.DateTime `json:"creation_date,omitempty"`

	// The unique identifier for the user or application that created the resource.
	CreatedBy *string `json:"created_by,omitempty"`

	// The date the resource was last modified.
	ModificationDate *strfmt.DateTime `json:"modification_date,omitempty"`

	// The unique identifier for the user or application that last modified the resource.
	ModifiedBy *string `json:"modified_by,omitempty"`

	// The number of scope attachments that are associated with the rule.
	NumberOfAttachments *int64 `json:"number_of_attachments,omitempty"`
}

// Constants associated with the Rule.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	RuleRuleTypeUserDefinedConst = "user_defined"
)

// UnmarshalRule unmarshals an instance of Rule from the specified map of raw messages.
func UnmarshalRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rule)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
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
	err = core.UnmarshalPrimitive(m, "rule_type", &obj.RuleType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTargetResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_config", &obj.RequiredConfig, UnmarshalRuleRequiredConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enforcement_actions", &obj.EnforcementActions, UnmarshalEnforcementAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rule_id", &obj.RuleID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creation_date", &obj.CreationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modification_date", &obj.ModificationDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_by", &obj.ModifiedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "number_of_attachments", &obj.NumberOfAttachments)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleCondition : RuleCondition struct
// Models which "extend" this model:
// - RuleConditionSingleProperty
// - RuleConditionOrLvl2
// - RuleConditionAndLvl2
type RuleCondition struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property,omitempty"`

	// The way in which the `property` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator,omitempty"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`

	Or []RuleSingleProperty `json:"or,omitempty"`

	And []RuleSingleProperty `json:"and,omitempty"`
}

// Constants associated with the RuleCondition.Operator property.
// The way in which the `property` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleConditionOperatorIpsInRangeConst           = "ips_in_range"
	RuleConditionOperatorIsEmptyConst              = "is_empty"
	RuleConditionOperatorIsFalseConst              = "is_false"
	RuleConditionOperatorIsNotEmptyConst           = "is_not_empty"
	RuleConditionOperatorIsTrueConst               = "is_true"
	RuleConditionOperatorNumEqualsConst            = "num_equals"
	RuleConditionOperatorNumGreaterThanConst       = "num_greater_than"
	RuleConditionOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleConditionOperatorNumLessThanConst          = "num_less_than"
	RuleConditionOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleConditionOperatorNumNotEqualsConst         = "num_not_equals"
	RuleConditionOperatorStringEqualsConst         = "string_equals"
	RuleConditionOperatorStringMatchConst          = "string_match"
	RuleConditionOperatorStringNotEqualsConst      = "string_not_equals"
	RuleConditionOperatorStringNotMatchConst       = "string_not_match"
	RuleConditionOperatorStringsInListConst        = "strings_in_list"
)

func (*RuleCondition) isaRuleCondition() bool {
	return true
}

type RuleConditionIntf interface {
	isaRuleCondition() bool
}

// UnmarshalRuleCondition unmarshals an instance of RuleCondition from the specified map of raw messages.
func UnmarshalRuleCondition(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleCondition)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
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
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleList : A list of rules.
type RuleList struct {
	// The requested offset for the returned items.
	Offset *int64 `json:"offset" validate:"required"`

	// The requested limit for the returned items.
	Limit *int64 `json:"limit" validate:"required"`

	// The total number of available items.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// The first page of available items.
	First *Link `json:"first" validate:"required"`

	// The last page of available items.
	Last *Link `json:"last" validate:"required"`

	// An array of rules.
	Rules []Rule `json:"rules" validate:"required"`
}

// UnmarshalRuleList unmarshals an instance of RuleList from the specified map of raw messages.
func UnmarshalRuleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleList)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalLink)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "last", &obj.Last, UnmarshalLink)
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

// RuleRequest : User-settable properties associated with a rule to be created or updated.
type RuleRequest struct {
	// Your IBM Cloud account ID.
	AccountID *string `json:"account_id,omitempty"`

	// A human-readable alias to assign to your rule.
	Name *string `json:"name" validate:"required"`

	// An extended description of your rule.
	Description *string `json:"description" validate:"required"`

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string `json:"rule_type,omitempty"`

	// The properties that describe the resource that you want to target
	// with the rule.
	Target *TargetResource `json:"target" validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `json:"required_config" validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `json:"enforcement_actions" validate:"required"`

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string `json:"labels,omitempty"`
}

// Constants associated with the RuleRequest.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	RuleRequestRuleTypeUserDefinedConst = "user_defined"
)

// NewRuleRequest : Instantiate RuleRequest (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequest(name string, description string, target *TargetResource, requiredConfig RuleRequiredConfigIntf, enforcementActions []EnforcementAction) (model *RuleRequest, err error) {
	model = &RuleRequest{
		Name:               core.StringPtr(name),
		Description:        core.StringPtr(description),
		Target:             target,
		RequiredConfig:     requiredConfig,
		EnforcementActions: enforcementActions,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleRequest unmarshals an instance of RuleRequest from the specified map of raw messages.
func UnmarshalRuleRequest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequest)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
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
	err = core.UnmarshalPrimitive(m, "rule_type", &obj.RuleType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "target", &obj.Target, UnmarshalTargetResource)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_config", &obj.RequiredConfig, UnmarshalRuleRequiredConfig)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "enforcement_actions", &obj.EnforcementActions, UnmarshalEnforcementAction)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "labels", &obj.Labels)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfig : RuleRequiredConfig struct
// Models which "extend" this model:
// - RuleRequiredConfigSingleProperty
// - RuleRequiredConfigMultipleProperties
type RuleRequiredConfig struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property,omitempty"`

	// The way in which the `property` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator,omitempty"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`

	Or []RuleConditionIntf `json:"or,omitempty"`

	And []RuleConditionIntf `json:"and,omitempty"`
}

// Constants associated with the RuleRequiredConfig.Operator property.
// The way in which the `property` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleRequiredConfigOperatorIpsInRangeConst           = "ips_in_range"
	RuleRequiredConfigOperatorIsEmptyConst              = "is_empty"
	RuleRequiredConfigOperatorIsFalseConst              = "is_false"
	RuleRequiredConfigOperatorIsNotEmptyConst           = "is_not_empty"
	RuleRequiredConfigOperatorIsTrueConst               = "is_true"
	RuleRequiredConfigOperatorNumEqualsConst            = "num_equals"
	RuleRequiredConfigOperatorNumGreaterThanConst       = "num_greater_than"
	RuleRequiredConfigOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleRequiredConfigOperatorNumLessThanConst          = "num_less_than"
	RuleRequiredConfigOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleRequiredConfigOperatorNumNotEqualsConst         = "num_not_equals"
	RuleRequiredConfigOperatorStringEqualsConst         = "string_equals"
	RuleRequiredConfigOperatorStringMatchConst          = "string_match"
	RuleRequiredConfigOperatorStringNotEqualsConst      = "string_not_equals"
	RuleRequiredConfigOperatorStringNotMatchConst       = "string_not_match"
	RuleRequiredConfigOperatorStringsInListConst        = "strings_in_list"
)

func (*RuleRequiredConfig) isaRuleRequiredConfig() bool {
	return true
}

type RuleRequiredConfigIntf interface {
	isaRuleRequiredConfig() bool
}

// UnmarshalRuleRequiredConfig unmarshals an instance of RuleRequiredConfig from the specified map of raw messages.
func UnmarshalRuleRequiredConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfig)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
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
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleResponseError : RuleResponseError struct
type RuleResponseError struct {
	// Specifies the problem that caused the error.
	Code *string `json:"code" validate:"required"`

	// Describes the problem.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalRuleResponseError unmarshals an instance of RuleResponseError from the specified map of raw messages.
func UnmarshalRuleResponseError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleResponseError)
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

// RuleScope : The extent at which the rule can be attached across your accounts.
type RuleScope struct {
	// A short description or alias to assign to the scope.
	Note *string `json:"note,omitempty"`

	// The ID of the scope, such as an enterprise, account, or account group, that you want to evaluate.
	ScopeID *string `json:"scope_id" validate:"required"`

	// The type of scope that you want to evaluate.
	ScopeType *string `json:"scope_type" validate:"required"`
}

// Constants associated with the RuleScope.ScopeType property.
// The type of scope that you want to evaluate.
const (
	RuleScopeScopeTypeAccountConst                = "account"
	RuleScopeScopeTypeAccountResourceGroupConst   = "account.resource_group"
	RuleScopeScopeTypeEnterpriseConst             = "enterprise"
	RuleScopeScopeTypeEnterpriseAccountConst      = "enterprise.account"
	RuleScopeScopeTypeEnterpriseAccountGroupConst = "enterprise.account_group"
)

// NewRuleScope : Instantiate RuleScope (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleScope(scopeID string, scopeType string) (model *RuleScope, err error) {
	model = &RuleScope{
		ScopeID:   core.StringPtr(scopeID),
		ScopeType: core.StringPtr(scopeType),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleScope unmarshals an instance of RuleScope from the specified map of raw messages.
func UnmarshalRuleScope(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleScope)
	err = core.UnmarshalPrimitive(m, "note", &obj.Note)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_id", &obj.ScopeID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope_type", &obj.ScopeType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
type RuleSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleSingleProperty : Instantiate RuleSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleSingleProperty(property string, operator string) (model *RuleSingleProperty, err error) {
	model = &RuleSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleSingleProperty unmarshals an instance of RuleSingleProperty from the specified map of raw messages.
func UnmarshalRuleSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
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

// RuleTargetAttribute : The attributes that are associated with a rule target.
type RuleTargetAttribute struct {
	Name *string `json:"name" validate:"required"`

	// The way in which the `name` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleTargetAttribute.Operator property.
// The way in which the `name` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleTargetAttributeOperatorIpsInRangeConst           = "ips_in_range"
	RuleTargetAttributeOperatorIsEmptyConst              = "is_empty"
	RuleTargetAttributeOperatorIsFalseConst              = "is_false"
	RuleTargetAttributeOperatorIsNotEmptyConst           = "is_not_empty"
	RuleTargetAttributeOperatorIsTrueConst               = "is_true"
	RuleTargetAttributeOperatorNumEqualsConst            = "num_equals"
	RuleTargetAttributeOperatorNumGreaterThanConst       = "num_greater_than"
	RuleTargetAttributeOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleTargetAttributeOperatorNumLessThanConst          = "num_less_than"
	RuleTargetAttributeOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleTargetAttributeOperatorNumNotEqualsConst         = "num_not_equals"
	RuleTargetAttributeOperatorStringEqualsConst         = "string_equals"
	RuleTargetAttributeOperatorStringMatchConst          = "string_match"
	RuleTargetAttributeOperatorStringNotEqualsConst      = "string_not_equals"
	RuleTargetAttributeOperatorStringNotMatchConst       = "string_not_match"
	RuleTargetAttributeOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleTargetAttribute : Instantiate RuleTargetAttribute (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleTargetAttribute(name string, operator string) (model *RuleTargetAttribute, err error) {
	model = &RuleTargetAttribute{
		Name:     core.StringPtr(name),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRuleTargetAttribute unmarshals an instance of RuleTargetAttribute from the specified map of raw messages.
func UnmarshalRuleTargetAttribute(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleTargetAttribute)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
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

// TargetResource : The properties that describe the resource that you want to target with the rule.
type TargetResource struct {
	// The programmatic name of the IBM Cloud service that you want to target with the rule.
	ServiceName *string `json:"service_name" validate:"required"`

	// The type of resource that you want to target.
	ResourceKind *string `json:"resource_kind" validate:"required"`

	// An extra qualifier for the resource kind. When you include additional attributes, only the resources that match the
	// definition are included in the rule.
	AdditionalTargetAttributes []RuleTargetAttribute `json:"additional_target_attributes,omitempty"`
}

// NewTargetResource : Instantiate TargetResource (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewTargetResource(serviceName string, resourceKind string) (model *TargetResource, err error) {
	model = &TargetResource{
		ServiceName:  core.StringPtr(serviceName),
		ResourceKind: core.StringPtr(resourceKind),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTargetResource unmarshals an instance of TargetResource from the specified map of raw messages.
func UnmarshalTargetResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TargetResource)
	err = core.UnmarshalPrimitive(m, "service_name", &obj.ServiceName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_kind", &obj.ResourceKind)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "additional_target_attributes", &obj.AdditionalTargetAttributes, UnmarshalRuleTargetAttribute)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateAttachmentOptions : The UpdateAttachment options.
type UpdateAttachmentOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// The UUID that uniquely identifies the attachment.
	AttachmentID *string `validate:"required,ne="`

	// Compares a supplied `Etag` value with the version that is stored for the requested resource. If the values match,
	// the server allows the request method to continue.
	//
	// To find the `Etag` value, run a GET request on the resource that you want to modify, and check the response headers.
	IfMatch *string `validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string `validate:"required"`

	// The extent at which the rule can be attached across your accounts.
	IncludedScope *RuleScope `validate:"required"`

	// The extent at which the rule can be excluded from the included scope.
	ExcludedScopes []RuleScope

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateAttachmentOptions : Instantiate UpdateAttachmentOptions
func (*ConfigurationGovernanceV1) NewUpdateAttachmentOptions(ruleID string, attachmentID string, ifMatch string, accountID string, includedScope *RuleScope) *UpdateAttachmentOptions {
	return &UpdateAttachmentOptions{
		RuleID:        core.StringPtr(ruleID),
		AttachmentID:  core.StringPtr(attachmentID),
		IfMatch:       core.StringPtr(ifMatch),
		AccountID:     core.StringPtr(accountID),
		IncludedScope: includedScope,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *UpdateAttachmentOptions) SetRuleID(ruleID string) *UpdateAttachmentOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetAttachmentID : Allow user to set AttachmentID
func (options *UpdateAttachmentOptions) SetAttachmentID(attachmentID string) *UpdateAttachmentOptions {
	options.AttachmentID = core.StringPtr(attachmentID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateAttachmentOptions) SetIfMatch(ifMatch string) *UpdateAttachmentOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateAttachmentOptions) SetAccountID(accountID string) *UpdateAttachmentOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetIncludedScope : Allow user to set IncludedScope
func (options *UpdateAttachmentOptions) SetIncludedScope(includedScope *RuleScope) *UpdateAttachmentOptions {
	options.IncludedScope = includedScope
	return options
}

// SetExcludedScopes : Allow user to set ExcludedScopes
func (options *UpdateAttachmentOptions) SetExcludedScopes(excludedScopes []RuleScope) *UpdateAttachmentOptions {
	options.ExcludedScopes = excludedScopes
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateAttachmentOptions) SetTransactionID(transactionID string) *UpdateAttachmentOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateAttachmentOptions) SetHeaders(param map[string]string) *UpdateAttachmentOptions {
	options.Headers = param
	return options
}

// UpdateRuleOptions : The UpdateRule options.
type UpdateRuleOptions struct {
	// The UUID that uniquely identifies the rule.
	RuleID *string `validate:"required,ne="`

	// Compares a supplied `Etag` value with the version that is stored for the requested resource. If the values match,
	// the server allows the request method to continue.
	//
	// To find the `Etag` value, run a GET request on the resource that you want to modify, and check the response headers.
	IfMatch *string `validate:"required"`

	// A human-readable alias to assign to your rule.
	Name *string `validate:"required"`

	// An extended description of your rule.
	Description *string `validate:"required"`

	// The properties that describe the resource that you want to target
	// with the rule.
	Target *TargetResource `validate:"required"`

	RequiredConfig RuleRequiredConfigIntf `validate:"required"`

	// The actions that the service must run on your behalf when a request to create or modify the target resource does not
	// comply with your conditions.
	EnforcementActions []EnforcementAction `validate:"required"`

	// Your IBM Cloud account ID.
	AccountID *string

	// The type of rule. Rules that you create are `user_defined`.
	RuleType *string

	// Labels that you can use to group and search for similar rules, such as those that help you to meet a specific
	// organization guideline.
	Labels []string

	// The unique identifier that is used to trace an entire request. If you omit this field, the service generates and
	// sends a transaction ID as a response header of the request. In the case of an error, the transaction ID is set in
	// the `trace` field of the response body.
	//
	// **Note:** To help with debugging logs, it is strongly recommended that you generate and supply a `Transaction-Id`
	// with each request.
	TransactionID *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateRuleOptions.RuleType property.
// The type of rule. Rules that you create are `user_defined`.
const (
	UpdateRuleOptionsRuleTypeUserDefinedConst = "user_defined"
)

// NewUpdateRuleOptions : Instantiate UpdateRuleOptions
func (*ConfigurationGovernanceV1) NewUpdateRuleOptions(ruleID string, ifMatch string, name string, description string, target *TargetResource, requiredConfig RuleRequiredConfigIntf, enforcementActions []EnforcementAction) *UpdateRuleOptions {
	return &UpdateRuleOptions{
		RuleID:             core.StringPtr(ruleID),
		IfMatch:            core.StringPtr(ifMatch),
		Name:               core.StringPtr(name),
		Description:        core.StringPtr(description),
		Target:             target,
		RequiredConfig:     requiredConfig,
		EnforcementActions: enforcementActions,
	}
}

// SetRuleID : Allow user to set RuleID
func (options *UpdateRuleOptions) SetRuleID(ruleID string) *UpdateRuleOptions {
	options.RuleID = core.StringPtr(ruleID)
	return options
}

// SetIfMatch : Allow user to set IfMatch
func (options *UpdateRuleOptions) SetIfMatch(ifMatch string) *UpdateRuleOptions {
	options.IfMatch = core.StringPtr(ifMatch)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateRuleOptions) SetName(name string) *UpdateRuleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateRuleOptions) SetDescription(description string) *UpdateRuleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetTarget : Allow user to set Target
func (options *UpdateRuleOptions) SetTarget(target *TargetResource) *UpdateRuleOptions {
	options.Target = target
	return options
}

// SetRequiredConfig : Allow user to set RequiredConfig
func (options *UpdateRuleOptions) SetRequiredConfig(requiredConfig RuleRequiredConfigIntf) *UpdateRuleOptions {
	options.RequiredConfig = requiredConfig
	return options
}

// SetEnforcementActions : Allow user to set EnforcementActions
func (options *UpdateRuleOptions) SetEnforcementActions(enforcementActions []EnforcementAction) *UpdateRuleOptions {
	options.EnforcementActions = enforcementActions
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *UpdateRuleOptions) SetAccountID(accountID string) *UpdateRuleOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetRuleType : Allow user to set RuleType
func (options *UpdateRuleOptions) SetRuleType(ruleType string) *UpdateRuleOptions {
	options.RuleType = core.StringPtr(ruleType)
	return options
}

// SetLabels : Allow user to set Labels
func (options *UpdateRuleOptions) SetLabels(labels []string) *UpdateRuleOptions {
	options.Labels = labels
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *UpdateRuleOptions) SetTransactionID(transactionID string) *UpdateRuleOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRuleOptions) SetHeaders(param map[string]string) *UpdateRuleOptions {
	options.Headers = param
	return options
}

// RuleConditionAndLvl2 : A condition with the `and` logical operator.
// This model "extends" RuleCondition
type RuleConditionAndLvl2 struct {
	Description *string `json:"description,omitempty"`

	And []RuleSingleProperty `json:"and" validate:"required"`
}

// NewRuleConditionAndLvl2 : Instantiate RuleConditionAndLvl2 (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionAndLvl2(and []RuleSingleProperty) (model *RuleConditionAndLvl2, err error) {
	model = &RuleConditionAndLvl2{
		And: and,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionAndLvl2) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionAndLvl2 unmarshals an instance of RuleConditionAndLvl2 from the specified map of raw messages.
func UnmarshalRuleConditionAndLvl2(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionAndLvl2)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleConditionOrLvl2 : A condition with the `or` logical operator.
// This model "extends" RuleCondition
type RuleConditionOrLvl2 struct {
	Description *string `json:"description,omitempty"`

	Or []RuleSingleProperty `json:"or" validate:"required"`
}

// NewRuleConditionOrLvl2 : Instantiate RuleConditionOrLvl2 (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionOrLvl2(or []RuleSingleProperty) (model *RuleConditionOrLvl2, err error) {
	model = &RuleConditionOrLvl2{
		Or: or,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionOrLvl2) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionOrLvl2 unmarshals an instance of RuleConditionOrLvl2 from the specified map of raw messages.
func UnmarshalRuleConditionOrLvl2(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionOrLvl2)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleSingleProperty)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleConditionSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
// This model "extends" RuleCondition
type RuleConditionSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleConditionSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleConditionSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleConditionSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleConditionSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleConditionSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleConditionSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleConditionSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleConditionSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleConditionSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleConditionSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleConditionSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleConditionSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleConditionSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleConditionSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleConditionSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleConditionSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleConditionSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleConditionSingleProperty : Instantiate RuleConditionSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleConditionSingleProperty(property string, operator string) (model *RuleConditionSingleProperty, err error) {
	model = &RuleConditionSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleConditionSingleProperty) isaRuleCondition() bool {
	return true
}

// UnmarshalRuleConditionSingleProperty unmarshals an instance of RuleConditionSingleProperty from the specified map of raw messages.
func UnmarshalRuleConditionSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleConditionSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
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

// RuleRequiredConfigMultipleProperties : The requirements that must be met to determine the resource's level of compliance in accordance with the rule.
//
// Use logical operators (`and`/`or`) to define multiple property checks and conditions. To define requirements for a
// rule, list one or more property check objects in the `and` array. To add conditions to a property check, use `or`.
// Models which "extend" this model:
// - RuleRequiredConfigMultiplePropertiesConditionOr
// - RuleRequiredConfigMultiplePropertiesConditionAnd
// This model "extends" RuleRequiredConfig
type RuleRequiredConfigMultipleProperties struct {
	Description *string `json:"description,omitempty"`

	Or []RuleConditionIntf `json:"or,omitempty"`

	And []RuleConditionIntf `json:"and,omitempty"`
}

func (*RuleRequiredConfigMultipleProperties) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

type RuleRequiredConfigMultiplePropertiesIntf interface {
	RuleRequiredConfigIntf
	isaRuleRequiredConfigMultipleProperties() bool
}

func (*RuleRequiredConfigMultipleProperties) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultipleProperties unmarshals an instance of RuleRequiredConfigMultipleProperties from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultipleProperties(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultipleProperties)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigSingleProperty : The requirement that must be met to determine the resource's level of compliance in accordance with the rule.
//
// To apply a single property check, define a configuration property and the desired value that you want to check
// against.
// This model "extends" RuleRequiredConfig
type RuleRequiredConfigSingleProperty struct {
	Description *string `json:"description,omitempty"`

	// A resource configuration variable that describes the property that you want to apply to the target resource.
	//
	// Available options depend on the target service and resource.
	Property *string `json:"property" validate:"required"`

	// The way in which the `property` field is compared to its value.
	//
	// There are three types of operators: string, numeric, and boolean.
	Operator *string `json:"operator" validate:"required"`

	// The way in which you want your property to be applied.
	//
	// Value options differ depending on the rule that you configure. If you use a boolean operator, you do not need to
	// input a value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the RuleRequiredConfigSingleProperty.Operator property.
// The way in which the `property` field is compared to its value.
//
// There are three types of operators: string, numeric, and boolean.
const (
	RuleRequiredConfigSinglePropertyOperatorIpsInRangeConst           = "ips_in_range"
	RuleRequiredConfigSinglePropertyOperatorIsEmptyConst              = "is_empty"
	RuleRequiredConfigSinglePropertyOperatorIsFalseConst              = "is_false"
	RuleRequiredConfigSinglePropertyOperatorIsNotEmptyConst           = "is_not_empty"
	RuleRequiredConfigSinglePropertyOperatorIsTrueConst               = "is_true"
	RuleRequiredConfigSinglePropertyOperatorNumEqualsConst            = "num_equals"
	RuleRequiredConfigSinglePropertyOperatorNumGreaterThanConst       = "num_greater_than"
	RuleRequiredConfigSinglePropertyOperatorNumGreaterThanEqualsConst = "num_greater_than_equals"
	RuleRequiredConfigSinglePropertyOperatorNumLessThanConst          = "num_less_than"
	RuleRequiredConfigSinglePropertyOperatorNumLessThanEqualsConst    = "num_less_than_equals"
	RuleRequiredConfigSinglePropertyOperatorNumNotEqualsConst         = "num_not_equals"
	RuleRequiredConfigSinglePropertyOperatorStringEqualsConst         = "string_equals"
	RuleRequiredConfigSinglePropertyOperatorStringMatchConst          = "string_match"
	RuleRequiredConfigSinglePropertyOperatorStringNotEqualsConst      = "string_not_equals"
	RuleRequiredConfigSinglePropertyOperatorStringNotMatchConst       = "string_not_match"
	RuleRequiredConfigSinglePropertyOperatorStringsInListConst        = "strings_in_list"
)

// NewRuleRequiredConfigSingleProperty : Instantiate RuleRequiredConfigSingleProperty (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigSingleProperty(property string, operator string) (model *RuleRequiredConfigSingleProperty, err error) {
	model = &RuleRequiredConfigSingleProperty{
		Property: core.StringPtr(property),
		Operator: core.StringPtr(operator),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigSingleProperty) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigSingleProperty unmarshals an instance of RuleRequiredConfigSingleProperty from the specified map of raw messages.
func UnmarshalRuleRequiredConfigSingleProperty(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigSingleProperty)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "property", &obj.Property)
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

// RuleRequiredConfigMultiplePropertiesConditionAnd : A condition with the `and` logical operator.
// This model "extends" RuleRequiredConfigMultipleProperties
type RuleRequiredConfigMultiplePropertiesConditionAnd struct {
	Description *string `json:"description,omitempty"`

	And []RuleConditionIntf `json:"and" validate:"required"`
}

// NewRuleRequiredConfigMultiplePropertiesConditionAnd : Instantiate RuleRequiredConfigMultiplePropertiesConditionAnd (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigMultiplePropertiesConditionAnd(and []RuleConditionIntf) (model *RuleRequiredConfigMultiplePropertiesConditionAnd, err error) {
	model = &RuleRequiredConfigMultiplePropertiesConditionAnd{
		And: and,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigMultiplePropertiesConditionAnd) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

func (*RuleRequiredConfigMultiplePropertiesConditionAnd) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultiplePropertiesConditionAnd unmarshals an instance of RuleRequiredConfigMultiplePropertiesConditionAnd from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultiplePropertiesConditionAnd(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultiplePropertiesConditionAnd)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "and", &obj.And, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RuleRequiredConfigMultiplePropertiesConditionOr : A condition with the `or` logical operator.
// This model "extends" RuleRequiredConfigMultipleProperties
type RuleRequiredConfigMultiplePropertiesConditionOr struct {
	Description *string `json:"description,omitempty"`

	Or []RuleConditionIntf `json:"or" validate:"required"`
}

// NewRuleRequiredConfigMultiplePropertiesConditionOr : Instantiate RuleRequiredConfigMultiplePropertiesConditionOr (Generic Model Constructor)
func (*ConfigurationGovernanceV1) NewRuleRequiredConfigMultiplePropertiesConditionOr(or []RuleConditionIntf) (model *RuleRequiredConfigMultiplePropertiesConditionOr, err error) {
	model = &RuleRequiredConfigMultiplePropertiesConditionOr{
		Or: or,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

func (*RuleRequiredConfigMultiplePropertiesConditionOr) isaRuleRequiredConfigMultipleProperties() bool {
	return true
}

func (*RuleRequiredConfigMultiplePropertiesConditionOr) isaRuleRequiredConfig() bool {
	return true
}

// UnmarshalRuleRequiredConfigMultiplePropertiesConditionOr unmarshals an instance of RuleRequiredConfigMultiplePropertiesConditionOr from the specified map of raw messages.
func UnmarshalRuleRequiredConfigMultiplePropertiesConditionOr(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RuleRequiredConfigMultiplePropertiesConditionOr)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "or", &obj.Or, UnmarshalRuleCondition)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
