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

// Package catalogmanagementv1 : Operations and models for the CatalogManagementV1 service
package catalogmanagementv1

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/go-openapi/strfmt"
	"reflect"
)

// CatalogManagementV1 : This is the API to use for managing private catalogs for IBM Cloud. Private catalogs provide a
// way to centrally manage access to products in the IBM Cloud catalog and your own catalogs.
//
// Version: 1.0
type CatalogManagementV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://cm.globalcatalog.cloud.ibm.com/api/v1-beta"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "catalog_management"

// CatalogManagementV1Options : Service options
type CatalogManagementV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewCatalogManagementV1UsingExternalConfig : constructs an instance of CatalogManagementV1 with passed in options and external configuration.
func NewCatalogManagementV1UsingExternalConfig(options *CatalogManagementV1Options) (catalogManagement *CatalogManagementV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	catalogManagement, err = NewCatalogManagementV1(options)
	if err != nil {
		return
	}

	err = catalogManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = catalogManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewCatalogManagementV1 : constructs an instance of CatalogManagementV1 with passed in options.
func NewCatalogManagementV1(options *CatalogManagementV1Options) (service *CatalogManagementV1, err error) {
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

	service = &CatalogManagementV1{
		Service: baseService,
	}

	return
}

// SetServiceURL sets the service URL
func (catalogManagement *CatalogManagementV1) SetServiceURL(url string) error {
	return catalogManagement.Service.SetServiceURL(url)
}

// GetCatalogAccount : Get the account settings
// Get the account level settings for the account for private catalog.
func (catalogManagement *CatalogManagementV1) GetCatalogAccount(getCatalogAccountOptions *GetCatalogAccountOptions) (result *Account, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCatalogAccountOptions, "getCatalogAccountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogaccount"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetCatalogAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccount)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateCatalogAccount : Set the account settings
func (catalogManagement *CatalogManagementV1) UpdateCatalogAccount(updateCatalogAccountOptions *UpdateCatalogAccountOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateCatalogAccountOptions, "updateCatalogAccountOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogaccount"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCatalogAccountOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "UpdateCatalogAccount")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCatalogAccountOptions.ID != nil {
		body["id"] = updateCatalogAccountOptions.ID
	}
	if updateCatalogAccountOptions.AccountFilters != nil {
		body["account_filters"] = updateCatalogAccountOptions.AccountFilters
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetCatalogAccountFilters : Get the accumulated filters of the account and of the catalogs you have access to
// Get the accumulated filters of the account and of the catalogs you have access to.
func (catalogManagement *CatalogManagementV1) GetCatalogAccountFilters(getCatalogAccountFiltersOptions *GetCatalogAccountFiltersOptions) (result *AccumulatedFilters, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCatalogAccountFiltersOptions, "getCatalogAccountFiltersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogaccount/filters"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogAccountFiltersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetCatalogAccountFilters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getCatalogAccountFiltersOptions.Catalog != nil {
		builder.AddQuery("catalog", fmt.Sprint(*getCatalogAccountFiltersOptions.Catalog))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAccumulatedFilters)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListCatalogs : Get list of catalogs
// List the available catalogs for a given account.
func (catalogManagement *CatalogManagementV1) ListCatalogs(listCatalogsOptions *ListCatalogsOptions) (result *CatalogSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCatalogsOptions, "listCatalogsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCatalogsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListCatalogs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalogSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateCatalog : Create a catalog
// Create a catalog for a given account.
func (catalogManagement *CatalogManagementV1) CreateCatalog(createCatalogOptions *CreateCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createCatalogOptions, "createCatalogOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CreateCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCatalogOptions.ID != nil {
		body["id"] = createCatalogOptions.ID
	}
	if createCatalogOptions.Label != nil {
		body["label"] = createCatalogOptions.Label
	}
	if createCatalogOptions.ShortDescription != nil {
		body["short_description"] = createCatalogOptions.ShortDescription
	}
	if createCatalogOptions.CatalogIconURL != nil {
		body["catalog_icon_url"] = createCatalogOptions.CatalogIconURL
	}
	if createCatalogOptions.Tags != nil {
		body["tags"] = createCatalogOptions.Tags
	}
	if createCatalogOptions.URL != nil {
		body["url"] = createCatalogOptions.URL
	}
	if createCatalogOptions.Crn != nil {
		body["crn"] = createCatalogOptions.Crn
	}
	if createCatalogOptions.OfferingsURL != nil {
		body["offerings_url"] = createCatalogOptions.OfferingsURL
	}
	if createCatalogOptions.Features != nil {
		body["features"] = createCatalogOptions.Features
	}
	if createCatalogOptions.Disabled != nil {
		body["disabled"] = createCatalogOptions.Disabled
	}
	if createCatalogOptions.Created != nil {
		body["created"] = createCatalogOptions.Created
	}
	if createCatalogOptions.Updated != nil {
		body["updated"] = createCatalogOptions.Updated
	}
	if createCatalogOptions.ResourceGroupID != nil {
		body["resource_group_id"] = createCatalogOptions.ResourceGroupID
	}
	if createCatalogOptions.OwningAccount != nil {
		body["owning_account"] = createCatalogOptions.OwningAccount
	}
	if createCatalogOptions.CatalogFilters != nil {
		body["catalog_filters"] = createCatalogOptions.CatalogFilters
	}
	if createCatalogOptions.SyndicationSettings != nil {
		body["syndication_settings"] = createCatalogOptions.SyndicationSettings
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCatalog : Get a catalog
// Get a catalog.
func (catalogManagement *CatalogManagementV1) GetCatalog(getCatalogOptions *GetCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCatalogOptions, "getCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCatalogOptions, "getCatalogOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs"}
	pathParameters := []string{*getCatalogOptions.CatalogID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceCatalog : Update a catalog
// Update a catalog.
func (catalogManagement *CatalogManagementV1) ReplaceCatalog(replaceCatalogOptions *ReplaceCatalogOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceCatalogOptions, "replaceCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceCatalogOptions, "replaceCatalogOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs"}
	pathParameters := []string{*replaceCatalogOptions.CatalogID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReplaceCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceCatalogOptions.ID != nil {
		body["id"] = replaceCatalogOptions.ID
	}
	if replaceCatalogOptions.Label != nil {
		body["label"] = replaceCatalogOptions.Label
	}
	if replaceCatalogOptions.ShortDescription != nil {
		body["short_description"] = replaceCatalogOptions.ShortDescription
	}
	if replaceCatalogOptions.CatalogIconURL != nil {
		body["catalog_icon_url"] = replaceCatalogOptions.CatalogIconURL
	}
	if replaceCatalogOptions.Tags != nil {
		body["tags"] = replaceCatalogOptions.Tags
	}
	if replaceCatalogOptions.URL != nil {
		body["url"] = replaceCatalogOptions.URL
	}
	if replaceCatalogOptions.Crn != nil {
		body["crn"] = replaceCatalogOptions.Crn
	}
	if replaceCatalogOptions.OfferingsURL != nil {
		body["offerings_url"] = replaceCatalogOptions.OfferingsURL
	}
	if replaceCatalogOptions.Features != nil {
		body["features"] = replaceCatalogOptions.Features
	}
	if replaceCatalogOptions.Disabled != nil {
		body["disabled"] = replaceCatalogOptions.Disabled
	}
	if replaceCatalogOptions.Created != nil {
		body["created"] = replaceCatalogOptions.Created
	}
	if replaceCatalogOptions.Updated != nil {
		body["updated"] = replaceCatalogOptions.Updated
	}
	if replaceCatalogOptions.ResourceGroupID != nil {
		body["resource_group_id"] = replaceCatalogOptions.ResourceGroupID
	}
	if replaceCatalogOptions.OwningAccount != nil {
		body["owning_account"] = replaceCatalogOptions.OwningAccount
	}
	if replaceCatalogOptions.CatalogFilters != nil {
		body["catalog_filters"] = replaceCatalogOptions.CatalogFilters
	}
	if replaceCatalogOptions.SyndicationSettings != nil {
		body["syndication_settings"] = replaceCatalogOptions.SyndicationSettings
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteCatalog : Delete a catalog
// Delete a catalog.
func (catalogManagement *CatalogManagementV1) DeleteCatalog(deleteCatalogOptions *DeleteCatalogOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCatalogOptions, "deleteCatalogOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCatalogOptions, "deleteCatalogOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs"}
	pathParameters := []string{*deleteCatalogOptions.CatalogID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCatalogOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeleteCatalog")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetEnterprise : Get the enterprise settings for the specified enterprise ID
// Get the enterprise settings for the specified enterprise ID.
func (catalogManagement *CatalogManagementV1) GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions) (result *Enterprise, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getEnterpriseOptions, "getEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getEnterpriseOptions, "getEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{*getEnterpriseOptions.EnterpriseID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEnterprise)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceEnterprise : Set the enterprise settings
func (catalogManagement *CatalogManagementV1) ReplaceEnterprise(replaceEnterpriseOptions *ReplaceEnterpriseOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceEnterpriseOptions, "replaceEnterpriseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceEnterpriseOptions, "replaceEnterpriseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"enterprises"}
	pathParameters := []string{*replaceEnterpriseOptions.EnterpriseID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceEnterpriseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReplaceEnterprise")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceEnterpriseOptions.ID != nil {
		body["id"] = replaceEnterpriseOptions.ID
	}
	if replaceEnterpriseOptions.AccountFilters != nil {
		body["account_filters"] = replaceEnterpriseOptions.AccountFilters
	}
	if replaceEnterpriseOptions.AccountGroups != nil {
		body["account_groups"] = replaceEnterpriseOptions.AccountGroups
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetConsumptionOfferings : Get list of offerings for consumption
// List the available offerings from both public and from the account that currently scoped for consumption. These
// copies cannot be used updating. They are not complete and only return what is visible to the caller.
func (catalogManagement *CatalogManagementV1) GetConsumptionOfferings(getConsumptionOfferingsOptions *GetConsumptionOfferingsOptions) (result *OfferingSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getConsumptionOfferingsOptions, "getConsumptionOfferingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"offerings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getConsumptionOfferingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetConsumptionOfferings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getConsumptionOfferingsOptions.Digest != nil {
		builder.AddQuery("digest", fmt.Sprint(*getConsumptionOfferingsOptions.Digest))
	}
	if getConsumptionOfferingsOptions.Catalog != nil {
		builder.AddQuery("catalog", fmt.Sprint(*getConsumptionOfferingsOptions.Catalog))
	}
	if getConsumptionOfferingsOptions.Select != nil {
		builder.AddQuery("select", fmt.Sprint(*getConsumptionOfferingsOptions.Select))
	}
	if getConsumptionOfferingsOptions.IncludeHidden != nil {
		builder.AddQuery("includeHidden", fmt.Sprint(*getConsumptionOfferingsOptions.IncludeHidden))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOfferingSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListOfferings : Get list of offerings
// List the available offerings in the specified catalog.
func (catalogManagement *CatalogManagementV1) ListOfferings(listOfferingsOptions *ListOfferingsOptions) (result *OfferingSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listOfferingsOptions, "listOfferingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listOfferingsOptions, "listOfferingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings"}
	pathParameters := []string{*listOfferingsOptions.CatalogID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listOfferingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListOfferings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listOfferingsOptions.Digest != nil {
		builder.AddQuery("digest", fmt.Sprint(*listOfferingsOptions.Digest))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOfferingSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateOffering : Create an offering
// Create an offering.
func (catalogManagement *CatalogManagementV1) CreateOffering(createOfferingOptions *CreateOfferingOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createOfferingOptions, "createOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createOfferingOptions, "createOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings"}
	pathParameters := []string{*createOfferingOptions.CatalogID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CreateOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createOfferingOptions.NewID != nil {
		body["id"] = createOfferingOptions.NewID
	}
	if createOfferingOptions.NewURL != nil {
		body["url"] = createOfferingOptions.NewURL
	}
	if createOfferingOptions.NewCrn != nil {
		body["crn"] = createOfferingOptions.NewCrn
	}
	if createOfferingOptions.NewLabel != nil {
		body["label"] = createOfferingOptions.NewLabel
	}
	if createOfferingOptions.NewName != nil {
		body["name"] = createOfferingOptions.NewName
	}
	if createOfferingOptions.NewOfferingIconURL != nil {
		body["offering_icon_url"] = createOfferingOptions.NewOfferingIconURL
	}
	if createOfferingOptions.NewOfferingDocsURL != nil {
		body["offering_docs_url"] = createOfferingOptions.NewOfferingDocsURL
	}
	if createOfferingOptions.NewOfferingSupportURL != nil {
		body["offering_support_url"] = createOfferingOptions.NewOfferingSupportURL
	}
	if createOfferingOptions.NewTags != nil {
		body["tags"] = createOfferingOptions.NewTags
	}
	if createOfferingOptions.NewRating != nil {
		body["rating"] = createOfferingOptions.NewRating
	}
	if createOfferingOptions.NewCreated != nil {
		body["created"] = createOfferingOptions.NewCreated
	}
	if createOfferingOptions.NewUpdated != nil {
		body["updated"] = createOfferingOptions.NewUpdated
	}
	if createOfferingOptions.NewShortDescription != nil {
		body["short_description"] = createOfferingOptions.NewShortDescription
	}
	if createOfferingOptions.NewLongDescription != nil {
		body["long_description"] = createOfferingOptions.NewLongDescription
	}
	if createOfferingOptions.NewFeatures != nil {
		body["features"] = createOfferingOptions.NewFeatures
	}
	if createOfferingOptions.NewKinds != nil {
		body["kinds"] = createOfferingOptions.NewKinds
	}
	if createOfferingOptions.NewPermitRequestIbmPublicPublish != nil {
		body["permit_request_ibm_public_publish"] = createOfferingOptions.NewPermitRequestIbmPublicPublish
	}
	if createOfferingOptions.NewIbmPublishApproved != nil {
		body["ibm_publish_approved"] = createOfferingOptions.NewIbmPublishApproved
	}
	if createOfferingOptions.NewPublicPublishApproved != nil {
		body["public_publish_approved"] = createOfferingOptions.NewPublicPublishApproved
	}
	if createOfferingOptions.NewPublicOriginalCrn != nil {
		body["public_original_crn"] = createOfferingOptions.NewPublicOriginalCrn
	}
	if createOfferingOptions.NewPublishPublicCrn != nil {
		body["publish_public_crn"] = createOfferingOptions.NewPublishPublicCrn
	}
	if createOfferingOptions.NewPortalApprovalRecord != nil {
		body["portal_approval_record"] = createOfferingOptions.NewPortalApprovalRecord
	}
	if createOfferingOptions.NewPortalUiURL != nil {
		body["portal_ui_url"] = createOfferingOptions.NewPortalUiURL
	}
	if createOfferingOptions.NewCatalogID != nil {
		body["catalog_id"] = createOfferingOptions.NewCatalogID
	}
	if createOfferingOptions.NewCatalogName != nil {
		body["catalog_name"] = createOfferingOptions.NewCatalogName
	}
	if createOfferingOptions.NewMetadata != nil {
		body["metadata"] = createOfferingOptions.NewMetadata
	}
	if createOfferingOptions.NewDisclaimer != nil {
		body["disclaimer"] = createOfferingOptions.NewDisclaimer
	}
	if createOfferingOptions.NewHidden != nil {
		body["hidden"] = createOfferingOptions.NewHidden
	}
	if createOfferingOptions.NewProvider != nil {
		body["provider"] = createOfferingOptions.NewProvider
	}
	if createOfferingOptions.NewRepoInfo != nil {
		body["repo_info"] = createOfferingOptions.NewRepoInfo
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportOfferingVersion : Import new version to offering from a tgz
// Import new version to offering from a tgz.
func (catalogManagement *CatalogManagementV1) ImportOfferingVersion(importOfferingVersionOptions *ImportOfferingVersionOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importOfferingVersionOptions, "importOfferingVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importOfferingVersionOptions, "importOfferingVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings", "version"}
	pathParameters := []string{*importOfferingVersionOptions.CatalogID, *importOfferingVersionOptions.OfferingID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importOfferingVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ImportOfferingVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("zipurl", fmt.Sprint(*importOfferingVersionOptions.Zipurl))
	if importOfferingVersionOptions.TargetVersion != nil {
		builder.AddQuery("targetVersion", fmt.Sprint(*importOfferingVersionOptions.TargetVersion))
	}
	if importOfferingVersionOptions.IncludeConfig != nil {
		builder.AddQuery("includeConfig", fmt.Sprint(*importOfferingVersionOptions.IncludeConfig))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ImportOffering : Import a new offering from a tgz
// Import a new offering from a tgz.
func (catalogManagement *CatalogManagementV1) ImportOffering(importOfferingOptions *ImportOfferingOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importOfferingOptions, "importOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importOfferingOptions, "importOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "import/offerings"}
	pathParameters := []string{*importOfferingOptions.CatalogID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range importOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ImportOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("zipurl", fmt.Sprint(*importOfferingOptions.Zipurl))
	if importOfferingOptions.OfferingID != nil {
		builder.AddQuery("offeringID", fmt.Sprint(*importOfferingOptions.OfferingID))
	}
	if importOfferingOptions.IncludeConfig != nil {
		builder.AddQuery("includeConfig", fmt.Sprint(*importOfferingOptions.IncludeConfig))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReloadOffering : Reload existing version in offering from a tgz
// Reload existing version in offering from a tgz.
func (catalogManagement *CatalogManagementV1) ReloadOffering(reloadOfferingOptions *ReloadOfferingOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(reloadOfferingOptions, "reloadOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(reloadOfferingOptions, "reloadOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings", "reload"}
	pathParameters := []string{*reloadOfferingOptions.CatalogID, *reloadOfferingOptions.OfferingID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range reloadOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReloadOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("zipurl", fmt.Sprint(*reloadOfferingOptions.Zipurl))
	builder.AddQuery("targetVersion", fmt.Sprint(*reloadOfferingOptions.TargetVersion))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetOffering : Get an offering
// Get an offering.
func (catalogManagement *CatalogManagementV1) GetOffering(getOfferingOptions *GetOfferingOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getOfferingOptions, "getOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getOfferingOptions, "getOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings"}
	pathParameters := []string{*getOfferingOptions.CatalogID, *getOfferingOptions.OfferingID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceOffering : Update an offering
// Update an offering.
func (catalogManagement *CatalogManagementV1) ReplaceOffering(replaceOfferingOptions *ReplaceOfferingOptions) (result *Catalog, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceOfferingOptions, "replaceOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceOfferingOptions, "replaceOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings"}
	pathParameters := []string{*replaceOfferingOptions.CatalogID, *replaceOfferingOptions.OfferingID}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReplaceOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if replaceOfferingOptions.NewID != nil {
		body["id"] = replaceOfferingOptions.NewID
	}
	if replaceOfferingOptions.NewURL != nil {
		body["url"] = replaceOfferingOptions.NewURL
	}
	if replaceOfferingOptions.NewCrn != nil {
		body["crn"] = replaceOfferingOptions.NewCrn
	}
	if replaceOfferingOptions.NewLabel != nil {
		body["label"] = replaceOfferingOptions.NewLabel
	}
	if replaceOfferingOptions.NewName != nil {
		body["name"] = replaceOfferingOptions.NewName
	}
	if replaceOfferingOptions.NewOfferingIconURL != nil {
		body["offering_icon_url"] = replaceOfferingOptions.NewOfferingIconURL
	}
	if replaceOfferingOptions.NewOfferingDocsURL != nil {
		body["offering_docs_url"] = replaceOfferingOptions.NewOfferingDocsURL
	}
	if replaceOfferingOptions.NewOfferingSupportURL != nil {
		body["offering_support_url"] = replaceOfferingOptions.NewOfferingSupportURL
	}
	if replaceOfferingOptions.NewTags != nil {
		body["tags"] = replaceOfferingOptions.NewTags
	}
	if replaceOfferingOptions.NewRating != nil {
		body["rating"] = replaceOfferingOptions.NewRating
	}
	if replaceOfferingOptions.NewCreated != nil {
		body["created"] = replaceOfferingOptions.NewCreated
	}
	if replaceOfferingOptions.NewUpdated != nil {
		body["updated"] = replaceOfferingOptions.NewUpdated
	}
	if replaceOfferingOptions.NewShortDescription != nil {
		body["short_description"] = replaceOfferingOptions.NewShortDescription
	}
	if replaceOfferingOptions.NewLongDescription != nil {
		body["long_description"] = replaceOfferingOptions.NewLongDescription
	}
	if replaceOfferingOptions.NewFeatures != nil {
		body["features"] = replaceOfferingOptions.NewFeatures
	}
	if replaceOfferingOptions.NewKinds != nil {
		body["kinds"] = replaceOfferingOptions.NewKinds
	}
	if replaceOfferingOptions.NewPermitRequestIbmPublicPublish != nil {
		body["permit_request_ibm_public_publish"] = replaceOfferingOptions.NewPermitRequestIbmPublicPublish
	}
	if replaceOfferingOptions.NewIbmPublishApproved != nil {
		body["ibm_publish_approved"] = replaceOfferingOptions.NewIbmPublishApproved
	}
	if replaceOfferingOptions.NewPublicPublishApproved != nil {
		body["public_publish_approved"] = replaceOfferingOptions.NewPublicPublishApproved
	}
	if replaceOfferingOptions.NewPublicOriginalCrn != nil {
		body["public_original_crn"] = replaceOfferingOptions.NewPublicOriginalCrn
	}
	if replaceOfferingOptions.NewPublishPublicCrn != nil {
		body["publish_public_crn"] = replaceOfferingOptions.NewPublishPublicCrn
	}
	if replaceOfferingOptions.NewPortalApprovalRecord != nil {
		body["portal_approval_record"] = replaceOfferingOptions.NewPortalApprovalRecord
	}
	if replaceOfferingOptions.NewPortalUiURL != nil {
		body["portal_ui_url"] = replaceOfferingOptions.NewPortalUiURL
	}
	if replaceOfferingOptions.NewCatalogID != nil {
		body["catalog_id"] = replaceOfferingOptions.NewCatalogID
	}
	if replaceOfferingOptions.NewCatalogName != nil {
		body["catalog_name"] = replaceOfferingOptions.NewCatalogName
	}
	if replaceOfferingOptions.NewMetadata != nil {
		body["metadata"] = replaceOfferingOptions.NewMetadata
	}
	if replaceOfferingOptions.NewDisclaimer != nil {
		body["disclaimer"] = replaceOfferingOptions.NewDisclaimer
	}
	if replaceOfferingOptions.NewHidden != nil {
		body["hidden"] = replaceOfferingOptions.NewHidden
	}
	if replaceOfferingOptions.NewProvider != nil {
		body["provider"] = replaceOfferingOptions.NewProvider
	}
	if replaceOfferingOptions.NewRepoInfo != nil {
		body["repo_info"] = replaceOfferingOptions.NewRepoInfo
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCatalog)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteOffering : Delete an offering
// Delete an offering.
func (catalogManagement *CatalogManagementV1) DeleteOffering(deleteOfferingOptions *DeleteOfferingOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteOfferingOptions, "deleteOfferingOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteOfferingOptions, "deleteOfferingOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings"}
	pathParameters := []string{*deleteOfferingOptions.CatalogID, *deleteOfferingOptions.OfferingID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteOfferingOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeleteOffering")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// ReplaceOfferingIcon : upload an icon for the offering
// upload an icon file to be stored in GC. File is uploaded as a binary payload - not as a form.
func (catalogManagement *CatalogManagementV1) ReplaceOfferingIcon(replaceOfferingIconOptions *ReplaceOfferingIconOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceOfferingIconOptions, "replaceOfferingIconOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceOfferingIconOptions, "replaceOfferingIconOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings", "icon"}
	pathParameters := []string{*replaceOfferingIconOptions.CatalogID, *replaceOfferingIconOptions.OfferingID, *replaceOfferingIconOptions.FileName}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceOfferingIconOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReplaceOfferingIcon")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateOfferingIbm : Approve offering to be permitted to publish to IBM Public Catalog (IBMers only or Everyone)
// Approve or disapprove the offering to be allowed to publish to the IBM Public Catalog in `ibm` (visible to IBM only)
// or `public` (visible to everyone). Can approve to only `ibm`, or it can be extended to `public`. If extended to
// `public` then `ibm` is automatically approved too. If disapprove `public`, then `ibm` approval will not  be changed.
// If disapprove `ibm` then `public` will automatically be disapproved. This is because the process steps always go
// first through `ibm` and then to `public`. `ibm` cannot be skipped. Only users with Approval IAM authority can use
// this.
func (catalogManagement *CatalogManagementV1) UpdateOfferingIbm(updateOfferingIbmOptions *UpdateOfferingIbmOptions) (result *ApprovalResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateOfferingIbmOptions, "updateOfferingIbmOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateOfferingIbmOptions, "updateOfferingIbmOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"catalogs", "offerings", "publish", ""}
	pathParameters := []string{*updateOfferingIbmOptions.CatalogID, *updateOfferingIbmOptions.OfferingID, *updateOfferingIbmOptions.ApprovalType, *updateOfferingIbmOptions.Approved}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateOfferingIbmOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "UpdateOfferingIbm")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApprovalResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetVersionAbout : Get the about information, in markdown, for the current version
// Get the about information, in markdown, for the current version.
func (catalogManagement *CatalogManagementV1) GetVersionAbout(getVersionAboutOptions *GetVersionAboutOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionAboutOptions, "getVersionAboutOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionAboutOptions, "getVersionAboutOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "about"}
	pathParameters := []string{*getVersionAboutOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionAboutOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersionAbout")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/markdown")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, &result)

	return
}

// GetVersionLicense : Get the license content for the specified license ID in the specified version
// Get the license content for the specified license ID in the specified version.
func (catalogManagement *CatalogManagementV1) GetVersionLicense(getVersionLicenseOptions *GetVersionLicenseOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionLicenseOptions, "getVersionLicenseOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionLicenseOptions, "getVersionLicenseOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "licenses"}
	pathParameters := []string{*getVersionLicenseOptions.VersionLocatorID, *getVersionLicenseOptions.LicenseID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionLicenseOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersionLicense")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetVersionContainerImages : Get get the list of container images associated with this version
// The "image_manifest_url" property of the version should be pointing the a URL for the image manifest, this api
// reflects that content.
func (catalogManagement *CatalogManagementV1) GetVersionContainerImages(getVersionContainerImagesOptions *GetVersionContainerImagesOptions) (result *ImageManifest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionContainerImagesOptions, "getVersionContainerImagesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionContainerImagesOptions, "getVersionContainerImagesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "containerImages"}
	pathParameters := []string{*getVersionContainerImagesOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionContainerImagesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersionContainerImages")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImageManifest)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeprecateVersion : Deprecate the specified version
// Deprecate the specified version.
func (catalogManagement *CatalogManagementV1) DeprecateVersion(deprecateVersionOptions *DeprecateVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deprecateVersionOptions, "deprecateVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deprecateVersionOptions, "deprecateVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "deprecate"}
	pathParameters := []string{*deprecateVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deprecateVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeprecateVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// AccountPublishVersion : Publish the specified version so it is viewable by account members
// Publish the specified version so it is viewable by account members.
func (catalogManagement *CatalogManagementV1) AccountPublishVersion(accountPublishVersionOptions *AccountPublishVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(accountPublishVersionOptions, "accountPublishVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(accountPublishVersionOptions, "accountPublishVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "account-publish"}
	pathParameters := []string{*accountPublishVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range accountPublishVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "AccountPublishVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// IbmPublishVersion : Publish the specified version so that it is visible to IBMers in the public catalog
// Publish the specified version so that it is visible to IBMers in the public catalog.
func (catalogManagement *CatalogManagementV1) IbmPublishVersion(ibmPublishVersionOptions *IbmPublishVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(ibmPublishVersionOptions, "ibmPublishVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(ibmPublishVersionOptions, "ibmPublishVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "ibm-publish"}
	pathParameters := []string{*ibmPublishVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range ibmPublishVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "IbmPublishVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// PublicPublishVersion : Publish the specified version so it is visible to all users in the public catalog
// Publish the specified version so it is visible to all users in the public catalog.
func (catalogManagement *CatalogManagementV1) PublicPublishVersion(publicPublishVersionOptions *PublicPublishVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(publicPublishVersionOptions, "publicPublishVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(publicPublishVersionOptions, "publicPublishVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "public-publish"}
	pathParameters := []string{*publicPublishVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range publicPublishVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "PublicPublishVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// CommitVersion : Commit a working copy of the specified version
// Commit a working copy of the specified version.
func (catalogManagement *CatalogManagementV1) CommitVersion(commitVersionOptions *CommitVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(commitVersionOptions, "commitVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(commitVersionOptions, "commitVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "commit"}
	pathParameters := []string{*commitVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range commitVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CommitVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetVersionWorkingCopy : Create a working copy of the specified version
// Create a working copy of the specified version.
func (catalogManagement *CatalogManagementV1) GetVersionWorkingCopy(getVersionWorkingCopyOptions *GetVersionWorkingCopyOptions) (result *Version, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionWorkingCopyOptions, "getVersionWorkingCopyOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionWorkingCopyOptions, "getVersionWorkingCopyOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "workingcopy"}
	pathParameters := []string{*getVersionWorkingCopyOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionWorkingCopyOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersionWorkingCopy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVersion)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetVersionUpdates : Get available updates for the specified version
// Get available updates for the specified version.
func (catalogManagement *CatalogManagementV1) GetVersionUpdates(getVersionUpdatesOptions *GetVersionUpdatesOptions) (result []VersionUpdateDescriptor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionUpdatesOptions, "getVersionUpdatesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionUpdatesOptions, "getVersionUpdatesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "updates"}
	pathParameters := []string{*getVersionUpdatesOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionUpdatesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersionUpdates")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getVersionUpdatesOptions.ClusterID != nil {
		builder.AddQuery("cluster_id", fmt.Sprint(*getVersionUpdatesOptions.ClusterID))
	}
	if getVersionUpdatesOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getVersionUpdatesOptions.Region))
	}
	if getVersionUpdatesOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*getVersionUpdatesOptions.ResourceGroupID))
	}
	if getVersionUpdatesOptions.Namespace != nil {
		builder.AddQuery("namespace", fmt.Sprint(*getVersionUpdatesOptions.Namespace))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse []json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVersionUpdateDescriptor)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetVersion : Get the Offering/Kind/Version 'branch' for the specified locator ID
// Get the Offering/Kind/Version 'branch' for the specified locator ID.
func (catalogManagement *CatalogManagementV1) GetVersion(getVersionOptions *GetVersionOptions) (result *Offering, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVersionOptions, "getVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVersionOptions, "getVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions"}
	pathParameters := []string{*getVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalOffering)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteVersion : Delete a version
// Delete a the specified version.  If the version is an active version with a working copy, the working copy will be
// deleted as well.
func (catalogManagement *CatalogManagementV1) DeleteVersion(deleteVersionOptions *DeleteVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVersionOptions, "deleteVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteVersionOptions, "deleteVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions"}
	pathParameters := []string{*deleteVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeleteVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// ListVersions : Search for versions
// [deprecated] use /search/license/versions api instead.   Search across all accounts for versions, requires global
// admin permission.
func (catalogManagement *CatalogManagementV1) ListVersions(listVersionsOptions *ListVersionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listVersionsOptions, "listVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listVersionsOptions, "listVersionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("q", fmt.Sprint(*listVersionsOptions.Q))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetRepos : List a repo's entries
// List the available entries from a given repo.
func (catalogManagement *CatalogManagementV1) GetRepos(getReposOptions *GetReposOptions) (result *HelmRepoList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getReposOptions, "getReposOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getReposOptions, "getReposOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"repo", "entries"}
	pathParameters := []string{*getReposOptions.Type}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getReposOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetRepos")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("repourl", fmt.Sprint(*getReposOptions.Repourl))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHelmRepoList)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetRepo : Get contents of a repo
// Get the contents of a given repo.
func (catalogManagement *CatalogManagementV1) GetRepo(getRepoOptions *GetRepoOptions) (result *HelmPackage, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRepoOptions, "getRepoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRepoOptions, "getRepoOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"repo"}
	pathParameters := []string{*getRepoOptions.Type}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRepoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetRepo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("charturl", fmt.Sprint(*getRepoOptions.Charturl))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalHelmPackage)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListClusters : List Kube clusters
// List Kube clusters.
func (catalogManagement *CatalogManagementV1) ListClusters(listClustersOptions *ListClustersOptions) (result *ClusterSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listClustersOptions, "listClustersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes/clusters"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listClustersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListClusters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listClustersOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listClustersOptions.Limit))
	}
	if listClustersOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listClustersOptions.Offset))
	}
	if listClustersOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listClustersOptions.Type))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClusterSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetCluster : Get Kube cluster
// Get Kube cluster.
func (catalogManagement *CatalogManagementV1) GetCluster(getClusterOptions *GetClusterOptions) (result *ClusterInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getClusterOptions, "getClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getClusterOptions, "getClusterOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes/clusters"}
	pathParameters := []string{*getClusterOptions.ClusterID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getClusterOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*getClusterOptions.XAuthRefreshToken))
	}

	builder.AddQuery("region", fmt.Sprint(*getClusterOptions.Region))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalClusterInfo)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetNamespaces : Get cluster namespaces
// Get cluster namespaces.
func (catalogManagement *CatalogManagementV1) GetNamespaces(getNamespacesOptions *GetNamespacesOptions) (result *NamespaceSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getNamespacesOptions, "getNamespacesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getNamespacesOptions, "getNamespacesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes/clusters", "namespaces"}
	pathParameters := []string{*getNamespacesOptions.ClusterID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getNamespacesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetNamespaces")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getNamespacesOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*getNamespacesOptions.XAuthRefreshToken))
	}

	builder.AddQuery("region", fmt.Sprint(*getNamespacesOptions.Region))
	if getNamespacesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*getNamespacesOptions.Limit))
	}
	if getNamespacesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*getNamespacesOptions.Offset))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNamespaceSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// InstallVersion : Create an install
// Create an install.
func (catalogManagement *CatalogManagementV1) InstallVersion(installVersionOptions *InstallVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(installVersionOptions, "installVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(installVersionOptions, "installVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "install"}
	pathParameters := []string{*installVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range installVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "InstallVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if installVersionOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*installVersionOptions.XAuthRefreshToken))
	}

	body := make(map[string]interface{})
	if installVersionOptions.NewClusterID != nil {
		body["cluster_id"] = installVersionOptions.NewClusterID
	}
	if installVersionOptions.NewRegion != nil {
		body["region"] = installVersionOptions.NewRegion
	}
	if installVersionOptions.NewNamespace != nil {
		body["namespace"] = installVersionOptions.NewNamespace
	}
	if installVersionOptions.NewOverrideValues != nil {
		body["override_values"] = installVersionOptions.NewOverrideValues
	}
	if installVersionOptions.NewEntitlementApikey != nil {
		body["entitlement_apikey"] = installVersionOptions.NewEntitlementApikey
	}
	if installVersionOptions.NewSchematics != nil {
		body["schematics"] = installVersionOptions.NewSchematics
	}
	if installVersionOptions.NewScript != nil {
		body["script"] = installVersionOptions.NewScript
	}
	if installVersionOptions.NewScriptID != nil {
		body["script_id"] = installVersionOptions.NewScriptID
	}
	if installVersionOptions.NewVersionLocatorID != nil {
		body["version_locator_id"] = installVersionOptions.NewVersionLocatorID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// PreinstallVersion : Create a preinstall
// Create a preinstall.
func (catalogManagement *CatalogManagementV1) PreinstallVersion(preinstallVersionOptions *PreinstallVersionOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(preinstallVersionOptions, "preinstallVersionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(preinstallVersionOptions, "preinstallVersionOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "preinstall"}
	pathParameters := []string{*preinstallVersionOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range preinstallVersionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "PreinstallVersion")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if preinstallVersionOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*preinstallVersionOptions.XAuthRefreshToken))
	}

	body := make(map[string]interface{})
	if preinstallVersionOptions.NewClusterID != nil {
		body["cluster_id"] = preinstallVersionOptions.NewClusterID
	}
	if preinstallVersionOptions.NewRegion != nil {
		body["region"] = preinstallVersionOptions.NewRegion
	}
	if preinstallVersionOptions.NewNamespace != nil {
		body["namespace"] = preinstallVersionOptions.NewNamespace
	}
	if preinstallVersionOptions.NewOverrideValues != nil {
		body["override_values"] = preinstallVersionOptions.NewOverrideValues
	}
	if preinstallVersionOptions.NewEntitlementApikey != nil {
		body["entitlement_apikey"] = preinstallVersionOptions.NewEntitlementApikey
	}
	if preinstallVersionOptions.NewSchematics != nil {
		body["schematics"] = preinstallVersionOptions.NewSchematics
	}
	if preinstallVersionOptions.NewScript != nil {
		body["script"] = preinstallVersionOptions.NewScript
	}
	if preinstallVersionOptions.NewScriptID != nil {
		body["script_id"] = preinstallVersionOptions.NewScriptID
	}
	if preinstallVersionOptions.NewVersionLocatorID != nil {
		body["version_locator_id"] = preinstallVersionOptions.NewVersionLocatorID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetPreinstall : Get a preinstall
// Get a preinstall.
func (catalogManagement *CatalogManagementV1) GetPreinstall(getPreinstallOptions *GetPreinstallOptions) (result *InstallStatus, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPreinstallOptions, "getPreinstallOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPreinstallOptions, "getPreinstallOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "preinstall"}
	pathParameters := []string{*getPreinstallOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPreinstallOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetPreinstall")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPreinstallOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*getPreinstallOptions.XAuthRefreshToken))
	}

	if getPreinstallOptions.ClusterID != nil {
		builder.AddQuery("cluster_id", fmt.Sprint(*getPreinstallOptions.ClusterID))
	}
	if getPreinstallOptions.Region != nil {
		builder.AddQuery("region", fmt.Sprint(*getPreinstallOptions.Region))
	}
	if getPreinstallOptions.Namespace != nil {
		builder.AddQuery("namespace", fmt.Sprint(*getPreinstallOptions.Namespace))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalInstallStatus)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ValidationInstall : Validate a offering
// Validate a offering.
func (catalogManagement *CatalogManagementV1) ValidationInstall(validationInstallOptions *ValidationInstallOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(validationInstallOptions, "validationInstallOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(validationInstallOptions, "validationInstallOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "validation/install"}
	pathParameters := []string{*validationInstallOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range validationInstallOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ValidationInstall")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")
	if validationInstallOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*validationInstallOptions.XAuthRefreshToken))
	}

	body := make(map[string]interface{})
	if validationInstallOptions.NewClusterID != nil {
		body["cluster_id"] = validationInstallOptions.NewClusterID
	}
	if validationInstallOptions.NewRegion != nil {
		body["region"] = validationInstallOptions.NewRegion
	}
	if validationInstallOptions.NewNamespace != nil {
		body["namespace"] = validationInstallOptions.NewNamespace
	}
	if validationInstallOptions.NewOverrideValues != nil {
		body["override_values"] = validationInstallOptions.NewOverrideValues
	}
	if validationInstallOptions.NewEntitlementApikey != nil {
		body["entitlement_apikey"] = validationInstallOptions.NewEntitlementApikey
	}
	if validationInstallOptions.NewSchematics != nil {
		body["schematics"] = validationInstallOptions.NewSchematics
	}
	if validationInstallOptions.NewScript != nil {
		body["script"] = validationInstallOptions.NewScript
	}
	if validationInstallOptions.NewScriptID != nil {
		body["script_id"] = validationInstallOptions.NewScriptID
	}
	if validationInstallOptions.NewVersionLocatorID != nil {
		body["version_locator_id"] = validationInstallOptions.NewVersionLocatorID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetValidationStatus : Returns the install status for the specified offering version
// Returns the install status for the specified offering version.
func (catalogManagement *CatalogManagementV1) GetValidationStatus(getValidationStatusOptions *GetValidationStatusOptions) (result *Validation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getValidationStatusOptions, "getValidationStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getValidationStatusOptions, "getValidationStatusOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "validation/install"}
	pathParameters := []string{*getValidationStatusOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getValidationStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetValidationStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getValidationStatusOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*getValidationStatusOptions.XAuthRefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalValidation)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetOverrideValues : Returns the override values that were used to validate the specified offering version
// Returns the override values that were used to validate the specified offering version.
func (catalogManagement *CatalogManagementV1) GetOverrideValues(getOverrideValuesOptions *GetOverrideValuesOptions) (result map[string]interface{}, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getOverrideValuesOptions, "getOverrideValuesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getOverrideValuesOptions, "getOverrideValuesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "validation/overridevalues"}
	pathParameters := []string{*getOverrideValuesOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getOverrideValuesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetOverrideValues")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, &result)

	return
}

// GetSchematicsWorkspaces : Returns the schematics workspaces for the specified offering version
// Returns the schematics workspaces for the specified offering version.
func (catalogManagement *CatalogManagementV1) GetSchematicsWorkspaces(getSchematicsWorkspacesOptions *GetSchematicsWorkspacesOptions) (result *SchematicsWorkspaceSearchResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSchematicsWorkspacesOptions, "getSchematicsWorkspacesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSchematicsWorkspacesOptions, "getSchematicsWorkspacesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "workspaces"}
	pathParameters := []string{*getSchematicsWorkspacesOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSchematicsWorkspacesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetSchematicsWorkspaces")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getSchematicsWorkspacesOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*getSchematicsWorkspacesOptions.XAuthRefreshToken))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSchematicsWorkspaceSearchResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CanDeploySchematics : Returns the schematics permissions for the specified user
// Returns the schematics permissions for the specified user.
func (catalogManagement *CatalogManagementV1) CanDeploySchematics(canDeploySchematicsOptions *CanDeploySchematicsOptions) (result *DeployRequirementsCheck, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(canDeploySchematicsOptions, "canDeploySchematicsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(canDeploySchematicsOptions, "canDeploySchematicsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"versions", "candeploy"}
	pathParameters := []string{*canDeploySchematicsOptions.VersionLocatorID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range canDeploySchematicsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CanDeploySchematics")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("cluster_id", fmt.Sprint(*canDeploySchematicsOptions.ClusterID))
	builder.AddQuery("region", fmt.Sprint(*canDeploySchematicsOptions.Region))
	if canDeploySchematicsOptions.Namespace != nil {
		builder.AddQuery("namespace", fmt.Sprint(*canDeploySchematicsOptions.Namespace))
	}
	if canDeploySchematicsOptions.ResourceGroupID != nil {
		builder.AddQuery("resource_group_id", fmt.Sprint(*canDeploySchematicsOptions.ResourceGroupID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeployRequirementsCheck)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetResourceGroups : Returns all active resource groups in the current account, where the current user has permission to create schematics workspaces
// Returns all active resource groups in the current account, where the current user has permission to create schematics
// workspaces.
func (catalogManagement *CatalogManagementV1) GetResourceGroups(getResourceGroupsOptions *GetResourceGroupsOptions) (result *ResourceGroups, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getResourceGroupsOptions, "getResourceGroupsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/schematics/resourcegroups"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceGroupsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetResourceGroups")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceGroups)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateKubeDeploy : Create a deployment on a Kube cluster
// Create a deployment on a Kube cluster.
func (catalogManagement *CatalogManagementV1) CreateKubeDeploy(createKubeDeployOptions *CreateKubeDeployOptions) (result *KubeDeployResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createKubeDeployOptions, "createKubeDeployOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createKubeDeployOptions, "createKubeDeployOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createKubeDeployOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CreateKubeDeploy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createKubeDeployOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*createKubeDeployOptions.XAuthRefreshToken))
	}

	body := make(map[string]interface{})
	if createKubeDeployOptions.ClusterID != nil {
		body["cluster_id"] = createKubeDeployOptions.ClusterID
	}
	if createKubeDeployOptions.Region != nil {
		body["region"] = createKubeDeployOptions.Region
	}
	if createKubeDeployOptions.Namespace != nil {
		body["namespace"] = createKubeDeployOptions.Namespace
	}
	if createKubeDeployOptions.VersionLocatorID != nil {
		body["version_locator_id"] = createKubeDeployOptions.VersionLocatorID
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKubeDeployResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListKubeDeploy : Get deployment from a Kube cluster
// Get deployment from a Kube cluster.
func (catalogManagement *CatalogManagementV1) ListKubeDeploy(listKubeDeployOptions *ListKubeDeployOptions) (result *KubeDeployResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listKubeDeployOptions, "listKubeDeployOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listKubeDeployOptions, "listKubeDeployOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listKubeDeployOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListKubeDeploy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listKubeDeployOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*listKubeDeployOptions.XAuthRefreshToken))
	}

	builder.AddQuery("cluster_id", fmt.Sprint(*listKubeDeployOptions.ClusterID))
	builder.AddQuery("namespace", fmt.Sprint(*listKubeDeployOptions.Namespace))
	builder.AddQuery("region", fmt.Sprint(*listKubeDeployOptions.Region))
	builder.AddQuery("version_locator_id", fmt.Sprint(*listKubeDeployOptions.VersionLocatorID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKubeDeployResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ReplaceKubeDeploy : Update deployment on Kube cluster
// Update deployment on Kube cluster.
func (catalogManagement *CatalogManagementV1) ReplaceKubeDeploy(replaceKubeDeployOptions *ReplaceKubeDeployOptions) (result *KubeDeployResult, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceKubeDeployOptions, "replaceKubeDeployOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceKubeDeployOptions, "replaceKubeDeployOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.PUT)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceKubeDeployOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ReplaceKubeDeploy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if replaceKubeDeployOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*replaceKubeDeployOptions.XAuthRefreshToken))
	}

	body := make(map[string]interface{})
	if replaceKubeDeployOptions.ClusterID != nil {
		body["cluster_id"] = replaceKubeDeployOptions.ClusterID
	}
	if replaceKubeDeployOptions.Region != nil {
		body["region"] = replaceKubeDeployOptions.Region
	}
	if replaceKubeDeployOptions.Namespace != nil {
		body["namespace"] = replaceKubeDeployOptions.Namespace
	}
	if replaceKubeDeployOptions.VersionLocatorID != nil {
		body["version_locator_id"] = replaceKubeDeployOptions.VersionLocatorID
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalKubeDeployResult)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteKubeDeploy : Delete deployment from Kube cluster
// Delete deployment from Kube cluster.
func (catalogManagement *CatalogManagementV1) DeleteKubeDeploy(deleteKubeDeployOptions *DeleteKubeDeployOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteKubeDeployOptions, "deleteKubeDeployOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteKubeDeployOptions, "deleteKubeDeployOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"deploy/kubernetes"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteKubeDeployOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeleteKubeDeploy")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteKubeDeployOptions.XAuthRefreshToken != nil {
		builder.AddHeader("X-Auth-Refresh-Token", fmt.Sprint(*deleteKubeDeployOptions.XAuthRefreshToken))
	}

	builder.AddQuery("cluster_id", fmt.Sprint(*deleteKubeDeployOptions.ClusterID))
	builder.AddQuery("namespace", fmt.Sprint(*deleteKubeDeployOptions.Namespace))
	builder.AddQuery("region", fmt.Sprint(*deleteKubeDeployOptions.Region))
	builder.AddQuery("version_locator_id", fmt.Sprint(*deleteKubeDeployOptions.VersionLocatorID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetLicenseProviders : Get license providers
// Get license providers.
func (catalogManagement *CatalogManagementV1) GetLicenseProviders(getLicenseProvidersOptions *GetLicenseProvidersOptions) (result *LicenseProviders, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getLicenseProvidersOptions, "getLicenseProvidersOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/license_providers"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLicenseProvidersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetLicenseProviders")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLicenseProviders)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ListLicenseEntitlements : Get license entitlements
// Get license entitlements bound to an account.
func (catalogManagement *CatalogManagementV1) ListLicenseEntitlements(listLicenseEntitlementsOptions *ListLicenseEntitlementsOptions) (result *LicenseEntitlements, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listLicenseEntitlementsOptions, "listLicenseEntitlementsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/entitlements"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLicenseEntitlementsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "ListLicenseEntitlements")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listLicenseEntitlementsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*listLicenseEntitlementsOptions.AccountID))
	}
	if listLicenseEntitlementsOptions.LicenseProductID != nil {
		builder.AddQuery("license_product_id", fmt.Sprint(*listLicenseEntitlementsOptions.LicenseProductID))
	}
	if listLicenseEntitlementsOptions.VersionID != nil {
		builder.AddQuery("version_id", fmt.Sprint(*listLicenseEntitlementsOptions.VersionID))
	}
	if listLicenseEntitlementsOptions.State != nil {
		builder.AddQuery("state", fmt.Sprint(*listLicenseEntitlementsOptions.State))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLicenseEntitlements)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateLicenseEntitlement : Create a license entitlement
// Create an entitlement for a Cloud account.  This is used to give an account an entitlement to a license.
func (catalogManagement *CatalogManagementV1) CreateLicenseEntitlement(createLicenseEntitlementOptions *CreateLicenseEntitlementOptions) (result *LicenseEntitlement, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createLicenseEntitlementOptions, "createLicenseEntitlementOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/entitlements"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.POST)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLicenseEntitlementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "CreateLicenseEntitlement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if createLicenseEntitlementOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*createLicenseEntitlementOptions.AccountID))
	}

	body := make(map[string]interface{})
	if createLicenseEntitlementOptions.Name != nil {
		body["name"] = createLicenseEntitlementOptions.Name
	}
	if createLicenseEntitlementOptions.EffectiveFrom != nil {
		body["effective_from"] = createLicenseEntitlementOptions.EffectiveFrom
	}
	if createLicenseEntitlementOptions.EffectiveTo != nil {
		body["effective_to"] = createLicenseEntitlementOptions.EffectiveTo
	}
	if createLicenseEntitlementOptions.VersionID != nil {
		body["version_id"] = createLicenseEntitlementOptions.VersionID
	}
	if createLicenseEntitlementOptions.LicenseID != nil {
		body["license_id"] = createLicenseEntitlementOptions.LicenseID
	}
	if createLicenseEntitlementOptions.LicenseOwnerID != nil {
		body["license_owner_id"] = createLicenseEntitlementOptions.LicenseOwnerID
	}
	if createLicenseEntitlementOptions.NamespaceRepository != nil {
		body["namespace_repository"] = createLicenseEntitlementOptions.NamespaceRepository
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
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLicenseEntitlement)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetLicenseEntitlements : Get entitlements for a specific license product ID
// Get an entitlements for a specific license product ID bound to an account.
func (catalogManagement *CatalogManagementV1) GetLicenseEntitlements(getLicenseEntitlementsOptions *GetLicenseEntitlementsOptions) (result *LicenseEntitlements, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLicenseEntitlementsOptions, "getLicenseEntitlementsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLicenseEntitlementsOptions, "getLicenseEntitlementsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/entitlements/productID"}
	pathParameters := []string{*getLicenseEntitlementsOptions.LicenseProductID}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLicenseEntitlementsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetLicenseEntitlements")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getLicenseEntitlementsOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getLicenseEntitlementsOptions.AccountID))
	}
	if getLicenseEntitlementsOptions.VersionID != nil {
		builder.AddQuery("version_id", fmt.Sprint(*getLicenseEntitlementsOptions.VersionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLicenseEntitlements)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteLicenseEntitlement : Delete license entitlement
// Delete a license entitlement that is bound to an account. Note that BSS will mark the entitlement field "state":
// "removed".
func (catalogManagement *CatalogManagementV1) DeleteLicenseEntitlement(deleteLicenseEntitlementOptions *DeleteLicenseEntitlementOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLicenseEntitlementOptions, "deleteLicenseEntitlementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLicenseEntitlementOptions, "deleteLicenseEntitlementOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/entitlements"}
	pathParameters := []string{*deleteLicenseEntitlementOptions.EntitlementID}

	builder := core.NewRequestBuilder(core.DELETE)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLicenseEntitlementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "DeleteLicenseEntitlement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	if deleteLicenseEntitlementOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*deleteLicenseEntitlementOptions.AccountID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// GetLicenses : Get licenses
// Retrieve available licenses from supported license subsystems.  This is used to get the list of available licenses
// that the user has.
func (catalogManagement *CatalogManagementV1) GetLicenses(getLicensesOptions *GetLicensesOptions) (result *Licenses, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLicensesOptions, "getLicensesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLicensesOptions, "getLicensesOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"license/licenses"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLicensesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "GetLicenses")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("license_provider_id", fmt.Sprint(*getLicensesOptions.LicenseProviderID))
	if getLicensesOptions.AccountID != nil {
		builder.AddQuery("account_id", fmt.Sprint(*getLicensesOptions.AccountID))
	}
	if getLicensesOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*getLicensesOptions.Name))
	}
	if getLicensesOptions.LicenseType != nil {
		builder.AddQuery("license_type", fmt.Sprint(*getLicensesOptions.LicenseType))
	}
	if getLicensesOptions.LicenseProductID != nil {
		builder.AddQuery("license_product_id", fmt.Sprint(*getLicensesOptions.LicenseProductID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = catalogManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLicenses)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// SearchLicenseVersions : Search for versions
// Search across accounts for all versions usig a particular license, requires global admin permission.
func (catalogManagement *CatalogManagementV1) SearchLicenseVersions(searchLicenseVersionsOptions *SearchLicenseVersionsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(searchLicenseVersionsOptions, "searchLicenseVersionsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(searchLicenseVersionsOptions, "searchLicenseVersionsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"search/license/versions"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range searchLicenseVersionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "SearchLicenseVersions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("q", fmt.Sprint(*searchLicenseVersionsOptions.Q))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// SearchLicenseOfferings : Search for Offerings
// Search across accounts for all offerings using a particular license, requires global admin permission.
func (catalogManagement *CatalogManagementV1) SearchLicenseOfferings(searchLicenseOfferingsOptions *SearchLicenseOfferingsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(searchLicenseOfferingsOptions, "searchLicenseOfferingsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(searchLicenseOfferingsOptions, "searchLicenseOfferingsOptions")
	if err != nil {
		return
	}

	pathSegments := []string{"search/license/offerings"}
	pathParameters := []string{}

	builder := core.NewRequestBuilder(core.GET)
	_, err = builder.ConstructHTTPURL(catalogManagement.Service.Options.URL, pathSegments, pathParameters)
	if err != nil {
		return
	}

	for headerName, headerValue := range searchLicenseOfferingsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("catalog_management", "V1", "SearchLicenseOfferings")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("q", fmt.Sprint(*searchLicenseOfferingsOptions.Q))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = catalogManagement.Service.Request(request, nil)

	return
}

// Account : Account information.
type Account struct {
	// Account identification.
	ID *string `json:"id,omitempty"`

	// Filters for account and catalog filters.
	AccountFilters *Filters `json:"account_filters,omitempty"`
}


// UnmarshalAccount unmarshals an instance of Account from the specified map of raw messages.
func UnmarshalAccount(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Account)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account_filters", &obj.AccountFilters, UnmarshalFilters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountGroup : Filters for an Account Group.
type AccountGroup struct {
	// Account group identification.
	ID *string `json:"id,omitempty"`

	// Filters for account and catalog filters.
	AccountFilters *Filters `json:"account_filters,omitempty"`
}


// UnmarshalAccountGroup unmarshals an instance of AccountGroup from the specified map of raw messages.
func UnmarshalAccountGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccountGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account_filters", &obj.AccountFilters, UnmarshalFilters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccountPublishVersionOptions : The AccountPublishVersion options.
type AccountPublishVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAccountPublishVersionOptions : Instantiate AccountPublishVersionOptions
func (*CatalogManagementV1) NewAccountPublishVersionOptions(versionLocatorID string) *AccountPublishVersionOptions {
	return &AccountPublishVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *AccountPublishVersionOptions) SetVersionLocatorID(versionLocatorID string) *AccountPublishVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *AccountPublishVersionOptions) SetHeaders(param map[string]string) *AccountPublishVersionOptions {
	options.Headers = param
	return options
}

// AccumulatedFilters : The accumulated filters for an account. This will return the account filters plus a filter for each catalog the user
// has access to.
type AccumulatedFilters struct {
	// Filters for accounts (at this time this will always be just one item array).
	AccountFilters []Filters `json:"account_filters,omitempty"`

	// The filters for all of the accessible catalogs.
	CatalogFilters []AccumulatedFiltersCatalogFiltersItem `json:"catalog_filters,omitempty"`
}


// UnmarshalAccumulatedFilters unmarshals an instance of AccumulatedFilters from the specified map of raw messages.
func UnmarshalAccumulatedFilters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccumulatedFilters)
	err = core.UnmarshalModel(m, "account_filters", &obj.AccountFilters, UnmarshalFilters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "catalog_filters", &obj.CatalogFilters, UnmarshalAccumulatedFiltersCatalogFiltersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccumulatedFiltersCatalogFiltersItem : AccumulatedFiltersCatalogFiltersItem struct
type AccumulatedFiltersCatalogFiltersItem struct {
	// Filters for catalog.
	Catalog *AccumulatedFiltersCatalogFiltersItemCatalog `json:"catalog,omitempty"`

	// Filters for account and catalog filters.
	Filters *Filters `json:"filters,omitempty"`
}


// UnmarshalAccumulatedFiltersCatalogFiltersItem unmarshals an instance of AccumulatedFiltersCatalogFiltersItem from the specified map of raw messages.
func UnmarshalAccumulatedFiltersCatalogFiltersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccumulatedFiltersCatalogFiltersItem)
	err = core.UnmarshalModel(m, "catalog", &obj.Catalog, UnmarshalAccumulatedFiltersCatalogFiltersItemCatalog)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "filters", &obj.Filters, UnmarshalFilters)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AccumulatedFiltersCatalogFiltersItemCatalog : Filters for catalog.
type AccumulatedFiltersCatalogFiltersItemCatalog struct {
	// The ID of the catalog.
	ID *string `json:"id,omitempty"`

	// The name of the catalog.
	Name *string `json:"name,omitempty"`
}


// UnmarshalAccumulatedFiltersCatalogFiltersItemCatalog unmarshals an instance of AccumulatedFiltersCatalogFiltersItemCatalog from the specified map of raw messages.
func UnmarshalAccumulatedFiltersCatalogFiltersItemCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AccumulatedFiltersCatalogFiltersItemCatalog)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApprovalResult : Result of approval.
type ApprovalResult struct {
	// Visible to IBM.
	Ibm *bool `json:"ibm,omitempty"`

	// Visible to everyone.
	Public *bool `json:"public,omitempty"`

	// Denotes whether approval has changed.
	Changed *bool `json:"changed,omitempty"`
}


// UnmarshalApprovalResult unmarshals an instance of ApprovalResult from the specified map of raw messages.
func UnmarshalApprovalResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApprovalResult)
	err = core.UnmarshalPrimitive(m, "ibm", &obj.Ibm)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public", &obj.Public)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "changed", &obj.Changed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CanDeploySchematicsOptions : The CanDeploySchematics options.
type CanDeploySchematicsOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// ID of the cluster.
	ClusterID *string `json:"cluster_id" validate:"required"`

	// Cluster region.
	Region *string `json:"region" validate:"required"`

	// Required if the version's pre-install scope is `namespace`.
	Namespace *string `json:"namespace,omitempty"`

	// Resource group identification.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCanDeploySchematicsOptions : Instantiate CanDeploySchematicsOptions
func (*CatalogManagementV1) NewCanDeploySchematicsOptions(versionLocatorID string, clusterID string, region string) *CanDeploySchematicsOptions {
	return &CanDeploySchematicsOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		ClusterID: core.StringPtr(clusterID),
		Region: core.StringPtr(region),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *CanDeploySchematicsOptions) SetVersionLocatorID(versionLocatorID string) *CanDeploySchematicsOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *CanDeploySchematicsOptions) SetClusterID(clusterID string) *CanDeploySchematicsOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *CanDeploySchematicsOptions) SetRegion(region string) *CanDeploySchematicsOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *CanDeploySchematicsOptions) SetNamespace(namespace string) *CanDeploySchematicsOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *CanDeploySchematicsOptions) SetResourceGroupID(resourceGroupID string) *CanDeploySchematicsOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CanDeploySchematicsOptions) SetHeaders(param map[string]string) *CanDeploySchematicsOptions {
	options.Headers = param
	return options
}

// Catalog : Catalog information.
type Catalog struct {
	// Unique ID.
	ID *string `json:"id,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// Description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// URL for an icon associated with this catalog.
	CatalogIconURL *string `json:"catalog_icon_url,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// The url for this specific catalog.
	URL *string `json:"url,omitempty"`

	// CRN associated with the catalog.
	Crn *string `json:"crn,omitempty"`

	// URL path to offerings.
	OfferingsURL *string `json:"offerings_url,omitempty"`

	// List of features associated with this catalog.
	Features []Feature `json:"features,omitempty"`

	// Denotes whether a catalog is disabled.
	Disabled *bool `json:"disabled,omitempty"`

	// The date'time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date'time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Resource group id the catalog is owned by.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Account that owns catalog.
	OwningAccount *string `json:"owning_account,omitempty"`

	// Filters for account and catalog filters.
	CatalogFilters *Filters `json:"catalog_filters,omitempty"`

	// Feature information.
	SyndicationSettings *SyndicationResource `json:"syndication_settings,omitempty"`
}


// UnmarshalCatalog unmarshals an instance of Catalog from the specified map of raw messages.
func UnmarshalCatalog(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Catalog)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_icon_url", &obj.CatalogIconURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offerings_url", &obj.OfferingsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disabled", &obj.Disabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owning_account", &obj.OwningAccount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "catalog_filters", &obj.CatalogFilters, UnmarshalFilters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "syndication_settings", &obj.SyndicationSettings, UnmarshalSyndicationResource)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CatalogSearchResult : Paginated catalog search result.
type CatalogSearchResult struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting objects.
	Resources []Catalog `json:"resources,omitempty"`
}


// UnmarshalCatalogSearchResult unmarshals an instance of CatalogSearchResult from the specified map of raw messages.
func UnmarshalCatalogSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CatalogSearchResult)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalCatalog)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CategoryFilter : Filter on a category. The filter will match against the values of the given category with include or exclude.
type CategoryFilter struct {
	// -> true - This is an include filter, false - this is an exclude filter.
	Include *bool `json:"include,omitempty"`

	// Offering filter terms.
	Filter *FilterTerms `json:"filter,omitempty"`
}


// UnmarshalCategoryFilter unmarshals an instance of CategoryFilter from the specified map of raw messages.
func UnmarshalCategoryFilter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CategoryFilter)
	err = core.UnmarshalPrimitive(m, "include", &obj.Include)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "filter", &obj.Filter, UnmarshalFilterTerms)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterInfo : Cluster information.
type ClusterInfo struct {
	// Resource Group ID.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Resource Group name.
	ResourceGroupName *string `json:"resource_group_name,omitempty"`

	// Cluster ID.
	ID *string `json:"id,omitempty"`

	// Cluster name.
	Name *string `json:"name,omitempty"`

	// Cluster region.
	Region *string `json:"region,omitempty"`
}


// UnmarshalClusterInfo unmarshals an instance of ClusterInfo from the specified map of raw messages.
func UnmarshalClusterInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterInfo)
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_name", &obj.ResourceGroupName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterRequestBodySchematics : Schematics workspace configuration.
type ClusterRequestBodySchematics struct {
	// Schematics workspace name.
	Name *string `json:"name,omitempty"`

	// Schematics workspace description.
	Description *string `json:"description,omitempty"`

	// Schematics workspace tags.
	Tags []string `json:"tags,omitempty"`

	// Resource group to use when creating the schematics workspace.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`
}


// UnmarshalClusterRequestBodySchematics unmarshals an instance of ClusterRequestBodySchematics from the specified map of raw messages.
func UnmarshalClusterRequestBodySchematics(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterRequestBodySchematics)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterSearchResult : Paginated cluster search result.
type ClusterSearchResult struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting objects.
	Resources []ClusterInfo `json:"resources,omitempty"`
}


// UnmarshalClusterSearchResult unmarshals an instance of ClusterSearchResult from the specified map of raw messages.
func UnmarshalClusterSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterSearchResult)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalClusterInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CommitVersionOptions : The CommitVersion options.
type CommitVersionOptions struct {
	// A dotted value of `catalogID`.`versionID` for either the working copy to commit or the active version with a working
	// copy to be committed.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCommitVersionOptions : Instantiate CommitVersionOptions
func (*CatalogManagementV1) NewCommitVersionOptions(versionLocatorID string) *CommitVersionOptions {
	return &CommitVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *CommitVersionOptions) SetVersionLocatorID(versionLocatorID string) *CommitVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CommitVersionOptions) SetHeaders(param map[string]string) *CommitVersionOptions {
	options.Headers = param
	return options
}

// Configuration : Configuration description.
type Configuration struct {
	// Configuration key.
	Key *string `json:"key,omitempty"`

	// Value type (string, boolean, int).
	Type *string `json:"type,omitempty"`

	// The default value.
	DefaultValue *string `json:"default_value,omitempty"`

	// Constraint associated with value, e.g., for string type - regx:[a-z].
	ValueConstraint *string `json:"value_constraint,omitempty"`

	// Key description.
	Description *string `json:"description,omitempty"`

	// Is key required to install.
	Required *bool `json:"required,omitempty"`

	// List of options of type.
	Options []interface{} `json:"options,omitempty"`

	// Hide values.
	Hidden *bool `json:"hidden,omitempty"`
}


// UnmarshalConfiguration unmarshals an instance of Configuration from the specified map of raw messages.
func UnmarshalConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Configuration)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_value", &obj.DefaultValue)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value_constraint", &obj.ValueConstraint)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "required", &obj.Required)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "options", &obj.Options)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hidden", &obj.Hidden)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCatalogOptions : The CreateCatalog options.
type CreateCatalogOptions struct {
	// Unique ID.
	ID *string `json:"id,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// Description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// URL for an icon associated with this catalog.
	CatalogIconURL *string `json:"catalog_icon_url,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// The url for this specific catalog.
	URL *string `json:"url,omitempty"`

	// CRN associated with the catalog.
	Crn *string `json:"crn,omitempty"`

	// URL path to offerings.
	OfferingsURL *string `json:"offerings_url,omitempty"`

	// List of features associated with this catalog.
	Features []Feature `json:"features,omitempty"`

	// Denotes whether a catalog is disabled.
	Disabled *bool `json:"disabled,omitempty"`

	// The date'time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date'time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Resource group id the catalog is owned by.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Account that owns catalog.
	OwningAccount *string `json:"owning_account,omitempty"`

	// Filters for account and catalog filters.
	CatalogFilters *Filters `json:"catalog_filters,omitempty"`

	// Feature information.
	SyndicationSettings *SyndicationResource `json:"syndication_settings,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCatalogOptions : Instantiate CreateCatalogOptions
func (*CatalogManagementV1) NewCreateCatalogOptions() *CreateCatalogOptions {
	return &CreateCatalogOptions{}
}

// SetID : Allow user to set ID
func (options *CreateCatalogOptions) SetID(id string) *CreateCatalogOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetLabel : Allow user to set Label
func (options *CreateCatalogOptions) SetLabel(label string) *CreateCatalogOptions {
	options.Label = core.StringPtr(label)
	return options
}

// SetShortDescription : Allow user to set ShortDescription
func (options *CreateCatalogOptions) SetShortDescription(shortDescription string) *CreateCatalogOptions {
	options.ShortDescription = core.StringPtr(shortDescription)
	return options
}

// SetCatalogIconURL : Allow user to set CatalogIconURL
func (options *CreateCatalogOptions) SetCatalogIconURL(catalogIconURL string) *CreateCatalogOptions {
	options.CatalogIconURL = core.StringPtr(catalogIconURL)
	return options
}

// SetTags : Allow user to set Tags
func (options *CreateCatalogOptions) SetTags(tags []string) *CreateCatalogOptions {
	options.Tags = tags
	return options
}

// SetURL : Allow user to set URL
func (options *CreateCatalogOptions) SetURL(url string) *CreateCatalogOptions {
	options.URL = core.StringPtr(url)
	return options
}

// SetCrn : Allow user to set Crn
func (options *CreateCatalogOptions) SetCrn(crn string) *CreateCatalogOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetOfferingsURL : Allow user to set OfferingsURL
func (options *CreateCatalogOptions) SetOfferingsURL(offeringsURL string) *CreateCatalogOptions {
	options.OfferingsURL = core.StringPtr(offeringsURL)
	return options
}

// SetFeatures : Allow user to set Features
func (options *CreateCatalogOptions) SetFeatures(features []Feature) *CreateCatalogOptions {
	options.Features = features
	return options
}

// SetDisabled : Allow user to set Disabled
func (options *CreateCatalogOptions) SetDisabled(disabled bool) *CreateCatalogOptions {
	options.Disabled = core.BoolPtr(disabled)
	return options
}

// SetCreated : Allow user to set Created
func (options *CreateCatalogOptions) SetCreated(created *strfmt.DateTime) *CreateCatalogOptions {
	options.Created = created
	return options
}

// SetUpdated : Allow user to set Updated
func (options *CreateCatalogOptions) SetUpdated(updated *strfmt.DateTime) *CreateCatalogOptions {
	options.Updated = updated
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *CreateCatalogOptions) SetResourceGroupID(resourceGroupID string) *CreateCatalogOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetOwningAccount : Allow user to set OwningAccount
func (options *CreateCatalogOptions) SetOwningAccount(owningAccount string) *CreateCatalogOptions {
	options.OwningAccount = core.StringPtr(owningAccount)
	return options
}

// SetCatalogFilters : Allow user to set CatalogFilters
func (options *CreateCatalogOptions) SetCatalogFilters(catalogFilters *Filters) *CreateCatalogOptions {
	options.CatalogFilters = catalogFilters
	return options
}

// SetSyndicationSettings : Allow user to set SyndicationSettings
func (options *CreateCatalogOptions) SetSyndicationSettings(syndicationSettings *SyndicationResource) *CreateCatalogOptions {
	options.SyndicationSettings = syndicationSettings
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCatalogOptions) SetHeaders(param map[string]string) *CreateCatalogOptions {
	options.Headers = param
	return options
}

// CreateKubeDeployOptions : The CreateKubeDeploy options.
type CreateKubeDeployOptions struct {
	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster ID.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Cluster region.
	Region *string `json:"region,omitempty"`

	// Kube namespace.
	Namespace *string `json:"namespace,omitempty"`

	// Catalog version locator ID.
	VersionLocatorID *string `json:"version_locator_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateKubeDeployOptions : Instantiate CreateKubeDeployOptions
func (*CatalogManagementV1) NewCreateKubeDeployOptions(xAuthRefreshToken string) *CreateKubeDeployOptions {
	return &CreateKubeDeployOptions{
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *CreateKubeDeployOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *CreateKubeDeployOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *CreateKubeDeployOptions) SetClusterID(clusterID string) *CreateKubeDeployOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *CreateKubeDeployOptions) SetRegion(region string) *CreateKubeDeployOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *CreateKubeDeployOptions) SetNamespace(namespace string) *CreateKubeDeployOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *CreateKubeDeployOptions) SetVersionLocatorID(versionLocatorID string) *CreateKubeDeployOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateKubeDeployOptions) SetHeaders(param map[string]string) *CreateKubeDeployOptions {
	options.Headers = param
	return options
}

// CreateLicenseEntitlementOptions : The CreateLicenseEntitlement options.
type CreateLicenseEntitlementOptions struct {
	// Entitlement name.
	Name *string `json:"name,omitempty"`

	// entitlement is good from this starting date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveFrom *string `json:"effective_from,omitempty"`

	// entitlement is good until this ending date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveTo *string `json:"effective_to,omitempty"`

	// Global Catalog ID of the version.
	VersionID *string `json:"version_id,omitempty"`

	// specific license entitlement ID from the license provider, eg. D1W3R4.
	LicenseID *string `json:"license_id,omitempty"`

	// IBM ID of the owner of this license entitlement.
	LicenseOwnerID *string `json:"license_owner_id,omitempty"`

	// location of the registry images, eg. cp/cp4d.
	NamespaceRepository *string `json:"namespace_repository,omitempty"`

	// if not specified the token's account will be used.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateLicenseEntitlementOptions : Instantiate CreateLicenseEntitlementOptions
func (*CatalogManagementV1) NewCreateLicenseEntitlementOptions() *CreateLicenseEntitlementOptions {
	return &CreateLicenseEntitlementOptions{}
}

// SetName : Allow user to set Name
func (options *CreateLicenseEntitlementOptions) SetName(name string) *CreateLicenseEntitlementOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetEffectiveFrom : Allow user to set EffectiveFrom
func (options *CreateLicenseEntitlementOptions) SetEffectiveFrom(effectiveFrom string) *CreateLicenseEntitlementOptions {
	options.EffectiveFrom = core.StringPtr(effectiveFrom)
	return options
}

// SetEffectiveTo : Allow user to set EffectiveTo
func (options *CreateLicenseEntitlementOptions) SetEffectiveTo(effectiveTo string) *CreateLicenseEntitlementOptions {
	options.EffectiveTo = core.StringPtr(effectiveTo)
	return options
}

// SetVersionID : Allow user to set VersionID
func (options *CreateLicenseEntitlementOptions) SetVersionID(versionID string) *CreateLicenseEntitlementOptions {
	options.VersionID = core.StringPtr(versionID)
	return options
}

// SetLicenseID : Allow user to set LicenseID
func (options *CreateLicenseEntitlementOptions) SetLicenseID(licenseID string) *CreateLicenseEntitlementOptions {
	options.LicenseID = core.StringPtr(licenseID)
	return options
}

// SetLicenseOwnerID : Allow user to set LicenseOwnerID
func (options *CreateLicenseEntitlementOptions) SetLicenseOwnerID(licenseOwnerID string) *CreateLicenseEntitlementOptions {
	options.LicenseOwnerID = core.StringPtr(licenseOwnerID)
	return options
}

// SetNamespaceRepository : Allow user to set NamespaceRepository
func (options *CreateLicenseEntitlementOptions) SetNamespaceRepository(namespaceRepository string) *CreateLicenseEntitlementOptions {
	options.NamespaceRepository = core.StringPtr(namespaceRepository)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *CreateLicenseEntitlementOptions) SetAccountID(accountID string) *CreateLicenseEntitlementOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLicenseEntitlementOptions) SetHeaders(param map[string]string) *CreateLicenseEntitlementOptions {
	options.Headers = param
	return options
}

// CreateOfferingOptions : The CreateOffering options.
type CreateOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// unique id.
	NewID *string `json:"new_id,omitempty"`

	// The url for this specific offering.
	NewURL *string `json:"new_url,omitempty"`

	// The crn for this specific offering.
	NewCrn *string `json:"new_crn,omitempty"`

	// Display Name in the requested language.
	NewLabel *string `json:"new_label,omitempty"`

	// The programmatic name of this offering.
	NewName *string `json:"new_name,omitempty"`

	// URL for an icon associated with this offering.
	NewOfferingIconURL *string `json:"new_offering_icon_url,omitempty"`

	// URL for an additional docs with this offering.
	NewOfferingDocsURL *string `json:"new_offering_docs_url,omitempty"`

	// URL to be displayed in the Consumption UI for getting support on this offering.
	NewOfferingSupportURL *string `json:"new_offering_support_url,omitempty"`

	// List of tags associated with this catalog.
	NewTags []string `json:"new_tags,omitempty"`

	// Repository info for offerings.
	NewRating *Rating `json:"new_rating,omitempty"`

	// The date and time this catalog was created.
	NewCreated *strfmt.DateTime `json:"new_created,omitempty"`

	// The date and time this catalog was last updated.
	NewUpdated *strfmt.DateTime `json:"new_updated,omitempty"`

	// Short description in the requested language.
	NewShortDescription *string `json:"new_short_description,omitempty"`

	// Long description in the requested language.
	NewLongDescription *string `json:"new_long_description,omitempty"`

	// list of features associated with this offering.
	NewFeatures []Feature `json:"new_features,omitempty"`

	// Array of kind.
	NewKinds []Kind `json:"new_kinds,omitempty"`

	// Is it permitted to request publishing to IBM or Public.
	NewPermitRequestIbmPublicPublish *bool `json:"new_permit_request_ibm_public_publish,omitempty"`

	// Indicates if this offering has been approved for use by all IBMers.
	NewIbmPublishApproved *bool `json:"new_ibm_publish_approved,omitempty"`

	// Indicates if this offering has been approved for use by all IBM Cloud users.
	NewPublicPublishApproved *bool `json:"new_public_publish_approved,omitempty"`

	// The original offering CRN that this publish entry came from.
	NewPublicOriginalCrn *string `json:"new_public_original_crn,omitempty"`

	// The crn of the public catalog entry of this offering.
	NewPublishPublicCrn *string `json:"new_publish_public_crn,omitempty"`

	// The portal's approval record ID.
	NewPortalApprovalRecord *string `json:"new_portal_approval_record,omitempty"`

	// The portal UI URL.
	NewPortalUiURL *string `json:"new_portal_ui_url,omitempty"`

	// The id of the catalog containing this offering.
	NewCatalogID *string `json:"new_catalog_id,omitempty"`

	// The name of the catalog.
	NewCatalogName *string `json:"new_catalog_name,omitempty"`

	// Map of metadata values for this offering.
	NewMetadata interface{} `json:"new_metadata,omitempty"`

	// A disclaimer for this offering.
	NewDisclaimer *string `json:"new_disclaimer,omitempty"`

	// Determine if this offering should be displayed in the Consumption UI.
	NewHidden *bool `json:"new_hidden,omitempty"`

	// Provider of this offering.
	NewProvider *string `json:"new_provider,omitempty"`

	// Repository info for offerings.
	NewRepoInfo *RepoInfo `json:"new_repo_info,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateOfferingOptions : Instantiate CreateOfferingOptions
func (*CatalogManagementV1) NewCreateOfferingOptions(catalogID string) *CreateOfferingOptions {
	return &CreateOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *CreateOfferingOptions) SetCatalogID(catalogID string) *CreateOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetNewID : Allow user to set NewID
func (options *CreateOfferingOptions) SetNewID(newID string) *CreateOfferingOptions {
	options.NewID = core.StringPtr(newID)
	return options
}

// SetNewURL : Allow user to set NewURL
func (options *CreateOfferingOptions) SetNewURL(newURL string) *CreateOfferingOptions {
	options.NewURL = core.StringPtr(newURL)
	return options
}

// SetNewCrn : Allow user to set NewCrn
func (options *CreateOfferingOptions) SetNewCrn(newCrn string) *CreateOfferingOptions {
	options.NewCrn = core.StringPtr(newCrn)
	return options
}

// SetNewLabel : Allow user to set NewLabel
func (options *CreateOfferingOptions) SetNewLabel(newLabel string) *CreateOfferingOptions {
	options.NewLabel = core.StringPtr(newLabel)
	return options
}

// SetNewName : Allow user to set NewName
func (options *CreateOfferingOptions) SetNewName(newName string) *CreateOfferingOptions {
	options.NewName = core.StringPtr(newName)
	return options
}

// SetNewOfferingIconURL : Allow user to set NewOfferingIconURL
func (options *CreateOfferingOptions) SetNewOfferingIconURL(newOfferingIconURL string) *CreateOfferingOptions {
	options.NewOfferingIconURL = core.StringPtr(newOfferingIconURL)
	return options
}

// SetNewOfferingDocsURL : Allow user to set NewOfferingDocsURL
func (options *CreateOfferingOptions) SetNewOfferingDocsURL(newOfferingDocsURL string) *CreateOfferingOptions {
	options.NewOfferingDocsURL = core.StringPtr(newOfferingDocsURL)
	return options
}

// SetNewOfferingSupportURL : Allow user to set NewOfferingSupportURL
func (options *CreateOfferingOptions) SetNewOfferingSupportURL(newOfferingSupportURL string) *CreateOfferingOptions {
	options.NewOfferingSupportURL = core.StringPtr(newOfferingSupportURL)
	return options
}

// SetNewTags : Allow user to set NewTags
func (options *CreateOfferingOptions) SetNewTags(newTags []string) *CreateOfferingOptions {
	options.NewTags = newTags
	return options
}

// SetNewRating : Allow user to set NewRating
func (options *CreateOfferingOptions) SetNewRating(newRating *Rating) *CreateOfferingOptions {
	options.NewRating = newRating
	return options
}

// SetNewCreated : Allow user to set NewCreated
func (options *CreateOfferingOptions) SetNewCreated(newCreated *strfmt.DateTime) *CreateOfferingOptions {
	options.NewCreated = newCreated
	return options
}

// SetNewUpdated : Allow user to set NewUpdated
func (options *CreateOfferingOptions) SetNewUpdated(newUpdated *strfmt.DateTime) *CreateOfferingOptions {
	options.NewUpdated = newUpdated
	return options
}

// SetNewShortDescription : Allow user to set NewShortDescription
func (options *CreateOfferingOptions) SetNewShortDescription(newShortDescription string) *CreateOfferingOptions {
	options.NewShortDescription = core.StringPtr(newShortDescription)
	return options
}

// SetNewLongDescription : Allow user to set NewLongDescription
func (options *CreateOfferingOptions) SetNewLongDescription(newLongDescription string) *CreateOfferingOptions {
	options.NewLongDescription = core.StringPtr(newLongDescription)
	return options
}

// SetNewFeatures : Allow user to set NewFeatures
func (options *CreateOfferingOptions) SetNewFeatures(newFeatures []Feature) *CreateOfferingOptions {
	options.NewFeatures = newFeatures
	return options
}

// SetNewKinds : Allow user to set NewKinds
func (options *CreateOfferingOptions) SetNewKinds(newKinds []Kind) *CreateOfferingOptions {
	options.NewKinds = newKinds
	return options
}

// SetNewPermitRequestIbmPublicPublish : Allow user to set NewPermitRequestIbmPublicPublish
func (options *CreateOfferingOptions) SetNewPermitRequestIbmPublicPublish(newPermitRequestIbmPublicPublish bool) *CreateOfferingOptions {
	options.NewPermitRequestIbmPublicPublish = core.BoolPtr(newPermitRequestIbmPublicPublish)
	return options
}

// SetNewIbmPublishApproved : Allow user to set NewIbmPublishApproved
func (options *CreateOfferingOptions) SetNewIbmPublishApproved(newIbmPublishApproved bool) *CreateOfferingOptions {
	options.NewIbmPublishApproved = core.BoolPtr(newIbmPublishApproved)
	return options
}

// SetNewPublicPublishApproved : Allow user to set NewPublicPublishApproved
func (options *CreateOfferingOptions) SetNewPublicPublishApproved(newPublicPublishApproved bool) *CreateOfferingOptions {
	options.NewPublicPublishApproved = core.BoolPtr(newPublicPublishApproved)
	return options
}

// SetNewPublicOriginalCrn : Allow user to set NewPublicOriginalCrn
func (options *CreateOfferingOptions) SetNewPublicOriginalCrn(newPublicOriginalCrn string) *CreateOfferingOptions {
	options.NewPublicOriginalCrn = core.StringPtr(newPublicOriginalCrn)
	return options
}

// SetNewPublishPublicCrn : Allow user to set NewPublishPublicCrn
func (options *CreateOfferingOptions) SetNewPublishPublicCrn(newPublishPublicCrn string) *CreateOfferingOptions {
	options.NewPublishPublicCrn = core.StringPtr(newPublishPublicCrn)
	return options
}

// SetNewPortalApprovalRecord : Allow user to set NewPortalApprovalRecord
func (options *CreateOfferingOptions) SetNewPortalApprovalRecord(newPortalApprovalRecord string) *CreateOfferingOptions {
	options.NewPortalApprovalRecord = core.StringPtr(newPortalApprovalRecord)
	return options
}

// SetNewPortalUiURL : Allow user to set NewPortalUiURL
func (options *CreateOfferingOptions) SetNewPortalUiURL(newPortalUiURL string) *CreateOfferingOptions {
	options.NewPortalUiURL = core.StringPtr(newPortalUiURL)
	return options
}

// SetNewCatalogID : Allow user to set NewCatalogID
func (options *CreateOfferingOptions) SetNewCatalogID(newCatalogID string) *CreateOfferingOptions {
	options.NewCatalogID = core.StringPtr(newCatalogID)
	return options
}

// SetNewCatalogName : Allow user to set NewCatalogName
func (options *CreateOfferingOptions) SetNewCatalogName(newCatalogName string) *CreateOfferingOptions {
	options.NewCatalogName = core.StringPtr(newCatalogName)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *CreateOfferingOptions) SetNewMetadata(newMetadata interface{}) *CreateOfferingOptions {
	options.NewMetadata = newMetadata
	return options
}

// SetNewDisclaimer : Allow user to set NewDisclaimer
func (options *CreateOfferingOptions) SetNewDisclaimer(newDisclaimer string) *CreateOfferingOptions {
	options.NewDisclaimer = core.StringPtr(newDisclaimer)
	return options
}

// SetNewHidden : Allow user to set NewHidden
func (options *CreateOfferingOptions) SetNewHidden(newHidden bool) *CreateOfferingOptions {
	options.NewHidden = core.BoolPtr(newHidden)
	return options
}

// SetNewProvider : Allow user to set NewProvider
func (options *CreateOfferingOptions) SetNewProvider(newProvider string) *CreateOfferingOptions {
	options.NewProvider = core.StringPtr(newProvider)
	return options
}

// SetNewRepoInfo : Allow user to set NewRepoInfo
func (options *CreateOfferingOptions) SetNewRepoInfo(newRepoInfo *RepoInfo) *CreateOfferingOptions {
	options.NewRepoInfo = newRepoInfo
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateOfferingOptions) SetHeaders(param map[string]string) *CreateOfferingOptions {
	options.Headers = param
	return options
}

// DeleteCatalogOptions : The DeleteCatalog options.
type DeleteCatalogOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCatalogOptions : Instantiate DeleteCatalogOptions
func (*CatalogManagementV1) NewDeleteCatalogOptions(catalogID string) *DeleteCatalogOptions {
	return &DeleteCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *DeleteCatalogOptions) SetCatalogID(catalogID string) *DeleteCatalogOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCatalogOptions) SetHeaders(param map[string]string) *DeleteCatalogOptions {
	options.Headers = param
	return options
}

// DeleteKubeDeployOptions : The DeleteKubeDeploy options.
type DeleteKubeDeployOptions struct {
	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster identification.
	ClusterID *string `json:"cluster_id" validate:"required"`

	// Kube namespace.
	Namespace *string `json:"namespace" validate:"required"`

	// Kube region.
	Region *string `json:"region" validate:"required"`

	// Catalog version locator identification.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteKubeDeployOptions : Instantiate DeleteKubeDeployOptions
func (*CatalogManagementV1) NewDeleteKubeDeployOptions(xAuthRefreshToken string, clusterID string, namespace string, region string, versionLocatorID string) *DeleteKubeDeployOptions {
	return &DeleteKubeDeployOptions{
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
		ClusterID: core.StringPtr(clusterID),
		Namespace: core.StringPtr(namespace),
		Region: core.StringPtr(region),
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *DeleteKubeDeployOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *DeleteKubeDeployOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *DeleteKubeDeployOptions) SetClusterID(clusterID string) *DeleteKubeDeployOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *DeleteKubeDeployOptions) SetNamespace(namespace string) *DeleteKubeDeployOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetRegion : Allow user to set Region
func (options *DeleteKubeDeployOptions) SetRegion(region string) *DeleteKubeDeployOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *DeleteKubeDeployOptions) SetVersionLocatorID(versionLocatorID string) *DeleteKubeDeployOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteKubeDeployOptions) SetHeaders(param map[string]string) *DeleteKubeDeployOptions {
	options.Headers = param
	return options
}

// DeleteLicenseEntitlementOptions : The DeleteLicenseEntitlement options.
type DeleteLicenseEntitlementOptions struct {
	// The specific entitlement ID (can be obtained from one of the license entitlement queries).
	EntitlementID *string `json:"entitlement_id" validate:"required"`

	// The account ID to query for the entitlement. Default is the account from the user's token.
	AccountID *string `json:"account_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLicenseEntitlementOptions : Instantiate DeleteLicenseEntitlementOptions
func (*CatalogManagementV1) NewDeleteLicenseEntitlementOptions(entitlementID string) *DeleteLicenseEntitlementOptions {
	return &DeleteLicenseEntitlementOptions{
		EntitlementID: core.StringPtr(entitlementID),
	}
}

// SetEntitlementID : Allow user to set EntitlementID
func (options *DeleteLicenseEntitlementOptions) SetEntitlementID(entitlementID string) *DeleteLicenseEntitlementOptions {
	options.EntitlementID = core.StringPtr(entitlementID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *DeleteLicenseEntitlementOptions) SetAccountID(accountID string) *DeleteLicenseEntitlementOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLicenseEntitlementOptions) SetHeaders(param map[string]string) *DeleteLicenseEntitlementOptions {
	options.Headers = param
	return options
}

// DeleteOfferingOptions : The DeleteOffering options.
type DeleteOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteOfferingOptions : Instantiate DeleteOfferingOptions
func (*CatalogManagementV1) NewDeleteOfferingOptions(catalogID string, offeringID string) *DeleteOfferingOptions {
	return &DeleteOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *DeleteOfferingOptions) SetCatalogID(catalogID string) *DeleteOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *DeleteOfferingOptions) SetOfferingID(offeringID string) *DeleteOfferingOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteOfferingOptions) SetHeaders(param map[string]string) *DeleteOfferingOptions {
	options.Headers = param
	return options
}

// DeleteVersionOptions : The DeleteVersion options.
type DeleteVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteVersionOptions : Instantiate DeleteVersionOptions
func (*CatalogManagementV1) NewDeleteVersionOptions(versionLocatorID string) *DeleteVersionOptions {
	return &DeleteVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *DeleteVersionOptions) SetVersionLocatorID(versionLocatorID string) *DeleteVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVersionOptions) SetHeaders(param map[string]string) *DeleteVersionOptions {
	options.Headers = param
	return options
}

// DeployRequirementsCheck : Failed deployment requirements.
type DeployRequirementsCheck struct {
	// Failed during pre-install.
	PreInstall interface{} `json:"pre_install,omitempty"`

	// Failed during install.
	Install interface{} `json:"install,omitempty"`
}


// UnmarshalDeployRequirementsCheck unmarshals an instance of DeployRequirementsCheck from the specified map of raw messages.
func UnmarshalDeployRequirementsCheck(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeployRequirementsCheck)
	err = core.UnmarshalPrimitive(m, "pre_install", &obj.PreInstall)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "install", &obj.Install)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Deployment : Deployment for offering.
type Deployment struct {
	// unique id.
	ID *string `json:"id,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// The programmatic name of this offering.
	Name *string `json:"name,omitempty"`

	// Short description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// Long description in the requested language.
	LongDescription *string `json:"long_description,omitempty"`

	// open ended metadata information.
	Metadata interface{} `json:"metadata,omitempty"`

	// list of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// the date'time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// the date'time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`
}


// UnmarshalDeployment unmarshals an instance of Deployment from the specified map of raw messages.
func UnmarshalDeployment(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Deployment)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "long_description", &obj.LongDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeprecateVersionOptions : The DeprecateVersion options.
type DeprecateVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeprecateVersionOptions : Instantiate DeprecateVersionOptions
func (*CatalogManagementV1) NewDeprecateVersionOptions(versionLocatorID string) *DeprecateVersionOptions {
	return &DeprecateVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *DeprecateVersionOptions) SetVersionLocatorID(versionLocatorID string) *DeprecateVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeprecateVersionOptions) SetHeaders(param map[string]string) *DeprecateVersionOptions {
	options.Headers = param
	return options
}

// Enterprise : Enterprise account information.
type Enterprise struct {
	// Enterprise identification.
	ID *string `json:"id,omitempty"`

	// Filters for account and catalog filters.
	AccountFilters *Filters `json:"account_filters,omitempty"`

	// Map of account group ids to AccountGroup objects.
	AccountGroups *EnterpriseAccountGroups `json:"account_groups,omitempty"`
}


// UnmarshalEnterprise unmarshals an instance of Enterprise from the specified map of raw messages.
func UnmarshalEnterprise(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Enterprise)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account_filters", &obj.AccountFilters, UnmarshalFilters)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "account_groups", &obj.AccountGroups, UnmarshalEnterpriseAccountGroups)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EnterpriseAccountGroups : Map of account group ids to AccountGroup objects.
type EnterpriseAccountGroups struct {
	// Filters for an Account Group.
	Keys *AccountGroup `json:"keys,omitempty"`
}


// UnmarshalEnterpriseAccountGroups unmarshals an instance of EnterpriseAccountGroups from the specified map of raw messages.
func UnmarshalEnterpriseAccountGroups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EnterpriseAccountGroups)
	err = core.UnmarshalModel(m, "keys", &obj.Keys, UnmarshalAccountGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Feature : Feature information.
type Feature struct {
	// Heading.
	Title *string `json:"title,omitempty"`

	// Feature description.
	Description *string `json:"description,omitempty"`
}


// UnmarshalFeature unmarshals an instance of Feature from the specified map of raw messages.
func UnmarshalFeature(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Feature)
	err = core.UnmarshalPrimitive(m, "title", &obj.Title)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FilterTerms : Offering filter terms.
type FilterTerms struct {
	// List of values to match against. If include is true, then if the offering has one of the values then the offering is
	// included. If include is false, then if the offering has one of the values then the offering is excluded.
	FilterTerms []string `json:"filter_terms,omitempty"`
}


// UnmarshalFilterTerms unmarshals an instance of FilterTerms from the specified map of raw messages.
func UnmarshalFilterTerms(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FilterTerms)
	err = core.UnmarshalPrimitive(m, "filter_terms", &obj.FilterTerms)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Filters : Filters for account and catalog filters.
type Filters struct {
	// -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some
	// offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some
	// offerings.
	IncludeAll *bool `json:"include_all,omitempty"`

	// Filter against offering properties.
	CategoryFilters map[string]CategoryFilter `json:"category_filters,omitempty"`

	// Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
	IdFilters *IDFilter `json:"id_filters,omitempty"`
}


// UnmarshalFilters unmarshals an instance of Filters from the specified map of raw messages.
func UnmarshalFilters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Filters)
	err = core.UnmarshalPrimitive(m, "include_all", &obj.IncludeAll)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "category_filters", &obj.CategoryFilters, UnmarshalCategoryFilter)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "id_filters", &obj.IdFilters, UnmarshalIDFilter)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetCatalogAccountFiltersOptions : The GetCatalogAccountFilters options.
