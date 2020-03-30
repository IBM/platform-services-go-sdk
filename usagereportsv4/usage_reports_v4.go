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

// Package usagereportsv4 : Operations and models for the UsageReportsV4 service
package usagereportsv4

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// UsageReportsV4 : Usage reports for IBM Cloud accounts
//
// Version: 4.0.2
type UsageReportsV4 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://metering-reporting.ng.bluemix.net"

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

// SetServiceURL sets the service URL
func (usageReports *UsageReportsV4) SetServiceURL(url string) error {
	return usageReports.Service.SetServiceURL(url)
}

// GetAccountSummary : Get account billing summary
// Returns the billing summary for the account for a given month. Users with the Administrator role on the Billing
// service can access this report.
func (usageReports *UsageReportsV4) GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions) (result *AccountSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountSummaryOptions, "getAccountSummaryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountSummaryOptions, "getAccountSummaryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "summary"}
	pathParameters := []string{*getAccountSummaryOptions.AccountID, *getAccountSummaryOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
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

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAccountSummary(m)
		response.Result = result
	}

	return
}

// GetAccountUsage : Get account usage
// Returns aggregated usage for all of the resources and plans in an account in a given month. Users with the
// Administrator role on the Billing service can access this report.
func (usageReports *UsageReportsV4) GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions) (result *AccountUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountUsageOptions, "getAccountUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountUsageOptions, "getAccountUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "usage"}
	pathParameters := []string{*getAccountUsageOptions.AccountID, *getAccountUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
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

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalAccountUsage(m)
		response.Result = result
	}

	return
}

// GetResourceGroupUsage : Get resource group usage
// Returns aggregated usage for all of the resources and plans in a resource group in a given month. Users with the
// Administrator role on the Billing service can access this report.
func (usageReports *UsageReportsV4) GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions) (result *ResourceGroupUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceGroupUsageOptions, "getResourceGroupUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceGroupUsageOptions, "getResourceGroupUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "resource_groups", "usage"}
	pathParameters := []string{*getResourceGroupUsageOptions.AccountID, *getResourceGroupUsageOptions.ResourceGroupID, *getResourceGroupUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
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

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResourceGroupUsage(m)
		response.Result = result
	}

	return
}

// GetOrganizationUsage : Get Cloud Foundry org usage
// Returns aggregated usage for all the resources and plans in a Cloud Foundry organization in a given month. Users with
// the Administrator role on the Billing service or the Billing Manager role on the Cloud Foundry org can access this
// report.
func (usageReports *UsageReportsV4) GetOrganizationUsage(getOrganizationUsageOptions *GetOrganizationUsageOptions) (result *OrganizationUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getOrganizationUsageOptions, "getOrganizationUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getOrganizationUsageOptions, "getOrganizationUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "organizations", "usage"}
	pathParameters := []string{*getOrganizationUsageOptions.AccountID, *getOrganizationUsageOptions.OrganizationID, *getOrganizationUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getOrganizationUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetOrganizationUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalOrganizationUsage(m)
		response.Result = result
	}

	return
}

// GetAccountInstancesUsage : Get resource instance usage in an account
// Returns instance-level usage for resources in an account. Filter the results with query parameters. Users with the
// Administrator role on the Billing service can access this report.
func (usageReports *UsageReportsV4) GetAccountInstancesUsage(getAccountInstancesUsageOptions *GetAccountInstancesUsageOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAccountInstancesUsageOptions, "getAccountInstancesUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAccountInstancesUsageOptions, "getAccountInstancesUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "resource_instances/usage"}
	pathParameters := []string{*getAccountInstancesUsageOptions.AccountID, *getAccountInstancesUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAccountInstancesUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetAccountInstancesUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getAccountInstancesUsageOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getAccountInstancesUsageOptions.Limit))
	}
	if getAccountInstancesUsageOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getAccountInstancesUsageOptions.Start))
	}
	if getAccountInstancesUsageOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*getAccountInstancesUsageOptions.ResourceGroupID))
	}
	if getAccountInstancesUsageOptions.OrganizationID != nil {
		builder.AddQuery("organization_id", fmt.Sprint(*getAccountInstancesUsageOptions.OrganizationID))
	}
	if getAccountInstancesUsageOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getAccountInstancesUsageOptions.ResourceInstanceID))
	}
	if getAccountInstancesUsageOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getAccountInstancesUsageOptions.ResourceID))
	}
	if getAccountInstancesUsageOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getAccountInstancesUsageOptions.PlanID))
	}
	if getAccountInstancesUsageOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getAccountInstancesUsageOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalInstancesUsage(m)
		response.Result = result
	}

	return
}

// GetResourceGroupInstancesUsage : Get resource instance usage in a resource group
// Returns instance-level usage for resources in a resource group. Filter the results with query parameters. Users with
// the Administrator role on the  Billing service can access this report.
func (usageReports *UsageReportsV4) GetResourceGroupInstancesUsage(getResourceGroupInstancesUsageOptions *GetResourceGroupInstancesUsageOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceGroupInstancesUsageOptions, "getResourceGroupInstancesUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceGroupInstancesUsageOptions, "getResourceGroupInstancesUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "resource_groups", "resource_instances/usage"}
	pathParameters := []string{*getResourceGroupInstancesUsageOptions.AccountID, *getResourceGroupInstancesUsageOptions.ResourceGroupID, *getResourceGroupInstancesUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceGroupInstancesUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetResourceGroupInstancesUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getResourceGroupInstancesUsageOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getResourceGroupInstancesUsageOptions.Limit))
	}
	if getResourceGroupInstancesUsageOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getResourceGroupInstancesUsageOptions.Start))
	}
	if getResourceGroupInstancesUsageOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getResourceGroupInstancesUsageOptions.ResourceInstanceID))
	}
	if getResourceGroupInstancesUsageOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getResourceGroupInstancesUsageOptions.ResourceID))
	}
	if getResourceGroupInstancesUsageOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getResourceGroupInstancesUsageOptions.PlanID))
	}
	if getResourceGroupInstancesUsageOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getResourceGroupInstancesUsageOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalInstancesUsage(m)
		response.Result = result
	}

	return
}

