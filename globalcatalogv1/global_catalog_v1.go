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

// Package globalcatalogv1 : Operations and models for the GlobalCatalogV1 service
package globalcatalogv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
	"io"
)

// GlobalCatalogV1 : The catalog service manages offerings across geographies as the system of record. The catalog
// supports a RESTful API where users can retrieve information about existing offerings and create, manage, and delete
// their offerings. Start with the base URL and use the endpoints to retrieve metadata about services in the catalog and
// manage service visbility. Depending on the kind of object, the metadata can include information about pricing,
// provisioning, regions, and more. For more information, see the [catalog
// documentation](https://cloud.ibm.com/docs/overview/catalog.html#global-catalog-overview).
//
// Version: 1.0.3
type GlobalCatalogV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://globalcatalog.cloud.ibm.com/api/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "global_catalog"

// GlobalCatalogV1Options : Service options
type GlobalCatalogV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewGlobalCatalogV1UsingExternalConfig : constructs an instance of GlobalCatalogV1 with passed in options and external configuration.
func NewGlobalCatalogV1UsingExternalConfig(options *GlobalCatalogV1Options) (globalCatalog *GlobalCatalogV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	globalCatalog, err = NewGlobalCatalogV1(options)
	if err != nil {
		return
	}

	err = globalCatalog.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = globalCatalog.Service.SetServiceURL(options.URL)
	}
	return
}

// NewGlobalCatalogV1 : constructs an instance of GlobalCatalogV1 with passed in options.
func NewGlobalCatalogV1(options *GlobalCatalogV1Options) (service *GlobalCatalogV1, err error) {
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

	service = &GlobalCatalogV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (globalCatalog *GlobalCatalogV1) SetServiceURL(url string) error {
	return globalCatalog.Service.SetServiceURL(url)
}

// ListCatalogEntries : Returns parent catalog entries
// Includes key information, such as ID, name, kind, CRN, tags, and provider. This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions) (result *SearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogEntriesOptions, "listCatalogEntriesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{""}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCatalogEntriesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "ListCatalogEntries")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listCatalogEntriesOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*listCatalogEntriesOptions.Account))
	}
	if listCatalogEntriesOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*listCatalogEntriesOptions.Include))
	}
	if listCatalogEntriesOptions.Q != nil {
		builder.AddQuery("q", fmt.Sprint(*listCatalogEntriesOptions.Q))
	}
	if listCatalogEntriesOptions.SortBy != nil {
		builder.AddQuery("sort-by", fmt.Sprint(*listCatalogEntriesOptions.SortBy))
	}
	if listCatalogEntriesOptions.Descending != nil {
		builder.AddQuery("descending", fmt.Sprint(*listCatalogEntriesOptions.Descending))
	}
	if listCatalogEntriesOptions.Languages != nil {
		builder.AddQuery("languages", fmt.Sprint(*listCatalogEntriesOptions.Languages))
	}
	if listCatalogEntriesOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*listCatalogEntriesOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalSearchResult(m)
		response.Result = result
	}

	return
}

// CreateCatalogEntry : Create a catalog entry
// The created catalog entry is restricted by default. You must have an administrator or editor role in the scope of the
// provided token. This API will return an ETag that can be used for standard ETag processing, except when depth query
// is used.
func (globalCatalog *GlobalCatalogV1) CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions) (result *CatalogEntry, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCatalogEntryOptions, "createCatalogEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCatalogEntryOptions, "createCatalogEntryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{""}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCatalogEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "CreateCatalogEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createCatalogEntryOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*createCatalogEntryOptions.Account))
	}

	body := make(map[string]interface{})
	if createCatalogEntryOptions.Name != nil {
		body["name"] = createCatalogEntryOptions.Name
	}
	if createCatalogEntryOptions.Kind != nil {
		body["kind"] = createCatalogEntryOptions.Kind
	}
	if createCatalogEntryOptions.OverviewUi != nil {
		body["overview_ui"] = createCatalogEntryOptions.OverviewUi
	}
	if createCatalogEntryOptions.Images != nil {
		body["images"] = createCatalogEntryOptions.Images
	}
	if createCatalogEntryOptions.Disabled != nil {
		body["disabled"] = createCatalogEntryOptions.Disabled
	}
	if createCatalogEntryOptions.Tags != nil {
		body["tags"] = createCatalogEntryOptions.Tags
	}
	if createCatalogEntryOptions.Provider != nil {
		body["provider"] = createCatalogEntryOptions.Provider
	}
	if createCatalogEntryOptions.ID != nil {
		body["id"] = createCatalogEntryOptions.ID
	}
	if createCatalogEntryOptions.ParentID != nil {
		body["parent_id"] = createCatalogEntryOptions.ParentID
	}
	if createCatalogEntryOptions.Group != nil {
		body["group"] = createCatalogEntryOptions.Group
	}
	if createCatalogEntryOptions.Active != nil {
		body["active"] = createCatalogEntryOptions.Active
	}
	if createCatalogEntryOptions.Metadata != nil {
		body["metadata"] = createCatalogEntryOptions.Metadata
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalCatalogEntry(m)
		response.Result = result
	}

	return
}

// GetCatalogEntry : Get a specific catalog object
// This endpoint returns a specific catalog entry using the object's unique identifier, for example
// `/_*service_name*?complete=true`. This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions) (result *CatalogEntry, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCatalogEntryOptions, "getCatalogEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCatalogEntryOptions, "getCatalogEntryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{""}
	pathParameters := []string{*getCatalogEntryOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetCatalogEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCatalogEntryOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getCatalogEntryOptions.Account))
	}
	if getCatalogEntryOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getCatalogEntryOptions.Include))
	}
	if getCatalogEntryOptions.Languages != nil {
		builder.AddQuery("languages", fmt.Sprint(*getCatalogEntryOptions.Languages))
	}
	if getCatalogEntryOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*getCatalogEntryOptions.Complete))
	}
	if getCatalogEntryOptions.Depth != nil {
		builder.AddQuery("depth", fmt.Sprint(*getCatalogEntryOptions.Depth))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalCatalogEntry(m)
		response.Result = result
	}

	return
}

// UpdateCatalogEntry : Update a catalog entry
// Update a catalog entry. The visibility of the catalog entry cannot be modified with this endpoint. You must be an
// administrator or editor in the scope of the provided token. This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions) (result *CatalogEntry, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCatalogEntryOptions, "updateCatalogEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCatalogEntryOptions, "updateCatalogEntryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{""}
	pathParameters := []string{*updateCatalogEntryOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCatalogEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "UpdateCatalogEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if updateCatalogEntryOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*updateCatalogEntryOptions.Account))
	}
	if updateCatalogEntryOptions.Move != nil {
		builder.AddQuery("move", fmt.Sprint(*updateCatalogEntryOptions.Move))
	}

	body := make(map[string]interface{})
	if updateCatalogEntryOptions.Name != nil {
		body["name"] = updateCatalogEntryOptions.Name
	}
	if updateCatalogEntryOptions.Kind != nil {
		body["kind"] = updateCatalogEntryOptions.Kind
	}
	if updateCatalogEntryOptions.OverviewUi != nil {
		body["overview_ui"] = updateCatalogEntryOptions.OverviewUi
	}
	if updateCatalogEntryOptions.Images != nil {
		body["images"] = updateCatalogEntryOptions.Images
	}
	if updateCatalogEntryOptions.Disabled != nil {
		body["disabled"] = updateCatalogEntryOptions.Disabled
	}
	if updateCatalogEntryOptions.Tags != nil {
		body["tags"] = updateCatalogEntryOptions.Tags
	}
	if updateCatalogEntryOptions.Provider != nil {
		body["provider"] = updateCatalogEntryOptions.Provider
	}
	if updateCatalogEntryOptions.ParentID != nil {
		body["parent_id"] = updateCatalogEntryOptions.ParentID
	}
	if updateCatalogEntryOptions.Group != nil {
		body["group"] = updateCatalogEntryOptions.Group
	}
	if updateCatalogEntryOptions.Active != nil {
		body["active"] = updateCatalogEntryOptions.Active
	}
	if updateCatalogEntryOptions.Metadata != nil {
		body["metadata"] = updateCatalogEntryOptions.Metadata
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalCatalogEntry(m)
		response.Result = result
	}

	return
}

// DeleteCatalogEntry : Delete a catalog entry
// Delete a catalog entry. This will archive the catalog entry for a minimum of two weeks. While archived, it can be
// restored using the PUT restore API. After two weeks, it will be deleted and cannot be restored. You must have
// administrator role in the scope of the provided token to modify it. This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) DeleteCatalogEntry(deleteCatalogEntryOptions *DeleteCatalogEntryOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCatalogEntryOptions, "deleteCatalogEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCatalogEntryOptions, "deleteCatalogEntryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{""}
	pathParameters := []string{*deleteCatalogEntryOptions.ID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCatalogEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "DeleteCatalogEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteCatalogEntryOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*deleteCatalogEntryOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// GetChildObjects : Get child catalog entries of a specific kind
// Fetch child catalog entries for a catalog entry with a specific id. This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions) (result *[]SearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getChildObjectsOptions, "getChildObjectsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getChildObjectsOptions, "getChildObjectsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", ""}
	pathParameters := []string{*getChildObjectsOptions.ID, *getChildObjectsOptions.Kind}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getChildObjectsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetChildObjects")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getChildObjectsOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getChildObjectsOptions.Account))
	}
	if getChildObjectsOptions.Include != nil {
		builder.AddQuery("include", fmt.Sprint(*getChildObjectsOptions.Include))
	}
	if getChildObjectsOptions.Q != nil {
		builder.AddQuery("q", fmt.Sprint(*getChildObjectsOptions.Q))
	}
	if getChildObjectsOptions.SortBy != nil {
		builder.AddQuery("sort-by", fmt.Sprint(*getChildObjectsOptions.SortBy))
	}
	if getChildObjectsOptions.Descending != nil {
		builder.AddQuery("descending", fmt.Sprint(*getChildObjectsOptions.Descending))
	}
	if getChildObjectsOptions.Languages != nil {
		builder.AddQuery("languages", fmt.Sprint(*getChildObjectsOptions.Languages))
	}
	if getChildObjectsOptions.Complete != nil {
		builder.AddQuery("complete", fmt.Sprint(*getChildObjectsOptions.Complete))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make([]map[string]interface{}, 1))
	if err == nil {
		s, ok := response.Result.([]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		slice, e := UnmarshalSearchResultSlice(s)
		result = &slice
		err = e
		response.Result = result
	}

	return
}

// RestoreCatalogEntry : Restore archived catalog entry
// Restore an archived catalog entry. You must have an administrator role in the scope of the provided token.
func (globalCatalog *GlobalCatalogV1) RestoreCatalogEntry(restoreCatalogEntryOptions *RestoreCatalogEntryOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(restoreCatalogEntryOptions, "restoreCatalogEntryOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(restoreCatalogEntryOptions, "restoreCatalogEntryOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "restore"}
	pathParameters := []string{*restoreCatalogEntryOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range restoreCatalogEntryOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "RestoreCatalogEntry")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if restoreCatalogEntryOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*restoreCatalogEntryOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// GetVisibility : Get the visibility constraints for an object
// This endpoint returns the visibility rules for this object. Overall visibility is determined by the parent objects
// and any further restrictions on this object. You must have an administrator role in the scope of the provided token.
// This endpoint is ETag enabled.
func (globalCatalog *GlobalCatalogV1) GetVisibility(getVisibilityOptions *GetVisibilityOptions) (result *Visibility, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVisibilityOptions, "getVisibilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVisibilityOptions, "getVisibilityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "visibility"}
	pathParameters := []string{*getVisibilityOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVisibilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetVisibility")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getVisibilityOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getVisibilityOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalVisibility(m)
		response.Result = result
	}

	return
}

// UpdateVisibility : Update visibility
// Update an Object's Visibility. You must have an administrator role in the scope of the provided token. This endpoint
// is ETag enabled.
func (globalCatalog *GlobalCatalogV1) UpdateVisibility(updateVisibilityOptions *UpdateVisibilityOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateVisibilityOptions, "updateVisibilityOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateVisibilityOptions, "updateVisibilityOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "visibility"}
	pathParameters := []string{*updateVisibilityOptions.ID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateVisibilityOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "UpdateVisibility")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	if updateVisibilityOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*updateVisibilityOptions.Account))
	}

	body := make(map[string]interface{})
	if updateVisibilityOptions.Include != nil {
		body["include"] = updateVisibilityOptions.Include
	}
	if updateVisibilityOptions.Exclude != nil {
		body["exclude"] = updateVisibilityOptions.Exclude
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// GetPricing : Get the pricing for an object
// This endpoint returns the pricing for an object. Static pricing is defined in the catalog. Dynamic pricing is stored
// in Bluemix Pricing Catalog.
func (globalCatalog *GlobalCatalogV1) GetPricing(getPricingOptions *GetPricingOptions) (result *PricingGet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPricingOptions, "getPricingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPricingOptions, "getPricingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "pricing"}
	pathParameters := []string{*getPricingOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPricingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetPricing")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getPricingOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getPricingOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalPricingGet(m)
		response.Result = result
	}

	return
}

// GetAuditLogs : Get the audit logs for an object
// This endpoint returns the audit logs for an object. Only administrators and editors can get logs.
func (globalCatalog *GlobalCatalogV1) GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions) (result *SearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getAuditLogsOptions, "getAuditLogsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getAuditLogsOptions, "getAuditLogsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "logs"}
	pathParameters := []string{*getAuditLogsOptions.ID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAuditLogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetAuditLogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getAuditLogsOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getAuditLogsOptions.Account))
	}
	if getAuditLogsOptions.Ascending != nil {
		builder.AddQuery("ascending", fmt.Sprint(*getAuditLogsOptions.Ascending))
	}
	if getAuditLogsOptions.Startat != nil {
		builder.AddQuery("startat", fmt.Sprint(*getAuditLogsOptions.Startat))
	}
	if getAuditLogsOptions.Offset != nil {
		builder.AddQuery("_offset", fmt.Sprint(*getAuditLogsOptions.Offset))
	}
	if getAuditLogsOptions.Limit != nil {
		builder.AddQuery("_limit", fmt.Sprint(*getAuditLogsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalSearchResult(m)
		response.Result = result
	}

	return
}

// ListArtifacts : Get artifacts
// This endpoint returns a list of artifacts for an object.
func (globalCatalog *GlobalCatalogV1) ListArtifacts(listArtifactsOptions *ListArtifactsOptions) (result *Artifacts, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listArtifactsOptions, "listArtifactsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listArtifactsOptions, "listArtifactsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "artifacts"}
	pathParameters := []string{*listArtifactsOptions.ObjectID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listArtifactsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "ListArtifacts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listArtifactsOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*listArtifactsOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, make(map[string]interface{}))
	if err == nil {
		m, ok := response.Result.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("an error occurred while processing the operation response")
			return
		}
		result, err = UnmarshalArtifacts(m)
		response.Result = result
	}

	return
}

// GetArtifact : Get artifact
// This endpoint returns the binary of an artifact.
func (globalCatalog *GlobalCatalogV1) GetArtifact(getArtifactOptions *GetArtifactOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getArtifactOptions, "getArtifactOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getArtifactOptions, "getArtifactOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "artifacts"}
	pathParameters := []string{*getArtifactOptions.ObjectID, *getArtifactOptions.ArtifactID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getArtifactOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "GetArtifact")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if getArtifactOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*getArtifactOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// UploadArtifact : Upload artifact
// This endpoint uploads the binary for an artifact. Only administrators and editors can upload artifacts.
func (globalCatalog *GlobalCatalogV1) UploadArtifact(uploadArtifactOptions *UploadArtifactOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(uploadArtifactOptions, "uploadArtifactOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(uploadArtifactOptions, "uploadArtifactOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "artifacts"}
	pathParameters := []string{*uploadArtifactOptions.ObjectID, *uploadArtifactOptions.ArtifactID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range uploadArtifactOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "UploadArtifact")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if uploadArtifactOptions.ContentType != nil {
		builder.AddHeader("Content-Type", fmt.Sprint(*uploadArtifactOptions.ContentType))
	}

	if uploadArtifactOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*uploadArtifactOptions.Account))
	}

	_, err = builder.SetBodyContent(core.StringNilMapper(uploadArtifactOptions.ContentType), nil, nil, uploadArtifactOptions.Artifact)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// DeleteArtifact : Delete artifact
