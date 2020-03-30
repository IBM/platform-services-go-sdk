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

package iamaccessgroupsv2_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`IamAccessGroupsV2`, func() {
	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions)`, func() {
		createAccessGroupPath := "/groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createAccessGroupPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
			}))
			It(`Invoke CreateAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.CreateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("testString")
				createAccessGroupOptionsModel.Description = core.StringPtr("testString")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions)`, func() {
		listAccessGroupsPath := "/groups"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAccessGroupsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))

				Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))


				// TODO: Add check for show_federated query parameter


				// TODO: Add check for hide_public_access query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}]}`)
			}))
			It(`Invoke ListAccessGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(true)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions)`, func() {
		getAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccessGroupPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

				// TODO: Add check for show_federated query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
			}))
			It(`Invoke GetAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions)`, func() {
		updateAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateAccessGroupPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
			}))
			It(`Invoke UpdateAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Description = core.StringPtr("testString")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions)`, func() {
		deleteAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(deleteAccessGroupPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

				// TODO: Add check for force query parameter

				res.WriteHeader(204)
			}))
			It(`Invoke DeleteAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.DeleteAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {
		getAccountSettingsPath := "/groups/settings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccountSettingsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
			}))
			It(`Invoke GetAccountSettings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {
		updateAccountSettingsPath := "/groups/settings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateAccountSettingsPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"account_id": "AccountID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
			}))
			It(`Invoke UpdateAccountSettings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.UpdateAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions)`, func() {
		isMemberOfAccessGroupPath := "/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(isMemberOfAccessGroupPath))
				Expect(req.Method).To(Equal("HEAD"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(204)
			}))
			It(`Invoke IsMemberOfAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.IsMemberOfAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the IsMemberOfAccessGroupOptions model
				isMemberOfAccessGroupOptionsModel := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				isMemberOfAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions)`, func() {
		addMembersToAccessGroupPath := "/groups/testString/members"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addMembersToAccessGroupPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(207)
				fmt.Fprintf(res, `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "CreatedAt", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
			}))
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddMembersToAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("testString")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("testString")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions)`, func() {
		listAccessGroupMembersPath := "/groups/testString/members"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAccessGroupMembersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

				// TODO: Add check for limit query parameter


				// TODO: Add check for offset query parameter

				Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))


				// TODO: Add check for verbose query parameter

				Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "CreatedAt", "created_by_id": "CreatedByID"}]}`)
			}))
			It(`Invoke ListAccessGroupMembers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccessGroupMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Limit = core.Float64Ptr(72.5)
				listAccessGroupMembersOptionsModel.Offset = core.Float64Ptr(72.5)
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(true)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions)`, func() {
		removeMemberFromAccessGroupPath := "/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeMemberFromAccessGroupPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(204)
			}))
			It(`Invoke RemoveMemberFromAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.RemoveMemberFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				removeMemberFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				removeMemberFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions)`, func() {
		removeMembersFromAccessGroupPath := "/groups/testString/members/delete"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeMembersFromAccessGroupPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(207)
				fmt.Fprintf(res, `{"access_group_id": "AccessGroupID", "members": [{"iam_id": "IamID", "trace": "Trace", "status_code": 10, "errors": [{"code": "Code", "message": "Message"}]}]}`)
			}))
			It(`Invoke RemoveMembersFromAccessGroup successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveMembersFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"testString"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions)`, func() {
		removeMemberFromAllAccessGroupsPath := "/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeMemberFromAllAccessGroupsPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(207)
				fmt.Fprintf(res, `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
			}))
			It(`Invoke RemoveMemberFromAllAccessGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.RemoveMemberFromAllAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions)`, func() {
		addMemberToMultipleAccessGroupsPath := "/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addMemberToMultipleAccessGroupsPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(207)
				fmt.Fprintf(res, `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
			}))
			It(`Invoke AddMemberToMultipleAccessGroups successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddMemberToMultipleAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"testString"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions)`, func() {
		addAccessGroupRulePath := "/groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(addAccessGroupRulePath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(201)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke AddAccessGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.AddAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("testString")
				ruleConditionsModel.Operator = core.StringPtr("testString")
				ruleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions)`, func() {
		listAccessGroupRulesPath := "/groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listAccessGroupRulesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}]}`)
			}))
			It(`Invoke ListAccessGroupRules successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListAccessGroupRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions)`, func() {
		getAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getAccessGroupRulePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke GetAccessGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions)`, func() {
		replaceAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(replaceAccessGroupRulePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["If-Match"]).ToNot(BeNil())
				Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "CreatedAt", "created_by_id": "CreatedByID", "last_modified_at": "LastModifiedAt", "last_modified_by_id": "LastModifiedByID"}`)
			}))
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ReplaceAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("testString")
				ruleConditionsModel.Operator = core.StringPtr("testString")
				ruleConditionsModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(38))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions)`, func() {
		removeAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeAccessGroupRulePath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
				Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
				res.WriteHeader(204)
			}))
			It(`Invoke RemoveAccessGroupRule successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.RemoveAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveAccessGroupRuleOptions model
				removeAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				removeAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a sample service client instance`, func() {
			testService, _ := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL:           "http://iamaccessgroupsv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddGroupMembersRequestMembersItem successfully`, func() {
				iamID := "testString"
				typeVar := "testString"
				model, err := testService.NewAddGroupMembersRequestMembersItem(iamID, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleConditions successfully`, func() {
				claim := "testString"
				operator := "testString"
				value := "testString"
				model, err := testService.NewRuleConditions(claim, operator, value)
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
