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

// Package openservicebrokerv1 : Operations and models for the OpenServiceBrokerV1 service
package openservicebrokerv1

import (
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.ibm.com/ibmcloud/platform-services-go-sdk/common"
)

// OpenServiceBrokerV1 : Contribute resources to the IBM Cloud catalog by implementing a `service broker` that conforms
// to the [Open Service Broker API](https://github.com/openservicebrokerapi/servicebroker/blob/master/spec.md) version
// 2.12  specification and provides enablement extensions for integration with IBM Cloud and the Resource Controller
// provisioning model.
//
// Version: 1.4
type OpenServiceBrokerV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = ""

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "open_service_broker"

// OpenServiceBrokerV1Options : Service options
type OpenServiceBrokerV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewOpenServiceBrokerV1UsingExternalConfig : constructs an instance of OpenServiceBrokerV1 with passed in options and external configuration.
func NewOpenServiceBrokerV1UsingExternalConfig(options *OpenServiceBrokerV1Options) (openServiceBroker *OpenServiceBrokerV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	openServiceBroker, err = NewOpenServiceBrokerV1(options)
	if err != nil {
		return
	}

	err = openServiceBroker.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = openServiceBroker.Service.SetServiceURL(options.URL)
	}
	return
}

// NewOpenServiceBrokerV1 : constructs an instance of OpenServiceBrokerV1 with passed in options.
func NewOpenServiceBrokerV1(options *OpenServiceBrokerV1Options) (service *OpenServiceBrokerV1, err error) {
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

	service = &OpenServiceBrokerV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (openServiceBroker *OpenServiceBrokerV1) SetServiceURL(url string) error {
	return openServiceBroker.Service.SetServiceURL(url)
}

// GetServiceInstanceState : Get the current state of the service instance
// Get the current state information associated with the service instance.
//
// As a service provider you need a way to manage provisioned service instances.  If an account comes past due, you may
// need a to disable the service (without deleting it), and when the account is settled re-enable the service.
//
// This endpoint allows both the provider and IBM Cloud to query for the state of a provisioned service instance.  For
// example, IBM Cloud may query the provider to figure out if a given service is disabled or not and present that state
// to the user.
func (openServiceBroker *OpenServiceBrokerV1) GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions) (result *Resp1874644Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getServiceInstanceStateOptions, "getServiceInstanceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getServiceInstanceStateOptions, "getServiceInstanceStateOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"bluemix_v1/service_instances"}
	pathParameters := []string{*getServiceInstanceStateOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getServiceInstanceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "GetServiceInstanceState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResp1874644Root(m)
		response.Result = result
	}

	return
}

// ReplaceState : Update the state of a provisioned service instance
// Update (disable or enable) the state of a provisioned service instance. As a service provider you need a way to
// manage provisioned service instances. If an account comes past due, you may need a to disable the service (without
// deleting it), and when the account is settled re-enable the service. This endpoint allows the provider to enable or
// disable the state of a provisioned service instance. It is the service provider's responsibility to disable access to
// the service instance when the disable endpoint is invoked and to re-enable that access when the enable endpoint is
// invoked. When your service broker receives an enable / disable request, it should take whatever action is necessary
// to enable / disable (respectively) the service.  Additionally, If a bind request comes in for a disabled service, the
// broker should reject that request with any code other than `204`, and provide a user-facing message in the
// description.
func (openServiceBroker *OpenServiceBrokerV1) ReplaceState(replaceStateOptions *ReplaceStateOptions) (result *Resp2448145Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceStateOptions, "replaceStateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceStateOptions, "replaceStateOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"bluemix_v1/service_instances"}
	pathParameters := []string{*replaceStateOptions.InstanceID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceStateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceState")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceStateOptions.Enabled != nil {
		body["enabled"] = replaceStateOptions.Enabled
	}
	if replaceStateOptions.InitiatorID != nil {
		body["initiator_id"] = replaceStateOptions.InitiatorID
	}
	if replaceStateOptions.ReasonCode != nil {
		body["reason_code"] = replaceStateOptions.ReasonCode
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResp2448145Root(m)
		response.Result = result
	}

	return
}

// ReplaceServiceInstance : Create (provision) a service instance
// Create a service instance with GUID. When your service broker receives a provision request from the IBM Cloud
// platform, it MUST take whatever action is necessary to create a new resource.
//
// When a user creates a service instance from the IBM Cloud console or the IBM Cloud CLI, the IBM Cloud platform
// validates that the user has permission to create the service instance using IBM Cloud IAM. After this validation
// occurs, your service broker's provision endpoint (PUT /v2/resource_instances/:instance_id) will be invoked. When
// provisioning occurs, the IBM Cloud platform provides the following values:
//
// - The IBM Cloud context is included in the context variable
// - The X-Broker-API-Originating-Identity will have the IBM IAM ID of the user that initiated the request
// - The parameters section will include the requested location (and additional parameters required by your service).
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) (result *Resp2079872Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceInstanceOptions, "replaceServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceInstanceOptions, "replaceServiceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances"}
	pathParameters := []string{*replaceServiceInstanceOptions.InstanceID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if replaceServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*replaceServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if replaceServiceInstanceOptions.Context != nil {
		body["context"] = replaceServiceInstanceOptions.Context
	}
	if replaceServiceInstanceOptions.OrganizationGuid != nil {
		body["organization_guid"] = replaceServiceInstanceOptions.OrganizationGuid
	}
	if replaceServiceInstanceOptions.Parameters != nil {
		body["parameters"] = replaceServiceInstanceOptions.Parameters
	}
	if replaceServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = replaceServiceInstanceOptions.PlanID
	}
	if replaceServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = replaceServiceInstanceOptions.ServiceID
	}
	if replaceServiceInstanceOptions.SpaceGuid != nil {
		body["space_guid"] = replaceServiceInstanceOptions.SpaceGuid
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResp2079872Root(m)
		response.Result = result
	}

	return
}

