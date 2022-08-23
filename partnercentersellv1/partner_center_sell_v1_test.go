/**
 * (C) Copyright IBM Corp. 2022.
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

package partnercentersellv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnercentersellv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PartnerCenterSellV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(partnerCenterSellService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(partnerCenterSellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "https://partnercentersellv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(partnerCenterSellService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_URL": "https://partnercentersellv1/api",
				"PARTNER_CENTER_SELL_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				})
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
					URL: "https://testService/api",
				})
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				})
				err := partnerCenterSellService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := partnerCenterSellService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != partnerCenterSellService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(partnerCenterSellService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(partnerCenterSellService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_URL": "https://partnercentersellv1/api",
				"PARTNER_CENTER_SELL_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerCenterSellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PARTNER_CENTER_SELL_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1UsingExternalConfig(&partnercentersellv1.PartnerCenterSellV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(partnerCenterSellService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = partnercentersellv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListProducts(listProductsOptions *ListProductsOptions) - Operation response error`, func() {
		listProductsPath := "/products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProducts with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := new(partnercentersellv1.ListProductsOptions)
				listProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListProducts(listProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListProducts(listProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProducts(listProductsOptions *ListProductsOptions)`, func() {
		listProductsPath := "/products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"errors": [{"message": "Message", "extensions": {"code": "Code", "serviceName": "ServiceName", "exception": {"anyKey": "anyValue"}, "trid": "Trid", "operationName": "OperationName"}}], "products": [{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}]}`)
				}))
			})
			It(`Invoke ListProducts successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := new(partnercentersellv1.ListProductsOptions)
				listProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListProductsWithContext(ctx, listProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListProducts(listProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListProductsWithContext(ctx, listProductsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProductsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"errors": [{"message": "Message", "extensions": {"code": "Code", "serviceName": "ServiceName", "exception": {"anyKey": "anyValue"}, "trid": "Trid", "operationName": "OperationName"}}], "products": [{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}]}`)
				}))
			})
			It(`Invoke ListProducts successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListProducts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := new(partnercentersellv1.ListProductsOptions)
				listProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListProducts(listProductsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProducts with error: Operation request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := new(partnercentersellv1.ListProductsOptions)
				listProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListProducts(listProductsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListProducts successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := new(partnercentersellv1.ListProductsOptions)
				listProductsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListProducts(listProductsOptionsModel)
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
	Describe(`CreateProduct(createProductOptions *CreateProductOptions) - Operation response error`, func() {
		createProductPath := "/products"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateProductOptions model
				createProductOptionsModel := new(partnercentersellv1.CreateProductOptions)
				createProductOptionsModel.ProductName = core.StringPtr("testString")
				createProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				createProductOptionsModel.ProductType = core.StringPtr("SOFTWARE")
				createProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				createProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateProduct(createProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateProduct(createProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateProduct(createProductOptions *CreateProductOptions)`, func() {
		createProductPath := "/products"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke CreateProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CreateProductOptions model
				createProductOptionsModel := new(partnercentersellv1.CreateProductOptions)
				createProductOptionsModel.ProductName = core.StringPtr("testString")
				createProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				createProductOptionsModel.ProductType = core.StringPtr("SOFTWARE")
				createProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				createProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateProductWithContext(ctx, createProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateProduct(createProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateProductWithContext(ctx, createProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke CreateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateProductOptions model
				createProductOptionsModel := new(partnercentersellv1.CreateProductOptions)
				createProductOptionsModel.ProductName = core.StringPtr("testString")
				createProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				createProductOptionsModel.ProductType = core.StringPtr("SOFTWARE")
				createProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				createProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateProduct(createProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateProductOptions model
				createProductOptionsModel := new(partnercentersellv1.CreateProductOptions)
				createProductOptionsModel.ProductName = core.StringPtr("testString")
				createProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				createProductOptionsModel.ProductType = core.StringPtr("SOFTWARE")
				createProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				createProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateProduct(createProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateProductOptions model with no property values
				createProductOptionsModelNew := new(partnercentersellv1.CreateProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateProduct(createProductOptionsModelNew)
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
			It(`Invoke CreateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreateProductOptions model
				createProductOptionsModel := new(partnercentersellv1.CreateProductOptions)
				createProductOptionsModel.ProductName = core.StringPtr("testString")
				createProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				createProductOptionsModel.ProductType = core.StringPtr("SOFTWARE")
				createProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				createProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateProduct(createProductOptionsModel)
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
	Describe(`GetProduct(getProductOptions *GetProductOptions) - Operation response error`, func() {
		getProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProductPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductOptions model
				getProductOptionsModel := new(partnercentersellv1.GetProductOptions)
				getProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetProduct(getProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetProduct(getProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetProduct(getProductOptions *GetProductOptions)`, func() {
		getProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke GetProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetProductOptions model
				getProductOptionsModel := new(partnercentersellv1.GetProductOptions)
				getProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetProductWithContext(ctx, getProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetProduct(getProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetProductWithContext(ctx, getProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getProductPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke GetProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetProductOptions model
				getProductOptionsModel := new(partnercentersellv1.GetProductOptions)
				getProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetProduct(getProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductOptions model
				getProductOptionsModel := new(partnercentersellv1.GetProductOptions)
				getProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetProduct(getProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetProductOptions model with no property values
				getProductOptionsModelNew := new(partnercentersellv1.GetProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetProduct(getProductOptionsModelNew)
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
			It(`Invoke GetProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetProductOptions model
				getProductOptionsModel := new(partnercentersellv1.GetProductOptions)
				getProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetProduct(getProductOptionsModel)
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
	Describe(`UpdateProduct(updateProductOptions *UpdateProductOptions) - Operation response error`, func() {
		updateProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProductPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdateProductOptions model
				updateProductOptionsModel := new(partnercentersellv1.UpdateProductOptions)
				updateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				updateProductOptionsModel.ProductName = core.StringPtr("testString")
				updateProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				updateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateProduct(updateProductOptions *UpdateProductOptions)`, func() {
		updateProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke UpdateProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the UpdateProductOptions model
				updateProductOptionsModel := new(partnercentersellv1.UpdateProductOptions)
				updateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				updateProductOptionsModel.ProductName = core.StringPtr("testString")
				updateProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				updateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateProductWithContext(ctx, updateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateProductWithContext(ctx, updateProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateProductPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke UpdateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateProductOptions model
				updateProductOptionsModel := new(partnercentersellv1.UpdateProductOptions)
				updateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				updateProductOptionsModel.ProductName = core.StringPtr("testString")
				updateProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				updateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdateProductOptions model
				updateProductOptionsModel := new(partnercentersellv1.UpdateProductOptions)
				updateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				updateProductOptionsModel.ProductName = core.StringPtr("testString")
				updateProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				updateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateProductOptions model with no property values
				updateProductOptionsModelNew := new(partnercentersellv1.UpdateProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateProduct(updateProductOptionsModelNew)
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
			It(`Invoke UpdateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdateProductOptions model
				updateProductOptionsModel := new(partnercentersellv1.UpdateProductOptions)
				updateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel.MaterialAgreement = core.BoolPtr(true)
				updateProductOptionsModel.ProductName = core.StringPtr("testString")
				updateProductOptionsModel.TaxAssessment = core.StringPtr("SOFTWARE")
				updateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateProduct(updateProductOptionsModel)
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
	Describe(`DeleteProduct(deleteProductOptions *DeleteProductOptions)`, func() {
		deleteProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteProductPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `false`)
				}))
			})
			It(`Invoke DeleteProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the DeleteProductOptions model
				deleteProductOptionsModel := new(partnercentersellv1.DeleteProductOptions)
				deleteProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.DeleteProductWithContext(ctx, deleteProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.DeleteProduct(deleteProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.DeleteProductWithContext(ctx, deleteProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteProductPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `false`)
				}))
			})
			It(`Invoke DeleteProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.DeleteProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteProductOptions model
				deleteProductOptionsModel := new(partnercentersellv1.DeleteProductOptions)
				deleteProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.DeleteProduct(deleteProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteProductOptions model
				deleteProductOptionsModel := new(partnercentersellv1.DeleteProductOptions)
				deleteProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.DeleteProduct(deleteProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteProductOptions model with no property values
				deleteProductOptionsModelNew := new(partnercentersellv1.DeleteProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.DeleteProduct(deleteProductOptionsModelNew)
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
			It(`Invoke DeleteProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeleteProductOptions model
				deleteProductOptionsModel := new(partnercentersellv1.DeleteProductOptions)
				deleteProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.DeleteProduct(deleteProductOptionsModel)
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
	Describe(`PublishProduct(publishProductOptions *PublishProductOptions) - Operation response error`, func() {
		publishProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/publish"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PublishProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PublishProductOptions model
				publishProductOptionsModel := new(partnercentersellv1.PublishProductOptions)
				publishProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.PublishProduct(publishProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.PublishProduct(publishProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PublishProduct(publishProductOptions *PublishProductOptions)`, func() {
		publishProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/publish"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(publishProductPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke PublishProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the PublishProductOptions model
				publishProductOptionsModel := new(partnercentersellv1.PublishProductOptions)
				publishProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.PublishProductWithContext(ctx, publishProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.PublishProduct(publishProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.PublishProductWithContext(ctx, publishProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(publishProductPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke PublishProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.PublishProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PublishProductOptions model
				publishProductOptionsModel := new(partnercentersellv1.PublishProductOptions)
				publishProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.PublishProduct(publishProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PublishProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PublishProductOptions model
				publishProductOptionsModel := new(partnercentersellv1.PublishProductOptions)
				publishProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.PublishProduct(publishProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PublishProductOptions model with no property values
				publishProductOptionsModelNew := new(partnercentersellv1.PublishProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.PublishProduct(publishProductOptionsModelNew)
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
			It(`Invoke PublishProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the PublishProductOptions model
				publishProductOptionsModel := new(partnercentersellv1.PublishProductOptions)
				publishProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.PublishProduct(publishProductOptionsModel)
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
	Describe(`SuspendProduct(suspendProductOptions *SuspendProductOptions) - Operation response error`, func() {
		suspendProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/suspend"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(suspendProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke SuspendProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SuspendProductOptions model
				suspendProductOptionsModel := new(partnercentersellv1.SuspendProductOptions)
				suspendProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsModel.Reason = core.StringPtr("testString")
				suspendProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`SuspendProduct(suspendProductOptions *SuspendProductOptions)`, func() {
		suspendProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/suspend"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(suspendProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke SuspendProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the SuspendProductOptions model
				suspendProductOptionsModel := new(partnercentersellv1.SuspendProductOptions)
				suspendProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsModel.Reason = core.StringPtr("testString")
				suspendProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.SuspendProductWithContext(ctx, suspendProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.SuspendProductWithContext(ctx, suspendProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(suspendProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke SuspendProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.SuspendProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SuspendProductOptions model
				suspendProductOptionsModel := new(partnercentersellv1.SuspendProductOptions)
				suspendProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsModel.Reason = core.StringPtr("testString")
				suspendProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke SuspendProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SuspendProductOptions model
				suspendProductOptionsModel := new(partnercentersellv1.SuspendProductOptions)
				suspendProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsModel.Reason = core.StringPtr("testString")
				suspendProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SuspendProductOptions model with no property values
				suspendProductOptionsModelNew := new(partnercentersellv1.SuspendProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.SuspendProduct(suspendProductOptionsModelNew)
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
			It(`Invoke SuspendProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SuspendProductOptions model
				suspendProductOptionsModel := new(partnercentersellv1.SuspendProductOptions)
				suspendProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsModel.Reason = core.StringPtr("testString")
				suspendProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.SuspendProduct(suspendProductOptionsModel)
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
	Describe(`DeprecateProduct(deprecateProductOptions *DeprecateProductOptions) - Operation response error`, func() {
		deprecateProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/deprecate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deprecateProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeprecateProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeprecateProductOptions model
				deprecateProductOptionsModel := new(partnercentersellv1.DeprecateProductOptions)
				deprecateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsModel.Reason = core.StringPtr("testString")
				deprecateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeprecateProduct(deprecateProductOptions *DeprecateProductOptions)`, func() {
		deprecateProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/deprecate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deprecateProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke DeprecateProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the DeprecateProductOptions model
				deprecateProductOptionsModel := new(partnercentersellv1.DeprecateProductOptions)
				deprecateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsModel.Reason = core.StringPtr("testString")
				deprecateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.DeprecateProductWithContext(ctx, deprecateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.DeprecateProductWithContext(ctx, deprecateProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deprecateProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke DeprecateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.DeprecateProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeprecateProductOptions model
				deprecateProductOptionsModel := new(partnercentersellv1.DeprecateProductOptions)
				deprecateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsModel.Reason = core.StringPtr("testString")
				deprecateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeprecateProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeprecateProductOptions model
				deprecateProductOptionsModel := new(partnercentersellv1.DeprecateProductOptions)
				deprecateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsModel.Reason = core.StringPtr("testString")
				deprecateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeprecateProductOptions model with no property values
				deprecateProductOptionsModelNew := new(partnercentersellv1.DeprecateProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModelNew)
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
			It(`Invoke DeprecateProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeprecateProductOptions model
				deprecateProductOptionsModel := new(partnercentersellv1.DeprecateProductOptions)
				deprecateProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsModel.Reason = core.StringPtr("testString")
				deprecateProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.DeprecateProduct(deprecateProductOptionsModel)
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
	Describe(`RestoreProduct(restoreProductOptions *RestoreProductOptions) - Operation response error`, func() {
		restoreProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/restore"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreProductPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RestoreProduct with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RestoreProductOptions model
				restoreProductOptionsModel := new(partnercentersellv1.RestoreProductOptions)
				restoreProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsModel.Reason = core.StringPtr("testString")
				restoreProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RestoreProduct(restoreProductOptions *RestoreProductOptions)`, func() {
		restoreProductPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/restore"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(restoreProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke RestoreProduct successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the RestoreProductOptions model
				restoreProductOptionsModel := new(partnercentersellv1.RestoreProductOptions)
				restoreProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsModel.Reason = core.StringPtr("testString")
				restoreProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.RestoreProductWithContext(ctx, restoreProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.RestoreProductWithContext(ctx, restoreProductOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(restoreProductPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke RestoreProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.RestoreProduct(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RestoreProductOptions model
				restoreProductOptionsModel := new(partnercentersellv1.RestoreProductOptions)
				restoreProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsModel.Reason = core.StringPtr("testString")
				restoreProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RestoreProduct with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RestoreProductOptions model
				restoreProductOptionsModel := new(partnercentersellv1.RestoreProductOptions)
				restoreProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsModel.Reason = core.StringPtr("testString")
				restoreProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RestoreProductOptions model with no property values
				restoreProductOptionsModelNew := new(partnercentersellv1.RestoreProductOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.RestoreProduct(restoreProductOptionsModelNew)
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
			It(`Invoke RestoreProduct successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RestoreProductOptions model
				restoreProductOptionsModel := new(partnercentersellv1.RestoreProductOptions)
				restoreProductOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsModel.Reason = core.StringPtr("testString")
				restoreProductOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.RestoreProduct(restoreProductOptionsModel)
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
	Describe(`ListBadges(listBadgesOptions *ListBadgesOptions) - Operation response error`, func() {
		listBadgesPath := "/products/badges"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBadgesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListBadges with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := new(partnercentersellv1.ListBadgesOptions)
				listBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListBadges(listBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListBadges(listBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListBadges(listBadgesOptions *ListBadgesOptions)`, func() {
		listBadgesPath := "/products/badges"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listBadgesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "learnMoreLinks": {"firstParty": "FirstParty", "thirdParty": "ThirdParty"}, "getStartedLink": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke ListBadges successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := new(partnercentersellv1.ListBadgesOptions)
				listBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListBadgesWithContext(ctx, listBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListBadges(listBadgesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListBadgesWithContext(ctx, listBadgesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listBadgesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "learnMoreLinks": {"firstParty": "FirstParty", "thirdParty": "ThirdParty"}, "getStartedLink": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke ListBadges successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListBadges(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := new(partnercentersellv1.ListBadgesOptions)
				listBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListBadges(listBadgesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListBadges with error: Operation request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := new(partnercentersellv1.ListBadgesOptions)
				listBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListBadges(listBadgesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListBadges successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := new(partnercentersellv1.ListBadgesOptions)
				listBadgesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListBadges(listBadgesOptionsModel)
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
	Describe(`GetBadge(getBadgeOptions *GetBadgeOptions) - Operation response error`, func() {
		getBadgePath := "/products/badges/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBadgePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetBadge with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetBadgeOptions model
				getBadgeOptionsModel := new(partnercentersellv1.GetBadgeOptions)
				getBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetBadge(getBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetBadge(getBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetBadge(getBadgeOptions *GetBadgeOptions)`, func() {
		getBadgePath := "/products/badges/9fab83da-98cb-4f18-a7ba-b6f0435c9673"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getBadgePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "learnMoreLinks": {"firstParty": "FirstParty", "thirdParty": "ThirdParty"}, "getStartedLink": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke GetBadge successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetBadgeOptions model
				getBadgeOptionsModel := new(partnercentersellv1.GetBadgeOptions)
				getBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetBadgeWithContext(ctx, getBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetBadge(getBadgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetBadgeWithContext(ctx, getBadgeOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getBadgePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "label": "Label", "description": "Description", "learnMoreLinks": {"firstParty": "FirstParty", "thirdParty": "ThirdParty"}, "getStartedLink": "GetStartedLink", "tag": "Tag"}`)
				}))
			})
			It(`Invoke GetBadge successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetBadge(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetBadgeOptions model
				getBadgeOptionsModel := new(partnercentersellv1.GetBadgeOptions)
				getBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetBadge(getBadgeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetBadge with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetBadgeOptions model
				getBadgeOptionsModel := new(partnercentersellv1.GetBadgeOptions)
				getBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetBadge(getBadgeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetBadgeOptions model with no property values
				getBadgeOptionsModelNew := new(partnercentersellv1.GetBadgeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetBadge(getBadgeOptionsModelNew)
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
			It(`Invoke GetBadge successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetBadgeOptions model
				getBadgeOptionsModel := new(partnercentersellv1.GetBadgeOptions)
				getBadgeOptionsModel.BadgeID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetBadge(getBadgeOptionsModel)
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
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions) - Operation response error`, func() {
		getCatalogPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalog with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(partnercentersellv1.GetCatalogOptions)
				getCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions)`, func() {
		getCatalogPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "catalogId": "CatalogID", "deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "description": "Description", "documentationUrl": "DocumentationURL", "editable": true, "highlights": [{"description": "Description", "description_i18n": {"anyKey": "anyValue"}, "title": "Title", "title_i18n": {"anyKey": "anyValue"}}], "iconUrl": "IconURL", "id": "ID", "keywords": ["Keywords"], "label": "Label", "label_i18n": {"anyKey": "anyValue"}, "longDescription": "LongDescription", "long_description_i18n": {"anyKey": "anyValue"}, "media": [{"caption": "Caption", "caption_i18n": {"anyKey": "anyValue"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "name": "Name", "pcManaged": false, "provider": "Provider", "publishedToAccessList": false, "publishedToIBM": true, "publishedToPublic": false, "short_description_i18n": {"anyKey": "anyValue"}, "tags": ["Tags"], "versions": [{"deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "id": "ID", "kindFormat": "Helm chart", "kindId": "KindID", "kindTarget": "iks", "packageVersion": "PackageVersion", "state": "deprecated", "stateChangeTime": "StateChangeTime", "validatedState": "ValidatedState", "version": "Version", "versionLocator": "VersionLocator", "allowlistedAccounts": ["AllowlistedAccounts"]}]}`)
				}))
			})
			It(`Invoke GetCatalog successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(partnercentersellv1.GetCatalogOptions)
				getCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetCatalogWithContext(ctx, getCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "catalogId": "CatalogID", "deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "description": "Description", "documentationUrl": "DocumentationURL", "editable": true, "highlights": [{"description": "Description", "description_i18n": {"anyKey": "anyValue"}, "title": "Title", "title_i18n": {"anyKey": "anyValue"}}], "iconUrl": "IconURL", "id": "ID", "keywords": ["Keywords"], "label": "Label", "label_i18n": {"anyKey": "anyValue"}, "longDescription": "LongDescription", "long_description_i18n": {"anyKey": "anyValue"}, "media": [{"caption": "Caption", "caption_i18n": {"anyKey": "anyValue"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "name": "Name", "pcManaged": false, "provider": "Provider", "publishedToAccessList": false, "publishedToIBM": true, "publishedToPublic": false, "short_description_i18n": {"anyKey": "anyValue"}, "tags": ["Tags"], "versions": [{"deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "id": "ID", "kindFormat": "Helm chart", "kindId": "KindID", "kindTarget": "iks", "packageVersion": "PackageVersion", "state": "deprecated", "stateChangeTime": "StateChangeTime", "validatedState": "ValidatedState", "version": "Version", "versionLocator": "VersionLocator", "allowlistedAccounts": ["AllowlistedAccounts"]}]}`)
				}))
			})
			It(`Invoke GetCatalog successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(partnercentersellv1.GetCatalogOptions)
				getCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCatalog with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(partnercentersellv1.GetCatalogOptions)
				getCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCatalogOptions model with no property values
				getCatalogOptionsModelNew := new(partnercentersellv1.GetCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetCatalog(getCatalogOptionsModelNew)
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
			It(`Invoke GetCatalog successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(partnercentersellv1.GetCatalogOptions)
				getCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetCatalog(getCatalogOptionsModel)
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
	Describe(`UpdateCatalog(updateCatalogOptions *UpdateCatalogOptions) - Operation response error`, func() {
		updateCatalogPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateCatalog with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogOptions model
				updateCatalogOptionsModel := new(partnercentersellv1.UpdateCatalogOptions)
				updateCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateCatalogOptionsModel.Description = core.StringPtr("testString")
				updateCatalogOptionsModel.IconURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Keywords = []string{"testString"}
				updateCatalogOptionsModel.PricingModel = core.StringPtr("free")
				updateCatalogOptionsModel.Category = core.StringPtr("testString")
				updateCatalogOptionsModel.ProviderType = []string{"ibm_community"}
				updateCatalogOptionsModel.Label = core.StringPtr("testString")
				updateCatalogOptionsModel.Name = core.StringPtr("testString")
				updateCatalogOptionsModel.Provider = core.StringPtr("testString")
				updateCatalogOptionsModel.Tags = []string{"testString"}
				updateCatalogOptionsModel.DocumentationURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Highlights = []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}
				updateCatalogOptionsModel.LongDescription = core.StringPtr("testString")
				updateCatalogOptionsModel.Media = []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}
				updateCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCatalog(updateCatalogOptions *UpdateCatalogOptions)`, func() {
		updateCatalogPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "catalogId": "CatalogID", "deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "description": "Description", "documentationUrl": "DocumentationURL", "editable": true, "highlights": [{"description": "Description", "description_i18n": {"anyKey": "anyValue"}, "title": "Title", "title_i18n": {"anyKey": "anyValue"}}], "iconUrl": "IconURL", "id": "ID", "keywords": ["Keywords"], "label": "Label", "label_i18n": {"anyKey": "anyValue"}, "longDescription": "LongDescription", "long_description_i18n": {"anyKey": "anyValue"}, "media": [{"caption": "Caption", "caption_i18n": {"anyKey": "anyValue"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "name": "Name", "pcManaged": false, "provider": "Provider", "publishedToAccessList": false, "publishedToIBM": true, "publishedToPublic": false, "short_description_i18n": {"anyKey": "anyValue"}, "tags": ["Tags"], "versions": [{"deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "id": "ID", "kindFormat": "Helm chart", "kindId": "KindID", "kindTarget": "iks", "packageVersion": "PackageVersion", "state": "deprecated", "stateChangeTime": "StateChangeTime", "validatedState": "ValidatedState", "version": "Version", "versionLocator": "VersionLocator", "allowlistedAccounts": ["AllowlistedAccounts"]}]}`)
				}))
			})
			It(`Invoke UpdateCatalog successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogOptions model
				updateCatalogOptionsModel := new(partnercentersellv1.UpdateCatalogOptions)
				updateCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateCatalogOptionsModel.Description = core.StringPtr("testString")
				updateCatalogOptionsModel.IconURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Keywords = []string{"testString"}
				updateCatalogOptionsModel.PricingModel = core.StringPtr("free")
				updateCatalogOptionsModel.Category = core.StringPtr("testString")
				updateCatalogOptionsModel.ProviderType = []string{"ibm_community"}
				updateCatalogOptionsModel.Label = core.StringPtr("testString")
				updateCatalogOptionsModel.Name = core.StringPtr("testString")
				updateCatalogOptionsModel.Provider = core.StringPtr("testString")
				updateCatalogOptionsModel.Tags = []string{"testString"}
				updateCatalogOptionsModel.DocumentationURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Highlights = []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}
				updateCatalogOptionsModel.LongDescription = core.StringPtr("testString")
				updateCatalogOptionsModel.Media = []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}
				updateCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateCatalogWithContext(ctx, updateCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateCatalogWithContext(ctx, updateCatalogOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateCatalogPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "catalogId": "CatalogID", "deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "description": "Description", "documentationUrl": "DocumentationURL", "editable": true, "highlights": [{"description": "Description", "description_i18n": {"anyKey": "anyValue"}, "title": "Title", "title_i18n": {"anyKey": "anyValue"}}], "iconUrl": "IconURL", "id": "ID", "keywords": ["Keywords"], "label": "Label", "label_i18n": {"anyKey": "anyValue"}, "longDescription": "LongDescription", "long_description_i18n": {"anyKey": "anyValue"}, "media": [{"caption": "Caption", "caption_i18n": {"anyKey": "anyValue"}, "thumbnail": "Thumbnail", "type": "image", "url": "URL"}], "name": "Name", "pcManaged": false, "provider": "Provider", "publishedToAccessList": false, "publishedToIBM": true, "publishedToPublic": false, "short_description_i18n": {"anyKey": "anyValue"}, "tags": ["Tags"], "versions": [{"deprecatePending": {"deprecateDate": "DeprecateDate", "deprecateState": "DeprecateState", "description": "Description"}, "id": "ID", "kindFormat": "Helm chart", "kindId": "KindID", "kindTarget": "iks", "packageVersion": "PackageVersion", "state": "deprecated", "stateChangeTime": "StateChangeTime", "validatedState": "ValidatedState", "version": "Version", "versionLocator": "VersionLocator", "allowlistedAccounts": ["AllowlistedAccounts"]}]}`)
				}))
			})
			It(`Invoke UpdateCatalog successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogOptions model
				updateCatalogOptionsModel := new(partnercentersellv1.UpdateCatalogOptions)
				updateCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateCatalogOptionsModel.Description = core.StringPtr("testString")
				updateCatalogOptionsModel.IconURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Keywords = []string{"testString"}
				updateCatalogOptionsModel.PricingModel = core.StringPtr("free")
				updateCatalogOptionsModel.Category = core.StringPtr("testString")
				updateCatalogOptionsModel.ProviderType = []string{"ibm_community"}
				updateCatalogOptionsModel.Label = core.StringPtr("testString")
				updateCatalogOptionsModel.Name = core.StringPtr("testString")
				updateCatalogOptionsModel.Provider = core.StringPtr("testString")
				updateCatalogOptionsModel.Tags = []string{"testString"}
				updateCatalogOptionsModel.DocumentationURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Highlights = []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}
				updateCatalogOptionsModel.LongDescription = core.StringPtr("testString")
				updateCatalogOptionsModel.Media = []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}
				updateCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateCatalog with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogOptions model
				updateCatalogOptionsModel := new(partnercentersellv1.UpdateCatalogOptions)
				updateCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateCatalogOptionsModel.Description = core.StringPtr("testString")
				updateCatalogOptionsModel.IconURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Keywords = []string{"testString"}
				updateCatalogOptionsModel.PricingModel = core.StringPtr("free")
				updateCatalogOptionsModel.Category = core.StringPtr("testString")
				updateCatalogOptionsModel.ProviderType = []string{"ibm_community"}
				updateCatalogOptionsModel.Label = core.StringPtr("testString")
				updateCatalogOptionsModel.Name = core.StringPtr("testString")
				updateCatalogOptionsModel.Provider = core.StringPtr("testString")
				updateCatalogOptionsModel.Tags = []string{"testString"}
				updateCatalogOptionsModel.DocumentationURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Highlights = []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}
				updateCatalogOptionsModel.LongDescription = core.StringPtr("testString")
				updateCatalogOptionsModel.Media = []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}
				updateCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateCatalogOptions model with no property values
				updateCatalogOptionsModelNew := new(partnercentersellv1.UpdateCatalogOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModelNew)
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
			It(`Invoke UpdateCatalog successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")

				// Construct an instance of the UpdateCatalogOptions model
				updateCatalogOptionsModel := new(partnercentersellv1.UpdateCatalogOptions)
				updateCatalogOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel.CatalogID = core.StringPtr("testString")
				updateCatalogOptionsModel.Description = core.StringPtr("testString")
				updateCatalogOptionsModel.IconURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Keywords = []string{"testString"}
				updateCatalogOptionsModel.PricingModel = core.StringPtr("free")
				updateCatalogOptionsModel.Category = core.StringPtr("testString")
				updateCatalogOptionsModel.ProviderType = []string{"ibm_community"}
				updateCatalogOptionsModel.Label = core.StringPtr("testString")
				updateCatalogOptionsModel.Name = core.StringPtr("testString")
				updateCatalogOptionsModel.Provider = core.StringPtr("testString")
				updateCatalogOptionsModel.Tags = []string{"testString"}
				updateCatalogOptionsModel.DocumentationURL = core.StringPtr("testString")
				updateCatalogOptionsModel.Highlights = []partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}
				updateCatalogOptionsModel.LongDescription = core.StringPtr("testString")
				updateCatalogOptionsModel.Media = []partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}
				updateCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateCatalog(updateCatalogOptionsModel)
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
	Describe(`RequestCatalogApproval(requestCatalogApprovalOptions *RequestCatalogApprovalOptions) - Operation response error`, func() {
		requestCatalogApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog/approvals"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestCatalogApprovalPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RequestCatalogApproval with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestCatalogApprovalOptions model
				requestCatalogApprovalOptionsModel := new(partnercentersellv1.RequestCatalogApprovalOptions)
				requestCatalogApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RequestCatalogApproval(requestCatalogApprovalOptions *RequestCatalogApprovalOptions)`, func() {
		requestCatalogApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/catalog/approvals"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestCatalogApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestCatalogApproval successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the RequestCatalogApprovalOptions model
				requestCatalogApprovalOptionsModel := new(partnercentersellv1.RequestCatalogApprovalOptions)
				requestCatalogApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.RequestCatalogApprovalWithContext(ctx, requestCatalogApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.RequestCatalogApprovalWithContext(ctx, requestCatalogApprovalOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(requestCatalogApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestCatalogApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.RequestCatalogApproval(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RequestCatalogApprovalOptions model
				requestCatalogApprovalOptionsModel := new(partnercentersellv1.RequestCatalogApprovalOptions)
				requestCatalogApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RequestCatalogApproval with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestCatalogApprovalOptions model
				requestCatalogApprovalOptionsModel := new(partnercentersellv1.RequestCatalogApprovalOptions)
				requestCatalogApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RequestCatalogApprovalOptions model with no property values
				requestCatalogApprovalOptionsModelNew := new(partnercentersellv1.RequestCatalogApprovalOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModelNew)
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
			It(`Invoke RequestCatalogApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestCatalogApprovalOptions model
				requestCatalogApprovalOptionsModel := new(partnercentersellv1.RequestCatalogApprovalOptions)
				requestCatalogApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.RequestCatalogApproval(requestCatalogApprovalOptionsModel)
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
	Describe(`ListPlans(listPlansOptions *ListPlansOptions) - Operation response error`, func() {
		listPlansPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPlansPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPlans with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListPlansOptions model
				listPlansOptionsModel := new(partnercentersellv1.ListPlansOptions)
				listPlansOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListPlans(listPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListPlans(listPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPlans(listPlansOptions *ListPlansOptions)`, func() {
		listPlansPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPlansPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke ListPlans successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListPlansOptions model
				listPlansOptionsModel := new(partnercentersellv1.ListPlansOptions)
				listPlansOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListPlansWithContext(ctx, listPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListPlans(listPlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListPlansWithContext(ctx, listPlansOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPlansPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke ListPlans successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListPlans(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPlansOptions model
				listPlansOptionsModel := new(partnercentersellv1.ListPlansOptions)
				listPlansOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListPlans(listPlansOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPlans with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListPlansOptions model
				listPlansOptionsModel := new(partnercentersellv1.ListPlansOptions)
				listPlansOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListPlans(listPlansOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPlansOptions model with no property values
				listPlansOptionsModelNew := new(partnercentersellv1.ListPlansOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.ListPlans(listPlansOptionsModelNew)
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
			It(`Invoke ListPlans successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListPlansOptions model
				listPlansOptionsModel := new(partnercentersellv1.ListPlansOptions)
				listPlansOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListPlans(listPlansOptionsModel)
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
	Describe(`CreatePlan(createPlanOptions *CreatePlanOptions) - Operation response error`, func() {
		createPlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPlanPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreatePlanOptions model
				createPlanOptionsModel := new(partnercentersellv1.CreatePlanOptions)
				createPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsModel.Description = core.StringPtr("testString")
				createPlanOptionsModel.Label = core.StringPtr("testString")
				createPlanOptionsModel.Type = core.StringPtr("byol")
				createPlanOptionsModel.URL = core.StringPtr("testString")
				createPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreatePlan(createPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreatePlan(createPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePlan(createPlanOptions *CreatePlanOptions)`, func() {
		createPlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPlanPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke CreatePlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the CreatePlanOptions model
				createPlanOptionsModel := new(partnercentersellv1.CreatePlanOptions)
				createPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsModel.Description = core.StringPtr("testString")
				createPlanOptionsModel.Label = core.StringPtr("testString")
				createPlanOptionsModel.Type = core.StringPtr("byol")
				createPlanOptionsModel.URL = core.StringPtr("testString")
				createPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreatePlanWithContext(ctx, createPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreatePlan(createPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreatePlanWithContext(ctx, createPlanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPlanPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke CreatePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreatePlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreatePlanOptions model
				createPlanOptionsModel := new(partnercentersellv1.CreatePlanOptions)
				createPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsModel.Description = core.StringPtr("testString")
				createPlanOptionsModel.Label = core.StringPtr("testString")
				createPlanOptionsModel.Type = core.StringPtr("byol")
				createPlanOptionsModel.URL = core.StringPtr("testString")
				createPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreatePlan(createPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreatePlanOptions model
				createPlanOptionsModel := new(partnercentersellv1.CreatePlanOptions)
				createPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsModel.Description = core.StringPtr("testString")
				createPlanOptionsModel.Label = core.StringPtr("testString")
				createPlanOptionsModel.Type = core.StringPtr("byol")
				createPlanOptionsModel.URL = core.StringPtr("testString")
				createPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreatePlan(createPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePlanOptions model with no property values
				createPlanOptionsModelNew := new(partnercentersellv1.CreatePlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreatePlan(createPlanOptionsModelNew)
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
			It(`Invoke CreatePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the CreatePlanOptions model
				createPlanOptionsModel := new(partnercentersellv1.CreatePlanOptions)
				createPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsModel.Description = core.StringPtr("testString")
				createPlanOptionsModel.Label = core.StringPtr("testString")
				createPlanOptionsModel.Type = core.StringPtr("byol")
				createPlanOptionsModel.URL = core.StringPtr("testString")
				createPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreatePlan(createPlanOptionsModel)
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
	Describe(`GetPlan(getPlanOptions *GetPlanOptions) - Operation response error`, func() {
		getPlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPlanPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetPlanOptions model
				getPlanOptionsModel := new(partnercentersellv1.GetPlanOptions)
				getPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getPlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				getPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetPlan(getPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetPlan(getPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPlan(getPlanOptions *GetPlanOptions)`, func() {
		getPlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPlanPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}`)
				}))
			})
			It(`Invoke GetPlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetPlanOptions model
				getPlanOptionsModel := new(partnercentersellv1.GetPlanOptions)
				getPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getPlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				getPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetPlanWithContext(ctx, getPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetPlan(getPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetPlanWithContext(ctx, getPlanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPlanPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}`)
				}))
			})
			It(`Invoke GetPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetPlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPlanOptions model
				getPlanOptionsModel := new(partnercentersellv1.GetPlanOptions)
				getPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getPlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				getPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetPlan(getPlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetPlanOptions model
				getPlanOptionsModel := new(partnercentersellv1.GetPlanOptions)
				getPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getPlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				getPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetPlan(getPlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPlanOptions model with no property values
				getPlanOptionsModelNew := new(partnercentersellv1.GetPlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetPlan(getPlanOptionsModelNew)
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
			It(`Invoke GetPlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetPlanOptions model
				getPlanOptionsModel := new(partnercentersellv1.GetPlanOptions)
				getPlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getPlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				getPlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetPlan(getPlanOptionsModel)
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
	Describe(`UpdatePlan(updatePlanOptions *UpdatePlanOptions) - Operation response error`, func() {
		updatePlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePlanPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdatePlanOptions model
				updatePlanOptionsModel := new(partnercentersellv1.UpdatePlanOptions)
				updatePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updatePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				updatePlanOptionsModel.Description = core.StringPtr("testString")
				updatePlanOptionsModel.Label = core.StringPtr("testString")
				updatePlanOptionsModel.Type = core.StringPtr("byol")
				updatePlanOptionsModel.URL = core.StringPtr("testString")
				updatePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePlan(updatePlanOptions *UpdatePlanOptions)`, func() {
		updatePlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePlanPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke UpdatePlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the UpdatePlanOptions model
				updatePlanOptionsModel := new(partnercentersellv1.UpdatePlanOptions)
				updatePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updatePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				updatePlanOptionsModel.Description = core.StringPtr("testString")
				updatePlanOptionsModel.Label = core.StringPtr("testString")
				updatePlanOptionsModel.Type = core.StringPtr("byol")
				updatePlanOptionsModel.URL = core.StringPtr("testString")
				updatePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdatePlanWithContext(ctx, updatePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdatePlanWithContext(ctx, updatePlanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updatePlanPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke UpdatePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdatePlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePlanOptions model
				updatePlanOptionsModel := new(partnercentersellv1.UpdatePlanOptions)
				updatePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updatePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				updatePlanOptionsModel.Description = core.StringPtr("testString")
				updatePlanOptionsModel.Label = core.StringPtr("testString")
				updatePlanOptionsModel.Type = core.StringPtr("byol")
				updatePlanOptionsModel.URL = core.StringPtr("testString")
				updatePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdatePlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdatePlanOptions model
				updatePlanOptionsModel := new(partnercentersellv1.UpdatePlanOptions)
				updatePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updatePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				updatePlanOptionsModel.Description = core.StringPtr("testString")
				updatePlanOptionsModel.Label = core.StringPtr("testString")
				updatePlanOptionsModel.Type = core.StringPtr("byol")
				updatePlanOptionsModel.URL = core.StringPtr("testString")
				updatePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePlanOptions model with no property values
				updatePlanOptionsModelNew := new(partnercentersellv1.UpdatePlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdatePlan(updatePlanOptionsModelNew)
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
			It(`Invoke UpdatePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the UpdatePlanOptions model
				updatePlanOptionsModel := new(partnercentersellv1.UpdatePlanOptions)
				updatePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updatePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				updatePlanOptionsModel.Description = core.StringPtr("testString")
				updatePlanOptionsModel.Label = core.StringPtr("testString")
				updatePlanOptionsModel.Type = core.StringPtr("byol")
				updatePlanOptionsModel.URL = core.StringPtr("testString")
				updatePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdatePlan(updatePlanOptionsModel)
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
	Describe(`DeletePlan(deletePlanOptions *DeletePlanOptions) - Operation response error`, func() {
		deletePlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePlanPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeletePlan with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeletePlanOptions model
				deletePlanOptionsModel := new(partnercentersellv1.DeletePlanOptions)
				deletePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deletePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				deletePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeletePlan(deletePlanOptions *DeletePlanOptions)`, func() {
		deletePlanPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/plans/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePlanPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke DeletePlan successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the DeletePlanOptions model
				deletePlanOptionsModel := new(partnercentersellv1.DeletePlanOptions)
				deletePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deletePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				deletePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.DeletePlanWithContext(ctx, deletePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.DeletePlanWithContext(ctx, deletePlanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deletePlanPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"plans": [{"description": "Description", "id": "ID", "label": "Label", "type": "byol", "url": "URL"}]}`)
				}))
			})
			It(`Invoke DeletePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.DeletePlan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeletePlanOptions model
				deletePlanOptionsModel := new(partnercentersellv1.DeletePlanOptions)
				deletePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deletePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				deletePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeletePlan with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeletePlanOptions model
				deletePlanOptionsModel := new(partnercentersellv1.DeletePlanOptions)
				deletePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deletePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				deletePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeletePlanOptions model with no property values
				deletePlanOptionsModelNew := new(partnercentersellv1.DeletePlanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.DeletePlan(deletePlanOptionsModelNew)
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
			It(`Invoke DeletePlan successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the DeletePlanOptions model
				deletePlanOptionsModel := new(partnercentersellv1.DeletePlanOptions)
				deletePlanOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deletePlanOptionsModel.PricingPlanID = core.StringPtr("testString")
				deletePlanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.DeletePlan(deletePlanOptionsModel)
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
	Describe(`GetSupport(getSupportOptions *GetSupportOptions) - Operation response error`, func() {
		getSupportPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSupport with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportOptions model
				getSupportOptionsModel := new(partnercentersellv1.GetSupportOptions)
				getSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetSupport(getSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetSupport(getSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSupport(getSupportOptions *GetSupportOptions)`, func() {
		getSupportPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}`)
				}))
			})
			It(`Invoke GetSupport successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetSupportOptions model
				getSupportOptionsModel := new(partnercentersellv1.GetSupportOptions)
				getSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetSupportWithContext(ctx, getSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetSupport(getSupportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetSupportWithContext(ctx, getSupportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSupportPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}`)
				}))
			})
			It(`Invoke GetSupport successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetSupport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSupportOptions model
				getSupportOptionsModel := new(partnercentersellv1.GetSupportOptions)
				getSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetSupport(getSupportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSupport with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportOptions model
				getSupportOptionsModel := new(partnercentersellv1.GetSupportOptions)
				getSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetSupport(getSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSupportOptions model with no property values
				getSupportOptionsModelNew := new(partnercentersellv1.GetSupportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetSupport(getSupportOptionsModelNew)
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
			It(`Invoke GetSupport successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportOptions model
				getSupportOptionsModel := new(partnercentersellv1.GetSupportOptions)
				getSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetSupport(getSupportOptionsModel)
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
	Describe(`UpdateSupport(updateSupportOptions *UpdateSupportOptions) - Operation response error`, func() {
		updateSupportPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSupport with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the UpdateSupportOptions model
				updateSupportOptionsModel := new(partnercentersellv1.UpdateSupportOptions)
				updateSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel.EscalationContacts = []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}
				updateSupportOptionsModel.Locations = []string{"US"}
				updateSupportOptionsModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				updateSupportOptionsModel.SupportEscalation = supportEscalationModel
				updateSupportOptionsModel.SupportType = core.StringPtr("third-party")
				updateSupportOptionsModel.URL = core.StringPtr("https://my-company.com/support")
				updateSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSupport(updateSupportOptions *UpdateSupportOptions)`, func() {
		updateSupportPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}`)
				}))
			})
			It(`Invoke UpdateSupport successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the UpdateSupportOptions model
				updateSupportOptionsModel := new(partnercentersellv1.UpdateSupportOptions)
				updateSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel.EscalationContacts = []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}
				updateSupportOptionsModel.Locations = []string{"US"}
				updateSupportOptionsModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				updateSupportOptionsModel.SupportEscalation = supportEscalationModel
				updateSupportOptionsModel.SupportType = core.StringPtr("third-party")
				updateSupportOptionsModel.URL = core.StringPtr("https://my-company.com/support")
				updateSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateSupportWithContext(ctx, updateSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateSupportWithContext(ctx, updateSupportOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}`)
				}))
			})
			It(`Invoke UpdateSupport successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateSupport(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the UpdateSupportOptions model
				updateSupportOptionsModel := new(partnercentersellv1.UpdateSupportOptions)
				updateSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel.EscalationContacts = []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}
				updateSupportOptionsModel.Locations = []string{"US"}
				updateSupportOptionsModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				updateSupportOptionsModel.SupportEscalation = supportEscalationModel
				updateSupportOptionsModel.SupportType = core.StringPtr("third-party")
				updateSupportOptionsModel.URL = core.StringPtr("https://my-company.com/support")
				updateSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSupport with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the UpdateSupportOptions model
				updateSupportOptionsModel := new(partnercentersellv1.UpdateSupportOptions)
				updateSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel.EscalationContacts = []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}
				updateSupportOptionsModel.Locations = []string{"US"}
				updateSupportOptionsModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				updateSupportOptionsModel.SupportEscalation = supportEscalationModel
				updateSupportOptionsModel.SupportType = core.StringPtr("third-party")
				updateSupportOptionsModel.URL = core.StringPtr("https://my-company.com/support")
				updateSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSupportOptions model with no property values
				updateSupportOptionsModelNew := new(partnercentersellv1.UpdateSupportOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateSupport(updateSupportOptionsModelNew)
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
			It(`Invoke UpdateSupport successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the UpdateSupportOptions model
				updateSupportOptionsModel := new(partnercentersellv1.UpdateSupportOptions)
				updateSupportOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel.EscalationContacts = []partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}
				updateSupportOptionsModel.Locations = []string{"US"}
				updateSupportOptionsModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				updateSupportOptionsModel.SupportEscalation = supportEscalationModel
				updateSupportOptionsModel.SupportType = core.StringPtr("third-party")
				updateSupportOptionsModel.URL = core.StringPtr("https://my-company.com/support")
				updateSupportOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateSupport(updateSupportOptionsModel)
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
	Describe(`ListSupportChangeRequests(listSupportChangeRequestsOptions *ListSupportChangeRequestsOptions) - Operation response error`, func() {
		listSupportChangeRequestsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSupportChangeRequests with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestsOptions model
				listSupportChangeRequestsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				listSupportChangeRequestsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSupportChangeRequests(listSupportChangeRequestsOptions *ListSupportChangeRequestsOptions)`, func() {
		listSupportChangeRequestsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"changes": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke ListSupportChangeRequests successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListSupportChangeRequestsOptions model
				listSupportChangeRequestsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				listSupportChangeRequestsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListSupportChangeRequestsWithContext(ctx, listSupportChangeRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListSupportChangeRequestsWithContext(ctx, listSupportChangeRequestsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"changes": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke ListSupportChangeRequests successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequests(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSupportChangeRequestsOptions model
				listSupportChangeRequestsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				listSupportChangeRequestsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSupportChangeRequests with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestsOptions model
				listSupportChangeRequestsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				listSupportChangeRequestsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSupportChangeRequestsOptions model with no property values
				listSupportChangeRequestsOptionsModelNew := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModelNew)
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
			It(`Invoke ListSupportChangeRequests successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestsOptions model
				listSupportChangeRequestsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestsOptions)
				listSupportChangeRequestsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequests(listSupportChangeRequestsOptionsModel)
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
	Describe(`CreateSupportChangeRequest(createSupportChangeRequestOptions *CreateSupportChangeRequestOptions) - Operation response error`, func() {
		createSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSupportChangeRequest with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the CreateSupportChangeRequestOptions model
				createSupportChangeRequestOptionsModel := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				createSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createSupportChangeRequestOptionsModel.Change = supportModel
				createSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSupportChangeRequest(createSupportChangeRequestOptions *CreateSupportChangeRequestOptions)`, func() {
		createSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke CreateSupportChangeRequest successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the CreateSupportChangeRequestOptions model
				createSupportChangeRequestOptionsModel := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				createSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createSupportChangeRequestOptionsModel.Change = supportModel
				createSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.CreateSupportChangeRequestWithContext(ctx, createSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.CreateSupportChangeRequestWithContext(ctx, createSupportChangeRequestOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke CreateSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.CreateSupportChangeRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the CreateSupportChangeRequestOptions model
				createSupportChangeRequestOptionsModel := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				createSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createSupportChangeRequestOptionsModel.Change = supportModel
				createSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSupportChangeRequest with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the CreateSupportChangeRequestOptions model
				createSupportChangeRequestOptionsModel := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				createSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createSupportChangeRequestOptionsModel.Change = supportModel
				createSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSupportChangeRequestOptions model with no property values
				createSupportChangeRequestOptionsModelNew := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModelNew)
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
			It(`Invoke CreateSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the CreateSupportChangeRequestOptions model
				createSupportChangeRequestOptionsModel := new(partnercentersellv1.CreateSupportChangeRequestOptions)
				createSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createSupportChangeRequestOptionsModel.Change = supportModel
				createSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.CreateSupportChangeRequest(createSupportChangeRequestOptionsModel)
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
	Describe(`GetSupportChangeRequest(getSupportChangeRequestOptions *GetSupportChangeRequestOptions) - Operation response error`, func() {
		getSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportChangeRequestPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSupportChangeRequest with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportChangeRequestOptions model
				getSupportChangeRequestOptionsModel := new(partnercentersellv1.GetSupportChangeRequestOptions)
				getSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				getSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSupportChangeRequest(getSupportChangeRequestOptions *GetSupportChangeRequestOptions)`, func() {
		getSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSupportChangeRequestPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}`)
				}))
			})
			It(`Invoke GetSupportChangeRequest successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the GetSupportChangeRequestOptions model
				getSupportChangeRequestOptionsModel := new(partnercentersellv1.GetSupportChangeRequestOptions)
				getSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				getSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.GetSupportChangeRequestWithContext(ctx, getSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.GetSupportChangeRequestWithContext(ctx, getSupportChangeRequestOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSupportChangeRequestPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}`)
				}))
			})
			It(`Invoke GetSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.GetSupportChangeRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSupportChangeRequestOptions model
				getSupportChangeRequestOptionsModel := new(partnercentersellv1.GetSupportChangeRequestOptions)
				getSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				getSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSupportChangeRequest with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportChangeRequestOptions model
				getSupportChangeRequestOptionsModel := new(partnercentersellv1.GetSupportChangeRequestOptions)
				getSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				getSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSupportChangeRequestOptions model with no property values
				getSupportChangeRequestOptionsModelNew := new(partnercentersellv1.GetSupportChangeRequestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModelNew)
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
			It(`Invoke GetSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the GetSupportChangeRequestOptions model
				getSupportChangeRequestOptionsModel := new(partnercentersellv1.GetSupportChangeRequestOptions)
				getSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				getSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.GetSupportChangeRequest(getSupportChangeRequestOptionsModel)
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
	Describe(`UpdateSupportChangeRequest(updateSupportChangeRequestOptions *UpdateSupportChangeRequestOptions) - Operation response error`, func() {
		updateSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportChangeRequestPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSupportChangeRequest with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				updateSupportChangeRequestOptionsModel := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				updateSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				updateSupportChangeRequestOptionsModel.Change = supportModel
				updateSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSupportChangeRequest(updateSupportChangeRequestOptions *UpdateSupportChangeRequestOptions)`, func() {
		updateSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportChangeRequestPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke UpdateSupportChangeRequest successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				updateSupportChangeRequestOptionsModel := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				updateSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				updateSupportChangeRequestOptionsModel.Change = supportModel
				updateSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.UpdateSupportChangeRequestWithContext(ctx, updateSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.UpdateSupportChangeRequestWithContext(ctx, updateSupportChangeRequestOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSupportChangeRequestPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke UpdateSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.UpdateSupportChangeRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				updateSupportChangeRequestOptionsModel := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				updateSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				updateSupportChangeRequestOptionsModel.Change = supportModel
				updateSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSupportChangeRequest with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				updateSupportChangeRequestOptionsModel := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				updateSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				updateSupportChangeRequestOptionsModel.Change = supportModel
				updateSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateSupportChangeRequestOptions model with no property values
				updateSupportChangeRequestOptionsModelNew := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModelNew)
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
			It(`Invoke UpdateSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				updateSupportChangeRequestOptionsModel := new(partnercentersellv1.UpdateSupportChangeRequestOptions)
				updateSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				updateSupportChangeRequestOptionsModel.Change = supportModel
				updateSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.UpdateSupportChangeRequest(updateSupportChangeRequestOptionsModel)
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
	Describe(`ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptions *ListSupportChangeRequestReviewsOptions) - Operation response error`, func() {
		listSupportChangeRequestReviewsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/reviews"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestReviewsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSupportChangeRequestReviews with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				listSupportChangeRequestReviewsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				listSupportChangeRequestReviewsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestReviewsOptionsModel.ChangeRequestID = core.StringPtr("testString")
				listSupportChangeRequestReviewsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptions *ListSupportChangeRequestReviewsOptions)`, func() {
		listSupportChangeRequestReviewsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/reviews"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestReviewsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke ListSupportChangeRequestReviews successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				listSupportChangeRequestReviewsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				listSupportChangeRequestReviewsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestReviewsOptionsModel.ChangeRequestID = core.StringPtr("testString")
				listSupportChangeRequestReviewsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListSupportChangeRequestReviewsWithContext(ctx, listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListSupportChangeRequestReviewsWithContext(ctx, listSupportChangeRequestReviewsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listSupportChangeRequestReviewsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke ListSupportChangeRequestReviews successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequestReviews(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				listSupportChangeRequestReviewsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				listSupportChangeRequestReviewsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestReviewsOptionsModel.ChangeRequestID = core.StringPtr("testString")
				listSupportChangeRequestReviewsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSupportChangeRequestReviews with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				listSupportChangeRequestReviewsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				listSupportChangeRequestReviewsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestReviewsOptionsModel.ChangeRequestID = core.StringPtr("testString")
				listSupportChangeRequestReviewsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListSupportChangeRequestReviewsOptions model with no property values
				listSupportChangeRequestReviewsOptionsModelNew := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModelNew)
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
			It(`Invoke ListSupportChangeRequestReviews successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				listSupportChangeRequestReviewsOptionsModel := new(partnercentersellv1.ListSupportChangeRequestReviewsOptions)
				listSupportChangeRequestReviewsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestReviewsOptionsModel.ChangeRequestID = core.StringPtr("testString")
				listSupportChangeRequestReviewsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListSupportChangeRequestReviews(listSupportChangeRequestReviewsOptionsModel)
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
	Describe(`RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptions *RequestSupportChangeRequestReviewOptions) - Operation response error`, func() {
		requestSupportChangeRequestReviewPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/reviews"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportChangeRequestReviewPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RequestSupportChangeRequestReview with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				requestSupportChangeRequestReviewOptionsModel := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				requestSupportChangeRequestReviewOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportChangeRequestReviewOptionsModel.ChangeRequestID = core.StringPtr("testString")
				requestSupportChangeRequestReviewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptions *RequestSupportChangeRequestReviewOptions)`, func() {
		requestSupportChangeRequestReviewPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/reviews"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportChangeRequestReviewPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestSupportChangeRequestReview successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				requestSupportChangeRequestReviewOptionsModel := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				requestSupportChangeRequestReviewOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportChangeRequestReviewOptionsModel.ChangeRequestID = core.StringPtr("testString")
				requestSupportChangeRequestReviewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.RequestSupportChangeRequestReviewWithContext(ctx, requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.RequestSupportChangeRequestReviewWithContext(ctx, requestSupportChangeRequestReviewOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportChangeRequestReviewPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestSupportChangeRequestReview successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.RequestSupportChangeRequestReview(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				requestSupportChangeRequestReviewOptionsModel := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				requestSupportChangeRequestReviewOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportChangeRequestReviewOptionsModel.ChangeRequestID = core.StringPtr("testString")
				requestSupportChangeRequestReviewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RequestSupportChangeRequestReview with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				requestSupportChangeRequestReviewOptionsModel := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				requestSupportChangeRequestReviewOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportChangeRequestReviewOptionsModel.ChangeRequestID = core.StringPtr("testString")
				requestSupportChangeRequestReviewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RequestSupportChangeRequestReviewOptions model with no property values
				requestSupportChangeRequestReviewOptionsModelNew := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModelNew)
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
			It(`Invoke RequestSupportChangeRequestReview successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				requestSupportChangeRequestReviewOptionsModel := new(partnercentersellv1.RequestSupportChangeRequestReviewOptions)
				requestSupportChangeRequestReviewOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportChangeRequestReviewOptionsModel.ChangeRequestID = core.StringPtr("testString")
				requestSupportChangeRequestReviewOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.RequestSupportChangeRequestReview(requestSupportChangeRequestReviewOptionsModel)
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
	Describe(`MergeSupportChangeRequest(mergeSupportChangeRequestOptions *MergeSupportChangeRequestOptions) - Operation response error`, func() {
		mergeSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/merge"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(mergeSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke MergeSupportChangeRequest with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the MergeSupportChangeRequestOptions model
				mergeSupportChangeRequestOptionsModel := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				mergeSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				mergeSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				mergeSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`MergeSupportChangeRequest(mergeSupportChangeRequestOptions *MergeSupportChangeRequestOptions)`, func() {
		mergeSupportChangeRequestPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/changes/testString/merge"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(mergeSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke MergeSupportChangeRequest successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the MergeSupportChangeRequestOptions model
				mergeSupportChangeRequestOptionsModel := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				mergeSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				mergeSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				mergeSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.MergeSupportChangeRequestWithContext(ctx, mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.MergeSupportChangeRequestWithContext(ctx, mergeSupportChangeRequestOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(mergeSupportChangeRequestPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"accountId": "AccountID", "createdAt": "CreatedAt", "id": "ID", "materialAgreement": false, "productType": "SOFTWARE", "productName": "ProductName", "publishedAt": "PublishedAt", "taxAssessment": "SOFTWARE", "updatedAt": "UpdatedAt", "changeRequests": [{"id": "ID", "createdAt": "CreatedAt", "initiator": "Initiator", "merged": "Merged", "change": {"locations": ["US"], "process": "Process", "process_i18n": {"anyKey": "anyValue"}, "support_details": [{"availability": {"always_available": false, "times": [{"day": 1, "end_time": "19:30", "start_time": "10:30"}], "timezone": "America/Los_Angeles"}, "contact": "Contact", "response_wait_time": {"type": "hour", "value": 2}, "type": "email"}], "support_escalation": {"contact": "Contact", "escalation_wait_time": {"type": "hour", "value": 2}, "response_wait_time": {"type": "hour", "value": 2}}, "support_type": "third-party", "url": "https://my-company.com/support"}}]}`)
				}))
			})
			It(`Invoke MergeSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.MergeSupportChangeRequest(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MergeSupportChangeRequestOptions model
				mergeSupportChangeRequestOptionsModel := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				mergeSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				mergeSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				mergeSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke MergeSupportChangeRequest with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the MergeSupportChangeRequestOptions model
				mergeSupportChangeRequestOptionsModel := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				mergeSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				mergeSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				mergeSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the MergeSupportChangeRequestOptions model with no property values
				mergeSupportChangeRequestOptionsModelNew := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModelNew)
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
			It(`Invoke MergeSupportChangeRequest successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the MergeSupportChangeRequestOptions model
				mergeSupportChangeRequestOptionsModel := new(partnercentersellv1.MergeSupportChangeRequestOptions)
				mergeSupportChangeRequestOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				mergeSupportChangeRequestOptionsModel.ChangeRequestID = core.StringPtr("testString")
				mergeSupportChangeRequestOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.MergeSupportChangeRequest(mergeSupportChangeRequestOptionsModel)
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
	Describe(`RequestSupportApproval(requestSupportApprovalOptions *RequestSupportApprovalOptions) - Operation response error`, func() {
		requestSupportApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/approvals"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportApprovalPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RequestSupportApproval with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportApprovalOptions model
				requestSupportApprovalOptionsModel := new(partnercentersellv1.RequestSupportApprovalOptions)
				requestSupportApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RequestSupportApproval(requestSupportApprovalOptions *RequestSupportApprovalOptions)`, func() {
		requestSupportApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/support/approvals"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestSupportApproval successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the RequestSupportApprovalOptions model
				requestSupportApprovalOptionsModel := new(partnercentersellv1.RequestSupportApprovalOptions)
				requestSupportApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.RequestSupportApprovalWithContext(ctx, requestSupportApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.RequestSupportApprovalWithContext(ctx, requestSupportApprovalOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(requestSupportApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestSupportApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.RequestSupportApproval(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RequestSupportApprovalOptions model
				requestSupportApprovalOptionsModel := new(partnercentersellv1.RequestSupportApprovalOptions)
				requestSupportApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RequestSupportApproval with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportApprovalOptions model
				requestSupportApprovalOptionsModel := new(partnercentersellv1.RequestSupportApprovalOptions)
				requestSupportApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RequestSupportApprovalOptions model with no property values
				requestSupportApprovalOptionsModelNew := new(partnercentersellv1.RequestSupportApprovalOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModelNew)
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
			It(`Invoke RequestSupportApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestSupportApprovalOptions model
				requestSupportApprovalOptionsModel := new(partnercentersellv1.RequestSupportApprovalOptions)
				requestSupportApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.RequestSupportApproval(requestSupportApprovalOptionsModel)
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
	Describe(`RequestProductApproval(requestProductApprovalOptions *RequestProductApprovalOptions) - Operation response error`, func() {
		requestProductApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/approvals"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestProductApprovalPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RequestProductApproval with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestProductApprovalOptions model
				requestProductApprovalOptionsModel := new(partnercentersellv1.RequestProductApprovalOptions)
				requestProductApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RequestProductApproval(requestProductApprovalOptions *RequestProductApprovalOptions)`, func() {
		requestProductApprovalPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/approvals"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(requestProductApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestProductApproval successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the RequestProductApprovalOptions model
				requestProductApprovalOptionsModel := new(partnercentersellv1.RequestProductApprovalOptions)
				requestProductApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.RequestProductApprovalWithContext(ctx, requestProductApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.RequestProductApprovalWithContext(ctx, requestProductApprovalOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(requestProductApprovalPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}`)
				}))
			})
			It(`Invoke RequestProductApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.RequestProductApproval(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RequestProductApprovalOptions model
				requestProductApprovalOptionsModel := new(partnercentersellv1.RequestProductApprovalOptions)
				requestProductApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke RequestProductApproval with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestProductApprovalOptions model
				requestProductApprovalOptionsModel := new(partnercentersellv1.RequestProductApprovalOptions)
				requestProductApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RequestProductApprovalOptions model with no property values
				requestProductApprovalOptionsModelNew := new(partnercentersellv1.RequestProductApprovalOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModelNew)
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
			It(`Invoke RequestProductApproval successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the RequestProductApprovalOptions model
				requestProductApprovalOptionsModel := new(partnercentersellv1.RequestProductApprovalOptions)
				requestProductApprovalOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.RequestProductApproval(requestProductApprovalOptionsModel)
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
	Describe(`ListProductApprovals(listProductApprovalsOptions *ListProductApprovalsOptions) - Operation response error`, func() {
		listProductApprovalsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/approvals"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductApprovalsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProductApprovals with error: Operation response processing error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductApprovalsOptions model
				listProductApprovalsOptionsModel := new(partnercentersellv1.ListProductApprovalsOptions)
				listProductApprovalsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				partnerCenterSellService.EnableRetries(0, 0)
				result, response, operationErr = partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListProductApprovals(listProductApprovalsOptions *ListProductApprovalsOptions)`, func() {
		listProductApprovalsPath := "/products/9fab83da-98cb-4f18-a7ba-b6f0435c9673/approvals"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProductApprovalsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"approvals": [{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}]}`)
				}))
			})
			It(`Invoke ListProductApprovals successfully with retries`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())
				partnerCenterSellService.EnableRetries(0, 0)

				// Construct an instance of the ListProductApprovalsOptions model
				listProductApprovalsOptionsModel := new(partnercentersellv1.ListProductApprovalsOptions)
				listProductApprovalsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := partnerCenterSellService.ListProductApprovalsWithContext(ctx, listProductApprovalsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				partnerCenterSellService.DisableRetries()
				result, response, operationErr := partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = partnerCenterSellService.ListProductApprovalsWithContext(ctx, listProductApprovalsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProductApprovalsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"approvals": [{"history": [{"at": "At", "displayName": "DisplayName", "event": "Event", "reason": "Reason", "username": "Username"}], "id": "ID", "nextEvents": ["NextEvents"], "state": {"name": "waitingForRequestForReview"}}]}`)
				}))
			})
			It(`Invoke ListProductApprovals successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := partnerCenterSellService.ListProductApprovals(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProductApprovalsOptions model
				listProductApprovalsOptionsModel := new(partnercentersellv1.ListProductApprovalsOptions)
				listProductApprovalsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProductApprovals with error: Operation validation and request error`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductApprovalsOptions model
				listProductApprovalsOptionsModel := new(partnercentersellv1.ListProductApprovalsOptions)
				listProductApprovalsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := partnerCenterSellService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProductApprovalsOptions model with no property values
				listProductApprovalsOptionsModelNew := new(partnercentersellv1.ListProductApprovalsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModelNew)
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
			It(`Invoke ListProductApprovals successfully`, func() {
				partnerCenterSellService, serviceErr := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(partnerCenterSellService).ToNot(BeNil())

				// Construct an instance of the ListProductApprovalsOptions model
				listProductApprovalsOptionsModel := new(partnercentersellv1.ListProductApprovalsOptions)
				listProductApprovalsOptionsModel.ProductID = CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := partnerCenterSellService.ListProductApprovals(listProductApprovalsOptionsModel)
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
			partnerCenterSellService, _ := partnercentersellv1.NewPartnerCenterSellV1(&partnercentersellv1.PartnerCenterSellV1Options{
				URL:           "http://partnercentersellv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreatePlanOptions successfully`, func() {
				// Construct an instance of the CreatePlanOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				createPlanOptionsDescription := "testString"
				createPlanOptionsLabel := "testString"
				createPlanOptionsType := "byol"
				createPlanOptionsURL := "testString"
				createPlanOptionsModel := partnerCenterSellService.NewCreatePlanOptions(productID, createPlanOptionsDescription, createPlanOptionsLabel, createPlanOptionsType, createPlanOptionsURL)
				createPlanOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				createPlanOptionsModel.SetDescription("testString")
				createPlanOptionsModel.SetLabel("testString")
				createPlanOptionsModel.SetType("byol")
				createPlanOptionsModel.SetURL("testString")
				createPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPlanOptionsModel).ToNot(BeNil())
				Expect(createPlanOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(createPlanOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createPlanOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(createPlanOptionsModel.Type).To(Equal(core.StringPtr("byol")))
				Expect(createPlanOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(createPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateProductOptions successfully`, func() {
				// Construct an instance of the CreateProductOptions model
				createProductOptionsProductName := "testString"
				createProductOptionsTaxAssessment := "SOFTWARE"
				createProductOptionsProductType := "SOFTWARE"
				createProductOptionsModel := partnerCenterSellService.NewCreateProductOptions(createProductOptionsProductName, createProductOptionsTaxAssessment, createProductOptionsProductType)
				createProductOptionsModel.SetProductName("testString")
				createProductOptionsModel.SetTaxAssessment("SOFTWARE")
				createProductOptionsModel.SetProductType("SOFTWARE")
				createProductOptionsModel.SetMaterialAgreement(true)
				createProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createProductOptionsModel).ToNot(BeNil())
				Expect(createProductOptionsModel.ProductName).To(Equal(core.StringPtr("testString")))
				Expect(createProductOptionsModel.TaxAssessment).To(Equal(core.StringPtr("SOFTWARE")))
				Expect(createProductOptionsModel.ProductType).To(Equal(core.StringPtr("SOFTWARE")))
				Expect(createProductOptionsModel.MaterialAgreement).To(Equal(core.BoolPtr(true)))
				Expect(createProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSupportChangeRequestOptions successfully`, func() {
				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				Expect(supportDetailsAvailabilityTimesModel).ToNot(BeNil())
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")
				Expect(supportDetailsAvailabilityTimesModel.Day).To(Equal(core.Int64Ptr(int64(1))))
				Expect(supportDetailsAvailabilityTimesModel.EndTime).To(Equal(core.StringPtr("19:30")))
				Expect(supportDetailsAvailabilityTimesModel.StartTime).To(Equal(core.StringPtr("10:30")))

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				Expect(supportDetailsAvailabilityModel).ToNot(BeNil())
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")
				Expect(supportDetailsAvailabilityModel.AlwaysAvailable).To(Equal(core.BoolPtr(true)))
				Expect(supportDetailsAvailabilityModel.Times).To(Equal([]partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}))
				Expect(supportDetailsAvailabilityModel.Timezone).To(Equal(core.StringPtr("America/Los_Angeles")))

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				Expect(supportResponseTimesModel).ToNot(BeNil())
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportResponseTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportResponseTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				Expect(supportDetailsModel).ToNot(BeNil())
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")
				Expect(supportDetailsModel.Availability).To(Equal(supportDetailsAvailabilityModel))
				Expect(supportDetailsModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))
				Expect(supportDetailsModel.Type).To(Equal(core.StringPtr("email")))

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				Expect(supportEscalationTimesModel).ToNot(BeNil())
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportEscalationTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportEscalationTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				Expect(supportEscalationModel).ToNot(BeNil())
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel
				Expect(supportEscalationModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportEscalationModel.EscalationWaitTime).To(Equal(supportEscalationTimesModel))
				Expect(supportEscalationModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				Expect(supportModel).ToNot(BeNil())
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")
				Expect(supportModel.Locations).To(Equal([]string{"US"}))
				Expect(supportModel.Process).To(Equal(core.StringPtr("testString")))
				Expect(supportModel.ProcessI18n).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(supportModel.SupportDetails).To(Equal([]partnercentersellv1.SupportDetails{*supportDetailsModel}))
				Expect(supportModel.SupportEscalation).To(Equal(supportEscalationModel))
				Expect(supportModel.SupportType).To(Equal(core.StringPtr("third-party")))
				Expect(supportModel.URL).To(Equal(core.StringPtr("https://my-company.com/support")))

				// Construct an instance of the CreateSupportChangeRequestOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				var createSupportChangeRequestOptionsChange *partnercentersellv1.Support = nil
				createSupportChangeRequestOptionsModel := partnerCenterSellService.NewCreateSupportChangeRequestOptions(productID, createSupportChangeRequestOptionsChange)
				createSupportChangeRequestOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				createSupportChangeRequestOptionsModel.SetChange(supportModel)
				createSupportChangeRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSupportChangeRequestOptionsModel).ToNot(BeNil())
				Expect(createSupportChangeRequestOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(createSupportChangeRequestOptionsModel.Change).To(Equal(supportModel))
				Expect(createSupportChangeRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePlanOptions successfully`, func() {
				// Construct an instance of the DeletePlanOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				pricingPlanID := "testString"
				deletePlanOptionsModel := partnerCenterSellService.NewDeletePlanOptions(productID, pricingPlanID)
				deletePlanOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deletePlanOptionsModel.SetPricingPlanID("testString")
				deletePlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePlanOptionsModel).ToNot(BeNil())
				Expect(deletePlanOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deletePlanOptionsModel.PricingPlanID).To(Equal(core.StringPtr("testString")))
				Expect(deletePlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteProductOptions successfully`, func() {
				// Construct an instance of the DeleteProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deleteProductOptionsModel := partnerCenterSellService.NewDeleteProductOptions(productID)
				deleteProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deleteProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteProductOptionsModel).ToNot(BeNil())
				Expect(deleteProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deleteProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeprecateProductOptions successfully`, func() {
				// Construct an instance of the DeprecateProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				deprecateProductOptionsReason := "testString"
				deprecateProductOptionsModel := partnerCenterSellService.NewDeprecateProductOptions(productID, deprecateProductOptionsReason)
				deprecateProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				deprecateProductOptionsModel.SetReason("testString")
				deprecateProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deprecateProductOptionsModel).ToNot(BeNil())
				Expect(deprecateProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(deprecateProductOptionsModel.Reason).To(Equal(core.StringPtr("testString")))
				Expect(deprecateProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetBadgeOptions successfully`, func() {
				// Construct an instance of the GetBadgeOptions model
				badgeID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getBadgeOptionsModel := partnerCenterSellService.NewGetBadgeOptions(badgeID)
				getBadgeOptionsModel.SetBadgeID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getBadgeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getBadgeOptionsModel).ToNot(BeNil())
				Expect(getBadgeOptionsModel.BadgeID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getBadgeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getCatalogOptionsModel := partnerCenterSellService.NewGetCatalogOptions(productID)
				getCatalogOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPlanOptions successfully`, func() {
				// Construct an instance of the GetPlanOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				pricingPlanID := "testString"
				getPlanOptionsModel := partnerCenterSellService.NewGetPlanOptions(productID, pricingPlanID)
				getPlanOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getPlanOptionsModel.SetPricingPlanID("testString")
				getPlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPlanOptionsModel).ToNot(BeNil())
				Expect(getPlanOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getPlanOptionsModel.PricingPlanID).To(Equal(core.StringPtr("testString")))
				Expect(getPlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetProductOptions successfully`, func() {
				// Construct an instance of the GetProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getProductOptionsModel := partnerCenterSellService.NewGetProductOptions(productID)
				getProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getProductOptionsModel).ToNot(BeNil())
				Expect(getProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSupportChangeRequestOptions successfully`, func() {
				// Construct an instance of the GetSupportChangeRequestOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				changeRequestID := "testString"
				getSupportChangeRequestOptionsModel := partnerCenterSellService.NewGetSupportChangeRequestOptions(productID, changeRequestID)
				getSupportChangeRequestOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getSupportChangeRequestOptionsModel.SetChangeRequestID("testString")
				getSupportChangeRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSupportChangeRequestOptionsModel).ToNot(BeNil())
				Expect(getSupportChangeRequestOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getSupportChangeRequestOptionsModel.ChangeRequestID).To(Equal(core.StringPtr("testString")))
				Expect(getSupportChangeRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSupportOptions successfully`, func() {
				// Construct an instance of the GetSupportOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				getSupportOptionsModel := partnerCenterSellService.NewGetSupportOptions(productID)
				getSupportOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				getSupportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSupportOptionsModel).ToNot(BeNil())
				Expect(getSupportOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(getSupportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListBadgesOptions successfully`, func() {
				// Construct an instance of the ListBadgesOptions model
				listBadgesOptionsModel := partnerCenterSellService.NewListBadgesOptions()
				listBadgesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listBadgesOptionsModel).ToNot(BeNil())
				Expect(listBadgesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPlansOptions successfully`, func() {
				// Construct an instance of the ListPlansOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listPlansOptionsModel := partnerCenterSellService.NewListPlansOptions(productID)
				listPlansOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listPlansOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPlansOptionsModel).ToNot(BeNil())
				Expect(listPlansOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listPlansOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProductApprovalsOptions successfully`, func() {
				// Construct an instance of the ListProductApprovalsOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listProductApprovalsOptionsModel := partnerCenterSellService.NewListProductApprovalsOptions(productID)
				listProductApprovalsOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listProductApprovalsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProductApprovalsOptionsModel).ToNot(BeNil())
				Expect(listProductApprovalsOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listProductApprovalsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProductsOptions successfully`, func() {
				// Construct an instance of the ListProductsOptions model
				listProductsOptionsModel := partnerCenterSellService.NewListProductsOptions()
				listProductsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProductsOptionsModel).ToNot(BeNil())
				Expect(listProductsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSupportChangeRequestReviewsOptions successfully`, func() {
				// Construct an instance of the ListSupportChangeRequestReviewsOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				changeRequestID := "testString"
				listSupportChangeRequestReviewsOptionsModel := partnerCenterSellService.NewListSupportChangeRequestReviewsOptions(productID, changeRequestID)
				listSupportChangeRequestReviewsOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listSupportChangeRequestReviewsOptionsModel.SetChangeRequestID("testString")
				listSupportChangeRequestReviewsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSupportChangeRequestReviewsOptionsModel).ToNot(BeNil())
				Expect(listSupportChangeRequestReviewsOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listSupportChangeRequestReviewsOptionsModel.ChangeRequestID).To(Equal(core.StringPtr("testString")))
				Expect(listSupportChangeRequestReviewsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSupportChangeRequestsOptions successfully`, func() {
				// Construct an instance of the ListSupportChangeRequestsOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				listSupportChangeRequestsOptionsModel := partnerCenterSellService.NewListSupportChangeRequestsOptions(productID)
				listSupportChangeRequestsOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				listSupportChangeRequestsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSupportChangeRequestsOptionsModel).ToNot(BeNil())
				Expect(listSupportChangeRequestsOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(listSupportChangeRequestsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMediaSectionInput successfully`, func() {
				caption := "testString"
				typeVar := "image"
				url := "testString"
				_model, err := partnerCenterSellService.NewMediaSectionInput(caption, typeVar, url)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewMergeSupportChangeRequestOptions successfully`, func() {
				// Construct an instance of the MergeSupportChangeRequestOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				changeRequestID := "testString"
				mergeSupportChangeRequestOptionsModel := partnerCenterSellService.NewMergeSupportChangeRequestOptions(productID, changeRequestID)
				mergeSupportChangeRequestOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				mergeSupportChangeRequestOptionsModel.SetChangeRequestID("testString")
				mergeSupportChangeRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(mergeSupportChangeRequestOptionsModel).ToNot(BeNil())
				Expect(mergeSupportChangeRequestOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(mergeSupportChangeRequestOptionsModel.ChangeRequestID).To(Equal(core.StringPtr("testString")))
				Expect(mergeSupportChangeRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPublishProductOptions successfully`, func() {
				// Construct an instance of the PublishProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				publishProductOptionsModel := partnerCenterSellService.NewPublishProductOptions(productID)
				publishProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				publishProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(publishProductOptionsModel).ToNot(BeNil())
				Expect(publishProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(publishProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRequestCatalogApprovalOptions successfully`, func() {
				// Construct an instance of the RequestCatalogApprovalOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestCatalogApprovalOptionsModel := partnerCenterSellService.NewRequestCatalogApprovalOptions(productID)
				requestCatalogApprovalOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				requestCatalogApprovalOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(requestCatalogApprovalOptionsModel).ToNot(BeNil())
				Expect(requestCatalogApprovalOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(requestCatalogApprovalOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRequestProductApprovalOptions successfully`, func() {
				// Construct an instance of the RequestProductApprovalOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestProductApprovalOptionsModel := partnerCenterSellService.NewRequestProductApprovalOptions(productID)
				requestProductApprovalOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				requestProductApprovalOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(requestProductApprovalOptionsModel).ToNot(BeNil())
				Expect(requestProductApprovalOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(requestProductApprovalOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRequestSupportApprovalOptions successfully`, func() {
				// Construct an instance of the RequestSupportApprovalOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				requestSupportApprovalOptionsModel := partnerCenterSellService.NewRequestSupportApprovalOptions(productID)
				requestSupportApprovalOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				requestSupportApprovalOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(requestSupportApprovalOptionsModel).ToNot(BeNil())
				Expect(requestSupportApprovalOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(requestSupportApprovalOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRequestSupportChangeRequestReviewOptions successfully`, func() {
				// Construct an instance of the RequestSupportChangeRequestReviewOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				changeRequestID := "testString"
				requestSupportChangeRequestReviewOptionsModel := partnerCenterSellService.NewRequestSupportChangeRequestReviewOptions(productID, changeRequestID)
				requestSupportChangeRequestReviewOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				requestSupportChangeRequestReviewOptionsModel.SetChangeRequestID("testString")
				requestSupportChangeRequestReviewOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(requestSupportChangeRequestReviewOptionsModel).ToNot(BeNil())
				Expect(requestSupportChangeRequestReviewOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(requestSupportChangeRequestReviewOptionsModel.ChangeRequestID).To(Equal(core.StringPtr("testString")))
				Expect(requestSupportChangeRequestReviewOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRestoreProductOptions successfully`, func() {
				// Construct an instance of the RestoreProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				restoreProductOptionsReason := "testString"
				restoreProductOptionsModel := partnerCenterSellService.NewRestoreProductOptions(productID, restoreProductOptionsReason)
				restoreProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				restoreProductOptionsModel.SetReason("testString")
				restoreProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(restoreProductOptionsModel).ToNot(BeNil())
				Expect(restoreProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(restoreProductOptionsModel.Reason).To(Equal(core.StringPtr("testString")))
				Expect(restoreProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewSuspendProductOptions successfully`, func() {
				// Construct an instance of the SuspendProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				suspendProductOptionsReason := "testString"
				suspendProductOptionsModel := partnerCenterSellService.NewSuspendProductOptions(productID, suspendProductOptionsReason)
				suspendProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				suspendProductOptionsModel.SetReason("testString")
				suspendProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(suspendProductOptionsModel).ToNot(BeNil())
				Expect(suspendProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(suspendProductOptionsModel.Reason).To(Equal(core.StringPtr("testString")))
				Expect(suspendProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateCatalogOptions successfully`, func() {
				// Construct an instance of the HighlightSectionInput model
				highlightSectionInputModel := new(partnercentersellv1.HighlightSectionInput)
				Expect(highlightSectionInputModel).ToNot(BeNil())
				highlightSectionInputModel.Description = core.StringPtr("testString")
				highlightSectionInputModel.Title = core.StringPtr("testString")
				Expect(highlightSectionInputModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(highlightSectionInputModel.Title).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the MediaSectionInput model
				mediaSectionInputModel := new(partnercentersellv1.MediaSectionInput)
				Expect(mediaSectionInputModel).ToNot(BeNil())
				mediaSectionInputModel.Caption = core.StringPtr("testString")
				mediaSectionInputModel.Thumbnail = core.StringPtr("testString")
				mediaSectionInputModel.Type = core.StringPtr("image")
				mediaSectionInputModel.URL = core.StringPtr("testString")
				Expect(mediaSectionInputModel.Caption).To(Equal(core.StringPtr("testString")))
				Expect(mediaSectionInputModel.Thumbnail).To(Equal(core.StringPtr("testString")))
				Expect(mediaSectionInputModel.Type).To(Equal(core.StringPtr("image")))
				Expect(mediaSectionInputModel.URL).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateCatalogOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateCatalogOptionsModel := partnerCenterSellService.NewUpdateCatalogOptions(productID)
				updateCatalogOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateCatalogOptionsModel.SetCatalogID("testString")
				updateCatalogOptionsModel.SetDescription("testString")
				updateCatalogOptionsModel.SetIconURL("testString")
				updateCatalogOptionsModel.SetKeywords([]string{"testString"})
				updateCatalogOptionsModel.SetPricingModel("free")
				updateCatalogOptionsModel.SetCategory("testString")
				updateCatalogOptionsModel.SetProviderType([]string{"ibm_community"})
				updateCatalogOptionsModel.SetLabel("testString")
				updateCatalogOptionsModel.SetName("testString")
				updateCatalogOptionsModel.SetProvider("testString")
				updateCatalogOptionsModel.SetTags([]string{"testString"})
				updateCatalogOptionsModel.SetDocumentationURL("testString")
				updateCatalogOptionsModel.SetHighlights([]partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel})
				updateCatalogOptionsModel.SetLongDescription("testString")
				updateCatalogOptionsModel.SetMedia([]partnercentersellv1.MediaSectionInput{*mediaSectionInputModel})
				updateCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCatalogOptionsModel).ToNot(BeNil())
				Expect(updateCatalogOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateCatalogOptionsModel.CatalogID).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.IconURL).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Keywords).To(Equal([]string{"testString"}))
				Expect(updateCatalogOptionsModel.PricingModel).To(Equal(core.StringPtr("free")))
				Expect(updateCatalogOptionsModel.Category).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.ProviderType).To(Equal([]string{"ibm_community"}))
				Expect(updateCatalogOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Provider).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(updateCatalogOptionsModel.DocumentationURL).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Highlights).To(Equal([]partnercentersellv1.HighlightSectionInput{*highlightSectionInputModel}))
				Expect(updateCatalogOptionsModel.LongDescription).To(Equal(core.StringPtr("testString")))
				Expect(updateCatalogOptionsModel.Media).To(Equal([]partnercentersellv1.MediaSectionInput{*mediaSectionInputModel}))
				Expect(updateCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePlanOptions successfully`, func() {
				// Construct an instance of the UpdatePlanOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				pricingPlanID := "testString"
				updatePlanOptionsDescription := "testString"
				updatePlanOptionsLabel := "testString"
				updatePlanOptionsType := "byol"
				updatePlanOptionsURL := "testString"
				updatePlanOptionsModel := partnerCenterSellService.NewUpdatePlanOptions(productID, pricingPlanID, updatePlanOptionsDescription, updatePlanOptionsLabel, updatePlanOptionsType, updatePlanOptionsURL)
				updatePlanOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updatePlanOptionsModel.SetPricingPlanID("testString")
				updatePlanOptionsModel.SetDescription("testString")
				updatePlanOptionsModel.SetLabel("testString")
				updatePlanOptionsModel.SetType("byol")
				updatePlanOptionsModel.SetURL("testString")
				updatePlanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePlanOptionsModel).ToNot(BeNil())
				Expect(updatePlanOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updatePlanOptionsModel.PricingPlanID).To(Equal(core.StringPtr("testString")))
				Expect(updatePlanOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updatePlanOptionsModel.Label).To(Equal(core.StringPtr("testString")))
				Expect(updatePlanOptionsModel.Type).To(Equal(core.StringPtr("byol")))
				Expect(updatePlanOptionsModel.URL).To(Equal(core.StringPtr("testString")))
				Expect(updatePlanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateProductOptions successfully`, func() {
				// Construct an instance of the UpdateProductOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateProductOptionsModel := partnerCenterSellService.NewUpdateProductOptions(productID)
				updateProductOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateProductOptionsModel.SetMaterialAgreement(true)
				updateProductOptionsModel.SetProductName("testString")
				updateProductOptionsModel.SetTaxAssessment("SOFTWARE")
				updateProductOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateProductOptionsModel).ToNot(BeNil())
				Expect(updateProductOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateProductOptionsModel.MaterialAgreement).To(Equal(core.BoolPtr(true)))
				Expect(updateProductOptionsModel.ProductName).To(Equal(core.StringPtr("testString")))
				Expect(updateProductOptionsModel.TaxAssessment).To(Equal(core.StringPtr("SOFTWARE")))
				Expect(updateProductOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSupportChangeRequestOptions successfully`, func() {
				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				Expect(supportDetailsAvailabilityTimesModel).ToNot(BeNil())
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")
				Expect(supportDetailsAvailabilityTimesModel.Day).To(Equal(core.Int64Ptr(int64(1))))
				Expect(supportDetailsAvailabilityTimesModel.EndTime).To(Equal(core.StringPtr("19:30")))
				Expect(supportDetailsAvailabilityTimesModel.StartTime).To(Equal(core.StringPtr("10:30")))

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				Expect(supportDetailsAvailabilityModel).ToNot(BeNil())
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")
				Expect(supportDetailsAvailabilityModel.AlwaysAvailable).To(Equal(core.BoolPtr(true)))
				Expect(supportDetailsAvailabilityModel.Times).To(Equal([]partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}))
				Expect(supportDetailsAvailabilityModel.Timezone).To(Equal(core.StringPtr("America/Los_Angeles")))

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				Expect(supportResponseTimesModel).ToNot(BeNil())
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportResponseTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportResponseTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				Expect(supportDetailsModel).ToNot(BeNil())
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")
				Expect(supportDetailsModel.Availability).To(Equal(supportDetailsAvailabilityModel))
				Expect(supportDetailsModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))
				Expect(supportDetailsModel.Type).To(Equal(core.StringPtr("email")))

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				Expect(supportEscalationTimesModel).ToNot(BeNil())
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportEscalationTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportEscalationTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				Expect(supportEscalationModel).ToNot(BeNil())
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel
				Expect(supportEscalationModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportEscalationModel.EscalationWaitTime).To(Equal(supportEscalationTimesModel))
				Expect(supportEscalationModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))

				// Construct an instance of the Support model
				supportModel := new(partnercentersellv1.Support)
				Expect(supportModel).ToNot(BeNil())
				supportModel.Locations = []string{"US"}
				supportModel.Process = core.StringPtr("testString")
				supportModel.ProcessI18n = map[string]interface{}{"anyKey": "anyValue"}
				supportModel.SupportDetails = []partnercentersellv1.SupportDetails{*supportDetailsModel}
				supportModel.SupportEscalation = supportEscalationModel
				supportModel.SupportType = core.StringPtr("third-party")
				supportModel.URL = core.StringPtr("https://my-company.com/support")
				Expect(supportModel.Locations).To(Equal([]string{"US"}))
				Expect(supportModel.Process).To(Equal(core.StringPtr("testString")))
				Expect(supportModel.ProcessI18n).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(supportModel.SupportDetails).To(Equal([]partnercentersellv1.SupportDetails{*supportDetailsModel}))
				Expect(supportModel.SupportEscalation).To(Equal(supportEscalationModel))
				Expect(supportModel.SupportType).To(Equal(core.StringPtr("third-party")))
				Expect(supportModel.URL).To(Equal(core.StringPtr("https://my-company.com/support")))

				// Construct an instance of the UpdateSupportChangeRequestOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				changeRequestID := "testString"
				var updateSupportChangeRequestOptionsChange *partnercentersellv1.Support = nil
				updateSupportChangeRequestOptionsModel := partnerCenterSellService.NewUpdateSupportChangeRequestOptions(productID, changeRequestID, updateSupportChangeRequestOptionsChange)
				updateSupportChangeRequestOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateSupportChangeRequestOptionsModel.SetChangeRequestID("testString")
				updateSupportChangeRequestOptionsModel.SetChange(supportModel)
				updateSupportChangeRequestOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSupportChangeRequestOptionsModel).ToNot(BeNil())
				Expect(updateSupportChangeRequestOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateSupportChangeRequestOptionsModel.ChangeRequestID).To(Equal(core.StringPtr("testString")))
				Expect(updateSupportChangeRequestOptionsModel.Change).To(Equal(supportModel))
				Expect(updateSupportChangeRequestOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSupportOptions successfully`, func() {
				// Construct an instance of the EscalationContactsUpdate model
				escalationContactsUpdateModel := new(partnercentersellv1.EscalationContactsUpdate)
				Expect(escalationContactsUpdateModel).ToNot(BeNil())
				escalationContactsUpdateModel.Email = core.StringPtr("testString")
				escalationContactsUpdateModel.Name = core.StringPtr("testString")
				Expect(escalationContactsUpdateModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(escalationContactsUpdateModel.Name).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SupportDetailsAvailabilityTimes model
				supportDetailsAvailabilityTimesModel := new(partnercentersellv1.SupportDetailsAvailabilityTimes)
				Expect(supportDetailsAvailabilityTimesModel).ToNot(BeNil())
				supportDetailsAvailabilityTimesModel.Day = core.Int64Ptr(int64(1))
				supportDetailsAvailabilityTimesModel.EndTime = core.StringPtr("19:30")
				supportDetailsAvailabilityTimesModel.StartTime = core.StringPtr("10:30")
				Expect(supportDetailsAvailabilityTimesModel.Day).To(Equal(core.Int64Ptr(int64(1))))
				Expect(supportDetailsAvailabilityTimesModel.EndTime).To(Equal(core.StringPtr("19:30")))
				Expect(supportDetailsAvailabilityTimesModel.StartTime).To(Equal(core.StringPtr("10:30")))

				// Construct an instance of the SupportDetailsAvailability model
				supportDetailsAvailabilityModel := new(partnercentersellv1.SupportDetailsAvailability)
				Expect(supportDetailsAvailabilityModel).ToNot(BeNil())
				supportDetailsAvailabilityModel.AlwaysAvailable = core.BoolPtr(true)
				supportDetailsAvailabilityModel.Times = []partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}
				supportDetailsAvailabilityModel.Timezone = core.StringPtr("America/Los_Angeles")
				Expect(supportDetailsAvailabilityModel.AlwaysAvailable).To(Equal(core.BoolPtr(true)))
				Expect(supportDetailsAvailabilityModel.Times).To(Equal([]partnercentersellv1.SupportDetailsAvailabilityTimes{*supportDetailsAvailabilityTimesModel}))
				Expect(supportDetailsAvailabilityModel.Timezone).To(Equal(core.StringPtr("America/Los_Angeles")))

				// Construct an instance of the SupportResponseTimes model
				supportResponseTimesModel := new(partnercentersellv1.SupportResponseTimes)
				Expect(supportResponseTimesModel).ToNot(BeNil())
				supportResponseTimesModel.Type = core.StringPtr("hour")
				supportResponseTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportResponseTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportResponseTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportDetails model
				supportDetailsModel := new(partnercentersellv1.SupportDetails)
				Expect(supportDetailsModel).ToNot(BeNil())
				supportDetailsModel.Availability = supportDetailsAvailabilityModel
				supportDetailsModel.Contact = core.StringPtr("testString")
				supportDetailsModel.ResponseWaitTime = supportResponseTimesModel
				supportDetailsModel.Type = core.StringPtr("email")
				Expect(supportDetailsModel.Availability).To(Equal(supportDetailsAvailabilityModel))
				Expect(supportDetailsModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportDetailsModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))
				Expect(supportDetailsModel.Type).To(Equal(core.StringPtr("email")))

				// Construct an instance of the SupportEscalationTimes model
				supportEscalationTimesModel := new(partnercentersellv1.SupportEscalationTimes)
				Expect(supportEscalationTimesModel).ToNot(BeNil())
				supportEscalationTimesModel.Type = core.StringPtr("hour")
				supportEscalationTimesModel.Value = core.Int64Ptr(int64(2))
				Expect(supportEscalationTimesModel.Type).To(Equal(core.StringPtr("hour")))
				Expect(supportEscalationTimesModel.Value).To(Equal(core.Int64Ptr(int64(2))))

				// Construct an instance of the SupportEscalation model
				supportEscalationModel := new(partnercentersellv1.SupportEscalation)
				Expect(supportEscalationModel).ToNot(BeNil())
				supportEscalationModel.Contact = core.StringPtr("testString")
				supportEscalationModel.EscalationWaitTime = supportEscalationTimesModel
				supportEscalationModel.ResponseWaitTime = supportResponseTimesModel
				Expect(supportEscalationModel.Contact).To(Equal(core.StringPtr("testString")))
				Expect(supportEscalationModel.EscalationWaitTime).To(Equal(supportEscalationTimesModel))
				Expect(supportEscalationModel.ResponseWaitTime).To(Equal(supportResponseTimesModel))

				// Construct an instance of the UpdateSupportOptions model
				productID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
				updateSupportOptionsModel := partnerCenterSellService.NewUpdateSupportOptions(productID)
				updateSupportOptionsModel.SetProductID(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673"))
				updateSupportOptionsModel.SetEscalationContacts([]partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel})
				updateSupportOptionsModel.SetLocations([]string{"US"})
				updateSupportOptionsModel.SetSupportDetails([]partnercentersellv1.SupportDetails{*supportDetailsModel})
				updateSupportOptionsModel.SetSupportEscalation(supportEscalationModel)
				updateSupportOptionsModel.SetSupportType("third-party")
				updateSupportOptionsModel.SetURL("https://my-company.com/support")
				updateSupportOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSupportOptionsModel).ToNot(BeNil())
				Expect(updateSupportOptionsModel.ProductID).To(Equal(CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")))
				Expect(updateSupportOptionsModel.EscalationContacts).To(Equal([]partnercentersellv1.EscalationContactsUpdate{*escalationContactsUpdateModel}))
				Expect(updateSupportOptionsModel.Locations).To(Equal([]string{"US"}))
				Expect(updateSupportOptionsModel.SupportDetails).To(Equal([]partnercentersellv1.SupportDetails{*supportDetailsModel}))
				Expect(updateSupportOptionsModel.SupportEscalation).To(Equal(supportEscalationModel))
				Expect(updateSupportOptionsModel.SupportType).To(Equal(core.StringPtr("third-party")))
				Expect(updateSupportOptionsModel.URL).To(Equal(core.StringPtr("https://my-company.com/support")))
				Expect(updateSupportOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
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
