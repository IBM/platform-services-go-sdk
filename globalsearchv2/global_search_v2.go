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

// Package globalsearchv2 : Operations and models for the GlobalSearchV2 service
package globalsearchv2

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"reflect"
	"strings"
)

// GlobalSearchV2 : Search for resources with the global and shared resource properties repository integrated in the IBM
// Cloud Platform. The search repository stores and searches cloud resources attributes, which categorize or classify
// resources. A resource is a physical or logical component that can be provisioned or reserved for an application or
// service instance and is owned by resource providers, such as Cloud Foundry, IBM containers, or Resource Controller,
// in the IBM Cloud platform. Resources are uniquely identified by a CRN (Cloud Resource Naming identifier) or by an IMS
// ID. The properties of a resource include tags and system properties. Both properties are defined in an IBM Cloud
// billing account, and span across many regions.
//
// Version: 2.0.1
type GlobalSearchV2 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.global-search-tagging.cloud.ibm.com/"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "global_search"

// GlobalSearchV2Options : Service options
type GlobalSearchV2Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewGlobalSearchV2UsingExternalConfig : constructs an instance of GlobalSearchV2 with passed in options and external configuration.
func NewGlobalSearchV2UsingExternalConfig(options *GlobalSearchV2Options) (globalSearch *GlobalSearchV2, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	globalSearch, err = NewGlobalSearchV2(options)
	if err != nil {
		return
	}

	err = globalSearch.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = globalSearch.Service.SetServiceURL(options.URL)
	}
	return
}

