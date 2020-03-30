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

// Package enterpriseusagereportsv1 : Operations and models for the EnterpriseUsageReportsV1 service
package enterpriseusagereportsv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// EnterpriseUsageReportsV1 : Usage reports for IBM Cloud enterprise entities
//
// Version: 1.0.0-beta.1
type EnterpriseUsageReportsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://enterprise.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "enterprise_usage_reports"

// EnterpriseUsageReportsV1Options : Service options
type EnterpriseUsageReportsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewEnterpriseUsageReportsV1UsingExternalConfig : constructs an instance of EnterpriseUsageReportsV1 with passed in options and external configuration.
func NewEnterpriseUsageReportsV1UsingExternalConfig(options *EnterpriseUsageReportsV1Options) (enterpriseUsageReports *EnterpriseUsageReportsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	enterpriseUsageReports, err = NewEnterpriseUsageReportsV1(options)
	if err != nil {
		return
	}

	err = enterpriseUsageReports.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = enterpriseUsageReports.Service.SetServiceURL(options.URL)
	}
	return
}

// NewEnterpriseUsageReportsV1 : constructs an instance of EnterpriseUsageReportsV1 with passed in options.
func NewEnterpriseUsageReportsV1(options *EnterpriseUsageReportsV1Options) (service *EnterpriseUsageReportsV1, err error) {
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

	service = &EnterpriseUsageReportsV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (enterpriseUsageReports *EnterpriseUsageReportsV1) SetServiceURL(url string) error {
	return enterpriseUsageReports.Service.SetServiceURL(url)
}

// ListResourceUsageReport : Get usage reports for enterprise entities
// Usage reports for entities in the IBM Cloud enterprise. These entities can be the enterprise, an account group, or an
// account.
func (enterpriseUsageReports *EnterpriseUsageReportsV1) ListResourceUsageReport(listResourceUsageReportOptions *ListResourceUsageReportOptions) (result *Reports, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listResourceUsageReportOptions, "listResourceUsageReportOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/resource-usage-reports"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseUsageReports.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourceUsageReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_usage_reports", "V1", "ListResourceUsageReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listResourceUsageReportOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*listResourceUsageReportOptions.EnterpriseID))
	}
	if listResourceUsageReportOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*listResourceUsageReportOptions.AccountGroupID))
	}
	if listResourceUsageReportOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listResourceUsageReportOptions.AccountID))
	}
	if listResourceUsageReportOptions.Children != nil {
		builder.AddQuery("children", fmt.Sprint(*listResourceUsageReportOptions.Children))
	}
	if listResourceUsageReportOptions.Month != nil {
		builder.AddQuery("month", fmt.Sprint(*listResourceUsageReportOptions.Month))
	}
	if listResourceUsageReportOptions.BillingUnitID != nil {
		builder.AddQuery("billing_unit_id", fmt.Sprint(*listResourceUsageReportOptions.BillingUnitID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseUsageReports.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalReports(m)
		response.Result = result
	}

	return
}

// ListResourceUsageReportOptions : The ListResourceUsageReport options.
type ListResourceUsageReportOptions struct {
	// The ID of the enterprise for which the reports are queried. This parameter cannot be used with the `account_id` or
	// `account_group_id` query parameters.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The ID of the account group for which the reports are queried. This parameter cannot be used with the `account_id`
	// or `enterprise_id` query parameters.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// The ID of the account for which the reports are queried. This parameter cannot be used with the `account_group_id`
	// or `enterprise_id` query parameters.
	AccountID *string `json:"account_id,omitempty"`

	// Returns the reports for the immediate child entities under the current account group or enterprise. This parameter
	// cannot be used with the `account_id` query parameter.
	Children *bool `json:"children,omitempty"`

	// The billing month for which the usage report is requested. The format is in yyyy-mm. Defaults to the month in which
	// the report is queried.
	Month *string `json:"month,omitempty"`

	// The ID of the billing unit by which to filter the reports.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourceUsageReportOptions : Instantiate ListResourceUsageReportOptions
func (*EnterpriseUsageReportsV1) NewListResourceUsageReportOptions() *ListResourceUsageReportOptions {
	return &ListResourceUsageReportOptions{}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *ListResourceUsageReportOptions) SetEnterpriseID(enterpriseID string) *ListResourceUsageReportOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *ListResourceUsageReportOptions) SetAccountGroupID(accountGroupID string) *ListResourceUsageReportOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *ListResourceUsageReportOptions) SetAccountID(accountID string) *ListResourceUsageReportOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetChildren : Allow user to set Children