// GetOrganizationInstancesUsage : Get resource instance usage in a Cloud Foundry org
// Returns instance-level usage for resources in a Cloud Foundry organization. Filter the results with query parameters.
// Users with the Administrator role on the Billing service or the Billing Manager role on the Cloud Foundry org can
// access this report.
func (usageReports *UsageReportsV4) GetOrganizationInstancesUsage(getOrganizationInstancesUsageOptions *GetOrganizationInstancesUsageOptions) (result *InstancesUsage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getOrganizationInstancesUsageOptions, "getOrganizationInstancesUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getOrganizationInstancesUsageOptions, "getOrganizationInstancesUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/accounts", "organizations", "resource_instances/usage"}
	pathParameters := []string{*getOrganizationInstancesUsageOptions.AccountID, *getOrganizationInstancesUsageOptions.OrganizationID, *getOrganizationInstancesUsageOptions.Billingmonth}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(usageReports.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getOrganizationInstancesUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_reports", "V4", "GetOrganizationInstancesUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getOrganizationInstancesUsageOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getOrganizationInstancesUsageOptions.Limit))
	}
	if getOrganizationInstancesUsageOptions.Start != nil {
		builder.AddQuery("_start", fmt.Sprint(*getOrganizationInstancesUsageOptions.Start))
	}
	if getOrganizationInstancesUsageOptions.ResourceInstanceID != nil {
		builder.AddQuery("resource_instance_id", fmt.Sprint(*getOrganizationInstancesUsageOptions.ResourceInstanceID))
	}
	if getOrganizationInstancesUsageOptions.ResourceID != nil {
		builder.AddQuery("resource_id", fmt.Sprint(*getOrganizationInstancesUsageOptions.ResourceID))
	}
	if getOrganizationInstancesUsageOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*getOrganizationInstancesUsageOptions.PlanID))
	}
	if getOrganizationInstancesUsageOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getOrganizationInstancesUsageOptions.Region))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalInstancesUsage(m)
		response.Result = result
	}

	return
}

// AccountResources : A summary of resource usage.
type AccountResources struct {
	// Billable charges for all cloud resources used in the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// Non-billable charges for all cloud resources used in the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`
}


// UnmarshalAccountResources constructs an instance of AccountResources from the specified map.
func UnmarshalAccountResources(m map[string]interface{}) (result *AccountResources, err error) {
	obj := new(AccountResources)
	obj.BillableCost, err = core.UnmarshalFloat64(m, "billable_cost")
	if err != nil {
		return
	}
	obj.NonBillableCost, err = core.UnmarshalFloat64(m, "non_billable_cost")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAccountResourcesSlice unmarshals a slice of AccountResources instances from the specified list of maps.
func UnmarshalAccountResourcesSlice(s []interface{}) (slice []AccountResources, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AccountResources'")
			return
		}
		obj, e := UnmarshalAccountResources(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAccountResourcesAsProperty unmarshals an instance of AccountResources that is stored as a property
// within the specified map.
func UnmarshalAccountResourcesAsProperty(m map[string]interface{}, propertyName string) (result *AccountResources, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AccountResources'", propertyName)
			return
		}
		result, err = UnmarshalAccountResources(objMap)
	}
	return
}

// UnmarshalAccountResourcesSliceAsProperty unmarshals a slice of AccountResources instances that are stored as a property
// within the specified map.
func UnmarshalAccountResourcesSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AccountResources, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AccountResources'", propertyName)
			return
		}
		slice, err = UnmarshalAccountResourcesSlice(vSlice)
	}
	return
}

// AccountSummary : Summary for the account for a given month.
type AccountSummary struct {
	// The ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// Month in which usages were incurred in `yyyy-mm` format.
	BillingMonth *string `json:"billing_month" validate:"required"`

	// Country.
	BillingCountryCode *string `json:"billing_country_code" validate:"required"`

	// Currency in which the account is billed.
	BillingCurrencyCode *string `json:"billing_currency_code" validate:"required"`

	// A summary of resource usage.
	Resources *AccountResources `json:"resources" validate:"required"`

	// List of offers applicable for the account for the month.
	Offers []Offer `json:"offers" validate:"required"`

	// Subscription information for an account.
	Subscription *SubscriptionInfo `json:"subscription" validate:"required"`

	// Account support information.
	Support []SupportInfo `json:"support" validate:"required"`
}


// UnmarshalAccountSummary constructs an instance of AccountSummary from the specified map.
func UnmarshalAccountSummary(m map[string]interface{}) (result *AccountSummary, err error) {
	obj := new(AccountSummary)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.BillingMonth, err = core.UnmarshalString(m, "billing_month")
	if err != nil {
		return
	}
	obj.BillingCountryCode, err = core.UnmarshalString(m, "billing_country_code")
	if err != nil {
		return
	}
	obj.BillingCurrencyCode, err = core.UnmarshalString(m, "billing_currency_code")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalAccountResourcesAsProperty(m, "resources")
	if err != nil {
		return
	}
	obj.Offers, err = UnmarshalOfferSliceAsProperty(m, "offers")
	if err != nil {
		return
	}
	obj.Subscription, err = UnmarshalSubscriptionInfoAsProperty(m, "subscription")
	if err != nil {
		return
	}
	obj.Support, err = UnmarshalSupportInfoSliceAsProperty(m, "support")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAccountSummarySlice unmarshals a slice of AccountSummary instances from the specified list of maps.
func UnmarshalAccountSummarySlice(s []interface{}) (slice []AccountSummary, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AccountSummary'")
			return
		}
		obj, e := UnmarshalAccountSummary(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAccountSummaryAsProperty unmarshals an instance of AccountSummary that is stored as a property
// within the specified map.
func UnmarshalAccountSummaryAsProperty(m map[string]interface{}, propertyName string) (result *AccountSummary, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AccountSummary'", propertyName)
			return
		}
		result, err = UnmarshalAccountSummary(objMap)
	}
	return
}

// UnmarshalAccountSummarySliceAsProperty unmarshals a slice of AccountSummary instances that are stored as a property
// within the specified map.
func UnmarshalAccountSummarySliceAsProperty(m map[string]interface{}, propertyName string) (slice []AccountSummary, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AccountSummary'", propertyName)
			return
		}
		slice, err = UnmarshalAccountSummarySlice(vSlice)
	}
	return
}

