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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-d753183b-20201209-163011
 */

// Package usagereportsv4 : Operations and models for the UsageReportsV4 service
package usagereportsv4

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

// UsageReportsV4 : Usage reports for IBM Cloud accounts
//
// Version: 4.0.5
type UsageReportsV4 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://billing.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "usage_reports"

// UsageReportsV4Options : Service options
type UsageReportsV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUsageReportsV4UsingExternalConfig : constructs an instance of UsageReportsV4 with passed in options and external configuration.
func NewUsageReportsV4UsingExternalConfig(options *UsageReportsV4Options) (usageReports *UsageReportsV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	usageReports, err = NewUsageReportsV4(options)
	if err != nil {
		return
	}

	err = usageReports.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = usageReports.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUsageReportsV4 : constructs an instance of UsageReportsV4 with passed in options.
func NewUsageReportsV4(options *UsageReportsV4Options) (service *UsageReportsV4, err error) {
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

	service = &UsageReportsV4{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "usageReports" suitable for processing requests.
func (usageReports *UsageReportsV4) Clone() *UsageReportsV4 {
	if core.IsNil(usageReports) {
		return nil
	}
	clone := *usageReports
	clone.Service = usageReports.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (usageReports *UsageReportsV4) SetServiceURL(url string) error {
	return usageReports.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (usageReports *UsageReportsV4) GetServiceURL() string {
	return usageReports.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (usageReports *UsageReportsV4) SetDefaultHeaders(headers http.Header) {
	usageReports.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (usageReports *UsageReportsV4) SetEnableGzipCompression(enableGzip bool) {
	usageReports.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (usageReports *UsageReportsV4) GetEnableGzipCompression() bool {
	return usageReports.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (usageReports *UsageReportsV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	usageReports.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (usageReports *UsageReportsV4) DisableRetries() {
	usageReports.Service.DisableRetries()
}

// GetAccountSummary : Get account summary
// Returns the summary for the account for a given month. Account billing managers are authorized to access this report.
func (usageReports *UsageReportsV4) GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions) (result *AccountSummary, response *core.DetailedResponse, err error) {
	return usageReports.GetAccountSummaryWithContext(context.Background(), getAccountSummaryOptions)
}

// GetAccountSummaryWithContext is an alternate form of the GetAccountSummary method which supports a Context parameter
func (usageReports *UsageReportsV4) GetAccountSummaryWithContext(ctx context.Context, getAccountSummaryOptions *GetAccountSummaryOptions) (result *AccountSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSummaryOptions, "getAccountSummaryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSummaryOptions, "getAccountSummaryOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":   *getAccountSummaryOptions.AccountID,
		"billingmonth": *getAccountSummaryOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/summary/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountSummaryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetAccountSummary")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountSummary)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetAccountUsage : Get account usage
// Usage for all the resources and plans in an account for a given month. Account billing managers are authorized to
// access this report.
func (usageReports *UsageReportsV4) GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions) (result *AccountUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetAccountUsageWithContext(context.Background(), getAccountUsageOptions)
}

// GetAccountUsageWithContext is an alternate form of the GetAccountUsage method which supports a Context parameter
func (usageReports *UsageReportsV4) GetAccountUsageWithContext(ctx context.Context, getAccountUsageOptions *GetAccountUsageOptions) (result *AccountUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountUsageOptions, "getAccountUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountUsageOptions, "getAccountUsageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":   *getAccountUsageOptions.AccountID,
		"billingmonth": *getAccountUsageOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetAccountUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getAccountUsageOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getAccountUsageOptions.AcceptLanguage))
	}

	if getAccountUsageOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getAccountUsageOptions.Names))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccountUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetResourceGroupUsage : Get resource group usage
// Usage for all the resources and plans in a resource group in a given month. Account billing managers or resource
// group billing managers are authorized to access this report.
func (usageReports *UsageReportsV4) GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions) (result *ResourceGroupUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetResourceGroupUsageWithContext(context.Background(), getResourceGroupUsageOptions)
}

// GetResourceGroupUsageWithContext is an alternate form of the GetResourceGroupUsage method which supports a Context parameter
func (usageReports *UsageReportsV4) GetResourceGroupUsageWithContext(ctx context.Context, getResourceGroupUsageOptions *GetResourceGroupUsageOptions) (result *ResourceGroupUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceGroupUsageOptions, "getResourceGroupUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceGroupUsageOptions, "getResourceGroupUsageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":        *getResourceGroupUsageOptions.AccountID,
		"resource_group_id": *getResourceGroupUsageOptions.ResourceGroupID,
		"billingmonth":      *getResourceGroupUsageOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/resource_groups/{resource_group_id}/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceGroupUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetResourceGroupUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getResourceGroupUsageOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getResourceGroupUsageOptions.AcceptLanguage))
	}

	if getResourceGroupUsageOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getResourceGroupUsageOptions.Names))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceGroupUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetOrgUsage : Get organization usage
