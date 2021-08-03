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

package atrackerv1_test

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
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`AtrackerV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(atrackerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(atrackerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
				URL: "https://atrackerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(atrackerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_URL": "https://atrackerv1/api",
				"ATRACKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1UsingExternalConfig(&atrackerv1.AtrackerV1Options{
				})
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1UsingExternalConfig(&atrackerv1.AtrackerV1Options{
					URL: "https://testService/api",
				})
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1UsingExternalConfig(&atrackerv1.AtrackerV1Options{
				})
				err := atrackerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := atrackerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != atrackerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(atrackerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(atrackerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_URL": "https://atrackerv1/api",
				"ATRACKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			atrackerService, serviceErr := atrackerv1.NewAtrackerV1UsingExternalConfig(&atrackerv1.AtrackerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(atrackerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ATRACKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			atrackerService, serviceErr := atrackerv1.NewAtrackerV1UsingExternalConfig(&atrackerv1.AtrackerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(atrackerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = atrackerv1.GetServiceURLForRegion("us-south")
			Expect(url).To(Equal("https://us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv1.GetServiceURLForRegion("private.us-south")
			Expect(url).To(Equal("https://private.us-south.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv1.GetServiceURLForRegion("us-east")
			Expect(url).To(Equal("https://us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv1.GetServiceURLForRegion("private.us-east")
			Expect(url).To(Equal("https://private.us-east.atracker.cloud.ibm.com"))
			Expect(err).To(BeNil())

			url, err = atrackerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateTarget(createTargetOptions *CreateTargetOptions) - Operation response error`, func() {
		createTargetPath := "/api/v1/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createTargetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv1.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModel)
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
		createTargetPath := "/api/v1/targets"
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv1.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.CreateTargetWithContext(ctx, createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.CreateTargetWithContext(ctx, createTargetOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.CreateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv1.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv1.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateTargetOptions model with no property values
				createTargetOptionsModelNew := new(atrackerv1.CreateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.CreateTarget(createTargetOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsModel := new(atrackerv1.CreateTargetOptions)
				createTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				createTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				createTargetOptionsModel.CosEndpoint = cosEndpointModel
				createTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.CreateTarget(createTargetOptionsModel)
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
		listTargetsPath := "/api/v1/targets"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listTargetsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListTargets with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv1.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ListTargets(listTargetsOptionsModel)
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
		listTargetsPath := "/api/v1/targets"
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
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv1.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ListTargetsWithContext(ctx, listTargetsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"targets": [{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListTargets successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ListTargets(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv1.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ListTargets(listTargetsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListTargets with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv1.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := new(atrackerv1.ListTargetsOptions)
				listTargetsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ListTargets(listTargetsOptionsModel)
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
		getTargetPath := "/api/v1/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getTargetPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv1.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModel)
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
		getTargetPath := "/api/v1/targets/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv1.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetTargetWithContext(ctx, getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetTargetWithContext(ctx, getTargetOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv1.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv1.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetTargetOptions model with no property values
				getTargetOptionsModelNew := new(atrackerv1.GetTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.GetTarget(getTargetOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetTargetOptions model
				getTargetOptionsModel := new(atrackerv1.GetTargetOptions)
				getTargetOptionsModel.ID = core.StringPtr("testString")
				getTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetTarget(getTargetOptionsModel)
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
	Describe(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) - Operation response error`, func() {
		replaceTargetPath := "/api/v1/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv1.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions)`, func() {
		replaceTargetPath := "/api/v1/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ReplaceTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv1.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ReplaceTargetWithContext(ctx, replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ReplaceTargetWithContext(ctx, replaceTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceTargetPath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ReplaceTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ReplaceTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv1.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv1.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceTargetOptions model with no property values
				replaceTargetOptionsModelNew := new(atrackerv1.ReplaceTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ReplaceTarget(replaceTargetOptionsModelNew)
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
			It(`Invoke ReplaceTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")

				// Construct an instance of the ReplaceTargetOptions model
				replaceTargetOptionsModel := new(atrackerv1.ReplaceTargetOptions)
				replaceTargetOptionsModel.ID = core.StringPtr("testString")
				replaceTargetOptionsModel.Name = core.StringPtr("my-cos-target")
				replaceTargetOptionsModel.TargetType = core.StringPtr("cloud_object_storage")
				replaceTargetOptionsModel.CosEndpoint = cosEndpointModel
				replaceTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ReplaceTarget(replaceTargetOptionsModel)
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
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions) - Operation response error`, func() {
		deleteTargetPath := "/api/v1/targets/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv1.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
		deleteTargetPath := "/api/v1/targets/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_code": 10, "trace": "Trace", "warnings": [{"code": "Code", "message": "Message"}]}`)
				}))
			})
			It(`Invoke DeleteTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv1.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.DeleteTargetWithContext(ctx, deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.DeleteTargetWithContext(ctx, deleteTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(deleteTargetPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"status_code": 10, "trace": "Trace", "warnings": [{"code": "Code", "message": "Message"}]}`)
				}))
			})
			It(`Invoke DeleteTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.DeleteTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv1.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv1.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteTargetOptions model with no property values
				deleteTargetOptionsModelNew := new(atrackerv1.DeleteTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.DeleteTarget(deleteTargetOptionsModelNew)
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
			It(`Invoke DeleteTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteTargetOptions model
				deleteTargetOptionsModel := new(atrackerv1.DeleteTargetOptions)
				deleteTargetOptionsModel.ID = core.StringPtr("testString")
				deleteTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.DeleteTarget(deleteTargetOptionsModel)
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
	Describe(`ValidateTarget(validateTargetOptions *ValidateTargetOptions) - Operation response error`, func() {
		validateTargetPath := "/api/v1/targets/testString/validate"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ValidateTarget with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv1.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ValidateTarget(validateTargetOptions *ValidateTargetOptions)`, func() {
		validateTargetPath := "/api/v1/targets/testString/validate"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ValidateTarget successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv1.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ValidateTargetWithContext(ctx, validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ValidateTargetWithContext(ctx, validateTargetOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(validateTargetPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "name": "a-cos-target-us-south", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6", "target_type": "cloud_object_storage", "encrypt_key": "REDACTED", "cos_endpoint": {"endpoint": "s3.private.us-east.cloud-object-storage.appdomain.cloud", "target_crn": "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::", "bucket": "my-atracker-bucket", "api_key": "xxxxxxxxxxxxxx"}, "cos_write_status": {"status": "success", "last_failure": "2021-05-18T20:15:12.353Z", "reason_for_last_failure": "Provided API key could not be found"}, "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ValidateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ValidateTarget(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv1.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ValidateTarget with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv1.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ValidateTargetOptions model with no property values
				validateTargetOptionsModelNew := new(atrackerv1.ValidateTargetOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ValidateTarget(validateTargetOptionsModelNew)
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
			It(`Invoke ValidateTarget successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ValidateTargetOptions model
				validateTargetOptionsModel := new(atrackerv1.ValidateTargetOptions)
				validateTargetOptionsModel.ID = core.StringPtr("testString")
				validateTargetOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ValidateTarget(validateTargetOptionsModel)
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
	Describe(`CreateRoute(createRouteOptions *CreateRouteOptions) - Operation response error`, func() {
		createRoutePath := "/api/v1/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRoutePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRoute with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv1.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				createRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModel)
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
		createRoutePath := "/api/v1/routes"
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv1.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				createRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.CreateRouteWithContext(ctx, createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.CreateRouteWithContext(ctx, createRouteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke CreateRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.CreateRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv1.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				createRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv1.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				createRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRouteOptions model with no property values
				createRouteOptionsModelNew := new(atrackerv1.CreateRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.CreateRoute(createRouteOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsModel := new(atrackerv1.CreateRouteOptions)
				createRouteOptionsModel.Name = core.StringPtr("my-route")
				createRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				createRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				createRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.CreateRoute(createRouteOptionsModel)
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
		listRoutesPath := "/api/v1/routes"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRoutesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoutes with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv1.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ListRoutes(listRoutesOptionsModel)
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
		listRoutesPath := "/api/v1/routes"
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
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv1.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ListRoutesWithContext(ctx, listRoutesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"routes": [{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}]}`)
				}))
			})
			It(`Invoke ListRoutes successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ListRoutes(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv1.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ListRoutes(listRoutesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoutes with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv1.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := new(atrackerv1.ListRoutesOptions)
				listRoutesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ListRoutes(listRoutesOptionsModel)
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
		getRoutePath := "/api/v1/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRoutePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRoute with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv1.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModel)
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
		getRoutePath := "/api/v1/routes/testString"
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv1.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetRouteWithContext(ctx, getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetRouteWithContext(ctx, getRouteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke GetRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv1.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv1.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRouteOptions model with no property values
				getRouteOptionsModelNew := new(atrackerv1.GetRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.GetRoute(getRouteOptionsModelNew)
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetRouteOptions model
				getRouteOptionsModel := new(atrackerv1.GetRouteOptions)
				getRouteOptionsModel.ID = core.StringPtr("testString")
				getRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetRoute(getRouteOptionsModel)
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
	Describe(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions) - Operation response error`, func() {
		replaceRoutePath := "/api/v1/routes/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRoute with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv1.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				replaceRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {
		replaceRoutePath := "/api/v1/routes/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ReplaceRoute successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv1.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				replaceRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.ReplaceRouteWithContext(ctx, replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.ReplaceRouteWithContext(ctx, replaceRouteOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRoutePath))
					Expect(req.Method).To(Equal("PUT"))

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
					fmt.Fprintf(res, "%s", `{"id": "c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "name": "my-route", "crn": "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4", "version": 0, "receive_global_events": false, "rules": [{"target_ids": ["c3af557f-fb0e-4476-85c3-0889e7fe7bc4"]}], "created": "2021-05-18T20:15:12.353Z", "updated": "2021-05-18T20:15:12.353Z"}`)
				}))
			})
			It(`Invoke ReplaceRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.ReplaceRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv1.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				replaceRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv1.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				replaceRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRouteOptions model with no property values
				replaceRouteOptionsModelNew := new(atrackerv1.ReplaceRouteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = atrackerService.ReplaceRoute(replaceRouteOptionsModelNew)
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
			It(`Invoke ReplaceRoute successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}

				// Construct an instance of the ReplaceRouteOptions model
				replaceRouteOptionsModel := new(atrackerv1.ReplaceRouteOptions)
				replaceRouteOptionsModel.ID = core.StringPtr("testString")
				replaceRouteOptionsModel.Name = core.StringPtr("my-route")
				replaceRouteOptionsModel.ReceiveGlobalEvents = core.BoolPtr(false)
				replaceRouteOptionsModel.Rules = []atrackerv1.Rule{*ruleModel}
				replaceRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.ReplaceRoute(replaceRouteOptionsModel)
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
		deleteRoutePath := "/api/v1/routes/testString"
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
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := atrackerService.DeleteRoute(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(atrackerv1.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = atrackerService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRoute with error: Operation validation and request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the DeleteRouteOptions model
				deleteRouteOptionsModel := new(atrackerv1.DeleteRouteOptions)
				deleteRouteOptionsModel.ID = core.StringPtr("testString")
				deleteRouteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := atrackerService.DeleteRoute(deleteRouteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRouteOptions model with no property values
				deleteRouteOptionsModelNew := new(atrackerv1.DeleteRouteOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = atrackerService.DeleteRoute(deleteRouteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions) - Operation response error`, func() {
		getEndpointsPath := "/api/v1/endpoints"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetEndpoints with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(atrackerv1.GetEndpointsOptions)
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetEndpoints(getEndpointsOptions *GetEndpointsOptions)`, func() {
		getEndpointsPath := "/api/v1/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_endpoint": {"public_url": "https://us-south.atracker.cloud.ibm.com", "public_enabled": true, "private_url": "https://private.us-south.atracker.cloud.ibm.com", "private_enabled": true}}`)
				}))
			})
			It(`Invoke GetEndpoints successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(atrackerv1.GetEndpointsOptions)
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.GetEndpointsWithContext(ctx, getEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.GetEndpointsWithContext(ctx, getEndpointsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getEndpointsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"api_endpoint": {"public_url": "https://us-south.atracker.cloud.ibm.com", "public_enabled": true, "private_url": "https://private.us-south.atracker.cloud.ibm.com", "private_enabled": true}}`)
				}))
			})
			It(`Invoke GetEndpoints successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.GetEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(atrackerv1.GetEndpointsOptions)
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.GetEndpoints(getEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetEndpoints with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(atrackerv1.GetEndpointsOptions)
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.GetEndpoints(getEndpointsOptionsModel)
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
			It(`Invoke GetEndpoints successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := new(atrackerv1.GetEndpointsOptions)
				getEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.GetEndpoints(getEndpointsOptionsModel)
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
	Describe(`PatchEndpoints(patchEndpointsOptions *PatchEndpointsOptions) - Operation response error`, func() {
		patchEndpointsPath := "/api/v1/endpoints"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchEndpointsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PatchEndpoints with error: Operation response processing error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := new(atrackerv1.PatchEndpointsOptions)
				patchEndpointsOptionsModel.APIEndpoint = endpointsRequestAPIEndpointModel
				patchEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				atrackerService.EnableRetries(0, 0)
				result, response, operationErr = atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchEndpoints(patchEndpointsOptions *PatchEndpointsOptions)`, func() {
		patchEndpointsPath := "/api/v1/endpoints"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchEndpointsPath))
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
					fmt.Fprintf(res, "%s", `{"api_endpoint": {"public_url": "https://us-south.atracker.cloud.ibm.com", "public_enabled": true, "private_url": "https://private.us-south.atracker.cloud.ibm.com", "private_enabled": true}}`)
				}))
			})
			It(`Invoke PatchEndpoints successfully with retries`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())
				atrackerService.EnableRetries(0, 0)

				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := new(atrackerv1.PatchEndpointsOptions)
				patchEndpointsOptionsModel.APIEndpoint = endpointsRequestAPIEndpointModel
				patchEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := atrackerService.PatchEndpointsWithContext(ctx, patchEndpointsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				atrackerService.DisableRetries()
				result, response, operationErr := atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = atrackerService.PatchEndpointsWithContext(ctx, patchEndpointsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(patchEndpointsPath))
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
					fmt.Fprintf(res, "%s", `{"api_endpoint": {"public_url": "https://us-south.atracker.cloud.ibm.com", "public_enabled": true, "private_url": "https://private.us-south.atracker.cloud.ibm.com", "private_enabled": true}}`)
				}))
			})
			It(`Invoke PatchEndpoints successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := atrackerService.PatchEndpoints(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := new(atrackerv1.PatchEndpointsOptions)
				patchEndpointsOptionsModel.APIEndpoint = endpointsRequestAPIEndpointModel
				patchEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PatchEndpoints with error: Operation request error`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := new(atrackerv1.PatchEndpointsOptions)
				patchEndpointsOptionsModel.APIEndpoint = endpointsRequestAPIEndpointModel
				patchEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := atrackerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
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
			It(`Invoke PatchEndpoints successfully`, func() {
				atrackerService, serviceErr := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(atrackerService).ToNot(BeNil())

				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := new(atrackerv1.PatchEndpointsOptions)
				patchEndpointsOptionsModel.APIEndpoint = endpointsRequestAPIEndpointModel
				patchEndpointsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := atrackerService.PatchEndpoints(patchEndpointsOptionsModel)
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
			atrackerService, _ := atrackerv1.NewAtrackerV1(&atrackerv1.AtrackerV1Options{
				URL:           "http://atrackerv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateRouteOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				Expect(ruleModel.TargetIds).To(Equal([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}))

				// Construct an instance of the CreateRouteOptions model
				createRouteOptionsName := "my-route"
				createRouteOptionsReceiveGlobalEvents := false
				createRouteOptionsRules := []atrackerv1.Rule{}
				createRouteOptionsModel := atrackerService.NewCreateRouteOptions(createRouteOptionsName, createRouteOptionsReceiveGlobalEvents, createRouteOptionsRules)
				createRouteOptionsModel.SetName("my-route")
				createRouteOptionsModel.SetReceiveGlobalEvents(false)
				createRouteOptionsModel.SetRules([]atrackerv1.Rule{*ruleModel})
				createRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRouteOptionsModel).ToNot(BeNil())
				Expect(createRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(createRouteOptionsModel.ReceiveGlobalEvents).To(Equal(core.BoolPtr(false)))
				Expect(createRouteOptionsModel.Rules).To(Equal([]atrackerv1.Rule{*ruleModel}))
				Expect(createRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateTargetOptions successfully`, func() {
				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				Expect(cosEndpointModel).ToNot(BeNil())
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				Expect(cosEndpointModel.Endpoint).To(Equal(core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")))
				Expect(cosEndpointModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(cosEndpointModel.Bucket).To(Equal(core.StringPtr("my-atracker-bucket")))
				Expect(cosEndpointModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))

				// Construct an instance of the CreateTargetOptions model
				createTargetOptionsName := "my-cos-target"
				createTargetOptionsTargetType := "cloud_object_storage"
				var createTargetOptionsCosEndpoint *atrackerv1.CosEndpoint = nil
				createTargetOptionsModel := atrackerService.NewCreateTargetOptions(createTargetOptionsName, createTargetOptionsTargetType, createTargetOptionsCosEndpoint)
				createTargetOptionsModel.SetName("my-cos-target")
				createTargetOptionsModel.SetTargetType("cloud_object_storage")
				createTargetOptionsModel.SetCosEndpoint(cosEndpointModel)
				createTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createTargetOptionsModel).ToNot(BeNil())
				Expect(createTargetOptionsModel.Name).To(Equal(core.StringPtr("my-cos-target")))
				Expect(createTargetOptionsModel.TargetType).To(Equal(core.StringPtr("cloud_object_storage")))
				Expect(createTargetOptionsModel.CosEndpoint).To(Equal(cosEndpointModel))
				Expect(createTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRouteOptions successfully`, func() {
				// Construct an instance of the DeleteRouteOptions model
				id := "testString"
				deleteRouteOptionsModel := atrackerService.NewDeleteRouteOptions(id)
				deleteRouteOptionsModel.SetID("testString")
				deleteRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRouteOptionsModel).ToNot(BeNil())
				Expect(deleteRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteTargetOptions successfully`, func() {
				// Construct an instance of the DeleteTargetOptions model
				id := "testString"
				deleteTargetOptionsModel := atrackerService.NewDeleteTargetOptions(id)
				deleteTargetOptionsModel.SetID("testString")
				deleteTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteTargetOptionsModel).ToNot(BeNil())
				Expect(deleteTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetEndpointsOptions successfully`, func() {
				// Construct an instance of the GetEndpointsOptions model
				getEndpointsOptionsModel := atrackerService.NewGetEndpointsOptions()
				getEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getEndpointsOptionsModel).ToNot(BeNil())
				Expect(getEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRouteOptions successfully`, func() {
				// Construct an instance of the GetRouteOptions model
				id := "testString"
				getRouteOptionsModel := atrackerService.NewGetRouteOptions(id)
				getRouteOptionsModel.SetID("testString")
				getRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRouteOptionsModel).ToNot(BeNil())
				Expect(getRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetTargetOptions successfully`, func() {
				// Construct an instance of the GetTargetOptions model
				id := "testString"
				getTargetOptionsModel := atrackerService.NewGetTargetOptions(id)
				getTargetOptionsModel.SetID("testString")
				getTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getTargetOptionsModel).ToNot(BeNil())
				Expect(getTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRoutesOptions successfully`, func() {
				// Construct an instance of the ListRoutesOptions model
				listRoutesOptionsModel := atrackerService.NewListRoutesOptions()
				listRoutesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRoutesOptionsModel).ToNot(BeNil())
				Expect(listRoutesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListTargetsOptions successfully`, func() {
				// Construct an instance of the ListTargetsOptions model
				listTargetsOptionsModel := atrackerService.NewListTargetsOptions()
				listTargetsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listTargetsOptionsModel).ToNot(BeNil())
				Expect(listTargetsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchEndpointsOptions successfully`, func() {
				// Construct an instance of the EndpointsRequestAPIEndpoint model
				endpointsRequestAPIEndpointModel := new(atrackerv1.EndpointsRequestAPIEndpoint)
				Expect(endpointsRequestAPIEndpointModel).ToNot(BeNil())
				endpointsRequestAPIEndpointModel.PublicEnabled = core.BoolPtr(true)
				Expect(endpointsRequestAPIEndpointModel.PublicEnabled).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PatchEndpointsOptions model
				patchEndpointsOptionsModel := atrackerService.NewPatchEndpointsOptions()
				patchEndpointsOptionsModel.SetAPIEndpoint(endpointsRequestAPIEndpointModel)
				patchEndpointsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchEndpointsOptionsModel).ToNot(BeNil())
				Expect(patchEndpointsOptionsModel.APIEndpoint).To(Equal(endpointsRequestAPIEndpointModel))
				Expect(patchEndpointsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceRouteOptions successfully`, func() {
				// Construct an instance of the Rule model
				ruleModel := new(atrackerv1.Rule)
				Expect(ruleModel).ToNot(BeNil())
				ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				Expect(ruleModel.TargetIds).To(Equal([]string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}))

				// Construct an instance of the ReplaceRouteOptions model
				id := "testString"
				replaceRouteOptionsName := "my-route"
				replaceRouteOptionsReceiveGlobalEvents := false
				replaceRouteOptionsRules := []atrackerv1.Rule{}
				replaceRouteOptionsModel := atrackerService.NewReplaceRouteOptions(id, replaceRouteOptionsName, replaceRouteOptionsReceiveGlobalEvents, replaceRouteOptionsRules)
				replaceRouteOptionsModel.SetID("testString")
				replaceRouteOptionsModel.SetName("my-route")
				replaceRouteOptionsModel.SetReceiveGlobalEvents(false)
				replaceRouteOptionsModel.SetRules([]atrackerv1.Rule{*ruleModel})
				replaceRouteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRouteOptionsModel).ToNot(BeNil())
				Expect(replaceRouteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRouteOptionsModel.Name).To(Equal(core.StringPtr("my-route")))
				Expect(replaceRouteOptionsModel.ReceiveGlobalEvents).To(Equal(core.BoolPtr(false)))
				Expect(replaceRouteOptionsModel.Rules).To(Equal([]atrackerv1.Rule{*ruleModel}))
				Expect(replaceRouteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceTargetOptions successfully`, func() {
				// Construct an instance of the CosEndpoint model
				cosEndpointModel := new(atrackerv1.CosEndpoint)
				Expect(cosEndpointModel).ToNot(BeNil())
				cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
				cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
				cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
				cosEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
				Expect(cosEndpointModel.Endpoint).To(Equal(core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")))
				Expect(cosEndpointModel.TargetCRN).To(Equal(core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")))
				Expect(cosEndpointModel.Bucket).To(Equal(core.StringPtr("my-atracker-bucket")))
				Expect(cosEndpointModel.APIKey).To(Equal(core.StringPtr("xxxxxxxxxxxxxx")))

				// Construct an instance of the ReplaceTargetOptions model
				id := "testString"
				replaceTargetOptionsName := "my-cos-target"
				replaceTargetOptionsTargetType := "cloud_object_storage"
				var replaceTargetOptionsCosEndpoint *atrackerv1.CosEndpoint = nil
				replaceTargetOptionsModel := atrackerService.NewReplaceTargetOptions(id, replaceTargetOptionsName, replaceTargetOptionsTargetType, replaceTargetOptionsCosEndpoint)
				replaceTargetOptionsModel.SetID("testString")
				replaceTargetOptionsModel.SetName("my-cos-target")
				replaceTargetOptionsModel.SetTargetType("cloud_object_storage")
				replaceTargetOptionsModel.SetCosEndpoint(cosEndpointModel)
				replaceTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceTargetOptionsModel).ToNot(BeNil())
				Expect(replaceTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceTargetOptionsModel.Name).To(Equal(core.StringPtr("my-cos-target")))
				Expect(replaceTargetOptionsModel.TargetType).To(Equal(core.StringPtr("cloud_object_storage")))
				Expect(replaceTargetOptionsModel.CosEndpoint).To(Equal(cosEndpointModel))
				Expect(replaceTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRule successfully`, func() {
				targetIds := []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
				_model, err := atrackerService.NewRule(targetIds)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewValidateTargetOptions successfully`, func() {
				// Construct an instance of the ValidateTargetOptions model
				id := "testString"
				validateTargetOptionsModel := atrackerService.NewValidateTargetOptions(id)
				validateTargetOptionsModel.SetID("testString")
				validateTargetOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(validateTargetOptionsModel).ToNot(BeNil())
				Expect(validateTargetOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(validateTargetOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCosEndpoint successfully`, func() {
				endpoint := "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				targetCRN := "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket := "my-atracker-bucket"
				apiKey := "xxxxxxxxxxxxxx"
				_model, err := atrackerService.NewCosEndpoint(endpoint, targetCRN, bucket, apiKey)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
