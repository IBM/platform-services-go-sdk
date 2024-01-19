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
 * IBM OpenAPI SDK Code Generator Version: 3.75.0-726bc7e3-20230713-221716
 */

// Package partnerbillingunitsv1 : Operations and models for the PartnerBillingUnitsV1 service
package partnerbillingunitsv1

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

// PartnerBillingUnitsV1 : Billing units for IBM Cloud partners
//
// API Version: 1.0.0
type PartnerBillingUnitsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://partner.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "partner_billing_units"

// PartnerBillingUnitsV1Options : Service options
type PartnerBillingUnitsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewPartnerBillingUnitsV1UsingExternalConfig : constructs an instance of PartnerBillingUnitsV1 with passed in options and external configuration.
func NewPartnerBillingUnitsV1UsingExternalConfig(options *PartnerBillingUnitsV1Options) (partnerBillingUnits *PartnerBillingUnitsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	partnerBillingUnits, err = NewPartnerBillingUnitsV1(options)
	if err != nil {
		return
	}

	err = partnerBillingUnits.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = partnerBillingUnits.Service.SetServiceURL(options.URL)
	}
	return
}

// NewPartnerBillingUnitsV1 : constructs an instance of PartnerBillingUnitsV1 with passed in options.
func NewPartnerBillingUnitsV1(options *PartnerBillingUnitsV1Options) (service *PartnerBillingUnitsV1, err error) {
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

	service = &PartnerBillingUnitsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "partnerBillingUnits" suitable for processing requests.
func (partnerBillingUnits *PartnerBillingUnitsV1) Clone() *PartnerBillingUnitsV1 {
	if core.IsNil(partnerBillingUnits) {
		return nil
	}
	clone := *partnerBillingUnits
	clone.Service = partnerBillingUnits.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (partnerBillingUnits *PartnerBillingUnitsV1) SetServiceURL(url string) error {
	return partnerBillingUnits.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (partnerBillingUnits *PartnerBillingUnitsV1) GetServiceURL() string {
	return partnerBillingUnits.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (partnerBillingUnits *PartnerBillingUnitsV1) SetDefaultHeaders(headers http.Header) {
	partnerBillingUnits.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (partnerBillingUnits *PartnerBillingUnitsV1) SetEnableGzipCompression(enableGzip bool) {
	partnerBillingUnits.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (partnerBillingUnits *PartnerBillingUnitsV1) GetEnableGzipCompression() bool {
	return partnerBillingUnits.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (partnerBillingUnits *PartnerBillingUnitsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	partnerBillingUnits.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (partnerBillingUnits *PartnerBillingUnitsV1) DisableRetries() {
	partnerBillingUnits.Service.DisableRetries()
}

// GetBillingOptions : Get customers billing options
// Returns the billing options for the requested customer for a given month.
func (partnerBillingUnits *PartnerBillingUnitsV1) GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions) (result *BillingOptionsSummary, response *core.DetailedResponse, err error) {
	return partnerBillingUnits.GetBillingOptionsWithContext(context.Background(), getBillingOptionsOptions)
}

// GetBillingOptionsWithContext is an alternate form of the GetBillingOptions method which supports a Context parameter
func (partnerBillingUnits *PartnerBillingUnitsV1) GetBillingOptionsWithContext(ctx context.Context, getBillingOptionsOptions *GetBillingOptionsOptions) (result *BillingOptionsSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBillingOptionsOptions, "getBillingOptionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBillingOptionsOptions, "getBillingOptionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerBillingUnits.Service.Options.URL, `/v1/billing-options`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBillingOptionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_billing_units", "V1", "GetBillingOptions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getBillingOptionsOptions.PartnerID))
	if getBillingOptionsOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getBillingOptionsOptions.CustomerID))
	}
	if getBillingOptionsOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getBillingOptionsOptions.ResellerID))
	}
	if getBillingOptionsOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getBillingOptionsOptions.Date))
	}
	if getBillingOptionsOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getBillingOptionsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalBillingOptionsSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCreditPoolsReport : Get subscription burn-down report
// Returns the subscription or commitment burn-down reports for the end customers for a given month.
func (partnerBillingUnits *PartnerBillingUnitsV1) GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions) (result *CreditPoolsReportSummary, response *core.DetailedResponse, err error) {
	return partnerBillingUnits.GetCreditPoolsReportWithContext(context.Background(), getCreditPoolsReportOptions)
}