func (options *ListResourceUsageReportOptions) SetChildren(children bool) *ListResourceUsageReportOptions {
	options.Children = core.BoolPtr(children)
	return options
}

// SetMonth : Allow user to set Month
func (options *ListResourceUsageReportOptions) SetMonth(month string) *ListResourceUsageReportOptions {
	options.Month = core.StringPtr(month)
	return options
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (options *ListResourceUsageReportOptions) SetBillingUnitID(billingUnitID string) *ListResourceUsageReportOptions {
	options.BillingUnitID = core.StringPtr(billingUnitID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourceUsageReportOptions) SetHeaders(param map[string]string) *ListResourceUsageReportOptions {
	options.Headers = param
	return options
}

// MetricUsageModel : An object that represents a metric.
type MetricUsageModel struct {
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
	Price []interface{} `json:"price,omitempty"`
}


// UnmarshalMetricUsageModel constructs an instance of MetricUsageModel from the specified map.
func UnmarshalMetricUsageModel(m map[string]interface{}) (result *MetricUsageModel, err error) {
	obj := new(MetricUsageModel)
	obj.Metric, err = core.UnmarshalString(m, "metric")
	if err != nil {
		return
	}
	obj.Unit, err = core.UnmarshalString(m, "unit")
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
	obj.RatedCost, err = core.UnmarshalFloat64(m, "rated_cost")
	if err != nil {
		return
	}
	obj.Price, err = core.UnmarshalAnySlice(m, "price")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalMetricUsageModelSlice unmarshals a slice of MetricUsageModel instances from the specified list of maps.
func UnmarshalMetricUsageModelSlice(s []interface{}) (slice []MetricUsageModel, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'MetricUsageModel'")
			return
		}
		obj, e := UnmarshalMetricUsageModel(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalMetricUsageModelAsProperty unmarshals an instance of MetricUsageModel that is stored as a property
// within the specified map.
func UnmarshalMetricUsageModelAsProperty(m map[string]interface{}, propertyName string) (result *MetricUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'MetricUsageModel'", propertyName)
			return
		}
		result, err = UnmarshalMetricUsageModel(objMap)
	}
	return
}

// UnmarshalMetricUsageModelSliceAsProperty unmarshals a slice of MetricUsageModel instances that are stored as a property
// within the specified map.
func UnmarshalMetricUsageModelSliceAsProperty(m map[string]interface{}, propertyName string) (slice []MetricUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'MetricUsageModel'", propertyName)
			return
		}
		slice, err = UnmarshalMetricUsageModelSlice(vSlice)
	}
	return
}

// PlanUsageModel : Aggregated values for the plan.
type PlanUsageModel struct {
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
	Usage []MetricUsageModel `json:"usage" validate:"required"`
}


// UnmarshalPlanUsageModel constructs an instance of PlanUsageModel from the specified map.
func UnmarshalPlanUsageModel(m map[string]interface{}) (result *PlanUsageModel, err error) {
	obj := new(PlanUsageModel)
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.PricingRegion, err = core.UnmarshalString(m, "pricing_region")
	if err != nil {
		return
	}
	obj.PricingPlanID, err = core.UnmarshalString(m, "pricing_plan_id")
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
	obj.RatedCost, err = core.UnmarshalFloat64(m, "rated_cost")
	if err != nil {
		return
	}
	obj.Usage, err = UnmarshalMetricUsageModelSliceAsProperty(m, "usage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPlanUsageModelSlice unmarshals a slice of PlanUsageModel instances from the specified list of maps.
func UnmarshalPlanUsageModelSlice(s []interface{}) (slice []PlanUsageModel, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PlanUsageModel'")
			return
		}
		obj, e := UnmarshalPlanUsageModel(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPlanUsageModelAsProperty unmarshals an instance of PlanUsageModel that is stored as a property
// within the specified map.
func UnmarshalPlanUsageModelAsProperty(m map[string]interface{}, propertyName string) (result *PlanUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'PlanUsageModel'", propertyName)
			return
		}
		result, err = UnmarshalPlanUsageModel(objMap)
	}
	return
}

// UnmarshalPlanUsageModelSliceAsProperty unmarshals a slice of PlanUsageModel instances that are stored as a property
// within the specified map.
func UnmarshalPlanUsageModelSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PlanUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PlanUsageModel'", propertyName)
			return
		}
		slice, err = UnmarshalPlanUsageModelSlice(vSlice)
	}
	return
}