// AccountUsage : Aggregated usage and charges for all the plans in the account.
type AccountUsage struct {
	// ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The country to use for pricing.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans, and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month of the reported usage and charges in `yyyy-mm` format.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}


// UnmarshalAccountUsage constructs an instance of AccountUsage from the specified map.
func UnmarshalAccountUsage(m map[string]interface{}) (result *AccountUsage, err error) {
	obj := new(AccountUsage)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.PricingCountry, err = core.UnmarshalString(m, "pricing_country")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.Month, err = core.UnmarshalString(m, "month")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalResourceSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAccountUsageSlice unmarshals a slice of AccountUsage instances from the specified list of maps.
func UnmarshalAccountUsageSlice(s []interface{}) (slice []AccountUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'AccountUsage'")
			return
		}
		obj, e := UnmarshalAccountUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAccountUsageAsProperty unmarshals an instance of AccountUsage that is stored as a property
// within the specified map.
func UnmarshalAccountUsageAsProperty(m map[string]interface{}, propertyName string) (result *AccountUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'AccountUsage'", propertyName)
			return
		}
		result, err = UnmarshalAccountUsage(objMap)
	}
	return
}

// UnmarshalAccountUsageSliceAsProperty unmarshals a slice of AccountUsage instances that are stored as a property
// within the specified map.
func UnmarshalAccountUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []AccountUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'AccountUsage'", propertyName)
			return
		}
		slice, err = UnmarshalAccountUsageSlice(vSlice)
	}
	return
}

// Credits : Credits.
type Credits struct {
	// Total credit available for the term.
	Total *float64 `json:"total,omitempty"`

	// Credit available in the offer at the beginning of the month.
	StartingBalance *float64 `json:"starting_balance" validate:"required"`

	// Credit used in this month.
	Used *float64 `json:"used" validate:"required"`

	// Remaining credit in the offer.
	Balance *float64 `json:"balance" validate:"required"`
}


// UnmarshalCredits constructs an instance of Credits from the specified map.
func UnmarshalCredits(m map[string]interface{}) (result *Credits, err error) {
	obj := new(Credits)
	obj.Total, err = core.UnmarshalFloat64(m, "total")
	if err != nil {
		return
	}
	obj.StartingBalance, err = core.UnmarshalFloat64(m, "starting_balance")
	if err != nil {
		return
	}
	obj.Used, err = core.UnmarshalFloat64(m, "used")
	if err != nil {
		return
	}
	obj.Balance, err = core.UnmarshalFloat64(m, "balance")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCreditsSlice unmarshals a slice of Credits instances from the specified list of maps.
func UnmarshalCreditsSlice(s []interface{}) (slice []Credits, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Credits'")
			return
		}
		obj, e := UnmarshalCredits(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCreditsAsProperty unmarshals an instance of Credits that is stored as a property
// within the specified map.
func UnmarshalCreditsAsProperty(m map[string]interface{}, propertyName string) (result *Credits, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Credits'", propertyName)
			return
		}
		result, err = UnmarshalCredits(objMap)
	}
	return
}

// UnmarshalCreditsSliceAsProperty unmarshals a slice of Credits instances that are stored as a property
// within the specified map.
func UnmarshalCreditsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Credits, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Credits'", propertyName)
			return
		}
		slice, err = UnmarshalCreditsSlice(vSlice)
	}
	return
}

