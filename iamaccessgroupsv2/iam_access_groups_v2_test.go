/**
 * (C) Copyright IBM Corp. 2022.
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
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
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

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
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

				clone := iamAccessGroupsService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamAccessGroupsService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamAccessGroupsService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamAccessGroupsService.Service.Options.Authenticator))
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
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = iamaccessgroupsv2.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateAccessGroup(createAccessGroupOptions *CreateAccessGroupOptions) - Operation response error`, func() {
		createAccessGroupPath := "/v2/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
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
		createAccessGroupPath := "/v2/groups"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke CreateAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.CreateAccessGroupWithContext(ctx, createAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke CreateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the CreateAccessGroupOptions model
				createAccessGroupOptionsModel := new(iamaccessgroupsv2.CreateAccessGroupOptions)
				createAccessGroupOptionsModel.AccountID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Name = core.StringPtr("Managers")
				createAccessGroupOptionsModel.Description = core.StringPtr("Group for managers")
				createAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				createAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) - Operation response error`, func() {
		listAccessGroupsPath := "/v2/groups"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
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
		listAccessGroupsPath := "/v2/groups"
		Context(`Using mock server endpoint with timeout`, func() {
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
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}]}`)
				}))
			})
			It(`Invoke ListAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupsOptions model
				listAccessGroupsOptionsModel := new(iamaccessgroupsv2.ListAccessGroupsOptions)
				listAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupsWithContext(ctx, listAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
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
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"name"}))
					// TODO: Add check for show_federated query parameter
					// TODO: Add check for hide_public_access query parameter
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "groups": [{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}]}`)
				}))
			})
			It(`Invoke ListAccessGroups successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAccessGroups successfully`, func() {
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
				listAccessGroupsOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupsOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupsOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupsOptionsModel.Sort = core.StringPtr("name")
				listAccessGroupsOptionsModel.ShowFederated = core.BoolPtr(false)
				listAccessGroupsOptionsModel.HidePublicAccess = core.BoolPtr(false)
				listAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroups(listAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupsList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"limit":1,"groups":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID","href":"Href","is_federated":false}]}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"limit":1,"groups":[{"id":"ID","name":"Name","description":"Description","account_id":"AccountID","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID","last_modified_at":"2019-01-01T12:00:00.000Z","last_modified_by_id":"LastModifiedByID","href":"Href","is_federated":false}]}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccessGroupsPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupsOptionsModel := &iamaccessgroupsv2.ListAccessGroupsOptions{
					AccountID: core.StringPtr("testString"),
					TransactionID: core.StringPtr("testString"),
					IamID: core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: core.StringPtr("name"),
					ShowFederated: core.BoolPtr(false),
					HidePublicAccess: core.BoolPtr(false),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.Group
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccessGroupsPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupsOptionsModel := &iamaccessgroupsv2.ListAccessGroupsOptions{
					AccountID: core.StringPtr("testString"),
					TransactionID: core.StringPtr("testString"),
					IamID: core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit: core.Int64Ptr(int64(10)),
					Sort: core.StringPtr("name"),
					ShowFederated: core.BoolPtr(false),
					HidePublicAccess: core.BoolPtr(false),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`GetAccessGroup(getAccessGroupOptions *GetAccessGroupOptions) - Operation response error`, func() {
		getAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
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
		getAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// TODO: Add check for show_federated query parameter
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke GetAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupWithContext(ctx, getAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke GetAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetAccessGroupOptions model
				getAccessGroupOptionsModel := new(iamaccessgroupsv2.GetAccessGroupOptions)
				getAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAccessGroup successfully`, func() {
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
				getAccessGroupOptionsModel.ShowFederated = core.BoolPtr(false)
				getAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccessGroup(updateAccessGroupOptions *UpdateAccessGroupOptions) - Operation response error`, func() {
		updateAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
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
		updateAccessGroupPath := "/v2/groups/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateAccessGroupWithContext(ctx, updateAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "description": "Description", "account_id": "AccountID", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href", "is_federated": false}`)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdateAccessGroupOptions model
				updateAccessGroupOptionsModel := new(iamaccessgroupsv2.UpdateAccessGroupOptions)
				updateAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.IfMatch = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateAccessGroup successfully`, func() {
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
				updateAccessGroupOptionsModel.Name = core.StringPtr("Awesome Managers")
				updateAccessGroupOptionsModel.Description = core.StringPtr("Group for awesome managers.")
				updateAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteAccessGroup(deleteAccessGroupOptions *DeleteAccessGroupOptions)`, func() {
		deleteAccessGroupPath := "/v2/groups/testString"
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

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamAccessGroupsService.DeleteAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteAccessGroupOptions model
				deleteAccessGroupOptionsModel := new(iamaccessgroupsv2.DeleteAccessGroupOptions)
				deleteAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(false)
				deleteAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
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
				deleteAccessGroupOptionsModel.Force = core.BoolPtr(false)
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
	Describe(`IsMemberOfAccessGroup(isMemberOfAccessGroupOptions *IsMemberOfAccessGroupOptions)`, func() {
		isMemberOfAccessGroupPath := "/v2/groups/testString/members/testString"
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
		addMembersToAccessGroupPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

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
		addMembersToAccessGroupPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddMembersToAccessGroupWithContext(ctx, addMembersToAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"members": [{"iam_id": "IamID", "type": "Type", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

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
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(207)
				}))
			})
			It(`Invoke AddMembersToAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")

				// Construct an instance of the AddMembersToAccessGroupOptions model
				addMembersToAccessGroupOptionsModel := new(iamaccessgroupsv2.AddMembersToAccessGroupOptions)
				addMembersToAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Members = []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{*addGroupMembersRequestMembersItemModel}
				addMembersToAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				addMembersToAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupMembers(listAccessGroupMembersOptions *ListAccessGroupMembersOptions) - Operation response error`, func() {
		listAccessGroupMembersPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
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
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
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
		listAccessGroupMembersPath := "/v2/groups/testString/members"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "membership_type": "MembershipType", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupMembersWithContext(ctx, listAccessGroupMembersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["membership_type"]).To(Equal([]string{"static"}))
					Expect(req.URL.Query()["limit"]).To(Equal([]string{fmt.Sprint(int64(10))}))
					Expect(req.URL.Query()["offset"]).To(Equal([]string{fmt.Sprint(int64(38))}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"testString"}))
					// TODO: Add check for verbose query parameter
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"limit": 5, "offset": 6, "total_count": 10, "first": {"href": "Href"}, "previous": {"href": "Href"}, "next": {"href": "Href"}, "last": {"href": "Href"}, "members": [{"iam_id": "IamID", "type": "Type", "membership_type": "MembershipType", "name": "Name", "email": "Email", "description": "Description", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAccessGroupMembersOptions model
				listAccessGroupMembersOptionsModel := new(iamaccessgroupsv2.ListAccessGroupMembersOptions)
				listAccessGroupMembersOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAccessGroupMembers successfully`, func() {
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
				listAccessGroupMembersOptionsModel.MembershipType = core.StringPtr("static")
				listAccessGroupMembersOptionsModel.Limit = core.Int64Ptr(int64(10))
				listAccessGroupMembersOptionsModel.Offset = core.Int64Ptr(int64(38))
				listAccessGroupMembersOptionsModel.Type = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Verbose = core.BoolPtr(false)
				listAccessGroupMembersOptionsModel.Sort = core.StringPtr("testString")
				listAccessGroupMembersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupMembers(listAccessGroupMembersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Test pagination helper method on response`, func() {
			It(`Invoke GetNextOffset successfully`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=135")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(Equal(core.Int64Ptr(int64(135))))
			})
			It(`Invoke GetNextOffset without a "Next" property in the response`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset without any query params in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).To(BeNil())
				Expect(value).To(BeNil())
			})
			It(`Invoke GetNextOffset with a non-integer query param in the "Next" URL`, func() {
				responseObject := new(iamaccessgroupsv2.GroupMembersList)
				nextObject := new(iamaccessgroupsv2.HrefStruct)
				nextObject.Href = core.StringPtr("ibm.com?offset=tiger")
				responseObject.Next = nextObject
	
				value, err := responseObject.GetNextOffset()
				Expect(err).NotTo(BeNil())
				Expect(value).To(BeNil())
			})
		})
		Context(`Using mock server endpoint - paginated response`, func() {
			BeforeEach(func() {
				var requestNumber int = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupMembersPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					requestNumber++
					if requestNumber == 1 {
						fmt.Fprintf(res, "%s", `{"next":{"href":"https://myhost.com/somePath?offset=1"},"total_count":2,"members":[{"iam_id":"IamID","type":"Type","membership_type":"MembershipType","name":"Name","email":"Email","description":"Description","href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID"}],"limit":1}`)
					} else if requestNumber == 2 {
						fmt.Fprintf(res, "%s", `{"total_count":2,"members":[{"iam_id":"IamID","type":"Type","membership_type":"MembershipType","name":"Name","email":"Email","description":"Description","href":"Href","created_at":"2019-01-01T12:00:00.000Z","created_by_id":"CreatedByID"}],"limit":1}`)
					} else {
						res.WriteHeader(400)
					}
				}))
			})
			It(`Use AccessGroupMembersPager.GetNext successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupMembersOptionsModel := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
					AccessGroupID: core.StringPtr("testString"),
					TransactionID: core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit: core.Int64Ptr(int64(10)),
					Type: core.StringPtr("testString"),
					Verbose: core.BoolPtr(false),
					Sort: core.StringPtr("testString"),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				var allResults []iamaccessgroupsv2.ListGroupMembersResponseMember
				for pager.HasNext() {
					nextPage, err := pager.GetNext()
					Expect(err).To(BeNil())
					Expect(nextPage).ToNot(BeNil())
					allResults = append(allResults, nextPage...)
				}
				Expect(len(allResults)).To(Equal(2))
			})
			It(`Use AccessGroupMembersPager.GetAll successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				listAccessGroupMembersOptionsModel := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
					AccessGroupID: core.StringPtr("testString"),
					TransactionID: core.StringPtr("testString"),
					MembershipType: core.StringPtr("static"),
					Limit: core.Int64Ptr(int64(10)),
					Type: core.StringPtr("testString"),
					Verbose: core.BoolPtr(false),
					Sort: core.StringPtr("testString"),
				}

				pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptionsModel)
				Expect(err).To(BeNil())
				Expect(pager).ToNot(BeNil())

				allResults, err := pager.GetAll()
				Expect(err).To(BeNil())
				Expect(allResults).ToNot(BeNil())
				Expect(len(allResults)).To(Equal(2))
			})
		})
	})
	Describe(`RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions *RemoveMemberFromAccessGroupOptions)`, func() {
		removeMemberFromAccessGroupPath := "/v2/groups/testString/members/testString"
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
		removeMembersFromAccessGroupPath := "/v2/groups/testString/members/delete"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
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
		removeMembersFromAccessGroupPath := "/v2/groups/testString/members/delete"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"access_group_id": "AccessGroupID", "members": [{"iam_id": "IamID", "trace": "Trace", "status_code": 10, "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroupWithContext(ctx, removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(207)
				}))
			})
			It(`Invoke RemoveMembersFromAccessGroup successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RemoveMembersFromAccessGroupOptions model
				removeMembersFromAccessGroupOptionsModel := new(iamaccessgroupsv2.RemoveMembersFromAccessGroupOptions)
				removeMembersFromAccessGroupOptionsModel.AccessGroupID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Members = []string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}
				removeMembersFromAccessGroupOptionsModel.TransactionID = core.StringPtr("testString")
				removeMembersFromAccessGroupOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.RemoveMembersFromAccessGroup(removeMembersFromAccessGroupOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions *RemoveMemberFromAllAccessGroupsOptions) - Operation response error`, func() {
		removeMemberFromAllAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
		removeMemberFromAllAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(removeMemberFromAllAccessGroupsPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RemoveMemberFromAllAccessGroupsOptions model
				removeMemberFromAllAccessGroupsOptionsModel := new(iamaccessgroupsv2.RemoveMemberFromAllAccessGroupsOptions)
				removeMemberFromAllAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				removeMemberFromAllAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.RemoveMemberFromAllAccessGroupsWithContext(ctx, removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(207)
				}))
			})
			It(`Invoke RemoveMemberFromAllAccessGroups successfully`, func() {
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

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions *AddMemberToMultipleAccessGroupsOptions) - Operation response error`, func() {
		addMemberToMultipleAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"access-group-id-1"}
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
		addMemberToMultipleAccessGroupsPath := "/v2/groups/_allgroups/members/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(207)
					fmt.Fprintf(res, "%s", `{"iam_id": "IamID", "groups": [{"access_group_id": "AccessGroupID", "status_code": 10, "trace": "Trace", "errors": [{"code": "Code", "message": "Message"}]}]}`)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"access-group-id-1"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroupsWithContext(ctx, addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				addMemberToMultipleAccessGroupsOptionsModel := new(iamaccessgroupsv2.AddMemberToMultipleAccessGroupsOptions)
				addMemberToMultipleAccessGroupsOptionsModel.AccountID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.IamID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"access-group-id-1"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"access-group-id-1"}
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(207)
				}))
			})
			It(`Invoke AddMemberToMultipleAccessGroups successfully`, func() {
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
				addMemberToMultipleAccessGroupsOptionsModel.Type = core.StringPtr("user")
				addMemberToMultipleAccessGroupsOptionsModel.Groups = []string{"access-group-id-1"}
				addMemberToMultipleAccessGroupsOptionsModel.TransactionID = core.StringPtr("testString")
				addMemberToMultipleAccessGroupsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddAccessGroupRule(addAccessGroupRuleOptions *AddAccessGroupRuleOptions) - Operation response error`, func() {
		addAccessGroupRulePath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
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
		addAccessGroupRulePath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.AddAccessGroupRuleWithContext(ctx, addAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the AddAccessGroupRuleOptions model
				addAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.AddAccessGroupRuleOptions)
				addAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				addAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				addAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				addAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				addAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListAccessGroupRules(listAccessGroupRulesOptions *ListAccessGroupRulesOptions) - Operation response error`, func() {
		listAccessGroupRulesPath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
		listAccessGroupRulesPath := "/v2/groups/testString/rules"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the ListAccessGroupRulesOptions model
				listAccessGroupRulesOptionsModel := new(iamaccessgroupsv2.ListAccessGroupRulesOptions)
				listAccessGroupRulesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.TransactionID = core.StringPtr("testString")
				listAccessGroupRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ListAccessGroupRulesWithContext(ctx, listAccessGroupRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAccessGroupRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"rules": [{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListAccessGroupRules successfully`, func() {
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

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetAccessGroupRule(getAccessGroupRuleOptions *GetAccessGroupRuleOptions) - Operation response error`, func() {
		getAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
		getAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccessGroupRuleOptions model
				getAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.GetAccessGroupRuleOptions)
				getAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				getAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccessGroupRuleWithContext(ctx, getAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccessGroupRulePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAccessGroupRule successfully`, func() {
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

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceAccessGroupRule(replaceAccessGroupRuleOptions *ReplaceAccessGroupRuleOptions) - Operation response error`, func() {
		replaceAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
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
		replaceAccessGroupRulePath := "/v2/groups/testString/rules/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.ReplaceAccessGroupRuleWithContext(ctx, replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "expiration": 10, "realm_name": "RealmName", "access_group_id": "AccessGroupID", "account_id": "AccountID", "conditions": [{"claim": "Claim", "operator": "EQUALS", "value": "Value"}], "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

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
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceAccessGroupRule successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				replaceAccessGroupRuleOptionsModel := new(iamaccessgroupsv2.ReplaceAccessGroupRuleOptions)
				replaceAccessGroupRuleOptionsModel.AccessGroupID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.RuleID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Expiration = core.Int64Ptr(int64(12))
				replaceAccessGroupRuleOptionsModel.RealmName = core.StringPtr("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.Conditions = []iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}
				replaceAccessGroupRuleOptionsModel.Name = core.StringPtr("Manager group rule")
				replaceAccessGroupRuleOptionsModel.TransactionID = core.StringPtr("testString")
				replaceAccessGroupRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`RemoveAccessGroupRule(removeAccessGroupRuleOptions *RemoveAccessGroupRuleOptions)`, func() {
		removeAccessGroupRulePath := "/v2/groups/testString/rules/testString"
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
	Describe(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) - Operation response error`, func() {
		getAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
		getAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getAccountSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the GetAccountSettingsOptions model
				getAccountSettingsOptionsModel := new(iamaccessgroupsv2.GetAccountSettingsOptions)
				getAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				getAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.GetAccountSettingsWithContext(ctx, getAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetAccountSettings successfully`, func() {
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

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateAccountSettings(updateAccountSettingsOptions *UpdateAccountSettingsOptions) - Operation response error`, func() {
		updateAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
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
					fmt.Fprint(res, `} this is not valid json {`)
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
		updateAccountSettingsPath := "/v2/groups/settings"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
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
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully with retries`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())
				iamAccessGroupsService.EnableRetries(0, 0)

				// Construct an instance of the UpdateAccountSettingsOptions model
				updateAccountSettingsOptionsModel := new(iamaccessgroupsv2.UpdateAccountSettingsOptions)
				updateAccountSettingsOptionsModel.AccountID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.PublicAccessEnabled = core.BoolPtr(true)
				updateAccountSettingsOptionsModel.TransactionID = core.StringPtr("testString")
				updateAccountSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamAccessGroupsService.DisableRetries()
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamAccessGroupsService.UpdateAccountSettingsWithContext(ctx, updateAccountSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"account_id": "AccountID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "public_access_enabled": false}`)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
				iamAccessGroupsService, serviceErr := iamaccessgroupsv2.NewIamAccessGroupsV2(&iamaccessgroupsv2.IamAccessGroupsV2Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamAccessGroupsService).ToNot(BeNil())

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
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateAccountSettings successfully`, func() {
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

				// Invoke operation
				result, response, operationErr := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
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
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("isManager")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("EQUALS")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("true")))

				// Construct an instance of the AddAccessGroupRuleOptions model
				accessGroupID := "testString"
				addAccessGroupRuleOptionsExpiration := int64(12)
				addAccessGroupRuleOptionsRealmName := "https://idp.example.org/SAML2"
				addAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				addAccessGroupRuleOptionsModel := iamAccessGroupsService.NewAddAccessGroupRuleOptions(accessGroupID, addAccessGroupRuleOptionsExpiration, addAccessGroupRuleOptionsRealmName, addAccessGroupRuleOptionsConditions)
				addAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				addAccessGroupRuleOptionsModel.SetExpiration(int64(12))
				addAccessGroupRuleOptionsModel.SetRealmName("https://idp.example.org/SAML2")
				addAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				addAccessGroupRuleOptionsModel.SetName("Manager group rule")
				addAccessGroupRuleOptionsModel.SetTransactionID("testString")
				addAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(addAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(addAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(addAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(addAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(addAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddGroupMembersRequestMembersItem successfully`, func() {
				iamID := "testString"
				typeVar := "testString"
				_model, err := iamAccessGroupsService.NewAddGroupMembersRequestMembersItem(iamID, typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewAddMemberToMultipleAccessGroupsOptions successfully`, func() {
				// Construct an instance of the AddMemberToMultipleAccessGroupsOptions model
				accountID := "testString"
				iamID := "testString"
				addMemberToMultipleAccessGroupsOptionsModel := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(accountID, iamID)
				addMemberToMultipleAccessGroupsOptionsModel.SetAccountID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetIamID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetType("user")
				addMemberToMultipleAccessGroupsOptionsModel.SetGroups([]string{"access-group-id-1"})
				addMemberToMultipleAccessGroupsOptionsModel.SetTransactionID("testString")
				addMemberToMultipleAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addMemberToMultipleAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(addMemberToMultipleAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Type).To(Equal(core.StringPtr("user")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Groups).To(Equal([]string{"access-group-id-1"}))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(addMemberToMultipleAccessGroupsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddMembersToAccessGroupOptions successfully`, func() {
				// Construct an instance of the AddGroupMembersRequestMembersItem model
				addGroupMembersRequestMembersItemModel := new(iamaccessgroupsv2.AddGroupMembersRequestMembersItem)
				Expect(addGroupMembersRequestMembersItemModel).ToNot(BeNil())
				addGroupMembersRequestMembersItemModel.IamID = core.StringPtr("IBMid-user1")
				addGroupMembersRequestMembersItemModel.Type = core.StringPtr("user")
				Expect(addGroupMembersRequestMembersItemModel.IamID).To(Equal(core.StringPtr("IBMid-user1")))
				Expect(addGroupMembersRequestMembersItemModel.Type).To(Equal(core.StringPtr("user")))

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
				createAccessGroupOptionsName := "Managers"
				createAccessGroupOptionsModel := iamAccessGroupsService.NewCreateAccessGroupOptions(accountID, createAccessGroupOptionsName)
				createAccessGroupOptionsModel.SetAccountID("testString")
				createAccessGroupOptionsModel.SetName("Managers")
				createAccessGroupOptionsModel.SetDescription("Group for managers")
				createAccessGroupOptionsModel.SetTransactionID("testString")
				createAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createAccessGroupOptionsModel).ToNot(BeNil())
				Expect(createAccessGroupOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("Managers")))
				Expect(createAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("Group for managers")))
				Expect(createAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteAccessGroupOptions successfully`, func() {
				// Construct an instance of the DeleteAccessGroupOptions model
				accessGroupID := "testString"
				deleteAccessGroupOptionsModel := iamAccessGroupsService.NewDeleteAccessGroupOptions(accessGroupID)
				deleteAccessGroupOptionsModel.SetAccessGroupID("testString")
				deleteAccessGroupOptionsModel.SetTransactionID("testString")
				deleteAccessGroupOptionsModel.SetForce(false)
				deleteAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteAccessGroupOptionsModel).ToNot(BeNil())
				Expect(deleteAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteAccessGroupOptionsModel.Force).To(Equal(core.BoolPtr(false)))
				Expect(deleteAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetAccessGroupOptions successfully`, func() {
				// Construct an instance of the GetAccessGroupOptions model
				accessGroupID := "testString"
				getAccessGroupOptionsModel := iamAccessGroupsService.NewGetAccessGroupOptions(accessGroupID)
				getAccessGroupOptionsModel.SetAccessGroupID("testString")
				getAccessGroupOptionsModel.SetTransactionID("testString")
				getAccessGroupOptionsModel.SetShowFederated(false)
				getAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getAccessGroupOptionsModel).ToNot(BeNil())
				Expect(getAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getAccessGroupOptionsModel.ShowFederated).To(Equal(core.BoolPtr(false)))
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
				listAccessGroupMembersOptionsModel.SetMembershipType("static")
				listAccessGroupMembersOptionsModel.SetLimit(int64(10))
				listAccessGroupMembersOptionsModel.SetOffset(int64(38))
				listAccessGroupMembersOptionsModel.SetType("testString")
				listAccessGroupMembersOptionsModel.SetVerbose(false)
				listAccessGroupMembersOptionsModel.SetSort("testString")
				listAccessGroupMembersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupMembersOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupMembersOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.MembershipType).To(Equal(core.StringPtr("static")))
				Expect(listAccessGroupMembersOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccessGroupMembersOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccessGroupMembersOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupMembersOptionsModel.Verbose).To(Equal(core.BoolPtr(false)))
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
				listAccessGroupsOptionsModel.SetMembershipType("static")
				listAccessGroupsOptionsModel.SetLimit(int64(10))
				listAccessGroupsOptionsModel.SetOffset(int64(38))
				listAccessGroupsOptionsModel.SetSort("name")
				listAccessGroupsOptionsModel.SetShowFederated(false)
				listAccessGroupsOptionsModel.SetHidePublicAccess(false)
				listAccessGroupsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAccessGroupsOptionsModel).ToNot(BeNil())
				Expect(listAccessGroupsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listAccessGroupsOptionsModel.MembershipType).To(Equal(core.StringPtr("static")))
				Expect(listAccessGroupsOptionsModel.Limit).To(Equal(core.Int64Ptr(int64(10))))
				Expect(listAccessGroupsOptionsModel.Offset).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAccessGroupsOptionsModel.Sort).To(Equal(core.StringPtr("name")))
				Expect(listAccessGroupsOptionsModel.ShowFederated).To(Equal(core.BoolPtr(false)))
				Expect(listAccessGroupsOptionsModel.HidePublicAccess).To(Equal(core.BoolPtr(false)))
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
				removeMembersFromAccessGroupOptionsModel.SetMembers([]string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"})
				removeMembersFromAccessGroupOptionsModel.SetTransactionID("testString")
				removeMembersFromAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(removeMembersFromAccessGroupOptionsModel).ToNot(BeNil())
				Expect(removeMembersFromAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Members).To(Equal([]string{"IBMId-user1", "iam-ServiceId-123", "iam-Profile-123"}))
				Expect(removeMembersFromAccessGroupOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(removeMembersFromAccessGroupOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceAccessGroupRuleOptions successfully`, func() {
				// Construct an instance of the RuleConditions model
				ruleConditionsModel := new(iamaccessgroupsv2.RuleConditions)
				Expect(ruleConditionsModel).ToNot(BeNil())
				ruleConditionsModel.Claim = core.StringPtr("isManager")
				ruleConditionsModel.Operator = core.StringPtr("EQUALS")
				ruleConditionsModel.Value = core.StringPtr("true")
				Expect(ruleConditionsModel.Claim).To(Equal(core.StringPtr("isManager")))
				Expect(ruleConditionsModel.Operator).To(Equal(core.StringPtr("EQUALS")))
				Expect(ruleConditionsModel.Value).To(Equal(core.StringPtr("true")))

				// Construct an instance of the ReplaceAccessGroupRuleOptions model
				accessGroupID := "testString"
				ruleID := "testString"
				ifMatch := "testString"
				replaceAccessGroupRuleOptionsExpiration := int64(12)
				replaceAccessGroupRuleOptionsRealmName := "https://idp.example.org/SAML2"
				replaceAccessGroupRuleOptionsConditions := []iamaccessgroupsv2.RuleConditions{}
				replaceAccessGroupRuleOptionsModel := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(accessGroupID, ruleID, ifMatch, replaceAccessGroupRuleOptionsExpiration, replaceAccessGroupRuleOptionsRealmName, replaceAccessGroupRuleOptionsConditions)
				replaceAccessGroupRuleOptionsModel.SetAccessGroupID("testString")
				replaceAccessGroupRuleOptionsModel.SetRuleID("testString")
				replaceAccessGroupRuleOptionsModel.SetIfMatch("testString")
				replaceAccessGroupRuleOptionsModel.SetExpiration(int64(12))
				replaceAccessGroupRuleOptionsModel.SetRealmName("https://idp.example.org/SAML2")
				replaceAccessGroupRuleOptionsModel.SetConditions([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel})
				replaceAccessGroupRuleOptionsModel.SetName("Manager group rule")
				replaceAccessGroupRuleOptionsModel.SetTransactionID("testString")
				replaceAccessGroupRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceAccessGroupRuleOptionsModel).ToNot(BeNil())
				Expect(replaceAccessGroupRuleOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.RuleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Expiration).To(Equal(core.Int64Ptr(int64(12))))
				Expect(replaceAccessGroupRuleOptionsModel.RealmName).To(Equal(core.StringPtr("https://idp.example.org/SAML2")))
				Expect(replaceAccessGroupRuleOptionsModel.Conditions).To(Equal([]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel}))
				Expect(replaceAccessGroupRuleOptionsModel.Name).To(Equal(core.StringPtr("Manager group rule")))
				Expect(replaceAccessGroupRuleOptionsModel.TransactionID).To(Equal(core.StringPtr("testString")))
				Expect(replaceAccessGroupRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewRuleConditions successfully`, func() {
				claim := "testString"
				operator := "EQUALS"
				value := "testString"
				_model, err := iamAccessGroupsService.NewRuleConditions(claim, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateAccessGroupOptions successfully`, func() {
				// Construct an instance of the UpdateAccessGroupOptions model
				accessGroupID := "testString"
				ifMatch := "testString"
				updateAccessGroupOptionsModel := iamAccessGroupsService.NewUpdateAccessGroupOptions(accessGroupID, ifMatch)
				updateAccessGroupOptionsModel.SetAccessGroupID("testString")
				updateAccessGroupOptionsModel.SetIfMatch("testString")
				updateAccessGroupOptionsModel.SetName("Awesome Managers")
				updateAccessGroupOptionsModel.SetDescription("Group for awesome managers.")
				updateAccessGroupOptionsModel.SetTransactionID("testString")
				updateAccessGroupOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateAccessGroupOptionsModel).ToNot(BeNil())
				Expect(updateAccessGroupOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateAccessGroupOptionsModel.Name).To(Equal(core.StringPtr("Awesome Managers")))
				Expect(updateAccessGroupOptionsModel.Description).To(Equal(core.StringPtr("Group for awesome managers.")))
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
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
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
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
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
