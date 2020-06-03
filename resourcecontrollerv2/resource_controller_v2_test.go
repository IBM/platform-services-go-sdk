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

package resourcecontrollerv2_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
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

var _ = Describe(`ResourceControllerV2`, func() {
	var testServer *httptest.Server
    Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions) - Operation response error`, func() {
		listResourceInstancesPath := "/v2/resource_instances"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceInstances with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.Guid = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
		listResourceInstancesPath := "/v2/resource_instances"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceInstancesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
				}))
			})
			It(`Invoke ListResourceInstances successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceInstances(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.Guid = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
 				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResourceInstances with error: Operation request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := new(resourcecontrollerv2.ListResourceInstancesOptions)
				listResourceInstancesOptionsModel.Guid = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Name = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.ResourcePlanID = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Type = core.StringPtr("testString")
				listResourceInstancesOptionsModel.SubType = core.StringPtr("testString")
				listResourceInstancesOptionsModel.Limit = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResourceInstances(listResourceInstancesOptionsModel)
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
	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions) - Operation response error`, func() {
		createResourceInstancePath := "/v2/resource_instances"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("my-instance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("bluemix-us-south")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("5c49eabc-f5e8-5881-a37e-2d100a33b3df")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("cloudant-standard")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				createResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				createResourceInstanceOptionsModel.EntityLock = core.StringPtr("testString")
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
		createResourceInstancePath := "/v2/resource_instances"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Entity-Lock"]).ToNot(BeNil())
					Expect(req.Header["Entity-Lock"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("my-instance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("bluemix-us-south")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("5c49eabc-f5e8-5881-a37e-2d100a33b3df")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("cloudant-standard")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				createResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				createResourceInstanceOptionsModel.EntityLock = core.StringPtr("testString")
 				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsModel := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				createResourceInstanceOptionsModel.Name = core.StringPtr("my-instance")
				createResourceInstanceOptionsModel.Target = core.StringPtr("bluemix-us-south")
				createResourceInstanceOptionsModel.ResourceGroup = core.StringPtr("5c49eabc-f5e8-5881-a37e-2d100a33b3df")
				createResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("cloudant-standard")
				createResourceInstanceOptionsModel.Tags = []string{"testString"}
				createResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				createResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				createResourceInstanceOptionsModel.EntityLock = core.StringPtr("testString")
				createResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceInstanceOptions model with no property values
				createResourceInstanceOptionsModelNew := new(resourcecontrollerv2.CreateResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResourceInstance(createResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions) - Operation response error`, func() {
		getResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceInstancePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
		getResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceInstancePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
 				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceInstanceOptions model
				getResourceInstanceOptionsModel := new(resourcecontrollerv2.GetResourceInstanceOptions)
				getResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				getResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceInstanceOptions model with no property values
				getResourceInstanceOptionsModelNew := new(resourcecontrollerv2.GetResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResourceInstance(getResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
		deleteResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")
 				deleteResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				deleteResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceInstanceOptions model with no property values
				deleteResourceInstanceOptionsModelNew := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteResourceInstance(deleteResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions) - Operation response error`, func() {
		updateResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("my-new-instance-name")
				updateResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
		updateResourceInstancePath := "/v2/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceInstancePath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("my-new-instance-name")
				updateResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
 				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceInstanceOptions model
				updateResourceInstanceOptionsModel := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				updateResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				updateResourceInstanceOptionsModel.Name = core.StringPtr("my-new-instance-name")
				updateResourceInstanceOptionsModel.Parameters = make(map[string]interface{})
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)
				updateResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceInstanceOptions model with no property values
				updateResourceInstanceOptionsModelNew := new(resourcecontrollerv2.UpdateResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateResourceInstance(updateResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`LockResourceInstance(lockResourceInstanceOptions *LockResourceInstanceOptions) - Operation response error`, func() {
		lockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(lockResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke LockResourceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`LockResourceInstance(lockResourceInstanceOptions *LockResourceInstanceOptions)`, func() {
		lockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(lockResourceInstancePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke LockResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.LockResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
 				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke LockResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the LockResourceInstanceOptions model
				lockResourceInstanceOptionsModel := new(resourcecontrollerv2.LockResourceInstanceOptions)
				lockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				lockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.LockResourceInstance(lockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the LockResourceInstanceOptions model with no property values
				lockResourceInstanceOptionsModelNew := new(resourcecontrollerv2.LockResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.LockResourceInstance(lockResourceInstanceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UnlockResourceInstance(unlockResourceInstanceOptions *UnlockResourceInstanceOptions) - Operation response error`, func() {
		unlockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(unlockResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UnlockResourceInstance with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UnlockResourceInstance(unlockResourceInstanceOptions *UnlockResourceInstanceOptions)`, func() {
		unlockResourceInstancePath := "/v2/resource_instances/testString/lock"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(unlockResourceInstancePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "locked": true, "last_operation": {"mapKey": "anyValue"}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UnlockResourceInstance successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UnlockResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
 				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UnlockResourceInstance with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UnlockResourceInstanceOptions model
				unlockResourceInstanceOptionsModel := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				unlockResourceInstanceOptionsModel.ID = core.StringPtr("testString")
				unlockResourceInstanceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UnlockResourceInstance(unlockResourceInstanceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UnlockResourceInstanceOptions model with no property values
				unlockResourceInstanceOptionsModelNew := new(resourcecontrollerv2.UnlockResourceInstanceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UnlockResourceInstance(unlockResourceInstanceOptionsModelNew)
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
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions) - Operation response error`, func() {
		listResourceKeysPath := "/v2/resource_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceKeys with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.Guid = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions)`, func() {
		listResourceKeysPath := "/v2/resource_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceKeysPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
				}))
			})
			It(`Invoke ListResourceKeys successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceKeys(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.Guid = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
 				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResourceKeys with error: Operation request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := new(resourcecontrollerv2.ListResourceKeysOptions)
				listResourceKeysOptionsModel.Guid = core.StringPtr("testString")
				listResourceKeysOptionsModel.Name = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceKeysOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceKeysOptionsModel.Limit = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResourceKeys(listResourceKeysOptionsModel)
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
	Describe(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions) - Operation response error`, func() {
		createResourceKeyPath := "/v2/resource_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceKeyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceKey with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("my-key")
				createResourceKeyOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions)`, func() {
		createResourceKeyPath := "/v2/resource_keys"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceKeyPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateResourceKey successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("my-key")
				createResourceKeyOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
 				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResourceKey with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				resourceKeyPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsModel := new(resourcecontrollerv2.CreateResourceKeyOptions)
				createResourceKeyOptionsModel.Name = core.StringPtr("my-key")
				createResourceKeyOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceKeyOptionsModel.Parameters = resourceKeyPostParametersModel
				createResourceKeyOptionsModel.Role = core.StringPtr("Writer")
				createResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceKeyOptions model with no property values
				createResourceKeyOptionsModelNew := new(resourcecontrollerv2.CreateResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResourceKey(createResourceKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions) - Operation response error`, func() {
		getResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceKeyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceKey with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions)`, func() {
		getResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceKeyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetResourceKey successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
 				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResourceKey with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceKeyOptions model
				getResourceKeyOptionsModel := new(resourcecontrollerv2.GetResourceKeyOptions)
				getResourceKeyOptionsModel.ID = core.StringPtr("testString")
				getResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceKeyOptions model with no property values
				getResourceKeyOptionsModelNew := new(resourcecontrollerv2.GetResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResourceKey(getResourceKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceKey(deleteResourceKeyOptions *DeleteResourceKeyOptions)`, func() {
		deleteResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteResourceKeyPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceKey successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceKeyOptions model
				deleteResourceKeyOptionsModel := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				deleteResourceKeyOptionsModel.ID = core.StringPtr("testString")
 				deleteResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceKey(deleteResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceKey with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceKeyOptions model
				deleteResourceKeyOptionsModel := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				deleteResourceKeyOptionsModel.ID = core.StringPtr("testString")
				deleteResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteResourceKey(deleteResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceKeyOptions model with no property values
				deleteResourceKeyOptionsModelNew := new(resourcecontrollerv2.DeleteResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteResourceKey(deleteResourceKeyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions) - Operation response error`, func() {
		updateResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceKeyPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceKey with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("my-new-key-name")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions)`, func() {
		updateResourceKeyPath := "/v2/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceKeyPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateResourceKey successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceKey(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("my-new-key-name")
 				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateResourceKey with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceKeyOptions model
				updateResourceKeyOptionsModel := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				updateResourceKeyOptionsModel.ID = core.StringPtr("testString")
				updateResourceKeyOptionsModel.Name = core.StringPtr("my-new-key-name")
				updateResourceKeyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceKeyOptions model with no property values
				updateResourceKeyOptionsModelNew := new(resourcecontrollerv2.UpdateResourceKeyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateResourceKey(updateResourceKeyOptionsModelNew)
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
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions) - Operation response error`, func() {
		listResourceBindingsPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceBindings with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.Guid = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions)`, func() {
		listResourceBindingsPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceBindingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
				}))
			})
			It(`Invoke ListResourceBindings successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceBindings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.Guid = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
 				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResourceBindings with error: Operation request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := new(resourcecontrollerv2.ListResourceBindingsOptions)
				listResourceBindingsOptionsModel.Guid = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Name = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.RegionBindingID = core.StringPtr("testString")
				listResourceBindingsOptionsModel.Limit = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResourceBindings(listResourceBindingsOptionsModel)
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
	Describe(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions) - Operation response error`, func() {
		createResourceBindingPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceBindingPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceBinding with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")
				createResourceBindingOptionsModel.Name = core.StringPtr("my-binding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions)`, func() {
		createResourceBindingPath := "/v2/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceBindingPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateResourceBinding successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")
				createResourceBindingOptionsModel.Name = core.StringPtr("my-binding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
 				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResourceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				resourceBindingPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsModel := new(resourcecontrollerv2.CreateResourceBindingOptions)
				createResourceBindingOptionsModel.Source = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")
				createResourceBindingOptionsModel.Name = core.StringPtr("my-binding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")
				createResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceBindingOptions model with no property values
				createResourceBindingOptionsModelNew := new(resourcecontrollerv2.CreateResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResourceBinding(createResourceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions) - Operation response error`, func() {
		getResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceBindingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceBinding with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions)`, func() {
		getResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceBindingPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetResourceBinding successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
 				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResourceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceBindingOptions model
				getResourceBindingOptionsModel := new(resourcecontrollerv2.GetResourceBindingOptions)
				getResourceBindingOptionsModel.ID = core.StringPtr("testString")
				getResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceBindingOptions model with no property values
				getResourceBindingOptionsModelNew := new(resourcecontrollerv2.GetResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResourceBinding(getResourceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceBinding(deleteResourceBindingOptions *DeleteResourceBindingOptions)`, func() {
		deleteResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteResourceBindingPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceBinding successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceBindingOptions model
				deleteResourceBindingOptionsModel := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				deleteResourceBindingOptionsModel.ID = core.StringPtr("testString")
 				deleteResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceBinding(deleteResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceBindingOptions model
				deleteResourceBindingOptionsModel := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				deleteResourceBindingOptionsModel.ID = core.StringPtr("testString")
				deleteResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteResourceBinding(deleteResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceBindingOptions model with no property values
				deleteResourceBindingOptionsModelNew := new(resourcecontrollerv2.DeleteResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteResourceBinding(deleteResourceBindingOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions) - Operation response error`, func() {
		updateResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceBindingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceBinding with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("my-new-binding-name")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions)`, func() {
		updateResourceBindingPath := "/v2/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceBindingPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateResourceBinding successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceBinding(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("my-new-binding-name")
 				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateResourceBinding with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceBindingOptions model
				updateResourceBindingOptionsModel := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				updateResourceBindingOptionsModel.ID = core.StringPtr("testString")
				updateResourceBindingOptionsModel.Name = core.StringPtr("my-new-binding-name")
				updateResourceBindingOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceBindingOptions model with no property values
				updateResourceBindingOptionsModelNew := new(resourcecontrollerv2.UpdateResourceBindingOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateResourceBinding(updateResourceBindingOptionsModelNew)
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
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions) - Operation response error`, func() {
		listResourceAliasesPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListResourceAliases with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.Guid = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions)`, func() {
		listResourceAliasesPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listResourceAliasesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"2019-01-08T00:00:00.000Z"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
				}))
			})
			It(`Invoke ListResourceAliases successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceAliases(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.Guid = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
 				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListResourceAliases with error: Operation request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := new(resourcecontrollerv2.ListResourceAliasesOptions)
				listResourceAliasesOptionsModel.Guid = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Name = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.RegionInstanceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.ResourceGroupID = core.StringPtr("testString")
				listResourceAliasesOptionsModel.Limit = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListResourceAliases(listResourceAliasesOptionsModel)
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
	Describe(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions) - Operation response error`, func() {
		createResourceAliasPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceAliasPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateResourceAlias with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("my-alias")
				createResourceAliasOptionsModel.Source = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions)`, func() {
		createResourceAliasPath := "/v2/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createResourceAliasPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke CreateResourceAlias successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("my-alias")
				createResourceAliasOptionsModel.Source = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")
 				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateResourceAlias with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsModel := new(resourcecontrollerv2.CreateResourceAliasOptions)
				createResourceAliasOptionsModel.Name = core.StringPtr("my-alias")
				createResourceAliasOptionsModel.Source = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")
				createResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateResourceAliasOptions model with no property values
				createResourceAliasOptionsModelNew := new(resourcecontrollerv2.CreateResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateResourceAlias(createResourceAliasOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions) - Operation response error`, func() {
		getResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceAliasPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetResourceAlias with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions)`, func() {
		getResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getResourceAliasPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke GetResourceAlias successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
 				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetResourceAlias with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetResourceAliasOptions model
				getResourceAliasOptionsModel := new(resourcecontrollerv2.GetResourceAliasOptions)
				getResourceAliasOptionsModel.ID = core.StringPtr("testString")
				getResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetResourceAliasOptions model with no property values
				getResourceAliasOptionsModelNew := new(resourcecontrollerv2.GetResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetResourceAlias(getResourceAliasOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteResourceAlias(deleteResourceAliasOptions *DeleteResourceAliasOptions)`, func() {
		deleteResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(deleteResourceAliasPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteResourceAlias successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceAliasOptions model
				deleteResourceAliasOptionsModel := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				deleteResourceAliasOptionsModel.ID = core.StringPtr("testString")
 				deleteResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceAlias(deleteResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteResourceAlias with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteResourceAliasOptions model
				deleteResourceAliasOptionsModel := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				deleteResourceAliasOptionsModel.ID = core.StringPtr("testString")
				deleteResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteResourceAlias(deleteResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteResourceAliasOptions model with no property values
				deleteResourceAliasOptionsModelNew := new(resourcecontrollerv2.DeleteResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteResourceAlias(deleteResourceAliasOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions) - Operation response error`, func() {
		updateResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceAliasPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateResourceAlias with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("my-new-alias-name")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions)`, func() {
		updateResourceAliasPath := "/v2/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateResourceAliasPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
				}))
			})
			It(`Invoke UpdateResourceAlias successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceAlias(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("my-new-alias-name")
 				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateResourceAlias with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateResourceAliasOptions model
				updateResourceAliasOptionsModel := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				updateResourceAliasOptionsModel.ID = core.StringPtr("testString")
				updateResourceAliasOptionsModel.Name = core.StringPtr("my-new-alias-name")
				updateResourceAliasOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateResourceAliasOptions model with no property values
				updateResourceAliasOptionsModelNew := new(resourcecontrollerv2.UpdateResourceAliasOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateResourceAlias(updateResourceAliasOptionsModelNew)
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
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "https://resourcecontrollerv2/api",
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_URL": "https://resourcecontrollerv2/api",
				"RESOURCE_CONTROLLER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
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
				"RESOURCE_CONTROLLER_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2UsingExternalConfig(&resourcecontrollerv2.ResourceControllerV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions) - Operation response error`, func() {
		listReclamationsPath := "/v1/reclamations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListReclamations with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions)`, func() {
		listReclamationsPath := "/v1/reclamations"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listReclamationsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"resources": [{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCrn", "resource_instance_id": "anyValue", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": "CustomProperties", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
				}))
			})
			It(`Invoke ListReclamations successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListReclamations(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
 				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListReclamations with error: Operation request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := new(resourcecontrollerv2.ListReclamationsOptions)
				listReclamationsOptionsModel.AccountID = core.StringPtr("testString")
				listReclamationsOptionsModel.ResourceInstanceID = core.StringPtr("testString")
				listReclamationsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListReclamations(listReclamationsOptionsModel)
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
	Describe(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions) - Operation response error`, func() {
		runReclamationActionPath := "/v1/reclamations/testString/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runReclamationActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RunReclamationAction with error: Operation response processing error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions)`, func() {
		runReclamationActionPath := "/v1/reclamations/testString/actions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(runReclamationActionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCrn", "resource_instance_id": "anyValue", "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": "CustomProperties", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}`)
				}))
			})
			It(`Invoke RunReclamationAction successfully`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RunReclamationAction(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
 				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke RunReclamationAction with error: Operation validation and request error`, func() {
				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the RunReclamationActionOptions model
				runReclamationActionOptionsModel := new(resourcecontrollerv2.RunReclamationActionOptions)
				runReclamationActionOptionsModel.ID = core.StringPtr("testString")
				runReclamationActionOptionsModel.ActionName = core.StringPtr("testString")
				runReclamationActionOptionsModel.RequestBy = core.StringPtr("testString")
				runReclamationActionOptionsModel.Comment = core.StringPtr("testString")
				runReclamationActionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RunReclamationActionOptions model with no property values
				runReclamationActionOptionsModelNew := new(resourcecontrollerv2.RunReclamationActionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.RunReclamationAction(runReclamationActionOptionsModelNew)
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
			testService, _ := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
				URL:           "http://resourcecontrollerv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreateResourceAliasOptions successfully`, func() {
				// Construct an instance of the CreateResourceAliasOptions model
				createResourceAliasOptionsName := "my-alias"
				createResourceAliasOptionsSource := "a8dff6d3-d287-4668-a81d-c87c55c2656d"
				createResourceAliasOptionsTarget := "crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5"
				createResourceAliasOptionsModel := testService.NewCreateResourceAliasOptions(createResourceAliasOptionsName, createResourceAliasOptionsSource, createResourceAliasOptionsTarget)
				createResourceAliasOptionsModel.SetName("my-alias")
				createResourceAliasOptionsModel.SetSource("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				createResourceAliasOptionsModel.SetTarget("crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")
				createResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceAliasOptionsModel).ToNot(BeNil())
				Expect(createResourceAliasOptionsModel.Name).To(Equal(core.StringPtr("my-alias")))
				Expect(createResourceAliasOptionsModel.Source).To(Equal(core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")))
				Expect(createResourceAliasOptionsModel.Target).To(Equal(core.StringPtr("crn:v1:staging:public:cf:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")))
				Expect(createResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceBindingOptions successfully`, func() {
				// Construct an instance of the ResourceBindingPostParameters model
				resourceBindingPostParametersModel := new(resourcecontrollerv2.ResourceBindingPostParameters)
				Expect(resourceBindingPostParametersModel).ToNot(BeNil())
				resourceBindingPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				Expect(resourceBindingPostParametersModel.ServiceidCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")))

				// Construct an instance of the CreateResourceBindingOptions model
				createResourceBindingOptionsSource := "25eba2a9-beef-450b-82cf-f5ad5e36c6dd"
				createResourceBindingOptionsTarget := "crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c"
				createResourceBindingOptionsModel := testService.NewCreateResourceBindingOptions(createResourceBindingOptionsSource, createResourceBindingOptionsTarget)
				createResourceBindingOptionsModel.SetSource("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceBindingOptionsModel.SetTarget("crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")
				createResourceBindingOptionsModel.SetName("my-binding")
				createResourceBindingOptionsModel.SetParameters(resourceBindingPostParametersModel)
				createResourceBindingOptionsModel.SetRole("Writer")
				createResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceBindingOptionsModel).ToNot(BeNil())
				Expect(createResourceBindingOptionsModel.Source).To(Equal(core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")))
				Expect(createResourceBindingOptionsModel.Target).To(Equal(core.StringPtr("crn:v1:bluemix:public:cf:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")))
				Expect(createResourceBindingOptionsModel.Name).To(Equal(core.StringPtr("my-binding")))
				Expect(createResourceBindingOptionsModel.Parameters).To(Equal(resourceBindingPostParametersModel))
				Expect(createResourceBindingOptionsModel.Role).To(Equal(core.StringPtr("Writer")))
				Expect(createResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceInstanceOptions successfully`, func() {
				// Construct an instance of the CreateResourceInstanceOptions model
				createResourceInstanceOptionsName := "my-instance"
				createResourceInstanceOptionsTarget := "bluemix-us-south"
				createResourceInstanceOptionsResourceGroup := "5c49eabc-f5e8-5881-a37e-2d100a33b3df"
				createResourceInstanceOptionsResourcePlanID := "cloudant-standard"
				createResourceInstanceOptionsModel := testService.NewCreateResourceInstanceOptions(createResourceInstanceOptionsName, createResourceInstanceOptionsTarget, createResourceInstanceOptionsResourceGroup, createResourceInstanceOptionsResourcePlanID)
				createResourceInstanceOptionsModel.SetName("my-instance")
				createResourceInstanceOptionsModel.SetTarget("bluemix-us-south")
				createResourceInstanceOptionsModel.SetResourceGroup("5c49eabc-f5e8-5881-a37e-2d100a33b3df")
				createResourceInstanceOptionsModel.SetResourcePlanID("cloudant-standard")
				createResourceInstanceOptionsModel.SetTags([]string{"testString"})
				createResourceInstanceOptionsModel.SetAllowCleanup(true)
				createResourceInstanceOptionsModel.SetParameters(make(map[string]interface{}))
				createResourceInstanceOptionsModel.SetEntityLock("testString")
				createResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(createResourceInstanceOptionsModel.Name).To(Equal(core.StringPtr("my-instance")))
				Expect(createResourceInstanceOptionsModel.Target).To(Equal(core.StringPtr("bluemix-us-south")))
				Expect(createResourceInstanceOptionsModel.ResourceGroup).To(Equal(core.StringPtr("5c49eabc-f5e8-5881-a37e-2d100a33b3df")))
				Expect(createResourceInstanceOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("cloudant-standard")))
				Expect(createResourceInstanceOptionsModel.Tags).To(Equal([]string{"testString"}))
				Expect(createResourceInstanceOptionsModel.AllowCleanup).To(Equal(core.BoolPtr(true)))
				Expect(createResourceInstanceOptionsModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(createResourceInstanceOptionsModel.EntityLock).To(Equal(core.StringPtr("testString")))
				Expect(createResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateResourceKeyOptions successfully`, func() {
				// Construct an instance of the ResourceKeyPostParameters model
				resourceKeyPostParametersModel := new(resourcecontrollerv2.ResourceKeyPostParameters)
				Expect(resourceKeyPostParametersModel).ToNot(BeNil())
				resourceKeyPostParametersModel.ServiceidCrn = core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")
				Expect(resourceKeyPostParametersModel.ServiceidCrn).To(Equal(core.StringPtr("crn:v1:bluemix:public:iam-identity::a/9fceaa56d1ab84893af6b9eec5ab81bb::serviceid:ServiceId-fe4c29b5-db13-410a-bacc-b5779a03d393")))

				// Construct an instance of the CreateResourceKeyOptions model
				createResourceKeyOptionsName := "my-key"
				createResourceKeyOptionsSource := "25eba2a9-beef-450b-82cf-f5ad5e36c6dd"
				createResourceKeyOptionsModel := testService.NewCreateResourceKeyOptions(createResourceKeyOptionsName, createResourceKeyOptionsSource)
				createResourceKeyOptionsModel.SetName("my-key")
				createResourceKeyOptionsModel.SetSource("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")
				createResourceKeyOptionsModel.SetParameters(resourceKeyPostParametersModel)
				createResourceKeyOptionsModel.SetRole("Writer")
				createResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createResourceKeyOptionsModel).ToNot(BeNil())
				Expect(createResourceKeyOptionsModel.Name).To(Equal(core.StringPtr("my-key")))
				Expect(createResourceKeyOptionsModel.Source).To(Equal(core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")))
				Expect(createResourceKeyOptionsModel.Parameters).To(Equal(resourceKeyPostParametersModel))
				Expect(createResourceKeyOptionsModel.Role).To(Equal(core.StringPtr("Writer")))
				Expect(createResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceAliasOptions successfully`, func() {
				// Construct an instance of the DeleteResourceAliasOptions model
				id := "testString"
				deleteResourceAliasOptionsModel := testService.NewDeleteResourceAliasOptions(id)
				deleteResourceAliasOptionsModel.SetID("testString")
				deleteResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceAliasOptionsModel).ToNot(BeNil())
				Expect(deleteResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceBindingOptions successfully`, func() {
				// Construct an instance of the DeleteResourceBindingOptions model
				id := "testString"
				deleteResourceBindingOptionsModel := testService.NewDeleteResourceBindingOptions(id)
				deleteResourceBindingOptionsModel.SetID("testString")
				deleteResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceBindingOptionsModel).ToNot(BeNil())
				Expect(deleteResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceInstanceOptions successfully`, func() {
				// Construct an instance of the DeleteResourceInstanceOptions model
				id := "testString"
				deleteResourceInstanceOptionsModel := testService.NewDeleteResourceInstanceOptions(id)
				deleteResourceInstanceOptionsModel.SetID("testString")
				deleteResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(deleteResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteResourceKeyOptions successfully`, func() {
				// Construct an instance of the DeleteResourceKeyOptions model
				id := "testString"
				deleteResourceKeyOptionsModel := testService.NewDeleteResourceKeyOptions(id)
				deleteResourceKeyOptionsModel.SetID("testString")
				deleteResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteResourceKeyOptionsModel).ToNot(BeNil())
				Expect(deleteResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceAliasOptions successfully`, func() {
				// Construct an instance of the GetResourceAliasOptions model
				id := "testString"
				getResourceAliasOptionsModel := testService.NewGetResourceAliasOptions(id)
				getResourceAliasOptionsModel.SetID("testString")
				getResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceAliasOptionsModel).ToNot(BeNil())
				Expect(getResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceBindingOptions successfully`, func() {
				// Construct an instance of the GetResourceBindingOptions model
				id := "testString"
				getResourceBindingOptionsModel := testService.NewGetResourceBindingOptions(id)
				getResourceBindingOptionsModel.SetID("testString")
				getResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceBindingOptionsModel).ToNot(BeNil())
				Expect(getResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceInstanceOptions successfully`, func() {
				// Construct an instance of the GetResourceInstanceOptions model
				id := "testString"
				getResourceInstanceOptionsModel := testService.NewGetResourceInstanceOptions(id)
				getResourceInstanceOptionsModel.SetID("testString")
				getResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(getResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetResourceKeyOptions successfully`, func() {
				// Construct an instance of the GetResourceKeyOptions model
				id := "testString"
				getResourceKeyOptionsModel := testService.NewGetResourceKeyOptions(id)
				getResourceKeyOptionsModel.SetID("testString")
				getResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getResourceKeyOptionsModel).ToNot(BeNil())
				Expect(getResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListReclamationsOptions successfully`, func() {
				// Construct an instance of the ListReclamationsOptions model
				listReclamationsOptionsModel := testService.NewListReclamationsOptions()
				listReclamationsOptionsModel.SetAccountID("testString")
				listReclamationsOptionsModel.SetResourceInstanceID("testString")
				listReclamationsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listReclamationsOptionsModel).ToNot(BeNil())
				Expect(listReclamationsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listReclamationsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceAliasesOptions successfully`, func() {
				// Construct an instance of the ListResourceAliasesOptions model
				listResourceAliasesOptionsModel := testService.NewListResourceAliasesOptions()
				listResourceAliasesOptionsModel.SetGuid("testString")
				listResourceAliasesOptionsModel.SetName("testString")
				listResourceAliasesOptionsModel.SetResourceInstanceID("testString")
				listResourceAliasesOptionsModel.SetRegionInstanceID("testString")
				listResourceAliasesOptionsModel.SetResourceID("testString")
				listResourceAliasesOptionsModel.SetResourceGroupID("testString")
				listResourceAliasesOptionsModel.SetLimit("testString")
				listResourceAliasesOptionsModel.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.SetUpdatedTo("2019-01-08T00:00:00.000Z")
				listResourceAliasesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceAliasesOptionsModel).ToNot(BeNil())
				Expect(listResourceAliasesOptionsModel.Guid).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.RegionInstanceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.Limit).To(Equal(core.StringPtr("testString")))
				Expect(listResourceAliasesOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceAliasesOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceAliasesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceBindingsOptions successfully`, func() {
				// Construct an instance of the ListResourceBindingsOptions model
				listResourceBindingsOptionsModel := testService.NewListResourceBindingsOptions()
				listResourceBindingsOptionsModel.SetGuid("testString")
				listResourceBindingsOptionsModel.SetName("testString")
				listResourceBindingsOptionsModel.SetResourceGroupID("testString")
				listResourceBindingsOptionsModel.SetResourceID("testString")
				listResourceBindingsOptionsModel.SetRegionBindingID("testString")
				listResourceBindingsOptionsModel.SetLimit("testString")
				listResourceBindingsOptionsModel.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.SetUpdatedTo("2019-01-08T00:00:00.000Z")
				listResourceBindingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceBindingsOptionsModel).ToNot(BeNil())
				Expect(listResourceBindingsOptionsModel.Guid).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.RegionBindingID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.Limit).To(Equal(core.StringPtr("testString")))
				Expect(listResourceBindingsOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceBindingsOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceBindingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceInstancesOptions successfully`, func() {
				// Construct an instance of the ListResourceInstancesOptions model
				listResourceInstancesOptionsModel := testService.NewListResourceInstancesOptions()
				listResourceInstancesOptionsModel.SetGuid("testString")
				listResourceInstancesOptionsModel.SetName("testString")
				listResourceInstancesOptionsModel.SetResourceGroupID("testString")
				listResourceInstancesOptionsModel.SetResourceID("testString")
				listResourceInstancesOptionsModel.SetResourcePlanID("testString")
				listResourceInstancesOptionsModel.SetType("testString")
				listResourceInstancesOptionsModel.SetSubType("testString")
				listResourceInstancesOptionsModel.SetLimit("testString")
				listResourceInstancesOptionsModel.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.SetUpdatedTo("2019-01-08T00:00:00.000Z")
				listResourceInstancesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceInstancesOptionsModel).ToNot(BeNil())
				Expect(listResourceInstancesOptionsModel.Guid).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.SubType).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.Limit).To(Equal(core.StringPtr("testString")))
				Expect(listResourceInstancesOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceInstancesOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceInstancesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListResourceKeysOptions successfully`, func() {
				// Construct an instance of the ListResourceKeysOptions model
				listResourceKeysOptionsModel := testService.NewListResourceKeysOptions()
				listResourceKeysOptionsModel.SetGuid("testString")
				listResourceKeysOptionsModel.SetName("testString")
				listResourceKeysOptionsModel.SetResourceGroupID("testString")
				listResourceKeysOptionsModel.SetResourceID("testString")
				listResourceKeysOptionsModel.SetLimit("testString")
				listResourceKeysOptionsModel.SetUpdatedFrom("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.SetUpdatedTo("2019-01-08T00:00:00.000Z")
				listResourceKeysOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listResourceKeysOptionsModel).ToNot(BeNil())
				Expect(listResourceKeysOptionsModel.Guid).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.ResourceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.ResourceID).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.Limit).To(Equal(core.StringPtr("testString")))
				Expect(listResourceKeysOptionsModel.UpdatedFrom).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceKeysOptionsModel.UpdatedTo).To(Equal(core.StringPtr("2019-01-08T00:00:00.000Z")))
				Expect(listResourceKeysOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockResourceInstanceOptions successfully`, func() {
				// Construct an instance of the LockResourceInstanceOptions model
				id := "testString"
				lockResourceInstanceOptionsModel := testService.NewLockResourceInstanceOptions(id)
				lockResourceInstanceOptionsModel.SetID("testString")
				lockResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(lockResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(lockResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(lockResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRunReclamationActionOptions successfully`, func() {
				// Construct an instance of the RunReclamationActionOptions model
				id := "testString"
				actionName := "testString"
				runReclamationActionOptionsModel := testService.NewRunReclamationActionOptions(id, actionName)
				runReclamationActionOptionsModel.SetID("testString")
				runReclamationActionOptionsModel.SetActionName("testString")
				runReclamationActionOptionsModel.SetRequestBy("testString")
				runReclamationActionOptionsModel.SetComment("testString")
				runReclamationActionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(runReclamationActionOptionsModel).ToNot(BeNil())
				Expect(runReclamationActionOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.ActionName).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.RequestBy).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.Comment).To(Equal(core.StringPtr("testString")))
				Expect(runReclamationActionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUnlockResourceInstanceOptions successfully`, func() {
				// Construct an instance of the UnlockResourceInstanceOptions model
				id := "testString"
				unlockResourceInstanceOptionsModel := testService.NewUnlockResourceInstanceOptions(id)
				unlockResourceInstanceOptionsModel.SetID("testString")
				unlockResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(unlockResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(unlockResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(unlockResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceAliasOptions successfully`, func() {
				// Construct an instance of the UpdateResourceAliasOptions model
				id := "testString"
				updateResourceAliasOptionsName := "my-new-alias-name"
				updateResourceAliasOptionsModel := testService.NewUpdateResourceAliasOptions(id, updateResourceAliasOptionsName)
				updateResourceAliasOptionsModel.SetID("testString")
				updateResourceAliasOptionsModel.SetName("my-new-alias-name")
				updateResourceAliasOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceAliasOptionsModel).ToNot(BeNil())
				Expect(updateResourceAliasOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceAliasOptionsModel.Name).To(Equal(core.StringPtr("my-new-alias-name")))
				Expect(updateResourceAliasOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceBindingOptions successfully`, func() {
				// Construct an instance of the UpdateResourceBindingOptions model
				id := "testString"
				updateResourceBindingOptionsName := "my-new-binding-name"
				updateResourceBindingOptionsModel := testService.NewUpdateResourceBindingOptions(id, updateResourceBindingOptionsName)
				updateResourceBindingOptionsModel.SetID("testString")
				updateResourceBindingOptionsModel.SetName("my-new-binding-name")
				updateResourceBindingOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceBindingOptionsModel).ToNot(BeNil())
				Expect(updateResourceBindingOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceBindingOptionsModel.Name).To(Equal(core.StringPtr("my-new-binding-name")))
				Expect(updateResourceBindingOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceInstanceOptions successfully`, func() {
				// Construct an instance of the UpdateResourceInstanceOptions model
				id := "testString"
				updateResourceInstanceOptionsModel := testService.NewUpdateResourceInstanceOptions(id)
				updateResourceInstanceOptionsModel.SetID("testString")
				updateResourceInstanceOptionsModel.SetName("my-new-instance-name")
				updateResourceInstanceOptionsModel.SetParameters(make(map[string]interface{}))
				updateResourceInstanceOptionsModel.SetResourcePlanID("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.SetAllowCleanup(true)
				updateResourceInstanceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceInstanceOptionsModel).ToNot(BeNil())
				Expect(updateResourceInstanceOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceInstanceOptionsModel.Name).To(Equal(core.StringPtr("my-new-instance-name")))
				Expect(updateResourceInstanceOptionsModel.Parameters).To(Equal(make(map[string]interface{})))
				Expect(updateResourceInstanceOptionsModel.ResourcePlanID).To(Equal(core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")))
				Expect(updateResourceInstanceOptionsModel.AllowCleanup).To(Equal(core.BoolPtr(true)))
				Expect(updateResourceInstanceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateResourceKeyOptions successfully`, func() {
				// Construct an instance of the UpdateResourceKeyOptions model
				id := "testString"
				updateResourceKeyOptionsName := "my-new-key-name"
				updateResourceKeyOptionsModel := testService.NewUpdateResourceKeyOptions(id, updateResourceKeyOptionsName)
				updateResourceKeyOptionsModel.SetID("testString")
				updateResourceKeyOptionsModel.SetName("my-new-key-name")
				updateResourceKeyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateResourceKeyOptionsModel).ToNot(BeNil())
				Expect(updateResourceKeyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateResourceKeyOptionsModel.Name).To(Equal(core.StringPtr("my-new-key-name")))
				Expect(updateResourceKeyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