// Usage for all the resources and plans in an organization in a given month. Account billing managers or organization
// billing managers are authorized to access this report.
func (usageReports *UsageReportsV4) GetOrgUsage(getOrgUsageOptions *GetOrgUsageOptions) (result *OrgUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetOrgUsageWithContext(context.Background(), getOrgUsageOptions)
}

// GetOrgUsageWithContext is an alternate form of the GetOrgUsage method which supports a Context parameter
func (usageReports *UsageReportsV4) GetOrgUsageWithContext(ctx context.Context, getOrgUsageOptions *GetOrgUsageOptions) (result *OrgUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getOrgUsageOptions, "getOrgUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getOrgUsageOptions, "getOrgUsageOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":      *getOrgUsageOptions.AccountID,
		"organization_id": *getOrgUsageOptions.OrganizationID,
		"billingmonth":    *getOrgUsageOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/organizations/{organization_id}/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getOrgUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetOrgUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getOrgUsageOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getOrgUsageOptions.AcceptLanguage))
	}

	if getOrgUsageOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getOrgUsageOptions.Names))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOrgUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetResourceUsageAccount : Get resource instance usage in an account
// Query for resource instance usage in an account. Filter the results with query parameters. Account billing
// administrator is authorized to access this report.
func (usageReports *UsageReportsV4) GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetResourceUsageAccountWithContext(context.Background(), getResourceUsageAccountOptions)
}

// GetResourceUsageAccountWithContext is an alternate form of the GetResourceUsageAccount method which supports a Context parameter
func (usageReports *UsageReportsV4) GetResourceUsageAccountWithContext(ctx context.Context, getResourceUsageAccountOptions *GetResourceUsageAccountOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceUsageAccountOptions, "getResourceUsageAccountOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceUsageAccountOptions, "getResourceUsageAccountOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":   *getResourceUsageAccountOptions.AccountID,
		"billingmonth": *getResourceUsageAccountOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/resource_instances/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceUsageAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetResourceUsageAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getResourceUsageAccountOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getResourceUsageAccountOptions.AcceptLanguage))
	}

	if getResourceUsageAccountOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getResourceUsageAccountOptions.Names))
	}
	if getResourceUsageAccountOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getResourceUsageAccountOptions.Limit))
	}
	if getResourceUsageAccountOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getResourceUsageAccountOptions.Start))
	}
	if getResourceUsageAccountOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*getResourceUsageAccountOptions.ResourceGroupID))
	}
	if getResourceUsageAccountOptions.OrganizationID != nil {
		builder.AddQuery("organization_id", fmt.Sprint(*getResourceUsageAccountOptions.OrganizationID))
	}
	if getResourceUsageAccountOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getResourceUsageAccountOptions.ResourceInstanceID))
	}
	if getResourceUsageAccountOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getResourceUsageAccountOptions.ResourceID))
	}
	if getResourceUsageAccountOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getResourceUsageAccountOptions.PlanID))
	}
	if getResourceUsageAccountOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getResourceUsageAccountOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancesUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetResourceUsageResourceGroup : Get resource instance usage in a resource group
// Query for resource instance usage in a resource group. Filter the results with query parameters. Account billing
// administrator and resource group billing administrators are authorized to access this report.
func (usageReports *UsageReportsV4) GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetResourceUsageResourceGroupWithContext(context.Background(), getResourceUsageResourceGroupOptions)
}

// GetResourceUsageResourceGroupWithContext is an alternate form of the GetResourceUsageResourceGroup method which supports a Context parameter
func (usageReports *UsageReportsV4) GetResourceUsageResourceGroupWithContext(ctx context.Context, getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceUsageResourceGroupOptions, "getResourceUsageResourceGroupOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceUsageResourceGroupOptions, "getResourceUsageResourceGroupOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":        *getResourceUsageResourceGroupOptions.AccountID,
		"resource_group_id": *getResourceUsageResourceGroupOptions.ResourceGroupID,
		"billingmonth":      *getResourceUsageResourceGroupOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/resource_groups/{resource_group_id}/resource_instances/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceUsageResourceGroupOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetResourceUsageResourceGroup")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getResourceUsageResourceGroupOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getResourceUsageResourceGroupOptions.AcceptLanguage))
	}

	if getResourceUsageResourceGroupOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getResourceUsageResourceGroupOptions.Names))
	}
	if getResourceUsageResourceGroupOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getResourceUsageResourceGroupOptions.Limit))
	}
	if getResourceUsageResourceGroupOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getResourceUsageResourceGroupOptions.Start))
	}
	if getResourceUsageResourceGroupOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getResourceUsageResourceGroupOptions.ResourceInstanceID))
	}
	if getResourceUsageResourceGroupOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getResourceUsageResourceGroupOptions.ResourceID))
	}
	if getResourceUsageResourceGroupOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getResourceUsageResourceGroupOptions.PlanID))
	}
	if getResourceUsageResourceGroupOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getResourceUsageResourceGroupOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancesUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetResourceUsageOrg : Get resource instance usage in an organization