// UpdateServiceInstance : Update a service instance
// Patch an instance by GUID. Enabling this endpoint allows your user to change plans and service parameters in a
// provisioned service instance. If your offering supports multiple plans, and you want users to be able to change plans
// for a provisioned instance, you will need to enable the ability for users to update their service instance.
//
// To enable support for the update of the plan, a broker MUST declare support per service by specifying
// `"plan_updateable": true` in your brokers' catalog.json.
func (openServiceBroker *OpenServiceBrokerV1) UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateServiceInstanceOptions, "updateServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateServiceInstanceOptions, "updateServiceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances"}
	pathParameters := []string{*updateServiceInstanceOptions.InstanceID}

	builder := core.NewRequestBuilder(core.PATCH)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "UpdateServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*updateServiceInstanceOptions.AcceptsIncomplete))
	}

	body := make(map[string]interface{})
	if updateServiceInstanceOptions.Context != nil {
		body["context"] = updateServiceInstanceOptions.Context
	}
	if updateServiceInstanceOptions.Parameters != nil {
		body["parameters"] = updateServiceInstanceOptions.Parameters
	}
	if updateServiceInstanceOptions.PlanID != nil {
		body["plan_id"] = updateServiceInstanceOptions.PlanID
	}
	if updateServiceInstanceOptions.PreviousValues != nil {
		body["previous_values"] = updateServiceInstanceOptions.PreviousValues
	}
	if updateServiceInstanceOptions.ServiceID != nil {
		body["service_id"] = updateServiceInstanceOptions.ServiceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, new(string))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*string)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// DeleteServiceInstance : Delete (deprovision) a service instance