type GetCatalogAccountFiltersOptions struct {
	// catalog id. Narrow down filters to the account and just the one catalog.
	Catalog *string `json:"catalog,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogAccountFiltersOptions : Instantiate GetCatalogAccountFiltersOptions
func (*CatalogManagementV1) NewGetCatalogAccountFiltersOptions() *GetCatalogAccountFiltersOptions {
	return &GetCatalogAccountFiltersOptions{}
}

// SetCatalog : Allow user to set Catalog
func (options *GetCatalogAccountFiltersOptions) SetCatalog(catalog string) *GetCatalogAccountFiltersOptions {
	options.Catalog = core.StringPtr(catalog)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogAccountFiltersOptions) SetHeaders(param map[string]string) *GetCatalogAccountFiltersOptions {
	options.Headers = param
	return options
}

// GetCatalogAccountOptions : The GetCatalogAccount options.
type GetCatalogAccountOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogAccountOptions : Instantiate GetCatalogAccountOptions
func (*CatalogManagementV1) NewGetCatalogAccountOptions() *GetCatalogAccountOptions {
	return &GetCatalogAccountOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogAccountOptions) SetHeaders(param map[string]string) *GetCatalogAccountOptions {
	options.Headers = param
	return options
}

// GetCatalogOptions : The GetCatalog options.
type GetCatalogOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCatalogOptions : Instantiate GetCatalogOptions
func (*CatalogManagementV1) NewGetCatalogOptions(catalogID string) *GetCatalogOptions {
	return &GetCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *GetCatalogOptions) SetCatalogID(catalogID string) *GetCatalogOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCatalogOptions) SetHeaders(param map[string]string) *GetCatalogOptions {
	options.Headers = param
	return options
}

// GetClusterOptions : The GetCluster options.
type GetClusterOptions struct {
	// ID of the cluster.
	ClusterID *string `json:"cluster_id" validate:"required"`

	// Region of the cluster.
	Region *string `json:"region" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetClusterOptions : Instantiate GetClusterOptions
func (*CatalogManagementV1) NewGetClusterOptions(clusterID string, region string, xAuthRefreshToken string) *GetClusterOptions {
	return &GetClusterOptions{
		ClusterID: core.StringPtr(clusterID),
		Region: core.StringPtr(region),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetClusterID : Allow user to set ClusterID
func (options *GetClusterOptions) SetClusterID(clusterID string) *GetClusterOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetClusterOptions) SetRegion(region string) *GetClusterOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *GetClusterOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *GetClusterOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetClusterOptions) SetHeaders(param map[string]string) *GetClusterOptions {
	options.Headers = param
	return options
}

// GetConsumptionOfferingsOptions : The GetConsumptionOfferings options.
type GetConsumptionOfferingsOptions struct {
	// true - Strip down the content of what is returned. For example don't return the readme. Makes the result much
	// smaller. Defaults to false.
	Digest *bool `json:"digest,omitempty"`

	// catalog id. Narrow search down to just a particular catalog. It will apply the catalog's public filters to the
	// public catalog offerings on the result.
	Catalog *string `json:"catalog,omitempty"`

	// What should be selected. Default is 'all' which will return both public and private offerings. 'public' returns only
	// the public offerings and 'private' returns only the private offerings.
	Select *string `json:"select,omitempty"`

	// true - include offerings which have been marked as hidden. The default is false and hidden offerings are not
	// returned.
	IncludeHidden *bool `json:"includeHidden,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetConsumptionOfferingsOptions.Select property.
// What should be selected. Default is 'all' which will return both public and private offerings. 'public' returns only
// the public offerings and 'private' returns only the private offerings.
const (
	GetConsumptionOfferingsOptions_Select_All = "all"
	GetConsumptionOfferingsOptions_Select_Private = "private"
	GetConsumptionOfferingsOptions_Select_Public = "public"
)

// NewGetConsumptionOfferingsOptions : Instantiate GetConsumptionOfferingsOptions
func (*CatalogManagementV1) NewGetConsumptionOfferingsOptions() *GetConsumptionOfferingsOptions {
	return &GetConsumptionOfferingsOptions{}
}

// SetDigest : Allow user to set Digest
func (options *GetConsumptionOfferingsOptions) SetDigest(digest bool) *GetConsumptionOfferingsOptions {
	options.Digest = core.BoolPtr(digest)
	return options
}

// SetCatalog : Allow user to set Catalog
func (options *GetConsumptionOfferingsOptions) SetCatalog(catalog string) *GetConsumptionOfferingsOptions {
	options.Catalog = core.StringPtr(catalog)
	return options
}

// SetSelect : Allow user to set Select
func (options *GetConsumptionOfferingsOptions) SetSelect(selectVar string) *GetConsumptionOfferingsOptions {
	options.Select = core.StringPtr(selectVar)
	return options
}

// SetIncludeHidden : Allow user to set IncludeHidden
func (options *GetConsumptionOfferingsOptions) SetIncludeHidden(includeHidden bool) *GetConsumptionOfferingsOptions {
	options.IncludeHidden = core.BoolPtr(includeHidden)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetConsumptionOfferingsOptions) SetHeaders(param map[string]string) *GetConsumptionOfferingsOptions {
	options.Headers = param
	return options
}

// GetEnterpriseOptions : The GetEnterprise options.
type GetEnterpriseOptions struct {
	// Enterprise identification.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetEnterpriseOptions : Instantiate GetEnterpriseOptions
func (*CatalogManagementV1) NewGetEnterpriseOptions(enterpriseID string) *GetEnterpriseOptions {
	return &GetEnterpriseOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *GetEnterpriseOptions) SetEnterpriseID(enterpriseID string) *GetEnterpriseOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetEnterpriseOptions) SetHeaders(param map[string]string) *GetEnterpriseOptions {
	options.Headers = param
	return options
}

// GetLicenseEntitlementsOptions : The GetLicenseEntitlements options.
type GetLicenseEntitlementsOptions struct {
	// The license product ID. If from PPA (Passport Advantage) this is a specific product Part number, eg. D1YGZLL.
	LicenseProductID *string `json:"license_product_id" validate:"required"`

	// The account ID to query for the entitlement. Default is the account from the user's token.
	AccountID *string `json:"account_id,omitempty"`

	// The GC ID of the specific offering version.
	VersionID *string `json:"version_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLicenseEntitlementsOptions : Instantiate GetLicenseEntitlementsOptions
func (*CatalogManagementV1) NewGetLicenseEntitlementsOptions(licenseProductID string) *GetLicenseEntitlementsOptions {
	return &GetLicenseEntitlementsOptions{
		LicenseProductID: core.StringPtr(licenseProductID),
	}
}

// SetLicenseProductID : Allow user to set LicenseProductID
func (options *GetLicenseEntitlementsOptions) SetLicenseProductID(licenseProductID string) *GetLicenseEntitlementsOptions {
	options.LicenseProductID = core.StringPtr(licenseProductID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *GetLicenseEntitlementsOptions) SetAccountID(accountID string) *GetLicenseEntitlementsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetVersionID : Allow user to set VersionID
func (options *GetLicenseEntitlementsOptions) SetVersionID(versionID string) *GetLicenseEntitlementsOptions {
	options.VersionID = core.StringPtr(versionID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLicenseEntitlementsOptions) SetHeaders(param map[string]string) *GetLicenseEntitlementsOptions {
	options.Headers = param
	return options
}

// GetLicenseProvidersOptions : The GetLicenseProviders options.
type GetLicenseProvidersOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLicenseProvidersOptions : Instantiate GetLicenseProvidersOptions
func (*CatalogManagementV1) NewGetLicenseProvidersOptions() *GetLicenseProvidersOptions {
	return &GetLicenseProvidersOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetLicenseProvidersOptions) SetHeaders(param map[string]string) *GetLicenseProvidersOptions {
	options.Headers = param
	return options
}

// GetLicensesOptions : The GetLicenses options.
type GetLicensesOptions struct {
	// ID of the license provider, ie. retrieved from GET license_providers.
	LicenseProviderID *string `json:"license_provider_id" validate:"required"`

	// If not specified the token's account will be used.
	AccountID *string `json:"account_id,omitempty"`

	// License name.
	Name *string `json:"name,omitempty"`

	// Type of license, if not specified, default is ibm-ppa.
	LicenseType *string `json:"license_type,omitempty"`

	// The license product ID. If from PPA (Passport Advantage) this is the product Part number, eg. D1YGZLL.
	LicenseProductID *string `json:"license_product_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLicensesOptions : Instantiate GetLicensesOptions
func (*CatalogManagementV1) NewGetLicensesOptions(licenseProviderID string) *GetLicensesOptions {
	return &GetLicensesOptions{
		LicenseProviderID: core.StringPtr(licenseProviderID),
	}
}

// SetLicenseProviderID : Allow user to set LicenseProviderID
func (options *GetLicensesOptions) SetLicenseProviderID(licenseProviderID string) *GetLicensesOptions {
	options.LicenseProviderID = core.StringPtr(licenseProviderID)
	return options
}

// SetAccountID : Allow user to set AccountID
func (options *GetLicensesOptions) SetAccountID(accountID string) *GetLicensesOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetName : Allow user to set Name
func (options *GetLicensesOptions) SetName(name string) *GetLicensesOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetLicenseType : Allow user to set LicenseType
func (options *GetLicensesOptions) SetLicenseType(licenseType string) *GetLicensesOptions {
	options.LicenseType = core.StringPtr(licenseType)
	return options
}

// SetLicenseProductID : Allow user to set LicenseProductID
func (options *GetLicensesOptions) SetLicenseProductID(licenseProductID string) *GetLicensesOptions {
	options.LicenseProductID = core.StringPtr(licenseProductID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetLicensesOptions) SetHeaders(param map[string]string) *GetLicensesOptions {
	options.Headers = param
	return options
}

// GetNamespacesOptions : The GetNamespaces options.
type GetNamespacesOptions struct {
	// ID of the cluster.
	ClusterID *string `json:"cluster_id" validate:"required"`

	// Cluster region.
	Region *string `json:"region" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// number or results to return.
	Limit *int64 `json:"limit,omitempty"`

	// number of results to skip before returning values.
	Offset *int64 `json:"offset,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetNamespacesOptions : Instantiate GetNamespacesOptions
func (*CatalogManagementV1) NewGetNamespacesOptions(clusterID string, region string, xAuthRefreshToken string) *GetNamespacesOptions {
	return &GetNamespacesOptions{
		ClusterID: core.StringPtr(clusterID),
		Region: core.StringPtr(region),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetClusterID : Allow user to set ClusterID
func (options *GetNamespacesOptions) SetClusterID(clusterID string) *GetNamespacesOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetNamespacesOptions) SetRegion(region string) *GetNamespacesOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *GetNamespacesOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *GetNamespacesOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetLimit : Allow user to set Limit
func (options *GetNamespacesOptions) SetLimit(limit int64) *GetNamespacesOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *GetNamespacesOptions) SetOffset(offset int64) *GetNamespacesOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetNamespacesOptions) SetHeaders(param map[string]string) *GetNamespacesOptions {
	options.Headers = param
	return options
}

// GetOfferingOptions : The GetOffering options.
type GetOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetOfferingOptions : Instantiate GetOfferingOptions
func (*CatalogManagementV1) NewGetOfferingOptions(catalogID string, offeringID string) *GetOfferingOptions {
	return &GetOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *GetOfferingOptions) SetCatalogID(catalogID string) *GetOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *GetOfferingOptions) SetOfferingID(offeringID string) *GetOfferingOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetOfferingOptions) SetHeaders(param map[string]string) *GetOfferingOptions {
	options.Headers = param
	return options
}

// GetOverrideValuesOptions : The GetOverrideValues options.
type GetOverrideValuesOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetOverrideValuesOptions : Instantiate GetOverrideValuesOptions
func (*CatalogManagementV1) NewGetOverrideValuesOptions(versionLocatorID string) *GetOverrideValuesOptions {
	return &GetOverrideValuesOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetOverrideValuesOptions) SetVersionLocatorID(versionLocatorID string) *GetOverrideValuesOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetOverrideValuesOptions) SetHeaders(param map[string]string) *GetOverrideValuesOptions {
	options.Headers = param
	return options
}

// GetPreinstallOptions : The GetPreinstall options.
type GetPreinstallOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// ID of the cluster.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Cluster region.
	Region *string `json:"region,omitempty"`

	// Required if the version's pre-install scope is `namespace`.
	Namespace *string `json:"namespace,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPreinstallOptions : Instantiate GetPreinstallOptions
func (*CatalogManagementV1) NewGetPreinstallOptions(versionLocatorID string, xAuthRefreshToken string) *GetPreinstallOptions {
	return &GetPreinstallOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetPreinstallOptions) SetVersionLocatorID(versionLocatorID string) *GetPreinstallOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *GetPreinstallOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *GetPreinstallOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *GetPreinstallOptions) SetClusterID(clusterID string) *GetPreinstallOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetPreinstallOptions) SetRegion(region string) *GetPreinstallOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *GetPreinstallOptions) SetNamespace(namespace string) *GetPreinstallOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetPreinstallOptions) SetHeaders(param map[string]string) *GetPreinstallOptions {
	options.Headers = param
	return options
}

// GetRepoOptions : The GetRepo options.
type GetRepoOptions struct {
	// The type of repo (valid repo types: helm).
	Type *string `json:"type" validate:"required"`

	// The URL for the repo's chart zip file (e.g
	// https://registry.bluemix.net/helm/ibm-charts/charts/ibm-redis-ha-dev-1.0.0.tgz).
	Charturl *string `json:"charturl" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRepoOptions : Instantiate GetRepoOptions
func (*CatalogManagementV1) NewGetRepoOptions(typeVar string, charturl string) *GetRepoOptions {
	return &GetRepoOptions{
		Type: core.StringPtr(typeVar),
		Charturl: core.StringPtr(charturl),
	}
}

// SetType : Allow user to set Type
func (options *GetRepoOptions) SetType(typeVar string) *GetRepoOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetCharturl : Allow user to set Charturl
func (options *GetRepoOptions) SetCharturl(charturl string) *GetRepoOptions {
	options.Charturl = core.StringPtr(charturl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRepoOptions) SetHeaders(param map[string]string) *GetRepoOptions {
	options.Headers = param
	return options
}

// GetReposOptions : The GetRepos options.
type GetReposOptions struct {
	// The type of repo (valid repo types: helm).
	Type *string `json:"type" validate:"required"`

	// The URL for the repo's root (e.g https://kubernetes-charts-incubator.storage.googleapis.com).
	Repourl *string `json:"repourl" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetReposOptions : Instantiate GetReposOptions
func (*CatalogManagementV1) NewGetReposOptions(typeVar string, repourl string) *GetReposOptions {
	return &GetReposOptions{
		Type: core.StringPtr(typeVar),
		Repourl: core.StringPtr(repourl),
	}
}

// SetType : Allow user to set Type
func (options *GetReposOptions) SetType(typeVar string) *GetReposOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetRepourl : Allow user to set Repourl
func (options *GetReposOptions) SetRepourl(repourl string) *GetReposOptions {
	options.Repourl = core.StringPtr(repourl)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetReposOptions) SetHeaders(param map[string]string) *GetReposOptions {
	options.Headers = param
	return options
}

// GetResourceGroupsOptions : The GetResourceGroups options.
type GetResourceGroupsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceGroupsOptions : Instantiate GetResourceGroupsOptions
func (*CatalogManagementV1) NewGetResourceGroupsOptions() *GetResourceGroupsOptions {
	return &GetResourceGroupsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceGroupsOptions) SetHeaders(param map[string]string) *GetResourceGroupsOptions {
	options.Headers = param
	return options
}

// GetSchematicsWorkspacesOptions : The GetSchematicsWorkspaces options.
type GetSchematicsWorkspacesOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSchematicsWorkspacesOptions : Instantiate GetSchematicsWorkspacesOptions
func (*CatalogManagementV1) NewGetSchematicsWorkspacesOptions(versionLocatorID string, xAuthRefreshToken string) *GetSchematicsWorkspacesOptions {
	return &GetSchematicsWorkspacesOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetSchematicsWorkspacesOptions) SetVersionLocatorID(versionLocatorID string) *GetSchematicsWorkspacesOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *GetSchematicsWorkspacesOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *GetSchematicsWorkspacesOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetSchematicsWorkspacesOptions) SetHeaders(param map[string]string) *GetSchematicsWorkspacesOptions {
	options.Headers = param
	return options
}

// GetValidationStatusOptions : The GetValidationStatus options.
type GetValidationStatusOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetValidationStatusOptions : Instantiate GetValidationStatusOptions
func (*CatalogManagementV1) NewGetValidationStatusOptions(versionLocatorID string, xAuthRefreshToken string) *GetValidationStatusOptions {
	return &GetValidationStatusOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetValidationStatusOptions) SetVersionLocatorID(versionLocatorID string) *GetValidationStatusOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *GetValidationStatusOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *GetValidationStatusOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetValidationStatusOptions) SetHeaders(param map[string]string) *GetValidationStatusOptions {
	options.Headers = param
	return options
}

