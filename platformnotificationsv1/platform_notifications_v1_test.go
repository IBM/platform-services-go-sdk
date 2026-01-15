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

package platformnotificationsv1_test

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
	"github.com/IBM/platform-services-go-sdk/platformnotificationsv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`PlatformNotificationsV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(platformNotificationsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(platformNotificationsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
				URL: "https://platformnotificationsv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(platformNotificationsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PLATFORM_NOTIFICATIONS_URL": "https://platformnotificationsv1/api",
				"PLATFORM_NOTIFICATIONS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(&platformnotificationsv1.PlatformNotificationsV1Options{
				})
				Expect(platformNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := platformNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != platformNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(platformNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(platformNotificationsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL: "https://testService/api",
				})
				Expect(platformNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := platformNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != platformNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(platformNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(platformNotificationsService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(&platformnotificationsv1.PlatformNotificationsV1Options{
				})
				err := platformNotificationsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := platformNotificationsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != platformNotificationsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(platformNotificationsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(platformNotificationsService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PLATFORM_NOTIFICATIONS_URL": "https://platformnotificationsv1/api",
				"PLATFORM_NOTIFICATIONS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(&platformnotificationsv1.PlatformNotificationsV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(platformNotificationsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"PLATFORM_NOTIFICATIONS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1UsingExternalConfig(&platformnotificationsv1.PlatformNotificationsV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(platformNotificationsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = platformnotificationsv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions) - Operation response error`, func() {
		listDistributionListDestinationsPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDistributionListDestinationsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDistributionListDestinations with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDistributionListDestinationsOptions model
				listDistributionListDestinationsOptionsModel := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				listDistributionListDestinationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDistributionListDestinations(listDistributionListDestinationsOptions *ListDistributionListDestinationsOptions)`, func() {
		listDistributionListDestinationsPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDistributionListDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destinations": [{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}]}`)
				}))
			})
			It(`Invoke ListDistributionListDestinations successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListDistributionListDestinationsOptions model
				listDistributionListDestinationsOptionsModel := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				listDistributionListDestinationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.ListDistributionListDestinationsWithContext(ctx, listDistributionListDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.ListDistributionListDestinationsWithContext(ctx, listDistributionListDestinationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listDistributionListDestinationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destinations": [{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}]}`)
				}))
			})
			It(`Invoke ListDistributionListDestinations successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.ListDistributionListDestinations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDistributionListDestinationsOptions model
				listDistributionListDestinationsOptionsModel := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				listDistributionListDestinationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDistributionListDestinations with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDistributionListDestinationsOptions model
				listDistributionListDestinationsOptionsModel := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				listDistributionListDestinationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDistributionListDestinationsOptions model with no property values
				listDistributionListDestinationsOptionsModelNew := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModelNew)
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
			It(`Invoke ListDistributionListDestinations successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListDistributionListDestinationsOptions model
				listDistributionListDestinationsOptionsModel := new(platformnotificationsv1.ListDistributionListDestinationsOptions)
				listDistributionListDestinationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.ListDistributionListDestinations(listDistributionListDestinationsOptionsModel)
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
	Describe(`CreateDistributionListDestination(createDistributionListDestinationOptions *CreateDistributionListDestinationOptions) - Operation response error`, func() {
		createDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDistributionListDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDistributionListDestination with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the CreateDistributionListDestinationOptions model
				createDistributionListDestinationOptionsModel := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				createDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.AddDestinationPrototype = addDestinationPrototypeModel
				createDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDistributionListDestination(createDistributionListDestinationOptions *CreateDistributionListDestinationOptions)`, func() {
		createDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDistributionListDestinationPath))
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
					fmt.Fprintf(res, "%s", `{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke CreateDistributionListDestination successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the CreateDistributionListDestinationOptions model
				createDistributionListDestinationOptionsModel := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				createDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.AddDestinationPrototype = addDestinationPrototypeModel
				createDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.CreateDistributionListDestinationWithContext(ctx, createDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.CreateDistributionListDestinationWithContext(ctx, createDistributionListDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDistributionListDestinationPath))
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
					fmt.Fprintf(res, "%s", `{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke CreateDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.CreateDistributionListDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the CreateDistributionListDestinationOptions model
				createDistributionListDestinationOptionsModel := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				createDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.AddDestinationPrototype = addDestinationPrototypeModel
				createDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDistributionListDestination with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the CreateDistributionListDestinationOptions model
				createDistributionListDestinationOptionsModel := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				createDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.AddDestinationPrototype = addDestinationPrototypeModel
				createDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDistributionListDestinationOptions model with no property values
				createDistributionListDestinationOptionsModelNew := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModelNew)
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
			It(`Invoke CreateDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")

				// Construct an instance of the CreateDistributionListDestinationOptions model
				createDistributionListDestinationOptionsModel := new(platformnotificationsv1.CreateDistributionListDestinationOptions)
				createDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.AddDestinationPrototype = addDestinationPrototypeModel
				createDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.CreateDistributionListDestination(createDistributionListDestinationOptionsModel)
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
	Describe(`GetDistributionListDestination(getDistributionListDestinationOptions *GetDistributionListDestinationOptions) - Operation response error`, func() {
		getDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDistributionListDestinationPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDistributionListDestination with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDistributionListDestinationOptions model
				getDistributionListDestinationOptionsModel := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				getDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDistributionListDestination(getDistributionListDestinationOptions *GetDistributionListDestinationOptions)`, func() {
		getDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDistributionListDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke GetDistributionListDestination successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetDistributionListDestinationOptions model
				getDistributionListDestinationOptionsModel := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				getDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.GetDistributionListDestinationWithContext(ctx, getDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.GetDistributionListDestinationWithContext(ctx, getDistributionListDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getDistributionListDestinationPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"destination_id": "12345678-1234-1234-1234-123456789012", "destination_type": "event_notifications"}`)
				}))
			})
			It(`Invoke GetDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.GetDistributionListDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDistributionListDestinationOptions model
				getDistributionListDestinationOptionsModel := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				getDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDistributionListDestination with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDistributionListDestinationOptions model
				getDistributionListDestinationOptionsModel := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				getDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDistributionListDestinationOptions model with no property values
				getDistributionListDestinationOptionsModelNew := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModelNew)
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
			It(`Invoke GetDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetDistributionListDestinationOptions model
				getDistributionListDestinationOptionsModel := new(platformnotificationsv1.GetDistributionListDestinationOptions)
				getDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.GetDistributionListDestination(getDistributionListDestinationOptionsModel)
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
	Describe(`DeleteDistributionListDestination(deleteDistributionListDestinationOptions *DeleteDistributionListDestinationOptions)`, func() {
		deleteDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDistributionListDestinationPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := platformNotificationsService.DeleteDistributionListDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteDistributionListDestinationOptions model
				deleteDistributionListDestinationOptionsModel := new(platformnotificationsv1.DeleteDistributionListDestinationOptions)
				deleteDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				deleteDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = platformNotificationsService.DeleteDistributionListDestination(deleteDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteDistributionListDestination with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteDistributionListDestinationOptions model
				deleteDistributionListDestinationOptionsModel := new(platformnotificationsv1.DeleteDistributionListDestinationOptions)
				deleteDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				deleteDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := platformNotificationsService.DeleteDistributionListDestination(deleteDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteDistributionListDestinationOptions model with no property values
				deleteDistributionListDestinationOptionsModelNew := new(platformnotificationsv1.DeleteDistributionListDestinationOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = platformNotificationsService.DeleteDistributionListDestination(deleteDistributionListDestinationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestDistributionListDestination(testDistributionListDestinationOptions *TestDistributionListDestinationOptions) - Operation response error`, func() {
		testDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012/test"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testDistributionListDestinationPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke TestDistributionListDestination with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDistributionListDestinationOptions model
				testDistributionListDestinationOptionsModel := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				testDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototypeModel
				testDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`TestDistributionListDestination(testDistributionListDestinationOptions *TestDistributionListDestinationOptions)`, func() {
		testDistributionListDestinationPath := "/v1/distribution_lists/a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6/destinations/12345678-1234-1234-1234-123456789012/test"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(testDistributionListDestinationPath))
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
					fmt.Fprintf(res, "%s", `{"message": "success"}`)
				}))
			})
			It(`Invoke TestDistributionListDestination successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDistributionListDestinationOptions model
				testDistributionListDestinationOptionsModel := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				testDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototypeModel
				testDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.TestDistributionListDestinationWithContext(ctx, testDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.TestDistributionListDestinationWithContext(ctx, testDistributionListDestinationOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(testDistributionListDestinationPath))
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
					fmt.Fprintf(res, "%s", `{"message": "success"}`)
				}))
			})
			It(`Invoke TestDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.TestDistributionListDestination(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDistributionListDestinationOptions model
				testDistributionListDestinationOptionsModel := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				testDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototypeModel
				testDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke TestDistributionListDestination with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDistributionListDestinationOptions model
				testDistributionListDestinationOptionsModel := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				testDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototypeModel
				testDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the TestDistributionListDestinationOptions model with no property values
				testDistributionListDestinationOptionsModelNew := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModelNew)
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
			It(`Invoke TestDistributionListDestination successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")

				// Construct an instance of the TestDistributionListDestinationOptions model
				testDistributionListDestinationOptionsModel := new(platformnotificationsv1.TestDistributionListDestinationOptions)
				testDistributionListDestinationOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.DestinationID = core.StringPtr("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype = testDestinationRequestBodyPrototypeModel
				testDistributionListDestinationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.TestDistributionListDestination(testDistributionListDestinationOptionsModel)
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
			platformNotificationsService, _ := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
				URL:           "http://platformnotificationsv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateDistributionListDestinationOptions successfully`, func() {
				// Construct an instance of the AddDestinationPrototypeEventNotificationDestinationPrototype model
				addDestinationPrototypeModel := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
				Expect(addDestinationPrototypeModel).ToNot(BeNil())
				addDestinationPrototypeModel.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
				addDestinationPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				Expect(addDestinationPrototypeModel.DestinationID).To(Equal(CreateMockUUID("12345678-1234-1234-1234-123456789012")))
				Expect(addDestinationPrototypeModel.DestinationType).To(Equal(core.StringPtr("event_notifications")))

				// Construct an instance of the CreateDistributionListDestinationOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
				var addDestinationPrototype platformnotificationsv1.AddDestinationPrototypeIntf = nil
				createDistributionListDestinationOptionsModel := platformNotificationsService.NewCreateDistributionListDestinationOptions(accountID, addDestinationPrototype)
				createDistributionListDestinationOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createDistributionListDestinationOptionsModel.SetAddDestinationPrototype(addDestinationPrototypeModel)
				createDistributionListDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDistributionListDestinationOptionsModel).ToNot(BeNil())
				Expect(createDistributionListDestinationOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(createDistributionListDestinationOptionsModel.AddDestinationPrototype).To(Equal(addDestinationPrototypeModel))
				Expect(createDistributionListDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDistributionListDestinationOptions successfully`, func() {
				// Construct an instance of the DeleteDistributionListDestinationOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
				destinationID := "12345678-1234-1234-1234-123456789012"
				deleteDistributionListDestinationOptionsModel := platformNotificationsService.NewDeleteDistributionListDestinationOptions(accountID, destinationID)
				deleteDistributionListDestinationOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteDistributionListDestinationOptionsModel.SetDestinationID("12345678-1234-1234-1234-123456789012")
				deleteDistributionListDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDistributionListDestinationOptionsModel).ToNot(BeNil())
				Expect(deleteDistributionListDestinationOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(deleteDistributionListDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("12345678-1234-1234-1234-123456789012")))
				Expect(deleteDistributionListDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDistributionListDestinationOptions successfully`, func() {
				// Construct an instance of the GetDistributionListDestinationOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
				destinationID := "12345678-1234-1234-1234-123456789012"
				getDistributionListDestinationOptionsModel := platformNotificationsService.NewGetDistributionListDestinationOptions(accountID, destinationID)
				getDistributionListDestinationOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getDistributionListDestinationOptionsModel.SetDestinationID("12345678-1234-1234-1234-123456789012")
				getDistributionListDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDistributionListDestinationOptionsModel).ToNot(BeNil())
				Expect(getDistributionListDestinationOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(getDistributionListDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("12345678-1234-1234-1234-123456789012")))
				Expect(getDistributionListDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDistributionListDestinationsOptions successfully`, func() {
				// Construct an instance of the ListDistributionListDestinationsOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
				listDistributionListDestinationsOptionsModel := platformNotificationsService.NewListDistributionListDestinationsOptions(accountID)
				listDistributionListDestinationsOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listDistributionListDestinationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDistributionListDestinationsOptionsModel).ToNot(BeNil())
				Expect(listDistributionListDestinationsOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(listDistributionListDestinationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTestDistributionListDestinationOptions successfully`, func() {
				// Construct an instance of the TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype model
				testDestinationRequestBodyPrototypeModel := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
				Expect(testDestinationRequestBodyPrototypeModel).ToNot(BeNil())
				testDestinationRequestBodyPrototypeModel.DestinationType = core.StringPtr("event_notifications")
				testDestinationRequestBodyPrototypeModel.NotificationType = core.StringPtr("incident")
				Expect(testDestinationRequestBodyPrototypeModel.DestinationType).To(Equal(core.StringPtr("event_notifications")))
				Expect(testDestinationRequestBodyPrototypeModel.NotificationType).To(Equal(core.StringPtr("incident")))

				// Construct an instance of the TestDistributionListDestinationOptions model
				accountID := "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
				destinationID := "12345678-1234-1234-1234-123456789012"
				var testDestinationRequestBodyPrototype platformnotificationsv1.TestDestinationRequestBodyPrototypeIntf = nil
				testDistributionListDestinationOptionsModel := platformNotificationsService.NewTestDistributionListDestinationOptions(accountID, destinationID, testDestinationRequestBodyPrototype)
				testDistributionListDestinationOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				testDistributionListDestinationOptionsModel.SetDestinationID("12345678-1234-1234-1234-123456789012")
				testDistributionListDestinationOptionsModel.SetTestDestinationRequestBodyPrototype(testDestinationRequestBodyPrototypeModel)
				testDistributionListDestinationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(testDistributionListDestinationOptionsModel).ToNot(BeNil())
				Expect(testDistributionListDestinationOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(testDistributionListDestinationOptionsModel.DestinationID).To(Equal(core.StringPtr("12345678-1234-1234-1234-123456789012")))
				Expect(testDistributionListDestinationOptionsModel.TestDestinationRequestBodyPrototype).To(Equal(testDestinationRequestBodyPrototypeModel))
				Expect(testDistributionListDestinationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddDestinationPrototypeEventNotificationDestinationPrototype successfully`, func() {
				destinationID := CreateMockUUID("12345678-1234-1234-1234-123456789012")
				destinationType := "event_notifications"
				_model, err := platformNotificationsService.NewAddDestinationPrototypeEventNotificationDestinationPrototype(destinationID, destinationType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype successfully`, func() {
				destinationType := "event_notifications"
				notificationType := "incident"
				_model, err := platformNotificationsService.NewTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype(destinationType, notificationType)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Model unmarshaling tests`, func() {
		It(`Invoke UnmarshalAddDestinationPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.AddDestinationPrototype)
			model.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
			model.DestinationType = core.StringPtr("event_notifications")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = platformnotificationsv1.UnmarshalAddDestinationPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalTestDestinationRequestBodyPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.TestDestinationRequestBodyPrototype)
			model.DestinationType = core.StringPtr("event_notifications")
			model.NotificationType = core.StringPtr("incident")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result interface{}
			err = platformnotificationsv1.UnmarshalTestDestinationRequestBodyPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
		})
		It(`Invoke UnmarshalAddDestinationPrototypeEventNotificationDestinationPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype)
			model.DestinationID = CreateMockUUID("12345678-1234-1234-1234-123456789012")
			model.DestinationType = core.StringPtr("event_notifications")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *platformnotificationsv1.AddDestinationPrototypeEventNotificationDestinationPrototype
			err = platformnotificationsv1.UnmarshalAddDestinationPrototypeEventNotificationDestinationPrototype(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype)
			model.DestinationType = core.StringPtr("event_notifications")
			model.NotificationType = core.StringPtr("incident")

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *platformnotificationsv1.TestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype
			err = platformnotificationsv1.UnmarshalTestDestinationRequestBodyPrototypeTestEventNotificationDestinationRequestBodyPrototype(raw, &result)
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