// Delete (deprovision) a service instance by GUID. When a service broker receives a deprovision request from the IBM
// Cloud platform, it MUST delete any resources it created during the provision. Usually this means that all resources
// are immediately reclaimed for future provisions.
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceInstanceOptions, "deleteServiceInstanceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceInstanceOptions, "deleteServiceInstanceOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances"}
	pathParameters := []string{*deleteServiceInstanceOptions.InstanceID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceInstanceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "DeleteServiceInstance")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("service_id", fmt.Sprint(*deleteServiceInstanceOptions.ServiceID))
	builder.AddQuery("plan_id", fmt.Sprint(*deleteServiceInstanceOptions.PlanID))
	if deleteServiceInstanceOptions.AcceptsIncomplete != nil {
		builder.AddQuery("accepts_incomplete", fmt.Sprint(*deleteServiceInstanceOptions.AcceptsIncomplete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, new(string))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*string)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// ListCatalog : Get the catalog metadata stored within the broker
// This endpoints defines the contract between the broker and the IBM Cloud platform for the services and plans that the
// broker supports. This endpoint returns the catalog metadata stored within your broker. These values define the
// minimal provisioning contract between your service and the IBM Cloud platform. All additional catalog metadata that
// is not required for provisioning is stored within the IBM Cloud catalog, and any updates to catalog display values
// that are used to render your dashboard like links, icons, and i18n translated metadata should be updated in the
// Resource Management Console (RMC), and not housed in your broker. None of metadata stored in your broker is displayed
// in the IBM Cloud console or the IBM Cloud CLI; the console and CLI will return what was set withn RMC and stored in
// the IBM Cloud catalog.
func (openServiceBroker *OpenServiceBrokerV1) ListCatalog(listCatalogOptions *ListCatalogOptions) (result *[]Services, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogOptions, "listCatalogOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/catalog"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ListCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, make([]map[string]interface{}, 1))
	if err == nil {
		s, ok := response.Result.([]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		slice, e := UnmarshalServicesSlice(s)
		result = &slice
		err = e
		response.Result = result
	}

	return
}

// ListLastOperation : Get the current status of a provision in-progress for a service instance
// Get `last_operation` for instance by GUID (for asynchronous provision calls). When a broker returns status code `202
// Accepted` during a provision, update, or deprovision call, the IBM Cloud platform will begin polling the
// `last_operation` endpoint to obtain the state of the last requested operation. The broker response MUST contain the
// field `state` and MAY contain the field `description`.
//
// Valid values for `state` are `in progress`, `succeeded`, and `failed`. The platform will poll the `last_operation
// `endpoint as long as the broker returns "state": "in progress". Returning "state": "succeeded" or "state": "failed"
// will cause the platform to cease polling. The value provided for description will be passed through to the platform
// API client and can be used to provide additional detail for users about the progress of the operation.
func (openServiceBroker *OpenServiceBrokerV1) ListLastOperation(listLastOperationOptions *ListLastOperationOptions) (result *Resp2079894Root, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLastOperationOptions, "listLastOperationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLastOperationOptions, "listLastOperationOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances", "last_operation"}
	pathParameters := []string{*listLastOperationOptions.InstanceID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLastOperationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ListLastOperation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listLastOperationOptions.Operation != nil {
		builder.AddQuery("operation", fmt.Sprint(*listLastOperationOptions.Operation))
	}
	if listLastOperationOptions.PlanID != nil {
		builder.AddQuery("plan_id", fmt.Sprint(*listLastOperationOptions.PlanID))
	}
	if listLastOperationOptions.ServiceID != nil {
		builder.AddQuery("service_id", fmt.Sprint(*listLastOperationOptions.ServiceID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalResp2079894Root(m)
		response.Result = result
	}

	return
}

// ReplaceServiceBinding : Bind a service instance to another resource
// Create binding by GUID on service instance.
//
// If your service can be bound to applications in IBM Cloud, `bindable:true` must be specified in the catalog.json of
// your service broker. If bindable, it must be able to return API endpoints and credentials to your service consumers.
//
// **Note:** Brokers that do not offer any bindable services do not need to implement the endpoint for bind requests.
//
// See the OSB 2.12 spec for more details on
// [binding](https://github.com/openservicebrokerapi/servicebroker/blob/v2.12/spec.md#binding).
func (openServiceBroker *OpenServiceBrokerV1) ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceServiceBindingOptions, "replaceServiceBindingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceServiceBindingOptions, "replaceServiceBindingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances", "service_bindings"}
	pathParameters := []string{*replaceServiceBindingOptions.BindingID, *replaceServiceBindingOptions.InstanceID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceServiceBindingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "ReplaceServiceBinding")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceServiceBindingOptions.BindResource != nil {
		body["bind_resource"] = replaceServiceBindingOptions.BindResource
	}
	if replaceServiceBindingOptions.Parameters != nil {
		body["parameters"] = replaceServiceBindingOptions.Parameters
	}
	if replaceServiceBindingOptions.PlanID != nil {
		body["plan_id"] = replaceServiceBindingOptions.PlanID
	}
	if replaceServiceBindingOptions.ServiceID != nil {
		body["service_id"] = replaceServiceBindingOptions.ServiceID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, new(string))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*string)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// DeleteServiceBinding : Delete (unbind) the credentials bound to a resource
// Delete instance binding by GUID.
//
// When a broker receives an unbind request from the IBM Cloud platform, it MUST delete any resources associated with
// the binding. In the case where credentials were generated, this might result in requests to the service instance
// failing to authenticate.
//
// **Note**: Brokers that do not provide any bindable services or plans do not need to implement this endpoint.
func (openServiceBroker *OpenServiceBrokerV1) DeleteServiceBinding(deleteServiceBindingOptions *DeleteServiceBindingOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteServiceBindingOptions, "deleteServiceBindingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteServiceBindingOptions, "deleteServiceBindingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/service_instances", "service_bindings"}
	pathParameters := []string{*deleteServiceBindingOptions.BindingID, *deleteServiceBindingOptions.InstanceID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(openServiceBroker.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteServiceBindingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("open_service_broker", "V1", "DeleteServiceBinding")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("plan_id", fmt.Sprint(*deleteServiceBindingOptions.PlanID))
	builder.AddQuery("service_id", fmt.Sprint(*deleteServiceBindingOptions.ServiceID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = openServiceBroker.Service.Request(request, new(string))
	if err == nil {
		var ok bool
		result, ok = response.Result.(*string)
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
		}
	}

	return
}

// DeleteServiceBindingOptions : The DeleteServiceBinding options.
type DeleteServiceBindingOptions struct {
	// The `binding_id` is the ID of a previously provisioned binding for that service instance.
	BindingID *string `json:"binding_id" validate:"required"`

	// The `instance_id` is the ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// The ID of the plan from the catalog.json in the broker. It MUST be a non-empty string and should be a GUID.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of the service from the catalog.json in the broker. It MUST be a non-empty string and should be a GUID.
	ServiceID *string `json:"service_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceBindingOptions : Instantiate DeleteServiceBindingOptions
func (*OpenServiceBrokerV1) NewDeleteServiceBindingOptions(bindingID string, instanceID string, planID string, serviceID string) *DeleteServiceBindingOptions {
	return &DeleteServiceBindingOptions{
		BindingID: core.StringPtr(bindingID),
		InstanceID: core.StringPtr(instanceID),
		PlanID: core.StringPtr(planID),
		ServiceID: core.StringPtr(serviceID),
	}
}

// SetBindingID : Allow user to set BindingID
func (options *DeleteServiceBindingOptions) SetBindingID(bindingID string) *DeleteServiceBindingOptions {
	options.BindingID = core.StringPtr(bindingID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteServiceBindingOptions) SetInstanceID(instanceID string) *DeleteServiceBindingOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *DeleteServiceBindingOptions) SetPlanID(planID string) *DeleteServiceBindingOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *DeleteServiceBindingOptions) SetServiceID(serviceID string) *DeleteServiceBindingOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceBindingOptions) SetHeaders(param map[string]string) *DeleteServiceBindingOptions {
	options.Headers = param
	return options
}

// DeleteServiceInstanceOptions : The DeleteServiceInstance options.
type DeleteServiceInstanceOptions struct {
	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. MUST be a non-empty
	// string.
	ServiceID *string `json:"service_id" validate:"required"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID. MUST be a non-empty string.
	PlanID *string `json:"plan_id" validate:"required"`

	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteServiceInstanceOptions : Instantiate DeleteServiceInstanceOptions
func (*OpenServiceBrokerV1) NewDeleteServiceInstanceOptions(serviceID string, planID string, instanceID string) *DeleteServiceInstanceOptions {
	return &DeleteServiceInstanceOptions{
		ServiceID: core.StringPtr(serviceID),
		PlanID: core.StringPtr(planID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetServiceID : Allow user to set ServiceID
func (options *DeleteServiceInstanceOptions) SetServiceID(serviceID string) *DeleteServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *DeleteServiceInstanceOptions) SetPlanID(planID string) *DeleteServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *DeleteServiceInstanceOptions) SetInstanceID(instanceID string) *DeleteServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *DeleteServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *DeleteServiceInstanceOptions {
	options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteServiceInstanceOptions) SetHeaders(param map[string]string) *DeleteServiceInstanceOptions {
	options.Headers = param
	return options
}

// GetServiceInstanceStateOptions : The GetServiceInstanceState options.
type GetServiceInstanceStateOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetServiceInstanceStateOptions : Instantiate GetServiceInstanceStateOptions
func (*OpenServiceBrokerV1) NewGetServiceInstanceStateOptions(instanceID string) *GetServiceInstanceStateOptions {
	return &GetServiceInstanceStateOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *GetServiceInstanceStateOptions) SetInstanceID(instanceID string) *GetServiceInstanceStateOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetServiceInstanceStateOptions) SetHeaders(param map[string]string) *GetServiceInstanceStateOptions {
	options.Headers = param
	return options
}

// ListCatalogOptions : The ListCatalog options.
type ListCatalogOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCatalogOptions : Instantiate ListCatalogOptions
func (*OpenServiceBrokerV1) NewListCatalogOptions() *ListCatalogOptions {
	return &ListCatalogOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogOptions) SetHeaders(param map[string]string) *ListCatalogOptions {
	options.Headers = param
	return options
}

// ListLastOperationOptions : The ListLastOperation options.
type ListLastOperationOptions struct {
	// The unique instance ID generated during provisioning by the IBM Cloud platform.
	InstanceID *string `json:"instance_id" validate:"required"`

	// A broker-provided identifier for the operation. When a value for operation is included with asynchronous responses
	// for provision and update, and deprovision requests, the IBM Cloud platform will provide the same value using this
	// query parameter as a URL-encoded string. If present, MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`

	// ID of the plan from the catalog.json in your broker. If present, MUST be a non-empty string.
	PlanID *string `json:"plan_id,omitempty"`

	// ID of the service from the catalog.json in your service broker. If present, MUST be a non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLastOperationOptions : Instantiate ListLastOperationOptions
func (*OpenServiceBrokerV1) NewListLastOperationOptions(instanceID string) *ListLastOperationOptions {
	return &ListLastOperationOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ListLastOperationOptions) SetInstanceID(instanceID string) *ListLastOperationOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetOperation : Allow user to set Operation
func (options *ListLastOperationOptions) SetOperation(operation string) *ListLastOperationOptions {
	options.Operation = core.StringPtr(operation)
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *ListLastOperationOptions) SetPlanID(planID string) *ListLastOperationOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *ListLastOperationOptions) SetServiceID(serviceID string) *ListLastOperationOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLastOperationOptions) SetHeaders(param map[string]string) *ListLastOperationOptions {
	options.Headers = param
	return options
}

// ReplaceServiceBindingOptions : The ReplaceServiceBinding options.
type ReplaceServiceBindingOptions struct {
	// The `binding_id` is provided by the IBM Cloud platform. This ID will be used for future unbind requests, so the
	// broker can use it to correlate the resource it creates.
	BindingID *string `json:"binding_id" validate:"required"`

	// The :`instance_id` is the ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// A JSON object that contains data for platform resources associated with the binding to be created.
	BindResource []BindResource `json:"bind_resource,omitempty"`

	// Configuration options for the service binding.
	Parameters interface{} `json:"parameters,omitempty"`

	// The ID of the plan from the catalog.json in your broker. If present, it MUST be a non-empty string.
	PlanID *string `json:"plan_id,omitempty"`

	// The ID of the service from the catalog.json in your broker. If present, it MUST be a non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceBindingOptions : Instantiate ReplaceServiceBindingOptions
func (*OpenServiceBrokerV1) NewReplaceServiceBindingOptions(bindingID string, instanceID string) *ReplaceServiceBindingOptions {
	return &ReplaceServiceBindingOptions{
		BindingID: core.StringPtr(bindingID),
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetBindingID : Allow user to set BindingID
func (options *ReplaceServiceBindingOptions) SetBindingID(bindingID string) *ReplaceServiceBindingOptions {
	options.BindingID = core.StringPtr(bindingID)
	return options
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceServiceBindingOptions) SetInstanceID(instanceID string) *ReplaceServiceBindingOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetBindResource : Allow user to set BindResource
func (options *ReplaceServiceBindingOptions) SetBindResource(bindResource []BindResource) *ReplaceServiceBindingOptions {
	options.BindResource = bindResource
	return options
}

// SetParameters : Allow user to set Parameters
func (options *ReplaceServiceBindingOptions) SetParameters(parameters interface{}) *ReplaceServiceBindingOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *ReplaceServiceBindingOptions) SetPlanID(planID string) *ReplaceServiceBindingOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *ReplaceServiceBindingOptions) SetServiceID(serviceID string) *ReplaceServiceBindingOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceBindingOptions) SetHeaders(param map[string]string) *ReplaceServiceBindingOptions {
	options.Headers = param
	return options
}

// ReplaceServiceInstanceOptions : The ReplaceServiceInstance options.
type ReplaceServiceInstanceOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Platform specific contextual information under which the service instance is to be provisioned.
	Context []Context `json:"context,omitempty"`

	// Deprecated in favor of `context`. The identifier for the project space within the IBM Cloud platform organization.
	// Although most brokers will not use this field, it might be helpful for executing operations on a user's behalf. It
	// MUST be a non-empty string.
	OrganizationGuid *string `json:"organization_guid,omitempty"`

	// A list of plans for this service that must contain at least one plan.
	Parameters []Parameters `json:"parameters,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID and it MUST be unique to a service.
	PlanID *string `json:"plan_id,omitempty"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID and it MUST be a
	// non-empty string.
	ServiceID *string `json:"service_id,omitempty"`

	// Deprecated in favor of `context`. The IBM Cloud platform GUID for the organization under which the service instance
	// is to be provisioned. Although most brokers will not use this field, it might be helpful for executing operations on
	// a user's behalf. It MUST be a non-empty string.
	SpaceGuid *string `json:"space_guid,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *bool `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceServiceInstanceOptions : Instantiate ReplaceServiceInstanceOptions
func (*OpenServiceBrokerV1) NewReplaceServiceInstanceOptions(instanceID string) *ReplaceServiceInstanceOptions {
	return &ReplaceServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceServiceInstanceOptions) SetInstanceID(instanceID string) *ReplaceServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetContext : Allow user to set Context
func (options *ReplaceServiceInstanceOptions) SetContext(context []Context) *ReplaceServiceInstanceOptions {
	options.Context = context
	return options
}

// SetOrganizationGuid : Allow user to set OrganizationGuid
func (options *ReplaceServiceInstanceOptions) SetOrganizationGuid(organizationGuid string) *ReplaceServiceInstanceOptions {
	options.OrganizationGuid = core.StringPtr(organizationGuid)
	return options
}

// SetParameters : Allow user to set Parameters
func (options *ReplaceServiceInstanceOptions) SetParameters(parameters []Parameters) *ReplaceServiceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *ReplaceServiceInstanceOptions) SetPlanID(planID string) *ReplaceServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *ReplaceServiceInstanceOptions) SetServiceID(serviceID string) *ReplaceServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetSpaceGuid : Allow user to set SpaceGuid
func (options *ReplaceServiceInstanceOptions) SetSpaceGuid(spaceGuid string) *ReplaceServiceInstanceOptions {
	options.SpaceGuid = core.StringPtr(spaceGuid)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *ReplaceServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete bool) *ReplaceServiceInstanceOptions {
	options.AcceptsIncomplete = core.BoolPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceServiceInstanceOptions) SetHeaders(param map[string]string) *ReplaceServiceInstanceOptions {
	options.Headers = param
	return options
}

// ReplaceStateOptions : The ReplaceState options.
type ReplaceStateOptions struct {
	// The `instance_id` of a service instance is provided by the IBM Cloud platform. This ID will be used for future
	// requests to bind and deprovision, so the broker can use it to correlate the resource it creates.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled,omitempty"`

	// Optional string that shows the user ID that is initiating the call.
	InitiatorID *string `json:"initiator_id,omitempty"`

	// Optional string that states the reason code for the service instance state change. Valid values are
	// `IBMCLOUD_ACCT_ACTIVATE`, `IBMCLOUD_RECLAMATION_RESTORE`, or `IBMCLOUD_SERVICE_INSTANCE_BELOW_CAP` for enable calls;
	// `IBMCLOUD_ACCT_SUSPEND`, `IBMCLOUD_RECLAMATION_SCHEDULE`, or `IBMCLOUD_SERVICE_INSTANCE_ABOVE_CAP` for disable
	// calls; and `IBMCLOUD_ADMIN_REQUEST` for enable and disable calls.<br/><br/>Previously accepted values had a `BMX_`
	// prefix, such as `BMX_ACCT_ACTIVATE`. These values are deprecated.
	ReasonCode *string `json:"reason_code,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceStateOptions : Instantiate ReplaceStateOptions
func (*OpenServiceBrokerV1) NewReplaceStateOptions(instanceID string) *ReplaceStateOptions {
	return &ReplaceStateOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *ReplaceStateOptions) SetInstanceID(instanceID string) *ReplaceStateOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetEnabled : Allow user to set Enabled
func (options *ReplaceStateOptions) SetEnabled(enabled bool) *ReplaceStateOptions {
	options.Enabled = core.BoolPtr(enabled)
	return options
}

// SetInitiatorID : Allow user to set InitiatorID
func (options *ReplaceStateOptions) SetInitiatorID(initiatorID string) *ReplaceStateOptions {
	options.InitiatorID = core.StringPtr(initiatorID)
	return options
}

// SetReasonCode : Allow user to set ReasonCode
func (options *ReplaceStateOptions) SetReasonCode(reasonCode string) *ReplaceStateOptions {
	options.ReasonCode = core.StringPtr(reasonCode)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceStateOptions) SetHeaders(param map[string]string) *ReplaceStateOptions {
	options.Headers = param
	return options
}

// Resp1874644Root : Check the active status of an enabled service.
type Resp1874644Root struct {
	// Indicates (from the viewpoint of the provider) whether the service instance is active and is meaningful if enabled
	// is true. The default value is true if not specified.
	Active *bool `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled,omitempty"`

	// Indicates when the service instance was last accessed/modified/etc., and is meaningful if enabled is true AND active
	// is false. Represented as milliseconds since the epoch, but does not need to be accurate to the second/hour.
	LastActive *float64 `json:"last_active,omitempty"`
}


// UnmarshalResp1874644Root constructs an instance of Resp1874644Root from the specified map.
func UnmarshalResp1874644Root(m map[string]interface{}) (result *Resp1874644Root, err error) {
	obj := new(Resp1874644Root)
	obj.Active, err = core.UnmarshalBool(m, "active")
	if err != nil {
		return
	}
	obj.Enabled, err = core.UnmarshalBool(m, "enabled")
	if err != nil {
		return
	}
	obj.LastActive, err = core.UnmarshalFloat64(m, "last_active")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResp1874644RootSlice unmarshals a slice of Resp1874644Root instances from the specified list of maps.
func UnmarshalResp1874644RootSlice(s []interface{}) (slice []Resp1874644Root, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resp1874644Root'")
			return
		}
		obj, e := UnmarshalResp1874644Root(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResp1874644RootAsProperty unmarshals an instance of Resp1874644Root that is stored as a property
// within the specified map.
func UnmarshalResp1874644RootAsProperty(m map[string]interface{}, propertyName string) (result *Resp1874644Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resp1874644Root'", propertyName)
			return
		}
		result, err = UnmarshalResp1874644Root(objMap)
	}
	return
}

// UnmarshalResp1874644RootSliceAsProperty unmarshals a slice of Resp1874644Root instances that are stored as a property
// within the specified map.
func UnmarshalResp1874644RootSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resp1874644Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resp1874644Root'", propertyName)
			return
		}
		slice, err = UnmarshalResp1874644RootSlice(vSlice)
	}
	return
}

// Resp2079872Root : OK - MUST be returned if the service instance already exists, is fully provisioned, and the requested parameters are
// identical to the existing service instance.
type Resp2079872Root struct {
	// The URL of a web-based management user interface for the service instance; we refer to this as a service dashboard.
	// The URL MUST contain enough information for the dashboard to identify the resource being accessed. Note: a broker
	// that wishes to return `dashboard_url` for a service instance MUST return it with the initial response to the
	// provision request, even if the service is provisioned asynchronously. If present, it MUST be a non-empty string.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// For asynchronous responses, service brokers MAY return an identifier representing the operation. The value of this
	// field MUST be provided by the platform with requests to the `last_operation` endpoint in a URL encoded query
	// parameter. If present, MUST be a non-empty string.
	Operation *string `json:"operation,omitempty"`
}


// UnmarshalResp2079872Root constructs an instance of Resp2079872Root from the specified map.
func UnmarshalResp2079872Root(m map[string]interface{}) (result *Resp2079872Root, err error) {
	obj := new(Resp2079872Root)
	obj.DashboardURL, err = core.UnmarshalString(m, "dashboard_url")
	if err != nil {
		return
	}
	obj.Operation, err = core.UnmarshalString(m, "operation")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResp2079872RootSlice unmarshals a slice of Resp2079872Root instances from the specified list of maps.
func UnmarshalResp2079872RootSlice(s []interface{}) (slice []Resp2079872Root, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resp2079872Root'")
			return
		}
		obj, e := UnmarshalResp2079872Root(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResp2079872RootAsProperty unmarshals an instance of Resp2079872Root that is stored as a property
// within the specified map.
func UnmarshalResp2079872RootAsProperty(m map[string]interface{}, propertyName string) (result *Resp2079872Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resp2079872Root'", propertyName)
			return
		}
		result, err = UnmarshalResp2079872Root(objMap)
	}
	return
}

// UnmarshalResp2079872RootSliceAsProperty unmarshals a slice of Resp2079872Root instances that are stored as a property
// within the specified map.
func UnmarshalResp2079872RootSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resp2079872Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resp2079872Root'", propertyName)
			return
		}
		slice, err = UnmarshalResp2079872RootSlice(vSlice)
	}
	return
}