// Query for resource instance usage in an organization. Filter the results with query parameters. Account billing
// administrator and organization billing administrators are authorized to access this report.
func (usageReports *UsageReportsV4) GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	return usageReports.GetResourceUsageOrgWithContext(context.Background(), getResourceUsageOrgOptions)
}

// GetResourceUsageOrgWithContext is an alternate form of the GetResourceUsageOrg method which supports a Context parameter
func (usageReports *UsageReportsV4) GetResourceUsageOrgWithContext(ctx context.Context, getResourceUsageOrgOptions *GetResourceUsageOrgOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceUsageOrgOptions, "getResourceUsageOrgOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceUsageOrgOptions, "getResourceUsageOrgOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"account_id":      *getResourceUsageOrgOptions.AccountID,
		"organization_id": *getResourceUsageOrgOptions.OrganizationID,
		"billingmonth":    *getResourceUsageOrgOptions.Billingmonth,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = usageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(usageReports.Service.Options.URL, `/v4/accounts/{account_id}/organizations/{organization_id}/resource_instances/usage/{billingmonth}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceUsageOrgOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetResourceUsageOrg")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getResourceUsageOrgOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getResourceUsageOrgOptions.AcceptLanguage))
	}

	if getResourceUsageOrgOptions.Names != nil {
		builder.AddQuery("_names", fmt.Sprint(*getResourceUsageOrgOptions.Names))
	}
	if getResourceUsageOrgOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getResourceUsageOrgOptions.Limit))
	}
	if getResourceUsageOrgOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getResourceUsageOrgOptions.Start))
	}
	if getResourceUsageOrgOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getResourceUsageOrgOptions.ResourceInstanceID))
	}
	if getResourceUsageOrgOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getResourceUsageOrgOptions.ResourceID))
	}
	if getResourceUsageOrgOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getResourceUsageOrgOptions.PlanID))
	}
	if getResourceUsageOrgOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getResourceUsageOrgOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = usageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstancesUsage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AccountSummary : A summary of charges and credits for an account.
type AccountSummary struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The month in which usages were incurred. Represented in yyyy-mm format.
	BillingMonth *string `json:"billing_month" validate:"required"`

	// Country.
	BillingCountryCode *string `json:"billing_country_code" validate:"required"`

	// The currency in which the account is billed.
	BillingCurrencyCode *string `json:"billing_currency_code" validate:"required"`

	// Charges related to cloud resources.
	Resources *ResourcesSummary `json:"resources" validate:"required"`

	// The list of offers applicable for the account for the month.
	Offers []Offer `json:"offers" validate:"required"`

	// Support-related charges.
	Support []SupportSummary `json:"support" validate:"required"`

	// A summary of charges and credits related to a subscription.
	Subscription *SubscriptionSummary `json:"subscription" validate:"required"`
}

// UnmarshalAccountSummary unmarshals an instance of AccountSummary from the specified map of raw messages.
func UnmarshalAccountSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountSummary)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_month", &obj.BillingMonth)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_country_code", &obj.BillingCountryCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_currency_code", &obj.BillingCurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourcesSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "offers", &obj.Offers, UnmarshalOffer)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "support", &obj.Support, UnmarshalSupportSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subscription", &obj.Subscription, UnmarshalSubscriptionSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountUsage : The aggregated usage and charges for all the plans in the account.
type AccountUsage struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The target country pricing that should be used.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}

// UnmarshalAccountUsage unmarshals an instance of AccountUsage from the specified map of raw messages.
func UnmarshalAccountUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountUsage)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_country", &obj.PricingCountry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Discount : Information about a discount that is associated with a metric.
type Discount struct {
	// The reference ID of the discount.
	Ref *string `json:"ref" validate:"required"`

	// The name of the discount indicating category.
	Name *string `json:"name" validate:"required"`

	// The name of the discount.
	DisplayName *string `json:"display_name" validate:"required"`

	// The discount percentage.
	Discount *float64 `json:"discount" validate:"required"`
}

// UnmarshalDiscount unmarshals an instance of Discount from the specified map of raw messages.
func UnmarshalDiscount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Discount)
	err = core.UnmarshalPrimitive(m, "ref", &obj.Ref)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "discount", &obj.Discount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetAccountSummaryOptions : The GetAccountSummary options.
type GetAccountSummaryOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountSummaryOptions : Instantiate GetAccountSummaryOptions
func (*UsageReportsV4) NewGetAccountSummaryOptions(accountID string, billingmonth string) *GetAccountSummaryOptions {
	return &GetAccountSummaryOptions{
		AccountID:    core.StringPtr(accountID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountSummaryOptions) SetAccountID(accountID string) *GetAccountSummaryOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetAccountSummaryOptions) SetBillingmonth(billingmonth string) *GetAccountSummaryOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountSummaryOptions) SetHeaders(param map[string]string) *GetAccountSummaryOptions {
	options.Headers = param
	return options
}

// GetAccountUsageOptions : The GetAccountUsage options.
type GetAccountUsageOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountUsageOptions : Instantiate GetAccountUsageOptions
func (*UsageReportsV4) NewGetAccountUsageOptions(accountID string, billingmonth string) *GetAccountUsageOptions {
	return &GetAccountUsageOptions{
		AccountID:    core.StringPtr(accountID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountUsageOptions) SetAccountID(accountID string) *GetAccountUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetAccountUsageOptions) SetBillingmonth(billingmonth string) *GetAccountUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetAccountUsageOptions) SetNames(names bool) *GetAccountUsageOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetAccountUsageOptions) SetAcceptLanguage(acceptLanguage string) *GetAccountUsageOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountUsageOptions) SetHeaders(param map[string]string) *GetAccountUsageOptions {
	options.Headers = param
	return options
}

// GetOrgUsageOptions : The GetOrgUsage options.
type GetOrgUsageOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetOrgUsageOptions : Instantiate GetOrgUsageOptions
func (*UsageReportsV4) NewGetOrgUsageOptions(accountID string, organizationID string, billingmonth string) *GetOrgUsageOptions {
	return &GetOrgUsageOptions{
		AccountID:      core.StringPtr(accountID),
		OrganizationID: core.StringPtr(organizationID),
		Billingmonth:   core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetOrgUsageOptions) SetAccountID(accountID string) *GetOrgUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetOrgUsageOptions) SetOrganizationID(organizationID string) *GetOrgUsageOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetOrgUsageOptions) SetBillingmonth(billingmonth string) *GetOrgUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetOrgUsageOptions) SetNames(names bool) *GetOrgUsageOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetOrgUsageOptions) SetAcceptLanguage(acceptLanguage string) *GetOrgUsageOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetOrgUsageOptions) SetHeaders(param map[string]string) *GetOrgUsageOptions {
	options.Headers = param
	return options
}

// GetResourceGroupUsageOptions : The GetResourceGroupUsage options.
type GetResourceGroupUsageOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// Resource group for which the usage report is requested.
	ResourceGroupID *string `json:"resource_group_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceGroupUsageOptions : Instantiate GetResourceGroupUsageOptions
func (*UsageReportsV4) NewGetResourceGroupUsageOptions(accountID string, resourceGroupID string, billingmonth string) *GetResourceGroupUsageOptions {
	return &GetResourceGroupUsageOptions{
		AccountID:       core.StringPtr(accountID),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		Billingmonth:    core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceGroupUsageOptions) SetAccountID(accountID string) *GetResourceGroupUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetResourceGroupUsageOptions) SetResourceGroupID(resourceGroupID string) *GetResourceGroupUsageOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetResourceGroupUsageOptions) SetBillingmonth(billingmonth string) *GetResourceGroupUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetResourceGroupUsageOptions) SetNames(names bool) *GetResourceGroupUsageOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetResourceGroupUsageOptions) SetAcceptLanguage(acceptLanguage string) *GetResourceGroupUsageOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceGroupUsageOptions) SetHeaders(param map[string]string) *GetResourceGroupUsageOptions {
	options.Headers = param
	return options
}