// This endpoint deletes an artifact. Only administrators and editors can delete artifacts.
func (globalCatalog *GlobalCatalogV1) DeleteArtifact(deleteArtifactOptions *DeleteArtifactOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteArtifactOptions, "deleteArtifactOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteArtifactOptions, "deleteArtifactOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"", "artifacts"}
	pathParameters := []string{*deleteArtifactOptions.ObjectID, *deleteArtifactOptions.ArtifactID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(globalCatalog.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteArtifactOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("global_catalog", "V1", "DeleteArtifact")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteArtifactOptions.Account != nil {
		builder.AddQuery("account", fmt.Sprint(*deleteArtifactOptions.Account))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = globalCatalog.Service.Request(request, nil)

	return
}

// Amount : Country-specific pricing information.
type Amount struct {
	// Country.
	Counrty *string `json:"counrty,omitempty"`

	// Currency.
	Currency *string `json:"currency,omitempty"`

	// See Price for nested fields.
	Prices []Price `json:"prices,omitempty"`
}


// UnmarshalAmount constructs an instance of Amount from the specified map.
func UnmarshalAmount(m map[string]interface{}) (result *Amount, err error) {
	obj := new(Amount)
	obj.Counrty, err = core.UnmarshalString(m, "counrty")
	if err != nil {
		return
	}
	obj.Currency, err = core.UnmarshalString(m, "currency")
	if err != nil {
		return
	}
	obj.Prices, err = UnmarshalPriceSliceAsProperty(m, "prices")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalAmountSlice unmarshals a slice of Amount instances from the specified list of maps.
func UnmarshalAmountSlice(s []interface{}) (slice []Amount, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Amount'")
			return
		}
		obj, e := UnmarshalAmount(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalAmountAsProperty unmarshals an instance of Amount that is stored as a property
// within the specified map.
func UnmarshalAmountAsProperty(m map[string]interface{}, propertyName string) (result *Amount, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Amount'", propertyName)
			return
		}
		result, err = UnmarshalAmount(objMap)
	}
	return
}

// UnmarshalAmountSliceAsProperty unmarshals a slice of Amount instances that are stored as a property
// within the specified map.
func UnmarshalAmountSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Amount, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Amount'", propertyName)
			return
		}
		slice, err = UnmarshalAmountSlice(vSlice)
	}
	return
}

// Artifact : Artifact Details.
type Artifact struct {
	// The name of the artifact.
	Name *string `json:"name,omitempty"`

	// The timestamp of the last update to the artifact.
	Updated *string `json:"updated,omitempty"`

	// The url for the artifact.
	URL *string `json:"url,omitempty"`

	// The etag of the artifact.
	Etag *string `json:"etag,omitempty"`

	// The content length of the artifact.
	Size *int64 `json:"size,omitempty"`
}


// UnmarshalArtifact constructs an instance of Artifact from the specified map.
func UnmarshalArtifact(m map[string]interface{}) (result *Artifact, err error) {
	obj := new(Artifact)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Updated, err = core.UnmarshalString(m, "updated")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "url")
	if err != nil {
		return
	}
	obj.Etag, err = core.UnmarshalString(m, "etag")
	if err != nil {
		return
	}
	obj.Size, err = core.UnmarshalInt64(m, "size")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalArtifactSlice unmarshals a slice of Artifact instances from the specified list of maps.
func UnmarshalArtifactSlice(s []interface{}) (slice []Artifact, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Artifact'")
			return
		}
		obj, e := UnmarshalArtifact(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalArtifactAsProperty unmarshals an instance of Artifact that is stored as a property
// within the specified map.
func UnmarshalArtifactAsProperty(m map[string]interface{}, propertyName string) (result *Artifact, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Artifact'", propertyName)
			return
		}
		result, err = UnmarshalArtifact(objMap)
	}
	return
}

// UnmarshalArtifactSliceAsProperty unmarshals a slice of Artifact instances that are stored as a property
// within the specified map.
func UnmarshalArtifactSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Artifact, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Artifact'", propertyName)
			return
		}
		slice, err = UnmarshalArtifactSlice(vSlice)
	}
	return
}

// Artifacts : Artifacts List.
type Artifacts struct {
	// The total number of artifacts.
	Count *int64 `json:"count,omitempty"`

	// The list of artifacts.
	Resources []Artifact `json:"resources,omitempty"`
}


// UnmarshalArtifacts constructs an instance of Artifacts from the specified map.
func UnmarshalArtifacts(m map[string]interface{}) (result *Artifacts, err error) {
	obj := new(Artifacts)
	obj.Count, err = core.UnmarshalInt64(m, "count")
	if err != nil {
		return
	}
	obj.Resources, err = UnmarshalArtifactSliceAsProperty(m, "resources")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalArtifactsSlice unmarshals a slice of Artifacts instances from the specified list of maps.
func UnmarshalArtifactsSlice(s []interface{}) (slice []Artifacts, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Artifacts'")
			return
		}
		obj, e := UnmarshalArtifacts(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalArtifactsAsProperty unmarshals an instance of Artifacts that is stored as a property
// within the specified map.
func UnmarshalArtifactsAsProperty(m map[string]interface{}, propertyName string) (result *Artifacts, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Artifacts'", propertyName)
			return
		}
		result, err = UnmarshalArtifacts(objMap)
	}
	return
}

// UnmarshalArtifactsSliceAsProperty unmarshals a slice of Artifacts instances that are stored as a property
// within the specified map.
func UnmarshalArtifactsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Artifacts, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Artifacts'", propertyName)
			return
		}
		slice, err = UnmarshalArtifactsSlice(vSlice)
	}
	return
}

// Bullets : Information related to list delimiters.
type Bullets struct {
	// The bullet title.
	Title *string `json:"title,omitempty"`

	// The bullet description.
	Description *string `json:"description,omitempty"`

	// The icon to use for rendering the bullet.
	Icon *string `json:"icon,omitempty"`

	// The bullet quantity.
	Quantity *string `json:"quantity,omitempty"`
}


// UnmarshalBullets constructs an instance of Bullets from the specified map.
func UnmarshalBullets(m map[string]interface{}) (result *Bullets, err error) {
	obj := new(Bullets)
	obj.Title, err = core.UnmarshalString(m, "title")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	obj.Icon, err = core.UnmarshalString(m, "icon")
	if err != nil {
		return
	}
	obj.Quantity, err = core.UnmarshalString(m, "quantity")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalBulletsSlice unmarshals a slice of Bullets instances from the specified list of maps.
func UnmarshalBulletsSlice(s []interface{}) (slice []Bullets, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Bullets'")
			return
		}
		obj, e := UnmarshalBullets(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalBulletsAsProperty unmarshals an instance of Bullets that is stored as a property
// within the specified map.
func UnmarshalBulletsAsProperty(m map[string]interface{}, propertyName string) (result *Bullets, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Bullets'", propertyName)
			return
		}
		result, err = UnmarshalBullets(objMap)
	}
	return
}

// UnmarshalBulletsSliceAsProperty unmarshals a slice of Bullets instances that are stored as a property
// within the specified map.
func UnmarshalBulletsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Bullets, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Bullets'", propertyName)
			return
		}
		slice, err = UnmarshalBulletsSlice(vSlice)
	}
	return
}

// Callbacks : Callback-related information associated with a catalog entry.
type Callbacks struct {
	// The URL of the deployment broker.
	BrokerUtl *string `json:"broker_utl,omitempty"`

	// The URL of the deployment broker SC proxy.
	BrokerProxyURL *string `json:"broker_proxy_url,omitempty"`

	// The URL of dashboard callback.
	DashboardURL *string `json:"dashboard_url,omitempty"`

	// The URL of dashboard data.
	DashboardDataURL *string `json:"dashboard_data_url,omitempty"`

	// The URL of the dashboard detail tab.
	DashboardDetailTabURL *string `json:"dashboard_detail_tab_url,omitempty"`

	// The URL of the dashboard detail tab extension.
	DashboardDetailTabExtURL *string `json:"dashboard_detail_tab_ext_url,omitempty"`

	// Service monitor API URL.
	ServiceMonitorApi *string `json:"service_monitor_api,omitempty"`

	// Service monitor app URL.
	ServiceMonitorApp *string `json:"service_monitor_app,omitempty"`

	// Service URL in staging.
	ServiceStagingURL *string `json:"service_staging_url,omitempty"`

	// Service URL in production.
	ServiceProductionURL *string `json:"service_production_url,omitempty"`
}


// UnmarshalCallbacks constructs an instance of Callbacks from the specified map.
func UnmarshalCallbacks(m map[string]interface{}) (result *Callbacks, err error) {
	obj := new(Callbacks)
	obj.BrokerUtl, err = core.UnmarshalString(m, "broker_utl")
	if err != nil {
		return
	}
	obj.BrokerProxyURL, err = core.UnmarshalString(m, "broker_proxy_url")
	if err != nil {
		return
	}
	obj.DashboardURL, err = core.UnmarshalString(m, "dashboard_url")
	if err != nil {
		return
	}
	obj.DashboardDataURL, err = core.UnmarshalString(m, "dashboard_data_url")
	if err != nil {
		return
	}
	obj.DashboardDetailTabURL, err = core.UnmarshalString(m, "dashboard_detail_tab_url")
	if err != nil {
		return
	}
	obj.DashboardDetailTabExtURL, err = core.UnmarshalString(m, "dashboard_detail_tab_ext_url")
	if err != nil {
		return
	}
	obj.ServiceMonitorApi, err = core.UnmarshalString(m, "service_monitor_api")
	if err != nil {
		return
	}
	obj.ServiceMonitorApp, err = core.UnmarshalString(m, "service_monitor_app")
	if err != nil {
		return
	}
	obj.ServiceStagingURL, err = core.UnmarshalString(m, "service_staging_url")
	if err != nil {
		return
	}
	obj.ServiceProductionURL, err = core.UnmarshalString(m, "service_production_url")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCallbacksSlice unmarshals a slice of Callbacks instances from the specified list of maps.
func UnmarshalCallbacksSlice(s []interface{}) (slice []Callbacks, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Callbacks'")
			return
		}
		obj, e := UnmarshalCallbacks(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCallbacksAsProperty unmarshals an instance of Callbacks that is stored as a property
// within the specified map.
func UnmarshalCallbacksAsProperty(m map[string]interface{}, propertyName string) (result *Callbacks, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Callbacks'", propertyName)
			return
		}
		result, err = UnmarshalCallbacks(objMap)
	}
	return
}

// UnmarshalCallbacksSliceAsProperty unmarshals a slice of Callbacks instances that are stored as a property
// within the specified map.
func UnmarshalCallbacksSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Callbacks, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Callbacks'", propertyName)
			return
		}
		slice, err = UnmarshalCallbacksSlice(vSlice)
	}
	return
}

// CatalogEntry : An entry in the global catalog.
type CatalogEntry struct {
	// Programmatic name for this catalog entry, which must be formatted like a CRN segment. See the display name in
	// OverviewUI for a user-readable name.
	Name *string `json:"name" validate:"required"`

	// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
	// object.
	Kind *string `json:"kind" validate:"required"`

	// Overview is nested in the top level. The key value pair is `[_language_]overview_ui`.
	OverviewUi *OverviewUI `json:"overview_ui" validate:"required"`

	// Image annotation for this catalog entry. The image is a URL.
	Images *Image `json:"images" validate:"required"`

	// The ID of the parent catalog entry if it exists.
	ParentID *string `json:"parent_id,omitempty"`

	// Boolean value that determines the global visibility for the catalog entry, and its children. If it is not enabled,
	// all plans are disabled.
	Disabled *bool `json:"disabled" validate:"required"`

	// A list of tags. For example, IBM, 3rd Party, Beta, GA, and Single Tenant.
	Tags []string `json:"tags" validate:"required"`

	// Boolean value that determines whether the catalog entry is a group.
	Group *bool `json:"group,omitempty"`

	// Information related to the provider associated with a catalog entry.
	Provider *Provider `json:"provider" validate:"required"`

	// Boolean value that describes whether the service is active.
	Active *bool `json:"active,omitempty"`

	// Model used to describe metadata object that can be set.
	Metadata *ObjectMetadataSet `json:"metadata,omitempty"`

	// Catalog entry's unique ID. It's the same across all catalog instances.
	ID *string `json:"id,omitempty"`

	CatalogCrn interface{} `json:"catalog_crn,omitempty"`

	URL interface{} `json:"url,omitempty"`

	ChildrenURL interface{} `json:"children_url,omitempty"`

	GeoTags interface{} `json:"geo_tags,omitempty"`

	PricingTags interface{} `json:"pricing_tags,omitempty"`

	Created interface{} `json:"created,omitempty"`

	Updated interface{} `json:"updated,omitempty"`
}

// Constants associated with the CatalogEntry.Kind property.
// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
// object.
const (
	CatalogEntry_Kind_Dashboard = "dashboard"
	CatalogEntry_Kind_Service = "service"
	CatalogEntry_Kind_Template = "template"
)


// UnmarshalCatalogEntry constructs an instance of CatalogEntry from the specified map.
func UnmarshalCatalogEntry(m map[string]interface{}) (result *CatalogEntry, err error) {
	obj := new(CatalogEntry)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Kind, err = core.UnmarshalString(m, "kind")
	if err != nil {
		return
	}
	obj.OverviewUi, err = UnmarshalOverviewUIAsProperty(m, "overview_ui")
	if err != nil {
		return
	}
	obj.Images, err = UnmarshalImageAsProperty(m, "images")
	if err != nil {
		return
	}
	obj.ParentID, err = core.UnmarshalString(m, "parent_id")
	if err != nil {
		return
	}
	obj.Disabled, err = core.UnmarshalBool(m, "disabled")
	if err != nil {
		return
	}
	obj.Tags, err = core.UnmarshalStringSlice(m, "tags")
	if err != nil {
		return
	}
	obj.Group, err = core.UnmarshalBool(m, "group")
	if err != nil {
		return
	}
	obj.Provider, err = UnmarshalProviderAsProperty(m, "provider")
	if err != nil {
		return
	}
	obj.Active, err = core.UnmarshalBool(m, "active")
	if err != nil {
		return
	}
	obj.Metadata, err = UnmarshalObjectMetadataSetAsProperty(m, "metadata")
	if err != nil {
		return
	}
	obj.ID, err = core.UnmarshalString(m, "id")
	if err != nil {
		return
	}
	obj.CatalogCrn, err = core.UnmarshalAny(m, "catalog_crn")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalAny(m, "url")
	if err != nil {
		return
	}
	obj.ChildrenURL, err = core.UnmarshalAny(m, "children_url")
	if err != nil {
		return
	}
	obj.GeoTags, err = core.UnmarshalAny(m, "geo_tags")
	if err != nil {
		return
	}
	obj.PricingTags, err = core.UnmarshalAny(m, "pricing_tags")
	if err != nil {
		return
	}
	obj.Created, err = core.UnmarshalAny(m, "created")
	if err != nil {
		return
	}
	obj.Updated, err = core.UnmarshalAny(m, "updated")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalCatalogEntrySlice unmarshals a slice of CatalogEntry instances from the specified list of maps.
func UnmarshalCatalogEntrySlice(s []interface{}) (slice []CatalogEntry, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'CatalogEntry'")
			return
		}
		obj, e := UnmarshalCatalogEntry(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalCatalogEntryAsProperty unmarshals an instance of CatalogEntry that is stored as a property
// within the specified map.
func UnmarshalCatalogEntryAsProperty(m map[string]interface{}, propertyName string) (result *CatalogEntry, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'CatalogEntry'", propertyName)
			return
		}
		result, err = UnmarshalCatalogEntry(objMap)
	}
	return
}

// UnmarshalCatalogEntrySliceAsProperty unmarshals a slice of CatalogEntry instances that are stored as a property
// within the specified map.
func UnmarshalCatalogEntrySliceAsProperty(m map[string]interface{}, propertyName string) (slice []CatalogEntry, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'CatalogEntry'", propertyName)
			return
		}
		slice, err = UnmarshalCatalogEntrySlice(vSlice)
	}
	return
}

