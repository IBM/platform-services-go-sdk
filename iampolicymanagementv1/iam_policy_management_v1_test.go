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
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}]}`)
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
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}]}`)
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
	Describe(`UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions) - Operation response error`, func() {
		updatePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePolicy with error: Operation response processing error`, func() {
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

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Description = core.StringPtr("testString")
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions)`, func() {
		updatePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyPath))
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
			It(`Invoke UpdatePolicy successfully with retries`, func() {
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

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Description = core.StringPtr("testString")
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.UpdatePolicyWithContext(ctx, updatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.UpdatePolicyWithContext(ctx, updatePolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updatePolicyPath))
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
			It(`Invoke UpdatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.UpdatePolicy(nil)
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

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Description = core.StringPtr("testString")
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdatePolicy with error: Operation validation and request error`, func() {
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

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Description = core.StringPtr("testString")
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePolicyOptions model with no property values
				updatePolicyOptionsModelNew := new(iampolicymanagementv1.UpdatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModelNew)
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
			It(`Invoke UpdatePolicy successfully`, func() {
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

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Description = core.StringPtr("testString")
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.UpdatePolicy(updatePolicyOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
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
	Describe(`PatchPolicy(patchPolicyOptions *PatchPolicyOptions) - Operation response error`, func() {
		patchPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchPolicyPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke PatchPolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PatchPolicyOptions model
				patchPolicyOptionsModel := new(iampolicymanagementv1.PatchPolicyOptions)
				patchPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				patchPolicyOptionsModel.IfMatch = core.StringPtr("testString")
				patchPolicyOptionsModel.State = core.StringPtr("active")
				patchPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`PatchPolicy(patchPolicyOptions *PatchPolicyOptions)`, func() {
		patchPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(patchPolicyPath))
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
			It(`Invoke PatchPolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the PatchPolicyOptions model
				patchPolicyOptionsModel := new(iampolicymanagementv1.PatchPolicyOptions)
				patchPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				patchPolicyOptionsModel.IfMatch = core.StringPtr("testString")
				patchPolicyOptionsModel.State = core.StringPtr("active")
				patchPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.PatchPolicyWithContext(ctx, patchPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.PatchPolicyWithContext(ctx, patchPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(patchPolicyPath))
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
			It(`Invoke PatchPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.PatchPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PatchPolicyOptions model
				patchPolicyOptionsModel := new(iampolicymanagementv1.PatchPolicyOptions)
				patchPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				patchPolicyOptionsModel.IfMatch = core.StringPtr("testString")
				patchPolicyOptionsModel.State = core.StringPtr("active")
				patchPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke PatchPolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PatchPolicyOptions model
				patchPolicyOptionsModel := new(iampolicymanagementv1.PatchPolicyOptions)
				patchPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				patchPolicyOptionsModel.IfMatch = core.StringPtr("testString")
				patchPolicyOptionsModel.State = core.StringPtr("active")
				patchPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the PatchPolicyOptions model with no property values
				patchPolicyOptionsModelNew := new(iampolicymanagementv1.PatchPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModelNew)
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
			It(`Invoke PatchPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PatchPolicyOptions model
				patchPolicyOptionsModel := new(iampolicymanagementv1.PatchPolicyOptions)
				patchPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				patchPolicyOptionsModel.IfMatch = core.StringPtr("testString")
				patchPolicyOptionsModel.State = core.StringPtr("active")
				patchPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.PatchPolicy(patchPolicyOptionsModel)
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
	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions) - Operation response error`, func() {
		updateRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRole with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions)`, func() {
		updateRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
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
			It(`Invoke UpdateRole successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.UpdateRoleWithContext(ctx, updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.UpdateRoleWithContext(ctx, updateRoleOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateRolePath))
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
			It(`Invoke UpdateRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.UpdateRole(nil)
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
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateRole with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRoleOptions model with no property values
				updateRoleOptionsModelNew := new(iampolicymanagementv1.UpdateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.UpdateRole(updateRoleOptionsModelNew)
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
			It(`Invoke UpdateRole successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.UpdateRole(updateRoleOptionsModel)
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
	Describe(`V2ListPolicies(v2ListPoliciesOptions *V2ListPoliciesOptions) - Operation response error`, func() {
		v2ListPoliciesPath := "/v2/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2ListPoliciesPath))
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
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke V2ListPolicies with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2ListPoliciesOptions model
				v2ListPoliciesOptionsModel := new(iampolicymanagementv1.V2ListPoliciesOptions)
				v2ListPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2ListPoliciesOptionsModel.IamID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Type = core.StringPtr("access")
				v2ListPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				v2ListPoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				v2ListPoliciesOptionsModel.State = core.StringPtr("active")
				v2ListPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`V2ListPolicies(v2ListPoliciesOptions *V2ListPoliciesOptions)`, func() {
		v2ListPoliciesPath := "/v2/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2ListPoliciesPath))
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
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}]}`)
				}))
			})
			It(`Invoke V2ListPolicies successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the V2ListPoliciesOptions model
				v2ListPoliciesOptionsModel := new(iampolicymanagementv1.V2ListPoliciesOptions)
				v2ListPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2ListPoliciesOptionsModel.IamID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Type = core.StringPtr("access")
				v2ListPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				v2ListPoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				v2ListPoliciesOptionsModel.State = core.StringPtr("active")
				v2ListPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.V2ListPoliciesWithContext(ctx, v2ListPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.V2ListPoliciesWithContext(ctx, v2ListPoliciesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(v2ListPoliciesPath))
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
					Expect(req.URL.Query()["format"]).To(Equal([]string{"include_last_permit"}))
					Expect(req.URL.Query()["state"]).To(Equal([]string{"active"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"policies": [{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}]}`)
				}))
			})
			It(`Invoke V2ListPolicies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.V2ListPolicies(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the V2ListPoliciesOptions model
				v2ListPoliciesOptionsModel := new(iampolicymanagementv1.V2ListPoliciesOptions)
				v2ListPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2ListPoliciesOptionsModel.IamID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Type = core.StringPtr("access")
				v2ListPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				v2ListPoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				v2ListPoliciesOptionsModel.State = core.StringPtr("active")
				v2ListPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke V2ListPolicies with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2ListPoliciesOptions model
				v2ListPoliciesOptionsModel := new(iampolicymanagementv1.V2ListPoliciesOptions)
				v2ListPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2ListPoliciesOptionsModel.IamID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Type = core.StringPtr("access")
				v2ListPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				v2ListPoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				v2ListPoliciesOptionsModel.State = core.StringPtr("active")
				v2ListPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the V2ListPoliciesOptions model with no property values
				v2ListPoliciesOptionsModelNew := new(iampolicymanagementv1.V2ListPoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModelNew)
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
			It(`Invoke V2ListPolicies successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2ListPoliciesOptions model
				v2ListPoliciesOptionsModel := new(iampolicymanagementv1.V2ListPoliciesOptions)
				v2ListPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2ListPoliciesOptionsModel.IamID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Type = core.StringPtr("access")
				v2ListPoliciesOptionsModel.ServiceType = core.StringPtr("service")
				v2ListPoliciesOptionsModel.ServiceName = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.ServiceGroupID = core.StringPtr("testString")
				v2ListPoliciesOptionsModel.Format = core.StringPtr("include_last_permit")
				v2ListPoliciesOptionsModel.State = core.StringPtr("active")
				v2ListPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.V2ListPolicies(v2ListPoliciesOptionsModel)
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
	Describe(`V2CreatePolicy(v2CreatePolicyOptions *V2CreatePolicyOptions) - Operation response error`, func() {
		v2CreatePolicyPath := "/v2/policies"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2CreatePolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "default")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke V2CreatePolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsModel := new(iampolicymanagementv1.V2CreatePolicyOptions)
				v2CreatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2CreatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2CreatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2CreatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2CreatePolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2CreatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`V2CreatePolicy(v2CreatePolicyOptions *V2CreatePolicyOptions)`, func() {
		v2CreatePolicyPath := "/v2/policies"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2CreatePolicyPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2CreatePolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsModel := new(iampolicymanagementv1.V2CreatePolicyOptions)
				v2CreatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2CreatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2CreatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2CreatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2CreatePolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2CreatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.V2CreatePolicyWithContext(ctx, v2CreatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.V2CreatePolicyWithContext(ctx, v2CreatePolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(v2CreatePolicyPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2CreatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.V2CreatePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsModel := new(iampolicymanagementv1.V2CreatePolicyOptions)
				v2CreatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2CreatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2CreatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2CreatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2CreatePolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2CreatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke V2CreatePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsModel := new(iampolicymanagementv1.V2CreatePolicyOptions)
				v2CreatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2CreatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2CreatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2CreatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2CreatePolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2CreatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the V2CreatePolicyOptions model with no property values
				v2CreatePolicyOptionsModelNew := new(iampolicymanagementv1.V2CreatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModelNew)
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
			It(`Invoke V2CreatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsModel := new(iampolicymanagementv1.V2CreatePolicyOptions)
				v2CreatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2CreatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2CreatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2CreatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2CreatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2CreatePolicyOptionsModel.AcceptLanguage = core.StringPtr("default")
				v2CreatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.V2CreatePolicy(v2CreatePolicyOptionsModel)
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
	Describe(`V2UpdatePolicy(v2UpdatePolicyOptions *V2UpdatePolicyOptions) - Operation response error`, func() {
		v2UpdatePolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2UpdatePolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke V2UpdatePolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2UpdatePolicyOptions model
				v2UpdatePolicyOptionsModel := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				v2UpdatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2UpdatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2UpdatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2UpdatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2UpdatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`V2UpdatePolicy(v2UpdatePolicyOptions *V2UpdatePolicyOptions)`, func() {
		v2UpdatePolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2UpdatePolicyPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2UpdatePolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2UpdatePolicyOptions model
				v2UpdatePolicyOptionsModel := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				v2UpdatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2UpdatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2UpdatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2UpdatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2UpdatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.V2UpdatePolicyWithContext(ctx, v2UpdatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.V2UpdatePolicyWithContext(ctx, v2UpdatePolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(v2UpdatePolicyPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subject": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "control": {"grant": {"roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}]}}, "resource": {"attributes": [{"key": "Key", "operator": "Operator", "value": "anyValue"}]}, "pattern": "Pattern", "rule": {"key": "Key", "operator": "Operator", "value": "anyValue"}, "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2UpdatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.V2UpdatePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2UpdatePolicyOptions model
				v2UpdatePolicyOptionsModel := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				v2UpdatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2UpdatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2UpdatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2UpdatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2UpdatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke V2UpdatePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2UpdatePolicyOptions model
				v2UpdatePolicyOptionsModel := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				v2UpdatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2UpdatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2UpdatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2UpdatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2UpdatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the V2UpdatePolicyOptions model with no property values
				v2UpdatePolicyOptionsModelNew := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModelNew)
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
			It(`Invoke V2UpdatePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")

				// Construct an instance of the V2UpdatePolicyOptions model
				v2UpdatePolicyOptionsModel := new(iampolicymanagementv1.V2UpdatePolicyOptions)
				v2UpdatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Type = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Control = v2PolicyBaseControlModel
				v2UpdatePolicyOptionsModel.Description = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Subject = v2PolicyBaseSubjectModel
				v2UpdatePolicyOptionsModel.Resource = v2PolicyBaseResourceModel
				v2UpdatePolicyOptionsModel.Pattern = core.StringPtr("testString")
				v2UpdatePolicyOptionsModel.Rule = v2PolicyBaseRuleModel
				v2UpdatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.V2UpdatePolicy(v2UpdatePolicyOptionsModel)
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
	Describe(`V2GetPolicy(v2GetPolicyOptions *V2GetPolicyOptions) - Operation response error`, func() {
		v2GetPolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2GetPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke V2GetPolicy with error: Operation response processing error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2GetPolicyOptions model
				v2GetPolicyOptionsModel := new(iampolicymanagementv1.V2GetPolicyOptions)
				v2GetPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2GetPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				iamPolicyManagementService.EnableRetries(0, 0)
				result, response, operationErr = iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`V2GetPolicy(v2GetPolicyOptions *V2GetPolicyOptions)`, func() {
		v2GetPolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2GetPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2GetPolicy successfully with retries`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())
				iamPolicyManagementService.EnableRetries(0, 0)

				// Construct an instance of the V2GetPolicyOptions model
				v2GetPolicyOptionsModel := new(iampolicymanagementv1.V2GetPolicyOptions)
				v2GetPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2GetPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := iamPolicyManagementService.V2GetPolicyWithContext(ctx, v2GetPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				iamPolicyManagementService.DisableRetries()
				result, response, operationErr := iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = iamPolicyManagementService.V2GetPolicyWithContext(ctx, v2GetPolicyOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(v2GetPolicyPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "type": "Type", "description": "Description", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}], "tags": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00.000Z", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00.000Z", "last_modified_by_id": "LastModifiedByID", "state": "active"}`)
				}))
			})
			It(`Invoke V2GetPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := iamPolicyManagementService.V2GetPolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the V2GetPolicyOptions model
				v2GetPolicyOptionsModel := new(iampolicymanagementv1.V2GetPolicyOptions)
				v2GetPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2GetPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke V2GetPolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2GetPolicyOptions model
				v2GetPolicyOptionsModel := new(iampolicymanagementv1.V2GetPolicyOptions)
				v2GetPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2GetPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the V2GetPolicyOptions model with no property values
				v2GetPolicyOptionsModelNew := new(iampolicymanagementv1.V2GetPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModelNew)
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
			It(`Invoke V2GetPolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2GetPolicyOptions model
				v2GetPolicyOptionsModel := new(iampolicymanagementv1.V2GetPolicyOptions)
				v2GetPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2GetPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := iamPolicyManagementService.V2GetPolicy(v2GetPolicyOptionsModel)
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
	Describe(`V2DeletePolicy(v2DeletePolicyOptions *V2DeletePolicyOptions)`, func() {
		v2DeletePolicyPath := "/v2/policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(v2DeletePolicyPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke V2DeletePolicy successfully`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := iamPolicyManagementService.V2DeletePolicy(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the V2DeletePolicyOptions model
				v2DeletePolicyOptionsModel := new(iampolicymanagementv1.V2DeletePolicyOptions)
				v2DeletePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2DeletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = iamPolicyManagementService.V2DeletePolicy(v2DeletePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke V2DeletePolicy with error: Operation validation and request error`, func() {
				iamPolicyManagementService, serviceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(iamPolicyManagementService).ToNot(BeNil())

				// Construct an instance of the V2DeletePolicyOptions model
				v2DeletePolicyOptionsModel := new(iampolicymanagementv1.V2DeletePolicyOptions)
				v2DeletePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				v2DeletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := iamPolicyManagementService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := iamPolicyManagementService.V2DeletePolicy(v2DeletePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the V2DeletePolicyOptions model with no property values
				v2DeletePolicyOptionsModelNew := new(iampolicymanagementv1.V2DeletePolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = iamPolicyManagementService.V2DeletePolicy(v2DeletePolicyOptionsModelNew)
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
			iamPolicyManagementService, _ := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL:           "http://iampolicymanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
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
			It(`Invoke NewListRolesOptions successfully`, func() {
				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := iamPolicyManagementService.NewListRolesOptions()
				listRolesOptionsModel.SetAcceptLanguage("default")
				listRolesOptionsModel.SetAccountID("testString")
				listRolesOptionsModel.SetServiceName("iam-groups")
				listRolesOptionsModel.SetSourceServiceName("iam-groups")
				listRolesOptionsModel.SetPolicyType("authorization")
				listRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRolesOptionsModel).ToNot(BeNil())
				Expect(listRolesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(listRolesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listRolesOptionsModel.ServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(listRolesOptionsModel.SourceServiceName).To(Equal(core.StringPtr("iam-groups")))
				Expect(listRolesOptionsModel.PolicyType).To(Equal(core.StringPtr("authorization")))
				Expect(listRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPatchPolicyOptions successfully`, func() {
				// Construct an instance of the PatchPolicyOptions model
				policyID := "testString"
				ifMatch := "testString"
				patchPolicyOptionsModel := iamPolicyManagementService.NewPatchPolicyOptions(policyID, ifMatch)
				patchPolicyOptionsModel.SetPolicyID("testString")
				patchPolicyOptionsModel.SetIfMatch("testString")
				patchPolicyOptionsModel.SetState("active")
				patchPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(patchPolicyOptionsModel).ToNot(BeNil())
				Expect(patchPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(patchPolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(patchPolicyOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(patchPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePolicyOptions successfully`, func() {
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

				// Construct an instance of the UpdatePolicyOptions model
				policyID := "testString"
				ifMatch := "testString"
				updatePolicyOptionsType := "testString"
				updatePolicyOptionsSubjects := []iampolicymanagementv1.PolicySubject{}
				updatePolicyOptionsRoles := []iampolicymanagementv1.PolicyRole{}
				updatePolicyOptionsResources := []iampolicymanagementv1.PolicyResource{}
				updatePolicyOptionsModel := iamPolicyManagementService.NewUpdatePolicyOptions(policyID, ifMatch, updatePolicyOptionsType, updatePolicyOptionsSubjects, updatePolicyOptionsRoles, updatePolicyOptionsResources)
				updatePolicyOptionsModel.SetPolicyID("testString")
				updatePolicyOptionsModel.SetIfMatch("testString")
				updatePolicyOptionsModel.SetType("testString")
				updatePolicyOptionsModel.SetSubjects([]iampolicymanagementv1.PolicySubject{*policySubjectModel})
				updatePolicyOptionsModel.SetRoles([]iampolicymanagementv1.PolicyRole{*policyRoleModel})
				updatePolicyOptionsModel.SetResources([]iampolicymanagementv1.PolicyResource{*policyResourceModel})
				updatePolicyOptionsModel.SetDescription("testString")
				updatePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePolicyOptionsModel).ToNot(BeNil())
				Expect(updatePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.Subjects).To(Equal([]iampolicymanagementv1.PolicySubject{*policySubjectModel}))
				Expect(updatePolicyOptionsModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))
				Expect(updatePolicyOptionsModel.Resources).To(Equal([]iampolicymanagementv1.PolicyResource{*policyResourceModel}))
				Expect(updatePolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRoleOptions successfully`, func() {
				// Construct an instance of the UpdateRoleOptions model
				roleID := "testString"
				ifMatch := "testString"
				updateRoleOptionsModel := iamPolicyManagementService.NewUpdateRoleOptions(roleID, ifMatch)
				updateRoleOptionsModel.SetRoleID("testString")
				updateRoleOptionsModel.SetIfMatch("testString")
				updateRoleOptionsModel.SetDisplayName("testString")
				updateRoleOptionsModel.SetDescription("testString")
				updateRoleOptionsModel.SetActions([]string{"testString"})
				updateRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateRoleOptionsModel).ToNot(BeNil())
				Expect(updateRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(updateRoleOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updateRoleOptionsModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(updateRoleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateRoleOptionsModel.Actions).To(Equal([]string{"testString"}))
				Expect(updateRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2CreatePolicyOptions successfully`, func() {
				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				Expect(v2PolicyBaseControlGrantModel).ToNot(BeNil())
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				Expect(v2PolicyBaseControlGrantModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				Expect(v2PolicyBaseControlModel).ToNot(BeNil())
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel
				Expect(v2PolicyBaseControlModel.Grant).To(Equal(v2PolicyBaseControlGrantModel))

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				Expect(v2PolicyAttributeModel).ToNot(BeNil())
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyAttributeModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				Expect(v2PolicyBaseSubjectModel).ToNot(BeNil())
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}
				Expect(v2PolicyBaseSubjectModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}))

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				Expect(v2PolicyBaseResourceModel).ToNot(BeNil())
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}
				Expect(v2PolicyBaseResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}))

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				Expect(v2PolicyBaseRuleModel).ToNot(BeNil())
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyBaseRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyBaseRuleModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyBaseRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2CreatePolicyOptions model
				v2CreatePolicyOptionsType := "testString"
				var v2CreatePolicyOptionsControl *iampolicymanagementv1.V2PolicyBaseControl = nil
				v2CreatePolicyOptionsModel := iamPolicyManagementService.NewV2CreatePolicyOptions(v2CreatePolicyOptionsType, v2CreatePolicyOptionsControl)
				v2CreatePolicyOptionsModel.SetType("testString")
				v2CreatePolicyOptionsModel.SetControl(v2PolicyBaseControlModel)
				v2CreatePolicyOptionsModel.SetDescription("testString")
				v2CreatePolicyOptionsModel.SetSubject(v2PolicyBaseSubjectModel)
				v2CreatePolicyOptionsModel.SetResource(v2PolicyBaseResourceModel)
				v2CreatePolicyOptionsModel.SetPattern("testString")
				v2CreatePolicyOptionsModel.SetRule(v2PolicyBaseRuleModel)
				v2CreatePolicyOptionsModel.SetAcceptLanguage("default")
				v2CreatePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v2CreatePolicyOptionsModel).ToNot(BeNil())
				Expect(v2CreatePolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(v2CreatePolicyOptionsModel.Control).To(Equal(v2PolicyBaseControlModel))
				Expect(v2CreatePolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(v2CreatePolicyOptionsModel.Subject).To(Equal(v2PolicyBaseSubjectModel))
				Expect(v2CreatePolicyOptionsModel.Resource).To(Equal(v2PolicyBaseResourceModel))
				Expect(v2CreatePolicyOptionsModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(v2CreatePolicyOptionsModel.Rule).To(Equal(v2PolicyBaseRuleModel))
				Expect(v2CreatePolicyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(v2CreatePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2DeletePolicyOptions successfully`, func() {
				// Construct an instance of the V2DeletePolicyOptions model
				policyID := "testString"
				v2DeletePolicyOptionsModel := iamPolicyManagementService.NewV2DeletePolicyOptions(policyID)
				v2DeletePolicyOptionsModel.SetPolicyID("testString")
				v2DeletePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v2DeletePolicyOptionsModel).ToNot(BeNil())
				Expect(v2DeletePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(v2DeletePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2GetPolicyOptions successfully`, func() {
				// Construct an instance of the V2GetPolicyOptions model
				policyID := "testString"
				v2GetPolicyOptionsModel := iamPolicyManagementService.NewV2GetPolicyOptions(policyID)
				v2GetPolicyOptionsModel.SetPolicyID("testString")
				v2GetPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v2GetPolicyOptionsModel).ToNot(BeNil())
				Expect(v2GetPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(v2GetPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2ListPoliciesOptions successfully`, func() {
				// Construct an instance of the V2ListPoliciesOptions model
				accountID := "testString"
				v2ListPoliciesOptionsModel := iamPolicyManagementService.NewV2ListPoliciesOptions(accountID)
				v2ListPoliciesOptionsModel.SetAccountID("testString")
				v2ListPoliciesOptionsModel.SetAcceptLanguage("default")
				v2ListPoliciesOptionsModel.SetIamID("testString")
				v2ListPoliciesOptionsModel.SetAccessGroupID("testString")
				v2ListPoliciesOptionsModel.SetType("access")
				v2ListPoliciesOptionsModel.SetServiceType("service")
				v2ListPoliciesOptionsModel.SetServiceName("testString")
				v2ListPoliciesOptionsModel.SetServiceGroupID("testString")
				v2ListPoliciesOptionsModel.SetFormat("include_last_permit")
				v2ListPoliciesOptionsModel.SetState("active")
				v2ListPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v2ListPoliciesOptionsModel).ToNot(BeNil())
				Expect(v2ListPoliciesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(v2ListPoliciesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("default")))
				Expect(v2ListPoliciesOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(v2ListPoliciesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(v2ListPoliciesOptionsModel.Type).To(Equal(core.StringPtr("access")))
				Expect(v2ListPoliciesOptionsModel.ServiceType).To(Equal(core.StringPtr("service")))
				Expect(v2ListPoliciesOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(v2ListPoliciesOptionsModel.ServiceGroupID).To(Equal(core.StringPtr("testString")))
				Expect(v2ListPoliciesOptionsModel.Format).To(Equal(core.StringPtr("include_last_permit")))
				Expect(v2ListPoliciesOptionsModel.State).To(Equal(core.StringPtr("active")))
				Expect(v2ListPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewV2PolicyBaseControl successfully`, func() {
				var grant *iampolicymanagementv1.V2PolicyBaseControlGrant = nil
				_, err := iamPolicyManagementService.NewV2PolicyBaseControl(grant)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewV2PolicyBaseControlGrant successfully`, func() {
				roles := []iampolicymanagementv1.PolicyRole{}
				_model, err := iamPolicyManagementService.NewV2PolicyBaseControlGrant(roles)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2UpdatePolicyOptions successfully`, func() {
				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyBaseControlGrant model
				v2PolicyBaseControlGrantModel := new(iampolicymanagementv1.V2PolicyBaseControlGrant)
				Expect(v2PolicyBaseControlGrantModel).ToNot(BeNil())
				v2PolicyBaseControlGrantModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				Expect(v2PolicyBaseControlGrantModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))

				// Construct an instance of the V2PolicyBaseControl model
				v2PolicyBaseControlModel := new(iampolicymanagementv1.V2PolicyBaseControl)
				Expect(v2PolicyBaseControlModel).ToNot(BeNil())
				v2PolicyBaseControlModel.Grant = v2PolicyBaseControlGrantModel
				Expect(v2PolicyBaseControlModel.Grant).To(Equal(v2PolicyBaseControlGrantModel))

				// Construct an instance of the V2PolicyAttribute model
				v2PolicyAttributeModel := new(iampolicymanagementv1.V2PolicyAttribute)
				Expect(v2PolicyAttributeModel).ToNot(BeNil())
				v2PolicyAttributeModel.Key = core.StringPtr("testString")
				v2PolicyAttributeModel.Operator = core.StringPtr("testString")
				v2PolicyAttributeModel.Value = core.StringPtr("testString")
				Expect(v2PolicyAttributeModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyAttributeModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2PolicyBaseSubject model
				v2PolicyBaseSubjectModel := new(iampolicymanagementv1.V2PolicyBaseSubject)
				Expect(v2PolicyBaseSubjectModel).ToNot(BeNil())
				v2PolicyBaseSubjectModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}
				Expect(v2PolicyBaseSubjectModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}))

				// Construct an instance of the V2PolicyBaseResource model
				v2PolicyBaseResourceModel := new(iampolicymanagementv1.V2PolicyBaseResource)
				Expect(v2PolicyBaseResourceModel).ToNot(BeNil())
				v2PolicyBaseResourceModel.Attributes = []iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}
				Expect(v2PolicyBaseResourceModel.Attributes).To(Equal([]iampolicymanagementv1.V2PolicyAttribute{*v2PolicyAttributeModel}))

				// Construct an instance of the V2PolicyBaseRuleV2PolicyAttribute model
				v2PolicyBaseRuleModel := new(iampolicymanagementv1.V2PolicyBaseRuleV2PolicyAttribute)
				Expect(v2PolicyBaseRuleModel).ToNot(BeNil())
				v2PolicyBaseRuleModel.Key = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Operator = core.StringPtr("testString")
				v2PolicyBaseRuleModel.Value = core.StringPtr("testString")
				Expect(v2PolicyBaseRuleModel.Key).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyBaseRuleModel.Operator).To(Equal(core.StringPtr("testString")))
				Expect(v2PolicyBaseRuleModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the V2UpdatePolicyOptions model
				policyID := "testString"
				ifMatch := "testString"
				v2UpdatePolicyOptionsType := "testString"
				var v2UpdatePolicyOptionsControl *iampolicymanagementv1.V2PolicyBaseControl = nil
				v2UpdatePolicyOptionsModel := iamPolicyManagementService.NewV2UpdatePolicyOptions(policyID, ifMatch, v2UpdatePolicyOptionsType, v2UpdatePolicyOptionsControl)
				v2UpdatePolicyOptionsModel.SetPolicyID("testString")
				v2UpdatePolicyOptionsModel.SetIfMatch("testString")
				v2UpdatePolicyOptionsModel.SetType("testString")
				v2UpdatePolicyOptionsModel.SetControl(v2PolicyBaseControlModel)
				v2UpdatePolicyOptionsModel.SetDescription("testString")
				v2UpdatePolicyOptionsModel.SetSubject(v2PolicyBaseSubjectModel)
				v2UpdatePolicyOptionsModel.SetResource(v2PolicyBaseResourceModel)
				v2UpdatePolicyOptionsModel.SetPattern("testString")
				v2UpdatePolicyOptionsModel.SetRule(v2PolicyBaseRuleModel)
				v2UpdatePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(v2UpdatePolicyOptionsModel).ToNot(BeNil())
				Expect(v2UpdatePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(v2UpdatePolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(v2UpdatePolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(v2UpdatePolicyOptionsModel.Control).To(Equal(v2PolicyBaseControlModel))
				Expect(v2UpdatePolicyOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(v2UpdatePolicyOptionsModel.Subject).To(Equal(v2PolicyBaseSubjectModel))
				Expect(v2UpdatePolicyOptionsModel.Resource).To(Equal(v2PolicyBaseResourceModel))
				Expect(v2UpdatePolicyOptionsModel.Pattern).To(Equal(core.StringPtr("testString")))
				Expect(v2UpdatePolicyOptionsModel.Rule).To(Equal(v2PolicyBaseRuleModel))
				Expect(v2UpdatePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPolicyRole successfully`, func() {
				roleID := "testString"
				_model, err := iamPolicyManagementService.NewPolicyRole(roleID)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
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
			It(`Invoke NewSubjectAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				_model, err := iamPolicyManagementService.NewSubjectAttribute(name, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyAttribute successfully`, func() {
				key := "testString"
				operator := "testString"
				value := core.StringPtr("testString")
				_model, err := iamPolicyManagementService.NewV2PolicyAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyBaseRuleV2PolicyAttribute successfully`, func() {
				key := "testString"
				operator := "testString"
				value := core.StringPtr("testString")
				_model, err := iamPolicyManagementService.NewV2PolicyBaseRuleV2PolicyAttribute(key, operator, value)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewV2PolicyBaseRuleV2RuleWithConditions successfully`, func() {
				operator := "and"
				conditions := []iampolicymanagementv1.V2PolicyAttribute{}
				_model, err := iamPolicyManagementService.NewV2PolicyBaseRuleV2RuleWithConditions(operator, conditions)
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
