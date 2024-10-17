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

package partnermanagementapisv1_test

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
	"github.com/IBM/platform-services-go-sdk/partnermanagementapisv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PartnerManagementAPIsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(partnerManagementAPIsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(partnerManagementAPIsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
				URL: "https://partnermanagementapisv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(partnerManagementAPIsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_MANAGEMENT_APIS_URL":       "https://partnermanagementapisv1/api",
				"PARTNER_MANAGEMENT_APIS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1UsingExternalConfig(&partnermanagementapisv1.PartnerManagementAPIsV1Options{})
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := partnerManagementAPIsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerManagementAPIsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerManagementAPIsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerManagementAPIsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1UsingExternalConfig(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL: "https://testService/api",
				})
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerManagementAPIsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerManagementAPIsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerManagementAPIsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerManagementAPIsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1UsingExternalConfig(&partnermanagementapisv1.PartnerManagementAPIsV1Options{})
				err := partnerManagementAPIsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerManagementAPIsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerManagementAPIsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerManagementAPIsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerManagementAPIsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_MANAGEMENT_APIS_URL":       "https://partnermanagementapisv1/api",
				"PARTNER_MANAGEMENT_APIS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1UsingExternalConfig(&partnermanagementapisv1.PartnerManagementAPIsV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(partnerManagementAPIsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_MANAGEMENT_APIS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1UsingExternalConfig(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerManagementAPIsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = partnermanagementapisv1.GetServiceURLForRegion("INVALID_REGION")
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
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2024-01"}))
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
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2024-01")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerManagementAPIsService.EnableRetries(0, 0)
				result, response, operationErr = partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
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
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["viewpoint"]).To(Equal([]string{"DISTRIBUTOR"}))
					// TODO: Add check for recurse query parameter
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "reports": [{"entity_id": "<distributor_enterprise_id>", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5", "entity_name": "Company", "entity_partner_type": "DISTRIBUTOR", "viewpoint": "DISTRIBUTOR", "month": "2024-01", "currency_code": "EUR", "country_code": "FRA", "billable_cost": 2331828.33275813, "billable_rated_cost": 3817593.35186263, "non_billable_cost": 0, "non_billable_rated_cost": 0, "resources": [{"resource_id": "cloudant", "resource_name": "Cloudant", "billable_cost": 75, "billable_rated_cost": 75, "non_billable_cost": 0, "non_billable_rated_cost": 0, "plans": [{"plan_id": "cloudant-standard", "pricing_region": "Standard", "pricing_plan_id": "billable:v4:cloudant-standard::1552694400000:", "billable": true, "cost": 75, "rated_cost": 75, "usage": [{"metric": "GB_STORAGE_ACCRUED_PER_MONTH", "unit": "GIGABYTE_MONTHS", "quantity": 10, "rateable_quantity": 10, "cost": 10, "rated_cost": 10, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully with retries`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				partnerManagementAPIsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2024-01")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerManagementAPIsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerManagementAPIsService.DisableRetries()
				result, response, operationErr := partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerManagementAPIsService.GetResourceUsageReportWithContext(ctx, getResourceUsageReportOptionsModel)
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
					Expect(req.URL.Query()["month"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["viewpoint"]).To(Equal([]string{"DISTRIBUTOR"}))
					// TODO: Add check for recurse query parameter
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "reports": [{"entity_id": "<distributor_enterprise_id>", "entity_type": "enterprise", "entity_crn": "crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5", "entity_name": "Company", "entity_partner_type": "DISTRIBUTOR", "viewpoint": "DISTRIBUTOR", "month": "2024-01", "currency_code": "EUR", "country_code": "FRA", "billable_cost": 2331828.33275813, "billable_rated_cost": 3817593.35186263, "non_billable_cost": 0, "non_billable_rated_cost": 0, "resources": [{"resource_id": "cloudant", "resource_name": "Cloudant", "billable_cost": 75, "billable_rated_cost": 75, "non_billable_cost": 0, "non_billable_rated_cost": 0, "plans": [{"plan_id": "cloudant-standard", "pricing_region": "Standard", "pricing_plan_id": "billable:v4:cloudant-standard::1552694400000:", "billable": true, "cost": 75, "rated_cost": 75, "usage": [{"metric": "GB_STORAGE_ACCRUED_PER_MONTH", "unit": "GIGABYTE_MONTHS", "quantity": 10, "rateable_quantity": 10, "cost": 10, "rated_cost": 10, "price": [{"anyKey": "anyValue"}]}]}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageReport successfully`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerManagementAPIsService.GetResourceUsageReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2024-01")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageReport with error: Operation validation and request error`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2024-01")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerManagementAPIsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceUsageReportOptions model with no property values
				getResourceUsageReportOptionsModelNew := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModelNew)
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
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageReportOptions model
				getResourceUsageReportOptionsModel := new(partnermanagementapisv1.GetResourceUsageReportOptions)
				getResourceUsageReportOptionsModel.PartnerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.ResellerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.CustomerID = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Children = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Month = core.StringPtr("2024-01")
				getResourceUsageReportOptionsModel.Viewpoint = core.StringPtr("DISTRIBUTOR")
				getResourceUsageReportOptionsModel.Recurse = core.BoolPtr(false)
				getResourceUsageReportOptionsModel.Limit = core.Int64Ptr(int64(10))
				getResourceUsageReportOptionsModel.Offset = core.StringPtr("testString")
				getResourceUsageReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerManagementAPIsService.GetResourceUsageReport(getResourceUsageReportOptionsModel)
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
				responseObject := new(partnermanagementapisv1.PartnerUsageReportSummary)
				nextObject := new(partnermanagementapisv1.PartnerUsageReportSummaryNext)
				nextObject.Offset = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(partnermanagementapisv1.PartnerUsageReportSummary)

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
						fmt.Fprintf(res, "%s", `{"next":{"offset":"1"},"reports":[{"entity_id":"<distributor_enterprise_id>","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5","entity_name":"Company","entity_partner_type":"DISTRIBUTOR","viewpoint":"DISTRIBUTOR","month":"2024-01","currency_code":"EUR","country_code":"FRA","billable_cost":2331828.33275813,"billable_rated_cost":3817593.35186263,"non_billable_cost":0,"non_billable_rated_cost":0,"resources":[{"resource_id":"cloudant","resource_name":"Cloudant","billable_cost":75,"billable_rated_cost":75,"non_billable_cost":0,"non_billable_rated_cost":0,"plans":[{"plan_id":"cloudant-standard","pricing_region":"Standard","pricing_plan_id":"billable:v4:cloudant-standard::1552694400000:","billable":true,"cost":75,"rated_cost":75,"usage":[{"metric":"GB_STORAGE_ACCRUED_PER_MONTH","unit":"GIGABYTE_MONTHS","quantity":10,"rateable_quantity":10,"cost":10,"rated_cost":10,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"reports":[{"entity_id":"<distributor_enterprise_id>","entity_type":"enterprise","entity_crn":"crn:v1:bluemix:public:enterprise::a/fa359b76ff2c41eda727aad47b7e4063::enterprise:33a7eb04e7d547cd9489e90c99d476a5","entity_name":"Company","entity_partner_type":"DISTRIBUTOR","viewpoint":"DISTRIBUTOR","month":"2024-01","currency_code":"EUR","country_code":"FRA","billable_cost":2331828.33275813,"billable_rated_cost":3817593.35186263,"non_billable_cost":0,"non_billable_rated_cost":0,"resources":[{"resource_id":"cloudant","resource_name":"Cloudant","billable_cost":75,"billable_rated_cost":75,"non_billable_cost":0,"non_billable_rated_cost":0,"plans":[{"plan_id":"cloudant-standard","pricing_region":"Standard","pricing_plan_id":"billable:v4:cloudant-standard::1552694400000:","billable":true,"cost":75,"rated_cost":75,"usage":[{"metric":"GB_STORAGE_ACCRUED_PER_MONTH","unit":"GIGABYTE_MONTHS","quantity":10,"rateable_quantity":10,"cost":10,"rated_cost":10,"price":[{"anyKey":"anyValue"}]}]}]}]}],"total_count":2,"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use GetResourceUsageReportPager.GetNext successfully`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &partnermanagementapisv1.GetResourceUsageReportOptions{
					PartnerID:  core.StringPtr("testString"),
					ResellerID: core.StringPtr("testString"),
					CustomerID: core.StringPtr("testString"),
					Children:   core.BoolPtr(false),
					Month:      core.StringPtr("2024-01"),
					Viewpoint:  core.StringPtr("DISTRIBUTOR"),
					Recurse:    core.BoolPtr(false),
					Limit:      core.Int64Ptr(int64(10)),
				}

				pager, err := partnerManagementAPIsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []partnermanagementapisv1.PartnerUsageReport
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use GetResourceUsageReportPager.GetAll successfully`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				getResourceUsageReportOptionsModel := &partnermanagementapisv1.GetResourceUsageReportOptions{
					PartnerID:  core.StringPtr("testString"),
					ResellerID: core.StringPtr("testString"),
					CustomerID: core.StringPtr("testString"),
					Children:   core.BoolPtr(false),
					Month:      core.StringPtr("2024-01"),
					Viewpoint:  core.StringPtr("DISTRIBUTOR"),
					Recurse:    core.BoolPtr(false),
					Limit:      core.Int64Ptr(int64(10)),
				}

				pager, err := partnerManagementAPIsService.NewGetResourceUsageReportPager(getResourceUsageReportOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBillingOptions with error: Operation response processing error`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnermanagementapisv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2024-01")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(200))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerManagementAPIsService.EnableRetries(0, 0)
				result, response, operationErr = partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": {"href": "Href", "offset": "Offset"}, "resources": [{"id": "CFL_JJKLVZ2I0JE-_MGU", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2024-01", "errors": [{"anyKey": "anyValue"}], "type": "SUBSCRIPTION", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "part_number": "<PART_NUMBER_1>", "catalog_id": "ibmcloud-platform-payg-commit", "order_id": "23wzpnpmh8", "po_number": "<PO_NUMBER_1>", "subscription_model": "4.0", "duration_in_months": 11, "monthly_amount": 8333.333333333334, "billing_system": {"anyKey": "anyValue"}, "country_code": "USA", "currency_code": "USD"}]}`)
				}))
			})
			It(`Invoke GetBillingOptions successfully with retries`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				partnerManagementAPIsService.EnableRetries(0, 0)

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnermanagementapisv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2024-01")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(200))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerManagementAPIsService.GetBillingOptionsWithContext(ctx, getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerManagementAPIsService.DisableRetries()
				result, response, operationErr := partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerManagementAPIsService.GetBillingOptionsWithContext(ctx, getBillingOptionsOptionsModel)
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(200))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": {"href": "Href", "offset": "Offset"}, "resources": [{"id": "CFL_JJKLVZ2I0JE-_MGU", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2024-01", "errors": [{"anyKey": "anyValue"}], "type": "SUBSCRIPTION", "start_date": "2019-05-01T00:00:00.000Z", "end_date": "2020-05-01T00:00:00.000Z", "state": "ACTIVE", "category": "PLATFORM", "payment_instrument": {"anyKey": "anyValue"}, "part_number": "<PART_NUMBER_1>", "catalog_id": "ibmcloud-platform-payg-commit", "order_id": "23wzpnpmh8", "po_number": "<PO_NUMBER_1>", "subscription_model": "4.0", "duration_in_months": 11, "monthly_amount": 8333.333333333334, "billing_system": {"anyKey": "anyValue"}, "country_code": "USA", "currency_code": "USD"}]}`)
				}))
			})
			It(`Invoke GetBillingOptions successfully`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerManagementAPIsService.GetBillingOptions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnermanagementapisv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2024-01")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(200))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBillingOptions with error: Operation validation and request error`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnermanagementapisv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2024-01")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(200))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerManagementAPIsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBillingOptionsOptions model with no property values
				getBillingOptionsOptionsModelNew := new(partnermanagementapisv1.GetBillingOptionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModelNew)
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
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetBillingOptionsOptions model
				getBillingOptionsOptionsModel := new(partnermanagementapisv1.GetBillingOptionsOptions)
				getBillingOptionsOptionsModel.PartnerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.CustomerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.ResellerID = core.StringPtr("testString")
				getBillingOptionsOptionsModel.Date = core.StringPtr("2024-01")
				getBillingOptionsOptionsModel.Limit = core.Int64Ptr(int64(200))
				getBillingOptionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerManagementAPIsService.GetBillingOptions(getBillingOptionsOptionsModel)
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCreditPoolsReport with error: Operation response processing error`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2024-01")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerManagementAPIsService.EnableRetries(0, 0)
				result, response, operationErr = partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": {"href": "Href", "offset": "Offset"}, "resources": [{"type": "PLATFORM", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2024-01", "currency_code": "USD", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "billing_option_model": "4.0", "category": "PLATFORM", "start_date": "2019-07-01T00:00:00.000Z", "end_date": "2019-08-31T23:59:59.999Z", "total_credits": 100000, "starting_balance": 100000, "used_credits": 0, "current_balance": 100000, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPoolsReport successfully with retries`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())
				partnerManagementAPIsService.EnableRetries(0, 0)

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2024-01")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerManagementAPIsService.GetCreditPoolsReportWithContext(ctx, getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerManagementAPIsService.DisableRetries()
				result, response, operationErr := partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerManagementAPIsService.GetCreditPoolsReportWithContext(ctx, getCreditPoolsReportOptionsModel)
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
					Expect(req.URL.Query()["date"]).To(Equal([]string{"2024-01"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(30))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rows_count": 9, "next_url": {"href": "Href", "offset": "Offset"}, "resources": [{"type": "PLATFORM", "billing_unit_id": "e19fa97c9bb34963a31a2008044d8b59", "customer_id": "<ford_account_id>", "customer_type": "ACCOUNT", "customer_name": "Ford", "reseller_id": "<techdata_enterprise_id>", "reseller_name": "TechData", "month": "2024-01", "currency_code": "USD", "term_credits": [{"billing_option_id": "JWX986YRGFSHACQUEFOI", "billing_option_model": "4.0", "category": "PLATFORM", "start_date": "2019-07-01T00:00:00.000Z", "end_date": "2019-08-31T23:59:59.999Z", "total_credits": 100000, "starting_balance": 100000, "used_credits": 0, "current_balance": 100000, "resources": [{"anyKey": "anyValue"}]}], "overage": {"cost": 500, "resources": [{"anyKey": "anyValue"}]}}]}`)
				}))
			})
			It(`Invoke GetCreditPoolsReport successfully`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerManagementAPIsService.GetCreditPoolsReport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2024-01")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCreditPoolsReport with error: Operation validation and request error`, func() {
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2024-01")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerManagementAPIsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCreditPoolsReportOptions model with no property values
				getCreditPoolsReportOptionsModelNew := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModelNew)
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
				partnerManagementAPIsService, serviceErr := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerManagementAPIsService).ToNot(BeNil())

				// Construct an instance of the GetCreditPoolsReportOptions model
				getCreditPoolsReportOptionsModel := new(partnermanagementapisv1.GetCreditPoolsReportOptions)
				getCreditPoolsReportOptionsModel.PartnerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.CustomerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.ResellerID = core.StringPtr("testString")
				getCreditPoolsReportOptionsModel.Date = core.StringPtr("2024-01")
				getCreditPoolsReportOptionsModel.Limit = core.Int64Ptr(int64(30))
				getCreditPoolsReportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerManagementAPIsService.GetCreditPoolsReport(getCreditPoolsReportOptionsModel)
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
			partnerManagementAPIsService, _ := partnermanagementapisv1.NewPartnerManagementAPIsV1(&partnermanagementapisv1.PartnerManagementAPIsV1Options{
				URL:           "http://partnermanagementapisv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetBillingOptionsOptions successfully`, func() {
				// Construct an instance of the GetBillingOptionsOptions model
				partnerID := "testString"
				billingMonth := "testString"
				getBillingOptionsOptionsModel := partnerManagementAPIsService.NewGetBillingOptionsOptions(partnerID, billingMonth)
				getBillingOptionsOptionsModel.SetPartnerID("testString")
				getBillingOptionsOptionsModel.SetCustomerID("testString")
				getBillingOptionsOptionsModel.SetResellerID("testString")
				getBillingOptionsOptionsModel.SetDate("2024-01")
				getBillingOptionsOptionsModel.SetLimit(int64(200))
				getBillingOptionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBillingOptionsOptionsModel).ToNot(BeNil())
				Expect(getBillingOptionsOptionsModel.PartnerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.ResellerID).To(Equal(core.StringPtr("testString")))
				Expect(getBillingOptionsOptionsModel.Date).To(Equal(core.StringPtr("2024-01")))
				Expect(getBillingOptionsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(200))))
				Expect(getBillingOptionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCreditPoolsReportOptions successfully`, func() {
				// Construct an instance of the GetCreditPoolsReportOptions model
				partnerID := "testString"
				billingMonth := "testString"
				getCreditPoolsReportOptionsModel := partnerManagementAPIsService.NewGetCreditPoolsReportOptions(partnerID, billingMonth)
				getCreditPoolsReportOptionsModel.SetPartnerID("testString")
				getCreditPoolsReportOptionsModel.SetCustomerID("testString")
				getCreditPoolsReportOptionsModel.SetResellerID("testString")
				getCreditPoolsReportOptionsModel.SetDate("2024-01")
				getCreditPoolsReportOptionsModel.SetLimit(int64(30))
				getCreditPoolsReportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCreditPoolsReportOptionsModel).ToNot(BeNil())
				Expect(getCreditPoolsReportOptionsModel.PartnerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.ResellerID).To(Equal(core.StringPtr("testString")))
				Expect(getCreditPoolsReportOptionsModel.Date).To(Equal(core.StringPtr("2024-01")))
				Expect(getCreditPoolsReportOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(30))))
				Expect(getCreditPoolsReportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceUsageReportOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageReportOptions model
				partnerID := "testString"
				getResourceUsageReportOptionsModel := partnerManagementAPIsService.NewGetResourceUsageReportOptions(partnerID)
				getResourceUsageReportOptionsModel.SetPartnerID("testString")
				getResourceUsageReportOptionsModel.SetResellerID("testString")
				getResourceUsageReportOptionsModel.SetCustomerID("testString")
				getResourceUsageReportOptionsModel.SetChildren(false)
				getResourceUsageReportOptionsModel.SetMonth("2024-01")
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
				Expect(getResourceUsageReportOptionsModel.Month).To(Equal(core.StringPtr("2024-01")))
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
	ba := []byte(mockData)
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
