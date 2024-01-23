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

package partnerusagereportsv1_test

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
	"github.com/IBM/platform-services-go-sdk/partnerusagereportsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PartnerUsageReportsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(partnerUsageReportsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(partnerUsageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				URL: "https://partnerusagereportsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(partnerUsageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_USAGE_REPORTS_URL": "https://partnerusagereportsv1/api",
				"PARTNER_USAGE_REPORTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				})
				Expect(partnerUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := partnerUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerUsageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL: "https://testService/api",
				})
				Expect(partnerUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerUsageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				})
				err := partnerUsageReportsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerUsageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerUsageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerUsageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerUsageReportsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_USAGE_REPORTS_URL": "https://partnerusagereportsv1/api",
				"PARTNER_USAGE_REPORTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(&partnerusagereportsv1.PartnerUsageReportsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerUsageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_USAGE_REPORTS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1UsingExternalConfig(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerUsageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = partnerusagereportsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) - Operation response error`, func() {
		getResourceUsageReportPath := "/v1/resource-usage-reports"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["viewpoint"]).To(Equal([]string{"DISTRIBUTOR"}))
					// TODO: Add check for recurse query parameter
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceUsageReport with error: Operation response processing error`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2022-04")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerUsageReportsService.EnableRetries(0, 0)
				result, response, operationErr = partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions)`, func() {
		getResourceUsageReportPath := "/v1/resource-usage-reports"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["viewpoint"]).To(Equal([]string{"DISTRIBUTOR"}))
					// TODO: Add check for recurse query parameter
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "reports": [{"entity_id": "<distributor_enterprise_id>", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5", "entity_name": "Arrow", "entity_partner_type": "DISTRIBUTOR", "viewpoint": "DISTRIBUTOR", "month": "2022-04", "currency_code": "EUR", "country_code": "FRA", "billable_cost": 2331828.33275813, "billable_rated_cost": 3817593.35186263, "non_billable_cost": 0, "non_billable_rated_cost": 0, "resources": [{"resource_id": "cloudant", "resource_name": "Cloudant", "billable_cost": 75, "billable_rated_cost": 75, "non_billable_cost": 0, "non_billable_rated_cost": 0, "plans": [{"plan_id": "cloudant-standard", "pricing_region": "Standard", "pricing_plan_id": "billable:v4:cloudant-standard::1552694400000:", "billable": true, "cost": 75, "rated_cost": 75, "usage": [{"metric": "GB_STORAGE_ACCRUED_PER_MONTH", "unit": "GIGABYTE_MONTHS", "quantity": 10, "rateable_quantity": 10, "cost": 10, "rated_cost": 10, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully with retries`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())
				partnerUsageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2022-04")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerUsageReportsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerUsageReportsService.DisableRetries()
				result, response, operationErr := partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerUsageReportsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["partner_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["reseller_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					// TODO: Add check for children query parameter
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2022-04"}))
					Expect(req.URL.Query()["viewpoint"]).To(Equal([]string{"DISTRIBUTOR"}))
					// TODO: Add check for recurse query parameter
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "reports": [{"entity_id": "<distributor_enterprise_id>", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5", "entity_name": "Arrow", "entity_partner_type": "DISTRIBUTOR", "viewpoint": "DISTRIBUTOR", "month": "2022-04", "currency_code": "EUR", "country_code": "FRA", "billable_cost": 2331828.33275813, "billable_rated_cost": 3817593.35186263, "non_billable_cost": 0, "non_billable_rated_cost": 0, "resources": [{"resource_id": "cloudant", "resource_name": "Cloudant", "billable_cost": 75, "billable_rated_cost": 75, "non_billable_cost": 0, "non_billable_rated_cost": 0, "plans": [{"plan_id": "cloudant-standard", "pricing_region": "Standard", "pricing_plan_id": "billable:v4:cloudant-standard::1552694400000:", "billable": true, "cost": 75, "rated_cost": 75, "usage": [{"metric": "GB_STORAGE_ACCRUED_PER_MONTH", "unit": "GIGABYTE_MONTHS", "quantity": 10, "rateable_quantity": 10, "cost": 10, "rated_cost": 10, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerUsageReportsService.GetResourceUsageReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2022-04")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageReport with error: Operation validation and request error`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2022-04")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerUsageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceUsageReportOptions model with no property values
				getResourceUsageReportOptionsModelNew := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModelNew)
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
			It(`Invoke GetResourceUsageReport successfully`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnerusagereportsv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2022-04")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(partnerusagereportsv1.PartnerUsageReportSummary)
				nextObject := new(partnerusagereportsv1.PartnerUsageReportSummaryNext)
				nextObject.Offset = core.StringPtr("abc-123")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(partnerusagereportsv1.PartnerUsageReportSummary)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageReportPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"offset":"1"},"reports":[{"entity_id":"<distributor_enterprise_id>","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5","entity_name":"Arrow","entity_partner_type":"DISTRIBUTOR","viewpoint":"DISTRIBUTOR","month":"2022-04","currency_code":"EUR","country_code":"FRA","billable_cost":2331828.33275813,"billable_rated_cost":3817593.35186263,"non_billable_cost":0,"non_billable_rated_cost":0,"resources":[{"resource_id":"cloudant","resource_name":"Cloudant","billable_cost":75,"billable_rated_cost":75,"non_billable_cost":0,"non_billable_rated_cost":0,"plans":[{"plan_id":"cloudant-standard","pricing_region":"Standard","pricing_plan_id":"billable:v4:cloudant-standard::1552694400000:","billable":true,"cost":75,"rated_cost":75,"usage":[{"metric":"GB_STORAGE_ACCRUED_PER_MONTH","unit":"GIGABYTE_MONTHS","quantity":10,"rateable_quantity":10,"cost":10,"rated_cost":10,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"reports":[{"entity_id":"<distributor_enterprise_id>","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5","entity_name":"Arrow","entity_partner_type":"DISTRIBUTOR","viewpoint":"DISTRIBUTOR","month":"2022-04","currency_code":"EUR","country_code":"FRA","billable_cost":2331828.33275813,"billable_rated_cost":3817593.35186263,"non_billable_cost":0,"non_billable_rated_cost":0,"resources":[{"resource_id":"cloudant","resource_name":"Cloudant","billable_cost":75,"billable_rated_cost":75,"non_billable_cost":0,"non_billable_rated_cost":0,"plans":[{"plan_id":"cloudant-standard","pricing_region":"Standard","pricing_plan_id":"billable:v4:cloudant-standard::1552694400000:","billable":true,"cost":75,"rated_cost":75,"usage":[{"metric":"GB_STORAGE_ACCRUED_PER_MONTH","unit":"GIGABYTE_MONTHS","quantity":10,"rateable_quantity":10,"cost":10,"rated_cost":10,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use GetResourceUsageReportPager.GetNext successfully`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &partnerusagereportsv1.GetResourceUsageReportOptions{
					PartnerID: core.StringPtr("testString"),
					ResellerID: core.StringPtr("testString"),
					CustomerID: core.StringPtr("testString"),
					Children: core.BoolPtr(false),
					Month: core.StringPtr("2022-04"),
					Viewpoint: core.StringPtr("DISTRIBUTOR"),
					Recurse: core.BoolPtr(false),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := partnerUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []partnerusagereportsv1.PartnerUsageReport
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use GetResourceUsageReportPager.GetAll successfully`, func() {
				partnerUsageReportsService, serviceErr := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerUsageReportsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &partnerusagereportsv1.GetResourceUsageReportOptions{
					PartnerID: core.StringPtr("testString"),
					ResellerID: core.StringPtr("testString"),
					CustomerID: core.StringPtr("testString"),
					Children: core.BoolPtr(false),
					Month: core.StringPtr("2022-04"),
					Viewpoint: core.StringPtr("DISTRIBUTOR"),
					Recurse: core.BoolPtr(false),
					Limit: core.Int64Ptr(int64(10)),
				}

				pager, err := partnerUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			partnerUsageReportsService, _ := partnerusagereportsv1.NewPartnerUsageReportsV1(&partnerusagereportsv1.PartnerUsageReportsV1Options{
				URL:           "http://partnerusagereportsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetResourceUsageReportOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageReportOptions model
				partnerID := "testString"
				getResourceUsageReportOptionsModel := partnerUsageReportsService.NewGetResourceUsageReportOptions(partnerID)
				getResourceUsageReportOptionsModel.SetPartnerID("testString")
				getResourceUsageReportOptionsModel.SetResellerID("testString")
				getResourceUsageReportOptionsModel.SetCustomerID("testString")
				getResourceUsageReportOptionsModel.SetChildren(false)
				getResourceUsageReportOptionsModel.SetMonth("2022-04")
				getResourceUsageReportOptionsModel.SetViewpoint("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.SetRecurse(false)
				getResourceUsageReportOptionsModel.SetLimit(int64(10))
				getResourceUsageReportOptionsModel.SetOffset("testString")
				getResourceUsageReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceUsageReportOptionsModel).ToNot(BeNil())
				Expect(getResourceUsageReportOptionsModel.PartnerID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.ResellerID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.Children).To(Equal(core.BoolPtr(false)))
				Expect(getResourceUsageReportOptionsModel.Month).To(Equal(core.StringPtr("2022-04")))
				Expect(getResourceUsageReportOptionsModel.Viewpoint).To(Equal(core.StringPtr("DISTRIBUTOR")))
				Expect(getResourceUsageReportOptionsModel.Recurse).To(Equal(core.BoolPtr(false)))
				Expect(getResourceUsageReportOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(getResourceUsageReportOptionsModel.Offset).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
