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
	Describe(`ChangeServiceInstanceState(changeServiceInstanceStateOptions *ChangeServiceInstanceStateOptions) - Operation response error`, func() {
		changeServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ChangeServiceInstanceState with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeServiceInstanceStateOptions model
				changeServiceInstanceStateOptionsModel := new(openservicebrokerv1.ChangeServiceInstanceStateOptions)
				changeServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				changeServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				changeServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				changeServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
				changeServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ChangeServiceInstanceState(changeServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ChangeServiceInstanceState(changeServiceInstanceStateOptions *ChangeServiceInstanceStateOptions)`, func() {
		changeServiceInstanceStatePath := "/bluemix_v1/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(changeServiceInstanceStatePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"active": true, "enabled": false, "last_active": 10}`)
				}))
			})
			It(`Invoke ChangeServiceInstanceState successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ChangeServiceInstanceState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ChangeServiceInstanceStateOptions model
				changeServiceInstanceStateOptionsModel := new(openservicebrokerv1.ChangeServiceInstanceStateOptions)
				changeServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				changeServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				changeServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				changeServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
 				changeServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ChangeServiceInstanceState(changeServiceInstanceStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ChangeServiceInstanceState with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ChangeServiceInstanceStateOptions model
				changeServiceInstanceStateOptionsModel := new(openservicebrokerv1.ChangeServiceInstanceStateOptions)
				changeServiceInstanceStateOptionsModel.InstanceID = core.StringPtr("testString")
				changeServiceInstanceStateOptionsModel.Enabled = core.BoolPtr(false)
				changeServiceInstanceStateOptionsModel.InitiatorID = core.StringPtr("null")
				changeServiceInstanceStateOptionsModel.ReasonCode = core.StringPtr("null")
				changeServiceInstanceStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ChangeServiceInstanceState(changeServiceInstanceStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ChangeServiceInstanceStateOptions model with no property values
				changeServiceInstanceStateOptionsModelNew := new(openservicebrokerv1.ChangeServiceInstanceStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ChangeServiceInstanceState(changeServiceInstanceStateOptionsModelNew)
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
	Describe(`CreateServiceInstance(createServiceInstanceOptions *CreateServiceInstanceOptions) - Operation response error`, func() {
		createServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateServiceInstance with error: Operation response processing error`, func() {
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

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(openservicebrokerv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Context = contextModel
				createServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.Parameters = make(map[string]string)
				createServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				createServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateServiceInstance(createServiceInstanceOptions *CreateServiceInstanceOptions)`, func() {
		createServiceInstancePath := "/v2/service_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createServiceInstancePath))
					Expect(req.Method).To(Equal("PUT"))

					// TODO: Add check for accepts_incomplete query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"dashboard_url": "DashboardURL", "operation": "Operation"}`)
				}))
			})
			It(`Invoke CreateServiceInstance successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(openservicebrokerv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Context = contextModel
				createServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.Parameters = make(map[string]string)
				createServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				createServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
 				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateServiceInstance with error: Operation validation and request error`, func() {
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

				// Construct an instance of the CreateServiceInstanceOptions model
				createServiceInstanceOptionsModel := new(openservicebrokerv1.CreateServiceInstanceOptions)
				createServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceInstanceOptionsModel.Context = contextModel
				createServiceInstanceOptionsModel.OrganizationGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.Parameters = make(map[string]string)
				createServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				createServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				createServiceInstanceOptionsModel.SpaceGuid = core.StringPtr("null")
				createServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				createServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateServiceInstance(createServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateServiceInstanceOptions model with no property values
				createServiceInstanceOptionsModelNew := new(openservicebrokerv1.CreateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateServiceInstance(createServiceInstanceOptionsModelNew)
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

					// TODO: Add check for accepts_incomplete query parameter

					res.WriteHeader(200)
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
				response, operationErr := testService.UpdateServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = contextModel
				updateServiceInstanceOptionsModel.Parameters = make(map[string]string)
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = make(map[string]string)
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
 				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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

				// Construct an instance of the UpdateServiceInstanceOptions model
				updateServiceInstanceOptionsModel := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				updateServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				updateServiceInstanceOptionsModel.Context = contextModel
				updateServiceInstanceOptionsModel.Parameters = make(map[string]string)
				updateServiceInstanceOptionsModel.PlanID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.PreviousValues = make(map[string]string)
				updateServiceInstanceOptionsModel.ServiceID = core.StringPtr("null")
				updateServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
				updateServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.UpdateServiceInstance(updateServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateServiceInstanceOptions model with no property values
				updateServiceInstanceOptionsModelNew := new(openservicebrokerv1.UpdateServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.UpdateServiceInstance(updateServiceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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

					res.WriteHeader(200)
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
				response, operationErr := testService.DeleteServiceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceInstanceOptions model
				deleteServiceInstanceOptionsModel := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				deleteServiceInstanceOptionsModel.ServiceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceInstanceOptionsModel.AcceptsIncomplete = core.BoolPtr(true)
 				deleteServiceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
				response, operationErr := testService.DeleteServiceInstance(deleteServiceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceInstanceOptions model with no property values
				deleteServiceInstanceOptionsModelNew := new(openservicebrokerv1.DeleteServiceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteServiceInstance(deleteServiceInstanceOptionsModelNew)
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
	Describe(`GetCatalog(getCatalogOptions *GetCatalogOptions) - Operation response error`, func() {
		getCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCatalog with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(openservicebrokerv1.GetCatalogOptions)
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetCatalog(getCatalogOptionsModel)
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
		getCatalogPath := "/v2/catalog"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getCatalogPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"services": [{"bindable": true, "description": "Description", "id": "ID", "name": "Name", "plan_updateable": true, "plans": [{"description": "Description", "free": true, "id": "ID", "name": "Name"}]}]}`)
				}))
			})
			It(`Invoke GetCatalog successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCatalog(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(openservicebrokerv1.GetCatalogOptions)
 				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCatalog(getCatalogOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetCatalog with error: Operation request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := new(openservicebrokerv1.GetCatalogOptions)
				getCatalogOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetCatalog(getCatalogOptionsModel)
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
	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions) - Operation response error`, func() {
		getLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLastOperation with error: Operation response processing error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLastOperation(getLastOperationOptions *GetLastOperationOptions)`, func() {
		getLastOperationPath := "/v2/service_instances/testString/last_operation"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getLastOperationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["operation"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["service_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"description": "Description", "state": "State"}`)
				}))
			})
			It(`Invoke GetLastOperation successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetLastOperation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
 				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetLastOperation with error: Operation validation and request error`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetLastOperationOptions model
				getLastOperationOptionsModel := new(openservicebrokerv1.GetLastOperationOptions)
				getLastOperationOptionsModel.InstanceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Operation = core.StringPtr("testString")
				getLastOperationOptionsModel.PlanID = core.StringPtr("testString")
				getLastOperationOptionsModel.ServiceID = core.StringPtr("testString")
				getLastOperationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetLastOperation(getLastOperationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLastOperationOptions model with no property values
				getLastOperationOptionsModelNew := new(openservicebrokerv1.GetLastOperationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetLastOperation(getLastOperationOptionsModelNew)
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
	Describe(`CreateServiceBinding(createServiceBindingOptions *CreateServiceBindingOptions) - Operation response error`, func() {
		createServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createServiceBindingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateServiceBinding with error: Operation response processing error`, func() {
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
				bindResourceModel.AppGuid = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the CreateServiceBindingOptions model
				createServiceBindingOptionsModel := new(openservicebrokerv1.CreateServiceBindingOptions)
				createServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				createServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceBindingOptionsModel.BindResource = bindResourceModel
				createServiceBindingOptionsModel.Parameters = make(map[string]string)
				createServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				createServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				createServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateServiceBinding(createServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateServiceBinding(createServiceBindingOptions *CreateServiceBindingOptions)`, func() {
		createServiceBindingPath := "/v2/service_instances/testString/service_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createServiceBindingPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"credentials": {"anyKey": "anyValue"}, "syslog_drain_url": "SyslogDrainURL", "route_service_url": "RouteServiceURL", "volume_mounts": [{"driver": "Driver", "container_dir": "ContainerDir", "mode": "Mode", "device_type": "DeviceType", "device": "Device"}]}`)
				}))
			})
			It(`Invoke CreateServiceBinding successfully`, func() {
				testService, testServiceErr := openservicebrokerv1.NewOpenServiceBrokerV1(&openservicebrokerv1.OpenServiceBrokerV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")
				bindResourceModel.AppGuid = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the CreateServiceBindingOptions model
				createServiceBindingOptionsModel := new(openservicebrokerv1.CreateServiceBindingOptions)
				createServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				createServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceBindingOptionsModel.BindResource = bindResourceModel
				createServiceBindingOptionsModel.Parameters = make(map[string]string)
				createServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				createServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
 				createServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateServiceBinding(createServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateServiceBinding with error: Operation validation and request error`, func() {
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
				bindResourceModel.AppGuid = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")

				// Construct an instance of the CreateServiceBindingOptions model
				createServiceBindingOptionsModel := new(openservicebrokerv1.CreateServiceBindingOptions)
				createServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				createServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				createServiceBindingOptionsModel.BindResource = bindResourceModel
				createServiceBindingOptionsModel.Parameters = make(map[string]string)
				createServiceBindingOptionsModel.PlanID = core.StringPtr("null")
				createServiceBindingOptionsModel.ServiceID = core.StringPtr("null")
				createServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateServiceBinding(createServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateServiceBindingOptions model with no property values
				createServiceBindingOptionsModelNew := new(openservicebrokerv1.CreateServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateServiceBinding(createServiceBindingOptionsModelNew)
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

					res.WriteHeader(200)
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
				response, operationErr := testService.DeleteServiceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteServiceBindingOptions model
				deleteServiceBindingOptionsModel := new(openservicebrokerv1.DeleteServiceBindingOptions)
				deleteServiceBindingOptionsModel.BindingID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.InstanceID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.PlanID = core.StringPtr("testString")
				deleteServiceBindingOptionsModel.ServiceID = core.StringPtr("testString")
 				deleteServiceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
				response, operationErr := testService.DeleteServiceBinding(deleteServiceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteServiceBindingOptions model with no property values
				deleteServiceBindingOptionsModelNew := new(openservicebrokerv1.DeleteServiceBindingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteServiceBinding(deleteServiceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
			It(`Invoke NewChangeServiceInstanceStateOptions successfully`, func() {
				// Construct an instance of the ChangeServiceInstanceStateOptions model
				instanceID := "testString"
				changeServiceInstanceStateOptionsModel := testService.NewChangeServiceInstanceStateOptions(instanceID)
				changeServiceInstanceStateOptionsModel.SetInstanceID("testString")
				changeServiceInstanceStateOptionsModel.SetEnabled(false)
				changeServiceInstanceStateOptionsModel.SetInitiatorID("null")
				changeServiceInstanceStateOptionsModel.SetReasonCode("null")
				changeServiceInstanceStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(changeServiceInstanceStateOptionsModel).ToNot(BeNil())
				Expect(changeServiceInstanceStateOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(changeServiceInstanceStateOptionsModel.Enabled).To(Equal(core.BoolPtr(false)))
				Expect(changeServiceInstanceStateOptionsModel.InitiatorID).To(Equal(core.StringPtr("null")))
				Expect(changeServiceInstanceStateOptionsModel.ReasonCode).To(Equal(core.StringPtr("null")))
				Expect(changeServiceInstanceStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateServiceBindingOptions successfully`, func() {
				// Construct an instance of the BindResource model
				bindResourceModel := new(openservicebrokerv1.BindResource)
				Expect(bindResourceModel).ToNot(BeNil())
				bindResourceModel.AccountID = core.StringPtr("null")
				bindResourceModel.ServiceidCrn = core.StringPtr("null")
				bindResourceModel.TargetCrn = core.StringPtr("null")
				bindResourceModel.AppGuid = core.StringPtr("null")
				bindResourceModel.Route = core.StringPtr("null")
				Expect(bindResourceModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.ServiceidCrn).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.TargetCrn).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.AppGuid).To(Equal(core.StringPtr("null")))
				Expect(bindResourceModel.Route).To(Equal(core.StringPtr("null")))

				// Construct an instance of the CreateServiceBindingOptions model
				bindingID := "testString"
				instanceID := "testString"
				createServiceBindingOptionsModel := testService.NewCreateServiceBindingOptions(bindingID, instanceID)
				createServiceBindingOptionsModel.SetBindingID("testString")
				createServiceBindingOptionsModel.SetInstanceID("testString")
				createServiceBindingOptionsModel.SetBindResource(bindResourceModel)
				createServiceBindingOptionsModel.SetParameters(make(map[string]string))
				createServiceBindingOptionsModel.SetPlanID("null")
				createServiceBindingOptionsModel.SetServiceID("null")
				createServiceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceBindingOptionsModel).ToNot(BeNil())
				Expect(createServiceBindingOptionsModel.BindingID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceBindingOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceBindingOptionsModel.BindResource).To(Equal(bindResourceModel))
				Expect(createServiceBindingOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(createServiceBindingOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(createServiceBindingOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(createServiceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateServiceInstanceOptions successfully`, func() {
				// Construct an instance of the Context model
				contextModel := new(openservicebrokerv1.Context)
				Expect(contextModel).ToNot(BeNil())
				contextModel.AccountID = core.StringPtr("null")
				contextModel.Crn = core.StringPtr("null")
				contextModel.Platform = core.StringPtr("null")
				Expect(contextModel.AccountID).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Crn).To(Equal(core.StringPtr("null")))
				Expect(contextModel.Platform).To(Equal(core.StringPtr("null")))

				// Construct an instance of the CreateServiceInstanceOptions model
				instanceID := "testString"
				createServiceInstanceOptionsModel := testService.NewCreateServiceInstanceOptions(instanceID)
				createServiceInstanceOptionsModel.SetInstanceID("testString")
				createServiceInstanceOptionsModel.SetContext(contextModel)
				createServiceInstanceOptionsModel.SetOrganizationGuid("null")
				createServiceInstanceOptionsModel.SetParameters(make(map[string]string))
				createServiceInstanceOptionsModel.SetPlanID("null")
				createServiceInstanceOptionsModel.SetServiceID("null")
				createServiceInstanceOptionsModel.SetSpaceGuid("null")
				createServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				createServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(createServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(createServiceInstanceOptionsModel.Context).To(Equal(contextModel))
				Expect(createServiceInstanceOptionsModel.OrganizationGuid).To(Equal(core.StringPtr("null")))
				Expect(createServiceInstanceOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(createServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(createServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(createServiceInstanceOptionsModel.SpaceGuid).To(Equal(core.StringPtr("null")))
				Expect(createServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
				Expect(createServiceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
			It(`Invoke NewGetCatalogOptions successfully`, func() {
				// Construct an instance of the GetCatalogOptions model
				getCatalogOptionsModel := testService.NewGetCatalogOptions()
				getCatalogOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCatalogOptionsModel).ToNot(BeNil())
				Expect(getCatalogOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLastOperationOptions successfully`, func() {
				// Construct an instance of the GetLastOperationOptions model
				instanceID := "testString"
				getLastOperationOptionsModel := testService.NewGetLastOperationOptions(instanceID)
				getLastOperationOptionsModel.SetInstanceID("testString")
				getLastOperationOptionsModel.SetOperation("testString")
				getLastOperationOptionsModel.SetPlanID("testString")
				getLastOperationOptionsModel.SetServiceID("testString")
				getLastOperationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLastOperationOptionsModel).ToNot(BeNil())
				Expect(getLastOperationOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.Operation).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.PlanID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.ServiceID).To(Equal(core.StringPtr("testString")))
				Expect(getLastOperationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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

				// Construct an instance of the UpdateServiceInstanceOptions model
				instanceID := "testString"
				updateServiceInstanceOptionsModel := testService.NewUpdateServiceInstanceOptions(instanceID)
				updateServiceInstanceOptionsModel.SetInstanceID("testString")
				updateServiceInstanceOptionsModel.SetContext(contextModel)
				updateServiceInstanceOptionsModel.SetParameters(make(map[string]string))
				updateServiceInstanceOptionsModel.SetPlanID("null")
				updateServiceInstanceOptionsModel.SetPreviousValues(make(map[string]string))
				updateServiceInstanceOptionsModel.SetServiceID("null")
				updateServiceInstanceOptionsModel.SetAcceptsIncomplete(true)
				updateServiceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateServiceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateServiceInstanceOptionsModel.InstanceID).To(Equal(core.StringPtr("testString")))
				Expect(updateServiceInstanceOptionsModel.Context).To(Equal(contextModel))
				Expect(updateServiceInstanceOptionsModel.Parameters).To(Equal(make(map[string]string)))
				Expect(updateServiceInstanceOptionsModel.PlanID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.PreviousValues).To(Equal(make(map[string]string)))
				Expect(updateServiceInstanceOptionsModel.ServiceID).To(Equal(core.StringPtr("null")))
				Expect(updateServiceInstanceOptionsModel.AcceptsIncomplete).To(Equal(core.BoolPtr(true)))
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
