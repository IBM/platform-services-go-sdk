/**
 * (C) Copyright IBM Corp. 2023.
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
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`IamPolicyManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(iamPolicyManagementService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(iamPolicyManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "https://iampolicymanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(iamPolicyManagementService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL": "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
				})
				Expect(iamPolicyManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := iamPolicyManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamPolicyManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamPolicyManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamPolicyManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(iamPolicyManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamPolicyManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamPolicyManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamPolicyManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamPolicyManagementService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
				})
				err := iamPolicyManagementService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := iamPolicyManagementService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != iamPolicyManagementService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(iamPolicyManagementService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(iamPolicyManagementService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL": "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamPolicyManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(iamPolicyManagementService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = iampolicymanagementv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListPolicies(listPoliciesOptions *ListPoliciesOptions) - Operation response error`, func() {
		listPoliciesPath := "/v1/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_value"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"id"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPolicies with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("access")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listPoliciesOptionsModel.TagName = core.StringPtr("testString")
				listPoliciesOptionsModel.TagValue = core.StringPtr("testString")
				listPoliciesOptionsModel.Sort = core.StringPtr("id")
				listPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listPoliciesOptionsModel.State = core.StringPtr("active")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicies(listPoliciesOptions *ListPoliciesOptions)`, func() {
		listPoliciesPath := "/v1/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_value"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"id"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}]}`)
				}))
			})
			It(`Invoke ListPolicies successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("access")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listPoliciesOptionsModel.TagName = core.StringPtr("testString")
				listPoliciesOptionsModel.TagValue = core.StringPtr("testString")
				listPoliciesOptionsModel.Sort = core.StringPtr("id")
				listPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listPoliciesOptionsModel.State = core.StringPtr("active")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListPoliciesWithContext(ctx, listPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListPoliciesWithContext(ctx, listPoliciesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["tag_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["tag_value"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"id"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}]}`)
				}))
			})
			It(`Invoke ListPolicies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("access")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listPoliciesOptionsModel.TagName = core.StringPtr("testString")
				listPoliciesOptionsModel.TagValue = core.StringPtr("testString")
				listPoliciesOptionsModel.Sort = core.StringPtr("id")
				listPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listPoliciesOptionsModel.State = core.StringPtr("active")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPolicies with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("access")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listPoliciesOptionsModel.TagName = core.StringPtr("testString")
				listPoliciesOptionsModel.TagValue = core.StringPtr("testString")
				listPoliciesOptionsModel.Sort = core.StringPtr("id")
				listPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listPoliciesOptionsModel.State = core.StringPtr("active")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPoliciesOptions model with no property values
				listPoliciesOptionsModelNew := new(iampolicymanagementv1.ListPoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ListPolicies(listPoliciesOptionsModelNew)
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
			It(`Invoke ListPolicies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("access")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listPoliciesOptionsModel.TagName = core.StringPtr("testString")
				listPoliciesOptionsModel.TagValue = core.StringPtr("testString")
				listPoliciesOptionsModel.Sort = core.StringPtr("id")
				listPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listPoliciesOptionsModel.State = core.StringPtr("active")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListPolicies(listPoliciesOptionsModel)
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
	Describe(`CreatePolicy(createPolicyOptions *CreatePolicyOptions) - Operation response error`, func() {
		createPolicyPath := "/v1/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.Description = core.StringPtr("testString")
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePolicy(createPolicyOptions *CreatePolicyOptions)`, func() {
		createPolicyPath := "/v1/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke CreatePolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.Description = core.StringPtr("testString")
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.CreatePolicyWithContext(ctx, createPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.CreatePolicyWithContext(ctx, createPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke CreatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.CreatePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.Description = core.StringPtr("testString")
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.Description = core.StringPtr("testString")
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePolicyOptions model with no property values
				createPolicyOptionsModelNew := new(iampolicymanagementv1.CreatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicy(createPolicyOptionsModelNew)
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
			It(`Invoke CreatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.Description = core.StringPtr("testString")
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.CreatePolicy(createPolicyOptionsModel)
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
	Describe(`ReplacePolicy(replacePolicyOptions *ReplacePolicyOptions) - Operation response error`, func() {
		replacePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplacePolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the ReplacePolicyOptions model
				replacePolicyOptionsModel := new(iampolicymanagementv1.ReplacePolicyOptions)
				replacePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				replacePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyOptionsModel.Type = core.StringPtr("testString")
				replacePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				replacePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				replacePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				replacePolicyOptionsModel.Description = core.StringPtr("testString")
				replacePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplacePolicy(replacePolicyOptions *ReplacePolicyOptions)`, func() {
		replacePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyPath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke ReplacePolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the ReplacePolicyOptions model
				replacePolicyOptionsModel := new(iampolicymanagementv1.ReplacePolicyOptions)
				replacePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				replacePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyOptionsModel.Type = core.StringPtr("testString")
				replacePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				replacePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				replacePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				replacePolicyOptionsModel.Description = core.StringPtr("testString")
				replacePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ReplacePolicyWithContext(ctx, replacePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ReplacePolicyWithContext(ctx, replacePolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyPath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke ReplacePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ReplacePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the ReplacePolicyOptions model
				replacePolicyOptionsModel := new(iampolicymanagementv1.ReplacePolicyOptions)
				replacePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				replacePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyOptionsModel.Type = core.StringPtr("testString")
				replacePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				replacePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				replacePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				replacePolicyOptionsModel.Description = core.StringPtr("testString")
				replacePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplacePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the ReplacePolicyOptions model
				replacePolicyOptionsModel := new(iampolicymanagementv1.ReplacePolicyOptions)
				replacePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				replacePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyOptionsModel.Type = core.StringPtr("testString")
				replacePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				replacePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				replacePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				replacePolicyOptionsModel.Description = core.StringPtr("testString")
				replacePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplacePolicyOptions model with no property values
				replacePolicyOptionsModelNew := new(iampolicymanagementv1.ReplacePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModelNew)
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
			It(`Invoke ReplacePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}

				// Construct an instance of the ReplacePolicyOptions model
				replacePolicyOptionsModel := new(iampolicymanagementv1.ReplacePolicyOptions)
				replacePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				replacePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyOptionsModel.Type = core.StringPtr("testString")
				replacePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				replacePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				replacePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				replacePolicyOptionsModel.Description = core.StringPtr("testString")
				replacePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ReplacePolicy(replacePolicyOptionsModel)
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
	Describe(`GetPolicy(getPolicyOptions *GetPolicyOptions) - Operation response error`, func() {
		getPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicy(getPolicyOptions *GetPolicyOptions)`, func() {
		getPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}`)
				}))
			})
			It(`Invoke GetPolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetPolicyWithContext(ctx, getPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetPolicyWithContext(ctx, getPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}`)
				}))
			})
			It(`Invoke GetPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPolicyOptions model with no property values
				getPolicyOptionsModelNew := new(iampolicymanagementv1.GetPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetPolicy(getPolicyOptionsModelNew)
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
			It(`Invoke GetPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetPolicy(getPolicyOptionsModel)
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
	Describe(`DeletePolicy(deletePolicyOptions *DeletePolicyOptions)`, func() {
		deletePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.DeletePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePolicyOptions model
				deletePolicyOptionsModel := new(iampolicymanagementv1.DeletePolicyOptions)
				deletePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.DeletePolicy(deletePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the DeletePolicyOptions model
				deletePolicyOptionsModel := new(iampolicymanagementv1.DeletePolicyOptions)
				deletePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.DeletePolicy(deletePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePolicyOptions model with no property values
				deletePolicyOptionsModelNew := new(iampolicymanagementv1.DeletePolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.DeletePolicy(deletePolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePolicyState(updatePolicyStateOptions *UpdatePolicyStateOptions) - Operation response error`, func() {
		updatePolicyStatePath := "/v1/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyStatePath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePolicyState with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdatePolicyStateOptions model
				updatePolicyStateOptionsModel := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				updatePolicyStateOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyStateOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyStateOptionsModel.State = core.StringPtr("active")
				updatePolicyStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePolicyState(updatePolicyStateOptions *UpdatePolicyStateOptions)`, func() {
		updatePolicyStatePath := "/v1/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyStatePath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke UpdatePolicyState successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdatePolicyStateOptions model
				updatePolicyStateOptionsModel := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				updatePolicyStateOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyStateOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyStateOptionsModel.State = core.StringPtr("active")
				updatePolicyStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.UpdatePolicyStateWithContext(ctx, updatePolicyStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.UpdatePolicyStateWithContext(ctx, updatePolicyStateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyStatePath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke UpdatePolicyState successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.UpdatePolicyState(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the UpdatePolicyStateOptions model
				updatePolicyStateOptionsModel := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				updatePolicyStateOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyStateOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyStateOptionsModel.State = core.StringPtr("active")
				updatePolicyStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdatePolicyState with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdatePolicyStateOptions model
				updatePolicyStateOptionsModel := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				updatePolicyStateOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyStateOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyStateOptionsModel.State = core.StringPtr("active")
				updatePolicyStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePolicyStateOptions model with no property values
				updatePolicyStateOptionsModelNew := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModelNew)
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
			It(`Invoke UpdatePolicyState successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdatePolicyStateOptions model
				updatePolicyStateOptionsModel := new(iampolicymanagementv1.UpdatePolicyStateOptions)
				updatePolicyStateOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyStateOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyStateOptionsModel.State = core.StringPtr("active")
				updatePolicyStateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.UpdatePolicyState(updatePolicyStateOptionsModel)
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
	Describe(`ListRoles(listRolesOptions *ListRolesOptions) - Operation response error`, func() {
		listRolesPath := "/v2/roles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["source_service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["policy_type"]).To(Equal([]string{"authorization"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"IAM"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoles with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.SourceServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.PolicyType = core.StringPtr("authorization")
				listRolesOptionsModel.ServiceGroupID = core.StringPtr("IAM")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListRoles(listRolesOptions *ListRolesOptions)`, func() {
		listRolesPath := "/v2/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["source_service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["policy_type"]).To(Equal([]string{"authorization"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"IAM"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"custom_roles": [{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}], "service_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN"}], "system_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN"}]}`)
				}))
			})
			It(`Invoke ListRoles successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.SourceServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.PolicyType = core.StringPtr("authorization")
				listRolesOptionsModel.ServiceGroupID = core.StringPtr("IAM")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListRolesWithContext(ctx, listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListRolesWithContext(ctx, listRolesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listRolesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["source_service_name"]).To(Equal([]string{"iam-groups"}))
					Expect(req.URL.Query()["policy_type"]).To(Equal([]string{"authorization"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"IAM"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"custom_roles": [{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}], "service_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN"}], "system_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN"}]}`)
				}))
			})
			It(`Invoke ListRoles successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListRoles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.SourceServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.PolicyType = core.StringPtr("authorization")
				listRolesOptionsModel.ServiceGroupID = core.StringPtr("IAM")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListRoles with error: Operation request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.SourceServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.PolicyType = core.StringPtr("authorization")
				listRolesOptionsModel.ServiceGroupID = core.StringPtr("IAM")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
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
			It(`Invoke ListRoles successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.SourceServiceName = core.StringPtr("iam-groups")
				listRolesOptionsModel.PolicyType = core.StringPtr("authorization")
				listRolesOptionsModel.ServiceGroupID = core.StringPtr("IAM")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListRoles(listRolesOptionsModel)
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
	Describe(`CreateRole(createRoleOptions *CreateRoleOptions) - Operation response error`, func() {
		createRolePath := "/v2/roles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRole with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("Developer")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("iam-groups")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("default")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRole(createRoleOptions *CreateRoleOptions)`, func() {
		createRolePath := "/v2/roles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateRole successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("Developer")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("iam-groups")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("default")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.CreateRoleWithContext(ctx, createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.CreateRoleWithContext(ctx, createRoleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createRolePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.CreateRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("Developer")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("iam-groups")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("default")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateRole with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("Developer")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("iam-groups")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("default")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRoleOptions model with no property values
				createRoleOptionsModelNew := new(iampolicymanagementv1.CreateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.CreateRole(createRoleOptionsModelNew)
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
			It(`Invoke CreateRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("Developer")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("iam-groups")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("default")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.CreateRole(createRoleOptionsModel)
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
	Describe(`ReplaceRole(replaceRoleOptions *ReplaceRoleOptions) - Operation response error`, func() {
		replaceRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRolePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceRole with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceRoleOptions model
				replaceRoleOptionsModel := new(iampolicymanagementv1.ReplaceRoleOptions)
				replaceRoleOptionsModel.RoleID = core.StringPtr("testString")
				replaceRoleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRoleOptionsModel.DisplayName = core.StringPtr("testString")
				replaceRoleOptionsModel.Actions = []string{"testString"}
				replaceRoleOptionsModel.Description = core.StringPtr("testString")
				replaceRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceRole(replaceRoleOptions *ReplaceRoleOptions)`, func() {
		replaceRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceRolePath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke ReplaceRole successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceRoleOptions model
				replaceRoleOptionsModel := new(iampolicymanagementv1.ReplaceRoleOptions)
				replaceRoleOptionsModel.RoleID = core.StringPtr("testString")
				replaceRoleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRoleOptionsModel.DisplayName = core.StringPtr("testString")
				replaceRoleOptionsModel.Actions = []string{"testString"}
				replaceRoleOptionsModel.Description = core.StringPtr("testString")
				replaceRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ReplaceRoleWithContext(ctx, replaceRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ReplaceRoleWithContext(ctx, replaceRoleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceRolePath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke ReplaceRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ReplaceRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceRoleOptions model
				replaceRoleOptionsModel := new(iampolicymanagementv1.ReplaceRoleOptions)
				replaceRoleOptionsModel.RoleID = core.StringPtr("testString")
				replaceRoleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRoleOptionsModel.DisplayName = core.StringPtr("testString")
				replaceRoleOptionsModel.Actions = []string{"testString"}
				replaceRoleOptionsModel.Description = core.StringPtr("testString")
				replaceRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceRole with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceRoleOptions model
				replaceRoleOptionsModel := new(iampolicymanagementv1.ReplaceRoleOptions)
				replaceRoleOptionsModel.RoleID = core.StringPtr("testString")
				replaceRoleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRoleOptionsModel.DisplayName = core.StringPtr("testString")
				replaceRoleOptionsModel.Actions = []string{"testString"}
				replaceRoleOptionsModel.Description = core.StringPtr("testString")
				replaceRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceRoleOptions model with no property values
				replaceRoleOptionsModelNew := new(iampolicymanagementv1.ReplaceRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModelNew)
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
			It(`Invoke ReplaceRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ReplaceRoleOptions model
				replaceRoleOptionsModel := new(iampolicymanagementv1.ReplaceRoleOptions)
				replaceRoleOptionsModel.RoleID = core.StringPtr("testString")
				replaceRoleOptionsModel.IfMatch = core.StringPtr("testString")
				replaceRoleOptionsModel.DisplayName = core.StringPtr("testString")
				replaceRoleOptionsModel.Actions = []string{"testString"}
				replaceRoleOptionsModel.Description = core.StringPtr("testString")
				replaceRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ReplaceRole(replaceRoleOptionsModel)
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
	Describe(`GetRole(getRoleOptions *GetRoleOptions) - Operation response error`, func() {
		getRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRole with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRole(getRoleOptions *GetRoleOptions)`, func() {
		getRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke GetRole successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetRoleWithContext(ctx, getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetRoleWithContext(ctx, getRoleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "CRN", "name": "Developer", "account_id": "AccountID", "service_name": "iam-groups", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke GetRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetRole with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRoleOptions model with no property values
				getRoleOptionsModelNew := new(iampolicymanagementv1.GetRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetRole(getRoleOptionsModelNew)
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
			It(`Invoke GetRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetRole(getRoleOptionsModel)
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
	Describe(`DeleteRole(deleteRoleOptions *DeleteRoleOptions)`, func() {
		deleteRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteRolePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.DeleteRole(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(iampolicymanagementv1.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRole with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(iampolicymanagementv1.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRoleOptions model with no property values
				deleteRoleOptionsModelNew := new(iampolicymanagementv1.DeleteRoleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.DeleteRole(deleteRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListV2Policies(listV2PoliciesOptions *ListV2PoliciesOptions) - Operation response error`, func() {
		listV2PoliciesPath := "/v2/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listV2PoliciesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListV2Policies with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListV2PoliciesOptions model
				listV2PoliciesOptionsModel := new(iampolicymanagementv1.ListV2PoliciesOptions)
				listV2PoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listV2PoliciesOptionsModel.IamID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Type = core.StringPtr("access")
				listV2PoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listV2PoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				listV2PoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Sort = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listV2PoliciesOptionsModel.State = core.StringPtr("active")
				listV2PoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListV2Policies(listV2PoliciesOptions *ListV2PoliciesOptions)`, func() {
		listV2PoliciesPath := "/v2/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listV2PoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19, "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}]}`)
				}))
			})
			It(`Invoke ListV2Policies successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListV2PoliciesOptions model
				listV2PoliciesOptionsModel := new(iampolicymanagementv1.ListV2PoliciesOptions)
				listV2PoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listV2PoliciesOptionsModel.IamID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Type = core.StringPtr("access")
				listV2PoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listV2PoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				listV2PoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Sort = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listV2PoliciesOptionsModel.State = core.StringPtr("active")
				listV2PoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListV2PoliciesWithContext(ctx, listV2PoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListV2PoliciesWithContext(ctx, listV2PoliciesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listV2PoliciesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["iam_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["access_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["type"]).To(Equal([]string{"access"}))
					Expect(req.URL.Query()["service_type"]).To(Equal([]string{"service"}))
					Expect(req.URL.Query()["service_name"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["service_group_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["sort"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19, "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}]}`)
				}))
			})
			It(`Invoke ListV2Policies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListV2Policies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListV2PoliciesOptions model
				listV2PoliciesOptionsModel := new(iampolicymanagementv1.ListV2PoliciesOptions)
				listV2PoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listV2PoliciesOptionsModel.IamID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Type = core.StringPtr("access")
				listV2PoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listV2PoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				listV2PoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Sort = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listV2PoliciesOptionsModel.State = core.StringPtr("active")
				listV2PoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListV2Policies with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListV2PoliciesOptions model
				listV2PoliciesOptionsModel := new(iampolicymanagementv1.ListV2PoliciesOptions)
				listV2PoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listV2PoliciesOptionsModel.IamID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Type = core.StringPtr("access")
				listV2PoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listV2PoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				listV2PoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Sort = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listV2PoliciesOptionsModel.State = core.StringPtr("active")
				listV2PoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListV2PoliciesOptions model with no property values
				listV2PoliciesOptionsModelNew := new(iampolicymanagementv1.ListV2PoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModelNew)
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
			It(`Invoke ListV2Policies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListV2PoliciesOptions model
				listV2PoliciesOptionsModel := new(iampolicymanagementv1.ListV2PoliciesOptions)
				listV2PoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listV2PoliciesOptionsModel.IamID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Type = core.StringPtr("access")
				listV2PoliciesOptionsModel.ServiceType = core.StringPtr("service")
				listV2PoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				listV2PoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Sort = core.StringPtr("testString")
				listV2PoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				listV2PoliciesOptionsModel.State = core.StringPtr("active")
				listV2PoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListV2Policies(listV2PoliciesOptionsModel)
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
	Describe(`CreateV2Policy(createV2PolicyOptions *CreateV2PolicyOptions) - Operation response error`, func() {
		createV2PolicyPath := "/v2/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createV2PolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateV2Policy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateV2PolicyOptions model
				createV2PolicyOptionsModel := new(iampolicymanagementv1.CreateV2PolicyOptions)
				createV2PolicyOptionsModel.Control = controlModel
				createV2PolicyOptionsModel.Type = core.StringPtr("access")
				createV2PolicyOptionsModel.Description = core.StringPtr("testString")
				createV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				createV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				createV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				createV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				createV2PolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateV2Policy(createV2PolicyOptions *CreateV2PolicyOptions)`, func() {
		createV2PolicyPath := "/v2/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createV2PolicyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19}`)
				}))
			})
			It(`Invoke CreateV2Policy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateV2PolicyOptions model
				createV2PolicyOptionsModel := new(iampolicymanagementv1.CreateV2PolicyOptions)
				createV2PolicyOptionsModel.Control = controlModel
				createV2PolicyOptionsModel.Type = core.StringPtr("access")
				createV2PolicyOptionsModel.Description = core.StringPtr("testString")
				createV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				createV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				createV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				createV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				createV2PolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.CreateV2PolicyWithContext(ctx, createV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.CreateV2PolicyWithContext(ctx, createV2PolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createV2PolicyPath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19}`)
				}))
			})
			It(`Invoke CreateV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.CreateV2Policy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateV2PolicyOptions model
				createV2PolicyOptionsModel := new(iampolicymanagementv1.CreateV2PolicyOptions)
				createV2PolicyOptionsModel.Control = controlModel
				createV2PolicyOptionsModel.Type = core.StringPtr("access")
				createV2PolicyOptionsModel.Description = core.StringPtr("testString")
				createV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				createV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				createV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				createV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				createV2PolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateV2Policy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateV2PolicyOptions model
				createV2PolicyOptionsModel := new(iampolicymanagementv1.CreateV2PolicyOptions)
				createV2PolicyOptionsModel.Control = controlModel
				createV2PolicyOptionsModel.Type = core.StringPtr("access")
				createV2PolicyOptionsModel.Description = core.StringPtr("testString")
				createV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				createV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				createV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				createV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				createV2PolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateV2PolicyOptions model with no property values
				createV2PolicyOptionsModelNew := new(iampolicymanagementv1.CreateV2PolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModelNew)
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
			It(`Invoke CreateV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the CreateV2PolicyOptions model
				createV2PolicyOptionsModel := new(iampolicymanagementv1.CreateV2PolicyOptions)
				createV2PolicyOptionsModel.Control = controlModel
				createV2PolicyOptionsModel.Type = core.StringPtr("access")
				createV2PolicyOptionsModel.Description = core.StringPtr("testString")
				createV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				createV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				createV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				createV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				createV2PolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				createV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.CreateV2Policy(createV2PolicyOptionsModel)
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
	Describe(`ReplaceV2Policy(replaceV2PolicyOptions *ReplaceV2PolicyOptions) - Operation response error`, func() {
		replaceV2PolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceV2PolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceV2Policy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceV2PolicyOptions model
				replaceV2PolicyOptionsModel := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				replaceV2PolicyOptionsModel.ID = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Control = controlModel
				replaceV2PolicyOptionsModel.Type = core.StringPtr("access")
				replaceV2PolicyOptionsModel.Description = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				replaceV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				replaceV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				replaceV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceV2Policy(replaceV2PolicyOptions *ReplaceV2PolicyOptions)`, func() {
		replaceV2PolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceV2PolicyPath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19}`)
				}))
			})
			It(`Invoke ReplaceV2Policy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceV2PolicyOptions model
				replaceV2PolicyOptionsModel := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				replaceV2PolicyOptionsModel.ID = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Control = controlModel
				replaceV2PolicyOptionsModel.Type = core.StringPtr("access")
				replaceV2PolicyOptionsModel.Description = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				replaceV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				replaceV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				replaceV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ReplaceV2PolicyWithContext(ctx, replaceV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ReplaceV2PolicyWithContext(ctx, replaceV2PolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replaceV2PolicyPath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19}`)
				}))
			})
			It(`Invoke ReplaceV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ReplaceV2Policy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceV2PolicyOptions model
				replaceV2PolicyOptionsModel := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				replaceV2PolicyOptionsModel.ID = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Control = controlModel
				replaceV2PolicyOptionsModel.Type = core.StringPtr("access")
				replaceV2PolicyOptionsModel.Description = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				replaceV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				replaceV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				replaceV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceV2Policy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceV2PolicyOptions model
				replaceV2PolicyOptionsModel := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				replaceV2PolicyOptionsModel.ID = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Control = controlModel
				replaceV2PolicyOptionsModel.Type = core.StringPtr("access")
				replaceV2PolicyOptionsModel.Description = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				replaceV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				replaceV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				replaceV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceV2PolicyOptions model with no property values
				replaceV2PolicyOptionsModelNew := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModelNew)
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
			It(`Invoke ReplaceV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the ReplaceV2PolicyOptions model
				replaceV2PolicyOptionsModel := new(iampolicymanagementv1.ReplaceV2PolicyOptions)
				replaceV2PolicyOptionsModel.ID = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.IfMatch = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Control = controlModel
				replaceV2PolicyOptionsModel.Type = core.StringPtr("access")
				replaceV2PolicyOptionsModel.Description = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Subject = v2PolicySubjectModel
				replaceV2PolicyOptionsModel.Resource = v2PolicyResourceModel
				replaceV2PolicyOptionsModel.Pattern = core.StringPtr("testString")
				replaceV2PolicyOptionsModel.Rule = v2PolicyRuleModel
				replaceV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ReplaceV2Policy(replaceV2PolicyOptionsModel)
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
	Describe(`GetV2Policy(getV2PolicyOptions *GetV2PolicyOptions) - Operation response error`, func() {
		getV2PolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getV2PolicyPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetV2Policy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetV2PolicyOptions model
				getV2PolicyOptionsModel := new(iampolicymanagementv1.GetV2PolicyOptions)
				getV2PolicyOptionsModel.ID = core.StringPtr("testString")
				getV2PolicyOptionsModel.Format = core.StringPtr("include_last_permit")
				getV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetV2Policy(getV2PolicyOptions *GetV2PolicyOptions)`, func() {
		getV2PolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getV2PolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19, "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}`)
				}))
			})
			It(`Invoke GetV2Policy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetV2PolicyOptions model
				getV2PolicyOptionsModel := new(iampolicymanagementv1.GetV2PolicyOptions)
				getV2PolicyOptionsModel.ID = core.StringPtr("testString")
				getV2PolicyOptionsModel.Format = core.StringPtr("include_last_permit")
				getV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetV2PolicyWithContext(ctx, getV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetV2PolicyWithContext(ctx, getV2PolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getV2PolicyPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"type": "access", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "Value"}]}, "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "id": "ID", "href": "Href", "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}, "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active", "last_permit_at": "LastPermitAt", "last_permit_frequency": 19, "template": {"id": "ID", "version": "Version", "assignment_id": "AssignmentID", "root_id": "RootID", "root_version": "RootVersion"}}`)
				}))
			})
			It(`Invoke GetV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetV2Policy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetV2PolicyOptions model
				getV2PolicyOptionsModel := new(iampolicymanagementv1.GetV2PolicyOptions)
				getV2PolicyOptionsModel.ID = core.StringPtr("testString")
				getV2PolicyOptionsModel.Format = core.StringPtr("include_last_permit")
				getV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetV2Policy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetV2PolicyOptions model
				getV2PolicyOptionsModel := new(iampolicymanagementv1.GetV2PolicyOptions)
				getV2PolicyOptionsModel.ID = core.StringPtr("testString")
				getV2PolicyOptionsModel.Format = core.StringPtr("include_last_permit")
				getV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetV2PolicyOptions model with no property values
				getV2PolicyOptionsModelNew := new(iampolicymanagementv1.GetV2PolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModelNew)
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
			It(`Invoke GetV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetV2PolicyOptions model
				getV2PolicyOptionsModel := new(iampolicymanagementv1.GetV2PolicyOptions)
				getV2PolicyOptionsModel.ID = core.StringPtr("testString")
				getV2PolicyOptionsModel.Format = core.StringPtr("include_last_permit")
				getV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetV2Policy(getV2PolicyOptionsModel)
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
	Describe(`DeleteV2Policy(deleteV2PolicyOptions *DeleteV2PolicyOptions)`, func() {
		deleteV2PolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteV2PolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteV2Policy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.DeleteV2Policy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteV2PolicyOptions model
				deleteV2PolicyOptionsModel := new(iampolicymanagementv1.DeleteV2PolicyOptions)
				deleteV2PolicyOptionsModel.ID = core.StringPtr("testString")
				deleteV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.DeleteV2Policy(deleteV2PolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteV2Policy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the DeleteV2PolicyOptions model
				deleteV2PolicyOptionsModel := new(iampolicymanagementv1.DeleteV2PolicyOptions)
				deleteV2PolicyOptionsModel.ID = core.StringPtr("testString")
				deleteV2PolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.DeleteV2Policy(deleteV2PolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteV2PolicyOptions model with no property values
				deleteV2PolicyOptionsModelNew := new(iampolicymanagementv1.DeleteV2PolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.DeleteV2Policy(deleteV2PolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicyTemplates(listPolicyTemplatesOptions *ListPolicyTemplatesOptions) - Operation response error`, func() {
		listPolicyTemplatesPath := "/v1/policy_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPolicyTemplates with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplatesOptions model
				listPolicyTemplatesOptionsModel := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				listPolicyTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyTemplatesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicyTemplates(listPolicyTemplatesOptions *ListPolicyTemplatesOptions)`, func() {
		listPolicyTemplatesPath := "/v1/policy_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policy_templates": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListPolicyTemplates successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListPolicyTemplatesOptions model
				listPolicyTemplatesOptionsModel := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				listPolicyTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyTemplatesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListPolicyTemplatesWithContext(ctx, listPolicyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListPolicyTemplatesWithContext(ctx, listPolicyTemplatesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplatesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policy_templates": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListPolicyTemplates successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplates(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPolicyTemplatesOptions model
				listPolicyTemplatesOptionsModel := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				listPolicyTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyTemplatesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPolicyTemplates with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplatesOptions model
				listPolicyTemplatesOptionsModel := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				listPolicyTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyTemplatesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPolicyTemplatesOptions model with no property values
				listPolicyTemplatesOptionsModelNew := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModelNew)
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
			It(`Invoke ListPolicyTemplates successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplatesOptions model
				listPolicyTemplatesOptionsModel := new(iampolicymanagementv1.ListPolicyTemplatesOptions)
				listPolicyTemplatesOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyTemplatesOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyTemplatesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplates(listPolicyTemplatesOptionsModel)
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
	Describe(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions) - Operation response error`, func() {
		createPolicyTemplatePath := "/v1/policy_templates"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplatePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePolicyTemplate with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				createPolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePolicyTemplate(createPolicyTemplateOptions *CreatePolicyTemplateOptions)`, func() {
		createPolicyTemplatePath := "/v1/policy_templates"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplatePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreatePolicyTemplate successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				createPolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.CreatePolicyTemplateWithContext(ctx, createPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.CreatePolicyTemplateWithContext(ctx, createPolicyTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplatePath))
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

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreatePolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				createPolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePolicyTemplate with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				createPolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePolicyTemplateOptions model with no property values
				createPolicyTemplateOptionsModelNew := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModelNew)
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
			It(`Invoke CreatePolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateOptions)
				createPolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.AccountID = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateOptionsModel.AcceptLanguage = core.StringPtr("default")
				createPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplate(createPolicyTemplateOptionsModel)
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
	Describe(`GetPolicyTemplate(getPolicyTemplateOptions *GetPolicyTemplateOptions) - Operation response error`, func() {
		getPolicyTemplatePath := "/v1/policy_templates/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplatePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicyTemplate with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateOptions model
				getPolicyTemplateOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				getPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyTemplate(getPolicyTemplateOptions *GetPolicyTemplateOptions)`, func() {
		getPolicyTemplatePath := "/v1/policy_templates/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetPolicyTemplate successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetPolicyTemplateOptions model
				getPolicyTemplateOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				getPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetPolicyTemplateWithContext(ctx, getPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetPolicyTemplateWithContext(ctx, getPolicyTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplatePath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetPolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyTemplateOptions model
				getPolicyTemplateOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				getPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPolicyTemplate with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateOptions model
				getPolicyTemplateOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				getPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPolicyTemplateOptions model with no property values
				getPolicyTemplateOptionsModelNew := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModelNew)
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
			It(`Invoke GetPolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateOptions model
				getPolicyTemplateOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateOptions)
				getPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplate(getPolicyTemplateOptionsModel)
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
	Describe(`DeletePolicyTemplate(deletePolicyTemplateOptions *DeletePolicyTemplateOptions)`, func() {
		deletePolicyTemplatePath := "/v1/policy_templates/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePolicyTemplatePath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.DeletePolicyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePolicyTemplateOptions model
				deletePolicyTemplateOptionsModel := new(iampolicymanagementv1.DeletePolicyTemplateOptions)
				deletePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				deletePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.DeletePolicyTemplate(deletePolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePolicyTemplate with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the DeletePolicyTemplateOptions model
				deletePolicyTemplateOptionsModel := new(iampolicymanagementv1.DeletePolicyTemplateOptions)
				deletePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				deletePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.DeletePolicyTemplate(deletePolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePolicyTemplateOptions model with no property values
				deletePolicyTemplateOptionsModelNew := new(iampolicymanagementv1.DeletePolicyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.DeletePolicyTemplate(deletePolicyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions *CreatePolicyTemplateVersionOptions) - Operation response error`, func() {
		createPolicyTemplateVersionPath := "/v1/policy_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplateVersionPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePolicyTemplateVersion with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				createPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				createPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePolicyTemplateVersion(createPolicyTemplateVersionOptions *CreatePolicyTemplateVersionOptions)`, func() {
		createPolicyTemplateVersionPath := "/v1/policy_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplateVersionPath))
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

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreatePolicyTemplateVersion successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				createPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				createPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersionWithContext(ctx, createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.CreatePolicyTemplateVersionWithContext(ctx, createPolicyTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createPolicyTemplateVersionPath))
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

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke CreatePolicyTemplateVersion successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				createPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				createPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreatePolicyTemplateVersion with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				createPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				createPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePolicyTemplateVersionOptions model with no property values
				createPolicyTemplateVersionOptionsModelNew := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModelNew)
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
			It(`Invoke CreatePolicyTemplateVersion successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				createPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.CreatePolicyTemplateVersionOptions)
				createPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Policy = templatePolicyModel
				createPolicyTemplateVersionOptionsModel.Name = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Description = core.StringPtr("testString")
				createPolicyTemplateVersionOptionsModel.Committed = core.BoolPtr(true)
				createPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.CreatePolicyTemplateVersion(createPolicyTemplateVersionOptionsModel)
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
	Describe(`ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions *ListPolicyTemplateVersionsOptions) - Operation response error`, func() {
		listPolicyTemplateVersionsPath := "/v1/policy_templates/testString/versions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPolicyTemplateVersions with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				listPolicyTemplateVersionsOptionsModel := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				listPolicyTemplateVersionsOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				listPolicyTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicyTemplateVersions(listPolicyTemplateVersionsOptions *ListPolicyTemplateVersionsOptions)`, func() {
		listPolicyTemplateVersionsPath := "/v1/policy_templates/testString/versions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"versions": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListPolicyTemplateVersions successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				listPolicyTemplateVersionsOptionsModel := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				listPolicyTemplateVersionsOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				listPolicyTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListPolicyTemplateVersionsWithContext(ctx, listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListPolicyTemplateVersionsWithContext(ctx, listPolicyTemplateVersionsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyTemplateVersionsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"versions": [{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}]}`)
				}))
			})
			It(`Invoke ListPolicyTemplateVersions successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplateVersions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				listPolicyTemplateVersionsOptionsModel := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				listPolicyTemplateVersionsOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				listPolicyTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPolicyTemplateVersions with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				listPolicyTemplateVersionsOptionsModel := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				listPolicyTemplateVersionsOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				listPolicyTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPolicyTemplateVersionsOptions model with no property values
				listPolicyTemplateVersionsOptionsModelNew := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModelNew)
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
			It(`Invoke ListPolicyTemplateVersions successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				listPolicyTemplateVersionsOptionsModel := new(iampolicymanagementv1.ListPolicyTemplateVersionsOptions)
				listPolicyTemplateVersionsOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				listPolicyTemplateVersionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListPolicyTemplateVersions(listPolicyTemplateVersionsOptionsModel)
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
	Describe(`ReplacePolicyTemplate(replacePolicyTemplateOptions *ReplacePolicyTemplateOptions) - Operation response error`, func() {
		replacePolicyTemplatePath := "/v1/policy_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyTemplatePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplacePolicyTemplate with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the ReplacePolicyTemplateOptions model
				replacePolicyTemplateOptionsModel := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				replacePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Policy = templatePolicyModel
				replacePolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				replacePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplacePolicyTemplate(replacePolicyTemplateOptions *ReplacePolicyTemplateOptions)`, func() {
		replacePolicyTemplatePath := "/v1/policy_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyTemplatePath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplacePolicyTemplate successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the ReplacePolicyTemplateOptions model
				replacePolicyTemplateOptionsModel := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				replacePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Policy = templatePolicyModel
				replacePolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				replacePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ReplacePolicyTemplateWithContext(ctx, replacePolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ReplacePolicyTemplateWithContext(ctx, replacePolicyTemplateOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(replacePolicyTemplatePath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke ReplacePolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ReplacePolicyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the ReplacePolicyTemplateOptions model
				replacePolicyTemplateOptionsModel := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				replacePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Policy = templatePolicyModel
				replacePolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				replacePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplacePolicyTemplate with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the ReplacePolicyTemplateOptions model
				replacePolicyTemplateOptionsModel := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				replacePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Policy = templatePolicyModel
				replacePolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				replacePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplacePolicyTemplateOptions model with no property values
				replacePolicyTemplateOptionsModelNew := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModelNew)
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
			It(`Invoke ReplacePolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				rolesModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				controlModel.Grant = grantModel

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel

				// Construct an instance of the ReplacePolicyTemplateOptions model
				replacePolicyTemplateOptionsModel := new(iampolicymanagementv1.ReplacePolicyTemplateOptions)
				replacePolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.IfMatch = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Policy = templatePolicyModel
				replacePolicyTemplateOptionsModel.Name = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Description = core.StringPtr("testString")
				replacePolicyTemplateOptionsModel.Committed = core.BoolPtr(true)
				replacePolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ReplacePolicyTemplate(replacePolicyTemplateOptionsModel)
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
	Describe(`DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptions *DeletePolicyTemplateVersionOptions)`, func() {
		deletePolicyTemplateVersionPath := "/v1/policy_templates/testString/versions/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deletePolicyTemplateVersionPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePolicyTemplateVersion successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.DeletePolicyTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeletePolicyTemplateVersionOptions model
				deletePolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.DeletePolicyTemplateVersionOptions)
				deletePolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				deletePolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deletePolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePolicyTemplateVersion with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the DeletePolicyTemplateVersionOptions model
				deletePolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.DeletePolicyTemplateVersionOptions)
				deletePolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				deletePolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				deletePolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePolicyTemplateVersionOptions model with no property values
				deletePolicyTemplateVersionOptionsModelNew := new(iampolicymanagementv1.DeletePolicyTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.DeletePolicyTemplateVersion(deletePolicyTemplateVersionOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyTemplateVersion(getPolicyTemplateVersionOptions *GetPolicyTemplateVersionOptions) - Operation response error`, func() {
		getPolicyTemplateVersionPath := "/v1/policy_templates/testString/versions/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicyTemplateVersion with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateVersionOptions model
				getPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				getPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyTemplateVersion(getPolicyTemplateVersionOptions *GetPolicyTemplateVersionOptions)`, func() {
		getPolicyTemplateVersionPath := "/v1/policy_templates/testString/versions/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetPolicyTemplateVersion successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetPolicyTemplateVersionOptions model
				getPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				getPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetPolicyTemplateVersionWithContext(ctx, getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetPolicyTemplateVersionWithContext(ctx, getPolicyTemplateVersionOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyTemplateVersionPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "description": "Description", "account_id": "AccountID", "version": "Version", "committed": false, "policy": {"type": "access", "description": "Description", "resource": {"attributes": [{"key": "Key", "operator": "stringEquals", "value": "anyValue"}], "tags": [{"key": "Key", "value": "Value", "operator": "stringEquals"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "timeLessThan", "value": "anyValue"}, "control": {"grant": {"roles": [{"role_id": "RoleID"}]}}}, "id": "ID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetPolicyTemplateVersion successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplateVersion(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyTemplateVersionOptions model
				getPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				getPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPolicyTemplateVersion with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateVersionOptions model
				getPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				getPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPolicyTemplateVersionOptions model with no property values
				getPolicyTemplateVersionOptionsModelNew := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModelNew)
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
			It(`Invoke GetPolicyTemplateVersion successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyTemplateVersionOptions model
				getPolicyTemplateVersionOptionsModel := new(iampolicymanagementv1.GetPolicyTemplateVersionOptions)
				getPolicyTemplateVersionOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Version = core.StringPtr("testString")
				getPolicyTemplateVersionOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetPolicyTemplateVersion(getPolicyTemplateVersionOptionsModel)
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
	Describe(`CommitPolicyTemplate(commitPolicyTemplateOptions *CommitPolicyTemplateOptions)`, func() {
		commitPolicyTemplatePath := "/v1/policy_templates/testString/versions/testString/commit"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(commitPolicyTemplatePath))
					Expect(req.Method).To(Equal("POST"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke CommitPolicyTemplate successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.CommitPolicyTemplate(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the CommitPolicyTemplateOptions model
				commitPolicyTemplateOptionsModel := new(iampolicymanagementv1.CommitPolicyTemplateOptions)
				commitPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				commitPolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				commitPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.CommitPolicyTemplate(commitPolicyTemplateOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke CommitPolicyTemplate with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the CommitPolicyTemplateOptions model
				commitPolicyTemplateOptionsModel := new(iampolicymanagementv1.CommitPolicyTemplateOptions)
				commitPolicyTemplateOptionsModel.PolicyTemplateID = core.StringPtr("testString")
				commitPolicyTemplateOptionsModel.Version = core.StringPtr("testString")
				commitPolicyTemplateOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.CommitPolicyTemplate(commitPolicyTemplateOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the CommitPolicyTemplateOptions model with no property values
				commitPolicyTemplateOptionsModelNew := new(iampolicymanagementv1.CommitPolicyTemplateOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.CommitPolicyTemplate(commitPolicyTemplateOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicyAssignments(listPolicyAssignmentsOptions *ListPolicyAssignmentsOptions) - Operation response error`, func() {
		listPolicyAssignmentsPath := "/v1/policy_assignments"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPolicyAssignments with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyAssignmentsOptions model
				listPolicyAssignmentsOptionsModel := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				listPolicyAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPolicyAssignments(listPolicyAssignmentsOptions *ListPolicyAssignmentsOptions)`, func() {
		listPolicyAssignmentsPath := "/v1/policy_assignments"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"assignments": [{"template_id": "TemplateID", "template_version": "TemplateVersion", "assignment_id": "AssignmentID", "target_type": "Account", "target": "Target", "options": [{"subject_type": "iam_id", "subject_id": "SubjectID", "root_requester_id": "RootRequesterID", "root_template_id": "RootTemplateID", "root_template_version": "RootTemplateVersion"}], "id": "ID", "account_id": "AccountID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "resources": [{"target": "Target", "policy": {"resource_created": {"id": "ID"}, "status": "Status", "error_message": {"trace": "Trace", "errors": [{"code": "insufficent_permissions", "message": "Message", "details": {"conflicts_with": {"etag": "Etag", "role": "Role", "policy": "Policy"}}, "more_info": "MoreInfo"}], "status_code": 10}}}], "status": "in_progress"}]}`)
				}))
			})
			It(`Invoke ListPolicyAssignments successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the ListPolicyAssignmentsOptions model
				listPolicyAssignmentsOptionsModel := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				listPolicyAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.ListPolicyAssignmentsWithContext(ctx, listPolicyAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.ListPolicyAssignmentsWithContext(ctx, listPolicyAssignmentsOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listPolicyAssignmentsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					Expect(req.URL.Query()["account_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_id"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["template_version"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"assignments": [{"template_id": "TemplateID", "template_version": "TemplateVersion", "assignment_id": "AssignmentID", "target_type": "Account", "target": "Target", "options": [{"subject_type": "iam_id", "subject_id": "SubjectID", "root_requester_id": "RootRequesterID", "root_template_id": "RootTemplateID", "root_template_version": "RootTemplateVersion"}], "id": "ID", "account_id": "AccountID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "resources": [{"target": "Target", "policy": {"resource_created": {"id": "ID"}, "status": "Status", "error_message": {"trace": "Trace", "errors": [{"code": "insufficent_permissions", "message": "Message", "details": {"conflicts_with": {"etag": "Etag", "role": "Role", "policy": "Policy"}}, "more_info": "MoreInfo"}], "status_code": 10}}}], "status": "in_progress"}]}`)
				}))
			})
			It(`Invoke ListPolicyAssignments successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.ListPolicyAssignments(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPolicyAssignmentsOptions model
				listPolicyAssignmentsOptionsModel := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				listPolicyAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPolicyAssignments with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyAssignmentsOptions model
				listPolicyAssignmentsOptionsModel := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				listPolicyAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPolicyAssignmentsOptions model with no property values
				listPolicyAssignmentsOptionsModelNew := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModelNew)
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
			It(`Invoke ListPolicyAssignments successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the ListPolicyAssignmentsOptions model
				listPolicyAssignmentsOptionsModel := new(iampolicymanagementv1.ListPolicyAssignmentsOptions)
				listPolicyAssignmentsOptionsModel.AccountID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.AcceptLanguage = core.StringPtr("default")
				listPolicyAssignmentsOptionsModel.TemplateID = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.TemplateVersion = core.StringPtr("testString")
				listPolicyAssignmentsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.ListPolicyAssignments(listPolicyAssignmentsOptionsModel)
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
	Describe(`GetPolicyAssignment(getPolicyAssignmentOptions *GetPolicyAssignmentOptions) - Operation response error`, func() {
		getPolicyAssignmentPath := "/v1/policy_assignments/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyAssignmentPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicyAssignment with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyAssignmentOptions model
				getPolicyAssignmentOptionsModel := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				getPolicyAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getPolicyAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicyAssignment(getPolicyAssignmentOptions *GetPolicyAssignmentOptions)`, func() {
		getPolicyAssignmentPath := "/v1/policy_assignments/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"template_id": "TemplateID", "template_version": "TemplateVersion", "assignment_id": "AssignmentID", "target_type": "Account", "target": "Target", "options": [{"subject_type": "iam_id", "subject_id": "SubjectID", "root_requester_id": "RootRequesterID", "root_template_id": "RootTemplateID", "root_template_version": "RootTemplateVersion"}], "id": "ID", "account_id": "AccountID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "resources": [{"target": "Target", "policy": {"resource_created": {"id": "ID"}, "status": "Status", "error_message": {"trace": "Trace", "errors": [{"code": "insufficent_permissions", "message": "Message", "details": {"conflicts_with": {"etag": "Etag", "role": "Role", "policy": "Policy"}}, "more_info": "MoreInfo"}], "status_code": 10}}}], "status": "in_progress"}`)
				}))
			})
			It(`Invoke GetPolicyAssignment successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the GetPolicyAssignmentOptions model
				getPolicyAssignmentOptionsModel := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				getPolicyAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getPolicyAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.GetPolicyAssignmentWithContext(ctx, getPolicyAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.GetPolicyAssignmentWithContext(ctx, getPolicyAssignmentOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(getPolicyAssignmentPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"template_id": "TemplateID", "template_version": "TemplateVersion", "assignment_id": "AssignmentID", "target_type": "Account", "target": "Target", "options": [{"subject_type": "iam_id", "subject_id": "SubjectID", "root_requester_id": "RootRequesterID", "root_template_id": "RootTemplateID", "root_template_version": "RootTemplateVersion"}], "id": "ID", "account_id": "AccountID", "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "resources": [{"target": "Target", "policy": {"resource_created": {"id": "ID"}, "status": "Status", "error_message": {"trace": "Trace", "errors": [{"code": "insufficent_permissions", "message": "Message", "details": {"conflicts_with": {"etag": "Etag", "role": "Role", "policy": "Policy"}}, "more_info": "MoreInfo"}], "status_code": 10}}}], "status": "in_progress"}`)
				}))
			})
			It(`Invoke GetPolicyAssignment successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.GetPolicyAssignment(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPolicyAssignmentOptions model
				getPolicyAssignmentOptionsModel := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				getPolicyAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getPolicyAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPolicyAssignment with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyAssignmentOptions model
				getPolicyAssignmentOptionsModel := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				getPolicyAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getPolicyAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPolicyAssignmentOptions model with no property values
				getPolicyAssignmentOptionsModelNew := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModelNew)
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
			It(`Invoke GetPolicyAssignment successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the GetPolicyAssignmentOptions model
				getPolicyAssignmentOptionsModel := new(iampolicymanagementv1.GetPolicyAssignmentOptions)
				getPolicyAssignmentOptionsModel.AssignmentID = core.StringPtr("testString")
				getPolicyAssignmentOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.GetPolicyAssignment(getPolicyAssignmentOptionsModel)
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
			iamPolicyManagementService, _ := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL:           "http://iampolicymanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCommitPolicyTemplateOptions successfully`, func() {
				// Construct an instance of the CommitPolicyTemplateOptions model
				policyTemplateID := "testString"
				version := "testString"
				commitPolicyTemplateOptionsModel := iamPolicyManagementService.NewCommitPolicyTemplateOptions(policyTemplateID, version)
				commitPolicyTemplateOptionsModel.SetPolicyTemplateID("testString")
				commitPolicyTemplateOptionsModel.SetVersion("testString")
				commitPolicyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(commitPolicyTemplateOptionsModel).ToNot(BeNil())
				Expect(commitPolicyTemplateOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(commitPolicyTemplateOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(commitPolicyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewControl successfully`, func() {
				var grant *iampolicymanagementv1.Grant = nil
				_, err := iamPolicyManagementService.NewControl(grant)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreatePolicyOptions successfully`, func() {
				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				Expect(subjectAttributeModel).ToNot(BeNil())
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")
				Expect(subjectAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(subjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				Expect(policySubjectModel).ToNot(BeNil())
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}
				Expect(policySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}))

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				Expect(resourceTagModel).ToNot(BeNil())
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")
				Expect(resourceTagModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				Expect(policyResourceModel).ToNot(BeNil())
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}
				Expect(policyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}))
				Expect(policyResourceModel.Tags).To(Equal([]iampolicymanagementv1.ResourceTag{*resourceTagModel}))

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsType := "testString"
				createPolicyOptionsSubjects := []iampolicymanagementv1.PolicySubject{}
				createPolicyOptionsRoles := []iampolicymanagementv1.PolicyRole{}
				createPolicyOptionsResources := []iampolicymanagementv1.PolicyResource{}
				createPolicyOptionsModel := iamPolicyManagementService.NewCreatePolicyOptions(createPolicyOptionsType, createPolicyOptionsSubjects, createPolicyOptionsRoles, createPolicyOptionsResources)
				createPolicyOptionsModel.SetType("testString")
				createPolicyOptionsModel.SetSubjects([]iampolicymanagementv1.PolicySubject{*policySubjectModel})
				createPolicyOptionsModel.SetRoles([]iampolicymanagementv1.PolicyRole{*policyRoleModel})
				createPolicyOptionsModel.SetResources([]iampolicymanagementv1.PolicyResource{*policyResourceModel})
				createPolicyOptionsModel.SetDescription("testString")
				createPolicyOptionsModel.SetAcceptLanguage("default")
				createPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPolicyOptionsModel).ToNot(BeNil())
				Expect(createPolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyOptionsModel.Subjects).To(Equal([]iampolicymanagementv1.PolicySubject{*policySubjectModel}))
				Expect(createPolicyOptionsModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))
				Expect(createPolicyOptionsModel.Resources).To(Equal([]iampolicymanagementv1.PolicyResource{*policyResourceModel}))
				Expect(createPolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(createPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePolicyTemplateOptions successfully`, func() {
				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				Expect(v2PolicyResourceAttributeModel).ToNot(BeNil())
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyResourceAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicyResourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				Expect(v2PolicyResourceTagModel).ToNot(BeNil())
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")
				Expect(v2PolicyResourceTagModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Operator).To(Equal(core.StringPtr("stringEquals")))

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				Expect(v2PolicyResourceModel).ToNot(BeNil())
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}
				Expect(v2PolicyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}))
				Expect(v2PolicyResourceModel.Tags).To(Equal([]iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}))

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				Expect(v2PolicyRuleModel).ToNot(BeNil())
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyRuleModel.Operator).To(Equal(core.StringPtr("timeLessThan")))
				Expect(v2PolicyRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				Expect(rolesModel).ToNot(BeNil())
				rolesModel.RoleID = core.StringPtr("testString")
				Expect(rolesModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				Expect(grantModel).ToNot(BeNil())
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}
				Expect(grantModel.Roles).To(Equal([]iampolicymanagementv1.Roles{*rolesModel}))

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				Expect(controlModel).ToNot(BeNil())
				controlModel.Grant = grantModel
				Expect(controlModel.Grant).To(Equal(grantModel))

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				Expect(templatePolicyModel).ToNot(BeNil())
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel
				Expect(templatePolicyModel.Type).To(Equal(core.StringPtr("access")))
				Expect(templatePolicyModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Resource).To(Equal(v2PolicyResourceModel))
				Expect(templatePolicyModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Rule).To(Equal(v2PolicyRuleModel))
				Expect(templatePolicyModel.Control).To(Equal(controlModel))

				// Construct an instance of the CreatePolicyTemplateOptions model
				createPolicyTemplateOptionsName := "testString"
				createPolicyTemplateOptionsAccountID := "testString"
				var createPolicyTemplateOptionsPolicy *iampolicymanagementv1.TemplatePolicy = nil
				createPolicyTemplateOptionsModel := iamPolicyManagementService.NewCreatePolicyTemplateOptions(createPolicyTemplateOptionsName, createPolicyTemplateOptionsAccountID, createPolicyTemplateOptionsPolicy)
				createPolicyTemplateOptionsModel.SetName("testString")
				createPolicyTemplateOptionsModel.SetAccountID("testString")
				createPolicyTemplateOptionsModel.SetPolicy(templatePolicyModel)
				createPolicyTemplateOptionsModel.SetDescription("testString")
				createPolicyTemplateOptionsModel.SetCommitted(true)
				createPolicyTemplateOptionsModel.SetAcceptLanguage("default")
				createPolicyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPolicyTemplateOptionsModel).ToNot(BeNil())
				Expect(createPolicyTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateOptionsModel.Policy).To(Equal(templatePolicyModel))
				Expect(createPolicyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateOptionsModel.Committed).To(Equal(core.BoolPtr(true)))
				Expect(createPolicyTemplateOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(createPolicyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreatePolicyTemplateVersionOptions successfully`, func() {
				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				Expect(v2PolicyResourceAttributeModel).ToNot(BeNil())
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyResourceAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicyResourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				Expect(v2PolicyResourceTagModel).ToNot(BeNil())
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")
				Expect(v2PolicyResourceTagModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Operator).To(Equal(core.StringPtr("stringEquals")))

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				Expect(v2PolicyResourceModel).ToNot(BeNil())
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}
				Expect(v2PolicyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}))
				Expect(v2PolicyResourceModel.Tags).To(Equal([]iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}))

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				Expect(v2PolicyRuleModel).ToNot(BeNil())
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyRuleModel.Operator).To(Equal(core.StringPtr("timeLessThan")))
				Expect(v2PolicyRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				Expect(rolesModel).ToNot(BeNil())
				rolesModel.RoleID = core.StringPtr("testString")
				Expect(rolesModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				Expect(grantModel).ToNot(BeNil())
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}
				Expect(grantModel.Roles).To(Equal([]iampolicymanagementv1.Roles{*rolesModel}))

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				Expect(controlModel).ToNot(BeNil())
				controlModel.Grant = grantModel
				Expect(controlModel.Grant).To(Equal(grantModel))

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				Expect(templatePolicyModel).ToNot(BeNil())
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel
				Expect(templatePolicyModel.Type).To(Equal(core.StringPtr("access")))
				Expect(templatePolicyModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Resource).To(Equal(v2PolicyResourceModel))
				Expect(templatePolicyModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Rule).To(Equal(v2PolicyRuleModel))
				Expect(templatePolicyModel.Control).To(Equal(controlModel))

				// Construct an instance of the CreatePolicyTemplateVersionOptions model
				policyTemplateID := "testString"
				var createPolicyTemplateVersionOptionsPolicy *iampolicymanagementv1.TemplatePolicy = nil
				createPolicyTemplateVersionOptionsModel := iamPolicyManagementService.NewCreatePolicyTemplateVersionOptions(policyTemplateID, createPolicyTemplateVersionOptionsPolicy)
				createPolicyTemplateVersionOptionsModel.SetPolicyTemplateID("testString")
				createPolicyTemplateVersionOptionsModel.SetPolicy(templatePolicyModel)
				createPolicyTemplateVersionOptionsModel.SetName("testString")
				createPolicyTemplateVersionOptionsModel.SetDescription("testString")
				createPolicyTemplateVersionOptionsModel.SetCommitted(true)
				createPolicyTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPolicyTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(createPolicyTemplateVersionOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateVersionOptionsModel.Policy).To(Equal(templatePolicyModel))
				Expect(createPolicyTemplateVersionOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateVersionOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyTemplateVersionOptionsModel.Committed).To(Equal(core.BoolPtr(true)))
				Expect(createPolicyTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRoleOptions successfully`, func() {
				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsDisplayName := "testString"
				createRoleOptionsActions := []string{"testString"}
				createRoleOptionsName := "Developer"
				createRoleOptionsAccountID := "testString"
				createRoleOptionsServiceName := "iam-groups"
				createRoleOptionsModel := iamPolicyManagementService.NewCreateRoleOptions(createRoleOptionsDisplayName, createRoleOptionsActions, createRoleOptionsName, createRoleOptionsAccountID, createRoleOptionsServiceName)
				createRoleOptionsModel.SetDisplayName("testString")
				createRoleOptionsModel.SetActions([]string{"testString"})
				createRoleOptionsModel.SetName("Developer")
				createRoleOptionsModel.SetAccountID("testString")
				createRoleOptionsModel.SetServiceName("iam-groups")
				createRoleOptionsModel.SetDescription("testString")
				createRoleOptionsModel.SetAcceptLanguage("default")
				createRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRoleOptionsModel).ToNot(BeNil())
				Expect(createRoleOptionsModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.Actions).To(Equal([]string{"testString"}))
				Expect(createRoleOptionsModel.Name).To(Equal(core.StringPtr("Developer")))
				Expect(createRoleOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(createRoleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(createRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateV2PolicyOptions successfully`, func() {
				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				Expect(rolesModel).ToNot(BeNil())
				rolesModel.RoleID = core.StringPtr("testString")
				Expect(rolesModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				Expect(grantModel).ToNot(BeNil())
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}
				Expect(grantModel.Roles).To(Equal([]iampolicymanagementv1.Roles{*rolesModel}))

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				Expect(controlModel).ToNot(BeNil())
				controlModel.Grant = grantModel
				Expect(controlModel.Grant).To(Equal(grantModel))

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				Expect(v2PolicySubjectAttributeModel).ToNot(BeNil())
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicySubjectAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicySubjectAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicySubjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				Expect(v2PolicySubjectModel).ToNot(BeNil())
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}
				Expect(v2PolicySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}))

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				Expect(v2PolicyResourceAttributeModel).ToNot(BeNil())
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyResourceAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicyResourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				Expect(v2PolicyResourceTagModel).ToNot(BeNil())
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")
				Expect(v2PolicyResourceTagModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Operator).To(Equal(core.StringPtr("stringEquals")))

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				Expect(v2PolicyResourceModel).ToNot(BeNil())
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}
				Expect(v2PolicyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}))
				Expect(v2PolicyResourceModel.Tags).To(Equal([]iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}))

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				Expect(v2PolicyRuleModel).ToNot(BeNil())
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyRuleModel.Operator).To(Equal(core.StringPtr("timeLessThan")))
				Expect(v2PolicyRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateV2PolicyOptions model
				var createV2PolicyOptionsControl *iampolicymanagementv1.Control = nil
				createV2PolicyOptionsType := "access"
				createV2PolicyOptionsModel := iamPolicyManagementService.NewCreateV2PolicyOptions(createV2PolicyOptionsControl, createV2PolicyOptionsType)
				createV2PolicyOptionsModel.SetControl(controlModel)
				createV2PolicyOptionsModel.SetType("access")
				createV2PolicyOptionsModel.SetDescription("testString")
				createV2PolicyOptionsModel.SetSubject(v2PolicySubjectModel)
				createV2PolicyOptionsModel.SetResource(v2PolicyResourceModel)
				createV2PolicyOptionsModel.SetPattern("testString")
				createV2PolicyOptionsModel.SetRule(v2PolicyRuleModel)
				createV2PolicyOptionsModel.SetAcceptLanguage("default")
				createV2PolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createV2PolicyOptionsModel).ToNot(BeNil())
				Expect(createV2PolicyOptionsModel.Control).To(Equal(controlModel))
				Expect(createV2PolicyOptionsModel.Type).To(Equal(core.StringPtr("access")))
				Expect(createV2PolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createV2PolicyOptionsModel.Subject).To(Equal(v2PolicySubjectModel))
				Expect(createV2PolicyOptionsModel.Resource).To(Equal(v2PolicyResourceModel))
				Expect(createV2PolicyOptionsModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(createV2PolicyOptionsModel.Rule).To(Equal(v2PolicyRuleModel))
				Expect(createV2PolicyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(createV2PolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePolicyOptions successfully`, func() {
				// Construct an instance of the DeletePolicyOptions model
				policyID := "testString"
				deletePolicyOptionsModel := iamPolicyManagementService.NewDeletePolicyOptions(policyID)
				deletePolicyOptionsModel.SetPolicyID("testString")
				deletePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePolicyOptionsModel).ToNot(BeNil())
				Expect(deletePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(deletePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePolicyTemplateOptions successfully`, func() {
				// Construct an instance of the DeletePolicyTemplateOptions model
				policyTemplateID := "testString"
				deletePolicyTemplateOptionsModel := iamPolicyManagementService.NewDeletePolicyTemplateOptions(policyTemplateID)
				deletePolicyTemplateOptionsModel.SetPolicyTemplateID("testString")
				deletePolicyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePolicyTemplateOptionsModel).ToNot(BeNil())
				Expect(deletePolicyTemplateOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deletePolicyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePolicyTemplateVersionOptions successfully`, func() {
				// Construct an instance of the DeletePolicyTemplateVersionOptions model
				policyTemplateID := "testString"
				version := "testString"
				deletePolicyTemplateVersionOptionsModel := iamPolicyManagementService.NewDeletePolicyTemplateVersionOptions(policyTemplateID, version)
				deletePolicyTemplateVersionOptionsModel.SetPolicyTemplateID("testString")
				deletePolicyTemplateVersionOptionsModel.SetVersion("testString")
				deletePolicyTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePolicyTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(deletePolicyTemplateVersionOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(deletePolicyTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(deletePolicyTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRoleOptions successfully`, func() {
				// Construct an instance of the DeleteRoleOptions model
				roleID := "testString"
				deleteRoleOptionsModel := iamPolicyManagementService.NewDeleteRoleOptions(roleID)
				deleteRoleOptionsModel.SetRoleID("testString")
				deleteRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRoleOptionsModel).ToNot(BeNil())
				Expect(deleteRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteV2PolicyOptions successfully`, func() {
				// Construct an instance of the DeleteV2PolicyOptions model
				id := "testString"
				deleteV2PolicyOptionsModel := iamPolicyManagementService.NewDeleteV2PolicyOptions(id)
				deleteV2PolicyOptionsModel.SetID("testString")
				deleteV2PolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteV2PolicyOptionsModel).ToNot(BeNil())
				Expect(deleteV2PolicyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteV2PolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyAssignmentOptions successfully`, func() {
				// Construct an instance of the GetPolicyAssignmentOptions model
				assignmentID := "testString"
				getPolicyAssignmentOptionsModel := iamPolicyManagementService.NewGetPolicyAssignmentOptions(assignmentID)
				getPolicyAssignmentOptionsModel.SetAssignmentID("testString")
				getPolicyAssignmentOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyAssignmentOptionsModel).ToNot(BeNil())
				Expect(getPolicyAssignmentOptionsModel.AssignmentID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyAssignmentOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyOptions successfully`, func() {
				// Construct an instance of the GetPolicyOptions model
				policyID := "testString"
				getPolicyOptionsModel := iamPolicyManagementService.NewGetPolicyOptions(policyID)
				getPolicyOptionsModel.SetPolicyID("testString")
				getPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyOptionsModel).ToNot(BeNil())
				Expect(getPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyTemplateOptions successfully`, func() {
				// Construct an instance of the GetPolicyTemplateOptions model
				policyTemplateID := "testString"
				getPolicyTemplateOptionsModel := iamPolicyManagementService.NewGetPolicyTemplateOptions(policyTemplateID)
				getPolicyTemplateOptionsModel.SetPolicyTemplateID("testString")
				getPolicyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyTemplateOptionsModel).ToNot(BeNil())
				Expect(getPolicyTemplateOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyTemplateVersionOptions successfully`, func() {
				// Construct an instance of the GetPolicyTemplateVersionOptions model
				policyTemplateID := "testString"
				version := "testString"
				getPolicyTemplateVersionOptionsModel := iamPolicyManagementService.NewGetPolicyTemplateVersionOptions(policyTemplateID, version)
				getPolicyTemplateVersionOptionsModel.SetPolicyTemplateID("testString")
				getPolicyTemplateVersionOptionsModel.SetVersion("testString")
				getPolicyTemplateVersionOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyTemplateVersionOptionsModel).ToNot(BeNil())
				Expect(getPolicyTemplateVersionOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyTemplateVersionOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyTemplateVersionOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRoleOptions successfully`, func() {
				// Construct an instance of the GetRoleOptions model
				roleID := "testString"
				getRoleOptionsModel := iamPolicyManagementService.NewGetRoleOptions(roleID)
				getRoleOptionsModel.SetRoleID("testString")
				getRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRoleOptionsModel).ToNot(BeNil())
				Expect(getRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(getRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetV2PolicyOptions successfully`, func() {
				// Construct an instance of the GetV2PolicyOptions model
				id := "testString"
				getV2PolicyOptionsModel := iamPolicyManagementService.NewGetV2PolicyOptions(id)
				getV2PolicyOptionsModel.SetID("testString")
				getV2PolicyOptionsModel.SetFormat("include_last_permit")
				getV2PolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getV2PolicyOptionsModel).ToNot(BeNil())
				Expect(getV2PolicyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getV2PolicyOptionsModel.Format).To(Equal(core.StringPtr("include_last_permit")))
				Expect(getV2PolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGrant successfully`, func() {
				roles := []iampolicymanagementv1.Roles{}
				_model, err := iamPolicyManagementService.NewGrant(roles)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListPoliciesOptions successfully`, func() {
				// Construct an instance of the ListPoliciesOptions model
				accountID := "testString"
				listPoliciesOptionsModel := iamPolicyManagementService.NewListPoliciesOptions(accountID)
				listPoliciesOptionsModel.SetAccountID("testString")
				listPoliciesOptionsModel.SetAcceptLanguage("default")
				listPoliciesOptionsModel.SetIamID("testString")
				listPoliciesOptionsModel.SetAccessGroupID("testString")
				listPoliciesOptionsModel.SetType("access")
				listPoliciesOptionsModel.SetServiceType("service")
				listPoliciesOptionsModel.SetTagName("testString")
				listPoliciesOptionsModel.SetTagValue("testString")
				listPoliciesOptionsModel.SetSort("id")
				listPoliciesOptionsModel.SetFormat("include_last_permit")
				listPoliciesOptionsModel.SetState("active")
				listPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPoliciesOptionsModel).ToNot(BeNil())
				Expect(listPoliciesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listPoliciesOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.Type).To(Equal(core.StringPtr("access")))
				Expect(listPoliciesOptionsModel.ServiceType).To(Equal(core.StringPtr("service")))
				Expect(listPoliciesOptionsModel.TagName).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.TagValue).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.Sort).To(Equal(core.StringPtr("id")))
				Expect(listPoliciesOptionsModel.Format).To(Equal(core.StringPtr("include_last_permit")))
				Expect(listPoliciesOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(listPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPolicyAssignmentsOptions successfully`, func() {
				// Construct an instance of the ListPolicyAssignmentsOptions model
				accountID := "testString"
				listPolicyAssignmentsOptionsModel := iamPolicyManagementService.NewListPolicyAssignmentsOptions(accountID)
				listPolicyAssignmentsOptionsModel.SetAccountID("testString")
				listPolicyAssignmentsOptionsModel.SetAcceptLanguage("default")
				listPolicyAssignmentsOptionsModel.SetTemplateID("testString")
				listPolicyAssignmentsOptionsModel.SetTemplateVersion("testString")
				listPolicyAssignmentsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPolicyAssignmentsOptionsModel).ToNot(BeNil())
				Expect(listPolicyAssignmentsOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listPolicyAssignmentsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listPolicyAssignmentsOptionsModel.TemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listPolicyAssignmentsOptionsModel.TemplateVersion).To(Equal(core.StringPtr("testString")))
				Expect(listPolicyAssignmentsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPolicyTemplateVersionsOptions successfully`, func() {
				// Construct an instance of the ListPolicyTemplateVersionsOptions model
				policyTemplateID := "testString"
				listPolicyTemplateVersionsOptionsModel := iamPolicyManagementService.NewListPolicyTemplateVersionsOptions(policyTemplateID)
				listPolicyTemplateVersionsOptionsModel.SetPolicyTemplateID("testString")
				listPolicyTemplateVersionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPolicyTemplateVersionsOptionsModel).ToNot(BeNil())
				Expect(listPolicyTemplateVersionsOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(listPolicyTemplateVersionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPolicyTemplatesOptions successfully`, func() {
				// Construct an instance of the ListPolicyTemplatesOptions model
				accountID := "testString"
				listPolicyTemplatesOptionsModel := iamPolicyManagementService.NewListPolicyTemplatesOptions(accountID)
				listPolicyTemplatesOptionsModel.SetAccountID("testString")
				listPolicyTemplatesOptionsModel.SetAcceptLanguage("default")
				listPolicyTemplatesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPolicyTemplatesOptionsModel).ToNot(BeNil())
				Expect(listPolicyTemplatesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listPolicyTemplatesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listPolicyTemplatesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRolesOptions successfully`, func() {
				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := iamPolicyManagementService.NewListRolesOptions()
				listRolesOptionsModel.SetAcceptLanguage("default")
				listRolesOptionsModel.SetAccountID("testString")
				listRolesOptionsModel.SetServiceName("iam-groups")
				listRolesOptionsModel.SetSourceServiceName("iam-groups")
				listRolesOptionsModel.SetPolicyType("authorization")
				listRolesOptionsModel.SetServiceGroupID("IAM")
				listRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRolesOptionsModel).ToNot(BeNil())
				Expect(listRolesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listRolesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listRolesOptionsModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(listRolesOptionsModel.SourceServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(listRolesOptionsModel.PolicyType).To(Equal(core.StringPtr("authorization")))
				Expect(listRolesOptionsModel.ServiceGroupID).To(Equal(core.StringPtr("IAM")))
				Expect(listRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListV2PoliciesOptions successfully`, func() {
				// Construct an instance of the ListV2PoliciesOptions model
				accountID := "testString"
				listV2PoliciesOptionsModel := iamPolicyManagementService.NewListV2PoliciesOptions(accountID)
				listV2PoliciesOptionsModel.SetAccountID("testString")
				listV2PoliciesOptionsModel.SetAcceptLanguage("default")
				listV2PoliciesOptionsModel.SetIamID("testString")
				listV2PoliciesOptionsModel.SetAccessGroupID("testString")
				listV2PoliciesOptionsModel.SetType("access")
				listV2PoliciesOptionsModel.SetServiceType("service")
				listV2PoliciesOptionsModel.SetServiceName("testString")
				listV2PoliciesOptionsModel.SetServiceGroupID("testString")
				listV2PoliciesOptionsModel.SetSort("testString")
				listV2PoliciesOptionsModel.SetFormat("include_last_permit")
				listV2PoliciesOptionsModel.SetState("active")
				listV2PoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listV2PoliciesOptionsModel).ToNot(BeNil())
				Expect(listV2PoliciesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listV2PoliciesOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.Type).To(Equal(core.StringPtr("access")))
				Expect(listV2PoliciesOptionsModel.ServiceType).To(Equal(core.StringPtr("service")))
				Expect(listV2PoliciesOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.ServiceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.Sort).To(Equal(core.StringPtr("testString")))
				Expect(listV2PoliciesOptionsModel.Format).To(Equal(core.StringPtr("include_last_permit")))
				Expect(listV2PoliciesOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(listV2PoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPolicyRole successfully`, func() {
				roleID := "testString"
				_model, err := iamPolicyManagementService.NewPolicyRole(roleID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplacePolicyOptions successfully`, func() {
				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				Expect(subjectAttributeModel).ToNot(BeNil())
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")
				Expect(subjectAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(subjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				Expect(policySubjectModel).ToNot(BeNil())
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}
				Expect(policySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}))

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ResourceTag model
				resourceTagModel := new(iampolicymanagementv1.ResourceTag)
				Expect(resourceTagModel).ToNot(BeNil())
				resourceTagModel.Name = core.StringPtr("testString")
				resourceTagModel.Value = core.StringPtr("testString")
				resourceTagModel.Operator = core.StringPtr("testString")
				Expect(resourceTagModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceTagModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				Expect(policyResourceModel).ToNot(BeNil())
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				policyResourceModel.Tags = []iampolicymanagementv1.ResourceTag{*resourceTagModel}
				Expect(policyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}))
				Expect(policyResourceModel.Tags).To(Equal([]iampolicymanagementv1.ResourceTag{*resourceTagModel}))

				// Construct an instance of the ReplacePolicyOptions model
				policyID := "testString"
				ifMatch := "testString"
				replacePolicyOptionsType := "testString"
				replacePolicyOptionsSubjects := []iampolicymanagementv1.PolicySubject{}
				replacePolicyOptionsRoles := []iampolicymanagementv1.PolicyRole{}
				replacePolicyOptionsResources := []iampolicymanagementv1.PolicyResource{}
				replacePolicyOptionsModel := iamPolicyManagementService.NewReplacePolicyOptions(policyID, ifMatch, replacePolicyOptionsType, replacePolicyOptionsSubjects, replacePolicyOptionsRoles, replacePolicyOptionsResources)
				replacePolicyOptionsModel.SetPolicyID("testString")
				replacePolicyOptionsModel.SetIfMatch("testString")
				replacePolicyOptionsModel.SetType("testString")
				replacePolicyOptionsModel.SetSubjects([]iampolicymanagementv1.PolicySubject{*policySubjectModel})
				replacePolicyOptionsModel.SetRoles([]iampolicymanagementv1.PolicyRole{*policyRoleModel})
				replacePolicyOptionsModel.SetResources([]iampolicymanagementv1.PolicyResource{*policyResourceModel})
				replacePolicyOptionsModel.SetDescription("testString")
				replacePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replacePolicyOptionsModel).ToNot(BeNil())
				Expect(replacePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyOptionsModel.Subjects).To(Equal([]iampolicymanagementv1.PolicySubject{*policySubjectModel}))
				Expect(replacePolicyOptionsModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))
				Expect(replacePolicyOptionsModel.Resources).To(Equal([]iampolicymanagementv1.PolicyResource{*policyResourceModel}))
				Expect(replacePolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplacePolicyTemplateOptions successfully`, func() {
				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				Expect(v2PolicyResourceAttributeModel).ToNot(BeNil())
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyResourceAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicyResourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				Expect(v2PolicyResourceTagModel).ToNot(BeNil())
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")
				Expect(v2PolicyResourceTagModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Operator).To(Equal(core.StringPtr("stringEquals")))

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				Expect(v2PolicyResourceModel).ToNot(BeNil())
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}
				Expect(v2PolicyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}))
				Expect(v2PolicyResourceModel.Tags).To(Equal([]iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}))

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				Expect(v2PolicyRuleModel).ToNot(BeNil())
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyRuleModel.Operator).To(Equal(core.StringPtr("timeLessThan")))
				Expect(v2PolicyRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				Expect(rolesModel).ToNot(BeNil())
				rolesModel.RoleID = core.StringPtr("testString")
				Expect(rolesModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				Expect(grantModel).ToNot(BeNil())
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}
				Expect(grantModel.Roles).To(Equal([]iampolicymanagementv1.Roles{*rolesModel}))

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				Expect(controlModel).ToNot(BeNil())
				controlModel.Grant = grantModel
				Expect(controlModel.Grant).To(Equal(grantModel))

				// Construct an instance of the TemplatePolicy model
				templatePolicyModel := new(iampolicymanagementv1.TemplatePolicy)
				Expect(templatePolicyModel).ToNot(BeNil())
				templatePolicyModel.Type = core.StringPtr("access")
				templatePolicyModel.Description = core.StringPtr("testString")
				templatePolicyModel.Resource = v2PolicyResourceModel
				templatePolicyModel.Pattern = core.StringPtr("testString")
				templatePolicyModel.Rule = v2PolicyRuleModel
				templatePolicyModel.Control = controlModel
				Expect(templatePolicyModel.Type).To(Equal(core.StringPtr("access")))
				Expect(templatePolicyModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Resource).To(Equal(v2PolicyResourceModel))
				Expect(templatePolicyModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(templatePolicyModel.Rule).To(Equal(v2PolicyRuleModel))
				Expect(templatePolicyModel.Control).To(Equal(controlModel))

				// Construct an instance of the ReplacePolicyTemplateOptions model
				policyTemplateID := "testString"
				version := "testString"
				ifMatch := "testString"
				var replacePolicyTemplateOptionsPolicy *iampolicymanagementv1.TemplatePolicy = nil
				replacePolicyTemplateOptionsModel := iamPolicyManagementService.NewReplacePolicyTemplateOptions(policyTemplateID, version, ifMatch, replacePolicyTemplateOptionsPolicy)
				replacePolicyTemplateOptionsModel.SetPolicyTemplateID("testString")
				replacePolicyTemplateOptionsModel.SetVersion("testString")
				replacePolicyTemplateOptionsModel.SetIfMatch("testString")
				replacePolicyTemplateOptionsModel.SetPolicy(templatePolicyModel)
				replacePolicyTemplateOptionsModel.SetName("testString")
				replacePolicyTemplateOptionsModel.SetDescription("testString")
				replacePolicyTemplateOptionsModel.SetCommitted(true)
				replacePolicyTemplateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replacePolicyTemplateOptionsModel).ToNot(BeNil())
				Expect(replacePolicyTemplateOptionsModel.PolicyTemplateID).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyTemplateOptionsModel.Version).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyTemplateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyTemplateOptionsModel.Policy).To(Equal(templatePolicyModel))
				Expect(replacePolicyTemplateOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyTemplateOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replacePolicyTemplateOptionsModel.Committed).To(Equal(core.BoolPtr(true)))
				Expect(replacePolicyTemplateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceRoleOptions successfully`, func() {
				// Construct an instance of the ReplaceRoleOptions model
				roleID := "testString"
				ifMatch := "testString"
				replaceRoleOptionsDisplayName := "testString"
				replaceRoleOptionsActions := []string{"testString"}
				replaceRoleOptionsModel := iamPolicyManagementService.NewReplaceRoleOptions(roleID, ifMatch, replaceRoleOptionsDisplayName, replaceRoleOptionsActions)
				replaceRoleOptionsModel.SetRoleID("testString")
				replaceRoleOptionsModel.SetIfMatch("testString")
				replaceRoleOptionsModel.SetDisplayName("testString")
				replaceRoleOptionsModel.SetActions([]string{"testString"})
				replaceRoleOptionsModel.SetDescription("testString")
				replaceRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceRoleOptionsModel).ToNot(BeNil())
				Expect(replaceRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(replaceRoleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceRoleOptionsModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(replaceRoleOptionsModel.Actions).To(Equal([]string{"testString"}))
				Expect(replaceRoleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewReplaceV2PolicyOptions successfully`, func() {
				// Construct an instance of the Roles model
				rolesModel := new(iampolicymanagementv1.Roles)
				Expect(rolesModel).ToNot(BeNil())
				rolesModel.RoleID = core.StringPtr("testString")
				Expect(rolesModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the Grant model
				grantModel := new(iampolicymanagementv1.Grant)
				Expect(grantModel).ToNot(BeNil())
				grantModel.Roles = []iampolicymanagementv1.Roles{*rolesModel}
				Expect(grantModel.Roles).To(Equal([]iampolicymanagementv1.Roles{*rolesModel}))

				// Construct an instance of the Control model
				controlModel := new(iampolicymanagementv1.Control)
				Expect(controlModel).ToNot(BeNil())
				controlModel.Grant = grantModel
				Expect(controlModel.Grant).To(Equal(grantModel))

				// Construct an instance of the V2PolicySubjectAttribute model
				v2PolicySubjectAttributeModel := new(iampolicymanagementv1.V2PolicySubjectAttribute)
				Expect(v2PolicySubjectAttributeModel).ToNot(BeNil())
				v2PolicySubjectAttributeModel.Key = core.StringPtr("testString")
				v2PolicySubjectAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicySubjectAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicySubjectAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicySubjectAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicySubjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicySubject model
				v2PolicySubjectModel := new(iampolicymanagementv1.V2PolicySubject)
				Expect(v2PolicySubjectModel).ToNot(BeNil())
				v2PolicySubjectModel.Attributes = []iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}
				Expect(v2PolicySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicySubjectAttribute{*v2PolicySubjectAttributeModel}))

				// Construct an instance of the V2PolicyResourceAttribute model
				v2PolicyResourceAttributeModel := new(iampolicymanagementv1.V2PolicyResourceAttribute)
				Expect(v2PolicyResourceAttributeModel).ToNot(BeNil())
				v2PolicyResourceAttributeModel.Key = core.StringPtr("testString")
				v2PolicyResourceAttributeModel.Operator = core.StringPtr("stringEquals")
				v2PolicyResourceAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyResourceAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceAttributeModel.Operator).To(Equal(core.StringPtr("stringEquals")))
				Expect(v2PolicyResourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyResourceTag model
				v2PolicyResourceTagModel := new(iampolicymanagementv1.V2PolicyResourceTag)
				Expect(v2PolicyResourceTagModel).ToNot(BeNil())
				v2PolicyResourceTagModel.Key = core.StringPtr("testString")
				v2PolicyResourceTagModel.Value = core.StringPtr("testString")
				v2PolicyResourceTagModel.Operator = core.StringPtr("stringEquals")
				Expect(v2PolicyResourceTagModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyResourceTagModel.Operator).To(Equal(core.StringPtr("stringEquals")))

				// Construct an instance of the V2PolicyResource model
				v2PolicyResourceModel := new(iampolicymanagementv1.V2PolicyResource)
				Expect(v2PolicyResourceModel).ToNot(BeNil())
				v2PolicyResourceModel.Attributes = []iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}
				v2PolicyResourceModel.Tags = []iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}
				Expect(v2PolicyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyResourceAttribute{*v2PolicyResourceAttributeModel}))
				Expect(v2PolicyResourceModel.Tags).To(Equal([]iampolicymanagementv1.V2PolicyResourceTag{*v2PolicyResourceTagModel}))

				// Construct an instance of the V2PolicyRuleRuleAttribute model
				v2PolicyRuleModel := new(iampolicymanagementv1.V2PolicyRuleRuleAttribute)
				Expect(v2PolicyRuleModel).ToNot(BeNil())
				v2PolicyRuleModel.Key = core.StringPtr("testString")
				v2PolicyRuleModel.Operator = core.StringPtr("timeLessThan")
				v2PolicyRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyRuleModel.Operator).To(Equal(core.StringPtr("timeLessThan")))
				Expect(v2PolicyRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ReplaceV2PolicyOptions model
				id := "testString"
				ifMatch := "testString"
				var replaceV2PolicyOptionsControl *iampolicymanagementv1.Control = nil
				replaceV2PolicyOptionsType := "access"
				replaceV2PolicyOptionsModel := iamPolicyManagementService.NewReplaceV2PolicyOptions(id, ifMatch, replaceV2PolicyOptionsControl, replaceV2PolicyOptionsType)
				replaceV2PolicyOptionsModel.SetID("testString")
				replaceV2PolicyOptionsModel.SetIfMatch("testString")
				replaceV2PolicyOptionsModel.SetControl(controlModel)
				replaceV2PolicyOptionsModel.SetType("access")
				replaceV2PolicyOptionsModel.SetDescription("testString")
				replaceV2PolicyOptionsModel.SetSubject(v2PolicySubjectModel)
				replaceV2PolicyOptionsModel.SetResource(v2PolicyResourceModel)
				replaceV2PolicyOptionsModel.SetPattern("testString")
				replaceV2PolicyOptionsModel.SetRule(v2PolicyRuleModel)
				replaceV2PolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceV2PolicyOptionsModel).ToNot(BeNil())
				Expect(replaceV2PolicyOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(replaceV2PolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(replaceV2PolicyOptionsModel.Control).To(Equal(controlModel))
				Expect(replaceV2PolicyOptionsModel.Type).To(Equal(core.StringPtr("access")))
				Expect(replaceV2PolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(replaceV2PolicyOptionsModel.Subject).To(Equal(v2PolicySubjectModel))
				Expect(replaceV2PolicyOptionsModel.Resource).To(Equal(v2PolicyResourceModel))
				Expect(replaceV2PolicyOptionsModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(replaceV2PolicyOptionsModel.Rule).To(Equal(v2PolicyRuleModel))
				Expect(replaceV2PolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewResourceAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := iamPolicyManagementService.NewResourceAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceTag successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := iamPolicyManagementService.NewResourceTag(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRole successfully`, func() {
				displayName := "testString"
				actions := []string{"testString"}
				_model, err := iamPolicyManagementService.NewRole(displayName, actions)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRoles successfully`, func() {
				roleID := "testString"
				_model, err := iamPolicyManagementService.NewRoles(roleID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewRuleAttribute successfully`, func() {
				key := "testString"
				operator := "timeLessThan"
				value := core.StringPtr("testString")
				_model, err := iamPolicyManagementService.NewRuleAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubjectAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := iamPolicyManagementService.NewSubjectAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewTemplatePolicy successfully`, func() {
				typeVar := "access"
				var control *iampolicymanagementv1.Control = nil
				_, err := iamPolicyManagementService.NewTemplatePolicy(typeVar, control)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewUpdatePolicyStateOptions successfully`, func() {
				// Construct an instance of the UpdatePolicyStateOptions model
				policyID := "testString"
				ifMatch := "testString"
				updatePolicyStateOptionsModel := iamPolicyManagementService.NewUpdatePolicyStateOptions(policyID, ifMatch)
				updatePolicyStateOptionsModel.SetPolicyID("testString")
				updatePolicyStateOptionsModel.SetIfMatch("testString")
				updatePolicyStateOptionsModel.SetState("active")
				updatePolicyStateOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePolicyStateOptionsModel).ToNot(BeNil())
				Expect(updatePolicyStateOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyStateOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyStateOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(updatePolicyStateOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2PolicyResource successfully`, func() {
				attributes := []iampolicymanagementv1.V2PolicyResourceAttribute{}
				_model, err := iamPolicyManagementService.NewV2PolicyResource(attributes)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyResourceAttribute successfully`, func() {
				key := "testString"
				operator := "stringEquals"
				value := core.StringPtr("testString")
				_model, err := iamPolicyManagementService.NewV2PolicyResourceAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyResourceTag successfully`, func() {
				key := "testString"
				value := "testString"
				operator := "stringEquals"
				_model, err := iamPolicyManagementService.NewV2PolicyResourceTag(key, value, operator)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicySubject successfully`, func() {
				attributes := []iampolicymanagementv1.V2PolicySubjectAttribute{}
				_model, err := iamPolicyManagementService.NewV2PolicySubject(attributes)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicySubjectAttribute successfully`, func() {
				key := "testString"
				operator := "stringEquals"
				value := "testString"
				_model, err := iamPolicyManagementService.NewV2PolicySubjectAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyRuleRuleAttribute successfully`, func() {
				key := "testString"
				operator := "timeLessThan"
				value := core.StringPtr("testString")
				_model, err := iamPolicyManagementService.NewV2PolicyRuleRuleAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyRuleRuleWithConditions successfully`, func() {
				operator := "and"
				conditions := []iampolicymanagementv1.RuleAttribute{}
				_model, err := iamPolicyManagementService.NewV2PolicyRuleRuleWithConditions(operator, conditions)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