// Reports : Resource Usage Reports API response.
type Reports struct {
	// The maximum number of reports in the response.
	Limit *float64 `json:"limit,omitempty"`

	// An object that contains the link to the first page of the search query.
	First *ReportsFirst `json:"first,omitempty"`

	// An object that contains the link to the first page of the search query.
	Next *ReportsNext `json:"next,omitempty"`

	// The list of usage reports.
	Reports []ResourceUsageReportModel `json:"reports,omitempty"`
}


// UnmarshalReports constructs an instance of Reports from the specified map.
func UnmarshalReports(m map[string]interface{}) (result *Reports, err error) {
	obj := new(Reports)
	obj.Limit, err = core.UnmarshalFloat64(m, "limit")
	if err != nil {
		return
	}
	obj.First, err = UnmarshalReportsFirstAsProperty(m, "first")
	if err != nil {
		return
	}
	obj.Next, err = UnmarshalReportsNextAsProperty(m, "next")
	if err != nil {
		return
	}
	obj.Reports, err = UnmarshalResourceUsageReportModelSliceAsProperty(m, "reports")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalReportsSlice unmarshals a slice of Reports instances from the specified list of maps.
func UnmarshalReportsSlice(s []interface{}) (slice []Reports, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Reports'")
			return
		}
		obj, e := UnmarshalReports(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalReportsAsProperty unmarshals an instance of Reports that is stored as a property
// within the specified map.
func UnmarshalReportsAsProperty(m map[string]interface{}, propertyName string) (result *Reports, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Reports'", propertyName)
			return
		}
		result, err = UnmarshalReports(objMap)
	}
	return
}

// UnmarshalReportsSliceAsProperty unmarshals a slice of Reports instances that are stored as a property
// within the specified map.
func UnmarshalReportsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Reports, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Reports'", propertyName)
			return
		}
		slice, err = UnmarshalReportsSlice(vSlice)
	}
	return
}

// ReportsFirst : An object that contains the link to the first page of the search query.
type ReportsFirst struct {
	// A link to the first page of the search query.
	Href *string `json:"href,omitempty"`
}


// UnmarshalReportsFirst constructs an instance of ReportsFirst from the specified map.
func UnmarshalReportsFirst(m map[string]interface{}) (result *ReportsFirst, err error) {
	obj := new(ReportsFirst)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalReportsFirstSlice unmarshals a slice of ReportsFirst instances from the specified list of maps.
func UnmarshalReportsFirstSlice(s []interface{}) (slice []ReportsFirst, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ReportsFirst'")
			return
		}
		obj, e := UnmarshalReportsFirst(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalReportsFirstAsProperty unmarshals an instance of ReportsFirst that is stored as a property
// within the specified map.
func UnmarshalReportsFirstAsProperty(m map[string]interface{}, propertyName string) (result *ReportsFirst, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ReportsFirst'", propertyName)
			return
		}
		result, err = UnmarshalReportsFirst(objMap)
	}
	return
}

// UnmarshalReportsFirstSliceAsProperty unmarshals a slice of ReportsFirst instances that are stored as a property
// within the specified map.
func UnmarshalReportsFirstSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ReportsFirst, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ReportsFirst'", propertyName)
			return
		}
		slice, err = UnmarshalReportsFirstSlice(vSlice)
	}
	return
}

// ReportsNext : An object that contains the link to the first page of the search query.
type ReportsNext struct {
	// A link to the first page of the search query.
	Href *string `json:"href,omitempty"`
}


