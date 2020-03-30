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

// Package usagemeteringv4 : Operations and models for the UsageMeteringV4 service
package usagemeteringv4

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// UsageMeteringV4 : IBM Cloud Usage Metering is a platform service that enables service providers to submit metrics
// collected for  resource instances provisioned by IBM Cloud users. IBM and third-party service providers that are
// delivering  an integrated billing service in IBM Cloud are required to submit usage for all active service instances
// every hour.  This is important because inability to report usage can lead to loss of revenue collection for IBM,  in
// turn causing loss of revenue share for the service providers.
//
// Version: 4.0.8
type UsageMeteringV4 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://billing.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "usage_metering"

// UsageMeteringV4Options : Service options
type UsageMeteringV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewUsageMeteringV4UsingExternalConfig : constructs an instance of UsageMeteringV4 with passed in options and external configuration.
func NewUsageMeteringV4UsingExternalConfig(options *UsageMeteringV4Options) (usageMetering *UsageMeteringV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	usageMetering, err = NewUsageMeteringV4(options)
	if err != nil {
		return
	}

	err = usageMetering.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = usageMetering.Service.SetServiceURL(options.URL)
	}
	return
}

// NewUsageMeteringV4 : constructs an instance of UsageMeteringV4 with passed in options.
func NewUsageMeteringV4(options *UsageMeteringV4Options) (service *UsageMeteringV4, err error) {
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

	service = &UsageMeteringV4{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (usageMetering *UsageMeteringV4) SetServiceURL(url string) error {
	return usageMetering.Service.SetServiceURL(url)
}

// ReportResourceUsage : Report Resource Controller resource usage
// Report usage for resource instances that were provisioned through the resource controller.
func (usageMetering *UsageMeteringV4) ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions) (result *ResponseAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reportResourceUsageOptions, "reportResourceUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reportResourceUsageOptions, "reportResourceUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v4/metering/resources", "usage"}
	pathParameters := []string{*reportResourceUsageOptions.ResourceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(usageMetering.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range reportResourceUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_metering", "V4", "ReportResourceUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(reportResourceUsageOptions.ResourceInstanceUsage)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageMetering.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResponseAccepted(m)
		response.Result = result
	}

	return
}

// ReportCfresourceUsage : Report Cloud Foundry resource usage
// Report usage for resource instances that were provisioned by using the Cloud Foundry provisioning broker.
func (usageMetering *UsageMeteringV4) ReportCfresourceUsage(reportCfresourceUsageOptions *ReportCfresourceUsageOptions) (result *ResponseAccepted, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reportCfresourceUsageOptions, "reportCfresourceUsageOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reportCfresourceUsageOptions, "reportCfresourceUsageOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v1/metering/resources", "usage"}
	pathParameters := []string{*reportCfresourceUsageOptions.ResourceID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(usageMetering.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range reportCfresourceUsageOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("usage_metering", "V4", "ReportCfresourceUsage")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(reportCfresourceUsageOptions.CfResourceInstanceUsage)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = usageMetering.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResponseAccepted(m)
		response.Result = result
	}

	return
}

// ReportCfresourceUsageOptions : The ReportCfresourceUsage options.
type ReportCfresourceUsageOptions struct {
	// The resource for which the usage is submitted.
	ResourceID *string `json:"resource_id" validate:"required"`

	// An array of usage records.
	CfResourceInstanceUsage []CfResourceInstanceUsage `json:"CfResourceInstanceUsage" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReportCfresourceUsageOptions : Instantiate ReportCfresourceUsageOptions
func (*UsageMeteringV4) NewReportCfresourceUsageOptions(resourceID string, cfResourceInstanceUsage []CfResourceInstanceUsage) *ReportCfresourceUsageOptions {
	return &ReportCfresourceUsageOptions{
		ResourceID: core.StringPtr(resourceID),
		CfResourceInstanceUsage: cfResourceInstanceUsage,
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *ReportCfresourceUsageOptions) SetResourceID(resourceID string) *ReportCfresourceUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetCfResourceInstanceUsage : Allow user to set CfResourceInstanceUsage
func (options *ReportCfresourceUsageOptions) SetCfResourceInstanceUsage(cfResourceInstanceUsage []CfResourceInstanceUsage) *ReportCfresourceUsageOptions {
	options.CfResourceInstanceUsage = cfResourceInstanceUsage
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReportCfresourceUsageOptions) SetHeaders(param map[string]string) *ReportCfresourceUsageOptions {
	options.Headers = param
	return options
}

// ReportResourceUsageOptions : The ReportResourceUsage options.
type ReportResourceUsageOptions struct {
	// The resource for which the usage is submitted.
	ResourceID *string `json:"resource_id" validate:"required"`

	// Array of usage records.
	ResourceInstanceUsage []ResourceInstanceUsage `json:"ResourceInstanceUsage" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReportResourceUsageOptions : Instantiate ReportResourceUsageOptions
func (*UsageMeteringV4) NewReportResourceUsageOptions(resourceID string, resourceInstanceUsage []ResourceInstanceUsage) *ReportResourceUsageOptions {
	return &ReportResourceUsageOptions{
		ResourceID: core.StringPtr(resourceID),
		ResourceInstanceUsage: resourceInstanceUsage,
	}
}

// SetResourceID : Allow user to set ResourceID
func (options *ReportResourceUsageOptions) SetResourceID(resourceID string) *ReportResourceUsageOptions {
	options.ResourceID = core.StringPtr(resourceID)
	return options
}

// SetResourceInstanceUsage : Allow user to set ResourceInstanceUsage
func (options *ReportResourceUsageOptions) SetResourceInstanceUsage(resourceInstanceUsage []ResourceInstanceUsage) *ReportResourceUsageOptions {
	options.ResourceInstanceUsage = resourceInstanceUsage
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReportResourceUsageOptions) SetHeaders(param map[string]string) *ReportResourceUsageOptions {
	options.Headers = param
	return options
}

// CfResourceInstanceUsage : Usage information for a Cloud Foundry resource instance.
type CfResourceInstanceUsage struct {
	// The ID of the orgnization to which the instance belongs. If the usage is for the instance is in a public cloud
	// region, then the organization ID format is `{mccp_region_name}:{guid}`. If the usage is for a syndicated instance,
	// then the organization ID format is `{mccp_region_id}:{guid}`.
	OrganizationID *string `json:"organization_id" validate:"required"`

	// The ID of the space under which the instance is provisioned.
	SpaceID *string `json:"space_id" validate:"required"`

	// The ID of the instance provisioned using the Cloud Foundry service controller.
	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	// The plan to meter the instance's usage with.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region for the usage. For public regions, the region parameter value is the mccp region name in which
	// the organization is present. For syndicated usage, use `us-south`.
	Region *string `json:"region" validate:"required"`

	// The time from which the resource instance was metered in the format milliseconds since epoch.
	Start *int64 `json:"start" validate:"required"`

	// The time until which the resource instance was metered in the format milliseconds since epoch. This value is the
	// same as start value for event-based submissions.
	End *int64 `json:"end" validate:"required"`

	// Usage measurements for the resource instance.
	MeasuredUsage []MeasureAndQuantity `json:"measured_usage" validate:"required"`

	// If an instance's usage should be aggregated at the consumer level, specify the ID of the consumer. Usage is
	// accumulated to instance-consumer combination.
	ConsumerID *string `json:"consumer_id,omitempty"`
}


// NewCfResourceInstanceUsage : Instantiate CfResourceInstanceUsage (Generic Model Constructor)
func (*UsageMeteringV4) NewCfResourceInstanceUsage(organizationID string, spaceID string, resourceInstanceID string, planID string, region string, start int64, end int64, measuredUsage []MeasureAndQuantity) (model *CfResourceInstanceUsage, err error) {
	model = &CfResourceInstanceUsage{
		OrganizationID: core.StringPtr(organizationID),
		SpaceID: core.StringPtr(spaceID),
		ResourceInstanceID: core.StringPtr(resourceInstanceID),
		PlanID: core.StringPtr(planID),
		Region: core.StringPtr(region),
		Start: core.Int64Ptr(start),
		End: core.Int64Ptr(end),
		MeasuredUsage: measuredUsage,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCfResourceInstanceUsage constructs an instance of CfResourceInstanceUsage from the specified map.
func UnmarshalCfResourceInstanceUsage(m map[string]interface{}) (result *CfResourceInstanceUsage, err error) {
	obj := new(CfResourceInstanceUsage)
	obj.OrganizationID, err = core.UnmarshalString(m, "organization_id")
	if err != nil {
		return
	}
	obj.SpaceID, err = core.UnmarshalString(m, "space_id")
	if err != nil {
		return
	}
	obj.ResourceInstanceID, err = core.UnmarshalString(m, "resource_instance_id")
	if err != nil {
		return
	}
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.Region, err = core.UnmarshalString(m, "region")
	if err != nil {
		return
	}
	obj.Start, err = core.UnmarshalInt64(m, "start")
	if err != nil {
		return
	}
	obj.End, err = core.UnmarshalInt64(m, "end")
	if err != nil {
		return
	}
	obj.MeasuredUsage, err = UnmarshalMeasureAndQuantitySliceAsProperty(m, "measured_usage")
	if err != nil {
		return
	}
	obj.ConsumerID, err = core.UnmarshalString(m, "consumer_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCfResourceInstanceUsageSlice unmarshals a slice of CfResourceInstanceUsage instances from the specified list of maps.
func UnmarshalCfResourceInstanceUsageSlice(s []interface{}) (slice []CfResourceInstanceUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'CfResourceInstanceUsage'")
			return
		}
		obj, e := UnmarshalCfResourceInstanceUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCfResourceInstanceUsageAsProperty unmarshals an instance of CfResourceInstanceUsage that is stored as a property
// within the specified map.
func UnmarshalCfResourceInstanceUsageAsProperty(m map[string]interface{}, propertyName string) (result *CfResourceInstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'CfResourceInstanceUsage'", propertyName)
			return
		}
		result, err = UnmarshalCfResourceInstanceUsage(objMap)
	}
	return
}

// UnmarshalCfResourceInstanceUsageSliceAsProperty unmarshals a slice of CfResourceInstanceUsage instances that are stored as a property
// within the specified map.
func UnmarshalCfResourceInstanceUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []CfResourceInstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'CfResourceInstanceUsage'", propertyName)
			return
		}
		slice, err = UnmarshalCfResourceInstanceUsageSlice(vSlice)
	}
	return
}

// MeasureAndQuantity : A usage measurement.
type MeasureAndQuantity struct {
	// The name of the measure.
	Measure *string `json:"measure" validate:"required"`

	// For consumption-based submissions, `quantity` can be a double or integer value. For event-based submissions that do
	// not have binary states, previous and current values are required, such as `{ "previous": 1, "current": 2 }`.
	Quantity interface{} `json:"quantity" validate:"required"`
}


// NewMeasureAndQuantity : Instantiate MeasureAndQuantity (Generic Model Constructor)
func (*UsageMeteringV4) NewMeasureAndQuantity(measure string, quantity interface{}) (model *MeasureAndQuantity, err error) {
	model = &MeasureAndQuantity{
		Measure: core.StringPtr(measure),
		Quantity: quantity,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalMeasureAndQuantity constructs an instance of MeasureAndQuantity from the specified map.
func UnmarshalMeasureAndQuantity(m map[string]interface{}) (result *MeasureAndQuantity, err error) {
	obj := new(MeasureAndQuantity)
	obj.Measure, err = core.UnmarshalString(m, "measure")
	if err != nil {
		return
	}
	obj.Quantity, err = core.UnmarshalAny(m, "quantity")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalMeasureAndQuantitySlice unmarshals a slice of MeasureAndQuantity instances from the specified list of maps.
func UnmarshalMeasureAndQuantitySlice(s []interface{}) (slice []MeasureAndQuantity, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'MeasureAndQuantity'")
			return
		}
		obj, e := UnmarshalMeasureAndQuantity(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalMeasureAndQuantityAsProperty unmarshals an instance of MeasureAndQuantity that is stored as a property
// within the specified map.
func UnmarshalMeasureAndQuantityAsProperty(m map[string]interface{}, propertyName string) (result *MeasureAndQuantity, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'MeasureAndQuantity'", propertyName)
			return
		}
		result, err = UnmarshalMeasureAndQuantity(objMap)
	}
	return
}

// UnmarshalMeasureAndQuantitySliceAsProperty unmarshals a slice of MeasureAndQuantity instances that are stored as a property
// within the specified map.
func UnmarshalMeasureAndQuantitySliceAsProperty(m map[string]interface{}, propertyName string) (slice []MeasureAndQuantity, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'MeasureAndQuantity'", propertyName)
			return
		}
		slice, err = UnmarshalMeasureAndQuantitySlice(vSlice)
	}
	return
}

// ResourceInstanceUsage : Usage information for a resource instance.
type ResourceInstanceUsage struct {
	// The ID of the instance that incurred the usage. The ID is a CRN for instances that are provisioned with the resource
	// controller.
	ResourceInstanceID *string `json:"resource_instance_id" validate:"required"`

	// The plan with which the instance's usage should be metered.
	PlanID *string `json:"plan_id" validate:"required"`

	// The pricing region to which the usage must be aggregated. This field is required if the ID is not a CRN or if the
	// CRN does not have a region.
	Region *string `json:"region,omitempty"`

	// The time from which the resource instance was metered in the format milliseconds since epoch.
	Start *int64 `json:"start" validate:"required"`

	// The time until which the resource instance was metered in the format milliseconds since epoch. This value is the
	// same as start value for event-based submissions.
	End *int64 `json:"end" validate:"required"`

	// Usage measurements for the resource instance.
	MeasuredUsage []MeasureAndQuantity `json:"measured_usage" validate:"required"`

	// If an instance's usage should be aggregated at the consumer level, specify the ID of the consumer. Usage is
	// accumulated to the instance-consumer combination.
	ConsumerID *string `json:"consumer_id,omitempty"`
}


// NewResourceInstanceUsage : Instantiate ResourceInstanceUsage (Generic Model Constructor)
func (*UsageMeteringV4) NewResourceInstanceUsage(resourceInstanceID string, planID string, start int64, end int64, measuredUsage []MeasureAndQuantity) (model *ResourceInstanceUsage, err error) {
	model = &ResourceInstanceUsage{
		ResourceInstanceID: core.StringPtr(resourceInstanceID),
		PlanID: core.StringPtr(planID),
		Start: core.Int64Ptr(start),
		End: core.Int64Ptr(end),
		MeasuredUsage: measuredUsage,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResourceInstanceUsage constructs an instance of ResourceInstanceUsage from the specified map.
func UnmarshalResourceInstanceUsage(m map[string]interface{}) (result *ResourceInstanceUsage, err error) {
	obj := new(ResourceInstanceUsage)
	obj.ResourceInstanceID, err = core.UnmarshalString(m, "resource_instance_id")
	if err != nil {
		return
	}
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.Region, err = core.UnmarshalString(m, "region")
	if err != nil {
		return
	}
	obj.Start, err = core.UnmarshalInt64(m, "start")
	if err != nil {
		return
	}
	obj.End, err = core.UnmarshalInt64(m, "end")
	if err != nil {
		return
	}
	obj.MeasuredUsage, err = UnmarshalMeasureAndQuantitySliceAsProperty(m, "measured_usage")
	if err != nil {
		return
	}
	obj.ConsumerID, err = core.UnmarshalString(m, "consumer_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceInstanceUsageSlice unmarshals a slice of ResourceInstanceUsage instances from the specified list of maps.
func UnmarshalResourceInstanceUsageSlice(s []interface{}) (slice []ResourceInstanceUsage, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceInstanceUsage'")
			return
		}
		obj, e := UnmarshalResourceInstanceUsage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceInstanceUsageAsProperty unmarshals an instance of ResourceInstanceUsage that is stored as a property
// within the specified map.
func UnmarshalResourceInstanceUsageAsProperty(m map[string]interface{}, propertyName string) (result *ResourceInstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceInstanceUsage'", propertyName)
			return
		}
		result, err = UnmarshalResourceInstanceUsage(objMap)
	}
	return
}

// UnmarshalResourceInstanceUsageSliceAsProperty unmarshals a slice of ResourceInstanceUsage instances that are stored as a property
// within the specified map.
func UnmarshalResourceInstanceUsageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceInstanceUsage, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceInstanceUsage'", propertyName)
			return
		}
		slice, err = UnmarshalResourceInstanceUsageSlice(vSlice)
	}
	return
}

// ResourceUsageDetails : Resource usage details.
type ResourceUsageDetails struct {
	// A response code similar to HTTP status codes.
	Status *float64 `json:"status" validate:"required"`

	// The location of the usage.
	Location *string `json:"location" validate:"required"`

	// The error code that was encountered.
	Code *string `json:"code,omitempty"`

	// A description of the error.
	Message *string `json:"message,omitempty"`
}


// UnmarshalResourceUsageDetails constructs an instance of ResourceUsageDetails from the specified map.
func UnmarshalResourceUsageDetails(m map[string]interface{}) (result *ResourceUsageDetails, err error) {
	obj := new(ResourceUsageDetails)
	obj.Status, err = core.UnmarshalFloat64(m, "status")
	if err != nil {
		return
	}
	obj.Location, err = core.UnmarshalString(m, "location")
	if err != nil {
		return
	}
	obj.Code, err = core.UnmarshalString(m, "code")
	if err != nil {
		return
	}
	obj.Message, err = core.UnmarshalString(m, "message")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResourceUsageDetailsSlice unmarshals a slice of ResourceUsageDetails instances from the specified list of maps.
func UnmarshalResourceUsageDetailsSlice(s []interface{}) (slice []ResourceUsageDetails, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResourceUsageDetails'")
			return
		}
		obj, e := UnmarshalResourceUsageDetails(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResourceUsageDetailsAsProperty unmarshals an instance of ResourceUsageDetails that is stored as a property
// within the specified map.
func UnmarshalResourceUsageDetailsAsProperty(m map[string]interface{}, propertyName string) (result *ResourceUsageDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResourceUsageDetails'", propertyName)
			return
		}
		result, err = UnmarshalResourceUsageDetails(objMap)
	}
	return
}

// UnmarshalResourceUsageDetailsSliceAsProperty unmarshals a slice of ResourceUsageDetails instances that are stored as a property
// within the specified map.
func UnmarshalResourceUsageDetailsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResourceUsageDetails, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResourceUsageDetails'", propertyName)
			return
		}
		slice, err = UnmarshalResourceUsageDetailsSlice(vSlice)
	}
	return
}

