/**
 * (C) Copyright IBM Corp. 2021.
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

package usagereportsv4_test

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
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`UsageReportsV4`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(usageReportsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(usageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
				URL: "https://usagereportsv4/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(usageReportsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_REPORTS_URL":       "https://usagereportsv4/api",
				"USAGE_REPORTS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4UsingExternalConfig(&usagereportsv4.UsageReportsV4Options{})
				Expect(usageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := usageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4UsingExternalConfig(&usagereportsv4.UsageReportsV4Options{
					URL: "https://testService/api",
				})
				Expect(usageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := usageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageReportsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4UsingExternalConfig(&usagereportsv4.UsageReportsV4Options{})
				err := usageReportsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := usageReportsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != usageReportsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(usageReportsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(usageReportsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_REPORTS_URL":       "https://usagereportsv4/api",
				"USAGE_REPORTS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4UsingExternalConfig(&usagereportsv4.UsageReportsV4Options{})

			It(`Instantiate service client with error`, func() {
				Expect(usageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USAGE_REPORTS_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4UsingExternalConfig(&usagereportsv4.UsageReportsV4Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(usageReportsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = usagereportsv4.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions) - Operation response error`, func() {
		getAccountSummaryPath := "/v4/accounts/testString/summary/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSummaryPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSummary with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions)`, func() {
		getAccountSummaryPath := "/v4/accounts/testString/summary/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "billing_month": "BillingMonth", "billing_country_code": "BillingCountryCode", "billing_currency_code": "BillingCurrencyCode", "resources": {"billable_cost": 12, "non_billable_cost": 15}, "offers": [{"offer_id": "OfferID", "credits_total": 12, "offer_template": "OfferTemplate", "valid_from": "2019-01-01T12:00:00.000Z", "expires_on": "2019-01-01T12:00:00.000Z", "credits": {"starting_balance": 15, "used": 4, "balance": 7}}], "support": [{"cost": 4, "type": "Type", "overage": 7}], "subscription": {"overage": 7, "subscriptions": [{"subscription_id": "SubscriptionID", "charge_agreement_number": "ChargeAgreementNumber", "type": "Type", "subscription_amount": 18, "start": "2019-01-01T12:00:00.000Z", "end": "2019-01-01T12:00:00.000Z", "credits_total": 12, "terms": [{"start": "2019-01-01T12:00:00.000Z", "end": "2019-01-01T12:00:00.000Z", "credits": {"total": 5, "starting_balance": 15, "used": 4, "balance": 7}}]}]}}`)
				}))
			})
			It(`Invoke GetAccountSummary successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetAccountSummaryWithContext(ctx, getAccountSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetAccountSummaryWithContext(ctx, getAccountSummaryOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSummaryPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "billing_month": "BillingMonth", "billing_country_code": "BillingCountryCode", "billing_currency_code": "BillingCurrencyCode", "resources": {"billable_cost": 12, "non_billable_cost": 15}, "offers": [{"offer_id": "OfferID", "credits_total": 12, "offer_template": "OfferTemplate", "valid_from": "2019-01-01T12:00:00.000Z", "expires_on": "2019-01-01T12:00:00.000Z", "credits": {"starting_balance": 15, "used": 4, "balance": 7}}], "support": [{"cost": 4, "type": "Type", "overage": 7}], "subscription": {"overage": 7, "subscriptions": [{"subscription_id": "SubscriptionID", "charge_agreement_number": "ChargeAgreementNumber", "type": "Type", "subscription_amount": 18, "start": "2019-01-01T12:00:00.000Z", "end": "2019-01-01T12:00:00.000Z", "credits_total": 12, "terms": [{"start": "2019-01-01T12:00:00.000Z", "end": "2019-01-01T12:00:00.000Z", "credits": {"total": 5, "starting_balance": 15, "used": 4, "balance": 7}}]}]}}`)
				}))
			})
			It(`Invoke GetAccountSummary successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetAccountSummary(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountSummary with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSummaryOptions model with no property values
				getAccountSummaryOptionsModelNew := new(usagereportsv4.GetAccountSummaryOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetAccountSummary(getAccountSummaryOptionsModelNew)
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
			It(`Invoke GetAccountSummary successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSummaryOptions model
				getAccountSummaryOptionsModel := new(usagereportsv4.GetAccountSummaryOptions)
				getAccountSummaryOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountSummaryOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetAccountSummary(getAccountSummaryOptionsModel)
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
	Describe(`GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions) - Operation response error`, func() {
		getAccountUsagePath := "/v4/accounts/testString/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountUsagePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountUsage with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountUsageOptionsModel.Names = core.BoolPtr(true)
				getAccountUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getAccountUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions)`, func() {
		getAccountUsagePath := "/v4/accounts/testString/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetAccountUsage successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountUsageOptionsModel.Names = core.BoolPtr(true)
				getAccountUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getAccountUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetAccountUsageWithContext(ctx, getAccountUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetAccountUsageWithContext(ctx, getAccountUsageOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAccountUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetAccountUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetAccountUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountUsageOptionsModel.Names = core.BoolPtr(true)
				getAccountUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getAccountUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAccountUsage with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountUsageOptionsModel.Names = core.BoolPtr(true)
				getAccountUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getAccountUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountUsageOptions model with no property values
				getAccountUsageOptionsModelNew := new(usagereportsv4.GetAccountUsageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetAccountUsage(getAccountUsageOptionsModelNew)
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
			It(`Invoke GetAccountUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetAccountUsageOptions model
				getAccountUsageOptionsModel := new(usagereportsv4.GetAccountUsageOptions)
				getAccountUsageOptionsModel.AccountID = core.StringPtr("testString")
				getAccountUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getAccountUsageOptionsModel.Names = core.BoolPtr(true)
				getAccountUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getAccountUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetAccountUsage(getAccountUsageOptionsModel)
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
	Describe(`GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions) - Operation response error`, func() {
		getResourceGroupUsagePath := "/v4/accounts/testString/resource_groups/testString/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupUsagePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceGroupUsage with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Names = core.BoolPtr(true)
				getResourceGroupUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions)`, func() {
		getResourceGroupUsagePath := "/v4/accounts/testString/resource_groups/testString/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetResourceGroupUsage successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Names = core.BoolPtr(true)
				getResourceGroupUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetResourceGroupUsageWithContext(ctx, getResourceGroupUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetResourceGroupUsageWithContext(ctx, getResourceGroupUsageOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetResourceGroupUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetResourceGroupUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Names = core.BoolPtr(true)
				getResourceGroupUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceGroupUsage with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Names = core.BoolPtr(true)
				getResourceGroupUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceGroupUsageOptions model with no property values
				getResourceGroupUsageOptionsModelNew := new(usagereportsv4.GetResourceGroupUsageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModelNew)
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
			It(`Invoke GetResourceGroupUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupUsageOptions model
				getResourceGroupUsageOptionsModel := new(usagereportsv4.GetResourceGroupUsageOptions)
				getResourceGroupUsageOptionsModel.AccountID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Names = core.BoolPtr(true)
				getResourceGroupUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceGroupUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptionsModel)
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
	Describe(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) - Operation response error`, func() {
		getResourceUsageAccountPath := "/v4/accounts/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageAccountPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["organization_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceUsageAccount with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageAccountOptions model
				getResourceUsageAccountOptionsModel := new(usagereportsv4.GetResourceUsageAccountOptions)
				getResourceUsageAccountOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageAccountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageAccountOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions)`, func() {
		getResourceUsageAccountPath := "/v4/accounts/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageAccountPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["organization_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageAccount successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageAccountOptions model
				getResourceUsageAccountOptionsModel := new(usagereportsv4.GetResourceUsageAccountOptions)
				getResourceUsageAccountOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageAccountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageAccountOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetResourceUsageAccountWithContext(ctx, getResourceUsageAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetResourceUsageAccountWithContext(ctx, getResourceUsageAccountOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageAccountPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["organization_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageAccount successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetResourceUsageAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageAccountOptions model
				getResourceUsageAccountOptionsModel := new(usagereportsv4.GetResourceUsageAccountOptions)
				getResourceUsageAccountOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageAccountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageAccountOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageAccount with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageAccountOptions model
				getResourceUsageAccountOptionsModel := new(usagereportsv4.GetResourceUsageAccountOptions)
				getResourceUsageAccountOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageAccountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageAccountOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceUsageAccountOptions model with no property values
				getResourceUsageAccountOptionsModelNew := new(usagereportsv4.GetResourceUsageAccountOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModelNew)
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
			It(`Invoke GetResourceUsageAccount successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageAccountOptions model
				getResourceUsageAccountOptionsModel := new(usagereportsv4.GetResourceUsageAccountOptions)
				getResourceUsageAccountOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageAccountOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageAccountOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageAccountOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptionsModel)
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
	Describe(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) - Operation response error`, func() {
		getResourceUsageResourceGroupPath := "/v4/accounts/testString/resource_groups/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceUsageResourceGroup with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				getResourceUsageResourceGroupOptionsModel := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				getResourceUsageResourceGroupOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageResourceGroupOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageResourceGroupOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions)`, func() {
		getResourceUsageResourceGroupPath := "/v4/accounts/testString/resource_groups/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageResourceGroup successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				getResourceUsageResourceGroupOptionsModel := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				getResourceUsageResourceGroupOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageResourceGroupOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageResourceGroupOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetResourceUsageResourceGroupWithContext(ctx, getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetResourceUsageResourceGroupWithContext(ctx, getResourceUsageResourceGroupOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageResourceGroup successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetResourceUsageResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				getResourceUsageResourceGroupOptionsModel := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				getResourceUsageResourceGroupOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageResourceGroupOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageResourceGroupOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageResourceGroup with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				getResourceUsageResourceGroupOptionsModel := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				getResourceUsageResourceGroupOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageResourceGroupOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageResourceGroupOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceUsageResourceGroupOptions model with no property values
				getResourceUsageResourceGroupOptionsModelNew := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModelNew)
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
			It(`Invoke GetResourceUsageResourceGroup successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				getResourceUsageResourceGroupOptionsModel := new(usagereportsv4.GetResourceUsageResourceGroupOptions)
				getResourceUsageResourceGroupOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceGroupID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageResourceGroupOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageResourceGroupOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptionsModel)
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
	Describe(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) - Operation response error`, func() {
		getResourceUsageOrgPath := "/v4/accounts/testString/organizations/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageOrgPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceUsageOrg with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageOrgOptions model
				getResourceUsageOrgOptionsModel := new(usagereportsv4.GetResourceUsageOrgOptions)
				getResourceUsageOrgOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageOrgOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageOrgOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions)`, func() {
		getResourceUsageOrgPath := "/v4/accounts/testString/organizations/testString/resource_instances/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageOrgPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageOrg successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetResourceUsageOrgOptions model
				getResourceUsageOrgOptionsModel := new(usagereportsv4.GetResourceUsageOrgOptions)
				getResourceUsageOrgOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageOrgOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageOrgOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetResourceUsageOrgWithContext(ctx, getResourceUsageOrgOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetResourceUsageOrgWithContext(ctx, getResourceUsageOrgOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getResourceUsageOrgPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					Expect(req.URL.Query()["_limit"]).To(Equal([]string{fmt.Sprint(int64(1))}))
					Expect(req.URL.Query()["_start"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["region"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "count": 5, "first": {"href": "Href"}, "next": {"href": "Href", "offset": "Offset"}, "resources": [{"account_id": "AccountID", "resource_instance_id": "ResourceInstanceID", "resource_instance_name": "ResourceInstanceName", "resource_id": "ResourceID", "resource_name": "ResourceName", "resource_group_id": "ResourceGroupID", "resource_group_name": "ResourceGroupName", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "space_id": "SpaceID", "space_name": "SpaceName", "consumer_id": "ConsumerID", "region": "Region", "pricing_region": "PricingRegion", "pricing_country": "USA", "currency_code": "USD", "billable": true, "plan_id": "PlanID", "plan_name": "PlanName", "month": "2017-08", "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}]}`)
				}))
			})
			It(`Invoke GetResourceUsageOrg successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetResourceUsageOrg(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceUsageOrgOptions model
				getResourceUsageOrgOptionsModel := new(usagereportsv4.GetResourceUsageOrgOptions)
				getResourceUsageOrgOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageOrgOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageOrgOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetResourceUsageOrg with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageOrgOptions model
				getResourceUsageOrgOptionsModel := new(usagereportsv4.GetResourceUsageOrgOptions)
				getResourceUsageOrgOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageOrgOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageOrgOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceUsageOrgOptions model with no property values
				getResourceUsageOrgOptionsModelNew := new(usagereportsv4.GetResourceUsageOrgOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModelNew)
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
			It(`Invoke GetResourceUsageOrg successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetResourceUsageOrgOptions model
				getResourceUsageOrgOptionsModel := new(usagereportsv4.GetResourceUsageOrgOptions)
				getResourceUsageOrgOptionsModel.AccountID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.OrganizationID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Billingmonth = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Names = core.BoolPtr(true)
				getResourceUsageOrgOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Limit = core.Int64Ptr(int64(1))
				getResourceUsageOrgOptionsModel.Start = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.ResourceID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.PlanID = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Region = core.StringPtr("testString")
				getResourceUsageOrgOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptionsModel)
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
	Describe(`GetOrgUsage(getOrgUsageOptions *GetOrgUsageOptions) - Operation response error`, func() {
		getOrgUsagePath := "/v4/accounts/testString/organizations/testString/usage/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOrgUsagePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetOrgUsage with error: Operation response processing error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetOrgUsageOptions model
				getOrgUsageOptionsModel := new(usagereportsv4.GetOrgUsageOptions)
				getOrgUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrgUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrgUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrgUsageOptionsModel.Names = core.BoolPtr(true)
				getOrgUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getOrgUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				usageReportsService.EnableRetries(0, 0)
				result, response, operationErr = usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetOrgUsage(getOrgUsageOptions *GetOrgUsageOptions)`, func() {
		getOrgUsagePath := "/v4/accounts/testString/organizations/testString/usage/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getOrgUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetOrgUsage successfully with retries`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())
				usageReportsService.EnableRetries(0, 0)

				// Construct an instance of the GetOrgUsageOptions model
				getOrgUsageOptionsModel := new(usagereportsv4.GetOrgUsageOptions)
				getOrgUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrgUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrgUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrgUsageOptionsModel.Names = core.BoolPtr(true)
				getOrgUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getOrgUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := usageReportsService.GetOrgUsageWithContext(ctx, getOrgUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				usageReportsService.DisableRetries()
				result, response, operationErr := usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = usageReportsService.GetOrgUsageWithContext(ctx, getOrgUsageOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getOrgUsagePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for _names query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "organization_id": "OrganizationID", "organization_name": "OrganizationName", "pricing_country": "USA", "currency_code": "USD", "month": "2017-08", "resources": [{"resource_id": "ResourceID", "resource_name": "ResourceName", "billable_cost": 12, "billable_rated_cost": 17, "non_billable_cost": 15, "non_billable_rated_cost": 20, "plans": [{"plan_id": "PlanID", "plan_name": "PlanName", "pricing_region": "PricingRegion", "billable": true, "cost": 4, "rated_cost": 9, "usage": [{"metric": "UP-TIME", "metric_name": "UP-TIME", "quantity": 711.11, "rateable_quantity": 700, "cost": 123.45, "rated_cost": 130.0, "price": ["anyValue"], "unit": "HOURS", "unit_name": "HOURS", "non_chargeable": true, "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}], "discounts": [{"ref": "Discount-d27beddb-111b-4bbf-8cb1-b770f531c1a9", "name": "platform-discount", "display_name": "Platform Service Discount", "discount": 5}]}]}`)
				}))
			})
			It(`Invoke GetOrgUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := usageReportsService.GetOrgUsage(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetOrgUsageOptions model
				getOrgUsageOptionsModel := new(usagereportsv4.GetOrgUsageOptions)
				getOrgUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrgUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrgUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrgUsageOptionsModel.Names = core.BoolPtr(true)
				getOrgUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getOrgUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetOrgUsage with error: Operation validation and request error`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetOrgUsageOptions model
				getOrgUsageOptionsModel := new(usagereportsv4.GetOrgUsageOptions)
				getOrgUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrgUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrgUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrgUsageOptionsModel.Names = core.BoolPtr(true)
				getOrgUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getOrgUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := usageReportsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetOrgUsageOptions model with no property values
				getOrgUsageOptionsModelNew := new(usagereportsv4.GetOrgUsageOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = usageReportsService.GetOrgUsage(getOrgUsageOptionsModelNew)
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
			It(`Invoke GetOrgUsage successfully`, func() {
				usageReportsService, serviceErr := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(usageReportsService).ToNot(BeNil())

				// Construct an instance of the GetOrgUsageOptions model
				getOrgUsageOptionsModel := new(usagereportsv4.GetOrgUsageOptions)
				getOrgUsageOptionsModel.AccountID = core.StringPtr("testString")
				getOrgUsageOptionsModel.OrganizationID = core.StringPtr("testString")
				getOrgUsageOptionsModel.Billingmonth = core.StringPtr("testString")
				getOrgUsageOptionsModel.Names = core.BoolPtr(true)
				getOrgUsageOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getOrgUsageOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := usageReportsService.GetOrgUsage(getOrgUsageOptionsModel)
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
			usageReportsService, _ := usagereportsv4.NewUsageReportsV4(&usagereportsv4.UsageReportsV4Options{
				URL:           "http://usagereportsv4modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetAccountSummaryOptions successfully`, func() {
				// Construct an instance of the GetAccountSummaryOptions model
				accountID := "testString"
				billingmonth := "testString"
				getAccountSummaryOptionsModel := usageReportsService.NewGetAccountSummaryOptions(accountID, billingmonth)
				getAccountSummaryOptionsModel.SetAccountID("testString")
				getAccountSummaryOptionsModel.SetBillingmonth("testString")
				getAccountSummaryOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSummaryOptionsModel).ToNot(BeNil())
				Expect(getAccountSummaryOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSummaryOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSummaryOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountUsageOptions successfully`, func() {
				// Construct an instance of the GetAccountUsageOptions model
				accountID := "testString"
				billingmonth := "testString"
				getAccountUsageOptionsModel := usageReportsService.NewGetAccountUsageOptions(accountID, billingmonth)
				getAccountUsageOptionsModel.SetAccountID("testString")
				getAccountUsageOptionsModel.SetBillingmonth("testString")
				getAccountUsageOptionsModel.SetNames(true)
				getAccountUsageOptionsModel.SetAcceptLanguage("testString")
				getAccountUsageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountUsageOptionsModel).ToNot(BeNil())
				Expect(getAccountUsageOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountUsageOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getAccountUsageOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getAccountUsageOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getAccountUsageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetOrgUsageOptions successfully`, func() {
				// Construct an instance of the GetOrgUsageOptions model
				accountID := "testString"
				organizationID := "testString"
				billingmonth := "testString"
				getOrgUsageOptionsModel := usageReportsService.NewGetOrgUsageOptions(accountID, organizationID, billingmonth)
				getOrgUsageOptionsModel.SetAccountID("testString")
				getOrgUsageOptionsModel.SetOrganizationID("testString")
				getOrgUsageOptionsModel.SetBillingmonth("testString")
				getOrgUsageOptionsModel.SetNames(true)
				getOrgUsageOptionsModel.SetAcceptLanguage("testString")
				getOrgUsageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getOrgUsageOptionsModel).ToNot(BeNil())
				Expect(getOrgUsageOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getOrgUsageOptionsModel.OrganizationID).To(Equal(core.StringPtr("testString")))
				Expect(getOrgUsageOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getOrgUsageOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getOrgUsageOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getOrgUsageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceGroupUsageOptions successfully`, func() {
				// Construct an instance of the GetResourceGroupUsageOptions model
				accountID := "testString"
				resourceGroupID := "testString"
				billingmonth := "testString"
				getResourceGroupUsageOptionsModel := usageReportsService.NewGetResourceGroupUsageOptions(accountID, resourceGroupID, billingmonth)
				getResourceGroupUsageOptionsModel.SetAccountID("testString")
				getResourceGroupUsageOptionsModel.SetResourceGroupID("testString")
				getResourceGroupUsageOptionsModel.SetBillingmonth("testString")
				getResourceGroupUsageOptionsModel.SetNames(true)
				getResourceGroupUsageOptionsModel.SetAcceptLanguage("testString")
				getResourceGroupUsageOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceGroupUsageOptionsModel).ToNot(BeNil())
				Expect(getResourceGroupUsageOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceGroupUsageOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceGroupUsageOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getResourceGroupUsageOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getResourceGroupUsageOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getResourceGroupUsageOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceUsageAccountOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageAccountOptions model
				accountID := "testString"
				billingmonth := "testString"
				getResourceUsageAccountOptionsModel := usageReportsService.NewGetResourceUsageAccountOptions(accountID, billingmonth)
				getResourceUsageAccountOptionsModel.SetAccountID("testString")
				getResourceUsageAccountOptionsModel.SetBillingmonth("testString")
				getResourceUsageAccountOptionsModel.SetNames(true)
				getResourceUsageAccountOptionsModel.SetAcceptLanguage("testString")
				getResourceUsageAccountOptionsModel.SetLimit(int64(1))
				getResourceUsageAccountOptionsModel.SetStart("testString")
				getResourceUsageAccountOptionsModel.SetResourceGroupID("testString")
				getResourceUsageAccountOptionsModel.SetOrganizationID("testString")
				getResourceUsageAccountOptionsModel.SetResourceInstanceID("testString")
				getResourceUsageAccountOptionsModel.SetResourceID("testString")
				getResourceUsageAccountOptionsModel.SetPlanID("testString")
				getResourceUsageAccountOptionsModel.SetRegion("testString")
				getResourceUsageAccountOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceUsageAccountOptionsModel).ToNot(BeNil())
				Expect(getResourceUsageAccountOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getResourceUsageAccountOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getResourceUsageAccountOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.OrganizationID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageAccountOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceUsageOrgOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageOrgOptions model
				accountID := "testString"
				organizationID := "testString"
				billingmonth := "testString"
				getResourceUsageOrgOptionsModel := usageReportsService.NewGetResourceUsageOrgOptions(accountID, organizationID, billingmonth)
				getResourceUsageOrgOptionsModel.SetAccountID("testString")
				getResourceUsageOrgOptionsModel.SetOrganizationID("testString")
				getResourceUsageOrgOptionsModel.SetBillingmonth("testString")
				getResourceUsageOrgOptionsModel.SetNames(true)
				getResourceUsageOrgOptionsModel.SetAcceptLanguage("testString")
				getResourceUsageOrgOptionsModel.SetLimit(int64(1))
				getResourceUsageOrgOptionsModel.SetStart("testString")
				getResourceUsageOrgOptionsModel.SetResourceInstanceID("testString")
				getResourceUsageOrgOptionsModel.SetResourceID("testString")
				getResourceUsageOrgOptionsModel.SetPlanID("testString")
				getResourceUsageOrgOptionsModel.SetRegion("testString")
				getResourceUsageOrgOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceUsageOrgOptionsModel).ToNot(BeNil())
				Expect(getResourceUsageOrgOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.OrganizationID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getResourceUsageOrgOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getResourceUsageOrgOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageOrgOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceUsageResourceGroupOptions successfully`, func() {
				// Construct an instance of the GetResourceUsageResourceGroupOptions model
				accountID := "testString"
				resourceGroupID := "testString"
				billingmonth := "testString"
				getResourceUsageResourceGroupOptionsModel := usageReportsService.NewGetResourceUsageResourceGroupOptions(accountID, resourceGroupID, billingmonth)
				getResourceUsageResourceGroupOptionsModel.SetAccountID("testString")
				getResourceUsageResourceGroupOptionsModel.SetResourceGroupID("testString")
				getResourceUsageResourceGroupOptionsModel.SetBillingmonth("testString")
				getResourceUsageResourceGroupOptionsModel.SetNames(true)
				getResourceUsageResourceGroupOptionsModel.SetAcceptLanguage("testString")
				getResourceUsageResourceGroupOptionsModel.SetLimit(int64(1))
				getResourceUsageResourceGroupOptionsModel.SetStart("testString")
				getResourceUsageResourceGroupOptionsModel.SetResourceInstanceID("testString")
				getResourceUsageResourceGroupOptionsModel.SetResourceID("testString")
				getResourceUsageResourceGroupOptionsModel.SetPlanID("testString")
				getResourceUsageResourceGroupOptionsModel.SetRegion("testString")
				getResourceUsageResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceUsageResourceGroupOptionsModel).ToNot(BeNil())
				Expect(getResourceUsageResourceGroupOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.Billingmonth).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.Names).To(Equal(core.BoolPtr(true)))
				Expect(getResourceUsageResourceGroupOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(1))))
				Expect(getResourceUsageResourceGroupOptionsModel.Start).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.Region).To(Equal(core.StringPtr("testString")))
				Expect(getResourceUsageResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