// GetAccountInstancesUsageOptions : The GetAccountInstancesUsage options.
type GetAccountInstancesUsageOptions struct {
	// Account ID to which the resource instances belong to.
	AccountID *string `json:"account_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

	// Number of usage records returned. The default value is 10. Maximum value is 20.
	Limit *int64 `json:"_limit,omitempty"`

	// The offset from which the records must be fetched. Offset information is included in the response.
	Start *string `json:"_start,omitempty"`

	// Filter by resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Filter by organization_id.
	OrganizationID *string `json:"organization_id,omitempty"`

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

// NewGetAccountInstancesUsageOptions : Instantiate GetAccountInstancesUsageOptions
func (*UsageReportsV4) NewGetAccountInstancesUsageOptions(accountID string, billingmonth string) *GetAccountInstancesUsageOptions {
	return &GetAccountInstancesUsageOptions{
		AccountID: core.StringPtr(accountID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetAccountInstancesUsageOptions) SetAccountID(accountID string) *GetAccountInstancesUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetAccountInstancesUsageOptions) SetBillingmonth(billingmonth string) *GetAccountInstancesUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetAccountInstancesUsageOptions) SetLimit(limit int64) *GetAccountInstancesUsageOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetAccountInstancesUsageOptions) SetStart(start string) *GetAccountInstancesUsageOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetAccountInstancesUsageOptions) SetResourceGroupID(resourceGroupID string) *GetAccountInstancesUsageOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetAccountInstancesUsageOptions) SetOrganizationID(organizationID string) *GetAccountInstancesUsageOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetAccountInstancesUsageOptions) SetResourceInstanceID(resourceInstanceID string) *GetAccountInstancesUsageOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetAccountInstancesUsageOptions) SetResourceID(resourceID string) *GetAccountInstancesUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetAccountInstancesUsageOptions) SetPlanID(planID string) *GetAccountInstancesUsageOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetAccountInstancesUsageOptions) SetRegion(region string) *GetAccountInstancesUsageOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAccountInstancesUsageOptions) SetHeaders(param map[string]string) *GetAccountInstancesUsageOptions {
	options.Headers = param
	return options
}

// GetAccountSummaryOptions : The GetAccountSummary options.
type GetAccountSummaryOptions struct {
	// Account ID for which the summary is requested.
	AccountID *string `json:"account_id" validate:"required"`

	// Billing month for which the summary is requested in `yyyy-mm` format.
	Billingmonth *string `json:"billingmonth" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountSummaryOptions : Instantiate GetAccountSummaryOptions
func (*UsageReportsV4) NewGetAccountSummaryOptions(accountID string, billingmonth string) *GetAccountSummaryOptions {
	return &GetAccountSummaryOptions{
		AccountID: core.StringPtr(accountID),
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
	// Account ID for which the usage is requested.
	AccountID *string `json:"account_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAccountUsageOptions : Instantiate GetAccountUsageOptions
func (*UsageReportsV4) NewGetAccountUsageOptions(accountID string, billingmonth string) *GetAccountUsageOptions {
	return &GetAccountUsageOptions{
		AccountID: core.StringPtr(accountID),
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

// SetHeaders : Allow user to set Headers
func (options *GetAccountUsageOptions) SetHeaders(param map[string]string) *GetAccountUsageOptions {
	options.Headers = param
	return options
}

// GetOrganizationInstancesUsageOptions : The GetOrganizationInstancesUsage options.
type GetOrganizationInstancesUsageOptions struct {
	// ID of the account under which the organization is present.
	AccountID *string `json:"account_id" validate:"required"`

	// ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

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

// NewGetOrganizationInstancesUsageOptions : Instantiate GetOrganizationInstancesUsageOptions
func (*UsageReportsV4) NewGetOrganizationInstancesUsageOptions(accountID string, organizationID string, billingmonth string) *GetOrganizationInstancesUsageOptions {
	return &GetOrganizationInstancesUsageOptions{
		AccountID: core.StringPtr(accountID),
		OrganizationID: core.StringPtr(organizationID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetOrganizationInstancesUsageOptions) SetAccountID(accountID string) *GetOrganizationInstancesUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetOrganizationInstancesUsageOptions) SetOrganizationID(organizationID string) *GetOrganizationInstancesUsageOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetOrganizationInstancesUsageOptions) SetBillingmonth(billingmonth string) *GetOrganizationInstancesUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetOrganizationInstancesUsageOptions) SetLimit(limit int64) *GetOrganizationInstancesUsageOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetOrganizationInstancesUsageOptions) SetStart(start string) *GetOrganizationInstancesUsageOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetOrganizationInstancesUsageOptions) SetResourceInstanceID(resourceInstanceID string) *GetOrganizationInstancesUsageOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetOrganizationInstancesUsageOptions) SetResourceID(resourceID string) *GetOrganizationInstancesUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetOrganizationInstancesUsageOptions) SetPlanID(planID string) *GetOrganizationInstancesUsageOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetOrganizationInstancesUsageOptions) SetRegion(region string) *GetOrganizationInstancesUsageOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetOrganizationInstancesUsageOptions) SetHeaders(param map[string]string) *GetOrganizationInstancesUsageOptions {
	options.Headers = param
	return options
}

// GetOrganizationUsageOptions : The GetOrganizationUsage options.
type GetOrganizationUsageOptions struct {
	// ID of the account containing the organization.
	AccountID *string `json:"account_id" validate:"required"`

	// ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetOrganizationUsageOptions : Instantiate GetOrganizationUsageOptions
func (*UsageReportsV4) NewGetOrganizationUsageOptions(accountID string, organizationID string, billingmonth string) *GetOrganizationUsageOptions {
	return &GetOrganizationUsageOptions{
		AccountID: core.StringPtr(accountID),
		OrganizationID: core.StringPtr(organizationID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetOrganizationUsageOptions) SetAccountID(accountID string) *GetOrganizationUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetOrganizationID : Allow user to set OrganizationID
func (options *GetOrganizationUsageOptions) SetOrganizationID(organizationID string) *GetOrganizationUsageOptions {
	options.OrganizationID = core.StringPtr(organizationID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetOrganizationUsageOptions) SetBillingmonth(billingmonth string) *GetOrganizationUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetOrganizationUsageOptions) SetHeaders(param map[string]string) *GetOrganizationUsageOptions {
	options.Headers = param
	return options
}

// GetResourceGroupInstancesUsageOptions : The GetResourceGroupInstancesUsage options.
type GetResourceGroupInstancesUsageOptions struct {
	// ID of the account in which the resource group is present.
	AccountID *string `json:"account_id" validate:"required"`

	// ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

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

// NewGetResourceGroupInstancesUsageOptions : Instantiate GetResourceGroupInstancesUsageOptions
func (*UsageReportsV4) NewGetResourceGroupInstancesUsageOptions(accountID string, resourceGroupID string, billingmonth string) *GetResourceGroupInstancesUsageOptions {
	return &GetResourceGroupInstancesUsageOptions{
		AccountID: core.StringPtr(accountID),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		Billingmonth: core.StringPtr(billingmonth),
	}
}

// SetAccountID : Allow user to set AccountID
func (options *GetResourceGroupInstancesUsageOptions) SetAccountID(accountID string) *GetResourceGroupInstancesUsageOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetResourceGroupInstancesUsageOptions) SetResourceGroupID(resourceGroupID string) *GetResourceGroupInstancesUsageOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetBillingmonth : Allow user to set Billingmonth
func (options *GetResourceGroupInstancesUsageOptions) SetBillingmonth(billingmonth string) *GetResourceGroupInstancesUsageOptions {
	options.Billingmonth = core.StringPtr(billingmonth)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetResourceGroupInstancesUsageOptions) SetLimit(limit int64) *GetResourceGroupInstancesUsageOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetStart : Allow user to set Start
func (options *GetResourceGroupInstancesUsageOptions) SetStart(start string) *GetResourceGroupInstancesUsageOptions {
	options.Start = core.StringPtr(start)
	return options
}

// SetResourceInstanceID : Allow user to set ResourceInstanceID
func (options *GetResourceGroupInstancesUsageOptions) SetResourceInstanceID(resourceInstanceID string) *GetResourceGroupInstancesUsageOptions {
	options.ResourceInstanceID = core.StringPtr(resourceInstanceID)
	return options
}

// SetResourceID : Allow user to set ResourceID
func (options *GetResourceGroupInstancesUsageOptions) SetResourceID(resourceID string) *GetResourceGroupInstancesUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *GetResourceGroupInstancesUsageOptions) SetPlanID(planID string) *GetResourceGroupInstancesUsageOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetResourceGroupInstancesUsageOptions) SetRegion(region string) *GetResourceGroupInstancesUsageOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceGroupInstancesUsageOptions) SetHeaders(param map[string]string) *GetResourceGroupInstancesUsageOptions {
	options.Headers = param
	return options
}

// GetResourceGroupUsageOptions : The GetResourceGroupUsage options.
type GetResourceGroupUsageOptions struct {
	// Account ID containing the resource group.
	AccountID *string `json:"account_id" validate:"required"`

	// Resource group for which the usage is requested.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// Month for which the usage is requested.
	Billingmonth *string `json:"billingmonth" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceGroupUsageOptions : Instantiate GetResourceGroupUsageOptions
func (*UsageReportsV4) NewGetResourceGroupUsageOptions(accountID string, resourceGroupID string, billingmonth string) *GetResourceGroupUsageOptions {
	return &GetResourceGroupUsageOptions{
		AccountID: core.StringPtr(accountID),
		ResourceGroupID: core.StringPtr(resourceGroupID),
		Billingmonth: core.StringPtr(billingmonth),
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

// SetHeaders : Allow user to set Headers
func (options *GetResourceGroupUsageOptions) SetHeaders(param map[string]string) *GetResourceGroupUsageOptions {
	options.Headers = param
	return options
}

// InstanceUsage : Aggregated usage and charges for an instance.
type InstanceUsage struct {
	// ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// The ID of the resource instance.
	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The ID of the organization.
	OrganizationID *string `json:"organization_id,omitempty"`

	// The ID of the space.
	Space *string `json:"space,omitempty"`

	// The ID of the consumer.
	ConsumerID *string `json:"consumer_id,omitempty"`

	// The region in which the instance was provisioned.
	Region *string `json:"region,omitempty"`

	// The pricing region in which the usage submitted was rated.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// The country to use for pricing.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans, and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// Whether the cost is charged to the account.
	Billable *bool `json:"billable" validate:"required"`

	// The ID of the plan with which the instance was provisioned and rated.
	PlanID *string `json:"plan_id" validate:"required"`

	// The month of the reported usage and charges in `yyyy-mm` format.
	Month *string `json:"month" validate:"required"`

	// All of the resources used in the account.
	Usage []Metric `json:"usage" validate:"required"`
}


// UnmarshalInstanceUsage constructs an instance of InstanceUsage from the specified map.
func UnmarshalInstanceUsage(m map[string]interface{}) (result *InstanceUsage, err error) {
	obj := new(InstanceUsage)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.ResourceInstanceID, err = core.UnmarshalString(m, "resource_instance_id")
	if err != nil {
		return
	}
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.ResourceGroupID, err = core.UnmarshalString(m, "resource_group_id")
	if err != nil {
		return
	}
	obj.OrganizationID, err = core.UnmarshalString(m, "organization_id")
	if err != nil {
		return
	}
	obj.Space, err = core.UnmarshalString(m, "space")
	if err != nil {
		return
	}
	obj.ConsumerID, err = core.UnmarshalString(m, "consumer_id")
	if err != nil {
		return
	}
	obj.Region, err = core.UnmarshalString(m, "region")
	if err != nil {
		return
	}
	obj.PricingRegion, err = core.UnmarshalString(m, "pricing_region")
	if err != nil {
		return
	}
	obj.PricingCountry, err = core.UnmarshalString(m, "pricing_country")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.Billable, err = core.UnmarshalBool(m, "billable")
	if err != nil {
		return
	}
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.Month, err = core.UnmarshalString(m, "month")
	if err != nil {
		return
	}
	obj.Usage, err = UnmarshalMetricSliceAsProperty(m, "usage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalInstanceUsageSlice unmarshals a slice of InstanceUsage instances from the specified list of maps.
func UnmarshalInstanceUsageSlice(s []interface{}) (slice []InstanceUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'InstanceUsage'")
			return
		}
		obj, e := UnmarshalInstanceUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalInstanceUsageAsProperty unmarshals an instance of InstanceUsage that is stored as a property
// within the specified map.
func UnmarshalInstanceUsageAsProperty(m map[string]interface{}, propertyName string) (result *InstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'InstanceUsage'", propertyName)
			return
		}
		result, err = UnmarshalInstanceUsage(objMap)
	}
	return
}

// UnmarshalInstanceUsageSliceAsProperty unmarshals a slice of InstanceUsage instances that are stored as a property
// within the specified map.
func UnmarshalInstanceUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []InstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'InstanceUsage'", propertyName)
			return
		}
		slice, err = UnmarshalInstanceUsageSlice(vSlice)
	}
	return
}

// InstancesUsage : List of instnace usage reports.
type InstancesUsage struct {
	// The maximum number of reports in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The number of reports in the response.
	Count *int64 `json:"count,omitempty"`

	// Link to the first page of the search query.
	First *PageLink `json:"first,omitempty"`

	// Link to the next page of the search query.
	Next *PageLink `json:"next,omitempty"`

	// The list of instance usage reports.
	Resources []InstanceUsage `json:"resources,omitempty"`
}


// UnmarshalInstancesUsage constructs an instance of InstancesUsage from the specified map.
func UnmarshalInstancesUsage(m map[string]interface{}) (result *InstancesUsage, err error) {
	obj := new(InstancesUsage)
	obj.Limit, err = core.UnmarshalInt64(m, "limit")
	if err != nil {
		return
	}
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalPageLinkAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalPageLinkAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalInstanceUsageSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalInstancesUsageSlice unmarshals a slice of InstancesUsage instances from the specified list of maps.
func UnmarshalInstancesUsageSlice(s []interface{}) (slice []InstancesUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'InstancesUsage'")
			return
		}
		obj, e := UnmarshalInstancesUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalInstancesUsageAsProperty unmarshals an instance of InstancesUsage that is stored as a property
// within the specified map.
func UnmarshalInstancesUsageAsProperty(m map[string]interface{}, propertyName string) (result *InstancesUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'InstancesUsage'", propertyName)
			return
		}
		result, err = UnmarshalInstancesUsage(objMap)
	}
	return
}

// UnmarshalInstancesUsageSliceAsProperty unmarshals a slice of InstancesUsage instances that are stored as a property
// within the specified map.
func UnmarshalInstancesUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []InstancesUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'InstancesUsage'", propertyName)
			return
		}
		slice, err = UnmarshalInstancesUsageSlice(vSlice)
	}
	return
}

