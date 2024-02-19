/**
 * (C) Copyright IBM Corp. 2024.
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
 * IBM OpenAPI SDK Code Generator Version: 3.85.0-75c38f8f-20240206-210220
 */

// Package partnerusagereportsv1 : Operations and models for the PartnerUsageReportsV1 service
package partnerusagereportsv1

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

// PartnerUsageReportsV1 : Usage reports for IBM Cloud partner entities
//
// API Version: 1.0.0
type PartnerUsageReportsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://partner.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "partner_usage_reports"

// PartnerUsageReportsV1Options : Service options
type PartnerUsageReportsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPartnerUsageReportsV1UsingExternalConfig : constructs an instance of PartnerUsageReportsV1 with passed in options and external configuration.
func NewPartnerUsageReportsV1UsingExternalConfig(options *PartnerUsageReportsV1Options) (partnerUsageReports *PartnerUsageReportsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	partnerUsageReports, err = NewPartnerUsageReportsV1(options)
	if err != nil {
		return
	}

	err = partnerUsageReports.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = partnerUsageReports.Service.SetServiceURL(options.URL)
	}
	return
}

// NewPartnerUsageReportsV1 : constructs an instance of PartnerUsageReportsV1 with passed in options.
func NewPartnerUsageReportsV1(options *PartnerUsageReportsV1Options) (service *PartnerUsageReportsV1, err error) {
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

	service = &PartnerUsageReportsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "partnerUsageReports" suitable for processing requests.
func (partnerUsageReports *PartnerUsageReportsV1) Clone() *PartnerUsageReportsV1 {
	if core.IsNil(partnerUsageReports) {
		return nil
	}
	clone := *partnerUsageReports
	clone.Service = partnerUsageReports.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (partnerUsageReports *PartnerUsageReportsV1) SetServiceURL(url string) error {
	return partnerUsageReports.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (partnerUsageReports *PartnerUsageReportsV1) GetServiceURL() string {
	return partnerUsageReports.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (partnerUsageReports *PartnerUsageReportsV1) SetDefaultHeaders(headers http.Header) {
	partnerUsageReports.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (partnerUsageReports *PartnerUsageReportsV1) SetEnableGzipCompression(enableGzip bool) {
	partnerUsageReports.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (partnerUsageReports *PartnerUsageReportsV1) GetEnableGzipCompression() bool {
	return partnerUsageReports.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (partnerUsageReports *PartnerUsageReportsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	partnerUsageReports.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (partnerUsageReports *PartnerUsageReportsV1) DisableRetries() {
	partnerUsageReports.Service.DisableRetries()
}

// GetResourceUsageReport : Get partner resource usage report
// Returns the summary for the partner for a given month. Partner billing managers are authorized to access this report.
func (partnerUsageReports *PartnerUsageReportsV1) GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *PartnerUsageReportSummary, response *core.DetailedResponse, err error) {
	return partnerUsageReports.GetResourceUsageReportWithContext(context.Background(), getResourceUsageReportOptions)
}

// GetResourceUsageReportWithContext is an alternate form of the GetResourceUsageReport method which supports a Context parameter
func (partnerUsageReports *PartnerUsageReportsV1) GetResourceUsageReportWithContext(ctx context.Context, getResourceUsageReportOptions *GetResourceUsageReportOptions) (result *PartnerUsageReportSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceUsageReportOptions, "getResourceUsageReportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceUsageReportOptions, "getResourceUsageReportOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerUsageReports.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerUsageReports.Service.Options.URL, `/v1/resource-usage-reports`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceUsageReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_usage_reports", "V1", "GetResourceUsageReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getResourceUsageReportOptions.PartnerID))
	if getResourceUsageReportOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getResourceUsageReportOptions.ResellerID))
	}
	if getResourceUsageReportOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getResourceUsageReportOptions.CustomerID))
	}
	if getResourceUsageReportOptions.Children != nil {
		builder.AddQuery("children", fmt.Sprint(*getResourceUsageReportOptions.Children))
	}
	if getResourceUsageReportOptions.Month != nil {
		builder.AddQuery("month", fmt.Sprint(*getResourceUsageReportOptions.Month))
	}
	if getResourceUsageReportOptions.Viewpoint != nil {
		builder.AddQuery("viewpoint", fmt.Sprint(*getResourceUsageReportOptions.Viewpoint))
	}
	if getResourceUsageReportOptions.Recurse != nil {
		builder.AddQuery("recurse", fmt.Sprint(*getResourceUsageReportOptions.Recurse))
	}
	if getResourceUsageReportOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getResourceUsageReportOptions.Limit))
	}
	if getResourceUsageReportOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getResourceUsageReportOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerUsageReports.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPartnerUsageReportSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetResourceUsageReportOptions : The GetResourceUsageReport options.
type GetResourceUsageReportOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Enterprise ID of the reseller for which the report is requested. This parameter cannot be used along with
	// `customer_id` query parameter.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Enterprise ID of the child customer for which the report is requested. This parameter cannot be used along with
	// `reseller_id` query parameter.
	CustomerID *string `json:"customer_id,omitempty"`

	// Get report rolled-up to the direct children of the requested entity. Defaults to false. This parameter cannot be
	// used along with `customer_id` query parameter.
	Children *bool `json:"children,omitempty"`

	// The billing month for which the usage report is requested. Format is `yyyy-mm`. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// Enables partner to view the cost of provisioned services as applicable at each level of the hierarchy. Defaults to
	// the type of the calling partner. The valid values are `DISTRIBUTOR`, `RESELLER` and `END_CUSTOMER`.
	Viewpoint *string `json:"viewpoint,omitempty"`

	// Get usage report rolled-up to the end customers of the requested entity. Defaults to false. This parameter cannot be
	// used along with `reseller_id` query parameter or `customer_id` query parameter.
	Recurse *bool `json:"recurse,omitempty"`

	// Number of usage records to be returned. The default value is 30. Maximum value is 200.
	Limit *int64 `json:"limit,omitempty"`

	// An opaque value representing the offset of the first item to be returned by a search query. If not specified, then
	// the first page of results is returned. To retrieve the next page of search results, use the 'offset' query parameter
	// value within the 'next.href' URL found within a prior search query response.
	Offset *string `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetResourceUsageReportOptions.Viewpoint property.
// Enables partner to view the cost of provisioned services as applicable at each level of the hierarchy. Defaults to
// the type of the calling partner. The valid values are `DISTRIBUTOR`, `RESELLER` and `END_CUSTOMER`.
const (
	GetResourceUsageReportOptionsViewpointDistributorConst = "DISTRIBUTOR"
	GetResourceUsageReportOptionsViewpointEndCustomerConst = "END_CUSTOMER"
	GetResourceUsageReportOptionsViewpointResellerConst    = "RESELLER"
)

// NewGetResourceUsageReportOptions : Instantiate GetResourceUsageReportOptions
func (*PartnerUsageReportsV1) NewGetResourceUsageReportOptions(partnerID string) *GetResourceUsageReportOptions {
	return &GetResourceUsageReportOptions{
		PartnerID: core.StringPtr(partnerID),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetResourceUsageReportOptions) SetPartnerID(partnerID string) *GetResourceUsageReportOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetResourceUsageReportOptions) SetResellerID(resellerID string) *GetResourceUsageReportOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetResourceUsageReportOptions) SetCustomerID(customerID string) *GetResourceUsageReportOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetChildren : Allow user to set Children
func (_options *GetResourceUsageReportOptions) SetChildren(children bool) *GetResourceUsageReportOptions {
	_options.Children = core.BoolPtr(children)
	return _options
}

// SetMonth : Allow user to set Month
func (_options *GetResourceUsageReportOptions) SetMonth(month string) *GetResourceUsageReportOptions {
	_options.Month = core.StringPtr(month)
	return _options
}

// SetViewpoint : Allow user to set Viewpoint
func (_options *GetResourceUsageReportOptions) SetViewpoint(viewpoint string) *GetResourceUsageReportOptions {
	_options.Viewpoint = core.StringPtr(viewpoint)
	return _options
}

// SetRecurse : Allow user to set Recurse
func (_options *GetResourceUsageReportOptions) SetRecurse(recurse bool) *GetResourceUsageReportOptions {
	_options.Recurse = core.BoolPtr(recurse)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetResourceUsageReportOptions) SetLimit(limit int64) *GetResourceUsageReportOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *GetResourceUsageReportOptions) SetOffset(offset string) *GetResourceUsageReportOptions {
	_options.Offset = core.StringPtr(offset)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceUsageReportOptions) SetHeaders(param map[string]string) *GetResourceUsageReportOptions {
	options.Headers = param
	return options
}

// MetricUsage : An object that represents a metric.
type MetricUsage struct {
	// The name of the metric.
	Metric *string `json:"metric" validate:"required"`

	// A unit to qualify the quantity.
	Unit *string `json:"unit" validate:"required"`

	// The aggregated value for the metric.
	Quantity *float64 `json:"quantity" validate:"required"`

	// The quantity that is used for calculating charges.
	RateableQuantity *float64 `json:"rateable_quantity" validate:"required"`

	// The cost that was incurred by the metric.
	Cost *float64 `json:"cost" validate:"required"`

	// The pre-discounted cost that was incurred by the metric.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// The price with which cost was calculated.
	Price []map[string]interface{} `json:"price,omitempty"`
}

// UnmarshalMetricUsage unmarshals an instance of MetricUsage from the specified map of raw messages.
func UnmarshalMetricUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MetricUsage)
	err = core.UnmarshalPrimitive(m, "metric", &obj.Metric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "unit", &obj.Unit)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummaryFirst : The link to the first page of the search query.
type PartnerUsageReportSummaryFirst struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalPartnerUsageReportSummaryFirst unmarshals an instance of PartnerUsageReportSummaryFirst from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummaryFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummaryFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummaryNext : The link to the next page of the search query.
type PartnerUsageReportSummaryNext struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the next page.
	Offset *string `json:"offset,omitempty"`
}

// UnmarshalPartnerUsageReportSummaryNext unmarshals an instance of PartnerUsageReportSummaryNext from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummaryNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummaryNext)
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

// PartnerUsageReport : Aggregated usage report of a partner.
type PartnerUsageReport struct {
	// The ID of the entity.
	EntityID *string `json:"entity_id,omitempty"`

	// The entity type.
	EntityType *string `json:"entity_type,omitempty"`

	// The Cloud Resource Name (CRN) of the entity towards which the resource usages were rolled up.
	EntityCRN *string `json:"entity_crn,omitempty"`

	// A user-defined name for the entity, such as the enterprise name or account group name.
	EntityName *string `json:"entity_name,omitempty"`

	// Role of the `entity_id` for which the usage report is fetched.
	EntityPartnerType *string `json:"entity_partner_type,omitempty"`

	// Enables partner to view the cost of provisioned services as applicable at each level of the hierarchy.
	Viewpoint *string `json:"viewpoint,omitempty"`

	// The billing month for which the usage report is requested. Format is yyyy-mm.
	Month *string `json:"month,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The country code of the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// Billable charges that are aggregated from all entities in the report.
	BillableCost *float64 `json:"billable_cost,omitempty"`

	// Aggregated billable charges before discounts.
	BillableRatedCost *float64 `json:"billable_rated_cost,omitempty"`

	// Non-billable charges that are aggregated from all entities in the report.
	NonBillableCost *float64 `json:"non_billable_cost,omitempty"`

	// Aggregated non-billable charges before discounts.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost,omitempty"`

	Resources []ResourceUsage `json:"resources,omitempty"`
}

