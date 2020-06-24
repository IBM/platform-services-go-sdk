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

package openservicebrokerv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/openservicebrokerv1"
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

var _ = Describe(`OpenServiceBrokerV1`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions) - Operation response error`, func() {
		getServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getServiceInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetServiceInstanceState with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetServiceInstanceState(getServiceInstanceStateOptions *GetServiceInstanceStateOptions)`, func() {
		getServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getServiceInstanceStatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"active": true, "enabled": false, "last_active": 10}`)
				}))
			})
			It(`Invoke GetServiceInstanceState successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
 				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetServiceInstanceState with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetServiceInstanceStateOptions model
				getServiceInstanceStateOptionsModel := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				getServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				getServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetServiceInstanceState(getServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetServiceInstanceStateOptions model with no property values
				getServiceInstanceStateOptionsModelNew := new(openservicebrokerv1.GetServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetServiceInstanceState(getServiceInstanceStateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceState(replaceStateOptions *ReplaceStateOptions) - Operation response error`, func() {
		replaceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(replaceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceState with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ReplaceStateOptions model
				replaceStateOptionsModel := new(openservicebrokerv1.ReplaceStateOptions)
				replaceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceStateOptionsModel.ReasonCode = core.StringPtr("null")
				replaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ReplaceState(replaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceState(replaceStateOptions *ReplaceStateOptions)`, func() {
		replaceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(replaceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"active": true, "enabled": false, "last_active": 10}`)
				}))
			})
			It(`Invoke ReplaceState successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceStateOptions model
				replaceStateOptionsModel := new(openservicebrokerv1.ReplaceStateOptions)
				replaceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceStateOptionsModel.ReasonCode = core.StringPtr("null")
 				replaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceState(replaceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ReplaceState with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ReplaceStateOptions model
				replaceStateOptionsModel := new(openservicebrokerv1.ReplaceStateOptions)
				replaceStateOptionsModel.InstanceID = core.StringPtr("testString")
				replaceStateOptionsModel.Enabled = core.BoolPtr(false)
				replaceStateOptionsModel.InitiatorID = core.StringPtr("null")
				replaceStateOptionsModel.ReasonCode = core.StringPtr("null")
				replaceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ReplaceState(replaceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceStateOptions model with no property values
				replaceStateOptionsModelNew := new(openservicebrokerv1.ReplaceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ReplaceState(replaceStateOptionsModelNew)
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
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions) - Operation response error`, func() {
		replaceServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(replaceServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceServiceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				replaceServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = []openservicebrokerv1.Parameters{*parametersModel}
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceServiceInstance(replaceServiceInstanceOptions *ReplaceServiceInstanceOptions)`, func() {
		replaceServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(replaceServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke ReplaceServiceInstance successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				replaceServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = []openservicebrokerv1.Parameters{*parametersModel}
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
 				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ReplaceServiceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceInstanceOptions model
				replaceServiceInstanceOptionsModel := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				replaceServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				replaceServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.Parameters = []openservicebrokerv1.Parameters{*parametersModel}
				replaceServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				replaceServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				replaceServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ReplaceServiceInstance(replaceServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceInstanceOptions model with no property values
				replaceServiceInstanceOptionsModelNew := new(openservicebrokerv1.ReplaceServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ReplaceServiceInstance(replaceServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateServiceInstance(updateServiceInstanceOptions *UpdateServiceInstanceOptions)`, func() {
		updateServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateServiceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.URL.Query()["accepts_incomplete"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"OperationResponse"`)
				}))
			})
			It(`Invoke UpdateServiceInstance successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				updateServiceInstanceOptionsModel.Parameters = parametersModel
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.StringPtr("testString")
 				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateServiceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = []openservicebrokerv1.Context{*contextModel}
				updateServiceInstanceOptionsModel.Parameters = parametersModel
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = []string{"testString"}
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteServiceInstance(deleteServiceInstanceOptions *DeleteServiceInstanceOptions)`, func() {
		deleteServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteServiceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))


					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"OperationResponse"`)
				}))
			})
			It(`Invoke DeleteServiceInstance successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
 				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteServiceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
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
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListCatalog(listCatalogOptions *ListCatalogOptions) - Operation response error`, func() {
		listCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCatalog with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListCatalog(listCatalogOptions *ListCatalogOptions)`, func() {
		listCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `[{"bindable": true, "description": "Description", "id": "ID", "name": "Name", "plan_updateable": true, "plans": [{"description": "Description", "free": true, "id": "ID", "name": "Name"}]}]`)
				}))
			})
			It(`Invoke ListCatalog successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
 				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListCatalog(listCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListCatalog with error: Operation request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := new(openservicebrokerv1.ListCatalogOptions)
				listCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListCatalog(listCatalogOptionsModel)
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
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListLastOperation(listLastOperationOptions *ListLastOperationOptions) - Operation response error`, func() {
		listLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListLastOperation with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListLastOperationOptions model
				listLastOperationOptionsModel := new(openservicebrokerv1.ListLastOperationOptions)
				listLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Operation = core.StringPtr("testString")
				listLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				listLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListLastOperation(listLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListLastOperation(listLastOperationOptions *ListLastOperationOptions)`, func() {
		listLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"description": "Description", "state": "State"}`)
				}))
			})
			It(`Invoke ListLastOperation successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListLastOperationOptions model
				listLastOperationOptionsModel := new(openservicebrokerv1.ListLastOperationOptions)
				listLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Operation = core.StringPtr("testString")
				listLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				listLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
 				listLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListLastOperation(listLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListLastOperation with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListLastOperationOptions model
				listLastOperationOptionsModel := new(openservicebrokerv1.ListLastOperationOptions)
				listLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Operation = core.StringPtr("testString")
				listLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				listLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				listLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListLastOperation(listLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListLastOperationOptions model with no property values
				listLastOperationOptionsModelNew := new(openservicebrokerv1.ListLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListLastOperation(listLastOperationOptionsModelNew)
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
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "https://openservicebrokerv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_URL": "https://openservicebrokerv1/api",
				"OPEN_SERVICE_BROKER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"OPEN_SERVICE_BROKER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1UsingExternalConfig(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`ReplaceServiceBinding(replaceServiceBindingOptions *ReplaceServiceBindingOptions)`, func() {
		replaceServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(replaceServiceBindingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"OperationResponse"`)
				}))
			})
			It(`Invoke ReplaceServiceBinding successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = []openservicebrokerv1.BindResource{*bindResourceModel}
				replaceServiceBindingOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
 				replaceServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ReplaceServiceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")

				// Construct an instance of the ReplaceServiceBindingOptions model
				replaceServiceBindingOptionsModel := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				replaceServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				replaceServiceBindingOptionsModel.BindResource = []openservicebrokerv1.BindResource{*bindResourceModel}
				replaceServiceBindingOptionsModel.Parameters = map[string]interface{}{"anyKey": "anyValue"}
				replaceServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				replaceServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ReplaceServiceBinding(replaceServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceServiceBindingOptions model with no property values
				replaceServiceBindingOptionsModelNew := new(openservicebrokerv1.ReplaceServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ReplaceServiceBinding(replaceServiceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteServiceBinding(deleteServiceBindingOptions *DeleteServiceBindingOptions)`, func() {
		deleteServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteServiceBindingPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `"OperationResponse"`)
				}))
			})
			It(`Invoke DeleteServiceBinding successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")
 				deleteServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke DeleteServiceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteServiceBindingOptions model with no property values
				deleteServiceBindingOptionsModelNew := new(openservicebrokerv1.DeleteServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.DeleteServiceBinding(deleteServiceBindingOptionsModelNew)
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
			testService, _ := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
				URL:           "http://openservicebrokerv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewDeleteServiceBindingOptions successfully`, func() {
				// Construct an instance of the DeleteServiceBindingOptions model
				bindingID := "testString"
				instanceID := "testString"
				planID := "testString"
				serviceID := "testString"
				deleteServiceBindingOptionsModel := testService.NewDeleteServiceBindingOptions(bindingID, instanceID, planID, serviceID)
				deleteServiceBindingOptionsModel.SetBindingID("testString")
				deleteServiceBindingOptionsModel.SetInstanceID("testString")
				deleteServiceBindingOptionsModel.SetPlanID("testString")
				deleteServiceBindingOptionsModel.SetServiceID("testString")
				deleteServiceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceBindingOptionsModel).ToNot(BeNil())
				Expect(deleteServiceBindingOptionsModel.BindingID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteServiceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteServiceInstanceOptions model
				serviceID := "testString"
				planID := "testString"
				instanceID := "testString"
				deleteServiceInstanceOptionsModel := testService.NewDeleteServiceInstanceOptions(serviceID, planID, instanceID)
				deleteServiceInstanceOptionsModel.SetServiceID("testString")
				deleteServiceInstanceOptionsModel.SetPlanID("testString")
				deleteServiceInstanceOptionsModel.SetInstanceID("testString")
				deleteServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				deleteServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(deleteServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(deleteServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the GetServiceInstanceStateOptions model
				instanceID := "testString"
				getServiceInstanceStateOptionsModel := testService.NewGetServiceInstanceStateOptions(instanceID)
				getServiceInstanceStateOptionsModel.SetInstanceID("testString")
				getServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(getServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCatalogOptions successfully`, func() {
				// Construct an instance of the ListCatalogOptions model
				listCatalogOptionsModel := testService.NewListCatalogOptions()
				listCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCatalogOptionsModel).ToNot(BeNil())
				Expect(listCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListLastOperationOptions successfully`, func() {
				// Construct an instance of the ListLastOperationOptions model
				instanceID := "testString"
				listLastOperationOptionsModel := testService.NewListLastOperationOptions(instanceID)
				listLastOperationOptionsModel.SetInstanceID("testString")
				listLastOperationOptionsModel.SetOperation("testString")
				listLastOperationOptionsModel.SetPlanID("testString")
				listLastOperationOptionsModel.SetServiceID("testString")
				listLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listLastOperationOptionsModel).ToNot(BeNil())
				Expect(listLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listLastOperationOptionsModel.Operation).To(Equal(core.StringPtr("testString")))
				Expect(listLastOperationOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(listLastOperationOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(listLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceBindingOptions successfully`, func() {
				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				Expect(bindResourceModel).ToNot(BeNil())
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")
				Expect(bindResourceModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.ServiceidCrn).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.TargetCrn).To(Equal(core.StringPtr("null")))

				// Construct an instance of the ReplaceServiceBindingOptions model
				bindingID := "testString"
				instanceID := "testString"
				replaceServiceBindingOptionsModel := testService.NewReplaceServiceBindingOptions(bindingID, instanceID)
				replaceServiceBindingOptionsModel.SetBindingID("testString")
				replaceServiceBindingOptionsModel.SetInstanceID("testString")
				replaceServiceBindingOptionsModel.SetBindResource([]openservicebrokerv1.BindResource{*bindResourceModel})
				replaceServiceBindingOptionsModel.SetParameters(map[string]interface{}{"anyKey": "anyValue"})
				replaceServiceBindingOptionsModel.SetPlanID("null")
				replaceServiceBindingOptionsModel.SetServiceID("null")
				replaceServiceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceBindingOptionsModel).ToNot(BeNil())
				Expect(replaceServiceBindingOptionsModel.BindingID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceBindingOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceBindingOptionsModel.BindResource).To(Equal([]openservicebrokerv1.BindResource{*bindResourceModel}))
				Expect(replaceServiceBindingOptionsModel.Parameters).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(replaceServiceBindingOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceBindingOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceServiceInstanceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")
				Expect(contextModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Crn).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Platform).To(Equal(core.StringPtr("null")))

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				Expect(parametersModel).ToNot(BeNil())
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")
				Expect(parametersModel.Parameter1).To(Equal(core.Int64Ptr(int64(38))))
				Expect(parametersModel.Parameter2).To(Equal(core.StringPtr("null")))

				// Construct an instance of the ReplaceServiceInstanceOptions model
				instanceID := "testString"
				replaceServiceInstanceOptionsModel := testService.NewReplaceServiceInstanceOptions(instanceID)
				replaceServiceInstanceOptionsModel.SetInstanceID("testString")
				replaceServiceInstanceOptionsModel.SetContext([]openservicebrokerv1.Context{*contextModel})
				replaceServiceInstanceOptionsModel.SetOrganizationGuid("null")
				replaceServiceInstanceOptionsModel.SetParameters([]openservicebrokerv1.Parameters{*parametersModel})
				replaceServiceInstanceOptionsModel.SetPlanID("null")
				replaceServiceInstanceOptionsModel.SetServiceID("null")
				replaceServiceInstanceOptionsModel.SetSpaceGuid("null")
				replaceServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				replaceServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(replaceServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceServiceInstanceOptionsModel.Context).To(Equal([]openservicebrokerv1.Context{*contextModel}))
				Expect(replaceServiceInstanceOptionsModel.OrganizationGuid).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.Parameters).To(Equal([]openservicebrokerv1.Parameters{*parametersModel}))
				Expect(replaceServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.SpaceGuid).To(Equal(core.StringPtr("null")))
				Expect(replaceServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(replaceServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceStateOptions successfully`, func() {
				// Construct an instance of the ReplaceStateOptions model
				instanceID := "testString"
				replaceStateOptionsModel := testService.NewReplaceStateOptions(instanceID)
				replaceStateOptionsModel.SetInstanceID("testString")
				replaceStateOptionsModel.SetEnabled(false)
				replaceStateOptionsModel.SetInitiatorID("null")
				replaceStateOptionsModel.SetReasonCode("null")
				replaceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceStateOptionsModel).ToNot(BeNil())
				Expect(replaceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(replaceStateOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(replaceStateOptionsModel.InitiatorID).To(Equal(core.StringPtr("null")))
				Expect(replaceStateOptionsModel.ReasonCode).To(Equal(core.StringPtr("null")))
				Expect(replaceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")
				Expect(contextModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Crn).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Platform).To(Equal(core.StringPtr("null")))

				// Construct an instance of the Parameters model
				parametersModel := new(openservicebrokerv1.Parameters)
				Expect(parametersModel).ToNot(BeNil())
				parametersModel.Parameter1 = core.Int64Ptr(int64(38))
				parametersModel.Parameter2 = core.StringPtr("null")
				Expect(parametersModel.Parameter1).To(Equal(core.Int64Ptr(int64(38))))
				Expect(parametersModel.Parameter2).To(Equal(core.StringPtr("null")))

				// Construct an instance of the UpdateServiceInstanceOptions model
				instanceID := "testString"
				updateServiceInstanceOptionsModel := testService.NewUpdateServiceInstanceOptions(instanceID)
				updateServiceInstanceOptionsModel.SetInstanceID("testString")
				updateServiceInstanceOptionsModel.SetContext([]openservicebrokerv1.Context{*contextModel})
				updateServiceInstanceOptionsModel.SetParameters(parametersModel)
				updateServiceInstanceOptionsModel.SetPlanID("null")
				updateServiceInstanceOptionsModel.SetPreviousValues([]string{"testString"})
				updateServiceInstanceOptionsModel.SetServiceID("null")
				updateServiceInstanceOptionsModel.SetAcceptsIncomplete("testString")
				updateServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceOptionsModel.Context).To(Equal([]openservicebrokerv1.Context{*contextModel}))
				Expect(updateServiceInstanceOptionsModel.Parameters).To(Equal(parametersModel))
				Expect(updateServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.PreviousValues).To(Equal([]string{"testString"}))
				Expect(updateServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