// GetCreditPoolsReportWithContext is an alternate form of the GetCreditPoolsReport method which supports a Context parameter
func (partnerBillingUnits *PartnerBillingUnitsV1) GetCreditPoolsReportWithContext(ctx context.Context, getCreditPoolsReportOptions *GetCreditPoolsReportOptions) (result *CreditPoolsReportSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCreditPoolsReportOptions, "getCreditPoolsReportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCreditPoolsReportOptions, "getCreditPoolsReportOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = partnerBillingUnits.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(partnerBillingUnits.Service.Options.URL, `/v1/credit-pools`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCreditPoolsReportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("partner_billing_units", "V1", "GetCreditPoolsReport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("partner_id", fmt.Sprint(*getCreditPoolsReportOptions.PartnerID))
	if getCreditPoolsReportOptions.CustomerID != nil {
		builder.AddQuery("customer_id", fmt.Sprint(*getCreditPoolsReportOptions.CustomerID))
	}
	if getCreditPoolsReportOptions.ResellerID != nil {
		builder.AddQuery("reseller_id", fmt.Sprint(*getCreditPoolsReportOptions.ResellerID))
	}
	if getCreditPoolsReportOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getCreditPoolsReportOptions.Date))
	}
	if getCreditPoolsReportOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getCreditPoolsReportOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = partnerBillingUnits.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreditPoolsReportSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// BillingOptionsSummaryFirst : The link to the first page of the search query.
type BillingOptionsSummaryFirst struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalBillingOptionsSummaryFirst unmarshals an instance of BillingOptionsSummaryFirst from the specified map of raw messages.
func UnmarshalBillingOptionsSummaryFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOptionsSummaryFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingOptionsSummaryNext : The link to the next page of the search query.
type BillingOptionsSummaryNext struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the next page.
	Offset *string `json:"offset,omitempty"`
}

// UnmarshalBillingOptionsSummaryNext unmarshals an instance of BillingOptionsSummaryNext from the specified map of raw messages.
func UnmarshalBillingOptionsSummaryNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOptionsSummaryNext)
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

// BillingOption : Billing options report for the end customers.
type BillingOption struct {
	// The ID of the billing option.
	ID *string `json:"id,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Account ID of the customer.
	CustomerID *string `json:"customer_id,omitempty"`

	// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
	CustomerType *string `json:"customer_type,omitempty"`

	// A user-defined name for the customer.
	CustomerName *string `json:"customer_name,omitempty"`

	// ID of the reseller in the heirarchy of the requested customer.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Name of the reseller in the heirarchy of the requested customer.
	ResellerName *string `json:"reseller_name,omitempty"`

	// The billing month for which the burn-down report is requested. Format is yyyy-mm. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// Errors in the billing.
	Errors []map[string]interface{} `json:"errors,omitempty"`

	// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
	Type *string `json:"type,omitempty"`

	// The start date of billing option.
	StartDate *strfmt.DateTime `json:"start_date,omitempty"`

	// The end date of billing option.
	EndDate *strfmt.DateTime `json:"end_date,omitempty"`

	// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
	State *string `json:"state,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The payment method for support.
	PaymentInstrument map[string]interface{} `json:"payment_instrument,omitempty"`

	// Part number of the offering.
	PartNumber *string `json:"part_number,omitempty"`

	// ID of the catalog containing this offering.
	CatalogID *string `json:"catalog_id,omitempty"`

	// ID of the order containing this offering.
	OrderID *string `json:"order_id,omitempty"`

	// PO Number of the offering.
	PoNumber *string `json:"po_number,omitempty"`

	// Subscription model.
	SubscriptionModel *string `json:"subscription_model,omitempty"`

	// The duration of the billing options in months.
	DurationInMonths *int64 `json:"duration_in_months,omitempty"`

	// Amount billed monthly for this offering.
	MonthlyAmount *float64 `json:"monthly_amount,omitempty"`

	// The support billing system.
	BillingSystem map[string]interface{} `json:"billing_system,omitempty"`

	// The country code for the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`
}

