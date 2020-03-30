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

package enterprisemanagementv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/enterprisemanagementv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`EnterpriseManagementV1`, func() {
	Describe(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions)`, func() {
		createAccountGroupPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createAccountGroupPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"account_group_id": "AccountGroupID"}`)
			}))
			It(`Invoke CreateAccountGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountGroupOptions model
				createAccountGroupOptionsModel := new(enterprisemanagementv1.CreateAccountGroupOptions)
				createAccountGroupOptionsModel.Parent = core.StringPtr("testString")
				createAccountGroupOptionsModel.Name = core.StringPtr("testString")
				createAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateAccountGroup(createAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {
		listAccountGroupsPath := "/account-groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAccountGroupsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["parent_account_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}]}`)
			}))
			It(`Invoke ListAccountGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccountGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountGroupsOptions model
				listAccountGroupsOptionsModel := new(enterprisemanagementv1.ListAccountGroupsOptions)
				listAccountGroupsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.ParentAccountGroupID = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Parent = core.StringPtr("testString")
				listAccountGroupsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccountGroups(listAccountGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountGroupByID(getAccountGroupByIdOptions *GetAccountGroupByIdOptions)`, func() {
		getAccountGroupByIDPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountGroupByIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}`)
			}))
			It(`Invoke GetAccountGroupByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountGroupByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountGroupByIdOptions model
				getAccountGroupByIdOptionsModel := new(enterprisemanagementv1.GetAccountGroupByIdOptions)
				getAccountGroupByIdOptionsModel.AccountGroupID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountGroupByID(getAccountGroupByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions)`, func() {
		updateAccountGroupPath := "/account-groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateAccountGroupPath))
				Expect(req.Method).To(Equal("PATCH"))
				res.WriteHeader(204)
			}))
			It(`Invoke UpdateAccountGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateAccountGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountGroupOptions model
				updateAccountGroupOptionsModel := new(enterprisemanagementv1.UpdateAccountGroupOptions)
				updateAccountGroupOptionsModel.AccountGroupID = core.StringPtr("testString")
				updateAccountGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccountGroupOptionsModel.PrimaryContactIamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateAccountGroup(updateAccountGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountGroupPermissibleActions(getAccountGroupPermissibleActionsOptions *GetAccountGroupPermissibleActionsOptions)`, func() {
		getAccountGroupPermissibleActionsPath := "/account-groups/testString/permissible-actions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountGroupPermissibleActionsPath))
				Expect(req.Method).To(Equal("POST"))
				res.WriteHeader(200)
			}))
			It(`Invoke GetAccountGroupPermissibleActions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetAccountGroupPermissibleActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetAccountGroupPermissibleActionsOptions model
				getAccountGroupPermissibleActionsOptionsModel := new(enterprisemanagementv1.GetAccountGroupPermissibleActionsOptions)
				getAccountGroupPermissibleActionsOptionsModel.AccountGroupID = core.StringPtr("testString")
				getAccountGroupPermissibleActionsOptionsModel.Actions = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetAccountGroupPermissibleActions(getAccountGroupPermissibleActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ImportAccountToEnterprise(importAccountToEnterpriseOptions *ImportAccountToEnterpriseOptions)`, func() {
		importAccountToEnterprisePath := "/enterprises/testString/import/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(importAccountToEnterprisePath))
				Expect(req.Method).To(Equal("PUT"))
				res.WriteHeader(202)
			}))
			It(`Invoke ImportAccountToEnterprise successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.ImportAccountToEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the ImportAccountToEnterpriseOptions model
				importAccountToEnterpriseOptionsModel := new(enterprisemanagementv1.ImportAccountToEnterpriseOptions)
				importAccountToEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.AccountID = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.Parent = core.StringPtr("testString")
				importAccountToEnterpriseOptionsModel.BillingUnitID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.ImportAccountToEnterprise(importAccountToEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateAccount(createAccountOptions *CreateAccountOptions)`, func() {
		createAccountPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createAccountPath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"account_group_id": "AccountGroupID"}`)
			}))
			It(`Invoke CreateAccount successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccountOptions model
				createAccountOptionsModel := new(enterprisemanagementv1.CreateAccountOptions)
				createAccountOptionsModel.Parent = core.StringPtr("testString")
				createAccountOptionsModel.Name = core.StringPtr("testString")
				createAccountOptionsModel.OwnerIamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateAccount(createAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {
		listAccountsPath := "/accounts"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAccountsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["enterprise_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["parent"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}]}`)
			}))
			It(`Invoke ListAccounts successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccounts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccountsOptions model
				listAccountsOptionsModel := new(enterprisemanagementv1.ListAccountsOptions)
				listAccountsOptionsModel.EnterpriseID = core.StringPtr("testString")
				listAccountsOptionsModel.AccountGroupID = core.StringPtr("testString")
				listAccountsOptionsModel.Parent = core.StringPtr("testString")
				listAccountsOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccounts(listAccountsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountByID(getAccountByIdOptions *GetAccountByIdOptions)`, func() {
		getAccountByIDPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountByIDPath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "URL", "id": "ID", "crn": "Crn", "parent": "Parent", "enterprise_account_id": "EnterpriseAccountID", "enterprise_id": "EnterpriseID", "enterprise_path": "EnterprisePath", "name": "Name", "state": "State", "owner_iam_id": "OwnerIamID", "paid": true, "owner_email": "OwnerEmail", "is_enterprise_account": false, "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}`)
			}))
			It(`Invoke GetAccountByID successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountByID(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountByIdOptions model
				getAccountByIdOptionsModel := new(enterprisemanagementv1.GetAccountByIdOptions)
				getAccountByIdOptionsModel.AccountID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountByID(getAccountByIdOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateAccount(updateAccountOptions *UpdateAccountOptions)`, func() {
		updateAccountPath := "/accounts/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateAccountPath))
				Expect(req.Method).To(Equal("PATCH"))
				res.WriteHeader(204)
			}))
			It(`Invoke UpdateAccount successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateAccountOptions model
				updateAccountOptionsModel := new(enterprisemanagementv1.UpdateAccountOptions)
				updateAccountOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountOptionsModel.Parent = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateAccount(updateAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountPermissibleActions(getAccountPermissibleActionsOptions *GetAccountPermissibleActionsOptions)`, func() {
		getAccountPermissibleActionsPath := "/accounts/testString/permissible-actions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountPermissibleActionsPath))
				Expect(req.Method).To(Equal("POST"))
				res.WriteHeader(200)
			}))
			It(`Invoke GetAccountPermissibleActions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetAccountPermissibleActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetAccountPermissibleActionsOptions model
				getAccountPermissibleActionsOptionsModel := new(enterprisemanagementv1.GetAccountPermissibleActionsOptions)
				getAccountPermissibleActionsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountPermissibleActionsOptionsModel.Actions = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetAccountPermissibleActions(getAccountPermissibleActionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateEnterprise(createEnterpriseOptions *CreateEnterpriseOptions)`, func() {
		createEnterprisePath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createEnterprisePath))
				Expect(req.Method).To(Equal("POST"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(202)
				fmt.Fprintf(res, `{"enterprise_id": "EnterpriseID", "enterprise_account_id": "EnterpriseAccountID"}`)
			}))
			It(`Invoke CreateEnterprise successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateEnterpriseOptions model
				createEnterpriseOptionsModel := new(enterprisemanagementv1.CreateEnterpriseOptions)
				createEnterpriseOptionsModel.SourceAccountID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Name = core.StringPtr("testString")
				createEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")
				createEnterpriseOptionsModel.Domain = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateEnterprise(createEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions)`, func() {
		listEnterprisesPath := "/enterprises"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listEnterprisesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.URL.Query()["enterprise_account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_group_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rows_count": 9, "next_url": "NextURL", "resources": [{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "Crn", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}]}`)
			}))
			It(`Invoke ListEnterprises successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListEnterprises(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListEnterprisesOptions model
				listEnterprisesOptionsModel := new(enterprisemanagementv1.ListEnterprisesOptions)
				listEnterprisesOptionsModel.EnterpriseAccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountGroupID = core.StringPtr("testString")
				listEnterprisesOptionsModel.AccountID = core.StringPtr("testString")
				listEnterprisesOptionsModel.Limit = core.Int64Ptr(int64(38))

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListEnterprises(listEnterprisesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {
		getEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEnterprisePath))
				Expect(req.Method).To(Equal("GET"))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"url": "URL", "id": "ID", "enterprise_account_id": "EnterpriseAccountID", "crn": "Crn", "name": "Name", "domain": "Domain", "state": "State", "primary_contact_iam_id": "PrimaryContactIamID", "primary_contact_email": "PrimaryContactEmail", "created_at": "CreatedAt", "created_by": "CreatedBy", "updated_at": "UpdatedAt", "updated_by": "UpdatedBy"}`)
			}))
			It(`Invoke GetEnterprise successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetEnterpriseOptions model
				getEnterpriseOptionsModel := new(enterprisemanagementv1.GetEnterpriseOptions)
				getEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetEnterprise(getEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {
		updateEnterprisePath := "/enterprises/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateEnterprisePath))
				Expect(req.Method).To(Equal("PATCH"))
				res.WriteHeader(204)
			}))
			It(`Invoke UpdateEnterprise successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateEnterprise(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateEnterpriseOptions model
				updateEnterpriseOptionsModel := new(enterprisemanagementv1.UpdateEnterpriseOptions)
				updateEnterpriseOptionsModel.EnterpriseID = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Name = core.StringPtr("testString")
				updateEnterpriseOptionsModel.Domain = core.StringPtr("testString")
				updateEnterpriseOptionsModel.PrimaryContactIamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateEnterprise(updateEnterpriseOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetEnterprisePermissibleActions(getEnterprisePermissibleActionsOptions *GetEnterprisePermissibleActionsOptions)`, func() {
		getEnterprisePermissibleActionsPath := "/enterprises/testString/permissible-actions"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getEnterprisePermissibleActionsPath))
				Expect(req.Method).To(Equal("POST"))
				res.WriteHeader(200)
			}))
			It(`Invoke GetEnterprisePermissibleActions successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := enterprisemanagementv1.NewEnterpriseManagementV1(&enterprisemanagementv1.EnterpriseManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.GetEnterprisePermissibleActions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the GetEnterprisePermissibleActionsOptions model
				getEnterprisePermissibleActionsOptionsModel := new(enterprisemanagementv1.GetEnterprisePermissibleActionsOptions)
				getEnterprisePermissibleActionsOptionsModel.EnterpriseID = core.StringPtr("testString")
				getEnterprisePermissibleActionsOptionsModel.Actions = []string{"testString"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.GetEnterprisePermissibleActions(getEnterprisePermissibleActionsOptionsModel)
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
