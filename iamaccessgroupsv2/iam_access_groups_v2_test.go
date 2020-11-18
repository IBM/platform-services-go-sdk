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
	"context"
	"fmt"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
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

var _ = Describe(`IamAccessGroupsV2`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamAccessGroupsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "https://iamaccessgroupsv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL: "https://testService/api",
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				err := iamAccessGroupsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions) - Operation response error`, func() {
		createAccessGroupPath := "/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("testString")
				createAccessGroupOptionsModel.Description = core.StringPtr("testString")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions)`, func() {
		createAccessGroupPath := "/groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke CreateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("testString")
				createAccessGroupOptionsModel.Description = core.StringPtr("testString")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("testString")
				createAccessGroupOptionsModel.Description = core.StringPtr("testString")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateAccessGroupOptions model with no property values
				createAccessGroupOptionsModelNew := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) - Operation response error`, func() {
		listAccessGroupsPath := "/groups"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions)`, func() {
		listAccessGroupsPath := "/groups"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}]}`)
				}))
			})
			It(`Invoke ListAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(nil)
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
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupsOptions model with no property values
				listAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions) - Operation response error`, func() {
		getAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for show_federated query parameter

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(true)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions)`, func() {
		getAccessGroupPath := "/groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for show_federated query parameter

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke GetAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(true)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(true)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessGroupOptions model with no property values
				getAccessGroupOptionsModelNew := new(iamaccessgroupsv2.GetAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions) - Operation response error`, func() {
		updateAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessGroupPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Description = core.StringPtr("testString")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions)`, func() {
		updateAccessGroupPath := "/groups/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccessGroupPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(nil)
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
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Description = core.StringPtr("testString")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccessGroupOptions model with no property values
				updateAccessGroupOptionsModelNew := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions)`, func() {
		deleteAccessGroupPath := "/groups/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteAccessGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for force query parameter

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(true)
				deleteAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				response, operationErr = iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(true)
				deleteAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteAccessGroupOptions model with no property values
				deleteAccessGroupOptionsModelNew := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptionsModelNew)
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
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamAccessGroupsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "https://iamaccessgroupsv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL: "https://testService/api",
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				err := iamAccessGroupsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/groups/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccountSettings with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {
		getAccountSettingsPath := "/groups/settings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAccountSettings with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccountSettingsOptions model with no property values
				getAccountSettingsOptionsModelNew := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) - Operation response error`, func() {
		updateAccountSettingsPath := "/groups/settings"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateAccountSettings with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions)`, func() {
		updateAccountSettingsPath := "/groups/settings"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateAccountSettingsPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateAccountSettings with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateAccountSettingsOptions model with no property values
				updateAccountSettingsOptionsModelNew := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModelNew)
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
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamAccessGroupsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "https://iamaccessgroupsv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL: "https://testService/api",
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				err := iamAccessGroupsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})

	Describe(`IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions)`, func() {
		isMemberOfAccessGroupPath := "/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(isMemberOfAccessGroupPath))
					Expect(req.Method).To(Equal("HEAD"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke IsMemberOfAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.IsMemberOfAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the IsMemberOfAccessGroupOptions model
				isMemberOfAccessGroupOptionsModel := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				isMemberOfAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				response, operationErr = iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke IsMemberOfAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the IsMemberOfAccessGroupOptions model
				isMemberOfAccessGroupOptionsModel := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				isMemberOfAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				isMemberOfAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the IsMemberOfAccessGroupOptions model with no property values
				isMemberOfAccessGroupOptionsModelNew := new(iamaccessgroupsv2.IsMemberOfAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions) - Operation response error`, func() {
		addMembersToAccessGroupPath := "/groups/testString/members"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMembersToAccessGroupPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("testString")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("testString")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddMembersToAccessGroup(addMembersToAccessGroupOptions *AddMembersToAccessGroupOptions)`, func() {
		addMembersToAccessGroupPath := "/groups/testString/members"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMembersToAccessGroupPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(nil)
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
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddMembersToAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("testString")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("testString")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddMembersToAccessGroupOptions model with no property values
				addMembersToAccessGroupOptionsModelNew := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions) - Operation response error`, func() {
		listAccessGroupMembersPath := "/groups/testString/members"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroupMembers with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Limit = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Offset = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(true)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions)`, func() {
		listAccessGroupMembersPath := "/groups/testString/members"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))

					// TODO: Add check for limit query parameter


					// TODO: Add check for offset query parameter

					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))


					// TODO: Add check for verbose query parameter

					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Limit = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Offset = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(true)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAccessGroupMembers with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Limit = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Offset = core.Float64Ptr(float64(72.5))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(true)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupMembersOptions model with no property values
				listAccessGroupMembersOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions)`, func() {
		removeMemberFromAccessGroupPath := "/groups/testString/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAccessGroupPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveMemberFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.RemoveMemberFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				removeMemberFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				removeMemberFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				response, operationErr = iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveMemberFromAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				removeMemberFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				removeMemberFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveMemberFromAccessGroupOptions model with no property values
				removeMemberFromAccessGroupOptionsModelNew := new(iamaccessgroupsv2.RemoveMemberFromAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions) - Operation response error`, func() {
		removeMembersFromAccessGroupPath := "/groups/testString/members/delete"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMembersFromAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"testString"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptions *RemoveMembersFromAccessGroupOptions)`, func() {
		removeMembersFromAccessGroupPath := "/groups/testString/members/delete"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMembersFromAccessGroupPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"access_group_id": "AccessGroupID", "members": [{"iam_id": "IamID", "trace": "Trace", "status_code": 10, "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"testString"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RemoveMembersFromAccessGroup with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"testString"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveMembersFromAccessGroupOptions model with no property values
				removeMembersFromAccessGroupOptionsModelNew := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) - Operation response error`, func() {
		removeMemberFromAllAccessGroupsPath := "/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions)`, func() {
		removeMemberFromAllAccessGroupsPath := "/groups/_allgroups/members/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke RemoveMemberFromAllAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the RemoveMemberFromAllAccessGroupsOptions model with no property values
				removeMemberFromAllAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) - Operation response error`, func() {
		addMemberToMultipleAccessGroupsPath := "/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMemberToMultipleAccessGroupsPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"testString"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions)`, func() {
		addMemberToMultipleAccessGroupsPath := "/groups/_allgroups/members/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addMemberToMultipleAccessGroupsPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(nil)
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
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddMemberToMultipleAccessGroups with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"testString"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddMemberToMultipleAccessGroupsOptions model with no property values
				addMemberToMultipleAccessGroupsOptionsModelNew := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModelNew)
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
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamAccessGroupsService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "https://iamaccessgroupsv2/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamAccessGroupsService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL: "https://testService/api",
				})
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				})
				err := iamAccessGroupsService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_URL": "https://iamaccessgroupsv2/api",
				"IAM_ACCESS_GROUPS_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_ACCESS_GROUPS_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamAccessGroupsService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions) - Operation response error`, func() {
		addAccessGroupRulePath := "/groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAccessGroupRulePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions)`, func() {
		addAccessGroupRulePath := "/groups/testString/rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addAccessGroupRulePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(nil)
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
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke AddAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddAccessGroupRuleOptions model with no property values
				addAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions) - Operation response error`, func() {
		listAccessGroupRulesPath := "/groups/testString/rules"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAccessGroupRules with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions)`, func() {
		listAccessGroupRulesPath := "/groups/testString/rules"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAccessGroupRules with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListAccessGroupRulesOptions model with no property values
				listAccessGroupRulesOptionsModelNew := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions) - Operation response error`, func() {
		getAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions)`, func() {
		getAccessGroupRulePath := "/groups/testString/rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetAccessGroupRuleOptions model with no property values
				getAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) - Operation response error`, func() {
		replaceAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAccessGroupRulePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule with error: Operation response processing error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamAccessGroupsService.EnableRetries(0, 0)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions)`, func() {
		replaceAccessGroupRulePath := "/groups/testString/rules/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceAccessGroupRulePath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "Operator", "value": "Value"}], "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(nil)
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
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ReplaceAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceAccessGroupRuleOptions model with no property values
				replaceAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions)`, func() {
		removeAccessGroupRulePath := "/groups/testString/rules/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeAccessGroupRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke RemoveAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.RemoveAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the RemoveAccessGroupRuleOptions model
				removeAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				removeAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				response, operationErr = iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke RemoveAccessGroupRule with error: Operation validation and request error`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveAccessGroupRuleOptions model
				removeAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				removeAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				removeAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamAccessGroupsService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the RemoveAccessGroupRuleOptions model with no property values
				removeAccessGroupRuleOptionsModelNew := new(iamaccessgroupsv2.RemoveAccessGroupRuleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptionsModelNew)
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
			iamAccessGroupsService, _ := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
				URL:           "http://iamaccessgroupsv2modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				Expect(ruleConditionsModel).ToNot(BeNil())
				ruleConditionsModel.Claim = core.StringPtr("testString")
				ruleConditionsModel.Operator = core.StringPtr("testString")
				ruleConditionsModel.Value = core.StringPtr("testString")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddAccessGroupRuleOptions model
				accessGroupID := "testString"
				addAccessGroupRuleOptionsExpiration := int64(38)
				addAccessGroupRuleOptionsRealmName := "testString"
				addAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				addAccessGroupRuleOptionsModel := iamAccessGroupsService.NewAddAccessGroupRuleOptions(accessGroupID, addAccessGroupRuleOptionsExpiration, addAccessGroupRuleOptionsRealmName, addAccessGroupRuleOptionsConditions)
				addAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				addAccessGroupRuleOptionsModel.SetExpiration(int64(38))
				addAccessGroupRuleOptionsModel.SetRealmName("testString")
				addAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				addAccessGroupRuleOptionsModel.SetName("testString")
				addAccessGroupRuleOptionsModel.SetTransactionID("testString")
				addAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(addAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(addAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(addAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddGroupMembersRequestMembersItem successfully`, func() {
				iamID := "testString"
				typeVar := "testString"
				model, err := iamAccessGroupsService.NewAddGroupMembersRequestMembersItem(iamID, typeVar)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddMemberToMultipleAccessGroupsOptions successfully`, func() {
				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				accountID := "testString"
				iamID := "testString"
				addMemberToMultipleAccessGroupsOptionsModel := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(accountID, iamID)
				addMemberToMultipleAccessGroupsOptionsModel.SetAccountID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetIamID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetType("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetGroups([]string{"testString"})
				addMemberToMultipleAccessGroupsOptionsModel.SetTransactionID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMemberToMultipleAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(addMemberToMultipleAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Groups).To(Equal([]string{"testString"}))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddMembersToAccessGroupOptions successfully`, func() {
				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				Expect(addGroupMembersRequestMembersItemModel).ToNot(BeNil())
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("testString")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("testString")
				Expect(addGroupMembersRequestMembersItemModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(addGroupMembersRequestMembersItemModel.Type).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddMembersToAccessGroupOptions model
				accessGroupID := "testString"
				addMembersToAccessGroupOptionsModel := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(accessGroupID)
				addMembersToAccessGroupOptionsModel.SetAccessGroupID("testString")
				addMembersToAccessGroupOptionsModel.SetMembers([]iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel})
				addMembersToAccessGroupOptionsModel.SetTransactionID("testString")
				addMembersToAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMembersToAccessGroupOptionsModel).ToNot(BeNil())
				Expect(addMembersToAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(addMembersToAccessGroupOptionsModel.Members).To(Equal([]iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}))
				Expect(addMembersToAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addMembersToAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateAccessGroupOptions successfully`, func() {
				// Construct an instance of the CreateAccessGroupOptions model
				accountID := "testString"
				createAccessGroupOptionsName := "testString"
				createAccessGroupOptionsModel := iamAccessGroupsService.NewCreateAccessGroupOptions(accountID, createAccessGroupOptionsName)
				createAccessGroupOptionsModel.SetAccountID("testString")
				createAccessGroupOptionsModel.SetName("testString")
				createAccessGroupOptionsModel.SetDescription("testString")
				createAccessGroupOptionsModel.SetTransactionID("testString")
				createAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessGroupOptionsModel).ToNot(BeNil())
				Expect(createAccessGroupOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessGroupOptions successfully`, func() {
				// Construct an instance of the DeleteAccessGroupOptions model
				accessGroupID := "testString"
				deleteAccessGroupOptionsModel := iamAccessGroupsService.NewDeleteAccessGroupOptions(accessGroupID)
				deleteAccessGroupOptionsModel.SetAccessGroupID("testString")
				deleteAccessGroupOptionsModel.SetTransactionID("testString")
				deleteAccessGroupOptionsModel.SetForce(true)
				deleteAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessGroupOptionsModel).ToNot(BeNil())
				Expect(deleteAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.Force).To(Equal(core.BoolPtr(true)))
				Expect(deleteAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessGroupOptions successfully`, func() {
				// Construct an instance of the GetAccessGroupOptions model
				accessGroupID := "testString"
				getAccessGroupOptionsModel := iamAccessGroupsService.NewGetAccessGroupOptions(accessGroupID)
				getAccessGroupOptionsModel.SetAccessGroupID("testString")
				getAccessGroupOptionsModel.SetTransactionID("testString")
				getAccessGroupOptionsModel.SetShowFederated(true)
				getAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessGroupOptionsModel).ToNot(BeNil())
				Expect(getAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.ShowFederated).To(Equal(core.BoolPtr(true)))
				Expect(getAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the GetAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				getAccessGroupRuleOptionsModel := iamAccessGroupsService.NewGetAccessGroupRuleOptions(accessGroupID, ruleID)
				getAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				getAccessGroupRuleOptionsModel.SetRuleID("testString")
				getAccessGroupRuleOptionsModel.SetTransactionID("testString")
				getAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(getAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccountSettingsOptions successfully`, func() {
				// Construct an instance of the GetAccountSettingsOptions model
				accountID := "testString"
				getAccountSettingsOptionsModel := iamAccessGroupsService.NewGetAccountSettingsOptions(accountID)
				getAccountSettingsOptionsModel.SetAccountID("testString")
				getAccountSettingsOptionsModel.SetTransactionID("testString")
				getAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(getAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewIsMemberOfAccessGroupOptions successfully`, func() {
				// Construct an instance of the IsMemberOfAccessGroupOptions model
				accessGroupID := "testString"
				iamID := "testString"
				isMemberOfAccessGroupOptionsModel := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(accessGroupID, iamID)
				isMemberOfAccessGroupOptionsModel.SetAccessGroupID("testString")
				isMemberOfAccessGroupOptionsModel.SetIamID("testString")
				isMemberOfAccessGroupOptionsModel.SetTransactionID("testString")
				isMemberOfAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(isMemberOfAccessGroupOptionsModel).ToNot(BeNil())
				Expect(isMemberOfAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(isMemberOfAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupMembersOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupMembersOptions model
				accessGroupID := "testString"
				listAccessGroupMembersOptionsModel := iamAccessGroupsService.NewListAccessGroupMembersOptions(accessGroupID)
				listAccessGroupMembersOptionsModel.SetAccessGroupID("testString")
				listAccessGroupMembersOptionsModel.SetTransactionID("testString")
				listAccessGroupMembersOptionsModel.SetLimit(float64(72.5))
				listAccessGroupMembersOptionsModel.SetOffset(float64(72.5))
				listAccessGroupMembersOptionsModel.SetType("testString")
				listAccessGroupMembersOptionsModel.SetVerbose(true)
				listAccessGroupMembersOptionsModel.SetSort("testString")
				listAccessGroupMembersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupMembersOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupMembersOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Limit).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(listAccessGroupMembersOptionsModel.Offset).To(Equal(core.Float64Ptr(float64(72.5))))
				Expect(listAccessGroupMembersOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Verbose).To(Equal(core.BoolPtr(true)))
				Expect(listAccessGroupMembersOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupRulesOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupRulesOptions model
				accessGroupID := "testString"
				listAccessGroupRulesOptionsModel := iamAccessGroupsService.NewListAccessGroupRulesOptions(accessGroupID)
				listAccessGroupRulesOptionsModel.SetAccessGroupID("testString")
				listAccessGroupRulesOptionsModel.SetTransactionID("testString")
				listAccessGroupRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupRulesOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupRulesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupRulesOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAccessGroupsOptions successfully`, func() {
				// Construct an instance of the ListAccessGroupsOptions model
				accountID := "testString"
				listAccessGroupsOptionsModel := iamAccessGroupsService.NewListAccessGroupsOptions(accountID)
				listAccessGroupsOptionsModel.SetAccountID("testString")
				listAccessGroupsOptionsModel.SetTransactionID("testString")
				listAccessGroupsOptionsModel.SetIamID("testString")
				listAccessGroupsOptionsModel.SetLimit(int64(38))
				listAccessGroupsOptionsModel.SetOffset(int64(38))
				listAccessGroupsOptionsModel.SetSort("testString")
				listAccessGroupsOptionsModel.SetShowFederated(true)
				listAccessGroupsOptionsModel.SetHidePublicAccess(true)
				listAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccessGroupsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccessGroupsOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.ShowFederated).To(Equal(core.BoolPtr(true)))
				Expect(listAccessGroupsOptionsModel.HidePublicAccess).To(Equal(core.BoolPtr(true)))
				Expect(listAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RemoveAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				removeAccessGroupRuleOptionsModel := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(accessGroupID, ruleID)
				removeAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				removeAccessGroupRuleOptionsModel.SetRuleID("testString")
				removeAccessGroupRuleOptionsModel.SetTransactionID("testString")
				removeAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(removeAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMemberFromAccessGroupOptions successfully`, func() {
				// Construct an instance of the RemoveMemberFromAccessGroupOptions model
				accessGroupID := "testString"
				iamID := "testString"
				removeMemberFromAccessGroupOptionsModel := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(accessGroupID, iamID)
				removeMemberFromAccessGroupOptionsModel.SetAccessGroupID("testString")
				removeMemberFromAccessGroupOptionsModel.SetIamID("testString")
				removeMemberFromAccessGroupOptionsModel.SetTransactionID("testString")
				removeMemberFromAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMemberFromAccessGroupOptionsModel).ToNot(BeNil())
				Expect(removeMemberFromAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMemberFromAllAccessGroupsOptions successfully`, func() {
				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				accountID := "testString"
				iamID := "testString"
				removeMemberFromAllAccessGroupsOptionsModel := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(accountID, iamID)
				removeMemberFromAllAccessGroupsOptionsModel.SetAccountID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetIamID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetTransactionID("testString")
				removeMemberFromAllAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMemberFromAllAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(removeMemberFromAllAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMemberFromAllAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRemoveMembersFromAccessGroupOptions successfully`, func() {
				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				accessGroupID := "testString"
				removeMembersFromAccessGroupOptionsModel := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(accessGroupID)
				removeMembersFromAccessGroupOptionsModel.SetAccessGroupID("testString")
				removeMembersFromAccessGroupOptionsModel.SetMembers([]string{"testString"})
				removeMembersFromAccessGroupOptionsModel.SetTransactionID("testString")
				removeMembersFromAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMembersFromAccessGroupOptionsModel).ToNot(BeNil())
				Expect(removeMembersFromAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Members).To(Equal([]string{"testString"}))
				Expect(removeMembersFromAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				Expect(ruleConditionsModel).ToNot(BeNil())
				ruleConditionsModel.Claim = core.StringPtr("testString")
				ruleConditionsModel.Operator = core.StringPtr("testString")
				ruleConditionsModel.Value = core.StringPtr("testString")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				ifMatch := "testString"
				replaceAccessGroupRuleOptionsExpiration := int64(38)
				replaceAccessGroupRuleOptionsRealmName := "testString"
				replaceAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				replaceAccessGroupRuleOptionsModel := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(accessGroupID, ruleID, ifMatch, replaceAccessGroupRuleOptionsExpiration, replaceAccessGroupRuleOptionsRealmName, replaceAccessGroupRuleOptionsConditions)
				replaceAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				replaceAccessGroupRuleOptionsModel.SetRuleID("testString")
				replaceAccessGroupRuleOptionsModel.SetIfMatch("testString")
				replaceAccessGroupRuleOptionsModel.SetExpiration(int64(38))
				replaceAccessGroupRuleOptionsModel.SetRealmName("testString")
				replaceAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				replaceAccessGroupRuleOptionsModel.SetName("testString")
				replaceAccessGroupRuleOptionsModel.SetTransactionID("testString")
				replaceAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(replaceAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(38))))
				Expect(replaceAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(replaceAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleConditions successfully`, func() {
				claim := "testString"
				operator := "testString"
				value := "testString"
				model, err := iamAccessGroupsService.NewRuleConditions(claim, operator, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAccessGroupOptions successfully`, func() {
				// Construct an instance of the UpdateAccessGroupOptions model
				accessGroupID := "testString"
				ifMatch := "testString"
				updateAccessGroupOptionsModel := iamAccessGroupsService.NewUpdateAccessGroupOptions(accessGroupID, ifMatch)
				updateAccessGroupOptionsModel.SetAccessGroupID("testString")
				updateAccessGroupOptionsModel.SetIfMatch("testString")
				updateAccessGroupOptionsModel.SetName("testString")
				updateAccessGroupOptionsModel.SetDescription("testString")
				updateAccessGroupOptionsModel.SetTransactionID("testString")
				updateAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessGroupOptionsModel).ToNot(BeNil())
				Expect(updateAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateAccountSettingsOptions successfully`, func() {
				// Construct an instance of the UpdateAccountSettingsOptions model
				accountID := "testString"
				updateAccountSettingsOptionsModel := iamAccessGroupsService.NewUpdateAccountSettingsOptions(accountID)
				updateAccountSettingsOptionsModel.SetAccountID("testString")
				updateAccountSettingsOptionsModel.SetPublicAccessEnabled(true)
				updateAccountSettingsOptionsModel.SetTransactionID("testString")
				updateAccountSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccountSettingsOptionsModel).ToNot(BeNil())
				Expect(updateAccountSettingsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.PublicAccessEnabled).To(Equal(core.BoolPtr(true)))
				Expect(updateAccountSettingsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccountSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