// Constants associated with the BillingOption.CustomerType property.
// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
const (
	BillingOptionCustomerTypeAccountConst      = "ACCOUNT"
	BillingOptionCustomerTypeAccountGroupConst = "ACCOUNT_GROUP"
	BillingOptionCustomerTypeEnterpriseConst   = "ENTERPRISE"
)

// Constants associated with the BillingOption.Type property.
// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
const (
	BillingOptionTypeOfferConst        = "OFFER"
	BillingOptionTypeSubscriptionConst = "SUBSCRIPTION"
)

// Constants associated with the BillingOption.State property.
// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
const (
	BillingOptionStateActiveConst    = "ACTIVE"
	BillingOptionStateCanceledConst  = "CANCELED"
	BillingOptionStateSuspendedConst = "SUSPENDED"
)

// Constants associated with the BillingOption.Category property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
const (
	BillingOptionCategoryPlatformConst = "PLATFORM"
	BillingOptionCategoryServiceConst  = "SERVICE"
	BillingOptionCategorySupportConst  = "SUPPORT"
)

// UnmarshalBillingOption unmarshals an instance of BillingOption from the specified map of raw messages.
func UnmarshalBillingOption(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOption)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_type", &obj.CustomerType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_name", &obj.CustomerName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_id", &obj.ResellerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_name", &obj.ResellerName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "month", &obj.Month)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "payment_instrument", &obj.PaymentInstrument)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_number", &obj.PartNumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "order_id", &obj.OrderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "po_number", &obj.PoNumber)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subscription_model", &obj.SubscriptionModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "duration_in_months", &obj.DurationInMonths)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monthly_amount", &obj.MonthlyAmount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_system", &obj.BillingSystem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "country_code", &obj.CountryCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency_code", &obj.CurrencyCode)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// BillingOptionsSummary : The billing options report for the customer.
type BillingOptionsSummary struct {
	// The max number of reports in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The link to the first page of the search query.
	First *BillingOptionsSummaryFirst `json:"first,omitempty"`

	// The link to the next page of the search query.
	Next *BillingOptionsSummaryNext `json:"next,omitempty"`

	// Aggregated usage report of all requested partners.
	Resources []BillingOption `json:"resources,omitempty"`
}

// UnmarshalBillingOptionsSummary unmarshals an instance of BillingOptionsSummary from the specified map of raw messages.
func UnmarshalBillingOptionsSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(BillingOptionsSummary)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalBillingOptionsSummaryFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalBillingOptionsSummaryNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalBillingOption)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsReportSummaryFirst : The link to the first page of the search query.
type CreditPoolsReportSummaryFirst struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`
}

// UnmarshalCreditPoolsReportSummaryFirst unmarshals an instance of CreditPoolsReportSummaryFirst from the specified map of raw messages.
func UnmarshalCreditPoolsReportSummaryFirst(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReportSummaryFirst)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsReportSummaryNext : The link to the next page of the search query.
type CreditPoolsReportSummaryNext struct {
	// A link to a page of query results.
	Href *string `json:"href,omitempty"`

	// The value of the `_start` query parameter to fetch the next page.
	Offset *string `json:"offset,omitempty"`
}

// UnmarshalCreditPoolsReportSummaryNext unmarshals an instance of CreditPoolsReportSummaryNext from the specified map of raw messages.
func UnmarshalCreditPoolsReportSummaryNext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReportSummaryNext)
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

// CreditPoolsReport : Aggregated subscription burn-down report for the end customers.
type CreditPoolsReport struct {
	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE` and `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// Account ID of the customer.
	CustomerID *string `json:"customer_id,omitempty"`

	// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
	CustomerType *string `json:"customer_type,omitempty"`

	// A user-defined name for the customer.
	CustomerName *string `json:"customer_name,omitempty"`

	// ID of the reseller in the heirarchy of the requested customer.
	ResellerID *string `json:"reseller_id,omitempty"`

	// Name of the reseller in the heirarchy of the requested customer.
	ResellerName *string `json:"reseller_name,omitempty"`

	// The billing month for which the burn-down report is requested. Format is yyyy-mm. Defaults to current month.
	Month *string `json:"month,omitempty"`

	// The currency code of the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// A list of active subscription terms available within a credit.
	TermCredits []TermCredits `json:"term_credits,omitempty"`

	// Overage that was generated on the credit pool.
	Overage *Overage `json:"overage,omitempty"`
}

// Constants associated with the CreditPoolsReport.Type property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE` and `SUPPORT`.
const (
	CreditPoolsReportTypePlatformConst = "PLATFORM"
	CreditPoolsReportTypeServiceConst  = "SERVICE"
	CreditPoolsReportTypeSupportConst  = "SUPPORT"
)