// GetVersionAboutOptions : The GetVersionAbout options.
type GetVersionAboutOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionAboutOptions : Instantiate GetVersionAboutOptions
func (*CatalogManagementV1) NewGetVersionAboutOptions(versionLocatorID string) *GetVersionAboutOptions {
	return &GetVersionAboutOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionAboutOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionAboutOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionAboutOptions) SetHeaders(param map[string]string) *GetVersionAboutOptions {
	options.Headers = param
	return options
}

// GetVersionContainerImagesOptions : The GetVersionContainerImages options.
type GetVersionContainerImagesOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionContainerImagesOptions : Instantiate GetVersionContainerImagesOptions
func (*CatalogManagementV1) NewGetVersionContainerImagesOptions(versionLocatorID string) *GetVersionContainerImagesOptions {
	return &GetVersionContainerImagesOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionContainerImagesOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionContainerImagesOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionContainerImagesOptions) SetHeaders(param map[string]string) *GetVersionContainerImagesOptions {
	options.Headers = param
	return options
}

// GetVersionLicenseOptions : The GetVersionLicense options.
type GetVersionLicenseOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// The ID of the license, which maps to the file name in the 'licenses' directory of this verions tgz file.
	LicenseID *string `json:"license_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionLicenseOptions : Instantiate GetVersionLicenseOptions
func (*CatalogManagementV1) NewGetVersionLicenseOptions(versionLocatorID string, licenseID string) *GetVersionLicenseOptions {
	return &GetVersionLicenseOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		LicenseID: core.StringPtr(licenseID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionLicenseOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionLicenseOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetLicenseID : Allow user to set LicenseID
func (options *GetVersionLicenseOptions) SetLicenseID(licenseID string) *GetVersionLicenseOptions {
	options.LicenseID = core.StringPtr(licenseID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionLicenseOptions) SetHeaders(param map[string]string) *GetVersionLicenseOptions {
	options.Headers = param
	return options
}

// GetVersionOptions : The GetVersion options.
type GetVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionOptions : Instantiate GetVersionOptions
func (*CatalogManagementV1) NewGetVersionOptions(versionLocatorID string) *GetVersionOptions {
	return &GetVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionOptions) SetHeaders(param map[string]string) *GetVersionOptions {
	options.Headers = param
	return options
}

// GetVersionUpdatesOptions : The GetVersionUpdates options.
type GetVersionUpdatesOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// The id of the cluster where this version was installed.
	ClusterID *string `json:"cluster_id,omitempty"`

	// The region of the cluster where this version was installed.
	Region *string `json:"region,omitempty"`

	// The resource group id of the cluster where this version was installed.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The namespace of the cluster where this version was installed.
	Namespace *string `json:"namespace,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionUpdatesOptions : Instantiate GetVersionUpdatesOptions
