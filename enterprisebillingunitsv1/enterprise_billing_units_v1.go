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

// Package enterprisebillingunitsv1 : Operations and models for the EnterpriseBillingUnitsV1 service
package enterprisebillingunitsv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// EnterpriseBillingUnitsV1 : Billing units for IBM Cloud enterprises
//
// Version: 1.0.0
type EnterpriseBillingUnitsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://billing.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "enterprise_billing_units"

// EnterpriseBillingUnitsV1Options : Service options
type EnterpriseBillingUnitsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewEnterpriseBillingUnitsV1UsingExternalConfig : constructs an instance of EnterpriseBillingUnitsV1 with passed in options and external configuration.
func NewEnterpriseBillingUnitsV1UsingExternalConfig(options *EnterpriseBillingUnitsV1Options) (enterpriseBillingUnits *EnterpriseBillingUnitsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	enterpriseBillingUnits, err = NewEnterpriseBillingUnitsV1(options)
	if err != nil {
		return
	}

	err = enterpriseBillingUnits.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = enterpriseBillingUnits.Service.SetServiceURL(options.URL)
	}
	return
}

// NewEnterpriseBillingUnitsV1 : constructs an instance of EnterpriseBillingUnitsV1 with passed in options.
func NewEnterpriseBillingUnitsV1(options *EnterpriseBillingUnitsV1Options) (service *EnterpriseBillingUnitsV1, err error) {
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

	service = &EnterpriseBillingUnitsV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) SetServiceURL(url string) error {
	return enterpriseBillingUnits.Service.SetServiceURL(url)
}

// GetBillingOptionByQuery : Get billing options by query
// Return matching billing options if any exist. Show subscriptions and promotional offers that are available to a
// billing unit.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetBillingOptionByQuery(getBillingOptionByQueryOptions *GetBillingOptionByQueryOptions) (result *BillingOption, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBillingOptionByQueryOptions, "getBillingOptionByQueryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBillingOptionByQueryOptions, "getBillingOptionByQueryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/billing-options"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseBillingUnits.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBillingOptionByQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetBillingOptionByQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("billing_unit_id", fmt.Sprint(*getBillingOptionByQueryOptions.BillingUnitID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseBillingUnits.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalBillingOption(m)
		response.Result = result
	}

	return
}

// GetBillingUnitByID : Get billing unit by ID
// Return the billing unit information if it exists.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetBillingUnitByID(getBillingUnitByIdOptions *GetBillingUnitByIdOptions) (result *BillingUnit, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getBillingUnitByIdOptions, "getBillingUnitByIdOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getBillingUnitByIdOptions, "getBillingUnitByIdOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/billing-units"}
	pathParameters := []string{*getBillingUnitByIdOptions.BillingUnitID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseBillingUnits.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBillingUnitByIdOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetBillingUnitByID")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseBillingUnits.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalBillingUnit(m)
		response.Result = result
	}

	return
}