// CreateCatalogEntryOptions : The CreateCatalogEntry options.
type CreateCatalogEntryOptions struct {
	// Programmatic name for this catalog entry, which must be formatted like a CRN segment. See the display name in
	// OverviewUI for a user-readable name.
	Name *string `json:"name" validate:"required"`

	// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
	// object.
	Kind *string `json:"kind" validate:"required"`

	// Overview is nested in the top level. The key value pair is `[_language_]overview_ui`.
	OverviewUi *OverviewUI `json:"overview_ui" validate:"required"`

	// Image annotation for this catalog entry. The image is a URL.
	Images *Image `json:"images" validate:"required"`

	// Boolean value that determines the global visibility for the catalog entry, and its children. If it is not enabled,
	// all plans are disabled.
	Disabled *bool `json:"disabled" validate:"required"`

	// A list of tags. For example, IBM, 3rd Party, Beta, GA, and Single Tenant.
	Tags []string `json:"tags" validate:"required"`

	// Information related to the provider associated with a catalog entry.
	Provider *Provider `json:"provider" validate:"required"`

	// Catalog entry's unique ID. It's the same across all catalog instances.
	ID *string `json:"id" validate:"required"`

	// The ID of the parent catalog entry if it exists.
	ParentID *string `json:"parent_id,omitempty"`

	// Boolean value that determines whether the catalog entry is a group.
	Group *bool `json:"group,omitempty"`

	// Boolean value that describes whether the service is active.
	Active *bool `json:"active,omitempty"`

	// Model used to describe metadata object that can be set.
	Metadata *ObjectMetadataSet `json:"metadata,omitempty"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCatalogEntryOptions.Kind property.
// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
// object.
const (
	CreateCatalogEntryOptions_Kind_Dashboard = "dashboard"
	CreateCatalogEntryOptions_Kind_Service = "service"
	CreateCatalogEntryOptions_Kind_Template = "template"
)

// NewCreateCatalogEntryOptions : Instantiate CreateCatalogEntryOptions
func (*GlobalCatalogV1) NewCreateCatalogEntryOptions(name string, kind string, overviewUi *OverviewUI, images *Image, disabled bool, tags []string, provider *Provider, id string) *CreateCatalogEntryOptions {
	return &CreateCatalogEntryOptions{
		Name: core.StringPtr(name),
		Kind: core.StringPtr(kind),
		OverviewUi: overviewUi,
		Images: images,
		Disabled: core.BoolPtr(disabled),
		Tags: tags,
		Provider: provider,
		ID: core.StringPtr(id),
	}
}

// SetName : Allow user to set Name
func (options *CreateCatalogEntryOptions) SetName(name string) *CreateCatalogEntryOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetKind : Allow user to set Kind
func (options *CreateCatalogEntryOptions) SetKind(kind string) *CreateCatalogEntryOptions {
	options.Kind = core.StringPtr(kind)
	return options
}

// SetOverviewUi : Allow user to set OverviewUi
func (options *CreateCatalogEntryOptions) SetOverviewUi(overviewUi *OverviewUI) *CreateCatalogEntryOptions {
	options.OverviewUi = overviewUi
	return options
}

// SetImages : Allow user to set Images
func (options *CreateCatalogEntryOptions) SetImages(images *Image) *CreateCatalogEntryOptions {
	options.Images = images
	return options
}

// SetDisabled : Allow user to set Disabled
func (options *CreateCatalogEntryOptions) SetDisabled(disabled bool) *CreateCatalogEntryOptions {
	options.Disabled = core.BoolPtr(disabled)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateCatalogEntryOptions) SetTags(tags []string) *CreateCatalogEntryOptions {
	options.Tags = tags
	return options
}

// SetProvider : Allow user to set Provider
func (options *CreateCatalogEntryOptions) SetProvider(provider *Provider) *CreateCatalogEntryOptions {
	options.Provider = provider
	return options
}

// SetID : Allow user to set ID
func (options *CreateCatalogEntryOptions) SetID(id string) *CreateCatalogEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetParentID : Allow user to set ParentID
func (options *CreateCatalogEntryOptions) SetParentID(parentID string) *CreateCatalogEntryOptions {
	options.ParentID = core.StringPtr(parentID)
	return options
}

// SetGroup : Allow user to set Group
func (options *CreateCatalogEntryOptions) SetGroup(group bool) *CreateCatalogEntryOptions {
	options.Group = core.BoolPtr(group)
	return options
}

// SetActive : Allow user to set Active
func (options *CreateCatalogEntryOptions) SetActive(active bool) *CreateCatalogEntryOptions {
	options.Active = core.BoolPtr(active)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *CreateCatalogEntryOptions) SetMetadata(metadata *ObjectMetadataSet) *CreateCatalogEntryOptions {
	options.Metadata = metadata
	return options
}

// SetAccount : Allow user to set Account
func (options *CreateCatalogEntryOptions) SetAccount(account string) *CreateCatalogEntryOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCatalogEntryOptions) SetHeaders(param map[string]string) *CreateCatalogEntryOptions {
	options.Headers = param
	return options
}

// DeleteArtifactOptions : The DeleteArtifact options.
type DeleteArtifactOptions struct {
	// The object's unique ID.
	ObjectID *string `json:"object_id" validate:"required"`

	// The artifact's ID.
	ArtifactID *string `json:"artifact_id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteArtifactOptions : Instantiate DeleteArtifactOptions
func (*GlobalCatalogV1) NewDeleteArtifactOptions(objectID string, artifactID string) *DeleteArtifactOptions {
	return &DeleteArtifactOptions{
		ObjectID: core.StringPtr(objectID),
		ArtifactID: core.StringPtr(artifactID),
	}
}

// SetObjectID : Allow user to set ObjectID
func (options *DeleteArtifactOptions) SetObjectID(objectID string) *DeleteArtifactOptions {
	options.ObjectID = core.StringPtr(objectID)
	return options
}

// SetArtifactID : Allow user to set ArtifactID
func (options *DeleteArtifactOptions) SetArtifactID(artifactID string) *DeleteArtifactOptions {
	options.ArtifactID = core.StringPtr(artifactID)
	return options
}

// SetAccount : Allow user to set Account
func (options *DeleteArtifactOptions) SetAccount(account string) *DeleteArtifactOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteArtifactOptions) SetHeaders(param map[string]string) *DeleteArtifactOptions {
	options.Headers = param
	return options
}

// DeleteCatalogEntryOptions : The DeleteCatalogEntry options.
type DeleteCatalogEntryOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCatalogEntryOptions : Instantiate DeleteCatalogEntryOptions
func (*GlobalCatalogV1) NewDeleteCatalogEntryOptions(id string) *DeleteCatalogEntryOptions {
	return &DeleteCatalogEntryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *DeleteCatalogEntryOptions) SetID(id string) *DeleteCatalogEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *DeleteCatalogEntryOptions) SetAccount(account string) *DeleteCatalogEntryOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCatalogEntryOptions) SetHeaders(param map[string]string) *DeleteCatalogEntryOptions {
	options.Headers = param
	return options
}

// DeploymentBase : Deployment-related metadata.
type DeploymentBase struct {
	// Describes the region where the service is located.
	Location *string `json:"location,omitempty"`

	// A CRN that describes the deployment. crn:v1:[cname]:[ctype]:[location]:[scope]::[resource-type]:[resource].
	TargetCrn *string `json:"target_crn,omitempty"`

	// The broker associated with a catalog entry.
	Broker *DeploymentBaseBroker `json:"broker,omitempty"`

	// This deployment not only supports RC but is ready to migrate and support the RC broker for a location.
	SupportsRcMigration *bool `json:"supports_rc_migration,omitempty"`

	// network to use during deployment.
	TargetNetwork *string `json:"target_network,omitempty"`
}


// UnmarshalDeploymentBase constructs an instance of DeploymentBase from the specified map.
func UnmarshalDeploymentBase(m map[string]interface{}) (result *DeploymentBase, err error) {
	obj := new(DeploymentBase)
	obj.Location, err = core.UnmarshalString(m, "location")
	if err != nil {
		return
	}
	obj.TargetCrn, err = core.UnmarshalString(m, "target_crn")
	if err != nil {
		return
	}
	obj.Broker, err = UnmarshalDeploymentBaseBrokerAsProperty(m, "broker")
	if err != nil {
		return
	}
	obj.SupportsRcMigration, err = core.UnmarshalBool(m, "supports_rc_migration")
	if err != nil {
		return
	}
	obj.TargetNetwork, err = core.UnmarshalString(m, "target_network")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeploymentBaseSlice unmarshals a slice of DeploymentBase instances from the specified list of maps.
func UnmarshalDeploymentBaseSlice(s []interface{}) (slice []DeploymentBase, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeploymentBase'")
			return
		}
		obj, e := UnmarshalDeploymentBase(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeploymentBaseAsProperty unmarshals an instance of DeploymentBase that is stored as a property
// within the specified map.
func UnmarshalDeploymentBaseAsProperty(m map[string]interface{}, propertyName string) (result *DeploymentBase, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeploymentBase'", propertyName)
			return
		}
		result, err = UnmarshalDeploymentBase(objMap)
	}
	return
}

// UnmarshalDeploymentBaseSliceAsProperty unmarshals a slice of DeploymentBase instances that are stored as a property
// within the specified map.
func UnmarshalDeploymentBaseSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeploymentBase, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeploymentBase'", propertyName)
			return
		}
		slice, err = UnmarshalDeploymentBaseSlice(vSlice)
	}
	return
}

// DeploymentBaseBroker : The broker associated with a catalog entry.
type DeploymentBaseBroker struct {
	// Broker name.
	Name *string `json:"name,omitempty"`

	// Broker guid.
	Guid *string `json:"guid,omitempty"`
}


// UnmarshalDeploymentBaseBroker constructs an instance of DeploymentBaseBroker from the specified map.
func UnmarshalDeploymentBaseBroker(m map[string]interface{}) (result *DeploymentBaseBroker, err error) {
	obj := new(DeploymentBaseBroker)
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Guid, err = core.UnmarshalString(m, "guid")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalDeploymentBaseBrokerSlice unmarshals a slice of DeploymentBaseBroker instances from the specified list of maps.
func UnmarshalDeploymentBaseBrokerSlice(s []interface{}) (slice []DeploymentBaseBroker, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'DeploymentBaseBroker'")
			return
		}
		obj, e := UnmarshalDeploymentBaseBroker(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalDeploymentBaseBrokerAsProperty unmarshals an instance of DeploymentBaseBroker that is stored as a property
// within the specified map.
func UnmarshalDeploymentBaseBrokerAsProperty(m map[string]interface{}, propertyName string) (result *DeploymentBaseBroker, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'DeploymentBaseBroker'", propertyName)
			return
		}
		result, err = UnmarshalDeploymentBaseBroker(objMap)
	}
	return
}

// UnmarshalDeploymentBaseBrokerSliceAsProperty unmarshals a slice of DeploymentBaseBroker instances that are stored as a property
// within the specified map.
func UnmarshalDeploymentBaseBrokerSliceAsProperty(m map[string]interface{}, propertyName string) (slice []DeploymentBaseBroker, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'DeploymentBaseBroker'", propertyName)
			return
		}
		slice, err = UnmarshalDeploymentBaseBrokerSlice(vSlice)
	}
	return
}

// GetArtifactOptions : The GetArtifact options.
type GetArtifactOptions struct {
	// The object's unique ID.
	ObjectID *string `json:"object_id" validate:"required"`

	// The artifact's ID.
	ArtifactID *string `json:"artifact_id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetArtifactOptions : Instantiate GetArtifactOptions
func (*GlobalCatalogV1) NewGetArtifactOptions(objectID string, artifactID string) *GetArtifactOptions {
	return &GetArtifactOptions{
		ObjectID: core.StringPtr(objectID),
		ArtifactID: core.StringPtr(artifactID),
	}
}

// SetObjectID : Allow user to set ObjectID
func (options *GetArtifactOptions) SetObjectID(objectID string) *GetArtifactOptions {
	options.ObjectID = core.StringPtr(objectID)
	return options
}

// SetArtifactID : Allow user to set ArtifactID
func (options *GetArtifactOptions) SetArtifactID(artifactID string) *GetArtifactOptions {
	options.ArtifactID = core.StringPtr(artifactID)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetArtifactOptions) SetAccount(account string) *GetArtifactOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetArtifactOptions) SetHeaders(param map[string]string) *GetArtifactOptions {
	options.Headers = param
	return options
}

// GetAuditLogsOptions : The GetAuditLogs options.
type GetAuditLogsOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Sets the sort order. False is descending.
	Ascending *string `json:"ascending,omitempty"`

	// Starting time for the logs. If it's descending then the entries will be equal or earlier. The default is latest. For
	// ascending it will entries equal or later. The default is earliest. It can be either a number or a string. If a
	// number then it is in the format of Unix timestamps. If it is a string then it is a date in the format
	// YYYY-MM-DDTHH:MM:SSZ  and the time is UTC. The T and the Z are required. For example: 2017-12-24T12:00:00Z for Noon
	// UTC on Dec 24, 2017.
	Startat *string `json:"startat,omitempty"`

	// Count of number of log entries to skip before returning logs. The default is zero.
	Offset *int64 `json:"_offset,omitempty"`

	// Count of number of entries to return. The default is fifty. The maximum value is two hundred.
	Limit *int64 `json:"_limit,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAuditLogsOptions : Instantiate GetAuditLogsOptions
func (*GlobalCatalogV1) NewGetAuditLogsOptions(id string) *GetAuditLogsOptions {
	return &GetAuditLogsOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetAuditLogsOptions) SetID(id string) *GetAuditLogsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetAuditLogsOptions) SetAccount(account string) *GetAuditLogsOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetAscending : Allow user to set Ascending
func (options *GetAuditLogsOptions) SetAscending(ascending string) *GetAuditLogsOptions {
	options.Ascending = core.StringPtr(ascending)
	return options
}

// SetStartat : Allow user to set Startat
func (options *GetAuditLogsOptions) SetStartat(startat string) *GetAuditLogsOptions {
	options.Startat = core.StringPtr(startat)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetAuditLogsOptions) SetOffset(offset int64) *GetAuditLogsOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetAuditLogsOptions) SetLimit(limit int64) *GetAuditLogsOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetAuditLogsOptions) SetHeaders(param map[string]string) *GetAuditLogsOptions {
	options.Headers = param
	return options
}