// Resp2079894Root : OK - MUST be returned upon successful processing of this request.
type Resp2079894Root struct {
	// A user-facing message displayed to the platform API client. Can be used to tell the user details about the status of
	// the operation. If present, MUST be a non-empty string.
	Description *string `json:"description,omitempty"`

	// Valid values are `in progress`, `succeeded`, and `failed`. While `â€ state": "in progressâ€ `, the platform SHOULD
	// continue polling. A response with `â€ state": "succeededâ€ ` or `â€ state": "failedâ€ ` MUST cause the platform to
	// cease polling.
	State *string `json:"state" validate:"required"`
}


// UnmarshalResp2079894Root constructs an instance of Resp2079894Root from the specified map.
func UnmarshalResp2079894Root(m map[string]interface{}) (result *Resp2079894Root, err error) {
	obj := new(Resp2079894Root)
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResp2079894RootSlice unmarshals a slice of Resp2079894Root instances from the specified list of maps.
func UnmarshalResp2079894RootSlice(s []interface{}) (slice []Resp2079894Root, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resp2079894Root'")
			return
		}
		obj, e := UnmarshalResp2079894Root(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResp2079894RootAsProperty unmarshals an instance of Resp2079894Root that is stored as a property
// within the specified map.
func UnmarshalResp2079894RootAsProperty(m map[string]interface{}, propertyName string) (result *Resp2079894Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resp2079894Root'", propertyName)
			return
		}
		result, err = UnmarshalResp2079894Root(objMap)
	}
	return
}