// GetBillingUnitByQuery : Get billing unit by query
// Return matching billing unit information if any exists. Omits internal properties and enterprise account ID from the
// billing unit.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetBillingUnitByQuery(getBillingUnitByQueryOptions *GetBillingUnitByQueryOptions) (result *ResponseBillingUnitsByQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getBillingUnitByQueryOptions, "getBillingUnitByQueryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/billing-units"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseBillingUnits.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getBillingUnitByQueryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetBillingUnitByQuery")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getBillingUnitByQueryOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getBillingUnitByQueryOptions.AccountID))
	}
	if getBillingUnitByQueryOptions.EnterpriseID != nil {
		builder.AddQuery("enterprise_id", fmt.Sprint(*getBillingUnitByQueryOptions.EnterpriseID))
	}
	if getBillingUnitByQueryOptions.AccountGroupID != nil {
		builder.AddQuery("account_group_id", fmt.Sprint(*getBillingUnitByQueryOptions.AccountGroupID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseBillingUnits.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResponseBillingUnitsByQuery(m)
		response.Result = result
	}

	return
}

// GetCreditPools : Get credit pools
// Get credit pools for a billing unit. Credit pools can be either platform or support credit pools. The platform credit
// pool contains credit from platform subscriptions and promotional offers. The support credit pool contains credit from
// support subscriptions.
func (enterpriseBillingUnits *EnterpriseBillingUnitsV1) GetCreditPools(getCreditPoolsOptions *GetCreditPoolsOptions) (result *ResponseCreditPoolsByQuery, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCreditPoolsOptions, "getCreditPoolsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCreditPoolsOptions, "getCreditPoolsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/credit-pools"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(enterpriseBillingUnits.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCreditPoolsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("enterprise_billing_units", "V1", "GetCreditPools")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("billing_unit_id", fmt.Sprint(*getCreditPoolsOptions.BillingUnitID))
	if getCreditPoolsOptions.Date != nil {
		builder.AddQuery("date", fmt.Sprint(*getCreditPoolsOptions.Date))
	}
	if getCreditPoolsOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*getCreditPoolsOptions.Type))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = enterpriseBillingUnits.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResponseCreditPoolsByQuery(m)
		response.Result = result
	}

	return
}

// BillingOption : A container for all the plans in the resource.
type BillingOption struct {
	// The ID of the billing option.
	ID *string `json:"id,omitempty"`

	// The ID of the billing unit that's associated with the billing option.
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// The start date of billing option.
	StartDate *string `json:"start_date,omitempty"`

	// The end date of billing option.
	EndDate *string `json:"end_date,omitempty"`

	// The state of the billing option. The valid values include `ACTIVE, `SUSPENDED`, and `CANCELED`.
	State *string `json:"state,omitempty"`

	// The type of billing option. The valid values are `SUBSCRIPTION` and `OFFER`.
	Type *string `json:"type,omitempty"`

	// The category of the billing option. The valid values are `PLATFORM`, `SERVICE`, and `SUPPORT`.
	Category *string `json:"category,omitempty"`

	// The payment method for support.
	PaymentInstrument interface{} `json:"payment_instrument,omitempty"`

	// The duration of the billing options in months.
	DurationInMonths *float64 `json:"duration_in_months,omitempty"`

	// The line item ID for support.
	LineItemID *float64 `json:"line_item_id,omitempty"`

	// The support billing system.
	BillingSystem interface{} `json:"billing_system,omitempty"`

	// The renewal code for support. This code denotes whether the subscription automatically renews, is assessed monthly,
	// and so on.
	RenewalModeCode *string `json:"renewal_mode_code,omitempty"`

	// The date when the billing option was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}


// UnmarshalBillingOption constructs an instance of BillingOption from the specified map.
func UnmarshalBillingOption(m map[string]interface{}) (result *BillingOption, err error) {
	obj := new(BillingOption)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.BillingUnitID, err = core.UnmarshalString(m, "billing_unit_id")
	if err != nil {
		return
	}
	obj.StartDate, err = core.UnmarshalString(m, "start_date")
	if err != nil {
		return
	}
	obj.EndDate, err = core.UnmarshalString(m, "end_date")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Category, err = core.UnmarshalString(m, "category")
	if err != nil {
		return
	}
	obj.PaymentInstrument, err = core.UnmarshalAny(m, "payment_instrument")
	if err != nil {
		return
	}
	obj.DurationInMonths, err = core.UnmarshalFloat64(m, "duration_in_months")
	if err != nil {
		return
	}
	obj.LineItemID, err = core.UnmarshalFloat64(m, "line_item_id")
	if err != nil {
		return
	}
	obj.BillingSystem, err = core.UnmarshalAny(m, "billing_system")
	if err != nil {
		return
	}
	obj.RenewalModeCode, err = core.UnmarshalString(m, "renewal_mode_code")
	if err != nil {
		return
	}
	obj.UpdatedAt, err = core.UnmarshalString(m, "updated_at")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalBillingOptionSlice unmarshals a slice of BillingOption instances from the specified list of maps.
func UnmarshalBillingOptionSlice(s []interface{}) (slice []BillingOption, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'BillingOption'")
			return
		}
		obj, e := UnmarshalBillingOption(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalBillingOptionAsProperty unmarshals an instance of BillingOption that is stored as a property
// within the specified map.
func UnmarshalBillingOptionAsProperty(m map[string]interface{}, propertyName string) (result *BillingOption, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'BillingOption'", propertyName)
			return
		}
		result, err = UnmarshalBillingOption(objMap)
	}
	return
}

// UnmarshalBillingOptionSliceAsProperty unmarshals a slice of BillingOption instances that are stored as a property
// within the specified map.
func UnmarshalBillingOptionSliceAsProperty(m map[string]interface{}, propertyName string) (slice []BillingOption, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'BillingOption'", propertyName)
			return
		}
		slice, err = UnmarshalBillingOptionSlice(vSlice)
	}
	return
}

