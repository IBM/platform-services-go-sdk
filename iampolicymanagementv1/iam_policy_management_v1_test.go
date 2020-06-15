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
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
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

var _ = Describe(`IamPolicyManagementV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "https://iampolicymanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL":       "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL":       "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListPolicies(listPoliciesOptions *ListPoliciesOptions) - Operation response error`, func() {
		listPoliciesPath := "/v1/policies"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPolicies with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("testString")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("testString")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListPolicies(listPoliciesOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
			})
			It(`Invoke ListPolicies successfully`, func() {
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
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListPolicies with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListPoliciesOptions model
				listPoliciesOptionsModel := new(iampolicymanagementv1.ListPoliciesOptions)
				listPoliciesOptionsModel.AccountID = core.StringPtr("testString")
				listPoliciesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPoliciesOptionsModel.IamID = core.StringPtr("testString")
				listPoliciesOptionsModel.AccessGroupID = core.StringPtr("testString")
				listPoliciesOptionsModel.Type = core.StringPtr("testString")
				listPoliciesOptionsModel.ServiceType = core.StringPtr("testString")
				listPoliciesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListPolicies(listPoliciesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListPoliciesOptions model with no property values
				listPoliciesOptionsModelNew := new(iampolicymanagementv1.ListPoliciesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.ListPolicies(listPoliciesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreatePolicy(createPolicyOptions *CreatePolicyOptions) - Operation response error`, func() {
		createPolicyPath := "/v1/policies"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createPolicyPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreatePolicy with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreatePolicy(createPolicyOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
			})
			It(`Invoke CreatePolicy successfully`, func() {
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

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreatePolicy with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsModel := new(iampolicymanagementv1.CreatePolicyOptions)
				createPolicyOptionsModel.Type = core.StringPtr("testString")
				createPolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				createPolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				createPolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				createPolicyOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreatePolicy(createPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreatePolicyOptions model with no property values
				createPolicyOptionsModelNew := new(iampolicymanagementv1.CreatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreatePolicy(createPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdatePolicy(updatePolicyOptions *UpdatePolicyOptions) - Operation response error`, func() {
		updatePolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updatePolicyPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdatePolicy with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdatePolicy(updatePolicyOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
			})
			It(`Invoke UpdatePolicy successfully`, func() {
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

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdatePolicy with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				policyRoleModel.RoleID = core.StringPtr("testString")

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}

				// Construct an instance of the UpdatePolicyOptions model
				updatePolicyOptionsModel := new(iampolicymanagementv1.UpdatePolicyOptions)
				updatePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				updatePolicyOptionsModel.IfMatch = core.StringPtr("testString")
				updatePolicyOptionsModel.Type = core.StringPtr("testString")
				updatePolicyOptionsModel.Subjects = []iampolicymanagementv1.PolicySubject{*policySubjectModel}
				updatePolicyOptionsModel.Roles = []iampolicymanagementv1.PolicyRole{*policyRoleModel}
				updatePolicyOptionsModel.Resources = []iampolicymanagementv1.PolicyResource{*policyResourceModel}
				updatePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdatePolicy(updatePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdatePolicyOptions model with no property values
				updatePolicyOptionsModelNew := new(iampolicymanagementv1.UpdatePolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdatePolicy(updatePolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPolicy(getPolicyOptions *GetPolicyOptions) - Operation response error`, func() {
		getPolicyPath := "/v1/policies/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPolicy with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetPolicy(getPolicyOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getPolicyPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "type": "Type", "subjects": [{"attributes": [{"name": "Name", "value": "Value"}]}], "roles": [{"role_id": "RoleID", "display_name": "DisplayName", "description": "Description"}], "resources": [{"attributes": [{"name": "Name", "value": "Value", "operator": "Operator"}]}], "href": "Href", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID"}`)
				}))
			})
			It(`Invoke GetPolicy successfully`, func() {
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
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetPolicy with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetPolicyOptions model
				getPolicyOptionsModel := new(iampolicymanagementv1.GetPolicyOptions)
				getPolicyOptionsModel.PolicyID = core.StringPtr("testString")
				getPolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetPolicy(getPolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPolicyOptions model with no property values
				getPolicyOptionsModelNew := new(iampolicymanagementv1.GetPolicyOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetPolicy(getPolicyOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
					Expect(req.URL.Path).To(Equal(deletePolicyPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeletePolicy successfully`, func() {
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
				deletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeletePolicy(deletePolicyOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeletePolicy with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeletePolicyOptions model
				deletePolicyOptionsModel := new(iampolicymanagementv1.DeletePolicyOptions)
				deletePolicyOptionsModel.PolicyID = core.StringPtr("testString")
				deletePolicyOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeletePolicy(deletePolicyOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeletePolicyOptions model with no property values
				deletePolicyOptionsModelNew := new(iampolicymanagementv1.DeletePolicyOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeletePolicy(deletePolicyOptionsModelNew)
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
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(testService).ToNot(BeNil())
			Expect(testServiceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "https://iampolicymanagementv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(testService).To(BeNil())
			Expect(testServiceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL":       "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL: "https://testService/api",
				})
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})
				err := testService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(testService).ToNot(BeNil())
				Expect(testServiceErr).To(BeNil())
				Expect(testService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_URL":       "https://iampolicymanagementv1/api",
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"IAM_POLICY_MANAGEMENT_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1UsingExternalConfig(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(testService).To(BeNil())
				Expect(testServiceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`ListRoles(listRolesOptions *ListRolesOptions) - Operation response error`, func() {
		listRolesPath := "/v2/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListRoles with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("testString")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.ListRoles(listRolesOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
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
					fmt.Fprintf(res, `{"custom_roles": [{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}], "service_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn"}], "system_roles": [{"display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn"}]}`)
				}))
			})
			It(`Invoke ListRoles successfully`, func() {
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
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke ListRoles with error: Operation request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := new(iampolicymanagementv1.ListRolesOptions)
				listRolesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listRolesOptionsModel.AccountID = core.StringPtr("testString")
				listRolesOptionsModel.ServiceName = core.StringPtr("testString")
				listRolesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.ListRoles(listRolesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateRole(createRoleOptions *CreateRoleOptions) - Operation response error`, func() {
		createRolePath := "/v2/roles"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateRole with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("testString")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("testString")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.CreateRole(createRoleOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(createRolePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke CreateRole successfully`, func() {
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
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("testString")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("testString")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke CreateRole with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsModel := new(iampolicymanagementv1.CreateRoleOptions)
				createRoleOptionsModel.DisplayName = core.StringPtr("testString")
				createRoleOptionsModel.Actions = []string{"testString"}
				createRoleOptionsModel.Name = core.StringPtr("testString")
				createRoleOptionsModel.AccountID = core.StringPtr("testString")
				createRoleOptionsModel.ServiceName = core.StringPtr("testString")
				createRoleOptionsModel.Description = core.StringPtr("testString")
				createRoleOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.CreateRole(createRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateRoleOptions model with no property values
				createRoleOptionsModelNew := new(iampolicymanagementv1.CreateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.CreateRole(createRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateRole(updateRoleOptions *UpdateRoleOptions) - Operation response error`, func() {
		updateRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateRole with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.UpdateRole(updateRoleOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(updateRolePath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.Header["If-Match"]).ToNot(BeNil())
					Expect(req.Header["If-Match"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke UpdateRole successfully`, func() {
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
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke UpdateRole with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the UpdateRoleOptions model
				updateRoleOptionsModel := new(iampolicymanagementv1.UpdateRoleOptions)
				updateRoleOptionsModel.RoleID = core.StringPtr("testString")
				updateRoleOptionsModel.IfMatch = core.StringPtr("testString")
				updateRoleOptionsModel.DisplayName = core.StringPtr("testString")
				updateRoleOptionsModel.Description = core.StringPtr("testString")
				updateRoleOptionsModel.Actions = []string{"testString"}
				updateRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.UpdateRole(updateRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateRoleOptions model with no property values
				updateRoleOptionsModelNew := new(iampolicymanagementv1.UpdateRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.UpdateRole(updateRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetRole(getRoleOptions *GetRoleOptions) - Operation response error`, func() {
		getRolePath := "/v2/roles/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetRole with error: Operation response processing error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := testService.GetRole(getRoleOptionsModel)
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
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.Path).To(Equal(getRolePath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `{"id": "ID", "display_name": "DisplayName", "description": "Description", "actions": ["Actions"], "crn": "Crn", "name": "Name", "account_id": "AccountID", "service_name": "ServiceName", "created_at": "2019-01-01T12:00:00", "created_by_id": "CreatedByID", "last_modified_at": "2019-01-01T12:00:00", "last_modified_by_id": "LastModifiedByID", "href": "Href"}`)
				}))
			})
			It(`Invoke GetRole successfully`, func() {
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
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = testService.GetRole(getRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
			It(`Invoke GetRole with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the GetRoleOptions model
				getRoleOptionsModel := new(iampolicymanagementv1.GetRoleOptions)
				getRoleOptionsModel.RoleID = core.StringPtr("testString")
				getRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := testService.GetRole(getRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetRoleOptions model with no property values
				getRoleOptionsModelNew := new(iampolicymanagementv1.GetRoleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = testService.GetRole(getRoleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
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
					Expect(req.URL.Path).To(Equal(deleteRolePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteRole successfully`, func() {
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
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = testService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteRole with error: Operation validation and request error`, func() {
				testService, testServiceErr := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(testServiceErr).To(BeNil())
				Expect(testService).ToNot(BeNil())

				// Construct an instance of the DeleteRoleOptions model
				deleteRoleOptionsModel := new(iampolicymanagementv1.DeleteRoleOptions)
				deleteRoleOptionsModel.RoleID = core.StringPtr("testString")
				deleteRoleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := testService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := testService.DeleteRole(deleteRoleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteRoleOptions model with no property values
				deleteRoleOptionsModelNew := new(iampolicymanagementv1.DeleteRoleOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = testService.DeleteRole(deleteRoleOptionsModelNew)
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
			testService, _ := iampolicymanagementv1.NewIamPolicyManagementV1(&iampolicymanagementv1.IamPolicyManagementV1Options{
				URL:           "http://iampolicymanagementv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewCreatePolicyOptions successfully`, func() {
				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				Expect(subjectAttributeModel).ToNot(BeNil())
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")
				Expect(subjectAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(subjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				Expect(policyResourceModel).ToNot(BeNil())
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				Expect(policyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}))

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				Expect(policySubjectModel).ToNot(BeNil())
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}
				Expect(policySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}))

				// Construct an instance of the CreatePolicyOptions model
				createPolicyOptionsType := "testString"
				createPolicyOptionsSubjects := []iampolicymanagementv1.PolicySubject{}
				createPolicyOptionsRoles := []iampolicymanagementv1.PolicyRole{}
				createPolicyOptionsResources := []iampolicymanagementv1.PolicyResource{}
				createPolicyOptionsModel := testService.NewCreatePolicyOptions(createPolicyOptionsType, createPolicyOptionsSubjects, createPolicyOptionsRoles, createPolicyOptionsResources)
				createPolicyOptionsModel.SetType("testString")
				createPolicyOptionsModel.SetSubjects([]iampolicymanagementv1.PolicySubject{*policySubjectModel})
				createPolicyOptionsModel.SetRoles([]iampolicymanagementv1.PolicyRole{*policyRoleModel})
				createPolicyOptionsModel.SetResources([]iampolicymanagementv1.PolicyResource{*policyResourceModel})
				createPolicyOptionsModel.SetAcceptLanguage("testString")
				createPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createPolicyOptionsModel).ToNot(BeNil())
				Expect(createPolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyOptionsModel.Subjects).To(Equal([]iampolicymanagementv1.PolicySubject{*policySubjectModel}))
				Expect(createPolicyOptionsModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))
				Expect(createPolicyOptionsModel.Resources).To(Equal([]iampolicymanagementv1.PolicyResource{*policyResourceModel}))
				Expect(createPolicyOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateRoleOptions successfully`, func() {
				// Construct an instance of the CreateRoleOptions model
				createRoleOptionsDisplayName := "testString"
				createRoleOptionsActions := []string{"testString"}
				createRoleOptionsName := "testString"
				createRoleOptionsAccountID := "testString"
				createRoleOptionsServiceName := "testString"
				createRoleOptionsModel := testService.NewCreateRoleOptions(createRoleOptionsDisplayName, createRoleOptionsActions, createRoleOptionsName, createRoleOptionsAccountID, createRoleOptionsServiceName)
				createRoleOptionsModel.SetDisplayName("testString")
				createRoleOptionsModel.SetActions([]string{"testString"})
				createRoleOptionsModel.SetName("testString")
				createRoleOptionsModel.SetAccountID("testString")
				createRoleOptionsModel.SetServiceName("testString")
				createRoleOptionsModel.SetDescription("testString")
				createRoleOptionsModel.SetAcceptLanguage("testString")
				createRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createRoleOptionsModel).ToNot(BeNil())
				Expect(createRoleOptionsModel.DisplayName).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.Actions).To(Equal([]string{"testString"}))
				Expect(createRoleOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeletePolicyOptions successfully`, func() {
				// Construct an instance of the DeletePolicyOptions model
				policyID := "testString"
				deletePolicyOptionsModel := testService.NewDeletePolicyOptions(policyID)
				deletePolicyOptionsModel.SetPolicyID("testString")
				deletePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deletePolicyOptionsModel).ToNot(BeNil())
				Expect(deletePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(deletePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteRoleOptions successfully`, func() {
				// Construct an instance of the DeleteRoleOptions model
				roleID := "testString"
				deleteRoleOptionsModel := testService.NewDeleteRoleOptions(roleID)
				deleteRoleOptionsModel.SetRoleID("testString")
				deleteRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteRoleOptionsModel).ToNot(BeNil())
				Expect(deleteRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(deleteRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPolicyOptions successfully`, func() {
				// Construct an instance of the GetPolicyOptions model
				policyID := "testString"
				getPolicyOptionsModel := testService.NewGetPolicyOptions(policyID)
				getPolicyOptionsModel.SetPolicyID("testString")
				getPolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPolicyOptionsModel).ToNot(BeNil())
				Expect(getPolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(getPolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetRoleOptions successfully`, func() {
				// Construct an instance of the GetRoleOptions model
				roleID := "testString"
				getRoleOptionsModel := testService.NewGetRoleOptions(roleID)
				getRoleOptionsModel.SetRoleID("testString")
				getRoleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getRoleOptionsModel).ToNot(BeNil())
				Expect(getRoleOptionsModel.RoleID).To(Equal(core.StringPtr("testString")))
				Expect(getRoleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPoliciesOptions successfully`, func() {
				// Construct an instance of the ListPoliciesOptions model
				accountID := "testString"
				listPoliciesOptionsModel := testService.NewListPoliciesOptions(accountID)
				listPoliciesOptionsModel.SetAccountID("testString")
				listPoliciesOptionsModel.SetAcceptLanguage("testString")
				listPoliciesOptionsModel.SetIamID("testString")
				listPoliciesOptionsModel.SetAccessGroupID("testString")
				listPoliciesOptionsModel.SetType("testString")
				listPoliciesOptionsModel.SetServiceType("testString")
				listPoliciesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPoliciesOptionsModel).ToNot(BeNil())
				Expect(listPoliciesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.IamID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.AccessGroupID).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.ServiceType).To(Equal(core.StringPtr("testString")))
				Expect(listPoliciesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListRolesOptions successfully`, func() {
				// Construct an instance of the ListRolesOptions model
				listRolesOptionsModel := testService.NewListRolesOptions()
				listRolesOptionsModel.SetAcceptLanguage("testString")
				listRolesOptionsModel.SetAccountID("testString")
				listRolesOptionsModel.SetServiceName("testString")
				listRolesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listRolesOptionsModel).ToNot(BeNil())
				Expect(listRolesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listRolesOptionsModel.AccountID).To(Equal(core.StringPtr("testString")))
				Expect(listRolesOptionsModel.ServiceName).To(Equal(core.StringPtr("testString")))
				Expect(listRolesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdatePolicyOptions successfully`, func() {
				// Construct an instance of the ResourceAttribute model
				resourceAttributeModel := new(iampolicymanagementv1.ResourceAttribute)
				Expect(resourceAttributeModel).ToNot(BeNil())
				resourceAttributeModel.Name = core.StringPtr("testString")
				resourceAttributeModel.Value = core.StringPtr("testString")
				resourceAttributeModel.Operator = core.StringPtr("testString")
				Expect(resourceAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Value).To(Equal(core.StringPtr("testString")))
				Expect(resourceAttributeModel.Operator).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the SubjectAttribute model
				subjectAttributeModel := new(iampolicymanagementv1.SubjectAttribute)
				Expect(subjectAttributeModel).ToNot(BeNil())
				subjectAttributeModel.Name = core.StringPtr("testString")
				subjectAttributeModel.Value = core.StringPtr("testString")
				Expect(subjectAttributeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(subjectAttributeModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicyResource model
				policyResourceModel := new(iampolicymanagementv1.PolicyResource)
				Expect(policyResourceModel).ToNot(BeNil())
				policyResourceModel.Attributes = []iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}
				Expect(policyResourceModel.Attributes).To(Equal([]iampolicymanagementv1.ResourceAttribute{*resourceAttributeModel}))

				// Construct an instance of the PolicyRole model
				policyRoleModel := new(iampolicymanagementv1.PolicyRole)
				Expect(policyRoleModel).ToNot(BeNil())
				policyRoleModel.RoleID = core.StringPtr("testString")
				Expect(policyRoleModel.RoleID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PolicySubject model
				policySubjectModel := new(iampolicymanagementv1.PolicySubject)
				Expect(policySubjectModel).ToNot(BeNil())
				policySubjectModel.Attributes = []iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}
				Expect(policySubjectModel.Attributes).To(Equal([]iampolicymanagementv1.SubjectAttribute{*subjectAttributeModel}))

				// Construct an instance of the UpdatePolicyOptions model
				policyID := "testString"
				ifMatch := "testString"
				updatePolicyOptionsType := "testString"
				updatePolicyOptionsSubjects := []iampolicymanagementv1.PolicySubject{}
				updatePolicyOptionsRoles := []iampolicymanagementv1.PolicyRole{}
				updatePolicyOptionsResources := []iampolicymanagementv1.PolicyResource{}
				updatePolicyOptionsModel := testService.NewUpdatePolicyOptions(policyID, ifMatch, updatePolicyOptionsType, updatePolicyOptionsSubjects, updatePolicyOptionsRoles, updatePolicyOptionsResources)
				updatePolicyOptionsModel.SetPolicyID("testString")
				updatePolicyOptionsModel.SetIfMatch("testString")
				updatePolicyOptionsModel.SetType("testString")
				updatePolicyOptionsModel.SetSubjects([]iampolicymanagementv1.PolicySubject{*policySubjectModel})
				updatePolicyOptionsModel.SetRoles([]iampolicymanagementv1.PolicyRole{*policyRoleModel})
				updatePolicyOptionsModel.SetResources([]iampolicymanagementv1.PolicyResource{*policyResourceModel})
				updatePolicyOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updatePolicyOptionsModel).ToNot(BeNil())
				Expect(updatePolicyOptionsModel.PolicyID).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.IfMatch).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.Type).To(Equal(core.StringPtr("testString")))
				Expect(updatePolicyOptionsModel.Subjects).To(Equal([]iampolicymanagementv1.PolicySubject{*policySubjectModel}))
				Expect(updatePolicyOptionsModel.Roles).To(Equal([]iampolicymanagementv1.PolicyRole{*policyRoleModel}))
				Expect(updatePolicyOptionsModel.Resources).To(Equal([]iampolicymanagementv1.PolicyResource{*policyResourceModel}))
				Expect(updatePolicyOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateRoleOptions successfully`, func() {
				// Construct an instance of the UpdateRoleOptions model
				roleID := "testString"
				ifMatch := "testString"
				updateRoleOptionsModel := testService.NewUpdateRoleOptions(roleID, ifMatch)
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
			It(`Invoke NewPolicyRole successfully`, func() {
				roleID := "testString"
				model, err := testService.NewPolicyRole(roleID)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				model, err := testService.NewResourceAttribute(name, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSubjectAttribute successfully`, func() {
				name := "testString"
				value := "testString"
				model, err := testService.NewSubjectAttribute(name, value)
				Expect(model).ToNot(BeNil())
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
