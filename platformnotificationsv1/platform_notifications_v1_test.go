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
	Describe(`CreatePreferences(createPreferencesOptions *CreatePreferencesOptions) - Operation response error`, func() {
		createPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPreferencesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePreferences with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the CreatePreferencesOptions model
				createPreferencesOptionsModel := new(platformnotificationsv1.CreatePreferencesOptions)
				createPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				createPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePreferences(createPreferencesOptions *CreatePreferencesOptions)`, func() {
		createPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPreferencesPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke CreatePreferences successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the CreatePreferencesOptions model
				createPreferencesOptionsModel := new(platformnotificationsv1.CreatePreferencesOptions)
				createPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				createPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.CreatePreferencesWithContext(ctx, createPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.CreatePreferencesWithContext(ctx, createPreferencesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPreferencesPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke CreatePreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.CreatePreferences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the CreatePreferencesOptions model
				createPreferencesOptionsModel := new(platformnotificationsv1.CreatePreferencesOptions)
				createPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				createPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePreferences with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the CreatePreferencesOptions model
				createPreferencesOptionsModel := new(platformnotificationsv1.CreatePreferencesOptions)
				createPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				createPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePreferencesOptions model with no property values
				createPreferencesOptionsModelNew := new(platformnotificationsv1.CreatePreferencesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.CreatePreferences(createPreferencesOptionsModelNew)
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
			It(`Invoke CreatePreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the CreatePreferencesOptions model
				createPreferencesOptionsModel := new(platformnotificationsv1.CreatePreferencesOptions)
				createPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				createPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				createPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.CreatePreferences(createPreferencesOptionsModel)
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
	Describe(`GetPreferences(getPreferencesOptions *GetPreferencesOptions) - Operation response error`, func() {
		getPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPreferencesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPreferences with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetPreferencesOptions model
				getPreferencesOptionsModel := new(platformnotificationsv1.GetPreferencesOptions)
				getPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				getPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPreferences(getPreferencesOptions *GetPreferencesOptions)`, func() {
		getPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPreferencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke GetPreferences successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetPreferencesOptions model
				getPreferencesOptionsModel := new(platformnotificationsv1.GetPreferencesOptions)
				getPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				getPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.GetPreferencesWithContext(ctx, getPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.GetPreferencesWithContext(ctx, getPreferencesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPreferencesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke GetPreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.GetPreferences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPreferencesOptions model
				getPreferencesOptionsModel := new(platformnotificationsv1.GetPreferencesOptions)
				getPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				getPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPreferences with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetPreferencesOptions model
				getPreferencesOptionsModel := new(platformnotificationsv1.GetPreferencesOptions)
				getPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				getPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPreferencesOptions model with no property values
				getPreferencesOptionsModelNew := new(platformnotificationsv1.GetPreferencesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.GetPreferences(getPreferencesOptionsModelNew)
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
			It(`Invoke GetPreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetPreferencesOptions model
				getPreferencesOptionsModel := new(platformnotificationsv1.GetPreferencesOptions)
				getPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				getPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.GetPreferences(getPreferencesOptionsModel)
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
	Describe(`ReplaceNotificationPreferences(replaceNotificationPreferencesOptions *ReplaceNotificationPreferencesOptions) - Operation response error`, func() {
		replaceNotificationPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationPreferencesPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceNotificationPreferences with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				replaceNotificationPreferencesOptionsModel := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				replaceNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceNotificationPreferences(replaceNotificationPreferencesOptions *ReplaceNotificationPreferencesOptions)`, func() {
		replaceNotificationPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationPreferencesPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke ReplaceNotificationPreferences successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				replaceNotificationPreferencesOptionsModel := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				replaceNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.ReplaceNotificationPreferencesWithContext(ctx, replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.ReplaceNotificationPreferencesWithContext(ctx, replaceNotificationPreferencesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationPreferencesPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"incident_severity1": {"channels": ["email"], "updates": true}, "incident_severity2": {"channels": ["email"], "updates": true}, "incident_severity3": {"channels": ["email"], "updates": true}, "incident_severity4": {"channels": ["email"], "updates": true}, "maintenance_high": {"channels": ["email"], "updates": true}, "maintenance_medium": {"channels": ["email"], "updates": true}, "maintenance_low": {"channels": ["email"], "updates": true}, "announcements_major": {"channels": ["email"]}, "announcements_minor": {"channels": ["email"]}, "security_normal": {"channels": ["email"]}, "account_normal": {"channels": ["email"]}, "billing_and_usage_order": {"channels": ["email"]}, "billing_and_usage_invoices": {"channels": ["email"]}, "billing_and_usage_payments": {"channels": ["email"]}, "billing_and_usage_subscriptions_and_promo_codes": {"channels": ["email"]}, "billing_and_usage_spending_alerts": {"channels": ["email"]}, "resourceactivity_normal": {"channels": ["email"]}, "ordering_review": {"channels": ["email"]}, "ordering_approved": {"channels": ["email"]}, "ordering_approved_vsi": {"channels": ["email"]}, "ordering_approved_server": {"channels": ["email"]}, "provisioning_reload_complete": {"channels": ["email"]}, "provisioning_complete_vsi": {"channels": ["email"]}, "provisioning_complete_server": {"channels": ["email"]}}`)
				}))
			})
			It(`Invoke ReplaceNotificationPreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.ReplaceNotificationPreferences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				replaceNotificationPreferencesOptionsModel := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				replaceNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceNotificationPreferences with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				replaceNotificationPreferencesOptionsModel := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				replaceNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceNotificationPreferencesOptions model with no property values
				replaceNotificationPreferencesOptionsModelNew := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModelNew)
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
			It(`Invoke ReplaceNotificationPreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				replaceNotificationPreferencesOptionsModel := new(platformnotificationsv1.ReplaceNotificationPreferencesOptions)
				replaceNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.IncidentSeverity1 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity2 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity3 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.IncidentSeverity4 = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceHigh = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceMedium = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.MaintenanceLow = preferenceValueWithUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMajor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AnnouncementsMinor = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.SecurityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ResourceactivityNormal = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingReview = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApproved = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.OrderingApprovedServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer = preferenceValueWithoutUpdatesModel
				replaceNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.ReplaceNotificationPreferences(replaceNotificationPreferencesOptionsModel)
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
	Describe(`DeleteNotificationPreferences(deleteNotificationPreferencesOptions *DeleteNotificationPreferencesOptions)`, func() {
		deleteNotificationPreferencesPath := "/v1/notifications/IBMid-1234567890/preferences"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteNotificationPreferencesPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteNotificationPreferences successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := platformNotificationsService.DeleteNotificationPreferences(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteNotificationPreferencesOptions model
				deleteNotificationPreferencesOptionsModel := new(platformnotificationsv1.DeleteNotificationPreferencesOptions)
				deleteNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				deleteNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = platformNotificationsService.DeleteNotificationPreferences(deleteNotificationPreferencesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteNotificationPreferences with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the DeleteNotificationPreferencesOptions model
				deleteNotificationPreferencesOptionsModel := new(platformnotificationsv1.DeleteNotificationPreferencesOptions)
				deleteNotificationPreferencesOptionsModel.IamID = core.StringPtr("IBMid-1234567890")
				deleteNotificationPreferencesOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteNotificationPreferencesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := platformNotificationsService.DeleteNotificationPreferences(deleteNotificationPreferencesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteNotificationPreferencesOptions model with no property values
				deleteNotificationPreferencesOptionsModelNew := new(platformnotificationsv1.DeleteNotificationPreferencesOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = platformNotificationsService.DeleteNotificationPreferences(deleteNotificationPreferencesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions) - Operation response error`, func() {
		listNotificationsPath := "/v1/notifications"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"3fe78a36b9aa7f26"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListNotifications with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(platformnotificationsv1.ListNotificationsOptions)
				listNotificationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.Start = core.StringPtr("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListNotifications(listNotificationsOptions *ListNotificationsOptions)`, func() {
		listNotificationsPath := "/v1/notifications"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"3fe78a36b9aa7f26"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 50, "total_count": 232, "first": {"href": "https://api.example.com/v1/notifications?limit=50"}, "previous": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "next": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "last": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "notifications": [{"title": "System Maintenance Scheduled", "body": "Scheduled maintenance will occur on March 15th from 10:00 AM to 11:00 AM UTC.", "id": "12345", "category": "maintenance", "component_names": ["ComponentNames"], "start_time": 1771791490, "is_global": false, "state": "new", "regions": ["Regions"], "crn_masks": ["CrnMasks"], "record_id": "rec-67890", "source_id": "src-11111", "completion_code": "successful", "end_time": 1771791490, "update_time": 1771791490, "severity": 2, "lucene_query": "region:us-south AND service_name:event-notifications", "resource_link": "https://cloud.ibm.com/status/incident/12345"}]}`)
				}))
			})
			It(`Invoke ListNotifications successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(platformnotificationsv1.ListNotificationsOptions)
				listNotificationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.Start = core.StringPtr("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.ListNotificationsWithContext(ctx, listNotificationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.ListNotificationsWithContext(ctx, listNotificationsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"}))
					Expect(req.URL.Query()["start"]).To(Equal([]string{"3fe78a36b9aa7f26"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(50))}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 50, "total_count": 232, "first": {"href": "https://api.example.com/v1/notifications?limit=50"}, "previous": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "next": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "last": {"href": "https://api.example.com/v1/notifications?start=3fe78a36b9aa7f26&limit=50", "start": "3fe78a36b9aa7f26"}, "notifications": [{"title": "System Maintenance Scheduled", "body": "Scheduled maintenance will occur on March 15th from 10:00 AM to 11:00 AM UTC.", "id": "12345", "category": "maintenance", "component_names": ["ComponentNames"], "start_time": 1771791490, "is_global": false, "state": "new", "regions": ["Regions"], "crn_masks": ["CrnMasks"], "record_id": "rec-67890", "source_id": "src-11111", "completion_code": "successful", "end_time": 1771791490, "update_time": 1771791490, "severity": 2, "lucene_query": "region:us-south AND service_name:event-notifications", "resource_link": "https://cloud.ibm.com/status/incident/12345"}]}`)
				}))
			})
			It(`Invoke ListNotifications successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.ListNotifications(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(platformnotificationsv1.ListNotificationsOptions)
				listNotificationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.Start = core.StringPtr("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListNotifications with error: Operation request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(platformnotificationsv1.ListNotificationsOptions)
				listNotificationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.Start = core.StringPtr("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
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
			It(`Invoke ListNotifications successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := new(platformnotificationsv1.ListNotificationsOptions)
				listNotificationsOptionsModel.AccountID = core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.Start = core.StringPtr("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.Limit = core.Int64Ptr(int64(50))
				listNotificationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.ListNotifications(listNotificationsOptionsModel)
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
			It(`Invoke GetNextStart successfully`, func() {
				responseObject := new(platformnotificationsv1.NotificationCollection)
				nextObject := new(platformnotificationsv1.PaginationLinkWithToken)
				nextObject.Start = core.StringPtr("abc-123")
				responseObject.Next = nextObject

				value, err := responseObject.GetNextStart()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.StringPtr("abc-123")))
			})
			It(`Invoke GetNextStart without a "Next" property in the response`, func() {
				responseObject := new(platformnotificationsv1.NotificationCollection)

				value, err := responseObject.GetNextStart()
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
					Expect(req.URL.EscapedPath()).To(Equal(listNotificationsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"start":"1"},"total_count":2,"limit":1,"notifications":[{"title":"System Maintenance Scheduled","body":"Scheduled maintenance will occur on March 15th from 10:00 AM to 11:00 AM UTC.","id":"12345","category":"maintenance","component_names":["ComponentNames"],"start_time":1771791490,"is_global":false,"state":"new","regions":["Regions"],"crn_masks":["CrnMasks"],"record_id":"rec-67890","source_id":"src-11111","completion_code":"successful","end_time":1771791490,"update_time":1771791490,"severity":2,"lucene_query":"region:us-south AND service_name:event-notifications","resource_link":"https://cloud.ibm.com/status/incident/12345"}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"notifications":[{"title":"System Maintenance Scheduled","body":"Scheduled maintenance will occur on March 15th from 10:00 AM to 11:00 AM UTC.","id":"12345","category":"maintenance","component_names":["ComponentNames"],"start_time":1771791490,"is_global":false,"state":"new","regions":["Regions"],"crn_masks":["CrnMasks"],"record_id":"rec-67890","source_id":"src-11111","completion_code":"successful","end_time":1771791490,"update_time":1771791490,"severity":2,"lucene_query":"region:us-south AND service_name:event-notifications","resource_link":"https://cloud.ibm.com/status/incident/12345"}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use NotificationsPager.GetNext successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				listNotificationsOptionsModel := &platformnotificationsv1.ListNotificationsOptions{
					AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
					Limit: core.Int64Ptr(int64(50)),
				}

				pager, err := platformNotificationsService.NewNotificationsPager(listNotificationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []platformnotificationsv1.Notification
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use NotificationsPager.GetAll successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				listNotificationsOptionsModel := &platformnotificationsv1.ListNotificationsOptions{
					AccountID: core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"),
					Limit: core.Int64Ptr(int64(50)),
				}

				pager, err := platformNotificationsService.NewNotificationsPager(listNotificationsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAcknowledgment(getAcknowledgmentOptions *GetAcknowledgmentOptions) - Operation response error`, func() {
		getAcknowledgmentPath := "/v1/notifications/acknowledgment"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAcknowledgmentPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAcknowledgment with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := new(platformnotificationsv1.GetAcknowledgmentOptions)
				getAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAcknowledgment(getAcknowledgmentOptions *GetAcknowledgmentOptions)`, func() {
		getAcknowledgmentPath := "/v1/notifications/acknowledgment"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAcknowledgmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"has_unread": true, "latest_notification_id": "1678901234000", "last_acknowledged_id": "1678800000000"}`)
				}))
			})
			It(`Invoke GetAcknowledgment successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := new(platformnotificationsv1.GetAcknowledgmentOptions)
				getAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.GetAcknowledgmentWithContext(ctx, getAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.GetAcknowledgmentWithContext(ctx, getAcknowledgmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getAcknowledgmentPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"has_unread": true, "latest_notification_id": "1678901234000", "last_acknowledged_id": "1678800000000"}`)
				}))
			})
			It(`Invoke GetAcknowledgment successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.GetAcknowledgment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := new(platformnotificationsv1.GetAcknowledgmentOptions)
				getAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetAcknowledgment with error: Operation request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := new(platformnotificationsv1.GetAcknowledgmentOptions)
				getAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
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
			It(`Invoke GetAcknowledgment successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := new(platformnotificationsv1.GetAcknowledgmentOptions)
				getAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.GetAcknowledgment(getAcknowledgmentOptionsModel)
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
	Describe(`ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions *ReplaceNotificationAcknowledgmentOptions) - Operation response error`, func() {
		replaceNotificationAcknowledgmentPath := "/v1/notifications/acknowledgment"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationAcknowledgmentPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceNotificationAcknowledgment with error: Operation response processing error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsModel := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID = core.StringPtr("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				platformNotificationsService.EnableRetries(0, 0)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptions *ReplaceNotificationAcknowledgmentOptions)`, func() {
		replaceNotificationAcknowledgmentPath := "/v1/notifications/acknowledgment"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationAcknowledgmentPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"has_unread": true, "latest_notification_id": "1678901234000", "last_acknowledged_id": "1678800000000"}`)
				}))
			})
			It(`Invoke ReplaceNotificationAcknowledgment successfully with retries`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())
				platformNotificationsService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsModel := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID = core.StringPtr("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgmentWithContext(ctx, replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				platformNotificationsService.DisableRetries()
				result, response, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = platformNotificationsService.ReplaceNotificationAcknowledgmentWithContext(ctx, replaceNotificationAcknowledgmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceNotificationAcknowledgmentPath))
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

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"1369339417d906e5620b8d861d40cfd7"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"has_unread": true, "latest_notification_id": "1678901234000", "last_acknowledged_id": "1678800000000"}`)
				}))
			})
			It(`Invoke ReplaceNotificationAcknowledgment successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsModel := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID = core.StringPtr("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceNotificationAcknowledgment with error: Operation validation and request error`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsModel := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID = core.StringPtr("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := platformNotificationsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceNotificationAcknowledgmentOptions model with no property values
				replaceNotificationAcknowledgmentOptionsModelNew := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModelNew)
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
			It(`Invoke ReplaceNotificationAcknowledgment successfully`, func() {
				platformNotificationsService, serviceErr := platformnotificationsv1.NewPlatformNotificationsV1(&platformnotificationsv1.PlatformNotificationsV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(platformNotificationsService).ToNot(BeNil())

				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsModel := new(platformnotificationsv1.ReplaceNotificationAcknowledgmentOptions)
				replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID = core.StringPtr("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.AccountID = core.StringPtr("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := platformNotificationsService.ReplaceNotificationAcknowledgment(replaceNotificationAcknowledgmentOptionsModel)
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
			It(`Invoke NewCreatePreferencesOptions successfully`, func() {
				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				Expect(preferenceValueWithUpdatesModel).ToNot(BeNil())
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)
				Expect(preferenceValueWithUpdatesModel.Channels).To(Equal([]string{"email"}))
				Expect(preferenceValueWithUpdatesModel.Updates).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				Expect(preferenceValueWithoutUpdatesModel).ToNot(BeNil())
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}
				Expect(preferenceValueWithoutUpdatesModel.Channels).To(Equal([]string{"email"}))

				// Construct an instance of the CreatePreferencesOptions model
				iamID := "IBMid-1234567890"
				createPreferencesOptionsModel := platformNotificationsService.NewCreatePreferencesOptions(iamID)
				createPreferencesOptionsModel.SetIamID("IBMid-1234567890")
				createPreferencesOptionsModel.SetIncidentSeverity1(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetIncidentSeverity2(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetIncidentSeverity3(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetIncidentSeverity4(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetMaintenanceHigh(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetMaintenanceMedium(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetMaintenanceLow(preferenceValueWithUpdatesModel)
				createPreferencesOptionsModel.SetAnnouncementsMajor(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetAnnouncementsMinor(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetSecurityNormal(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetAccountNormal(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetBillingAndUsageOrder(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetBillingAndUsageInvoices(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetBillingAndUsagePayments(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetBillingAndUsageSubscriptionsAndPromoCodes(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetBillingAndUsageSpendingAlerts(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetResourceactivityNormal(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetOrderingReview(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetOrderingApproved(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetOrderingApprovedVsi(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetOrderingApprovedServer(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetProvisioningReloadComplete(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetProvisioningCompleteVsi(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetProvisioningCompleteServer(preferenceValueWithoutUpdatesModel)
				createPreferencesOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				createPreferencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPreferencesOptionsModel).ToNot(BeNil())
				Expect(createPreferencesOptionsModel.IamID).To(Equal(core.StringPtr("IBMid-1234567890")))
				Expect(createPreferencesOptionsModel.IncidentSeverity1).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.IncidentSeverity2).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.IncidentSeverity3).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.IncidentSeverity4).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.MaintenanceHigh).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.MaintenanceMedium).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.MaintenanceLow).To(Equal(preferenceValueWithUpdatesModel))
				Expect(createPreferencesOptionsModel.AnnouncementsMajor).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.AnnouncementsMinor).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.SecurityNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.AccountNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.BillingAndUsageOrder).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.BillingAndUsageInvoices).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.BillingAndUsagePayments).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.BillingAndUsageSpendingAlerts).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.ResourceactivityNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.OrderingReview).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.OrderingApproved).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.OrderingApprovedVsi).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.OrderingApprovedServer).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.ProvisioningReloadComplete).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.ProvisioningCompleteVsi).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.ProvisioningCompleteServer).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(createPreferencesOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(createPreferencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewDeleteNotificationPreferencesOptions successfully`, func() {
				// Construct an instance of the DeleteNotificationPreferencesOptions model
				iamID := "IBMid-1234567890"
				deleteNotificationPreferencesOptionsModel := platformNotificationsService.NewDeleteNotificationPreferencesOptions(iamID)
				deleteNotificationPreferencesOptionsModel.SetIamID("IBMid-1234567890")
				deleteNotificationPreferencesOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				deleteNotificationPreferencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteNotificationPreferencesOptionsModel).ToNot(BeNil())
				Expect(deleteNotificationPreferencesOptionsModel.IamID).To(Equal(core.StringPtr("IBMid-1234567890")))
				Expect(deleteNotificationPreferencesOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(deleteNotificationPreferencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAcknowledgmentOptions successfully`, func() {
				// Construct an instance of the GetAcknowledgmentOptions model
				getAcknowledgmentOptionsModel := platformNotificationsService.NewGetAcknowledgmentOptions()
				getAcknowledgmentOptionsModel.SetAccountID("1369339417d906e5620b8d861d40cfd7")
				getAcknowledgmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAcknowledgmentOptionsModel).ToNot(BeNil())
				Expect(getAcknowledgmentOptionsModel.AccountID).To(Equal(core.StringPtr("1369339417d906e5620b8d861d40cfd7")))
				Expect(getAcknowledgmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetPreferencesOptions successfully`, func() {
				// Construct an instance of the GetPreferencesOptions model
				iamID := "IBMid-1234567890"
				getPreferencesOptionsModel := platformNotificationsService.NewGetPreferencesOptions(iamID)
				getPreferencesOptionsModel.SetIamID("IBMid-1234567890")
				getPreferencesOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				getPreferencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPreferencesOptionsModel).ToNot(BeNil())
				Expect(getPreferencesOptionsModel.IamID).To(Equal(core.StringPtr("IBMid-1234567890")))
				Expect(getPreferencesOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(getPreferencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewListNotificationsOptions successfully`, func() {
				// Construct an instance of the ListNotificationsOptions model
				listNotificationsOptionsModel := platformNotificationsService.NewListNotificationsOptions()
				listNotificationsOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				listNotificationsOptionsModel.SetStart("3fe78a36b9aa7f26")
				listNotificationsOptionsModel.SetLimit(int64(50))
				listNotificationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listNotificationsOptionsModel).ToNot(BeNil())
				Expect(listNotificationsOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(listNotificationsOptionsModel.Start).To(Equal(core.StringPtr("3fe78a36b9aa7f26")))
				Expect(listNotificationsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(50))))
				Expect(listNotificationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPreferenceValueWithUpdates successfully`, func() {
				channels := []string{"email"}
				_model, err := platformNotificationsService.NewPreferenceValueWithUpdates(channels)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPreferenceValueWithoutUpdates successfully`, func() {
				channels := []string{"email"}
				_model, err := platformNotificationsService.NewPreferenceValueWithoutUpdates(channels)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceNotificationAcknowledgmentOptions successfully`, func() {
				// Construct an instance of the ReplaceNotificationAcknowledgmentOptions model
				replaceNotificationAcknowledgmentOptionsLastAcknowledgedID := "1772804159452"
				replaceNotificationAcknowledgmentOptionsModel := platformNotificationsService.NewReplaceNotificationAcknowledgmentOptions(replaceNotificationAcknowledgmentOptionsLastAcknowledgedID)
				replaceNotificationAcknowledgmentOptionsModel.SetLastAcknowledgedID("1772804159452")
				replaceNotificationAcknowledgmentOptionsModel.SetAccountID("1369339417d906e5620b8d861d40cfd7")
				replaceNotificationAcknowledgmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceNotificationAcknowledgmentOptionsModel).ToNot(BeNil())
				Expect(replaceNotificationAcknowledgmentOptionsModel.LastAcknowledgedID).To(Equal(core.StringPtr("1772804159452")))
				Expect(replaceNotificationAcknowledgmentOptionsModel.AccountID).To(Equal(core.StringPtr("1369339417d906e5620b8d861d40cfd7")))
				Expect(replaceNotificationAcknowledgmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceNotificationPreferencesOptions successfully`, func() {
				// Construct an instance of the PreferenceValueWithUpdates model
				preferenceValueWithUpdatesModel := new(platformnotificationsv1.PreferenceValueWithUpdates)
				Expect(preferenceValueWithUpdatesModel).ToNot(BeNil())
				preferenceValueWithUpdatesModel.Channels = []string{"email"}
				preferenceValueWithUpdatesModel.Updates = core.BoolPtr(true)
				Expect(preferenceValueWithUpdatesModel.Channels).To(Equal([]string{"email"}))
				Expect(preferenceValueWithUpdatesModel.Updates).To(Equal(core.BoolPtr(true)))

				// Construct an instance of the PreferenceValueWithoutUpdates model
				preferenceValueWithoutUpdatesModel := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
				Expect(preferenceValueWithoutUpdatesModel).ToNot(BeNil())
				preferenceValueWithoutUpdatesModel.Channels = []string{"email"}
				Expect(preferenceValueWithoutUpdatesModel.Channels).To(Equal([]string{"email"}))

				// Construct an instance of the ReplaceNotificationPreferencesOptions model
				iamID := "IBMid-1234567890"
				replaceNotificationPreferencesOptionsModel := platformNotificationsService.NewReplaceNotificationPreferencesOptions(iamID)
				replaceNotificationPreferencesOptionsModel.SetIamID("IBMid-1234567890")
				replaceNotificationPreferencesOptionsModel.SetIncidentSeverity1(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetIncidentSeverity2(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetIncidentSeverity3(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetIncidentSeverity4(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetMaintenanceHigh(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetMaintenanceMedium(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetMaintenanceLow(preferenceValueWithUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetAnnouncementsMajor(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetAnnouncementsMinor(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetSecurityNormal(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetAccountNormal(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetBillingAndUsageOrder(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetBillingAndUsageInvoices(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetBillingAndUsagePayments(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetBillingAndUsageSubscriptionsAndPromoCodes(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetBillingAndUsageSpendingAlerts(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetResourceactivityNormal(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetOrderingReview(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetOrderingApproved(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetOrderingApprovedVsi(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetOrderingApprovedServer(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetProvisioningReloadComplete(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetProvisioningCompleteVsi(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetProvisioningCompleteServer(preferenceValueWithoutUpdatesModel)
				replaceNotificationPreferencesOptionsModel.SetAccountID("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")
				replaceNotificationPreferencesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceNotificationPreferencesOptionsModel).ToNot(BeNil())
				Expect(replaceNotificationPreferencesOptionsModel.IamID).To(Equal(core.StringPtr("IBMid-1234567890")))
				Expect(replaceNotificationPreferencesOptionsModel.IncidentSeverity1).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.IncidentSeverity2).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.IncidentSeverity3).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.IncidentSeverity4).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.MaintenanceHigh).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.MaintenanceMedium).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.MaintenanceLow).To(Equal(preferenceValueWithUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.AnnouncementsMajor).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.AnnouncementsMinor).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.SecurityNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.AccountNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.BillingAndUsageOrder).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.BillingAndUsageInvoices).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.BillingAndUsagePayments).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.BillingAndUsageSubscriptionsAndPromoCodes).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.BillingAndUsageSpendingAlerts).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.ResourceactivityNormal).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.OrderingReview).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.OrderingApproved).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.OrderingApprovedVsi).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.OrderingApprovedServer).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.ProvisioningReloadComplete).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.ProvisioningCompleteVsi).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.ProvisioningCompleteServer).To(Equal(preferenceValueWithoutUpdatesModel))
				Expect(replaceNotificationPreferencesOptionsModel.AccountID).To(Equal(core.StringPtr("a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6")))
				Expect(replaceNotificationPreferencesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
		It(`Invoke UnmarshalPreferenceValueWithUpdates successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.PreferenceValueWithUpdates)
			model.Channels = []string{"email"}
			model.Updates = core.BoolPtr(true)

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *platformnotificationsv1.PreferenceValueWithUpdates
			err = platformnotificationsv1.UnmarshalPreferenceValueWithUpdates(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPreferenceValueWithoutUpdates successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.PreferenceValueWithoutUpdates)
			model.Channels = []string{"email"}

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *platformnotificationsv1.PreferenceValueWithoutUpdates
			err = platformnotificationsv1.UnmarshalPreferenceValueWithoutUpdates(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
		})
		It(`Invoke UnmarshalPreferencesObject successfully`, func() {
			// Construct an instance of the model.
			model := new(platformnotificationsv1.PreferencesObject)
			model.IncidentSeverity1 = nil
			model.IncidentSeverity2 = nil
			model.IncidentSeverity3 = nil
			model.IncidentSeverity4 = nil
			model.MaintenanceHigh = nil
			model.MaintenanceMedium = nil
			model.MaintenanceLow = nil
			model.AnnouncementsMajor = nil
			model.AnnouncementsMinor = nil
			model.SecurityNormal = nil
			model.AccountNormal = nil
			model.BillingAndUsageOrder = nil
			model.BillingAndUsageInvoices = nil
			model.BillingAndUsagePayments = nil
			model.BillingAndUsageSubscriptionsAndPromoCodes = nil
			model.BillingAndUsageSpendingAlerts = nil
			model.ResourceactivityNormal = nil
			model.OrderingReview = nil
			model.OrderingApproved = nil
			model.OrderingApprovedVsi = nil
			model.OrderingApprovedServer = nil
			model.ProvisioningReloadComplete = nil
			model.ProvisioningCompleteVsi = nil
			model.ProvisioningCompleteServer = nil

			b, err := json.Marshal(model)
			Expect(err).To(BeNil())

			var raw map[string]json.RawMessage
			err = json.Unmarshal(b, &raw)
			Expect(err).To(BeNil())

			var result *platformnotificationsv1.PreferencesObject
			err = platformnotificationsv1.UnmarshalPreferencesObject(raw, &result)
			Expect(err).To(BeNil())
			Expect(result).ToNot(BeNil())
			Expect(result).To(Equal(model))
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