// Metric : Metric for reporting usage.
type Metric struct {
	// Name of the metric.
	Metric *string `json:"metric" validate:"required"`

	// Aggregated value for the metric.
	Quantity *float64 `json:"quantity" validate:"required"`

	// Quantity that is used for calculating charges.
	RateableQuantity *float64 `json:"rateable_quantity,omitempty"`

	// Cost incurred by the metric.
	Cost *float64 `json:"cost" validate:"required"`

	// The price with which cost was calculated.
	Price []map[string]interface{} `json:"price,omitempty"`

	// Unit qualifying the quantity.
	Unit *string `json:"unit,omitempty"`

	// When set to `true`, the cost is for informational purpose and is not included while calculating the plan charges.
	NonChargeable *bool `json:"non_chargeable,omitempty"`
}


// UnmarshalMetric constructs an instance of Metric from the specified map.
func UnmarshalMetric(m map[string]interface{}) (result *Metric, err error) {
	obj := new(Metric)
	obj.Metric, err = core.UnmarshalString(m, "metric")
	if err != nil {
		return
	}
	obj.Quantity, err = core.UnmarshalFloat64(m, "quantity")
	if err != nil {
		return
	}
	obj.RateableQuantity, err = core.UnmarshalFloat64(m, "rateable_quantity")
	if err != nil {
		return
	}
	obj.Cost, err = core.UnmarshalFloat64(m, "cost")
	if err != nil {
		return
	}
	obj.Unit, err = core.UnmarshalString(m, "unit")
	if err != nil {
		return
	}
	obj.NonChargeable, err = core.UnmarshalBool(m, "non_chargeable")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalMetricSlice unmarshals a slice of Metric instances from the specified list of maps.
func UnmarshalMetricSlice(s []interface{}) (slice []Metric, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Metric'")
			return
		}
		obj, e := UnmarshalMetric(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalMetricAsProperty unmarshals an instance of Metric that is stored as a property
// within the specified map.
func UnmarshalMetricAsProperty(m map[string]interface{}, propertyName string) (result *Metric, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Metric'", propertyName)
			return
		}
		result, err = UnmarshalMetric(objMap)
	}
	return
}

