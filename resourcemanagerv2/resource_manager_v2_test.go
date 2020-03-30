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
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/resourcemanagerv2"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`ResourceManagerV2`, func() {
	Describe(`GetAccountQuotaList(getAccountQuotaListOptions *GetAccountQuotaListOptions)`, func() {
		bearerToken := "0ui9876453"
		getAccountQuotaListPath := "/quota_definitions/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountQuotaListPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "type": "Type", "number_of_apps": 12, "number_of_service_instances": 24, "default_number_of_instances_per_lite_plan": 35, "instances_per_app": 15, "instance_memory": "InstanceMemory", "total_app_memory": "TotalAppMemory", "vsi_limit": 8, "resource_quotas": {"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}, "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
			}))
			It(`Invoke GetAccountQuotaList successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountQuotaList(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountQuotaListOptions model
				getAccountQuotaListOptionsModel := new(resourcemanagerv2.GetAccountQuotaListOptions)
				getAccountQuotaListOptionsModel.AccountID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountQuotaList(getAccountQuotaListOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceQuota(getResourceQuotaOptions *GetResourceQuotaOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceQuotaPath := "/quota_definitions/accounts/testString/resource_types/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceQuotaPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}`)
			}))
			It(`Invoke GetResourceQuota successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceQuotaOptions model
				getResourceQuotaOptionsModel := new(resourcemanagerv2.GetResourceQuotaOptions)
				getResourceQuotaOptionsModel.AccountID = core.StringPtr("testString")
				getResourceQuotaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceQuota(getResourceQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceQuota(updateResourceQuotaOptions *UpdateResourceQuotaOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceQuotaPath := "/quota_definitions/accounts/testString/resource_types/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceQuotaPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"error_code": "RG-CloudResourceGroupErrorResponse", "message": "Message", "status_code": "StatusCode", "transaction_id": "TransactionID"}`)
			}))
			It(`Invoke UpdateResourceQuota successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceQuotaOptions model
				updateResourceQuotaOptionsModel := new(resourcemanagerv2.UpdateResourceQuotaOptions)
				updateResourceQuotaOptionsModel.AccountID = core.StringPtr("testString")
				updateResourceQuotaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceQuota(updateResourceQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceQuota(deleteResourceQuotaOptions *DeleteResourceQuotaOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceQuotaPath := "/quota_definitions/accounts/testString/resource_types/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceQuotaPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"error_code": "RG-CloudResourceGroupErrorResponse", "message": "Message", "status_code": "StatusCode", "transaction_id": "TransactionID"}`)
			}))
			It(`Invoke DeleteResourceQuota successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.DeleteResourceQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteResourceQuotaOptions model
				deleteResourceQuotaOptionsModel := new(resourcemanagerv2.DeleteResourceQuotaOptions)
				deleteResourceQuotaOptionsModel.AccountID = core.StringPtr("testString")
				deleteResourceQuotaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.DeleteResourceQuota(deleteResourceQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateDefaultResourceQuota(createDefaultResourceQuotaOptions *CreateDefaultResourceQuotaOptions)`, func() {
		bearerToken := "0ui9876453"
		createDefaultResourceQuotaPath := "/quota_definitions/resource_types/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createDefaultResourceQuotaPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"error_code": "RG-CloudResourceGroupErrorResponse", "message": "Message", "status_code": "StatusCode", "transaction_id": "TransactionID"}`)
			}))
			It(`Invoke CreateDefaultResourceQuota successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateDefaultResourceQuota(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateDefaultResourceQuotaOptions model
				createDefaultResourceQuotaOptionsModel := new(resourcemanagerv2.CreateDefaultResourceQuotaOptions)
				createDefaultResourceQuotaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateDefaultResourceQuota(createDefaultResourceQuotaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateSchema(createSchemaOptions *CreateSchemaOptions)`, func() {
		bearerToken := "0ui9876453"
		createSchemaPath := "/quota_definitions/resource_types/testString/schemas"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createSchemaPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"error_code": "RG-CloudResourceGroupErrorResponse", "message": "Message", "status_code": "StatusCode", "transaction_id": "TransactionID"}`)
			}))
			It(`Invoke CreateSchema successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSchemaOptions model
				createSchemaOptionsModel := new(resourcemanagerv2.CreateSchemaOptions)
				createSchemaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateSchema(createSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetSchema(getSchemaOptions *GetSchemaOptions)`, func() {
		bearerToken := "0ui9876453"
		getSchemaPath := "/quota_definitions/resource_types/testString/schemas"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getSchemaPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}`)
			}))
			It(`Invoke GetSchema successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetSchema(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSchemaOptions model
				getSchemaOptionsModel := new(resourcemanagerv2.GetSchemaOptions)
				getSchemaOptionsModel.ResourceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetSchema(getSchemaOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListQuotaDefinitions(listQuotaDefinitionsOptions *ListQuotaDefinitionsOptions)`, func() {
		bearerToken := "0ui9876453"
		listQuotaDefinitionsPath := "/quota_definitions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listQuotaDefinitionsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"resources": [{"id": "ID", "name": "Name", "type": "Type", "number_of_apps": 12, "number_of_service_instances": 24, "default_number_of_instances_per_lite_plan": 35, "instances_per_app": 15, "instance_memory": "InstanceMemory", "total_app_memory": "TotalAppMemory", "vsi_limit": 8, "resource_quotas": {"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}, "created_at": "CreatedAt", "updated_at": "UpdatedAt"}]}`)
			}))
			It(`Invoke ListQuotaDefinitions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListQuotaDefinitions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListQuotaDefinitionsOptions model
				listQuotaDefinitionsOptionsModel := new(resourcemanagerv2.ListQuotaDefinitionsOptions)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListQuotaDefinitions(listQuotaDefinitionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetQuotaDefinition(getQuotaDefinitionOptions *GetQuotaDefinitionOptions)`, func() {
		bearerToken := "0ui9876453"
		getQuotaDefinitionPath := "/quota_definitions/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getQuotaDefinitionPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "type": "Type", "number_of_apps": 12, "number_of_service_instances": 24, "default_number_of_instances_per_lite_plan": 35, "instances_per_app": 15, "instance_memory": "InstanceMemory", "total_app_memory": "TotalAppMemory", "vsi_limit": 8, "resource_quotas": {"_id": "ID", "resource_id": "ResourceID", "crn": "Crn", "limit": 5}, "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
			}))
			It(`Invoke GetQuotaDefinition successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetQuotaDefinition(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetQuotaDefinitionOptions model
				getQuotaDefinitionOptionsModel := new(resourcemanagerv2.GetQuotaDefinitionOptions)
				getQuotaDefinitionOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetQuotaDefinition(getQuotaDefinitionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListResourceGroups(listResourceGroupsOptions *ListResourceGroupsOptions)`, func() {
		bearerToken := "0ui9876453"
		listResourceGroupsPath := "/resource_groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listResourceGroupsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["date"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"resources": [{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}]}`)
			}))
			It(`Invoke ListResourceGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListResourceGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListResourceGroupsOptions model
				listResourceGroupsOptionsModel := new(resourcemanagerv2.ListResourceGroupsOptions)
				listResourceGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listResourceGroupsOptionsModel.Date = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListResourceGroups(listResourceGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateResourceGroup(createResourceGroupOptions *CreateResourceGroupOptions)`, func() {
		bearerToken := "0ui9876453"
		createResourceGroupPath := "/resource_groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createResourceGroupPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "crn": "Crn"}`)
			}))
			It(`Invoke CreateResourceGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateResourceGroupOptions model
				createResourceGroupOptionsModel := new(resourcemanagerv2.CreateResourceGroupOptions)
				createResourceGroupOptionsModel.Name = core.StringPtr("test1")
				createResourceGroupOptionsModel.AccountID = core.StringPtr("25eba2a9-beef-450b-82cf-f5ad5e36c6dd")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateResourceGroup(createResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetResourceGroup(getResourceGroupOptions *GetResourceGroupOptions)`, func() {
		bearerToken := "0ui9876453"
		getResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getResourceGroupPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
			}))
			It(`Invoke GetResourceGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetResourceGroupOptions model
				getResourceGroupOptionsModel := new(resourcemanagerv2.GetResourceGroupOptions)
				getResourceGroupOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetResourceGroup(getResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateResourceGroup(updateResourceGroupOptions *UpdateResourceGroupOptions)`, func() {
		bearerToken := "0ui9876453"
		updateResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateResourceGroupPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "crn": "Crn", "account_id": "AccountID", "name": "Name", "state": "State", "default": false, "quota_id": "QuotaID", "quota_url": "QuotaURL", "payment_methods_url": "PaymentMethodsURL", "resource_linkages": [{"anyKey": "anyValue"}], "teams_url": "TeamsURL", "created_at": "CreatedAt", "updated_at": "UpdatedAt"}`)
			}))
			It(`Invoke UpdateResourceGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateResourceGroupOptions model
				updateResourceGroupOptionsModel := new(resourcemanagerv2.UpdateResourceGroupOptions)
				updateResourceGroupOptionsModel.ID = core.StringPtr("testString")
				updateResourceGroupOptionsModel.Name = core.StringPtr("testString")
				updateResourceGroupOptionsModel.State = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateResourceGroup(updateResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteResourceGroup(deleteResourceGroupOptions *DeleteResourceGroupOptions)`, func() {
		bearerToken := "0ui9876453"
		deleteResourceGroupPath := "/resource_groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteResourceGroupPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteResourceGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := resourcemanagerv2.NewResourceManagerV2(&resourcemanagerv2.ResourceManagerV2Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteResourceGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteResourceGroupOptions model
				deleteResourceGroupOptionsModel := new(resourcemanagerv2.DeleteResourceGroupOptions)
				deleteResourceGroupOptionsModel.ID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteResourceGroup(deleteResourceGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
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