// NewGlobalSearchV2 : constructs an instance of GlobalSearchV2 with passed in options.
func NewGlobalSearchV2(options *GlobalSearchV2Options) (service *GlobalSearchV2, err error) {
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

	service = &GlobalSearchV2{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (globalSearch *GlobalSearchV2) SetServiceURL(url string) error {
	return globalSearch.Service.SetServiceURL(url)
}

// Search : Find instances of resources
// 'Find cloud foundry resources, resource controlled enabled resources, or storage and network resources running on
// classic infrastructure in a specific account ID. You can apply query strings if necessary. To filter results, you can
// insert a string using the Lucene syntax and the query string is parsed into a series of terms and operators. A term
// can be a single word or a phrase, in which case the search is performed for all the words, in the same order. To
// filter for a specific value regardless of the property that contains it, use an asterisk as the key name. Only
// resources that belong to the account ID and that are accessible by the client are returned. You must use this
// operation when you need to fetch more than `10000` resource items. The `/v2/resources/search` prohibits paginating
// through such a big number. On the first call, the operation returns a live cursor on the data that you must use on
// all the subsequent calls to get the next batch of results until you get the empty result set. By default, the fields
// returned for every resources are: "crn", "name", "family", "type", "account_id". You can specify the subset of the
// fields you want in your request.''.
func (globalSearch *GlobalSearchV2) Search(searchOptions *SearchOptions) (result *ScanResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(searchOptions, "searchOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(searchOptions, "searchOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v3/resources/search"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(globalSearch.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range searchOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_search", "V2", "Search")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if searchOptions.TransactionID != nil {
		builder.AddHeader("transaction-id", fmt.Sprint(*searchOptions.TransactionID))
	}

	if searchOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*searchOptions.AccountID))
	}
	if searchOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*searchOptions.Limit))
	}
	if searchOptions.Timeout != nil {
		builder.AddQuery("timeout", fmt.Sprint(*searchOptions.Timeout))
	}
	if searchOptions.Sort != nil {
		builder.AddQuery("sort", strings.Join(searchOptions.Sort, ","))
	}

	body := make(map[string]interface{})
	if searchOptions.Query != nil {
		body["query"] = searchOptions.Query
	}
	if searchOptions.Fields != nil {
		body["fields"] = searchOptions.Fields
	}
	if searchOptions.SearchCursor != nil {
		body["search_cursor"] = searchOptions.SearchCursor
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
	response, err = globalSearch.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalScanResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSupportedTypes : Get all supported resource types
// Retrieves a list of all the resource types supported by GhoST.
func (globalSearch *GlobalSearchV2) GetSupportedTypes(getSupportedTypesOptions *GetSupportedTypesOptions) (result *SupportedTypesList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSupportedTypesOptions, "getSupportedTypesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"v2/resources/supported_types"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalSearch.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSupportedTypesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_search", "V2", "GetSupportedTypes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = globalSearch.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSupportedTypesList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetSupportedTypesOptions : The GetSupportedTypes options.
type GetSupportedTypesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSupportedTypesOptions : Instantiate GetSupportedTypesOptions
func (*GlobalSearchV2) NewGetSupportedTypesOptions() *GetSupportedTypesOptions {
	return &GetSupportedTypesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSupportedTypesOptions) SetHeaders(param map[string]string) *GetSupportedTypesOptions {
	options.Headers = param
	return options
}

// ResultItem : A resource returned in a search result.
type ResultItem struct {
	// Resource identifier in CRN format.
	Crn *string `json:"crn,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}


// SetProperty allows the user to set an arbitrary property on an instance of ResultItem
func (o *ResultItem) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of ResultItem
func (o *ResultItem) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of ResultItem
func (o *ResultItem) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of ResultItem
func (o *ResultItem) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Crn != nil {
		m["crn"] = o.Crn
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalResultItem unmarshals an instance of ResultItem from the specified map of raw messages.
func UnmarshalResultItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultItem)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	delete(m, "crn")
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

// ScanResult : The search scan response.
type ScanResult struct {
	// The search cursor to use on all calls after the first one.
	SearchCursor *string `json:"search_cursor" validate:"required"`

	// Value of the limit parameter specified by the user.
	Limit *float64 `json:"limit,omitempty"`

	// The array of results. Each item represents a resource. An empty array signals the end of the result set, there are
	// no more hits to fetch.
	Items []ResultItem `json:"items" validate:"required"`
}


// UnmarshalScanResult unmarshals an instance of ScanResult from the specified map of raw messages.
func UnmarshalScanResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ScanResult)
	err = core.UnmarshalPrimitive(m, "search_cursor", &obj.SearchCursor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalResultItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchOptions : The Search options.
type SearchOptions struct {
	// The Lucene-formatted query string. Default to '*' if not set.
	Query *string `json:"query,omitempty"`

	// The list of the fields returned by the search. Defaults to all. `crn` is always returned.
	Fields []string `json:"fields,omitempty"`

	// An opaque search cursor that is returned on each operation call and that must be set on next call.
	SearchCursor *string `json:"search_cursor,omitempty"`

	// An aplhanumeric string that can be used to trace a request across services. If not specified it will be
	// automatically generated with the prefix "gst-".
	TransactionID *string `json:"transaction-id,omitempty"`

	// The account ID to filter resources.
	AccountID *string `json:"account_id,omitempty"`

	// The maximum number of hits to return. Defaults to 10.
	Limit *int64 `json:"limit,omitempty"`

	// A search timeout, bounding the search request to be executed within the specified time value and bail with the hits
	// accumulated up to that point when expired. Defaults to the system defined timeout.
	Timeout *int64 `json:"timeout,omitempty"`

	// Comma separated properties names used for sorting.
	Sort []string `json:"sort,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSearchOptions : Instantiate SearchOptions
func (*GlobalSearchV2) NewSearchOptions() *SearchOptions {
	return &SearchOptions{}
}

// SetQuery : Allow user to set Query
func (options *SearchOptions) SetQuery(query string) *SearchOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetFields : Allow user to set Fields
func (options *SearchOptions) SetFields(fields []string) *SearchOptions {
	options.Fields = fields
	return options
}

// SetSearchCursor : Allow user to set SearchCursor
func (options *SearchOptions) SetSearchCursor(searchCursor string) *SearchOptions {
	options.SearchCursor = core.StringPtr(searchCursor)
	return options
}

// SetTransactionID : Allow user to set TransactionID
func (options *SearchOptions) SetTransactionID(transactionID string) *SearchOptions {
	options.TransactionID = core.StringPtr(transactionID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *SearchOptions) SetAccountID(accountID string) *SearchOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetLimit : Allow user to set Limit
func (options *SearchOptions) SetLimit(limit int64) *SearchOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetTimeout : Allow user to set Timeout
func (options *SearchOptions) SetTimeout(timeout int64) *SearchOptions {
	options.Timeout = core.Int64Ptr(timeout)
	return options
}

// SetSort : Allow user to set Sort
func (options *SearchOptions) SetSort(sort []string) *SearchOptions {
	options.Sort = sort
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SearchOptions) SetHeaders(param map[string]string) *SearchOptions {
	options.Headers = param
	return options
}

// SupportedTypesList : A list of resource types supported by GhoST.
type SupportedTypesList struct {
	// A list of resource types supported by GhoST.
	SupportedTypes []string `json:"supported_types,omitempty"`
}


// UnmarshalSupportedTypesList unmarshals an instance of SupportedTypesList from the specified map of raw messages.
func UnmarshalSupportedTypesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SupportedTypesList)
	err = core.UnmarshalPrimitive(m, "supported_types", &obj.SupportedTypes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