// UnmarshalReportsNext constructs an instance of ReportsNext from the specified map.
func UnmarshalReportsNext(m map[string]interface{}) (result *ReportsNext, err error) {
	obj := new(ReportsNext)
	obj.Href, err = core.UnmarshalString(m, "href")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalReportsNextSlice unmarshals a slice of ReportsNext instances from the specified list of maps.
func UnmarshalReportsNextSlice(s []interface{}) (slice []ReportsNext, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ReportsNext'")
			return
		}
		obj, e := UnmarshalReportsNext(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalReportsNextAsProperty unmarshals an instance of ReportsNext that is stored as a property
// within the specified map.
func UnmarshalReportsNextAsProperty(m map[string]interface{}, propertyName string) (result *ReportsNext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ReportsNext'", propertyName)
			return
		}
		result, err = UnmarshalReportsNext(objMap)
	}
	return
}

// UnmarshalReportsNextSliceAsProperty unmarshals a slice of ReportsNext instances that are stored as a property
// within the specified map.
func UnmarshalReportsNextSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ReportsNext, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ReportsNext'", propertyName)
			return
		}
		slice, err = UnmarshalReportsNextSlice(vSlice)
	}
	return
}

// ResourceUsageModel : A container for all the plans in the resource.
type ResourceUsageModel struct {
	// The ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The billable charges for the account.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// The pre-discounted billable charges for the account.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// The non-billable charges for the account.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// The pre-discounted, non-billable charges for the account.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// All of the plans in the resource.
	Plans []PlanUsageModel `json:"plans" validate:"required"`
}


// UnmarshalResourceUsageModel constructs an instance of ResourceUsageModel from the specified map.
func UnmarshalResourceUsageModel(m map[string]interface{}) (result *ResourceUsageModel, err error) {
	obj := new(ResourceUsageModel)
	obj.ResourceID, err = core.UnmarshalString(m, "resource_id")
	if err != nil {
		return
	}
	obj.BillableCost, err = core.UnmarshalFloat64(m, "billable_cost")
	if err != nil {
		return
	}
	obj.BillableRatedCost, err = core.UnmarshalFloat64(m, "billable_rated_cost")
	if err != nil {
		return
	}
	obj.NonBillableCost, err = core.UnmarshalFloat64(m, "non_billable_cost")
	if err != nil {
		return
	}
	obj.NonBillableRatedCost, err = core.UnmarshalFloat64(m, "non_billable_rated_cost")
	if err != nil {
		return
	}
	obj.Plans, err = UnmarshalPlanUsageModelSliceAsProperty(m, "plans")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceUsageModelSlice unmarshals a slice of ResourceUsageModel instances from the specified list of maps.
func UnmarshalResourceUsageModelSlice(s []interface{}) (slice []ResourceUsageModel, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceUsageModel'")
			return
		}
		obj, e := UnmarshalResourceUsageModel(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceUsageModelAsProperty unmarshals an instance of ResourceUsageModel that is stored as a property
// within the specified map.
func UnmarshalResourceUsageModelAsProperty(m map[string]interface{}, propertyName string) (result *ResourceUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceUsageModel'", propertyName)
			return
		}
		result, err = UnmarshalResourceUsageModel(objMap)
	}
	return
}

// UnmarshalResourceUsageModelSliceAsProperty unmarshals a slice of ResourceUsageModel instances that are stored as a property
// within the specified map.
func UnmarshalResourceUsageModelSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceUsageModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceUsageModel'", propertyName)
			return
		}
		slice, err = UnmarshalResourceUsageModelSlice(vSlice)
	}
	return
}