// Constants associated with the CreditPoolsReport.CustomerType property.
// The customer type. The valid values are `ENTERPRISE`, `ACCOUNT`, and `ACCOUNT_GROUP`.
const (
	CreditPoolsReportCustomerTypeAccountConst      = "ACCOUNT"
	CreditPoolsReportCustomerTypeAccountGroupConst = "ACCOUNT_GROUP"
	CreditPoolsReportCustomerTypeEnterpriseConst   = "ENTERPRISE"
)

// UnmarshalCreditPoolsReport unmarshals an instance of CreditPoolsReport from the specified map of raw messages.
func UnmarshalCreditPoolsReport(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReport)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_unit_id", &obj.BillingUnitID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_id", &obj.CustomerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_type", &obj.CustomerType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customer_name", &obj.CustomerName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_id", &obj.ResellerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "reseller_name", &obj.ResellerName)
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
	err = core.UnmarshalModel(m, "term_credits", &obj.TermCredits, UnmarshalTermCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "overage", &obj.Overage, UnmarshalOverage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreditPoolsReportSummary : The aggregated credit pools report.
type CreditPoolsReportSummary struct {
	// The max number of reports in the response.
	Limit *int64 `json:"limit,omitempty"`

	// The link to the first page of the search query.
	First *CreditPoolsReportSummaryFirst `json:"first,omitempty"`

	// The link to the next page of the search query.
	Next *CreditPoolsReportSummaryNext `json:"next,omitempty"`

	// Aggregated usage report of all requested partners.
	Resources []CreditPoolsReport `json:"resources,omitempty"`
}

// UnmarshalCreditPoolsReportSummary unmarshals an instance of CreditPoolsReportSummary from the specified map of raw messages.
func UnmarshalCreditPoolsReportSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreditPoolsReportSummary)
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalCreditPoolsReportSummaryFirst)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalCreditPoolsReportSummaryNext)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalCreditPoolsReport)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetBillingOptionsOptions : The GetBillingOptions options.
type GetBillingOptionsOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Enterprise ID of the customer for which the report is requested.
	CustomerID *string `json:"customer_id,omitempty"`

	// Enterprise ID of the reseller for which the report is requested.
	ResellerID *string `json:"reseller_id,omitempty"`

	// The billing month for which the usage report is requested. Format is yyyy-mm. Defaults to current month.
	Date *string `json:"date,omitempty"`

	// Number of usage records returned. The default value is 30. Maximum value is 200.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBillingOptionsOptions : Instantiate GetBillingOptionsOptions
func (*PartnerBillingUnitsV1) NewGetBillingOptionsOptions(partnerID string, billingMonth string) *GetBillingOptionsOptions {
	return &GetBillingOptionsOptions{
		PartnerID: core.StringPtr(partnerID),
		Date:      core.StringPtr(billingMonth),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetBillingOptionsOptions) SetPartnerID(partnerID string) *GetBillingOptionsOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetBillingOptionsOptions) SetCustomerID(customerID string) *GetBillingOptionsOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetBillingOptionsOptions) SetResellerID(resellerID string) *GetBillingOptionsOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetDate : Allow user to set Date
func (_options *GetBillingOptionsOptions) SetDate(date string) *GetBillingOptionsOptions {
	_options.Date = core.StringPtr(date)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetBillingOptionsOptions) SetLimit(limit int64) *GetBillingOptionsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingOptionsOptions) SetHeaders(param map[string]string) *GetBillingOptionsOptions {
	options.Headers = param
	return options
}