// BillingUnit : Aggregated usage and charges for all the plans in the account.
type BillingUnit struct {
	// The ID of the billing unit, which is a globally unique identifier (GUID).
	ID *string `json:"id,omitempty"`

	// The Cloud Resource Name (CRN) of the billing unit, scoped to the enterprise account ID.
	Crn *string `json:"crn,omitempty"`

	// The name of the billing unit.
	Name *string `json:"name,omitempty"`

	// The ID of the enterprise to which the billing unit is associated.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The currency code for the billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The country code for the billing unit.
	CountryCode *string `json:"country_code,omitempty"`

	// A flag that indicates whether this billing unit is the primary billing mechanism for the enterprise.
	Master *bool `json:"master,omitempty"`

	// The creation date of the billing unit.
	CreatedAt *string `json:"created_at,omitempty"`
}


// UnmarshalBillingUnit constructs an instance of BillingUnit from the specified map.
func UnmarshalBillingUnit(m map[string]interface{}) (result *BillingUnit, err error) {
	obj := new(BillingUnit)
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.EnterpriseID, err = core.UnmarshalString(m, "enterprise_id")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.CountryCode, err = core.UnmarshalString(m, "country_code")
	if err != nil {
		return
	}
	obj.Master, err = core.UnmarshalBool(m, "master")
	if err != nil {
		return
	}
	obj.CreatedAt, err = core.UnmarshalString(m, "created_at")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalBillingUnitSlice unmarshals a slice of BillingUnit instances from the specified list of maps.
func UnmarshalBillingUnitSlice(s []interface{}) (slice []BillingUnit, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'BillingUnit'")
			return
		}
		obj, e := UnmarshalBillingUnit(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalBillingUnitAsProperty unmarshals an instance of BillingUnit that is stored as a property
// within the specified map.
func UnmarshalBillingUnitAsProperty(m map[string]interface{}, propertyName string) (result *BillingUnit, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'BillingUnit'", propertyName)
			return
		}
		result, err = UnmarshalBillingUnit(objMap)
	}
	return
}

// UnmarshalBillingUnitSliceAsProperty unmarshals a slice of BillingUnit instances that are stored as a property
// within the specified map.
func UnmarshalBillingUnitSliceAsProperty(m map[string]interface{}, propertyName string) (slice []BillingUnit, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'BillingUnit'", propertyName)
			return
		}
		slice, err = UnmarshalBillingUnitSlice(vSlice)
	}
	return
}

// CreditPoolOverage : Overage that was generated on the credit pool.
type CreditPoolOverage struct {
	// The number of credits used as overage.
	Cost *float64 `json:"cost,omitempty"`

	// A list of resources that generated overage.
	Resources []interface{} `json:"resources,omitempty"`
}