// GetCatalogEntryOptions : The GetCatalogEntry options.
type GetCatalogEntryOptions struct {
	// The catalog entry's unqiue ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// A GET call by default returns a basic set of properties. To include other properties, you must add this parameter. A
	// wildcard (`*`) includes all properties for an object, for example `GET /id?include=*`. To include specific metadata
	// fields, separate each field with a colon (:), for example `GET /id?include=metadata.ui:metadata.pricing`.
	Include *string `json:"include,omitempty"`

	// Return the data strings in the specified langauge. By default the strings returned are of the language preferred by
	// your browser through the Accept-Langauge header, which allows an override of the header. Languages are specified in
	// standard form, such as `en-us`. To include all languages use a wildcard (*).
	Languages *string `json:"languages,omitempty"`

	// Returns all available fields for all languages. Use the value `?complete=true` as shortcut for
	// ?include=*&languages=*.
	Complete *string `json:"complete,omitempty"`

	// Return the children down to the requested depth. Use * to include the entire children tree. If there are more
	// children than the maximum permitted an error will be returned. Be judicious with this as it can cause a large number
	// of database accesses and can result in a large amount of data returned.
	Depth *int64 `json:"depth,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogEntryOptions : Instantiate GetCatalogEntryOptions
func (*GlobalCatalogV1) NewGetCatalogEntryOptions(id string) *GetCatalogEntryOptions {
	return &GetCatalogEntryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetCatalogEntryOptions) SetID(id string) *GetCatalogEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetCatalogEntryOptions) SetAccount(account string) *GetCatalogEntryOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetCatalogEntryOptions) SetInclude(include string) *GetCatalogEntryOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetLanguages : Allow user to set Languages
func (options *GetCatalogEntryOptions) SetLanguages(languages string) *GetCatalogEntryOptions {
	options.Languages = core.StringPtr(languages)
	return options
}

// SetComplete : Allow user to set Complete
func (options *GetCatalogEntryOptions) SetComplete(complete string) *GetCatalogEntryOptions {
	options.Complete = core.StringPtr(complete)
	return options
}

// SetDepth : Allow user to set Depth
func (options *GetCatalogEntryOptions) SetDepth(depth int64) *GetCatalogEntryOptions {
	options.Depth = core.Int64Ptr(depth)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogEntryOptions) SetHeaders(param map[string]string) *GetCatalogEntryOptions {
	options.Headers = param
	return options
}

// GetChildObjectsOptions : The GetChildObjects options.
type GetChildObjectsOptions struct {
	// The parent catalog entry's ID.
	ID *string `json:"id" validate:"required"`

	// The **kind** of child catalog entries to search for. A wildcard (*) includes all child catalog entries for all
	// kinds, for example `GET /service_name/_*`.
	Kind *string `json:"kind" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// A colon (:) separated list of properties to include. A GET call by defaults return a limited set of properties. To
	// include other properties, you must add the include parameter.  A wildcard (*) includes all properties.
	Include *string `json:"include,omitempty"`

	// A query filter, for example, `q=kind:iaas IBM`  will filter on entries of **kind** iaas that has `IBM` in their
	// name, display name, or description.
	Q *string `json:"q,omitempty"`

	// The field on which to sort the output. By default by name. Available fields are **name**, **kind**, and
	// **provider**.
	SortBy *string `json:"sort-by,omitempty"`

	// The sort order. The default is false, which is ascending.
	Descending *string `json:"descending,omitempty"`

	// Return the data strings in the specified langauge. By default the strings returned are of the language preferred by
	// your browser through the Accept-Langauge header. This allows an override of the header. Languages are specified in
	// standard form, such as `en-us`. To include all languages use the wildcard (*).
	Languages *string `json:"languages,omitempty"`

	// Use the value `?complete=true` as shortcut for ?include=*&languages=*.
	Complete *string `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetChildObjectsOptions : Instantiate GetChildObjectsOptions
func (*GlobalCatalogV1) NewGetChildObjectsOptions(id string, kind string) *GetChildObjectsOptions {
	return &GetChildObjectsOptions{
		ID: core.StringPtr(id),
		Kind: core.StringPtr(kind),
	}
}

// SetID : Allow user to set ID
func (options *GetChildObjectsOptions) SetID(id string) *GetChildObjectsOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetKind : Allow user to set Kind
func (options *GetChildObjectsOptions) SetKind(kind string) *GetChildObjectsOptions {
	options.Kind = core.StringPtr(kind)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetChildObjectsOptions) SetAccount(account string) *GetChildObjectsOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetInclude : Allow user to set Include
func (options *GetChildObjectsOptions) SetInclude(include string) *GetChildObjectsOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetQ : Allow user to set Q
func (options *GetChildObjectsOptions) SetQ(q string) *GetChildObjectsOptions {
	options.Q = core.StringPtr(q)
	return options
}

// SetSortBy : Allow user to set SortBy
func (options *GetChildObjectsOptions) SetSortBy(sortBy string) *GetChildObjectsOptions {
	options.SortBy = core.StringPtr(sortBy)
	return options
}

// SetDescending : Allow user to set Descending
func (options *GetChildObjectsOptions) SetDescending(descending string) *GetChildObjectsOptions {
	options.Descending = core.StringPtr(descending)
	return options
}

// SetLanguages : Allow user to set Languages
func (options *GetChildObjectsOptions) SetLanguages(languages string) *GetChildObjectsOptions {
	options.Languages = core.StringPtr(languages)
	return options
}

// SetComplete : Allow user to set Complete
func (options *GetChildObjectsOptions) SetComplete(complete string) *GetChildObjectsOptions {
	options.Complete = core.StringPtr(complete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetChildObjectsOptions) SetHeaders(param map[string]string) *GetChildObjectsOptions {
	options.Headers = param
	return options
}

// GetPricingOptions : The GetPricing options.
type GetPricingOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPricingOptions : Instantiate GetPricingOptions
func (*GlobalCatalogV1) NewGetPricingOptions(id string) *GetPricingOptions {
	return &GetPricingOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetPricingOptions) SetID(id string) *GetPricingOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetPricingOptions) SetAccount(account string) *GetPricingOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPricingOptions) SetHeaders(param map[string]string) *GetPricingOptions {
	options.Headers = param
	return options
}

// GetVisibilityOptions : The GetVisibility options.
type GetVisibilityOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVisibilityOptions : Instantiate GetVisibilityOptions
func (*GlobalCatalogV1) NewGetVisibilityOptions(id string) *GetVisibilityOptions {
	return &GetVisibilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetVisibilityOptions) SetID(id string) *GetVisibilityOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *GetVisibilityOptions) SetAccount(account string) *GetVisibilityOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVisibilityOptions) SetHeaders(param map[string]string) *GetVisibilityOptions {
	options.Headers = param
	return options
}

// I18N : Language specific translation of translation properties, like label and description.
type I18N struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]*Strings
}


// SetProperty allows the user to set an arbitrary property on an instance of I18N
func (o *I18N) SetProperty(key string, value *Strings) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]*Strings)
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of I18N
func (o *I18N) GetProperty(key string) *Strings {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of I18N
func (o *I18N) GetProperties() map[string]*Strings {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of I18N
func (o *I18N) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)	
	return
}

// UnmarshalI18N constructs an instance of I18N from the specified map.
func UnmarshalI18N(m map[string]interface{}) (result *I18N, err error) {
	m = core.CopyMap(m)
	obj := new(I18N)
	for k := range m {
		v, e := UnmarshalStringsAsProperty(m, k)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	result = obj
	return
}

// UnmarshalI18NSlice unmarshals a slice of I18N instances from the specified list of maps.
func UnmarshalI18NSlice(s []interface{}) (slice []I18N, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'I18N'")
			return
		}
		obj, e := UnmarshalI18N(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalI18NAsProperty unmarshals an instance of I18N that is stored as a property
// within the specified map.
func UnmarshalI18NAsProperty(m map[string]interface{}, propertyName string) (result *I18N, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'I18N'", propertyName)
			return
		}
		result, err = UnmarshalI18N(objMap)
	}
	return
}

// UnmarshalI18NSliceAsProperty unmarshals a slice of I18N instances that are stored as a property
// within the specified map.
func UnmarshalI18NSliceAsProperty(m map[string]interface{}, propertyName string) (slice []I18N, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'I18N'", propertyName)
			return
		}
		slice, err = UnmarshalI18NSlice(vSlice)
	}
	return
}

// Image : Image annotation for this catalog entry. The image is a URL.
type Image struct {
	// URL for the large, default image.
	Image *string `json:"image" validate:"required"`

	// URL for a small image.
	SmallImage *string `json:"small_image,omitempty"`

	// URL for a medium image.
	MediumImage *string `json:"medium_image,omitempty"`

	// URL for a featured image.
	FeatureImage *string `json:"feature_image,omitempty"`
}


// NewImage : Instantiate Image (Generic Model Constructor)
func (*GlobalCatalogV1) NewImage(image string) (model *Image, err error) {
	model = &Image{
		Image: core.StringPtr(image),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalImage constructs an instance of Image from the specified map.
func UnmarshalImage(m map[string]interface{}) (result *Image, err error) {
	obj := new(Image)
	obj.Image, err = core.UnmarshalString(m, "image")
	if err != nil {
		return
	}
	obj.SmallImage, err = core.UnmarshalString(m, "small_image")
	if err != nil {
		return
	}
	obj.MediumImage, err = core.UnmarshalString(m, "medium_image")
	if err != nil {
		return
	}
	obj.FeatureImage, err = core.UnmarshalString(m, "feature_image")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalImageSlice unmarshals a slice of Image instances from the specified list of maps.
func UnmarshalImageSlice(s []interface{}) (slice []Image, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Image'")
			return
		}
		obj, e := UnmarshalImage(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalImageAsProperty unmarshals an instance of Image that is stored as a property
// within the specified map.
func UnmarshalImageAsProperty(m map[string]interface{}, propertyName string) (result *Image, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Image'", propertyName)
			return
		}
		result, err = UnmarshalImage(objMap)
	}
	return
}

// UnmarshalImageSliceAsProperty unmarshals a slice of Image instances that are stored as a property
// within the specified map.
func UnmarshalImageSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Image, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Image'", propertyName)
			return
		}
		slice, err = UnmarshalImageSlice(vSlice)
	}
	return
}

// ListArtifactsOptions : The ListArtifacts options.
type ListArtifactsOptions struct {
	// The object's unique ID.
	ObjectID *string `json:"object_id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListArtifactsOptions : Instantiate ListArtifactsOptions
func (*GlobalCatalogV1) NewListArtifactsOptions(objectID string) *ListArtifactsOptions {
	return &ListArtifactsOptions{
		ObjectID: core.StringPtr(objectID),
	}
}

// SetObjectID : Allow user to set ObjectID
func (options *ListArtifactsOptions) SetObjectID(objectID string) *ListArtifactsOptions {
	options.ObjectID = core.StringPtr(objectID)
	return options
}

// SetAccount : Allow user to set Account
func (options *ListArtifactsOptions) SetAccount(account string) *ListArtifactsOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListArtifactsOptions) SetHeaders(param map[string]string) *ListArtifactsOptions {
	options.Headers = param
	return options
}

// ListCatalogEntriesOptions : The ListCatalogEntries options.
type ListCatalogEntriesOptions struct {
	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// A GET call by default returns a basic set of properties. To include other properties, you must add this parameter. A
	// wildcard (`*`) includes all properties for an object, for example `GET /?include=*`. To include specific metadata
	// fields, separate each field with a colon (:), for example `GET /?include=metadata.ui:metadata.pricing`.
	Include *string `json:"include,omitempty"`

	// Searches the catalog entries for keywords. Add filters to refine your search. A query filter, for example,
	// `q=kind:iaas service_name rc:true`, filters entries of kind iaas with metadata.service.rc_compatible set to true and
	//  have a service name is in their name, display name, or description.  Valid tags are **kind**:<string>,
	// **tag**:<strging>, **rc**:[true|false], **iam**:[true|false], **active**:[true|false], **geo**:<string>, and
	// **price**:<string>.
	Q *string `json:"q,omitempty"`

	// The field on which the output is sorted. Sorts by default by **name** property. Available fields are **name**,
	// **displayname** (overview_ui.display_name), **kind**, **provider** (provider.name), **sbsindex**
	// (metadata.ui.side_by_side_index), and the time **created**, and **updated**.
	SortBy *string `json:"sort-by,omitempty"`

	// Sets the sort order. The default is false, which is ascending.
	Descending *string `json:"descending,omitempty"`

	// Return the data strings in a specified langauge. By default, the strings returned are of the language preferred by
	// your browser through the Accept-Langauge header, which allows an override of the header. Languages are specified in
	// standard form, such as `en-us`. To include all languages use a wildcard (*).
	Languages *string `json:"languages,omitempty"`

	// Returns all available fields for all languages. Use the value `?complete=true` as shortcut for
	// ?include=*&languages=*.
	Complete *string `json:"complete,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCatalogEntriesOptions : Instantiate ListCatalogEntriesOptions
func (*GlobalCatalogV1) NewListCatalogEntriesOptions() *ListCatalogEntriesOptions {
	return &ListCatalogEntriesOptions{}
}

// SetAccount : Allow user to set Account
func (options *ListCatalogEntriesOptions) SetAccount(account string) *ListCatalogEntriesOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetInclude : Allow user to set Include
func (options *ListCatalogEntriesOptions) SetInclude(include string) *ListCatalogEntriesOptions {
	options.Include = core.StringPtr(include)
	return options
}

// SetQ : Allow user to set Q
func (options *ListCatalogEntriesOptions) SetQ(q string) *ListCatalogEntriesOptions {
	options.Q = core.StringPtr(q)
	return options
}

// SetSortBy : Allow user to set SortBy
func (options *ListCatalogEntriesOptions) SetSortBy(sortBy string) *ListCatalogEntriesOptions {
	options.SortBy = core.StringPtr(sortBy)
	return options
}

// SetDescending : Allow user to set Descending
func (options *ListCatalogEntriesOptions) SetDescending(descending string) *ListCatalogEntriesOptions {
	options.Descending = core.StringPtr(descending)
	return options
}

// SetLanguages : Allow user to set Languages
func (options *ListCatalogEntriesOptions) SetLanguages(languages string) *ListCatalogEntriesOptions {
	options.Languages = core.StringPtr(languages)
	return options
}

// SetComplete : Allow user to set Complete
func (options *ListCatalogEntriesOptions) SetComplete(complete string) *ListCatalogEntriesOptions {
	options.Complete = core.StringPtr(complete)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogEntriesOptions) SetHeaders(param map[string]string) *ListCatalogEntriesOptions {
	options.Headers = param
	return options
}

// Metrics : Plan-specific cost metrics information.
type Metrics struct {
	// The metric ID or part number.
	MetricID *string `json:"metric_id,omitempty"`

	// The tier model.
	TierModel *string `json:"tier_model,omitempty"`

	// The charge unit name.
	ChargeUnitName *string `json:"charge_unit_name,omitempty"`

	// The charge unit quantity.
	ChargeUnitQuantity *string `json:"charge_unit_quantity,omitempty"`

	// Display name of the resource.
	ResourceDisplayName *string `json:"resource_display_name,omitempty"`

	// Display name of the charge unit.
	ChargeUnitDisplayName *string `json:"charge_unit_display_name,omitempty"`

	// Usage limit for the metric.
	UsageCapQty *int64 `json:"usage_cap_qty,omitempty"`

	// The pricing per metric by country and currency.
	Amounts []Amount `json:"amounts,omitempty"`
}


// UnmarshalMetrics constructs an instance of Metrics from the specified map.
func UnmarshalMetrics(m map[string]interface{}) (result *Metrics, err error) {
	obj := new(Metrics)
	obj.MetricID, err = core.UnmarshalString(m, "metric_id")
	if err != nil {
		return
	}
	obj.TierModel, err = core.UnmarshalString(m, "tier_model")
	if err != nil {
		return
	}
	obj.ChargeUnitName, err = core.UnmarshalString(m, "charge_unit_name")
	if err != nil {
		return
	}
	obj.ChargeUnitQuantity, err = core.UnmarshalString(m, "charge_unit_quantity")
	if err != nil {
		return
	}
	obj.ResourceDisplayName, err = core.UnmarshalString(m, "resource_display_name")
	if err != nil {
		return
	}
	obj.ChargeUnitDisplayName, err = core.UnmarshalString(m, "charge_unit_display_name")
	if err != nil {
		return
	}
	obj.UsageCapQty, err = core.UnmarshalInt64(m, "usage_cap_qty")
	if err != nil {
		return
	}
	obj.Amounts, err = UnmarshalAmountSliceAsProperty(m, "amounts")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalMetricsSlice unmarshals a slice of Metrics instances from the specified list of maps.
func UnmarshalMetricsSlice(s []interface{}) (slice []Metrics, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Metrics'")
			return
		}
		obj, e := UnmarshalMetrics(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalMetricsAsProperty unmarshals an instance of Metrics that is stored as a property
// within the specified map.
func UnmarshalMetricsAsProperty(m map[string]interface{}, propertyName string) (result *Metrics, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Metrics'", propertyName)
			return
		}
		result, err = UnmarshalMetrics(objMap)
	}
	return
}

// UnmarshalMetricsSliceAsProperty unmarshals a slice of Metrics instances that are stored as a property
// within the specified map.
func UnmarshalMetricsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Metrics, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Metrics'", propertyName)
			return
		}
		slice, err = UnmarshalMetricsSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseAlias : Alias-related metadata.
type ObjectMetadataBaseAlias struct {
	// Type of alias.
	Type *string `json:"type,omitempty"`

	// Points to the plan that this object is an alias for.
	PlanID *string `json:"plan_id,omitempty"`
}


// UnmarshalObjectMetadataBaseAlias constructs an instance of ObjectMetadataBaseAlias from the specified map.
func UnmarshalObjectMetadataBaseAlias(m map[string]interface{}) (result *ObjectMetadataBaseAlias, err error) {
	obj := new(ObjectMetadataBaseAlias)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseAliasSlice unmarshals a slice of ObjectMetadataBaseAlias instances from the specified list of maps.
func UnmarshalObjectMetadataBaseAliasSlice(s []interface{}) (slice []ObjectMetadataBaseAlias, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseAlias'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseAlias(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseAliasAsProperty unmarshals an instance of ObjectMetadataBaseAlias that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseAliasAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseAlias, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseAlias'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseAlias(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseAliasSliceAsProperty unmarshals a slice of ObjectMetadataBaseAlias instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseAliasSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseAlias, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseAlias'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseAliasSlice(vSlice)
	}
	return
}

// ObjectMetadataBasePlan : Plan-related metadata.
type ObjectMetadataBasePlan struct {
	// Boolean value that describes whether the service can be bound to an application.
	Bindable *bool `json:"bindable,omitempty"`

	// Boolean value that describes whether the service can be reserved.
	Reservable *bool `json:"reservable,omitempty"`

	// Boolean value that describes whether the service can be used internally.
	AllowInternalUsers *bool `json:"allow_internal_users,omitempty"`

	// Boolean value that describes whether the service can be provisioned asynchronously.
	AsyncProvisioningSupported *bool `json:"async_provisioning_supported,omitempty"`

	// Boolean value that describes whether the service can be unprovisioned asynchronously.
	AsyncUnprovisioningSupported *bool `json:"async_unprovisioning_supported,omitempty"`

	// Test check interval.
	TestCheckInterval *int64 `json:"test_check_interval,omitempty"`

	// Single scope instance.
	SingleScopeInstance *string `json:"single_scope_instance,omitempty"`

	// Boolean value that describes whether the service check is enabled.
	ServiceCheckEnabled *bool `json:"service_check_enabled,omitempty"`

	// If the field is imported from Cloud Foundry, the Cloud Foundry region's GUID. This is a required field. For example,
	// `us-south=123`.
	CfGuid *string `json:"cf_guid,omitempty"`
}


// UnmarshalObjectMetadataBasePlan constructs an instance of ObjectMetadataBasePlan from the specified map.
func UnmarshalObjectMetadataBasePlan(m map[string]interface{}) (result *ObjectMetadataBasePlan, err error) {
	obj := new(ObjectMetadataBasePlan)
	obj.Bindable, err = core.UnmarshalBool(m, "bindable")
	if err != nil {
		return
	}
	obj.Reservable, err = core.UnmarshalBool(m, "reservable")
	if err != nil {
		return
	}
	obj.AllowInternalUsers, err = core.UnmarshalBool(m, "allow_internal_users")
	if err != nil {
		return
	}
	obj.AsyncProvisioningSupported, err = core.UnmarshalBool(m, "async_provisioning_supported")
	if err != nil {
		return
	}
	obj.AsyncUnprovisioningSupported, err = core.UnmarshalBool(m, "async_unprovisioning_supported")
	if err != nil {
		return
	}
	obj.TestCheckInterval, err = core.UnmarshalInt64(m, "test_check_interval")
	if err != nil {
		return
	}
	obj.SingleScopeInstance, err = core.UnmarshalString(m, "single_scope_instance")
	if err != nil {
		return
	}
	obj.ServiceCheckEnabled, err = core.UnmarshalBool(m, "service_check_enabled")
	if err != nil {
		return
	}
	obj.CfGuid, err = core.UnmarshalString(m, "cf_guid")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBasePlanSlice unmarshals a slice of ObjectMetadataBasePlan instances from the specified list of maps.
func UnmarshalObjectMetadataBasePlanSlice(s []interface{}) (slice []ObjectMetadataBasePlan, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBasePlan'")
			return
		}
		obj, e := UnmarshalObjectMetadataBasePlan(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBasePlanAsProperty unmarshals an instance of ObjectMetadataBasePlan that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBasePlanAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBasePlan, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBasePlan'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBasePlan(objMap)
	}
	return
}

// UnmarshalObjectMetadataBasePlanSliceAsProperty unmarshals a slice of ObjectMetadataBasePlan instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBasePlanSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBasePlan, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBasePlan'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBasePlanSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseService : Service-related metadata.
type ObjectMetadataBaseService struct {
	// Type of service.
	Type *string `json:"type,omitempty"`

	// Boolean value that describes whether the service is compatible with Identity and Access Management.
	IamCompatible *bool `json:"iam_compatible,omitempty"`

	// Boolean value that describes whether the service has a unique API key.
	UniqueApiKey *bool `json:"unique_api_key,omitempty"`

	// Boolean value that describes whether the service is provisionable or not. You may need sales or support to create
	// this service.
	Provisionable *bool `json:"provisionable,omitempty"`

	// Boolean value that describes whether the service supports asynchronous provisioning.
	AsyncProvisioningSupported *bool `json:"async_provisioning_supported,omitempty"`

	// Boolean value that describes whether the service supports asynchronous unprovisioning.
	AsyncUnprovisioningSupported *bool `json:"async_unprovisioning_supported,omitempty"`

	// If the field is imported from Cloud Foundry, the Cloud Foundry region's GUID. This is a required field. For example,
	// `us-south=123`.
	CfGuid *string `json:"cf_guid,omitempty"`

	// Boolean value that describes whether you can create bindings for this service.
	Bindable *bool `json:"bindable,omitempty"`

	// Service dependencies.
	Requires []string `json:"requires,omitempty"`

	// Boolean value that describes whether the service supports upgrade or downgrade for some plans.
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`

	// String that describes whether the service is active or inactive.
	State *string `json:"state,omitempty"`

	// Boolean value that describes whether the service check is enabled.
	ServiceCheckEnabled *bool `json:"service_check_enabled,omitempty"`

	// Test check interval.
	TestCheckInterval *int64 `json:"test_check_interval,omitempty"`

	// Boolean value that describes whether the service supports service keys.
	ServiceKeySupported *bool `json:"service_key_supported,omitempty"`
}