// GetResourceUsageAccountOptions : The GetResourceUsageAccount options.
type GetResourceUsageAccountOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Number of usage records returned. The default value is 10. Maximum value is 20.
	Limit *int64 `json:"_limit,omitempty"`

	// The offset from which the records must be fetched. Offset information is included in the response.
	Start *string `json:"_start,omitempty"`

	// Filter by resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Filter by organization_id.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Filter by resource instance_id.
	ResourceInstanceID *string `json:"resource_instance_id,omitempty"`

	// Filter by resource_id.
	ResourceID *string `json:"resource_id,omitempty"`

	// Filter by plan_id.
	PlanID *string `json:"plan_id,omitempty"`

	// Region in which the resource instance is provisioned.
	Region *string `json:"region,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceUsageAccountOptions : Instantiate GetResourceUsageAccountOptions
func (*UsageReportsV4) NewGetResourceUsageAccountOptions(accountID string, billingmonth string) *GetResourceUsageAccountOptions {
	return &GetResourceUsageAccountOptions{
		AccountID:    core.StringPtr(accountID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceUsageAccountOptions) SetAccountID(accountID string) *GetResourceUsageAccountOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetResourceUsageAccountOptions) SetBillingmonth(billingmonth string) *GetResourceUsageAccountOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetResourceUsageAccountOptions) SetNames(names bool) *GetResourceUsageAccountOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetResourceUsageAccountOptions) SetAcceptLanguage(acceptLanguage string) *GetResourceUsageAccountOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetResourceUsageAccountOptions) SetLimit(limit int64) *GetResourceUsageAccountOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetResourceUsageAccountOptions) SetStart(start string) *GetResourceUsageAccountOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetResourceUsageAccountOptions) SetResourceGroupID(resourceGroupID string) *GetResourceUsageAccountOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetResourceUsageAccountOptions) SetOrganizationID(organizationID string) *GetResourceUsageAccountOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetResourceUsageAccountOptions) SetResourceInstanceID(resourceInstanceID string) *GetResourceUsageAccountOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetResourceUsageAccountOptions) SetResourceID(resourceID string) *GetResourceUsageAccountOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetResourceUsageAccountOptions) SetPlanID(planID string) *GetResourceUsageAccountOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetResourceUsageAccountOptions) SetRegion(region string) *GetResourceUsageAccountOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceUsageAccountOptions) SetHeaders(param map[string]string) *GetResourceUsageAccountOptions {
	options.Headers = param
	return options
}

// GetResourceUsageOrgOptions : The GetResourceUsageOrg options.
type GetResourceUsageOrgOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Number of usage records returned. The default value is 10. Maximum value is 20.
	Limit *int64 `json:"_limit,omitempty"`

	// The offset from which the records must be fetched. Offset information is included in the response.
	Start *string `json:"_start,omitempty"`

	// Filter by resource instance id.
	ResourceInstanceID *string `json:"resource_instance_id,omitempty"`

	// Filter by resource_id.
	ResourceID *string `json:"resource_id,omitempty"`

	// Filter by plan_id.
	PlanID *string `json:"plan_id,omitempty"`

	// Region in which the resource instance is provisioned.
	Region *string `json:"region,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceUsageOrgOptions : Instantiate GetResourceUsageOrgOptions
