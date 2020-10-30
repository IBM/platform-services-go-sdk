/**
 * (C) Copyright IBM Corp. 2020.
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

package resourcemanagerv2_test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
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

var _ = Describe(`ResourceManagerV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(resourceManagerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(resourceManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "https://resourcemanagerv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(resourceManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_URL": "https://resourcemanagerv2/api",
				"RESOURCE_MANAGER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				})
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
					URL: "https://testService/api",
				})
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				})
				err := resourceManagerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_URL": "https://resourcemanagerv2/api",
				"RESOURCE_MANAGER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceGroups(listResourceGroupsOptions *ListResourceGroupsOptions) - Operation response error`, func() {
		listResourceGroupsPath := "/resource_groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceGroupsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceGroups with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the ListResourceGroupsOptions model
				listResourceGroupsOptionsModel := new(resourcemanagerv2.ListResourceGroupsOptions)
				listResourceGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Date = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceGroups(listResourceGroupsOptions *ListResourceGroupsOptions)`, func() {
		listResourceGroupsPath := "/resource_groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listResourceGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListResourceGroups successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.ListResourceGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceGroupsOptions model
				listResourceGroupsOptionsModel := new(resourcemanagerv2.ListResourceGroupsOptions)
				listResourceGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Date = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.ListResourceGroupsWithContext(ctx, listResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.ListResourceGroupsWithContext(ctx, listResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListResourceGroups with error: Operation request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the ListResourceGroupsOptions model
				listResourceGroupsOptionsModel := new(resourcemanagerv2.ListResourceGroupsOptions)
				listResourceGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Date = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateResourceGroup(createResourceGroupOptions *CreateResourceGroupOptions) - Operation response error`, func() {
		createResourceGroupPath := "/resource_groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceGroupPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceGroup with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceGroupOptions model
				createResourceGroupOptionsModel := new(resourcemanagerv2.CreateResourceGroupOptions)
				createResourceGroupOptionsModel.Name = core.StringPtr("test1")
				createResourceGroupOptionsModel.AccountID = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceGroup(createResourceGroupOptions *CreateResourceGroupOptions)`, func() {
		createResourceGroupPath := "/resource_groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createResourceGroupPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "Crn"}`)
				}))
			})
			It(`Invoke CreateResourceGroup successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.CreateResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceGroupOptions model
				createResourceGroupOptionsModel := new(resourcemanagerv2.CreateResourceGroupOptions)
				createResourceGroupOptionsModel.Name = core.StringPtr("test1")
				createResourceGroupOptionsModel.AccountID = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.CreateResourceGroupWithContext(ctx, createResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.CreateResourceGroupWithContext(ctx, createResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateResourceGroup with error: Operation request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the CreateResourceGroupOptions model
				createResourceGroupOptionsModel := new(resourcemanagerv2.CreateResourceGroupOptions)
				createResourceGroupOptionsModel.Name = core.StringPtr("test1")
				createResourceGroupOptionsModel.AccountID = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceGroup(getResourceGroupOptions *GetResourceGroupOptions) - Operation response error`, func() {
		getResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceGroup with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupOptions model
				getResourceGroupOptionsModel := new(resourcemanagerv2.GetResourceGroupOptions)
				getResourceGroupOptionsModel.ID = core.StringPtr("testString")
				getResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceGroup(getResourceGroupOptions *GetResourceGroupOptions)`, func() {
		getResourceGroupPath := "/resource_groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getResourceGroupPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetResourceGroup successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.GetResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupOptions model
				getResourceGroupOptionsModel := new(resourcemanagerv2.GetResourceGroupOptions)
				getResourceGroupOptionsModel.ID = core.StringPtr("testString")
				getResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.GetResourceGroupWithContext(ctx, getResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.GetResourceGroupWithContext(ctx, getResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetResourceGroup with error: Operation validation and request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the GetResourceGroupOptions model
				getResourceGroupOptionsModel := new(resourcemanagerv2.GetResourceGroupOptions)
				getResourceGroupOptionsModel.ID = core.StringPtr("testString")
				getResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceGroupOptions model with no property values
				getResourceGroupOptionsModelNew := new(resourcemanagerv2.GetResourceGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceManagerService.GetResourceGroup(getResourceGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceGroup(updateResourceGroupOptions *UpdateResourceGroupOptions) - Operation response error`, func() {
		updateResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceGroup with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceGroupOptions model
				updateResourceGroupOptionsModel := new(resourcemanagerv2.UpdateResourceGroupOptions)
				updateResourceGroupOptionsModel.ID = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Name = core.StringPtr("testString")
				updateResourceGroupOptionsModel.State = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceGroup(updateResourceGroupOptions *UpdateResourceGroupOptions)`, func() {
		updateResourceGroupPath := "/resource_groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateResourceGroupPath))
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
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateResourceGroup successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.UpdateResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceGroupOptions model
				updateResourceGroupOptionsModel := new(resourcemanagerv2.UpdateResourceGroupOptions)
				updateResourceGroupOptionsModel.ID = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Name = core.StringPtr("testString")
				updateResourceGroupOptionsModel.State = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.UpdateResourceGroupWithContext(ctx, updateResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.UpdateResourceGroupWithContext(ctx, updateResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateResourceGroup with error: Operation validation and request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceGroupOptions model
				updateResourceGroupOptionsModel := new(resourcemanagerv2.UpdateResourceGroupOptions)
				updateResourceGroupOptionsModel.ID = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Name = core.StringPtr("testString")
				updateResourceGroupOptionsModel.State = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceGroupOptions model with no property values
				updateResourceGroupOptionsModelNew := new(resourcemanagerv2.UpdateResourceGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceManagerService.UpdateResourceGroup(updateResourceGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceGroup(deleteResourceGroupOptions *DeleteResourceGroupOptions)`, func() {
		deleteResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteResourceGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceGroup successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := resourceManagerService.DeleteResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceGroupOptions model
				deleteResourceGroupOptionsModel := new(resourcemanagerv2.DeleteResourceGroupOptions)
				deleteResourceGroupOptionsModel.ID = core.StringPtr("testString")
				deleteResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = resourceManagerService.DeleteResourceGroup(deleteResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				response, operationErr = resourceManagerService.DeleteResourceGroup(deleteResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceGroup with error: Operation validation and request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceGroupOptions model
				deleteResourceGroupOptionsModel := new(resourcemanagerv2.DeleteResourceGroupOptions)
				deleteResourceGroupOptionsModel.ID = core.StringPtr("testString")
				deleteResourceGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := resourceManagerService.DeleteResourceGroup(deleteResourceGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceGroupOptions model with no property values
				deleteResourceGroupOptionsModelNew := new(resourcemanagerv2.DeleteResourceGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = resourceManagerService.DeleteResourceGroup(deleteResourceGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(resourceManagerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(resourceManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "https://resourcemanagerv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(resourceManagerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_URL": "https://resourcemanagerv2/api",
				"RESOURCE_MANAGER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				})
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
					URL: "https://testService/api",
				})
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				})
				err := resourceManagerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_URL": "https://resourcemanagerv2/api",
				"RESOURCE_MANAGER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"RESOURCE_MANAGER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2UsingExternalConfig(&resourcemanagerv2.ResourceManagerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(resourceManagerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListQuotaDefinitions(listQuotaDefinitionsOptions *ListQuotaDefinitionsOptions) - Operation response error`, func() {
		listQuotaDefinitionsPath := "/quota_definitions"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listQuotaDefinitionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListQuotaDefinitions with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the ListQuotaDefinitionsOptions model
				listQuotaDefinitionsOptionsModel := new(resourcemanagerv2.ListQuotaDefinitionsOptions)
				listQuotaDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListQuotaDefinitions(listQuotaDefinitionsOptions *ListQuotaDefinitionsOptions)`, func() {
		listQuotaDefinitionsPath := "/quota_definitions"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listQuotaDefinitionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"resources": [{"id": "ID", "name": "Name", "type": "Type", "number_of_apps": 12, "number_of_service_instances": 24, "default_number_of_instances_per_lite_plan": 35, "instances_per_app": 15, "instance_memory": "InstanceMemory", "total_app_memory": "TotalAppMemory", "vsi_limit": 8, "resource_quotas": [{"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}], "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00"}]}`)
				}))
			})
			It(`Invoke ListQuotaDefinitions successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.ListQuotaDefinitions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListQuotaDefinitionsOptions model
				listQuotaDefinitionsOptionsModel := new(resourcemanagerv2.ListQuotaDefinitionsOptions)
				listQuotaDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.ListQuotaDefinitionsWithContext(ctx, listQuotaDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.ListQuotaDefinitionsWithContext(ctx, listQuotaDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListQuotaDefinitions with error: Operation request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the ListQuotaDefinitionsOptions model
				listQuotaDefinitionsOptionsModel := new(resourcemanagerv2.ListQuotaDefinitionsOptions)
				listQuotaDefinitionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetQuotaDefinition(getQuotaDefinitionOptions *GetQuotaDefinitionOptions) - Operation response error`, func() {
		getQuotaDefinitionPath := "/quota_definitions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQuotaDefinitionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetQuotaDefinition with error: Operation response processing error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the GetQuotaDefinitionOptions model
				getQuotaDefinitionOptionsModel := new(resourcemanagerv2.GetQuotaDefinitionOptions)
				getQuotaDefinitionOptionsModel.ID = core.StringPtr("testString")
				getQuotaDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				resourceManagerService.EnableRetries(0, 0)
				result, response, operationErr = resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetQuotaDefinition(getQuotaDefinitionOptions *GetQuotaDefinitionOptions)`, func() {
		getQuotaDefinitionPath := "/quota_definitions/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getQuotaDefinitionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "type": "Type", "number_of_apps": 12, "number_of_service_instances": 24, "default_number_of_instances_per_lite_plan": 35, "instances_per_app": 15, "instance_memory": "InstanceMemory", "total_app_memory": "TotalAppMemory", "vsi_limit": 8, "resource_quotas": [{"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}], "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetQuotaDefinition successfully`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())
				resourceManagerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := resourceManagerService.GetQuotaDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetQuotaDefinitionOptions model
				getQuotaDefinitionOptionsModel := new(resourcemanagerv2.GetQuotaDefinitionOptions)
				getQuotaDefinitionOptionsModel.ID = core.StringPtr("testString")
				getQuotaDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.GetQuotaDefinitionWithContext(ctx, getQuotaDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				resourceManagerService.DisableRetries()
				result, response, operationErr = resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = resourceManagerService.GetQuotaDefinitionWithContext(ctx, getQuotaDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetQuotaDefinition with error: Operation validation and request error`, func() {
				resourceManagerService, serviceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(resourceManagerService).ToNot(BeNil())

				// Construct an instance of the GetQuotaDefinitionOptions model
				getQuotaDefinitionOptionsModel := new(resourcemanagerv2.GetQuotaDefinitionOptions)
				getQuotaDefinitionOptionsModel.ID = core.StringPtr("testString")
				getQuotaDefinitionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := resourceManagerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetQuotaDefinitionOptions model with no property values
				getQuotaDefinitionOptionsModelNew := new(resourcemanagerv2.GetQuotaDefinitionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = resourceManagerService.GetQuotaDefinition(getQuotaDefinitionOptionsModelNew)
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
			resourceManagerService, _ := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
				URL:           "http://resourcemanagerv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateResourceGroupOptions successfully`, func() {
				// Construct an instance of the CreateResourceGroupOptions model
				createResourceGroupOptionsModel := resourceManagerService.NewCreateResourceGroupOptions()
				createResourceGroupOptionsModel.SetName("test1")
				createResourceGroupOptionsModel.SetAccountID("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceGroupOptionsModel).ToNot(BeNil())
				Expect(createResourceGroupOptionsModel.Name).To(Equal(core.StringPtr("test1")))
				Expect(createResourceGroupOptionsModel.AccountID).To(Equal(core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")))
				Expect(createResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceGroupOptions successfully`, func() {
				// Construct an instance of the DeleteResourceGroupOptions model
				id := "testString"
				deleteResourceGroupOptionsModel := resourceManagerService.NewDeleteResourceGroupOptions(id)
				deleteResourceGroupOptionsModel.SetID("testString")
				deleteResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceGroupOptionsModel).ToNot(BeNil())
				Expect(deleteResourceGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetQuotaDefinitionOptions successfully`, func() {
				// Construct an instance of the GetQuotaDefinitionOptions model
				id := "testString"
				getQuotaDefinitionOptionsModel := resourceManagerService.NewGetQuotaDefinitionOptions(id)
				getQuotaDefinitionOptionsModel.SetID("testString")
				getQuotaDefinitionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getQuotaDefinitionOptionsModel).ToNot(BeNil())
				Expect(getQuotaDefinitionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getQuotaDefinitionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceGroupOptions successfully`, func() {
				// Construct an instance of the GetResourceGroupOptions model
				id := "testString"
				getResourceGroupOptionsModel := resourceManagerService.NewGetResourceGroupOptions(id)
				getResourceGroupOptionsModel.SetID("testString")
				getResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceGroupOptionsModel).ToNot(BeNil())
				Expect(getResourceGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListQuotaDefinitionsOptions successfully`, func() {
				// Construct an instance of the ListQuotaDefinitionsOptions model
				listQuotaDefinitionsOptionsModel := resourceManagerService.NewListQuotaDefinitionsOptions()
				listQuotaDefinitionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listQuotaDefinitionsOptionsModel).ToNot(BeNil())
				Expect(listQuotaDefinitionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceGroupsOptions successfully`, func() {
				// Construct an instance of the ListResourceGroupsOptions model
				listResourceGroupsOptionsModel := resourceManagerService.NewListResourceGroupsOptions()
				listResourceGroupsOptionsModel.SetAccountID("testString")
				listResourceGroupsOptionsModel.SetDate("testString")
				listResourceGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceGroupsOptionsModel).ToNot(BeNil())
				Expect(listResourceGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceGroupsOptionsModel.Date).To(Equal(core.StringPtr("testString")))
				Expect(listResourceGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceGroupOptions successfully`, func() {
				// Construct an instance of the UpdateResourceGroupOptions model
				id := "testString"
				updateResourceGroupOptionsModel := resourceManagerService.NewUpdateResourceGroupOptions(id)
				updateResourceGroupOptionsModel.SetID("testString")
				updateResourceGroupOptionsModel.SetName("testString")
				updateResourceGroupOptionsModel.SetState("testString")
				updateResourceGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceGroupOptionsModel).ToNot(BeNil())
				Expect(updateResourceGroupOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceGroupOptionsModel.State).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
