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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
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

var _ = Describe(`UserManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(userManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "https://usermanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL": "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				})
				err := userManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL": "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions) - Operation response error`, func() {
		getUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserSettings with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserSettings(getUserSettingsOptions *GetUserSettingsOptions)`, func() {
		getUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "notification_language": "NotificationLanguage", "allowed_ip_addresses": "32.96.110.50,172.16.254.1", "self_manage": true}`)
				}))
			})
			It(`Invoke GetUserSettings successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.GetUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUserSettings with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserSettingsOptions model
				getUserSettingsOptionsModel := new(usermanagementv1.GetUserSettingsOptions)
				getUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				getUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.GetUserSettings(getUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserSettingsOptions model with no property values
				getUserSettingsOptionsModelNew := new(usermanagementv1.GetUserSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.GetUserSettings(getUserSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions) - Operation response error`, func() {
		updateUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateUserSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateUserSettings with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIpAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)
				updateUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserSettings(updateUserSettingsOptions *UpdateUserSettingsOptions)`, func() {
		updateUserSettingsPath := "/v2/accounts/testString/users/testString/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateUserSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"language": "Language", "notification_language": "NotificationLanguage", "allowed_ip_addresses": "32.96.110.50,172.16.254.1", "self_manage": true}`)
				}))
			})
			It(`Invoke UpdateUserSettings successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.UpdateUserSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIpAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)
				updateUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateUserSettings with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserSettingsOptions model
				updateUserSettingsOptionsModel := new(usermanagementv1.UpdateUserSettingsOptions)
				updateUserSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.IamID = core.StringPtr("testString")
				updateUserSettingsOptionsModel.Language = core.StringPtr("testString")
				updateUserSettingsOptionsModel.NotificationLanguage = core.StringPtr("testString")
				updateUserSettingsOptionsModel.AllowedIpAddresses = core.StringPtr("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SelfManage = core.BoolPtr(true)
				updateUserSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.UpdateUserSettings(updateUserSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateUserSettingsOptions model with no property values
				updateUserSettingsOptionsModelNew := new(usermanagementv1.UpdateUserSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.UpdateUserSettings(updateUserSettingsOptionsModelNew)
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
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(userManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL: "https://usermanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(userManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL": "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				})
				err := userManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_URL": "https://usermanagementv1/api",
				"USER_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"USER_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			userManagementService, serviceErr := usermanagementv1.NewUserManagementV1UsingExternalConfig(&usermanagementv1.UserManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(userManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListUsers(listUsersOptions *ListUsersOptions) - Operation response error`, func() {
		listUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListUsers with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.State = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListUsers(listUsersOptions *ListUsersOptions)`, func() {
		listUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(listUsersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID"}]}`)
				}))
			})
			It(`Invoke ListUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.ListUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.State = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListUsers with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the ListUsersOptions model
				listUsersOptionsModel := new(usermanagementv1.ListUsersOptions)
				listUsersOptionsModel.AccountID = core.StringPtr("testString")
				listUsersOptionsModel.State = core.StringPtr("testString")
				listUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.ListUsers(listUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListUsersOptions model with no property values
				listUsersOptionsModelNew := new(usermanagementv1.ListUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.ListUsers(listUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`InviteUsers(inviteUsersOptions *InviteUsersOptions) - Operation response error`, func() {
		inviteUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(inviteUsersPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke InviteUsers with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`InviteUsers(inviteUsersOptions *InviteUsersOptions)`, func() {
		inviteUsersPath := "/v2/accounts/testString/users"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(inviteUsersPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"total_results": 12, "limit": 5, "first_url": "FirstURL", "next_url": "NextURL", "resources": [{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID"}]}`)
				}))
			})
			It(`Invoke InviteUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.InviteUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke InviteUsers with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				roleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}

				// Construct an instance of the InviteUsersOptions model
				inviteUsersOptionsModel := new(usermanagementv1.InviteUsersOptions)
				inviteUsersOptionsModel.AccountID = core.StringPtr("testString")
				inviteUsersOptionsModel.Users = []usermanagementv1.InviteUser{*inviteUserModel}
				inviteUsersOptionsModel.IamPolicy = []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}
				inviteUsersOptionsModel.AccessGroups = []string{"testString"}
				inviteUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.InviteUsers(inviteUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the InviteUsersOptions model with no property values
				inviteUsersOptionsModelNew := new(usermanagementv1.InviteUsersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.InviteUsers(inviteUsersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions) - Operation response error`, func() {
		getUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserProfilePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetUserProfile with error: Operation response processing error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetUserProfile(getUserProfileOptions *GetUserProfileOptions)`, func() {
		getUserProfilePath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getUserProfilePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "iam_id": "IamID", "realm": "Realm", "user_id": "UserID", "firstname": "Firstname", "lastname": "Lastname", "state": "State", "email": "Email", "phonenumber": "Phonenumber", "altphonenumber": "Altphonenumber", "photo": "Photo", "account_id": "AccountID"}`)
				}))
			})
			It(`Invoke GetUserProfile successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := userManagementService.GetUserProfile(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetUserProfile with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the GetUserProfileOptions model
				getUserProfileOptionsModel := new(usermanagementv1.GetUserProfileOptions)
				getUserProfileOptionsModel.AccountID = core.StringPtr("testString")
				getUserProfileOptionsModel.IamID = core.StringPtr("testString")
				getUserProfileOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := userManagementService.GetUserProfile(getUserProfileOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetUserProfileOptions model with no property values
				getUserProfileOptionsModelNew := new(usermanagementv1.GetUserProfileOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = userManagementService.GetUserProfile(getUserProfileOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateUserProfiles(updateUserProfilesOptions *UpdateUserProfilesOptions)`, func() {
		updateUserProfilesPath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateUserProfilesPath))
					Expect(req.Method).To(Equal("PATCH"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke UpdateUserProfiles successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.UpdateUserProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the UpdateUserProfilesOptions model
				updateUserProfilesOptionsModel := new(usermanagementv1.UpdateUserProfilesOptions)
				updateUserProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateUserProfilesOptionsModel.IamID = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Firstname = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Lastname = core.StringPtr("testString")
				updateUserProfilesOptionsModel.State = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Email = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Phonenumber = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Altphonenumber = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Photo = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.UpdateUserProfiles(updateUserProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateUserProfiles with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateUserProfilesOptions model
				updateUserProfilesOptionsModel := new(usermanagementv1.UpdateUserProfilesOptions)
				updateUserProfilesOptionsModel.AccountID = core.StringPtr("testString")
				updateUserProfilesOptionsModel.IamID = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Firstname = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Lastname = core.StringPtr("testString")
				updateUserProfilesOptionsModel.State = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Email = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Phonenumber = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Altphonenumber = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Photo = core.StringPtr("testString")
				updateUserProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.UpdateUserProfiles(updateUserProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateUserProfilesOptions model with no property values
				updateUserProfilesOptionsModelNew := new(usermanagementv1.UpdateUserProfilesOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.UpdateUserProfiles(updateUserProfilesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveUsers(removeUsersOptions *RemoveUsersOptions)`, func() {
		removeUsersPath := "/v2/accounts/testString/users/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(removeUsersPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveUsers successfully`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := userManagementService.RemoveUsers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveUsersOptions model
				removeUsersOptionsModel := new(usermanagementv1.RemoveUsersOptions)
				removeUsersOptionsModel.AccountID = core.StringPtr("testString")
				removeUsersOptionsModel.IamID = core.StringPtr("testString")
				removeUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = userManagementService.RemoveUsers(removeUsersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveUsers with error: Operation validation and request error`, func() {
				userManagementService, serviceErr := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(userManagementService).ToNot(BeNil())

				// Construct an instance of the RemoveUsersOptions model
				removeUsersOptionsModel := new(usermanagementv1.RemoveUsersOptions)
				removeUsersOptionsModel.AccountID = core.StringPtr("testString")
				removeUsersOptionsModel.IamID = core.StringPtr("testString")
				removeUsersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := userManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := userManagementService.RemoveUsers(removeUsersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveUsersOptions model with no property values
				removeUsersOptionsModelNew := new(usermanagementv1.RemoveUsersOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = userManagementService.RemoveUsers(removeUsersOptionsModelNew)
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
			userManagementService, _ := usermanagementv1.NewUserManagementV1(&usermanagementv1.UserManagementV1Options{
				URL:           "http://usermanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewGetUserProfileOptions successfully`, func() {
				// Construct an instance of the GetUserProfileOptions model
				accountID := "testString"
				iamID := "testString"
				getUserProfileOptionsModel := userManagementService.NewGetUserProfileOptions(accountID, iamID)
				getUserProfileOptionsModel.SetAccountID("testString")
				getUserProfileOptionsModel.SetIamID("testString")
				getUserProfileOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserProfileOptionsModel).ToNot(BeNil())
				Expect(getUserProfileOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getUserProfileOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(getUserProfileOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetUserSettingsOptions successfully`, func() {
				// Construct an instance of the GetUserSettingsOptions model
				accountID := "testString"
				iamID := "testString"
				getUserSettingsOptionsModel := userManagementService.NewGetUserSettingsOptions(accountID, iamID)
				getUserSettingsOptionsModel.SetAccountID("testString")
				getUserSettingsOptionsModel.SetIamID("testString")
				getUserSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getUserSettingsOptionsModel).ToNot(BeNil())
				Expect(getUserSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getUserSettingsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(getUserSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewInviteUsersOptions successfully`, func() {
				// Construct an instance of the Attribute model
				attributeModel := new(usermanagementv1.Attribute)
				Expect(attributeModel).ToNot(BeNil())
				attributeModel.Name = core.StringPtr("testString")
				attributeModel.Value = core.StringPtr("testString")
				Expect(attributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(attributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Resource model
				resourceModel := new(usermanagementv1.Resource)
				Expect(resourceModel).ToNot(BeNil())
				resourceModel.Attributes = []usermanagementv1.Attribute{*attributeModel}
				Expect(resourceModel.Attributes).To(Equal([]usermanagementv1.Attribute{*attributeModel}))

				// Construct an instance of the Role model
				roleModel := new(usermanagementv1.Role)
				Expect(roleModel).ToNot(BeNil())
				roleModel.RoleID = core.StringPtr("testString")
				Expect(roleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InviteUser model
				inviteUserModel := new(usermanagementv1.InviteUser)
				Expect(inviteUserModel).ToNot(BeNil())
				inviteUserModel.Email = core.StringPtr("testString")
				inviteUserModel.AccountRole = core.StringPtr("testString")
				Expect(inviteUserModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(inviteUserModel.AccountRole).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the InviteUserIamPolicy model
				inviteUserIamPolicyModel := new(usermanagementv1.InviteUserIamPolicy)
				Expect(inviteUserIamPolicyModel).ToNot(BeNil())
				inviteUserIamPolicyModel.Roles = []usermanagementv1.Role{*roleModel}
				inviteUserIamPolicyModel.Resources = []usermanagementv1.Resource{*resourceModel}
				Expect(inviteUserIamPolicyModel.Roles).To(Equal([]usermanagementv1.Role{*roleModel}))
				Expect(inviteUserIamPolicyModel.Resources).To(Equal([]usermanagementv1.Resource{*resourceModel}))

				// Construct an instance of the InviteUsersOptions model
				accountID := "testString"
				inviteUsersOptionsModel := userManagementService.NewInviteUsersOptions(accountID)
				inviteUsersOptionsModel.SetAccountID("testString")
				inviteUsersOptionsModel.SetUsers([]usermanagementv1.InviteUser{*inviteUserModel})
				inviteUsersOptionsModel.SetIamPolicy([]usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel})
				inviteUsersOptionsModel.SetAccessGroups([]string{"testString"})
				inviteUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(inviteUsersOptionsModel).ToNot(BeNil())
				Expect(inviteUsersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(inviteUsersOptionsModel.Users).To(Equal([]usermanagementv1.InviteUser{*inviteUserModel}))
				Expect(inviteUsersOptionsModel.IamPolicy).To(Equal([]usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel}))
				Expect(inviteUsersOptionsModel.AccessGroups).To(Equal([]string{"testString"}))
				Expect(inviteUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListUsersOptions successfully`, func() {
				// Construct an instance of the ListUsersOptions model
				accountID := "testString"
				listUsersOptionsModel := userManagementService.NewListUsersOptions(accountID)
				listUsersOptionsModel.SetAccountID("testString")
				listUsersOptionsModel.SetState("testString")
				listUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listUsersOptionsModel).ToNot(BeNil())
				Expect(listUsersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.State).To(Equal(core.StringPtr("testString")))
				Expect(listUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveUsersOptions successfully`, func() {
				// Construct an instance of the RemoveUsersOptions model
				accountID := "testString"
				iamID := "testString"
				removeUsersOptionsModel := userManagementService.NewRemoveUsersOptions(accountID, iamID)
				removeUsersOptionsModel.SetAccountID("testString")
				removeUsersOptionsModel.SetIamID("testString")
				removeUsersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeUsersOptionsModel).ToNot(BeNil())
				Expect(removeUsersOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(removeUsersOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeUsersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserProfilesOptions successfully`, func() {
				// Construct an instance of the UpdateUserProfilesOptions model
				accountID := "testString"
				iamID := "testString"
				updateUserProfilesOptionsModel := userManagementService.NewUpdateUserProfilesOptions(accountID, iamID)
				updateUserProfilesOptionsModel.SetAccountID("testString")
				updateUserProfilesOptionsModel.SetIamID("testString")
				updateUserProfilesOptionsModel.SetFirstname("testString")
				updateUserProfilesOptionsModel.SetLastname("testString")
				updateUserProfilesOptionsModel.SetState("testString")
				updateUserProfilesOptionsModel.SetEmail("testString")
				updateUserProfilesOptionsModel.SetPhonenumber("testString")
				updateUserProfilesOptionsModel.SetAltphonenumber("testString")
				updateUserProfilesOptionsModel.SetPhoto("testString")
				updateUserProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserProfilesOptionsModel).ToNot(BeNil())
				Expect(updateUserProfilesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Firstname).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Lastname).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.State).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Email).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Phonenumber).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Altphonenumber).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Photo).To(Equal(core.StringPtr("testString")))
				Expect(updateUserProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateUserSettingsOptions successfully`, func() {
				// Construct an instance of the UpdateUserSettingsOptions model
				accountID := "testString"
				iamID := "testString"
				updateUserSettingsOptionsModel := userManagementService.NewUpdateUserSettingsOptions(accountID, iamID)
				updateUserSettingsOptionsModel.SetAccountID("testString")
				updateUserSettingsOptionsModel.SetIamID("testString")
				updateUserSettingsOptionsModel.SetLanguage("testString")
				updateUserSettingsOptionsModel.SetNotificationLanguage("testString")
				updateUserSettingsOptionsModel.SetAllowedIpAddresses("32.96.110.50,172.16.254.1")
				updateUserSettingsOptionsModel.SetSelfManage(true)
				updateUserSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateUserSettingsOptionsModel).ToNot(BeNil())
				Expect(updateUserSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.Language).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.NotificationLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateUserSettingsOptionsModel.AllowedIpAddresses).To(Equal(core.StringPtr("32.96.110.50,172.16.254.1")))
				Expect(updateUserSettingsOptionsModel.SelfManage).To(Equal(core.BoolPtr(true)))
				Expect(updateUserSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
