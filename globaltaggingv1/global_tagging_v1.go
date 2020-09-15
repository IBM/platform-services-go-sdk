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
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-ef5e13c2-20200915-144510
 */
 

// Package globaltaggingv1 : Operations and models for the GlobalTaggingV1 service
package globaltaggingv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"reflect"
	"strings"
)

// GlobalTaggingV1 : Manage your tags with the Tagging API in IBM Cloud. You can attach, detach, delete a tag or list
// all tags in your billing account with the Tagging API. The tag name must be unique within a billing account. You can
// create tags in two formats: `key:value` or `label`.
//
// Version: 1.0.3
type GlobalTaggingV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://tags.global-search-tagging.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "global_tagging"

// GlobalTaggingV1Options : Service options
type GlobalTaggingV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewGlobalTaggingV1UsingExternalConfig : constructs an instance of GlobalTaggingV1 with passed in options and external configuration.
func NewGlobalTaggingV1UsingExternalConfig(options *GlobalTaggingV1Options) (globalTagging *GlobalTaggingV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	globalTagging, err = NewGlobalTaggingV1(options)
	if err != nil {
		return
	}

	err = globalTagging.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = globalTagging.Service.SetServiceURL(options.URL)
	}
	return
}

// NewGlobalTaggingV1 : constructs an instance of GlobalTaggingV1 with passed in options.
func NewGlobalTaggingV1(options *GlobalTaggingV1Options) (service *GlobalTaggingV1, err error) {
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

	service = &GlobalTaggingV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (globalTagging *GlobalTaggingV1) SetServiceURL(url string) error {
	return globalTagging.Service.SetServiceURL(url)
}

// ListTags : Get all tags
// Lists all tags in a billing account. Use the `attached_to` parameter to return the list of tags attached to the
// specified resource.
func (globalTagging *GlobalTaggingV1) ListTags(listTagsOptions *ListTagsOptions) (result *TagList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listTagsOptions, "listTagsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ResolveRequestURL(globalTagging.Service.Options.URL, `/v3/tags`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listTagsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_tagging", "V1", "ListTags")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listTagsOptions.Providers != nil {
		builder.AddQuery("providers", strings.Join(listTagsOptions.Providers, ","))
	}
	if listTagsOptions.AttachedTo != nil {
		builder.AddQuery("attached_to", fmt.Sprint(*listTagsOptions.AttachedTo))
	}
	if listTagsOptions.FullData != nil {
		builder.AddQuery("full_data", fmt.Sprint(*listTagsOptions.FullData))
	}
	if listTagsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listTagsOptions.Offset))
	}
	if listTagsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listTagsOptions.Limit))
	}
	if listTagsOptions.OrderByName != nil {
		builder.AddQuery("order_by_name", fmt.Sprint(*listTagsOptions.OrderByName))
	}
	if listTagsOptions.Timeout != nil {
		builder.AddQuery("timeout", fmt.Sprint(*listTagsOptions.Timeout))
	}
	if listTagsOptions.AttachedOnly != nil {
		builder.AddQuery("attached_only", fmt.Sprint(*listTagsOptions.AttachedOnly))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalTagging.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTagList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteTagAll : Delete unused tags
// Delete the tags that are not attatched to any resource.
func (globalTagging *GlobalTaggingV1) DeleteTagAll(deleteTagAllOptions *DeleteTagAllOptions) (result *DeleteTagsResult, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(deleteTagAllOptions, "deleteTagAllOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(globalTagging.Service.Options.URL, `/v3/tags`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTagAllOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_tagging", "V1", "DeleteTagAll")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if deleteTagAllOptions.Providers != nil {
		builder.AddQuery("providers", fmt.Sprint(*deleteTagAllOptions.Providers))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalTagging.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteTagsResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteTag : Delete a tag
// Delete an existing tag. A tag can be deleted only if it is not attached to any resource.
func (globalTagging *GlobalTaggingV1) DeleteTag(deleteTagOptions *DeleteTagOptions) (result *DeleteTagResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTagOptions, "deleteTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTagOptions, "deleteTagOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tag_name": *deleteTagOptions.TagName,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ResolveRequestURL(globalTagging.Service.Options.URL, `/v3/tags/{tag_name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_tagging", "V1", "DeleteTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if deleteTagOptions.Providers != nil {
		builder.AddQuery("providers", strings.Join(deleteTagOptions.Providers, ","))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalTagging.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteTagResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AttachTag : Attach one or more tags
// Attaches one or more tags to one or more resources. To attach a tag to a resource managed by the Resource Controller,
// you must be an editor on the resource. To attach a tag to a Cloud Foundry resource, you must have space developer
// role. To attach a tag to IMS resources, depending on the resource, you need either `view hardware details`, `view
// virtual server details` or `manage storage` permission.
func (globalTagging *GlobalTaggingV1) AttachTag(attachTagOptions *AttachTagOptions) (result *TagResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(attachTagOptions, "attachTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(attachTagOptions, "attachTagOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(globalTagging.Service.Options.URL, `/v3/tags/attach`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range attachTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_tagging", "V1", "AttachTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if attachTagOptions.Resources != nil {
		body["resources"] = attachTagOptions.Resources
	}
	if attachTagOptions.TagName != nil {
		body["tag_name"] = attachTagOptions.TagName
	}
	if attachTagOptions.TagNames != nil {
		body["tag_names"] = attachTagOptions.TagNames
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
	response, err = globalTagging.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTagResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DetachTag : Detach one or more tags
// Detach one or more tags from one or more resources. To detach a tag from a Resource Controller managed resource, you
// must be an editor on the resource. To detach a tag to a Cloud Foundry resource, you must have `space developer` role.
//   To detach a tag to IMS resources, depending on the resource, you need either `view hardware details`, `view virtual
// server details` or `storage manage` permission.
func (globalTagging *GlobalTaggingV1) DetachTag(detachTagOptions *DetachTagOptions) (result *TagResults, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(detachTagOptions, "detachTagOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(detachTagOptions, "detachTagOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ResolveRequestURL(globalTagging.Service.Options.URL, `/v3/tags/detach`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range detachTagOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_tagging", "V1", "DetachTag")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if detachTagOptions.Resources != nil {
		body["resources"] = detachTagOptions.Resources
	}
	if detachTagOptions.TagName != nil {
		body["tag_name"] = detachTagOptions.TagName
	}
	if detachTagOptions.TagNames != nil {
		body["tag_names"] = detachTagOptions.TagNames
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
	response, err = globalTagging.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTagResults)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// AttachTagOptions : The AttachTag options.
type AttachTagOptions struct {
	// List of resources on which the tag or tags should be attached.
	Resources []Resource `json:"resources" validate:"required"`

	// The name of the tag to attach.
	TagName *string `json:"tag_name,omitempty"`

	// An array of tag names to attach.
	TagNames []string `json:"tag_names,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAttachTagOptions : Instantiate AttachTagOptions
func (*GlobalTaggingV1) NewAttachTagOptions(resources []Resource) *AttachTagOptions {
	return &AttachTagOptions{
		Resources: resources,
	}
}

// SetResources : Allow user to set Resources
func (options *AttachTagOptions) SetResources(resources []Resource) *AttachTagOptions {
	options.Resources = resources
	return options
}

// SetTagName : Allow user to set TagName
func (options *AttachTagOptions) SetTagName(tagName string) *AttachTagOptions {
	options.TagName = core.StringPtr(tagName)
	return options
}

// SetTagNames : Allow user to set TagNames
func (options *AttachTagOptions) SetTagNames(tagNames []string) *AttachTagOptions {
	options.TagNames = tagNames
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AttachTagOptions) SetHeaders(param map[string]string) *AttachTagOptions {
	options.Headers = param
	return options
}

// DeleteTagAllOptions : The DeleteTagAll options.
type DeleteTagAllOptions struct {
	// Select a provider. Supported values are `ghost` and `ims`.
	Providers *string `json:"providers,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteTagAllOptions.Providers property.
// Select a provider. Supported values are `ghost` and `ims`.
const (
	DeleteTagAllOptions_Providers_Ghost = "ghost"
	DeleteTagAllOptions_Providers_Ims = "ims"
)

// NewDeleteTagAllOptions : Instantiate DeleteTagAllOptions
func (*GlobalTaggingV1) NewDeleteTagAllOptions() *DeleteTagAllOptions {
	return &DeleteTagAllOptions{}
}

// SetProviders : Allow user to set Providers
func (options *DeleteTagAllOptions) SetProviders(providers string) *DeleteTagAllOptions {
	options.Providers = core.StringPtr(providers)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTagAllOptions) SetHeaders(param map[string]string) *DeleteTagAllOptions {
	options.Headers = param
	return options
}

// DeleteTagOptions : The DeleteTag options.
type DeleteTagOptions struct {
	// The name of tag to be deleted.
	TagName *string `json:"tag_name" validate:"required,ne="`

	// Select a provider. Supported values are `ghost` and `ims`. To delete tag both in GhoST in IMS, use `ghost,ims`.
	Providers []string `json:"providers,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteTagOptions.Providers property.
const (
	DeleteTagOptions_Providers_Ghost = "ghost"
	DeleteTagOptions_Providers_Ims = "ims"
)

// NewDeleteTagOptions : Instantiate DeleteTagOptions
func (*GlobalTaggingV1) NewDeleteTagOptions(tagName string) *DeleteTagOptions {
	return &DeleteTagOptions{
		TagName: core.StringPtr(tagName),
	}
}

// SetTagName : Allow user to set TagName
func (options *DeleteTagOptions) SetTagName(tagName string) *DeleteTagOptions {
	options.TagName = core.StringPtr(tagName)
	return options
}

// SetProviders : Allow user to set Providers
func (options *DeleteTagOptions) SetProviders(providers []string) *DeleteTagOptions {
	options.Providers = providers
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTagOptions) SetHeaders(param map[string]string) *DeleteTagOptions {
	options.Headers = param
	return options
}

// DeleteTagResults : Results of a delete_tag request.
type DeleteTagResults struct {
	// Array of results of a delete_tag request.
	Results []DeleteTagResultsItem `json:"results,omitempty"`
}


// UnmarshalDeleteTagResults unmarshals an instance of DeleteTagResults from the specified map of raw messages.
func UnmarshalDeleteTagResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteTagResults)
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalDeleteTagResultsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteTagResultsItem : Result of a delete_tag request.
type DeleteTagResultsItem struct {
	// The provider of the tag.
	Provider *string `json:"provider,omitempty"`

	// It is `true` if the operation exits with an error.
	IsError *bool `json:"is_error,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// Constants associated with the DeleteTagResultsItem.Provider property.
// The provider of the tag.
const (
	DeleteTagResultsItem_Provider_Ghost = "ghost"
	DeleteTagResultsItem_Provider_Ims = "ims"
)


// SetProperty allows the user to set an arbitrary property on an instance of DeleteTagResultsItem
func (o *DeleteTagResultsItem) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of DeleteTagResultsItem
func (o *DeleteTagResultsItem) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of DeleteTagResultsItem
func (o *DeleteTagResultsItem) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of DeleteTagResultsItem
func (o *DeleteTagResultsItem) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Provider != nil {
		m["provider"] = o.Provider
	}
	if o.IsError != nil {
		m["is_error"] = o.IsError
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalDeleteTagResultsItem unmarshals an instance of DeleteTagResultsItem from the specified map of raw messages.
func UnmarshalDeleteTagResultsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteTagResultsItem)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	delete(m, "provider")
	err = core.UnmarshalPrimitive(m, "is_error", &obj.IsError)
	if err != nil {
		return
	}
	delete(m, "is_error")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteTagsResult : The results of a deleting unattatched tags.
type DeleteTagsResult struct {
	// The number of tags deleted in the account.
	TotalCount *int64 `json:"total_count,omitempty"`

	// An indicator that is set to true if there was an error deleting some of the tags.
	Errors *bool `json:"errors,omitempty"`

	Items []DeleteTagsResultItem `json:"items,omitempty"`
}


// UnmarshalDeleteTagsResult unmarshals an instance of DeleteTagsResult from the specified map of raw messages.
func UnmarshalDeleteTagsResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteTagsResult)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalDeleteTagsResultItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteTagsResultItem : Result of deleting one unattached tag.
type DeleteTagsResultItem struct {
	// The name of the tag that was deleted.
	TagName *string `json:"tag_name,omitempty"`

	// An indicator that is set to true if there was an error deleting the tag.
	IsError *bool `json:"is_error,omitempty"`
}


// UnmarshalDeleteTagsResultItem unmarshals an instance of DeleteTagsResultItem from the specified map of raw messages.
func UnmarshalDeleteTagsResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteTagsResultItem)
	err = core.UnmarshalPrimitive(m, "tag_name", &obj.TagName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_error", &obj.IsError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DetachTagOptions : The DetachTag options.
type DetachTagOptions struct {
	// List of resources on which the tag or tags should be detached.
	Resources []Resource `json:"resources" validate:"required"`

	// The name of the tag to detach.
	TagName *string `json:"tag_name,omitempty"`

	// An array of tag names to detach.
	TagNames []string `json:"tag_names,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDetachTagOptions : Instantiate DetachTagOptions
func (*GlobalTaggingV1) NewDetachTagOptions(resources []Resource) *DetachTagOptions {
	return &DetachTagOptions{
		Resources: resources,
	}
}

// SetResources : Allow user to set Resources
func (options *DetachTagOptions) SetResources(resources []Resource) *DetachTagOptions {
	options.Resources = resources
	return options
}

// SetTagName : Allow user to set TagName
func (options *DetachTagOptions) SetTagName(tagName string) *DetachTagOptions {
	options.TagName = core.StringPtr(tagName)
	return options
}

// SetTagNames : Allow user to set TagNames
func (options *DetachTagOptions) SetTagNames(tagNames []string) *DetachTagOptions {
	options.TagNames = tagNames
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DetachTagOptions) SetHeaders(param map[string]string) *DetachTagOptions {
	options.Headers = param
	return options
}

// ListTagsOptions : The ListTags options.
type ListTagsOptions struct {
	// Select a provider. Supported values are `ghost` and `ims`. To list GhoST tags and infrastructure tags use
	// `ghost,ims`.
	Providers []string `json:"providers,omitempty"`

	// If you want to return only the list of tags attached to a specified resource, pass here the ID of the resource. For
	// GhoST onboarded resources, the resource ID is the CRN; for IMS resources, it is the IMS ID. When using this
	// parameter it is mandatory to specify the appropriate provider (`ims` or `ghost`).
	AttachedTo *string `json:"attached_to,omitempty"`

	// If set to `true`, this query returns the provider, `ghost`, `ims` or `ghost,ims`, where the tag exists and the
	// number of attached resources.
	FullData *bool `json:"full_data,omitempty"`

	// The offset is the index of the item from which you want to start returning data from.
	Offset *int64 `json:"offset,omitempty"`

	// The number of tags to return.
	Limit *int64 `json:"limit,omitempty"`

	// Order the output by tag name.
	OrderByName *string `json:"order_by_name,omitempty"`

	// The search timeout bounds the search request to be executed within the specified time value. It returns the hits
	// accumulated until time runs out.
	Timeout *int64 `json:"timeout,omitempty"`

	// Filter on attached tags. If true, returns only tags that are attached to one or more resources. If false returns all
	// tags.
	AttachedOnly *bool `json:"attached_only,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListTagsOptions.Providers property.
const (
	ListTagsOptions_Providers_Ghost = "ghost"
	ListTagsOptions_Providers_Ims = "ims"
)

// Constants associated with the ListTagsOptions.OrderByName property.
// Order the output by tag name.
const (
	ListTagsOptions_OrderByName_Asc = "asc"
	ListTagsOptions_OrderByName_Desc = "desc"
)

// NewListTagsOptions : Instantiate ListTagsOptions
func (*GlobalTaggingV1) NewListTagsOptions() *ListTagsOptions {
	return &ListTagsOptions{}
}

// SetProviders : Allow user to set Providers
func (options *ListTagsOptions) SetProviders(providers []string) *ListTagsOptions {
	options.Providers = providers
	return options
}

// SetAttachedTo : Allow user to set AttachedTo
func (options *ListTagsOptions) SetAttachedTo(attachedTo string) *ListTagsOptions {
	options.AttachedTo = core.StringPtr(attachedTo)
	return options
}

// SetFullData : Allow user to set FullData
func (options *ListTagsOptions) SetFullData(fullData bool) *ListTagsOptions {
	options.FullData = core.BoolPtr(fullData)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListTagsOptions) SetOffset(offset int64) *ListTagsOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *ListTagsOptions) SetLimit(limit int64) *ListTagsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOrderByName : Allow user to set OrderByName
func (options *ListTagsOptions) SetOrderByName(orderByName string) *ListTagsOptions {
	options.OrderByName = core.StringPtr(orderByName)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *ListTagsOptions) SetTimeout(timeout int64) *ListTagsOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetAttachedOnly : Allow user to set AttachedOnly
func (options *ListTagsOptions) SetAttachedOnly(attachedOnly bool) *ListTagsOptions {
	options.AttachedOnly = core.BoolPtr(attachedOnly)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListTagsOptions) SetHeaders(param map[string]string) *ListTagsOptions {
	options.Headers = param
	return options
}

// Resource : A resource that may have attached tags.
type Resource struct {
	// The CRN or IMS ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// The IMS resource type of the resource.
	ResourceType *string `json:"resource_type,omitempty"`
}


// NewResource : Instantiate Resource (Generic Model Constructor)
func (*GlobalTaggingV1) NewResource(resourceID string) (model *Resource, err error) {
	model = &Resource{
		ResourceID: core.StringPtr(resourceID),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_type", &obj.ResourceType)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Tag : A tag.
type Tag struct {
	// This is the name of the tag.
	Name *string `json:"name" validate:"required"`
}


// UnmarshalTag unmarshals an instance of Tag from the specified map of raw messages.
func UnmarshalTag(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Tag)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TagList : A list of tags.
type TagList struct {
	// The number of tags defined in the account.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The offset specific at input time.
	Offset *int64 `json:"offset,omitempty"`

	// The limit specified at input time.
	Limit *int64 `json:"limit,omitempty"`

	// This is an array of output results.
	Items []Tag `json:"items,omitempty"`
}


// UnmarshalTagList unmarshals an instance of TagList from the specified map of raw messages.
func UnmarshalTagList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TagList)
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalTag)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TagResults : Results of an attach_tag or detach_tag request.
type TagResults struct {
	// Array of results of an attach_tag or detach_tag request.
	Results []TagResultsItem `json:"results,omitempty"`
}


// UnmarshalTagResults unmarshals an instance of TagResults from the specified map of raw messages.
func UnmarshalTagResults(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TagResults)
	err = core.UnmarshalModel(m, "results", &obj.Results, UnmarshalTagResultsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TagResultsItem : Result of an attach_tag or detach_tag request for a tagged resource.
type TagResultsItem struct {
	// The CRN or IMS ID of the resource.
	ResourceID *string `json:"resource_id" validate:"required"`

	// It is `true` if the operation exits with an error.
	IsError *bool `json:"is_error,omitempty"`
}


// UnmarshalTagResultsItem unmarshals an instance of TagResultsItem from the specified map of raw messages.
func UnmarshalTagResultsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TagResultsItem)
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "is_error", &obj.IsError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