func (*CatalogManagementV1) NewGetVersionUpdatesOptions(versionLocatorID string) *GetVersionUpdatesOptions {
	return &GetVersionUpdatesOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionUpdatesOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionUpdatesOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *GetVersionUpdatesOptions) SetClusterID(clusterID string) *GetVersionUpdatesOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *GetVersionUpdatesOptions) SetRegion(region string) *GetVersionUpdatesOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *GetVersionUpdatesOptions) SetResourceGroupID(resourceGroupID string) *GetVersionUpdatesOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *GetVersionUpdatesOptions) SetNamespace(namespace string) *GetVersionUpdatesOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionUpdatesOptions) SetHeaders(param map[string]string) *GetVersionUpdatesOptions {
	options.Headers = param
	return options
}

// GetVersionWorkingCopyOptions : The GetVersionWorkingCopy options.
type GetVersionWorkingCopyOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVersionWorkingCopyOptions : Instantiate GetVersionWorkingCopyOptions
func (*CatalogManagementV1) NewGetVersionWorkingCopyOptions(versionLocatorID string) *GetVersionWorkingCopyOptions {
	return &GetVersionWorkingCopyOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *GetVersionWorkingCopyOptions) SetVersionLocatorID(versionLocatorID string) *GetVersionWorkingCopyOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetVersionWorkingCopyOptions) SetHeaders(param map[string]string) *GetVersionWorkingCopyOptions {
	options.Headers = param
	return options
}