// UnmarshalObjectMetadataBaseService constructs an instance of ObjectMetadataBaseService from the specified map.
func UnmarshalObjectMetadataBaseService(m map[string]interface{}) (result *ObjectMetadataBaseService, err error) {
	obj := new(ObjectMetadataBaseService)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.IamCompatible, err = core.UnmarshalBool(m, "iam_compatible")
	if err != nil {
		return
	}
	obj.UniqueApiKey, err = core.UnmarshalBool(m, "unique_api_key")
	if err != nil {
		return
	}
	obj.Provisionable, err = core.UnmarshalBool(m, "provisionable")
	if err != nil {
		return
	}
	obj.AsyncProvisioningSupported, err = core.UnmarshalBool(m, "async_provisioning_supported")
	if err != nil {
		return
	}
	obj.AsyncUnprovisioningSupported, err = core.UnmarshalBool(m, "async_unprovisioning_supported")
	if err != nil {
		return
	}
	obj.CfGuid, err = core.UnmarshalString(m, "cf_guid")
	if err != nil {
		return
	}
	obj.Bindable, err = core.UnmarshalBool(m, "bindable")
	if err != nil {
		return
	}
	obj.Requires, err = core.UnmarshalStringSlice(m, "requires")
	if err != nil {
		return
	}
	obj.PlanUpdateable, err = core.UnmarshalBool(m, "plan_updateable")
	if err != nil {
		return
	}
	obj.State, err = core.UnmarshalString(m, "state")
	if err != nil {
		return
	}
	obj.ServiceCheckEnabled, err = core.UnmarshalBool(m, "service_check_enabled")
	if err != nil {
		return
	}
	obj.TestCheckInterval, err = core.UnmarshalInt64(m, "test_check_interval")
	if err != nil {
		return
	}
	obj.ServiceKeySupported, err = core.UnmarshalBool(m, "service_key_supported")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseServiceSlice unmarshals a slice of ObjectMetadataBaseService instances from the specified list of maps.
func UnmarshalObjectMetadataBaseServiceSlice(s []interface{}) (slice []ObjectMetadataBaseService, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseService'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseService(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseServiceAsProperty unmarshals an instance of ObjectMetadataBaseService that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseServiceAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseService, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseService'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseService(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseServiceSliceAsProperty unmarshals a slice of ObjectMetadataBaseService instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseServiceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseService, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseService'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseServiceSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseSla : Service Level Agreement related metadata.
type ObjectMetadataBaseSla struct {
	// Required Service License Agreement Terms of Use.
	Terms *string `json:"terms,omitempty"`

	// Required deployment type. Valid values are dedicated, local, or public. It can be Single or Multi tennancy, more
	// specifically on a Server, VM, Physical, or Pod.
	Tenancy *string `json:"tenancy,omitempty"`

	// Provisioning reliability, for example, 99.95.
	Provisioning *string `json:"provisioning,omitempty"`

	// Uptime reliability of the service, for example, 99.95.
	Responsiveness *string `json:"responsiveness,omitempty"`

	// SLA Disaster Recovery-related metadata.
	Dr *ObjectMetadataBaseSlaDr `json:"dr,omitempty"`
}


// UnmarshalObjectMetadataBaseSla constructs an instance of ObjectMetadataBaseSla from the specified map.
func UnmarshalObjectMetadataBaseSla(m map[string]interface{}) (result *ObjectMetadataBaseSla, err error) {
	obj := new(ObjectMetadataBaseSla)
	obj.Terms, err = core.UnmarshalString(m, "terms")
	if err != nil {
		return
	}
	obj.Tenancy, err = core.UnmarshalString(m, "tenancy")
	if err != nil {
		return
	}
	obj.Provisioning, err = core.UnmarshalString(m, "provisioning")
	if err != nil {
		return
	}
	obj.Responsiveness, err = core.UnmarshalString(m, "responsiveness")
	if err != nil {
		return
	}
	obj.Dr, err = UnmarshalObjectMetadataBaseSlaDrAsProperty(m, "dr")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseSlaSlice unmarshals a slice of ObjectMetadataBaseSla instances from the specified list of maps.
func UnmarshalObjectMetadataBaseSlaSlice(s []interface{}) (slice []ObjectMetadataBaseSla, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseSla'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseSla(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseSlaAsProperty unmarshals an instance of ObjectMetadataBaseSla that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseSlaAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseSla, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseSla'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseSla(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseSlaSliceAsProperty unmarshals a slice of ObjectMetadataBaseSla instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseSlaSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseSla, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseSla'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseSlaSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseSlaDr : SLA Disaster Recovery-related metadata.
type ObjectMetadataBaseSlaDr struct {
	// Required boolean value that describes whether disaster recovery is on.
	Dr *bool `json:"dr,omitempty"`

	// Description of the disaster recovery implementation.
	Description *string `json:"description,omitempty"`
}


// UnmarshalObjectMetadataBaseSlaDr constructs an instance of ObjectMetadataBaseSlaDr from the specified map.
func UnmarshalObjectMetadataBaseSlaDr(m map[string]interface{}) (result *ObjectMetadataBaseSlaDr, err error) {
	obj := new(ObjectMetadataBaseSlaDr)
	obj.Dr, err = core.UnmarshalBool(m, "dr")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseSlaDrSlice unmarshals a slice of ObjectMetadataBaseSlaDr instances from the specified list of maps.
func UnmarshalObjectMetadataBaseSlaDrSlice(s []interface{}) (slice []ObjectMetadataBaseSlaDr, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseSlaDr'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseSlaDr(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseSlaDrAsProperty unmarshals an instance of ObjectMetadataBaseSlaDr that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseSlaDrAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseSlaDr, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseSlaDr'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseSlaDr(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseSlaDrSliceAsProperty unmarshals a slice of ObjectMetadataBaseSlaDr instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseSlaDrSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseSlaDr, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseSlaDr'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseSlaDrSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseTemplate : Template-related metadata.
type ObjectMetadataBaseTemplate struct {
	// List of required offering or plan IDs.
	Services []string `json:"services,omitempty"`

	// Cloud Foundry instance memory value.
	DefaultMemory *int64 `json:"default_memory,omitempty"`

	// Start Command.
	StartCmd *string `json:"start_cmd,omitempty"`

	// Location of your applications source files.
	Source *ObjectMetadataBaseTemplateSource `json:"source,omitempty"`

	// ID of the runtime.
	RuntimeCatalogID *string `json:"runtime_catalog_id,omitempty"`

	// ID of the Cloud Foundry runtime.
	CfRuntimeID *string `json:"cf_runtime_id,omitempty"`

	// ID of the boilerplate or template.
	TemplateID *string `json:"template_id,omitempty"`

	// File path to the executable file for the template.
	ExecutableFile *string `json:"executable_file,omitempty"`

	// ID of the buildpack used by the template.
	Buildpack *string `json:"buildpack,omitempty"`

	// Environment variables for the template.
	EnvironmentVariables *ObjectMetadataBaseTemplateEnvironmentVariables `json:"environment_variables,omitempty"`
}


// UnmarshalObjectMetadataBaseTemplate constructs an instance of ObjectMetadataBaseTemplate from the specified map.
func UnmarshalObjectMetadataBaseTemplate(m map[string]interface{}) (result *ObjectMetadataBaseTemplate, err error) {
	obj := new(ObjectMetadataBaseTemplate)
	obj.Services, err = core.UnmarshalStringSlice(m, "services")
	if err != nil {
		return
	}
	obj.DefaultMemory, err = core.UnmarshalInt64(m, "default_memory")
	if err != nil {
		return
	}
	obj.StartCmd, err = core.UnmarshalString(m, "start_cmd")
	if err != nil {
		return
	}
	obj.Source, err = UnmarshalObjectMetadataBaseTemplateSourceAsProperty(m, "source")
	if err != nil {
		return
	}
	obj.RuntimeCatalogID, err = core.UnmarshalString(m, "runtime_catalog_id")
	if err != nil {
		return
	}
	obj.CfRuntimeID, err = core.UnmarshalString(m, "cf_runtime_id")
	if err != nil {
		return
	}
	obj.TemplateID, err = core.UnmarshalString(m, "template_id")
	if err != nil {
		return
	}
	obj.ExecutableFile, err = core.UnmarshalString(m, "executable_file")
	if err != nil {
		return
	}
	obj.Buildpack, err = core.UnmarshalString(m, "buildpack")
	if err != nil {
		return
	}
	obj.EnvironmentVariables, err = UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesAsProperty(m, "environment_variables")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseTemplateSlice unmarshals a slice of ObjectMetadataBaseTemplate instances from the specified list of maps.
func UnmarshalObjectMetadataBaseTemplateSlice(s []interface{}) (slice []ObjectMetadataBaseTemplate, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseTemplate'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseTemplate(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateAsProperty unmarshals an instance of ObjectMetadataBaseTemplate that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseTemplate, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseTemplate'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseTemplate(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateSliceAsProperty unmarshals a slice of ObjectMetadataBaseTemplate instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseTemplate, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseTemplate'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseTemplateSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseTemplateEnvironmentVariables : Environment variables for the template.
type ObjectMetadataBaseTemplateEnvironmentVariables struct {
	// Key is the editable first string in a key:value pair of environment variables.
	Key *string `json:"_key_,omitempty"`
}


// UnmarshalObjectMetadataBaseTemplateEnvironmentVariables constructs an instance of ObjectMetadataBaseTemplateEnvironmentVariables from the specified map.
func UnmarshalObjectMetadataBaseTemplateEnvironmentVariables(m map[string]interface{}) (result *ObjectMetadataBaseTemplateEnvironmentVariables, err error) {
	obj := new(ObjectMetadataBaseTemplateEnvironmentVariables)
	obj.Key, err = core.UnmarshalString(m, "_key_")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesSlice unmarshals a slice of ObjectMetadataBaseTemplateEnvironmentVariables instances from the specified list of maps.
func UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesSlice(s []interface{}) (slice []ObjectMetadataBaseTemplateEnvironmentVariables, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseTemplateEnvironmentVariables'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseTemplateEnvironmentVariables(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesAsProperty unmarshals an instance of ObjectMetadataBaseTemplateEnvironmentVariables that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseTemplateEnvironmentVariables, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseTemplateEnvironmentVariables'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseTemplateEnvironmentVariables(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesSliceAsProperty unmarshals a slice of ObjectMetadataBaseTemplateEnvironmentVariables instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseTemplateEnvironmentVariables, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseTemplateEnvironmentVariables'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseTemplateEnvironmentVariablesSlice(vSlice)
	}
	return
}

// ObjectMetadataBaseTemplateSource : Location of your applications source files.
type ObjectMetadataBaseTemplateSource struct {
	// Path to your application.
	Path *string `json:"path,omitempty"`

	// Type of source, for example, git.
	Type *string `json:"type,omitempty"`

	// URL to source.
	URL *string `json:"url,omitempty"`
}


// UnmarshalObjectMetadataBaseTemplateSource constructs an instance of ObjectMetadataBaseTemplateSource from the specified map.
func UnmarshalObjectMetadataBaseTemplateSource(m map[string]interface{}) (result *ObjectMetadataBaseTemplateSource, err error) {
	obj := new(ObjectMetadataBaseTemplateSource)
	obj.Path, err = core.UnmarshalString(m, "path")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "url")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataBaseTemplateSourceSlice unmarshals a slice of ObjectMetadataBaseTemplateSource instances from the specified list of maps.
func UnmarshalObjectMetadataBaseTemplateSourceSlice(s []interface{}) (slice []ObjectMetadataBaseTemplateSource, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataBaseTemplateSource'")
			return
		}
		obj, e := UnmarshalObjectMetadataBaseTemplateSource(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateSourceAsProperty unmarshals an instance of ObjectMetadataBaseTemplateSource that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateSourceAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataBaseTemplateSource, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataBaseTemplateSource'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataBaseTemplateSource(objMap)
	}
	return
}

// UnmarshalObjectMetadataBaseTemplateSourceSliceAsProperty unmarshals a slice of ObjectMetadataBaseTemplateSource instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataBaseTemplateSourceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataBaseTemplateSource, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataBaseTemplateSource'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataBaseTemplateSourceSlice(vSlice)
	}
	return
}

// ObjectMetadataSet : Model used to describe metadata object that can be set.
type ObjectMetadataSet struct {
	// Boolean value that describes whether the service is compatible with the Resource Controller.
	RcCompatible *bool `json:"rc_compatible,omitempty"`

	// Information related to the UI presentation associated with a catalog entry.
	Ui *UIMetaData `json:"ui,omitempty"`

	// Compliance information for HIPAA and PCI.
	Compliance []string `json:"compliance,omitempty"`

	// Service-related metadata.
	Service *ObjectMetadataBaseService `json:"service,omitempty"`

	// Plan-related metadata.
	Plan *ObjectMetadataBasePlan `json:"plan,omitempty"`

	// Template-related metadata.
	Template *ObjectMetadataBaseTemplate `json:"template,omitempty"`

	// Alias-related metadata.
	Alias *ObjectMetadataBaseAlias `json:"alias,omitempty"`

	// Service Level Agreement related metadata.
	Sla *ObjectMetadataBaseSla `json:"sla,omitempty"`

	// Callback-related information associated with a catalog entry.
	Callbacks *Callbacks `json:"callbacks,omitempty"`

	// Optional version of the object.
	Version *string `json:"version,omitempty"`

	// The original name of the object.
	OriginalName *string `json:"original_name,omitempty"`

	// Additional information.
	Other interface{} `json:"other,omitempty"`

	// Pricing-related information.
	Pricing *PricingSet `json:"pricing,omitempty"`

	// Deployment-related metadata.
	Deployment *DeploymentBase `json:"deployment,omitempty"`
}


// UnmarshalObjectMetadataSet constructs an instance of ObjectMetadataSet from the specified map.
func UnmarshalObjectMetadataSet(m map[string]interface{}) (result *ObjectMetadataSet, err error) {
	obj := new(ObjectMetadataSet)
	obj.RcCompatible, err = core.UnmarshalBool(m, "rc_compatible")
	if err != nil {
		return
	}
	obj.Ui, err = UnmarshalUIMetaDataAsProperty(m, "ui")
	if err != nil {
		return
	}
	obj.Compliance, err = core.UnmarshalStringSlice(m, "compliance")
	if err != nil {
		return
	}
	obj.Service, err = UnmarshalObjectMetadataBaseServiceAsProperty(m, "service")
	if err != nil {
		return
	}
	obj.Plan, err = UnmarshalObjectMetadataBasePlanAsProperty(m, "plan")
	if err != nil {
		return
	}
	obj.Template, err = UnmarshalObjectMetadataBaseTemplateAsProperty(m, "template")
	if err != nil {
		return
	}
	obj.Alias, err = UnmarshalObjectMetadataBaseAliasAsProperty(m, "alias")
	if err != nil {
		return
	}
	obj.Sla, err = UnmarshalObjectMetadataBaseSlaAsProperty(m, "sla")
	if err != nil {
		return
	}
	obj.Callbacks, err = UnmarshalCallbacksAsProperty(m, "callbacks")
	if err != nil {
		return
	}
	obj.Version, err = core.UnmarshalString(m, "version")
	if err != nil {
		return
	}
	obj.OriginalName, err = core.UnmarshalString(m, "original_name")
	if err != nil {
		return
	}
	obj.Other, err = core.UnmarshalAny(m, "other")
	if err != nil {
		return
	}
	obj.Pricing, err = UnmarshalPricingSetAsProperty(m, "pricing")
	if err != nil {
		return
	}
	obj.Deployment, err = UnmarshalDeploymentBaseAsProperty(m, "deployment")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalObjectMetadataSetSlice unmarshals a slice of ObjectMetadataSet instances from the specified list of maps.
func UnmarshalObjectMetadataSetSlice(s []interface{}) (slice []ObjectMetadataSet, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'ObjectMetadataSet'")
			return
		}
		obj, e := UnmarshalObjectMetadataSet(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalObjectMetadataSetAsProperty unmarshals an instance of ObjectMetadataSet that is stored as a property
// within the specified map.
func UnmarshalObjectMetadataSetAsProperty(m map[string]interface{}, propertyName string) (result *ObjectMetadataSet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'ObjectMetadataSet'", propertyName)
			return
		}
		result, err = UnmarshalObjectMetadataSet(objMap)
	}
	return
}

// UnmarshalObjectMetadataSetSliceAsProperty unmarshals a slice of ObjectMetadataSet instances that are stored as a property
// within the specified map.
func UnmarshalObjectMetadataSetSliceAsProperty(m map[string]interface{}, propertyName string) (slice []ObjectMetadataSet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'ObjectMetadataSet'", propertyName)
			return
		}
		slice, err = UnmarshalObjectMetadataSetSlice(vSlice)
	}
	return
}

// Overview : Overview is nested in the top level. The key value pair is `[_language_]overview_ui`.
type Overview struct {
	// The translated display name.
	DisplayName *string `json:"display_name" validate:"required"`

	// The translated long description.
	LongDescription *string `json:"long_description" validate:"required"`

	// The translated description.
	Description *string `json:"description" validate:"required"`
}


// NewOverview : Instantiate Overview (Generic Model Constructor)
func (*GlobalCatalogV1) NewOverview(displayName string, longDescription string, description string) (model *Overview, err error) {
	model = &Overview{
		DisplayName: core.StringPtr(displayName),
		LongDescription: core.StringPtr(longDescription),
		Description: core.StringPtr(description),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalOverview constructs an instance of Overview from the specified map.
func UnmarshalOverview(m map[string]interface{}) (result *Overview, err error) {
	obj := new(Overview)
	obj.DisplayName, err = core.UnmarshalString(m, "display_name")
	if err != nil {
		return
	}
	obj.LongDescription, err = core.UnmarshalString(m, "long_description")
	if err != nil {
		return
	}
	obj.Description, err = core.UnmarshalString(m, "description")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalOverviewSlice unmarshals a slice of Overview instances from the specified list of maps.
func UnmarshalOverviewSlice(s []interface{}) (slice []Overview, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Overview'")
			return
		}
		obj, e := UnmarshalOverview(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalOverviewAsProperty unmarshals an instance of Overview that is stored as a property
// within the specified map.
func UnmarshalOverviewAsProperty(m map[string]interface{}, propertyName string) (result *Overview, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Overview'", propertyName)
			return
		}
		result, err = UnmarshalOverview(objMap)
	}
	return
}

// UnmarshalOverviewSliceAsProperty unmarshals a slice of Overview instances that are stored as a property
// within the specified map.
func UnmarshalOverviewSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Overview, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Overview'", propertyName)
			return
		}
		slice, err = UnmarshalOverviewSlice(vSlice)
	}
	return
}

// OverviewUI : Overview is nested in the top level. The key value pair is `[_language_]overview_ui`.
type OverviewUI struct {

	// Allows users to set arbitrary properties
	additionalProperties map[string]*Overview
}


// SetProperty allows the user to set an arbitrary property on an instance of OverviewUI
func (o *OverviewUI) SetProperty(key string, value *Overview) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]*Overview)
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of OverviewUI
func (o *OverviewUI) GetProperty(key string) *Overview {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of OverviewUI
func (o *OverviewUI) GetProperties() map[string]*Overview {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of OverviewUI
func (o *OverviewUI) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	buffer, err = json.Marshal(m)	
	return
}

// UnmarshalOverviewUI constructs an instance of OverviewUI from the specified map.
func UnmarshalOverviewUI(m map[string]interface{}) (result *OverviewUI, err error) {
	m = core.CopyMap(m)
	obj := new(OverviewUI)
	for k := range m {
		v, e := UnmarshalOverviewAsProperty(m, k)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	result = obj
	return
}

// UnmarshalOverviewUISlice unmarshals a slice of OverviewUI instances from the specified list of maps.
func UnmarshalOverviewUISlice(s []interface{}) (slice []OverviewUI, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'OverviewUI'")
			return
		}
		obj, e := UnmarshalOverviewUI(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalOverviewUIAsProperty unmarshals an instance of OverviewUI that is stored as a property
// within the specified map.
func UnmarshalOverviewUIAsProperty(m map[string]interface{}, propertyName string) (result *OverviewUI, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'OverviewUI'", propertyName)
			return
		}
		result, err = UnmarshalOverviewUI(objMap)
	}
	return
}

// UnmarshalOverviewUISliceAsProperty unmarshals a slice of OverviewUI instances that are stored as a property
// within the specified map.
func UnmarshalOverviewUISliceAsProperty(m map[string]interface{}, propertyName string) (slice []OverviewUI, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'OverviewUI'", propertyName)
			return
		}
		slice, err = UnmarshalOverviewUISlice(vSlice)
	}
	return
}

// Price : Pricing-related information.
type Price struct {
	// Pricing tier.
	QuantityTier *int64 `json:"quantity_tier,omitempty"`

	// Price in the selected currency.
	Price *float64 `json:"Price,omitempty"`
}


// UnmarshalPrice constructs an instance of Price from the specified map.
func UnmarshalPrice(m map[string]interface{}) (result *Price, err error) {
	obj := new(Price)
	obj.QuantityTier, err = core.UnmarshalInt64(m, "quantity_tier")
	if err != nil {
		return
	}
	obj.Price, err = core.UnmarshalFloat64(m, "Price")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPriceSlice unmarshals a slice of Price instances from the specified list of maps.
func UnmarshalPriceSlice(s []interface{}) (slice []Price, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Price'")
			return
		}
		obj, e := UnmarshalPrice(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPriceAsProperty unmarshals an instance of Price that is stored as a property
// within the specified map.
func UnmarshalPriceAsProperty(m map[string]interface{}, propertyName string) (result *Price, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Price'", propertyName)
			return
		}
		result, err = UnmarshalPrice(objMap)
	}
	return
}

// UnmarshalPriceSliceAsProperty unmarshals a slice of Price instances that are stored as a property
// within the specified map.
func UnmarshalPriceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Price, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Price'", propertyName)
			return
		}
		slice, err = UnmarshalPriceSlice(vSlice)
	}
	return
}

// PricingGet : Pricing-related information.
type PricingGet struct {
	// Type of plan. Valid values are `free`, `trial`, `paygo`, `bluemix-subscription`, and `ibm-subscription`.
	Type *string `json:"type,omitempty"`

	// Defines where the pricing originates.
	Origin *string `json:"origin,omitempty"`

	// Plan-specific starting price information.
	StartingPrice *StartingPrice `json:"starting_price,omitempty"`

	// Plan-specific cost metric structure.
	Metrics []Metrics `json:"metrics,omitempty"`
}


// UnmarshalPricingGet constructs an instance of PricingGet from the specified map.
func UnmarshalPricingGet(m map[string]interface{}) (result *PricingGet, err error) {
	obj := new(PricingGet)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Origin, err = core.UnmarshalString(m, "origin")
	if err != nil {
		return
	}
	obj.StartingPrice, err = UnmarshalStartingPriceAsProperty(m, "starting_price")
	if err != nil {
		return
	}
	obj.Metrics, err = UnmarshalMetricsSliceAsProperty(m, "metrics")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPricingGetSlice unmarshals a slice of PricingGet instances from the specified list of maps.
func UnmarshalPricingGetSlice(s []interface{}) (slice []PricingGet, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PricingGet'")
			return
		}
		obj, e := UnmarshalPricingGet(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPricingGetAsProperty unmarshals an instance of PricingGet that is stored as a property
// within the specified map.
func UnmarshalPricingGetAsProperty(m map[string]interface{}, propertyName string) (result *PricingGet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'PricingGet'", propertyName)
			return
		}
		result, err = UnmarshalPricingGet(objMap)
	}
	return
}

// UnmarshalPricingGetSliceAsProperty unmarshals a slice of PricingGet instances that are stored as a property
// within the specified map.
func UnmarshalPricingGetSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PricingGet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PricingGet'", propertyName)
			return
		}
		slice, err = UnmarshalPricingGetSlice(vSlice)
	}
	return
}

