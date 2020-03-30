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

package usermanagementv1_test

import (
	"bytes"
	"fmt"
	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/platform-services-go-sdk/usermanagementv1"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"
)

var _ = Describe(`UserManagementV1`, func() {
	Describe(`GetUserLinkages(getUserLinkagesOptions *GetUserLinkagesOptions)`, func() {
		bearerToken := "0ui9876453"
		getUserLinkagesPath := "/v2/accounts/testString/users/testString/linkages"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getUserLinkagesPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"linkages": [{"origin": "Origin", "id": "ID"}]}`)
			}))
			It(`Invoke GetUserLinkages successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetUserLinkages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserLinkagesOptions model
				getUserLinkagesOptionsModel := new(usermanagementv1.GetUserLinkagesOptions)
				getUserLinkagesOptionsModel.AccountID = core.StringPtr("testString")
				getUserLinkagesOptionsModel.IamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetUserLinkages(getUserLinkagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateUserLinkages(createUserLinkagesOptions *CreateUserLinkagesOptions)`, func() {
		bearerToken := "0ui9876453"
		createUserLinkagesPath := "/v2/accounts/testString/users/testString/linkages/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createUserLinkagesPath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke CreateUserLinkages successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateUserLinkages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateUserLinkagesOptions model
				createUserLinkagesOptionsModel := new(usermanagementv1.CreateUserLinkagesOptions)
				createUserLinkagesOptionsModel.AccountID = core.StringPtr("testString")
				createUserLinkagesOptionsModel.IamID = core.StringPtr("testString")
				createUserLinkagesOptionsModel.Origin = core.StringPtr("testString")
				createUserLinkagesOptionsModel.IdFromOrigin = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateUserLinkages(createUserLinkagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveUserLinkages(removeUserLinkagesOptions *RemoveUserLinkagesOptions)`, func() {
		bearerToken := "0ui9876453"
		removeUserLinkagesPath := "/v2/accounts/testString/users/testString/linkages/testString/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeUserLinkagesPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke RemoveUserLinkages successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.RemoveUserLinkages(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveUserLinkagesOptions model
				removeUserLinkagesOptionsModel := new(usermanagementv1.RemoveUserLinkagesOptions)
				removeUserLinkagesOptionsModel.AccountID = core.StringPtr("testString")
				removeUserLinkagesOptionsModel.IamID = core.StringPtr("testString")
				removeUserLinkagesOptionsModel.Origin = core.StringPtr("testString")
				removeUserLinkagesOptionsModel.IdFromOrigin = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.RemoveUserLinkages(removeUserLinkagesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {
		bearerToken := "0ui9876453"
		getUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getUserProfilePath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))

				// TODO: Add check for include_linkages query parameter

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "linkages": [{"origin": "Origin", "id": "ID"}]}`)
			}))
			It(`Invoke GetUserProfile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.IncludeLinkages = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`CreateUserProfile(createUserProfileOptions *CreateUserProfileOptions)`, func() {
		bearerToken := "0ui9876453"
		createUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(createUserProfilePath))
				Expect(req.Method).To(Equal("PUT"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke CreateUserProfile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.CreateUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CreateUserProfileOptions model
				createUserProfileOptionsModel := new(usermanagementv1.CreateUserProfileOptions)
				createUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				createUserProfileOptionsModel.IamID = core.StringPtr("testString")
				createUserProfileOptionsModel.Realm = core.StringPtr("IBMid")
				createUserProfileOptionsModel.UserID = core.StringPtr("example@ibm.com")
				createUserProfileOptionsModel.Firstname = core.StringPtr("testString")
				createUserProfileOptionsModel.Lastname = core.StringPtr("testString")
				createUserProfileOptionsModel.State = core.StringPtr("testString")
				createUserProfileOptionsModel.Email = core.StringPtr("testString")
				createUserProfileOptionsModel.Phonenumber = core.StringPtr("testString")
				createUserProfileOptionsModel.Altphonenumber = core.StringPtr("testString")
				createUserProfileOptionsModel.Photo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.CreateUserProfile(createUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateUserProfile(updateUserProfileOptions *UpdateUserProfileOptions)`, func() {
		bearerToken := "0ui9876453"
		updateUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateUserProfilePath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke UpdateUserProfile successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserProfileOptions model
				updateUserProfileOptionsModel := new(usermanagementv1.UpdateUserProfileOptions)
				updateUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				updateUserProfileOptionsModel.IamID = core.StringPtr("testString")
				updateUserProfileOptionsModel.UserID = core.StringPtr("testString")
				updateUserProfileOptionsModel.Firstname = core.StringPtr("testString")
				updateUserProfileOptionsModel.Lastname = core.StringPtr("testString")
				updateUserProfileOptionsModel.State = core.StringPtr("testString")
				updateUserProfileOptionsModel.Email = core.StringPtr("testString")
				updateUserProfileOptionsModel.Phonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Altphonenumber = core.StringPtr("testString")
				updateUserProfileOptionsModel.Photo = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateUserProfile(updateUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {
		bearerToken := "0ui9876453"
		getUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getUserSettingsPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"language": "Language", "notification_language": "NotificationLanguage", "allowed_ip_addresses": "32.96.110.50,172.16.254.1", "self_manage": true}`)
			}))
			It(`Invoke GetUserSettings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions)`, func() {
		bearerToken := "0ui9876453"
		updateUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(updateUserSettingsPath))
				Expect(req.Method).To(Equal("PATCH"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke UpdateUserSettings successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.UpdateUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIpAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`RemoveUserFromAccount(removeUserFromAccountOptions *RemoveUserFromAccountOptions)`, func() {
		bearerToken := "0ui9876453"
		removeUserFromAccountPath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(removeUserFromAccountPath))
				Expect(req.Method).To(Equal("DELETE"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(204)
			}))
			It(`Invoke RemoveUserFromAccount successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.RemoveUserFromAccount(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveUserFromAccountOptions model
				removeUserFromAccountOptionsModel := new(usermanagementv1.RemoveUserFromAccountOptions)
				removeUserFromAccountOptionsModel.AccountID = core.StringPtr("testString")
				removeUserFromAccountOptionsModel.IamID = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.RemoveUserFromAccount(removeUserFromAccountOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {
		bearerToken := "0ui9876453"
		listUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(listUsersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["IAMid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["firstname"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["lastname"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["email"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["realm"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "linkages": [{"origin": "Origin", "id": "ID"}]}]}`)
			}))
			It(`Invoke ListUsers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.ListUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.IAMid = core.StringPtr("testString")
				listUsersOptionsModel.Firstname = core.StringPtr("testString")
				listUsersOptionsModel.Lastname = core.StringPtr("testString")
				listUsersOptionsModel.Email = core.StringPtr("testString")
				listUsersOptionsModel.State = core.StringPtr("testString")
				listUsersOptionsModel.Realm = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`InviteUsers(inviteUsersOptions *InviteUsersOptions)`, func() {
		bearerToken := "0ui9876453"
		inviteUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(inviteUsersPath))
				Expect(req.Method).To(Equal("POST"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.WriteHeader(202)
			}))
			It(`Invoke InviteUsers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := testService.InviteUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
		})
	})
	Describe(`GetImsUsers(getImsUsersOptions *GetImsUsersOptions)`, func() {
		bearerToken := "0ui9876453"
		getImsUsersPath := "/v2/accounts/testString/ims/users"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getImsUsersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				Expect(req.URL.Query()["IAMid"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["firstname"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["lastname"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["email"]).To(Equal([]string{"testString"}))

				Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "linkages": [{"origin": "Origin", "id": "ID"}]}]}`)
			}))
			It(`Invoke GetImsUsers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetImsUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetImsUsersOptions model
				getImsUsersOptionsModel := new(usermanagementv1.GetImsUsersOptions)
				getImsUsersOptionsModel.AccountID = core.StringPtr("testString")
				getImsUsersOptionsModel.IAMid = core.StringPtr("testString")
				getImsUsersOptionsModel.Firstname = core.StringPtr("testString")
				getImsUsersOptionsModel.Lastname = core.StringPtr("testString")
				getImsUsersOptionsModel.Email = core.StringPtr("testString")
				getImsUsersOptionsModel.State = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetImsUsers(getImsUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
	Describe(`GetCfUsers(getCfUsersOptions *GetCfUsersOptions)`, func() {
		bearerToken := "0ui9876453"
		getCfUsersPath := "/v2/accounts/testString/organizations/testString/users"
		Context(`Using mock server endpoint`, func() {
			testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				defer GinkgoRecover()

				// Verify the contents of the request
				Expect(req.URL.Path).To(Equal(getCfUsersPath))
				Expect(req.Method).To(Equal("GET"))
				Expect(req.Header["Authorization"]).ToNot(BeNil())
				Expect(req.Header["Authorization"][0]).To(Equal("Bearer " + bearerToken))
				res.Header().Set("Content-type", "application/json")
				res.WriteHeader(200)
				fmt.Fprintf(res, `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID", "linkages": [{"origin": "Origin", "id": "ID"}]}]}`)
			}))
			It(`Invoke GetCfUsers successfully`, func() {
				defer testServer.Close()

				testService, testServiceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.BearerTokenAuthenticator{
						BearerToken: bearerToken,
					},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := testService.GetCfUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCfUsersOptions model
				getCfUsersOptionsModel := new(usermanagementv1.GetCfUsersOptions)
				getCfUsersOptionsModel.AccountID = core.StringPtr("testString")
				getCfUsersOptionsModel.OrganizationGuid = core.StringPtr("testString")

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetCfUsers(getCfUsersOptionsModel)
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