func (*UsageReportsV4) NewGetResourceUsageOrgOptions(accountID string, organizationID string, billingmonth string) *GetResourceUsageOrgOptions {
	return &GetResourceUsageOrgOptions{
		AccountID:      core.StringPtr(accountID),
		OrganizationID: core.StringPtr(organizationID),
		Billingmonth:   core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceUsageOrgOptions) SetAccountID(accountID string) *GetResourceUsageOrgOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetResourceUsageOrgOptions) SetOrganizationID(organizationID string) *GetResourceUsageOrgOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetResourceUsageOrgOptions) SetBillingmonth(billingmonth string) *GetResourceUsageOrgOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetResourceUsageOrgOptions) SetNames(names bool) *GetResourceUsageOrgOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetResourceUsageOrgOptions) SetAcceptLanguage(acceptLanguage string) *GetResourceUsageOrgOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetResourceUsageOrgOptions) SetLimit(limit int64) *GetResourceUsageOrgOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetResourceUsageOrgOptions) SetStart(start string) *GetResourceUsageOrgOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetResourceUsageOrgOptions) SetResourceInstanceID(resourceInstanceID string) *GetResourceUsageOrgOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetResourceUsageOrgOptions) SetResourceID(resourceID string) *GetResourceUsageOrgOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetResourceUsageOrgOptions) SetPlanID(planID string) *GetResourceUsageOrgOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetResourceUsageOrgOptions) SetRegion(region string) *GetResourceUsageOrgOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceUsageOrgOptions) SetHeaders(param map[string]string) *GetResourceUsageOrgOptions {
	options.Headers = param
	return options
}