// ResponseAccepted : Response when usage submitted is accepted.
type ResponseAccepted struct {
	// Response body that contains the status of each submitted usage record.
	Resources []ResourceUsageDetails `json:"resources" validate:"required"`
}


// UnmarshalResponseAccepted constructs an instance of ResponseAccepted from the specified map.
func UnmarshalResponseAccepted(m map[string]interface{}) (result *ResponseAccepted, err error) {
	obj := new(ResponseAccepted)
	obj.Resources, err = UnmarshalResourceUsageDetailsSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResponseAcceptedSlice unmarshals a slice of ResponseAccepted instances from the specified list of maps.
func UnmarshalResponseAcceptedSlice(s []interface{}) (slice []ResponseAccepted, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ResponseAccepted'")
			return
		}
		obj, e := UnmarshalResponseAccepted(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResponseAcceptedAsProperty unmarshals an instance of ResponseAccepted that is stored as a property
// within the specified map.
func UnmarshalResponseAcceptedAsProperty(m map[string]interface{}, propertyName string) (result *ResponseAccepted, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ResponseAccepted'", propertyName)
			return
		}
		result, err = UnmarshalResponseAccepted(objMap)
	}
	return
}

// UnmarshalResponseAcceptedSliceAsProperty unmarshals a slice of ResponseAccepted instances that are stored as a property
// within the specified map.
func UnmarshalResponseAcceptedSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ResponseAccepted, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ResponseAccepted'", propertyName)
			return
		}
		slice, err = UnmarshalResponseAcceptedSlice(vSlice)
	}
	return
}