// UnmarshalCreditPoolOverage constructs an instance of CreditPoolOverage from the specified map.
func UnmarshalCreditPoolOverage(m map[string]interface{}) (result *CreditPoolOverage, err error) {
	obj := new(CreditPoolOverage)
	obj.Cost, err = core.UnmarshalFloat64(m, "cost")
	if err != nil {
		return
	}
	obj.Resources, err = core.UnmarshalAnySlice(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCreditPoolOverageSlice unmarshals a slice of CreditPoolOverage instances from the specified list of maps.
func UnmarshalCreditPoolOverageSlice(s []interface{}) (slice []CreditPoolOverage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'CreditPoolOverage'")
			return
		}
		obj, e := UnmarshalCreditPoolOverage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCreditPoolOverageAsProperty unmarshals an instance of CreditPoolOverage that is stored as a property
// within the specified map.
func UnmarshalCreditPoolOverageAsProperty(m map[string]interface{}, propertyName string) (result *CreditPoolOverage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'CreditPoolOverage'", propertyName)
			return
		}
		result, err = UnmarshalCreditPoolOverage(objMap)
	}
	return
}

// UnmarshalCreditPoolOverageSliceAsProperty unmarshals a slice of CreditPoolOverage instances that are stored as a property
// within the specified map.
func UnmarshalCreditPoolOverageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []CreditPoolOverage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'CreditPoolOverage'", propertyName)
			return
		}
		slice, err = UnmarshalCreditPoolOverageSlice(vSlice)
	}
	return
}

// CreditPool : The credit pool for a billing unit.
type CreditPool struct {
	// The type of credit, either `PLATFORM` or `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// The currency code of the associated billing unit.
	CurrencyCode *string `json:"currency_code,omitempty"`

	// The ID of the billing unit that's associated with the credit pool. This value is a globally unique identifier
	// (GUID).
	BillingUnitID *string `json:"billing_unit_id,omitempty"`

	// A list of active subscription terms available within a credit pool.
	TermCredits []TermCredits `json:"term_credits,omitempty"`

	// Overage that was generated on the credit pool.
	Overage *CreditPoolOverage `json:"overage,omitempty"`
}


// UnmarshalCreditPool constructs an instance of CreditPool from the specified map.
func UnmarshalCreditPool(m map[string]interface{}) (result *CreditPool, err error) {
	obj := new(CreditPool)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.CurrencyCode, err = core.UnmarshalString(m, "currency_code")
	if err != nil {
		return
	}
	obj.BillingUnitID, err = core.UnmarshalString(m, "billing_unit_id")
	if err != nil {
		return
	}
	obj.TermCredits, err = UnmarshalTermCreditsSliceAsProperty(m, "term_credits")
	if err != nil {
		return
	}
	obj.Overage, err = UnmarshalCreditPoolOverageAsProperty(m, "overage")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCreditPoolSlice unmarshals a slice of CreditPool instances from the specified list of maps.
func UnmarshalCreditPoolSlice(s []interface{}) (slice []CreditPool, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'CreditPool'")
			return
		}
		obj, e := UnmarshalCreditPool(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCreditPoolAsProperty unmarshals an instance of CreditPool that is stored as a property
// within the specified map.
func UnmarshalCreditPoolAsProperty(m map[string]interface{}, propertyName string) (result *CreditPool, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'CreditPool'", propertyName)
			return
		}
		result, err = UnmarshalCreditPool(objMap)
	}
	return
}

// UnmarshalCreditPoolSliceAsProperty unmarshals a slice of CreditPool instances that are stored as a property
// within the specified map.
func UnmarshalCreditPoolSliceAsProperty(m map[string]interface{}, propertyName string) (slice []CreditPool, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'CreditPool'", propertyName)
			return
		}
		slice, err = UnmarshalCreditPoolSlice(vSlice)
	}
	return
}