// UnmarshalPartnerUsageReport unmarshals an instance of PartnerUsageReport from the specified map of raw messages.
func UnmarshalPartnerUsageReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReport)
	err = core.UnmarshalPrimitive(m, "entity_id", &obj.EntityID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_type", &obj.EntityType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_crn", &obj.EntityCRN)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_name", &obj.EntityName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "entity_partner_type", &obj.EntityPartnerType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "viewpoint", &obj.Viewpoint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
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
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PartnerUsageReportSummary : The aggregated partner usage report.
type PartnerUsageReportSummary struct {
	// The maximum number of usage records in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The link to the first page of the search query.
	First *PartnerUsageReportSummaryFirst `json:"first,omitempty"`

	// The link to the next page of the search query.
	Next *PartnerUsageReportSummaryNext `json:"next,omitempty"`

	// Aggregated usage report of all requested partners.
	Reports []PartnerUsageReport `json:"reports,omitempty"`
}

// UnmarshalPartnerUsageReportSummary unmarshals an instance of PartnerUsageReportSummary from the specified map of raw messages.
func UnmarshalPartnerUsageReportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PartnerUsageReportSummary)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalPartnerUsageReportSummaryFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalPartnerUsageReportSummaryNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reports", &obj.Reports, UnmarshalPartnerUsageReport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *PartnerUsageReportSummary) GetNextOffset() (*string, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	return resp.Next.Offset, nil
}