// GetResourceUsageResourceGroupOptions : The GetResourceUsageResourceGroup options.
type GetResourceUsageResourceGroupOptions struct {
	// Account ID for which the usage report is requested.
	AccountID *string `json:"account_id" validate:"required,ne="`

	// Resource group for which the usage report is requested.
	ResourceGroupID *string `json:"resource_group_id" validate:"required,ne="`

	// The billing month for which the usage report is requested.  Format is yyyy-mm.
	Billingmonth *string `json:"billingmonth" validate:"required,ne="`

	// Include the name of every resource, plan, resource instance, organization, and resource group.
	Names *bool `json:"_names,omitempty"`

	// Prioritize the names returned in the order of the specified languages. Language will default to English.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Number of usage records returned. The default value is 10. Maximum value is 20.
	Limit *int64 `json:"_limit,omitempty"`

	// The offset from which the records must be fetched. Offset information is included in the response.
	Start *string `json:"_start,omitempty"`

	// Filter by resource instance id.
	ResourceInstanceID *string `json:"resource_instance_id,omitempty"`

	// Filter by resource_id.
	ResourceID *string `json:"resource_id,omitempty"`

	// Filter by plan_id.
	PlanID *string `json:"plan_id,omitempty"`

	// Region in which the resource instance is provisioned.
	Region *string `json:"region,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceUsageResourceGroupOptions : Instantiate GetResourceUsageResourceGroupOptions
func (*UsageReportsV4) NewGetResourceUsageResourceGroupOptions(accountID string, resourceGroupID string, billingmonth string) *GetResourceUsageResourceGroupOptions {
	return &GetResourceUsageResourceGroupOptions{
		AccountID:       core.StringPtr(accountID),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		Billingmonth:    core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceUsageResourceGroupOptions) SetAccountID(accountID string) *GetResourceUsageResourceGroupOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetResourceUsageResourceGroupOptions) SetResourceGroupID(resourceGroupID string) *GetResourceUsageResourceGroupOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetResourceUsageResourceGroupOptions) SetBillingmonth(billingmonth string) *GetResourceUsageResourceGroupOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetNames : Allow user to set Names
func (options *GetResourceUsageResourceGroupOptions) SetNames(names bool) *GetResourceUsageResourceGroupOptions {
	options.Names = core.BoolPtr(names)
	return options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (options *GetResourceUsageResourceGroupOptions) SetAcceptLanguage(acceptLanguage string) *GetResourceUsageResourceGroupOptions {
	options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetResourceUsageResourceGroupOptions) SetLimit(limit int64) *GetResourceUsageResourceGroupOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetResourceUsageResourceGroupOptions) SetStart(start string) *GetResourceUsageResourceGroupOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetResourceUsageResourceGroupOptions) SetResourceInstanceID(resourceInstanceID string) *GetResourceUsageResourceGroupOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetResourceUsageResourceGroupOptions) SetResourceID(resourceID string) *GetResourceUsageResourceGroupOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetResourceUsageResourceGroupOptions) SetPlanID(planID string) *GetResourceUsageResourceGroupOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetResourceUsageResourceGroupOptions) SetRegion(region string) *GetResourceUsageResourceGroupOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceUsageResourceGroupOptions) SetHeaders(param map[string]string) *GetResourceUsageResourceGroupOptions {
	options.Headers = param
	return options
}

// InstanceUsage : The aggregated usage and charges for an instance.
type InstanceUsage struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the resource instance.
	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The ID of the organization.
	OrganizationID *string `json:"organization_id,omitempty"`

	// The ID of the space.
	Space *string `json:"space,omitempty"`

	// The ID of the consumer.
	ConsumerID *string `json:"consumer_id,omitempty"`

	// The region where instance was provisioned.
	Region *string `json:"region,omitempty"`

	// The pricing region where the usage that was submitted was rated.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// The target country pricing that should be used.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// Is the cost charged to the account.
	Billable *bool `json:"billable" validate:"required"`

	// The ID of the plan where the instance was provisioned and rated.
	PlanID *string `json:"plan_id" validate:"required"`

	// The month.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Usage []Metric `json:"usage" validate:"required"`
}

// UnmarshalInstanceUsage unmarshals an instance of InstanceUsage from the specified map of raw messages.
func UnmarshalInstanceUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstanceUsage)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_instance_id", &obj.ResourceInstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "organization_id", &obj.OrganizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "space", &obj.Space)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "consumer_id", &obj.ConsumerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_region", &obj.PricingRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_country", &obj.PricingCountry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable", &obj.Billable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalMetric)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancesUsageFirst : The link to the first page of the search query.
type InstancesUsageFirst struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalInstancesUsageFirst unmarshals an instance of InstancesUsageFirst from the specified map of raw messages.
func UnmarshalInstancesUsageFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstancesUsageFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancesUsageNext : The link to the next page of the search query.
type InstancesUsageNext struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the next page.
	Offset *string `json:"offset,omitempty"`
}

// UnmarshalInstancesUsageNext unmarshals an instance of InstancesUsageNext from the specified map of raw messages.
func UnmarshalInstancesUsageNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstancesUsageNext)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstancesUsage : The list of instance usage reports.
type InstancesUsage struct {
	// The max number of reports in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The number of reports in the response.
	Count *int64 `json:"count,omitempty"`

	// The link to the first page of the search query.
	First *InstancesUsageFirst `json:"first,omitempty"`

	// The link to the next page of the search query.
	Next *InstancesUsageNext `json:"next,omitempty"`

	// The list of instance usage reports.
	Resources []InstanceUsage `json:"resources,omitempty"`
}

// UnmarshalInstancesUsage unmarshals an instance of InstancesUsage from the specified map of raw messages.
func UnmarshalInstancesUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstancesUsage)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalInstancesUsageFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalInstancesUsageNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalInstanceUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Metric : Information about a metric.
type Metric struct {
	// The name of the metric.
	Metric *string `json:"metric" validate:"required"`

	// The aggregated value for the metric.
	Quantity *float64 `json:"quantity" validate:"required"`

	// The quantity that is used for calculating charges.
	RateableQuantity *float64 `json:"rateable_quantity,omitempty"`

	// The cost incurred by the metric.
	Cost *float64 `json:"cost" validate:"required"`

	// Pre-discounted cost incurred by the metric.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// The price with which the cost was calculated.
	Price []interface{} `json:"price,omitempty"`

	// The unit that qualifies the quantity.
	Unit *string `json:"unit,omitempty"`

	// When set to `true`, the cost is for informational purpose and is not included while calculating the plan charges.
	NonChargeable *bool `json:"non_chargeable,omitempty"`

	// All the discounts applicable to the metric.
	Discounts []Discount `json:"discounts" validate:"required"`
}

// UnmarshalMetric unmarshals an instance of Metric from the specified map of raw messages.
func UnmarshalMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Metric)
	err = core.UnmarshalPrimitive(m, "metric", &obj.Metric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity", &obj.Quantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rateable_quantity", &obj.RateableQuantity)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_chargeable", &obj.NonChargeable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "discounts", &obj.Discounts, UnmarshalDiscount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Offer : Information about an individual offer.
type Offer struct {
	// The ID of the offer.
	OfferID *string `json:"offer_id" validate:"required"`

	// The total credits before applying the offer.
	CreditsTotal *float64 `json:"credits_total" validate:"required"`

	// The template with which the offer was generated.
	OfferTemplate *string `json:"offer_template" validate:"required"`

	// The date from which the offer is valid.
	ValidFrom *strfmt.DateTime `json:"valid_from" validate:"required"`

	// The date until the offer is valid.
	ExpiresOn *strfmt.DateTime `json:"expires_on" validate:"required"`

	// Credit information related to an offer.
	Credits *OfferCredits `json:"credits" validate:"required"`
}

// UnmarshalOffer unmarshals an instance of Offer from the specified map of raw messages.
func UnmarshalOffer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Offer)
	err = core.UnmarshalPrimitive(m, "offer_id", &obj.OfferID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credits_total", &obj.CreditsTotal)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offer_template", &obj.OfferTemplate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "valid_from", &obj.ValidFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expires_on", &obj.ExpiresOn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "credits", &obj.Credits, UnmarshalOfferCredits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferCredits : Credit information related to an offer.
type OfferCredits struct {
	// The available credits in the offer at the beginning of the month.
	StartingBalance *float64 `json:"starting_balance" validate:"required"`

	// The credits used in this month.
	Used *float64 `json:"used" validate:"required"`

	// The remaining credits in the offer.
	Balance *float64 `json:"balance" validate:"required"`
}

// UnmarshalOfferCredits unmarshals an instance of OfferCredits from the specified map of raw messages.
func UnmarshalOfferCredits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferCredits)
	err = core.UnmarshalPrimitive(m, "starting_balance", &obj.StartingBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "used", &obj.Used)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "balance", &obj.Balance)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OrgUsage : The aggregated usage and charges for all the plans in the org.
type OrgUsage struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required"`

	// The target country pricing that should be used.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}

// UnmarshalOrgUsage unmarshals an instance of OrgUsage from the specified map of raw messages.
func UnmarshalOrgUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OrgUsage)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "organization_id", &obj.OrganizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_country", &obj.PricingCountry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Plan : The aggregated values for the plan.
type Plan struct {
	// The ID of the plan.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region for the plan.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// Indicates if the plan charges are billed to the customer.
	Billable *bool `json:"billable" validate:"required"`

	// The total cost incurred by the plan.
	Cost *float64 `json:"cost" validate:"required"`

	// Total pre-discounted cost incurred by the plan.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// All the metrics in the plan.
	Usage []Metric `json:"usage" validate:"required"`

	// All the discounts applicable to the plan.
	Discounts []Discount `json:"discounts" validate:"required"`
}