// PricingSet : Pricing-related information.
type PricingSet struct {
	// Type of plan. Valid values are `free`, `trial`, `paygo`, `bluemix-subscription`, and `ibm-subscription`.
	Type *string `json:"type,omitempty"`

	// Defines where the pricing originates.
	Origin *string `json:"origin,omitempty"`

	// Plan-specific starting price information.
	StartingPrice *StartingPrice `json:"starting_price,omitempty"`
}


// UnmarshalPricingSet constructs an instance of PricingSet from the specified map.
func UnmarshalPricingSet(m map[string]interface{}) (result *PricingSet, err error) {
	obj := new(PricingSet)
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.Origin, err = core.UnmarshalString(m, "origin")
	if err != nil {
		return
	}
	obj.StartingPrice, err = UnmarshalStartingPriceAsProperty(m, "starting_price")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalPricingSetSlice unmarshals a slice of PricingSet instances from the specified list of maps.
func UnmarshalPricingSetSlice(s []interface{}) (slice []PricingSet, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'PricingSet'")
			return
		}
		obj, e := UnmarshalPricingSet(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalPricingSetAsProperty unmarshals an instance of PricingSet that is stored as a property
// within the specified map.
func UnmarshalPricingSetAsProperty(m map[string]interface{}, propertyName string) (result *PricingSet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'PricingSet'", propertyName)
			return
		}
		result, err = UnmarshalPricingSet(objMap)
	}
	return
}

// UnmarshalPricingSetSliceAsProperty unmarshals a slice of PricingSet instances that are stored as a property
// within the specified map.
func UnmarshalPricingSetSliceAsProperty(m map[string]interface{}, propertyName string) (slice []PricingSet, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'PricingSet'", propertyName)
			return
		}
		slice, err = UnmarshalPricingSetSlice(vSlice)
	}
	return
}

// Provider : Information related to the provider associated with a catalog entry.
type Provider struct {
	// Provider's email address for this catalog entry.
	Email *string `json:"email" validate:"required"`

	// Provider's name, for example, IBM.
	Name *string `json:"name" validate:"required"`

	// Provider's contact name.
	Contact *string `json:"contact,omitempty"`

	// Provider's support email.
	SupportEmail *string `json:"support_email,omitempty"`

	// Provider's contact phone.
	Phone *string `json:"phone,omitempty"`
}


// NewProvider : Instantiate Provider (Generic Model Constructor)
func (*GlobalCatalogV1) NewProvider(email string, name string) (model *Provider, err error) {
	model = &Provider{
		Email: core.StringPtr(email),
		Name: core.StringPtr(name),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalProvider constructs an instance of Provider from the specified map.
func UnmarshalProvider(m map[string]interface{}) (result *Provider, err error) {
	obj := new(Provider)
	obj.Email, err = core.UnmarshalString(m, "email")
	if err != nil {
		return
	}
	obj.Name, err = core.UnmarshalString(m, "name")
	if err != nil {
		return
	}
	obj.Contact, err = core.UnmarshalString(m, "contact")
	if err != nil {
		return
	}
	obj.SupportEmail, err = core.UnmarshalString(m, "support_email")
	if err != nil {
		return
	}
	obj.Phone, err = core.UnmarshalString(m, "phone")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalProviderSlice unmarshals a slice of Provider instances from the specified list of maps.
func UnmarshalProviderSlice(s []interface{}) (slice []Provider, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Provider'")
			return
		}
		obj, e := UnmarshalProvider(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalProviderAsProperty unmarshals an instance of Provider that is stored as a property
// within the specified map.
func UnmarshalProviderAsProperty(m map[string]interface{}, propertyName string) (result *Provider, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Provider'", propertyName)
			return
		}
		result, err = UnmarshalProvider(objMap)
	}
	return
}

// UnmarshalProviderSliceAsProperty unmarshals a slice of Provider instances that are stored as a property
// within the specified map.
func UnmarshalProviderSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Provider, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Provider'", propertyName)
			return
		}
		slice, err = UnmarshalProviderSlice(vSlice)
	}
	return
}

