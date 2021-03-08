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

package securitycompliancev1_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/securitycompliancev1"
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

var _ = Describe(`SecurityComplianceV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(securityComplianceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "https://securitycompliancev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
					URL: "https://testService/api",
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				err := securityComplianceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = securitycompliancev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateValidationScan(createValidationScanOptions *CreateValidationScanOptions) - Operation response error`, func() {
		createValidationScanPath := "/posture/v1/scans/validation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValidationScanPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateValidationScan with error: Operation response processing error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the CreateValidationScanOptions model
				createValidationScanOptionsModel := new(securitycompliancev1.CreateValidationScanOptions)
				createValidationScanOptionsModel.AccountID = core.StringPtr("testString")
				createValidationScanOptionsModel.ScopeID = core.Int64Ptr(int64(1))
				createValidationScanOptionsModel.ProfileID = core.Int64Ptr(int64(6))
				createValidationScanOptionsModel.GroupProfileID = core.Int64Ptr(int64(13))
				createValidationScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityComplianceService.CreateValidationScan(createValidationScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityComplianceService.EnableRetries(0, 0)
				result, response, operationErr = securityComplianceService.CreateValidationScan(createValidationScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateValidationScan(createValidationScanOptions *CreateValidationScanOptions)`, func() {
		createValidationScanPath := "/posture/v1/scans/validation"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createValidationScanPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": true, "message": "Message"}`)
				}))
			})
			It(`Invoke CreateValidationScan successfully with retries`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				securityComplianceService.EnableRetries(0, 0)

				// Construct an instance of the CreateValidationScanOptions model
				createValidationScanOptionsModel := new(securitycompliancev1.CreateValidationScanOptions)
				createValidationScanOptionsModel.AccountID = core.StringPtr("testString")
				createValidationScanOptionsModel.ScopeID = core.Int64Ptr(int64(1))
				createValidationScanOptionsModel.ProfileID = core.Int64Ptr(int64(6))
				createValidationScanOptionsModel.GroupProfileID = core.Int64Ptr(int64(13))
				createValidationScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityComplianceService.CreateValidationScanWithContext(ctx, createValidationScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityComplianceService.DisableRetries()
				result, response, operationErr := securityComplianceService.CreateValidationScan(createValidationScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityComplianceService.CreateValidationScanWithContext(ctx, createValidationScanOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createValidationScanPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"result": true, "message": "Message"}`)
				}))
			})
			It(`Invoke CreateValidationScan successfully`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityComplianceService.CreateValidationScan(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateValidationScanOptions model
				createValidationScanOptionsModel := new(securitycompliancev1.CreateValidationScanOptions)
				createValidationScanOptionsModel.AccountID = core.StringPtr("testString")
				createValidationScanOptionsModel.ScopeID = core.Int64Ptr(int64(1))
				createValidationScanOptionsModel.ProfileID = core.Int64Ptr(int64(6))
				createValidationScanOptionsModel.GroupProfileID = core.Int64Ptr(int64(13))
				createValidationScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityComplianceService.CreateValidationScan(createValidationScanOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateValidationScan with error: Operation validation and request error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the CreateValidationScanOptions model
				createValidationScanOptionsModel := new(securitycompliancev1.CreateValidationScanOptions)
				createValidationScanOptionsModel.AccountID = core.StringPtr("testString")
				createValidationScanOptionsModel.ScopeID = core.Int64Ptr(int64(1))
				createValidationScanOptionsModel.ProfileID = core.Int64Ptr(int64(6))
				createValidationScanOptionsModel.GroupProfileID = core.Int64Ptr(int64(13))
				createValidationScanOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityComplianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityComplianceService.CreateValidationScan(createValidationScanOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateValidationScanOptions model with no property values
				createValidationScanOptionsModelNew := new(securitycompliancev1.CreateValidationScanOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityComplianceService.CreateValidationScan(createValidationScanOptionsModelNew)
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
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(securityComplianceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "https://securitycompliancev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
					URL: "https://testService/api",
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				err := securityComplianceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = securitycompliancev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListProfile(listProfileOptions *ListProfileOptions) - Operation response error`, func() {
		listProfilePath := "/posture/v1/profiles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListProfile with error: Operation response processing error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the ListProfileOptions model
				listProfileOptionsModel := new(securitycompliancev1.ListProfileOptions)
				listProfileOptionsModel.AccountID = core.StringPtr("testString")
				listProfileOptionsModel.Name = core.StringPtr("testString")
				listProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityComplianceService.ListProfile(listProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityComplianceService.EnableRetries(0, 0)
				result, response, operationErr = securityComplianceService.ListProfile(listProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListProfile(listProfileOptions *ListProfileOptions)`, func() {
		listProfilePath := "/posture/v1/profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "no_of_goals": 58, "description": "CIS IBM Foundations Benchmark 1.0.0", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["[IBM Cloud]"], "resource": ["[My_example_bucket]"], "environment_category": ["[Cloud]"], "resource_category": ["[Storage]"], "resource_type": ["Bucket"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "profile_id": 3045, "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "profile_type": "predefined", "created_time": "2021-02-26T04:07:25Z", "modified_time": "2021-02-26T04:07:25Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfile successfully with retries`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				securityComplianceService.EnableRetries(0, 0)

				// Construct an instance of the ListProfileOptions model
				listProfileOptionsModel := new(securitycompliancev1.ListProfileOptions)
				listProfileOptionsModel.AccountID = core.StringPtr("testString")
				listProfileOptionsModel.Name = core.StringPtr("testString")
				listProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityComplianceService.ListProfileWithContext(ctx, listProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityComplianceService.DisableRetries()
				result, response, operationErr := securityComplianceService.ListProfile(listProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityComplianceService.ListProfileWithContext(ctx, listProfileOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listProfilePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"profiles": [{"name": "CIS IBM Foundations Benchmark 1.0.0", "no_of_goals": 58, "description": "CIS IBM Foundations Benchmark 1.0.0", "version": 1, "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "reason_for_delete": "ReasonForDelete", "applicability_criteria": {"environment": ["[IBM Cloud]"], "resource": ["[My_example_bucket]"], "environment_category": ["[Cloud]"], "resource_category": ["[Storage]"], "resource_type": ["Bucket"], "software_details": {"anyKey": "anyValue"}, "os_details": {"anyKey": "anyValue"}, "additional_details": {"anyKey": "anyValue"}, "environment_category_description": {"mapKey": "Cloud"}, "environment_description": {"mapKey": "IBM Cloud"}, "resource_category_description": {"mapKey": "Storage"}, "resource_type_description": {"mapKey": "Bucket"}, "resource_description": {"mapKey": "My_specific_bucket"}}, "profile_id": 3045, "base_profile": "CIS IBM Foundations Benchmark 1.0.0", "profile_type": "predefined", "created_time": "2021-02-26T04:07:25Z", "modified_time": "2021-02-26T04:07:25Z", "enabled": true}]}`)
				}))
			})
			It(`Invoke ListProfile successfully`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityComplianceService.ListProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListProfileOptions model
				listProfileOptionsModel := new(securitycompliancev1.ListProfileOptions)
				listProfileOptionsModel.AccountID = core.StringPtr("testString")
				listProfileOptionsModel.Name = core.StringPtr("testString")
				listProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityComplianceService.ListProfile(listProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListProfile with error: Operation validation and request error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the ListProfileOptions model
				listProfileOptionsModel := new(securitycompliancev1.ListProfileOptions)
				listProfileOptionsModel.AccountID = core.StringPtr("testString")
				listProfileOptionsModel.Name = core.StringPtr("testString")
				listProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityComplianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityComplianceService.ListProfile(listProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListProfileOptions model with no property values
				listProfileOptionsModelNew := new(securitycompliancev1.ListProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityComplianceService.ListProfile(listProfileOptionsModelNew)
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
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(securityComplianceService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "https://securitycompliancev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(securityComplianceService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
					URL: "https://testService/api",
				})
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				})
				err := securityComplianceService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := securityComplianceService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != securityComplianceService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(securityComplianceService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(securityComplianceService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_URL": "https://securitycompliancev1/api",
				"SECURITY_COMPLIANCE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"SECURITY_COMPLIANCE_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1UsingExternalConfig(&securitycompliancev1.SecurityComplianceV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(securityComplianceService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = securitycompliancev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListScopes(listScopesOptions *ListScopesOptions) - Operation response error`, func() {
		listScopesPath := "/posture/v1/scopes"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListScopes with error: Operation response processing error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(securitycompliancev1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := securityComplianceService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				securityComplianceService.EnableRetries(0, 0)
				result, response, operationErr = securityComplianceService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListScopes(listScopesOptions *ListScopesOptions)`, func() {
		listScopesPath := "/posture/v1/scopes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "scope_id": 1, "name": "My_Example_Scope", "enabled": true, "environment_type": "ibm", "created_time": "2021-02-26T04:07:25Z", "modified_time": "2021-02-26T04:07:25Z", "last_scan_type": "fact_collection", "last_scan_type_description": "Fact collection", "last_scan_status_updated_time": "2021-02-26T04:07:25Z", "collectors_id": [12], "scans": [{"scan_id": 235, "discover_id": 49, "status": "validation_completed", "status_message": "The collector aborted the task during upgrade."}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully with retries`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())
				securityComplianceService.EnableRetries(0, 0)

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(securitycompliancev1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := securityComplianceService.ListScopesWithContext(ctx, listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				securityComplianceService.DisableRetries()
				result, response, operationErr := securityComplianceService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = securityComplianceService.ListScopesWithContext(ctx, listScopesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listScopesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"scopes": [{"description": "This scope targets all of the resources that are available in our IBM Cloud staging environment.", "created_by": "IBMid-5500081P68", "modified_by": "IBMid-5500081P68", "scope_id": 1, "name": "My_Example_Scope", "enabled": true, "environment_type": "ibm", "created_time": "2021-02-26T04:07:25Z", "modified_time": "2021-02-26T04:07:25Z", "last_scan_type": "fact_collection", "last_scan_type_description": "Fact collection", "last_scan_status_updated_time": "2021-02-26T04:07:25Z", "collectors_id": [12], "scans": [{"scan_id": 235, "discover_id": 49, "status": "validation_completed", "status_message": "The collector aborted the task during upgrade."}]}]}`)
				}))
			})
			It(`Invoke ListScopes successfully`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := securityComplianceService.ListScopes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(securitycompliancev1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = securityComplianceService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListScopes with error: Operation validation and request error`, func() {
				securityComplianceService, serviceErr := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(securityComplianceService).ToNot(BeNil())

				// Construct an instance of the ListScopesOptions model
				listScopesOptionsModel := new(securitycompliancev1.ListScopesOptions)
				listScopesOptionsModel.AccountID = core.StringPtr("testString")
				listScopesOptionsModel.Name = core.StringPtr("testString")
				listScopesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := securityComplianceService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := securityComplianceService.ListScopes(listScopesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListScopesOptions model with no property values
				listScopesOptionsModelNew := new(securitycompliancev1.ListScopesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = securityComplianceService.ListScopes(listScopesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			securityComplianceService, _ := securitycompliancev1.NewSecurityComplianceV1(&securitycompliancev1.SecurityComplianceV1Options{
				URL:           "http://securitycompliancev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateValidationScanOptions successfully`, func() {
				// Construct an instance of the CreateValidationScanOptions model
				accountID := "testString"
				createValidationScanOptionsModel := securityComplianceService.NewCreateValidationScanOptions(accountID)
				createValidationScanOptionsModel.SetAccountID("testString")
				createValidationScanOptionsModel.SetScopeID(int64(1))
				createValidationScanOptionsModel.SetProfileID(int64(6))
				createValidationScanOptionsModel.SetGroupProfileID(int64(13))
				createValidationScanOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createValidationScanOptionsModel).ToNot(BeNil())
				Expect(createValidationScanOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createValidationScanOptionsModel.ScopeID).To(Equal(core.Int64Ptr(int64(1))))
				Expect(createValidationScanOptionsModel.ProfileID).To(Equal(core.Int64Ptr(int64(6))))
				Expect(createValidationScanOptionsModel.GroupProfileID).To(Equal(core.Int64Ptr(int64(13))))
				Expect(createValidationScanOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListProfileOptions successfully`, func() {
				// Construct an instance of the ListProfileOptions model
				accountID := "testString"
				listProfileOptionsModel := securityComplianceService.NewListProfileOptions(accountID)
				listProfileOptionsModel.SetAccountID("testString")
				listProfileOptionsModel.SetName("testString")
				listProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listProfileOptionsModel).ToNot(BeNil())
				Expect(listProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listProfileOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListScopesOptions successfully`, func() {
				// Construct an instance of the ListScopesOptions model
				accountID := "testString"
				listScopesOptionsModel := securityComplianceService.NewListScopesOptions(accountID)
				listScopesOptionsModel.SetAccountID("testString")
				listScopesOptionsModel.SetName("testString")
				listScopesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listScopesOptionsModel).ToNot(BeNil())
				Expect(listScopesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listScopesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