// UnmarshalResp2079894RootSliceAsProperty unmarshals a slice of Resp2079894Root instances that are stored as a property
// within the specified map.
func UnmarshalResp2079894RootSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resp2079894Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resp2079894Root'", propertyName)
			return
		}
		slice, err = UnmarshalResp2079894RootSlice(vSlice)
	}
	return
}

// Resp2448145Root : Check the enabled status of active service.
type Resp2448145Root struct {
	// Indicates (from the viewpoint of the provider) whether the service instance is active and is meaningful if `enabled`
	// is true.  The default value is true if not specified.
	Active *bool `json:"active,omitempty"`

	// Indicates the current state of the service instance.
	Enabled *bool `json:"enabled" validate:"required"`

	// Indicates when the service instance was last accessed or modified, and is meaningful if `enabled` is true AND
	// `active` is false.  Represented as milliseconds since the epoch, but does not need to be accurate to the
	// second/hour.
	LastActive *int64 `json:"last_active,omitempty"`
}


// UnmarshalResp2448145Root constructs an instance of Resp2448145Root from the specified map.
func UnmarshalResp2448145Root(m map[string]interface{}) (result *Resp2448145Root, err error) {
	obj := new(Resp2448145Root)
	obj.Active, err = core.UnmarshalBool(m, "active")
	if err != nil {
		return
	}
	obj.Enabled, err = core.UnmarshalBool(m, "enabled")
	if err != nil {
		return
	}
	obj.LastActive, err = core.UnmarshalInt64(m, "last_active")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalResp2448145RootSlice unmarshals a slice of Resp2448145Root instances from the specified list of maps.
func UnmarshalResp2448145RootSlice(s []interface{}) (slice []Resp2448145Root, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Resp2448145Root'")
			return
		}
		obj, e := UnmarshalResp2448145Root(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalResp2448145RootAsProperty unmarshals an instance of Resp2448145Root that is stored as a property
// within the specified map.
func UnmarshalResp2448145RootAsProperty(m map[string]interface{}, propertyName string) (result *Resp2448145Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Resp2448145Root'", propertyName)
			return
		}
		result, err = UnmarshalResp2448145Root(objMap)
	}
	return
}