// RestoreCatalogEntryOptions : The RestoreCatalogEntry options.
type RestoreCatalogEntryOptions struct {
	// The catalog entry's unique ID.
	ID *string `json:"id" validate:"required"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRestoreCatalogEntryOptions : Instantiate RestoreCatalogEntryOptions
func (*GlobalCatalogV1) NewRestoreCatalogEntryOptions(id string) *RestoreCatalogEntryOptions {
	return &RestoreCatalogEntryOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *RestoreCatalogEntryOptions) SetID(id string) *RestoreCatalogEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccount : Allow user to set Account
func (options *RestoreCatalogEntryOptions) SetAccount(account string) *RestoreCatalogEntryOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RestoreCatalogEntryOptions) SetHeaders(param map[string]string) *RestoreCatalogEntryOptions {
	options.Headers = param
	return options
}

// SearchResult : The results obtained by performing a search.
type SearchResult struct {
	// Returned Page Number.
	Page *string `json:"page,omitempty"`

	// Results Per Page – if the page is full.
	ResultsPerPage *string `json:"results_per_page,omitempty"`

	// Total number of results.
	TotalResults *string `json:"total_results,omitempty"`

	// Resulting objects.
	Resources []interface{} `json:"resources,omitempty"`
}


// UnmarshalSearchResult constructs an instance of SearchResult from the specified map.
func UnmarshalSearchResult(m map[string]interface{}) (result *SearchResult, err error) {
	obj := new(SearchResult)
	obj.Page, err = core.UnmarshalString(m, "page")
	if err != nil {
		return
	}
	obj.ResultsPerPage, err = core.UnmarshalString(m, "results_per_page")
	if err != nil {
		return
	}
	obj.TotalResults, err = core.UnmarshalString(m, "total_results")
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

// UnmarshalSearchResultSlice unmarshals a slice of SearchResult instances from the specified list of maps.
func UnmarshalSearchResultSlice(s []interface{}) (slice []SearchResult, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'SearchResult'")
			return
		}
		obj, e := UnmarshalSearchResult(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalSearchResultAsProperty unmarshals an instance of SearchResult that is stored as a property
// within the specified map.
func UnmarshalSearchResultAsProperty(m map[string]interface{}, propertyName string) (result *SearchResult, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'SearchResult'", propertyName)
			return
		}
		result, err = UnmarshalSearchResult(objMap)
	}
	return
}

// UnmarshalSearchResultSliceAsProperty unmarshals a slice of SearchResult instances that are stored as a property
// within the specified map.
func UnmarshalSearchResultSliceAsProperty(m map[string]interface{}, propertyName string) (slice []SearchResult, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'SearchResult'", propertyName)
			return
		}
		slice, err = UnmarshalSearchResultSlice(vSlice)
	}
	return
}

// StartingPrice : Plan-specific starting price information.
type StartingPrice struct {
	// ID of the plan the starting price is calculated.
	PlanID *string `json:"plan_id,omitempty"`

	// ID of the deployment the starting price is calculated.
	DeploymentID *string `json:"deployment_id,omitempty"`

	// The pricing per metric by country and currency.
	Amount []Amount `json:"amount,omitempty"`
}


// UnmarshalStartingPrice constructs an instance of StartingPrice from the specified map.
func UnmarshalStartingPrice(m map[string]interface{}) (result *StartingPrice, err error) {
	obj := new(StartingPrice)
	obj.PlanID, err = core.UnmarshalString(m, "plan_id")
	if err != nil {
		return
	}
	obj.DeploymentID, err = core.UnmarshalString(m, "deployment_id")
	if err != nil {
		return
	}
	obj.Amount, err = UnmarshalAmountSliceAsProperty(m, "amount")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalStartingPriceSlice unmarshals a slice of StartingPrice instances from the specified list of maps.
func UnmarshalStartingPriceSlice(s []interface{}) (slice []StartingPrice, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'StartingPrice'")
			return
		}
		obj, e := UnmarshalStartingPrice(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalStartingPriceAsProperty unmarshals an instance of StartingPrice that is stored as a property
// within the specified map.
func UnmarshalStartingPriceAsProperty(m map[string]interface{}, propertyName string) (result *StartingPrice, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'StartingPrice'", propertyName)
			return
		}
		result, err = UnmarshalStartingPrice(objMap)
	}
	return
}

// UnmarshalStartingPriceSliceAsProperty unmarshals a slice of StartingPrice instances that are stored as a property
// within the specified map.
func UnmarshalStartingPriceSliceAsProperty(m map[string]interface{}, propertyName string) (slice []StartingPrice, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'StartingPrice'", propertyName)
			return
		}
		slice, err = UnmarshalStartingPriceSlice(vSlice)
	}
	return
}

// Strings : Information related to a translated text message.
type Strings struct {
	// Presentation information related to list delimiters.
	Bullets []Bullets `json:"bullets,omitempty"`

	// Media-related metadata.
	Media []UIMetaMedia `json:"media,omitempty"`

	// Warning that a message is not creatable.
	NotCreatableMsg *string `json:"not_creatable_msg,omitempty"`

	// Warning that a robot message is not creatable.
	NotCreatableRobotMsg *string `json:"not_creatable__robot_msg,omitempty"`

	// Warning for deprecation.
	DeprecationWarning *string `json:"deprecation_warning,omitempty"`

	// Popup warning message.
	PopupWarningMessage *string `json:"popup_warning_message,omitempty"`

	// Instructions for UI strings.
	Instruction *string `json:"instruction,omitempty"`
}


// UnmarshalStrings constructs an instance of Strings from the specified map.
func UnmarshalStrings(m map[string]interface{}) (result *Strings, err error) {
	obj := new(Strings)
	obj.Bullets, err = UnmarshalBulletsSliceAsProperty(m, "bullets")
	if err != nil {
		return
	}
	obj.Media, err = UnmarshalUIMetaMediaSliceAsProperty(m, "media")
	if err != nil {
		return
	}
	obj.NotCreatableMsg, err = core.UnmarshalString(m, "not_creatable_msg")
	if err != nil {
		return
	}
	obj.NotCreatableRobotMsg, err = core.UnmarshalString(m, "not_creatable__robot_msg")
	if err != nil {
		return
	}
	obj.DeprecationWarning, err = core.UnmarshalString(m, "deprecation_warning")
	if err != nil {
		return
	}
	obj.PopupWarningMessage, err = core.UnmarshalString(m, "popup_warning_message")
	if err != nil {
		return
	}
	obj.Instruction, err = core.UnmarshalString(m, "instruction")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalStringsSlice unmarshals a slice of Strings instances from the specified list of maps.
func UnmarshalStringsSlice(s []interface{}) (slice []Strings, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Strings'")
			return
		}
		obj, e := UnmarshalStrings(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalStringsAsProperty unmarshals an instance of Strings that is stored as a property
// within the specified map.
func UnmarshalStringsAsProperty(m map[string]interface{}, propertyName string) (result *Strings, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Strings'", propertyName)
			return
		}
		result, err = UnmarshalStrings(objMap)
	}
	return
}

// UnmarshalStringsSliceAsProperty unmarshals a slice of Strings instances that are stored as a property
// within the specified map.
func UnmarshalStringsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []Strings, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Strings'", propertyName)
			return
		}
		slice, err = UnmarshalStringsSlice(vSlice)
	}
	return
}

// UIMetaData : Information related to the UI presentation associated with a catalog entry.
type UIMetaData struct {
	// Language specific translation of translation properties, like label and description.
	Strings *I18N `json:"strings,omitempty"`

	// UI based URLs.
	Urls *URLS `json:"urls,omitempty"`

	// Describes how the embeddable dashboard is rendered.
	EmbeddableDashboard *string `json:"embeddable_dashboard,omitempty"`

	// Describes whether the embeddable dashboard is rendered at the full width.
	EmbeddableDashboardFullWidth *bool `json:"embeddable_dashboard_full_width,omitempty"`

	// Defines the order of information presented.
	NavigationOrder []string `json:"navigation_order,omitempty"`

	// Describes whether this entry is able to be created from the UI element or CLI.
	NotCreatable *bool `json:"not_creatable,omitempty"`

	// Describes whether a plan or flavor is reservable.
	Reservable *bool `json:"reservable,omitempty"`

	// ID of the primary offering for a group.
	PrimaryOfferingID *string `json:"primary_offering_id,omitempty"`

	// Alert to ACE to allow instance UI to be accessible while the provisioning state of instance is in progress.
	AccessibleDuringProvision *bool `json:"accessible_during_provision,omitempty"`

	// Specifies a side by side ordering weight to the UI.
	SideBySideIndex *int64 `json:"side_by_side_index,omitempty"`

	// Date and time the service will no longer be available.
	EndOfServiceTime *strfmt.DateTime `json:"end_of_service_time,omitempty"`
}


// UnmarshalUIMetaData constructs an instance of UIMetaData from the specified map.
func UnmarshalUIMetaData(m map[string]interface{}) (result *UIMetaData, err error) {
	obj := new(UIMetaData)
	obj.Strings, err = UnmarshalI18NAsProperty(m, "strings")
	if err != nil {
		return
	}
	obj.Urls, err = UnmarshalURLSAsProperty(m, "urls")
	if err != nil {
		return
	}
	obj.EmbeddableDashboard, err = core.UnmarshalString(m, "embeddable_dashboard")
	if err != nil {
		return
	}
	obj.EmbeddableDashboardFullWidth, err = core.UnmarshalBool(m, "embeddable_dashboard_full_width")
	if err != nil {
		return
	}
	obj.NavigationOrder, err = core.UnmarshalStringSlice(m, "navigation_order")
	if err != nil {
		return
	}
	obj.NotCreatable, err = core.UnmarshalBool(m, "not_creatable")
	if err != nil {
		return
	}
	obj.Reservable, err = core.UnmarshalBool(m, "reservable")
	if err != nil {
		return
	}
	obj.PrimaryOfferingID, err = core.UnmarshalString(m, "primary_offering_id")
	if err != nil {
		return
	}
	obj.AccessibleDuringProvision, err = core.UnmarshalBool(m, "accessible_during_provision")
	if err != nil {
		return
	}
	obj.SideBySideIndex, err = core.UnmarshalInt64(m, "side_by_side_index")
	if err != nil {
		return
	}
	obj.EndOfServiceTime, err = core.UnmarshalDateTime(m, "end_of_service_time")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUIMetaDataSlice unmarshals a slice of UIMetaData instances from the specified list of maps.
func UnmarshalUIMetaDataSlice(s []interface{}) (slice []UIMetaData, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UIMetaData'")
			return
		}
		obj, e := UnmarshalUIMetaData(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUIMetaDataAsProperty unmarshals an instance of UIMetaData that is stored as a property
// within the specified map.
func UnmarshalUIMetaDataAsProperty(m map[string]interface{}, propertyName string) (result *UIMetaData, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UIMetaData'", propertyName)
			return
		}
		result, err = UnmarshalUIMetaData(objMap)
	}
	return
}

// UnmarshalUIMetaDataSliceAsProperty unmarshals a slice of UIMetaData instances that are stored as a property
// within the specified map.
func UnmarshalUIMetaDataSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UIMetaData, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UIMetaData'", propertyName)
			return
		}
		slice, err = UnmarshalUIMetaDataSlice(vSlice)
	}
	return
}

// UIMetaMedia : Media-related metadata.
type UIMetaMedia struct {
	// Caption for an image.
	Caption *string `json:"caption,omitempty"`

	// URL for thumbnail image.
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`

	// Type of media.
	Type *string `json:"type,omitempty"`

	// URL for media.
	URL *string `json:"URL,omitempty"`

	// Information related to list delimiters.
	Source *Bullets `json:"source,omitempty"`
}


// UnmarshalUIMetaMedia constructs an instance of UIMetaMedia from the specified map.
func UnmarshalUIMetaMedia(m map[string]interface{}) (result *UIMetaMedia, err error) {
	obj := new(UIMetaMedia)
	obj.Caption, err = core.UnmarshalString(m, "caption")
	if err != nil {
		return
	}
	obj.ThumbnailURL, err = core.UnmarshalString(m, "thumbnail_url")
	if err != nil {
		return
	}
	obj.Type, err = core.UnmarshalString(m, "type")
	if err != nil {
		return
	}
	obj.URL, err = core.UnmarshalString(m, "URL")
	if err != nil {
		return
	}
	obj.Source, err = UnmarshalBulletsAsProperty(m, "source")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalUIMetaMediaSlice unmarshals a slice of UIMetaMedia instances from the specified list of maps.
func UnmarshalUIMetaMediaSlice(s []interface{}) (slice []UIMetaMedia, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'UIMetaMedia'")
			return
		}
		obj, e := UnmarshalUIMetaMedia(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalUIMetaMediaAsProperty unmarshals an instance of UIMetaMedia that is stored as a property
// within the specified map.
func UnmarshalUIMetaMediaAsProperty(m map[string]interface{}, propertyName string) (result *UIMetaMedia, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'UIMetaMedia'", propertyName)
			return
		}
		result, err = UnmarshalUIMetaMedia(objMap)
	}
	return
}

// UnmarshalUIMetaMediaSliceAsProperty unmarshals a slice of UIMetaMedia instances that are stored as a property
// within the specified map.
func UnmarshalUIMetaMediaSliceAsProperty(m map[string]interface{}, propertyName string) (slice []UIMetaMedia, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'UIMetaMedia'", propertyName)
			return
		}
		slice, err = UnmarshalUIMetaMediaSlice(vSlice)
	}
	return
}

// URLS : UI based URLs.
type URLS struct {
	// URL for documentation.
	DocURL *string `json:"doc_url,omitempty"`

	// URL for usage instructions.
	InstructionsURL *string `json:"instructions_url,omitempty"`

	// API URL.
	ApiURL *string `json:"api_url,omitempty"`

	// URL Creation UI / API.
	CreateURL *string `json:"create_url,omitempty"`

	// URL to downlaod an SDK.
	SdkDownloadURL *string `json:"sdk_download_url,omitempty"`

	// URL to the terms of use for your service.
	TermsURL *string `json:"terms_url,omitempty"`

	// URL to the custom create page for your serivce.
	CustomCreatePageURL *string `json:"custom_create_page_url,omitempty"`

	// URL to the catalog details page for your serivce.
	CatalogDetailsURL *string `json:"catalog_details_url,omitempty"`

	// URL for deprecation documentation.
	DeprecationDocURL *string `json:"deprecation_doc_url,omitempty"`
}


// UnmarshalURLS constructs an instance of URLS from the specified map.
func UnmarshalURLS(m map[string]interface{}) (result *URLS, err error) {
	obj := new(URLS)
	obj.DocURL, err = core.UnmarshalString(m, "doc_url")
	if err != nil {
		return
	}
	obj.InstructionsURL, err = core.UnmarshalString(m, "instructions_url")
	if err != nil {
		return
	}
	obj.ApiURL, err = core.UnmarshalString(m, "api_url")
	if err != nil {
		return
	}
	obj.CreateURL, err = core.UnmarshalString(m, "create_url")
	if err != nil {
		return
	}
	obj.SdkDownloadURL, err = core.UnmarshalString(m, "sdk_download_url")
	if err != nil {
		return
	}
	obj.TermsURL, err = core.UnmarshalString(m, "terms_url")
	if err != nil {
		return
	}
	obj.CustomCreatePageURL, err = core.UnmarshalString(m, "custom_create_page_url")
	if err != nil {
		return
	}
	obj.CatalogDetailsURL, err = core.UnmarshalString(m, "catalog_details_url")
	if err != nil {
		return
	}
	obj.DeprecationDocURL, err = core.UnmarshalString(m, "deprecation_doc_url")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalURLSSlice unmarshals a slice of URLS instances from the specified list of maps.
func UnmarshalURLSSlice(s []interface{}) (slice []URLS, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'URLS'")
			return
		}
		obj, e := UnmarshalURLS(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalURLSAsProperty unmarshals an instance of URLS that is stored as a property
// within the specified map.
func UnmarshalURLSAsProperty(m map[string]interface{}, propertyName string) (result *URLS, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'URLS'", propertyName)
			return
		}
		result, err = UnmarshalURLS(objMap)
	}
	return
}