// GetCreditPoolsReportOptions : The GetCreditPoolsReport options.
type GetCreditPoolsReportOptions struct {
	// Enterprise ID of the distributor or reseller for which the report is requested.
	PartnerID *string `json:"partner_id" validate:"required"`

	// Enterprise ID of the customer for which the report is requested.
	CustomerID *string `json:"customer_id,omitempty"`

	// Enterprise ID of the reseller for which the report is requested.
	ResellerID *string `json:"reseller_id,omitempty"`

	// The billing month for which the usage report is requested. Format is yyyy-mm. Defaults to current month.
	Date *string `json:"date,omitempty"`

	// Number of usage records returned. The default value is 30. Maximum value is 200.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCreditPoolsReportOptions : Instantiate GetCreditPoolsReportOptions
func (*PartnerBillingUnitsV1) NewGetCreditPoolsReportOptions(partnerID string, billingMonth string) *GetCreditPoolsReportOptions {
	return &GetCreditPoolsReportOptions{
		PartnerID: core.StringPtr(partnerID),
		Date:      core.StringPtr(billingMonth),
	}
}

// SetPartnerID : Allow user to set PartnerID
func (_options *GetCreditPoolsReportOptions) SetPartnerID(partnerID string) *GetCreditPoolsReportOptions {
	_options.PartnerID = core.StringPtr(partnerID)
	return _options
}

// SetCustomerID : Allow user to set CustomerID
func (_options *GetCreditPoolsReportOptions) SetCustomerID(customerID string) *GetCreditPoolsReportOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetResellerID : Allow user to set ResellerID
func (_options *GetCreditPoolsReportOptions) SetResellerID(resellerID string) *GetCreditPoolsReportOptions {
	_options.ResellerID = core.StringPtr(resellerID)
	return _options
}

// SetDate : Allow user to set Date
func (_options *GetCreditPoolsReportOptions) SetDate(date string) *GetCreditPoolsReportOptions {
	_options.Date = core.StringPtr(date)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *GetCreditPoolsReportOptions) SetLimit(limit int64) *GetCreditPoolsReportOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCreditPoolsReportOptions) SetHeaders(param map[string]string) *GetCreditPoolsReportOptions {
	options.Headers = param
	return options
}

// Overage : Overage that was generated on the credit pool.
type Overage struct {
	// The number of credits used as overage.
	Cost *float64 `json:"cost,omitempty"`

	// A list of resources that generated overage.
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// UnmarshalOverage unmarshals an instance of Overage from the specified map of raw messages.
func UnmarshalOverage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Overage)
	err = core.UnmarshalPrimitive(m, "cost", &obj.Cost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TermCredits : The subscription term that is active in the requested month.
type TermCredits struct {
	// The ID of the billing option from which the subscription term is derived.
	BillingOptionID *string `json:"billing_option_id,omitempty"`

	// Billing option model.
	BillingOptionModel *string `json:"billing_option_model,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The start date of the term in ISO format.
	StartDate *strfmt.DateTime `json:"start_date,omitempty"`

	// The end date of the term in ISO format.
	EndDate *strfmt.DateTime `json:"end_date,omitempty"`

	// The total credit available in this term.
	TotalCredits *float64 `json:"total_credits,omitempty"`

	// The balance of available credit at the start of the current month.
	StartingBalance *float64 `json:"starting_balance,omitempty"`

	// The amount of credit used during the current month.
	UsedCredits *float64 `json:"used_credits,omitempty"`

	// The balance of remaining credit in the subscription term.
	CurrentBalance *float64 `json:"current_balance,omitempty"`

	// A list of resources that used credit during the month.
	Resources []map[string]interface{} `json:"resources,omitempty"`
}

// Constants associated with the TermCredits.Category property.
// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
const (
	TermCreditsCategoryPlatformConst = "PLATFORM"
	TermCreditsCategoryServiceConst  = "SERVICE"
	TermCreditsCategorySupportConst  = "SUPPORT"
)

// UnmarshalTermCredits unmarshals an instance of TermCredits from the specified map of raw messages.
func UnmarshalTermCredits(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TermCredits)
	err = core.UnmarshalPrimitive(m, "billing_option_id", &obj.BillingOptionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_option_model", &obj.BillingOptionModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "category", &obj.Category)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "start_date", &obj.StartDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "end_date", &obj.EndDate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_credits", &obj.TotalCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "starting_balance", &obj.StartingBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "used_credits", &obj.UsedCredits)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current_balance", &obj.CurrentBalance)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resources", &obj.Resources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