// GetBillingOptionByQueryOptions : The GetBillingOptionByQuery options.
type GetBillingOptionByQueryOptions struct {
	// The billing unit ID.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBillingOptionByQueryOptions : Instantiate GetBillingOptionByQueryOptions
func (*EnterpriseBillingUnitsV1) NewGetBillingOptionByQueryOptions(billingUnitID string) *GetBillingOptionByQueryOptions {
	return &GetBillingOptionByQueryOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (options *GetBillingOptionByQueryOptions) SetBillingUnitID(billingUnitID string) *GetBillingOptionByQueryOptions {
	options.BillingUnitID = core.StringPtr(billingUnitID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingOptionByQueryOptions) SetHeaders(param map[string]string) *GetBillingOptionByQueryOptions {
	options.Headers = param
	return options
}

// GetBillingUnitByIdOptions : The GetBillingUnitByID options.
type GetBillingUnitByIdOptions struct {
	// The ID of the requested billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBillingUnitByIdOptions : Instantiate GetBillingUnitByIdOptions
func (*EnterpriseBillingUnitsV1) NewGetBillingUnitByIdOptions(billingUnitID string) *GetBillingUnitByIdOptions {
	return &GetBillingUnitByIdOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (options *GetBillingUnitByIdOptions) SetBillingUnitID(billingUnitID string) *GetBillingUnitByIdOptions {
	options.BillingUnitID = core.StringPtr(billingUnitID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingUnitByIdOptions) SetHeaders(param map[string]string) *GetBillingUnitByIdOptions {
	options.Headers = param
	return options
}

// GetBillingUnitByQueryOptions : The GetBillingUnitByQuery options.
type GetBillingUnitByQueryOptions struct {
	// The enterprise account ID.
	AccountID *string `json:"account_id,omitempty"`

	// The enterprise ID.
	EnterpriseID *string `json:"enterprise_id,omitempty"`

	// The account group ID.
	AccountGroupID *string `json:"account_group_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetBillingUnitByQueryOptions : Instantiate GetBillingUnitByQueryOptions
func (*EnterpriseBillingUnitsV1) NewGetBillingUnitByQueryOptions() *GetBillingUnitByQueryOptions {
	return &GetBillingUnitByQueryOptions{}
}

// SetAccountID : Allow user to set AccountID
func (options *GetBillingUnitByQueryOptions) SetAccountID(accountID string) *GetBillingUnitByQueryOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *GetBillingUnitByQueryOptions) SetEnterpriseID(enterpriseID string) *GetBillingUnitByQueryOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetAccountGroupID : Allow user to set AccountGroupID
func (options *GetBillingUnitByQueryOptions) SetAccountGroupID(accountGroupID string) *GetBillingUnitByQueryOptions {
	options.AccountGroupID = core.StringPtr(accountGroupID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetBillingUnitByQueryOptions) SetHeaders(param map[string]string) *GetBillingUnitByQueryOptions {
	options.Headers = param
	return options
}

// GetCreditPoolsOptions : The GetCreditPools options.
type GetCreditPoolsOptions struct {
	// The ID of the billing unit.
	BillingUnitID *string `json:"billing_unit_id" validate:"required"`

	// The date in the format of YYYY-MM.
	Date *string `json:"date,omitempty"`

	// Filters the credit pool by type, either `PLATFORM` or `SUPPORT`.
	Type *string `json:"type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCreditPoolsOptions : Instantiate GetCreditPoolsOptions
func (*EnterpriseBillingUnitsV1) NewGetCreditPoolsOptions(billingUnitID string) *GetCreditPoolsOptions {
	return &GetCreditPoolsOptions{
		BillingUnitID: core.StringPtr(billingUnitID),
	}
}

// SetBillingUnitID : Allow user to set BillingUnitID
func (options *GetCreditPoolsOptions) SetBillingUnitID(billingUnitID string) *GetCreditPoolsOptions {
	options.BillingUnitID = core.StringPtr(billingUnitID)
	return options
}

// SetDate : Allow user to set Date
func (options *GetCreditPoolsOptions) SetDate(date string) *GetCreditPoolsOptions {
	options.Date = core.StringPtr(date)
	return options
}

// SetType : Allow user to set Type
func (options *GetCreditPoolsOptions) SetType(typeVar string) *GetCreditPoolsOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCreditPoolsOptions) SetHeaders(param map[string]string) *GetCreditPoolsOptions {
	options.Headers = param
	return options
}

// ResponseBillingUnitsByQuery : Aggregated usage and charges for all the plans in the account.
type ResponseBillingUnitsByQuery struct {
	// A count of the billing units that were found by the query.
	RowsCount *float64 `json:"rows_count,omitempty"`

	// Bookmark URL to query for next batch of billing units. This returns `null` if no additional pages are required.
	NextURL *string `json:"next_url,omitempty"`

	// A list of billing units found.
	Resources []BillingUnit `json:"resources,omitempty"`
}


// UnmarshalResponseBillingUnitsByQuery constructs an instance of ResponseBillingUnitsByQuery from the specified map.
func UnmarshalResponseBillingUnitsByQuery(m map[string]interface{}) (result *ResponseBillingUnitsByQuery, err error) {
	obj := new(ResponseBillingUnitsByQuery)
	obj.RowsCount, err = core.UnmarshalFloat64(m, "rows_count")
	if err != nil {
		return
	}
	obj.NextURL, err = core.UnmarshalString(m, "next_url")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalBillingUnitSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResponseBillingUnitsByQuerySlice unmarshals a slice of ResponseBillingUnitsByQuery instances from the specified list of maps.
func UnmarshalResponseBillingUnitsByQuerySlice(s []interface{}) (slice []ResponseBillingUnitsByQuery, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResponseBillingUnitsByQuery'")
			return
		}
		obj, e := UnmarshalResponseBillingUnitsByQuery(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResponseBillingUnitsByQueryAsProperty unmarshals an instance of ResponseBillingUnitsByQuery that is stored as a property
// within the specified map.
func UnmarshalResponseBillingUnitsByQueryAsProperty(m map[string]interface{}, propertyName string) (result *ResponseBillingUnitsByQuery, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResponseBillingUnitsByQuery'", propertyName)
			return
		}
		result, err = UnmarshalResponseBillingUnitsByQuery(objMap)
	}
	return
}

// UnmarshalResponseBillingUnitsByQuerySliceAsProperty unmarshals a slice of ResponseBillingUnitsByQuery instances that are stored as a property
// within the specified map.
func UnmarshalResponseBillingUnitsByQuerySliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResponseBillingUnitsByQuery, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResponseBillingUnitsByQuery'", propertyName)
			return
		}
		slice, err = UnmarshalResponseBillingUnitsByQuerySlice(vSlice)
	}
	return
}

// ResponseCreditPoolsByQuery : Aggregated usage and charges for all the plans in the account.
type ResponseCreditPoolsByQuery struct {
	// The number of credit pools that were found by the query.
	RowsCount *float64 `json:"rows_count,omitempty"`

	// A bookmark URL to the query for the next batch of billing units. Use a value of `null` if no additional pages are
	// required.
	NextURL *string `json:"next_url,omitempty"`

	// A list of credit pools found by the query.
	Resources []CreditPool `json:"resources,omitempty"`
}


// UnmarshalResponseCreditPoolsByQuery constructs an instance of ResponseCreditPoolsByQuery from the specified map.
func UnmarshalResponseCreditPoolsByQuery(m map[string]interface{}) (result *ResponseCreditPoolsByQuery, err error) {
	obj := new(ResponseCreditPoolsByQuery)
	obj.RowsCount, err = core.UnmarshalFloat64(m, "rows_count")
	if err != nil {
		return
	}
	obj.NextURL, err = core.UnmarshalString(m, "next_url")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalCreditPoolSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResponseCreditPoolsByQuerySlice unmarshals a slice of ResponseCreditPoolsByQuery instances from the specified list of maps.
func UnmarshalResponseCreditPoolsByQuerySlice(s []interface{}) (slice []ResponseCreditPoolsByQuery, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResponseCreditPoolsByQuery'")
			return
		}
		obj, e := UnmarshalResponseCreditPoolsByQuery(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResponseCreditPoolsByQueryAsProperty unmarshals an instance of ResponseCreditPoolsByQuery that is stored as a property
// within the specified map.
func UnmarshalResponseCreditPoolsByQueryAsProperty(m map[string]interface{}, propertyName string) (result *ResponseCreditPoolsByQuery, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResponseCreditPoolsByQuery'", propertyName)
			return
		}
		result, err = UnmarshalResponseCreditPoolsByQuery(objMap)
	}
	return
}

// UnmarshalResponseCreditPoolsByQuerySliceAsProperty unmarshals a slice of ResponseCreditPoolsByQuery instances that are stored as a property
// within the specified map.
func UnmarshalResponseCreditPoolsByQuerySliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResponseCreditPoolsByQuery, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResponseCreditPoolsByQuery'", propertyName)
			return
		}
		slice, err = UnmarshalResponseCreditPoolsByQuerySlice(vSlice)
	}
	return
}