// HelmChart : Helm chart.
type HelmChart struct {
	// Chart name.
	Name *string `json:"name,omitempty"`

	// Chart description.
	Description *string `json:"description,omitempty"`

	// Chart icon.
	Icon *string `json:"icon,omitempty"`

	// Chart version.
	Version *string `json:"version,omitempty"`

	// Chart app version.
	AppVersion *string `json:"appVersion,omitempty"`
}


// UnmarshalHelmChart unmarshals an instance of HelmChart from the specified map of raw messages.
func UnmarshalHelmChart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HelmChart)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icon", &obj.Icon)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appVersion", &obj.AppVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HelmPackage : Helm package.
type HelmPackage struct {
	// The name of the requested chart, or the name of a nested chart within the requested chart.
	Chart *HelmPackageChart `json:"chart,omitempty"`
}


// UnmarshalHelmPackage unmarshals an instance of HelmPackage from the specified map of raw messages.
func UnmarshalHelmPackage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HelmPackage)
	err = core.UnmarshalModel(m, "chart", &obj.Chart, UnmarshalHelmPackageChart)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HelmPackageChart : The name of the requested chart, or the name of a nested chart within the requested chart.
type HelmPackageChart struct {
	// Helm chart.
	ChartYaml *HelmChart `json:"Chart.yaml,omitempty"`

	// Project SHA.
	Sha interface{} `json:"sha,omitempty"`

	// Helm chart description.
	READMEMd *string `json:"README.md,omitempty"`

	// Values metadata.
	ValuesMetadata interface{} `json:"values-metadata,omitempty"`

	// License metadata.
	LicenseMetadata interface{} `json:"license-metadata,omitempty"`
}


// UnmarshalHelmPackageChart unmarshals an instance of HelmPackageChart from the specified map of raw messages.
func UnmarshalHelmPackageChart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HelmPackageChart)
	err = core.UnmarshalModel(m, "Chart.yaml", &obj.ChartYaml, UnmarshalHelmChart)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sha", &obj.Sha)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "README.md", &obj.READMEMd)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "values-metadata", &obj.ValuesMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license-metadata", &obj.LicenseMetadata)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HelmRepoList : Helm repository metadata.
type HelmRepoList struct {
	// A chart entry in the repo. This response will contain many chart names.
	Chart *HelmRepoListChart `json:"chart,omitempty"`
}


// UnmarshalHelmRepoList unmarshals an instance of HelmRepoList from the specified map of raw messages.
func UnmarshalHelmRepoList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HelmRepoList)
	err = core.UnmarshalModel(m, "chart", &obj.Chart, UnmarshalHelmRepoListChart)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HelmRepoListChart : A chart entry in the repo. This response will contain many chart names.
type HelmRepoListChart struct {
	// API version.
	ApiVersion *string `json:"api_version,omitempty"`

	// Date and time created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Description of Helm repo entry.
	Description *string `json:"description,omitempty"`

	// Denotes whether repo entry is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`

	// Digest of entry.
	Digest *string `json:"digest,omitempty"`

	// Location of repo entry.
	Home *string `json:"home,omitempty"`

	// Entry icon.
	Icon *string `json:"icon,omitempty"`

	// List of keywords.
	Keywords []string `json:"keywords,omitempty"`

	// Emails and names of repo maintainers.
	Maintainers []Maintainers `json:"maintainers,omitempty"`

	// Entry name.
	Name *string `json:"name,omitempty"`

	// Helm server version.
	TillerVersion *string `json:"tiller_version,omitempty"`

	// Array of URLs.
	Urls []string `json:"urls,omitempty"`

	// Array of sources.
	Sources []string `json:"sources,omitempty"`

	// Entry version.
	Version *string `json:"version,omitempty"`

	// Application version.
	AppVersion *string `json:"appVersion,omitempty"`
}


// UnmarshalHelmRepoListChart unmarshals an instance of HelmRepoListChart from the specified map of raw messages.
func UnmarshalHelmRepoListChart(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HelmRepoListChart)
	err = core.UnmarshalPrimitive(m, "api_version", &obj.ApiVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deprecated", &obj.Deprecated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "digest", &obj.Digest)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "home", &obj.Home)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "icon", &obj.Icon)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "keywords", &obj.Keywords)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "maintainers", &obj.Maintainers, UnmarshalMaintainers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tiller_version", &obj.TillerVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "urls", &obj.Urls)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sources", &obj.Sources)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "appVersion", &obj.AppVersion)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IDFilter : Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
type IDFilter struct {
	// Offering filter terms.
	Include *FilterTerms `json:"include,omitempty"`

	// Offering filter terms.
	Exclude *FilterTerms `json:"exclude,omitempty"`
}


// UnmarshalIDFilter unmarshals an instance of IDFilter from the specified map of raw messages.
func UnmarshalIDFilter(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(IDFilter)
	err = core.UnmarshalModel(m, "include", &obj.Include, UnmarshalFilterTerms)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "exclude", &obj.Exclude, UnmarshalFilterTerms)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// IbmPublishVersionOptions : The IbmPublishVersion options.
type IbmPublishVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewIbmPublishVersionOptions : Instantiate IbmPublishVersionOptions
func (*CatalogManagementV1) NewIbmPublishVersionOptions(versionLocatorID string) *IbmPublishVersionOptions {
	return &IbmPublishVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *IbmPublishVersionOptions) SetVersionLocatorID(versionLocatorID string) *IbmPublishVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *IbmPublishVersionOptions) SetHeaders(param map[string]string) *IbmPublishVersionOptions {
	options.Headers = param
	return options
}

// Image : Image.
type Image struct {
	// Image.
	Image *string `json:"image,omitempty"`
}


// UnmarshalImage unmarshals an instance of Image from the specified map of raw messages.
func UnmarshalImage(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Image)
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImageManifest : Image Manifest.
type ImageManifest struct {
	// Image manifest description.
	Description *string `json:"description,omitempty"`

	// List of images.
	Images []Image `json:"images,omitempty"`
}


// UnmarshalImageManifest unmarshals an instance of ImageManifest from the specified map of raw messages.
func UnmarshalImageManifest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImageManifest)
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "images", &obj.Images, UnmarshalImage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportOfferingOptions : The ImportOffering options.
type ImportOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// URL path to zip location.
	Zipurl *string `json:"zipurl" validate:"required"`

	// Re-use the specified offeringID during import.
	OfferingID *string `json:"offeringID,omitempty"`

	// Add all possible configuration items when creating this version.
	IncludeConfig *bool `json:"includeConfig,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportOfferingOptions : Instantiate ImportOfferingOptions
func (*CatalogManagementV1) NewImportOfferingOptions(catalogID string, zipurl string) *ImportOfferingOptions {
	return &ImportOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
		Zipurl: core.StringPtr(zipurl),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ImportOfferingOptions) SetCatalogID(catalogID string) *ImportOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetZipurl : Allow user to set Zipurl
func (options *ImportOfferingOptions) SetZipurl(zipurl string) *ImportOfferingOptions {
	options.Zipurl = core.StringPtr(zipurl)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *ImportOfferingOptions) SetOfferingID(offeringID string) *ImportOfferingOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetIncludeConfig : Allow user to set IncludeConfig
func (options *ImportOfferingOptions) SetIncludeConfig(includeConfig bool) *ImportOfferingOptions {
	options.IncludeConfig = core.BoolPtr(includeConfig)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportOfferingOptions) SetHeaders(param map[string]string) *ImportOfferingOptions {
	options.Headers = param
	return options
}

// ImportOfferingVersionOptions : The ImportOfferingVersion options.
type ImportOfferingVersionOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// URL path to zip location.
	Zipurl *string `json:"zipurl" validate:"required"`

	// The semver value for this new version, if not found in the zip url package content.
	TargetVersion *string `json:"targetVersion,omitempty"`

	// Add all possible configuration values to this version when importing.
	IncludeConfig *bool `json:"includeConfig,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportOfferingVersionOptions : Instantiate ImportOfferingVersionOptions
func (*CatalogManagementV1) NewImportOfferingVersionOptions(catalogID string, offeringID string, zipurl string) *ImportOfferingVersionOptions {
	return &ImportOfferingVersionOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
		Zipurl: core.StringPtr(zipurl),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ImportOfferingVersionOptions) SetCatalogID(catalogID string) *ImportOfferingVersionOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *ImportOfferingVersionOptions) SetOfferingID(offeringID string) *ImportOfferingVersionOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetZipurl : Allow user to set Zipurl
func (options *ImportOfferingVersionOptions) SetZipurl(zipurl string) *ImportOfferingVersionOptions {
	options.Zipurl = core.StringPtr(zipurl)
	return options
}

// SetTargetVersion : Allow user to set TargetVersion
func (options *ImportOfferingVersionOptions) SetTargetVersion(targetVersion string) *ImportOfferingVersionOptions {
	options.TargetVersion = core.StringPtr(targetVersion)
	return options
}

// SetIncludeConfig : Allow user to set IncludeConfig
func (options *ImportOfferingVersionOptions) SetIncludeConfig(includeConfig bool) *ImportOfferingVersionOptions {
	options.IncludeConfig = core.BoolPtr(includeConfig)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ImportOfferingVersionOptions) SetHeaders(param map[string]string) *ImportOfferingVersionOptions {
	options.Headers = param
	return options
}

// InstallStatus : Installation status.
type InstallStatus struct {
	// Installation status metadata.
	Metadata *InstallStatusMetadata `json:"metadata,omitempty"`

	// Release information.
	Release *InstallStatusRelease `json:"release,omitempty"`

	// Content management information.
	ContentMgmt *InstallStatusContentMgmt `json:"content_mgmt,omitempty"`
}


// UnmarshalInstallStatus unmarshals an instance of InstallStatus from the specified map of raw messages.
func UnmarshalInstallStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstallStatus)
	err = core.UnmarshalModel(m, "metadata", &obj.Metadata, UnmarshalInstallStatusMetadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "release", &obj.Release, UnmarshalInstallStatusRelease)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "content_mgmt", &obj.ContentMgmt, UnmarshalInstallStatusContentMgmt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallStatusContentMgmt : Content management information.
type InstallStatusContentMgmt struct {
	// Pods.
	Pods []interface{} `json:"pods,omitempty"`

	// Errors.
	Errors []interface{} `json:"errors,omitempty"`
}


// UnmarshalInstallStatusContentMgmt unmarshals an instance of InstallStatusContentMgmt from the specified map of raw messages.
func UnmarshalInstallStatusContentMgmt(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstallStatusContentMgmt)
	err = core.UnmarshalPrimitive(m, "pods", &obj.Pods)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallStatusMetadata : Installation status metadata.
type InstallStatusMetadata struct {
	// Cluster ID.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Cluster region.
	Region *string `json:"region,omitempty"`

	// Cluster namespace.
	Namespace *string `json:"namespace,omitempty"`

	// Workspace ID.
	WorkspaceID *string `json:"workspace_id,omitempty"`

	// Workspace name.
	WorkspaceName *string `json:"workspace_name,omitempty"`
}


// UnmarshalInstallStatusMetadata unmarshals an instance of InstallStatusMetadata from the specified map of raw messages.
func UnmarshalInstallStatusMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstallStatusMetadata)
	err = core.UnmarshalPrimitive(m, "cluster_id", &obj.ClusterID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "namespace", &obj.Namespace)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_id", &obj.WorkspaceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "workspace_name", &obj.WorkspaceName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallStatusRelease : Release information.
type InstallStatusRelease struct {
	// Kube deployments.
	Deployments []interface{} `json:"deployments,omitempty"`

	// Kube replica sets.
	Replicasets []interface{} `json:"replicasets,omitempty"`

	// Kube stateful sets.
	Statefulsets []interface{} `json:"statefulsets,omitempty"`

	// Kube pods.
	Pods []interface{} `json:"pods,omitempty"`

	// Kube errors.
	Errors []interface{} `json:"errors,omitempty"`
}


// UnmarshalInstallStatusRelease unmarshals an instance of InstallStatusRelease from the specified map of raw messages.
func UnmarshalInstallStatusRelease(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(InstallStatusRelease)
	err = core.UnmarshalPrimitive(m, "deployments", &obj.Deployments)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "replicasets", &obj.Replicasets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "statefulsets", &obj.Statefulsets)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pods", &obj.Pods)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// InstallVersionOptions : The InstallVersion options.
type InstallVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster ID.
	NewClusterID *string `json:"new_cluster_id,omitempty"`

	// Cluster region.
	NewRegion *string `json:"new_region,omitempty"`

	// Kube namespace.
	NewNamespace *string `json:"new_namespace,omitempty"`

	// Object containing Helm chart override values.
	NewOverrideValues interface{} `json:"new_override_values,omitempty"`

	// Entitlement API Key for this offering.
	NewEntitlementApikey *string `json:"new_entitlement_apikey,omitempty"`

	// Schematics workspace configuration.
	NewSchematics *ClusterRequestBodySchematics `json:"new_schematics,omitempty"`

	// Script.
	NewScript *string `json:"new_script,omitempty"`

	// Script ID.
	NewScriptID *string `json:"new_script_id,omitempty"`

	// Version locator ID.
	NewVersionLocatorID *string `json:"new_version_locator_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewInstallVersionOptions : Instantiate InstallVersionOptions
func (*CatalogManagementV1) NewInstallVersionOptions(versionLocatorID string, xAuthRefreshToken string) *InstallVersionOptions {
	return &InstallVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *InstallVersionOptions) SetVersionLocatorID(versionLocatorID string) *InstallVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *InstallVersionOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *InstallVersionOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetNewClusterID : Allow user to set NewClusterID
func (options *InstallVersionOptions) SetNewClusterID(newClusterID string) *InstallVersionOptions {
	options.NewClusterID = core.StringPtr(newClusterID)
	return options
}

// SetNewRegion : Allow user to set NewRegion
func (options *InstallVersionOptions) SetNewRegion(newRegion string) *InstallVersionOptions {
	options.NewRegion = core.StringPtr(newRegion)
	return options
}

// SetNewNamespace : Allow user to set NewNamespace
func (options *InstallVersionOptions) SetNewNamespace(newNamespace string) *InstallVersionOptions {
	options.NewNamespace = core.StringPtr(newNamespace)
	return options
}

// SetNewOverrideValues : Allow user to set NewOverrideValues
func (options *InstallVersionOptions) SetNewOverrideValues(newOverrideValues interface{}) *InstallVersionOptions {
	options.NewOverrideValues = newOverrideValues
	return options
}

// SetNewEntitlementApikey : Allow user to set NewEntitlementApikey
func (options *InstallVersionOptions) SetNewEntitlementApikey(newEntitlementApikey string) *InstallVersionOptions {
	options.NewEntitlementApikey = core.StringPtr(newEntitlementApikey)
	return options
}

// SetNewSchematics : Allow user to set NewSchematics
func (options *InstallVersionOptions) SetNewSchematics(newSchematics *ClusterRequestBodySchematics) *InstallVersionOptions {
	options.NewSchematics = newSchematics
	return options
}

// SetNewScript : Allow user to set NewScript
func (options *InstallVersionOptions) SetNewScript(newScript string) *InstallVersionOptions {
	options.NewScript = core.StringPtr(newScript)
	return options
}

// SetNewScriptID : Allow user to set NewScriptID
func (options *InstallVersionOptions) SetNewScriptID(newScriptID string) *InstallVersionOptions {
	options.NewScriptID = core.StringPtr(newScriptID)
	return options
}

// SetNewVersionLocatorID : Allow user to set NewVersionLocatorID
func (options *InstallVersionOptions) SetNewVersionLocatorID(newVersionLocatorID string) *InstallVersionOptions {
	options.NewVersionLocatorID = core.StringPtr(newVersionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *InstallVersionOptions) SetHeaders(param map[string]string) *InstallVersionOptions {
	options.Headers = param
	return options
}

// Kind : Offering kind.
type Kind struct {
	// Unique ID.
	ID *string `json:"id,omitempty"`

	// content kind, e.g., helm, vm image.
	FormatKind *string `json:"format_kind,omitempty"`

	// target cloud to install, e.g., iks, open_shift_iks.
	TargetKind *string `json:"target_kind,omitempty"`

	// Open ended metadata information.
	Metadata interface{} `json:"metadata,omitempty"`

	// Installation instruction.
	InstallDescription *string `json:"install_description,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// List of features associated with this offering.
	AdditionalFeatures []Feature `json:"additional_features,omitempty"`

	// The date and time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date and time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// list of versions.
	Versions []Version `json:"versions,omitempty"`

	// list of plans.
	Plans []Plan `json:"plans,omitempty"`
}


// UnmarshalKind unmarshals an instance of Kind from the specified map of raw messages.
func UnmarshalKind(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Kind)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "format_kind", &obj.FormatKind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target_kind", &obj.TargetKind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "install_description", &obj.InstallDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "additional_features", &obj.AdditionalFeatures, UnmarshalFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "versions", &obj.Versions, UnmarshalVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "plans", &obj.Plans, UnmarshalPlan)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// KubeDeployResult : Kube deploy result.
type KubeDeployResult struct {
	// Status phase.
	Phase *string `json:"phase,omitempty"`

	// Status message.
	Message *string `json:"message,omitempty"`

	// Kube API path.
	Link *string `json:"link,omitempty"`

	// Name of deployment.
	Name *string `json:"name,omitempty"`

	// Deployment version.
	Version *string `json:"version,omitempty"`

	// Kube namespace.
	Namespace *string `json:"namespace,omitempty"`

	// Package deployment exists in.
	PackageName *string `json:"package_name,omitempty"`

	// Catalog identification.
	CatalogID *string `json:"catalog_id,omitempty"`
}


// UnmarshalKubeDeployResult unmarshals an instance of KubeDeployResult from the specified map of raw messages.
func UnmarshalKubeDeployResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(KubeDeployResult)
	err = core.UnmarshalPrimitive(m, "phase", &obj.Phase)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "link", &obj.Link)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "namespace", &obj.Namespace)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "package_name", &obj.PackageName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// License : BSS license.
type License struct {
	// License ID.
	ID *string `json:"id,omitempty"`

	// license name.
	Name *string `json:"name,omitempty"`

	// type of license e.g., Apache xxx.
	Type *string `json:"type,omitempty"`

	// URL for the license text.
	URL *string `json:"url,omitempty"`

	// License description.
	Description *string `json:"description,omitempty"`
}


