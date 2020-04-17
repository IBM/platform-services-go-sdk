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
	"github.com/IBM/go-sdk-core/v3/core"
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
                URL: "{{BAD_URL_STRING",
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
                URL: "{{BAD_URL_STRING",
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"page": "Page", "results_per_page": "ResultsPerPage", "total_results": "TotalResults", "resources": [{"anyKey": "anyValue"}]}`)
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

					res.Header().Set("Content-type", "text/plain")
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Kind = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetaDataModel
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
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
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": []}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": []}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "metrics": []}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack"}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness"}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "compliance": ["Compliance"], "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "kind": "Kind", "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": []}]}]}]}]}]}]}]}]}]}`)
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Kind = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetaDataModel
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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
				createCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Name = core.StringPtr("testString")
				createCatalogEntryOptionsModel.OverviewUi = overviewUiModel
				createCatalogEntryOptionsModel.Kind = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Images = imageModel
				createCatalogEntryOptionsModel.Disabled = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Tags = []string{"testString"}
				createCatalogEntryOptionsModel.Provider = providerModel
				createCatalogEntryOptionsModel.ParentID = core.StringPtr("testString")
				createCatalogEntryOptionsModel.Group = core.BoolPtr(true)
				createCatalogEntryOptionsModel.Metadata = objectMetaDataModel
				createCatalogEntryOptionsModel.Active = core.BoolPtr(true)
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": []}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": []}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "metrics": []}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack"}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness"}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "compliance": ["Compliance"], "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "kind": "Kind", "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": []}]}]}]}]}]}]}]}]}]}`)
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

					res.Header().Set("Content-type", "text/plain")
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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

				// Construct an instance of the CatalogEntry model
				catalogEntryModel := new(globalcatalogv1.CatalogEntry)
				catalogEntryModel.Name = core.StringPtr("testString")
				catalogEntryModel.OverviewUi = overviewUiModel
				catalogEntryModel.Kind = core.StringPtr("testString")
				catalogEntryModel.Images = imageModel
				catalogEntryModel.ParentID = core.StringPtr("testString")
				catalogEntryModel.Disabled = core.BoolPtr(true)
				catalogEntryModel.Tags = []string{"testString"}
				catalogEntryModel.Group = core.BoolPtr(true)
				catalogEntryModel.Provider = providerModel
				catalogEntryModel.Metadata = objectMetaDataModel
				catalogEntryModel.Active = core.BoolPtr(true)
				catalogEntryModel.ParentURL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewName = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewOverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.NewKind = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewImages = imageModel
				updateCatalogEntryOptionsModel.NewDisabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewProvider = providerModel
				updateCatalogEntryOptionsModel.NewParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGroup = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewMetadata = objectMetaDataModel
				updateCatalogEntryOptionsModel.NewActive = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewCatalogCrn = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewChildrenURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewParentURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGeoTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewPricingTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewCreated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewUpdated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewChildren = []globalcatalogv1.CatalogEntry{*catalogEntryModel}
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
					fmt.Fprintf(res, `{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": []}]}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid", "password": {"text": "Text", "key": "Key", "iv": "Iv"}}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"strings": {}, "urls": {"doc_url": "DocURL", "instructions_url": "InstructionsURL", "api_url": "ApiURL", "create_url": "CreateURL", "sdk_download_url": "SdkDownloadURL", "terms_url": "TermsURL", "custom_create_page_url": "CustomCreatePageURL", "catalog_details_url": "CatalogDetailsURL", "deprecation_doc_url": "DeprecationDocURL"}, "embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": []}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": []}]}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "source": {"path": "Path", "type": "Type", "url": "URL"}, "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack", "environment_variables": {"_key_": "Key"}}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "broker": {"name": "Name", "guid": "Guid"}, "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness", "dr": {"dr": true, "description": "Description"}}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "ui": {"embeddable_dashboard": "EmbeddableDashboard", "embeddable_dashboard_full_width": true, "navigation_order": ["NavigationOrder"], "not_creatable": true, "reservable": true, "primary_offering_id": "PrimaryOfferingID", "accessible_during_provision": false, "side_by_side_index": 15, "end_of_service_time": "2019-01-01T12:00:00"}, "pricing": {"type": "Type", "origin": "Origin", "metrics": []}, "compliance": ["Compliance"], "service": {"type": "Type", "iam_compatible": false, "unique_api_key": true, "provisionable": false, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "cf_guid": "CfGuid", "bindable": true, "requires": ["Requires"], "plan_updateable": true, "state": "State", "service_check_enabled": false, "test_check_interval": 17, "service_key_supported": false}, "plan": {"bindable": true, "reservable": true, "allow_internal_users": true, "async_provisioning_supported": true, "async_unprovisioning_supported": true, "test_check_interval": 17, "single_scope_instance": "SingleScopeInstance", "service_check_enabled": false, "cf_guid": "CfGuid"}, "template": {"services": ["Services"], "default_memory": 13, "start_cmd": "StartCmd", "runtime_catalog_id": "RuntimeCatalogID", "cf_runtime_id": "CfRuntimeID", "template_id": "TemplateID", "executable_file": "ExecutableFile", "buildpack": "Buildpack"}, "deployment": {"location": "Location", "location_url": "LocationURL", "target_crn": "TargetCrn", "supports_rc_migration": false}, "alias": {"type": "Type", "plan_id": "PlanID"}, "sla": {"terms": "Terms", "tenancy": "Tenancy", "provisioning": "Provisioning", "responsiveness": "Responsiveness"}, "callbacks": {"broker_utl": "BrokerUtl", "broker_proxy_url": "BrokerProxyURL", "dashboard_url": "DashboardURL", "dashboard_data_url": "DashboardDataURL", "dashboard_detail_tab_url": "DashboardDetailTabURL", "dashboard_detail_tab_ext_url": "DashboardDetailTabExtURL", "service_monitor_api": "ServiceMonitorApi", "service_monitor_app": "ServiceMonitorApp", "service_staging_url": "ServiceStagingURL", "service_production_url": "ServiceProductionURL"}, "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "overview_ui": {}, "kind": "Kind", "images": {"image": "Image", "small_image": "SmallImage", "medium_image": "MediumImage", "feature_image": "FeatureImage"}, "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "provider": {"email": "Email", "name": "Name", "contact": "Contact", "support_email": "SupportEmail", "phone": "Phone"}, "metadata": {"rc_compatible": true, "compliance": ["Compliance"], "version": "Version", "original_name": "OriginalName", "other": {"anyKey": "anyValue"}}, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": [{"id": "ID", "name": "Name", "kind": "Kind", "parent_id": "ParentID", "disabled": true, "tags": ["Tags"], "group": false, "active": true, "catalog_crn": "CatalogCrn", "url": "URL", "children_url": "ChildrenURL", "parent_url": "ParentURL", "geo_tags": ["GeoTags"], "pricing_tags": ["PricingTags"], "created": "2019-01-01T12:00:00", "updated": "2019-01-01T12:00:00", "children": []}]}]}]}]}]}]}]}]}]}`)
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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

				// Construct an instance of the CatalogEntry model
				catalogEntryModel := new(globalcatalogv1.CatalogEntry)
				catalogEntryModel.Name = core.StringPtr("testString")
				catalogEntryModel.OverviewUi = overviewUiModel
				catalogEntryModel.Kind = core.StringPtr("testString")
				catalogEntryModel.Images = imageModel
				catalogEntryModel.ParentID = core.StringPtr("testString")
				catalogEntryModel.Disabled = core.BoolPtr(true)
				catalogEntryModel.Tags = []string{"testString"}
				catalogEntryModel.Group = core.BoolPtr(true)
				catalogEntryModel.Provider = providerModel
				catalogEntryModel.Metadata = objectMetaDataModel
				catalogEntryModel.Active = core.BoolPtr(true)
				catalogEntryModel.ParentURL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewName = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewOverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.NewKind = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewImages = imageModel
				updateCatalogEntryOptionsModel.NewDisabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewProvider = providerModel
				updateCatalogEntryOptionsModel.NewParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGroup = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewMetadata = objectMetaDataModel
				updateCatalogEntryOptionsModel.NewActive = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewCatalogCrn = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewChildrenURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewParentURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGeoTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewPricingTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewCreated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewUpdated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewChildren = []globalcatalogv1.CatalogEntry{*catalogEntryModel}
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
				bulletsModel.Quantity = core.StringPtr("testString")

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
				amountModel.Counrty = core.StringPtr("testString")
				amountModel.Currency = core.StringPtr("testString")
				amountModel.Prices = []globalcatalogv1.Price{*priceModel}

				// Construct an instance of the ObjectMetaDataDeploymentBrokerPassword model
				objectMetaDataDeploymentBrokerPasswordModel := new(globalcatalogv1.ObjectMetaDataDeploymentBrokerPassword)
				objectMetaDataDeploymentBrokerPasswordModel.Text = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Key = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerPasswordModel.Iv = core.StringPtr("testString")

				// Construct an instance of the Strings model
				stringsModel := new(globalcatalogv1.Strings)
				stringsModel.Bullets = []globalcatalogv1.Bullets{*bulletsModel}
				stringsModel.Media = []globalcatalogv1.UIMetaMedia{*uiMetaMediaModel}
				stringsModel.NotCreatableMsg = core.StringPtr("testString")
				stringsModel.NotCreatableRobotMsg = core.StringPtr("testString")
				stringsModel.DeprecationWarning = core.StringPtr("testString")
				stringsModel.PopupWarningMessage = core.StringPtr("testString")
				stringsModel.Instruction = core.StringPtr("testString")

				// Construct an instance of the I18N model
				i18NModel := new(globalcatalogv1.I18N)
				i18NModel.SetProperty("foo", stringsModel)

				// Construct an instance of the Metrics model
				metricsModel := new(globalcatalogv1.Metrics)
				metricsModel.MetricID = core.StringPtr("testString")
				metricsModel.TierModel = core.StringPtr("testString")
				metricsModel.ChargeUnitName = core.StringPtr("testString")
				metricsModel.ChargeUnitQuantity = core.StringPtr("testString")
				metricsModel.ResourceDisplayName = core.StringPtr("testString")
				metricsModel.ChargeUnitDisplayName = core.StringPtr("testString")
				metricsModel.UsageCapQty = core.Int64Ptr(int64(38))
				metricsModel.Amounts = []globalcatalogv1.Amount{*amountModel}

				// Construct an instance of the ObjectMetaDataDeploymentBroker model
				objectMetaDataDeploymentBrokerModel := new(globalcatalogv1.ObjectMetaDataDeploymentBroker)
				objectMetaDataDeploymentBrokerModel.Name = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Guid = core.StringPtr("testString")
				objectMetaDataDeploymentBrokerModel.Password = objectMetaDataDeploymentBrokerPasswordModel

				// Construct an instance of the ObjectMetaDataSlaDr model
				objectMetaDataSlaDrModel := new(globalcatalogv1.ObjectMetaDataSlaDr)
				objectMetaDataSlaDrModel.Dr = core.BoolPtr(true)
				objectMetaDataSlaDrModel.Description = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateEnvironmentVariables model
				objectMetaDataTemplateEnvironmentVariablesModel := new(globalcatalogv1.ObjectMetaDataTemplateEnvironmentVariables)
				objectMetaDataTemplateEnvironmentVariablesModel.Key = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataTemplateSource model
				objectMetaDataTemplateSourceModel := new(globalcatalogv1.ObjectMetaDataTemplateSource)
				objectMetaDataTemplateSourceModel.Path = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.Type = core.StringPtr("testString")
				objectMetaDataTemplateSourceModel.URL = core.StringPtr("testString")

				// Construct an instance of the StartingPrice model
				startingPriceModel := new(globalcatalogv1.StartingPrice)
				startingPriceModel.PlanID = core.StringPtr("testString")
				startingPriceModel.DeploymentID = core.StringPtr("testString")
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

				// Construct an instance of the Callbacks model
				callbacksModel := new(globalcatalogv1.Callbacks)
				callbacksModel.BrokerUtl = core.StringPtr("testString")
				callbacksModel.BrokerProxyURL = core.StringPtr("testString")
				callbacksModel.DashboardURL = core.StringPtr("testString")
				callbacksModel.DashboardDataURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabURL = core.StringPtr("testString")
				callbacksModel.DashboardDetailTabExtURL = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApi = core.StringPtr("testString")
				callbacksModel.ServiceMonitorApp = core.StringPtr("testString")
				callbacksModel.ServiceStagingURL = core.StringPtr("testString")
				callbacksModel.ServiceProductionURL = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataAlias model
				objectMetaDataAliasModel := new(globalcatalogv1.ObjectMetaDataAlias)
				objectMetaDataAliasModel.Type = core.StringPtr("testString")
				objectMetaDataAliasModel.PlanID = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataDeployment model
				objectMetaDataDeploymentModel := new(globalcatalogv1.ObjectMetaDataDeployment)
				objectMetaDataDeploymentModel.Location = core.StringPtr("testString")
				objectMetaDataDeploymentModel.TargetCrn = core.StringPtr("testString")
				objectMetaDataDeploymentModel.Broker = objectMetaDataDeploymentBrokerModel
				objectMetaDataDeploymentModel.SupportsRcMigration = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataPlan model
				objectMetaDataPlanModel := new(globalcatalogv1.ObjectMetaDataPlan)
				objectMetaDataPlanModel.Bindable = core.BoolPtr(true)
				objectMetaDataPlanModel.Reservable = core.BoolPtr(true)
				objectMetaDataPlanModel.AllowInternalUsers = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataPlanModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataPlanModel.SingleScopeInstance = core.StringPtr("testString")
				objectMetaDataPlanModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataPlanModel.CfGuid = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaDataService model
				objectMetaDataServiceModel := new(globalcatalogv1.ObjectMetaDataService)
				objectMetaDataServiceModel.Type = core.StringPtr("testString")
				objectMetaDataServiceModel.IamCompatible = core.BoolPtr(true)
				objectMetaDataServiceModel.UniqueApiKey = core.BoolPtr(true)
				objectMetaDataServiceModel.Provisionable = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncProvisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.AsyncUnprovisioningSupported = core.BoolPtr(true)
				objectMetaDataServiceModel.CfGuid = core.StringPtr("testString")
				objectMetaDataServiceModel.Bindable = core.BoolPtr(true)
				objectMetaDataServiceModel.Requires = []string{"testString"}
				objectMetaDataServiceModel.PlanUpdateable = core.BoolPtr(true)
				objectMetaDataServiceModel.State = core.StringPtr("testString")
				objectMetaDataServiceModel.ServiceCheckEnabled = core.BoolPtr(true)
				objectMetaDataServiceModel.TestCheckInterval = core.Int64Ptr(int64(38))
				objectMetaDataServiceModel.ServiceKeySupported = core.BoolPtr(true)

				// Construct an instance of the ObjectMetaDataSla model
				objectMetaDataSlaModel := new(globalcatalogv1.ObjectMetaDataSla)
				objectMetaDataSlaModel.Terms = core.StringPtr("testString")
				objectMetaDataSlaModel.Tenancy = core.StringPtr("testString")
				objectMetaDataSlaModel.Provisioning = core.StringPtr("testString")
				objectMetaDataSlaModel.Responsiveness = core.StringPtr("testString")
				objectMetaDataSlaModel.Dr = objectMetaDataSlaDrModel

				// Construct an instance of the ObjectMetaDataTemplate model
				objectMetaDataTemplateModel := new(globalcatalogv1.ObjectMetaDataTemplate)
				objectMetaDataTemplateModel.Services = []string{"testString"}
				objectMetaDataTemplateModel.DefaultMemory = core.Int64Ptr(int64(38))
				objectMetaDataTemplateModel.StartCmd = core.StringPtr("testString")
				objectMetaDataTemplateModel.Source = objectMetaDataTemplateSourceModel
				objectMetaDataTemplateModel.RuntimeCatalogID = core.StringPtr("testString")
				objectMetaDataTemplateModel.CfRuntimeID = core.StringPtr("testString")
				objectMetaDataTemplateModel.TemplateID = core.StringPtr("testString")
				objectMetaDataTemplateModel.ExecutableFile = core.StringPtr("testString")
				objectMetaDataTemplateModel.Buildpack = core.StringPtr("testString")
				objectMetaDataTemplateModel.EnvironmentVariables = objectMetaDataTemplateEnvironmentVariablesModel

				// Construct an instance of the Overview model
				overviewModel := new(globalcatalogv1.Overview)
				overviewModel.DisplayName = core.StringPtr("testString")
				overviewModel.LongDescription = core.StringPtr("testString")
				overviewModel.Description = core.StringPtr("testString")

				// Construct an instance of the Pricing model
				pricingModel := new(globalcatalogv1.Pricing)
				pricingModel.Type = core.StringPtr("testString")
				pricingModel.Origin = core.StringPtr("testString")
				pricingModel.StartingPrice = startingPriceModel
				pricingModel.Metrics = []globalcatalogv1.Metrics{*metricsModel}

				// Construct an instance of the UIMetaData model
				uiMetaDataModel := new(globalcatalogv1.UIMetaData)
				uiMetaDataModel.Strings = i18NModel
				uiMetaDataModel.Urls = urlsModel
				uiMetaDataModel.EmbeddableDashboard = core.StringPtr("testString")
				uiMetaDataModel.EmbeddableDashboardFullWidth = core.BoolPtr(true)
				uiMetaDataModel.NavigationOrder = []string{"testString"}
				uiMetaDataModel.NotCreatable = core.BoolPtr(true)
				uiMetaDataModel.Reservable = core.BoolPtr(true)
				uiMetaDataModel.PrimaryOfferingID = core.StringPtr("testString")
				uiMetaDataModel.AccessibleDuringProvision = core.BoolPtr(true)
				uiMetaDataModel.SideBySideIndex = core.Int64Ptr(int64(38))
				uiMetaDataModel.EndOfServiceTime = CreateMockDateTime()

				// Construct an instance of the Image model
				imageModel := new(globalcatalogv1.Image)
				imageModel.Image = core.StringPtr("testString")
				imageModel.SmallImage = core.StringPtr("testString")
				imageModel.MediumImage = core.StringPtr("testString")
				imageModel.FeatureImage = core.StringPtr("testString")

				// Construct an instance of the ObjectMetaData model
				objectMetaDataModel := new(globalcatalogv1.ObjectMetaData)
				objectMetaDataModel.RcCompatible = core.BoolPtr(true)
				objectMetaDataModel.Ui = uiMetaDataModel
				objectMetaDataModel.Pricing = pricingModel
				objectMetaDataModel.Compliance = []string{"testString"}
				objectMetaDataModel.Service = objectMetaDataServiceModel
				objectMetaDataModel.Plan = objectMetaDataPlanModel
				objectMetaDataModel.Template = objectMetaDataTemplateModel
				objectMetaDataModel.Deployment = objectMetaDataDeploymentModel
				objectMetaDataModel.Alias = objectMetaDataAliasModel
				objectMetaDataModel.Sla = objectMetaDataSlaModel
				objectMetaDataModel.Callbacks = callbacksModel
				objectMetaDataModel.Version = core.StringPtr("testString")
				objectMetaDataModel.OriginalName = core.StringPtr("testString")
				objectMetaDataModel.Other = CreateMockMap()

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

				// Construct an instance of the CatalogEntry model
				catalogEntryModel := new(globalcatalogv1.CatalogEntry)
				catalogEntryModel.Name = core.StringPtr("testString")
				catalogEntryModel.OverviewUi = overviewUiModel
				catalogEntryModel.Kind = core.StringPtr("testString")
				catalogEntryModel.Images = imageModel
				catalogEntryModel.ParentID = core.StringPtr("testString")
				catalogEntryModel.Disabled = core.BoolPtr(true)
				catalogEntryModel.Tags = []string{"testString"}
				catalogEntryModel.Group = core.BoolPtr(true)
				catalogEntryModel.Provider = providerModel
				catalogEntryModel.Metadata = objectMetaDataModel
				catalogEntryModel.Active = core.BoolPtr(true)
				catalogEntryModel.ParentURL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogEntryOptions model
				updateCatalogEntryOptionsModel := new(globalcatalogv1.UpdateCatalogEntryOptions)
				updateCatalogEntryOptionsModel.ID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewName = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewOverviewUi = overviewUiModel
				updateCatalogEntryOptionsModel.NewKind = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewImages = imageModel
				updateCatalogEntryOptionsModel.NewDisabled = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewProvider = providerModel
				updateCatalogEntryOptionsModel.NewParentID = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGroup = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewMetadata = objectMetaDataModel
				updateCatalogEntryOptionsModel.NewActive = core.BoolPtr(true)
				updateCatalogEntryOptionsModel.NewCatalogCrn = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewChildrenURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewParentURL = core.StringPtr("testString")
				updateCatalogEntryOptionsModel.NewGeoTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewPricingTags = []string{"testString"}
				updateCatalogEntryOptionsModel.NewCreated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewUpdated = CreateMockDateTime()
				updateCatalogEntryOptionsModel.NewChildren = []globalcatalogv1.CatalogEntry{*catalogEntryModel}
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `[{"page": "Page", "results_per_page": "ResultsPerPage", "total_results": "TotalResults", "resources": [{"anyKey": "anyValue"}]}]`)
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
                URL: "{{BAD_URL_STRING",
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
                URL: "{{BAD_URL_STRING",
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"restrictions": "Restrictions", "owner": {"type": "Type", "value": "Value"}, "include": {"accounts": {"_accountid_": "Accountid"}}, "exclude": {"accounts": {"_accountid_": "Accountid"}}, "approved": true}`)
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

				// Construct an instance of the Scope model
				scopeModel := new(globalcatalogv1.Scope)
				scopeModel.Type = core.StringPtr("testString")
				scopeModel.Value = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Owner = scopeModel
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

				// Construct an instance of the Scope model
				scopeModel := new(globalcatalogv1.Scope)
				scopeModel.Type = core.StringPtr("testString")
				scopeModel.Value = core.StringPtr("testString")

				// Construct an instance of the VisibilityDetail model
				visibilityDetailModel := new(globalcatalogv1.VisibilityDetail)
				visibilityDetailModel.Accounts = visibilityDetailAccountsModel

				// Construct an instance of the UpdateVisibilityOptions model
				updateVisibilityOptionsModel := new(globalcatalogv1.UpdateVisibilityOptions)
				updateVisibilityOptionsModel.ID = core.StringPtr("testString")
				updateVisibilityOptionsModel.Owner = scopeModel
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
                URL: "{{BAD_URL_STRING",
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
                URL: "{{BAD_URL_STRING",
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"type": "Type", "origin": "Origin", "starting_price": {"plan_id": "PlanID", "deployment_id": "DeploymentID", "amount": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}, "metrics": [{"metric_id": "MetricID", "tier_model": "TierModel", "charge_unit_name": "ChargeUnitName", "charge_unit_quantity": "ChargeUnitQuantity", "resource_display_name": "ResourceDisplayName", "charge_unit_display_name": "ChargeUnitDisplayName", "usage_cap_qty": 11, "amounts": [{"counrty": "Counrty", "currency": "Currency", "prices": [{"quantity_tier": 12, "Price": 5}]}]}]}`)
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
                URL: "{{BAD_URL_STRING",
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
                URL: "{{BAD_URL_STRING",
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"page": "Page", "results_per_page": "ResultsPerPage", "total_results": "TotalResults", "resources": [{"anyKey": "anyValue"}]}`)
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
                URL: "{{BAD_URL_STRING",
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
                URL: "{{BAD_URL_STRING",
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

					res.Header().Set("Content-type", "text/plain")
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
					fmt.Fprintf(res, `{"count": 5, "resources": [{"name": "Name", "updated": "Updated", "url": "URL", "etag": "Etag", "size": 4}]}`)
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
					Expect(req.URL.Query()["account"]).To(Equal([]string{"testString"}))

					res.WriteHeader(200)
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
				response, operationErr := testService.GetArtifact(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetArtifactOptions model
				getArtifactOptionsModel := new(globalcatalogv1.GetArtifactOptions)
				getArtifactOptionsModel.ObjectID = core.StringPtr("testString")
				getArtifactOptionsModel.ArtifactID = core.StringPtr("testString")
				getArtifactOptionsModel.Account = core.StringPtr("testString")
 				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
				getArtifactOptionsModel.Account = core.StringPtr("testString")
				getArtifactOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.GetArtifact(getArtifactOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the GetArtifactOptions model with no property values
				getArtifactOptionsModelNew := new(globalcatalogv1.GetArtifactOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.GetArtifact(getArtifactOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
			It(`Invoke NewCatalogEntry successfully`, func() {
				id := "testString"
				name := "testString"
				var overviewUi *globalcatalogv1.OverviewUI = nil
				kind := "testString"
				var images *globalcatalogv1.Image = nil
				disabled := true
				tags := []string{"testString"}
				var provider *globalcatalogv1.Provider = nil
				_, err := testService.NewCatalogEntry(id, name, overviewUi, kind, images, disabled, tags, provider)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewImage successfully`, func() {
				image := "testString"
				model, err := testService.NewImage(image)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewVisibilityDetail successfully`, func() {
				var accounts *globalcatalogv1.VisibilityDetailAccounts = nil
				_, err := testService.NewVisibilityDetail(accounts)
				Expect(err).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
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

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
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
