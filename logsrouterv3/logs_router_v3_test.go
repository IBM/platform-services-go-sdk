/**
 * (C) Copyright IBM Corp. 2026.
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

package logsrouterv3_test

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`LogsRouterV3`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(logsRouterService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(logsRouterService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
				URL: "https://logsrouterv3/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(logsRouterService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTER_URL": "https://logsrouterv3/api",
				"LOGS_ROUTER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3UsingExternalConfig(&logsrouterv3.LogsRouterV3Options{
				})
				Expect(logsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := logsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRouterService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3UsingExternalConfig(&logsrouterv3.LogsRouterV3Options{
					URL: "https://testService/api",
				})
				Expect(logsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRouterService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3UsingExternalConfig(&logsrouterv3.LogsRouterV3Options{
				})
				err := logsRouterService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := logsRouterService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != logsRouterService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(logsRouterService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(logsRouterService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTER_URL": "https://logsrouterv3/api",
				"LOGS_ROUTER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3UsingExternalConfig(&logsrouterv3.LogsRouterV3Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(logsRouterService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"LOGS_ROUTER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3UsingExternalConfig(&logsrouterv3.LogsRouterV3Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(logsRouterService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = logsrouterv3.GetServiceURLForRegion("au-syd")
			Expect(url).To(Equal("https://api.au-syd.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.au-syd")
			Expect(url).To(Equal("https://api.private.au-syd.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("br-sao")
			Expect(url).To(Equal("https://api.br-sao.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.br-sao")
			Expect(url).To(Equal("https://api.private.br-sao.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("ca-mon")
			Expect(url).To(Equal("https://api.ca-mon.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.ca-mon")
			Expect(url).To(Equal("https://api.private.ca-mon.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("ca-tor")
			Expect(url).To(Equal("https://api.ca-tor.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.ca-tor")
			Expect(url).To(Equal("https://api.private.ca-tor.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("eu-de")
			Expect(url).To(Equal("https://api.eu-de.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.eu-de")
			Expect(url).To(Equal("https://api.private.eu-de.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("eu-es")
			Expect(url).To(Equal("https://api.eu-es.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.eu-es")
			Expect(url).To(Equal("https://api.private.eu-es.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("eu-fr2")
			Expect(url).To(Equal("https://api.eu-fr2.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.eu-fr2")
			Expect(url).To(Equal("https://api.private.eu-fr2.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("eu-gb")
			Expect(url).To(Equal("https://api.eu-gb.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.eu-gb")
			Expect(url).To(Equal("https://api.private.eu-gb.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("jp-osa")
			Expect(url).To(Equal("https://api.jp-osa.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.jp-osa")
			Expect(url).To(Equal("https://api.private.jp-osa.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("jp-tok")
			Expect(url).To(Equal("https://api.jp-tok.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.jp-tok")
			Expect(url).To(Equal("https://api.private.jp-tok.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://api.us-east.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.us-east")
			Expect(url).To(Equal("https://api.private.us-east.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://api.us-south.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("private.us-south")
			Expect(url).To(Equal("https://api.private.us-south.logs-router.cloud.ibm.com/v3"))
			Expect(err).To(BeNil())

			url, err = logsrouterv3.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions) - Operation response error`, func() {
		createTargetPath := "/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTarget with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
		createTargetPath := "/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.CreateTargetWithContext(ctx, createTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(logsrouterv3.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.CreateTarget(createTargetOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(logsrouterv3.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				createTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.Region = core.StringPtr("us-south")
				createTargetOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.CreateTarget(createTargetOptionsModel)
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
	Describe(`ListTargets(listTargetsOptions *ListTargetsOptions) - Operation response error`, func() {
		listTargetsPath := "/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTargets with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(logsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
		listTargetsPath := "/targets"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(logsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.ListTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(logsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTargets with error: Operation request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(logsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.ListTargets(listTargetsOptionsModel)
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
			It(`Invoke ListTargets successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(logsrouterv3.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.ListTargets(listTargetsOptionsModel)
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
	Describe(`GetTarget(getTargetOptions *GetTargetOptions) - Operation response error`, func() {
		getTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTarget with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(logsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
		getTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke GetTarget successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(logsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.GetTargetWithContext(ctx, getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.GetTargetWithContext(ctx, getTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke GetTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.GetTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(logsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTarget with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(logsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTargetOptions model with no property values
				getTargetOptionsModelNew := new(logsrouterv3.GetTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.GetTarget(getTargetOptionsModelNew)
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
			It(`Invoke GetTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(logsrouterv3.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.GetTarget(getTargetOptionsModel)
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
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions) - Operation response error`, func() {
		updateTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateTarget with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
		updateTargetPath := "/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.UpdateTargetWithContext(ctx, updateTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateTargetPath))
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-lr-target-us-south", "crn": "crn:v1:bluemix:public:logs-router:us-south:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "destination_crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "target_type": "cloud_logs", "region": "us-south", "write_status": {"status": "success", "last_failure": "2025-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke UpdateTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.UpdateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateTarget with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.UpdateTarget(updateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateTargetOptions model with no property values
				updateTargetOptionsModelNew := new(logsrouterv3.UpdateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.UpdateTarget(updateTargetOptionsModelNew)
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
			It(`Invoke UpdateTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the UpdateTargetOptions model
				updateTargetOptionsModel := new(logsrouterv3.UpdateTargetOptions)
				updateTargetOptionsModel.ID = core.StringPtr("testString")
				updateTargetOptionsModel.Name = core.StringPtr("my-lr-target")
				updateTargetOptionsModel.DestinationCRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.UpdateTarget(updateTargetOptionsModel)
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
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
		deleteTargetPath := "/targets/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := logsRouterService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(logsrouterv3.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = logsRouterService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(logsrouterv3.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := logsRouterService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(logsrouterv3.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = logsRouterService.DeleteTarget(deleteTargetOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions) - Operation response error`, func() {
		createRoutePath := "/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRoute with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(logsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {
		createRoutePath := "/routes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke CreateRoute successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(logsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.CreateRouteWithContext(ctx, createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.CreateRouteWithContext(ctx, createRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
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
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke CreateRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.CreateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(logsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRoute with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(logsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRouteOptions model with no property values
				createRouteOptionsModelNew := new(logsrouterv3.CreateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.CreateRoute(createRouteOptionsModelNew)
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
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(logsrouterv3.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				createRouteOptionsModel.ManagedBy = core.StringPtr("enterprise")
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.CreateRoute(createRouteOptionsModel)
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
	Describe(`ListRoutes(listRoutesOptions *ListRoutesOptions) - Operation response error`, func() {
		listRoutesPath := "/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoutes with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(logsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {
		listRoutesPath := "/routes"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(logsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.ListRoutes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(logsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoutes with error: Operation request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(logsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.ListRoutes(listRoutesOptionsModel)
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
			It(`Invoke ListRoutes successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(logsrouterv3.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.ListRoutes(listRoutesOptionsModel)
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
	Describe(`GetRoute(getRouteOptions *GetRouteOptions) - Operation response error`, func() {
		getRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRoute with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(logsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {
		getRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke GetRoute successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(logsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.GetRouteWithContext(ctx, getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.GetRouteWithContext(ctx, getRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke GetRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.GetRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(logsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRoute with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(logsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRouteOptions model with no property values
				getRouteOptionsModelNew := new(logsrouterv3.GetRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.GetRoute(getRouteOptionsModelNew)
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
			It(`Invoke GetRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(logsrouterv3.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.GetRoute(getRouteOptionsModel)
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
	Describe(`UpdateRoute(updateRouteOptions *UpdateRouteOptions) - Operation response error`, func() {
		updateRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRoute with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(logsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRoute(updateRouteOptions *UpdateRouteOptions)`, func() {
		updateRoutePath := "/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke UpdateRoute successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(logsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.UpdateRouteWithContext(ctx, updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.UpdateRouteWithContext(ctx, updateRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRoutePath))
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "rules": [{"action": "send", "targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "inclusion_filters": [{"operand": "location", "operator": "is", "values": ["us-south"]}]}], "created_at": "2021-05-18T20:15:12.353Z", "updated_at": "2021-05-18T20:15:12.353Z", "managed_by": "enterprise"}`)
				}))
			})
			It(`Invoke UpdateRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.UpdateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(logsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRoute with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(logsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.UpdateRoute(updateRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRouteOptions model with no property values
				updateRouteOptionsModelNew := new(logsrouterv3.UpdateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.UpdateRoute(updateRouteOptionsModelNew)
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
			It(`Invoke UpdateRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

				// Construct an instance of the UpdateRouteOptions model
				updateRouteOptionsModel := new(logsrouterv3.UpdateRouteOptions)
				updateRouteOptionsModel.ID = core.StringPtr("testString")
				updateRouteOptionsModel.Name = core.StringPtr("my-route")
				updateRouteOptionsModel.Rules = []logsrouterv3.RulePrototype{*rulePrototypeModel}
				updateRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.UpdateRoute(updateRouteOptionsModel)
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
	Describe(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {
		deleteRoutePath := "/routes/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRoutePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRoute successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := logsRouterService.DeleteRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(logsrouterv3.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = logsRouterService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRoute with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(logsrouterv3.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := logsRouterService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRouteOptions model with no property values
				deleteRouteOptionsModelNew := new(logsrouterv3.DeleteRouteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = logsRouterService.DeleteRoute(deleteRouteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions) - Operation response error`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSettings with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(logsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
		getSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false, "api_version": 1}`)
				}))
			})
			It(`Invoke GetSettings successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(logsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.GetSettingsWithContext(ctx, getSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false, "api_version": 1}`)
				}))
			})
			It(`Invoke GetSettings successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.GetSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(logsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.GetSettings(getSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSettings with error: Operation request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(logsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.GetSettings(getSettingsOptionsModel)
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
			It(`Invoke GetSettings successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := new(logsrouterv3.GetSettingsOptions)
				getSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.GetSettings(getSettingsOptionsModel)
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
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions) - Operation response error`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateSettings with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(logsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
		updateSettingsPath := "/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false, "api_version": 1}`)
				}))
			})
			It(`Invoke UpdateSettings successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(logsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.UpdateSettingsWithContext(ctx, updateSettingsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateSettingsPath))
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
					fmt.Fprintf(res, "%s", `{"default_targets": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "crn": "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::", "name": "a-lr-target-us-south", "target_type": "cloud_logs"}], "permitted_target_regions": ["us-south"], "primary_metadata_region": "us-south", "backup_metadata_region": "us-east", "private_api_endpoint_only": false, "api_version": 1}`)
				}))
			})
			It(`Invoke UpdateSettings successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.UpdateSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(logsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.UpdateSettings(updateSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateSettings with error: Operation request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(logsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.UpdateSettings(updateSettingsOptionsModel)
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
			It(`Invoke UpdateSettings successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := new(logsrouterv3.UpdateSettingsOptions)
				updateSettingsOptionsModel.DefaultTargets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				updateSettingsOptionsModel.PermittedTargetRegions = []string{"us-south"}
				updateSettingsOptionsModel.PrimaryMetadataRegion = core.StringPtr("us-south")
				updateSettingsOptionsModel.BackupMetadataRegion = core.StringPtr("us-east")
				updateSettingsOptionsModel.PrivateAPIEndpointOnly = core.BoolPtr(false)
				updateSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.UpdateSettings(updateSettingsOptionsModel)
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
	Describe(`QueryDestinations(queryDestinationsOptions *QueryDestinationsOptions) - Operation response error`, func() {
		queryDestinationsPath := "/destinations/query"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryDestinationsPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke QueryDestinations with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsModel := new(logsrouterv3.QueryDestinationsOptions)
				queryDestinationsOptionsModel.Crns = []logsrouterv3.CRNPrototype{*crnPrototypeModel}
				queryDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`QueryDestinations(queryDestinationsOptions *QueryDestinationsOptions)`, func() {
		queryDestinationsPath := "/destinations/query"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(queryDestinationsPath))
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
					fmt.Fprintf(res, "%s", `{"destinations": {"mapKey": [{"account": "d26e70b9a57f4388a68b1e03888e82a9", "action": "send", "associated_targets": [{"crn": "crn:v1:bluemix:public:logs-router:us-south:a/d26e70b9a57f4388a68b1e03888e82a9::target:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7", "id": "2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7", "name": "my-logs-target", "association_reasons": [{"type": "route", "name": "my-dallas-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/d26e70b9a57f4388a68b1e03888e82a9::route:37bf4abf-4479-4992-b3c8-dbbab387e445"}]}], "crn": "crn:v1:bluemix:public:logs:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::", "region": "us-south", "service_name": "logs"}]}}`)
				}))
			})
			It(`Invoke QueryDestinations successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsModel := new(logsrouterv3.QueryDestinationsOptions)
				queryDestinationsOptionsModel.Crns = []logsrouterv3.CRNPrototype{*crnPrototypeModel}
				queryDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.QueryDestinationsWithContext(ctx, queryDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.QueryDestinationsWithContext(ctx, queryDestinationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(queryDestinationsPath))
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
					fmt.Fprintf(res, "%s", `{"destinations": {"mapKey": [{"account": "d26e70b9a57f4388a68b1e03888e82a9", "action": "send", "associated_targets": [{"crn": "crn:v1:bluemix:public:logs-router:us-south:a/d26e70b9a57f4388a68b1e03888e82a9::target:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7", "id": "2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7", "name": "my-logs-target", "association_reasons": [{"type": "route", "name": "my-dallas-route", "crn": "crn:v1:bluemix:public:logs-router:global:a/d26e70b9a57f4388a68b1e03888e82a9::route:37bf4abf-4479-4992-b3c8-dbbab387e445"}]}], "crn": "crn:v1:bluemix:public:logs:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::", "region": "us-south", "service_name": "logs"}]}}`)
				}))
			})
			It(`Invoke QueryDestinations successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.QueryDestinations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsModel := new(logsrouterv3.QueryDestinationsOptions)
				queryDestinationsOptionsModel.Crns = []logsrouterv3.CRNPrototype{*crnPrototypeModel}
				queryDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke QueryDestinations with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsModel := new(logsrouterv3.QueryDestinationsOptions)
				queryDestinationsOptionsModel.Crns = []logsrouterv3.CRNPrototype{*crnPrototypeModel}
				queryDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the QueryDestinationsOptions model with no property values
				queryDestinationsOptionsModelNew := new(logsrouterv3.QueryDestinationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.QueryDestinations(queryDestinationsOptionsModelNew)
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
			It(`Invoke QueryDestinations successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsModel := new(logsrouterv3.QueryDestinationsOptions)
				queryDestinationsOptionsModel.Crns = []logsrouterv3.CRNPrototype{*crnPrototypeModel}
				queryDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.QueryDestinations(queryDestinationsOptionsModel)
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
	Describe(`MigrateActions(migrateActionsOptions *MigrateActionsOptions) - Operation response error`, func() {
		migrateActionsPath := "/migrate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(migrateActionsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["action"]).To(Equal([]string{"complete"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke MigrateActions with error: Operation response processing error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the MigrateActionsOptions model
				migrateActionsOptionsModel := new(logsrouterv3.MigrateActionsOptions)
				migrateActionsOptionsModel.Action = core.StringPtr("complete")
				migrateActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := logsRouterService.MigrateActions(migrateActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				logsRouterService.EnableRetries(0, 0)
				result, response, operationErr = logsRouterService.MigrateActions(migrateActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`MigrateActions(migrateActionsOptions *MigrateActionsOptions)`, func() {
		migrateActionsPath := "/migrate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(migrateActionsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["action"]).To(Equal([]string{"complete"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_version": 3, "message": "Message"}`)
				}))
			})
			It(`Invoke MigrateActions successfully with retries`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())
				logsRouterService.EnableRetries(0, 0)

				// Construct an instance of the MigrateActionsOptions model
				migrateActionsOptionsModel := new(logsrouterv3.MigrateActionsOptions)
				migrateActionsOptionsModel.Action = core.StringPtr("complete")
				migrateActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := logsRouterService.MigrateActionsWithContext(ctx, migrateActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				logsRouterService.DisableRetries()
				result, response, operationErr := logsRouterService.MigrateActions(migrateActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = logsRouterService.MigrateActionsWithContext(ctx, migrateActionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(migrateActionsPath))
					Expect(req.Method).To(Equal("POST"))

					Expect(req.URL.Query()["action"]).To(Equal([]string{"complete"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_version": 3, "message": "Message"}`)
				}))
			})
			It(`Invoke MigrateActions successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := logsRouterService.MigrateActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the MigrateActionsOptions model
				migrateActionsOptionsModel := new(logsrouterv3.MigrateActionsOptions)
				migrateActionsOptionsModel.Action = core.StringPtr("complete")
				migrateActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = logsRouterService.MigrateActions(migrateActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke MigrateActions with error: Operation validation and request error`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the MigrateActionsOptions model
				migrateActionsOptionsModel := new(logsrouterv3.MigrateActionsOptions)
				migrateActionsOptionsModel.Action = core.StringPtr("complete")
				migrateActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := logsRouterService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := logsRouterService.MigrateActions(migrateActionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the MigrateActionsOptions model with no property values
				migrateActionsOptionsModelNew := new(logsrouterv3.MigrateActionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = logsRouterService.MigrateActions(migrateActionsOptionsModelNew)
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
			It(`Invoke MigrateActions successfully`, func() {
				logsRouterService, serviceErr := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(logsRouterService).ToNot(BeNil())

				// Construct an instance of the MigrateActionsOptions model
				migrateActionsOptionsModel := new(logsrouterv3.MigrateActionsOptions)
				migrateActionsOptionsModel.Action = core.StringPtr("complete")
				migrateActionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := logsRouterService.MigrateActions(migrateActionsOptionsModel)
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
			logsRouterService, _ := logsrouterv3.NewLogsRouterV3(&logsrouterv3.LogsRouterV3Options{
				URL:           "http://logsrouterv3modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCRNPrototype successfully`, func() {
				crn := "crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::"
				_model, err := logsRouterService.NewCRNPrototype(crn)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewCreateRouteOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				Expect(inclusionFilterPrototypeModel).ToNot(BeNil())
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}
				Expect(inclusionFilterPrototypeModel.Operand).To(Equal(core.StringPtr("location")))
				Expect(inclusionFilterPrototypeModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(inclusionFilterPrototypeModel.Values).To(Equal([]string{"us-south"}))

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}
				Expect(rulePrototypeModel.Action).To(Equal(core.StringPtr("send")))
				Expect(rulePrototypeModel.Targets).To(Equal([]logsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(rulePrototypeModel.InclusionFilters).To(Equal([]logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}))

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsName := "my-route"
				createRouteOptionsRules := []logsrouterv3.RulePrototype{}
				createRouteOptionsModel := logsRouterService.NewCreateRouteOptions(createRouteOptionsName, createRouteOptionsRules)
				createRouteOptionsModel.SetName("my-route")
				createRouteOptionsModel.SetRules([]logsrouterv3.RulePrototype{*rulePrototypeModel})
				createRouteOptionsModel.SetManagedBy("enterprise")
				createRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRouteOptionsModel).ToNot(BeNil())
				Expect(createRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(createRouteOptionsModel.Rules).To(Equal([]logsrouterv3.RulePrototype{*rulePrototypeModel}))
				Expect(createRouteOptionsModel.ManagedBy).To(Equal(core.StringPtr("enterprise")))
				Expect(createRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsName := "my-lr-target"
				createTargetOptionsDestinationCRN := "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
				createTargetOptionsModel := logsRouterService.NewCreateTargetOptions(createTargetOptionsName, createTargetOptionsDestinationCRN)
				createTargetOptionsModel.SetName("my-lr-target")
				createTargetOptionsModel.SetDestinationCRN("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				createTargetOptionsModel.SetRegion("us-south")
				createTargetOptionsModel.SetManagedBy("enterprise")
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.Name).To(Equal(core.StringPtr("my-lr-target")))
				Expect(createTargetOptionsModel.DestinationCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")))
				Expect(createTargetOptionsModel.Region).To(Equal(core.StringPtr("us-south")))
				Expect(createTargetOptionsModel.ManagedBy).To(Equal(core.StringPtr("enterprise")))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRouteOptions successfully`, func() {
				// Construct an instance of the DeleteRouteOptions model
				id := "testString"
				deleteRouteOptionsModel := logsRouterService.NewDeleteRouteOptions(id)
				deleteRouteOptionsModel.SetID("testString")
				deleteRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRouteOptionsModel).ToNot(BeNil())
				Expect(deleteRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				id := "testString"
				deleteTargetOptionsModel := logsRouterService.NewDeleteTargetOptions(id)
				deleteTargetOptionsModel.SetID("testString")
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRouteOptions successfully`, func() {
				// Construct an instance of the GetRouteOptions model
				id := "testString"
				getRouteOptionsModel := logsRouterService.NewGetRouteOptions(id)
				getRouteOptionsModel.SetID("testString")
				getRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRouteOptionsModel).ToNot(BeNil())
				Expect(getRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSettingsOptions successfully`, func() {
				// Construct an instance of the GetSettingsOptions model
				getSettingsOptionsModel := logsRouterService.NewGetSettingsOptions()
				getSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSettingsOptionsModel).ToNot(BeNil())
				Expect(getSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTargetOptions successfully`, func() {
				// Construct an instance of the GetTargetOptions model
				id := "testString"
				getTargetOptionsModel := logsRouterService.NewGetTargetOptions(id)
				getTargetOptionsModel.SetID("testString")
				getTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTargetOptionsModel).ToNot(BeNil())
				Expect(getTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInclusionFilterPrototype successfully`, func() {
				operand := "location"
				operator := "is"
				values := []string{"us-south"}
				_model, err := logsRouterService.NewInclusionFilterPrototype(operand, operator, values)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListRoutesOptions successfully`, func() {
				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := logsRouterService.NewListRoutesOptions()
				listRoutesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRoutesOptionsModel).ToNot(BeNil())
				Expect(listRoutesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTargetsOptions successfully`, func() {
				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := logsRouterService.NewListTargetsOptions()
				listTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTargetsOptionsModel).ToNot(BeNil())
				Expect(listTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewMigrateActionsOptions successfully`, func() {
				// Construct an instance of the MigrateActionsOptions model
				action := "complete"
				migrateActionsOptionsModel := logsRouterService.NewMigrateActionsOptions(action)
				migrateActionsOptionsModel.SetAction("complete")
				migrateActionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(migrateActionsOptionsModel).ToNot(BeNil())
				Expect(migrateActionsOptionsModel.Action).To(Equal(core.StringPtr("complete")))
				Expect(migrateActionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewQueryDestinationsOptions successfully`, func() {
				// Construct an instance of the CRNPrototype model
				crnPrototypeModel := new(logsrouterv3.CRNPrototype)
				Expect(crnPrototypeModel).ToNot(BeNil())
				crnPrototypeModel.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")
				Expect(crnPrototypeModel.CRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")))

				// Construct an instance of the QueryDestinationsOptions model
				queryDestinationsOptionsCrns := []logsrouterv3.CRNPrototype{}
				queryDestinationsOptionsModel := logsRouterService.NewQueryDestinationsOptions(queryDestinationsOptionsCrns)
				queryDestinationsOptionsModel.SetCrns([]logsrouterv3.CRNPrototype{*crnPrototypeModel})
				queryDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(queryDestinationsOptionsModel).ToNot(BeNil())
				Expect(queryDestinationsOptionsModel.Crns).To(Equal([]logsrouterv3.CRNPrototype{*crnPrototypeModel}))
				Expect(queryDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRulePrototype successfully`, func() {
				targets := []logsrouterv3.TargetIdentity{}
				inclusionFilters := []logsrouterv3.InclusionFilterPrototype{}
				_model, err := logsRouterService.NewRulePrototype(targets, inclusionFilters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTargetIdentity successfully`, func() {
				id := "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
				_model, err := logsRouterService.NewTargetIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateRouteOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the InclusionFilterPrototype model
				inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
				Expect(inclusionFilterPrototypeModel).ToNot(BeNil())
				inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
				inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
				inclusionFilterPrototypeModel.Values = []string{"us-south"}
				Expect(inclusionFilterPrototypeModel.Operand).To(Equal(core.StringPtr("location")))
				Expect(inclusionFilterPrototypeModel.Operator).To(Equal(core.StringPtr("is")))
				Expect(inclusionFilterPrototypeModel.Values).To(Equal([]string{"us-south"}))

				// Construct an instance of the RulePrototype model
				rulePrototypeModel := new(logsrouterv3.RulePrototype)
				Expect(rulePrototypeModel).ToNot(BeNil())
				rulePrototypeModel.Action = core.StringPtr("send")
				rulePrototypeModel.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
				rulePrototypeModel.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}
				Expect(rulePrototypeModel.Action).To(Equal(core.StringPtr("send")))
				Expect(rulePrototypeModel.Targets).To(Equal([]logsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(rulePrototypeModel.InclusionFilters).To(Equal([]logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}))

				// Construct an instance of the UpdateRouteOptions model
				id := "testString"
				updateRouteOptionsModel := logsRouterService.NewUpdateRouteOptions(id)
				updateRouteOptionsModel.SetID("testString")
				updateRouteOptionsModel.SetName("my-route")
				updateRouteOptionsModel.SetRules([]logsrouterv3.RulePrototype{*rulePrototypeModel})
				updateRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRouteOptionsModel).ToNot(BeNil())
				Expect(updateRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(updateRouteOptionsModel.Rules).To(Equal([]logsrouterv3.RulePrototype{*rulePrototypeModel}))
				Expect(updateRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateSettingsOptions successfully`, func() {
				// Construct an instance of the TargetIdentity model
				targetIdentityModel := new(logsrouterv3.TargetIdentity)
				Expect(targetIdentityModel).ToNot(BeNil())
				targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
				Expect(targetIdentityModel.ID).To(Equal(core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")))

				// Construct an instance of the UpdateSettingsOptions model
				updateSettingsOptionsModel := logsRouterService.NewUpdateSettingsOptions()
				updateSettingsOptionsModel.SetDefaultTargets([]logsrouterv3.TargetIdentity{*targetIdentityModel})
				updateSettingsOptionsModel.SetPermittedTargetRegions([]string{"us-south"})
				updateSettingsOptionsModel.SetPrimaryMetadataRegion("us-south")
				updateSettingsOptionsModel.SetBackupMetadataRegion("us-east")
				updateSettingsOptionsModel.SetPrivateAPIEndpointOnly(false)
				updateSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateSettingsOptionsModel).ToNot(BeNil())
				Expect(updateSettingsOptionsModel.DefaultTargets).To(Equal([]logsrouterv3.TargetIdentity{*targetIdentityModel}))
				Expect(updateSettingsOptionsModel.PermittedTargetRegions).To(Equal([]string{"us-south"}))
				Expect(updateSettingsOptionsModel.PrimaryMetadataRegion).To(Equal(core.StringPtr("us-south")))
				Expect(updateSettingsOptionsModel.BackupMetadataRegion).To(Equal(core.StringPtr("us-east")))
				Expect(updateSettingsOptionsModel.PrivateAPIEndpointOnly).To(Equal(core.BoolPtr(false)))
				Expect(updateSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateTargetOptions successfully`, func() {
				// Construct an instance of the UpdateTargetOptions model
				id := "testString"
				updateTargetOptionsModel := logsRouterService.NewUpdateTargetOptions(id)
				updateTargetOptionsModel.SetID("testString")
				updateTargetOptionsModel.SetName("my-lr-target")
				updateTargetOptionsModel.SetDestinationCRN("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
				updateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateTargetOptionsModel).ToNot(BeNil())
				Expect(updateTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateTargetOptionsModel.Name).To(Equal(core.StringPtr("my-lr-target")))
				Expect(updateTargetOptionsModel.DestinationCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")))
				Expect(updateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalCRNPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsrouterv3.CRNPrototype)
			model.CRN = core.StringPtr("crn:v1:bluemix:public:codeengine:us-south:a/d26e70b9a57f4388a68b1e03888e82a9:2c6a54f8-4afe-4b8b-b55a-9d31a8e890c7::")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsrouterv3.CRNPrototype
			err = logsrouterv3.UnmarshalCRNPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalInclusionFilterPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsrouterv3.InclusionFilterPrototype)
			model.Operand = core.StringPtr("location")
			model.Operator = core.StringPtr("is")
			model.Values = []string{"us-south"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsrouterv3.InclusionFilterPrototype
			err = logsrouterv3.UnmarshalInclusionFilterPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalRulePrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(logsrouterv3.RulePrototype)
			model.Action = core.StringPtr("send")
			model.Targets = nil
			model.InclusionFilters = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsrouterv3.RulePrototype
			err = logsrouterv3.UnmarshalRulePrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTargetIdentity successfully`, func() {
			// Construct an instance of the model.
			model := new(logsrouterv3.TargetIdentity)
			model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *logsrouterv3.TargetIdentity
			err = logsrouterv3.UnmarshalTargetIdentity(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("VGhpcyBpcyBhIHRlc3Qgb2YgdGhlIGVtZXJnZW5jeSBicm9hZGNhc3Qgc3lzdGVt")
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

func CreateMockByteArray(encodedString string) *[]byte {
	ba, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err)
	}
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