// UnmarshalResp2448145RootSliceAsProperty unmarshals a slice of Resp2448145Root instances that are stored as a property
// within the specified map.
func UnmarshalResp2448145RootSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Resp2448145Root, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Resp2448145Root'", propertyName)
			return
		}
		slice, err = UnmarshalResp2448145RootSlice(vSlice)
	}
	return
}

// UpdateServiceInstanceOptions : The UpdateServiceInstance options.
type UpdateServiceInstanceOptions struct {
	// The ID of a previously provisioned service instance.
	InstanceID *string `json:"instance_id" validate:"required"`

	// Contextual data under which the service instance is created.
	Context []Context `json:"context,omitempty"`

	// Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
	// ensure that the client has provided valid configuration parameters and values for the operation. If this field is
	// not present in the request message, then the broker MUST NOT change the parameters of the instance as a result of
	// this request.
	Parameters *Parameters `json:"parameters,omitempty"`

	// The ID of the plan for which the service instance has been requested, which is stored in the catalog.json of your
	// broker. This value should be a GUID. MUST be unique to a service. If present, MUST be a non-empty string. If this
	// field is not present in the request message, then the broker MUST NOT change the plan of the instance as a result of
	// this request.
	PlanID *string `json:"plan_id,omitempty"`

	// Information about the service instance prior to the update.
	PreviousValues []string `json:"previous_values,omitempty"`

	// The ID of the service stored in the catalog.json of your broker. This value should be a GUID. It MUST be a non-empty
	// string.
	ServiceID *string `json:"service_id,omitempty"`

	// A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous
	// deprovisioning. If this parameter is not included in the request, and the broker can only deprovision a service
	// instance of the requested plan asynchronously, the broker MUST reject the request with a `422` Unprocessable Entity.
	AcceptsIncomplete *string `json:"accepts_incomplete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateServiceInstanceOptions : Instantiate UpdateServiceInstanceOptions
func (*OpenServiceBrokerV1) NewUpdateServiceInstanceOptions(instanceID string) *UpdateServiceInstanceOptions {
	return &UpdateServiceInstanceOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (options *UpdateServiceInstanceOptions) SetInstanceID(instanceID string) *UpdateServiceInstanceOptions {
	options.InstanceID = core.StringPtr(instanceID)
	return options
}

// SetContext : Allow user to set Context
func (options *UpdateServiceInstanceOptions) SetContext(context []Context) *UpdateServiceInstanceOptions {
	options.Context = context
	return options
}

// SetParameters : Allow user to set Parameters
func (options *UpdateServiceInstanceOptions) SetParameters(parameters *Parameters) *UpdateServiceInstanceOptions {
	options.Parameters = parameters
	return options
}

// SetPlanID : Allow user to set PlanID
func (options *UpdateServiceInstanceOptions) SetPlanID(planID string) *UpdateServiceInstanceOptions {
	options.PlanID = core.StringPtr(planID)
	return options
}

// SetPreviousValues : Allow user to set PreviousValues
func (options *UpdateServiceInstanceOptions) SetPreviousValues(previousValues []string) *UpdateServiceInstanceOptions {
	options.PreviousValues = previousValues
	return options
}

// SetServiceID : Allow user to set ServiceID
func (options *UpdateServiceInstanceOptions) SetServiceID(serviceID string) *UpdateServiceInstanceOptions {
	options.ServiceID = core.StringPtr(serviceID)
	return options
}

// SetAcceptsIncomplete : Allow user to set AcceptsIncomplete
func (options *UpdateServiceInstanceOptions) SetAcceptsIncomplete(acceptsIncomplete string) *UpdateServiceInstanceOptions {
	options.AcceptsIncomplete = core.StringPtr(acceptsIncomplete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateServiceInstanceOptions) SetHeaders(param map[string]string) *UpdateServiceInstanceOptions {
	options.Headers = param
	return options
}

// BindResource : Bind a resource.
type BindResource struct {
	// Account owner of resource to bind.
	AccountID *string `json:"account_id,omitempty"`

	// Service ID of resource to bind.
	ServiceidCrn *string `json:"serviceid_crn,omitempty"`

	// Target ID of resource to bind.
	TargetCrn *string `json:"target_crn,omitempty"`
}


// UnmarshalBindResource constructs an instance of BindResource from the specified map.
func UnmarshalBindResource(m map[string]interface{}) (result *BindResource, err error) {
	obj := new(BindResource)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.ServiceidCrn, err = core.UnmarshalString(m, "serviceid_crn")
	if err != nil {
		return
	}
	obj.TargetCrn, err = core.UnmarshalString(m, "target_crn")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalBindResourceSlice unmarshals a slice of BindResource instances from the specified list of maps.
func UnmarshalBindResourceSlice(s []interface{}) (slice []BindResource, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'BindResource'")
			return
		}
		obj, e := UnmarshalBindResource(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalBindResourceAsProperty unmarshals an instance of BindResource that is stored as a property
// within the specified map.
func UnmarshalBindResourceAsProperty(m map[string]interface{}, propertyName string) (result *BindResource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'BindResource'", propertyName)
			return
		}
		result, err = UnmarshalBindResource(objMap)
	}
	return
}

// UnmarshalBindResourceSliceAsProperty unmarshals a slice of BindResource instances that are stored as a property
// within the specified map.
func UnmarshalBindResourceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []BindResource, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'BindResource'", propertyName)
			return
		}
		slice, err = UnmarshalBindResourceSlice(vSlice)
	}
	return
}

// Context : Contextual data under which the service instance is created.
type Context struct {
	// Returns the ID of the account in IBM Cloud that is provisioning the service instance.
	AccountID *string `json:"account_id,omitempty"`

	// When a customer provisions your service in IBM Cloud, a service instance is created and this instance is identified
	// by its IBM Cloud Resource Name (CRN). The CRN is utilized in all aspects of the interaction with IBM Cloud including
	// provisioning, binding (creating credentials and endpoints), metering, dashboard display, and access control. From a
	// service provider perspective, the CRN can largely be treated as an opaque string to be utilized with the IBM Cloud
	// APIs, but it can also be decomposed via the following structure:
	// `crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource`.
	Crn *string `json:"crn,omitempty"`

	// Identifies the platform as "ibmcloud".
	Platform *string `json:"platform,omitempty"`
}


// UnmarshalContext constructs an instance of Context from the specified map.
func UnmarshalContext(m map[string]interface{}) (result *Context, err error) {
	obj := new(Context)
	obj.AccountID, err = core.UnmarshalString(m, "account_id")
	if err != nil {
		return
	}
	obj.Crn, err = core.UnmarshalString(m, "crn")
	if err != nil {
		return
	}
	obj.Platform, err = core.UnmarshalString(m, "platform")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalContextSlice unmarshals a slice of Context instances from the specified list of maps.
func UnmarshalContextSlice(s []interface{}) (slice []Context, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Context'")
			return
		}
		obj, e := UnmarshalContext(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalContextAsProperty unmarshals an instance of Context that is stored as a property
// within the specified map.
func UnmarshalContextAsProperty(m map[string]interface{}, propertyName string) (result *Context, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Context'", propertyName)
			return
		}
		result, err = UnmarshalContext(objMap)
	}
	return
}

// UnmarshalContextSliceAsProperty unmarshals a slice of Context instances that are stored as a property
// within the specified map.
func UnmarshalContextSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Context, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Context'", propertyName)
			return
		}
		slice, err = UnmarshalContextSlice(vSlice)
	}
	return
}

// Parameters : Configuration options for the service instance. An opaque object, controller treats this as a blob. Brokers should
// ensure that the client has provided valid configuration parameters and values for the operation. If this field is not
// present in the request message, then the broker MUST NOT change the parameters of the instance as a result of this
// request.
type Parameters struct {
	// a custom integer or string within the parameters JSON object.
	Parameter1 *int64 `json:"parameter1,omitempty"`

	// a custom integer or string within the parameters JSON object.
	Parameter2 *string `json:"parameter2,omitempty"`
}


// UnmarshalParameters constructs an instance of Parameters from the specified map.
func UnmarshalParameters(m map[string]interface{}) (result *Parameters, err error) {
	obj := new(Parameters)
	obj.Parameter1, err = core.UnmarshalInt64(m, "parameter1")
	if err != nil {
		return
	}
	obj.Parameter2, err = core.UnmarshalString(m, "parameter2")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalParametersSlice unmarshals a slice of Parameters instances from the specified list of maps.
func UnmarshalParametersSlice(s []interface{}) (slice []Parameters, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Parameters'")
			return
		}
		obj, e := UnmarshalParameters(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalParametersAsProperty unmarshals an instance of Parameters that is stored as a property
// within the specified map.
func UnmarshalParametersAsProperty(m map[string]interface{}, propertyName string) (result *Parameters, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Parameters'", propertyName)
			return
		}
		result, err = UnmarshalParameters(objMap)
	}
	return
}

// UnmarshalParametersSliceAsProperty unmarshals a slice of Parameters instances that are stored as a property
// within the specified map.
func UnmarshalParametersSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Parameters, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Parameters'", propertyName)
			return
		}
		slice, err = UnmarshalParametersSlice(vSlice)
	}
	return
}

// Plans : Where is this in the source?.
type Plans struct {
	// A short description of the plan. It MUST be a non-empty string. The description is NOT displayed in the IBM Cloud
	// catalog or IBM Cloud CLI.
	Description *string `json:"description" validate:"required"`

	// When false, service instances of this plan have a cost. The default is true.
	Free *bool `json:"free,omitempty"`

	// An identifier used to correlate this plan in future requests to the broker.  This MUST be globally unique within a
	// platform marketplace. It MUST be a non-empty string and using a GUID is RECOMMENDED. If you define your service in
	// the RMC, it will create a unique GUID for you to use. It is recommended to use the RMC to define and generate these
	// values and then use them in your catalog.json metadata in your broker. This value is NOT displayed in the IBM Cloud
	// catalog or IBM Cloud CLI.
	ID *string `json:"id" validate:"required"`

	// The programmatic name of the plan. It MUST be unique within the service. All lowercase, no spaces. It MUST be a
	// non-empty string, and it's NOT displayed in the IBM Cloud catalog or IBM Cloud CLI.
	Name *string `json:"name" validate:"required"`
}


// UnmarshalPlans constructs an instance of Plans from the specified map.
func UnmarshalPlans(m map[string]interface{}) (result *Plans, err error) {
	obj := new(Plans)
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Free, err = core.UnmarshalBool(m, "free")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPlansSlice unmarshals a slice of Plans instances from the specified list of maps.
func UnmarshalPlansSlice(s []interface{}) (slice []Plans, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Plans'")
			return
		}
		obj, e := UnmarshalPlans(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPlansAsProperty unmarshals an instance of Plans that is stored as a property
// within the specified map.
func UnmarshalPlansAsProperty(m map[string]interface{}, propertyName string) (result *Plans, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Plans'", propertyName)
			return
		}
		result, err = UnmarshalPlans(objMap)
	}
	return
}

// UnmarshalPlansSliceAsProperty unmarshals a slice of Plans instances that are stored as a property
// within the specified map.
func UnmarshalPlansSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Plans, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Plans'", propertyName)
			return
		}
		slice, err = UnmarshalPlansSlice(vSlice)
	}
	return
}

// Services : The service object that describes the properties of your service.
type Services struct {
	// Specifies whether or not your service can be bound to applications in IBM Cloud. If bindable, it must be able to
	// return API endpoints and credentials to your service consumers.
	Bindable *bool `json:"bindable" validate:"required"`

	// A short description of the service. It MUST be a non-empty string. Note that this description is not displayed by
	// the the IBM Cloud console or IBM Cloud CLI.
	Description *string `json:"description" validate:"required"`

	// An identifier used to correlate this service in future requests to the broker. This MUST be globally unique within
	// the IBM Cloud platform. It MUST be a non-empty string, and using a GUID is recommended. Recommended: If you define
	// your service in the RMC, the RMC will generate a globally unique GUID service ID that you can use in your service
	// broker.
	ID *string `json:"id" validate:"required"`

	// The service name is not your display name. Your service name must follow the follow these rules:
	//  - It must be all lowercase.
	//  - It can't include spaces but may include hyphens (`-`).
	//  - It must be less than 32 characters.
	//  Your service name should include your company name. If your company has more then one offering your service name
	// should include both company and offering as part of the name. For example, the Compose company has offerings for
	// Redis and Elasticsearch. Sample service names on IBM Cloud for these offerings would be `compose-redis` and
	// `compose-elasticsearch`.  Each of these service names have associated display names that are shown in the IBM Cloud
	// catalog: *Compose Redis* and *Compose Elasticsearch*. Another company (e.g. FastJetMail) may only have the single
	// JetMail offering, in which case the service name should be `fastjetmail`. Recommended: If you define your service in
	// RMC, you can export a catalog.json that will include the service name you defined within the RMC.
	Name *string `json:"name" validate:"required"`

	// The Default is false. This specifices whether or not you support plan changes for provisioned instances. If your
	// offering supports multiple plans, and you want users to be able to change plans for a provisioned instance, you will
	// need to enable the ability for users to update their service instance by using /v2/service_instances/{instance_id}
	// PATCH.
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`

	// A list of plans for this service that must contain at least one plan.
	Plans []Plans `json:"plans" validate:"required"`
}