// UnmarshalURLSSliceAsProperty unmarshals a slice of URLS instances that are stored as a property
// within the specified map.
func UnmarshalURLSSliceAsProperty(m map[string]interface{}, propertyName string) (slice []URLS, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'URLS'", propertyName)
			return
		}
		slice, err = UnmarshalURLSSlice(vSlice)
	}
	return
}

// UpdateCatalogEntryOptions : The UpdateCatalogEntry options.
type UpdateCatalogEntryOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// Programmatic name for this catalog entry, which must be formatted like a CRN segment. See the display name in
	// OverviewUI for a user-readable name.
	Name *string `json:"name" validate:"required"`

	// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
	// object.
	Kind *string `json:"kind" validate:"required"`

	// Overview is nested in the top level. The key value pair is `[_language_]overview_ui`.
	OverviewUi *OverviewUI `json:"overview_ui" validate:"required"`

	// Image annotation for this catalog entry. The image is a URL.
	Images *Image `json:"images" validate:"required"`

	// Boolean value that determines the global visibility for the catalog entry, and its children. If it is not enabled,
	// all plans are disabled.
	Disabled *bool `json:"disabled" validate:"required"`

	// A list of tags. For example, IBM, 3rd Party, Beta, GA, and Single Tenant.
	Tags []string `json:"tags" validate:"required"`

	// Information related to the provider associated with a catalog entry.
	Provider *Provider `json:"provider" validate:"required"`

	// The ID of the parent catalog entry if it exists.
	ParentID *string `json:"parent_id,omitempty"`

	// Boolean value that determines whether the catalog entry is a group.
	Group *bool `json:"group,omitempty"`

	// Boolean value that describes whether the service is active.
	Active *bool `json:"active,omitempty"`

	// Model used to describe metadata object that can be set.
	Metadata *ObjectMetadataSet `json:"metadata,omitempty"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Reparenting object. In the body set the parent_id to a different parent. Or remove the parent_id field to reparent
	// to the root of the catalog. If this is not set to 'true' then changing the parent_id in the body of the request will
	// not be permitted. If this is 'true' and no change to parent_id then this is also error. This is to prevent
	// accidental changing of parent.
	Move *string `json:"move,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateCatalogEntryOptions.Kind property.
// The type of catalog entry, **service**, **template**, **dashboard**, which determines the type and shape of the
// object.
const (
	UpdateCatalogEntryOptions_Kind_Dashboard = "dashboard"
	UpdateCatalogEntryOptions_Kind_Service = "service"
	UpdateCatalogEntryOptions_Kind_Template = "template"
)

// NewUpdateCatalogEntryOptions : Instantiate UpdateCatalogEntryOptions
func (*GlobalCatalogV1) NewUpdateCatalogEntryOptions(id string, name string, kind string, overviewUi *OverviewUI, images *Image, disabled bool, tags []string, provider *Provider) *UpdateCatalogEntryOptions {
	return &UpdateCatalogEntryOptions{
		ID: core.StringPtr(id),
		Name: core.StringPtr(name),
		Kind: core.StringPtr(kind),
		OverviewUi: overviewUi,
		Images: images,
		Disabled: core.BoolPtr(disabled),
		Tags: tags,
		Provider: provider,
	}
}

// SetID : Allow user to set ID
func (options *UpdateCatalogEntryOptions) SetID(id string) *UpdateCatalogEntryOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateCatalogEntryOptions) SetName(name string) *UpdateCatalogEntryOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetKind : Allow user to set Kind
func (options *UpdateCatalogEntryOptions) SetKind(kind string) *UpdateCatalogEntryOptions {
	options.Kind = core.StringPtr(kind)
	return options
}

// SetOverviewUi : Allow user to set OverviewUi
func (options *UpdateCatalogEntryOptions) SetOverviewUi(overviewUi *OverviewUI) *UpdateCatalogEntryOptions {
	options.OverviewUi = overviewUi
	return options
}

// SetImages : Allow user to set Images
func (options *UpdateCatalogEntryOptions) SetImages(images *Image) *UpdateCatalogEntryOptions {
	options.Images = images
	return options
}

// SetDisabled : Allow user to set Disabled
func (options *UpdateCatalogEntryOptions) SetDisabled(disabled bool) *UpdateCatalogEntryOptions {
	options.Disabled = core.BoolPtr(disabled)
	return options
}

// SetTags : Allow user to set Tags
func (options *UpdateCatalogEntryOptions) SetTags(tags []string) *UpdateCatalogEntryOptions {
	options.Tags = tags
	return options
}

// SetProvider : Allow user to set Provider
func (options *UpdateCatalogEntryOptions) SetProvider(provider *Provider) *UpdateCatalogEntryOptions {
	options.Provider = provider
	return options
}

// SetParentID : Allow user to set ParentID
func (options *UpdateCatalogEntryOptions) SetParentID(parentID string) *UpdateCatalogEntryOptions {
	options.ParentID = core.StringPtr(parentID)
	return options
}

// SetGroup : Allow user to set Group
func (options *UpdateCatalogEntryOptions) SetGroup(group bool) *UpdateCatalogEntryOptions {
	options.Group = core.BoolPtr(group)
	return options
}

// SetActive : Allow user to set Active
func (options *UpdateCatalogEntryOptions) SetActive(active bool) *UpdateCatalogEntryOptions {
	options.Active = core.BoolPtr(active)
	return options
}

// SetMetadata : Allow user to set Metadata
func (options *UpdateCatalogEntryOptions) SetMetadata(metadata *ObjectMetadataSet) *UpdateCatalogEntryOptions {
	options.Metadata = metadata
	return options
}

// SetAccount : Allow user to set Account
func (options *UpdateCatalogEntryOptions) SetAccount(account string) *UpdateCatalogEntryOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetMove : Allow user to set Move
func (options *UpdateCatalogEntryOptions) SetMove(move string) *UpdateCatalogEntryOptions {
	options.Move = core.StringPtr(move)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCatalogEntryOptions) SetHeaders(param map[string]string) *UpdateCatalogEntryOptions {
	options.Headers = param
	return options
}

// UpdateVisibilityOptions : The UpdateVisibility options.
type UpdateVisibilityOptions struct {
	// The object's unique ID.
	ID *string `json:"id" validate:"required"`

	// Visibility details related to a catalog entry.
	Include *VisibilityDetail `json:"include,omitempty"`

	// Visibility details related to a catalog entry.
	Exclude *VisibilityDetail `json:"exclude,omitempty"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateVisibilityOptions : Instantiate UpdateVisibilityOptions
func (*GlobalCatalogV1) NewUpdateVisibilityOptions(id string) *UpdateVisibilityOptions {
	return &UpdateVisibilityOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UpdateVisibilityOptions) SetID(id string) *UpdateVisibilityOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetInclude : Allow user to set Include
func (options *UpdateVisibilityOptions) SetInclude(include *VisibilityDetail) *UpdateVisibilityOptions {
	options.Include = include
	return options
}

// SetExclude : Allow user to set Exclude
func (options *UpdateVisibilityOptions) SetExclude(exclude *VisibilityDetail) *UpdateVisibilityOptions {
	options.Exclude = exclude
	return options
}

// SetAccount : Allow user to set Account
func (options *UpdateVisibilityOptions) SetAccount(account string) *UpdateVisibilityOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateVisibilityOptions) SetHeaders(param map[string]string) *UpdateVisibilityOptions {
	options.Headers = param
	return options
}

// UploadArtifactOptions : The UploadArtifact options.
type UploadArtifactOptions struct {
	// The object's unique ID.
	ObjectID *string `json:"object_id" validate:"required"`

	// The artifact's ID.
	ArtifactID *string `json:"artifact_id" validate:"required"`

	Artifact io.ReadCloser `json:"artifact,omitempty"`

	// The type of the input.
	ContentType *string `json:"Content-Type,omitempty"`

	// This changes the scope of the request regardless of the authorization header. Example scopes are `account` and
	// `global`. `account=global` is reqired if operating with a service ID that has a global admin policy, for example
	// `GET /?account=global`.
	Account *string `json:"account,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUploadArtifactOptions : Instantiate UploadArtifactOptions
func (*GlobalCatalogV1) NewUploadArtifactOptions(objectID string, artifactID string) *UploadArtifactOptions {
	return &UploadArtifactOptions{
		ObjectID: core.StringPtr(objectID),
		ArtifactID: core.StringPtr(artifactID),
	}
}

// SetObjectID : Allow user to set ObjectID
func (options *UploadArtifactOptions) SetObjectID(objectID string) *UploadArtifactOptions {
	options.ObjectID = core.StringPtr(objectID)
	return options
}

// SetArtifactID : Allow user to set ArtifactID
func (options *UploadArtifactOptions) SetArtifactID(artifactID string) *UploadArtifactOptions {
	options.ArtifactID = core.StringPtr(artifactID)
	return options
}

// SetArtifact : Allow user to set Artifact
func (options *UploadArtifactOptions) SetArtifact(artifact io.ReadCloser) *UploadArtifactOptions {
	options.Artifact = artifact
	return options
}

// SetContentType : Allow user to set ContentType
func (options *UploadArtifactOptions) SetContentType(contentType string) *UploadArtifactOptions {
	options.ContentType = core.StringPtr(contentType)
	return options
}

// SetAccount : Allow user to set Account
func (options *UploadArtifactOptions) SetAccount(account string) *UploadArtifactOptions {
	options.Account = core.StringPtr(account)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UploadArtifactOptions) SetHeaders(param map[string]string) *UploadArtifactOptions {
	options.Headers = param
	return options
}

// Visibility : Information related to the visibility of a catalog entry.
type Visibility struct {
	// This controls the overall visibility. It is an enum of *public*, *ibm_only*, and *private*. public means it is
	// visible to all. ibm_only means it is visible to all IBM unless their account is explicitly excluded. private means
	// it is visible only to the included accounts.
	Restrictions *string `json:"restrictions,omitempty"`

	// IAM Scope-related information associated with a catalog entry.
	Owner *string `json:"owner,omitempty"`

	// Visibility details related to a catalog entry.
	Include *VisibilityDetail `json:"include,omitempty"`

	// Visibility details related to a catalog entry.
	Exclude *VisibilityDetail `json:"exclude,omitempty"`

	// Determines whether the owning account has full control over the visibility of the entry such as adding non-IBM
	// accounts to the whitelist and making entries `private`, `ibm_only` or `public`.
	Approved *bool `json:"approved,omitempty"`
}


// UnmarshalVisibility constructs an instance of Visibility from the specified map.
func UnmarshalVisibility(m map[string]interface{}) (result *Visibility, err error) {
	obj := new(Visibility)
	obj.Restrictions, err = core.UnmarshalString(m, "restrictions")
	if err != nil {
		return
	}
	obj.Owner, err = core.UnmarshalString(m, "owner")
	if err != nil {
		return
	}
	obj.Include, err = UnmarshalVisibilityDetailAsProperty(m, "include")
	if err != nil {
		return
	}
	obj.Exclude, err = UnmarshalVisibilityDetailAsProperty(m, "exclude")
	if err != nil {
		return
	}
	obj.Approved, err = core.UnmarshalBool(m, "approved")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalVisibilitySlice unmarshals a slice of Visibility instances from the specified list of maps.
func UnmarshalVisibilitySlice(s []interface{}) (slice []Visibility, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'Visibility'")
			return
		}
		obj, e := UnmarshalVisibility(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalVisibilityAsProperty unmarshals an instance of Visibility that is stored as a property
// within the specified map.
func UnmarshalVisibilityAsProperty(m map[string]interface{}, propertyName string) (result *Visibility, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'Visibility'", propertyName)
			return
		}
		result, err = UnmarshalVisibility(objMap)
	}
	return
}

// UnmarshalVisibilitySliceAsProperty unmarshals a slice of Visibility instances that are stored as a property
// within the specified map.
func UnmarshalVisibilitySliceAsProperty(m map[string]interface{}, propertyName string) (slice []Visibility, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'Visibility'", propertyName)
			return
		}
		slice, err = UnmarshalVisibilitySlice(vSlice)
	}
	return
}

// VisibilityDetail : Visibility details related to a catalog entry.
type VisibilityDetail struct {
	// Information related to the accounts for which a catalog entry is visible.
	Accounts *VisibilityDetailAccounts `json:"accounts" validate:"required"`
}


// NewVisibilityDetail : Instantiate VisibilityDetail (Generic Model Constructor)
func (*GlobalCatalogV1) NewVisibilityDetail(accounts *VisibilityDetailAccounts) (model *VisibilityDetail, err error) {
	model = &VisibilityDetail{
		Accounts: accounts,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalVisibilityDetail constructs an instance of VisibilityDetail from the specified map.
func UnmarshalVisibilityDetail(m map[string]interface{}) (result *VisibilityDetail, err error) {
	obj := new(VisibilityDetail)
	obj.Accounts, err = UnmarshalVisibilityDetailAccountsAsProperty(m, "accounts")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalVisibilityDetailSlice unmarshals a slice of VisibilityDetail instances from the specified list of maps.
func UnmarshalVisibilityDetailSlice(s []interface{}) (slice []VisibilityDetail, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'VisibilityDetail'")
			return
		}
		obj, e := UnmarshalVisibilityDetail(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalVisibilityDetailAsProperty unmarshals an instance of VisibilityDetail that is stored as a property
// within the specified map.
func UnmarshalVisibilityDetailAsProperty(m map[string]interface{}, propertyName string) (result *VisibilityDetail, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'VisibilityDetail'", propertyName)
			return
		}
		result, err = UnmarshalVisibilityDetail(objMap)
	}
	return
}

// UnmarshalVisibilityDetailSliceAsProperty unmarshals a slice of VisibilityDetail instances that are stored as a property
// within the specified map.
func UnmarshalVisibilityDetailSliceAsProperty(m map[string]interface{}, propertyName string) (slice []VisibilityDetail, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'VisibilityDetail'", propertyName)
			return
		}
		slice, err = UnmarshalVisibilityDetailSlice(vSlice)
	}
	return
}

// VisibilityDetailAccounts : Information related to the accounts for which a catalog entry is visible.
type VisibilityDetailAccounts struct {
	// (_accountid_) is the GUID of the account and the value is the scope of who set it. For setting visibility use "" as
	// the value. It is replaced with the owner scope when saved.
	Accountid *string `json:"_accountid_,omitempty"`
}


// UnmarshalVisibilityDetailAccounts constructs an instance of VisibilityDetailAccounts from the specified map.
func UnmarshalVisibilityDetailAccounts(m map[string]interface{}) (result *VisibilityDetailAccounts, err error) {
	obj := new(VisibilityDetailAccounts)
	obj.Accountid, err = core.UnmarshalString(m, "_accountid_")
	if err != nil {
		return
	}
	result = obj
	return
}

// UnmarshalVisibilityDetailAccountsSlice unmarshals a slice of VisibilityDetailAccounts instances from the specified list of maps.
func UnmarshalVisibilityDetailAccountsSlice(s []interface{}) (slice []VisibilityDetailAccounts, err error) {
	for _, v := range s {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("slice element should be a map containing an instance of 'VisibilityDetailAccounts'")
			return
		}
		obj, e := UnmarshalVisibilityDetailAccounts(objMap)
		if e != nil {
			err = e
			return
		}
		slice = append(slice, *obj)
	}
	return
}

// UnmarshalVisibilityDetailAccountsAsProperty unmarshals an instance of VisibilityDetailAccounts that is stored as a property
// within the specified map.
func UnmarshalVisibilityDetailAccountsAsProperty(m map[string]interface{}, propertyName string) (result *VisibilityDetailAccounts, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		objMap, ok := v.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a map containing an instance of 'VisibilityDetailAccounts'", propertyName)
			return
		}
		result, err = UnmarshalVisibilityDetailAccounts(objMap)
	}
	return
}

// UnmarshalVisibilityDetailAccountsSliceAsProperty unmarshals a slice of VisibilityDetailAccounts instances that are stored as a property
// within the specified map.
func UnmarshalVisibilityDetailAccountsSliceAsProperty(m map[string]interface{}, propertyName string) (slice []VisibilityDetailAccounts, err error) {
	v, foundIt := m[propertyName]
	if foundIt && v != nil {
		vSlice, ok := v.([]interface{})
		if !ok {
			err = fmt.Errorf("map property '%s' should be a slice of maps, each containing an instance of 'VisibilityDetailAccounts'", propertyName)
			return
		}
		slice, err = UnmarshalVisibilityDetailAccountsSlice(vSlice)
	}
	return
}