// UnmarshalPlan unmarshals an instance of Plan from the specified map of raw messages.
func UnmarshalPlan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Plan)
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_region", &obj.PricingRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable", &obj.Billable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rated_cost", &obj.RatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalMetric)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "discounts", &obj.Discounts, UnmarshalDiscount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resource : The container for all the plans in the resource.
type Resource struct {
	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The billable charges for the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The pre-discounted billable charges for the account.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// The non-billable charges for the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// The pre-discounted non-billable charges for the account.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// All the plans in the resource.
	Plans []Plan `json:"plans" validate:"required"`

	// All the discounts applicable to the resource.
	Discounts []Discount `json:"discounts" validate:"required"`
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billable_rated_cost", &obj.BillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_rated_cost", &obj.NonBillableRatedCost)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlan)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "discounts", &obj.Discounts, UnmarshalDiscount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceGroupUsage : The aggregated usage and charges for all the plans in the resource group.
type ResourceGroupUsage struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// The target country pricing that should be used.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}

// UnmarshalResourceGroupUsage unmarshals an instance of ResourceGroupUsage from the specified map of raw messages.
func UnmarshalResourceGroupUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceGroupUsage)
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_country", &obj.PricingCountry)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourcesSummary : Charges related to cloud resources.
type ResourcesSummary struct {
	// The billable charges for all cloud resources used in the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// Non-billable charges for all cloud resources used in the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`
}