// ResourceUsageReportModel : An object that represents a usage report.
type ResourceUsageReportModel struct {
	// The ID of the entity.
	EntityID *string `json:"entity_id" validate:"required"`

	// The entity type.
	EntityType *string `json:"entity_type" validate:"required"`

	// The Cloud Resource Name (CRN) of the entity towards which the resource usages were rolled up.
	EntityCrn *string `json:"entity_crn" validate:"required"`

	// A user-defined name for the entity, such as the enterprise name or account group name.
	EntityName *string `json:"entity_name" validate:"required"`

	// The ID of the billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// The CRN of the billing unit.
	BillingUnitCrn *string `json:"billing_unit_crn" validate:"required"`

	// The name of the billing unit.
	BillingUnitName *string `json:"billing_unit_name" validate:"required"`

	// The country code of the billing unit.
	CountryCode *string `json:"country_code" validate:"required"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code" validate:"required"`

	// Billing month.
	Month *string `json:"month" validate:"required"`

	// Billable charges that are aggregated from all entities in the report.
	BillableCost *float64 `json:"billable_cost" validate:"required"`

	// Non-billable charges that are aggregated from all entities in the report.
	NonBillableCost *float64 `json:"non_billable_cost" validate:"required"`

	// Aggregated billable charges before discounts.
	BillableRatedCost *float64 `json:"billable_rated_cost" validate:"required"`

	// Aggregated non-billable charges before discounts.
	NonBillableRatedCost *float64 `json:"non_billable_rated_cost" validate:"required"`

	// Details about all the resources that are included in the aggregated charges.
	Resources []ResourceUsageModel `json:"resources" validate:"required"`
}

// Constants associated with the ResourceUsageReportModel.EntityType property.
// The entity type.
const (
	ResourceUsageReportModel_EntityType_Account = "account"
	ResourceUsageReportModel_EntityType_AccountGroup = "account-group"
	ResourceUsageReportModel_EntityType_Enterprise = "enterprise"
)


// UnmarshalResourceUsageReportModel constructs an instance of ResourceUsageReportModel from the specified map.
func UnmarshalResourceUsageReportModel(m map[string]interface{}) (result *ResourceUsageReportModel, err error) {
	obj := new(ResourceUsageReportModel)
	obj.EntityID, err = core.UnmarshalString(m, "entity_id")
	if err != nil {
		return
	}
	obj.EntityType, err = core.UnmarshalString(m, "entity_type")
	if err != nil {
		return
	}
	obj.EntityCrn, err = core.UnmarshalString(m, "entity_crn")
	if err != nil {
		return
	}
	obj.EntityName, err = core.UnmarshalString(m, "entity_name")
	if err != nil {
		return
	}
	obj.BillingUnitID, err = core.UnmarshalString(m, "billing_unit_id")
	if err != nil {
		return
	}
	obj.BillingUnitCrn, err = core.UnmarshalString(m, "billing_unit_crn")
	if err != nil {
		return
	}
	obj.BillingUnitName, err = core.UnmarshalString(m, "billing_unit_name")
	if err != nil {
		return
	}
	obj.CountryCode, err = core.UnmarshalString(m, "country_code")
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
	obj.BillableCost, err = core.UnmarshalFloat64(m, "billable_cost")
	if err != nil {
		return
	}
	obj.NonBillableCost, err = core.UnmarshalFloat64(m, "non_billable_cost")
	if err != nil {
		return
	}
	obj.BillableRatedCost, err = core.UnmarshalFloat64(m, "billable_rated_cost")
	if err != nil {
		return
	}
	obj.NonBillableRatedCost, err = core.UnmarshalFloat64(m, "non_billable_rated_cost")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalResourceUsageModelSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceUsageReportModelSlice unmarshals a slice of ResourceUsageReportModel instances from the specified list of maps.
func UnmarshalResourceUsageReportModelSlice(s []interface{}) (slice []ResourceUsageReportModel, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceUsageReportModel'")
			return
		}
		obj, e := UnmarshalResourceUsageReportModel(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceUsageReportModelAsProperty unmarshals an instance of ResourceUsageReportModel that is stored as a property
// within the specified map.
func UnmarshalResourceUsageReportModelAsProperty(m map[string]interface{}, propertyName string) (result *ResourceUsageReportModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceUsageReportModel'", propertyName)
			return
		}
		result, err = UnmarshalResourceUsageReportModel(objMap)
	}
	return
}

// UnmarshalResourceUsageReportModelSliceAsProperty unmarshals a slice of ResourceUsageReportModel instances that are stored as a property
// within the specified map.
func UnmarshalResourceUsageReportModelSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceUsageReportModel, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceUsageReportModel'", propertyName)
			return
		}
		slice, err = UnmarshalResourceUsageReportModelSlice(vSlice)
	}
	return
}
