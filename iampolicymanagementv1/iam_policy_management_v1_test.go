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

package iampolicymanagementv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/iampolicymanagementv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`IamPolicyManagementV1`, func() {
	Describe(`ListPolicies(listPoliciesOptions *ListPoliciesOptions)`, func() {
		listPoliciesPath := "/v1/policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listPoliciesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Accept-Language"]).ToNot(BeNil())
				Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["service_type"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"policies": [{"id": "ID", "type": "Type", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}]}`)
			}))
			It(`Invoke ListPolicies successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("testString")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreatePolicy(createPolicyOptions *CreatePolicyOptions)`, func() {
		createPolicyPath := "/v1/policies"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createPolicyPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Accept-Language"]).ToNot(BeNil())
				Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "type": "Type", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke CreatePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreatePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRequestResourcesItemAttributesItem model
				policyRequestResourcesItemAttributesItemModel := new(iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem)
				policyRequestResourcesItemAttributesItemModel.Name = core.StringPtr("testString")
				policyRequestResourcesItemAttributesItemModel.Value = core.StringPtr("testString")
				policyRequestResourcesItemAttributesItemModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestSubjectsItemAttributesItem model
				policyRequestSubjectsItemAttributesItemModel := new(iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem)
				policyRequestSubjectsItemAttributesItemModel.Name = core.StringPtr("testString")
				policyRequestSubjectsItemAttributesItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestResourcesItem model
				policyRequestResourcesItemModel := new(iampolicymanagementv1.PolicyRequestResourcesItem)
				policyRequestResourcesItemModel.Attributes = []iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{*policyRequestResourcesItemAttributesItemModel}

				// Construct an instance of the PolicyRequestRolesItem model
				policyRequestRolesItemModel := new(iampolicymanagementv1.PolicyRequestRolesItem)
				policyRequestRolesItemModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestSubjectsItem model
				policyRequestSubjectsItemModel := new(iampolicymanagementv1.PolicyRequestSubjectsItem)
				policyRequestSubjectsItemModel.Attributes = []iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{*policyRequestSubjectsItemAttributesItemModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicyRequestSubjectsItem{*policyRequestSubjectsItemModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRequestRolesItem{*policyRequestRolesItemModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyRequestResourcesItem{*policyRequestResourcesItemModel}
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions)`, func() {
		updatePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updatePolicyPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "type": "Type", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke UpdatePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdatePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRequestResourcesItemAttributesItem model
				policyRequestResourcesItemAttributesItemModel := new(iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem)
				policyRequestResourcesItemAttributesItemModel.Name = core.StringPtr("testString")
				policyRequestResourcesItemAttributesItemModel.Value = core.StringPtr("testString")
				policyRequestResourcesItemAttributesItemModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestSubjectsItemAttributesItem model
				policyRequestSubjectsItemAttributesItemModel := new(iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem)
				policyRequestSubjectsItemAttributesItemModel.Name = core.StringPtr("testString")
				policyRequestSubjectsItemAttributesItemModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestResourcesItem model
				policyRequestResourcesItemModel := new(iampolicymanagementv1.PolicyRequestResourcesItem)
				policyRequestResourcesItemModel.Attributes = []iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{*policyRequestResourcesItemAttributesItemModel}

				// Construct an instance of the PolicyRequestRolesItem model
				policyRequestRolesItemModel := new(iampolicymanagementv1.PolicyRequestRolesItem)
				policyRequestRolesItemModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicyRequestSubjectsItem model
				policyRequestSubjectsItemModel := new(iampolicymanagementv1.PolicyRequestSubjectsItem)
				policyRequestSubjectsItemModel.Attributes = []iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{*policyRequestSubjectsItemAttributesItemModel}

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicyRequestSubjectsItem{*policyRequestSubjectsItemModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRequestRolesItem{*policyRequestRolesItemModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyRequestResourcesItem{*policyRequestResourcesItemModel}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetPolicy(getPolicyOptions *GetPolicyOptions)`, func() {
		getPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getPolicyPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "type": "Type", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke GetPolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeletePolicy(deletePolicyOptions *DeletePolicyOptions)`, func() {
		deletePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deletePolicyPath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(204)
			}))
			It(`Invoke DeletePolicy successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeletePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePolicyOptions model
				deletePolicyOptionsModel := new(iampolicymanagementv1.DeletePolicyOptions)
				deletePolicyOptionsModel.PolicyID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeletePolicy(deletePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListRoles(listRolesOptions *ListRolesOptions)`, func() {
		listRolesPath := "/v2/roles"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listRolesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Accept-Language"]).ToNot(BeNil())
				Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"custom_roles": [{"id": "ID", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}], "service_roles": [{"crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"]}], "system_roles": [{"crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"]}]}`)
			}))
			It(`Invoke ListRoles successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateRole(createRoleOptions *CreateRoleOptions)`, func() {
		createRolePath := "/v2/roles"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createRolePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Accept-Language"]).ToNot(BeNil())
				Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
			}))
			It(`Invoke CreateRole successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.Name = core.StringPtr("testString")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("testString")
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions)`, func() {
		updateRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateRolePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
			}))
			It(`Invoke UpdateRole successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetRole(getRoleOptions *GetRoleOptions)`, func() {
		getRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getRolePath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "crn": "Crn", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
			}))
			It(`Invoke GetRole successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteRole(deleteRoleOptions *DeleteRoleOptions)`, func() {
		deleteRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteRolePath))
				Expect(req.Method).To(Equal("DELETE"))
				res.WriteHeader(204)
			}))
			It(`Invoke DeleteRole successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(iampolicymanagementv1.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL:           "http://iampolicymanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewPolicyRequestResourcesItem successfully`, func() {
				attributes := []iampolicymanagementv1.PolicyRequestResourcesItemAttributesItem{}
				model, err := testService.NewPolicyRequestResourcesItem(attributes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPolicyRequestResourcesItemAttributesItem successfully`, func() {
				name := "testString"
				value := "testString"
				model, err := testService.NewPolicyRequestResourcesItemAttributesItem(name, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPolicyRequestRolesItem successfully`, func() {
				roleID := "testString"
				model, err := testService.NewPolicyRequestRolesItem(roleID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPolicyRequestSubjectsItem successfully`, func() {
				attributes := []iampolicymanagementv1.PolicyRequestSubjectsItemAttributesItem{}
				model, err := testService.NewPolicyRequestSubjectsItem(attributes)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPolicyRequestSubjectsItemAttributesItem successfully`, func() {
				name := "testString"
				value := "testString"
				model, err := testService.NewPolicyRequestSubjectsItemAttributesItem(name, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
