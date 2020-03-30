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
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/resourcecontrollerv2"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`ResourceControllerV2`, func() {
	Describe(`ListResourceInstances(listResourceInstancesOptions *ListResourceInstancesOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceInstancesPath := "/resource_instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceInstancesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_plan_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["sub_type"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {"mapKey": {"anyKey": "anyValue"}}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
			}))
			It(`Invoke ListResourceInstances successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				listResourceInstancesOptionsModel.UpdatedFrom = core.StringPtr("testString")
				listResourceInstancesOptionsModel.UpdatedTo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceInstances(listResourceInstancesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceInstance(createResourceInstanceOptions *CreateResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceInstancePath := "/resource_instances"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceInstancePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {"mapKey": {"anyKey": "anyValue"}}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke CreateResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				createResourceInstanceOptionsModel.Parameters = CreateMockMap()

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceInstance(createResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceInstance(getResourceInstanceOptions *GetResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceInstancePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {"mapKey": {"anyKey": "anyValue"}}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke GetResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceInstance(getResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceInstance(deleteResourceInstanceOptions *DeleteResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceInstancePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {"mapKey": {"anyKey": "anyValue"}}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke DeleteResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteResourceInstance(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteResourceInstanceOptions model
				deleteResourceInstanceOptionsModel := new(resourcecontrollerv2.DeleteResourceInstanceOptions)
				deleteResourceInstanceOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteResourceInstance(deleteResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceInstance(updateResourceInstanceOptions *UpdateResourceInstanceOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceInstancePath := "/resource_instances/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceInstancePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "resource_id": "ResourceID", "resource_plan_id": "ResourcePlanID", "target_crn": "TargetCrn", "state": "State", "type": "Type", "sub_type": "SubType", "allow_cleanup": true, "last_operation": {"mapKey": {"anyKey": "anyValue"}}, "dashboard_url": "DashboardURL", "plan_history": [{"resource_plan_id": "ResourcePlanID", "start_date": "2019-01-01T12:00:00"}], "resource_aliases_url": "ResourceAliasesURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke UpdateResourceInstance successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				updateResourceInstanceOptionsModel.Parameters = CreateMockMap()
				updateResourceInstanceOptionsModel.ResourcePlanID = core.StringPtr("a8dff6d3-d287-4668-a81d-c87c55c2656d")
				updateResourceInstanceOptionsModel.AllowCleanup = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceInstance(updateResourceInstanceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListResourceKeys(listResourceKeysOptions *ListResourceKeysOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceKeysPath := "/resource_keys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceKeysPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
			}))
			It(`Invoke ListResourceKeys successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				listResourceKeysOptionsModel.UpdatedFrom = core.StringPtr("testString")
				listResourceKeysOptionsModel.UpdatedTo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceKeys(listResourceKeysOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceKey(createResourceKeyOptions *CreateResourceKeyOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceKeyPath := "/resource_keys"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceKeyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke CreateResourceKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceKey(createResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceKey(getResourceKeyOptions *GetResourceKeyOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceKeyPath := "/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceKeyPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke GetResourceKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceKey(getResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceKey(deleteResourceKeyOptions *DeleteResourceKeyOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceKeyPath := "/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceKeyPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceKey(deleteResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceKey(updateResourceKeyOptions *UpdateResourceKeyOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceKeyPath := "/resource_keys/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceKeyPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_instance_url": "ResourceInstanceURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke UpdateResourceKey successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceKey(updateResourceKeyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListResourceBindings(listResourceBindingsOptions *ListResourceBindingsOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceBindingsPath := "/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceBindingsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["region_binding_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
			}))
			It(`Invoke ListResourceBindings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				listResourceBindingsOptionsModel.UpdatedFrom = core.StringPtr("testString")
				listResourceBindingsOptionsModel.UpdatedTo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceBindings(listResourceBindingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceBinding(createResourceBindingOptions *CreateResourceBindingOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceBindingPath := "/resource_bindings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceBindingPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke CreateResourceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				createResourceBindingOptionsModel.Target = core.StringPtr("crn:v1:bluemix:public:bluemix:us-south:s/0ba4dba0-a120-4a1e-a124-5a249a904b76::cf-application:a1caa40b-2c24-4da8-8267-ac2c1a42ad0c")
				createResourceBindingOptionsModel.Name = core.StringPtr("my-binding")
				createResourceBindingOptionsModel.Parameters = resourceBindingPostParametersModel
				createResourceBindingOptionsModel.Role = core.StringPtr("Writer")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceBinding(createResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceBinding(getResourceBindingOptions *GetResourceBindingOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceBindingPath := "/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceBindingPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke GetResourceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceBinding(getResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceBinding(deleteResourceBindingOptions *DeleteResourceBindingOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceBindingPath := "/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceBindingPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceBinding(deleteResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceBinding(updateResourceBindingOptions *UpdateResourceBindingOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceBindingPath := "/resource_bindings/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceBindingPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "source_crn": "SourceCrn", "target_crn": "TargetCrn", "region_binding_id": "RegionBindingID", "state": "State", "credentials": {"apikey": "Apikey", "iam_apikey_description": "IamApikeyDescription", "iam_apikey_name": "IamApikeyName", "iam_role_crn": "IamRoleCrn", "iam_serviceid_crn": "IamServiceidCrn"}, "iam_compatible": false, "resource_alias_url": "ResourceAliasURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke UpdateResourceBinding successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceBinding(updateResourceBindingOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListResourceAliases(listResourceAliasesOptions *ListResourceAliasesOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceAliasesPath := "/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceAliasesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["guid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["name"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["region_instance_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_from"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["updated_to"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"next_url": "NextURL", "resources": [{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}], "rows_count": 9}`)
			}))
			It(`Invoke ListResourceAliases successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				listResourceAliasesOptionsModel.UpdatedFrom = core.StringPtr("testString")
				listResourceAliasesOptionsModel.UpdatedTo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceAliases(listResourceAliasesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceAlias(createResourceAliasOptions *CreateResourceAliasOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceAliasPath := "/resource_aliases"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceAliasPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke CreateResourceAlias successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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
				createResourceAliasOptionsModel.Target = core.StringPtr("crn:v1:staging:public:bluemix:us-south:o/5e939cd5-6377-4383-b9e0-9db22cd11753::cf-space:66c8b915-101a-406c-a784-e6636676e4f5")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceAlias(createResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceAlias(getResourceAliasOptions *GetResourceAliasOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceAliasPath := "/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceAliasPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke GetResourceAlias successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceAlias(getResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceAlias(deleteResourceAliasOptions *DeleteResourceAliasOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceAliasPath := "/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceAliasPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceAlias successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceAlias(deleteResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceAlias(updateResourceAliasOptions *UpdateResourceAliasOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceAliasPath := "/resource_aliases/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceAliasPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "guid": "Guid", "crn": "Crn", "url": "URL", "name": "Name", "account_id": "AccountID", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "target_crn": "TargetCrn", "state": "State", "resource_instance_id": "ResourceInstanceID", "region_instance_id": "RegionInstanceID", "resource_instance_url": "ResourceInstanceURL", "resource_bindings_url": "ResourceBindingsURL", "resource_keys_url": "ResourceKeysURL", "created_at": "2019-01-01T12:00:00", "updated_at": "2019-01-01T12:00:00", "deleted_at": "2019-01-01T12:00:00"}`)
			}))
			It(`Invoke UpdateResourceAlias successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceAlias(updateResourceAliasOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListReclamations(listReclamationsOptions *ListReclamationsOptions)`, func() {
		bearerToken := "0ui9876453"
		listReclamationsPath := "/v1/reclamations"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listReclamationsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["resource_instance_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"resources": [{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCrn", "resource_instance_id": {"anyKey": "anyValue"}, "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": "CustomProperties", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}]}`)
			}))
			It(`Invoke ListReclamations successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListReclamations(listReclamationsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RunReclamationAction(runReclamationActionOptions *RunReclamationActionOptions)`, func() {
		bearerToken := "0ui9876453"
		runReclamationActionPath := "/v1/reclamations/testString/actions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(runReclamationActionPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "entity_id": "EntityID", "entity_type_id": "EntityTypeID", "entity_crn": "EntityCrn", "resource_instance_id": {"anyKey": "anyValue"}, "resource_group_id": "ResourceGroupID", "account_id": "AccountID", "policy_id": "PolicyID", "state": "State", "target_time": "TargetTime", "custom_properties": "CustomProperties", "created_at": "2019-01-01T12:00:00", "created_by": "CreatedBy", "updated_at": "2019-01-01T12:00:00", "updated_by": "UpdatedBy"}`)
			}))
			It(`Invoke RunReclamationAction successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcecontrollerv2.NewResourceControllerV2(&resourcecontrollerv2.ResourceControllerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
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

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RunReclamationAction(runReclamationActionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockMap() successfully`, func() {
			mockMap := CreateMockMap()
			Expect(mockMap).ToNot(BeNil())
		})
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

func CreateMockMap() map[string]interface{} {
	m := make(map[string]interface{})
	return m
}

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, len(mockData))
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
	d := strfmt.Date(time.Now())
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Now())
	return &d
}