// UnmarshalMetricSliceAsProperty unmarshals a slice of Metric instances that are stored as a property
// within the specified map.
func UnmarshalMetricSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Metric, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Metric'", propertyName)
			return
		}
		slice, err = UnmarshalMetricSlice(vSlice)
	}
	return
}

// Offer : Offer.
type Offer struct {
	// ID of the offer.
	OfferID *string `json:"offer_id" validate:"required"`

	// Total credits before applying the offer.
	CreditsTotal *float64 `json:"credits_total" validate:"required"`

	// Template with which the offer was generated.
	OfferTemplate *string `json:"offer_template" validate:"required"`

	// Date from which the offer is valid.
	ValidFrom *strfmt.DateTime `json:"valid_from" validate:"required"`

	// Date until which the offer is valid.
	ExpiresOn *strfmt.DateTime `json:"expires_on" validate:"required"`

	// Credits.
	Credits *Credits `json:"credits" validate:"required"`
}


// UnmarshalOffer constructs an instance of Offer from the specified map.
func UnmarshalOffer(m map[string]interface{}) (result *Offer, err error) {
	obj := new(Offer)
	obj.OfferID, err = core.UnmarshalString(m, "offer_id")
	if err != nil {
		return
	}
	obj.CreditsTotal, err = core.UnmarshalFloat64(m, "credits_total")
	if err != nil {
		return
	}
	obj.OfferTemplate, err = core.UnmarshalString(m, "offer_template")
	if err != nil {
		return
	}
	obj.ValidFrom, err = core.UnmarshalDateTime(m, "valid_from")
	if err != nil {
		return
	}
	obj.ExpiresOn, err = core.UnmarshalDateTime(m, "expires_on")
	if err != nil {
		return
	}
	obj.Credits, err = UnmarshalCreditsAsProperty(m, "credits")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalOfferSlice unmarshals a slice of Offer instances from the specified list of maps.
func UnmarshalOfferSlice(s []interface{}) (slice []Offer, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Offer'")
			return
		}
		obj, e := UnmarshalOffer(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalOfferAsProperty unmarshals an instance of Offer that is stored as a property
// within the specified map.
func UnmarshalOfferAsProperty(m map[string]interface{}, propertyName string) (result *Offer, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Offer'", propertyName)
			return
		}
		result, err = UnmarshalOffer(objMap)
	}
	return
}

// UnmarshalOfferSliceAsProperty unmarshals a slice of Offer instances that are stored as a property
// within the specified map.
func UnmarshalOfferSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Offer, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Offer'", propertyName)
			return
		}
		slice, err = UnmarshalOfferSlice(vSlice)
	}
	return
}

// OrganizationUsage : Aggregated usage and charges for all the plans in the org.
type OrganizationUsage struct {
	// ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// ID of the organization.
	OrganizationID *string `json:"organization_id" validate:"required"`

	// The country to use for pricing.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans, and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month of the reported usage and charges in `yyyy-mm` format.
	Month *string `json:"month" validate:"required"`

	// All the resource used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}


// UnmarshalOrganizationUsage constructs an instance of OrganizationUsage from the specified map.
func UnmarshalOrganizationUsage(m map[string]interface{}) (result *OrganizationUsage, err error) {
	obj := new(OrganizationUsage)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.OrganizationID, err = core.UnmarshalString(m, "organization_id")
	if err != nil {
		return
	}
	obj.PricingCountry, err = core.UnmarshalString(m, "pricing_country")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.Month, err = core.UnmarshalString(m, "month")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalResourceSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalOrganizationUsageSlice unmarshals a slice of OrganizationUsage instances from the specified list of maps.
func UnmarshalOrganizationUsageSlice(s []interface{}) (slice []OrganizationUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'OrganizationUsage'")
			return
		}
		obj, e := UnmarshalOrganizationUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalOrganizationUsageAsProperty unmarshals an instance of OrganizationUsage that is stored as a property
// within the specified map.
func UnmarshalOrganizationUsageAsProperty(m map[string]interface{}, propertyName string) (result *OrganizationUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'OrganizationUsage'", propertyName)
			return
		}
		result, err = UnmarshalOrganizationUsage(objMap)
	}
	return
}

// UnmarshalOrganizationUsageSliceAsProperty unmarshals a slice of OrganizationUsage instances that are stored as a property
// within the specified map.
func UnmarshalOrganizationUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []OrganizationUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'OrganizationUsage'", propertyName)
			return
		}
		slice, err = UnmarshalOrganizationUsageSlice(vSlice)
	}
	return
}

// PageLink : Link to a page of a paginated list.
type PageLink struct {
	// The URL for the page.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the page.
	Offset *string `json:"offset,omitempty"`
}


// UnmarshalPageLink constructs an instance of PageLink from the specified map.
func UnmarshalPageLink(m map[string]interface{}) (result *PageLink, err error) {
	obj := new(PageLink)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	obj.Offset, err = core.UnmarshalString(m, "offset")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPageLinkSlice unmarshals a slice of PageLink instances from the specified list of maps.
func UnmarshalPageLinkSlice(s []interface{}) (slice []PageLink, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PageLink'")
			return
		}
		obj, e := UnmarshalPageLink(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPageLinkAsProperty unmarshals an instance of PageLink that is stored as a property
// within the specified map.
func UnmarshalPageLinkAsProperty(m map[string]interface{}, propertyName string) (result *PageLink, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'PageLink'", propertyName)
			return
		}
		result, err = UnmarshalPageLink(objMap)
	}
	return
}

// UnmarshalPageLinkSliceAsProperty unmarshals a slice of PageLink instances that are stored as a property
// within the specified map.
func UnmarshalPageLinkSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PageLink, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PageLink'", propertyName)
			return
		}
		slice, err = UnmarshalPageLinkSlice(vSlice)
	}
	return
}

// Plan : Aggregated values for the plan.
type Plan struct {
	// The ID of the plan.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region for the plan.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// Whether the plan charges are billed to the customer.
	Billable *bool `json:"billable" validate:"required"`

	// Total cost incurred by the plan.
	Cost *float64 `json:"cost" validate:"required"`

	// All of the metrics in the plan.
	Usage []Metric `json:"usage" validate:"required"`
}