// UnmarshalLicense unmarshals an instance of License from the specified map of raw messages.
func UnmarshalLicense(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(License)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseEntitlement : License entitlement.
type LicenseEntitlement struct {
	// Entitlement name.
	Name *string `json:"name,omitempty"`

	// Entitlement ID.
	ID *string `json:"id,omitempty"`

	// Entitlement CRN.
	Crn *string `json:"crn,omitempty"`

	// URL for the BSS entitlement, e.g. /v1/licensing/entitlements/:id.
	URL *string `json:"url,omitempty"`

	// Entitlement offering type.
	OfferingType *string `json:"offering_type,omitempty"`

	// State of the BSS entitlement, e.g. 'active' or if it's been deleted, 'removed'.
	State *string `json:"state,omitempty"`

	// Entitlement is good from this starting date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveFrom *string `json:"effective_from,omitempty"`

	// Entitlement is good until this ending date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveUntil *string `json:"effective_until,omitempty"`

	// Account ID where this entitlement is bound to.
	AccountID *string `json:"account_id,omitempty"`

	// Account ID of owner.
	OwnerID *string `json:"owner_id,omitempty"`

	// GC ID of the specific offering version.
	VersionID *string `json:"version_id,omitempty"`

	// Marketplace offering ID for this license entitlement.
	LicenseOfferingID *string `json:"license_offering_id,omitempty"`

	// Specific license entitlement ID from the license provider, eg. D1W3R4.
	LicenseID *string `json:"license_id,omitempty"`

	// IBM ID of the owner of this license entitlement.
	LicenseOwnerID *string `json:"license_owner_id,omitempty"`

	// Type of license entitlement, e.g. ibm-ppa.
	LicenseType *string `json:"license_type,omitempty"`

	// ID of the license provider.
	LicenseProviderID *string `json:"license_provider_id,omitempty"`

	// URL for the BSS license provider, e.g. /v1/licensing/license_providers/:license_provider_id.
	LicenseProviderURL *string `json:"license_provider_url,omitempty"`

	// Specific license entitlement ID from the license provider, eg. D1W3R4.
	LicenseProductID *string `json:"license_product_id,omitempty"`

	// Location of the registry images, eg. cp/cp4d.
	NamespaceRepository *string `json:"namespace_repository,omitempty"`

	// API key for access to the license entitlement.
	Apikey *string `json:"apikey,omitempty"`

	// IBM ID.
	CreateBy *string `json:"create_by,omitempty"`

	// IBM ID.
	UpdateBy *string `json:"update_by,omitempty"`

	// Creation date, eg. '2019-07-17T21:21:47.6794935Z'.
	CreateAt *string `json:"create_at,omitempty"`

	// Date last updated, eg. '2019-07-17T21:21:47.6794935Z'.
	UpdatedAt *string `json:"updated_at,omitempty"`

	// Entitlement history.
	History []LicenseEntitlementHistoryItem `json:"history,omitempty"`

	// Array of license offering references.
	OfferingList []LicenseOfferingReference `json:"offering_list,omitempty"`
}


// UnmarshalLicenseEntitlement unmarshals an instance of LicenseEntitlement from the specified map of raw messages.
func UnmarshalLicenseEntitlement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseEntitlement)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_type", &obj.OfferingType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_from", &obj.EffectiveFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_until", &obj.EffectiveUntil)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_id", &obj.OwnerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_id", &obj.VersionID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_offering_id", &obj.LicenseOfferingID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_id", &obj.LicenseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_owner_id", &obj.LicenseOwnerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_type", &obj.LicenseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_provider_id", &obj.LicenseProviderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_provider_url", &obj.LicenseProviderURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_product_id", &obj.LicenseProductID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "namespace_repository", &obj.NamespaceRepository)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "apikey", &obj.Apikey)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "create_by", &obj.CreateBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "update_by", &obj.UpdateBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "create_at", &obj.CreateAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated_at", &obj.UpdatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "history", &obj.History, UnmarshalLicenseEntitlementHistoryItem)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "offering_list", &obj.OfferingList, UnmarshalLicenseOfferingReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseEntitlementHistoryItem : LicenseEntitlementHistoryItem struct
type LicenseEntitlementHistoryItem struct {
	// Eg. create.
	Action *string `json:"action,omitempty"`

	// Eg. IBM ID of user.
	User *string `json:"user,omitempty"`

	// Date of action, eg. '2019-07-17T21:21:47.6794935Z'.
	Date *string `json:"date,omitempty"`
}


// UnmarshalLicenseEntitlementHistoryItem unmarshals an instance of LicenseEntitlementHistoryItem from the specified map of raw messages.
func UnmarshalLicenseEntitlementHistoryItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseEntitlementHistoryItem)
	err = core.UnmarshalPrimitive(m, "action", &obj.Action)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "user", &obj.User)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "date", &obj.Date)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseEntitlements : Paginated list of license entitlements.
type LicenseEntitlements struct {
	// Total number of results.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Total number of pages.
	TotalPages *int64 `json:"total_pages,omitempty"`

	// Previous URL.
	PrevURL *string `json:"prev_url,omitempty"`

	// Next URL.
	NextURL *string `json:"next_url,omitempty"`

	// Resulting Entitlement objects.
	Resources []LicenseEntitlement `json:"resources,omitempty"`
}


// UnmarshalLicenseEntitlements unmarshals an instance of LicenseEntitlements from the specified map of raw messages.
func UnmarshalLicenseEntitlements(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseEntitlements)
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_pages", &obj.TotalPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev_url", &obj.PrevURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalLicenseEntitlement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseObject : License information.
type LicenseObject struct {
	// License name.
	Name *string `json:"name,omitempty"`

	// Type of offering.
	OfferingType *string `json:"offering_type,omitempty"`

	// Number of seats allowed for license.
	SeatsAllowed *string `json:"seats_allowed,omitempty"`

	// Number of seats used for license.
	SeatsUsed *string `json:"seats_used,omitempty"`

	// ID of license owner.
	OwnerID *string `json:"owner_id,omitempty"`

	// Marketplace offering ID for this license.
	LicenseOfferingID *string `json:"license_offering_id,omitempty"`

	// specific license entitlement ID from the license provider, eg. D1W3R4.
	LicenseID *string `json:"license_id,omitempty"`

	// IBM ID of the owner of this license entitlement.
	LicenseOwnerID *string `json:"license_owner_id,omitempty"`

	// type of license entitlement, e.g. ibm-ppa.
	LicenseType *string `json:"license_type,omitempty"`

	// ID of the license provider.
	LicenseProviderID *string `json:"license_provider_id,omitempty"`

	// specific license entitlement ID from the license provider, eg. D1W3R4.
	LicenseProductID *string `json:"license_product_id,omitempty"`

	// URL for the BSS license provider, e.g. /v1/licensing/license_providers/:license_provider_id.
	LicenseProviderURL *string `json:"license_provider_url,omitempty"`

	// license is good from this starting date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveFrom *string `json:"effective_from,omitempty"`

	// license is good until this ending date. eg. '2019-07-17T21:21:47.6794935Z'.
	EffectiveUntil *string `json:"effective_until,omitempty"`

	// If true, this will allow use of this license by all IBMers.
	Internal *bool `json:"internal,omitempty"`

	// Array of license offering references.
	OfferingList []LicenseOfferingReference `json:"offering_list,omitempty"`
}


// UnmarshalLicenseObject unmarshals an instance of LicenseObject from the specified map of raw messages.
func UnmarshalLicenseObject(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseObject)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_type", &obj.OfferingType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "seats_allowed", &obj.SeatsAllowed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "seats_used", &obj.SeatsUsed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner_id", &obj.OwnerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_offering_id", &obj.LicenseOfferingID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_id", &obj.LicenseID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_owner_id", &obj.LicenseOwnerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_type", &obj.LicenseType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_provider_id", &obj.LicenseProviderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_product_id", &obj.LicenseProductID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "license_provider_url", &obj.LicenseProviderURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_from", &obj.EffectiveFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "effective_until", &obj.EffectiveUntil)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "internal", &obj.Internal)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "offering_list", &obj.OfferingList, UnmarshalLicenseOfferingReference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseOfferingReference : License offering reference.
type LicenseOfferingReference struct {
	// Offering ID.
	ID *string `json:"id,omitempty"`

	// Offering name.
	Name *string `json:"name,omitempty"`

	// Offering label'.
	Label *string `json:"label,omitempty"`

	// URL to offering icon.
	OfferingIconURL *string `json:"offering_icon_url,omitempty"`

	// Account ID associated with offering.
	AccountID *string `json:"account_id,omitempty"`

	// Catalog ID associated with offering.
	CatalogID *string `json:"catalog_id,omitempty"`
}


// UnmarshalLicenseOfferingReference unmarshals an instance of LicenseOfferingReference from the specified map of raw messages.
func UnmarshalLicenseOfferingReference(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseOfferingReference)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_icon_url", &obj.OfferingIconURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseProvider : BSS License provider.
type LicenseProvider struct {
	// Provider name, eg. IBM Passport Advantage.
	Name *string `json:"name,omitempty"`

	// Short description of license provider.
	ShortDescription *string `json:"short_description,omitempty"`

	// Provider ID.
	ID *string `json:"id,omitempty"`

	// Type of license entitlement, e.g. ibm-ppa.
	LicenceType *string `json:"licence_type,omitempty"`

	// Type of offering.
	OfferingType *string `json:"offering_type,omitempty"`

	// URL of the license provider for where to create/get a license, e.g.
	// https://www.ibm.com/software/passportadvantage/aboutpassport.html.
	CreateURL *string `json:"create_url,omitempty"`

	// URL of the license provider for additional info, e.g. https://www.ibm.com/software/passportadvantage.
	InfoURL *string `json:"info_url,omitempty"`

	// URL for the BSS license provider, e.g. /v1/licensing/license_providers/:id.
	URL *string `json:"url,omitempty"`

	// Provider CRN.
	Crn *string `json:"crn,omitempty"`

	// State of license provider.
	State *string `json:"state,omitempty"`
}


// UnmarshalLicenseProvider unmarshals an instance of LicenseProvider from the specified map of raw messages.
func UnmarshalLicenseProvider(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseProvider)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "licence_type", &obj.LicenceType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_type", &obj.OfferingType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "create_url", &obj.CreateURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "info_url", &obj.InfoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LicenseProviders : Paginated list of license providers.
type LicenseProviders struct {
	// Total number of results.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Total number of pages.
	TotalPages *int64 `json:"total_pages,omitempty"`

	// Previous URL.
	PrevURL *string `json:"prev_url,omitempty"`

	// Next URL.
	NextURL *string `json:"next_url,omitempty"`

	// Resulting License Provider objects.
	Resources []LicenseProvider `json:"resources,omitempty"`
}


// UnmarshalLicenseProviders unmarshals an instance of LicenseProviders from the specified map of raw messages.
func UnmarshalLicenseProviders(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LicenseProviders)
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_pages", &obj.TotalPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev_url", &obj.PrevURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalLicenseProvider)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Licenses : Paginated list of licenses.
type Licenses struct {
	// Total number of results.
	TotalResults *int64 `json:"total_results,omitempty"`

	// Total number of pages.
	TotalPages *int64 `json:"total_pages,omitempty"`

	// Previous URL.
	PrevURL *string `json:"prev_url,omitempty"`

	// Next URL.
	NextURL *string `json:"next_url,omitempty"`

	// Resulting License objects.
	Resources []LicenseObject `json:"resources,omitempty"`
}


// UnmarshalLicenses unmarshals an instance of Licenses from the specified map of raw messages.
func UnmarshalLicenses(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Licenses)
	err = core.UnmarshalPrimitive(m, "total_results", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_pages", &obj.TotalPages)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev_url", &obj.PrevURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next_url", &obj.NextURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalLicenseObject)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListCatalogsOptions : The ListCatalogs options.
type ListCatalogsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCatalogsOptions : Instantiate ListCatalogsOptions
func (*CatalogManagementV1) NewListCatalogsOptions() *ListCatalogsOptions {
	return &ListCatalogsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListCatalogsOptions) SetHeaders(param map[string]string) *ListCatalogsOptions {
	options.Headers = param
	return options
}

// ListClustersOptions : The ListClusters options.
type ListClustersOptions struct {
	// number or results to return.
	Limit *int64 `json:"limit,omitempty"`

	// number of results to skip before returning values.
	Offset *int64 `json:"offset,omitempty"`

	// Kubernetes or OpenShift.  Default is kubernetes.
	Type *string `json:"type,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListClustersOptions : Instantiate ListClustersOptions
func (*CatalogManagementV1) NewListClustersOptions() *ListClustersOptions {
	return &ListClustersOptions{}
}

// SetLimit : Allow user to set Limit
func (options *ListClustersOptions) SetLimit(limit int64) *ListClustersOptions {
	options.Limit = core.Int64Ptr(limit)
	return options
}

// SetOffset : Allow user to set Offset
func (options *ListClustersOptions) SetOffset(offset int64) *ListClustersOptions {
	options.Offset = core.Int64Ptr(offset)
	return options
}

// SetType : Allow user to set Type
func (options *ListClustersOptions) SetType(typeVar string) *ListClustersOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListClustersOptions) SetHeaders(param map[string]string) *ListClustersOptions {
	options.Headers = param
	return options
}

// ListKubeDeployOptions : The ListKubeDeploy options.
type ListKubeDeployOptions struct {
	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster identification.
	ClusterID *string `json:"cluster_id" validate:"required"`

	// Kube namespace.
	Namespace *string `json:"namespace" validate:"required"`

	// Kube region.
	Region *string `json:"region" validate:"required"`

	// Catalog version locator identification.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListKubeDeployOptions : Instantiate ListKubeDeployOptions
func (*CatalogManagementV1) NewListKubeDeployOptions(xAuthRefreshToken string, clusterID string, namespace string, region string, versionLocatorID string) *ListKubeDeployOptions {
	return &ListKubeDeployOptions{
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
		ClusterID: core.StringPtr(clusterID),
		Namespace: core.StringPtr(namespace),
		Region: core.StringPtr(region),
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *ListKubeDeployOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *ListKubeDeployOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *ListKubeDeployOptions) SetClusterID(clusterID string) *ListKubeDeployOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *ListKubeDeployOptions) SetNamespace(namespace string) *ListKubeDeployOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetRegion : Allow user to set Region
func (options *ListKubeDeployOptions) SetRegion(region string) *ListKubeDeployOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *ListKubeDeployOptions) SetVersionLocatorID(versionLocatorID string) *ListKubeDeployOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListKubeDeployOptions) SetHeaders(param map[string]string) *ListKubeDeployOptions {
	options.Headers = param
	return options
}

// ListLicenseEntitlementsOptions : The ListLicenseEntitlements options.
type ListLicenseEntitlementsOptions struct {
	// The account ID to query for the entitlement. Default is the account from the user's token.
	AccountID *string `json:"account_id,omitempty"`

	// The license product ID. If from PPA (Passport Advantage) this is the product Part number(s) which can be one or more
	// IDs, eg. D1YGZLL,5737L09.
	LicenseProductID *string `json:"license_product_id,omitempty"`

	// The GC ID of the specific offering version.
	VersionID *string `json:"version_id,omitempty"`

	// The state of the license entitlement. eg. usually 'active', or if it's been deleted will show as 'removed'.
	State *string `json:"state,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLicenseEntitlementsOptions : Instantiate ListLicenseEntitlementsOptions
func (*CatalogManagementV1) NewListLicenseEntitlementsOptions() *ListLicenseEntitlementsOptions {
	return &ListLicenseEntitlementsOptions{}
}

// SetAccountID : Allow user to set AccountID
func (options *ListLicenseEntitlementsOptions) SetAccountID(accountID string) *ListLicenseEntitlementsOptions {
	options.AccountID = core.StringPtr(accountID)
	return options
}

// SetLicenseProductID : Allow user to set LicenseProductID
func (options *ListLicenseEntitlementsOptions) SetLicenseProductID(licenseProductID string) *ListLicenseEntitlementsOptions {
	options.LicenseProductID = core.StringPtr(licenseProductID)
	return options
}

// SetVersionID : Allow user to set VersionID
func (options *ListLicenseEntitlementsOptions) SetVersionID(versionID string) *ListLicenseEntitlementsOptions {
	options.VersionID = core.StringPtr(versionID)
	return options
}

// SetState : Allow user to set State
func (options *ListLicenseEntitlementsOptions) SetState(state string) *ListLicenseEntitlementsOptions {
	options.State = core.StringPtr(state)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListLicenseEntitlementsOptions) SetHeaders(param map[string]string) *ListLicenseEntitlementsOptions {
	options.Headers = param
	return options
}

// ListOfferingsOptions : The ListOfferings options.
type ListOfferingsOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// true - Strip down the content of what is returned. For example don't return the readme. Makes the result much
	// smaller. Defaults to false.
	Digest *bool `json:"digest,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListOfferingsOptions : Instantiate ListOfferingsOptions
func (*CatalogManagementV1) NewListOfferingsOptions(catalogID string) *ListOfferingsOptions {
	return &ListOfferingsOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ListOfferingsOptions) SetCatalogID(catalogID string) *ListOfferingsOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetDigest : Allow user to set Digest
func (options *ListOfferingsOptions) SetDigest(digest bool) *ListOfferingsOptions {
	options.Digest = core.BoolPtr(digest)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListOfferingsOptions) SetHeaders(param map[string]string) *ListOfferingsOptions {
	options.Headers = param
	return options
}

// ListVersionsOptions : The ListVersions options.
type ListVersionsOptions struct {
	// query, for now only "q=entitlement_key:<some-key>" is supported.
	Q *string `json:"q" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVersionsOptions : Instantiate ListVersionsOptions
func (*CatalogManagementV1) NewListVersionsOptions(q string) *ListVersionsOptions {
	return &ListVersionsOptions{
		Q: core.StringPtr(q),
	}
}

// SetQ : Allow user to set Q
func (options *ListVersionsOptions) SetQ(q string) *ListVersionsOptions {
	options.Q = core.StringPtr(q)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListVersionsOptions) SetHeaders(param map[string]string) *ListVersionsOptions {
	options.Headers = param
	return options
}

// Maintainers : Repo maintainers.
type Maintainers struct {
	// Maintainer email address.
	Email *string `json:"email,omitempty"`

	// Name of maintainer.
	Name *string `json:"name,omitempty"`
}


// UnmarshalMaintainers unmarshals an instance of Maintainers from the specified map of raw messages.
func UnmarshalMaintainers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Maintainers)
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NamespaceSearchResult : Paginated list of namespace search results.
type NamespaceSearchResult struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting objects.
	Resources []string `json:"resources,omitempty"`
}


// UnmarshalNamespaceSearchResult unmarshals an instance of NamespaceSearchResult from the specified map of raw messages.
func UnmarshalNamespaceSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NamespaceSearchResult)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
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

// Offering : Offering information.
type Offering struct {
	// unique id.
	ID *string `json:"id,omitempty"`

	// The url for this specific offering.
	URL *string `json:"url,omitempty"`

	// The crn for this specific offering.
	Crn *string `json:"crn,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// The programmatic name of this offering.
	Name *string `json:"name,omitempty"`

	// URL for an icon associated with this offering.
	OfferingIconURL *string `json:"offering_icon_url,omitempty"`

	// URL for an additional docs with this offering.
	OfferingDocsURL *string `json:"offering_docs_url,omitempty"`

	// URL to be displayed in the Consumption UI for getting support on this offering.
	OfferingSupportURL *string `json:"offering_support_url,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// Repository info for offerings.
	Rating *Rating `json:"rating,omitempty"`

	// The date and time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date and time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Short description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// Long description in the requested language.
	LongDescription *string `json:"long_description,omitempty"`

	// list of features associated with this offering.
	Features []Feature `json:"features,omitempty"`

	// Array of kind.
	Kinds []Kind `json:"kinds,omitempty"`

	// Is it permitted to request publishing to IBM or Public.
	PermitRequestIbmPublicPublish *bool `json:"permit_request_ibm_public_publish,omitempty"`

	// Indicates if this offering has been approved for use by all IBMers.
	IbmPublishApproved *bool `json:"ibm_publish_approved,omitempty"`

	// Indicates if this offering has been approved for use by all IBM Cloud users.
	PublicPublishApproved *bool `json:"public_publish_approved,omitempty"`

	// The original offering CRN that this publish entry came from.
	PublicOriginalCrn *string `json:"public_original_crn,omitempty"`

	// The crn of the public catalog entry of this offering.
	PublishPublicCrn *string `json:"publish_public_crn,omitempty"`

	// The portal's approval record ID.
	PortalApprovalRecord *string `json:"portal_approval_record,omitempty"`

	// The portal UI URL.
	PortalUiURL *string `json:"portal_ui_url,omitempty"`

	// The id of the catalog containing this offering.
	CatalogID *string `json:"catalog_id,omitempty"`

	// The name of the catalog.
	CatalogName *string `json:"catalog_name,omitempty"`

	// Map of metadata values for this offering.
	Metadata interface{} `json:"metadata,omitempty"`

	// A disclaimer for this offering.
	Disclaimer *string `json:"disclaimer,omitempty"`

	// Determine if this offering should be displayed in the Consumption UI.
	Hidden *bool `json:"hidden,omitempty"`

	// Provider of this offering.
	Provider *string `json:"provider,omitempty"`

	// Repository info for offerings.
	RepoInfo *RepoInfo `json:"repo_info,omitempty"`
}


// UnmarshalOffering unmarshals an instance of Offering from the specified map of raw messages.
func UnmarshalOffering(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Offering)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_icon_url", &obj.OfferingIconURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_docs_url", &obj.OfferingDocsURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_support_url", &obj.OfferingSupportURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "rating", &obj.Rating, UnmarshalRating)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "long_description", &obj.LongDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "features", &obj.Features, UnmarshalFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "kinds", &obj.Kinds, UnmarshalKind)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "permit_request_ibm_public_publish", &obj.PermitRequestIbmPublicPublish)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ibm_publish_approved", &obj.IbmPublishApproved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_publish_approved", &obj.PublicPublishApproved)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_original_crn", &obj.PublicOriginalCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "publish_public_crn", &obj.PublishPublicCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "portal_approval_record", &obj.PortalApprovalRecord)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "portal_ui_url", &obj.PortalUiURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_name", &obj.CatalogName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "disclaimer", &obj.Disclaimer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hidden", &obj.Hidden)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "repo_info", &obj.RepoInfo, UnmarshalRepoInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OfferingSearchResult : Paginated offering search result.
type OfferingSearchResult struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting objects.
	Resources []Offering `json:"resources,omitempty"`
}