// TermCredits : The subscription term that is active in the current month.
type TermCredits struct {
	// The ID of the billing option from which the subscription term is derived.
	BillingOptionID *string `json:"billing_option_id,omitempty"`

	// The category of the credit pool. The valid values are `PLATFORM`, `OFFER`, or `SERVICE` for platform credit and
	// `SUPPORT` for support credit.
	Category *string `json:"category,omitempty"`

	// The start date of the term in ISO format.
	StartDate *string `json:"start_date,omitempty"`

	// The end date of the term in ISO format.
	EndDate *string `json:"end_date,omitempty"`

	// The total credit available in this term.
	TotalCredits *float64 `json:"total_credits,omitempty"`

	// The balance of available credit at the start of the current month.
	StartingBalance *float64 `json:"starting_balance,omitempty"`

	// The amount of credit used during the current month.
	UsedCredits *float64 `json:"used_credits,omitempty"`

	// The balance of remaining credit in the subscription term.
	CurrentBalance *float64 `json:"current_balance,omitempty"`

	// A list of resources that used credit during the month.
	Resources []interface{} `json:"resources,omitempty"`
}


// UnmarshalTermCredits constructs an instance of TermCredits from the specified map.
func UnmarshalTermCredits(m map[string]interface{}) (result *TermCredits, err error) {
	obj := new(TermCredits)
	obj.BillingOptionID, err = core.UnmarshalString(m, "billing_option_id")
	if err != nil {
		return
	}
	obj.Category, err = core.UnmarshalString(m, "category")
	if err != nil {
		return
	}
	obj.StartDate, err = core.UnmarshalString(m, "start_date")
	if err != nil {
		return
	}
	obj.EndDate, err = core.UnmarshalString(m, "end_date")
	if err != nil {
		return
	}
	obj.TotalCredits, err = core.UnmarshalFloat64(m, "total_credits")
	if err != nil {
		return
	}
	obj.StartingBalance, err = core.UnmarshalFloat64(m, "starting_balance")
	if err != nil {
		return
	}
	obj.UsedCredits, err = core.UnmarshalFloat64(m, "used_credits")
	if err != nil {
		return
	}
	obj.CurrentBalance, err = core.UnmarshalFloat64(m, "current_balance")
	if err != nil {
		return
	}
	obj.Resources, err = core.UnmarshalAnySlice(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalTermCreditsSlice unmarshals a slice of TermCredits instances from the specified list of maps.
func UnmarshalTermCreditsSlice(s []interface{}) (slice []TermCredits, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'TermCredits'")
			return
		}
		obj, e := UnmarshalTermCredits(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalTermCreditsAsProperty unmarshals an instance of TermCredits that is stored as a property
// within the specified map.
func UnmarshalTermCreditsAsProperty(m map[string]interface{}, propertyName string) (result *TermCredits, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'TermCredits'", propertyName)
			return
		}
		result, err = UnmarshalTermCredits(objMap)
	}
	return
}

// UnmarshalTermCreditsSliceAsProperty unmarshals a slice of TermCredits instances that are stored as a property
// within the specified map.
func UnmarshalTermCreditsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []TermCredits, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'TermCredits'", propertyName)
			return
		}
		slice, err = UnmarshalTermCreditsSlice(vSlice)
	}
	return
}
