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

package globalcatalogv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/globalcatalogv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"
)

var _ = Describe(`GlobalCatalogV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions) - Operation response error`, func() {
		listCatalogEntriesPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCatalogEntriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalogEntries with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Complete = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCatalogEntries(listCatalogEntriesOptions *ListCatalogEntriesOptions)`, func() {
		listCatalogEntriesPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCatalogEntriesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"mapKey": {"anyKey": "anyValue"}}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCrn", "service_crn": "ServiceCrn", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "anyValue", "url": "anyValue", "children_url": "anyValue", "geo_tags": "anyValue", "pricing_tags": "anyValue", "created": "anyValue", "updated": "anyValue"}]}`)
				}))
			})
			It(`Invoke ListCatalogEntries successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListCatalogEntries(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Complete = core.StringPtr("testString")
 				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListCatalogEntries with error: Operation request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := new(globalcatalogv1.ListCatalogEntriesOptions)
				listCatalogEntriesOptionsModel.Account = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Include = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Q = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.SortBy = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Descending = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Languages = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Complete = core.StringPtr("testString")
				listCatalogEntriesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListCatalogEntries(listCatalogEntriesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions) - Operation response error`, func() {
		createCatalogEntryPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCatalogEntryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCatalogEntry with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateCatalogEntry(createCatalogEntryOptions *CreateCatalogEntryOptions)`, func() {
		createCatalogEntryPath := "/"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createCatalogEntryPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"name": "Name", "kind": "service", "overview_ui": {}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"mapKey": {"anyKey": "anyValue"}}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCrn", "service_crn": "ServiceCrn", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "anyValue", "url": "anyValue", "children_url": "anyValue", "geo_tags": "anyValue", "pricing_tags": "anyValue", "created": "anyValue", "updated": "anyValue"}`)
				}))
			})
			It(`Invoke CreateCatalogEntry successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
 				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateCatalogEntry with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsModel := new(globalcatalogv1.CreateCatalogEntryOptions)
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				createCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateCatalogEntry(createCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCatalogEntryOptions model with no property values
				createCatalogEntryOptionsModelNew := new(globalcatalogv1.CreateCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateCatalogEntry(createCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions) - Operation response error`, func() {
		getCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCatalogEntryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalogEntry with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetCatalogEntry(getCatalogEntryOptions *GetCatalogEntryOptions)`, func() {
		getCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCatalogEntryPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["depth"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Name", "kind": "service", "overview_ui": {}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"mapKey": {"anyKey": "anyValue"}}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCrn", "service_crn": "ServiceCrn", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "anyValue", "url": "anyValue", "children_url": "anyValue", "geo_tags": "anyValue", "pricing_tags": "anyValue", "created": "anyValue", "updated": "anyValue"}`)
				}))
			})
			It(`Invoke GetCatalogEntry successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
 				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCatalogEntry with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCatalogEntryOptions model
				getCatalogEntryOptionsModel := new(globalcatalogv1.GetCatalogEntryOptions)
				getCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Include = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Languages = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Complete = core.StringPtr("testString")
				getCatalogEntryOptionsModel.Depth = core.Int64Ptr(int64(38))
				getCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCatalogEntry(getCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogEntryOptions model with no property values
				getCatalogEntryOptionsModelNew := new(globalcatalogv1.GetCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetCatalogEntry(getCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions) - Operation response error`, func() {
		updateCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCatalogEntryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["move"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalogEntry with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateCatalogEntry(updateCatalogEntryOptions *UpdateCatalogEntryOptions)`, func() {
		updateCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateCatalogEntryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["move"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"name": "Name", "kind": "service", "overview_ui": {}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"mapKey": {"anyKey": "anyValue"}}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCrn", "service_crn": "ServiceCrn", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "anyValue", "url": "anyValue", "children_url": "anyValue", "geo_tags": "anyValue", "pricing_tags": "anyValue", "created": "anyValue", "updated": "anyValue"}`)
				}))
			})
			It(`Invoke UpdateCatalogEntry successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
 				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateCatalogEntry with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				overviewUiModel.SetProperty("foo", overviewModel)

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Kind = core.StringPtr("service")
				updateCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.Images = imageModel
				updateCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Tags = []string{"testString"}
				updateCatalogEntryOptionsModel.Provider = providerModel
				updateCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Active = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.Metadata = objectMetadataSetModel
				updateCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Move = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateCatalogEntry(updateCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogEntryOptions model with no property values
				updateCatalogEntryOptionsModelNew := new(globalcatalogv1.UpdateCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateCatalogEntry(updateCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteCatalogEntry(deleteCatalogEntryOptions *DeleteCatalogEntryOptions)`, func() {
		deleteCatalogEntryPath := "/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteCatalogEntryPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))


					// TODO: Add check for force query parameter

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCatalogEntry successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCatalogEntryOptions model
				deleteCatalogEntryOptionsModel := new(globalcatalogv1.DeleteCatalogEntryOptions)
				deleteCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Force = core.BoolPtr(true)
 				deleteCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteCatalogEntry(deleteCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCatalogEntry with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteCatalogEntryOptions model
				deleteCatalogEntryOptionsModel := new(globalcatalogv1.DeleteCatalogEntryOptions)
				deleteCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				deleteCatalogEntryOptionsModel.Force = core.BoolPtr(true)
				deleteCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteCatalogEntry(deleteCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCatalogEntryOptions model with no property values
				deleteCatalogEntryOptionsModelNew := new(globalcatalogv1.DeleteCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteCatalogEntry(deleteCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions) - Operation response error`, func() {
		getChildObjectsPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getChildObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetChildObjects with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.StringPtr("testString")
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetChildObjects(getChildObjectsOptions *GetChildObjectsOptions)`, func() {
		getChildObjectsPath := "/testString/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getChildObjectsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["include"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["q"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sort-by"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["descending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["languages"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["complete"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"name": "Name", "kind": "service", "overview_ui": {}, "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "active": true, "metadata": {"rc_compatible": true, "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "bindable": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false, "cf_guid": {"mapKey": "Inner"}}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": {"mapKey": "Inner"}}, "alias": {"type": "Type", "plan_id": "PlanID"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"mapKey": "Inner"}}, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL", "dashboard_url": "DashboardURL", "registration_url": "RegistrationURL", "apidocsurl": "Apidocsurl"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00", "hidden": true, "hide_lite_metering": true, "no_upgrade_next_step": false}, "compliance": ["Compliance"], "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"controller_url": "ControllerURL", "broker_url": "BrokerURL", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "api_endpoint": {"mapKey": "Inner"}}, "original_name": "OriginalName", "version": "Version", "other": {"mapKey": {"anyKey": "anyValue"}}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "deployment": {"location": "Location", "location_url": "LocationURL", "original_location": "OriginalLocation", "target_crn": "TargetCrn", "service_crn": "ServiceCrn", "mccp_id": "MccpID", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false, "target_network": "TargetNetwork"}}, "id": "ID", "catalog_crn": "anyValue", "url": "anyValue", "children_url": "anyValue", "geo_tags": "anyValue", "pricing_tags": "anyValue", "created": "anyValue", "updated": "anyValue"}]}`)
				}))
			})
			It(`Invoke GetChildObjects successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetChildObjects(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.StringPtr("testString")
 				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetChildObjects with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetChildObjectsOptions model
				getChildObjectsOptionsModel := new(globalcatalogv1.GetChildObjectsOptions)
				getChildObjectsOptionsModel.ID = core.StringPtr("testString")
				getChildObjectsOptionsModel.Kind = core.StringPtr("testString")
				getChildObjectsOptionsModel.Account = core.StringPtr("testString")
				getChildObjectsOptionsModel.Include = core.StringPtr("testString")
				getChildObjectsOptionsModel.Q = core.StringPtr("testString")
				getChildObjectsOptionsModel.SortBy = core.StringPtr("testString")
				getChildObjectsOptionsModel.Descending = core.StringPtr("testString")
				getChildObjectsOptionsModel.Languages = core.StringPtr("testString")
				getChildObjectsOptionsModel.Complete = core.StringPtr("testString")
				getChildObjectsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetChildObjects(getChildObjectsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetChildObjectsOptions model with no property values
				getChildObjectsOptionsModelNew := new(globalcatalogv1.GetChildObjectsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetChildObjects(getChildObjectsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RestoreCatalogEntry(restoreCatalogEntryOptions *RestoreCatalogEntryOptions)`, func() {
		restoreCatalogEntryPath := "/testString/restore"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(restoreCatalogEntryPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke RestoreCatalogEntry successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.RestoreCatalogEntry(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RestoreCatalogEntryOptions model
				restoreCatalogEntryOptionsModel := new(globalcatalogv1.RestoreCatalogEntryOptions)
				restoreCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Account = core.StringPtr("testString")
 				restoreCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.RestoreCatalogEntry(restoreCatalogEntryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RestoreCatalogEntry with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RestoreCatalogEntryOptions model
				restoreCatalogEntryOptionsModel := new(globalcatalogv1.RestoreCatalogEntryOptions)
				restoreCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Account = core.StringPtr("testString")
				restoreCatalogEntryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.RestoreCatalogEntry(restoreCatalogEntryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RestoreCatalogEntryOptions model with no property values
				restoreCatalogEntryOptionsModelNew := new(globalcatalogv1.RestoreCatalogEntryOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.RestoreCatalogEntry(restoreCatalogEntryOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetVisibility(getVisibilityOptions *GetVisibilityOptions) - Operation response error`, func() {
		getVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getVisibilityPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVisibility with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetVisibility(getVisibilityOptions *GetVisibilityOptions)`, func() {
		getVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getVisibilityPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}`)
				}))
			})
			It(`Invoke GetVisibility successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetVisibility(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
 				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetVisibility with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetVisibilityOptions model
				getVisibilityOptionsModel := new(globalcatalogv1.GetVisibilityOptions)
				getVisibilityOptionsModel.ID = core.StringPtr("testString")
				getVisibilityOptionsModel.Account = core.StringPtr("testString")
				getVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetVisibility(getVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVisibilityOptions model with no property values
				getVisibilityOptionsModelNew := new(globalcatalogv1.GetVisibilityOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetVisibility(getVisibilityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateVisibility(updateVisibilityOptions *UpdateVisibilityOptions)`, func() {
		updateVisibilityPath := "/testString/visibility"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateVisibilityPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateVisibility successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateVisibility(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Extendable = core.BoolPtr(true)
				updateVisibilityOptionsModel.Include = visibilityDetailModel
				updateVisibilityOptionsModel.Exclude = visibilityDetailModel
				updateVisibilityOptionsModel.Account = core.StringPtr("testString")
 				updateVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateVisibility(updateVisibilityOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateVisibility with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Extendable = core.BoolPtr(true)
				updateVisibilityOptionsModel.Include = visibilityDetailModel
				updateVisibilityOptionsModel.Exclude = visibilityDetailModel
				updateVisibilityOptionsModel.Account = core.StringPtr("testString")
				updateVisibilityOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateVisibility(updateVisibilityOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateVisibilityOptions model with no property values
				updateVisibilityOptionsModelNew := new(globalcatalogv1.UpdateVisibilityOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateVisibility(updateVisibilityOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetPricing(getPricingOptions *GetPricingOptions) - Operation response error`, func() {
		getPricingPath := "/testString/pricing"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPricingPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPricing with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetPricing(getPricingOptions *GetPricingOptions)`, func() {
		getPricingPath := "/testString/pricing"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPricingPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "unit": "Unit", "amount": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"part_ref": "PartRef", "metric_id": "MetricID", "tier_model": "TierModel", "charge_unit": "ChargeUnit", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "display_cap": 10, "effective_from": "2019-01-01T12:00:00", "effective_until": "2019-01-01T12:00:00", "amounts": [{"country": "Country", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetPricing successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPricing(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
 				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPricing with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPricingOptions model
				getPricingOptionsModel := new(globalcatalogv1.GetPricingOptions)
				getPricingOptionsModel.ID = core.StringPtr("testString")
				getPricingOptionsModel.Account = core.StringPtr("testString")
				getPricingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetPricing(getPricingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPricingOptions model with no property values
				getPricingOptionsModelNew := new(globalcatalogv1.GetPricingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetPricing(getPricingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions) - Operation response error`, func() {
		getAuditLogsPath := "/testString/logs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAuditLogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ascending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["startat"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAuditLogs with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("testString")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAuditLogs(getAuditLogsOptions *GetAuditLogsOptions)`, func() {
		getAuditLogsPath := "/testString/logs"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getAuditLogsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["ascending"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["startat"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["_offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"offset": 6, "limit": 5, "count": 5, "resource_count": 13, "first": "First", "last": "Last", "prev": "Prev", "next": "Next", "resources": [{"id": "ID", "effective": {"restrictions": "Restrictions", "owner": "Owner", "extendable": true, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}, "time": "2019-01-01T12:00:00", "who_id": "WhoID", "who_name": "WhoName", "who_email": "WhoEmail", "instance": "Instance", "gid": "Gid", "type": "Type", "message": "Message", "data": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke GetAuditLogs successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAuditLogs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("testString")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(38))
 				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetAuditLogs with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetAuditLogsOptions model
				getAuditLogsOptionsModel := new(globalcatalogv1.GetAuditLogsOptions)
				getAuditLogsOptionsModel.ID = core.StringPtr("testString")
				getAuditLogsOptionsModel.Account = core.StringPtr("testString")
				getAuditLogsOptionsModel.Ascending = core.StringPtr("testString")
				getAuditLogsOptionsModel.Startat = core.StringPtr("testString")
				getAuditLogsOptionsModel.Offset = core.Int64Ptr(int64(38))
				getAuditLogsOptionsModel.Limit = core.Int64Ptr(int64(38))
				getAuditLogsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetAuditLogs(getAuditLogsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAuditLogsOptions model with no property values
				getAuditLogsOptionsModelNew := new(globalcatalogv1.GetAuditLogsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetAuditLogs(getAuditLogsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "https://globalcatalogv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_URL": "https://globalcatalogv1/api",
				"GLOBAL_CATALOG_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_CATALOG_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1UsingExternalConfig(&globalcatalogv1.GlobalCatalogV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListArtifacts(listArtifactsOptions *ListArtifactsOptions) - Operation response error`, func() {
		listArtifactsPath := "/testString/artifacts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			 		defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listArtifactsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListArtifacts with error: Operation response processing error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListArtifacts(listArtifactsOptions *ListArtifactsOptions)`, func() {
		listArtifactsPath := "/testString/artifacts"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listArtifactsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"count": 5, "resources": [{"name": "Name", "updated": "2019-01-01T12:00:00", "url": "URL", "etag": "Etag", "size": 4}]}`)
				}))
			})
			It(`Invoke ListArtifacts successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListArtifacts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
 				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListArtifacts with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListArtifactsOptions model
				listArtifactsOptionsModel := new(globalcatalogv1.ListArtifactsOptions)
				listArtifactsOptionsModel.ObjectID = core.StringPtr("testString")
				listArtifactsOptionsModel.Account = core.StringPtr("testString")
				listArtifactsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListArtifacts(listArtifactsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListArtifactsOptions model with no property values
				listArtifactsOptionsModelNew := new(globalcatalogv1.ListArtifactsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListArtifacts(listArtifactsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetArtifact(getArtifactOptions *GetArtifactOptions)`, func() {
		getArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getArtifactPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "*/*")
					res.WriteHeader(200)
					fmt.Fprintf(res, `Contents of response byte-stream...`)
				}))
			})
			It(`Invoke GetArtifact successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
 				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetArtifact with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Accept = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetArtifactOptions model with no property values
				getArtifactOptionsModelNew := new(globalcatalogv1.GetArtifactOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetArtifact(getArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UploadArtifact(uploadArtifactOptions *UploadArtifactOptions)`, func() {
		uploadArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(uploadArtifactPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Content-Type"]).ToNot(BeNil())
					Expect(req.Header["Content-Type"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UploadArtifact successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UploadArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UploadArtifactOptions model
				uploadArtifactOptionsModel := new(globalcatalogv1.UploadArtifactOptions)
				uploadArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				uploadArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				uploadArtifactOptionsModel.Artifact = CreateMockReader("This is a mock file.")
				uploadArtifactOptionsModel.ContentType = core.StringPtr("testString")
				uploadArtifactOptionsModel.Account = core.StringPtr("testString")
 				uploadArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UploadArtifact(uploadArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UploadArtifact with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UploadArtifactOptions model
				uploadArtifactOptionsModel := new(globalcatalogv1.UploadArtifactOptions)
				uploadArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				uploadArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				uploadArtifactOptionsModel.Artifact = CreateMockReader("This is a mock file.")
				uploadArtifactOptionsModel.ContentType = core.StringPtr("testString")
				uploadArtifactOptionsModel.Account = core.StringPtr("testString")
				uploadArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UploadArtifact(uploadArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UploadArtifactOptions model with no property values
				uploadArtifactOptionsModelNew := new(globalcatalogv1.UploadArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UploadArtifact(uploadArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteArtifact(deleteArtifactOptions *DeleteArtifactOptions)`, func() {
		deleteArtifactPath := "/testString/artifacts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteArtifactPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteArtifact successfully`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteArtifactOptions model
				deleteArtifactOptionsModel := new(globalcatalogv1.DeleteArtifactOptions)
				deleteArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				deleteArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				deleteArtifactOptionsModel.Account = core.StringPtr("testString")
 				deleteArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteArtifact(deleteArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteArtifact with error: Operation validation and request error`, func() {
				testService, testServiceErr := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteArtifactOptions model
				deleteArtifactOptionsModel := new(globalcatalogv1.DeleteArtifactOptions)
				deleteArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				deleteArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				deleteArtifactOptionsModel.Account = core.StringPtr("testString")
				deleteArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteArtifact(deleteArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteArtifactOptions model with no property values
				deleteArtifactOptionsModelNew := new(globalcatalogv1.DeleteArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteArtifact(deleteArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			testService, _ := globalcatalogv1.NewGlobalCatalogV1(&globalcatalogv1.GlobalCatalogV1Options{
				URL:           "http://globalcatalogv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateCatalogEntryOptions successfully`, func() {
				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				Expect(bulletsModel).ToNot(BeNil())
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))
				Expect(bulletsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				Expect(priceModel).ToNot(BeNil())
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)
				Expect(priceModel.QuantityTier).To(Equal(core.Int64Ptr(int64(38))))
				Expect(priceModel.Price).To(Equal(core.Float64Ptr(72.5)))

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				Expect(uiMetaMediaModel).ToNot(BeNil())
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel
				Expect(uiMetaMediaModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.ThumbnailURL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Source).To(Equal(bulletsModel))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				Expect(amountModel).ToNot(BeNil())
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}
				Expect(amountModel.Country).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Currency).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Prices).To(Equal([]globalcatalogv1.Price{*priceModel}))

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				Expect(stringsModel).ToNot(BeNil())
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")
				Expect(stringsModel.Bullets).To(Equal([]globalcatalogv1.Bullets{*bulletsModel}))
				Expect(stringsModel.Media).To(Equal([]globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}))
				Expect(stringsModel.NotCreatableMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.NotCreatableRobotMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.DeprecationWarning).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.PopupWarningMessage).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.Instruction).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				Expect(brokerModel).ToNot(BeNil())
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")
				Expect(brokerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(brokerModel.Guid).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				Expect(drMetaDataModel).ToNot(BeNil())
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")
				Expect(drMetaDataModel.Dr).To(Equal(core.BoolPtr(true)))
				Expect(drMetaDataModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				Expect(i18NModel).ToNot(BeNil())
				i18NModel.SetProperty("foo", stringsModel)
				Expect(i18NModel.GetProperty("foo")).To(Equal(stringsModel))
				Expect(i18NModel.GetProperties()).ToNot(BeEmpty())

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				Expect(sourceMetaDataModel).ToNot(BeNil())
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(sourceMetaDataModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				Expect(startingPriceModel).ToNot(BeNil())
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}
				Expect(startingPriceModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Unit).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Amount).To(Equal([]globalcatalogv1.Amount{*amountModel}))

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				Expect(urlsModel).ToNot(BeNil())
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")
				Expect(urlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.ApiURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CreateURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.SdkDownloadURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DeprecationDocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.RegistrationURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.Apidocsurl).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				Expect(aliasMetaDataModel).ToNot(BeNil())
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")
				Expect(aliasMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(aliasMetaDataModel.PlanID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				Expect(cfMetaDataModel).ToNot(BeNil())
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)
				Expect(cfMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.UniqueApiKey).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Provisionable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Requires).To(Equal([]string{"testString"}))
				Expect(cfMetaDataModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.State).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(cfMetaDataModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CfGuid).To(Equal(make(map[string]string)))

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				Expect(callbacksModel).ToNot(BeNil())
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)
				Expect(callbacksModel.ControllerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerProxyURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDataURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabExtURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApi).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApp).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ApiEndpoint).To(Equal(make(map[string]string)))

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				Expect(deploymentBaseModel).ToNot(BeNil())
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")
				Expect(deploymentBaseModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.LocationURL).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.OriginalLocation).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.TargetCrn).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.ServiceCrn).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.MccpID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.Broker).To(Equal(brokerModel))
				Expect(deploymentBaseModel.SupportsRcMigration).To(Equal(core.BoolPtr(true)))
				Expect(deploymentBaseModel.TargetNetwork).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				Expect(overviewModel).ToNot(BeNil())
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")
				Expect(overviewModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.FeaturedDescription).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				Expect(planMetaDataModel).ToNot(BeNil())
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)
				Expect(planMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.Reservable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AllowInternalUsers).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(planMetaDataModel.SingleScopeInstance).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.CfGuid).To(Equal(make(map[string]string)))

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				Expect(pricingSetModel).ToNot(BeNil())
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel
				Expect(pricingSetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.Origin).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.StartingPrice).To(Equal(startingPriceModel))

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				Expect(slaMetaDataModel).ToNot(BeNil())
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel
				Expect(slaMetaDataModel.Terms).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Tenancy).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Provisioning).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Responsiveness).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Dr).To(Equal(drMetaDataModel))

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				Expect(templateMetaDataModel).ToNot(BeNil())
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)
				Expect(templateMetaDataModel.Services).To(Equal([]string{"testString"}))
				Expect(templateMetaDataModel.DefaultMemory).To(Equal(core.Int64Ptr(int64(38))))
				Expect(templateMetaDataModel.StartCmd).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Source).To(Equal(sourceMetaDataModel))
				Expect(templateMetaDataModel.RuntimeCatalogID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.CfRuntimeID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.ExecutableFile).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Buildpack).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.EnvironmentVariables).To(Equal(make(map[string]string)))

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				Expect(uiMetaDataModel).ToNot(BeNil())
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				Expect(uiMetaDataModel.Strings).To(Equal(i18NModel))
				Expect(uiMetaDataModel.Urls).To(Equal(urlsModel))
				Expect(uiMetaDataModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.EmbeddableDashboardFullWidth).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NavigationOrder).To(Equal([]string{"testString"}))
				Expect(uiMetaDataModel.NotCreatable).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.PrimaryOfferingID).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.AccessibleDuringProvision).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.SideBySideIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uiMetaDataModel.EndOfServiceTime).To(Equal(CreateMockDateTime()))
				Expect(uiMetaDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.HideLiteMetering).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NoUpgradeNextStep).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				Expect(imageModel).ToNot(BeNil())
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")
				Expect(imageModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.SmallImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.MediumImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.FeatureImage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				Expect(objectMetadataSetModel).ToNot(BeNil())
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel
				Expect(objectMetadataSetModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(objectMetadataSetModel.Service).To(Equal(cfMetaDataModel))
				Expect(objectMetadataSetModel.Plan).To(Equal(planMetaDataModel))
				Expect(objectMetadataSetModel.Alias).To(Equal(aliasMetaDataModel))
				Expect(objectMetadataSetModel.Template).To(Equal(templateMetaDataModel))
				Expect(objectMetadataSetModel.Ui).To(Equal(uiMetaDataModel))
				Expect(objectMetadataSetModel.Compliance).To(Equal([]string{"testString"}))
				Expect(objectMetadataSetModel.Sla).To(Equal(slaMetaDataModel))
				Expect(objectMetadataSetModel.Callbacks).To(Equal(callbacksModel))
				Expect(objectMetadataSetModel.OriginalName).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Other).To(Equal(make(map[string]interface{})))
				Expect(objectMetadataSetModel.Pricing).To(Equal(pricingSetModel))
				Expect(objectMetadataSetModel.Deployment).To(Equal(deploymentBaseModel))

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				Expect(overviewUiModel).ToNot(BeNil())
				overviewUiModel.SetProperty("foo", overviewModel)
				Expect(overviewUiModel.GetProperty("foo")).To(Equal(overviewModel))
				Expect(overviewUiModel.GetProperties()).ToNot(BeEmpty())

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				Expect(providerModel).ToNot(BeNil())
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")
				Expect(providerModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.SupportEmail).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Phone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateCatalogEntryOptions model
				createCatalogEntryOptionsName := "testString"
				createCatalogEntryOptionsKind := "service"
				var createCatalogEntryOptionsOverviewUi *globalcatalogv1.OverviewUI = nil
				var createCatalogEntryOptionsImages *globalcatalogv1.Image = nil
				createCatalogEntryOptionsDisabled := true
				createCatalogEntryOptionsTags := []string{"testString"}
				var createCatalogEntryOptionsProvider *globalcatalogv1.Provider = nil
				createCatalogEntryOptionsID := "testString"
				createCatalogEntryOptionsModel := testService.NewCreateCatalogEntryOptions(createCatalogEntryOptionsName, createCatalogEntryOptionsKind, createCatalogEntryOptionsOverviewUi, createCatalogEntryOptionsImages, createCatalogEntryOptionsDisabled, createCatalogEntryOptionsTags, createCatalogEntryOptionsProvider, createCatalogEntryOptionsID)
				createCatalogEntryOptionsModel.SetName("testString")
				createCatalogEntryOptionsModel.SetKind("service")
				createCatalogEntryOptionsModel.SetOverviewUi(overviewUiModel)
				createCatalogEntryOptionsModel.SetImages(imageModel)
				createCatalogEntryOptionsModel.SetDisabled(true)
				createCatalogEntryOptionsModel.SetTags([]string{"testString"})
				createCatalogEntryOptionsModel.SetProvider(providerModel)
				createCatalogEntryOptionsModel.SetID("testString")
				createCatalogEntryOptionsModel.SetParentID("testString")
				createCatalogEntryOptionsModel.SetGroup(true)
				createCatalogEntryOptionsModel.SetActive(true)
				createCatalogEntryOptionsModel.SetMetadata(objectMetadataSetModel)
				createCatalogEntryOptionsModel.SetAccount("testString")
				createCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(createCatalogEntryOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(createCatalogEntryOptionsModel.OverviewUi).To(Equal(overviewUiModel))
				Expect(createCatalogEntryOptionsModel.Images).To(Equal(imageModel))
				Expect(createCatalogEntryOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createCatalogEntryOptionsModel.Provider).To(Equal(providerModel))
				Expect(createCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(createCatalogEntryOptionsModel.Metadata).To(Equal(objectMetadataSetModel))
				Expect(createCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(createCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteArtifactOptions successfully`, func() {
				// Construct an instance of the DeleteArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				deleteArtifactOptionsModel := testService.NewDeleteArtifactOptions(objectID, artifactID)
				deleteArtifactOptionsModel.SetObjectID("testString")
				deleteArtifactOptionsModel.SetArtifactID("testString")
				deleteArtifactOptionsModel.SetAccount("testString")
				deleteArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteArtifactOptionsModel).ToNot(BeNil())
				Expect(deleteArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(deleteArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCatalogEntryOptions successfully`, func() {
				// Construct an instance of the DeleteCatalogEntryOptions model
				id := "testString"
				deleteCatalogEntryOptionsModel := testService.NewDeleteCatalogEntryOptions(id)
				deleteCatalogEntryOptionsModel.SetID("testString")
				deleteCatalogEntryOptionsModel.SetAccount("testString")
				deleteCatalogEntryOptionsModel.SetForce(true)
				deleteCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(deleteCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(deleteCatalogEntryOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(deleteCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetArtifactOptions successfully`, func() {
				// Construct an instance of the GetArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				getArtifactOptionsModel := testService.NewGetArtifactOptions(objectID, artifactID)
				getArtifactOptionsModel.SetObjectID("testString")
				getArtifactOptionsModel.SetArtifactID("testString")
				getArtifactOptionsModel.SetAccept("testString")
				getArtifactOptionsModel.SetAccount("testString")
				getArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getArtifactOptionsModel).ToNot(BeNil())
				Expect(getArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Accept).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAuditLogsOptions successfully`, func() {
				// Construct an instance of the GetAuditLogsOptions model
				id := "testString"
				getAuditLogsOptionsModel := testService.NewGetAuditLogsOptions(id)
				getAuditLogsOptionsModel.SetID("testString")
				getAuditLogsOptionsModel.SetAccount("testString")
				getAuditLogsOptionsModel.SetAscending("testString")
				getAuditLogsOptionsModel.SetStartat("testString")
				getAuditLogsOptionsModel.SetOffset(int64(38))
				getAuditLogsOptionsModel.SetLimit(int64(38))
				getAuditLogsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAuditLogsOptionsModel).ToNot(BeNil())
				Expect(getAuditLogsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Ascending).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Startat).To(Equal(core.StringPtr("testString")))
				Expect(getAuditLogsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getAuditLogsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getAuditLogsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogEntryOptions successfully`, func() {
				// Construct an instance of the GetCatalogEntryOptions model
				id := "testString"
				getCatalogEntryOptionsModel := testService.NewGetCatalogEntryOptions(id)
				getCatalogEntryOptionsModel.SetID("testString")
				getCatalogEntryOptionsModel.SetAccount("testString")
				getCatalogEntryOptionsModel.SetInclude("testString")
				getCatalogEntryOptionsModel.SetLanguages("testString")
				getCatalogEntryOptionsModel.SetComplete("testString")
				getCatalogEntryOptionsModel.SetDepth(int64(38))
				getCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(getCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Complete).To(Equal(core.StringPtr("testString")))
				Expect(getCatalogEntryOptionsModel.Depth).To(Equal(core.Int64Ptr(int64(38))))
				Expect(getCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetChildObjectsOptions successfully`, func() {
				// Construct an instance of the GetChildObjectsOptions model
				id := "testString"
				kind := "testString"
				getChildObjectsOptionsModel := testService.NewGetChildObjectsOptions(id, kind)
				getChildObjectsOptionsModel.SetID("testString")
				getChildObjectsOptionsModel.SetKind("testString")
				getChildObjectsOptionsModel.SetAccount("testString")
				getChildObjectsOptionsModel.SetInclude("testString")
				getChildObjectsOptionsModel.SetQ("testString")
				getChildObjectsOptionsModel.SetSortBy("testString")
				getChildObjectsOptionsModel.SetDescending("testString")
				getChildObjectsOptionsModel.SetLanguages("testString")
				getChildObjectsOptionsModel.SetComplete("testString")
				getChildObjectsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getChildObjectsOptionsModel).ToNot(BeNil())
				Expect(getChildObjectsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Kind).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.SortBy).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Descending).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Complete).To(Equal(core.StringPtr("testString")))
				Expect(getChildObjectsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPricingOptions successfully`, func() {
				// Construct an instance of the GetPricingOptions model
				id := "testString"
				getPricingOptionsModel := testService.NewGetPricingOptions(id)
				getPricingOptionsModel.SetID("testString")
				getPricingOptionsModel.SetAccount("testString")
				getPricingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPricingOptionsModel).ToNot(BeNil())
				Expect(getPricingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getPricingOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getPricingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVisibilityOptions successfully`, func() {
				// Construct an instance of the GetVisibilityOptions model
				id := "testString"
				getVisibilityOptionsModel := testService.NewGetVisibilityOptions(id)
				getVisibilityOptionsModel.SetID("testString")
				getVisibilityOptionsModel.SetAccount("testString")
				getVisibilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVisibilityOptionsModel).ToNot(BeNil())
				Expect(getVisibilityOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getVisibilityOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(getVisibilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewImage successfully`, func() {
				image := "testString"
				model, err := testService.NewImage(image)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListArtifactsOptions successfully`, func() {
				// Construct an instance of the ListArtifactsOptions model
				objectID := "testString"
				listArtifactsOptionsModel := testService.NewListArtifactsOptions(objectID)
				listArtifactsOptionsModel.SetObjectID("testString")
				listArtifactsOptionsModel.SetAccount("testString")
				listArtifactsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listArtifactsOptionsModel).ToNot(BeNil())
				Expect(listArtifactsOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(listArtifactsOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(listArtifactsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogEntriesOptions successfully`, func() {
				// Construct an instance of the ListCatalogEntriesOptions model
				listCatalogEntriesOptionsModel := testService.NewListCatalogEntriesOptions()
				listCatalogEntriesOptionsModel.SetAccount("testString")
				listCatalogEntriesOptionsModel.SetInclude("testString")
				listCatalogEntriesOptionsModel.SetQ("testString")
				listCatalogEntriesOptionsModel.SetSortBy("testString")
				listCatalogEntriesOptionsModel.SetDescending("testString")
				listCatalogEntriesOptionsModel.SetLanguages("testString")
				listCatalogEntriesOptionsModel.SetComplete("testString")
				listCatalogEntriesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogEntriesOptionsModel).ToNot(BeNil())
				Expect(listCatalogEntriesOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Include).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Q).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.SortBy).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Descending).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Languages).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Complete).To(Equal(core.StringPtr("testString")))
				Expect(listCatalogEntriesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewOverview successfully`, func() {
				displayName := "testString"
				longDescription := "testString"
				description := "testString"
				model, err := testService.NewOverview(displayName, longDescription, description)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewProvider successfully`, func() {
				email := "testString"
				name := "testString"
				model, err := testService.NewProvider(email, name)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRestoreCatalogEntryOptions successfully`, func() {
				// Construct an instance of the RestoreCatalogEntryOptions model
				id := "testString"
				restoreCatalogEntryOptionsModel := testService.NewRestoreCatalogEntryOptions(id)
				restoreCatalogEntryOptionsModel.SetID("testString")
				restoreCatalogEntryOptionsModel.SetAccount("testString")
				restoreCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(restoreCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(restoreCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(restoreCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogEntryOptions successfully`, func() {
				// Construct an instance of the Bullets model
				bulletsModel := new(globalcatalogv1.Bullets)
				Expect(bulletsModel).ToNot(BeNil())
				bulletsModel.Title = core.StringPtr("testString")
				bulletsModel.Description = core.StringPtr("testString")
				bulletsModel.Icon = core.StringPtr("testString")
				bulletsModel.Quantity = core.Int64Ptr(int64(38))
				Expect(bulletsModel.Title).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Icon).To(Equal(core.StringPtr("testString")))
				Expect(bulletsModel.Quantity).To(Equal(core.Int64Ptr(int64(38))))

				// Construct an instance of the Price model
				priceModel := new(globalcatalogv1.Price)
				Expect(priceModel).ToNot(BeNil())
				priceModel.QuantityTier = core.Int64Ptr(int64(38))
				priceModel.Price = core.Float64Ptr(72.5)
				Expect(priceModel.QuantityTier).To(Equal(core.Int64Ptr(int64(38))))
				Expect(priceModel.Price).To(Equal(core.Float64Ptr(72.5)))

				// Construct an instance of the UIMetaMedia model
				uiMetaMediaModel := new(globalcatalogv1.UIMetaMedia)
				Expect(uiMetaMediaModel).ToNot(BeNil())
				uiMetaMediaModel.Caption = core.StringPtr("testString")
				uiMetaMediaModel.ThumbnailURL = core.StringPtr("testString")
				uiMetaMediaModel.Type = core.StringPtr("testString")
				uiMetaMediaModel.URL = core.StringPtr("testString")
				uiMetaMediaModel.Source = bulletsModel
				Expect(uiMetaMediaModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.ThumbnailURL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaMediaModel.Source).To(Equal(bulletsModel))

				// Construct an instance of the Amount model
				amountModel := new(globalcatalogv1.Amount)
				Expect(amountModel).ToNot(BeNil())
				amountModel.Country = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}
				Expect(amountModel.Country).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Currency).To(Equal(core.StringPtr("testString")))
				Expect(amountModel.Prices).To(Equal([]globalcatalogv1.Price{*priceModel}))

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				Expect(stringsModel).ToNot(BeNil())
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")
				Expect(stringsModel.Bullets).To(Equal([]globalcatalogv1.Bullets{*bulletsModel}))
				Expect(stringsModel.Media).To(Equal([]globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}))
				Expect(stringsModel.NotCreatableMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.NotCreatableRobotMsg).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.DeprecationWarning).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.PopupWarningMessage).To(Equal(core.StringPtr("testString")))
				Expect(stringsModel.Instruction).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Broker model
				brokerModel := new(globalcatalogv1.Broker)
				Expect(brokerModel).ToNot(BeNil())
				brokerModel.Name = core.StringPtr("testString")
				brokerModel.Guid = core.StringPtr("testString")
				Expect(brokerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(brokerModel.Guid).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the DRMetaData model
				drMetaDataModel := new(globalcatalogv1.DRMetaData)
				Expect(drMetaDataModel).ToNot(BeNil())
				drMetaDataModel.Dr = core.BoolPtr(true)
				drMetaDataModel.Description = core.StringPtr("testString")
				Expect(drMetaDataModel.Dr).To(Equal(core.BoolPtr(true)))
				Expect(drMetaDataModel.Description).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				Expect(i18NModel).ToNot(BeNil())
				i18NModel.SetProperty("foo", stringsModel)
				Expect(i18NModel.GetProperty("foo")).To(Equal(stringsModel))
				Expect(i18NModel.GetProperties()).ToNot(BeEmpty())

				// Construct an instance of the SourceMetaData model
				sourceMetaDataModel := new(globalcatalogv1.SourceMetaData)
				Expect(sourceMetaDataModel).ToNot(BeNil())
				sourceMetaDataModel.Path = core.StringPtr("testString")
				sourceMetaDataModel.Type = core.StringPtr("testString")
				sourceMetaDataModel.URL = core.StringPtr("testString")
				Expect(sourceMetaDataModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(sourceMetaDataModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				Expect(startingPriceModel).ToNot(BeNil())
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
				startingPriceModel.Unit = core.StringPtr("testString")
				startingPriceModel.Amount = []globalcatalogv1.Amount{*amountModel}
				Expect(startingPriceModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.DeploymentID).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Unit).To(Equal(core.StringPtr("testString")))
				Expect(startingPriceModel.Amount).To(Equal([]globalcatalogv1.Amount{*amountModel}))

				// Construct an instance of the URLS model
				urlsModel := new(globalcatalogv1.URLS)
				Expect(urlsModel).ToNot(BeNil())
				urlsModel.DocURL = core.StringPtr("testString")
				urlsModel.InstructionsURL = core.StringPtr("testString")
				urlsModel.ApiURL = core.StringPtr("testString")
				urlsModel.CreateURL = core.StringPtr("testString")
				urlsModel.SdkDownloadURL = core.StringPtr("testString")
				urlsModel.TermsURL = core.StringPtr("testString")
				urlsModel.CustomCreatePageURL = core.StringPtr("testString")
				urlsModel.CatalogDetailsURL = core.StringPtr("testString")
				urlsModel.DeprecationDocURL = core.StringPtr("testString")
				urlsModel.DashboardURL = core.StringPtr("testString")
				urlsModel.RegistrationURL = core.StringPtr("testString")
				urlsModel.Apidocsurl = core.StringPtr("testString")
				Expect(urlsModel.DocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.InstructionsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.ApiURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CreateURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.SdkDownloadURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.TermsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CustomCreatePageURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.CatalogDetailsURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DeprecationDocURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.RegistrationURL).To(Equal(core.StringPtr("testString")))
				Expect(urlsModel.Apidocsurl).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AliasMetaData model
				aliasMetaDataModel := new(globalcatalogv1.AliasMetaData)
				Expect(aliasMetaDataModel).ToNot(BeNil())
				aliasMetaDataModel.Type = core.StringPtr("testString")
				aliasMetaDataModel.PlanID = core.StringPtr("testString")
				Expect(aliasMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(aliasMetaDataModel.PlanID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CFMetaData model
				cfMetaDataModel := new(globalcatalogv1.CFMetaData)
				Expect(cfMetaDataModel).ToNot(BeNil())
				cfMetaDataModel.Type = core.StringPtr("testString")
				cfMetaDataModel.IamCompatible = core.BoolPtr(true)
				cfMetaDataModel.UniqueApiKey = core.BoolPtr(true)
				cfMetaDataModel.Provisionable = core.BoolPtr(true)
				cfMetaDataModel.Bindable = core.BoolPtr(true)
				cfMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				cfMetaDataModel.Requires = []string{"testString"}
				cfMetaDataModel.PlanUpdateable = core.BoolPtr(true)
				cfMetaDataModel.State = core.StringPtr("testString")
				cfMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				cfMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				cfMetaDataModel.ServiceKeySupported = core.BoolPtr(true)
				cfMetaDataModel.CfGuid = make(map[string]string)
				Expect(cfMetaDataModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.IamCompatible).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.UniqueApiKey).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Provisionable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.Requires).To(Equal([]string{"testString"}))
				Expect(cfMetaDataModel.PlanUpdateable).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.State).To(Equal(core.StringPtr("testString")))
				Expect(cfMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(cfMetaDataModel.ServiceKeySupported).To(Equal(core.BoolPtr(true)))
				Expect(cfMetaDataModel.CfGuid).To(Equal(make(map[string]string)))

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				Expect(callbacksModel).ToNot(BeNil())
				callbacksModel.ControllerURL = core.StringPtr("testString")
				callbacksModel.BrokerURL = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ApiEndpoint = make(map[string]string)
				Expect(callbacksModel.ControllerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.BrokerProxyURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDataURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.DashboardDetailTabExtURL).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApi).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ServiceMonitorApp).To(Equal(core.StringPtr("testString")))
				Expect(callbacksModel.ApiEndpoint).To(Equal(make(map[string]string)))

				// Construct an instance of the DeploymentBase model
				deploymentBaseModel := new(globalcatalogv1.DeploymentBase)
				Expect(deploymentBaseModel).ToNot(BeNil())
				deploymentBaseModel.Location = core.StringPtr("testString")
				deploymentBaseModel.LocationURL = core.StringPtr("testString")
				deploymentBaseModel.OriginalLocation = core.StringPtr("testString")
				deploymentBaseModel.TargetCrn = core.StringPtr("testString")
				deploymentBaseModel.ServiceCrn = core.StringPtr("testString")
				deploymentBaseModel.MccpID = core.StringPtr("testString")
				deploymentBaseModel.Broker = brokerModel
				deploymentBaseModel.SupportsRcMigration = core.BoolPtr(true)
				deploymentBaseModel.TargetNetwork = core.StringPtr("testString")
				Expect(deploymentBaseModel.Location).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.LocationURL).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.OriginalLocation).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.TargetCrn).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.ServiceCrn).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.MccpID).To(Equal(core.StringPtr("testString")))
				Expect(deploymentBaseModel.Broker).To(Equal(brokerModel))
				Expect(deploymentBaseModel.SupportsRcMigration).To(Equal(core.BoolPtr(true)))
				Expect(deploymentBaseModel.TargetNetwork).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				Expect(overviewModel).ToNot(BeNil())
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")
				overviewModel.FeaturedDescription = core.StringPtr("testString")
				Expect(overviewModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(overviewModel.FeaturedDescription).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PlanMetaData model
				planMetaDataModel := new(globalcatalogv1.PlanMetaData)
				Expect(planMetaDataModel).ToNot(BeNil())
				planMetaDataModel.Bindable = core.BoolPtr(true)
				planMetaDataModel.Reservable = core.BoolPtr(true)
				planMetaDataModel.AllowInternalUsers = core.BoolPtr(true)
				planMetaDataModel.AsyncProvisioningSupported = core.BoolPtr(true)
				planMetaDataModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				planMetaDataModel.TestCheckInterval = core.Int64Ptr(int64(38))
				planMetaDataModel.SingleScopeInstance = core.StringPtr("testString")
				planMetaDataModel.ServiceCheckEnabled = core.BoolPtr(true)
				planMetaDataModel.CfGuid = make(map[string]string)
				Expect(planMetaDataModel.Bindable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.Reservable).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AllowInternalUsers).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncProvisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.AsyncUnprovisioningSupported).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.TestCheckInterval).To(Equal(core.Int64Ptr(int64(38))))
				Expect(planMetaDataModel.SingleScopeInstance).To(Equal(core.StringPtr("testString")))
				Expect(planMetaDataModel.ServiceCheckEnabled).To(Equal(core.BoolPtr(true)))
				Expect(planMetaDataModel.CfGuid).To(Equal(make(map[string]string)))

				// Construct an instance of the PricingSet model
				pricingSetModel := new(globalcatalogv1.PricingSet)
				Expect(pricingSetModel).ToNot(BeNil())
				pricingSetModel.Type = core.StringPtr("testString")
				pricingSetModel.Origin = core.StringPtr("testString")
				pricingSetModel.StartingPrice = startingPriceModel
				Expect(pricingSetModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.Origin).To(Equal(core.StringPtr("testString")))
				Expect(pricingSetModel.StartingPrice).To(Equal(startingPriceModel))

				// Construct an instance of the SLAMetaData model
				slaMetaDataModel := new(globalcatalogv1.SLAMetaData)
				Expect(slaMetaDataModel).ToNot(BeNil())
				slaMetaDataModel.Terms = core.StringPtr("testString")
				slaMetaDataModel.Tenancy = core.StringPtr("testString")
				slaMetaDataModel.Provisioning = core.StringPtr("testString")
				slaMetaDataModel.Responsiveness = core.StringPtr("testString")
				slaMetaDataModel.Dr = drMetaDataModel
				Expect(slaMetaDataModel.Terms).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Tenancy).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Provisioning).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Responsiveness).To(Equal(core.StringPtr("testString")))
				Expect(slaMetaDataModel.Dr).To(Equal(drMetaDataModel))

				// Construct an instance of the TemplateMetaData model
				templateMetaDataModel := new(globalcatalogv1.TemplateMetaData)
				Expect(templateMetaDataModel).ToNot(BeNil())
				templateMetaDataModel.Services = []string{"testString"}
				templateMetaDataModel.DefaultMemory = core.Int64Ptr(int64(38))
				templateMetaDataModel.StartCmd = core.StringPtr("testString")
				templateMetaDataModel.Source = sourceMetaDataModel
				templateMetaDataModel.RuntimeCatalogID = core.StringPtr("testString")
				templateMetaDataModel.CfRuntimeID = core.StringPtr("testString")
				templateMetaDataModel.TemplateID = core.StringPtr("testString")
				templateMetaDataModel.ExecutableFile = core.StringPtr("testString")
				templateMetaDataModel.Buildpack = core.StringPtr("testString")
				templateMetaDataModel.EnvironmentVariables = make(map[string]string)
				Expect(templateMetaDataModel.Services).To(Equal([]string{"testString"}))
				Expect(templateMetaDataModel.DefaultMemory).To(Equal(core.Int64Ptr(int64(38))))
				Expect(templateMetaDataModel.StartCmd).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Source).To(Equal(sourceMetaDataModel))
				Expect(templateMetaDataModel.RuntimeCatalogID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.CfRuntimeID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.ExecutableFile).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.Buildpack).To(Equal(core.StringPtr("testString")))
				Expect(templateMetaDataModel.EnvironmentVariables).To(Equal(make(map[string]string)))

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				Expect(uiMetaDataModel).ToNot(BeNil())
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()
				uiMetaDataModel.Hidden = core.BoolPtr(true)
				uiMetaDataModel.HideLiteMetering = core.BoolPtr(true)
				uiMetaDataModel.NoUpgradeNextStep = core.BoolPtr(true)
				Expect(uiMetaDataModel.Strings).To(Equal(i18NModel))
				Expect(uiMetaDataModel.Urls).To(Equal(urlsModel))
				Expect(uiMetaDataModel.EmbeddableDashboard).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.EmbeddableDashboardFullWidth).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NavigationOrder).To(Equal([]string{"testString"}))
				Expect(uiMetaDataModel.NotCreatable).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.PrimaryOfferingID).To(Equal(core.StringPtr("testString")))
				Expect(uiMetaDataModel.AccessibleDuringProvision).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.SideBySideIndex).To(Equal(core.Int64Ptr(int64(38))))
				Expect(uiMetaDataModel.EndOfServiceTime).To(Equal(CreateMockDateTime()))
				Expect(uiMetaDataModel.Hidden).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.HideLiteMetering).To(Equal(core.BoolPtr(true)))
				Expect(uiMetaDataModel.NoUpgradeNextStep).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				Expect(imageModel).ToNot(BeNil())
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")
				Expect(imageModel.Image).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.SmallImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.MediumImage).To(Equal(core.StringPtr("testString")))
				Expect(imageModel.FeatureImage).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ObjectMetadataSet model
				objectMetadataSetModel := new(globalcatalogv1.ObjectMetadataSet)
				Expect(objectMetadataSetModel).ToNot(BeNil())
				objectMetadataSetModel.RcCompatible = core.BoolPtr(true)
				objectMetadataSetModel.Service = cfMetaDataModel
				objectMetadataSetModel.Plan = planMetaDataModel
				objectMetadataSetModel.Alias = aliasMetaDataModel
				objectMetadataSetModel.Template = templateMetaDataModel
				objectMetadataSetModel.Ui = uiMetaDataModel
				objectMetadataSetModel.Compliance = []string{"testString"}
				objectMetadataSetModel.Sla = slaMetaDataModel
				objectMetadataSetModel.Callbacks = callbacksModel
				objectMetadataSetModel.OriginalName = core.StringPtr("testString")
				objectMetadataSetModel.Version = core.StringPtr("testString")
				objectMetadataSetModel.Other = make(map[string]interface{})
				objectMetadataSetModel.Pricing = pricingSetModel
				objectMetadataSetModel.Deployment = deploymentBaseModel
				Expect(objectMetadataSetModel.RcCompatible).To(Equal(core.BoolPtr(true)))
				Expect(objectMetadataSetModel.Service).To(Equal(cfMetaDataModel))
				Expect(objectMetadataSetModel.Plan).To(Equal(planMetaDataModel))
				Expect(objectMetadataSetModel.Alias).To(Equal(aliasMetaDataModel))
				Expect(objectMetadataSetModel.Template).To(Equal(templateMetaDataModel))
				Expect(objectMetadataSetModel.Ui).To(Equal(uiMetaDataModel))
				Expect(objectMetadataSetModel.Compliance).To(Equal([]string{"testString"}))
				Expect(objectMetadataSetModel.Sla).To(Equal(slaMetaDataModel))
				Expect(objectMetadataSetModel.Callbacks).To(Equal(callbacksModel))
				Expect(objectMetadataSetModel.OriginalName).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(objectMetadataSetModel.Other).To(Equal(make(map[string]interface{})))
				Expect(objectMetadataSetModel.Pricing).To(Equal(pricingSetModel))
				Expect(objectMetadataSetModel.Deployment).To(Equal(deploymentBaseModel))

				// Construct an instance of the OverviewUI model
				overviewUiModel := new(globalcatalogv1.OverviewUI)
				Expect(overviewUiModel).ToNot(BeNil())
				overviewUiModel.SetProperty("foo", overviewModel)
				Expect(overviewUiModel.GetProperty("foo")).To(Equal(overviewModel))
				Expect(overviewUiModel.GetProperties()).ToNot(BeEmpty())

				// Construct an instance of the Provider model
				providerModel := new(globalcatalogv1.Provider)
				Expect(providerModel).ToNot(BeNil())
				providerModel.Email = core.StringPtr("testString")
				providerModel.Name = core.StringPtr("testString")
				providerModel.Contact = core.StringPtr("testString")
				providerModel.SupportEmail = core.StringPtr("testString")
				providerModel.Phone = core.StringPtr("testString")
				Expect(providerModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.SupportEmail).To(Equal(core.StringPtr("testString")))
				Expect(providerModel.Phone).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateCatalogEntryOptions model
				id := "testString"
				updateCatalogEntryOptionsName := "testString"
				updateCatalogEntryOptionsKind := "service"
				var updateCatalogEntryOptionsOverviewUi *globalcatalogv1.OverviewUI = nil
				var updateCatalogEntryOptionsImages *globalcatalogv1.Image = nil
				updateCatalogEntryOptionsDisabled := true
				updateCatalogEntryOptionsTags := []string{"testString"}
				var updateCatalogEntryOptionsProvider *globalcatalogv1.Provider = nil
				updateCatalogEntryOptionsModel := testService.NewUpdateCatalogEntryOptions(id, updateCatalogEntryOptionsName, updateCatalogEntryOptionsKind, updateCatalogEntryOptionsOverviewUi, updateCatalogEntryOptionsImages, updateCatalogEntryOptionsDisabled, updateCatalogEntryOptionsTags, updateCatalogEntryOptionsProvider)
				updateCatalogEntryOptionsModel.SetID("testString")
				updateCatalogEntryOptionsModel.SetName("testString")
				updateCatalogEntryOptionsModel.SetKind("service")
				updateCatalogEntryOptionsModel.SetOverviewUi(overviewUiModel)
				updateCatalogEntryOptionsModel.SetImages(imageModel)
				updateCatalogEntryOptionsModel.SetDisabled(true)
				updateCatalogEntryOptionsModel.SetTags([]string{"testString"})
				updateCatalogEntryOptionsModel.SetProvider(providerModel)
				updateCatalogEntryOptionsModel.SetParentID("testString")
				updateCatalogEntryOptionsModel.SetGroup(true)
				updateCatalogEntryOptionsModel.SetActive(true)
				updateCatalogEntryOptionsModel.SetMetadata(objectMetadataSetModel)
				updateCatalogEntryOptionsModel.SetAccount("testString")
				updateCatalogEntryOptionsModel.SetMove("testString")
				updateCatalogEntryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogEntryOptionsModel).ToNot(BeNil())
				Expect(updateCatalogEntryOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Kind).To(Equal(core.StringPtr("service")))
				Expect(updateCatalogEntryOptionsModel.OverviewUi).To(Equal(overviewUiModel))
				Expect(updateCatalogEntryOptionsModel.Images).To(Equal(imageModel))
				Expect(updateCatalogEntryOptionsModel.Disabled).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(updateCatalogEntryOptionsModel.Provider).To(Equal(providerModel))
				Expect(updateCatalogEntryOptionsModel.ParentID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Group).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.Active).To(Equal(core.BoolPtr(true)))
				Expect(updateCatalogEntryOptionsModel.Metadata).To(Equal(objectMetadataSetModel))
				Expect(updateCatalogEntryOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Move).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogEntryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVisibilityOptions successfully`, func() {
				// Construct an instance of the VisibilityDetailAccounts model
				visibilityDetailAccountsModel := new(globalcatalogv1.VisibilityDetailAccounts)
				Expect(visibilityDetailAccountsModel).ToNot(BeNil())
				visibilityDetailAccountsModel.Accountid = core.StringPtr("testString")
				Expect(visibilityDetailAccountsModel.Accountid).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				Expect(visibilityDetailModel).ToNot(BeNil())
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel
				Expect(visibilityDetailModel.Accounts).To(Equal(visibilityDetailAccountsModel))

				// Construct an instance of the UpdateVisibilityOptions model
				id := "testString"
				updateVisibilityOptionsModel := testService.NewUpdateVisibilityOptions(id)
				updateVisibilityOptionsModel.SetID("testString")
				updateVisibilityOptionsModel.SetExtendable(true)
				updateVisibilityOptionsModel.SetInclude(visibilityDetailModel)
				updateVisibilityOptionsModel.SetExclude(visibilityDetailModel)
				updateVisibilityOptionsModel.SetAccount("testString")
				updateVisibilityOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVisibilityOptionsModel).ToNot(BeNil())
				Expect(updateVisibilityOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVisibilityOptionsModel.Extendable).To(Equal(core.BoolPtr(true)))
				Expect(updateVisibilityOptionsModel.Include).To(Equal(visibilityDetailModel))
				Expect(updateVisibilityOptionsModel.Exclude).To(Equal(visibilityDetailModel))
				Expect(updateVisibilityOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(updateVisibilityOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUploadArtifactOptions successfully`, func() {
				// Construct an instance of the UploadArtifactOptions model
				objectID := "testString"
				artifactID := "testString"
				uploadArtifactOptionsModel := testService.NewUploadArtifactOptions(objectID, artifactID)
				uploadArtifactOptionsModel.SetObjectID("testString")
				uploadArtifactOptionsModel.SetArtifactID("testString")
				uploadArtifactOptionsModel.SetArtifact(CreateMockReader("This is a mock file."))
				uploadArtifactOptionsModel.SetContentType("testString")
				uploadArtifactOptionsModel.SetAccount("testString")
				uploadArtifactOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(uploadArtifactOptionsModel).ToNot(BeNil())
				Expect(uploadArtifactOptionsModel.ObjectID).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.ArtifactID).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Artifact).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(uploadArtifactOptionsModel.ContentType).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Account).To(Equal(core.StringPtr("testString")))
				Expect(uploadArtifactOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVisibilityDetail successfully`, func() {
				var accounts *globalcatalogv1.VisibilityDetailAccounts = nil
				_, err := testService.NewVisibilityDetail(accounts)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