// UnmarshalServices constructs an instance of Services from the specified map.
func UnmarshalServices(m map[string]interface{}) (result *Services, err error) {
	obj := new(Services)
	obj.Bindable, err = core.UnmarshalBool(m, "bindable")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.PlanUpdateable, err = core.UnmarshalBool(m, "plan_updateable")
	if err != nil {
		return
	}
	obj.Plans, err = UnmarshalPlansSliceAsProperty(m, "plans")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalServicesSlice unmarshals a slice of Services instances from the specified list of maps.
func UnmarshalServicesSlice(s []interface{}) (slice []Services, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Services'")
			return
		}
		obj, e := UnmarshalServices(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalServicesAsProperty unmarshals an instance of Services that is stored as a property
// within the specified map.
func UnmarshalServicesAsProperty(m map[string]interface{}, propertyName string) (result *Services, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Services'", propertyName)
			return
		}
		result, err = UnmarshalServices(objMap)
	}
	return
}

// UnmarshalServicesSliceAsProperty unmarshals a slice of Services instances that are stored as a property
// within the specified map.
func UnmarshalServicesSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Services, err error) {
	v, foundIt := m[propertyName]
	if foundIt {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Services'", propertyName)
			return
		}
		slice, err = UnmarshalServicesSlice(vSlice)
	}
	return
}
