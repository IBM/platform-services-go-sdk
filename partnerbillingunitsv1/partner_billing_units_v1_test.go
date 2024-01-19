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

package partnerbillingunitsv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnerbillingunitsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PartnerBillingUnitsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(partnerBillingUnitsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(partnerBillingUnitsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
				URL: "https://partnerbillingunitsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(partnerBillingUnitsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_BILLING_UNITS_URL":       "https://partnerbillingunitsv1/api",
				"PARTNER_BILLING_UNITS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{})
				Expect(partnerBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := partnerBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerBillingUnitsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL: "https://testService/api",
				})
				Expect(partnerBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerBillingUnitsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{})
				err := partnerBillingUnitsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerBillingUnitsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerBillingUnitsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerBillingUnitsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerBillingUnitsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_BILLING_UNITS_URL":       "https://partnerbillingunitsv1/api",
				"PARTNER_BILLING_UNITS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(partnerBillingUnitsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_BILLING_UNITS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1UsingExternalConfig(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerBillingUnitsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = partnerbillingunitsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions) - Operation response error`, func() {
		getBillingOptionsPath := "/v1/billing-options"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBillingOptions with error: Operation response processing error`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2022-04")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(30))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
		getBillingOptionsPath := "/v1/billing-options"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"id": "CFL_JJKLVZ2I0JE-_MGU", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2022-04", "errors": [{"anyKey": "anyValue"}], "type": "SUBSCRIPTION", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "part_number": "<PART_NUMBER_1>", "catalog_id": "ibmcloud-platform-payg-commit", "order_id": "23wzpnpmh8", "po_number": "<PO_NUMBER_1>", "subscription_model": "4.0", "duration_in_months": 11, "monthly_amount": 8333.333333333334, "billing_system": {"anyKey": "anyValue"}, "country_code": "USA", "currency_code": "USD"}]}`)
				}))
			})
			It(`Invoke GetBillingOptions successfully with retries`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())
				partnerBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2022-04")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(30))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerBillingUnitsService.GetBillingOptionsWithContext(ctx, getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerBillingUnitsService.DisableRetries()
				result, response, operationErr := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerBillingUnitsService.GetBillingOptionsWithContext(ctx, getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBillingOptionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"id": "CFL_JJKLVZ2I0JE-_MGU", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2022-04", "errors": [{"anyKey": "anyValue"}], "type": "SUBSCRIPTION", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "part_number": "<PART_NUMBER_1>", "catalog_id": "ibmcloud-platform-payg-commit", "order_id": "23wzpnpmh8", "po_number": "<PO_NUMBER_1>", "subscription_model": "4.0", "duration_in_months": 11, "monthly_amount": 8333.333333333334, "billing_system": {"anyKey": "anyValue"}, "country_code": "USA", "currency_code": "USD"}]}`)
				}))
			})
			It(`Invoke GetBillingOptions successfully`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerBillingUnitsService.GetBillingOptions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2022-04")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(30))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBillingOptions with error: Operation validation and request error`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2022-04")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(30))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBillingOptionsOptions model with no property values
				getBillingOptionsOptionsModelNew := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetBillingOptions successfully`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnerbillingunitsv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2022-04")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(30))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerBillingUnitsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions) - Operation response error`, func() {
		getCreditPoolsReportPath := "/v1/credit-pools"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCreditPoolsReport with error: Operation response processing error`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2022-04")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerBillingUnitsService.EnableRetries(0, 0)
				result, response, operationErr = partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
		getCreditPoolsReportPath := "/v1/credit-pools"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"type": "PLATFORM", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2022-04", "currency_code": "USD", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "billing_option_model": "4.0", "category": "PLATFORM", "start_date": "2019-07-01T00:00:00.000Z", "end_date": "2019-08-31T23:59:59.999Z", "total_credits": 100000, "starting_balance": 100000, "used_credits": 0, "current_balance": 100000, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPoolsReport successfully with retries`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())
				partnerBillingUnitsService.EnableRetries(0, 0)

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2022-04")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerBillingUnitsService.GetCreditPoolsReportWithContext(ctx, getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerBillingUnitsService.DisableRetries()
				result, response, operationErr := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerBillingUnitsService.GetCreditPoolsReportWithContext(ctx, getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCreditPoolsReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"type": "PLATFORM", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2022-04", "currency_code": "USD", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "billing_option_model": "4.0", "category": "PLATFORM", "start_date": "2019-07-01T00:00:00.000Z", "end_date": "2019-08-31T23:59:59.999Z", "total_credits": 100000, "starting_balance": 100000, "used_credits": 0, "current_balance": 100000, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPoolsReport successfully`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerBillingUnitsService.GetCreditPoolsReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2022-04")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCreditPoolsReport with error: Operation validation and request error`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2022-04")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerBillingUnitsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCreditPoolsReportOptions model with no property values
				getCreditPoolsReportOptionsModelNew := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCreditPoolsReport successfully`, func() {
				partnerBillingUnitsService, serviceErr := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerBillingUnitsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnerbillingunitsv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2022-04")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerBillingUnitsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			partnerBillingUnitsService, _ := partnerbillingunitsv1.NewPartnerBillingUnitsV1(&partnerbillingunitsv1.PartnerBillingUnitsV1Options{
				URL:           "http://partnerbillingunitsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetBillingOptionsOptions successfully`, func() {
				// Construct an instance of the GetBillingOptionsOptions model
				partnerID := "testString"
				billingMonth := "testString"
				getBillingOptionsOptionsModel := partnerBillingUnitsService.NewGetBillingOptionsOptions(partnerID, billingMonth)
				getBillingOptionsOptionsModel.SetPartnerID("testString")
				getBillingOptionsOptionsModel.SetCustomerID("testString")
				getBillingOptionsOptionsModel.SetResellerID("testString")
				getBillingOptionsOptionsModel.SetDate("2022-04")
				getBillingOptionsOptionsModel.SetLimit(int64(30))
				getBillingOptionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBillingOptionsOptionsModel).ToNot(BeNil())
				Expect(getBillingOptionsOptionsModel.PartnerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.ResellerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.Date).To(Equal(core.StringPtr("2022-04")))
				Expect(getBillingOptionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(30))))
				Expect(getBillingOptionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCreditPoolsReportOptions successfully`, func() {
				// Construct an instance of the GetCreditPoolsReportOptions model
				partnerID := "testString"
				billingMonth := "testString"
				getCreditPoolsReportOptionsModel := partnerBillingUnitsService.NewGetCreditPoolsReportOptions(partnerID, billingMonth)
				getCreditPoolsReportOptionsModel.SetPartnerID("testString")
				getCreditPoolsReportOptionsModel.SetCustomerID("testString")
				getCreditPoolsReportOptionsModel.SetResellerID("testString")
				getCreditPoolsReportOptionsModel.SetDate("2022-04")
				getCreditPoolsReportOptionsModel.SetLimit(int64(30))
				getCreditPoolsReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCreditPoolsReportOptionsModel).ToNot(BeNil())
				Expect(getCreditPoolsReportOptionsModel.PartnerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.ResellerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.Date).To(Equal(core.StringPtr("2022-04")))
				Expect(getCreditPoolsReportOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(30))))
				Expect(getCreditPoolsReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