// PlanUsage : Aggregated values for the plan.
type PlanUsage struct {
	// The ID of the plan.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region for the plan.
	PricingRegion *string `json:"pricing_region,omitempty"`

	// The pricing plan with which the usage was rated.
	PricingPlanID *string `json:"pricing_plan_id,omitempty"`

	// Whether the plan charges are billed to the customer.
	Billable *bool `json:"billable" validate:"required"`

	// The total cost that was incurred by the plan.
	Cost *float64 `json:"cost" validate:"required"`

	// The total pre-discounted cost that was incurred by the plan.
	RatedCost *float64 `json:"rated_cost" validate:"required"`

	// All of the metrics in the plan.
	Usage []MetricUsage `json:"usage" validate:"required"`
}

// UnmarshalPlanUsage unmarshals an instance of PlanUsage from the specified map of raw messages.
func UnmarshalPlanUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PlanUsage)
	err = core.UnmarshalPrimitive(m, "plan_id", &obj.PlanID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_region", &obj.PricingRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pricing_plan_id", &obj.PricingPlanID)
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
	err = core.UnmarshalModel(m, "usage", &obj.Usage, UnmarshalMetricUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceUsage : A container for all the plans in the resource.
type ResourceUsage struct {
	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The name of the resource.
	ResourceName *string `json:"resource_name,omitempty"`

	// The billable charges for the partner.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The pre-discounted billable charges for the partner.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// The non-billable charges for the partner.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// The pre-discounted, non-billable charges for the partner.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// All of the plans in the resource.
	Plans []PlanUsage `json:"plans" validate:"required"`
}

// UnmarshalResourceUsage unmarshals an instance of ResourceUsage from the specified map of raw messages.
func UnmarshalResourceUsage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceUsage)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_name", &obj.ResourceName)
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
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlanUsage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetResourceUsageReportPager can be used to simplify the use of the "GetResourceUsageReport" method.
type GetResourceUsageReportPager struct {
	hasNext     bool
	options     *GetResourceUsageReportOptions
	client      *PartnerUsageReportsV1
	pageContext struct {
		next *string
	}
}

// NewGetResourceUsageReportPager returns a new GetResourceUsageReportPager instance.
func (partnerUsageReports *PartnerUsageReportsV1) NewGetResourceUsageReportPager(options *GetResourceUsageReportOptions) (pager *GetResourceUsageReportPager, err error) {
	if options.Offset != nil && *options.Offset != "" {
		err = fmt.Errorf("the 'options.Offset' field should not be set")
		return
	}

	var optionsCopy GetResourceUsageReportOptions = *options
	pager = &GetResourceUsageReportPager{
		hasNext: true,
		options: &optionsCopy,
		client:  partnerUsageReports,
	}
	return
}

// HasNext returns true if there are potentially more results to be retrieved.
func (pager *GetResourceUsageReportPager) HasNext() bool {
	return pager.hasNext
}

// GetNextWithContext returns the next page of results using the specified Context.
func (pager *GetResourceUsageReportPager) GetNextWithContext(ctx context.Context) (page []PartnerUsageReport, err error) {
	if !pager.HasNext() {
		return nil, fmt.Errorf("no more results available")
	}

	pager.options.Offset = pager.pageContext.next

	result, _, err := pager.client.GetResourceUsageReportWithContext(ctx, pager.options)
	if err != nil {
		return
	}

	var next *string
	if result.Next != nil {
		next = result.Next.Offset
	}
	pager.pageContext.next = next
	pager.hasNext = (pager.pageContext.next != nil)
	page = result.Reports

	return
}

// GetAllWithContext returns all results by invoking GetNextWithContext() repeatedly
// until all pages of results have been retrieved.
func (pager *GetResourceUsageReportPager) GetAllWithContext(ctx context.Context) (allItems []PartnerUsageReport, err error) {
	for pager.HasNext() {
		var nextPage []PartnerUsageReport
		nextPage, err = pager.GetNextWithContext(ctx)
		if err != nil {
			return
		}
		allItems = append(allItems, nextPage...)
	}
	return
}

// GetNext invokes GetNextWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetNext() (page []PartnerUsageReport, err error) {
	return pager.GetNextWithContext(context.Background())
}

// GetAll invokes GetAllWithContext() using context.Background() as the Context parameter.
func (pager *GetResourceUsageReportPager) GetAll() (allItems []PartnerUsageReport, err error) {
	return pager.GetAllWithContext(context.Background())
}