// UnmarshalPlan constructs an instance of Plan from the specified map.
func UnmarshalPlan(m map[string]interface{}) (result *Plan, err error) {
	obj := new(Plan)
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.PricingRegion, err = core.UnmarshalString(m, "pricing_region")
	if err != nil {
		return
	}
	obj.Billable, err = core.UnmarshalBool(m, "billable")
	if err != nil {
		return
	}
	obj.Cost, err = core.UnmarshalFloat64(m, "cost")
	if err != nil {
		return
	}
	obj.Usage, err = UnmarshalMetricSliceAsProperty(m, "usage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPlanSlice unmarshals a slice of Plan instances from the specified list of maps.
func UnmarshalPlanSlice(s []interface{}) (slice []Plan, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Plan'")
			return
		}
		obj, e := UnmarshalPlan(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPlanAsProperty unmarshals an instance of Plan that is stored as a property
// within the specified map.
func UnmarshalPlanAsProperty(m map[string]interface{}, propertyName string) (result *Plan, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Plan'", propertyName)
			return
		}
		result, err = UnmarshalPlan(objMap)
	}
	return
}

// UnmarshalPlanSliceAsProperty unmarshals a slice of Plan instances that are stored as a property
// within the specified map.
func UnmarshalPlanSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Plan, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Plan'", propertyName)
			return
		}
		slice, err = UnmarshalPlanSlice(vSlice)
	}
	return
}

// Resource : Container for all the plans in the resource.
type Resource struct {
	// ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The billable charges for the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The non billable charges for the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// All the plans in the resource.
	Plans []Plan `json:"plans" validate:"required"`
}


// UnmarshalResource constructs an instance of Resource from the specified map.
func UnmarshalResource(m map[string]interface{}) (result *Resource, err error) {
	obj := new(Resource)
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.BillableCost, err = core.UnmarshalFloat64(m, "billable_cost")
	if err != nil {
		return
	}
	obj.NonBillableCost, err = core.UnmarshalFloat64(m, "non_billable_cost")
	if err != nil {
		return
	}
	obj.Plans, err = UnmarshalPlanSliceAsProperty(m, "plans")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceSlice unmarshals a slice of Resource instances from the specified list of maps.
func UnmarshalResourceSlice(s []interface{}) (slice []Resource, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resource'")
			return
		}
		obj, e := UnmarshalResource(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceAsProperty unmarshals an instance of Resource that is stored as a property
// within the specified map.
func UnmarshalResourceAsProperty(m map[string]interface{}, propertyName string) (result *Resource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resource'", propertyName)
			return
		}
		result, err = UnmarshalResource(objMap)
	}
	return
}

// UnmarshalResourceSliceAsProperty unmarshals a slice of Resource instances that are stored as a property
// within the specified map.
func UnmarshalResourceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resource'", propertyName)
			return
		}
		slice, err = UnmarshalResourceSlice(vSlice)
	}
	return
}

// ResourceGroupUsage : Aggregated usage and charges for all the plans in the resource group.
type ResourceGroupUsage struct {
	// ID of the account.
	AccountID *string `json:"account_id" validate:"required"`

	// ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id" validate:"required"`

	// The country to use for pricing.
	PricingCountry *string `json:"pricing_country" validate:"required"`

	// The currency for the cost fields in the resources, plans, and metrics.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// The month of the reported usage and charges in `yyyy-mm` format.
	Month *string `json:"month" validate:"required"`

	// All of the resources used in the account.
	Resources []Resource `json:"resources" validate:"required"`
}


// UnmarshalResourceGroupUsage constructs an instance of ResourceGroupUsage from the specified map.
func UnmarshalResourceGroupUsage(m map[string]interface{}) (result *ResourceGroupUsage, err error) {
	obj := new(ResourceGroupUsage)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.ResourceGroupID, err = core.UnmarshalString(m, "resource_group_id")
	if err != nil {
		return
	}
	obj.PricingCountry, err = core.UnmarshalString(m, "pricing_country")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.Month, err = core.UnmarshalString(m, "month")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalResourceSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceGroupUsageSlice unmarshals a slice of ResourceGroupUsage instances from the specified list of maps.
func UnmarshalResourceGroupUsageSlice(s []interface{}) (slice []ResourceGroupUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceGroupUsage'")
			return
		}
		obj, e := UnmarshalResourceGroupUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceGroupUsageAsProperty unmarshals an instance of ResourceGroupUsage that is stored as a property
// within the specified map.
func UnmarshalResourceGroupUsageAsProperty(m map[string]interface{}, propertyName string) (result *ResourceGroupUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceGroupUsage'", propertyName)
			return
		}
		result, err = UnmarshalResourceGroupUsage(objMap)
	}
	return
}

// UnmarshalResourceGroupUsageSliceAsProperty unmarshals a slice of ResourceGroupUsage instances that are stored as a property
// within the specified map.
func UnmarshalResourceGroupUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceGroupUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceGroupUsage'", propertyName)
			return
		}
		slice, err = UnmarshalResourceGroupUsageSlice(vSlice)
	}
	return
}

// Subscription : Subscription information.
type Subscription struct {
	// ID of the subscription.
	SubscriptionID *string `json:"subscription_id" validate:"required"`

	// Charge agreement number of the subsciption.
	ChargeAgreementNumber *string `json:"charge_agreement_number" validate:"required"`

	// Type of the subscription.
	Type *string `json:"type" validate:"required"`

	// Credit available in the subscription for the month.
	SubscriptionAmount *float64 `json:"subscription_amount" validate:"required"`

	// Date from which the subscription was active.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// The date until which the subscription is active. This date does not apply to Pay-As-You-Go accounts.
	End *strfmt.DateTime `json:"end,omitempty"`

	// Total credit available in the subscription.
	CreditsTotal *float64 `json:"credits_total" validate:"required"`

	// Separate periods of time within the overall subscription term. Longer subscriptions might be divided into multiple
	// terms.
	Terms []Terms `json:"terms" validate:"required"`
}