// UnmarshalOfferingSearchResult unmarshals an instance of OfferingSearchResult from the specified map of raw messages.
func UnmarshalOfferingSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OfferingSearchResult)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalOffering)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Plan : Offering plan.
type Plan struct {
	// unique id.
	ID *string `json:"id,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// The programmatic name of this offering.
	Name *string `json:"name,omitempty"`

	// Short description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// Long description in the requested language.
	LongDescription *string `json:"long_description,omitempty"`

	// open ended metadata information.
	Metadata interface{} `json:"metadata,omitempty"`

	// list of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// list of features associated with this offering.
	AdditionalFeatures []Feature `json:"additional_features,omitempty"`

	// the date'time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// the date'time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// list of deployments.
	Deployments []Deployment `json:"deployments,omitempty"`
}


// UnmarshalPlan unmarshals an instance of Plan from the specified map of raw messages.
func UnmarshalPlan(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Plan)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "short_description", &obj.ShortDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "long_description", &obj.LongDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "additional_features", &obj.AdditionalFeatures, UnmarshalFeature)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "deployments", &obj.Deployments, UnmarshalDeployment)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PreinstallVersionOptions : The PreinstallVersion options.
type PreinstallVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster ID.
	NewClusterID *string `json:"new_cluster_id,omitempty"`

	// Cluster region.
	NewRegion *string `json:"new_region,omitempty"`

	// Kube namespace.
	NewNamespace *string `json:"new_namespace,omitempty"`

	// Object containing Helm chart override values.
	NewOverrideValues interface{} `json:"new_override_values,omitempty"`

	// Entitlement API Key for this offering.
	NewEntitlementApikey *string `json:"new_entitlement_apikey,omitempty"`

	// Schematics workspace configuration.
	NewSchematics *ClusterRequestBodySchematics `json:"new_schematics,omitempty"`

	// Script.
	NewScript *string `json:"new_script,omitempty"`

	// Script ID.
	NewScriptID *string `json:"new_script_id,omitempty"`

	// Version locator ID.
	NewVersionLocatorID *string `json:"new_version_locator_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPreinstallVersionOptions : Instantiate PreinstallVersionOptions
func (*CatalogManagementV1) NewPreinstallVersionOptions(versionLocatorID string, xAuthRefreshToken string) *PreinstallVersionOptions {
	return &PreinstallVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *PreinstallVersionOptions) SetVersionLocatorID(versionLocatorID string) *PreinstallVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *PreinstallVersionOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *PreinstallVersionOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetNewClusterID : Allow user to set NewClusterID
func (options *PreinstallVersionOptions) SetNewClusterID(newClusterID string) *PreinstallVersionOptions {
	options.NewClusterID = core.StringPtr(newClusterID)
	return options
}

// SetNewRegion : Allow user to set NewRegion
func (options *PreinstallVersionOptions) SetNewRegion(newRegion string) *PreinstallVersionOptions {
	options.NewRegion = core.StringPtr(newRegion)
	return options
}

// SetNewNamespace : Allow user to set NewNamespace
func (options *PreinstallVersionOptions) SetNewNamespace(newNamespace string) *PreinstallVersionOptions {
	options.NewNamespace = core.StringPtr(newNamespace)
	return options
}

// SetNewOverrideValues : Allow user to set NewOverrideValues
func (options *PreinstallVersionOptions) SetNewOverrideValues(newOverrideValues interface{}) *PreinstallVersionOptions {
	options.NewOverrideValues = newOverrideValues
	return options
}

// SetNewEntitlementApikey : Allow user to set NewEntitlementApikey
func (options *PreinstallVersionOptions) SetNewEntitlementApikey(newEntitlementApikey string) *PreinstallVersionOptions {
	options.NewEntitlementApikey = core.StringPtr(newEntitlementApikey)
	return options
}

// SetNewSchematics : Allow user to set NewSchematics
func (options *PreinstallVersionOptions) SetNewSchematics(newSchematics *ClusterRequestBodySchematics) *PreinstallVersionOptions {
	options.NewSchematics = newSchematics
	return options
}

// SetNewScript : Allow user to set NewScript
func (options *PreinstallVersionOptions) SetNewScript(newScript string) *PreinstallVersionOptions {
	options.NewScript = core.StringPtr(newScript)
	return options
}

// SetNewScriptID : Allow user to set NewScriptID
func (options *PreinstallVersionOptions) SetNewScriptID(newScriptID string) *PreinstallVersionOptions {
	options.NewScriptID = core.StringPtr(newScriptID)
	return options
}

// SetNewVersionLocatorID : Allow user to set NewVersionLocatorID
func (options *PreinstallVersionOptions) SetNewVersionLocatorID(newVersionLocatorID string) *PreinstallVersionOptions {
	options.NewVersionLocatorID = core.StringPtr(newVersionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PreinstallVersionOptions) SetHeaders(param map[string]string) *PreinstallVersionOptions {
	options.Headers = param
	return options
}

// PublicPublishVersionOptions : The PublicPublishVersion options.
type PublicPublishVersionOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPublicPublishVersionOptions : Instantiate PublicPublishVersionOptions
func (*CatalogManagementV1) NewPublicPublishVersionOptions(versionLocatorID string) *PublicPublishVersionOptions {
	return &PublicPublishVersionOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *PublicPublishVersionOptions) SetVersionLocatorID(versionLocatorID string) *PublicPublishVersionOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PublicPublishVersionOptions) SetHeaders(param map[string]string) *PublicPublishVersionOptions {
	options.Headers = param
	return options
}

// Rating : Repository info for offerings.
type Rating struct {
	// One start rating.
	OneStarCount *int64 `json:"one_star_count,omitempty"`

	// Two start rating.
	TwoStarCount *int64 `json:"two_star_count,omitempty"`

	// Three start rating.
	ThreeStarCount *int64 `json:"three_star_count,omitempty"`

	// Four start rating.
	FourStarCount *int64 `json:"four_star_count,omitempty"`
}


// UnmarshalRating unmarshals an instance of Rating from the specified map of raw messages.
func UnmarshalRating(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Rating)
	err = core.UnmarshalPrimitive(m, "one_star_count", &obj.OneStarCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "two_star_count", &obj.TwoStarCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "three_star_count", &obj.ThreeStarCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "four_star_count", &obj.FourStarCount)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReloadOfferingOptions : The ReloadOffering options.
type ReloadOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// URL path to zip location.
	Zipurl *string `json:"zipurl" validate:"required"`

	// The semver value for this new version.
	TargetVersion *string `json:"targetVersion" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReloadOfferingOptions : Instantiate ReloadOfferingOptions
func (*CatalogManagementV1) NewReloadOfferingOptions(catalogID string, offeringID string, zipurl string, targetVersion string) *ReloadOfferingOptions {
	return &ReloadOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
		Zipurl: core.StringPtr(zipurl),
		TargetVersion: core.StringPtr(targetVersion),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ReloadOfferingOptions) SetCatalogID(catalogID string) *ReloadOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *ReloadOfferingOptions) SetOfferingID(offeringID string) *ReloadOfferingOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetZipurl : Allow user to set Zipurl
func (options *ReloadOfferingOptions) SetZipurl(zipurl string) *ReloadOfferingOptions {
	options.Zipurl = core.StringPtr(zipurl)
	return options
}

// SetTargetVersion : Allow user to set TargetVersion
func (options *ReloadOfferingOptions) SetTargetVersion(targetVersion string) *ReloadOfferingOptions {
	options.TargetVersion = core.StringPtr(targetVersion)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReloadOfferingOptions) SetHeaders(param map[string]string) *ReloadOfferingOptions {
	options.Headers = param
	return options
}

// ReplaceCatalogOptions : The ReplaceCatalog options.
type ReplaceCatalogOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Unique ID.
	ID *string `json:"id,omitempty"`

	// Display Name in the requested language.
	Label *string `json:"label,omitempty"`

	// Description in the requested language.
	ShortDescription *string `json:"short_description,omitempty"`

	// URL for an icon associated with this catalog.
	CatalogIconURL *string `json:"catalog_icon_url,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// The url for this specific catalog.
	URL *string `json:"url,omitempty"`

	// CRN associated with the catalog.
	Crn *string `json:"crn,omitempty"`

	// URL path to offerings.
	OfferingsURL *string `json:"offerings_url,omitempty"`

	// List of features associated with this catalog.
	Features []Feature `json:"features,omitempty"`

	// Denotes whether a catalog is disabled.
	Disabled *bool `json:"disabled,omitempty"`

	// The date'time this catalog was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date'time this catalog was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Resource group id the catalog is owned by.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// Account that owns catalog.
	OwningAccount *string `json:"owning_account,omitempty"`

	// Filters for account and catalog filters.
	CatalogFilters *Filters `json:"catalog_filters,omitempty"`

	// Feature information.
	SyndicationSettings *SyndicationResource `json:"syndication_settings,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceCatalogOptions : Instantiate ReplaceCatalogOptions
func (*CatalogManagementV1) NewReplaceCatalogOptions(catalogID string) *ReplaceCatalogOptions {
	return &ReplaceCatalogOptions{
		CatalogID: core.StringPtr(catalogID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ReplaceCatalogOptions) SetCatalogID(catalogID string) *ReplaceCatalogOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetID : Allow user to set ID
func (options *ReplaceCatalogOptions) SetID(id string) *ReplaceCatalogOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetLabel : Allow user to set Label
func (options *ReplaceCatalogOptions) SetLabel(label string) *ReplaceCatalogOptions {
	options.Label = core.StringPtr(label)
	return options
}

// SetShortDescription : Allow user to set ShortDescription
func (options *ReplaceCatalogOptions) SetShortDescription(shortDescription string) *ReplaceCatalogOptions {
	options.ShortDescription = core.StringPtr(shortDescription)
	return options
}

// SetCatalogIconURL : Allow user to set CatalogIconURL
func (options *ReplaceCatalogOptions) SetCatalogIconURL(catalogIconURL string) *ReplaceCatalogOptions {
	options.CatalogIconURL = core.StringPtr(catalogIconURL)
	return options
}

// SetTags : Allow user to set Tags
func (options *ReplaceCatalogOptions) SetTags(tags []string) *ReplaceCatalogOptions {
	options.Tags = tags
	return options
}

// SetURL : Allow user to set URL
func (options *ReplaceCatalogOptions) SetURL(url string) *ReplaceCatalogOptions {
	options.URL = core.StringPtr(url)
	return options
}

// SetCrn : Allow user to set Crn
func (options *ReplaceCatalogOptions) SetCrn(crn string) *ReplaceCatalogOptions {
	options.Crn = core.StringPtr(crn)
	return options
}

// SetOfferingsURL : Allow user to set OfferingsURL
func (options *ReplaceCatalogOptions) SetOfferingsURL(offeringsURL string) *ReplaceCatalogOptions {
	options.OfferingsURL = core.StringPtr(offeringsURL)
	return options
}

// SetFeatures : Allow user to set Features
func (options *ReplaceCatalogOptions) SetFeatures(features []Feature) *ReplaceCatalogOptions {
	options.Features = features
	return options
}

// SetDisabled : Allow user to set Disabled
func (options *ReplaceCatalogOptions) SetDisabled(disabled bool) *ReplaceCatalogOptions {
	options.Disabled = core.BoolPtr(disabled)
	return options
}

// SetCreated : Allow user to set Created
func (options *ReplaceCatalogOptions) SetCreated(created *strfmt.DateTime) *ReplaceCatalogOptions {
	options.Created = created
	return options
}

// SetUpdated : Allow user to set Updated
func (options *ReplaceCatalogOptions) SetUpdated(updated *strfmt.DateTime) *ReplaceCatalogOptions {
	options.Updated = updated
	return options
}

// SetResourceGroupID : Allow user to set ResourceGroupID
func (options *ReplaceCatalogOptions) SetResourceGroupID(resourceGroupID string) *ReplaceCatalogOptions {
	options.ResourceGroupID = core.StringPtr(resourceGroupID)
	return options
}

// SetOwningAccount : Allow user to set OwningAccount
func (options *ReplaceCatalogOptions) SetOwningAccount(owningAccount string) *ReplaceCatalogOptions {
	options.OwningAccount = core.StringPtr(owningAccount)
	return options
}

// SetCatalogFilters : Allow user to set CatalogFilters
func (options *ReplaceCatalogOptions) SetCatalogFilters(catalogFilters *Filters) *ReplaceCatalogOptions {
	options.CatalogFilters = catalogFilters
	return options
}

// SetSyndicationSettings : Allow user to set SyndicationSettings
func (options *ReplaceCatalogOptions) SetSyndicationSettings(syndicationSettings *SyndicationResource) *ReplaceCatalogOptions {
	options.SyndicationSettings = syndicationSettings
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceCatalogOptions) SetHeaders(param map[string]string) *ReplaceCatalogOptions {
	options.Headers = param
	return options
}

// ReplaceEnterpriseOptions : The ReplaceEnterprise options.
type ReplaceEnterpriseOptions struct {
	// Enterprise identification.
	EnterpriseID *string `json:"enterprise_id" validate:"required"`

	// Enterprise identification.
	ID *string `json:"id,omitempty"`

	// Filters for account and catalog filters.
	AccountFilters *Filters `json:"account_filters,omitempty"`

	// Map of account group ids to AccountGroup objects.
	AccountGroups *EnterpriseAccountGroups `json:"account_groups,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceEnterpriseOptions : Instantiate ReplaceEnterpriseOptions
func (*CatalogManagementV1) NewReplaceEnterpriseOptions(enterpriseID string) *ReplaceEnterpriseOptions {
	return &ReplaceEnterpriseOptions{
		EnterpriseID: core.StringPtr(enterpriseID),
	}
}

// SetEnterpriseID : Allow user to set EnterpriseID
func (options *ReplaceEnterpriseOptions) SetEnterpriseID(enterpriseID string) *ReplaceEnterpriseOptions {
	options.EnterpriseID = core.StringPtr(enterpriseID)
	return options
}

// SetID : Allow user to set ID
func (options *ReplaceEnterpriseOptions) SetID(id string) *ReplaceEnterpriseOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccountFilters : Allow user to set AccountFilters
func (options *ReplaceEnterpriseOptions) SetAccountFilters(accountFilters *Filters) *ReplaceEnterpriseOptions {
	options.AccountFilters = accountFilters
	return options
}

// SetAccountGroups : Allow user to set AccountGroups
func (options *ReplaceEnterpriseOptions) SetAccountGroups(accountGroups *EnterpriseAccountGroups) *ReplaceEnterpriseOptions {
	options.AccountGroups = accountGroups
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceEnterpriseOptions) SetHeaders(param map[string]string) *ReplaceEnterpriseOptions {
	options.Headers = param
	return options
}

// ReplaceKubeDeployOptions : The ReplaceKubeDeploy options.
type ReplaceKubeDeployOptions struct {
	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster ID.
	ClusterID *string `json:"cluster_id,omitempty"`

	// Cluster region.
	Region *string `json:"region,omitempty"`

	// Kube namespace.
	Namespace *string `json:"namespace,omitempty"`

	// Catalog version locator ID.
	VersionLocatorID *string `json:"version_locator_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceKubeDeployOptions : Instantiate ReplaceKubeDeployOptions
func (*CatalogManagementV1) NewReplaceKubeDeployOptions(xAuthRefreshToken string) *ReplaceKubeDeployOptions {
	return &ReplaceKubeDeployOptions{
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *ReplaceKubeDeployOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *ReplaceKubeDeployOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetClusterID : Allow user to set ClusterID
func (options *ReplaceKubeDeployOptions) SetClusterID(clusterID string) *ReplaceKubeDeployOptions {
	options.ClusterID = core.StringPtr(clusterID)
	return options
}

// SetRegion : Allow user to set Region
func (options *ReplaceKubeDeployOptions) SetRegion(region string) *ReplaceKubeDeployOptions {
	options.Region = core.StringPtr(region)
	return options
}

// SetNamespace : Allow user to set Namespace
func (options *ReplaceKubeDeployOptions) SetNamespace(namespace string) *ReplaceKubeDeployOptions {
	options.Namespace = core.StringPtr(namespace)
	return options
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *ReplaceKubeDeployOptions) SetVersionLocatorID(versionLocatorID string) *ReplaceKubeDeployOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceKubeDeployOptions) SetHeaders(param map[string]string) *ReplaceKubeDeployOptions {
	options.Headers = param
	return options
}

// ReplaceOfferingIconOptions : The ReplaceOfferingIcon options.
type ReplaceOfferingIconOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// Name of the file name that is being uploaded.
	FileName *string `json:"file_name" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceOfferingIconOptions : Instantiate ReplaceOfferingIconOptions
func (*CatalogManagementV1) NewReplaceOfferingIconOptions(catalogID string, offeringID string, fileName string) *ReplaceOfferingIconOptions {
	return &ReplaceOfferingIconOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
		FileName: core.StringPtr(fileName),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ReplaceOfferingIconOptions) SetCatalogID(catalogID string) *ReplaceOfferingIconOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *ReplaceOfferingIconOptions) SetOfferingID(offeringID string) *ReplaceOfferingIconOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetFileName : Allow user to set FileName
func (options *ReplaceOfferingIconOptions) SetFileName(fileName string) *ReplaceOfferingIconOptions {
	options.FileName = core.StringPtr(fileName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceOfferingIconOptions) SetHeaders(param map[string]string) *ReplaceOfferingIconOptions {
	options.Headers = param
	return options
}

// ReplaceOfferingOptions : The ReplaceOffering options.
type ReplaceOfferingOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// unique id.
	NewID *string `json:"new_id,omitempty"`

	// The url for this specific offering.
	NewURL *string `json:"new_url,omitempty"`

	// The crn for this specific offering.
	NewCrn *string `json:"new_crn,omitempty"`

	// Display Name in the requested language.
	NewLabel *string `json:"new_label,omitempty"`

	// The programmatic name of this offering.
	NewName *string `json:"new_name,omitempty"`

	// URL for an icon associated with this offering.
	NewOfferingIconURL *string `json:"new_offering_icon_url,omitempty"`

	// URL for an additional docs with this offering.
	NewOfferingDocsURL *string `json:"new_offering_docs_url,omitempty"`

	// URL to be displayed in the Consumption UI for getting support on this offering.
	NewOfferingSupportURL *string `json:"new_offering_support_url,omitempty"`

	// List of tags associated with this catalog.
	NewTags []string `json:"new_tags,omitempty"`

	// Repository info for offerings.
	NewRating *Rating `json:"new_rating,omitempty"`

	// The date and time this catalog was created.
	NewCreated *strfmt.DateTime `json:"new_created,omitempty"`

	// The date and time this catalog was last updated.
	NewUpdated *strfmt.DateTime `json:"new_updated,omitempty"`

	// Short description in the requested language.
	NewShortDescription *string `json:"new_short_description,omitempty"`

	// Long description in the requested language.
	NewLongDescription *string `json:"new_long_description,omitempty"`

	// list of features associated with this offering.
	NewFeatures []Feature `json:"new_features,omitempty"`

	// Array of kind.
	NewKinds []Kind `json:"new_kinds,omitempty"`

	// Is it permitted to request publishing to IBM or Public.
	NewPermitRequestIbmPublicPublish *bool `json:"new_permit_request_ibm_public_publish,omitempty"`

	// Indicates if this offering has been approved for use by all IBMers.
	NewIbmPublishApproved *bool `json:"new_ibm_publish_approved,omitempty"`

	// Indicates if this offering has been approved for use by all IBM Cloud users.
	NewPublicPublishApproved *bool `json:"new_public_publish_approved,omitempty"`

	// The original offering CRN that this publish entry came from.
	NewPublicOriginalCrn *string `json:"new_public_original_crn,omitempty"`

	// The crn of the public catalog entry of this offering.
	NewPublishPublicCrn *string `json:"new_publish_public_crn,omitempty"`

	// The portal's approval record ID.
	NewPortalApprovalRecord *string `json:"new_portal_approval_record,omitempty"`

	// The portal UI URL.
	NewPortalUiURL *string `json:"new_portal_ui_url,omitempty"`

	// The id of the catalog containing this offering.
	NewCatalogID *string `json:"new_catalog_id,omitempty"`

	// The name of the catalog.
	NewCatalogName *string `json:"new_catalog_name,omitempty"`

	// Map of metadata values for this offering.
	NewMetadata interface{} `json:"new_metadata,omitempty"`

	// A disclaimer for this offering.
	NewDisclaimer *string `json:"new_disclaimer,omitempty"`

	// Determine if this offering should be displayed in the Consumption UI.
	NewHidden *bool `json:"new_hidden,omitempty"`

	// Provider of this offering.
	NewProvider *string `json:"new_provider,omitempty"`

	// Repository info for offerings.
	NewRepoInfo *RepoInfo `json:"new_repo_info,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceOfferingOptions : Instantiate ReplaceOfferingOptions
func (*CatalogManagementV1) NewReplaceOfferingOptions(catalogID string, offeringID string) *ReplaceOfferingOptions {
	return &ReplaceOfferingOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *ReplaceOfferingOptions) SetCatalogID(catalogID string) *ReplaceOfferingOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *ReplaceOfferingOptions) SetOfferingID(offeringID string) *ReplaceOfferingOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetNewID : Allow user to set NewID
func (options *ReplaceOfferingOptions) SetNewID(newID string) *ReplaceOfferingOptions {
	options.NewID = core.StringPtr(newID)
	return options
}

// SetNewURL : Allow user to set NewURL
func (options *ReplaceOfferingOptions) SetNewURL(newURL string) *ReplaceOfferingOptions {
	options.NewURL = core.StringPtr(newURL)
	return options
}

// SetNewCrn : Allow user to set NewCrn
func (options *ReplaceOfferingOptions) SetNewCrn(newCrn string) *ReplaceOfferingOptions {
	options.NewCrn = core.StringPtr(newCrn)
	return options
}

// SetNewLabel : Allow user to set NewLabel
func (options *ReplaceOfferingOptions) SetNewLabel(newLabel string) *ReplaceOfferingOptions {
	options.NewLabel = core.StringPtr(newLabel)
	return options
}

// SetNewName : Allow user to set NewName
func (options *ReplaceOfferingOptions) SetNewName(newName string) *ReplaceOfferingOptions {
	options.NewName = core.StringPtr(newName)
	return options
}

// SetNewOfferingIconURL : Allow user to set NewOfferingIconURL
func (options *ReplaceOfferingOptions) SetNewOfferingIconURL(newOfferingIconURL string) *ReplaceOfferingOptions {
	options.NewOfferingIconURL = core.StringPtr(newOfferingIconURL)
	return options
}

// SetNewOfferingDocsURL : Allow user to set NewOfferingDocsURL
func (options *ReplaceOfferingOptions) SetNewOfferingDocsURL(newOfferingDocsURL string) *ReplaceOfferingOptions {
	options.NewOfferingDocsURL = core.StringPtr(newOfferingDocsURL)
	return options
}

// SetNewOfferingSupportURL : Allow user to set NewOfferingSupportURL
func (options *ReplaceOfferingOptions) SetNewOfferingSupportURL(newOfferingSupportURL string) *ReplaceOfferingOptions {
	options.NewOfferingSupportURL = core.StringPtr(newOfferingSupportURL)
	return options
}

// SetNewTags : Allow user to set NewTags
func (options *ReplaceOfferingOptions) SetNewTags(newTags []string) *ReplaceOfferingOptions {
	options.NewTags = newTags
	return options
}

// SetNewRating : Allow user to set NewRating
func (options *ReplaceOfferingOptions) SetNewRating(newRating *Rating) *ReplaceOfferingOptions {
	options.NewRating = newRating
	return options
}

// SetNewCreated : Allow user to set NewCreated
func (options *ReplaceOfferingOptions) SetNewCreated(newCreated *strfmt.DateTime) *ReplaceOfferingOptions {
	options.NewCreated = newCreated
	return options
}

// SetNewUpdated : Allow user to set NewUpdated
func (options *ReplaceOfferingOptions) SetNewUpdated(newUpdated *strfmt.DateTime) *ReplaceOfferingOptions {
	options.NewUpdated = newUpdated
	return options
}

// SetNewShortDescription : Allow user to set NewShortDescription
func (options *ReplaceOfferingOptions) SetNewShortDescription(newShortDescription string) *ReplaceOfferingOptions {
	options.NewShortDescription = core.StringPtr(newShortDescription)
	return options
}

// SetNewLongDescription : Allow user to set NewLongDescription
func (options *ReplaceOfferingOptions) SetNewLongDescription(newLongDescription string) *ReplaceOfferingOptions {
	options.NewLongDescription = core.StringPtr(newLongDescription)
	return options
}

// SetNewFeatures : Allow user to set NewFeatures
func (options *ReplaceOfferingOptions) SetNewFeatures(newFeatures []Feature) *ReplaceOfferingOptions {
	options.NewFeatures = newFeatures
	return options
}

// SetNewKinds : Allow user to set NewKinds
func (options *ReplaceOfferingOptions) SetNewKinds(newKinds []Kind) *ReplaceOfferingOptions {
	options.NewKinds = newKinds
	return options
}

// SetNewPermitRequestIbmPublicPublish : Allow user to set NewPermitRequestIbmPublicPublish
func (options *ReplaceOfferingOptions) SetNewPermitRequestIbmPublicPublish(newPermitRequestIbmPublicPublish bool) *ReplaceOfferingOptions {
	options.NewPermitRequestIbmPublicPublish = core.BoolPtr(newPermitRequestIbmPublicPublish)
	return options
}

// SetNewIbmPublishApproved : Allow user to set NewIbmPublishApproved
func (options *ReplaceOfferingOptions) SetNewIbmPublishApproved(newIbmPublishApproved bool) *ReplaceOfferingOptions {
	options.NewIbmPublishApproved = core.BoolPtr(newIbmPublishApproved)
	return options
}

// SetNewPublicPublishApproved : Allow user to set NewPublicPublishApproved
func (options *ReplaceOfferingOptions) SetNewPublicPublishApproved(newPublicPublishApproved bool) *ReplaceOfferingOptions {
	options.NewPublicPublishApproved = core.BoolPtr(newPublicPublishApproved)
	return options
}

// SetNewPublicOriginalCrn : Allow user to set NewPublicOriginalCrn
func (options *ReplaceOfferingOptions) SetNewPublicOriginalCrn(newPublicOriginalCrn string) *ReplaceOfferingOptions {
	options.NewPublicOriginalCrn = core.StringPtr(newPublicOriginalCrn)
	return options
}

// SetNewPublishPublicCrn : Allow user to set NewPublishPublicCrn
func (options *ReplaceOfferingOptions) SetNewPublishPublicCrn(newPublishPublicCrn string) *ReplaceOfferingOptions {
	options.NewPublishPublicCrn = core.StringPtr(newPublishPublicCrn)
	return options
}

// SetNewPortalApprovalRecord : Allow user to set NewPortalApprovalRecord
func (options *ReplaceOfferingOptions) SetNewPortalApprovalRecord(newPortalApprovalRecord string) *ReplaceOfferingOptions {
	options.NewPortalApprovalRecord = core.StringPtr(newPortalApprovalRecord)
	return options
}

// SetNewPortalUiURL : Allow user to set NewPortalUiURL
func (options *ReplaceOfferingOptions) SetNewPortalUiURL(newPortalUiURL string) *ReplaceOfferingOptions {
	options.NewPortalUiURL = core.StringPtr(newPortalUiURL)
	return options
}

// SetNewCatalogID : Allow user to set NewCatalogID
func (options *ReplaceOfferingOptions) SetNewCatalogID(newCatalogID string) *ReplaceOfferingOptions {
	options.NewCatalogID = core.StringPtr(newCatalogID)
	return options
}

// SetNewCatalogName : Allow user to set NewCatalogName
func (options *ReplaceOfferingOptions) SetNewCatalogName(newCatalogName string) *ReplaceOfferingOptions {
	options.NewCatalogName = core.StringPtr(newCatalogName)
	return options
}

// SetNewMetadata : Allow user to set NewMetadata
func (options *ReplaceOfferingOptions) SetNewMetadata(newMetadata interface{}) *ReplaceOfferingOptions {
	options.NewMetadata = newMetadata
	return options
}

// SetNewDisclaimer : Allow user to set NewDisclaimer
func (options *ReplaceOfferingOptions) SetNewDisclaimer(newDisclaimer string) *ReplaceOfferingOptions {
	options.NewDisclaimer = core.StringPtr(newDisclaimer)
	return options
}

// SetNewHidden : Allow user to set NewHidden
func (options *ReplaceOfferingOptions) SetNewHidden(newHidden bool) *ReplaceOfferingOptions {
	options.NewHidden = core.BoolPtr(newHidden)
	return options
}

// SetNewProvider : Allow user to set NewProvider
func (options *ReplaceOfferingOptions) SetNewProvider(newProvider string) *ReplaceOfferingOptions {
	options.NewProvider = core.StringPtr(newProvider)
	return options
}

// SetNewRepoInfo : Allow user to set NewRepoInfo
func (options *ReplaceOfferingOptions) SetNewRepoInfo(newRepoInfo *RepoInfo) *ReplaceOfferingOptions {
	options.NewRepoInfo = newRepoInfo
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceOfferingOptions) SetHeaders(param map[string]string) *ReplaceOfferingOptions {
	options.Headers = param
	return options
}

// RepoInfo : Repository info for offerings.
type RepoInfo struct {
	// Token for private repos.
	Token *string `json:"token,omitempty"`

	// Public or enterprise GitHub.
	Type *string `json:"type,omitempty"`
}


// UnmarshalRepoInfo unmarshals an instance of RepoInfo from the specified map of raw messages.
func UnmarshalRepoInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RepoInfo)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Resource : Resource requirements.
type Resource struct {
	// Type of requirement.
	Type *string `json:"type,omitempty"`

	// mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.
	Value *string `json:"value,omitempty"`
}

// Constants associated with the Resource.Type property.
// Type of requirement.
const (
	Resource_Type_Cores = "cores"
	Resource_Type_Disk = "disk"
	Resource_Type_Mem = "mem"
	Resource_Type_Nodes = "nodes"
	Resource_Type_Targetversion = "targetVersion"
)


// UnmarshalResource unmarshals an instance of Resource from the specified map of raw messages.
func UnmarshalResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Resource)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceGroup : Resource group details.
type ResourceGroup struct {
	// Resource Group ID.
	ID *string `json:"id,omitempty"`

	// Provider name, eg. IBM Passport Advantage.
	Name *string `json:"name,omitempty"`

	// Provider CRN.
	Crn *string `json:"crn,omitempty"`

	// Account ID for this Resource Group.
	AccountID *string `json:"account_id,omitempty"`

	// State of this Resource Group.
	State *string `json:"state,omitempty"`

	// Indicates if this Resource Group is active or not.
	Default *bool `json:"default,omitempty"`
}


// UnmarshalResourceGroup unmarshals an instance of ResourceGroup from the specified map of raw messages.
func UnmarshalResourceGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "account_id", &obj.AccountID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default", &obj.Default)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceGroups : Resource groups details.
type ResourceGroups struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting Resource Group objects.
	Resources []ResourceGroup `json:"resources,omitempty"`
}


// UnmarshalResourceGroups unmarshals an instance of ResourceGroups from the specified map of raw messages.
func UnmarshalResourceGroups(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceGroups)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalResourceGroup)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceLocator : Account information.
type ResourceLocator struct {
	// Catalog ID.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Resource ID.
	ResourceID *string `json:"resource_id,omitempty"`
}


// UnmarshalResourceLocator unmarshals an instance of ResourceLocator from the specified map of raw messages.
func UnmarshalResourceLocator(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceLocator)
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_id", &obj.ResourceID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspace : Schematics workspace information.
type SchematicsWorkspace struct {
	// Workspace ID.
	ID *string `json:"id,omitempty"`

	// Workspace name.
	Name *string `json:"name,omitempty"`

	// Workspace types.
	Type []string `json:"type,omitempty"`

	// Workspace description.
	Description *string `json:"description,omitempty"`

	// Workspace tags.
	Tags []string `json:"tags,omitempty"`

	// Workspace creation date and time.
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty"`

	// Email address of user that created the ID.
	CreatedBy *string `json:"created_by,omitempty"`

	// Workspace apply status.
	Status *string `json:"status,omitempty"`

	// Workspace frozen/locked status.
	WorkspaceStatus *SchematicsWorkspaceWorkspaceStatus `json:"workspace_status,omitempty"`

	// Template reference.
	TemplateRef *string `json:"template_ref,omitempty"`

	// Template repository.
	TemplateRepo *SchematicsWorkspaceTemplateRepo `json:"template_repo,omitempty"`

	// Map of template data.
	TemplateData []interface{} `json:"template_data,omitempty"`

	// Data describing runtime.
	RuntimeData *SchematicsWorkspaceRuntimeData `json:"runtime_data,omitempty"`

	// Map of shared data.
	SharedData interface{} `json:"shared_data,omitempty"`

	// Catalog reference.
	CatalogRef *SchematicsWorkspaceCatalogRef `json:"catalog_ref,omitempty"`
}


// UnmarshalSchematicsWorkspace unmarshals an instance of SchematicsWorkspace from the specified map of raw messages.
func UnmarshalSchematicsWorkspace(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspace)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_at", &obj.CreatedAt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_by", &obj.CreatedBy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "workspace_status", &obj.WorkspaceStatus, UnmarshalSchematicsWorkspaceWorkspaceStatus)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_ref", &obj.TemplateRef)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "template_repo", &obj.TemplateRepo, UnmarshalSchematicsWorkspaceTemplateRepo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "template_data", &obj.TemplateData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "runtime_data", &obj.RuntimeData, UnmarshalSchematicsWorkspaceRuntimeData)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "shared_data", &obj.SharedData)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "catalog_ref", &obj.CatalogRef, UnmarshalSchematicsWorkspaceCatalogRef)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspaceCatalogRef : Catalog reference.
type SchematicsWorkspaceCatalogRef struct {
	// Version locator ID.
	ItemID *string `json:"item_id,omitempty"`

	// The name of the offering that generated this Blueprint.
	ItemName *string `json:"item_name,omitempty"`

	// Relative Dashboard URL for content that generated this Blueprint.
	ItemURL *string `json:"item_url,omitempty"`
}