// UnmarshalResourcesSummary unmarshals an instance of ResourcesSummary from the specified map of raw messages.
func UnmarshalResourcesSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourcesSummary)
	err = core.UnmarshalPrimitive(m, "billable_cost", &obj.BillableCost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "non_billable_cost", &obj.NonBillableCost)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Subscription : Subscription struct
type Subscription struct {
	// The ID of the subscription.
	SubscriptionID *string `json:"subscription_id" validate:"required"`

	// The charge agreement number of the subsciption.
	ChargeAgreementNumber *string `json:"charge_agreement_number" validate:"required"`

	// Type of the subscription.
	Type *string `json:"type" validate:"required"`

	// The credits available in the subscription for the month.
	SubscriptionAmount *float64 `json:"subscription_amount" validate:"required"`

	// The date from which the subscription was active.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// The date until which the subscription is active. End time is unavailable for PayGO accounts.
	End *strfmt.DateTime `json:"end,omitempty"`

	// The total credits available in the subscription.
	CreditsTotal *float64 `json:"credits_total" validate:"required"`

	// The terms through which the subscription is split into.
	Terms []SubscriptionTerm `json:"terms" validate:"required"`
}

// UnmarshalSubscription unmarshals an instance of Subscription from the specified map of raw messages.
func UnmarshalSubscription(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Subscription)
	err = core.UnmarshalPrimitive(m, "subscription_id", &obj.SubscriptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "charge_agreement_number", &obj.ChargeAgreementNumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subscription_amount", &obj.SubscriptionAmount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "credits_total", &obj.CreditsTotal)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "terms", &obj.Terms, UnmarshalSubscriptionTerm)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionSummary : A summary of charges and credits related to a subscription.
type SubscriptionSummary struct {
	// The charges after exhausting subscription credits and offers credits.
	Overage *float64 `json:"overage,omitempty"`

	// The list of subscriptions applicable for the month.
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

// UnmarshalSubscriptionSummary unmarshals an instance of SubscriptionSummary from the specified map of raw messages.
func UnmarshalSubscriptionSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionSummary)
	err = core.UnmarshalPrimitive(m, "overage", &obj.Overage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "subscriptions", &obj.Subscriptions, UnmarshalSubscription)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionTerm : SubscriptionTerm struct
type SubscriptionTerm struct {
	// The start date of the term.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// The end date of the term.
	End *strfmt.DateTime `json:"end" validate:"required"`

	// Information about credits related to a subscription.
	Credits *SubscriptionTermCredits `json:"credits" validate:"required"`
}

// UnmarshalSubscriptionTerm unmarshals an instance of SubscriptionTerm from the specified map of raw messages.
func UnmarshalSubscriptionTerm(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionTerm)
	err = core.UnmarshalPrimitive(m, "start", &obj.Start)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end", &obj.End)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "credits", &obj.Credits, UnmarshalSubscriptionTermCredits)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SubscriptionTermCredits : Information about credits related to a subscription.
type SubscriptionTermCredits struct {
	// The total credits available for the term.
	Total *float64 `json:"total" validate:"required"`

	// The unused credits in the term at the beginning of the month.
	StartingBalance *float64 `json:"starting_balance" validate:"required"`

	// The credits used in this month.
	Used *float64 `json:"used" validate:"required"`

	// The remaining credits in this term.
	Balance *float64 `json:"balance" validate:"required"`
}

// UnmarshalSubscriptionTermCredits unmarshals an instance of SubscriptionTermCredits from the specified map of raw messages.
func UnmarshalSubscriptionTermCredits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SubscriptionTermCredits)
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "starting_balance", &obj.StartingBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "used", &obj.Used)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "balance", &obj.Balance)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SupportSummary : SupportSummary struct
type SupportSummary struct {
	// The monthly support cost.
	Cost *float64 `json:"cost" validate:"required"`

	// The type of support.
	Type *string `json:"type" validate:"required"`

	// Additional support cost for the month.
	Overage *float64 `json:"overage" validate:"required"`
}

// UnmarshalSupportSummary unmarshals an instance of SupportSummary from the specified map of raw messages.
func UnmarshalSupportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SupportSummary)
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "overage", &obj.Overage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