// UnmarshalSubscription constructs an instance of Subscription from the specified map.
func UnmarshalSubscription(m map[string]interface{}) (result *Subscription, err error) {
	obj := new(Subscription)
	obj.SubscriptionID, err = core.UnmarshalString(m, "subscription_id")
	if err != nil {
		return
	}
	obj.ChargeAgreementNumber, err = core.UnmarshalString(m, "charge_agreement_number")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.SubscriptionAmount, err = core.UnmarshalFloat64(m, "subscription_amount")
	if err != nil {
		return
	}
	obj.Start, err = core.UnmarshalDateTime(m, "start")
	if err != nil {
		return
	}
	obj.End, err = core.UnmarshalDateTime(m, "end")
	if err != nil {
		return
	}
	obj.CreditsTotal, err = core.UnmarshalFloat64(m, "credits_total")
	if err != nil {
		return
	}
	obj.Terms, err = UnmarshalTermsSliceAsProperty(m, "terms")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalSubscriptionSlice unmarshals a slice of Subscription instances from the specified list of maps.
func UnmarshalSubscriptionSlice(s []interface{}) (slice []Subscription, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Subscription'")
			return
		}
		obj, e := UnmarshalSubscription(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalSubscriptionAsProperty unmarshals an instance of Subscription that is stored as a property
// within the specified map.
func UnmarshalSubscriptionAsProperty(m map[string]interface{}, propertyName string) (result *Subscription, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Subscription'", propertyName)
			return
		}
		result, err = UnmarshalSubscription(objMap)
	}
	return
}

// UnmarshalSubscriptionSliceAsProperty unmarshals a slice of Subscription instances that are stored as a property
// within the specified map.
func UnmarshalSubscriptionSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Subscription, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Subscription'", propertyName)
			return
		}
		slice, err = UnmarshalSubscriptionSlice(vSlice)
	}
	return
}

// SubscriptionInfo : Subscription information for an account.
type SubscriptionInfo struct {
	// Charges for usage that exceeds the available credit from subscriptions or offers.
	Overage *float64 `json:"overage,omitempty"`

	// List of subscriptions that applied to the month.
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}


// UnmarshalSubscriptionInfo constructs an instance of SubscriptionInfo from the specified map.
func UnmarshalSubscriptionInfo(m map[string]interface{}) (result *SubscriptionInfo, err error) {
	obj := new(SubscriptionInfo)
	obj.Overage, err = core.UnmarshalFloat64(m, "overage")
	if err != nil {
		return
	}
	obj.Subscriptions, err = UnmarshalSubscriptionSliceAsProperty(m, "subscriptions")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalSubscriptionInfoSlice unmarshals a slice of SubscriptionInfo instances from the specified list of maps.
func UnmarshalSubscriptionInfoSlice(s []interface{}) (slice []SubscriptionInfo, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'SubscriptionInfo'")
			return
		}
		obj, e := UnmarshalSubscriptionInfo(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalSubscriptionInfoAsProperty unmarshals an instance of SubscriptionInfo that is stored as a property
// within the specified map.
func UnmarshalSubscriptionInfoAsProperty(m map[string]interface{}, propertyName string) (result *SubscriptionInfo, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'SubscriptionInfo'", propertyName)
			return
		}
		result, err = UnmarshalSubscriptionInfo(objMap)
	}
	return
}

// UnmarshalSubscriptionInfoSliceAsProperty unmarshals a slice of SubscriptionInfo instances that are stored as a property
// within the specified map.
func UnmarshalSubscriptionInfoSliceAsProperty(m map[string]interface{}, propertyName string) (slice []SubscriptionInfo, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'SubscriptionInfo'", propertyName)
			return
		}
		slice, err = UnmarshalSubscriptionInfoSlice(vSlice)
	}
	return
}

// SupportInfo : Account support information.
type SupportInfo struct {
	// Monthly support cost.
	Cost *float64 `json:"cost" validate:"required"`

	// Type of support.
	Type *string `json:"type" validate:"required"`

	// Additional support cost for the month.
	Overage *float64 `json:"overage" validate:"required"`
}


// UnmarshalSupportInfo constructs an instance of SupportInfo from the specified map.
func UnmarshalSupportInfo(m map[string]interface{}) (result *SupportInfo, err error) {
	obj := new(SupportInfo)
	obj.Cost, err = core.UnmarshalFloat64(m, "cost")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Overage, err = core.UnmarshalFloat64(m, "overage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalSupportInfoSlice unmarshals a slice of SupportInfo instances from the specified list of maps.
func UnmarshalSupportInfoSlice(s []interface{}) (slice []SupportInfo, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'SupportInfo'")
			return
		}
		obj, e := UnmarshalSupportInfo(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalSupportInfoAsProperty unmarshals an instance of SupportInfo that is stored as a property
// within the specified map.
func UnmarshalSupportInfoAsProperty(m map[string]interface{}, propertyName string) (result *SupportInfo, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'SupportInfo'", propertyName)
			return
		}
		result, err = UnmarshalSupportInfo(objMap)
	}
	return
}

// UnmarshalSupportInfoSliceAsProperty unmarshals a slice of SupportInfo instances that are stored as a property
// within the specified map.
func UnmarshalSupportInfoSliceAsProperty(m map[string]interface{}, propertyName string) (slice []SupportInfo, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'SupportInfo'", propertyName)
			return
		}
		slice, err = UnmarshalSupportInfoSlice(vSlice)
	}
	return
}

// Terms : Terms.
type Terms struct {
	// Start date of the term.
	Start *strfmt.DateTime `json:"start" validate:"required"`

	// End date of the term.
	End *strfmt.DateTime `json:"end" validate:"required"`

	// Credits.
	Credits *Credits `json:"credits" validate:"required"`
}


// UnmarshalTerms constructs an instance of Terms from the specified map.
func UnmarshalTerms(m map[string]interface{}) (result *Terms, err error) {
	obj := new(Terms)
	obj.Start, err = core.UnmarshalDateTime(m, "start")
	if err != nil {
		return
	}
	obj.End, err = core.UnmarshalDateTime(m, "end")
	if err != nil {
		return
	}
	obj.Credits, err = UnmarshalCreditsAsProperty(m, "credits")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalTermsSlice unmarshals a slice of Terms instances from the specified list of maps.
func UnmarshalTermsSlice(s []interface{}) (slice []Terms, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Terms'")
			return
		}
		obj, e := UnmarshalTerms(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalTermsAsProperty unmarshals an instance of Terms that is stored as a property
// within the specified map.
func UnmarshalTermsAsProperty(m map[string]interface{}, propertyName string) (result *Terms, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Terms'", propertyName)
			return
		}
		result, err = UnmarshalTerms(objMap)
	}
	return
}

// UnmarshalTermsSliceAsProperty unmarshals a slice of Terms instances that are stored as a property
// within the specified map.
func UnmarshalTermsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Terms, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Terms'", propertyName)
			return
		}
		slice, err = UnmarshalTermsSlice(vSlice)
	}
	return
}