// UnmarshalSchematicsWorkspaceCatalogRef unmarshals an instance of SchematicsWorkspaceCatalogRef from the specified map of raw messages.
func UnmarshalSchematicsWorkspaceCatalogRef(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspaceCatalogRef)
	err = core.UnmarshalPrimitive(m, "item_id", &obj.ItemID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "item_name", &obj.ItemName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "item_url", &obj.ItemURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspaceRuntimeData : Data describing runtime.
type SchematicsWorkspaceRuntimeData struct {
	// Runtime ID.
	ID *string `json:"id,omitempty"`

	// Engine name.
	EngineName *string `json:"engine_name,omitempty"`

	// Engine version.
	EngineVersion *string `json:"engine_version,omitempty"`

	// URL path to state store.
	StateStoreURL *string `json:"state_store_url,omitempty"`

	// URL path to log store.
	LogStoreURL *string `json:"log_store_url,omitempty"`
}


// UnmarshalSchematicsWorkspaceRuntimeData unmarshals an instance of SchematicsWorkspaceRuntimeData from the specified map of raw messages.
func UnmarshalSchematicsWorkspaceRuntimeData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspaceRuntimeData)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_name", &obj.EngineName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "engine_version", &obj.EngineVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state_store_url", &obj.StateStoreURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "log_store_url", &obj.LogStoreURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspaceSearchResult : Result of schematics workspace search.
type SchematicsWorkspaceSearchResult struct {
	// The offset (origin 0) of the first resource in this page of search results.
	Offset *int64 `json:"offset,omitempty"`

	// The maximum number of resources returned in each page of search results.
	Limit *int64 `json:"limit,omitempty"`

	// The overall total number of resources in the search result set.
	TotalCount *int64 `json:"total_count,omitempty"`

	// The number of resources returned in this page of search results.
	ResourceCount *int64 `json:"resource_count,omitempty"`

	// A URL for retrieving the first page of search results.
	First *string `json:"first,omitempty"`

	// A URL for retrieving the last page of search results.
	Last *string `json:"last,omitempty"`

	// A URL for retrieving the previous page of search results.
	Prev *string `json:"prev,omitempty"`

	// A URL for retrieving the next page of search results.
	Next *string `json:"next,omitempty"`

	// Resulting objects.
	Resources []SchematicsWorkspace `json:"resources,omitempty"`
}


// UnmarshalSchematicsWorkspaceSearchResult unmarshals an instance of SchematicsWorkspaceSearchResult from the specified map of raw messages.
func UnmarshalSchematicsWorkspaceSearchResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspaceSearchResult)
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_count", &obj.ResourceCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "first", &obj.First)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last", &obj.Last)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prev", &obj.Prev)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "next", &obj.Next)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "resources", &obj.Resources, UnmarshalSchematicsWorkspace)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspaceTemplateRepo : Template repository.
type SchematicsWorkspaceTemplateRepo struct {
	// The fully qualified path of the tgz used in the deploy.
	RepoURL *string `json:"repo_url,omitempty"`

	// Name of chart.
	ChartName *string `json:"chart_name,omitempty"`

	// Name of script.
	ScriptName *string `json:"script_name,omitempty"`

	// Name of uninstall script.
	UninstallScriptName *string `json:"uninstall_script_name,omitempty"`

	// Name of folder.
	FolderName *string `json:"folder_name,omitempty"`

	// Digest of project.
	RepoShaValue *string `json:"repo_sha_value,omitempty"`
}


// UnmarshalSchematicsWorkspaceTemplateRepo unmarshals an instance of SchematicsWorkspaceTemplateRepo from the specified map of raw messages.
func UnmarshalSchematicsWorkspaceTemplateRepo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspaceTemplateRepo)
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "chart_name", &obj.ChartName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "script_name", &obj.ScriptName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uninstall_script_name", &obj.UninstallScriptName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "folder_name", &obj.FolderName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repo_sha_value", &obj.RepoShaValue)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SchematicsWorkspaceWorkspaceStatus : Workspace frozen/locked status.
type SchematicsWorkspaceWorkspaceStatus struct {
	// Workspace frozen state.
	Frozen *bool `json:"frozen,omitempty"`

	// Workspace locked state.
	Locked *bool `json:"locked,omitempty"`
}


// UnmarshalSchematicsWorkspaceWorkspaceStatus unmarshals an instance of SchematicsWorkspaceWorkspaceStatus from the specified map of raw messages.
func UnmarshalSchematicsWorkspaceWorkspaceStatus(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SchematicsWorkspaceWorkspaceStatus)
	err = core.UnmarshalPrimitive(m, "frozen", &obj.Frozen)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locked", &obj.Locked)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Script : Script information.
type Script struct {
	// Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this
	// version.
	Instructions *string `json:"instructions,omitempty"`

	// Optional script that needs to be run post any pre-condition script.
	Script *string `json:"script,omitempty"`

	// Optional iam permissions that are required on the target cluster to run this script.
	ScriptPermission *string `json:"script_permission,omitempty"`

	// Optional script that if run will remove the installed version.
	DeleteScript *string `json:"delete_script,omitempty"`

	// Optional value indicating if this script is scoped to a namespace or the entire cluster.
	Scope *string `json:"scope,omitempty"`
}


// UnmarshalScript unmarshals an instance of Script from the specified map of raw messages.
func UnmarshalScript(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Script)
	err = core.UnmarshalPrimitive(m, "instructions", &obj.Instructions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "script", &obj.Script)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "script_permission", &obj.ScriptPermission)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "delete_script", &obj.DeleteScript)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scope", &obj.Scope)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SearchLicenseOfferingsOptions : The SearchLicenseOfferings options.
type SearchLicenseOfferingsOptions struct {
	// query, for now only "q=entitlement_key:<some-key>" is supported.
	Q *string `json:"q" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSearchLicenseOfferingsOptions : Instantiate SearchLicenseOfferingsOptions
func (*CatalogManagementV1) NewSearchLicenseOfferingsOptions(q string) *SearchLicenseOfferingsOptions {
	return &SearchLicenseOfferingsOptions{
		Q: core.StringPtr(q),
	}
}

// SetQ : Allow user to set Q
func (options *SearchLicenseOfferingsOptions) SetQ(q string) *SearchLicenseOfferingsOptions {
	options.Q = core.StringPtr(q)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SearchLicenseOfferingsOptions) SetHeaders(param map[string]string) *SearchLicenseOfferingsOptions {
	options.Headers = param
	return options
}

// SearchLicenseVersionsOptions : The SearchLicenseVersions options.
type SearchLicenseVersionsOptions struct {
	// query, for now only "q=entitlement_key:<some-key>" is supported.
	Q *string `json:"q" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSearchLicenseVersionsOptions : Instantiate SearchLicenseVersionsOptions
func (*CatalogManagementV1) NewSearchLicenseVersionsOptions(q string) *SearchLicenseVersionsOptions {
	return &SearchLicenseVersionsOptions{
		Q: core.StringPtr(q),
	}
}

// SetQ : Allow user to set Q
func (options *SearchLicenseVersionsOptions) SetQ(q string) *SearchLicenseVersionsOptions {
	options.Q = core.StringPtr(q)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SearchLicenseVersionsOptions) SetHeaders(param map[string]string) *SearchLicenseVersionsOptions {
	options.Headers = param
	return options
}

// State : Offering state.
type State struct {
	// one of: new, validated, account-published, ibm-published, public-published.
	Current *string `json:"current,omitempty"`

	// Date and time of current request.
	CurrentEntered *strfmt.DateTime `json:"current_entered,omitempty"`

	// one of: new, validated, account-published, ibm-published, public-published.
	Pending *string `json:"pending,omitempty"`

	// Date and time of pending request.
	PendingRequested *strfmt.DateTime `json:"pending_requested,omitempty"`

	// one of: new, validated, account-published, ibm-published, public-published.
	Previous *string `json:"previous,omitempty"`
}


// UnmarshalState unmarshals an instance of State from the specified map of raw messages.
func UnmarshalState(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(State)
	err = core.UnmarshalPrimitive(m, "current", &obj.Current)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "current_entered", &obj.CurrentEntered)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pending", &obj.Pending)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pending_requested", &obj.PendingRequested)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "previous", &obj.Previous)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyndicationAuthorization : Feature information.
type SyndicationAuthorization struct {
	// Array of syndicated namespaces.
	Token *string `json:"token,omitempty"`

	// Date and time last updated.
	LastRun *strfmt.DateTime `json:"last_run,omitempty"`
}


// UnmarshalSyndicationAuthorization unmarshals an instance of SyndicationAuthorization from the specified map of raw messages.
func UnmarshalSyndicationAuthorization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyndicationAuthorization)
	err = core.UnmarshalPrimitive(m, "token", &obj.Token)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_run", &obj.LastRun)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyndicationCluster : Feature information.
type SyndicationCluster struct {
	// Cluster region.
	Region *string `json:"region,omitempty"`

	// Cluster ID.
	ID *string `json:"id,omitempty"`

	// Cluster name.
	Name *string `json:"name,omitempty"`

	// Resource group ID.
	ResourceGroupName *string `json:"resource_group_name,omitempty"`

	// Syndication type.
	Type *string `json:"type,omitempty"`

	// Syndicated namespaces.
	Namespaces []string `json:"namespaces,omitempty"`

	// Syndicated to all namespaces on cluster.
	AllNamespaces *bool `json:"all_namespaces,omitempty"`
}


// UnmarshalSyndicationCluster unmarshals an instance of SyndicationCluster from the specified map of raw messages.
func UnmarshalSyndicationCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyndicationCluster)
	err = core.UnmarshalPrimitive(m, "region", &obj.Region)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_name", &obj.ResourceGroupName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "namespaces", &obj.Namespaces)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "all_namespaces", &obj.AllNamespaces)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyndicationHistory : Feature information.
type SyndicationHistory struct {
	// Array of syndicated namespaces.
	Namespaces []string `json:"namespaces,omitempty"`

	// Array of syndicated namespaces.
	Clusters []SyndicationCluster `json:"clusters,omitempty"`

	// Date and time last syndicated.
	LastRun *strfmt.DateTime `json:"last_run,omitempty"`
}


// UnmarshalSyndicationHistory unmarshals an instance of SyndicationHistory from the specified map of raw messages.
func UnmarshalSyndicationHistory(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyndicationHistory)
	err = core.UnmarshalPrimitive(m, "namespaces", &obj.Namespaces)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalSyndicationCluster)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_run", &obj.LastRun)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SyndicationResource : Feature information.
type SyndicationResource struct {
	// Remove related components.
	RemoveRelatedComponents *bool `json:"remove_related_components,omitempty"`

	// Syndication clusters.
	Clusters []SyndicationCluster `json:"clusters,omitempty"`

	// Feature information.
	History *SyndicationHistory `json:"history,omitempty"`

	// Feature information.
	Authorization *SyndicationAuthorization `json:"authorization,omitempty"`
}


// UnmarshalSyndicationResource unmarshals an instance of SyndicationResource from the specified map of raw messages.
func UnmarshalSyndicationResource(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SyndicationResource)
	err = core.UnmarshalPrimitive(m, "remove_related_components", &obj.RemoveRelatedComponents)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalSyndicationCluster)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "history", &obj.History, UnmarshalSyndicationHistory)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authorization", &obj.Authorization, UnmarshalSyndicationAuthorization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCatalogAccountOptions : The UpdateCatalogAccount options.
type UpdateCatalogAccountOptions struct {
	// Account identification.
	ID *string `json:"id,omitempty"`

	// Filters for account and catalog filters.
	AccountFilters *Filters `json:"account_filters,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCatalogAccountOptions : Instantiate UpdateCatalogAccountOptions
func (*CatalogManagementV1) NewUpdateCatalogAccountOptions() *UpdateCatalogAccountOptions {
	return &UpdateCatalogAccountOptions{}
}

// SetID : Allow user to set ID
func (options *UpdateCatalogAccountOptions) SetID(id string) *UpdateCatalogAccountOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAccountFilters : Allow user to set AccountFilters
func (options *UpdateCatalogAccountOptions) SetAccountFilters(accountFilters *Filters) *UpdateCatalogAccountOptions {
	options.AccountFilters = accountFilters
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCatalogAccountOptions) SetHeaders(param map[string]string) *UpdateCatalogAccountOptions {
	options.Headers = param
	return options
}

// UpdateOfferingIbmOptions : The UpdateOfferingIbm options.
type UpdateOfferingIbmOptions struct {
	// Catalog identification.
	CatalogID *string `json:"catalog_id" validate:"required"`

	// Offering identification.
	OfferingID *string `json:"offering_id" validate:"required"`

	// Type of approval, ibm or public.
	ApprovalType *string `json:"approval_type" validate:"required"`

	// Approve (true) or disapprove (false).
	Approved *string `json:"approved" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateOfferingIbmOptions.ApprovalType property.
// Type of approval, ibm or public.
const (
	UpdateOfferingIbmOptions_ApprovalType_Ibm = "ibm"
	UpdateOfferingIbmOptions_ApprovalType_Public = "public"
)

// Constants associated with the UpdateOfferingIbmOptions.Approved property.
// Approve (true) or disapprove (false).
const (
	UpdateOfferingIbmOptions_Approved_False = "false"
	UpdateOfferingIbmOptions_Approved_True = "true"
)

// NewUpdateOfferingIbmOptions : Instantiate UpdateOfferingIbmOptions
func (*CatalogManagementV1) NewUpdateOfferingIbmOptions(catalogID string, offeringID string, approvalType string, approved string) *UpdateOfferingIbmOptions {
	return &UpdateOfferingIbmOptions{
		CatalogID: core.StringPtr(catalogID),
		OfferingID: core.StringPtr(offeringID),
		ApprovalType: core.StringPtr(approvalType),
		Approved: core.StringPtr(approved),
	}
}

// SetCatalogID : Allow user to set CatalogID
func (options *UpdateOfferingIbmOptions) SetCatalogID(catalogID string) *UpdateOfferingIbmOptions {
	options.CatalogID = core.StringPtr(catalogID)
	return options
}

// SetOfferingID : Allow user to set OfferingID
func (options *UpdateOfferingIbmOptions) SetOfferingID(offeringID string) *UpdateOfferingIbmOptions {
	options.OfferingID = core.StringPtr(offeringID)
	return options
}

// SetApprovalType : Allow user to set ApprovalType
func (options *UpdateOfferingIbmOptions) SetApprovalType(approvalType string) *UpdateOfferingIbmOptions {
	options.ApprovalType = core.StringPtr(approvalType)
	return options
}

// SetApproved : Allow user to set Approved
func (options *UpdateOfferingIbmOptions) SetApproved(approved string) *UpdateOfferingIbmOptions {
	options.Approved = core.StringPtr(approved)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateOfferingIbmOptions) SetHeaders(param map[string]string) *UpdateOfferingIbmOptions {
	options.Headers = param
	return options
}

// Validation : Validation response.
type Validation struct {
	// Date and time of last successful validation.
	Validated *strfmt.DateTime `json:"validated,omitempty"`

	// Date and time of last validation was requested.
	Requested *strfmt.DateTime `json:"requested,omitempty"`

	// Current validation state - <empty>, in_progress, valid, invalid, expired.
	State *string `json:"state,omitempty"`

	// Last operation (e.g. submit_deployment, generate_installer, install_offering.
	LastOperation *string `json:"last_operation,omitempty"`

	// Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.
	Target interface{} `json:"target,omitempty"`
}


// UnmarshalValidation unmarshals an instance of Validation from the specified map of raw messages.
func UnmarshalValidation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Validation)
	err = core.UnmarshalPrimitive(m, "validated", &obj.Validated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "requested", &obj.Requested)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_operation", &obj.LastOperation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ValidationInstallOptions : The ValidationInstall options.
type ValidationInstallOptions struct {
	// A dotted value of `catalogID`.`versionID`.
	VersionLocatorID *string `json:"version_locator_id" validate:"required"`

	// IAM Refresh token.
	XAuthRefreshToken *string `json:"X-Auth-Refresh-Token" validate:"required"`

	// Cluster ID.
	NewClusterID *string `json:"new_cluster_id,omitempty"`

	// Cluster region.
	NewRegion *string `json:"new_region,omitempty"`

	// Kube namespace.
	NewNamespace *string `json:"new_namespace,omitempty"`

	// Object containing Helm chart override values.
	NewOverrideValues interface{} `json:"new_override_values,omitempty"`

	// Entitlement API Key for this offering.
	NewEntitlementApikey *string `json:"new_entitlement_apikey,omitempty"`

	// Schematics workspace configuration.
	NewSchematics *ClusterRequestBodySchematics `json:"new_schematics,omitempty"`

	// Script.
	NewScript *string `json:"new_script,omitempty"`

	// Script ID.
	NewScriptID *string `json:"new_script_id,omitempty"`

	// Version locator ID.
	NewVersionLocatorID *string `json:"new_version_locator_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewValidationInstallOptions : Instantiate ValidationInstallOptions
func (*CatalogManagementV1) NewValidationInstallOptions(versionLocatorID string, xAuthRefreshToken string) *ValidationInstallOptions {
	return &ValidationInstallOptions{
		VersionLocatorID: core.StringPtr(versionLocatorID),
		XAuthRefreshToken: core.StringPtr(xAuthRefreshToken),
	}
}

// SetVersionLocatorID : Allow user to set VersionLocatorID
func (options *ValidationInstallOptions) SetVersionLocatorID(versionLocatorID string) *ValidationInstallOptions {
	options.VersionLocatorID = core.StringPtr(versionLocatorID)
	return options
}

// SetXAuthRefreshToken : Allow user to set XAuthRefreshToken
func (options *ValidationInstallOptions) SetXAuthRefreshToken(xAuthRefreshToken string) *ValidationInstallOptions {
	options.XAuthRefreshToken = core.StringPtr(xAuthRefreshToken)
	return options
}

// SetNewClusterID : Allow user to set NewClusterID
func (options *ValidationInstallOptions) SetNewClusterID(newClusterID string) *ValidationInstallOptions {
	options.NewClusterID = core.StringPtr(newClusterID)
	return options
}

// SetNewRegion : Allow user to set NewRegion
func (options *ValidationInstallOptions) SetNewRegion(newRegion string) *ValidationInstallOptions {
	options.NewRegion = core.StringPtr(newRegion)
	return options
}

// SetNewNamespace : Allow user to set NewNamespace
func (options *ValidationInstallOptions) SetNewNamespace(newNamespace string) *ValidationInstallOptions {
	options.NewNamespace = core.StringPtr(newNamespace)
	return options
}

// SetNewOverrideValues : Allow user to set NewOverrideValues
func (options *ValidationInstallOptions) SetNewOverrideValues(newOverrideValues interface{}) *ValidationInstallOptions {
	options.NewOverrideValues = newOverrideValues
	return options
}

// SetNewEntitlementApikey : Allow user to set NewEntitlementApikey
func (options *ValidationInstallOptions) SetNewEntitlementApikey(newEntitlementApikey string) *ValidationInstallOptions {
	options.NewEntitlementApikey = core.StringPtr(newEntitlementApikey)
	return options
}

// SetNewSchematics : Allow user to set NewSchematics
func (options *ValidationInstallOptions) SetNewSchematics(newSchematics *ClusterRequestBodySchematics) *ValidationInstallOptions {
	options.NewSchematics = newSchematics
	return options
}

// SetNewScript : Allow user to set NewScript
func (options *ValidationInstallOptions) SetNewScript(newScript string) *ValidationInstallOptions {
	options.NewScript = core.StringPtr(newScript)
	return options
}

// SetNewScriptID : Allow user to set NewScriptID
func (options *ValidationInstallOptions) SetNewScriptID(newScriptID string) *ValidationInstallOptions {
	options.NewScriptID = core.StringPtr(newScriptID)
	return options
}

// SetNewVersionLocatorID : Allow user to set NewVersionLocatorID
func (options *ValidationInstallOptions) SetNewVersionLocatorID(newVersionLocatorID string) *ValidationInstallOptions {
	options.NewVersionLocatorID = core.StringPtr(newVersionLocatorID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ValidationInstallOptions) SetHeaders(param map[string]string) *ValidationInstallOptions {
	options.Headers = param
	return options
}

// Version : Offering version information.
type Version struct {
	// Unique ID.
	ID *string `json:"id,omitempty"`

	// Version's CRN.
	Crn *string `json:"crn,omitempty"`

	// Version of content type.
	Version *string `json:"version,omitempty"`

	// hash of the content.
	Sha *string `json:"sha,omitempty"`

	// The date and time this version was created.
	Created *strfmt.DateTime `json:"created,omitempty"`

	// The date and time this version was last updated.
	Updated *strfmt.DateTime `json:"updated,omitempty"`

	// Offering ID.
	OfferingID *string `json:"offering_id,omitempty"`

	// Catalog ID.
	CatalogID *string `json:"catalog_id,omitempty"`

	// Kind ID.
	KindID *string `json:"kind_id,omitempty"`

	// List of tags associated with this catalog.
	Tags []string `json:"tags,omitempty"`

	// Content's repo URL.
	RepoURL *string `json:"repo_url,omitempty"`

	// Content's source URL (e.g git repo).
	SourceURL *string `json:"source_url,omitempty"`

	// File used to on-board this version.
	TgzURL *string `json:"tgz_url,omitempty"`

	// List of user solicited overrides.
	Configuration []Configuration `json:"configuration,omitempty"`

	// Open ended metadata information.
	Metadata interface{} `json:"metadata,omitempty"`

	// Validation response.
	Validation *Validation `json:"validation,omitempty"`

	// Resource requirements.
	RequiredResources *Resource `json:"required_resources,omitempty"`

	// Denotes if single instance can be deployed to a given cluster.
	SingleInstance *bool `json:"single_instance,omitempty"`

	// Script information.
	Install *Script `json:"install,omitempty"`

	// Optional pre-install instructions.
	PreInstall []Script `json:"pre_install,omitempty"`

	// Entitlement license info.
	Entitlement *VersionEntitlement `json:"entitlement,omitempty"`

	// List of licenses the product was built with.
	Licenses []License `json:"licenses,omitempty"`

	// If set, denotes a url to a YAML file with list of container images used by this version.
	ImageManifestURL *string `json:"image_manifest_url,omitempty"`

	// read only field, indicating if this version is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`

	// Version of the package used to create this version.
	PackageVersion *string `json:"package_version,omitempty"`

	// Offering state.
	State *State `json:"state,omitempty"`

	// Version locator ID.
	VersionLocator *string `json:"version_locator,omitempty"`

	// Version locator ID.
	ConsoleURL *string `json:"console_url,omitempty"`

	// Version's description.
	LongDescription *string `json:"long_description,omitempty"`

	// Version's description.
	WhitelistedAccounts []string `json:"whitelisted_accounts,omitempty"`
}


// UnmarshalVersion unmarshals an instance of Version from the specified map of raw messages.
func UnmarshalVersion(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Version)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sha", &obj.Sha)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "updated", &obj.Updated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offering_id", &obj.OfferingID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "catalog_id", &obj.CatalogID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "kind_id", &obj.KindID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tags", &obj.Tags)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "repo_url", &obj.RepoURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "source_url", &obj.SourceURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tgz_url", &obj.TgzURL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "configuration", &obj.Configuration, UnmarshalConfiguration)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "metadata", &obj.Metadata)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "validation", &obj.Validation, UnmarshalValidation)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_resources", &obj.RequiredResources, UnmarshalResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "single_instance", &obj.SingleInstance)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "install", &obj.Install, UnmarshalScript)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pre_install", &obj.PreInstall, UnmarshalScript)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "entitlement", &obj.Entitlement, UnmarshalVersionEntitlement)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "licenses", &obj.Licenses, UnmarshalLicense)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_manifest_url", &obj.ImageManifestURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deprecated", &obj.Deprecated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "package_version", &obj.PackageVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "state", &obj.State, UnmarshalState)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version_locator", &obj.VersionLocator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "console_url", &obj.ConsoleURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "long_description", &obj.LongDescription)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "whitelisted_accounts", &obj.WhitelistedAccounts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VersionEntitlement : Entitlement license info.
type VersionEntitlement struct {
	// Provider name.
	ProviderName *string `json:"provider_name,omitempty"`

	// Provider ID.
	ProviderID *string `json:"provider_id,omitempty"`

	// Product ID.
	ProductID *string `json:"product_id,omitempty"`

	// list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.
	PartNumbers []string `json:"part_numbers,omitempty"`

	// Image repository name.
	ImageRepoName *string `json:"image_repo_name,omitempty"`
}


// UnmarshalVersionEntitlement unmarshals an instance of VersionEntitlement from the specified map of raw messages.
func UnmarshalVersionEntitlement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VersionEntitlement)
	err = core.UnmarshalPrimitive(m, "provider_name", &obj.ProviderName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider_id", &obj.ProviderID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "product_id", &obj.ProductID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_numbers", &obj.PartNumbers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "image_repo_name", &obj.ImageRepoName)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VersionUpdateDescriptor : Indicates if the current version can be upgraded to the version identified by the descriptor.
type VersionUpdateDescriptor struct {
	// Account information.
	VersionLocator *ResourceLocator `json:"version_locator,omitempty"`

	// the version number of this version.
	Version *string `json:"version,omitempty"`

	// Offering state.
	State *State `json:"state,omitempty"`

	// Array of required resources.
	RequiredResources []Resource `json:"required_resources,omitempty"`

	// Version of package.
	PackageVersion *string `json:"package_version,omitempty"`

	// true if the current version can be upgraded to this version, false otherwise.
	CanUpdate *bool `json:"can_update,omitempty"`

	// If can_update is false, this map will contain messages for each failed check, otherwise it will be omitted.
	// Possible keys include nodes, cores, mem, disk, targetVersion, and install-permission-check.
	Messages interface{} `json:"messages,omitempty"`
}


// UnmarshalVersionUpdateDescriptor unmarshals an instance of VersionUpdateDescriptor from the specified map of raw messages.
func UnmarshalVersionUpdateDescriptor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VersionUpdateDescriptor)
	err = core.UnmarshalModel(m, "version_locator", &obj.VersionLocator, UnmarshalResourceLocator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "version", &obj.Version)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "state", &obj.State, UnmarshalState)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "required_resources", &obj.RequiredResources, UnmarshalResource)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "package_version", &obj.PackageVersion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "can_update", &obj.CanUpdate)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
